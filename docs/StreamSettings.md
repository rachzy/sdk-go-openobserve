# StreamSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApproxPartition** | Pointer to **bool** |  | [optional] 
**BloomFilterFields** | Pointer to **[]string** |  | [optional] 
**DataRetention** | Pointer to **int64** |  | [optional] 
**DefinedSchemaFields** | Pointer to **[]string** |  | [optional] 
**DistinctValueFields** | Pointer to [**[]DistinctField**](DistinctField.md) |  | [optional] 
**ExtendedRetentionDays** | Pointer to [**[]TimeRange**](TimeRange.md) |  | [optional] 
**FlattenLevel** | Pointer to **NullableInt64** |  | [optional] 
**FullTextSearchKeys** | Pointer to **[]string** |  | [optional] 
**IndexFields** | Pointer to **[]string** |  | [optional] 
**IndexUpdatedAt** | Pointer to **int64** |  | [optional] 
**MaxQueryRange** | Pointer to **int64** |  | [optional] 
**PartitionKeys** | Pointer to [**[]StreamPartition**](StreamPartition.md) |  | [optional] 
**PartitionTimeLevel** | Pointer to [**NullablePartitionTimeLevel**](PartitionTimeLevel.md) |  | [optional] 
**StoreOriginalData** | Pointer to **bool** |  | [optional] 

## Methods

### NewStreamSettings

`func NewStreamSettings() *StreamSettings`

NewStreamSettings instantiates a new StreamSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamSettingsWithDefaults

`func NewStreamSettingsWithDefaults() *StreamSettings`

NewStreamSettingsWithDefaults instantiates a new StreamSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproxPartition

`func (o *StreamSettings) GetApproxPartition() bool`

GetApproxPartition returns the ApproxPartition field if non-nil, zero value otherwise.

### GetApproxPartitionOk

`func (o *StreamSettings) GetApproxPartitionOk() (*bool, bool)`

GetApproxPartitionOk returns a tuple with the ApproxPartition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproxPartition

`func (o *StreamSettings) SetApproxPartition(v bool)`

SetApproxPartition sets ApproxPartition field to given value.

### HasApproxPartition

`func (o *StreamSettings) HasApproxPartition() bool`

HasApproxPartition returns a boolean if a field has been set.

### GetBloomFilterFields

`func (o *StreamSettings) GetBloomFilterFields() []string`

GetBloomFilterFields returns the BloomFilterFields field if non-nil, zero value otherwise.

### GetBloomFilterFieldsOk

`func (o *StreamSettings) GetBloomFilterFieldsOk() (*[]string, bool)`

GetBloomFilterFieldsOk returns a tuple with the BloomFilterFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBloomFilterFields

`func (o *StreamSettings) SetBloomFilterFields(v []string)`

SetBloomFilterFields sets BloomFilterFields field to given value.

### HasBloomFilterFields

`func (o *StreamSettings) HasBloomFilterFields() bool`

HasBloomFilterFields returns a boolean if a field has been set.

### GetDataRetention

`func (o *StreamSettings) GetDataRetention() int64`

GetDataRetention returns the DataRetention field if non-nil, zero value otherwise.

### GetDataRetentionOk

`func (o *StreamSettings) GetDataRetentionOk() (*int64, bool)`

GetDataRetentionOk returns a tuple with the DataRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRetention

`func (o *StreamSettings) SetDataRetention(v int64)`

SetDataRetention sets DataRetention field to given value.

### HasDataRetention

`func (o *StreamSettings) HasDataRetention() bool`

HasDataRetention returns a boolean if a field has been set.

### GetDefinedSchemaFields

`func (o *StreamSettings) GetDefinedSchemaFields() []string`

GetDefinedSchemaFields returns the DefinedSchemaFields field if non-nil, zero value otherwise.

### GetDefinedSchemaFieldsOk

`func (o *StreamSettings) GetDefinedSchemaFieldsOk() (*[]string, bool)`

GetDefinedSchemaFieldsOk returns a tuple with the DefinedSchemaFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinedSchemaFields

`func (o *StreamSettings) SetDefinedSchemaFields(v []string)`

SetDefinedSchemaFields sets DefinedSchemaFields field to given value.

### HasDefinedSchemaFields

`func (o *StreamSettings) HasDefinedSchemaFields() bool`

