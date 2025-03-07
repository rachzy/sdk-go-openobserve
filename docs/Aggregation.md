# Aggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Function** | [**AggFunction**](AggFunction.md) |  | 
**GroupBy** | Pointer to **[]string** |  | [optional] 
**Having** | [**Condition**](Condition.md) |  | 

## Methods

### NewAggregation

`func NewAggregation(function AggFunction, having Condition, ) *Aggregation`

NewAggregation instantiates a new Aggregation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationWithDefaults

`func NewAggregationWithDefaults() *Aggregation`

NewAggregationWithDefaults instantiates a new Aggregation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunction

`func (o *Aggregation) GetFunction() AggFunction`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *Aggregation) GetFunctionOk() (*AggFunction, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *Aggregation) SetFunction(v AggFunction)`

SetFunction sets Function field to given value.


### GetGroupBy

`func (o *Aggregation) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *Aggregation) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *Aggregation) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *Aggregation) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### SetGroupByNil

`func (o *Aggregation) SetGroupByNil(b bool)`

 SetGroupByNil sets the value for GroupBy to be an explicit nil

### UnsetGroupBy
`func (o *Aggregation) UnsetGroupBy()`

UnsetGroupBy ensures that no value is present for GroupBy, not even an explicit nil
### GetHaving

`func (o *Aggregation) GetHaving() Condition`

GetHaving returns the Having field if non-nil, zero value otherwise.

### GetHavingOk

`func (o *Aggregation) GetHavingOk() (*Condition, bool)`

GetHavingOk returns a tuple with the Having field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaving

`func (o *Aggregation) SetHaving(v Condition)`

SetHaving sets Having field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


