# \MetricsAPI

All URIs are relative to _http://localhost_

| Method                                                           | HTTP request                                                      | Description                                  |
| ---------------------------------------------------------------- | ----------------------------------------------------------------- | -------------------------------------------- |
| [**MetricsIngestionJson**](MetricsAPI.md#MetricsIngestionJson)   | **Post** /api/{org_id}/ingest/metrics/\_json                      | \_json ingestion API                         |
| [**PrometheusFormatQuery**](MetricsAPI.md#PrometheusFormatQuery) | **Get** /api/{org_id}/prometheus/api/v1/format_query              | prometheus formatting query expressions      |
| [**PrometheusLabelValues**](MetricsAPI.md#PrometheusLabelValues) | **Get** /api/{org_id}/prometheus/api/v1/label/{label_name}/values | prometheus query label values                |
| [**PrometheusLabels**](MetricsAPI.md#PrometheusLabels)           | **Get** /api/{org_id}/prometheus/api/v1/labels                    | prometheus getting label names               |
| [**PrometheusMetadata**](MetricsAPI.md#PrometheusMetadata)       | **Get** /api/{org_id}/prometheus/api/v1/metadata                  | prometheus query metric metadata             |
| [**PrometheusQuery**](MetricsAPI.md#PrometheusQuery)             | **Get** /api/{org_id}/prometheus/api/v1/query                     | prometheus instant queries                   |
| [**PrometheusRangeQuery**](MetricsAPI.md#PrometheusRangeQuery)   | **Get** /api/{org_id}/prometheus/api/v1/query_range               | prometheus range queries                     |
| [**PrometheusRemoteWrite**](MetricsAPI.md#PrometheusRemoteWrite) | **Post** /api/{org_id}/prometheus/api/v1/write                    | prometheus remote-write endpoint for metrics |
| [**PrometheusSeries**](MetricsAPI.md#PrometheusSeries)           | **Get** /api/{org_id}/prometheus/api/v1/series                    | prometheus finding series by label matchers  |

## MetricsIngestionJson

> IngestionResponse MetricsIngestionJson(ctx, orgId).Body(body).Execute()

\_json ingestion API

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
	body := "[{"__name__":"metrics stream name","__type__":"counter / gauge / histogram / summary","_timestamp":1687175143,"label_name1":"label_value1","label_name2":"label_value2","value":1.2}]" // string | Ingest data (json array)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.MetricsIngestionJson(context.Background(), orgId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.MetricsIngestionJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetricsIngestionJson`: IngestionResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.MetricsIngestionJson`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiMetricsIngestionJsonRequest struct via the builder pattern

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

## PrometheusFormatQuery

> HttpResponse PrometheusFormatQuery(ctx, orgId).Query(query).Execute()

prometheus formatting query expressions

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
	query := "query_example" // string | Prometheus expression query string.
	orgId := "orgId_example" // string |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.PrometheusFormatQuery(context.Background(), orgId).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.PrometheusFormatQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrometheusFormatQuery`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.PrometheusFormatQuery`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiPrometheusFormatQueryRequest struct via the builder pattern

| Name      | Type       | Description                         | Notes |
| --------- | ---------- | ----------------------------------- | ----- |
| **query** | **string** | Prometheus expression query string. |

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

## PrometheusLabelValues

> HttpResponse PrometheusLabelValues(ctx, orgId, labelName).Match(match).Start(start).End(end).Execute()

prometheus query label values

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
	labelName := "labelName_example" // string | Label name
	match := "match_example" // string | Series selector argument that selects the series from which to read the label values
	start := "start_example" // string | <rfc3339 | unix_timestamp>: Start timestamp (optional)
	end := "end_example" // string | <rfc3339 | unix_timestamp>: End timestamp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.PrometheusLabelValues(context.Background(), orgId, labelName).Match(match).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.PrometheusLabelValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrometheusLabelValues`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.PrometheusLabelValues`: %v\n", resp)
}
```

### Path Parameters

| Name          | Type                | Description                                                                 | Notes |
| ------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**       | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**     | **string**          | Organization name                                                           |
| **labelName** | **string**          | Label name                                                                  |

### Other Parameters

Other parameters are passed through a pointer to a apiPrometheusLabelValuesRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**match** | **string** | Series selector argument that selects the series from which to read the label values |
**start** | **string** | &lt;rfc3339 | unix_timestamp&gt;: Start timestamp |
**end** | **string** | &lt;rfc3339 | unix_timestamp&gt;: End timestamp |

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

## PrometheusLabels

> HttpResponse PrometheusLabels(ctx, orgId).Match(match).Start(start).End(end).Execute()

prometheus getting label names

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
	match := "match_example" // string | Series selector argument that selects the series from which to read the label names
	start := "start_example" // string | <rfc3339 | unix_timestamp>: Start timestamp (optional)
	end := "end_example" // string | <rfc3339 | unix_timestamp>: End timestamp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.PrometheusLabels(context.Background(), orgId).Match(match).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.PrometheusLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrometheusLabels`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.PrometheusLabels`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiPrometheusLabelsRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**match** | **string** | Series selector argument that selects the series from which to read the label names |
**start** | **string** | &lt;rfc3339 | unix_timestamp&gt;: Start timestamp |
**end** | **string** | &lt;rfc3339 | unix_timestamp&gt;: End timestamp |

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

## PrometheusMetadata

> HttpResponse PrometheusMetadata(ctx, orgId).Limit(limit).Metric(metric).Execute()

prometheus query metric metadata

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
	limit := "limit_example" // string | Maximum number of metrics to return
	metric := "metric_example" // string | A metric name to filter metadata for. All metric metadata is retrieved if left empty (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.PrometheusMetadata(context.Background(), orgId).Limit(limit).Metric(metric).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.PrometheusMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrometheusMetadata`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.PrometheusMetadata`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiPrometheusMetadataRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**limit** | **string** | Maximum number of metrics to return |
**metric** | **string** | A metric name to filter metadata for. All metric metadata is retrieved if left empty |

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

## PrometheusQuery

> HttpResponse PrometheusQuery(ctx, orgId).Query(query).Time(time).Timeout(timeout).Execute()

prometheus instant queries

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
	query := "query_example" // string | Prometheus expression query string
	time := "time_example" // string | <rfc3339 | unix_timestamp>: Evaluation timestamp. Optional (optional)
	timeout := "timeout_example" // string | Evaluation timeout (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.PrometheusQuery(context.Background(), orgId).Query(query).Time(time).Timeout(timeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.PrometheusQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrometheusQuery`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.PrometheusQuery`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiPrometheusQueryRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**query** | **string** | Prometheus expression query string |
**time** | **string** | &lt;rfc3339 | unix_timestamp&gt;: Evaluation timestamp. Optional |
**timeout** | **string** | Evaluation timeout |

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

## PrometheusRangeQuery

> HttpResponse PrometheusRangeQuery(ctx, orgId).Query(query).Start(start).End(end).Step(step).Timeout(timeout).Execute()

prometheus range queries

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
	query := "query_example" // string | Prometheus expression query string
	start := "start_example" // string | <rfc3339 | unix_timestamp>: Start timestamp, inclusive
	end := "end_example" // string | <rfc3339 | unix_timestamp>: End timestamp, inclusive
	step := "step_example" // string | Query resolution step width in duration format or float number of seconds (optional)
	timeout := "timeout_example" // string | Evaluation timeout (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.PrometheusRangeQuery(context.Background(), orgId).Query(query).Start(start).End(end).Step(step).Timeout(timeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.PrometheusRangeQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrometheusRangeQuery`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.PrometheusRangeQuery`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiPrometheusRangeQueryRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**query** | **string** | Prometheus expression query string |
**start** | **string** | &lt;rfc3339 | unix_timestamp&gt;: Start timestamp, inclusive |
**end** | **string** | &lt;rfc3339 | unix_timestamp&gt;: End timestamp, inclusive |
**step** | **string** | Query resolution step width in duration format or float number of seconds |
**timeout** | **string** | Evaluation timeout |

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

## PrometheusRemoteWrite

> IngestionResponse PrometheusRemoteWrite(ctx, orgId).Body(body).Execute()

prometheus remote-write endpoint for metrics

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
	body := "body_example" // string | prometheus WriteRequest

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.PrometheusRemoteWrite(context.Background(), orgId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.PrometheusRemoteWrite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrometheusRemoteWrite`: IngestionResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.PrometheusRemoteWrite`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiPrometheusRemoteWriteRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**body** | **string** | prometheus WriteRequest |

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

## PrometheusSeries

> HttpResponse PrometheusSeries(ctx, orgId).Match(match).Start(start).End(end).Execute()

prometheus finding series by label matchers

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
	match := "match_example" // string | <series_selector>: Series selector argument that selects the series to return
	start := "start_example" // string | <rfc3339 | unix_timestamp>: Start timestamp (optional)
	end := "end_example" // string | <rfc3339 | unix_timestamp>: End timestamp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.PrometheusSeries(context.Background(), orgId).Match(match).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.PrometheusSeries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrometheusSeries`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.PrometheusSeries`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiPrometheusSeriesRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**match** | **string** | &lt;series_selector&gt;: Series selector argument that selects the series to return |
**start** | **string** | &lt;rfc3339 | unix_timestamp&gt;: Start timestamp |
**end** | **string** | &lt;rfc3339 | unix_timestamp&gt;: End timestamp |

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
