# \SearchAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchAround**](SearchAPI.md#SearchAround) | **Get** /api/{org_id}/{stream_name}/_around | SearchAround
[**SearchAroundV2**](SearchAPI.md#SearchAroundV2) | **Post** /api/{org_id}/{stream_name}/_around | SearchAroundV2
[**SearchHistory**](SearchAPI.md#SearchHistory) | **Post** /api/{org_id}/_search_history | Search History
[**SearchPartition**](SearchAPI.md#SearchPartition) | **Post** /api/{org_id}/_search_partition | SearchStreamPartition
[**SearchSQL**](SearchAPI.md#SearchSQL) | **Post** /api/{org_id}/_search | SearchStreamData
[**SearchValues**](SearchAPI.md#SearchValues) | **Get** /api/{org_id}/{stream_name}/_values | SearchTopNValues



## SearchAround

> SearchResponse SearchAround(ctx, orgId, streamName).Key(key).Size(size).Regions(regions).Timeout(timeout).Execute()

SearchAround

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
	streamName := "streamName_example" // string | stream_name name
	key := int64(789) // int64 | around key
	size := int64(789) // int64 | around size
	regions := "regions_example" // string | regions, split by comma (optional)
	timeout := int64(789) // int64 | timeout, seconds (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchAround(context.Background(), orgId, streamName).Key(key).Size(size).Regions(regions).Timeout(timeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchAround``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAround`: SearchResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchAround`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization name | 
**streamName** | **string** | stream_name name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchAroundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **key** | **int64** | around key | 
 **size** | **int64** | around size | 
 **regions** | **string** | regions, split by comma | 
 **timeout** | **int64** | timeout, seconds | 

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAroundV2

> SearchResponse SearchAroundV2(ctx, orgId, streamName).Size(size).Body(body).Regions(regions).Timeout(timeout).Execute()

SearchAroundV2

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
	streamName := "streamName_example" // string | stream_name name
	size := int64(789) // int64 | around size
	body := "{"_timestamp":1675182660872049,"container_image":"dkr.ecr.us-west-2.amazonaws.com/openobserve:v0.0.3","container_name":"openobserve","docker_id":"eb0983bdb9ff9360d227e6a0b268fe3b24a0868c2c2d725a1516c11e88bf5789","host":"ip.us-east-2.compute.internal","namespace_name":"openobserve","pod_id":"35a0421f-9203-4d73-9663-9ff0ce26d409","pod_name":"openobserve-ingester-0"}" // string | around record data
	regions := "regions_example" // string | regions, split by comma (optional)
	timeout := int64(789) // int64 | timeout, seconds (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchAroundV2(context.Background(), orgId, streamName).Size(size).Body(body).Regions(regions).Timeout(timeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchAroundV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAroundV2`: SearchResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchAroundV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization name | 
**streamName** | **string** | stream_name name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchAroundV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **size** | **int64** | around size | 
 **body** | **string** | around record data | 
 **regions** | **string** | regions, split by comma | 
 **timeout** | **int64** | timeout, seconds | 

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchHistory

> SearchResponse SearchHistory(ctx, orgId).SearchHistoryRequest(searchHistoryRequest).Execute()

Search History

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
	orgId := "orgId_example" // string | Organization ID
	searchHistoryRequest := *openapiclient.NewSearchHistoryRequest(int64(123), int64(123)) // SearchHistoryRequest | Search history request parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchHistory(context.Background(), orgId).SearchHistoryRequest(searchHistoryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchHistory`: SearchResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchHistoryRequest** | [**SearchHistoryRequest**](SearchHistoryRequest.md) | Search history request parameters | 

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchPartition

> SearchResponse SearchPartition(ctx, orgId).SearchRequest(searchRequest).Execute()

SearchStreamPartition

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
	searchRequest := *openapiclient.NewSearchRequest(*openapiclient.NewSearchQuery("Sql_example")) // SearchRequest | Search query

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchPartition(context.Background(), orgId).SearchRequest(searchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchPartition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchPartition`: SearchResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchPartition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchPartitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchRequest** | [**SearchRequest**](SearchRequest.md) | Search query | 

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSQL

> SearchResponse SearchSQL(ctx, orgId).SearchRequest(searchRequest).Execute()

SearchStreamData

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
	searchRequest := *openapiclient.NewSearchRequest(*openapiclient.NewSearchQuery("Sql_example")) // SearchRequest | Search query

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchSQL(context.Background(), orgId).SearchRequest(searchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchSQL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSQL`: SearchResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchSQL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSQLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchRequest** | [**SearchRequest**](SearchRequest.md) | Search query | 

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchValues

> SearchResponse SearchValues(ctx, orgId, streamName).Fields(fields).Size(size).StartTime(startTime).EndTime(endTime).Filter(filter).Keyword(keyword).Regions(regions).Timeout(timeout).NoCount(noCount).Execute()

SearchTopNValues

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
	streamName := "streamName_example" // string | stream_name name
	fields := "fields_example" // string | fields, split by comma
	size := int64(789) // int64 | size
	startTime := int64(789) // int64 | start time
	endTime := int64(789) // int64 | end time
	filter := "filter_example" // string | filter, eg: a=b (optional)
	keyword := "keyword_example" // string | keyword, eg: abc (optional)
	regions := "regions_example" // string | regions, split by comma (optional)
	timeout := int64(789) // int64 | timeout, seconds (optional)
	noCount := true // bool | no need count, true of false (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchValues(context.Background(), orgId, streamName).Fields(fields).Size(size).StartTime(startTime).EndTime(endTime).Filter(filter).Keyword(keyword).Regions(regions).Timeout(timeout).NoCount(noCount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchValues`: SearchResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization name | 
**streamName** | **string** | stream_name name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | fields, split by comma | 
 **size** | **int64** | size | 
 **startTime** | **int64** | start time | 
 **endTime** | **int64** | end time | 
 **filter** | **string** | filter, eg: a&#x3D;b | 
 **keyword** | **string** | keyword, eg: abc | 
 **regions** | **string** | regions, split by comma | 
 **timeout** | **int64** | timeout, seconds | 
 **noCount** | **bool** | no need count, true of false | 

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

