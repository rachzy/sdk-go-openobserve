# \FoldersAPI

All URIs are relative to _http://localhost_

| Method                                                   | HTTP request                                                      | Description     |
| -------------------------------------------------------- | ----------------------------------------------------------------- | --------------- |
| [**CreateFolder**](FoldersAPI.md#CreateFolder)           | **Post** /api/v2/{org_id}/folders/{folder_type}                   | CreateFolder    |
| [**CreateFolder_0**](FoldersAPI.md#CreateFolder_0)       | **Post** /api/{org_id}/folders                                    | CreateFolder    |
| [**DeleteFolder**](FoldersAPI.md#DeleteFolder)           | **Delete** /api/v2/{org_id}/folders/{folder_type}/{folder_id}     | DeleteFolder    |
| [**DeleteFolder_0**](FoldersAPI.md#DeleteFolder_0)       | **Delete** /api/{org_id}/folders/{folder_id}                      | DeleteFolder    |
| [**GetFolder**](FoldersAPI.md#GetFolder)                 | **Get** /api/v2/{org_id}/folders/{folder_type}/{folder_id}        | GetFolder       |
| [**GetFolderByName**](FoldersAPI.md#GetFolderByName)     | **Get** /api/v2/{org_id}/folders/{folder_type}/name/{folder_name} | GetFolderByName |
| [**GetFolderByName_0**](FoldersAPI.md#GetFolderByName_0) | **Get** /api/{org_id}/folders/name/{folder_name}                  | GetFolderByName |
| [**GetFolder_0**](FoldersAPI.md#GetFolder_0)             | **Get** /api/{org_id}/folders/{folder_id}                         | GetFolder       |
| [**ListFolders**](FoldersAPI.md#ListFolders)             | **Get** /api/v2/{org_id}/folders/{folder_type}                    | ListFolders     |
| [**ListFolders_0**](FoldersAPI.md#ListFolders_0)         | **Get** /api/{org_id}/folders                                     | ListFolders     |
| [**UpdateFolder**](FoldersAPI.md#UpdateFolder)           | **Put** /api/v2/{org_id}/folders/{folder_type}/{folder_id}        | UpdateFolder    |
| [**UpdateFolder_0**](FoldersAPI.md#UpdateFolder_0)       | **Put** /api/{org_id}/folders/{folder_id}                         | UpdateFolder    |

## CreateFolder

> Folder CreateFolder(ctx, orgId, folderType).CreateFolderRequestBody(createFolderRequestBody).Execute()

CreateFolder

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
	folderType := openapiclient.FolderType("dashboards") // FolderType | Type of data the folder can contain
	createFolderRequestBody := *openapiclient.NewCreateFolderRequestBody("Description_example", "Name_example") // CreateFolderRequestBody | Folder details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.CreateFolder(context.Background(), orgId, folderType).CreateFolderRequestBody(createFolderRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.CreateFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFolder`: Folder
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.CreateFolder`: %v\n", resp)
}
```

### Path Parameters

| Name           | Type                  | Description                                                                 | Notes |
| -------------- | --------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**        | **context.Context**   | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**      | **string**            | Organization name                                                           |
| **folderType** | [**FolderType**](.md) | Type of data the folder can contain                                         |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFolderRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**createFolderRequestBody** | [**CreateFolderRequestBody**](CreateFolderRequestBody.md) | Folder details |

### Return type

[**Folder**](Folder.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateFolder_0

> Folder CreateFolder_0(ctx, orgId).CreateFolderRequestBody(createFolderRequestBody).Execute()

CreateFolder

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
	createFolderRequestBody := *openapiclient.NewCreateFolderRequestBody("Description_example", "Name_example") // CreateFolderRequestBody | Folder details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.CreateFolder_0(context.Background(), orgId).CreateFolderRequestBody(createFolderRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.CreateFolder_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFolder_0`: Folder
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.CreateFolder_0`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFolder_1Request struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**createFolderRequestBody** | [**CreateFolderRequestBody**](CreateFolderRequestBody.md) | Folder details |

### Return type

[**Folder**](Folder.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteFolder

> HttpResponse DeleteFolder(ctx, orgId, folderType, folderId).Execute()

DeleteFolder

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
	folderType := openapiclient.FolderType("dashboards") // FolderType | Type of data the folder can contain
	folderId := "folderId_example" // string | Folder ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.DeleteFolder(context.Background(), orgId, folderType, folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.DeleteFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFolder`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.DeleteFolder`: %v\n", resp)
}
```

### Path Parameters

| Name           | Type                  | Description                                                                 | Notes |
| -------------- | --------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**        | **context.Context**   | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**      | **string**            | Organization name                                                           |
| **folderType** | [**FolderType**](.md) | Type of data the folder can contain                                         |
| **folderId**   | **string**            | Folder ID                                                                   |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFolderRequest struct via the builder pattern

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

## DeleteFolder_0

> HttpResponse DeleteFolder_0(ctx, orgId, folderId).Execute()

DeleteFolder

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
	folderId := "folderId_example" // string | Folder ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.DeleteFolder_0(context.Background(), orgId, folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.DeleteFolder_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFolder_0`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.DeleteFolder_0`: %v\n", resp)
}
```

### Path Parameters

| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**    | **string**          | Organization name                                                           |
| **folderId** | **string**          | Folder ID                                                                   |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFolder_2Request struct via the builder pattern

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

## GetFolder

> Folder GetFolder(ctx, orgId, folderType, folderId).Execute()

GetFolder

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
	folderType := openapiclient.FolderType("dashboards") // FolderType | Type of data the folder can contain
	folderId := "folderId_example" // string | Folder ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.GetFolder(context.Background(), orgId, folderType, folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.GetFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFolder`: Folder
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.GetFolder`: %v\n", resp)
}
```

### Path Parameters

| Name           | Type                  | Description                                                                 | Notes |
| -------------- | --------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**        | **context.Context**   | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**      | **string**            | Organization name                                                           |
| **folderType** | [**FolderType**](.md) | Type of data the folder can contain                                         |
| **folderId**   | **string**            | Folder ID                                                                   |

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**Folder**](Folder.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetFolderByName

> Folder GetFolderByName(ctx, orgId, folderType, folderName).Execute()

GetFolderByName

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
	folderType := openapiclient.FolderType("dashboards") // FolderType | Type of data the folder can contain
	folderName := "folderName_example" // string | Folder Name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.GetFolderByName(context.Background(), orgId, folderType, folderName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.GetFolderByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFolderByName`: Folder
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.GetFolderByName`: %v\n", resp)
}
```

### Path Parameters

| Name           | Type                  | Description                                                                 | Notes |
| -------------- | --------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**        | **context.Context**   | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**      | **string**            | Organization name                                                           |
| **folderType** | [**FolderType**](.md) | Type of data the folder can contain                                         |
| **folderName** | **string**            | Folder Name                                                                 |

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderByNameRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**Folder**](Folder.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetFolderByName_0

> Folder GetFolderByName_0(ctx, orgId, folderName).Execute()

GetFolderByName

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
	folderName := "folderName_example" // string | Folder Name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.GetFolderByName_0(context.Background(), orgId, folderName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.GetFolderByName_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFolderByName_0`: Folder
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.GetFolderByName_0`: %v\n", resp)
}
```

### Path Parameters

| Name           | Type                | Description                                                                 | Notes |
| -------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**        | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**      | **string**          | Organization name                                                           |
| **folderName** | **string**          | Folder Name                                                                 |

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderByName_3Request struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**Folder**](Folder.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetFolder_0

> Folder GetFolder_0(ctx, orgId, folderId).Execute()

GetFolder

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
	folderId := "folderId_example" // string | Folder ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.GetFolder_0(context.Background(), orgId, folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.GetFolder_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFolder_0`: Folder
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.GetFolder_0`: %v\n", resp)
}
```

### Path Parameters

| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**    | **string**          | Organization name                                                           |
| **folderId** | **string**          | Folder ID                                                                   |

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolder_4Request struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**Folder**](Folder.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListFolders

> ListFoldersResponseBody ListFolders(ctx, orgId, folderType).Execute()

ListFolders

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
	folderType := openapiclient.FolderType("dashboards") // FolderType | Type of data the folder can contain

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.ListFolders(context.Background(), orgId, folderType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.ListFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFolders`: ListFoldersResponseBody
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.ListFolders`: %v\n", resp)
}
```

### Path Parameters

| Name           | Type                  | Description                                                                 | Notes |
| -------------- | --------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**        | **context.Context**   | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**      | **string**            | Organization name                                                           |
| **folderType** | [**FolderType**](.md) | Type of data the folder can contain                                         |

### Other Parameters

Other parameters are passed through a pointer to a apiListFoldersRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**ListFoldersResponseBody**](ListFoldersResponseBody.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListFolders_0

> ListFoldersResponseBody ListFolders_0(ctx, orgId).Execute()

ListFolders

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
	resp, r, err := apiClient.FoldersAPI.ListFolders_0(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.ListFolders_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFolders_0`: ListFoldersResponseBody
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.ListFolders_0`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId** | **string**          | Organization name                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiListFolders_5Request struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**ListFoldersResponseBody**](ListFoldersResponseBody.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateFolder

> HttpResponse UpdateFolder(ctx, orgId, folderType, folderId).Body(body).Execute()

UpdateFolder

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
	folderType := openapiclient.FolderType("dashboards") // FolderType | Type of data the folder can contain
	folderId := "folderId_example" // string | Folder name
	body := Folder({"description":"Traffic patterns and network performance of the infrastructure","title":"Infra"}) // Folder | Folder details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.UpdateFolder(context.Background(), orgId, folderType, folderId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.UpdateFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFolder`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.UpdateFolder`: %v\n", resp)
}
```

### Path Parameters

| Name           | Type                  | Description                                                                 | Notes |
| -------------- | --------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**        | **context.Context**   | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**      | **string**            | Organization name                                                           |
| **folderType** | [**FolderType**](.md) | Type of data the folder can contain                                         |
| **folderId**   | **string**            | Folder name                                                                 |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFolderRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**body** | **Folder** | Folder details |

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

## UpdateFolder_0

> HttpResponse UpdateFolder_0(ctx, orgId, folderId).Body(body).Execute()

UpdateFolder

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
	folderId := "folderId_example" // string | Folder name
	body := Folder({"description":"Traffic patterns and network performance of the infrastructure","title":"Infra"}) // Folder | Folder details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.UpdateFolder_0(context.Background(), orgId, folderId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.UpdateFolder_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFolder_0`: HttpResponse
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.UpdateFolder_0`: %v\n", resp)
}
```

### Path Parameters

| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **orgId**    | **string**          | Organization name                                                           |
| **folderId** | **string**          | Folder name                                                                 |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFolder_6Request struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**body** | **Folder** | Folder details |

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
