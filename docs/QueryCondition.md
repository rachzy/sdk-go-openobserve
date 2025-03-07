# QueryCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | Pointer to [**NullableAggregation**](Aggregation.md) |  | [optional] 
**Conditions** | Pointer to [**[]Condition**](Condition.md) |  | [optional] 
**MultiTimeRange** | Pointer to [**[]CompareHistoricData**](CompareHistoricData.md) |  | [optional] 
**Promql** | Pointer to **NullableString** |  | [optional] 
**PromqlCondition** | Pointer to [**NullableCondition**](Condition.md) |  | [optional] 
**SearchEventType** | Pointer to [**NullableSearchEventType**](SearchEventType.md) |  | [optional] 
**Sql** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**QueryType**](QueryType.md) |  | [optional] 
**VrlFunction** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewQueryCondition

`func NewQueryCondition() *QueryCondition`

NewQueryCondition instantiates a new QueryCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryConditionWithDefaults

`func NewQueryConditionWithDefaults() *QueryCondition`

NewQueryConditionWithDefaults instantiates a new QueryCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *QueryCondition) GetAggregation() Aggregation`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *QueryCondition) GetAggregationOk() (*Aggregation, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *QueryCondition) SetAggregation(v Aggregation)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *QueryCondition) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### SetAggregationNil

`func (o *QueryCondition) SetAggregationNil(b bool)`

 SetAggregationNil sets the value for Aggregation to be an explicit nil

### UnsetAggregation
`func (o *QueryCondition) UnsetAggregation()`

UnsetAggregation ensures that no value is present for Aggregation, not even an explicit nil
### GetConditions

`func (o *QueryCondition) GetConditions() []Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *QueryCondition) GetConditionsOk() (*[]Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *QueryCondition) SetConditions(v []Condition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *QueryCondition) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *QueryCondition) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *QueryCondition) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetMultiTimeRange

`func (o *QueryCondition) GetMultiTimeRange() []CompareHistoricData`

GetMultiTimeRange returns the MultiTimeRange field if non-nil, zero value otherwise.

### GetMultiTimeRangeOk

`func (o *QueryCondition) GetMultiTimeRangeOk() (*[]CompareHistoricData, bool)`

GetMultiTimeRangeOk returns a tuple with the MultiTimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiTimeRange

`func (o *QueryCondition) SetMultiTimeRange(v []CompareHistoricData)`

SetMultiTimeRange sets MultiTimeRange field to given value.

### HasMultiTimeRange

`func (o *QueryCondition) HasMultiTimeRange() bool`

HasMultiTimeRange returns a boolean if a field has been set.

### SetMultiTimeRangeNil

`func (o *QueryCondition) SetMultiTimeRangeNil(b bool)`

 SetMultiTimeRangeNil sets the value for MultiTimeRange to be an explicit nil

### UnsetMultiTimeRange
`func (o *QueryCondition) UnsetMultiTimeRange()`

UnsetMultiTimeRange ensures that no value is present for MultiTimeRange, not even an explicit nil
### GetPromql

`func (o *QueryCondition) GetPromql() string`

GetPromql returns the Promql field if non-nil, zero value otherwise.

### GetPromqlOk

`func (o *QueryCondition) GetPromqlOk() (*string, bool)`

GetPromqlOk returns a tuple with the Promql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromql

`func (o *QueryCondition) SetPromql(v string)`

SetPromql sets Promql field to given value.

### HasPromql

`func (o *QueryCondition) HasPromql() bool`

HasPromql returns a boolean if a field has been set.

### SetPromqlNil

`func (o *QueryCondition) SetPromqlNil(b bool)`

 SetPromqlNil sets the value for Promql to be an explicit nil

### UnsetPromql
`func (o *QueryCondition) UnsetPromql()`

UnsetPromql ensures that no value is present for Promql, not even an explicit nil
### GetPromqlCondition

