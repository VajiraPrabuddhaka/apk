// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package gen

import (
	"encoding/json"
	"fmt"
)

// Defines values for APIType.
const (
	ASYNC   APIType = "ASYNC"
	GRAPHQL APIType = "GRAPHQL"
	HTTP    APIType = "HTTP"
	SSE     APIType = "SSE"
	WEBHOOK APIType = "WEBHOOK"
	WEBSUB  APIType = "WEBSUB"
	WS      APIType = "WS"
)

// API defines model for API.
type API struct {
	Context     string  `json:"context"`
	CreatedTime *string `json:"createdTime,omitempty"`

	// Endpoint configuration of the API. This can be used to provide different types of endpoints including Simple REST Endpoints, Loadbalanced and Failover.
	//
	// `Simple REST Endpoint`
	//   {
	//     "endpoint_type": "http",
	//     "sandbox_endpoints":       {
	//        "url": "https://pizzashack-service:8080/am/sample/pizzashack/v3/api/"
	//     },
	//     "production_endpoints":       {
	//        "url": "https://pizzashack-service:8080/am/sample/pizzashack/v3/api/"
	//     }
	//   }
	EndpointConfig  *map[string]interface{} `json:"endpointConfig,omitempty"`
	LastUpdatedTime *string                 `json:"lastUpdatedTime,omitempty"`
	Name            string                  `json:"name"`
	Operations      *[]APIOperations        `json:"operations,omitempty"`
	ServiceInfo     *APIServiceInfo         `json:"serviceInfo,omitempty"`

	// The api creation type to be used. Accepted values are HTTP, WS, GRAPHQL, WEBSUB, SSE, WEBHOOK, ASYNC
	Type    *APIType `json:"type,omitempty"`
	Version string   `json:"version"`
}

// The api creation type to be used. Accepted values are HTTP, WS, GRAPHQL, WEBSUB, SSE, WEBHOOK, ASYNC
type APIType string

// APIDefinitionValidationResponse defines model for APIDefinitionValidationResponse.
type APIDefinitionValidationResponse struct {
	// OpenAPI definition content.
	Content *string `json:"content,omitempty"`

	// If there are more than one error list them out.
	// For example, list out validation errors by each field.
	Errors *[]ErrorListItem `json:"errors,omitempty"`

	// API definition information
	Info *APIDefinitionValidationResponseInfo `json:"info,omitempty"`

	// This attribute declares whether this definition is valid or not.
	IsValid bool `json:"isValid"`
}

// API definition information
type APIDefinitionValidationResponseInfo struct {
	Context     *string `json:"context,omitempty"`
	Description *string `json:"description,omitempty"`

	// contains host/servers specified in the API definition file/URL
	Endpoints      *[]string `json:"endpoints,omitempty"`
	Name           *string   `json:"name,omitempty"`
	OpenAPIVersion *string   `json:"openAPIVersion,omitempty"`
	Version        *string   `json:"version,omitempty"`
}

// APIInfo defines model for APIInfo.
type APIInfo struct {
	Context     *string `json:"context,omitempty"`
	CreatedTime *string `json:"createdTime,omitempty"`
	Name        *string `json:"name,omitempty"`
	Type        *string `json:"type,omitempty"`
	UpdatedTime *string `json:"updatedTime,omitempty"`
	Version     *string `json:"version,omitempty"`
}

