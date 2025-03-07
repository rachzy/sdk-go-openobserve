# SearchQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionId** | Pointer to **NullableString** |  | [optional] 
**EndTime** | Pointer to **int64** |  | [optional] 
**From** | Pointer to **int64** |  | [optional] 
**QueryFn** | Pointer to **NullableString** |  | [optional] 
**QueryType** | Pointer to **string** |  | [optional] 
**QuickMode** | Pointer to **bool** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**SkipWal** | Pointer to **bool** |  | [optional] 
**Sql** | **string** |  | 
**StartTime** | Pointer to **int64** |  | [optional] 
**StreamingId** | Pointer to **NullableString** |  | [optional] 
**StreamingOutput** | Pointer to **bool** |  | [optional] 
**TrackTotalHits** | Pointer to **bool** |  | [optional] 
**UsesZoFn** | Pointer to **bool** |  | [optional] 

## Methods

### NewSearchQuery

`func NewSearchQuery(sql string, ) *SearchQuery`

NewSearchQuery instantiates a new SearchQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchQueryWithDefaults

`func NewSearchQueryWithDefaults() *SearchQuery`

NewSearchQueryWithDefaults instantiates a new SearchQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionId

`func (o *SearchQuery) GetActionId() string`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *SearchQuery) GetActionIdOk() (*string, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *SearchQuery) SetActionId(v string)`

SetActionId sets ActionId field to given value.

### HasActionId

`func (o *SearchQuery) HasActionId() bool`

HasActionId returns a boolean if a field has been set.

### SetActionIdNil

`func (o *SearchQuery) SetActionIdNil(b bool)`

 SetActionIdNil sets the value for ActionId to be an explicit nil

### UnsetActionId
`func (o *SearchQuery) UnsetActionId()`

UnsetActionId ensures that no value is present for ActionId, not even an explicit nil
### GetEndTime

`func (o *SearchQuery) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SearchQuery) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SearchQuery) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *SearchQuery) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetFrom

`func (o *SearchQuery) GetFrom() int64`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SearchQuery) GetFromOk() (*int64, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SearchQuery) SetFrom(v int64)`

SetFrom sets From field to given value.

### HasFrom

`func (o *SearchQuery) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetQueryFn

`func (o *SearchQuery) GetQueryFn() string`

GetQueryFn returns the QueryFn field if non-nil, zero value otherwise.

### GetQueryFnOk

`func (o *SearchQuery) GetQueryFnOk() (*string, bool)`

GetQueryFnOk returns a tuple with the QueryFn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryFn

`func (o *SearchQuery) SetQueryFn(v string)`

SetQueryFn sets QueryFn field to given value.

### HasQueryFn

`func (o *SearchQuery) HasQueryFn() bool`

HasQueryFn returns a boolean if a field has been set.

### SetQueryFnNil

`func (o *SearchQuery) SetQueryFnNil(b bool)`

 SetQueryFnNil sets the value for QueryFn to be an explicit nil

### UnsetQueryFn
`func (o *SearchQuery) UnsetQueryFn()`

UnsetQueryFn ensures that no value is present for QueryFn, not even an explicit nil
### GetQueryType

`func (o *SearchQuery) GetQueryType() string`

GetQueryType returns the QueryType field if non-nil, zero value otherwise.

### GetQueryTypeOk

`func (o *SearchQuery) GetQueryTypeOk() (*string, bool)`

GetQueryTypeOk returns a tuple with the QueryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryType

`func (o *SearchQuery) SetQueryType(v string)`

SetQueryType sets QueryType field to given value.

### HasQueryType

`func (o *SearchQuery) HasQueryType() bool`

HasQueryType returns a boolean if a field has been set.

### GetQuickMode

`func (o *SearchQuery) GetQuickMode() bool`

GetQuickMode returns the QuickMode field if non-nil, zero value otherwise.

### GetQuickModeOk

`func (o *SearchQuery) GetQuickModeOk() (*bool, bool)`

GetQuickModeOk returns a tuple with the QuickMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickMode

`func (o *SearchQuery) SetQuickMode(v bool)`

