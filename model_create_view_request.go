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

// checks if the CreateViewRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateViewRequest{}

// CreateViewRequest struct for CreateViewRequest
type CreateViewRequest struct {
	// Base64 encoded string, containing all the data for a given view. This data is expected to be versioned so that the frontend can deserialize as required.
	Data interface{} `json:"data"`
	// User-readable name of the view, doesn't need to be unique.
	ViewName string `json:"view_name"`
}

type _CreateViewRequest CreateViewRequest

// NewCreateViewRequest instantiates a new CreateViewRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateViewRequest(data interface{}, viewName string) *CreateViewRequest {
	this := CreateViewRequest{}
	this.Data = data
	this.ViewName = viewName
	return &this
}

// NewCreateViewRequestWithDefaults instantiates a new CreateViewRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateViewRequestWithDefaults() *CreateViewRequest {
	this := CreateViewRequest{}
	return &this
}

// GetData returns the Data field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CreateViewRequest) GetData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateViewRequest) GetDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CreateViewRequest) SetData(v interface{}) {
	o.Data = v
}

// GetViewName returns the ViewName field value
func (o *CreateViewRequest) GetViewName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ViewName
}

// GetViewNameOk returns a tuple with the ViewName field value
// and a boolean to check if the value has been set.
func (o *CreateViewRequest) GetViewNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewName, true
}

// SetViewName sets field value
func (o *CreateViewRequest) SetViewName(v string) {
	o.ViewName = v
}

func (o CreateViewRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateViewRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	toSerialize["view_name"] = o.ViewName
	return toSerialize, nil
}

func (o *CreateViewRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"view_name",
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

	varCreateViewRequest := _CreateViewRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateViewRequest)

	if err != nil {
		return err
	}

	*o = CreateViewRequest(varCreateViewRequest)

	return err
}

type NullableCreateViewRequest struct {
	value *CreateViewRequest
	isSet bool
}

func (v NullableCreateViewRequest) Get() *CreateViewRequest {
	return v.value
}

func (v *NullableCreateViewRequest) Set(val *CreateViewRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateViewRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateViewRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateViewRequest(val *CreateViewRequest) *NullableCreateViewRequest {
	return &NullableCreateViewRequest{value: val, isSet: true}
}

func (v NullableCreateViewRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateViewRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


