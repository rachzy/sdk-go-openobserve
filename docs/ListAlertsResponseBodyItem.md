# ListAlertsResponseBodyItem

## Properties

| Name                | Type                                    | Description | Notes      |
| ------------------- | --------------------------------------- | ----------- | ---------- |
| **AlertId**         | [**string**](string.md)                 |             |
| **Condition**       | [**QueryCondition**](QueryCondition.md) |             |
| **Description**     | Pointer to **NullableString**           |             | [optional] |
| **Enabled**         | **bool**                                |             |
| **FolderId**        | **string**                              |             |
| **FolderName**      | **string**                              |             |
| **LastSatisfiedAt** | Pointer to **NullableInt64**            |             | [optional] |
| **LastTriggeredAt** | Pointer to **NullableInt64**            |             | [optional] |
| **Name**            | **string**                              |             |
| **Owner**           | Pointer to **NullableString**           |             | [optional] |

## Methods

### NewListAlertsResponseBodyItem

`func NewListAlertsResponseBodyItem(alertId string, condition QueryCondition, enabled bool, folderId string, folderName string, name string, ) *ListAlertsResponseBodyItem`

NewListAlertsResponseBodyItem instantiates a new ListAlertsResponseBodyItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAlertsResponseBodyItemWithDefaults

`func NewListAlertsResponseBodyItemWithDefaults() *ListAlertsResponseBodyItem`

NewListAlertsResponseBodyItemWithDefaults instantiates a new ListAlertsResponseBodyItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertId

`func (o *ListAlertsResponseBodyItem) GetAlertId() string`

GetAlertId returns the AlertId field if non-nil, zero value otherwise.

### GetAlertIdOk

`func (o *ListAlertsResponseBodyItem) GetAlertIdOk() (*string, bool)`

GetAlertIdOk returns a tuple with the AlertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertId

`func (o *ListAlertsResponseBodyItem) SetAlertId(v string)`

SetAlertId sets AlertId field to given value.

### GetCondition

`func (o *ListAlertsResponseBodyItem) GetCondition() QueryCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ListAlertsResponseBodyItem) GetConditionOk() (*QueryCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ListAlertsResponseBodyItem) SetCondition(v QueryCondition)`

SetCondition sets Condition field to given value.

### GetDescription

`func (o *ListAlertsResponseBodyItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListAlertsResponseBodyItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListAlertsResponseBodyItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListAlertsResponseBodyItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListAlertsResponseBodyItem) SetDescriptionNil(b bool)`

SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription

`func (o *ListAlertsResponseBodyItem) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

### GetEnabled

`func (o *ListAlertsResponseBodyItem) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListAlertsResponseBodyItem) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListAlertsResponseBodyItem) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### GetFolderId

`func (o *ListAlertsResponseBodyItem) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *ListAlertsResponseBodyItem) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *ListAlertsResponseBodyItem) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### GetFolderName

`func (o *ListAlertsResponseBodyItem) GetFolderName() string`

GetFolderName returns the FolderName field if non-nil, zero value otherwise.

### GetFolderNameOk

`func (o *ListAlertsResponseBodyItem) GetFolderNameOk() (*string, bool)`

GetFolderNameOk returns a tuple with the FolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderName

`func (o *ListAlertsResponseBodyItem) SetFolderName(v string)`

SetFolderName sets FolderName field to given value.

### GetLastSatisfiedAt

`func (o *ListAlertsResponseBodyItem) GetLastSatisfiedAt() int64`

GetLastSatisfiedAt returns the LastSatisfiedAt field if non-nil, zero value otherwise.

### GetLastSatisfiedAtOk

`func (o *ListAlertsResponseBodyItem) GetLastSatisfiedAtOk() (*int64, bool)`

GetLastSatisfiedAtOk returns a tuple with the LastSatisfiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSatisfiedAt

`func (o *ListAlertsResponseBodyItem) SetLastSatisfiedAt(v int64)`

SetLastSatisfiedAt sets LastSatisfiedAt field to given value.

### HasLastSatisfiedAt

`func (o *ListAlertsResponseBodyItem) HasLastSatisfiedAt() bool`

HasLastSatisfiedAt returns a boolean if a field has been set.

### SetLastSatisfiedAtNil

`func (o *ListAlertsResponseBodyItem) SetLastSatisfiedAtNil(b bool)`

SetLastSatisfiedAtNil sets the value for LastSatisfiedAt to be an explicit nil

### UnsetLastSatisfiedAt

`func (o *ListAlertsResponseBodyItem) UnsetLastSatisfiedAt()`

UnsetLastSatisfiedAt ensures that no value is present for LastSatisfiedAt, not even an explicit nil

### GetLastTriggeredAt

`func (o *ListAlertsResponseBodyItem) GetLastTriggeredAt() int64`

GetLastTriggeredAt returns the LastTriggeredAt field if non-nil, zero value otherwise.

### GetLastTriggeredAtOk

`func (o *ListAlertsResponseBodyItem) GetLastTriggeredAtOk() (*int64, bool)`

GetLastTriggeredAtOk returns a tuple with the LastTriggeredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTriggeredAt

`func (o *ListAlertsResponseBodyItem) SetLastTriggeredAt(v int64)`

SetLastTriggeredAt sets LastTriggeredAt field to given value.

### HasLastTriggeredAt

`func (o *ListAlertsResponseBodyItem) HasLastTriggeredAt() bool`

HasLastTriggeredAt returns a boolean if a field has been set.

### SetLastTriggeredAtNil

`func (o *ListAlertsResponseBodyItem) SetLastTriggeredAtNil(b bool)`

SetLastTriggeredAtNil sets the value for LastTriggeredAt to be an explicit nil

### UnsetLastTriggeredAt

`func (o *ListAlertsResponseBodyItem) UnsetLastTriggeredAt()`

UnsetLastTriggeredAt ensures that no value is present for LastTriggeredAt, not even an explicit nil

### GetName

`func (o *ListAlertsResponseBodyItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListAlertsResponseBodyItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListAlertsResponseBodyItem) SetName(v string)`

SetName sets Name field to given value.

### GetOwner

`func (o *ListAlertsResponseBodyItem) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListAlertsResponseBodyItem) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListAlertsResponseBodyItem) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListAlertsResponseBodyItem) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *ListAlertsResponseBodyItem) SetOwnerNil(b bool)`

SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner

`func (o *ListAlertsResponseBodyItem) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
