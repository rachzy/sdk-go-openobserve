# \LogsAPI

All URIs are relative to _http://localhost_

| Method                                                  | HTTP request                                 | Description                        |
| ------------------------------------------------------- | -------------------------------------------- | ---------------------------------- |
| [**LogsIngestionBulk**](LogsAPI.md#LogsIngestionBulk)   | **Post** /api/{org_id}/\_bulk                | \_bulk ES compatible ingestion API |
| [**LogsIngestionJson**](LogsAPI.md#LogsIngestionJson)   | **Post** /api/{org_id}/{stream_name}/\_json  | \_json ingestion API               |
| [**LogsIngestionMulti**](LogsAPI.md#LogsIngestionMulti) | **Post** /api/{org_id}/{stream_name}/\_multi | \_multi ingestion API              |

## LogsIngestionBulk

> BulkResponse LogsIngestionBulk(ctx, orgId).Body(body).Execute()

\_bulk ES compatible ingestion API

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
	body := "body_example" // string | Ingest data (ndjson)

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsAPI.LogsIngestionBulk(context.Background(), orgId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.LogsIngestionBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LogsIngestionBulk`: BulkResponse
	fmt.Fprintf(os.Stdout, "Response from `LogsAPI.LogsIngestionBulk`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiLogsIngestionBulkRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**body** | **string** | Ingest data (ndjson) |

### Return type

[**BulkResponse**](BulkResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## LogsIngestionJson

> IngestionResponse LogsIngestionJson(ctx, orgId, streamName).Body(body).Execute()

\_json ingestion API

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
	streamName := "streamName_example" // string | Stream name
	body := "[{"Athlete":"Alfred","City":"Athens","Country":"HUN","Discipline":"Swimming","Sport":"Aquatics","Year":1896},{"Athlete":"HERSCHMANN","City":"Athens","Country":"CHN","Discipline":"Swimming","Sport":"Aquatics","Year":1896}]" // string | Ingest data (json array)

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsAPI.LogsIngestionJson(context.Background(), orgId, streamName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.LogsIngestionJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LogsIngestionJson`: IngestionResponse
	fmt.Fprintf(os.Stdout, "Response from `LogsAPI.LogsIngestionJson`: %v\n", resp)
}
```

### Path Parameters

| Name           | Type                | Description                                                                 | Notes |
| -------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**        | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**      | **string**          | Organization name                                                           |
| **streamName** | **string**          | Stream name                                                                 |

### Other Parameters

Other parameters are passed through a pointer to a apiLogsIngestionJsonRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**body** | **string** | Ingest data (json array) |

### Return type

[**IngestionResponse**](IngestionResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## LogsIngestionMulti

> IngestionResponse LogsIngestionMulti(ctx, orgId, streamName).Body(body).Execute()

\_multi ingestion API

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
	streamName := "streamName_example" // string | Stream name
	body := "body_example" // string | Ingest data (multiple line json)

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsAPI.LogsIngestionMulti(context.Background(), orgId, streamName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.LogsIngestionMulti``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LogsIngestionMulti`: IngestionResponse
	fmt.Fprintf(os.Stdout, "Response from `LogsAPI.LogsIngestionMulti`: %v\n", resp)
}
```

### Path Parameters

| Name           | Type                | Description                                                                 | Notes |
| -------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**        | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**      | **string**          | Organization name                                                           |
| **streamName** | **string**          | Stream name                                                                 |

### Other Parameters

Other parameters are passed through a pointer to a apiLogsIngestionMultiRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**body** | **string** | Ingest data (multiple line json) |

### Return type

[**IngestionResponse**](IngestionResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
