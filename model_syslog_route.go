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

// checks if the SyslogRoute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyslogRoute{}

// SyslogRoute struct for SyslogRoute
type SyslogRoute struct {
	Id *string `json:"id,omitempty"`
	OrgId *string `json:"orgId,omitempty"`
	StreamName *string `json:"streamName,omitempty"`
	Subnets []string `json:"subnets,omitempty"`
}

// NewSyslogRoute instantiates a new SyslogRoute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyslogRoute() *SyslogRoute {
	this := SyslogRoute{}
	return &this
}

// NewSyslogRouteWithDefaults instantiates a new SyslogRoute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyslogRouteWithDefaults() *SyslogRoute {
	this := SyslogRoute{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SyslogRoute) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogRoute) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SyslogRoute) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SyslogRoute) SetId(v string) {
	o.Id = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *SyslogRoute) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogRoute) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *SyslogRoute) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *SyslogRoute) SetOrgId(v string) {
	o.OrgId = &v
}

// GetStreamName returns the StreamName field value if set, zero value otherwise.
func (o *SyslogRoute) GetStreamName() string {
	if o == nil || IsNil(o.StreamName) {
		var ret string
		return ret
	}
	return *o.StreamName
}

// GetStreamNameOk returns a tuple with the StreamName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogRoute) GetStreamNameOk() (*string, bool) {
	if o == nil || IsNil(o.StreamName) {
		return nil, false
	}
	return o.StreamName, true
}

// HasStreamName returns a boolean if a field has been set.
func (o *SyslogRoute) HasStreamName() bool {
	if o != nil && !IsNil(o.StreamName) {
		return true
	}

	return false
}

// SetStreamName gets a reference to the given string and assigns it to the StreamName field.
func (o *SyslogRoute) SetStreamName(v string) {
	o.StreamName = &v
}

// GetSubnets returns the Subnets field value if set, zero value otherwise.
func (o *SyslogRoute) GetSubnets() []string {
	if o == nil || IsNil(o.Subnets) {
		var ret []string
		return ret
	}
	return o.Subnets
}

// GetSubnetsOk returns a tuple with the Subnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogRoute) GetSubnetsOk() ([]string, bool) {
	if o == nil || IsNil(o.Subnets) {
		return nil, false
	}
	return o.Subnets, true
}

// HasSubnets returns a boolean if a field has been set.
func (o *SyslogRoute) HasSubnets() bool {
	if o != nil && !IsNil(o.Subnets) {
		return true
	}

	return false
}

// SetSubnets gets a reference to the given []string and assigns it to the Subnets field.
func (o *SyslogRoute) SetSubnets(v []string) {
	o.Subnets = v
}

func (o SyslogRoute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyslogRoute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.OrgId) {
		toSerialize["orgId"] = o.OrgId
	}
	if !IsNil(o.StreamName) {
		toSerialize["streamName"] = o.StreamName
	}
	if !IsNil(o.Subnets) {
		toSerialize["subnets"] = o.Subnets
	}
	return toSerialize, nil
}

type NullableSyslogRoute struct {
	value *SyslogRoute
	isSet bool
}

func (v NullableSyslogRoute) Get() *SyslogRoute {
	return v.value
}

func (v *NullableSyslogRoute) Set(val *SyslogRoute) {
	v.value = val
	v.isSet = true
}

func (v NullableSyslogRoute) IsSet() bool {
	return v.isSet
}

func (v *NullableSyslogRoute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyslogRoute(val *SyslogRoute) *NullableSyslogRoute {
	return &NullableSyslogRoute{value: val, isSet: true}
}

func (v NullableSyslogRoute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyslogRoute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


