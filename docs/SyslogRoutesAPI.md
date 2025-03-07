# \SyslogRoutesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSyslogRoute**](SyslogRoutesAPI.md#CreateSyslogRoute) | **Post** /api/{org_id}/syslog-routes | CreateSyslogRoute
[**DeleteSyslogRoute**](SyslogRoutesAPI.md#DeleteSyslogRoute) | **Delete** /api/{org_id}/syslog-routes/{id} | DeleteSyslogRoute
[**ListSyslogRoutes**](SyslogRoutesAPI.md#ListSyslogRoutes) | **Get** /api/{org_id}/syslog-routes | ListSyslogRoutes
[**UpdateSyslogRoute**](SyslogRoutesAPI.md#UpdateSyslogRoute) | **Put** /api/{org_id}/syslog-routes/{id} | UpdateSyslogRoute



## CreateSyslogRoute

> SyslogRoute CreateSyslogRoute(ctx, orgId).SyslogRoute(syslogRoute).Execute()

CreateSyslogRoute

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	orgId := "orgId_example" // string | Organization name
	syslogRoute := *openapiclient.NewSyslogRoute() // SyslogRoute | SyslogRoute details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyslogRoutesAPI.CreateSyslogRoute(context.Background(), orgId).SyslogRoute(syslogRoute).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyslogRoutesAPI.CreateSyslogRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSyslogRoute`: SyslogRoute
	fmt.Fprintf(os.Stdout, "Response from `SyslogRoutesAPI.CreateSyslogRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSyslogRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **syslogRoute** | [**SyslogRoute**](SyslogRoute.md) | SyslogRoute details | 

### Return type

[**SyslogRoute**](SyslogRoute.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSyslogRoute

> HttpResponse DeleteSyslogRoute(ctx, orgId, id).Execute()

DeleteSyslogRoute

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	orgId := "orgId_example" // string | Organization name
	id := "id_example" // string | SyslogRoute Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyslogRoutesAPI.DeleteSyslogRoute(context.Background(), orgId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyslogRoutesAPI.DeleteSyslogRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSyslogRoute`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `SyslogRoutesAPI.DeleteSyslogRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization name | 
**id** | **string** | SyslogRoute Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSyslogRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HttpResponse**](HttpResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyslogRoutes

> SyslogRoutes ListSyslogRoutes(ctx, orgId).Execute()

ListSyslogRoutes

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	orgId := "orgId_example" // string | Organization name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyslogRoutesAPI.ListSyslogRoutes(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyslogRoutesAPI.ListSyslogRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSyslogRoutes`: SyslogRoutes
	fmt.Fprintf(os.Stdout, "Response from `SyslogRoutesAPI.ListSyslogRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSyslogRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SyslogRoutes**](SyslogRoutes.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyslogRoute

> SyslogRoute UpdateSyslogRoute(ctx, orgId, id).SyslogRoute(syslogRoute).Execute()

UpdateSyslogRoute

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	orgId := "orgId_example" // string | Organization name
	id := "id_example" // string | Route ID
	syslogRoute := *openapiclient.NewSyslogRoute() // SyslogRoute | SyslogRoute details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyslogRoutesAPI.UpdateSyslogRoute(context.Background(), orgId, id).SyslogRoute(syslogRoute).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyslogRoutesAPI.UpdateSyslogRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSyslogRoute`: SyslogRoute
	fmt.Fprintf(os.Stdout, "Response from `SyslogRoutesAPI.UpdateSyslogRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization name | 
**id** | **string** | Route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSyslogRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **syslogRoute** | [**SyslogRoute**](SyslogRoute.md) | SyslogRoute details | 

### Return type

[**SyslogRoute**](SyslogRoute.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

