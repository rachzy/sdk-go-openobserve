# QueryStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **int64** |  | 
**OrgId** | Pointer to **NullableString** |  | [optional] 
**Query** | Pointer to [**NullableQueryInfo**](QueryInfo.md) |  | [optional] 
**ScanStats** | Pointer to [**NullableScanStats**](ScanStats.md) |  | [optional] 
**SearchType** | Pointer to [**NullableSearchEventType**](SearchEventType.md) |  | [optional] 
**StartedAt** | **int64** |  | 
**Status** | **string** |  | 
**StreamType** | Pointer to **NullableString** |  | [optional] 
**TraceId** | **string** |  | 
**UserId** | Pointer to **NullableString** |  | [optional] 
**WorkGroup** | **string** |  | 

## Methods

### NewQueryStatus

`func NewQueryStatus(createdAt int64, startedAt int64, status string, traceId string, workGroup string, ) *QueryStatus`

NewQueryStatus instantiates a new QueryStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryStatusWithDefaults

`func NewQueryStatusWithDefaults() *QueryStatus`

NewQueryStatusWithDefaults instantiates a new QueryStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *QueryStatus) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QueryStatus) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QueryStatus) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetOrgId

`func (o *QueryStatus) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *QueryStatus) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *QueryStatus) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *QueryStatus) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *QueryStatus) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *QueryStatus) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
### GetQuery

`func (o *QueryStatus) GetQuery() QueryInfo`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *QueryStatus) GetQueryOk() (*QueryInfo, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *QueryStatus) SetQuery(v QueryInfo)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *QueryStatus) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *QueryStatus) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *QueryStatus) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetScanStats

`func (o *QueryStatus) GetScanStats() ScanStats`

GetScanStats returns the ScanStats field if non-nil, zero value otherwise.

### GetScanStatsOk

`func (o *QueryStatus) GetScanStatsOk() (*ScanStats, bool)`

GetScanStatsOk returns a tuple with the ScanStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanStats

`func (o *QueryStatus) SetScanStats(v ScanStats)`

SetScanStats sets ScanStats field to given value.

### HasScanStats

`func (o *QueryStatus) HasScanStats() bool`

HasScanStats returns a boolean if a field has been set.

### SetScanStatsNil

`func (o *QueryStatus) SetScanStatsNil(b bool)`

 SetScanStatsNil sets the value for ScanStats to be an explicit nil

### UnsetScanStats
`func (o *QueryStatus) UnsetScanStats()`

UnsetScanStats ensures that no value is present for ScanStats, not even an explicit nil
### GetSearchType

`func (o *QueryStatus) GetSearchType() SearchEventType`

GetSearchType returns the SearchType field if non-nil, zero value otherwise.

### GetSearchTypeOk

`func (o *QueryStatus) GetSearchTypeOk() (*SearchEventType, bool)`

GetSearchTypeOk returns a tuple with the SearchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchType

`func (o *QueryStatus) SetSearchType(v SearchEventType)`

SetSearchType sets SearchType field to given value.

### HasSearchType

`func (o *QueryStatus) HasSearchType() bool`

HasSearchType returns a boolean if a field has been set.

### SetSearchTypeNil

`func (o *QueryStatus) SetSearchTypeNil(b bool)`

 SetSearchTypeNil sets the value for SearchType to be an explicit nil

### UnsetSearchType
`func (o *QueryStatus) UnsetSearchType()`

UnsetSearchType ensures that no value is present for SearchType, not even an explicit nil
### GetStartedAt

`func (o *QueryStatus) GetStartedAt() int64`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *QueryStatus) GetStartedAtOk() (*int64, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *QueryStatus) SetStartedAt(v int64)`

SetStartedAt sets StartedAt field to given value.


### GetStatus

`func (o *QueryStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QueryStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QueryStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStreamType

`func (o *QueryStatus) GetStreamType() string`

GetStreamType returns the StreamType field if non-nil, zero value otherwise.

### GetStreamTypeOk

`func (o *QueryStatus) GetStreamTypeOk() (*string, bool)`

GetStreamTypeOk returns a tuple with the StreamType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamType

`func (o *QueryStatus) SetStreamType(v string)`

SetStreamType sets StreamType field to given value.

### HasStreamType

`func (o *QueryStatus) HasStreamType() bool`

HasStreamType returns a boolean if a field has been set.

### SetStreamTypeNil

`func (o *QueryStatus) SetStreamTypeNil(b bool)`

 SetStreamTypeNil sets the value for StreamType to be an explicit nil

### UnsetStreamType
`func (o *QueryStatus) UnsetStreamType()`

UnsetStreamType ensures that no value is present for StreamType, not even an explicit nil
### GetTraceId

`func (o *QueryStatus) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *QueryStatus) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *QueryStatus) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetUserId

`func (o *QueryStatus) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *QueryStatus) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *QueryStatus) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *QueryStatus) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *QueryStatus) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *QueryStatus) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetWorkGroup

`func (o *QueryStatus) GetWorkGroup() string`

GetWorkGroup returns the WorkGroup field if non-nil, zero value otherwise.

### GetWorkGroupOk

`func (o *QueryStatus) GetWorkGroupOk() (*string, bool)`

GetWorkGroupOk returns a tuple with the WorkGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkGroup

`func (o *QueryStatus) SetWorkGroup(v string)`

SetWorkGroup sets WorkGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


