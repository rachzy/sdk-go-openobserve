# Condition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Column** | **string** |  | 
**IgnoreCase** | Pointer to **bool** |  | [optional] 
**Operator** | [**Operator**](Operator.md) |  | 
**Value** | **map[string]interface{}** |  | 

## Methods

### NewCondition

`func NewCondition(column string, operator Operator, value map[string]interface{}, ) *Condition`

NewCondition instantiates a new Condition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionWithDefaults

`func NewConditionWithDefaults() *Condition`

NewConditionWithDefaults instantiates a new Condition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumn

`func (o *Condition) GetColumn() string`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *Condition) GetColumnOk() (*string, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *Condition) SetColumn(v string)`

SetColumn sets Column field to given value.


### GetIgnoreCase

`func (o *Condition) GetIgnoreCase() bool`

GetIgnoreCase returns the IgnoreCase field if non-nil, zero value otherwise.

### GetIgnoreCaseOk

`func (o *Condition) GetIgnoreCaseOk() (*bool, bool)`

GetIgnoreCaseOk returns a tuple with the IgnoreCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreCase

`func (o *Condition) SetIgnoreCase(v bool)`

SetIgnoreCase sets IgnoreCase field to given value.

### HasIgnoreCase

`func (o *Condition) HasIgnoreCase() bool`

HasIgnoreCase returns a boolean if a field has been set.

### GetOperator

`func (o *Condition) GetOperator() Operator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *Condition) GetOperatorOk() (*Operator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *Condition) SetOperator(v Operator)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *Condition) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Condition) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Condition) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


