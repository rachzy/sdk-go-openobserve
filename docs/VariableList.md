# VariableList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** |  | 
**Name** | **string** |  | 
**Options** | Pointer to [**[]CustomFieldsOption**](CustomFieldsOption.md) |  | [optional] 
**QueryData** | Pointer to [**NullableQueryData**](QueryData.md) |  | [optional] 
**Type** | **string** |  | 
**Value** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewVariableList

`func NewVariableList(label string, name string, type_ string, ) *VariableList`

NewVariableList instantiates a new VariableList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableListWithDefaults

`func NewVariableListWithDefaults() *VariableList`

NewVariableListWithDefaults instantiates a new VariableList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *VariableList) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *VariableList) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *VariableList) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *VariableList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VariableList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VariableList) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *VariableList) GetOptions() []CustomFieldsOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *VariableList) GetOptionsOk() (*[]CustomFieldsOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *VariableList) SetOptions(v []CustomFieldsOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *VariableList) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *VariableList) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *VariableList) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetQueryData

`func (o *VariableList) GetQueryData() QueryData`

GetQueryData returns the QueryData field if non-nil, zero value otherwise.

### GetQueryDataOk

`func (o *VariableList) GetQueryDataOk() (*QueryData, bool)`

GetQueryDataOk returns a tuple with the QueryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryData

`func (o *VariableList) SetQueryData(v QueryData)`

SetQueryData sets QueryData field to given value.

### HasQueryData

`func (o *VariableList) HasQueryData() bool`

HasQueryData returns a boolean if a field has been set.

### SetQueryDataNil

`func (o *VariableList) SetQueryDataNil(b bool)`

 SetQueryDataNil sets the value for QueryData to be an explicit nil

### UnsetQueryData
`func (o *VariableList) UnsetQueryData()`

UnsetQueryData ensures that no value is present for QueryData, not even an explicit nil
### GetType

`func (o *VariableList) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VariableList) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VariableList) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *VariableList) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VariableList) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VariableList) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *VariableList) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *VariableList) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *VariableList) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


