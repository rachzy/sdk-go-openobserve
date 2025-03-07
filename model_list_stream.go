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

// checks if the ListStream type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListStream{}

// ListStream struct for ListStream
type ListStream struct {
	List []Stream `json:"list"`
	Total int32 `json:"total"`
}

type _ListStream ListStream

// NewListStream instantiates a new ListStream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListStream(list []Stream, total int32) *ListStream {
	this := ListStream{}
	this.List = list
	this.Total = total
	return &this
}

// NewListStreamWithDefaults instantiates a new ListStream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListStreamWithDefaults() *ListStream {
	this := ListStream{}
	return &this
}

// GetList returns the List field value
func (o *ListStream) GetList() []Stream {
	if o == nil {
		var ret []Stream
		return ret
	}

	return o.List
}

// GetListOk returns a tuple with the List field value
// and a boolean to check if the value has been set.
func (o *ListStream) GetListOk() ([]Stream, bool) {
	if o == nil {
		return nil, false
	}
	return o.List, true
}

// SetList sets field value
func (o *ListStream) SetList(v []Stream) {
	o.List = v
}

// GetTotal returns the Total field value
func (o *ListStream) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ListStream) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ListStream) SetTotal(v int32) {
	o.Total = v
}

func (o ListStream) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListStream) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["list"] = o.List
	toSerialize["total"] = o.Total
	return toSerialize, nil
}

func (o *ListStream) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"list",
		"total",
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

	varListStream := _ListStream{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListStream)

	if err != nil {
		return err
	}

	*o = ListStream(varListStream)

	return err
}

type NullableListStream struct {
	value *ListStream
	isSet bool
}

func (v NullableListStream) Get() *ListStream {
	return v.value
}

func (v *NullableListStream) Set(val *ListStream) {
	v.value = val
	v.isSet = true
}

func (v NullableListStream) IsSet() bool {
	return v.isSet
}

func (v *NullableListStream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListStream(val *ListStream) *NullableListStream {
	return &NullableListStream{value: val, isSet: true}
}

func (v NullableListStream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListStream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


