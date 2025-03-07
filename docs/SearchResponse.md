# SearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CachedRatio** | **int32** |  | 
**Columns** | Pointer to **[]string** |  | [optional] 
**From** | **int64** |  | 
**FunctionError** | Pointer to **string** |  | [optional] 
**HistogramInterval** | Pointer to **NullableInt64** |  | [optional] 
**Hits** | **[]map[string]interface{}** |  | 
**IdxScanSize** | **int32** |  | 
**IsPartial** | Pointer to **bool** |  | [optional] 
**NewEndTime** | Pointer to **NullableInt64** |  | [optional] 
**NewStartTime** | Pointer to **NullableInt64** |  | [optional] 
**OrderBy** | Pointer to [**NullableOrderBy**](OrderBy.md) |  | [optional] 
**ResponseType** | Pointer to **string** |  | [optional] 
**ResultCacheRatio** | Pointer to **int32** |  | [optional] 
**ScanRecords** | **int32** |  | 
**ScanSize** | **int32** |  | 
**Size** | **int64** |  | 
**Took** | **int32** |  | 
**TookDetail** | Pointer to [**NullableResponseTook**](ResponseTook.md) |  | [optional] 
**Total** | **int32** |  | 
**TraceId** | Pointer to **string** |  | [optional] 
**WorkGroup** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSearchResponse

`func NewSearchResponse(cachedRatio int32, from int64, hits []map[string]interface{}, idxScanSize int32, scanRecords int32, scanSize int32, size int64, took int32, total int32, ) *SearchResponse`

NewSearchResponse instantiates a new SearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResponseWithDefaults

`func NewSearchResponseWithDefaults() *SearchResponse`

NewSearchResponseWithDefaults instantiates a new SearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCachedRatio

`func (o *SearchResponse) GetCachedRatio() int32`

GetCachedRatio returns the CachedRatio field if non-nil, zero value otherwise.

### GetCachedRatioOk

`func (o *SearchResponse) GetCachedRatioOk() (*int32, bool)`

GetCachedRatioOk returns a tuple with the CachedRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachedRatio

`func (o *SearchResponse) SetCachedRatio(v int32)`

SetCachedRatio sets CachedRatio field to given value.


### GetColumns

`func (o *SearchResponse) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *SearchResponse) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *SearchResponse) SetColumns(v []string)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *SearchResponse) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetFrom

`func (o *SearchResponse) GetFrom() int64`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SearchResponse) GetFromOk() (*int64, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SearchResponse) SetFrom(v int64)`

SetFrom sets From field to given value.


### GetFunctionError

`func (o *SearchResponse) GetFunctionError() string`

GetFunctionError returns the FunctionError field if non-nil, zero value otherwise.

### GetFunctionErrorOk

`func (o *SearchResponse) GetFunctionErrorOk() (*string, bool)`

GetFunctionErrorOk returns a tuple with the FunctionError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionError

`func (o *SearchResponse) SetFunctionError(v string)`

SetFunctionError sets FunctionError field to given value.

### HasFunctionError

`func (o *SearchResponse) HasFunctionError() bool`

HasFunctionError returns a boolean if a field has been set.

### GetHistogramInterval

`func (o *SearchResponse) GetHistogramInterval() int64`

GetHistogramInterval returns the HistogramInterval field if non-nil, zero value otherwise.

### GetHistogramIntervalOk

`func (o *SearchResponse) GetHistogramIntervalOk() (*int64, bool)`

GetHistogramIntervalOk returns a tuple with the HistogramInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogramInterval

`func (o *SearchResponse) SetHistogramInterval(v int64)`

SetHistogramInterval sets HistogramInterval field to given value.

### HasHistogramInterval

`func (o *SearchResponse) HasHistogramInterval() bool`

HasHistogramInterval returns a boolean if a field has been set.

### SetHistogramIntervalNil

`func (o *SearchResponse) SetHistogramIntervalNil(b bool)`

 SetHistogramIntervalNil sets the value for HistogramInterval to be an explicit nil

### UnsetHistogramInterval
`func (o *SearchResponse) UnsetHistogramInterval()`

UnsetHistogramInterval ensures that no value is present for HistogramInterval, not even an explicit nil
### GetHits

