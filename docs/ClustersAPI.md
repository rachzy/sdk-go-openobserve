# \ClustersAPI

All URIs are relative to _http://localhost_

| Method                                          | HTTP request          | Description  |
| ----------------------------------------------- | --------------------- | ------------ |
| [**ListClusters**](ClustersAPI.md#ListClusters) | **Get** /api/clusters | ListClusters |

## ListClusters

> map[string][]string ListClusters(ctx).Execute()

ListClusters

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

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.ListClusters(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.ListClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClusters`: map[string][]string
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.ListClusters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListClustersRequest struct via the builder pattern

### Return type

[**map[string][]string**](array.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
