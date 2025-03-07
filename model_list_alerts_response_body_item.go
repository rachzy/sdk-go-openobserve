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

// checks if the ListAlertsResponseBodyItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAlertsResponseBodyItem{}

// ListAlertsResponseBodyItem An item in the list returned by the `ListDashboards` endpoint.
type ListAlertsResponseBodyItem struct {
	AlertId Ksuid `json:"alert_id"`
	Condition QueryCondition `json:"condition"`
	Description NullableString `json:"description,omitempty"`
	Enabled bool `json:"enabled"`
	FolderId string `json:"folder_id"`
	FolderName string `json:"folder_name"`
	LastSatisfiedAt NullableInt64 `json:"last_satisfied_at,omitempty"`
	LastTriggeredAt NullableInt64 `json:"last_triggered_at,omitempty"`
	Name string `json:"name"`
	Owner NullableString `json:"owner,omitempty"`
}

type _ListAlertsResponseBodyItem ListAlertsResponseBodyItem

// NewListAlertsResponseBodyItem instantiates a new ListAlertsResponseBodyItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAlertsResponseBodyItem(alertId Ksuid, condition QueryCondition, enabled bool, folderId string, folderName string, name string) *ListAlertsResponseBodyItem {
	this := ListAlertsResponseBodyItem{}
	this.AlertId = alertId
	this.Condition = condition
	this.Enabled = enabled
	this.FolderId = folderId
	this.FolderName = folderName
	this.Name = name
	return &this
}

// NewListAlertsResponseBodyItemWithDefaults instantiates a new ListAlertsResponseBodyItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAlertsResponseBodyItemWithDefaults() *ListAlertsResponseBodyItem {
	this := ListAlertsResponseBodyItem{}
	return &this
}

// GetAlertId returns the AlertId field value
func (o *ListAlertsResponseBodyItem) GetAlertId() Ksuid {
	if o == nil {
		var ret Ksuid
		return ret
	}

	return o.AlertId
}

// GetAlertIdOk returns a tuple with the AlertId field value
// and a boolean to check if the value has been set.
func (o *ListAlertsResponseBodyItem) GetAlertIdOk() (*Ksuid, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlertId, true
}

// SetAlertId sets field value
func (o *ListAlertsResponseBodyItem) SetAlertId(v Ksuid) {
	o.AlertId = v
}

// GetCondition returns the Condition field value
func (o *ListAlertsResponseBodyItem) GetCondition() QueryCondition {
	if o == nil {
		var ret QueryCondition
		return ret
	}

	return o.Condition
}

// GetConditionOk returns a tuple with the Condition field value
// and a boolean to check if the value has been set.
func (o *ListAlertsResponseBodyItem) GetConditionOk() (*QueryCondition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Condition, true
}

// SetCondition sets field value
func (o *ListAlertsResponseBodyItem) SetCondition(v QueryCondition) {
	o.Condition = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListAlertsResponseBodyItem) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListAlertsResponseBodyItem) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ListAlertsResponseBodyItem) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ListAlertsResponseBodyItem) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ListAlertsResponseBodyItem) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ListAlertsResponseBodyItem) UnsetDescription() {
	o.Description.Unset()
}

// GetEnabled returns the Enabled field value
func (o *ListAlertsResponseBodyItem) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ListAlertsResponseBodyItem) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ListAlertsResponseBodyItem) SetEnabled(v bool) {
	o.Enabled = v
}

// GetFolderId returns the FolderId field value
func (o *ListAlertsResponseBodyItem) GetFolderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value
// and a boolean to check if the value has been set.
func (o *ListAlertsResponseBodyItem) GetFolderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FolderId, true
}

// SetFolderId sets field value
func (o *ListAlertsResponseBodyItem) SetFolderId(v string) {
	o.FolderId = v
}

// GetFolderName returns the FolderName field value
func (o *ListAlertsResponseBodyItem) GetFolderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FolderName
}

// GetFolderNameOk returns a tuple with the FolderName field value
// and a boolean to check if the value has been set.
func (o *ListAlertsResponseBodyItem) GetFolderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FolderName, true
}

// SetFolderName sets field value
func (o *ListAlertsResponseBodyItem) SetFolderName(v string) {
	o.FolderName = v
}

// GetLastSatisfiedAt returns the LastSatisfiedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListAlertsResponseBodyItem) GetLastSatisfiedAt() int64 {
	if o == nil || IsNil(o.LastSatisfiedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.LastSatisfiedAt.Get()
}

// GetLastSatisfiedAtOk returns a tuple with the LastSatisfiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListAlertsResponseBodyItem) GetLastSatisfiedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastSatisfiedAt.Get(), o.LastSatisfiedAt.IsSet()
}

// HasLastSatisfiedAt returns a boolean if a field has been set.
func (o *ListAlertsResponseBodyItem) HasLastSatisfiedAt() bool {
	if o != nil && o.LastSatisfiedAt.IsSet() {
		return true
	}

	return false
}

