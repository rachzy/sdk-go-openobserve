# UpdateStreamSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApproxPartition** | Pointer to **NullableBool** |  | [optional] 
**BloomFilterFields** | Pointer to [**UpdateSettingsWrapper**](UpdateSettingsWrapper.md) |  | [optional] 
**DataRetention** | Pointer to **NullableInt64** |  | [optional] 
**DefinedSchemaFields** | Pointer to [**UpdateSettingsWrapper**](UpdateSettingsWrapper.md) |  | [optional] 
**DistinctValueFields** | Pointer to [**UpdateSettingsWrapper**](UpdateSettingsWrapper.md) |  | [optional] 
**ExtendedRetentionDays** | Pointer to [**UpdateSettingsWrapper**](UpdateSettingsWrapper.md) |  | [optional] 
**FlattenLevel** | Pointer to **NullableInt64** |  | [optional] 
**FullTextSearchKeys** | Pointer to [**UpdateSettingsWrapper**](UpdateSettingsWrapper.md) |  | [optional] 
**IndexFields** | Pointer to [**UpdateSettingsWrapper**](UpdateSettingsWrapper.md) |  | [optional] 
**MaxQueryRange** | Pointer to **NullableInt64** |  | [optional] 
**PartitionKeys** | Pointer to [**UpdateSettingsWrapper**](UpdateSettingsWrapper.md) |  | [optional] 
**PartitionTimeLevel** | Pointer to [**NullablePartitionTimeLevel**](PartitionTimeLevel.md) |  | [optional] 
**StoreOriginalData** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewUpdateStreamSettings

`func NewUpdateStreamSettings() *UpdateStreamSettings`

NewUpdateStreamSettings instantiates a new UpdateStreamSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStreamSettingsWithDefaults

`func NewUpdateStreamSettingsWithDefaults() *UpdateStreamSettings`

NewUpdateStreamSettingsWithDefaults instantiates a new UpdateStreamSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproxPartition

`func (o *UpdateStreamSettings) GetApproxPartition() bool`

GetApproxPartition returns the ApproxPartition field if non-nil, zero value otherwise.

### GetApproxPartitionOk

`func (o *UpdateStreamSettings) GetApproxPartitionOk() (*bool, bool)`

GetApproxPartitionOk returns a tuple with the ApproxPartition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproxPartition

`func (o *UpdateStreamSettings) SetApproxPartition(v bool)`

SetApproxPartition sets ApproxPartition field to given value.

### HasApproxPartition

`func (o *UpdateStreamSettings) HasApproxPartition() bool`

HasApproxPartition returns a boolean if a field has been set.

### SetApproxPartitionNil

`func (o *UpdateStreamSettings) SetApproxPartitionNil(b bool)`

 SetApproxPartitionNil sets the value for ApproxPartition to be an explicit nil

### UnsetApproxPartition
`func (o *UpdateStreamSettings) UnsetApproxPartition()`

UnsetApproxPartition ensures that no value is present for ApproxPartition, not even an explicit nil
### GetBloomFilterFields

`func (o *UpdateStreamSettings) GetBloomFilterFields() UpdateSettingsWrapper`

GetBloomFilterFields returns the BloomFilterFields field if non-nil, zero value otherwise.

### GetBloomFilterFieldsOk

`func (o *UpdateStreamSettings) GetBloomFilterFieldsOk() (*UpdateSettingsWrapper, bool)`

GetBloomFilterFieldsOk returns a tuple with the BloomFilterFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBloomFilterFields

`func (o *UpdateStreamSettings) SetBloomFilterFields(v UpdateSettingsWrapper)`

SetBloomFilterFields sets BloomFilterFields field to given value.

### HasBloomFilterFields

`func (o *UpdateStreamSettings) HasBloomFilterFields() bool`

HasBloomFilterFields returns a boolean if a field has been set.

### GetDataRetention

`func (o *UpdateStreamSettings) GetDataRetention() int64`

GetDataRetention returns the DataRetention field if non-nil, zero value otherwise.

### GetDataRetentionOk

`func (o *UpdateStreamSettings) GetDataRetentionOk() (*int64, bool)`

GetDataRetentionOk returns a tuple with the DataRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRetention

`func (o *UpdateStreamSettings) SetDataRetention(v int64)`

SetDataRetention sets DataRetention field to given value.

### HasDataRetention

`func (o *UpdateStreamSettings) HasDataRetention() bool`

HasDataRetention returns a boolean if a field has been set.

### SetDataRetentionNil

`func (o *UpdateStreamSettings) SetDataRetentionNil(b bool)`

 SetDataRetentionNil sets the value for DataRetention to be an explicit nil

### UnsetDataRetention
`func (o *UpdateStreamSettings) UnsetDataRetention()`

UnsetDataRetention ensures that no value is present for DataRetention, not even an explicit nil
### GetDefinedSchemaFields

