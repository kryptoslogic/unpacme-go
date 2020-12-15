/*
 * UnpacMe
 *
 *  # Introduction Welcome to the UNPACME API! All the malware unpacking and file analysis features that you are familiar with on the [unpac.me](https://www.unpac.me/) website are available through our API. You can easily integrate our unpacker into your malware analysis pipeline and begin unpacking at scale!   # Authentication The public UNPACME API is publicly available and can be accessed without authentication.  In order to use the private UNPACME API you must sign up for an account with UNPACME. Once you have a valid user account you can view your personal API key in your user profile.   <SecurityDefinitions />  # Response Structure When interacting with the UNPACME API, if the request was correctly handled, a <b>200</b> HTTP status code will be returned. The body of the response will usually be a JSON object (except for file downloads).  ## Response Status Codes  Status Code  | Description | Notes ------------- | ------------- | - 200  | OK | The request was successful 400  | Bad Request | The request was somehow incorrect. This can be caused by missing arguments or arguments with wrong values. 401 | Unauthorized | The supplied credentials, if any, are not sufficient to access the resource 403 | Forbidden | The account does not have enough privileges to make the request. 404 | Not Found | The requested resource is not found 429  | Too Many Requests | The request frequency has exceeded one of the account quotas (minute, daily or monthly). Monthly quotas are reset on the 1st of the month at 00:00 UTC. 500 | Server Error | The server could not return the representation due to an internal server error   ## Error Response  If an error has occurred while handling the request an error status code will be returend along with a JSON error message with the following properties.   Property  | Description ------------- | ------------- Error  | The error type Description  | A more informative message  # Example Clients  The following clients can be used to interact with the UNPACME API directly and are provided as examples. These clients are community projects and are not maintained or developed by UNPACME. UNPACME makes no claim as to the safety of these clients, use at your own risk.    - [UnpacMe Client](https://github.com/larsborn/UnpacMeClient) (Python)   - [UnpacMe Library](https://github.com/R3MRUM/unpacme) (Python)  <br> 
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// PublicApiService PublicApi service
type PublicApiService service

type ApiGetPublicFeedRequest struct {
	ctx _context.Context
	ApiService *PublicApiService
}


func (r ApiGetPublicFeedRequest) Execute() (PublicFeed, *_nethttp.Response, error) {
	return r.ApiService.GetPublicFeedExecute(r)
}

/*
 * GetPublicFeed Get public feed
 * Returns public feed of unpacked samples
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetPublicFeedRequest
 */
func (a *PublicApiService) GetPublicFeed(ctx _context.Context) ApiGetPublicFeedRequest {
	return ApiGetPublicFeedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PublicFeed
 */
func (a *PublicApiService) GetPublicFeedExecute(r ApiGetPublicFeedRequest) (PublicFeed, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PublicFeed
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicApiService.GetPublicFeed")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/public/feed"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPublicResultsRequest struct {
	ctx _context.Context
	ApiService *PublicApiService
	unpackId string
}


func (r ApiGetPublicResultsRequest) Execute() (UnpackResults, *_nethttp.Response, error) {
	return r.ApiService.GetPublicResultsExecute(r)
}

/*
 * GetPublicResults Get unpack results by ID
 * Returns unpack results
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param unpackId ID of unpacking submission
 * @return ApiGetPublicResultsRequest
 */
func (a *PublicApiService) GetPublicResults(ctx _context.Context, unpackId string) ApiGetPublicResultsRequest {
	return ApiGetPublicResultsRequest{
		ApiService: a,
		ctx: ctx,
		unpackId: unpackId,
	}
}

/*
 * Execute executes the request
 * @return UnpackResults
 */
func (a *PublicApiService) GetPublicResultsExecute(r ApiGetPublicResultsRequest) (UnpackResults, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UnpackResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicApiService.GetPublicResults")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/public/results/{unpack_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"unpack_id"+"}", _neturl.PathEscape(parameterToString(r.unpackId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPublicUnpackStatusRequest struct {
	ctx _context.Context
	ApiService *PublicApiService
	unpackId string
}


func (r ApiGetPublicUnpackStatusRequest) Execute() (UnpackStatus, *_nethttp.Response, error) {
	return r.ApiService.GetPublicUnpackStatusExecute(r)
}

/*
 * GetPublicUnpackStatus Get unpack status by ID
 * Returns a submission status
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param unpackId ID of unpacking submission
 * @return ApiGetPublicUnpackStatusRequest
 */
func (a *PublicApiService) GetPublicUnpackStatus(ctx _context.Context, unpackId string) ApiGetPublicUnpackStatusRequest {
	return ApiGetPublicUnpackStatusRequest{
		ApiService: a,
		ctx: ctx,
		unpackId: unpackId,
	}
}

/*
 * Execute executes the request
 * @return UnpackStatus
 */
func (a *PublicApiService) GetPublicUnpackStatusExecute(r ApiGetPublicUnpackStatusRequest) (UnpackStatus, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UnpackStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicApiService.GetPublicUnpackStatus")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/public/status/{unpack_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"unpack_id"+"}", _neturl.PathEscape(parameterToString(r.unpackId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
