# MoveAlertsRequestBody

## Properties

| Name            | Type                      | Description                                           | Notes |
| --------------- | ------------------------- | ----------------------------------------------------- | ----- |
| **AlertIds**    | [**[]string**](string.md) | IDs of the alerts to move.                            |
| **DstFolderId** | **string**                | Indicates the folder to which alerts should be moved. |

## Methods

### NewMoveAlertsRequestBody

`func NewMoveAlertsRequestBody(alertIds []string, dstFolderId string, ) *MoveAlertsRequestBody`

NewMoveAlertsRequestBody instantiates a new MoveAlertsRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveAlertsRequestBodyWithDefaults

`func NewMoveAlertsRequestBodyWithDefaults() *MoveAlertsRequestBody`

NewMoveAlertsRequestBodyWithDefaults instantiates a new MoveAlertsRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertIds

`func (o *MoveAlertsRequestBody) GetAlertIds() []string`

GetAlertIds returns the AlertIds field if non-nil, zero value otherwise.

### GetAlertIdsOk

`func (o *MoveAlertsRequestBody) GetAlertIdsOk() (*[]string, bool)`

GetAlertIdsOk returns a tuple with the AlertIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertIds

`func (o *MoveAlertsRequestBody) SetAlertIds(v []string)`

SetAlertIds sets AlertIds field to given value.

### GetDstFolderId

`func (o *MoveAlertsRequestBody) GetDstFolderId() string`

GetDstFolderId returns the DstFolderId field if non-nil, zero value otherwise.

### GetDstFolderIdOk

`func (o *MoveAlertsRequestBody) GetDstFolderIdOk() (*string, bool)`

GetDstFolderIdOk returns a tuple with the DstFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstFolderId

`func (o *MoveAlertsRequestBody) SetDstFolderId(v string)`

SetDstFolderId sets DstFolderId field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