HasDefinedSchemaFields returns a boolean if a field has been set.

### SetDefinedSchemaFieldsNil

`func (o *StreamSettings) SetDefinedSchemaFieldsNil(b bool)`

 SetDefinedSchemaFieldsNil sets the value for DefinedSchemaFields to be an explicit nil

### UnsetDefinedSchemaFields
`func (o *StreamSettings) UnsetDefinedSchemaFields()`

UnsetDefinedSchemaFields ensures that no value is present for DefinedSchemaFields, not even an explicit nil
### GetDistinctValueFields

`func (o *StreamSettings) GetDistinctValueFields() []DistinctField`

GetDistinctValueFields returns the DistinctValueFields field if non-nil, zero value otherwise.

### GetDistinctValueFieldsOk

`func (o *StreamSettings) GetDistinctValueFieldsOk() (*[]DistinctField, bool)`

GetDistinctValueFieldsOk returns a tuple with the DistinctValueFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctValueFields

`func (o *StreamSettings) SetDistinctValueFields(v []DistinctField)`

SetDistinctValueFields sets DistinctValueFields field to given value.

### HasDistinctValueFields

`func (o *StreamSettings) HasDistinctValueFields() bool`

HasDistinctValueFields returns a boolean if a field has been set.

### GetExtendedRetentionDays

`func (o *StreamSettings) GetExtendedRetentionDays() []TimeRange`

GetExtendedRetentionDays returns the ExtendedRetentionDays field if non-nil, zero value otherwise.

### GetExtendedRetentionDaysOk

`func (o *StreamSettings) GetExtendedRetentionDaysOk() (*[]TimeRange, bool)`

GetExtendedRetentionDaysOk returns a tuple with the ExtendedRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedRetentionDays

`func (o *StreamSettings) SetExtendedRetentionDays(v []TimeRange)`

SetExtendedRetentionDays sets ExtendedRetentionDays field to given value.

### HasExtendedRetentionDays

`func (o *StreamSettings) HasExtendedRetentionDays() bool`

HasExtendedRetentionDays returns a boolean if a field has been set.

### GetFlattenLevel

`func (o *StreamSettings) GetFlattenLevel() int64`

GetFlattenLevel returns the FlattenLevel field if non-nil, zero value otherwise.

### GetFlattenLevelOk

`func (o *StreamSettings) GetFlattenLevelOk() (*int64, bool)`

GetFlattenLevelOk returns a tuple with the FlattenLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlattenLevel

`func (o *StreamSettings) SetFlattenLevel(v int64)`

SetFlattenLevel sets FlattenLevel field to given value.

### HasFlattenLevel

`func (o *StreamSettings) HasFlattenLevel() bool`

HasFlattenLevel returns a boolean if a field has been set.

### SetFlattenLevelNil

`func (o *StreamSettings) SetFlattenLevelNil(b bool)`

 SetFlattenLevelNil sets the value for FlattenLevel to be an explicit nil

### UnsetFlattenLevel
`func (o *StreamSettings) UnsetFlattenLevel()`

UnsetFlattenLevel ensures that no value is present for FlattenLevel, not even an explicit nil
### GetFullTextSearchKeys

`func (o *StreamSettings) GetFullTextSearchKeys() []string`

GetFullTextSearchKeys returns the FullTextSearchKeys field if non-nil, zero value otherwise.

### GetFullTextSearchKeysOk

`func (o *StreamSettings) GetFullTextSearchKeysOk() (*[]string, bool)`

GetFullTextSearchKeysOk returns a tuple with the FullTextSearchKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullTextSearchKeys

`func (o *StreamSettings) SetFullTextSearchKeys(v []string)`

SetFullTextSearchKeys sets FullTextSearchKeys field to given value.

### HasFullTextSearchKeys

`func (o *StreamSettings) HasFullTextSearchKeys() bool`

HasFullTextSearchKeys returns a boolean if a field has been set.

### GetIndexFields

`func (o *StreamSettings) GetIndexFields() []string`

GetIndexFields returns the IndexFields field if non-nil, zero value otherwise.

### GetIndexFieldsOk

`func (o *StreamSettings) GetIndexFieldsOk() (*[]string, bool)`

GetIndexFieldsOk returns a tuple with the IndexFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexFields

`func (o *StreamSettings) SetIndexFields(v []string)`

