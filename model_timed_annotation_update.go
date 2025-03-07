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

// checks if the TimedAnnotationUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimedAnnotationUpdate{}

// TimedAnnotationUpdate struct for TimedAnnotationUpdate
type TimedAnnotationUpdate struct {
	EndTime NullableInt64 `json:"end_time,omitempty"`
	Panels []string `json:"panels,omitempty"`
	StartTime NullableInt64 `json:"start_time,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Text NullableString `json:"text,omitempty"`
	Title NullableString `json:"title,omitempty"`
}

// NewTimedAnnotationUpdate instantiates a new TimedAnnotationUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimedAnnotationUpdate() *TimedAnnotationUpdate {
	this := TimedAnnotationUpdate{}
	return &this
}

// NewTimedAnnotationUpdateWithDefaults instantiates a new TimedAnnotationUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimedAnnotationUpdateWithDefaults() *TimedAnnotationUpdate {
	this := TimedAnnotationUpdate{}
	return &this
}

// GetEndTime returns the EndTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimedAnnotationUpdate) GetEndTime() int64 {
	if o == nil || IsNil(o.EndTime.Get()) {
		var ret int64
		return ret
	}
	return *o.EndTime.Get()
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimedAnnotationUpdate) GetEndTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndTime.Get(), o.EndTime.IsSet()
}

// HasEndTime returns a boolean if a field has been set.
func (o *TimedAnnotationUpdate) HasEndTime() bool {
	if o != nil && o.EndTime.IsSet() {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given NullableInt64 and assigns it to the EndTime field.
func (o *TimedAnnotationUpdate) SetEndTime(v int64) {
	o.EndTime.Set(&v)
}
// SetEndTimeNil sets the value for EndTime to be an explicit nil
func (o *TimedAnnotationUpdate) SetEndTimeNil() {
	o.EndTime.Set(nil)
}

// UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
func (o *TimedAnnotationUpdate) UnsetEndTime() {
	o.EndTime.Unset()
}

// GetPanels returns the Panels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimedAnnotationUpdate) GetPanels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Panels
}

// GetPanelsOk returns a tuple with the Panels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimedAnnotationUpdate) GetPanelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Panels) {
		return nil, false
	}
	return o.Panels, true
}

// HasPanels returns a boolean if a field has been set.
func (o *TimedAnnotationUpdate) HasPanels() bool {
	if o != nil && !IsNil(o.Panels) {
		return true
	}

	return false
}

// SetPanels gets a reference to the given []string and assigns it to the Panels field.
func (o *TimedAnnotationUpdate) SetPanels(v []string) {
	o.Panels = v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimedAnnotationUpdate) GetStartTime() int64 {
	if o == nil || IsNil(o.StartTime.Get()) {
		var ret int64
		return ret
	}
	return *o.StartTime.Get()
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimedAnnotationUpdate) GetStartTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartTime.Get(), o.StartTime.IsSet()
}

// HasStartTime returns a boolean if a field has been set.
func (o *TimedAnnotationUpdate) HasStartTime() bool {
	if o != nil && o.StartTime.IsSet() {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given NullableInt64 and assigns it to the StartTime field.
func (o *TimedAnnotationUpdate) SetStartTime(v int64) {
	o.StartTime.Set(&v)
}
// SetStartTimeNil sets the value for StartTime to be an explicit nil
func (o *TimedAnnotationUpdate) SetStartTimeNil() {
	o.StartTime.Set(nil)
}

// UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
func (o *TimedAnnotationUpdate) UnsetStartTime() {
	o.StartTime.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimedAnnotationUpdate) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimedAnnotationUpdate) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TimedAnnotationUpdate) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *TimedAnnotationUpdate) SetTags(v []string) {
	o.Tags = v
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimedAnnotationUpdate) GetText() string {
	if o == nil || IsNil(o.Text.Get()) {
		var ret string
		return ret
	}
	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimedAnnotationUpdate) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// HasText returns a boolean if a field has been set.
func (o *TimedAnnotationUpdate) HasText() bool {
	if o != nil && o.Text.IsSet() {
		return true
	}

	return false
}

// SetText gets a reference to the given NullableString and assigns it to the Text field.
func (o *TimedAnnotationUpdate) SetText(v string) {
	o.Text.Set(&v)
}
// SetTextNil sets the value for Text to be an explicit nil
func (o *TimedAnnotationUpdate) SetTextNil() {
	o.Text.Set(nil)
}

// UnsetText ensures that no value is present for Text, not even an explicit nil
func (o *TimedAnnotationUpdate) UnsetText() {
	o.Text.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimedAnnotationUpdate) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimedAnnotationUpdate) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *TimedAnnotationUpdate) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *TimedAnnotationUpdate) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *TimedAnnotationUpdate) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *TimedAnnotationUpdate) UnsetTitle() {
	o.Title.Unset()
}

func (o TimedAnnotationUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimedAnnotationUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.EndTime.IsSet() {
		toSerialize["end_time"] = o.EndTime.Get()
	}
	if o.Panels != nil {
		toSerialize["panels"] = o.Panels
	}
	if o.StartTime.IsSet() {
		toSerialize["start_time"] = o.StartTime.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Text.IsSet() {
		toSerialize["text"] = o.Text.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	return toSerialize, nil
}

type NullableTimedAnnotationUpdate struct {
	value *TimedAnnotationUpdate
	isSet bool
}

func (v NullableTimedAnnotationUpdate) Get() *TimedAnnotationUpdate {
	return v.value
}

func (v *NullableTimedAnnotationUpdate) Set(val *TimedAnnotationUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableTimedAnnotationUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableTimedAnnotationUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimedAnnotationUpdate(val *TimedAnnotationUpdate) *NullableTimedAnnotationUpdate {
	return &NullableTimedAnnotationUpdate{value: val, isSet: true}
}

func (v NullableTimedAnnotationUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimedAnnotationUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