`func (o *QueryCondition) GetPromqlCondition() Condition`

GetPromqlCondition returns the PromqlCondition field if non-nil, zero value otherwise.

### GetPromqlConditionOk

`func (o *QueryCondition) GetPromqlConditionOk() (*Condition, bool)`

GetPromqlConditionOk returns a tuple with the PromqlCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromqlCondition

`func (o *QueryCondition) SetPromqlCondition(v Condition)`

SetPromqlCondition sets PromqlCondition field to given value.

### HasPromqlCondition

`func (o *QueryCondition) HasPromqlCondition() bool`

HasPromqlCondition returns a boolean if a field has been set.

### SetPromqlConditionNil

`func (o *QueryCondition) SetPromqlConditionNil(b bool)`

 SetPromqlConditionNil sets the value for PromqlCondition to be an explicit nil

### UnsetPromqlCondition
`func (o *QueryCondition) UnsetPromqlCondition()`

UnsetPromqlCondition ensures that no value is present for PromqlCondition, not even an explicit nil
### GetSearchEventType

`func (o *QueryCondition) GetSearchEventType() SearchEventType`

GetSearchEventType returns the SearchEventType field if non-nil, zero value otherwise.

### GetSearchEventTypeOk

`func (o *QueryCondition) GetSearchEventTypeOk() (*SearchEventType, bool)`

GetSearchEventTypeOk returns a tuple with the SearchEventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchEventType

`func (o *QueryCondition) SetSearchEventType(v SearchEventType)`

SetSearchEventType sets SearchEventType field to given value.

### HasSearchEventType

`func (o *QueryCondition) HasSearchEventType() bool`

HasSearchEventType returns a boolean if a field has been set.

### SetSearchEventTypeNil

`func (o *QueryCondition) SetSearchEventTypeNil(b bool)`

 SetSearchEventTypeNil sets the value for SearchEventType to be an explicit nil

### UnsetSearchEventType
`func (o *QueryCondition) UnsetSearchEventType()`

UnsetSearchEventType ensures that no value is present for SearchEventType, not even an explicit nil
### GetSql

`func (o *QueryCondition) GetSql() string`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *QueryCondition) GetSqlOk() (*string, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *QueryCondition) SetSql(v string)`

SetSql sets Sql field to given value.

### HasSql

`func (o *QueryCondition) HasSql() bool`

HasSql returns a boolean if a field has been set.

### SetSqlNil

`func (o *QueryCondition) SetSqlNil(b bool)`

 SetSqlNil sets the value for Sql to be an explicit nil

### UnsetSql
`func (o *QueryCondition) UnsetSql()`

UnsetSql ensures that no value is present for Sql, not even an explicit nil
### GetType

`func (o *QueryCondition) GetType() QueryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueryCondition) GetTypeOk() (*QueryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueryCondition) SetType(v QueryType)`

SetType sets Type field to given value.

### HasType

`func (o *QueryCondition) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVrlFunction

`func (o *QueryCondition) GetVrlFunction() string`

GetVrlFunction returns the VrlFunction field if non-nil, zero value otherwise.

### GetVrlFunctionOk

`func (o *QueryCondition) GetVrlFunctionOk() (*string, bool)`

GetVrlFunctionOk returns a tuple with the VrlFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrlFunction

`func (o *QueryCondition) SetVrlFunction(v string)`

SetVrlFunction sets VrlFunction field to given value.

### HasVrlFunction

`func (o *QueryCondition) HasVrlFunction() bool`

HasVrlFunction returns a boolean if a field has been set.

### SetVrlFunctionNil

`func (o *QueryCondition) SetVrlFunctionNil(b bool)`

 SetVrlFunctionNil sets the value for VrlFunction to be an explicit nil

### UnsetVrlFunction
`func (o *QueryCondition) UnsetVrlFunction()`

UnsetVrlFunction ensures that no value is present for VrlFunction, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