SetIndexFields sets IndexFields field to given value.

### HasIndexFields

`func (o *StreamSettings) HasIndexFields() bool`

HasIndexFields returns a boolean if a field has been set.

### GetIndexUpdatedAt

`func (o *StreamSettings) GetIndexUpdatedAt() int64`

GetIndexUpdatedAt returns the IndexUpdatedAt field if non-nil, zero value otherwise.

### GetIndexUpdatedAtOk

`func (o *StreamSettings) GetIndexUpdatedAtOk() (*int64, bool)`

GetIndexUpdatedAtOk returns a tuple with the IndexUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexUpdatedAt

`func (o *StreamSettings) SetIndexUpdatedAt(v int64)`

SetIndexUpdatedAt sets IndexUpdatedAt field to given value.

### HasIndexUpdatedAt

`func (o *StreamSettings) HasIndexUpdatedAt() bool`

HasIndexUpdatedAt returns a boolean if a field has been set.

### GetMaxQueryRange

`func (o *StreamSettings) GetMaxQueryRange() int64`

GetMaxQueryRange returns the MaxQueryRange field if non-nil, zero value otherwise.

### GetMaxQueryRangeOk

`func (o *StreamSettings) GetMaxQueryRangeOk() (*int64, bool)`

GetMaxQueryRangeOk returns a tuple with the MaxQueryRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQueryRange

`func (o *StreamSettings) SetMaxQueryRange(v int64)`

SetMaxQueryRange sets MaxQueryRange field to given value.

### HasMaxQueryRange

`func (o *StreamSettings) HasMaxQueryRange() bool`

HasMaxQueryRange returns a boolean if a field has been set.

### GetPartitionKeys

`func (o *StreamSettings) GetPartitionKeys() []StreamPartition`

GetPartitionKeys returns the PartitionKeys field if non-nil, zero value otherwise.

### GetPartitionKeysOk

`func (o *StreamSettings) GetPartitionKeysOk() (*[]StreamPartition, bool)`

GetPartitionKeysOk returns a tuple with the PartitionKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionKeys

`func (o *StreamSettings) SetPartitionKeys(v []StreamPartition)`

SetPartitionKeys sets PartitionKeys field to given value.

### HasPartitionKeys

`func (o *StreamSettings) HasPartitionKeys() bool`

HasPartitionKeys returns a boolean if a field has been set.

### GetPartitionTimeLevel

`func (o *StreamSettings) GetPartitionTimeLevel() PartitionTimeLevel`

GetPartitionTimeLevel returns the PartitionTimeLevel field if non-nil, zero value otherwise.

### GetPartitionTimeLevelOk

`func (o *StreamSettings) GetPartitionTimeLevelOk() (*PartitionTimeLevel, bool)`

GetPartitionTimeLevelOk returns a tuple with the PartitionTimeLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionTimeLevel

`func (o *StreamSettings) SetPartitionTimeLevel(v PartitionTimeLevel)`

SetPartitionTimeLevel sets PartitionTimeLevel field to given value.

### HasPartitionTimeLevel

`func (o *StreamSettings) HasPartitionTimeLevel() bool`

HasPartitionTimeLevel returns a boolean if a field has been set.

### SetPartitionTimeLevelNil

`func (o *StreamSettings) SetPartitionTimeLevelNil(b bool)`

 SetPartitionTimeLevelNil sets the value for PartitionTimeLevel to be an explicit nil

### UnsetPartitionTimeLevel
`func (o *StreamSettings) UnsetPartitionTimeLevel()`

UnsetPartitionTimeLevel ensures that no value is present for PartitionTimeLevel, not even an explicit nil
### GetStoreOriginalData

`func (o *StreamSettings) GetStoreOriginalData() bool`

GetStoreOriginalData returns the StoreOriginalData field if non-nil, zero value otherwise.

### GetStoreOriginalDataOk

`func (o *StreamSettings) GetStoreOriginalDataOk() (*bool, bool)`

GetStoreOriginalDataOk returns a tuple with the StoreOriginalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreOriginalData

`func (o *StreamSettings) SetStoreOriginalData(v bool)`

SetStoreOriginalData sets StoreOriginalData field to given value.

### HasStoreOriginalData

`func (o *StreamSettings) HasStoreOriginalData() bool`

HasStoreOriginalData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


