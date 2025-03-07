# Alert

## Properties

| Name                  | Type                                                   | Description                                                               | Notes                 |
| --------------------- | ------------------------------------------------------ | ------------------------------------------------------------------------- | --------------------- |
| **ContextAttributes** | Pointer to **map[string]string**                       |                                                                           | [optional]            |
| **Description**       | Pointer to **string**                                  |                                                                           | [optional]            |
| **Destinations**      | **[]string**                                           |                                                                           |
| **Enabled**           | Pointer to **bool**                                    |                                                                           | [optional]            |
| **Id**                | Pointer to [**Nullablestring**](string.md)             |                                                                           | [optional]            |
| **IsRealTime**        | Pointer to **bool**                                    |                                                                           | [optional]            |
| **LastEditedBy**      | Pointer to **NullableString**                          |                                                                           | [optional] [readonly] |
| **LastSatisfiedAt**   | Pointer to **NullableInt64**                           | Time when alert was last satisfied. Unix timestamp.                       | [optional] [readonly] |
| **LastTriggeredAt**   | Pointer to **NullableInt64**                           | Time when alert was last triggered. Unix timestamp.                       | [optional] [readonly] |
| **Name**              | Pointer to **string**                                  |                                                                           | [optional]            |
| **OrgId**             | Pointer to **string**                                  |                                                                           | [optional]            |
| **Owner**             | Pointer to **NullableString**                          |                                                                           | [optional]            |
| **QueryCondition**    | Pointer to [**QueryCondition**](QueryCondition.md)     |                                                                           | [optional]            |
| **RowTemplate**       | Pointer to **string**                                  |                                                                           | [optional]            |
| **StreamName**        | Pointer to **string**                                  |                                                                           | [optional]            |
| **StreamType**        | Pointer to [**StreamType**](StreamType.md)             |                                                                           | [optional]            |
| **TriggerCondition**  | Pointer to [**TriggerCondition**](TriggerCondition.md) |                                                                           | [optional]            |
| **TzOffset**          | Pointer to **int32**                                   | Timezone offset in minutes. Negative seconds means the western hemisphere | [optional]            |
| **UpdatedAt**         | Pointer to **NullableInt64**                           | Time when alert was last updated. Unix timestamp.                         | [optional] [readonly] |

## Methods

### NewAlert

`func NewAlert(destinations []string, ) *Alert`

NewAlert instantiates a new Alert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertWithDefaults

`func NewAlertWithDefaults() *Alert`

NewAlertWithDefaults instantiates a new Alert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextAttributes

`func (o *Alert) GetContextAttributes() map[string]string`

GetContextAttributes returns the ContextAttributes field if non-nil, zero value otherwise.

### GetContextAttributesOk

`func (o *Alert) GetContextAttributesOk() (*map[string]string, bool)`

GetContextAttributesOk returns a tuple with the ContextAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextAttributes

`func (o *Alert) SetContextAttributes(v map[string]string)`

SetContextAttributes sets ContextAttributes field to given value.

### HasContextAttributes

`func (o *Alert) HasContextAttributes() bool`

HasContextAttributes returns a boolean if a field has been set.

### SetContextAttributesNil

`func (o *Alert) SetContextAttributesNil(b bool)`

SetContextAttributesNil sets the value for ContextAttributes to be an explicit nil

### UnsetContextAttributes

`func (o *Alert) UnsetContextAttributes()`

UnsetContextAttributes ensures that no value is present for ContextAttributes, not even an explicit nil

### GetDescription

`func (o *Alert) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Alert) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Alert) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Alert) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDestinations

`func (o *Alert) GetDestinations() []string`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *Alert) GetDestinationsOk() (*[]string, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *Alert) SetDestinations(v []string)`

SetDestinations sets Destinations field to given value.

### GetEnabled

`func (o *Alert) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Alert) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Alert) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Alert) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *Alert) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Alert) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Alert) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Alert) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Alert) SetIdNil(b bool)`

