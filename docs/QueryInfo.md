# QueryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | **int64** |  | 
**Sql** | **string** |  | 
**StartTime** | **int64** |  | 

## Methods

### NewQueryInfo

`func NewQueryInfo(endTime int64, sql string, startTime int64, ) *QueryInfo`

NewQueryInfo instantiates a new QueryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryInfoWithDefaults

`func NewQueryInfoWithDefaults() *QueryInfo`

NewQueryInfoWithDefaults instantiates a new QueryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *QueryInfo) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *QueryInfo) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *QueryInfo) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.


### GetSql

`func (o *QueryInfo) GetSql() string`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *QueryInfo) GetSqlOk() (*string, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *QueryInfo) SetSql(v string)`

SetSql sets Sql field to given value.


### GetStartTime

`func (o *QueryInfo) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *QueryInfo) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *QueryInfo) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


