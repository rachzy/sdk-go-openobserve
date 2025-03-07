# \RumAPI

All URIs are relative to _http://localhost_

| Method                                                   | HTTP request                     | Description                      |
| -------------------------------------------------------- | -------------------------------- | -------------------------------- |
| [**LogIngestionJson**](RumAPI.md#LogIngestionJson)       | **Post** /rum/v1/{org_id}/logs   | Rum log ingestion API            |
| [**ReplayIngestionJson**](RumAPI.md#ReplayIngestionJson) | **Post** /rum/v1/{org_id}/replay | Rum session-replay ingestion API |
| [**RumIngestionMulti**](RumAPI.md#RumIngestionMulti)     | **Post** /rum/v1/{org_id}/rum    | Rum data ingestion API           |

## LogIngestionJson

> IngestionResponse LogIngestionJson(ctx, orgId).Body(body).Execute()

Rum log ingestion API

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
	body := "[{"Athlete":"Alfred","City":"Athens","Country":"HUN","Discipline":"Swimming","Sport":"Aquatics","Year":1896},{"Athlete":"HERSCHMANN","City":"Athens","Country":"CHN","Discipline":"Swimming","Sport":"Aquatics","Year":1896}]" // string | Ingest data (json array)

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)
	resp, r, err := apiClient.RumAPI.LogIngestionJson(context.Background(), orgId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumAPI.LogIngestionJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LogIngestionJson`: IngestionResponse
	fmt.Fprintf(os.Stdout, "Response from `RumAPI.LogIngestionJson`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiLogIngestionJsonRequest struct via the builder pattern

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

## ReplayIngestionJson

> IngestionResponse ReplayIngestionJson(ctx, orgId).Body(body).Execute()

Rum session-replay ingestion API

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
	body := "[{"Athlete":"Alfred","City":"Athens","Country":"HUN","Discipline":"Swimming","Sport":"Aquatics","Year":1896},{"Athlete":"HERSCHMANN","City":"Athens","Country":"CHN","Discipline":"Swimming","Sport":"Aquatics","Year":1896}]" // string | Ingest data (json array)

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)
	resp, r, err := apiClient.RumAPI.ReplayIngestionJson(context.Background(), orgId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumAPI.ReplayIngestionJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplayIngestionJson`: IngestionResponse
	fmt.Fprintf(os.Stdout, "Response from `RumAPI.ReplayIngestionJson`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiReplayIngestionJsonRequest struct via the builder pattern

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

## RumIngestionMulti

> IngestionResponse RumIngestionMulti(ctx, orgId).Body(body).Execute()

Rum data ingestion API

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
	body := "body_example" // string | Ingest data (multiple line json)

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)
	resp, r, err := apiClient.RumAPI.RumIngestionMulti(context.Background(), orgId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumAPI.RumIngestionMulti``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RumIngestionMulti`: IngestionResponse
	fmt.Fprintf(os.Stdout, "Response from `RumAPI.RumIngestionMulti`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiRumIngestionMultiRequest struct via the builder pattern

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
