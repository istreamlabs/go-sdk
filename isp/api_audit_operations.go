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


type AuditOperationsApi interface {

	/*
	GetChannelTimeline Get Channel Timeline

	Returns up to twenty items from the event timeline for a channel, sorted in reverse-chronological order.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId Unique channel identifier
	@return ApiGetChannelTimelineRequest
	*/
	GetChannelTimeline(ctx context.Context, channelId string) ApiGetChannelTimelineRequest

	// GetChannelTimelineExecute executes the request
	//  @return []ChannelTimelineEntry
	GetChannelTimelineExecute(r ApiGetChannelTimelineRequest) ([]ChannelTimelineEntry, *http.Response, error)
}

// AuditOperationsApiService AuditOperationsApi service
type AuditOperationsApiService service

type ApiGetChannelTimelineRequest struct {
	ctx context.Context
	ApiService AuditOperationsApi
	channelId string
	offset *int32
	cursor *string
	pageSize *int32
}

// Number of items to skip when calling a paginated API
func (r ApiGetChannelTimelineRequest) Offset(offset int32) ApiGetChannelTimelineRequest {
	r.offset = &offset
	return r
}

// Current page cursor
func (r ApiGetChannelTimelineRequest) Cursor(cursor string) ApiGetChannelTimelineRequest {
	r.cursor = &cursor
	return r
}

// Number of items to return
func (r ApiGetChannelTimelineRequest) PageSize(pageSize int32) ApiGetChannelTimelineRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetChannelTimelineRequest) Execute() ([]ChannelTimelineEntry, *http.Response, error) {
	return r.ApiService.GetChannelTimelineExecute(r)
}

/*
GetChannelTimeline Get Channel Timeline

Returns up to twenty items from the event timeline for a channel, sorted in reverse-chronological order.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channelId Unique channel identifier
 @return ApiGetChannelTimelineRequest
*/
func (a *AuditOperationsApiService) GetChannelTimeline(ctx context.Context, channelId string) ApiGetChannelTimelineRequest {
	return ApiGetChannelTimelineRequest{
		ApiService: a,
		ctx: ctx,
		channelId: channelId,
	}
}

// Execute executes the request
//  @return []ChannelTimelineEntry
func (a *AuditOperationsApiService) GetChannelTimelineExecute(r ApiGetChannelTimelineRequest) ([]ChannelTimelineEntry, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ChannelTimelineEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditOperationsApiService.GetChannelTimeline")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/channels/{channel-id}/timeline"
	localVarPath = strings.Replace(localVarPath, "{"+"channel-id"+"}", url.PathEscape(parameterToString(r.channelId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.channelId) > 60 {
		return localVarReturnValue, nil, reportError("channelId must have less than 60 elements")
	}

	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
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
			localVarReturnValue = items.([]ChannelTimelineEntry)
			localVarHTTPResponse = resp
		}
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
