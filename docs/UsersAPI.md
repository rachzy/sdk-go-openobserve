# \UsersAPI

All URIs are relative to _http://localhost_

| Method                                                 | HTTP request                              | Description                |
| ------------------------------------------------------ | ----------------------------------------- | -------------------------- |
| [**AddUserToOrg**](UsersAPI.md#AddUserToOrg)           | **Post** /api/{org_id}/users/{email_id}   | AddUserToOrganization      |
| [**RemoveUserFromOrg**](UsersAPI.md#RemoveUserFromOrg) | **Delete** /api/{org_id}/users/{email_id} | RemoveUserFromOrganization |
| [**UserList**](UsersAPI.md#UserList)                   | **Get** /api/{org_id}/users               | ListUsers                  |
| [**UserSave**](UsersAPI.md#UserSave)                   | **Post** /api/{org_id}/users              | CreateUser                 |
| [**UserUpdate**](UsersAPI.md#UserUpdate)               | **Put** /api/{org_id}/users/{email_id}    | UpdateUser                 |

## AddUserToOrg

> HttpResponse AddUserToOrg(ctx, orgId, emailId).UserOrgRole(userOrgRole).Execute()

AddUserToOrganization

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
	emailId := "emailId_example" // string | User's email id
	userOrgRole := *openobserve.NewUserOrgRole(openobserve.UserRole("admin")) // UserOrgRole | User role

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AddUserToOrg(context.Background(), orgId, emailId).UserOrgRole(userOrgRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AddUserToOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddUserToOrg`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AddUserToOrg`: %v\n", resp)
}
```

### Path Parameters

| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**   | **string**          | Organization name                                                           |
| **emailId** | **string**          | User&#39;s email id                                                         |

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserToOrgRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**userOrgRole** | [**UserOrgRole**](UserOrgRole.md) | User role |

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

## RemoveUserFromOrg

> HttpResponse RemoveUserFromOrg(ctx, orgId, emailId).Execute()

RemoveUserFromOrganization

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
	emailId := "emailId_example" // string | User name

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.RemoveUserFromOrg(context.Background(), orgId, emailId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.RemoveUserFromOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveUserFromOrg`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.RemoveUserFromOrg`: %v\n", resp)
}
```

### Path Parameters

| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**   | **string**          | Organization name                                                           |
| **emailId** | **string**          | User name                                                                   |

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserFromOrgRequest struct via the builder pattern

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

## UserList

> UserList UserList(ctx, orgId).Execute()

ListUsers

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

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UserList(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UserList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserList`: UserList
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UserList`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiUserListRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**UserList**](UserList.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UserSave

> HttpResponse UserSave(ctx, orgId).UserRequest(userRequest).Execute()

CreateUser

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
	userRequest := *openobserve.NewUserRequest("Email_example", "Password_example") // UserRequest | User data

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UserSave(context.Background(), orgId).UserRequest(userRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UserSave``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserSave`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UserSave`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiUserSaveRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**userRequest** | [**UserRequest**](UserRequest.md) | User data |

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

## UserUpdate

> HttpResponse UserUpdate(ctx, orgId, emailId).UpdateUser(updateUser).Execute()

UpdateUser

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
	emailId := "emailId_example" // string | User's email id
	updateUser := *openobserve.NewUpdateUser() // UpdateUser | User data

	configuration := openobserve.NewConfiguration()
	apiClient := openobserve.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UserUpdate(context.Background(), orgId, emailId).UpdateUser(updateUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UserUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserUpdate`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UserUpdate`: %v\n", resp)
}
```

### Path Parameters

| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**   | **string**          | Organization name                                                           |
| **emailId** | **string**          | User&#39;s email id                                                         |

### Other Parameters

Other parameters are passed through a pointer to a apiUserUpdateRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**updateUser** | [**UpdateUser**](UpdateUser.md) | User data |

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
