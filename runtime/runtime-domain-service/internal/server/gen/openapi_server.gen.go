// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package gen

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi/v5"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Retrieve/Search APIs
	// (GET /apis)
	GetAllAPIs(w http.ResponseWriter, r *http.Request, params GetAllAPIsParams)
	// Create a New API
	// (POST /apis)
	CreateAPI(w http.ResponseWriter, r *http.Request)
	// Export an API
	// (GET /apis/export)
	ExportAPI(w http.ResponseWriter, r *http.Request, params ExportAPIParams)
	// Import an API
	// (POST /apis/import)
	ImportAPI(w http.ResponseWriter, r *http.Request, params ImportAPIParams)
	// Import an API Definition
	// (POST /apis/import-definition)
	ImportAPIDefinition(w http.ResponseWriter, r *http.Request)
	// Create API from a Service
	// (POST /apis/import-service)
	ImportService(w http.ResponseWriter, r *http.Request, params ImportServiceParams)
	// Check Given API Context Name already Exists
	// (POST /apis/validate)
	ValidateAPI(w http.ResponseWriter, r *http.Request)
	// Validate an OpenAPI Definition
	// (POST /apis/validate-definition)
	ValidateAPIDefinition(w http.ResponseWriter, r *http.Request, params ValidateAPIDefinitionParams)
	// Delete an API
	// (DELETE /apis/{apiId})
	DeleteAPI(w http.ResponseWriter, r *http.Request, apiId string)
	// Get Details of an API
	// (GET /apis/{apiId})
	GetAPI(w http.ResponseWriter, r *http.Request, apiId string)
	// Update an API
	// (PUT /apis/{apiId})
	UpdateAPI(w http.ResponseWriter, r *http.Request, apiId string)
	// Get API Definition
	// (GET /apis/{apiId}/definition)
	GetAPIDefinition(w http.ResponseWriter, r *http.Request, apiId string)
	// Update API Definition
	// (PUT /apis/{apiId}/definition)
	UpdateAPIDefinition(w http.ResponseWriter, r *http.Request, apiId string)
	// Retrieve/Search Policies
	// (GET /policies)
	GetAllPolicies(w http.ResponseWriter, r *http.Request, params GetAllPoliciesParams)
	// Get Details of a Policy
	// (GET /policies/{mediationPolicyId})
	GetPolicy(w http.ResponseWriter, r *http.Request, mediationPolicyId string)
	// Retrieve/search services
	// (GET /services)
	SearchServices(w http.ResponseWriter, r *http.Request, params SearchServicesParams)
	// Get details of a service
	// (GET /services/{serviceId})
	GetServiceById(w http.ResponseWriter, r *http.Request, serviceId string, params GetServiceByIdParams)
	// Retrieve the usage of service
	// (GET /services/{serviceId}/usage)
	GetServiceUsage(w http.ResponseWriter, r *http.Request, serviceId string, params GetServiceUsageParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// GetAllAPIs operation middleware
func (siw *ServerInterfaceWrapper) GetAllAPIs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetAllAPIsParams

	// ------------- Optional query parameter "limit" -------------
	if paramValue := r.URL.Query().Get("limit"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	// ------------- Optional query parameter "offset" -------------
	if paramValue := r.URL.Query().Get("offset"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "offset", r.URL.Query(), &params.Offset)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "offset", Err: err})
		return
	}

	// ------------- Optional query parameter "sortBy" -------------
	if paramValue := r.URL.Query().Get("sortBy"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "sortBy", r.URL.Query(), &params.SortBy)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "sortBy", Err: err})
		return
	}

	// ------------- Optional query parameter "sortOrder" -------------
	if paramValue := r.URL.Query().Get("sortOrder"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "sortOrder", r.URL.Query(), &params.SortOrder)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "sortOrder", Err: err})
		return
	}

	// ------------- Optional query parameter "query" -------------
	if paramValue := r.URL.Query().Get("query"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "query", r.URL.Query(), &params.Query)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "query", Err: err})
		return
	}

	headers := r.Header

	// ------------- Optional header parameter "Accept" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("Accept")]; found {
		var Accept string
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandlerFunc(w, r, &TooManyValuesForParamError{ParamName: "Accept", Count: n})
			return
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "Accept", runtime.ParamLocationHeader, valueList[0], &Accept)
		if err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "Accept", Err: err})
			return
		}

		params.Accept = &Accept

	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetAllAPIs(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// CreateAPI operation middleware
func (siw *ServerInterfaceWrapper) CreateAPI(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateAPI(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ExportAPI operation middleware
func (siw *ServerInterfaceWrapper) ExportAPI(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ExportAPIParams

	// ------------- Optional query parameter "apiId" -------------
	if paramValue := r.URL.Query().Get("apiId"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "apiId", r.URL.Query(), &params.ApiId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "apiId", Err: err})
		return
	}

	// ------------- Optional query parameter "name" -------------
	if paramValue := r.URL.Query().Get("name"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "name", r.URL.Query(), &params.Name)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "name", Err: err})
		return
	}

	// ------------- Optional query parameter "version" -------------
	if paramValue := r.URL.Query().Get("version"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "version", r.URL.Query(), &params.Version)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "version", Err: err})
		return
	}

	// ------------- Optional query parameter "format" -------------
	if paramValue := r.URL.Query().Get("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", r.URL.Query(), &params.Format)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "format", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ExportAPI(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ImportAPI operation middleware
func (siw *ServerInterfaceWrapper) ImportAPI(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ImportAPIParams

	// ------------- Optional query parameter "overwrite" -------------
	if paramValue := r.URL.Query().Get("overwrite"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "overwrite", r.URL.Query(), &params.Overwrite)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "overwrite", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ImportAPI(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ImportAPIDefinition operation middleware
func (siw *ServerInterfaceWrapper) ImportAPIDefinition(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ImportAPIDefinition(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ImportService operation middleware
func (siw *ServerInterfaceWrapper) ImportService(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ImportServiceParams

	// ------------- Required query parameter "serviceKey" -------------
	if paramValue := r.URL.Query().Get("serviceKey"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "serviceKey"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "serviceKey", r.URL.Query(), &params.ServiceKey)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "serviceKey", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ImportService(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ValidateAPI operation middleware
func (siw *ServerInterfaceWrapper) ValidateAPI(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ValidateAPI(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ValidateAPIDefinition operation middleware
func (siw *ServerInterfaceWrapper) ValidateAPIDefinition(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ValidateAPIDefinitionParams

	// ------------- Optional query parameter "returnContent" -------------
	if paramValue := r.URL.Query().Get("returnContent"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "returnContent", r.URL.Query(), &params.ReturnContent)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "returnContent", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ValidateAPIDefinition(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// DeleteAPI operation middleware
func (siw *ServerInterfaceWrapper) DeleteAPI(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "apiId" -------------
	var apiId string

	err = runtime.BindStyledParameter("simple", false, "apiId", chi.URLParam(r, "apiId"), &apiId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "apiId", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteAPI(w, r, apiId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetAPI operation middleware
func (siw *ServerInterfaceWrapper) GetAPI(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "apiId" -------------
	var apiId string

	err = runtime.BindStyledParameter("simple", false, "apiId", chi.URLParam(r, "apiId"), &apiId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "apiId", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetAPI(w, r, apiId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// UpdateAPI operation middleware
func (siw *ServerInterfaceWrapper) UpdateAPI(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "apiId" -------------
	var apiId string

	err = runtime.BindStyledParameter("simple", false, "apiId", chi.URLParam(r, "apiId"), &apiId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "apiId", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateAPI(w, r, apiId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetAPIDefinition operation middleware
func (siw *ServerInterfaceWrapper) GetAPIDefinition(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "apiId" -------------
	var apiId string

	err = runtime.BindStyledParameter("simple", false, "apiId", chi.URLParam(r, "apiId"), &apiId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "apiId", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetAPIDefinition(w, r, apiId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// UpdateAPIDefinition operation middleware
func (siw *ServerInterfaceWrapper) UpdateAPIDefinition(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "apiId" -------------
	var apiId string

	err = runtime.BindStyledParameter("simple", false, "apiId", chi.URLParam(r, "apiId"), &apiId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "apiId", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateAPIDefinition(w, r, apiId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetAllPolicies operation middleware
func (siw *ServerInterfaceWrapper) GetAllPolicies(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetAllPoliciesParams

	// ------------- Optional query parameter "limit" -------------
	if paramValue := r.URL.Query().Get("limit"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	// ------------- Optional query parameter "offset" -------------
	if paramValue := r.URL.Query().Get("offset"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "offset", r.URL.Query(), &params.Offset)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "offset", Err: err})
		return
	}

	// ------------- Optional query parameter "sortBy" -------------
	if paramValue := r.URL.Query().Get("sortBy"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "sortBy", r.URL.Query(), &params.SortBy)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "sortBy", Err: err})
		return
	}

	// ------------- Optional query parameter "sortOrder" -------------
	if paramValue := r.URL.Query().Get("sortOrder"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "sortOrder", r.URL.Query(), &params.SortOrder)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "sortOrder", Err: err})
		return
	}

	// ------------- Optional query parameter "query" -------------
	if paramValue := r.URL.Query().Get("query"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "query", r.URL.Query(), &params.Query)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "query", Err: err})
		return
	}

	headers := r.Header

	// ------------- Optional header parameter "Accept" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("Accept")]; found {
		var Accept string
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandlerFunc(w, r, &TooManyValuesForParamError{ParamName: "Accept", Count: n})
			return
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "Accept", runtime.ParamLocationHeader, valueList[0], &Accept)
		if err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "Accept", Err: err})
			return
		}

		params.Accept = &Accept

	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetAllPolicies(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetPolicy operation middleware
func (siw *ServerInterfaceWrapper) GetPolicy(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "mediationPolicyId" -------------
	var mediationPolicyId string

	err = runtime.BindStyledParameter("simple", false, "mediationPolicyId", chi.URLParam(r, "mediationPolicyId"), &mediationPolicyId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "mediationPolicyId", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetPolicy(w, r, mediationPolicyId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// SearchServices operation middleware
func (siw *ServerInterfaceWrapper) SearchServices(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params SearchServicesParams

	// ------------- Optional query parameter "name" -------------
	if paramValue := r.URL.Query().Get("name"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "name", r.URL.Query(), &params.Name)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "name", Err: err})
		return
	}

	// ------------- Optional query parameter "namespace" -------------
	if paramValue := r.URL.Query().Get("namespace"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "namespace", r.URL.Query(), &params.Namespace)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "namespace", Err: err})
		return
	}

	// ------------- Optional query parameter "sortBy" -------------
	if paramValue := r.URL.Query().Get("sortBy"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "sortBy", r.URL.Query(), &params.SortBy)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "sortBy", Err: err})
		return
	}

	// ------------- Optional query parameter "sortOrder" -------------
	if paramValue := r.URL.Query().Get("sortOrder"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "sortOrder", r.URL.Query(), &params.SortOrder)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "sortOrder", Err: err})
		return
	}

	// ------------- Optional query parameter "limit" -------------
	if paramValue := r.URL.Query().Get("limit"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	// ------------- Optional query parameter "offset" -------------
	if paramValue := r.URL.Query().Get("offset"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "offset", r.URL.Query(), &params.Offset)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "offset", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.SearchServices(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetServiceById operation middleware
func (siw *ServerInterfaceWrapper) GetServiceById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "serviceId" -------------
	var serviceId string

	err = runtime.BindStyledParameter("simple", false, "serviceId", chi.URLParam(r, "serviceId"), &serviceId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "serviceId", Err: err})
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetServiceByIdParams

	// ------------- Optional query parameter "namespace" -------------
	if paramValue := r.URL.Query().Get("namespace"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "namespace", r.URL.Query(), &params.Namespace)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "namespace", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetServiceById(w, r, serviceId, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetServiceUsage operation middleware
func (siw *ServerInterfaceWrapper) GetServiceUsage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "serviceId" -------------
	var serviceId string

	err = runtime.BindStyledParameter("simple", false, "serviceId", chi.URLParam(r, "serviceId"), &serviceId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "serviceId", Err: err})
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetServiceUsageParams

	// ------------- Optional query parameter "namespace" -------------
	if paramValue := r.URL.Query().Get("namespace"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "namespace", r.URL.Query(), &params.Namespace)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "namespace", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetServiceUsage(w, r, serviceId, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/apis", wrapper.GetAllAPIs)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/apis", wrapper.CreateAPI)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/apis/export", wrapper.ExportAPI)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/apis/import", wrapper.ImportAPI)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/apis/import-definition", wrapper.ImportAPIDefinition)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/apis/import-service", wrapper.ImportService)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/apis/validate", wrapper.ValidateAPI)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/apis/validate-definition", wrapper.ValidateAPIDefinition)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/apis/{apiId}", wrapper.DeleteAPI)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/apis/{apiId}", wrapper.GetAPI)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/apis/{apiId}", wrapper.UpdateAPI)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/apis/{apiId}/definition", wrapper.GetAPIDefinition)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/apis/{apiId}/definition", wrapper.UpdateAPIDefinition)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/policies", wrapper.GetAllPolicies)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/policies/{mediationPolicyId}", wrapper.GetPolicy)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/services", wrapper.SearchServices)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/services/{serviceId}", wrapper.GetServiceById)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/services/{serviceId}/usage", wrapper.GetServiceUsage)
	})

	return r
}
