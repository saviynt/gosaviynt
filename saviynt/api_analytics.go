/*
Saviynt Enterprise Identity Cloud (EIC) API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package saviynt

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// AnalyticsAPIService AnalyticsAPI service
type AnalyticsAPIService service

type ApiFetchControlAttributesRequest struct {
	ctx context.Context
	ApiService *AnalyticsAPIService
	fetchControlAttributesRequest *FetchControlAttributesRequest
}

func (r ApiFetchControlAttributesRequest) FetchControlAttributesRequest(fetchControlAttributesRequest FetchControlAttributesRequest) ApiFetchControlAttributesRequest {
	r.fetchControlAttributesRequest = &fetchControlAttributesRequest
	return r
}

func (r ApiFetchControlAttributesRequest) Execute() (*http.Response, error) {
	return r.ApiService.FetchControlAttributesExecute(r)
}

/*
FetchControlAttributes Fetch Control Attributes

This API is used for fetching the details of dynamic attributes associated with an analytics control. This can be used for both Elasticsearch and Database analytics.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFetchControlAttributesRequest
*/
func (a *AnalyticsAPIService) FetchControlAttributes(ctx context.Context) ApiFetchControlAttributesRequest {
	return ApiFetchControlAttributesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *AnalyticsAPIService) FetchControlAttributesExecute(r ApiFetchControlAttributesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsAPIService.FetchControlAttributes")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fetchControlAttributes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.fetchControlAttributesRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiFetchControlDetailsRequest struct {
	ctx context.Context
	ApiService *AnalyticsAPIService
	controlId *string
	max *string
	offset *string
}

func (r ApiFetchControlDetailsRequest) ControlId(controlId string) ApiFetchControlDetailsRequest {
	r.controlId = &controlId
	return r
}

func (r ApiFetchControlDetailsRequest) Max(max string) ApiFetchControlDetailsRequest {
	r.max = &max
	return r
}

func (r ApiFetchControlDetailsRequest) Offset(offset string) ApiFetchControlDetailsRequest {
	r.offset = &offset
	return r
}

func (r ApiFetchControlDetailsRequest) Execute() (*http.Response, error) {
	return r.ApiService.FetchControlDetailsExecute(r)
}

