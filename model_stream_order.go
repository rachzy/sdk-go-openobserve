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

// checks if the StreamOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StreamOrder{}

// StreamOrder struct for StreamOrder
type StreamOrder struct {
	ApplyBeforeFlattening *bool `json:"applyBeforeFlattening,omitempty"`
	IsRemoved *bool `json:"isRemoved,omitempty"`
	Order *int32 `json:"order,omitempty"`
	Stream *string `json:"stream,omitempty"`
	StreamType *StreamType `json:"streamType,omitempty"`
}

// NewStreamOrder instantiates a new StreamOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamOrder() *StreamOrder {
	this := StreamOrder{}
	return &this
}

// NewStreamOrderWithDefaults instantiates a new StreamOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamOrderWithDefaults() *StreamOrder {
	this := StreamOrder{}
	return &this
}

// GetApplyBeforeFlattening returns the ApplyBeforeFlattening field value if set, zero value otherwise.
func (o *StreamOrder) GetApplyBeforeFlattening() bool {
	if o == nil || IsNil(o.ApplyBeforeFlattening) {
		var ret bool
		return ret
	}
	return *o.ApplyBeforeFlattening
}

// GetApplyBeforeFlatteningOk returns a tuple with the ApplyBeforeFlattening field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamOrder) GetApplyBeforeFlatteningOk() (*bool, bool) {
	if o == nil || IsNil(o.ApplyBeforeFlattening) {
		return nil, false
	}
	return o.ApplyBeforeFlattening, true
}

// HasApplyBeforeFlattening returns a boolean if a field has been set.
func (o *StreamOrder) HasApplyBeforeFlattening() bool {
	if o != nil && !IsNil(o.ApplyBeforeFlattening) {
		return true
	}

	return false
}

// SetApplyBeforeFlattening gets a reference to the given bool and assigns it to the ApplyBeforeFlattening field.
func (o *StreamOrder) SetApplyBeforeFlattening(v bool) {
	o.ApplyBeforeFlattening = &v
}

// GetIsRemoved returns the IsRemoved field value if set, zero value otherwise.
func (o *StreamOrder) GetIsRemoved() bool {
	if o == nil || IsNil(o.IsRemoved) {
		var ret bool
		return ret
	}
	return *o.IsRemoved
}

// GetIsRemovedOk returns a tuple with the IsRemoved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamOrder) GetIsRemovedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRemoved) {
		return nil, false
	}
	return o.IsRemoved, true
}

// HasIsRemoved returns a boolean if a field has been set.
func (o *StreamOrder) HasIsRemoved() bool {
	if o != nil && !IsNil(o.IsRemoved) {
		return true
	}

	return false
}

// SetIsRemoved gets a reference to the given bool and assigns it to the IsRemoved field.
func (o *StreamOrder) SetIsRemoved(v bool) {
	o.IsRemoved = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *StreamOrder) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamOrder) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *StreamOrder) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *StreamOrder) SetOrder(v int32) {
	o.Order = &v
}

// GetStream returns the Stream field value if set, zero value otherwise.
func (o *StreamOrder) GetStream() string {
	if o == nil || IsNil(o.Stream) {
		var ret string
		return ret
	}
	return *o.Stream
}

// GetStreamOk returns a tuple with the Stream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamOrder) GetStreamOk() (*string, bool) {
	if o == nil || IsNil(o.Stream) {
		return nil, false
	}
	return o.Stream, true
}

// HasStream returns a boolean if a field has been set.
func (o *StreamOrder) HasStream() bool {
	if o != nil && !IsNil(o.Stream) {
		return true
	}

	return false
}

// SetStream gets a reference to the given string and assigns it to the Stream field.
func (o *StreamOrder) SetStream(v string) {
	o.Stream = &v
}

// GetStreamType returns the StreamType field value if set, zero value otherwise.
func (o *StreamOrder) GetStreamType() StreamType {
	if o == nil || IsNil(o.StreamType) {
		var ret StreamType
		return ret
	}
	return *o.StreamType
}

// GetStreamTypeOk returns a tuple with the StreamType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamOrder) GetStreamTypeOk() (*StreamType, bool) {
	if o == nil || IsNil(o.StreamType) {
		return nil, false
	}
	return o.StreamType, true
}

// HasStreamType returns a boolean if a field has been set.
func (o *StreamOrder) HasStreamType() bool {
	if o != nil && !IsNil(o.StreamType) {
		return true
	}

	return false
}

// SetStreamType gets a reference to the given StreamType and assigns it to the StreamType field.
func (o *StreamOrder) SetStreamType(v StreamType) {
	o.StreamType = &v
}

func (o StreamOrder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StreamOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplyBeforeFlattening) {
		toSerialize["applyBeforeFlattening"] = o.ApplyBeforeFlattening
	}
	if !IsNil(o.IsRemoved) {
		toSerialize["isRemoved"] = o.IsRemoved
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.Stream) {
		toSerialize["stream"] = o.Stream
	}
	if !IsNil(o.StreamType) {
		toSerialize["streamType"] = o.StreamType
	}
	return toSerialize, nil
}

type NullableStreamOrder struct {
	value *StreamOrder
	isSet bool
}

func (v NullableStreamOrder) Get() *StreamOrder {
	return v.value
}

func (v *NullableStreamOrder) Set(val *StreamOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamOrder(val *StreamOrder) *NullableStreamOrder {
	return &NullableStreamOrder{value: val, isSet: true}
}

func (v NullableStreamOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


