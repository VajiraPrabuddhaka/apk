/*
 *  Copyright (c) 2022, WSO2 LLC. (http://www.wso2.org).
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package xds

import (
	"context"
	"fmt"
	"github.com/wso2/apk/APKManagementServer/internal/database"
	"math/rand"
	"net"
	"sync"
	"time"

	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/wso2/apk/APKManagementServer/internal/config"
	"github.com/wso2/apk/APKManagementServer/internal/logger"
	internal_types "github.com/wso2/apk/APKManagementServer/internal/types"
	"github.com/wso2/apk/APKManagementServer/internal/xds/callbacks"
	apkmgt_application "github.com/wso2/apk/adapter/pkg/discovery/api/wso2/discovery/apkmgt"
	apkmgt_service "github.com/wso2/apk/adapter/pkg/discovery/api/wso2/discovery/service/apkmgt"
	wso2_cache "github.com/wso2/apk/adapter/pkg/discovery/protocol/cache/v3"
	wso2_resource "github.com/wso2/apk/adapter/pkg/discovery/protocol/resource/v3"
	wso2_server "github.com/wso2/apk/adapter/pkg/discovery/protocol/server/v3"
	"github.com/wso2/apk/adapter/pkg/logging"
	"google.golang.org/grpc"
)

var (
	apiCache         wso2_cache.SnapshotCache
	apiCacheMutex    sync.Mutex
	introducedLabels map[string]bool
	Sent             bool = true
)

const (
	maxRandomInt             int    = 999999999
	typeURL                  string = "wso2.discovery.apkmgt.Application"
	grpcMaxConcurrentStreams        = 1000000
)

// IDHash uses ID field as the node hash.
type IDHash struct{}

// ID uses the node ID field
func (IDHash) ID(node *corev3.Node) string {
	if node == nil {
		return "unknown"
	}
	return node.Id
}

var _ wso2_cache.NodeHash = IDHash{}

func init() {
	apiCache = wso2_cache.NewSnapshotCache(false, IDHash{}, nil)
	rand.Seed(time.Now().UnixNano())
}

// FeedData mock data
func FeedData() {
	logger.LoggerXdsServer.Debug("adding mock data")
	version := rand.Intn(maxRandomInt)
	applications := apkmgt_application.Application{
		Uuid: "apiUUID1",
		Name: "name1",
	}
	newSnapshot, _ := wso2_cache.NewSnapshot(fmt.Sprint(version), map[wso2_resource.Type][]types.Resource{
		wso2_resource.APKMgtApplicationType: {&applications},
	})
	apiCacheMutex.Lock()
	defer apiCacheMutex.Unlock()
	apiCache.SetSnapshot(context.Background(), "mine", newSnapshot)
}

func AddMultipleApplications(applicationEventArray []*internal_types.ApplicationEvent) {
	snapshotMap := make(map[string]*wso2_cache.Snapshot)
	version := rand.Intn(maxRandomInt)

	for _, event := range applicationEventArray {
		label := event.Label
		appUUID := event.UUID

		snapshotEntry, snapshotFound := snapshotMap[label]
		var newSnapshot wso2_cache.Snapshot
		application, _ := database.GetApplicationByUUID(appUUID)
		if !snapshotFound {
			newSnapshot, _ = wso2_cache.NewSnapshot(fmt.Sprint(version), map[wso2_resource.Type][]types.Resource{
				wso2_resource.APKMgtApplicationType: {application},
			})
			snapshotEntry = &newSnapshot
			snapshotMap[label] = &newSnapshot
		} else {
			// error occurs if no snapshot is under the provided label
			resourceMap := snapshotEntry.GetResourcesAndTTL(typeURL)
			resourceMap[appUUID] = types.ResourceWithTTL{
				Resource: application,
			}
			appResources := convertResourceMapToArray(resourceMap)
			newSnapshot, _ = wso2_cache.NewSnapshot(fmt.Sprint(version), map[wso2_resource.Type][]types.Resource{
				wso2_resource.APKMgtApplicationType: appResources,
			})
			snapshotMap[label] = &newSnapshot
		}
	}
	apiCacheMutex.Lock()
	defer apiCacheMutex.Unlock()
	for label, snapshotEntry := range snapshotMap {
		apiCache.SetSnapshot(context.Background(), label, *snapshotEntry)
		introducedLabels[label] = true
		logger.LoggerXds.Infof("Application Snaphsot is updated for label %s with the version %d.", label, version)
	}
}

func convertResourceMapToArray(resourceMap map[string]types.ResourceWithTTL) []types.Resource {
	var appResources []types.Resource
	for _, res := range resourceMap {
		appResources = append(appResources, res.Resource)
	}
	return appResources
}

func InitAPKMgtServer() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	apkMgtAPIDsSrv := wso2_server.NewServer(ctx, apiCache, &callbacks.Callbacks{})

	var grpcOptions []grpc.ServerOption
	grpcOptions = append(grpcOptions, grpc.MaxConcurrentStreams(grpcMaxConcurrentStreams))
	grpcServer := grpc.NewServer(grpcOptions...)
	apkmgt_service.RegisterAPKMgtDiscoveryServiceServer(grpcServer, apkMgtAPIDsSrv)
	config := config.ReadConfigs()
	port := config.ManagementServer.XDSPort

	//todo (amaliMatharaarachchi) handle error gracefully
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logger.LoggerServer.ErrorC(logging.ErrorDetails{
			Message:   fmt.Sprintf("Error while listening on port: %v. Error: %v", port, err.Error()),
			Severity:  logging.BLOCKER,
			ErrorCode: 1000,
		})
	}

	logger.LoggerServer.Infof("APK Management server XDS is starting on port %v.", port)
	if err = grpcServer.Serve(listener); err != nil {
		logger.LoggerServer.ErrorC(logging.ErrorDetails{
			Message:   fmt.Sprint("Error while starting APK Management server XDS server."),
			Severity:  logging.BLOCKER,
			ErrorCode: 1001,
		})
	}
}
