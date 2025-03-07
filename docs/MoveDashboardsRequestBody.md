# MoveDashboardsRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DashboardIds** | **[]string** | IDs of the dashboards to move. | 
**DstFolderId** | **string** | Indicates the folder to which dashboard should be moved. | 

## Methods

### NewMoveDashboardsRequestBody

`func NewMoveDashboardsRequestBody(dashboardIds []string, dstFolderId string, ) *MoveDashboardsRequestBody`

NewMoveDashboardsRequestBody instantiates a new MoveDashboardsRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveDashboardsRequestBodyWithDefaults

`func NewMoveDashboardsRequestBodyWithDefaults() *MoveDashboardsRequestBody`

NewMoveDashboardsRequestBodyWithDefaults instantiates a new MoveDashboardsRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDashboardIds

`func (o *MoveDashboardsRequestBody) GetDashboardIds() []string`

GetDashboardIds returns the DashboardIds field if non-nil, zero value otherwise.

### GetDashboardIdsOk

`func (o *MoveDashboardsRequestBody) GetDashboardIdsOk() (*[]string, bool)`

GetDashboardIdsOk returns a tuple with the DashboardIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardIds

`func (o *MoveDashboardsRequestBody) SetDashboardIds(v []string)`

SetDashboardIds sets DashboardIds field to given value.


### GetDstFolderId

`func (o *MoveDashboardsRequestBody) GetDstFolderId() string`

GetDstFolderId returns the DstFolderId field if non-nil, zero value otherwise.

### GetDstFolderIdOk

`func (o *MoveDashboardsRequestBody) GetDstFolderIdOk() (*string, bool)`

GetDstFolderIdOk returns a tuple with the DstFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstFolderId

`func (o *MoveDashboardsRequestBody) SetDstFolderId(v string)`

SetDstFolderId sets DstFolderId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


