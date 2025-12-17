# \SlatesForOrganizationApi

All URIs are relative to *https://int.api.istreamplanet.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrgSlate**](SlatesForOrganizationApi.md#DeleteOrgSlate) | **Delete** /v2/{org}/slates/{slate-id} | Delete Slate
[**GetOrgSlate**](SlatesForOrganizationApi.md#GetOrgSlate) | **Get** /v2/{org}/slates/{slate-id} | Get Slate
[**ListOrgSlates**](SlatesForOrganizationApi.md#ListOrgSlates) | **Get** /v2/{org}/slates | List Slates
[**PutOrgSlate**](SlatesForOrganizationApi.md#PutOrgSlate) | **Put** /v2/{org}/slates/{slate-id} | Create/Update Slate



## DeleteOrgSlate

> DeleteOrgSlate(ctx, slateId, org).Execute()

Delete Slate



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/istreamlabs/go-sdk/isp"
)

func main() {
    slateId := TODO // interface{} | Unique identifier for this slate
    org := TODO // interface{} | Organization name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SlatesForOrganizationApi.DeleteOrgSlate(context.Background(), slateId, org).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlatesForOrganizationApi.DeleteOrgSlate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slateId** | [**interface{}**](.md) | Unique identifier for this slate | 
**org** | [**interface{}**](.md) | Organization name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgSlateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[authcode](../README.md#authcode), [m2m](../README.md#m2m)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgSlate

> Slate GetOrgSlate(ctx, slateId, org).Execute()

Get Slate



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/istreamlabs/go-sdk/isp"
)

func main() {
    slateId := TODO // interface{} | Unique identifier for this slate
    org := TODO // interface{} | Organization name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlatesForOrganizationApi.GetOrgSlate(context.Background(), slateId, org).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlatesForOrganizationApi.GetOrgSlate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgSlate`: Slate
    fmt.Fprintf(os.Stdout, "Response from `SlatesForOrganizationApi.GetOrgSlate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slateId** | [**interface{}**](.md) | Unique identifier for this slate | 
**org** | [**interface{}**](.md) | Organization name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgSlateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Slate**](Slate.md)

### Authorization

[authcode](../README.md#authcode), [m2m](../README.md#m2m)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgSlates

> interface{} ListOrgSlates(ctx, org).Cursor(cursor).Limit(limit).Execute()

List Slates



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/istreamlabs/go-sdk/isp"
)

func main() {
    org := TODO // interface{} | Organization name
    cursor := TODO // interface{} | Current page cursor (optional)
    limit := TODO // interface{} | Number of items to return (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlatesForOrganizationApi.ListOrgSlates(context.Background(), org).Cursor(cursor).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlatesForOrganizationApi.ListOrgSlates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrgSlates`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `SlatesForOrganizationApi.ListOrgSlates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | [**interface{}**](.md) | Organization name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgSlatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | [**interface{}**](interface{}.md) | Current page cursor | 
 **limit** | [**interface{}**](interface{}.md) | Number of items to return | [default to 100]

### Return type

**interface{}**

### Authorization

[authcode](../README.md#authcode), [m2m](../README.md#m2m)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutOrgSlate

> Slate PutOrgSlate(ctx, slateId, org).SlateWithoutID(slateWithoutID).Execute()

Create/Update Slate



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/istreamlabs/go-sdk/isp"
)

func main() {
    slateId := TODO // interface{} | Unique identifier for this slate
    org := TODO // interface{} | Organization name
    slateWithoutID := *openapiclient.NewSlateWithoutID(interface{}(123), interface{}(123)) // SlateWithoutID | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlatesForOrganizationApi.PutOrgSlate(context.Background(), slateId, org).SlateWithoutID(slateWithoutID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlatesForOrganizationApi.PutOrgSlate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutOrgSlate`: Slate
    fmt.Fprintf(os.Stdout, "Response from `SlatesForOrganizationApi.PutOrgSlate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slateId** | [**interface{}**](.md) | Unique identifier for this slate | 
**org** | [**interface{}**](.md) | Organization name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutOrgSlateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **slateWithoutID** | [**SlateWithoutID**](SlateWithoutID.md) |  | 

### Return type

[**Slate**](Slate.md)

### Authorization

[authcode](../README.md#authcode), [m2m](../README.md#m2m)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