`func (o *UpdateStreamSettings) GetDefinedSchemaFields() UpdateSettingsWrapper`

GetDefinedSchemaFields returns the DefinedSchemaFields field if non-nil, zero value otherwise.

### GetDefinedSchemaFieldsOk

`func (o *UpdateStreamSettings) GetDefinedSchemaFieldsOk() (*UpdateSettingsWrapper, bool)`

GetDefinedSchemaFieldsOk returns a tuple with the DefinedSchemaFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinedSchemaFields

`func (o *UpdateStreamSettings) SetDefinedSchemaFields(v UpdateSettingsWrapper)`

SetDefinedSchemaFields sets DefinedSchemaFields field to given value.

### HasDefinedSchemaFields

`func (o *UpdateStreamSettings) HasDefinedSchemaFields() bool`

HasDefinedSchemaFields returns a boolean if a field has been set.

### GetDistinctValueFields

`func (o *UpdateStreamSettings) GetDistinctValueFields() UpdateSettingsWrapper`

GetDistinctValueFields returns the DistinctValueFields field if non-nil, zero value otherwise.

### GetDistinctValueFieldsOk

`func (o *UpdateStreamSettings) GetDistinctValueFieldsOk() (*UpdateSettingsWrapper, bool)`

GetDistinctValueFieldsOk returns a tuple with the DistinctValueFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctValueFields

`func (o *UpdateStreamSettings) SetDistinctValueFields(v UpdateSettingsWrapper)`

SetDistinctValueFields sets DistinctValueFields field to given value.

### HasDistinctValueFields

`func (o *UpdateStreamSettings) HasDistinctValueFields() bool`

HasDistinctValueFields returns a boolean if a field has been set.

### GetExtendedRetentionDays

`func (o *UpdateStreamSettings) GetExtendedRetentionDays() UpdateSettingsWrapper`

GetExtendedRetentionDays returns the ExtendedRetentionDays field if non-nil, zero value otherwise.

### GetExtendedRetentionDaysOk

`func (o *UpdateStreamSettings) GetExtendedRetentionDaysOk() (*UpdateSettingsWrapper, bool)`

GetExtendedRetentionDaysOk returns a tuple with the ExtendedRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedRetentionDays

`func (o *UpdateStreamSettings) SetExtendedRetentionDays(v UpdateSettingsWrapper)`

SetExtendedRetentionDays sets ExtendedRetentionDays field to given value.

### HasExtendedRetentionDays

`func (o *UpdateStreamSettings) HasExtendedRetentionDays() bool`

HasExtendedRetentionDays returns a boolean if a field has been set.

### GetFlattenLevel

`func (o *UpdateStreamSettings) GetFlattenLevel() int64`

GetFlattenLevel returns the FlattenLevel field if non-nil, zero value otherwise.

### GetFlattenLevelOk

`func (o *UpdateStreamSettings) GetFlattenLevelOk() (*int64, bool)`

GetFlattenLevelOk returns a tuple with the FlattenLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlattenLevel

`func (o *UpdateStreamSettings) SetFlattenLevel(v int64)`

SetFlattenLevel sets FlattenLevel field to given value.

### HasFlattenLevel

`func (o *UpdateStreamSettings) HasFlattenLevel() bool`

HasFlattenLevel returns a boolean if a field has been set.

### SetFlattenLevelNil

`func (o *UpdateStreamSettings) SetFlattenLevelNil(b bool)`

 SetFlattenLevelNil sets the value for FlattenLevel to be an explicit nil

### UnsetFlattenLevel
`func (o *UpdateStreamSettings) UnsetFlattenLevel()`

UnsetFlattenLevel ensures that no value is present for FlattenLevel, not even an explicit nil
### GetFullTextSearchKeys

`func (o *UpdateStreamSettings) GetFullTextSearchKeys() UpdateSettingsWrapper`

GetFullTextSearchKeys returns the FullTextSearchKeys field if non-nil, zero value otherwise.

### GetFullTextSearchKeysOk

`func (o *UpdateStreamSettings) GetFullTextSearchKeysOk() (*UpdateSettingsWrapper, bool)`

GetFullTextSearchKeysOk returns a tuple with the FullTextSearchKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullTextSearchKeys

`func (o *UpdateStreamSettings) SetFullTextSearchKeys(v UpdateSettingsWrapper)`

SetFullTextSearchKeys sets FullTextSearchKeys field to given value.

### HasFullTextSearchKeys

`func (o *UpdateStreamSettings) HasFullTextSearchKeys() bool`

HasFullTextSearchKeys returns a boolean if a field has been set.

### GetIndexFields

`func (o *UpdateStreamSettings) GetIndexFields() UpdateSettingsWrapper`

GetIndexFields returns the IndexFields field if non-nil, zero value otherwise.

### GetIndexFieldsOk

