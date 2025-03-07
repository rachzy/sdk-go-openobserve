# \DashboardsAPI

All URIs are relative to _http://localhost_

| Method                                                                                | HTTP request                                                                                | Description                    |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------ |
| [**CreateAnnotations**](DashboardsAPI.md#CreateAnnotations)                           | **Post** /api/{org_id}/dashboards/{dashboard_id}/annotations                                | Create Timed Annotations       |
| [**CreateDashboard**](DashboardsAPI.md#CreateDashboard)                               | **Post** /api/{org_id}/dashboards                                                           | CreateDashboard                |
| [**DeleteAnnotations**](DashboardsAPI.md#DeleteAnnotations)                           | **Delete** /api/{org_id}/dashboards/{dashboard_id}/annotations                              | Delete Timed Annotations       |
| [**DeleteDashboard**](DashboardsAPI.md#DeleteDashboard)                               | **Delete** /api/{org_id}/dashboards/{dashboard_id}                                          | DeleteDashboard                |
| [**GetAnnotations**](DashboardsAPI.md#GetAnnotations)                                 | **Get** /api/{org_id}/dashboards/{dashboard_id}/annotations                                 | Get Timed Annotations          |
| [**GetDashboard**](DashboardsAPI.md#GetDashboard)                                     | **Get** /api/{org_id}/dashboards/{dashboard_id}                                             | GetDashboard                   |
| [**ListDashboards**](DashboardsAPI.md#ListDashboards)                                 | **Get** /api/{org_id}/dashboards                                                            | ListDashboards                 |
| [**MoveDashboard**](DashboardsAPI.md#MoveDashboard)                                   | **Put** /api/{org_id}/folders/dashboards/{dashboard_id}                                     | MoveDashboard                  |
| [**MoveDashboards**](DashboardsAPI.md#MoveDashboards)                                 | **Patch** /api/{org_id}/dashboards/move                                                     |
| [**RemoveTimedAnnotationFromPanel**](DashboardsAPI.md#RemoveTimedAnnotationFromPanel) | **Delete** /api/{org_id}/dashboards/{dashboard_id}/annotations/panels/{timed_annotation_id} | Delete Timed Annotation Panels |
| [**UpdateAnnotations**](DashboardsAPI.md#UpdateAnnotations)                           | **Put** /api/{org_id}/dashboards/{dashboard_id}/annotations/{timed_annotation_id}           | Update Timed Annotations       |
| [**UpdateDashboard**](DashboardsAPI.md#UpdateDashboard)                               | **Put** /api/{org_id}/dashboards/{dashboard_id}                                             | UpdateDashboard                |

## CreateAnnotations

> []TimedAnnotation CreateAnnotations(ctx, orgId, dashboardId).TimedAnnotationReq(timedAnnotationReq).Execute()

Create Timed Annotations

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openobserve "github.com/rachzy/sdk-go-openobserve"
)

func main() {
	orgId := "orgId_example" // string |
	dashboardId := "dashboardId_example" // string |
	timedAnnotationReq := *openobserve.NewTimedAnnotationReq([]openobserve.TimedAnnotation{*openobserve.NewTimedAnnotation([]string{"Panels_example"}, int64(123), []string{"Tags_example"}, "Title_example")}) // TimedAnnotationReq | Timed annotation request payload

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.CreateAnnotations(context.Background(), orgId, dashboardId).TimedAnnotationReq(timedAnnotationReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.CreateAnnotations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAnnotations`: []TimedAnnotation
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.CreateAnnotations`: %v\n", resp)
}
```

### Path Parameters

| Name            | Type                | Description                                                                 | Notes |
| --------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**         | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**       | **string**          |                                                                             |
| **dashboardId** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAnnotationsRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**timedAnnotationReq** | [**TimedAnnotationReq**](TimedAnnotationReq.md) | Timed annotation request payload |

### Return type

[**[]TimedAnnotation**](TimedAnnotation.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateDashboard

> DashboardDetails CreateDashboard(ctx, orgId).Body(body).Execute()

CreateDashboard

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openobserve "github.com/rachzy/sdk-go-openobserve"
)

func main() {
	orgId := "orgId_example" // string | Organization name
	body := JsonValue({"description":"Traffic patterns and network performance of the infrastructure","title":"Network Traffic Overview"}) // JsonValue | Dashboard details

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.CreateDashboard(context.Background(), orgId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.CreateDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDashboard`: DashboardDetails
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.CreateDashboard`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDashboardRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**body** | **JsonValue** | Dashboard details |

### Return type

[**DashboardDetails**](DashboardDetails.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteAnnotations

> DeleteAnnotations(ctx, orgId, dashboardId).TimedAnnotationDelete(timedAnnotationDelete).Execute()

Delete Timed Annotations

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openobserve "github.com/rachzy/sdk-go-openobserve"
)

func main() {
	orgId := "orgId_example" // string |
	dashboardId := "dashboardId_example" // string |
	timedAnnotationDelete := *openobserve.NewTimedAnnotationDelete([]string{"AnnotationIds_example"}) // TimedAnnotationDelete | Timed annotation delete request payload

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)
	r, err := apiClient.DashboardsAPI.DeleteAnnotations(context.Background(), orgId, dashboardId).TimedAnnotationDelete(timedAnnotationDelete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.DeleteAnnotations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

| Name            | Type                | Description                                                                 | Notes |
| --------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**         | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**       | **string**          |                                                                             |
| **dashboardId** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnnotationsRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**timedAnnotationDelete** | [**TimedAnnotationDelete**](TimedAnnotationDelete.md) | Timed annotation delete request payload |

### Return type

(empty response body)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteDashboard

> HttpResponse DeleteDashboard(ctx, orgId, dashboardId).Execute()

DeleteDashboard

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openobserve "github.com/rachzy/sdk-go-openobserve"
)

func main() {
	orgId := "orgId_example" // string | Organization name
	dashboardId := "dashboardId_example" // string | Dashboard ID

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.DeleteDashboard(context.Background(), orgId, dashboardId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.DeleteDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDashboard`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.DeleteDashboard`: %v\n", resp)
}
```

### Path Parameters

| Name            | Type                | Description                                                                 | Notes |
| --------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**         | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**       | **string**          | Organization name                                                           |
| **dashboardId** | **string**          | Dashboard ID                                                                |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDashboardRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

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

## GetAnnotations

> []TimedAnnotation GetAnnotations(ctx, orgId, dashboardId).StartTime(startTime).EndTime(endTime).Panels(panels).Execute()

Get Timed Annotations

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openobserve "github.com/rachzy/sdk-go-openobserve"
)

func main() {
	startTime := int64(789) // int64 | Time in microseconds
	endTime := int64(789) // int64 | Time in microseconds
	orgId := "orgId_example" // string |
	dashboardId := "dashboardId_example" // string |
	panels := "panels_example" // string | Commas separated list of panels (optional)

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.GetAnnotations(context.Background(), orgId, dashboardId).StartTime(startTime).EndTime(endTime).Panels(panels).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.GetAnnotations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnnotations`: []TimedAnnotation
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.GetAnnotations`: %v\n", resp)
}
```

### Path Parameters

| Name            | Type                | Description                                                                 | Notes |
| --------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**         | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**       | **string**          |                                                                             |
| **dashboardId** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnnotationsRequest struct via the builder pattern

| Name          | Type      | Description          | Notes |
| ------------- | --------- | -------------------- | ----- |
| **startTime** | **int64** | Time in microseconds |
| **endTime**   | **int64** | Time in microseconds |

**panels** | **string** | Commas separated list of panels |

### Return type

[**[]TimedAnnotation**](TimedAnnotation.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetDashboard

> DashboardDetails GetDashboard(ctx, orgId, dashboardId).Execute()

GetDashboard

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openobserve "github.com/rachzy/sdk-go-openobserve"
)

func main() {
	orgId := "orgId_example" // string | Organization name
	dashboardId := "dashboardId_example" // string | Dashboard ID

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.GetDashboard(context.Background(), orgId, dashboardId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.GetDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDashboard`: DashboardDetails
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.GetDashboard`: %v\n", resp)
}
```

### Path Parameters

| Name            | Type                | Description                                                                 | Notes |
| --------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**         | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**       | **string**          | Organization name                                                           |
| **dashboardId** | **string**          | Dashboard ID                                                                |

### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**DashboardDetails**](DashboardDetails.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListDashboards

> ListDashboardsResponseBody ListDashboards(ctx, orgId).Folder(folder).Title(title).PageSize(pageSize).Execute()

ListDashboards

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openobserve "github.com/rachzy/sdk-go-openobserve"
)

func main() {
	orgId := "orgId_example" // string | Organization name
	folder := "folder_example" // string | Optional folder ID filter parameter  If neither `folder` nor any other filter parameter are set then this will search for all dashboards in the \"default\" folder.  If `folder` is not set and another filter parameter, such as `title`, is set then this will search for dashboards in all folders. (optional)
	title := "title_example" // string | The optional case-insensitive title substring with which to filter dashboards. (optional)
	pageSize := int64(789) // int64 | The optional number of dashboards to retrieve. If not set then all dashboards that match the query parameters will be returned.  Currently this parameter is only untilized by the API when the `title` parameter is also set. (optional)

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.ListDashboards(context.Background(), orgId).Folder(folder).Title(title).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.ListDashboards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDashboards`: ListDashboardsResponseBody
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.ListDashboards`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiListDashboardsRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**folder** | **string** | Optional folder ID filter parameter If neither &#x60;folder&#x60; nor any other filter parameter are set then this will search for all dashboards in the \&quot;default\&quot; folder. If &#x60;folder&#x60; is not set and another filter parameter, such as &#x60;title&#x60;, is set then this will search for dashboards in all folders. |
**title** | **string** | The optional case-insensitive title substring with which to filter dashboards. |
**pageSize** | **int64** | The optional number of dashboards to retrieve. If not set then all dashboards that match the query parameters will be returned. Currently this parameter is only untilized by the API when the &#x60;title&#x60; parameter is also set. |

### Return type

[**ListDashboardsResponseBody**](ListDashboardsResponseBody.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## MoveDashboard

> HttpResponse MoveDashboard(ctx, orgId, dashboardId).MoveDashboardRequestBody(moveDashboardRequestBody).Execute()

MoveDashboard

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openobserve "github.com/rachzy/sdk-go-openobserve"
)

func main() {
	orgId := "orgId_example" // string | Organization name
	dashboardId := "dashboardId_example" // string | Dashboard ID
	moveDashboardRequestBody := *openobserve.NewMoveDashboardRequestBody("From_example", "To_example") // MoveDashboardRequestBody | MoveDashboard details

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.MoveDashboard(context.Background(), orgId, dashboardId).MoveDashboardRequestBody(moveDashboardRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.MoveDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveDashboard`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.MoveDashboard`: %v\n", resp)
}
```

### Path Parameters

| Name            | Type                | Description                                                                 | Notes |
| --------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**         | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**       | **string**          | Organization name                                                           |
| **dashboardId** | **string**          | Dashboard ID                                                                |

### Other Parameters

Other parameters are passed through a pointer to a apiMoveDashboardRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**moveDashboardRequestBody** | [**MoveDashboardRequestBody**](MoveDashboardRequestBody.md) | MoveDashboard details |

### Return type

[**HttpResponse**](HttpResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## MoveDashboards

> HttpResponse MoveDashboards(ctx, orgId).MoveDashboardsRequestBody(moveDashboardsRequestBody).Execute()

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openobserve "github.com/rachzy/sdk-go-openobserve"
)

func main() {
	orgId := "orgId_example" // string | Organization name
	moveDashboardsRequestBody := *openobserve.NewMoveDashboardsRequestBody([]string{"DashboardIds_example"}, "DstFolderId_example") // MoveDashboardsRequestBody | Identifies dashboards and the destination folder

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.MoveDashboards(context.Background(), orgId).MoveDashboardsRequestBody(moveDashboardsRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.MoveDashboards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveDashboards`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.MoveDashboards`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiMoveDashboardsRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**moveDashboardsRequestBody** | [**MoveDashboardsRequestBody**](MoveDashboardsRequestBody.md) | Identifies dashboards and the destination folder |

### Return type

[**HttpResponse**](HttpResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## RemoveTimedAnnotationFromPanel

> RemoveTimedAnnotationFromPanel(ctx, orgId, dashboardId, timedAnnotationId).RequestBody(requestBody).Execute()

Delete Timed Annotation Panels

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openobserve "github.com/rachzy/sdk-go-openobserve"
)

func main() {
	orgId := "orgId_example" // string |
	dashboardId := "dashboardId_example" // string |
	timedAnnotationId := "timedAnnotationId_example" // string |
	requestBody := []string{"Property_example"} // []string | IDs of dashboard panels from which to remove the timed annotation

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)
	r, err := apiClient.DashboardsAPI.RemoveTimedAnnotationFromPanel(context.Background(), orgId, dashboardId, timedAnnotationId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.RemoveTimedAnnotationFromPanel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

| Name                  | Type                | Description                                                                 | Notes |
| --------------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**               | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**             | **string**          |                                                                             |
| **dashboardId**       | **string**          |                                                                             |
| **timedAnnotationId** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTimedAnnotationFromPanelRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**requestBody** | **[]string** | IDs of dashboard panels from which to remove the timed annotation |

### Return type

(empty response body)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateAnnotations

> UpdateAnnotations(ctx, orgId, dashboardId, timedAnnotationId).TimedAnnotation(timedAnnotation).Execute()

Update Timed Annotations

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openobserve "github.com/rachzy/sdk-go-openobserve"
)

func main() {
	orgId := "orgId_example" // string |
	dashboardId := "dashboardId_example" // string |
	timedAnnotationId := "timedAnnotationId_example" // string |
	timedAnnotation := *openobserve.NewTimedAnnotation([]string{"Panels_example"}, int64(123), []string{"Tags_example"}, "Title_example") // TimedAnnotation | Timed annotation update request payload

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)
	r, err := apiClient.DashboardsAPI.UpdateAnnotations(context.Background(), orgId, dashboardId, timedAnnotationId).TimedAnnotation(timedAnnotation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.UpdateAnnotations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

| Name                  | Type                | Description                                                                 | Notes |
| --------------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**               | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**             | **string**          |                                                                             |
| **dashboardId**       | **string**          |                                                                             |
| **timedAnnotationId** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAnnotationsRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**timedAnnotation** | [**TimedAnnotation**](TimedAnnotation.md) | Timed annotation update request payload |

### Return type

(empty response body)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateDashboard

> DashboardDetails UpdateDashboard(ctx, orgId, dashboardId).Body(body).Execute()

UpdateDashboard

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openobserve "github.com/rachzy/sdk-go-openobserve"
)

func main() {
	orgId := "orgId_example" // string | Organization name
	dashboardId := "dashboardId_example" // string | Dashboard ID
	body := JsonValue(987) // JsonValue | Dashboard details

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.UpdateDashboard(context.Background(), orgId, dashboardId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.UpdateDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDashboard`: DashboardDetails
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.UpdateDashboard`: %v\n", resp)
}
```

### Path Parameters

| Name            | Type                | Description                                                                 | Notes |
| --------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**         | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**       | **string**          | Organization name                                                           |
| **dashboardId** | **string**          | Dashboard ID                                                                |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDashboardRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**body** | **JsonValue** | Dashboard details |

### Return type

[**DashboardDetails**](DashboardDetails.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
