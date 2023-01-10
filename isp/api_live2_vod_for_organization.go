/*
 * iStreamPlanet Channels API
 *
 * API version: 0.0.0
 * Contact: support@istreamplanet.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package isp

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


type Live2VODForOrganizationApi interface {

	/*
	ListOrgClips List available clips

	List all clips for all VODs of the channel identified in the request.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param org Organization name
	@param channelId Unique channel identifier
	@return ApiListOrgClipsRequest
	*/
	ListOrgClips(ctx context.Context, org string, channelId string) ApiListOrgClipsRequest

	// ListOrgClipsExecute executes the request
	//  @return ListClipsResponse
	ListOrgClipsExecute(r ApiListOrgClipsRequest) (*ListClipsResponse, *http.Response, error)

	/*
	ListOrgVods List VODs

	VODs can be listed after an appropriately configured channel is turned on for the first time.
If a channel with existing VODs is turned off or deleted, the VODs will still be returned.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param org Organization name
	@param channelId Unique channel identifier
	@return ApiListOrgVodsRequest
	*/
	ListOrgVods(ctx context.Context, org string, channelId string) ApiListOrgVodsRequest

	// ListOrgVodsExecute executes the request
	//  @return []ListVODsResponse
	ListOrgVodsExecute(r ApiListOrgVodsRequest) ([]ListVODsResponse, *http.Response, error)

	/*
	PostOrgClip Make a clip

	Starts a job to create a clip for all VODs for the channel identified in the request.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param org Organization name
	@param channelId Unique channel identifier
	@return ApiPostOrgClipRequest
	*/
	PostOrgClip(ctx context.Context, org string, channelId string) ApiPostOrgClipRequest

	// PostOrgClipExecute executes the request
	//  @return PostClipResponse
	PostOrgClipExecute(r ApiPostOrgClipRequest) (*PostClipResponse, *http.Response, error)

	/*
	PostOrgClipArchive Archive a clip

	Uses archive settings configured in Live2VOD for the organization and on the channel.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param org Organization name
	@param channelId Unique channel identifier
	@param clipId ID for a Clip
	@return ApiPostOrgClipArchiveRequest
	*/
	PostOrgClipArchive(ctx context.Context, org string, channelId string, clipId string) ApiPostOrgClipArchiveRequest

	// PostOrgClipArchiveExecute executes the request
	//  @return PostClipArchiveResponse
	PostOrgClipArchiveExecute(r ApiPostOrgClipArchiveRequest) (*PostClipArchiveResponse, *http.Response, error)
}

// Live2VODForOrganizationApiService Live2VODForOrganizationApi service
type Live2VODForOrganizationApiService service

type ApiListOrgClipsRequest struct {
	ctx context.Context
	ApiService Live2VODForOrganizationApi
	org string
	channelId string
}

func (r ApiListOrgClipsRequest) Execute() (*ListClipsResponse, *http.Response, error) {
	return r.ApiService.ListOrgClipsExecute(r)
}

/*
ListOrgClips List available clips

List all clips for all VODs of the channel identified in the request.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param org Organization name
 @param channelId Unique channel identifier
 @return ApiListOrgClipsRequest
*/
func (a *Live2VODForOrganizationApiService) ListOrgClips(ctx context.Context, org string, channelId string) ApiListOrgClipsRequest {
	return ApiListOrgClipsRequest{
		ApiService: a,
		ctx: ctx,
		org: org,
		channelId: channelId,
	}
}