SetQuickMode sets QuickMode field to given value.

### HasQuickMode

`func (o *SearchQuery) HasQuickMode() bool`

HasQuickMode returns a boolean if a field has been set.

### GetSize

`func (o *SearchQuery) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SearchQuery) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SearchQuery) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *SearchQuery) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSkipWal

`func (o *SearchQuery) GetSkipWal() bool`

GetSkipWal returns the SkipWal field if non-nil, zero value otherwise.

### GetSkipWalOk

`func (o *SearchQuery) GetSkipWalOk() (*bool, bool)`

GetSkipWalOk returns a tuple with the SkipWal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipWal

`func (o *SearchQuery) SetSkipWal(v bool)`

SetSkipWal sets SkipWal field to given value.

### HasSkipWal

`func (o *SearchQuery) HasSkipWal() bool`

HasSkipWal returns a boolean if a field has been set.

### GetSql

`func (o *SearchQuery) GetSql() string`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *SearchQuery) GetSqlOk() (*string, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *SearchQuery) SetSql(v string)`

SetSql sets Sql field to given value.


### GetStartTime

`func (o *SearchQuery) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SearchQuery) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SearchQuery) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *SearchQuery) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStreamingId

`func (o *SearchQuery) GetStreamingId() string`

GetStreamingId returns the StreamingId field if non-nil, zero value otherwise.

### GetStreamingIdOk

`func (o *SearchQuery) GetStreamingIdOk() (*string, bool)`

GetStreamingIdOk returns a tuple with the StreamingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamingId

`func (o *SearchQuery) SetStreamingId(v string)`

SetStreamingId sets StreamingId field to given value.

### HasStreamingId

`func (o *SearchQuery) HasStreamingId() bool`

HasStreamingId returns a boolean if a field has been set.

### SetStreamingIdNil

`func (o *SearchQuery) SetStreamingIdNil(b bool)`

 SetStreamingIdNil sets the value for StreamingId to be an explicit nil

### UnsetStreamingId
`func (o *SearchQuery) UnsetStreamingId()`

UnsetStreamingId ensures that no value is present for StreamingId, not even an explicit nil
### GetStreamingOutput

`func (o *SearchQuery) GetStreamingOutput() bool`

GetStreamingOutput returns the StreamingOutput field if non-nil, zero value otherwise.

### GetStreamingOutputOk

`func (o *SearchQuery) GetStreamingOutputOk() (*bool, bool)`

GetStreamingOutputOk returns a tuple with the StreamingOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamingOutput

`func (o *SearchQuery) SetStreamingOutput(v bool)`

SetStreamingOutput sets StreamingOutput field to given value.

### HasStreamingOutput

`func (o *SearchQuery) HasStreamingOutput() bool`

HasStreamingOutput returns a boolean if a field has been set.

### GetTrackTotalHits

`func (o *SearchQuery) GetTrackTotalHits() bool`

GetTrackTotalHits returns the TrackTotalHits field if non-nil, zero value otherwise.

### GetTrackTotalHitsOk

`func (o *SearchQuery) GetTrackTotalHitsOk() (*bool, bool)`

GetTrackTotalHitsOk returns a tuple with the TrackTotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackTotalHits

`func (o *SearchQuery) SetTrackTotalHits(v bool)`

SetTrackTotalHits sets TrackTotalHits field to given value.

### HasTrackTotalHits

`func (o *SearchQuery) HasTrackTotalHits() bool`

HasTrackTotalHits returns a boolean if a field has been set.

### GetUsesZoFn

`func (o *SearchQuery) GetUsesZoFn() bool`

GetUsesZoFn returns the UsesZoFn field if non-nil, zero value otherwise.

### GetUsesZoFnOk

`func (o *SearchQuery) GetUsesZoFnOk() (*bool, bool)`

GetUsesZoFnOk returns a tuple with the UsesZoFn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesZoFn

`func (o *SearchQuery) SetUsesZoFn(v bool)`

SetUsesZoFn sets UsesZoFn field to given value.

### HasUsesZoFn

`func (o *SearchQuery) HasUsesZoFn() bool`

HasUsesZoFn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