// APIList defines model for APIList.
type APIList struct {
	// Number of APIs returned.
	Count      *int        `json:"count,omitempty"`
	List       *[]APIInfo  `json:"list,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// APIOperationPolicies defines model for APIOperationPolicies.
type APIOperationPolicies struct {
	Fault    *[]OperationPolicy `json:"fault,omitempty"`
	Request  *[]OperationPolicy `json:"request,omitempty"`
	Response *[]OperationPolicy `json:"response,omitempty"`
}

// APIOperations defines model for APIOperations.
type APIOperations struct {
	Id                *string               `json:"id,omitempty"`
	OperationPolicies *APIOperationPolicies `json:"operationPolicies,omitempty"`
	PayloadSchema     *string               `json:"payloadSchema,omitempty"`
	Scopes            *[]string             `json:"scopes,omitempty"`
	Target            *string               `json:"target,omitempty"`
	UriMapping        *string               `json:"uriMapping,omitempty"`
	Verb              *string               `json:"verb,omitempty"`
}

// APIServiceInfo defines model for API_serviceInfo.
type APIServiceInfo struct {
	Name *string `json:"name,omitempty"`
}

// Error defines model for Error.
type Error struct {
	Code int64 `json:"code"`

	// A detail description about the error message.
	Description *string `json:"description,omitempty"`

	// If there are more than one error list them out.
	// For example, list out validation errors by each field.
	Error *[]ErrorListItem `json:"error,omitempty"`

	// Error message.
	Message string `json:"message"`

	// Preferably an url with more details about the error.
	MoreInfo *string `json:"moreInfo,omitempty"`
}

// ErrorListItem defines model for ErrorListItem.
type ErrorListItem struct {
	Code string `json:"code"`

	// A detail description about the error message.
	Description *string `json:"description,omitempty"`

	// Description about individual errors occurred
	Message string `json:"message"`
}

// MediationPolicy defines model for MediationPolicy.
type MediationPolicy struct {
	Name string  `json:"name"`
	Type *string `json:"type,omitempty"`
}

// MediationPolicyList defines model for MediationPolicyList.
type MediationPolicyList struct {
	List       *[]MediationPolicy `json:"list,omitempty"`
	Pagination *Pagination        `json:"pagination,omitempty"`
}

// OperationPolicy defines model for OperationPolicy.
type OperationPolicy struct {
	Parameters    *OperationPolicy_Parameters `json:"parameters,omitempty"`
	PolicyId      *string                     `json:"policyId,omitempty"`
	PolicyName    string                      `json:"policyName"`
	PolicyVersion *string                     `json:"policyVersion,omitempty"`
}

// OperationPolicy_Parameters defines model for OperationPolicy.Parameters.
type OperationPolicy_Parameters struct {
	AdditionalProperties map[string]map[string]interface{} `json:"-"`
}

// Pagination defines model for Pagination.
type Pagination struct {
	Limit *int `json:"limit,omitempty"`

	// Link to the next subset of resources qualified.
	// Empty if no more resources are to be returned.
	Next   *string `json:"next,omitempty"`
	Offset *int    `json:"offset,omitempty"`

	// Link to the previous subset of resources qualified.
	// Empty if current subset is the first subset returned.
	Previous *string `json:"previous,omitempty"`
	Total    *int    `json:"total,omitempty"`
}

// PortMapping defines model for PortMapping.
type PortMapping struct {
	Name       string  `json:"name"`
	Port       int32   `json:"port"`
	Protocol   *string `json:"protocol,omitempty"`
	Targetport int32   `json:"targetport"`
}

// Service defines model for Service.
type Service struct {
	Name        string         `json:"name"`
	Namespace   string         `json:"namespace"`
	Portmapping *[]PortMapping `json:"portmapping,omitempty"`
	Type        string         `json:"type"`
}

// ServiceList defines model for ServiceList.
type ServiceList struct {
	List       *[]Service  `json:"list,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// ApiIdDefinitionBody defines model for apiId_definition_body.
type ApiIdDefinitionBody struct {
	// API definition of the API
	ApiDefinition *string `json:"apiDefinition,omitempty"`

	// API definitio as a file
	File *string `json:"file,omitempty"`

	// API definition URL of the API
	Url *string `json:"url,omitempty"`
}

// ApisImportBody defines model for apis_import_body.
type ApisImportBody struct {
	// Zip archive consisting on exported API configuration
	File string `json:"file"`
}

// ApisImportdefinitionBody defines model for apis_importdefinition_body.
type ApisImportdefinitionBody struct {
	// Additional attributes specified as a stringified JSON with API's schema
	AdditionalProperties *string `json:"additionalProperties,omitempty"`

	// Definition to upload as a file
	File *string `json:"file,omitempty"`

	// Inline content of the API definition
	InlineAPIDefinition *string `json:"inlineAPIDefinition,omitempty"`

	// Definition url
	Url *string `json:"url,omitempty"`
}

// ApisValidatedefinitionBody defines model for apis_validatedefinition_body.
type ApisValidatedefinitionBody struct {
	// API definition as a file
	File *string `json:"file,omitempty"`

	// Inline content of the API definition
	InlineAPIDefinition *string `json:"inlineAPIDefinition,omitempty"`

	// API definition type - OpenAPI/AsyncAPI/GraphQL
	Type *string `json:"type,omitempty"`

	// API definition definition url
	Url *string `json:"url,omitempty"`
}

// GetAllAPIsParams defines parameters for GetAllAPIs.
type GetAllAPIsParams struct {
	// Maximum size of resource array to return.
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`

	// Starting point within the complete list of items qualified.
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`

	// Criteria for sorting.
	SortBy *GetAllAPIsParamsSortBy `form:"sortBy,omitempty" json:"sortBy,omitempty"`

	// Order of sorting(ascending/descending).
	SortOrder *GetAllAPIsParamsSortOrder `form:"sortOrder,omitempty" json:"sortOrder,omitempty"`
	Query     *string                    `form:"query,omitempty" json:"query,omitempty"`

	// Media types acceptable for the response. Default is application/json.
	Accept *string `json:"Accept,omitempty"`
}

// GetAllAPIsParamsSortBy defines parameters for GetAllAPIs.
type GetAllAPIsParamsSortBy string

// GetAllAPIsParamsSortOrder defines parameters for GetAllAPIs.
type GetAllAPIsParamsSortOrder string

// CreateAPIJSONBody defines parameters for CreateAPI.
type CreateAPIJSONBody = API

// ExportAPIParams defines parameters for ExportAPI.
type ExportAPIParams struct {
	// Name of the API
	ApiId *string `form:"apiId,omitempty" json:"apiId,omitempty"`

	// API Name
	Name *string `form:"name,omitempty" json:"name,omitempty"`

	// Version of the API
	Version *string `form:"version,omitempty" json:"version,omitempty"`

	// Format of output documents. Can be YAML or JSON.
	Format *ExportAPIParamsFormat `form:"format,omitempty" json:"format,omitempty"`
}

// ExportAPIParamsFormat defines parameters for ExportAPI.
type ExportAPIParamsFormat string

// ImportAPIParams defines parameters for ImportAPI.
type ImportAPIParams struct {
	// Whether to update the API or not. This is used when updating already existing APIs
	Overwrite *bool `form:"overwrite,omitempty" json:"overwrite,omitempty"`
}

// ImportServiceJSONBody defines parameters for ImportService.
type ImportServiceJSONBody = API

// ImportServiceParams defines parameters for ImportService.
type ImportServiceParams struct {
	// ID of service that should be imported from Service Catalog
	ServiceKey string `form:"serviceKey" json:"serviceKey"`
}

// ValidateAPIDefinitionParams defines parameters for ValidateOpenAPIDefinition.
type ValidateAPIDefinitionParams struct {
	// Specify whether to return the full content of the OpenAPI definition in the response. This is only
	// applicable when using url based validation
	ReturnContent *bool `form:"returnContent,omitempty" json:"returnContent,omitempty"`
}

// UpdateAPIJSONBody defines parameters for UpdateAPI.
type UpdateAPIJSONBody = API

// GetAllPoliciesParams defines parameters for GetAllPolicies.
type GetAllPoliciesParams struct {
	// Maximum size of resource array to return.
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`

	// Starting point within the complete list of items qualified.
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`

	// Criteria for sorting.
	SortBy *GetAllPoliciesParamsSortBy `form:"sortBy,omitempty" json:"sortBy,omitempty"`

	// Order of sorting(ascending/descending).
	SortOrder *GetAllPoliciesParamsSortOrder `form:"sortOrder,omitempty" json:"sortOrder,omitempty"`
	Query     *string                        `form:"query,omitempty" json:"query,omitempty"`

	// Media types acceptable for the response. Default is application/json.
	Accept *string `json:"Accept,omitempty"`
}

// GetAllPoliciesParamsSortBy defines parameters for GetAllPolicies.
type GetAllPoliciesParamsSortBy string

// GetAllPoliciesParamsSortOrder defines parameters for GetAllPolicies.
type GetAllPoliciesParamsSortOrder string

// SearchServicesParams defines parameters for SearchServices.
type SearchServicesParams struct {
	// Filter services by the name of the service
	Name *string `form:"name,omitempty" json:"name,omitempty"`

	// Required namespace in the K8s cluster
	Namespace *string `form:"namespace,omitempty" json:"namespace,omitempty"`

	// Criteria for sorting.
	SortBy *SearchServicesParamsSortBy `form:"sortBy,omitempty" json:"sortBy,omitempty"`

	// Order of sorting(ascending/descending).
	SortOrder *SearchServicesParamsSortOrder `form:"sortOrder,omitempty" json:"sortOrder,omitempty"`

	// Maximum size of resource array to return.
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`

	// Starting point within the complete list of items qualified.
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`
}

// SearchServicesParamsSortBy defines parameters for SearchServices.
type SearchServicesParamsSortBy string

// SearchServicesParamsSortOrder defines parameters for SearchServices.
type SearchServicesParamsSortOrder string

// GetServiceByIdParams defines parameters for GetServiceById.
type GetServiceByIdParams struct {
	// Required namespace in the K8s cluster
	Namespace *string `form:"namespace,omitempty" json:"namespace,omitempty"`
}

// GetServiceUsageParams defines parameters for GetServiceUsage.
type GetServiceUsageParams struct {
	// Required namespace in the K8s cluster
	Namespace *string `form:"namespace,omitempty" json:"namespace,omitempty"`
}

// CreateAPIJSONRequestBody defines body for CreateAPI for application/json ContentType.
type CreateAPIJSONRequestBody = CreateAPIJSONBody

// ImportServiceJSONRequestBody defines body for ImportService for application/json ContentType.
type ImportServiceJSONRequestBody = ImportServiceJSONBody

// UpdateAPIJSONRequestBody defines body for UpdateAPI for application/json ContentType.
type UpdateAPIJSONRequestBody = UpdateAPIJSONBody

// Getter for additional properties for OperationPolicy_Parameters. Returns the specified
// element and whether it was found
func (a OperationPolicy_Parameters) Get(fieldName string) (value map[string]interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for OperationPolicy_Parameters
func (a *OperationPolicy_Parameters) Set(fieldName string, value map[string]interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for OperationPolicy_Parameters to handle AdditionalProperties
func (a *OperationPolicy_Parameters) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal map[string]interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for OperationPolicy_Parameters to handle AdditionalProperties
func (a OperationPolicy_Parameters) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}
