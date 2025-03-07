# PanelFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Column** | **string** |  | 
**Operator** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 
**Value** | Pointer to **NullableString** |  | [optional] 
**Values** | **[]string** |  | 

## Methods

### NewPanelFilter

`func NewPanelFilter(column string, type_ string, values []string, ) *PanelFilter`

NewPanelFilter instantiates a new PanelFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPanelFilterWithDefaults

`func NewPanelFilterWithDefaults() *PanelFilter`

NewPanelFilterWithDefaults instantiates a new PanelFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumn

`func (o *PanelFilter) GetColumn() string`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *PanelFilter) GetColumnOk() (*string, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *PanelFilter) SetColumn(v string)`

SetColumn sets Column field to given value.


### GetOperator

`func (o *PanelFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PanelFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PanelFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *PanelFilter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperatorNil

`func (o *PanelFilter) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *PanelFilter) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetType

`func (o *PanelFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PanelFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PanelFilter) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *PanelFilter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PanelFilter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PanelFilter) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PanelFilter) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *PanelFilter) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *PanelFilter) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetValues

`func (o *PanelFilter) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *PanelFilter) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *PanelFilter) SetValues(v []string)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