`func (o *SearchResponse) GetHits() []map[string]interface{}`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *SearchResponse) GetHitsOk() (*[]map[string]interface{}, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *SearchResponse) SetHits(v []map[string]interface{})`

SetHits sets Hits field to given value.


### GetIdxScanSize

`func (o *SearchResponse) GetIdxScanSize() int32`

GetIdxScanSize returns the IdxScanSize field if non-nil, zero value otherwise.

### GetIdxScanSizeOk

`func (o *SearchResponse) GetIdxScanSizeOk() (*int32, bool)`

GetIdxScanSizeOk returns a tuple with the IdxScanSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdxScanSize

`func (o *SearchResponse) SetIdxScanSize(v int32)`

SetIdxScanSize sets IdxScanSize field to given value.


### GetIsPartial

`func (o *SearchResponse) GetIsPartial() bool`

GetIsPartial returns the IsPartial field if non-nil, zero value otherwise.

### GetIsPartialOk

`func (o *SearchResponse) GetIsPartialOk() (*bool, bool)`

GetIsPartialOk returns a tuple with the IsPartial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPartial

`func (o *SearchResponse) SetIsPartial(v bool)`

SetIsPartial sets IsPartial field to given value.

### HasIsPartial

`func (o *SearchResponse) HasIsPartial() bool`

HasIsPartial returns a boolean if a field has been set.

### GetNewEndTime

`func (o *SearchResponse) GetNewEndTime() int64`

GetNewEndTime returns the NewEndTime field if non-nil, zero value otherwise.

### GetNewEndTimeOk

`func (o *SearchResponse) GetNewEndTimeOk() (*int64, bool)`

GetNewEndTimeOk returns a tuple with the NewEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewEndTime

`func (o *SearchResponse) SetNewEndTime(v int64)`

SetNewEndTime sets NewEndTime field to given value.

### HasNewEndTime

`func (o *SearchResponse) HasNewEndTime() bool`

HasNewEndTime returns a boolean if a field has been set.

### SetNewEndTimeNil

`func (o *SearchResponse) SetNewEndTimeNil(b bool)`

 SetNewEndTimeNil sets the value for NewEndTime to be an explicit nil

### UnsetNewEndTime
`func (o *SearchResponse) UnsetNewEndTime()`

UnsetNewEndTime ensures that no value is present for NewEndTime, not even an explicit nil
### GetNewStartTime

`func (o *SearchResponse) GetNewStartTime() int64`

GetNewStartTime returns the NewStartTime field if non-nil, zero value otherwise.

### GetNewStartTimeOk

`func (o *SearchResponse) GetNewStartTimeOk() (*int64, bool)`

GetNewStartTimeOk returns a tuple with the NewStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewStartTime

`func (o *SearchResponse) SetNewStartTime(v int64)`

SetNewStartTime sets NewStartTime field to given value.

### HasNewStartTime

`func (o *SearchResponse) HasNewStartTime() bool`

HasNewStartTime returns a boolean if a field has been set.

### SetNewStartTimeNil

`func (o *SearchResponse) SetNewStartTimeNil(b bool)`

 SetNewStartTimeNil sets the value for NewStartTime to be an explicit nil

### UnsetNewStartTime
`func (o *SearchResponse) UnsetNewStartTime()`

UnsetNewStartTime ensures that no value is present for NewStartTime, not even an explicit nil
### GetOrderBy

`func (o *SearchResponse) GetOrderBy() OrderBy`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *SearchResponse) GetOrderByOk() (*OrderBy, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *SearchResponse) SetOrderBy(v OrderBy)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *SearchResponse) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### SetOrderByNil

`func (o *SearchResponse) SetOrderByNil(b bool)`

 SetOrderByNil sets the value for OrderBy to be an explicit nil

### UnsetOrderBy
`func (o *SearchResponse) UnsetOrderBy()`

UnsetOrderBy ensures that no value is present for OrderBy, not even an explicit nil
### GetResponseType

`func (o *SearchResponse) GetResponseType() string`

GetResponseType returns the ResponseType field if non-nil, zero value otherwise.

### GetResponseTypeOk

`func (o *SearchResponse) GetResponseTypeOk() (*string, bool)`

GetResponseTypeOk returns a tuple with the ResponseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseType

`func (o *SearchResponse) SetResponseType(v string)`

SetResponseType sets ResponseType field to given value.

### HasResponseType

`func (o *SearchResponse) HasResponseType() bool`

HasResponseType returns a boolean if a field has been set.

### GetResultCacheRatio

`func (o *SearchResponse) GetResultCacheRatio() int32`

GetResultCacheRatio returns the ResultCacheRatio field if non-nil, zero value otherwise.

### GetResultCacheRatioOk

`func (o *SearchResponse) GetResultCacheRatioOk() (*int32, bool)`

GetResultCacheRatioOk returns a tuple with the ResultCacheRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCacheRatio

`func (o *SearchResponse) SetResultCacheRatio(v int32)`

SetResultCacheRatio sets ResultCacheRatio field to given value.

### HasResultCacheRatio

`func (o *SearchResponse) HasResultCacheRatio() bool`

HasResultCacheRatio returns a boolean if a field has been set.

### GetScanRecords

`func (o *SearchResponse) GetScanRecords() int32`

GetScanRecords returns the ScanRecords field if non-nil, zero value otherwise.

### GetScanRecordsOk

`func (o *SearchResponse) GetScanRecordsOk() (*int32, bool)`

GetScanRecordsOk returns a tuple with the ScanRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanRecords

`func (o *SearchResponse) SetScanRecords(v int32)`

SetScanRecords sets ScanRecords field to given value.


### GetScanSize

`func (o *SearchResponse) GetScanSize() int32`

GetScanSize returns the ScanSize field if non-nil, zero value otherwise.

### GetScanSizeOk

`func (o *SearchResponse) GetScanSizeOk() (*int32, bool)`

GetScanSizeOk returns a tuple with the ScanSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanSize

`func (o *SearchResponse) SetScanSize(v int32)`

SetScanSize sets ScanSize field to given value.


### GetSize

`func (o *SearchResponse) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SearchResponse) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SearchResponse) SetSize(v int64)`

SetSize sets Size field to given value.


### GetTook

`func (o *SearchResponse) GetTook() int32`

GetTook returns the Took field if non-nil, zero value otherwise.

### GetTookOk

`func (o *SearchResponse) GetTookOk() (*int32, bool)`

GetTookOk returns a tuple with the Took field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTook

`func (o *SearchResponse) SetTook(v int32)`

SetTook sets Took field to given value.


### GetTookDetail

`func (o *SearchResponse) GetTookDetail() ResponseTook`

GetTookDetail returns the TookDetail field if non-nil, zero value otherwise.

### GetTookDetailOk

`func (o *SearchResponse) GetTookDetailOk() (*ResponseTook, bool)`

GetTookDetailOk returns a tuple with the TookDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTookDetail

`func (o *SearchResponse) SetTookDetail(v ResponseTook)`

SetTookDetail sets TookDetail field to given value.

### HasTookDetail

`func (o *SearchResponse) HasTookDetail() bool`

HasTookDetail returns a boolean if a field has been set.

### SetTookDetailNil

`func (o *SearchResponse) SetTookDetailNil(b bool)`

 SetTookDetailNil sets the value for TookDetail to be an explicit nil

### UnsetTookDetail
`func (o *SearchResponse) UnsetTookDetail()`

UnsetTookDetail ensures that no value is present for TookDetail, not even an explicit nil
### GetTotal

`func (o *SearchResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SearchResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SearchResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetTraceId

`func (o *SearchResponse) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *SearchResponse) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *SearchResponse) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *SearchResponse) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetWorkGroup

`func (o *SearchResponse) GetWorkGroup() string`

GetWorkGroup returns the WorkGroup field if non-nil, zero value otherwise.

### GetWorkGroupOk

`func (o *SearchResponse) GetWorkGroupOk() (*string, bool)`

GetWorkGroupOk returns a tuple with the WorkGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkGroup

`func (o *SearchResponse) SetWorkGroup(v string)`

SetWorkGroup sets WorkGroup field to given value.

### HasWorkGroup

`func (o *SearchResponse) HasWorkGroup() bool`

HasWorkGroup returns a boolean if a field has been set.

### SetWorkGroupNil

`func (o *SearchResponse) SetWorkGroupNil(b bool)`

 SetWorkGroupNil sets the value for WorkGroup to be an explicit nil

### UnsetWorkGroup
`func (o *SearchResponse) UnsetWorkGroup()`

UnsetWorkGroup ensures that no value is present for WorkGroup, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


