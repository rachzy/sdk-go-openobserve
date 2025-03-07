# SearchPartitionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | Pointer to **[]string** |  | [optional] 
**Encoding** | Pointer to [**RequestEncoding**](RequestEncoding.md) |  | [optional] 
**EndTime** | **int64** |  | 
**QueryFn** | Pointer to **NullableString** |  | [optional] 
**Regions** | Pointer to **[]string** |  | [optional] 
**Sql** | **string** |  | 
**StartTime** | **int64** |  | 
**StreamingOutput** | Pointer to **bool** |  | [optional] 

## Methods

### NewSearchPartitionRequest

`func NewSearchPartitionRequest(endTime int64, sql string, startTime int64, ) *SearchPartitionRequest`

NewSearchPartitionRequest instantiates a new SearchPartitionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchPartitionRequestWithDefaults

`func NewSearchPartitionRequestWithDefaults() *SearchPartitionRequest`

NewSearchPartitionRequestWithDefaults instantiates a new SearchPartitionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusters

`func (o *SearchPartitionRequest) GetClusters() []string`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *SearchPartitionRequest) GetClustersOk() (*[]string, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *SearchPartitionRequest) SetClusters(v []string)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *SearchPartitionRequest) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetEncoding

`func (o *SearchPartitionRequest) GetEncoding() RequestEncoding`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *SearchPartitionRequest) GetEncodingOk() (*RequestEncoding, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *SearchPartitionRequest) SetEncoding(v RequestEncoding)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *SearchPartitionRequest) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetEndTime

`func (o *SearchPartitionRequest) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SearchPartitionRequest) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SearchPartitionRequest) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.


### GetQueryFn

`func (o *SearchPartitionRequest) GetQueryFn() string`

GetQueryFn returns the QueryFn field if non-nil, zero value otherwise.

### GetQueryFnOk

`func (o *SearchPartitionRequest) GetQueryFnOk() (*string, bool)`

GetQueryFnOk returns a tuple with the QueryFn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryFn

`func (o *SearchPartitionRequest) SetQueryFn(v string)`

SetQueryFn sets QueryFn field to given value.

### HasQueryFn

`func (o *SearchPartitionRequest) HasQueryFn() bool`

HasQueryFn returns a boolean if a field has been set.

### SetQueryFnNil

`func (o *SearchPartitionRequest) SetQueryFnNil(b bool)`

 SetQueryFnNil sets the value for QueryFn to be an explicit nil

### UnsetQueryFn
`func (o *SearchPartitionRequest) UnsetQueryFn()`

UnsetQueryFn ensures that no value is present for QueryFn, not even an explicit nil
### GetRegions

`func (o *SearchPartitionRequest) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *SearchPartitionRequest) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *SearchPartitionRequest) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *SearchPartitionRequest) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetSql

`func (o *SearchPartitionRequest) GetSql() string`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *SearchPartitionRequest) GetSqlOk() (*string, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *SearchPartitionRequest) SetSql(v string)`

SetSql sets Sql field to given value.


### GetStartTime

`func (o *SearchPartitionRequest) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SearchPartitionRequest) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SearchPartitionRequest) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.


### GetStreamingOutput

`func (o *SearchPartitionRequest) GetStreamingOutput() bool`

GetStreamingOutput returns the StreamingOutput field if non-nil, zero value otherwise.

### GetStreamingOutputOk

`func (o *SearchPartitionRequest) GetStreamingOutputOk() (*bool, bool)`

GetStreamingOutputOk returns a tuple with the StreamingOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamingOutput

`func (o *SearchPartitionRequest) SetStreamingOutput(v bool)`

SetStreamingOutput sets StreamingOutput field to given value.

### HasStreamingOutput

`func (o *SearchPartitionRequest) HasStreamingOutput() bool`

HasStreamingOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


