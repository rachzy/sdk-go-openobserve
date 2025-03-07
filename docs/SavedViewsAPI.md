# \SavedViewsAPI

All URIs are relative to _http://localhost_

| Method                                                    | HTTP request                                  | Description |
| --------------------------------------------------------- | --------------------------------------------- | ----------- |
| [**CreateSavedViews**](SavedViewsAPI.md#CreateSavedViews) | **Post** /api/{org_id}/savedviews             |
| [**DeleteSavedViews**](SavedViewsAPI.md#DeleteSavedViews) | **Delete** /api/{org_id}/savedviews/{view_id} |
| [**GetSavedView**](SavedViewsAPI.md#GetSavedView)         | **Get** /api/{org_id}/savedviews/{view_id}    |
| [**ListSavedViews**](SavedViewsAPI.md#ListSavedViews)     | **Get** /api/{org_id}/savedviews              |
| [**UpdateSavedViews**](SavedViewsAPI.md#UpdateSavedViews) | **Put** /api/{org_id}/savedviews/{view_id}    |

## CreateSavedViews

> CreateViewResponse CreateSavedViews(ctx, orgId).CreateViewRequest(createViewRequest).Execute()

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/rachzy/sdk-go-openobserve"
)

func main() {
	orgId := "orgId_example" // string | Organization name
	createViewRequest := *openapiclient.NewCreateViewRequest(interface{}(123), "ViewName_example") // CreateViewRequest | Create view data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SavedViewsAPI.CreateSavedViews(context.Background(), orgId).CreateViewRequest(createViewRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SavedViewsAPI.CreateSavedViews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSavedViews`: CreateViewResponse
	fmt.Fprintf(os.Stdout, "Response from `SavedViewsAPI.CreateSavedViews`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSavedViewsRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**createViewRequest** | [**CreateViewRequest**](CreateViewRequest.md) | Create view data |

### Return type

[**CreateViewResponse**](CreateViewResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteSavedViews

> DeleteViewResponse DeleteSavedViews(ctx, orgId, viewId).Execute()

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/rachzy/sdk-go-openobserve"
)

func main() {
	orgId := "orgId_example" // string | Organization name
	viewId := "viewId_example" // string | The view_id to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SavedViewsAPI.DeleteSavedViews(context.Background(), orgId, viewId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SavedViewsAPI.DeleteSavedViews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSavedViews`: DeleteViewResponse
	fmt.Fprintf(os.Stdout, "Response from `SavedViewsAPI.DeleteSavedViews`: %v\n", resp)
}
```

### Path Parameters

| Name       | Type                | Description                                                                 | Notes |
| ---------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**    | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**  | **string**          | Organization name                                                           |
| **viewId** | **string**          | The view_id to delete                                                       |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSavedViewsRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**DeleteViewResponse**](DeleteViewResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetSavedView

> View GetSavedView(ctx, orgId, viewId).Execute()

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/rachzy/sdk-go-openobserve"
)

func main() {
	orgId := "orgId_example" // string | Organization name
	viewId := "viewId_example" // string | The view_id which was stored

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SavedViewsAPI.GetSavedView(context.Background(), orgId, viewId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SavedViewsAPI.GetSavedView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSavedView`: View
	fmt.Fprintf(os.Stdout, "Response from `SavedViewsAPI.GetSavedView`: %v\n", resp)
}
```

### Path Parameters

| Name       | Type                | Description                                                                 | Notes |
| ---------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**    | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**  | **string**          | Organization name                                                           |
| **viewId** | **string**          | The view_id which was stored                                                |

### Other Parameters

Other parameters are passed through a pointer to a apiGetSavedViewRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**View**](View.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListSavedViews

> ViewsWithoutData ListSavedViews(ctx, orgId).Execute()

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/rachzy/sdk-go-openobserve"
)

func main() {
	orgId := "orgId_example" // string | Organization name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SavedViewsAPI.ListSavedViews(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SavedViewsAPI.ListSavedViews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSavedViews`: ViewsWithoutData
	fmt.Fprintf(os.Stdout, "Response from `SavedViewsAPI.ListSavedViews`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiListSavedViewsRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**ViewsWithoutData**](ViewsWithoutData.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateSavedViews

> View UpdateSavedViews(ctx, orgId, viewId).UpdateViewRequest(updateViewRequest).Execute()

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/rachzy/sdk-go-openobserve"
)

func main() {
	orgId := "orgId_example" // string | Organization name
	viewId := "viewId_example" // string | View id to be updated
	updateViewRequest := *openapiclient.NewUpdateViewRequest(interface{}(123), "ViewName_example") // UpdateViewRequest | Update view data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SavedViewsAPI.UpdateSavedViews(context.Background(), orgId, viewId).UpdateViewRequest(updateViewRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SavedViewsAPI.UpdateSavedViews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSavedViews`: View
	fmt.Fprintf(os.Stdout, "Response from `SavedViewsAPI.UpdateSavedViews`: %v\n", resp)
}
```

### Path Parameters

| Name       | Type                | Description                                                                 | Notes |
| ---------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**    | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**  | **string**          | Organization name                                                           |
| **viewId** | **string**          | View id to be updated                                                       |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSavedViewsRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**updateViewRequest** | [**UpdateViewRequest**](UpdateViewRequest.md) | Update view data |

### Return type

[**View**](View.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
