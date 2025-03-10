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

// checks if the TimedAnnotationDelete type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimedAnnotationDelete{}

// TimedAnnotationDelete struct for TimedAnnotationDelete
type TimedAnnotationDelete struct {
	AnnotationIds []string `json:"annotation_ids"`
}

type _TimedAnnotationDelete TimedAnnotationDelete

// NewTimedAnnotationDelete instantiates a new TimedAnnotationDelete object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimedAnnotationDelete(annotationIds []string) *TimedAnnotationDelete {
	this := TimedAnnotationDelete{}
	this.AnnotationIds = annotationIds
	return &this
}

// NewTimedAnnotationDeleteWithDefaults instantiates a new TimedAnnotationDelete object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimedAnnotationDeleteWithDefaults() *TimedAnnotationDelete {
	this := TimedAnnotationDelete{}
	return &this
}

// GetAnnotationIds returns the AnnotationIds field value
func (o *TimedAnnotationDelete) GetAnnotationIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AnnotationIds
}

// GetAnnotationIdsOk returns a tuple with the AnnotationIds field value
// and a boolean to check if the value has been set.
func (o *TimedAnnotationDelete) GetAnnotationIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AnnotationIds, true
}

// SetAnnotationIds sets field value
func (o *TimedAnnotationDelete) SetAnnotationIds(v []string) {
	o.AnnotationIds = v
}

func (o TimedAnnotationDelete) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimedAnnotationDelete) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["annotation_ids"] = o.AnnotationIds
	return toSerialize, nil
}

func (o *TimedAnnotationDelete) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"annotation_ids",
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

	varTimedAnnotationDelete := _TimedAnnotationDelete{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTimedAnnotationDelete)

	if err != nil {
		return err
	}

	*o = TimedAnnotationDelete(varTimedAnnotationDelete)

	return err
}

type NullableTimedAnnotationDelete struct {
	value *TimedAnnotationDelete
	isSet bool
}

func (v NullableTimedAnnotationDelete) Get() *TimedAnnotationDelete {
	return v.value
}

func (v *NullableTimedAnnotationDelete) Set(val *TimedAnnotationDelete) {
	v.value = val
	v.isSet = true
}

func (v NullableTimedAnnotationDelete) IsSet() bool {
	return v.isSet
}

func (v *NullableTimedAnnotationDelete) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimedAnnotationDelete(val *TimedAnnotationDelete) *NullableTimedAnnotationDelete {
	return &NullableTimedAnnotationDelete{value: val, isSet: true}
}

func (v NullableTimedAnnotationDelete) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimedAnnotationDelete) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


