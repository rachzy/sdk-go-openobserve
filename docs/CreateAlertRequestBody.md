# CreateAlertRequestBody

## Properties

| Name                  | Type                                                   | Description                                                                                                                        | Notes                 |
| --------------------- | ------------------------------------------------------ | ---------------------------------------------------------------------------------------------------------------------------------- | --------------------- |
| **ContextAttributes** | Pointer to **map[string]string**                       |                                                                                                                                    | [optional]            |
| **Description**       | Pointer to **string**                                  |                                                                                                                                    | [optional]            |
| **Destinations**      | **[]string**                                           |                                                                                                                                    |
| **Enabled**           | Pointer to **bool**                                    |                                                                                                                                    | [optional]            |
| **Id**                | Pointer to [**Nullablestring**](string.md)             |                                                                                                                                    | [optional]            |
| **IsRealTime**        | Pointer to **bool**                                    |                                                                                                                                    | [optional]            |
| **LastEditedBy**      | Pointer to **NullableString**                          |                                                                                                                                    | [optional] [readonly] |
| **LastSatisfiedAt**   | Pointer to **NullableInt64**                           | Time when alert was last satisfied. Unix timestamp.                                                                                | [optional] [readonly] |
| **LastTriggeredAt**   | Pointer to **NullableInt64**                           | Time when alert was last triggered. Unix timestamp.                                                                                | [optional] [readonly] |
| **Name**              | Pointer to **string**                                  |                                                                                                                                    | [optional]            |
| **OrgId**             | Pointer to **string**                                  |                                                                                                                                    | [optional]            |
| **Owner**             | Pointer to **NullableString**                          |                                                                                                                                    | [optional]            |
| **QueryCondition**    | Pointer to [**QueryCondition**](QueryCondition.md)     |                                                                                                                                    | [optional]            |
| **RowTemplate**       | Pointer to **string**                                  |                                                                                                                                    | [optional]            |
| **StreamName**        | Pointer to **string**                                  |                                                                                                                                    | [optional]            |
| **StreamType**        | Pointer to [**StreamType**](StreamType.md)             |                                                                                                                                    | [optional]            |
| **TriggerCondition**  | Pointer to [**TriggerCondition**](TriggerCondition.md) |                                                                                                                                    | [optional]            |
| **TzOffset**          | Pointer to **int32**                                   | Timezone offset in minutes. Negative seconds means the western hemisphere                                                          | [optional]            |
| **UpdatedAt**         | Pointer to **NullableInt64**                           | Time when alert was last updated. Unix timestamp.                                                                                  | [optional] [readonly] |
| **FolderId**          | Pointer to **NullableString**                          | Optional folder ID indicating the folder in which to create the alert. If omitted the alert will be created in the default folder. | [optional]            |

## Methods

### NewCreateAlertRequestBody

`func NewCreateAlertRequestBody(destinations []string, ) *CreateAlertRequestBody`

NewCreateAlertRequestBody instantiates a new CreateAlertRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAlertRequestBodyWithDefaults

`func NewCreateAlertRequestBodyWithDefaults() *CreateAlertRequestBody`

NewCreateAlertRequestBodyWithDefaults instantiates a new CreateAlertRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextAttributes

`func (o *CreateAlertRequestBody) GetContextAttributes() map[string]string`

GetContextAttributes returns the ContextAttributes field if non-nil, zero value otherwise.

### GetContextAttributesOk

`func (o *CreateAlertRequestBody) GetContextAttributesOk() (*map[string]string, bool)`

GetContextAttributesOk returns a tuple with the ContextAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextAttributes

`func (o *CreateAlertRequestBody) SetContextAttributes(v map[string]string)`

SetContextAttributes sets ContextAttributes field to given value.

### HasContextAttributes

`func (o *CreateAlertRequestBody) HasContextAttributes() bool`

HasContextAttributes returns a boolean if a field has been set.

### SetContextAttributesNil

`func (o *CreateAlertRequestBody) SetContextAttributesNil(b bool)`

SetContextAttributesNil sets the value for ContextAttributes to be an explicit nil

### UnsetContextAttributes

`func (o *CreateAlertRequestBody) UnsetContextAttributes()`

UnsetContextAttributes ensures that no value is present for ContextAttributes, not even an explicit nil

### GetDescription

`func (o *CreateAlertRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAlertRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAlertRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAlertRequestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDestinations

`func (o *CreateAlertRequestBody) GetDestinations() []string`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *CreateAlertRequestBody) GetDestinationsOk() (*[]string, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *CreateAlertRequestBody) SetDestinations(v []string)`

SetDestinations sets Destinations field to given value.

### GetEnabled

`func (o *CreateAlertRequestBody) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateAlertRequestBody) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateAlertRequestBody) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateAlertRequestBody) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *CreateAlertRequestBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateAlertRequestBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateAlertRequestBody) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateAlertRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CreateAlertRequestBody) SetIdNil(b bool)`

