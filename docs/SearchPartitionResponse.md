# SearchPartitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompressedSize** | **int32** |  | 
**FileNum** | **int32** |  | 
**HistogramInterval** | Pointer to **NullableInt64** |  | [optional] 
**MaxQueryRange** | **int64** |  | 
**OrderBy** | [**OrderBy**](OrderBy.md) |  | 
**OriginalSize** | **int32** |  | 
**Partitions** | **[][]int64** |  | 
**Records** | **int32** |  | 
**StreamingAggs** | **bool** |  | 
**StreamingId** | Pointer to **NullableString** |  | [optional] 
**StreamingOutput** | **bool** |  | 
**TraceId** | **string** |  | 

## Methods

### NewSearchPartitionResponse

`func NewSearchPartitionResponse(compressedSize int32, fileNum int32, maxQueryRange int64, orderBy OrderBy, originalSize int32, partitions [][]int64, records int32, streamingAggs bool, streamingOutput bool, traceId string, ) *SearchPartitionResponse`

NewSearchPartitionResponse instantiates a new SearchPartitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchPartitionResponseWithDefaults

`func NewSearchPartitionResponseWithDefaults() *SearchPartitionResponse`

NewSearchPartitionResponseWithDefaults instantiates a new SearchPartitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompressedSize

`func (o *SearchPartitionResponse) GetCompressedSize() int32`

GetCompressedSize returns the CompressedSize field if non-nil, zero value otherwise.

### GetCompressedSizeOk

`func (o *SearchPartitionResponse) GetCompressedSizeOk() (*int32, bool)`

GetCompressedSizeOk returns a tuple with the CompressedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressedSize

`func (o *SearchPartitionResponse) SetCompressedSize(v int32)`

SetCompressedSize sets CompressedSize field to given value.


### GetFileNum

`func (o *SearchPartitionResponse) GetFileNum() int32`

GetFileNum returns the FileNum field if non-nil, zero value otherwise.

### GetFileNumOk

`func (o *SearchPartitionResponse) GetFileNumOk() (*int32, bool)`

GetFileNumOk returns a tuple with the FileNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileNum

`func (o *SearchPartitionResponse) SetFileNum(v int32)`

SetFileNum sets FileNum field to given value.


### GetHistogramInterval

`func (o *SearchPartitionResponse) GetHistogramInterval() int64`

GetHistogramInterval returns the HistogramInterval field if non-nil, zero value otherwise.

### GetHistogramIntervalOk

`func (o *SearchPartitionResponse) GetHistogramIntervalOk() (*int64, bool)`

GetHistogramIntervalOk returns a tuple with the HistogramInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogramInterval

`func (o *SearchPartitionResponse) SetHistogramInterval(v int64)`

SetHistogramInterval sets HistogramInterval field to given value.

### HasHistogramInterval

`func (o *SearchPartitionResponse) HasHistogramInterval() bool`

HasHistogramInterval returns a boolean if a field has been set.

### SetHistogramIntervalNil

`func (o *SearchPartitionResponse) SetHistogramIntervalNil(b bool)`

 SetHistogramIntervalNil sets the value for HistogramInterval to be an explicit nil

### UnsetHistogramInterval
`func (o *SearchPartitionResponse) UnsetHistogramInterval()`

UnsetHistogramInterval ensures that no value is present for HistogramInterval, not even an explicit nil
### GetMaxQueryRange

`func (o *SearchPartitionResponse) GetMaxQueryRange() int64`

GetMaxQueryRange returns the MaxQueryRange field if non-nil, zero value otherwise.

### GetMaxQueryRangeOk

`func (o *SearchPartitionResponse) GetMaxQueryRangeOk() (*int64, bool)`

GetMaxQueryRangeOk returns a tuple with the MaxQueryRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQueryRange

`func (o *SearchPartitionResponse) SetMaxQueryRange(v int64)`

SetMaxQueryRange sets MaxQueryRange field to given value.


### GetOrderBy

`func (o *SearchPartitionResponse) GetOrderBy() OrderBy`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *SearchPartitionResponse) GetOrderByOk() (*OrderBy, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *SearchPartitionResponse) SetOrderBy(v OrderBy)`

SetOrderBy sets OrderBy field to given value.


### GetOriginalSize

`func (o *SearchPartitionResponse) GetOriginalSize() int32`

GetOriginalSize returns the OriginalSize field if non-nil, zero value otherwise.

### GetOriginalSizeOk

`func (o *SearchPartitionResponse) GetOriginalSizeOk() (*int32, bool)`

GetOriginalSizeOk returns a tuple with the OriginalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSize

`func (o *SearchPartitionResponse) SetOriginalSize(v int32)`

SetOriginalSize sets OriginalSize field to given value.


### GetPartitions

`func (o *SearchPartitionResponse) GetPartitions() [][]int64`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *SearchPartitionResponse) GetPartitionsOk() (*[][]int64, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *SearchPartitionResponse) SetPartitions(v [][]int64)`

SetPartitions sets Partitions field to given value.


### GetRecords

`func (o *SearchPartitionResponse) GetRecords() int32`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *SearchPartitionResponse) GetRecordsOk() (*int32, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *SearchPartitionResponse) SetRecords(v int32)`

SetRecords sets Records field to given value.


### GetStreamingAggs

`func (o *SearchPartitionResponse) GetStreamingAggs() bool`

GetStreamingAggs returns the StreamingAggs field if non-nil, zero value otherwise.

### GetStreamingAggsOk

`func (o *SearchPartitionResponse) GetStreamingAggsOk() (*bool, bool)`

GetStreamingAggsOk returns a tuple with the StreamingAggs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamingAggs

`func (o *SearchPartitionResponse) SetStreamingAggs(v bool)`

SetStreamingAggs sets StreamingAggs field to given value.


### GetStreamingId

`func (o *SearchPartitionResponse) GetStreamingId() string`

GetStreamingId returns the StreamingId field if non-nil, zero value otherwise.

### GetStreamingIdOk

`func (o *SearchPartitionResponse) GetStreamingIdOk() (*string, bool)`

GetStreamingIdOk returns a tuple with the StreamingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamingId

`func (o *SearchPartitionResponse) SetStreamingId(v string)`

SetStreamingId sets StreamingId field to given value.

### HasStreamingId

`func (o *SearchPartitionResponse) HasStreamingId() bool`

HasStreamingId returns a boolean if a field has been set.

### SetStreamingIdNil

`func (o *SearchPartitionResponse) SetStreamingIdNil(b bool)`

 SetStreamingIdNil sets the value for StreamingId to be an explicit nil

### UnsetStreamingId
`func (o *SearchPartitionResponse) UnsetStreamingId()`

UnsetStreamingId ensures that no value is present for StreamingId, not even an explicit nil
### GetStreamingOutput

`func (o *SearchPartitionResponse) GetStreamingOutput() bool`

GetStreamingOutput returns the StreamingOutput field if non-nil, zero value otherwise.

### GetStreamingOutputOk

`func (o *SearchPartitionResponse) GetStreamingOutputOk() (*bool, bool)`

GetStreamingOutputOk returns a tuple with the StreamingOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamingOutput

`func (o *SearchPartitionResponse) SetStreamingOutput(v bool)`

SetStreamingOutput sets StreamingOutput field to given value.


### GetTraceId

`func (o *SearchPartitionResponse) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *SearchPartitionResponse) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *SearchPartitionResponse) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


