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

// checks if the ListAlertsResponseBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAlertsResponseBody{}

// ListAlertsResponseBody HTTP response body for `ListAlerts` endpoint.
type ListAlertsResponseBody struct {
	List []ListAlertsResponseBodyItem `json:"list"`
}

type _ListAlertsResponseBody ListAlertsResponseBody

// NewListAlertsResponseBody instantiates a new ListAlertsResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAlertsResponseBody(list []ListAlertsResponseBodyItem) *ListAlertsResponseBody {
	this := ListAlertsResponseBody{}
	this.List = list
	return &this
}

// NewListAlertsResponseBodyWithDefaults instantiates a new ListAlertsResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAlertsResponseBodyWithDefaults() *ListAlertsResponseBody {
	this := ListAlertsResponseBody{}
	return &this
}

// GetList returns the List field value
func (o *ListAlertsResponseBody) GetList() []ListAlertsResponseBodyItem {
	if o == nil {
		var ret []ListAlertsResponseBodyItem
		return ret
	}

	return o.List
}

// GetListOk returns a tuple with the List field value
// and a boolean to check if the value has been set.
func (o *ListAlertsResponseBody) GetListOk() ([]ListAlertsResponseBodyItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.List, true
}

// SetList sets field value
func (o *ListAlertsResponseBody) SetList(v []ListAlertsResponseBodyItem) {
	o.List = v
}

func (o ListAlertsResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAlertsResponseBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["list"] = o.List
	return toSerialize, nil
}

func (o *ListAlertsResponseBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"list",
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

	varListAlertsResponseBody := _ListAlertsResponseBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListAlertsResponseBody)

	if err != nil {
		return err
	}

	*o = ListAlertsResponseBody(varListAlertsResponseBody)

	return err
}

type NullableListAlertsResponseBody struct {
	value *ListAlertsResponseBody
	isSet bool
}

func (v NullableListAlertsResponseBody) Get() *ListAlertsResponseBody {
	return v.value
}

func (v *NullableListAlertsResponseBody) Set(val *ListAlertsResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableListAlertsResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableListAlertsResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAlertsResponseBody(val *ListAlertsResponseBody) *NullableListAlertsResponseBody {
	return &NullableListAlertsResponseBody{value: val, isSet: true}
}

func (v NullableListAlertsResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAlertsResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