`func (o *UpdateStreamSettings) GetIndexFieldsOk() (*UpdateSettingsWrapper, bool)`

GetIndexFieldsOk returns a tuple with the IndexFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexFields

`func (o *UpdateStreamSettings) SetIndexFields(v UpdateSettingsWrapper)`

SetIndexFields sets IndexFields field to given value.

### HasIndexFields

`func (o *UpdateStreamSettings) HasIndexFields() bool`

HasIndexFields returns a boolean if a field has been set.

### GetMaxQueryRange

`func (o *UpdateStreamSettings) GetMaxQueryRange() int64`

GetMaxQueryRange returns the MaxQueryRange field if non-nil, zero value otherwise.

### GetMaxQueryRangeOk

`func (o *UpdateStreamSettings) GetMaxQueryRangeOk() (*int64, bool)`

GetMaxQueryRangeOk returns a tuple with the MaxQueryRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQueryRange

`func (o *UpdateStreamSettings) SetMaxQueryRange(v int64)`

SetMaxQueryRange sets MaxQueryRange field to given value.

### HasMaxQueryRange

`func (o *UpdateStreamSettings) HasMaxQueryRange() bool`

HasMaxQueryRange returns a boolean if a field has been set.

### SetMaxQueryRangeNil

`func (o *UpdateStreamSettings) SetMaxQueryRangeNil(b bool)`

 SetMaxQueryRangeNil sets the value for MaxQueryRange to be an explicit nil

### UnsetMaxQueryRange
`func (o *UpdateStreamSettings) UnsetMaxQueryRange()`

UnsetMaxQueryRange ensures that no value is present for MaxQueryRange, not even an explicit nil
### GetPartitionKeys

`func (o *UpdateStreamSettings) GetPartitionKeys() UpdateSettingsWrapper`

GetPartitionKeys returns the PartitionKeys field if non-nil, zero value otherwise.

### GetPartitionKeysOk

`func (o *UpdateStreamSettings) GetPartitionKeysOk() (*UpdateSettingsWrapper, bool)`

GetPartitionKeysOk returns a tuple with the PartitionKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionKeys

`func (o *UpdateStreamSettings) SetPartitionKeys(v UpdateSettingsWrapper)`

SetPartitionKeys sets PartitionKeys field to given value.

### HasPartitionKeys

`func (o *UpdateStreamSettings) HasPartitionKeys() bool`

HasPartitionKeys returns a boolean if a field has been set.

### GetPartitionTimeLevel

`func (o *UpdateStreamSettings) GetPartitionTimeLevel() PartitionTimeLevel`

GetPartitionTimeLevel returns the PartitionTimeLevel field if non-nil, zero value otherwise.

### GetPartitionTimeLevelOk

`func (o *UpdateStreamSettings) GetPartitionTimeLevelOk() (*PartitionTimeLevel, bool)`

GetPartitionTimeLevelOk returns a tuple with the PartitionTimeLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionTimeLevel

`func (o *UpdateStreamSettings) SetPartitionTimeLevel(v PartitionTimeLevel)`

SetPartitionTimeLevel sets PartitionTimeLevel field to given value.

### HasPartitionTimeLevel

`func (o *UpdateStreamSettings) HasPartitionTimeLevel() bool`

HasPartitionTimeLevel returns a boolean if a field has been set.

### SetPartitionTimeLevelNil

`func (o *UpdateStreamSettings) SetPartitionTimeLevelNil(b bool)`

 SetPartitionTimeLevelNil sets the value for PartitionTimeLevel to be an explicit nil

### UnsetPartitionTimeLevel
`func (o *UpdateStreamSettings) UnsetPartitionTimeLevel()`

UnsetPartitionTimeLevel ensures that no value is present for PartitionTimeLevel, not even an explicit nil
### GetStoreOriginalData

`func (o *UpdateStreamSettings) GetStoreOriginalData() bool`

GetStoreOriginalData returns the StoreOriginalData field if non-nil, zero value otherwise.

### GetStoreOriginalDataOk

`func (o *UpdateStreamSettings) GetStoreOriginalDataOk() (*bool, bool)`

GetStoreOriginalDataOk returns a tuple with the StoreOriginalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreOriginalData

`func (o *UpdateStreamSettings) SetStoreOriginalData(v bool)`

SetStoreOriginalData sets StoreOriginalData field to given value.

### HasStoreOriginalData

`func (o *UpdateStreamSettings) HasStoreOriginalData() bool`

HasStoreOriginalData returns a boolean if a field has been set.

### SetStoreOriginalDataNil

`func (o *UpdateStreamSettings) SetStoreOriginalDataNil(b bool)`

 SetStoreOriginalDataNil sets the value for StoreOriginalData to be an explicit nil

### UnsetStoreOriginalData
`func (o *UpdateStreamSettings) UnsetStoreOriginalData()`

UnsetStoreOriginalData ensures that no value is present for StoreOriginalData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