/*
FetchControlDetails Fetch Analytics Details

This web service API is used to fetch the details of analytics controls.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFetchControlDetailsRequest
*/
func (a *AnalyticsAPIService) FetchControlDetails(ctx context.Context) ApiFetchControlDetailsRequest {
	return ApiFetchControlDetailsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *AnalyticsAPIService) FetchControlDetailsExecute(r ApiFetchControlDetailsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsAPIService.FetchControlDetails")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fetchControlDetails"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.controlId == nil {
		return nil, reportError("controlId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "controlId", r.controlId, "")
	if r.max != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "max", r.max, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "offset", r.offset, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiFetchControlDetailsESRequest struct {
	ctx context.Context
	ApiService *AnalyticsAPIService
}

func (r ApiFetchControlDetailsESRequest) Execute() (*http.Response, error) {
	return r.ApiService.FetchControlDetailsESExecute(r)
}

/*
FetchControlDetailsES Fetch Analytics Details ES

This web service API is used to fetch the details of analytics controls ES, includes the Offset value in the response.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFetchControlDetailsESRequest
*/
func (a *AnalyticsAPIService) FetchControlDetailsES(ctx context.Context) ApiFetchControlDetailsESRequest {
	return ApiFetchControlDetailsESRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *AnalyticsAPIService) FetchControlDetailsESExecute(r ApiFetchControlDetailsESRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsAPIService.FetchControlDetailsES")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fetchControlDetailsES"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiFetchControlListRequest struct {
	ctx context.Context
	ApiService *AnalyticsAPIService
	fetchControlListRequest *FetchControlListRequest
}

func (r ApiFetchControlListRequest) FetchControlListRequest(fetchControlListRequest FetchControlListRequest) ApiFetchControlListRequest {
	r.fetchControlListRequest = &fetchControlListRequest
	return r
}

func (r ApiFetchControlListRequest) Execute() (*FetchControlListResponse, *http.Response, error) {
	return r.ApiService.FetchControlListExecute(r)
}

/*
FetchControlList Fetch List of Analytics



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFetchControlListRequest
*/
func (a *AnalyticsAPIService) FetchControlList(ctx context.Context) ApiFetchControlListRequest {
	return ApiFetchControlListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FetchControlListResponse
func (a *AnalyticsAPIService) FetchControlListExecute(r ApiFetchControlListRequest) (*FetchControlListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FetchControlListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsAPIService.FetchControlList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fetchControlList"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.fetchControlListRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiFetchControlListESRequest struct {
	ctx context.Context
	ApiService *AnalyticsAPIService
}

func (r ApiFetchControlListESRequest) Execute() (*http.Response, error) {
	return r.ApiService.FetchControlListESExecute(r)
}

/*
FetchControlListES Fetch List of Analytics ES

This method fetches a list of Analytic Controls in Elastic.

Note: Security is based on owner, delegate, `ROLE_ADMIN`, SAV_ROLE's analytics ES category.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFetchControlListESRequest
*/
func (a *AnalyticsAPIService) FetchControlListES(ctx context.Context) ApiFetchControlListESRequest {
	return ApiFetchControlListESRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *AnalyticsAPIService) FetchControlListESExecute(r ApiFetchControlListESRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsAPIService.FetchControlListES")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fetchControlListES"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiFetchRuntimeControlsDataRequest struct {
	ctx context.Context
	ApiService *AnalyticsAPIService
	fetchRuntimeControlsDataRequest *FetchRuntimeControlsDataRequest
}

func (r ApiFetchRuntimeControlsDataRequest) FetchRuntimeControlsDataRequest(fetchRuntimeControlsDataRequest FetchRuntimeControlsDataRequest) ApiFetchRuntimeControlsDataRequest {
	r.fetchRuntimeControlsDataRequest = &fetchRuntimeControlsDataRequest
	return r
}

func (r ApiFetchRuntimeControlsDataRequest) Execute() (*http.Response, error) {
	return r.ApiService.FetchRuntimeControlsDataExecute(r)
}

/*
FetchRuntimeControlsData Fetch Runtime Controls Data

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFetchRuntimeControlsDataRequest
*/
func (a *AnalyticsAPIService) FetchRuntimeControlsData(ctx context.Context) ApiFetchRuntimeControlsDataRequest {
	return ApiFetchRuntimeControlsDataRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *AnalyticsAPIService) FetchRuntimeControlsDataExecute(r ApiFetchRuntimeControlsDataRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsAPIService.FetchRuntimeControlsData")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fetchRuntimeControlsData"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.fetchRuntimeControlsDataRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiFetchRuntimeControlsDataV2Request struct {
	ctx context.Context
	ApiService *AnalyticsAPIService
	fetchRuntimeControlsDataV2Request *FetchRuntimeControlsDataV2Request
}

// 
func (r ApiFetchRuntimeControlsDataV2Request) FetchRuntimeControlsDataV2Request(fetchRuntimeControlsDataV2Request FetchRuntimeControlsDataV2Request) ApiFetchRuntimeControlsDataV2Request {
	r.fetchRuntimeControlsDataV2Request = &fetchRuntimeControlsDataV2Request
	return r
}

func (r ApiFetchRuntimeControlsDataV2Request) Execute() (*http.Response, error) {
	return r.ApiService.FetchRuntimeControlsDataV2Execute(r)
}

/*
FetchRuntimeControlsDataV2 Fetch Runtime Controls Data V2

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFetchRuntimeControlsDataV2Request
*/
func (a *AnalyticsAPIService) FetchRuntimeControlsDataV2(ctx context.Context) ApiFetchRuntimeControlsDataV2Request {
	return ApiFetchRuntimeControlsDataV2Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *AnalyticsAPIService) FetchRuntimeControlsDataV2Execute(r ApiFetchRuntimeControlsDataV2Request) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsAPIService.FetchRuntimeControlsDataV2")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fetchRuntimeControlsDataV2"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.fetchRuntimeControlsDataV2Request
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiRunAnalyticsControlsRequest struct {
	ctx context.Context
	ApiService *AnalyticsAPIService
	jobgroup *string
	jobname *string
	analyticsCategories *string
	analyticsApplications *string
	analyticssubcategories *string
}

func (r ApiRunAnalyticsControlsRequest) Jobgroup(jobgroup string) ApiRunAnalyticsControlsRequest {
	r.jobgroup = &jobgroup
	return r
}

func (r ApiRunAnalyticsControlsRequest) Jobname(jobname string) ApiRunAnalyticsControlsRequest {
	r.jobname = &jobname
	return r
}

func (r ApiRunAnalyticsControlsRequest) AnalyticsCategories(analyticsCategories string) ApiRunAnalyticsControlsRequest {
	r.analyticsCategories = &analyticsCategories
	return r
}

func (r ApiRunAnalyticsControlsRequest) AnalyticsApplications(analyticsApplications string) ApiRunAnalyticsControlsRequest {
	r.analyticsApplications = &analyticsApplications
	return r
}

func (r ApiRunAnalyticsControlsRequest) Analyticssubcategories(analyticssubcategories string) ApiRunAnalyticsControlsRequest {
	r.analyticssubcategories = &analyticssubcategories
	return r
}

func (r ApiRunAnalyticsControlsRequest) Execute() (*http.Response, error) {
	return r.ApiService.RunAnalyticsControlsExecute(r)
}

/*
RunAnalyticsControls Run Analytics Controls

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRunAnalyticsControlsRequest
*/
func (a *AnalyticsAPIService) RunAnalyticsControls(ctx context.Context) ApiRunAnalyticsControlsRequest {
	return ApiRunAnalyticsControlsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *AnalyticsAPIService) RunAnalyticsControlsExecute(r ApiRunAnalyticsControlsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsAPIService.RunAnalyticsControls")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/runAnalyticsControls"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.jobgroup == nil {
		return nil, reportError("jobgroup is required and must be specified")
	}
	if r.jobname == nil {
		return nil, reportError("jobname is required and must be specified")
	}
	if r.analyticsCategories == nil {
		return nil, reportError("analyticsCategories is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "jobgroup", r.jobgroup, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "jobname", r.jobname, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "analyticsCategories", r.analyticsCategories, "")
	if r.analyticsApplications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "analyticsApplications", r.analyticsApplications, "")
	}
	if r.analyticssubcategories != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "analyticssubcategories", r.analyticssubcategories, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
