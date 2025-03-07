# \KVAPI

All URIs are relative to _http://localhost_

| Method                                      | HTTP request                      | Description |
| ------------------------------------------- | --------------------------------- | ----------- |
| [**GetKVValue**](KVAPI.md#GetKVValue)       | **Get** /api/{org_id}/kv/{key}    | GetValue    |
| [**ListKVKeys**](KVAPI.md#ListKVKeys)       | **Get** /api/{org_id}/kv          | ListKeys    |
| [**RemoveKVValue**](KVAPI.md#RemoveKVValue) | **Delete** /api/{org_id}/kv/{key} | RemoveValue |
| [**SetKVValue**](KVAPI.md#SetKVValue)       | **Post** /api/{org_id}/kv/{key}   | SetValue    |

## GetKVValue

> string GetKVValue(ctx, orgId, key).Execute()

GetValue

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
	key := "key_example" // string | Key name

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.GetKVValue(context.Background(), orgId, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.GetKVValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKVValue`: string
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.GetKVValue`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |
| **key**   | **string**          | Key name                                                                    |

### Other Parameters

Other parameters are passed through a pointer to a apiGetKVValueRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

**string**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListKVKeys

> []string ListKVKeys(ctx, orgId).Prefix(prefix).Execute()

ListKeys

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
	prefix := "prefix_example" // string | Key prefix (optional)

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.ListKVKeys(context.Background(), orgId).Prefix(prefix).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.ListKVKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKVKeys`: []string
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.ListKVKeys`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiListKVKeysRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**prefix** | **string** | Key prefix |

### Return type

**[]string**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## RemoveKVValue

> string RemoveKVValue(ctx, orgId, key).Execute()

RemoveValue

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
	key := "key_example" // string | Key name

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.RemoveKVValue(context.Background(), orgId, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.RemoveKVValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveKVValue`: string
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.RemoveKVValue`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |
| **key**   | **string**          | Key name                                                                    |

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveKVValueRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

**string**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SetKVValue

> string SetKVValue(ctx, orgId, key).Body(body).Execute()

SetValue

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
	key := "key_example" // string | Key name
	body := "body_example" // string | Value of the key

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.SetKVValue(context.Background(), orgId, key).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.SetKVValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetKVValue`: string
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.SetKVValue`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |
| **key**   | **string**          | Key name                                                                    |

### Other Parameters

Other parameters are passed through a pointer to a apiSetKVValueRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**body** | **string** | Value of the key |

### Return type

**string**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