SetIdNil sets the value for Id to be an explicit nil

### UnsetId

`func (o *Alert) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil

### GetIsRealTime

`func (o *Alert) GetIsRealTime() bool`

GetIsRealTime returns the IsRealTime field if non-nil, zero value otherwise.

### GetIsRealTimeOk

`func (o *Alert) GetIsRealTimeOk() (*bool, bool)`

GetIsRealTimeOk returns a tuple with the IsRealTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRealTime

`func (o *Alert) SetIsRealTime(v bool)`

SetIsRealTime sets IsRealTime field to given value.

### HasIsRealTime

`func (o *Alert) HasIsRealTime() bool`

HasIsRealTime returns a boolean if a field has been set.

### GetLastEditedBy

`func (o *Alert) GetLastEditedBy() string`

GetLastEditedBy returns the LastEditedBy field if non-nil, zero value otherwise.

### GetLastEditedByOk

`func (o *Alert) GetLastEditedByOk() (*string, bool)`

GetLastEditedByOk returns a tuple with the LastEditedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditedBy

`func (o *Alert) SetLastEditedBy(v string)`

SetLastEditedBy sets LastEditedBy field to given value.

### HasLastEditedBy

`func (o *Alert) HasLastEditedBy() bool`

HasLastEditedBy returns a boolean if a field has been set.

### SetLastEditedByNil

`func (o *Alert) SetLastEditedByNil(b bool)`

SetLastEditedByNil sets the value for LastEditedBy to be an explicit nil

### UnsetLastEditedBy

`func (o *Alert) UnsetLastEditedBy()`

UnsetLastEditedBy ensures that no value is present for LastEditedBy, not even an explicit nil

### GetLastSatisfiedAt

`func (o *Alert) GetLastSatisfiedAt() int64`

GetLastSatisfiedAt returns the LastSatisfiedAt field if non-nil, zero value otherwise.

### GetLastSatisfiedAtOk

`func (o *Alert) GetLastSatisfiedAtOk() (*int64, bool)`

GetLastSatisfiedAtOk returns a tuple with the LastSatisfiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSatisfiedAt

`func (o *Alert) SetLastSatisfiedAt(v int64)`

SetLastSatisfiedAt sets LastSatisfiedAt field to given value.

### HasLastSatisfiedAt

`func (o *Alert) HasLastSatisfiedAt() bool`

HasLastSatisfiedAt returns a boolean if a field has been set.

### SetLastSatisfiedAtNil

`func (o *Alert) SetLastSatisfiedAtNil(b bool)`

SetLastSatisfiedAtNil sets the value for LastSatisfiedAt to be an explicit nil

### UnsetLastSatisfiedAt

`func (o *Alert) UnsetLastSatisfiedAt()`

UnsetLastSatisfiedAt ensures that no value is present for LastSatisfiedAt, not even an explicit nil

### GetLastTriggeredAt

`func (o *Alert) GetLastTriggeredAt() int64`

GetLastTriggeredAt returns the LastTriggeredAt field if non-nil, zero value otherwise.

### GetLastTriggeredAtOk

`func (o *Alert) GetLastTriggeredAtOk() (*int64, bool)`

GetLastTriggeredAtOk returns a tuple with the LastTriggeredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTriggeredAt

`func (o *Alert) SetLastTriggeredAt(v int64)`

SetLastTriggeredAt sets LastTriggeredAt field to given value.

### HasLastTriggeredAt

`func (o *Alert) HasLastTriggeredAt() bool`

HasLastTriggeredAt returns a boolean if a field has been set.

### SetLastTriggeredAtNil

`func (o *Alert) SetLastTriggeredAtNil(b bool)`

SetLastTriggeredAtNil sets the value for LastTriggeredAt to be an explicit nil

### UnsetLastTriggeredAt

`func (o *Alert) UnsetLastTriggeredAt()`

UnsetLastTriggeredAt ensures that no value is present for LastTriggeredAt, not even an explicit nil

### GetName

`func (o *Alert) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Alert) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Alert) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Alert) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *Alert) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Alert) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Alert) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Alert) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOwner

