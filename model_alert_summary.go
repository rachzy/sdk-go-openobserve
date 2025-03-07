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

// checks if the AlertSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertSummary{}

// AlertSummary struct for AlertSummary
type AlertSummary struct {
	NumRealtime int64 `json:"num_realtime"`
	NumScheduled int64 `json:"num_scheduled"`
}

type _AlertSummary AlertSummary

// NewAlertSummary instantiates a new AlertSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertSummary(numRealtime int64, numScheduled int64) *AlertSummary {
	this := AlertSummary{}
	this.NumRealtime = numRealtime
	this.NumScheduled = numScheduled
	return &this
}

// NewAlertSummaryWithDefaults instantiates a new AlertSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertSummaryWithDefaults() *AlertSummary {
	this := AlertSummary{}
	return &this
}

// GetNumRealtime returns the NumRealtime field value
func (o *AlertSummary) GetNumRealtime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NumRealtime
}

// GetNumRealtimeOk returns a tuple with the NumRealtime field value
// and a boolean to check if the value has been set.
func (o *AlertSummary) GetNumRealtimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumRealtime, true
}

// SetNumRealtime sets field value
func (o *AlertSummary) SetNumRealtime(v int64) {
	o.NumRealtime = v
}

// GetNumScheduled returns the NumScheduled field value
func (o *AlertSummary) GetNumScheduled() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NumScheduled
}

// GetNumScheduledOk returns a tuple with the NumScheduled field value
// and a boolean to check if the value has been set.
func (o *AlertSummary) GetNumScheduledOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumScheduled, true
}

// SetNumScheduled sets field value
func (o *AlertSummary) SetNumScheduled(v int64) {
	o.NumScheduled = v
}

func (o AlertSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["num_realtime"] = o.NumRealtime
	toSerialize["num_scheduled"] = o.NumScheduled
	return toSerialize, nil
}

func (o *AlertSummary) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"num_realtime",
		"num_scheduled",
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

	varAlertSummary := _AlertSummary{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlertSummary)

	if err != nil {
		return err
	}

	*o = AlertSummary(varAlertSummary)

	return err
}

type NullableAlertSummary struct {
	value *AlertSummary
	isSet bool
}

func (v NullableAlertSummary) Get() *AlertSummary {
	return v.value
}

func (v *NullableAlertSummary) Set(val *AlertSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertSummary(val *AlertSummary) *NullableAlertSummary {
	return &NullableAlertSummary{value: val, isSet: true}
}

func (v NullableAlertSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


