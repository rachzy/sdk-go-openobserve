# AxisItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationFunction** | Pointer to [**NullableAggregationFunc**](AggregationFunc.md) |  | [optional] 
**Alias** | **string** |  | 
**Color** | Pointer to **NullableString** |  | [optional] 
**Column** | **string** |  | 
**Label** | **string** |  | 

## Methods

### NewAxisItem

`func NewAxisItem(alias string, column string, label string, ) *AxisItem`

NewAxisItem instantiates a new AxisItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAxisItemWithDefaults

`func NewAxisItemWithDefaults() *AxisItem`

NewAxisItemWithDefaults instantiates a new AxisItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationFunction

`func (o *AxisItem) GetAggregationFunction() AggregationFunc`

GetAggregationFunction returns the AggregationFunction field if non-nil, zero value otherwise.

### GetAggregationFunctionOk

`func (o *AxisItem) GetAggregationFunctionOk() (*AggregationFunc, bool)`

GetAggregationFunctionOk returns a tuple with the AggregationFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationFunction

`func (o *AxisItem) SetAggregationFunction(v AggregationFunc)`

SetAggregationFunction sets AggregationFunction field to given value.

### HasAggregationFunction

`func (o *AxisItem) HasAggregationFunction() bool`

HasAggregationFunction returns a boolean if a field has been set.

### SetAggregationFunctionNil

`func (o *AxisItem) SetAggregationFunctionNil(b bool)`

 SetAggregationFunctionNil sets the value for AggregationFunction to be an explicit nil

### UnsetAggregationFunction
`func (o *AxisItem) UnsetAggregationFunction()`

UnsetAggregationFunction ensures that no value is present for AggregationFunction, not even an explicit nil
### GetAlias

`func (o *AxisItem) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *AxisItem) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *AxisItem) SetAlias(v string)`

SetAlias sets Alias field to given value.


### GetColor

`func (o *AxisItem) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *AxisItem) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *AxisItem) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *AxisItem) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *AxisItem) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *AxisItem) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetColumn

`func (o *AxisItem) GetColumn() string`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *AxisItem) GetColumnOk() (*string, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *AxisItem) SetColumn(v string)`

SetColumn sets Column field to given value.


### GetLabel

`func (o *AxisItem) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AxisItem) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AxisItem) SetLabel(v string)`

SetLabel sets Label field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


