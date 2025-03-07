# \ShortUrlAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Retrieve**](ShortUrlAPI.md#Retrieve) | **Get** /short/{org_id}/short/{short_id} | Retrieve the original URL from a short_id
[**Shorten**](ShortUrlAPI.md#Shorten) | **Post** /api/{org_id}/short | Shorten a URL



## Retrieve

> Retrieve(ctx, shortId, orgId).Execute()

Retrieve the original URL from a short_id

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
	shortId := "ddbffcea3ad44292" // string | The short ID to retrieve the original URL
	orgId := "orgId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ShortUrlAPI.Retrieve(context.Background(), shortId, orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShortUrlAPI.Retrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shortId** | **string** | The short ID to retrieve the original URL | 
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Shorten

> ShortenUrlResponse Shorten(ctx, orgId).ShortenUrlRequest(shortenUrlRequest).Execute()

Shorten a URL

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
	shortenUrlRequest := *openapiclient.NewShortenUrlRequest("OriginalUrl_example") // ShortenUrlRequest | The original URL to shorten

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShortUrlAPI.Shorten(context.Background(), orgId).ShortenUrlRequest(shortenUrlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShortUrlAPI.Shorten``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Shorten`: ShortenUrlResponse
	fmt.Fprintf(os.Stdout, "Response from `ShortUrlAPI.Shorten`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShortenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shortenUrlRequest** | [**ShortenUrlRequest**](ShortenUrlRequest.md) | The original URL to shorten | 

### Return type

[**ShortenUrlResponse**](ShortenUrlResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