SetIdNil sets the value for Id to be an explicit nil

### UnsetId

`func (o *CreateAlertRequestBody) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil

### GetIsRealTime

`func (o *CreateAlertRequestBody) GetIsRealTime() bool`

GetIsRealTime returns the IsRealTime field if non-nil, zero value otherwise.

### GetIsRealTimeOk

`func (o *CreateAlertRequestBody) GetIsRealTimeOk() (*bool, bool)`

GetIsRealTimeOk returns a tuple with the IsRealTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRealTime

`func (o *CreateAlertRequestBody) SetIsRealTime(v bool)`

SetIsRealTime sets IsRealTime field to given value.

### HasIsRealTime

`func (o *CreateAlertRequestBody) HasIsRealTime() bool`

HasIsRealTime returns a boolean if a field has been set.

### GetLastEditedBy

`func (o *CreateAlertRequestBody) GetLastEditedBy() string`

GetLastEditedBy returns the LastEditedBy field if non-nil, zero value otherwise.

### GetLastEditedByOk

`func (o *CreateAlertRequestBody) GetLastEditedByOk() (*string, bool)`

GetLastEditedByOk returns a tuple with the LastEditedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditedBy

`func (o *CreateAlertRequestBody) SetLastEditedBy(v string)`

SetLastEditedBy sets LastEditedBy field to given value.

### HasLastEditedBy

`func (o *CreateAlertRequestBody) HasLastEditedBy() bool`

HasLastEditedBy returns a boolean if a field has been set.

### SetLastEditedByNil

`func (o *CreateAlertRequestBody) SetLastEditedByNil(b bool)`

SetLastEditedByNil sets the value for LastEditedBy to be an explicit nil

### UnsetLastEditedBy

`func (o *CreateAlertRequestBody) UnsetLastEditedBy()`

UnsetLastEditedBy ensures that no value is present for LastEditedBy, not even an explicit nil

### GetLastSatisfiedAt

`func (o *CreateAlertRequestBody) GetLastSatisfiedAt() int64`

GetLastSatisfiedAt returns the LastSatisfiedAt field if non-nil, zero value otherwise.

### GetLastSatisfiedAtOk

`func (o *CreateAlertRequestBody) GetLastSatisfiedAtOk() (*int64, bool)`

GetLastSatisfiedAtOk returns a tuple with the LastSatisfiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSatisfiedAt

`func (o *CreateAlertRequestBody) SetLastSatisfiedAt(v int64)`

SetLastSatisfiedAt sets LastSatisfiedAt field to given value.

### HasLastSatisfiedAt

`func (o *CreateAlertRequestBody) HasLastSatisfiedAt() bool`

HasLastSatisfiedAt returns a boolean if a field has been set.

### SetLastSatisfiedAtNil

`func (o *CreateAlertRequestBody) SetLastSatisfiedAtNil(b bool)`

SetLastSatisfiedAtNil sets the value for LastSatisfiedAt to be an explicit nil

### UnsetLastSatisfiedAt

`func (o *CreateAlertRequestBody) UnsetLastSatisfiedAt()`

UnsetLastSatisfiedAt ensures that no value is present for LastSatisfiedAt, not even an explicit nil

### GetLastTriggeredAt

`func (o *CreateAlertRequestBody) GetLastTriggeredAt() int64`

GetLastTriggeredAt returns the LastTriggeredAt field if non-nil, zero value otherwise.

### GetLastTriggeredAtOk

`func (o *CreateAlertRequestBody) GetLastTriggeredAtOk() (*int64, bool)`

GetLastTriggeredAtOk returns a tuple with the LastTriggeredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTriggeredAt

`func (o *CreateAlertRequestBody) SetLastTriggeredAt(v int64)`

SetLastTriggeredAt sets LastTriggeredAt field to given value.

### HasLastTriggeredAt

`func (o *CreateAlertRequestBody) HasLastTriggeredAt() bool`

HasLastTriggeredAt returns a boolean if a field has been set.

### SetLastTriggeredAtNil

`func (o *CreateAlertRequestBody) SetLastTriggeredAtNil(b bool)`

SetLastTriggeredAtNil sets the value for LastTriggeredAt to be an explicit nil

### UnsetLastTriggeredAt

`func (o *CreateAlertRequestBody) UnsetLastTriggeredAt()`

UnsetLastTriggeredAt ensures that no value is present for LastTriggeredAt, not even an explicit nil

### GetName

`func (o *CreateAlertRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAlertRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAlertRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateAlertRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *CreateAlertRequestBody) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *CreateAlertRequestBody) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *CreateAlertRequestBody) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *CreateAlertRequestBody) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOwner

