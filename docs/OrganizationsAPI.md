# \OrganizationsAPI

All URIs are relative to _http://localhost_

| Method                                                                                               | HTTP request                    | Description                             |
| ---------------------------------------------------------------------------------------------------- | ------------------------------- | --------------------------------------- |
| [**CreateOrganizationUserRumIngestToken**](OrganizationsAPI.md#CreateOrganizationUserRumIngestToken) | **Post** /api/{org_id}/rumtoken | CreateRumIngestToken                    |
| [**GetOrganizationSummary**](OrganizationsAPI.md#GetOrganizationSummary)                             | **Get** /api/{org_id}/summary   | GetOrganizationSummary                  |
| [**GetOrganizationUserIngestToken**](OrganizationsAPI.md#GetOrganizationUserIngestToken)             | **Get** /api/{org_id}/passcode  | GetIngestToken                          |
| [**GetOrganizationUserRumIngestToken**](OrganizationsAPI.md#GetOrganizationUserRumIngestToken)       | **Get** /api/{org_id}/rumtoken  | GetRumIngestToken                       |
| [**GetUserOrganizations**](OrganizationsAPI.md#GetUserOrganizations)                                 | **Get** /api/organizations      | GetOrganizations                        |
| [**OrganizationSettingCreate**](OrganizationsAPI.md#OrganizationSettingCreate)                       | **Post** /api/{org_id}/settings | Organization specific settings          |
| [**OrganizationSettingGet**](OrganizationsAPI.md#OrganizationSettingGet)                             | **Get** /api/{org_id}/settings  | Retrieve organization specific settings |
| [**UpdateOrganizationUserIngestToken**](OrganizationsAPI.md#UpdateOrganizationUserIngestToken)       | **Put** /api/{org_id}/passcode  | UpdateIngestToken                       |
| [**UpdateOrganizationUserRumIngestToken**](OrganizationsAPI.md#UpdateOrganizationUserRumIngestToken) | **Put** /api/{org_id}/rumtoken  | UpdateRumIngestToken                    |

## CreateOrganizationUserRumIngestToken

> RumIngestionResponse CreateOrganizationUserRumIngestToken(ctx, orgId).Execute()

CreateRumIngestToken

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.CreateOrganizationUserRumIngestToken(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.CreateOrganizationUserRumIngestToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationUserRumIngestToken`: RumIngestionResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.CreateOrganizationUserRumIngestToken`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationUserRumIngestTokenRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**RumIngestionResponse**](RumIngestionResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetOrganizationSummary

> OrgSummary GetOrganizationSummary(ctx, orgId).Execute()

GetOrganizationSummary

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganizationSummary(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSummary`: OrgSummary
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizationSummary`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSummaryRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**OrgSummary**](OrgSummary.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetOrganizationUserIngestToken

> PasscodeResponse GetOrganizationUserIngestToken(ctx, orgId).Execute()

GetIngestToken

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganizationUserIngestToken(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationUserIngestToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationUserIngestToken`: PasscodeResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizationUserIngestToken`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationUserIngestTokenRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**PasscodeResponse**](PasscodeResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetOrganizationUserRumIngestToken

> RumIngestionResponse GetOrganizationUserRumIngestToken(ctx, orgId).Execute()

GetRumIngestToken

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganizationUserRumIngestToken(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationUserRumIngestToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationUserRumIngestToken`: RumIngestionResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizationUserRumIngestToken`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationUserRumIngestTokenRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**RumIngestionResponse**](RumIngestionResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserOrganizations

> OrganizationResponse GetUserOrganizations(ctx).Execute()

GetOrganizations

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetUserOrganizations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetUserOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserOrganizations`: OrganizationResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetUserOrganizations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserOrganizationsRequest struct via the builder pattern

### Return type

[**OrganizationResponse**](OrganizationResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## OrganizationSettingCreate

> HttpResponse OrganizationSettingCreate(ctx, orgId).Body(body).Execute()

Organization specific settings

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
	body := OrganizationSettingPayload(987) // OrganizationSettingPayload | Organization settings

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.OrganizationSettingCreate(context.Background(), orgId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationSettingCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationSettingCreate`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrganizationSettingCreate`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationSettingCreateRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**body** | **OrganizationSettingPayload** | Organization settings |

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

## OrganizationSettingGet

> HttpResponse OrganizationSettingGet(ctx, orgId).Execute()

Retrieve organization specific settings

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.OrganizationSettingGet(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationSettingGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationSettingGet`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrganizationSettingGet`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationSettingGetRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

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

## UpdateOrganizationUserIngestToken

> PasscodeResponse UpdateOrganizationUserIngestToken(ctx, orgId).Execute()

UpdateIngestToken

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.UpdateOrganizationUserIngestToken(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.UpdateOrganizationUserIngestToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationUserIngestToken`: PasscodeResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.UpdateOrganizationUserIngestToken`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationUserIngestTokenRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**PasscodeResponse**](PasscodeResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateOrganizationUserRumIngestToken

> RumIngestionResponse UpdateOrganizationUserRumIngestToken(ctx, orgId).Execute()

UpdateRumIngestToken

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.UpdateOrganizationUserRumIngestToken(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.UpdateOrganizationUserRumIngestToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationUserRumIngestToken`: RumIngestionResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.UpdateOrganizationUserRumIngestToken`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationUserRumIngestTokenRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**RumIngestionResponse**](RumIngestionResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
