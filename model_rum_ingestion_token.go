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

// checks if the RumIngestionToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RumIngestionToken{}

// RumIngestionToken struct for RumIngestionToken
type RumIngestionToken struct {
	RumToken NullableString `json:"rum_token,omitempty"`
	User string `json:"user"`
}

type _RumIngestionToken RumIngestionToken

// NewRumIngestionToken instantiates a new RumIngestionToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRumIngestionToken(user string) *RumIngestionToken {
	this := RumIngestionToken{}
	this.User = user
	return &this
}

// NewRumIngestionTokenWithDefaults instantiates a new RumIngestionToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRumIngestionTokenWithDefaults() *RumIngestionToken {
	this := RumIngestionToken{}
	return &this
}

// GetRumToken returns the RumToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RumIngestionToken) GetRumToken() string {
	if o == nil || IsNil(o.RumToken.Get()) {
		var ret string
		return ret
	}
	return *o.RumToken.Get()
}

// GetRumTokenOk returns a tuple with the RumToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RumIngestionToken) GetRumTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RumToken.Get(), o.RumToken.IsSet()
}

// HasRumToken returns a boolean if a field has been set.
func (o *RumIngestionToken) HasRumToken() bool {
	if o != nil && o.RumToken.IsSet() {
		return true
	}

	return false
}

// SetRumToken gets a reference to the given NullableString and assigns it to the RumToken field.
func (o *RumIngestionToken) SetRumToken(v string) {
	o.RumToken.Set(&v)
}
// SetRumTokenNil sets the value for RumToken to be an explicit nil
func (o *RumIngestionToken) SetRumTokenNil() {
	o.RumToken.Set(nil)
}

// UnsetRumToken ensures that no value is present for RumToken, not even an explicit nil
func (o *RumIngestionToken) UnsetRumToken() {
	o.RumToken.Unset()
}

// GetUser returns the User field value
func (o *RumIngestionToken) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *RumIngestionToken) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *RumIngestionToken) SetUser(v string) {
	o.User = v
}

func (o RumIngestionToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RumIngestionToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RumToken.IsSet() {
		toSerialize["rum_token"] = o.RumToken.Get()
	}
	toSerialize["user"] = o.User
	return toSerialize, nil
}

func (o *RumIngestionToken) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user",
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

	varRumIngestionToken := _RumIngestionToken{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRumIngestionToken)

	if err != nil {
		return err
	}

	*o = RumIngestionToken(varRumIngestionToken)

	return err
}

type NullableRumIngestionToken struct {
	value *RumIngestionToken
	isSet bool
}

func (v NullableRumIngestionToken) Get() *RumIngestionToken {
	return v.value
}

func (v *NullableRumIngestionToken) Set(val *RumIngestionToken) {
	v.value = val
	v.isSet = true
}

func (v NullableRumIngestionToken) IsSet() bool {
	return v.isSet
}

func (v *NullableRumIngestionToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRumIngestionToken(val *RumIngestionToken) *NullableRumIngestionToken {
	return &NullableRumIngestionToken{value: val, isSet: true}
}

func (v NullableRumIngestionToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRumIngestionToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