// SetLastSatisfiedAt gets a reference to the given NullableInt64 and assigns it to the LastSatisfiedAt field.
func (o *ListAlertsResponseBodyItem) SetLastSatisfiedAt(v int64) {
	o.LastSatisfiedAt.Set(&v)
}
// SetLastSatisfiedAtNil sets the value for LastSatisfiedAt to be an explicit nil
func (o *ListAlertsResponseBodyItem) SetLastSatisfiedAtNil() {
	o.LastSatisfiedAt.Set(nil)
}

// UnsetLastSatisfiedAt ensures that no value is present for LastSatisfiedAt, not even an explicit nil
func (o *ListAlertsResponseBodyItem) UnsetLastSatisfiedAt() {
	o.LastSatisfiedAt.Unset()
}

// GetLastTriggeredAt returns the LastTriggeredAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListAlertsResponseBodyItem) GetLastTriggeredAt() int64 {
	if o == nil || IsNil(o.LastTriggeredAt.Get()) {
		var ret int64
		return ret
	}
	return *o.LastTriggeredAt.Get()
}

// GetLastTriggeredAtOk returns a tuple with the LastTriggeredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListAlertsResponseBodyItem) GetLastTriggeredAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastTriggeredAt.Get(), o.LastTriggeredAt.IsSet()
}

// HasLastTriggeredAt returns a boolean if a field has been set.
func (o *ListAlertsResponseBodyItem) HasLastTriggeredAt() bool {
	if o != nil && o.LastTriggeredAt.IsSet() {
		return true
	}

	return false
}

// SetLastTriggeredAt gets a reference to the given NullableInt64 and assigns it to the LastTriggeredAt field.
func (o *ListAlertsResponseBodyItem) SetLastTriggeredAt(v int64) {
	o.LastTriggeredAt.Set(&v)
}
// SetLastTriggeredAtNil sets the value for LastTriggeredAt to be an explicit nil
func (o *ListAlertsResponseBodyItem) SetLastTriggeredAtNil() {
	o.LastTriggeredAt.Set(nil)
}

// UnsetLastTriggeredAt ensures that no value is present for LastTriggeredAt, not even an explicit nil
func (o *ListAlertsResponseBodyItem) UnsetLastTriggeredAt() {
	o.LastTriggeredAt.Unset()
}

// GetName returns the Name field value
func (o *ListAlertsResponseBodyItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ListAlertsResponseBodyItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ListAlertsResponseBodyItem) SetName(v string) {
	o.Name = v
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListAlertsResponseBodyItem) GetOwner() string {
	if o == nil || IsNil(o.Owner.Get()) {
		var ret string
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListAlertsResponseBodyItem) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *ListAlertsResponseBodyItem) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableString and assigns it to the Owner field.
func (o *ListAlertsResponseBodyItem) SetOwner(v string) {
	o.Owner.Set(&v)
}
// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *ListAlertsResponseBodyItem) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *ListAlertsResponseBodyItem) UnsetOwner() {
	o.Owner.Unset()
}

func (o ListAlertsResponseBodyItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAlertsResponseBodyItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alert_id"] = o.AlertId
	toSerialize["condition"] = o.Condition
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["folder_id"] = o.FolderId
	toSerialize["folder_name"] = o.FolderName
	if o.LastSatisfiedAt.IsSet() {
		toSerialize["last_satisfied_at"] = o.LastSatisfiedAt.Get()
	}
	if o.LastTriggeredAt.IsSet() {
		toSerialize["last_triggered_at"] = o.LastTriggeredAt.Get()
	}
	toSerialize["name"] = o.Name
	if o.Owner.IsSet() {
		toSerialize["owner"] = o.Owner.Get()
	}
	return toSerialize, nil
}

func (o *ListAlertsResponseBodyItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"alert_id",
		"condition",
		"enabled",
		"folder_id",
		"folder_name",
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

	varListAlertsResponseBodyItem := _ListAlertsResponseBodyItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListAlertsResponseBodyItem)

	if err != nil {
		return err
	}

	*o = ListAlertsResponseBodyItem(varListAlertsResponseBodyItem)

	return err
}

type NullableListAlertsResponseBodyItem struct {
	value *ListAlertsResponseBodyItem
	isSet bool
}

func (v NullableListAlertsResponseBodyItem) Get() *ListAlertsResponseBodyItem {
	return v.value
}

func (v *NullableListAlertsResponseBodyItem) Set(val *ListAlertsResponseBodyItem) {
	v.value = val
	v.isSet = true
}

func (v NullableListAlertsResponseBodyItem) IsSet() bool {
	return v.isSet
}

func (v *NullableListAlertsResponseBodyItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAlertsResponseBodyItem(val *ListAlertsResponseBodyItem) *NullableListAlertsResponseBodyItem {
	return &NullableListAlertsResponseBodyItem{value: val, isSet: true}
}

func (v NullableListAlertsResponseBodyItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAlertsResponseBodyItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


