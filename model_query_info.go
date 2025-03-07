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
	"bytes"
	"fmt"
)

// checks if the QueryInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryInfo{}

// QueryInfo struct for QueryInfo
type QueryInfo struct {
	EndTime int64 `json:"end_time"`
	Sql string `json:"sql"`
	StartTime int64 `json:"start_time"`
}

type _QueryInfo QueryInfo

// NewQueryInfo instantiates a new QueryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryInfo(endTime int64, sql string, startTime int64) *QueryInfo {
	this := QueryInfo{}
	this.EndTime = endTime
	this.Sql = sql
	this.StartTime = startTime
	return &this
}

// NewQueryInfoWithDefaults instantiates a new QueryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryInfoWithDefaults() *QueryInfo {
	this := QueryInfo{}
	return &this
}

// GetEndTime returns the EndTime field value
func (o *QueryInfo) GetEndTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *QueryInfo) GetEndTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *QueryInfo) SetEndTime(v int64) {
	o.EndTime = v
}

// GetSql returns the Sql field value
func (o *QueryInfo) GetSql() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sql
}

// GetSqlOk returns a tuple with the Sql field value
// and a boolean to check if the value has been set.
func (o *QueryInfo) GetSqlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sql, true
}

// SetSql sets field value
func (o *QueryInfo) SetSql(v string) {
	o.Sql = v
}

// GetStartTime returns the StartTime field value
func (o *QueryInfo) GetStartTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *QueryInfo) GetStartTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *QueryInfo) SetStartTime(v int64) {
	o.StartTime = v
}

func (o QueryInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["end_time"] = o.EndTime
	toSerialize["sql"] = o.Sql
	toSerialize["start_time"] = o.StartTime
	return toSerialize, nil
}

func (o *QueryInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"end_time",
		"sql",
		"start_time",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varQueryInfo := _QueryInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQueryInfo)

	if err != nil {
		return err
	}

	*o = QueryInfo(varQueryInfo)

	return err
}

type NullableQueryInfo struct {
	value *QueryInfo
	isSet bool
}

func (v NullableQueryInfo) Get() *QueryInfo {
	return v.value
}

func (v *NullableQueryInfo) Set(val *QueryInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryInfo(val *QueryInfo) *NullableQueryInfo {
	return &NullableQueryInfo{value: val, isSet: true}
}

func (v NullableQueryInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


