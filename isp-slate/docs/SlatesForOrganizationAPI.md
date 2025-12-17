# \SlatesForOrganizationAPI

All URIs are relative to *https://int.api.istreamplanet.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrgSlate**](SlatesForOrganizationAPI.md#DeleteOrgSlate) | **Delete** /v2/{org}/slates/{slate-id} | Delete Slate
[**GetOrgSlate**](SlatesForOrganizationAPI.md#GetOrgSlate) | **Get** /v2/{org}/slates/{slate-id} | Get Slate
[**ListOrgSlates**](SlatesForOrganizationAPI.md#ListOrgSlates) | **Get** /v2/{org}/slates | List Slates
[**PutOrgSlate**](SlatesForOrganizationAPI.md#PutOrgSlate) | **Put** /v2/{org}/slates/{slate-id} | Create/Update Slate



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
	slateId := "slate-id-1" // string | Unique identifier for this slate
	org := "org_example" // string | Organization name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SlatesForOrganizationAPI.DeleteOrgSlate(context.Background(), slateId, org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlatesForOrganizationAPI.DeleteOrgSlate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slateId** | **string** | Unique identifier for this slate | 
**org** | **string** | Organization name | 

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
	slateId := "slate-id-1" // string | Unique identifier for this slate
	org := "org_example" // string | Organization name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlatesForOrganizationAPI.GetOrgSlate(context.Background(), slateId, org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlatesForOrganizationAPI.GetOrgSlate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgSlate`: Slate
	fmt.Fprintf(os.Stdout, "Response from `SlatesForOrganizationAPI.GetOrgSlate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slateId** | **string** | Unique identifier for this slate | 
**org** | **string** | Organization name | 

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

> []Slate ListOrgSlates(ctx, org).Cursor(cursor).Limit(limit).Execute()

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
	org := "org_example" // string | Organization name
	cursor := "cursor_example" // string | Current page cursor (optional)
	limit := int32(56) // int32 | Number of items to return (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlatesForOrganizationAPI.ListOrgSlates(context.Background(), org).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlatesForOrganizationAPI.ListOrgSlates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgSlates`: []Slate
	fmt.Fprintf(os.Stdout, "Response from `SlatesForOrganizationAPI.ListOrgSlates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | Organization name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgSlatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Current page cursor | 
 **limit** | **int32** | Number of items to return | [default to 100]

### Return type

[**[]Slate**](Slate.md)

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
	slateId := "slate-id-1" // string | Unique identifier for this slate
	org := "org_example" // string | Organization name
	slateWithoutID := *openapiclient.NewSlateWithoutID("Description_example", "Url_example") // SlateWithoutID | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlatesForOrganizationAPI.PutOrgSlate(context.Background(), slateId, org).SlateWithoutID(slateWithoutID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlatesForOrganizationAPI.PutOrgSlate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutOrgSlate`: Slate
	fmt.Fprintf(os.Stdout, "Response from `SlatesForOrganizationAPI.PutOrgSlate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slateId** | **string** | Unique identifier for this slate | 
**org** | **string** | Organization name | 

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

