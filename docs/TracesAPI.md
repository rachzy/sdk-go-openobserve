# \TracesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLatestTraces**](TracesAPI.md#GetLatestTraces) | **Get** /api/{org_id}/{stream_name}/traces/latest | GetLatestTraces
[**PostTraces**](TracesAPI.md#PostTraces) | **Post** /api/{org_id}/traces | TracesIngest



## GetLatestTraces

> SearchResponse GetLatestTraces(ctx, orgId, streamName).From(from).Size(size).StartTime(startTime).EndTime(endTime).Filter(filter).Timeout(timeout).Execute()

GetLatestTraces

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
	from := int64(789) // int64 | from
	size := int64(789) // int64 | size
	startTime := int64(789) // int64 | start time
	endTime := int64(789) // int64 | end time
	filter := "filter_example" // string | filter, eg: a=b AND c=d (optional)
	timeout := int64(789) // int64 | timeout, seconds (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TracesAPI.GetLatestTraces(context.Background(), orgId, streamName).From(from).Size(size).StartTime(startTime).EndTime(endTime).Filter(filter).Timeout(timeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.GetLatestTraces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLatestTraces`: SearchResponse
	fmt.Fprintf(os.Stdout, "Response from `TracesAPI.GetLatestTraces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization name | 
**streamName** | **string** | Stream name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestTracesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **int64** | from | 
 **size** | **int64** | size | 
 **startTime** | **int64** | start time | 
 **endTime** | **int64** | end time | 
 **filter** | **string** | filter, eg: a&#x3D;b AND c&#x3D;d | 
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


## PostTraces

> IngestionResponse PostTraces(ctx, orgId).Body(body).Execute()

TracesIngest

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
	orgId := "orgId_example" // string | 
	body := "body_example" // string | ExportTraceServiceRequest

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TracesAPI.PostTraces(context.Background(), orgId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.PostTraces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTraces`: IngestionResponse
	fmt.Fprintf(os.Stdout, "Response from `TracesAPI.PostTraces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTracesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | ExportTraceServiceRequest | 

### Return type

[**IngestionResponse**](IngestionResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/x-protobuf
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

