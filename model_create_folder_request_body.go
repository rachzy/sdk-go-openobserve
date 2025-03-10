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

// checks if the CreateFolderRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFolderRequestBody{}

// CreateFolderRequestBody HTTP request body for `CreateFolder` endpoint.
type CreateFolderRequestBody struct {
	Description string `json:"description"`
	Name string `json:"name"`
}

type _CreateFolderRequestBody CreateFolderRequestBody

// NewCreateFolderRequestBody instantiates a new CreateFolderRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFolderRequestBody(description string, name string) *CreateFolderRequestBody {
	this := CreateFolderRequestBody{}
	this.Description = description
	this.Name = name
	return &this
}

// NewCreateFolderRequestBodyWithDefaults instantiates a new CreateFolderRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFolderRequestBodyWithDefaults() *CreateFolderRequestBody {
	this := CreateFolderRequestBody{}
	return &this
}

// GetDescription returns the Description field value
func (o *CreateFolderRequestBody) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateFolderRequestBody) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateFolderRequestBody) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value
func (o *CreateFolderRequestBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateFolderRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateFolderRequestBody) SetName(v string) {
	o.Name = v
}

func (o CreateFolderRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFolderRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *CreateFolderRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"name",
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

	varCreateFolderRequestBody := _CreateFolderRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateFolderRequestBody)

	if err != nil {
		return err
	}

	*o = CreateFolderRequestBody(varCreateFolderRequestBody)

	return err
}

type NullableCreateFolderRequestBody struct {
	value *CreateFolderRequestBody
	isSet bool
}

func (v NullableCreateFolderRequestBody) Get() *CreateFolderRequestBody {
	return v.value
}

func (v *NullableCreateFolderRequestBody) Set(val *CreateFolderRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFolderRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFolderRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFolderRequestBody(val *CreateFolderRequestBody) *NullableCreateFolderRequestBody {
	return &NullableCreateFolderRequestBody{value: val, isSet: true}
}

func (v NullableCreateFolderRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFolderRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


