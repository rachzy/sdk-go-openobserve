/*
openobserve

OpenObserve API documents [https://openobserve.ai/docs/](https://openobserve.ai/docs/)

API version: 0.14.5
Contact: hello@zinclabs.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the StreamSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StreamSettings{}

// StreamSettings struct for StreamSettings
type StreamSettings struct {
	ApproxPartition *bool `json:"approx_partition,omitempty"`
	BloomFilterFields []string `json:"bloom_filter_fields,omitempty"`
	DataRetention *int64 `json:"data_retention,omitempty"`
	DefinedSchemaFields []string `json:"defined_schema_fields,omitempty"`
	DistinctValueFields []CustomFieldsOption `json:"distinct_value_fields,omitempty"`
	ExtendedRetentionDays []any `json:"extended_retention_days,omitempty"`
	FlattenLevel NullableInt64 `json:"flatten_level,omitempty"`
	FullTextSearchKeys []string `json:"full_text_search_keys,omitempty"`
	IndexFields []string `json:"index_fields,omitempty"`
	IndexUpdatedAt *int64 `json:"index_updated_at,omitempty"`
	MaxQueryRange *int64 `json:"max_query_range,omitempty"`
	PartitionKeys []StreamPartition `json:"partition_keys,omitempty"`
	PartitionTimeLevel NullablePartitionTimeLevel `json:"partition_time_level,omitempty"`
	StoreOriginalData *bool `json:"store_original_data,omitempty"`
}

// NewStreamSettings instantiates a new StreamSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamSettings() *StreamSettings {
	this := StreamSettings{}
	return &this
}

// NewStreamSettingsWithDefaults instantiates a new StreamSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamSettingsWithDefaults() *StreamSettings {
	this := StreamSettings{}
	return &this
}

// GetApproxPartition returns the ApproxPartition field value if set, zero value otherwise.
func (o *StreamSettings) GetApproxPartition() bool {
	if o == nil || IsNil(o.ApproxPartition) {
		var ret bool
		return ret
	}
	return *o.ApproxPartition
}

// GetApproxPartitionOk returns a tuple with the ApproxPartition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamSettings) GetApproxPartitionOk() (*bool, bool) {
	if o == nil || IsNil(o.ApproxPartition) {
		return nil, false
	}
	return o.ApproxPartition, true
}

// HasApproxPartition returns a boolean if a field has been set.
func (o *StreamSettings) HasApproxPartition() bool {
	if o != nil && !IsNil(o.ApproxPartition) {
		return true
	}

	return false
}

// SetApproxPartition gets a reference to the given bool and assigns it to the ApproxPartition field.
func (o *StreamSettings) SetApproxPartition(v bool) {
	o.ApproxPartition = &v
}

// GetBloomFilterFields returns the BloomFilterFields field value if set, zero value otherwise.
func (o *StreamSettings) GetBloomFilterFields() []string {
	if o == nil || IsNil(o.BloomFilterFields) {
		var ret []string
		return ret
	}
	return o.BloomFilterFields
}

// GetBloomFilterFieldsOk returns a tuple with the BloomFilterFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamSettings) GetBloomFilterFieldsOk() ([]string, bool) {
	if o == nil || IsNil(o.BloomFilterFields) {
		return nil, false
	}
	return o.BloomFilterFields, true
}

// HasBloomFilterFields returns a boolean if a field has been set.
func (o *StreamSettings) HasBloomFilterFields() bool {
	if o != nil && !IsNil(o.BloomFilterFields) {
		return true
	}

	return false
}

// SetBloomFilterFields gets a reference to the given []string and assigns it to the BloomFilterFields field.
func (o *StreamSettings) SetBloomFilterFields(v []string) {
	o.BloomFilterFields = v
}

// GetDataRetention returns the DataRetention field value if set, zero value otherwise.
func (o *StreamSettings) GetDataRetention() int64 {
	if o == nil || IsNil(o.DataRetention) {
		var ret int64
		return ret
	}
	return *o.DataRetention
}

// GetDataRetentionOk returns a tuple with the DataRetention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamSettings) GetDataRetentionOk() (*int64, bool) {
	if o == nil || IsNil(o.DataRetention) {
		return nil, false
	}
	return o.DataRetention, true
}

// HasDataRetention returns a boolean if a field has been set.
func (o *StreamSettings) HasDataRetention() bool {
	if o != nil && !IsNil(o.DataRetention) {
		return true
	}

	return false
}

// SetDataRetention gets a reference to the given int64 and assigns it to the DataRetention field.
func (o *StreamSettings) SetDataRetention(v int64) {
	o.DataRetention = &v
}

// GetDefinedSchemaFields returns the DefinedSchemaFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StreamSettings) GetDefinedSchemaFields() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DefinedSchemaFields
}

// GetDefinedSchemaFieldsOk returns a tuple with the DefinedSchemaFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StreamSettings) GetDefinedSchemaFieldsOk() ([]string, bool) {
	if o == nil || IsNil(o.DefinedSchemaFields) {
		return nil, false
	}
	return o.DefinedSchemaFields, true
}

// HasDefinedSchemaFields returns a boolean if a field has been set.
func (o *StreamSettings) HasDefinedSchemaFields() bool {
	if o != nil && !IsNil(o.DefinedSchemaFields) {
		return true
	}

	return false
}

// SetDefinedSchemaFields gets a reference to the given []string and assigns it to the DefinedSchemaFields field.
func (o *StreamSettings) SetDefinedSchemaFields(v []string) {
	o.DefinedSchemaFields = v
}

// GetDistinctValueFields returns the DistinctValueFields field value if set, zero value otherwise.
func (o *StreamSettings) GetDistinctValueFields() []CustomFieldsOption {
	if o == nil || IsNil(o.DistinctValueFields) {
		var ret []CustomFieldsOption
		return ret
	}
	return o.DistinctValueFields
}

// GetDistinctValueFieldsOk returns a tuple with the DistinctValueFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamSettings) GetDistinctValueFieldsOk() ([]CustomFieldsOption, bool) {
	if o == nil || IsNil(o.DistinctValueFields) {
		return nil, false
	}
	return o.DistinctValueFields, true
}

// HasDistinctValueFields returns a boolean if a field has been set.
func (o *StreamSettings) HasDistinctValueFields() bool {
	if o != nil && !IsNil(o.DistinctValueFields) {
		return true
	}

	return false
}

// SetDistinctValueFields gets a reference to the given []CustomFieldsOption and assigns it to the DistinctValueFields field.
func (o *StreamSettings) SetDistinctValueFields(v []CustomFieldsOption) {
	o.DistinctValueFields = v
}

// GetExtendedRetentionDays returns the ExtendedRetentionDays field value if set, zero value otherwise.
func (o *StreamSettings) GetExtendedRetentionDays() []any {
	if o == nil || IsNil(o.ExtendedRetentionDays) {
		var ret []any
		return ret
	}
	return o.ExtendedRetentionDays
}

// GetExtendedRetentionDaysOk returns a tuple with the ExtendedRetentionDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamSettings) GetExtendedRetentionDaysOk() ([]any, bool) {
	if o == nil || IsNil(o.ExtendedRetentionDays) {
		return nil, false
	}
	return o.ExtendedRetentionDays, true
}

// HasExtendedRetentionDays returns a boolean if a field has been set.
func (o *StreamSettings) HasExtendedRetentionDays() bool {
	if o != nil && !IsNil(o.ExtendedRetentionDays) {
		return true
	}

	return false
}

// SetExtendedRetentionDays gets a reference to the given []any and assigns it to the ExtendedRetentionDays field.
func (o *StreamSettings) SetExtendedRetentionDays(v []any) {
	o.ExtendedRetentionDays = v
}

// GetFlattenLevel returns the FlattenLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StreamSettings) GetFlattenLevel() int64 {
	if o == nil || IsNil(o.FlattenLevel.Get()) {
		var ret int64
		return ret
	}
	return *o.FlattenLevel.Get()
}

// GetFlattenLevelOk returns a tuple with the FlattenLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StreamSettings) GetFlattenLevelOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlattenLevel.Get(), o.FlattenLevel.IsSet()
}

// HasFlattenLevel returns a boolean if a field has been set.
func (o *StreamSettings) HasFlattenLevel() bool {
	if o != nil && o.FlattenLevel.IsSet() {
		return true
	}

	return false
}

// SetFlattenLevel gets a reference to the given NullableInt64 and assigns it to the FlattenLevel field.
func (o *StreamSettings) SetFlattenLevel(v int64) {
	o.FlattenLevel.Set(&v)
}
// SetFlattenLevelNil sets the value for FlattenLevel to be an explicit nil
func (o *StreamSettings) SetFlattenLevelNil() {
	o.FlattenLevel.Set(nil)
}

// UnsetFlattenLevel ensures that no value is present for FlattenLevel, not even an explicit nil
func (o *StreamSettings) UnsetFlattenLevel() {
	o.FlattenLevel.Unset()
}

// GetFullTextSearchKeys returns the FullTextSearchKeys field value if set, zero value otherwise.
func (o *StreamSettings) GetFullTextSearchKeys() []string {
	if o == nil || IsNil(o.FullTextSearchKeys) {
		var ret []string
		return ret
	}
	return o.FullTextSearchKeys
}

// GetFullTextSearchKeysOk returns a tuple with the FullTextSearchKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamSettings) GetFullTextSearchKeysOk() ([]string, bool) {
	if o == nil || IsNil(o.FullTextSearchKeys) {
		return nil, false
	}
	return o.FullTextSearchKeys, true
}

// HasFullTextSearchKeys returns a boolean if a field has been set.
func (o *StreamSettings) HasFullTextSearchKeys() bool {
	if o != nil && !IsNil(o.FullTextSearchKeys) {
		return true
	}

	return false
}

// SetFullTextSearchKeys gets a reference to the given []string and assigns it to the FullTextSearchKeys field.
func (o *StreamSettings) SetFullTextSearchKeys(v []string) {
	o.FullTextSearchKeys = v
}

// GetIndexFields returns the IndexFields field value if set, zero value otherwise.
func (o *StreamSettings) GetIndexFields() []string {
	if o == nil || IsNil(o.IndexFields) {
		var ret []string
		return ret
	}
	return o.IndexFields
}

// GetIndexFieldsOk returns a tuple with the IndexFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamSettings) GetIndexFieldsOk() ([]string, bool) {
	if o == nil || IsNil(o.IndexFields) {
		return nil, false
	}
	return o.IndexFields, true
}

// HasIndexFields returns a boolean if a field has been set.
func (o *StreamSettings) HasIndexFields() bool {
	if o != nil && !IsNil(o.IndexFields) {
		return true
	}

	return false
}

// SetIndexFields gets a reference to the given []string and assigns it to the IndexFields field.
func (o *StreamSettings) SetIndexFields(v []string) {
	o.IndexFields = v
}

// GetIndexUpdatedAt returns the IndexUpdatedAt field value if set, zero value otherwise.
func (o *StreamSettings) GetIndexUpdatedAt() int64 {
	if o == nil || IsNil(o.IndexUpdatedAt) {
		var ret int64
		return ret
	}
	return *o.IndexUpdatedAt
}

// GetIndexUpdatedAtOk returns a tuple with the IndexUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamSettings) GetIndexUpdatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.IndexUpdatedAt) {
		return nil, false
	}
	return o.IndexUpdatedAt, true
}

// HasIndexUpdatedAt returns a boolean if a field has been set.
func (o *StreamSettings) HasIndexUpdatedAt() bool {
	if o != nil && !IsNil(o.IndexUpdatedAt) {
		return true
	}

	return false
}

// SetIndexUpdatedAt gets a reference to the given int64 and assigns it to the IndexUpdatedAt field.
func (o *StreamSettings) SetIndexUpdatedAt(v int64) {
	o.IndexUpdatedAt = &v
}

// GetMaxQueryRange returns the MaxQueryRange field value if set, zero value otherwise.
func (o *StreamSettings) GetMaxQueryRange() int64 {
	if o == nil || IsNil(o.MaxQueryRange) {
		var ret int64
		return ret
	}
	return *o.MaxQueryRange
}

// GetMaxQueryRangeOk returns a tuple with the MaxQueryRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamSettings) GetMaxQueryRangeOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxQueryRange) {
		return nil, false
	}
	return o.MaxQueryRange, true
}

// HasMaxQueryRange returns a boolean if a field has been set.
func (o *StreamSettings) HasMaxQueryRange() bool {
	if o != nil && !IsNil(o.MaxQueryRange) {
		return true
	}

	return false
}

// SetMaxQueryRange gets a reference to the given int64 and assigns it to the MaxQueryRange field.
func (o *StreamSettings) SetMaxQueryRange(v int64) {
	o.MaxQueryRange = &v
}

// GetPartitionKeys returns the PartitionKeys field value if set, zero value otherwise.
func (o *StreamSettings) GetPartitionKeys() []StreamPartition {
	if o == nil || IsNil(o.PartitionKeys) {
		var ret []StreamPartition
		return ret
	}
	return o.PartitionKeys
}

// GetPartitionKeysOk returns a tuple with the PartitionKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamSettings) GetPartitionKeysOk() ([]StreamPartition, bool) {
	if o == nil || IsNil(o.PartitionKeys) {
		return nil, false
	}
	return o.PartitionKeys, true
}

// HasPartitionKeys returns a boolean if a field has been set.
func (o *StreamSettings) HasPartitionKeys() bool {
	if o != nil && !IsNil(o.PartitionKeys) {
		return true
	}

	return false
}

// SetPartitionKeys gets a reference to the given []StreamPartition and assigns it to the PartitionKeys field.
func (o *StreamSettings) SetPartitionKeys(v []StreamPartition) {
	o.PartitionKeys = v
}

// GetPartitionTimeLevel returns the PartitionTimeLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StreamSettings) GetPartitionTimeLevel() PartitionTimeLevel {
	if o == nil || IsNil(o.PartitionTimeLevel.Get()) {
		var ret PartitionTimeLevel
		return ret
	}
	return *o.PartitionTimeLevel.Get()
}

// GetPartitionTimeLevelOk returns a tuple with the PartitionTimeLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StreamSettings) GetPartitionTimeLevelOk() (*PartitionTimeLevel, bool) {
	if o == nil {
		return nil, false
	}
	return o.PartitionTimeLevel.Get(), o.PartitionTimeLevel.IsSet()
}

// HasPartitionTimeLevel returns a boolean if a field has been set.
func (o *StreamSettings) HasPartitionTimeLevel() bool {
	if o != nil && o.PartitionTimeLevel.IsSet() {
		return true
	}

	return false
}

// SetPartitionTimeLevel gets a reference to the given NullablePartitionTimeLevel and assigns it to the PartitionTimeLevel field.
func (o *StreamSettings) SetPartitionTimeLevel(v PartitionTimeLevel) {
	o.PartitionTimeLevel.Set(&v)
}
// SetPartitionTimeLevelNil sets the value for PartitionTimeLevel to be an explicit nil
func (o *StreamSettings) SetPartitionTimeLevelNil() {
	o.PartitionTimeLevel.Set(nil)
}

// UnsetPartitionTimeLevel ensures that no value is present for PartitionTimeLevel, not even an explicit nil
func (o *StreamSettings) UnsetPartitionTimeLevel() {
	o.PartitionTimeLevel.Unset()
}

// GetStoreOriginalData returns the StoreOriginalData field value if set, zero value otherwise.
func (o *StreamSettings) GetStoreOriginalData() bool {
	if o == nil || IsNil(o.StoreOriginalData) {
		var ret bool
		return ret
	}
	return *o.StoreOriginalData
}

// GetStoreOriginalDataOk returns a tuple with the StoreOriginalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamSettings) GetStoreOriginalDataOk() (*bool, bool) {
	if o == nil || IsNil(o.StoreOriginalData) {
		return nil, false
	}
	return o.StoreOriginalData, true
}

// HasStoreOriginalData returns a boolean if a field has been set.
func (o *StreamSettings) HasStoreOriginalData() bool {
	if o != nil && !IsNil(o.StoreOriginalData) {
		return true
	}

	return false
}

// SetStoreOriginalData gets a reference to the given bool and assigns it to the StoreOriginalData field.
func (o *StreamSettings) SetStoreOriginalData(v bool) {
	o.StoreOriginalData = &v
}

func (o StreamSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StreamSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApproxPartition) {
		toSerialize["approx_partition"] = o.ApproxPartition
	}
	if !IsNil(o.BloomFilterFields) {
		toSerialize["bloom_filter_fields"] = o.BloomFilterFields
	}
	if !IsNil(o.DataRetention) {
		toSerialize["data_retention"] = o.DataRetention
	}
	if o.DefinedSchemaFields != nil {
		toSerialize["defined_schema_fields"] = o.DefinedSchemaFields
	}
	if !IsNil(o.DistinctValueFields) {
		toSerialize["distinct_value_fields"] = o.DistinctValueFields
	}
	if !IsNil(o.ExtendedRetentionDays) {
		toSerialize["extended_retention_days"] = o.ExtendedRetentionDays
	}
	if o.FlattenLevel.IsSet() {
		toSerialize["flatten_level"] = o.FlattenLevel.Get()
	}
	if !IsNil(o.FullTextSearchKeys) {
		toSerialize["full_text_search_keys"] = o.FullTextSearchKeys
	}
	if !IsNil(o.IndexFields) {
		toSerialize["index_fields"] = o.IndexFields
	}
	if !IsNil(o.IndexUpdatedAt) {
		toSerialize["index_updated_at"] = o.IndexUpdatedAt
	}
	if !IsNil(o.MaxQueryRange) {
		toSerialize["max_query_range"] = o.MaxQueryRange
	}
	if !IsNil(o.PartitionKeys) {
		toSerialize["partition_keys"] = o.PartitionKeys
	}
	if o.PartitionTimeLevel.IsSet() {
		toSerialize["partition_time_level"] = o.PartitionTimeLevel.Get()
	}
	if !IsNil(o.StoreOriginalData) {
		toSerialize["store_original_data"] = o.StoreOriginalData
	}
	return toSerialize, nil
}

type NullableStreamSettings struct {
	value *StreamSettings
	isSet bool
}

func (v NullableStreamSettings) Get() *StreamSettings {
	return v.value
}

func (v *NullableStreamSettings) Set(val *StreamSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamSettings(val *StreamSettings) *NullableStreamSettings {
	return &NullableStreamSettings{value: val, isSet: true}
}

func (v NullableStreamSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


