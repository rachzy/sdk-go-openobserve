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

// checks if the ListDashboardsResponseBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDashboardsResponseBody{}

// ListDashboardsResponseBody HTTP response body for `ListDashboards` endpoint.
type ListDashboardsResponseBody struct {
	Dashboards []ListDashboardsResponseBodyItem `json:"dashboards"`
}

type _ListDashboardsResponseBody ListDashboardsResponseBody

// NewListDashboardsResponseBody instantiates a new ListDashboardsResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDashboardsResponseBody(dashboards []ListDashboardsResponseBodyItem) *ListDashboardsResponseBody {
	this := ListDashboardsResponseBody{}
	this.Dashboards = dashboards
	return &this
}

// NewListDashboardsResponseBodyWithDefaults instantiates a new ListDashboardsResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDashboardsResponseBodyWithDefaults() *ListDashboardsResponseBody {
	this := ListDashboardsResponseBody{}
	return &this
}

// GetDashboards returns the Dashboards field value
func (o *ListDashboardsResponseBody) GetDashboards() []ListDashboardsResponseBodyItem {
	if o == nil {
		var ret []ListDashboardsResponseBodyItem
		return ret
	}

	return o.Dashboards
}

// GetDashboardsOk returns a tuple with the Dashboards field value
// and a boolean to check if the value has been set.
func (o *ListDashboardsResponseBody) GetDashboardsOk() ([]ListDashboardsResponseBodyItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dashboards, true
}

// SetDashboards sets field value
func (o *ListDashboardsResponseBody) SetDashboards(v []ListDashboardsResponseBodyItem) {
	o.Dashboards = v
}

func (o ListDashboardsResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListDashboardsResponseBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dashboards"] = o.Dashboards
	return toSerialize, nil
}

func (o *ListDashboardsResponseBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dashboards",
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

	varListDashboardsResponseBody := _ListDashboardsResponseBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListDashboardsResponseBody)

	if err != nil {
		return err
	}

	*o = ListDashboardsResponseBody(varListDashboardsResponseBody)

	return err
}

type NullableListDashboardsResponseBody struct {
	value *ListDashboardsResponseBody
	isSet bool
}

func (v NullableListDashboardsResponseBody) Get() *ListDashboardsResponseBody {
	return v.value
}

func (v *NullableListDashboardsResponseBody) Set(val *ListDashboardsResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableListDashboardsResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableListDashboardsResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDashboardsResponseBody(val *ListDashboardsResponseBody) *NullableListDashboardsResponseBody {
	return &NullableListDashboardsResponseBody{value: val, isSet: true}
}

func (v NullableListDashboardsResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDashboardsResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