`func (o *CreateAlertRequestBody) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateAlertRequestBody) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateAlertRequestBody) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CreateAlertRequestBody) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *CreateAlertRequestBody) SetOwnerNil(b bool)`

SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner

`func (o *CreateAlertRequestBody) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil

### GetQueryCondition

`func (o *CreateAlertRequestBody) GetQueryCondition() QueryCondition`

GetQueryCondition returns the QueryCondition field if non-nil, zero value otherwise.

### GetQueryConditionOk

`func (o *CreateAlertRequestBody) GetQueryConditionOk() (*QueryCondition, bool)`

GetQueryConditionOk returns a tuple with the QueryCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryCondition

`func (o *CreateAlertRequestBody) SetQueryCondition(v QueryCondition)`

SetQueryCondition sets QueryCondition field to given value.

### HasQueryCondition

`func (o *CreateAlertRequestBody) HasQueryCondition() bool`

HasQueryCondition returns a boolean if a field has been set.

### GetRowTemplate

`func (o *CreateAlertRequestBody) GetRowTemplate() string`

GetRowTemplate returns the RowTemplate field if non-nil, zero value otherwise.

### GetRowTemplateOk

`func (o *CreateAlertRequestBody) GetRowTemplateOk() (*string, bool)`

GetRowTemplateOk returns a tuple with the RowTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowTemplate

`func (o *CreateAlertRequestBody) SetRowTemplate(v string)`

SetRowTemplate sets RowTemplate field to given value.

### HasRowTemplate

`func (o *CreateAlertRequestBody) HasRowTemplate() bool`

HasRowTemplate returns a boolean if a field has been set.

### GetStreamName

`func (o *CreateAlertRequestBody) GetStreamName() string`

GetStreamName returns the StreamName field if non-nil, zero value otherwise.

### GetStreamNameOk

`func (o *CreateAlertRequestBody) GetStreamNameOk() (*string, bool)`

GetStreamNameOk returns a tuple with the StreamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamName

`func (o *CreateAlertRequestBody) SetStreamName(v string)`

SetStreamName sets StreamName field to given value.

### HasStreamName

`func (o *CreateAlertRequestBody) HasStreamName() bool`

HasStreamName returns a boolean if a field has been set.

### GetStreamType

`func (o *CreateAlertRequestBody) GetStreamType() StreamType`

GetStreamType returns the StreamType field if non-nil, zero value otherwise.

### GetStreamTypeOk

`func (o *CreateAlertRequestBody) GetStreamTypeOk() (*StreamType, bool)`

GetStreamTypeOk returns a tuple with the StreamType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamType

`func (o *CreateAlertRequestBody) SetStreamType(v StreamType)`

SetStreamType sets StreamType field to given value.

### HasStreamType

`func (o *CreateAlertRequestBody) HasStreamType() bool`

HasStreamType returns a boolean if a field has been set.

### GetTriggerCondition

`func (o *CreateAlertRequestBody) GetTriggerCondition() TriggerCondition`

GetTriggerCondition returns the TriggerCondition field if non-nil, zero value otherwise.

### GetTriggerConditionOk

`func (o *CreateAlertRequestBody) GetTriggerConditionOk() (*TriggerCondition, bool)`

GetTriggerConditionOk returns a tuple with the TriggerCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerCondition

`func (o *CreateAlertRequestBody) SetTriggerCondition(v TriggerCondition)`

SetTriggerCondition sets TriggerCondition field to given value.

### HasTriggerCondition

`func (o *CreateAlertRequestBody) HasTriggerCondition() bool`

HasTriggerCondition returns a boolean if a field has been set.

### GetTzOffset

`func (o *CreateAlertRequestBody) GetTzOffset() int32`

GetTzOffset returns the TzOffset field if non-nil, zero value otherwise.

### GetTzOffsetOk

`func (o *CreateAlertRequestBody) GetTzOffsetOk() (*int32, bool)`

GetTzOffsetOk returns a tuple with the TzOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTzOffset

`func (o *CreateAlertRequestBody) SetTzOffset(v int32)`

SetTzOffset sets TzOffset field to given value.

### HasTzOffset

`func (o *CreateAlertRequestBody) HasTzOffset() bool`

HasTzOffset returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CreateAlertRequestBody) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateAlertRequestBody) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateAlertRequestBody) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CreateAlertRequestBody) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *CreateAlertRequestBody) SetUpdatedAtNil(b bool)`

SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt

`func (o *CreateAlertRequestBody) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

### GetFolderId

`func (o *CreateAlertRequestBody) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *CreateAlertRequestBody) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *CreateAlertRequestBody) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *CreateAlertRequestBody) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *CreateAlertRequestBody) SetFolderIdNil(b bool)`

SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId

`func (o *CreateAlertRequestBody) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
