# \StreamsAPI

All URIs are relative to _http://localhost_

| Method                                                         | HTTP request                                              | Description          |
| -------------------------------------------------------------- | --------------------------------------------------------- | -------------------- |
| [**StreamDelete**](StreamsAPI.md#StreamDelete)                 | **Delete** /api/{org_id}/streams/{stream_name}            | DeleteStream         |
| [**StreamDeleteFields**](StreamsAPI.md#StreamDeleteFields)     | **Put** /api/{org_id}/streams/{stream_name}/delete_fields | DeleteStreamFields   |
| [**StreamList**](StreamsAPI.md#StreamList)                     | **Get** /api/{org_id}/streams                             | ListStreams          |
| [**StreamSchema**](StreamsAPI.md#StreamSchema)                 | **Get** /api/{org_id}/streams/{stream_name}/schema        | GetSchema            |
| [**StreamSettings**](StreamsAPI.md#StreamSettings)             | **Post** /api/{org_id}/streams/{stream_name}/settings     | CreateStreamSettings |
| [**UpdateStreamSettings**](StreamsAPI.md#UpdateStreamSettings) | **Put** /api/{org_id}/streams/{stream_name}/settings      | UpdateStreamSettings |

## StreamDelete

> HttpResponse StreamDelete(ctx, orgId, streamName).Type*(type*).Execute()

DeleteStream

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
	streamName := "streamName_example" // string | Stream name
	type_ := "type__example" // string | Stream type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamsAPI.StreamDelete(context.Background(), orgId, streamName).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.StreamDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamDelete`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.StreamDelete`: %v\n", resp)
}
```

### Path Parameters

| Name           | Type                | Description                                                                 | Notes |
| -------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**        | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**      | **string**          | Organization name                                                           |
| **streamName** | **string**          | Stream name                                                                 |

### Other Parameters

Other parameters are passed through a pointer to a apiStreamDeleteRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**type\_** | **string** | Stream type |

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

## StreamDeleteFields

> HttpResponse StreamDeleteFields(ctx, orgId, streamName).Type*(type*).StreamDeleteFields(streamDeleteFields).Execute()

DeleteStreamFields

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
	streamName := "streamName_example" // string | Stream name
	type_ := "type__example" // string | Stream type
	streamDeleteFields := *openapiclient.NewStreamDeleteFields([]string{"Fields_example"}) // StreamDeleteFields | Stream delete fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamsAPI.StreamDeleteFields(context.Background(), orgId, streamName).Type_(type_).StreamDeleteFields(streamDeleteFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.StreamDeleteFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamDeleteFields`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.StreamDeleteFields`: %v\n", resp)
}
```

### Path Parameters

| Name           | Type                | Description                                                                 | Notes |
| -------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**        | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**      | **string**          | Organization name                                                           |
| **streamName** | **string**          | Stream name                                                                 |

### Other Parameters

Other parameters are passed through a pointer to a apiStreamDeleteFieldsRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**type\_** | **string** | Stream type |
**streamDeleteFields** | [**StreamDeleteFields**](StreamDeleteFields.md) | Stream delete fields |

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

## StreamList

> ListStream StreamList(ctx, orgId).Type*(type*).Keyword(keyword).Offset(offset).Limit(limit).Sort(sort).Execute()

ListStreams

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
	type_ := "type__example" // string | Stream type
	keyword := "keyword_example" // string | Keyword
	offset := int32(56) // int32 | Offset
	limit := int32(56) // int32 | Limit
	sort := "sort_example" // string | Sort

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamsAPI.StreamList(context.Background(), orgId).Type_(type_).Keyword(keyword).Offset(offset).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.StreamList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamList`: ListStream
	fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.StreamList`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiStreamListRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**type\_** | **string** | Stream type |
**keyword** | **string** | Keyword |
**offset** | **int32** | Offset |
**limit** | **int32** | Limit |
**sort** | **string** | Sort |

### Return type

[**ListStream**](ListStream.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## StreamSchema

> Stream StreamSchema(ctx, orgId, streamName).Type*(type*).Keyword(keyword).Offset(offset).Limit(limit).Execute()

GetSchema

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
	streamName := "streamName_example" // string | Stream name
	type_ := "type__example" // string | Stream type
	keyword := "keyword_example" // string | Keyword
	offset := int32(56) // int32 | Offset
	limit := int32(56) // int32 | Limit

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamsAPI.StreamSchema(context.Background(), orgId, streamName).Type_(type_).Keyword(keyword).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.StreamSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamSchema`: Stream
	fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.StreamSchema`: %v\n", resp)
}
```

### Path Parameters

| Name           | Type                | Description                                                                 | Notes |
| -------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**        | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**      | **string**          | Organization name                                                           |
| **streamName** | **string**          | Stream name                                                                 |

### Other Parameters

Other parameters are passed through a pointer to a apiStreamSchemaRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**type\_** | **string** | Stream type |
**keyword** | **string** | Keyword |
**offset** | **int32** | Offset |
**limit** | **int32** | Limit |

### Return type

[**Stream**](Stream.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## StreamSettings

> HttpResponse StreamSettings(ctx, orgId, streamName).Type*(type*).StreamSettings(streamSettings).Execute()

CreateStreamSettings

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
	streamName := "streamName_example" // string | Stream name
	type_ := "type__example" // string | Stream type
	streamSettings := *openapiclient.NewStreamSettings() // StreamSettings | Stream settings

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamsAPI.StreamSettings(context.Background(), orgId, streamName).Type_(type_).StreamSettings(streamSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.StreamSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamSettings`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.StreamSettings`: %v\n", resp)
}
```

### Path Parameters

| Name           | Type                | Description                                                                 | Notes |
| -------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**        | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**      | **string**          | Organization name                                                           |
| **streamName** | **string**          | Stream name                                                                 |

### Other Parameters

Other parameters are passed through a pointer to a apiStreamSettingsRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**type\_** | **string** | Stream type |
**streamSettings** | [**StreamSettings**](StreamSettings.md) | Stream settings |

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

## UpdateStreamSettings

> HttpResponse UpdateStreamSettings(ctx, orgId, streamName).Type*(type*).UpdateStreamSettings(updateStreamSettings).Execute()

UpdateStreamSettings

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
	streamName := "streamName_example" // string | Stream name
	type_ := "type__example" // string | Stream type
	updateStreamSettings := *openapiclient.NewUpdateStreamSettings() // UpdateStreamSettings | Stream settings

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamsAPI.UpdateStreamSettings(context.Background(), orgId, streamName).Type_(type_).UpdateStreamSettings(updateStreamSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.UpdateStreamSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStreamSettings`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.UpdateStreamSettings`: %v\n", resp)
}
```

### Path Parameters

| Name           | Type                | Description                                                                 | Notes |
| -------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**        | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**      | **string**          | Organization name                                                           |
| **streamName** | **string**          | Stream name                                                                 |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStreamSettingsRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**type\_** | **string** | Stream type |
**updateStreamSettings** | [**UpdateStreamSettings**](UpdateStreamSettings.md) | Stream settings |

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
