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


type SourcesApi interface {

	/*
	GetSource Get Source

	<b>This route is deprecated and is subject to removal any time after `Thu, 06 Apr 2023 19:00:00 UTC`. Use [get-org-source](#get-/v2/-org-/sources/-source-id-) instead.</b>

Get a source's configuration

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sourceId Unique source identifier
	@return ApiGetSourceRequest

	Deprecated
	*/
	GetSource(ctx context.Context, sourceId string) ApiGetSourceRequest

	// GetSourceExecute executes the request
	//  @return Source
	// Deprecated
	GetSourceExecute(r ApiGetSourceRequest) (*Source, *http.Response, error)

	/*
	ListSources List Sources

	<b>This route is deprecated and is subject to removal any time after `Thu, 06 Apr 2023 19:00:00 UTC`. Use [list-org-sources](#get-/v2/-org-/sources) instead.</b>

Get a list of sources that are used to create channels.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListSourcesRequest

	Deprecated
	*/
	ListSources(ctx context.Context) ApiListSourcesRequest

	// ListSourcesExecute executes the request
	//  @return []Summary
	// Deprecated
	ListSourcesExecute(r ApiListSourcesRequest) ([]Summary, *http.Response, error)
}

// SourcesApiService SourcesApi service
type SourcesApiService service

type ApiGetSourceRequest struct {
	ctx context.Context
	ApiService SourcesApi
	sourceId string
}

func (r ApiGetSourceRequest) Execute() (*Source, *http.Response, error) {
	return r.ApiService.GetSourceExecute(r)
}

/*
GetSource Get Source

<b>This route is deprecated and is subject to removal any time after `Thu, 06 Apr 2023 19:00:00 UTC`. Use [get-org-source](#get-/v2/-org-/sources/-source-id-) instead.</b>

Get a source's configuration

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param sourceId Unique source identifier
 @return ApiGetSourceRequest

Deprecated
*/
func (a *SourcesApiService) GetSource(ctx context.Context, sourceId string) ApiGetSourceRequest {
	return ApiGetSourceRequest{
		ApiService: a,
		ctx: ctx,
		sourceId: sourceId,
	}
}

// Execute executes the request
//  @return Source
// Deprecated
func (a *SourcesApiService) GetSourceExecute(r ApiGetSourceRequest) (*Source, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Source
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SourcesApiService.GetSource")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/sources/{source-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"source-id"+"}", url.PathEscape(parameterToString(r.sourceId, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 499 {
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
			localVarReturnValue = items.(*Source)
			localVarHTTPResponse = resp
		}
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListSourcesRequest struct {
	ctx context.Context
	ApiService SourcesApi
	cursor *string
	pageSize *int32
}

// Current page cursor
func (r ApiListSourcesRequest) Cursor(cursor string) ApiListSourcesRequest {
	r.cursor = &cursor
	return r
}

// Number of items to return
func (r ApiListSourcesRequest) PageSize(pageSize int32) ApiListSourcesRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiListSourcesRequest) Execute() ([]Summary, *http.Response, error) {
	return r.ApiService.ListSourcesExecute(r)
}

/*
ListSources List Sources

<b>This route is deprecated and is subject to removal any time after `Thu, 06 Apr 2023 19:00:00 UTC`. Use [list-org-sources](#get-/v2/-org-/sources) instead.</b>

Get a list of sources that are used to create channels.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListSourcesRequest

Deprecated
*/
func (a *SourcesApiService) ListSources(ctx context.Context) ApiListSourcesRequest {
	return ApiListSourcesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Summary
// Deprecated
func (a *SourcesApiService) ListSourcesExecute(r ApiListSourcesRequest) ([]Summary, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Summary
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SourcesApiService.ListSources")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/sources"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
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
		if localVarHTTPResponse.StatusCode == 401 {
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
		if localVarHTTPResponse.StatusCode == 499 {
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
			localVarReturnValue = items.([]Summary)
			localVarHTTPResponse = resp
		}
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