`func (o *Alert) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Alert) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Alert) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Alert) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *Alert) SetOwnerNil(b bool)`

SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner

`func (o *Alert) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil

### GetQueryCondition

`func (o *Alert) GetQueryCondition() QueryCondition`

GetQueryCondition returns the QueryCondition field if non-nil, zero value otherwise.

### GetQueryConditionOk

`func (o *Alert) GetQueryConditionOk() (*QueryCondition, bool)`

GetQueryConditionOk returns a tuple with the QueryCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryCondition

`func (o *Alert) SetQueryCondition(v QueryCondition)`

SetQueryCondition sets QueryCondition field to given value.

### HasQueryCondition

`func (o *Alert) HasQueryCondition() bool`

HasQueryCondition returns a boolean if a field has been set.

### GetRowTemplate

`func (o *Alert) GetRowTemplate() string`

GetRowTemplate returns the RowTemplate field if non-nil, zero value otherwise.

### GetRowTemplateOk

`func (o *Alert) GetRowTemplateOk() (*string, bool)`

GetRowTemplateOk returns a tuple with the RowTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowTemplate

`func (o *Alert) SetRowTemplate(v string)`

SetRowTemplate sets RowTemplate field to given value.

### HasRowTemplate

`func (o *Alert) HasRowTemplate() bool`

HasRowTemplate returns a boolean if a field has been set.

### GetStreamName

`func (o *Alert) GetStreamName() string`

GetStreamName returns the StreamName field if non-nil, zero value otherwise.

### GetStreamNameOk

`func (o *Alert) GetStreamNameOk() (*string, bool)`

GetStreamNameOk returns a tuple with the StreamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamName

`func (o *Alert) SetStreamName(v string)`

SetStreamName sets StreamName field to given value.

### HasStreamName

`func (o *Alert) HasStreamName() bool`

HasStreamName returns a boolean if a field has been set.

### GetStreamType

`func (o *Alert) GetStreamType() StreamType`

GetStreamType returns the StreamType field if non-nil, zero value otherwise.

### GetStreamTypeOk

`func (o *Alert) GetStreamTypeOk() (*StreamType, bool)`

GetStreamTypeOk returns a tuple with the StreamType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamType

`func (o *Alert) SetStreamType(v StreamType)`

SetStreamType sets StreamType field to given value.

### HasStreamType

`func (o *Alert) HasStreamType() bool`

HasStreamType returns a boolean if a field has been set.

### GetTriggerCondition

`func (o *Alert) GetTriggerCondition() TriggerCondition`

GetTriggerCondition returns the TriggerCondition field if non-nil, zero value otherwise.

### GetTriggerConditionOk

`func (o *Alert) GetTriggerConditionOk() (*TriggerCondition, bool)`

GetTriggerConditionOk returns a tuple with the TriggerCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerCondition

`func (o *Alert) SetTriggerCondition(v TriggerCondition)`

SetTriggerCondition sets TriggerCondition field to given value.

### HasTriggerCondition

`func (o *Alert) HasTriggerCondition() bool`

HasTriggerCondition returns a boolean if a field has been set.

### GetTzOffset

`func (o *Alert) GetTzOffset() int32`

GetTzOffset returns the TzOffset field if non-nil, zero value otherwise.

### GetTzOffsetOk

`func (o *Alert) GetTzOffsetOk() (*int32, bool)`

GetTzOffsetOk returns a tuple with the TzOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTzOffset

`func (o *Alert) SetTzOffset(v int32)`

SetTzOffset sets TzOffset field to given value.

### HasTzOffset

`func (o *Alert) HasTzOffset() bool`

HasTzOffset returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Alert) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Alert) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Alert) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Alert) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *Alert) SetUpdatedAtNil(b bool)`

SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt

`func (o *Alert) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
