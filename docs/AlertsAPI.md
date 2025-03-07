# \AlertsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlert**](AlertsAPI.md#CreateAlert) | **Post** /api/v2/{org_id}/alerts | CreateAlert
[**CreateDestination**](AlertsAPI.md#CreateDestination) | **Post** /api/{org_id}/alerts/destinations | CreateDestination
[**DeleteAlert**](AlertsAPI.md#DeleteAlert) | **Delete** /api/v2/{org_id}/alerts/{alert_id} | DeleteAlert
[**DeleteAlertDestination**](AlertsAPI.md#DeleteAlertDestination) | **Delete** /api/{org_id}/alerts/destinations/{destination_name} | DeleteDestination
[**DeleteAlert_0**](AlertsAPI.md#DeleteAlert_0) | **Delete** /api/{org_id}/{stream_name}/alerts/{alert_name} | DeleteAlert
[**EnableAlert**](AlertsAPI.md#EnableAlert) | **Patch** /api/v2/{org_id}/alerts/{alert_id}/enable | EnableAlert
[**EnableAlert_0**](AlertsAPI.md#EnableAlert_0) | **Put** /api/{org_id}/{stream_name}/alerts/{alert_name}/enable | EnableAlert
[**GetAlert**](AlertsAPI.md#GetAlert) | **Get** /api/v2/{org_id}/alerts/{alert_id} | GetAlert
[**GetAlert_0**](AlertsAPI.md#GetAlert_0) | **Get** /api/{org_id}/{stream_name}/alerts/{alert_name} | GetAlertByName
[**GetDestination**](AlertsAPI.md#GetDestination) | **Get** /api/{org_id}/alerts/destinations/{destination_name} | GetDestination
[**ListAlerts**](AlertsAPI.md#ListAlerts) | **Get** /api/v2/{org_id}/alerts | ListAlerts
[**ListAlerts_0**](AlertsAPI.md#ListAlerts_0) | **Get** /api/{org_id}/alerts | ListAlerts
[**ListDestinations**](AlertsAPI.md#ListDestinations) | **Get** /api/{org_id}/alerts/destinations | ListDestinations
[**ListStreamAlerts**](AlertsAPI.md#ListStreamAlerts) | **Get** /api/{org_id}/{stream_name}/alerts | ListStreamAlerts
[**MoveAlerts**](AlertsAPI.md#MoveAlerts) | **Patch** /api/v2/{org_id}/alerts/move | MoveAlerts
[**SaveAlert**](AlertsAPI.md#SaveAlert) | **Post** /api/{org_id}/{stream_name}/alerts | CreateAlert
[**TriggerAlert**](AlertsAPI.md#TriggerAlert) | **Patch** /api/v2/{org_id}/alerts/{alert_id}/trigger | TriggerAlert
[**TriggerAlert_0**](AlertsAPI.md#TriggerAlert_0) | **Put** /api/{org_id}/{stream_name}/alerts/{alert_name}/trigger | TriggerAlert
[**UpdateAlert**](AlertsAPI.md#UpdateAlert) | **Put** /api/v2/{org_id}/alerts/{alert_id} | UpdateAlert
[**UpdateAlert_0**](AlertsAPI.md#UpdateAlert_0) | **Put** /api/{org_id}/{stream_name}/alerts/{alert_name} | UpdateAlert
[**UpdateDestination**](AlertsAPI.md#UpdateDestination) | **Put** /api/{org_id}/alerts/destinations/{destination_name} | UpdateDestination



## CreateAlert

> HttpResponse CreateAlert(ctx, orgId).CreateAlertRequestBody(createAlertRequestBody).Execute()

CreateAlert

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
	createAlertRequestBody := *openapiclient.NewCreateAlertRequestBody([]string{"Destinations_example"}) // CreateAlertRequestBody | Alert data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.CreateAlert(context.Background(), orgId).CreateAlertRequestBody(createAlertRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.CreateAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAlert`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.CreateAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAlertRequestBody** | [**CreateAlertRequestBody**](CreateAlertRequestBody.md) | Alert data | 

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


## CreateDestination

> HttpResponse CreateDestination(ctx, orgId).Destination(destination).Execute()

CreateDestination

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
	destination := *openapiclient.NewDestination() // Destination | Destination data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.CreateDestination(context.Background(), orgId).Destination(destination).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.CreateDestination``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDestination`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.CreateDestination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **destination** | [**Destination**](Destination.md) | Destination data | 

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


## DeleteAlert

> HttpResponse DeleteAlert(ctx, orgId, alertId).Execute()

DeleteAlert

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
	alertId := TODO // Ksuid | Alert ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.DeleteAlert(context.Background(), orgId, alertId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.DeleteAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAlert`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.DeleteAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization name | 
**alertId** | [**Ksuid**](.md) | Alert ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertRequest struct via the builder pattern


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


## DeleteAlertDestination

> HttpResponse DeleteAlertDestination(ctx, orgId, destinationName).Execute()

DeleteDestination

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
	destinationName := "destinationName_example" // string | Destination name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.DeleteAlertDestination(context.Background(), orgId, destinationName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.DeleteAlertDestination``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAlertDestination`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.DeleteAlertDestination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization name | 
**destinationName** | **string** | Destination name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertDestinationRequest struct via the builder pattern


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


## DeleteAlert_0

> HttpResponse DeleteAlert_0(ctx, orgId, streamName, alertName).Execute()

DeleteAlert

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
	streamName := "streamName_example" // string | Stream name
	alertName := "alertName_example" // string | Alert name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.DeleteAlert_0(context.Background(), orgId, streamName, alertName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.DeleteAlert_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAlert_0`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.DeleteAlert_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization name | 
**streamName** | **string** | Stream name | 
**alertName** | **string** | Alert name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlert_1Request struct via the builder pattern


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


## EnableAlert

> HttpResponse EnableAlert(ctx, orgId, alertId).Value(value).Execute()

EnableAlert

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
	alertId := TODO // Ksuid | Alert ID
	value := true // bool | Set to `true` to enable the alert or `false` to disable the alert.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.EnableAlert(context.Background(), orgId, alertId).Value(value).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.EnableAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableAlert`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.EnableAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization name | 
**alertId** | [**Ksuid**](.md) | Alert ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **value** | **bool** | Set to &#x60;true&#x60; to enable the alert or &#x60;false&#x60; to disable the alert. | 

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


## EnableAlert_0

> HttpResponse EnableAlert_0(ctx, orgId, streamName, alertName).Value(value).Execute()

EnableAlert

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
	streamName := "streamName_example" // string | Stream name
	alertName := "alertName_example" // string | Alert name
	value := true // bool | Enable or disable alert

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.EnableAlert_0(context.Background(), orgId, streamName, alertName).Value(value).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.EnableAlert_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableAlert_0`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.EnableAlert_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization name | 
**streamName** | **string** | Stream name | 
**alertName** | **string** | Alert name | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableAlert_2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **value** | **bool** | Enable or disable alert | 

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


## GetAlert

> Alert GetAlert(ctx, orgId, alertId).Execute()

GetAlert

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
	alertId := TODO // Ksuid | Alert ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetAlert(context.Background(), orgId, alertId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlert`: Alert
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization name | 
**alertId** | [**Ksuid**](.md) | Alert ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Alert**](Alert.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlert_0

> Alert GetAlert_0(ctx, orgId, streamName, alertName).Execute()

GetAlertByName

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
	streamName := "streamName_example" // string | Stream name
	alertName := "alertName_example" // string | Alert name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetAlert_0(context.Background(), orgId, streamName, alertName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetAlert_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlert_0`: Alert
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetAlert_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization name | 
**streamName** | **string** | Stream name | 
**alertName** | **string** | Alert name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlert_3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Alert**](Alert.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDestination

> Destination GetDestination(ctx, orgId, destinationName).Execute()

GetDestination

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
	destinationName := "destinationName_example" // string | Destination name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetDestination(context.Background(), orgId, destinationName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetDestination``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDestination`: Destination
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetDestination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization name | 
**destinationName** | **string** | Destination name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Destination**](Destination.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlerts

> ListAlertsResponseBody ListAlerts(ctx, orgId).Folder(folder).StreamType(streamType).StreamName(streamName).AlertNameSubstring(alertNameSubstring).Owner(owner).Enabled(enabled).PageSize(pageSize).PageIdx(pageIdx).Execute()

ListAlerts

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
	folder := "folder_example" // string | Optional folder ID filter parameter. (optional)
	streamType := openapiclient.StreamType("logs") // StreamType | Optional stream type filter parameter. (optional)
	streamName := "streamName_example" // string | Optional stream name filter parameter.  This parameter is only used if `stream_type` is also provided. (optional)
	alertNameSubstring := "alertNameSubstring_example" // string | Optional case-insensitive name substring filter parameter. (optional)
	owner := "owner_example" // string | Optional owner user filter parameter. (optional)
	enabled := true // bool | Optional enabled filter parameter. (optional)
	pageSize := int64(789) // int64 | The optional number of alerts to retrieve. If not set then all alerts that match the query parameters will be returned. (optional)
	pageIdx := int64(789) // int64 | The optional page index. If not set then defaults to `0`.  This parameter is only used if `page_size` is also set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.ListAlerts(context.Background(), orgId).Folder(folder).StreamType(streamType).StreamName(streamName).AlertNameSubstring(alertNameSubstring).Owner(owner).Enabled(enabled).PageSize(pageSize).PageIdx(pageIdx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.ListAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAlerts`: ListAlertsResponseBody
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.ListAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **folder** | **string** | Optional folder ID filter parameter. | 
 **streamType** | [**StreamType**](StreamType.md) | Optional stream type filter parameter. | 
 **streamName** | **string** | Optional stream name filter parameter.  This parameter is only used if &#x60;stream_type&#x60; is also provided. | 
 **alertNameSubstring** | **string** | Optional case-insensitive name substring filter parameter. | 
 **owner** | **string** | Optional owner user filter parameter. | 
 **enabled** | **bool** | Optional enabled filter parameter. | 
 **pageSize** | **int64** | The optional number of alerts to retrieve. If not set then all alerts that match the query parameters will be returned. | 
 **pageIdx** | **int64** | The optional page index. If not set then defaults to &#x60;0&#x60;.  This parameter is only used if &#x60;page_size&#x60; is also set. | 

### Return type

[**ListAlertsResponseBody**](ListAlertsResponseBody.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlerts_0

> HttpResponse ListAlerts_0(ctx, orgId).Execute()

ListAlerts

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
	resp, r, err := apiClient.AlertsAPI.ListAlerts_0(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.ListAlerts_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAlerts_0`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.ListAlerts_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAlerts_4Request struct via the builder pattern


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


## ListDestinations

> []Destination ListDestinations(ctx, orgId).Module(module).Execute()

ListDestinations

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
	module := "module_example" // string | Destination module filter, none, alert, or pipeline (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.ListDestinations(context.Background(), orgId).Module(module).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.ListDestinations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDestinations`: []Destination
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.ListDestinations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDestinationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **module** | **string** | Destination module filter, none, alert, or pipeline | 

### Return type

[**[]Destination**](Destination.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStreamAlerts

> HttpResponse ListStreamAlerts(ctx, orgId, streamName).Execute()

ListStreamAlerts

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
	streamName := "streamName_example" // string | Stream name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.ListStreamAlerts(context.Background(), orgId, streamName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.ListStreamAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStreamAlerts`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.ListStreamAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization name | 
**streamName** | **string** | Stream name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStreamAlertsRequest struct via the builder pattern


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


## MoveAlerts

> HttpResponse MoveAlerts(ctx, orgId).MoveAlertsRequestBody(moveAlertsRequestBody).Execute()

MoveAlerts

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
	moveAlertsRequestBody := *openapiclient.NewMoveAlertsRequestBody([]Ksuid{"TODO"}, "DstFolderId_example") // MoveAlertsRequestBody | Identifies alerts and the destination folder

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.MoveAlerts(context.Background(), orgId).MoveAlertsRequestBody(moveAlertsRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.MoveAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveAlerts`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.MoveAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization name | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **moveAlertsRequestBody** | [**MoveAlertsRequestBody**](MoveAlertsRequestBody.md) | Identifies alerts and the destination folder | 

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


## SaveAlert

> HttpResponse SaveAlert(ctx, orgId, streamName).Alert(alert).Execute()

CreateAlert

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
	streamName := "streamName_example" // string | Stream name
	alert := *openapiclient.NewAlert([]string{"Destinations_example"}) // Alert | Alert data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.SaveAlert(context.Background(), orgId, streamName).Alert(alert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.SaveAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveAlert`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.SaveAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization name | 
**streamName** | **string** | Stream name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **alert** | [**Alert**](Alert.md) | Alert data | 

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


## TriggerAlert

> HttpResponse TriggerAlert(ctx, orgId, alertId).Execute()

TriggerAlert

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
	alertId := TODO // Ksuid | Alert ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.TriggerAlert(context.Background(), orgId, alertId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.TriggerAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggerAlert`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.TriggerAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization name | 
**alertId** | [**Ksuid**](.md) | Alert ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTriggerAlertRequest struct via the builder pattern


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


## TriggerAlert_0

> HttpResponse TriggerAlert_0(ctx, orgId, streamName, alertName).Execute()

TriggerAlert

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
	streamName := "streamName_example" // string | Stream name
	alertName := "alertName_example" // string | Alert name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.TriggerAlert_0(context.Background(), orgId, streamName, alertName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.TriggerAlert_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggerAlert_0`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.TriggerAlert_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization name | 
**streamName** | **string** | Stream name | 
**alertName** | **string** | Alert name | 

### Other Parameters

Other parameters are passed through a pointer to a apiTriggerAlert_5Request struct via the builder pattern


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


## UpdateAlert

> HttpResponse UpdateAlert(ctx, orgId, alertId).Body(body).Execute()

UpdateAlert

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
	alertId := TODO // Ksuid | Alert ID
	body := Alert(987) // Alert | Alert data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.UpdateAlert(context.Background(), orgId, alertId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.UpdateAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAlert`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.UpdateAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization name | 
**alertId** | [**Ksuid**](.md) | Alert ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **Alert** | Alert data | 

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


## UpdateAlert_0

> HttpResponse UpdateAlert_0(ctx, orgId, streamName, alertName).Alert(alert).Execute()

UpdateAlert

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
	streamName := "streamName_example" // string | Stream name
	alertName := "alertName_example" // string | Alert name
	alert := *openapiclient.NewAlert([]string{"Destinations_example"}) // Alert | Alert data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.UpdateAlert_0(context.Background(), orgId, streamName, alertName).Alert(alert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.UpdateAlert_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAlert_0`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.UpdateAlert_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization name | 
**streamName** | **string** | Stream name | 
**alertName** | **string** | Alert name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlert_6Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **alert** | [**Alert**](Alert.md) | Alert data | 

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


## UpdateDestination

> HttpResponse UpdateDestination(ctx, orgId, destinationName).Destination(destination).Execute()

UpdateDestination

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
	destinationName := "destinationName_example" // string | Destination name
	destination := *openapiclient.NewDestination() // Destination | Destination data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.UpdateDestination(context.Background(), orgId, destinationName).Destination(destination).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.UpdateDestination``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDestination`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.UpdateDestination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization name | 
**destinationName** | **string** | Destination name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **destination** | [**Destination**](Destination.md) | Destination data | 

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