// Execute executes the request
//  @return ListClipsResponse
func (a *Live2VODForOrganizationApiService) ListOrgClipsExecute(r ApiListOrgClipsRequest) (*ListClipsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListClipsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Live2VODForOrganizationApiService.ListOrgClips")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/{org}/channels/{channel-id}/clips"
	localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", url.PathEscape(parameterToString(r.org, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"channel-id"+"}", url.PathEscape(parameterToString(r.channelId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.channelId) > 60 {
		return localVarReturnValue, nil, reportError("channelId must have less than 60 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 501 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

	if disablePaging := r.ctx.Value(ContextDisablePaging); disablePaging == nil {
		if uri := GetLink(localVarHTTPResponse, RelNext); uri != nil {
			// This response is paginated. Read all the pages and append the items.
			items, resp, err := getAllPages(a.client, localVarReturnValue, localVarHTTPResponse)
			if err.Error() != "" {
				return localVarReturnValue, localVarHTTPResponse, err
			}
			localVarReturnValue = items.(*ListClipsResponse)
			localVarHTTPResponse = resp
		}
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListOrgVodsRequest struct {
	ctx context.Context
	ApiService Live2VODForOrganizationApi
	org string
	channelId string
}

func (r ApiListOrgVodsRequest) Execute() ([]ListVODsResponse, *http.Response, error) {
	return r.ApiService.ListOrgVodsExecute(r)
}

/*
ListOrgVods List VODs

VODs can be listed after an appropriately configured channel is turned on for the first time.
If a channel with existing VODs is turned off or deleted, the VODs will still be returned.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param org Organization name
 @param channelId Unique channel identifier
 @return ApiListOrgVodsRequest
*/
func (a *Live2VODForOrganizationApiService) ListOrgVods(ctx context.Context, org string, channelId string) ApiListOrgVodsRequest {
	return ApiListOrgVodsRequest{
		ApiService: a,
		ctx: ctx,
		org: org,
		channelId: channelId,
	}
}

// Execute executes the request
//  @return []ListVODsResponse
func (a *Live2VODForOrganizationApiService) ListOrgVodsExecute(r ApiListOrgVodsRequest) ([]ListVODsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ListVODsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Live2VODForOrganizationApiService.ListOrgVods")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/{org}/channels/{channel-id}/vods"
	localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", url.PathEscape(parameterToString(r.org, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"channel-id"+"}", url.PathEscape(parameterToString(r.channelId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.channelId) > 60 {
		return localVarReturnValue, nil, reportError("channelId must have less than 60 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 501 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

	if disablePaging := r.ctx.Value(ContextDisablePaging); disablePaging == nil {
		if uri := GetLink(localVarHTTPResponse, RelNext); uri != nil {
			// This response is paginated. Read all the pages and append the items.
			items, resp, err := getAllPages(a.client, localVarReturnValue, localVarHTTPResponse)
			if err.Error() != "" {
				return localVarReturnValue, localVarHTTPResponse, err
			}
			localVarReturnValue = items.([]ListVODsResponse)
			localVarHTTPResponse = resp
		}
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostOrgClipRequest struct {
	ctx context.Context
	ApiService Live2VODForOrganizationApi
	org string
	channelId string
	postClipRequest *PostClipRequest
}

func (r ApiPostOrgClipRequest) PostClipRequest(postClipRequest PostClipRequest) ApiPostOrgClipRequest {
	r.postClipRequest = &postClipRequest
	return r
}

func (r ApiPostOrgClipRequest) Execute() (*PostClipResponse, *http.Response, error) {
	return r.ApiService.PostOrgClipExecute(r)
}

/*
PostOrgClip Make a clip

Starts a job to create a clip for all VODs for the channel identified in the request.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param org Organization name
 @param channelId Unique channel identifier
 @return ApiPostOrgClipRequest
*/
func (a *Live2VODForOrganizationApiService) PostOrgClip(ctx context.Context, org string, channelId string) ApiPostOrgClipRequest {
	return ApiPostOrgClipRequest{
		ApiService: a,
		ctx: ctx,
		org: org,
		channelId: channelId,
	}
}

// Execute executes the request
//  @return PostClipResponse
func (a *Live2VODForOrganizationApiService) PostOrgClipExecute(r ApiPostOrgClipRequest) (*PostClipResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostClipResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Live2VODForOrganizationApiService.PostOrgClip")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/{org}/channels/{channel-id}/clips"
	localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", url.PathEscape(parameterToString(r.org, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"channel-id"+"}", url.PathEscape(parameterToString(r.channelId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.channelId) > 60 {
		return localVarReturnValue, nil, reportError("channelId must have less than 60 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.postClipRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 408 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 413 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 501 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

	if disablePaging := r.ctx.Value(ContextDisablePaging); disablePaging == nil {
		if uri := GetLink(localVarHTTPResponse, RelNext); uri != nil {
			// This response is paginated. Read all the pages and append the items.
			items, resp, err := getAllPages(a.client, localVarReturnValue, localVarHTTPResponse)
			if err.Error() != "" {
				return localVarReturnValue, localVarHTTPResponse, err
			}
			localVarReturnValue = items.(*PostClipResponse)
			localVarHTTPResponse = resp
		}
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostOrgClipArchiveRequest struct {
	ctx context.Context
	ApiService Live2VODForOrganizationApi
	org string
	channelId string
	clipId string
}

func (r ApiPostOrgClipArchiveRequest) Execute() (*PostClipArchiveResponse, *http.Response, error) {
	return r.ApiService.PostOrgClipArchiveExecute(r)
}

/*
PostOrgClipArchive Archive a clip

Uses archive settings configured in Live2VOD for the organization and on the channel.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param org Organization name
 @param channelId Unique channel identifier
 @param clipId ID for a Clip
 @return ApiPostOrgClipArchiveRequest
*/
func (a *Live2VODForOrganizationApiService) PostOrgClipArchive(ctx context.Context, org string, channelId string, clipId string) ApiPostOrgClipArchiveRequest {
	return ApiPostOrgClipArchiveRequest{
		ApiService: a,
		ctx: ctx,
		org: org,
		channelId: channelId,
		clipId: clipId,
	}
}

// Execute executes the request
//  @return PostClipArchiveResponse
func (a *Live2VODForOrganizationApiService) PostOrgClipArchiveExecute(r ApiPostOrgClipArchiveRequest) (*PostClipArchiveResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostClipArchiveResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Live2VODForOrganizationApiService.PostOrgClipArchive")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/{org}/channels/{channel-id}/clips/{clip-id}/archive"
	localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", url.PathEscape(parameterToString(r.org, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"channel-id"+"}", url.PathEscape(parameterToString(r.channelId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clip-id"+"}", url.PathEscape(parameterToString(r.clipId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.channelId) > 60 {
		return localVarReturnValue, nil, reportError("channelId must have less than 60 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 501 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

	if disablePaging := r.ctx.Value(ContextDisablePaging); disablePaging == nil {
		if uri := GetLink(localVarHTTPResponse, RelNext); uri != nil {
			// This response is paginated. Read all the pages and append the items.
			items, resp, err := getAllPages(a.client, localVarReturnValue, localVarHTTPResponse)
			if err.Error() != "" {
				return localVarReturnValue, localVarHTTPResponse, err
			}
			localVarReturnValue = items.(*PostClipArchiveResponse)
			localVarHTTPResponse = resp
		}
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
