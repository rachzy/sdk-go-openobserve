# ScanStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompressedSize** | **int64** |  | 
**Files** | **int64** |  | 
**IdxScanSize** | **int64** |  | 
**IdxTook** | **int64** |  | 
**OriginalSize** | **int64** |  | 
**QuerierDiskCachedFiles** | **int64** |  | 
**QuerierFiles** | **int64** |  | 
**QuerierMemoryCachedFiles** | **int64** |  | 
**Records** | **int64** |  | 

## Methods

### NewScanStats

`func NewScanStats(compressedSize int64, files int64, idxScanSize int64, idxTook int64, originalSize int64, querierDiskCachedFiles int64, querierFiles int64, querierMemoryCachedFiles int64, records int64, ) *ScanStats`

NewScanStats instantiates a new ScanStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanStatsWithDefaults

`func NewScanStatsWithDefaults() *ScanStats`

NewScanStatsWithDefaults instantiates a new ScanStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompressedSize

`func (o *ScanStats) GetCompressedSize() int64`

GetCompressedSize returns the CompressedSize field if non-nil, zero value otherwise.

### GetCompressedSizeOk

`func (o *ScanStats) GetCompressedSizeOk() (*int64, bool)`

GetCompressedSizeOk returns a tuple with the CompressedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressedSize

`func (o *ScanStats) SetCompressedSize(v int64)`

SetCompressedSize sets CompressedSize field to given value.


### GetFiles

`func (o *ScanStats) GetFiles() int64`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ScanStats) GetFilesOk() (*int64, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ScanStats) SetFiles(v int64)`

SetFiles sets Files field to given value.


### GetIdxScanSize

`func (o *ScanStats) GetIdxScanSize() int64`

GetIdxScanSize returns the IdxScanSize field if non-nil, zero value otherwise.

### GetIdxScanSizeOk

`func (o *ScanStats) GetIdxScanSizeOk() (*int64, bool)`

GetIdxScanSizeOk returns a tuple with the IdxScanSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdxScanSize

`func (o *ScanStats) SetIdxScanSize(v int64)`

SetIdxScanSize sets IdxScanSize field to given value.


### GetIdxTook

`func (o *ScanStats) GetIdxTook() int64`

GetIdxTook returns the IdxTook field if non-nil, zero value otherwise.

### GetIdxTookOk

`func (o *ScanStats) GetIdxTookOk() (*int64, bool)`

GetIdxTookOk returns a tuple with the IdxTook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdxTook

`func (o *ScanStats) SetIdxTook(v int64)`

SetIdxTook sets IdxTook field to given value.


### GetOriginalSize

`func (o *ScanStats) GetOriginalSize() int64`

GetOriginalSize returns the OriginalSize field if non-nil, zero value otherwise.

### GetOriginalSizeOk

`func (o *ScanStats) GetOriginalSizeOk() (*int64, bool)`

GetOriginalSizeOk returns a tuple with the OriginalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSize

`func (o *ScanStats) SetOriginalSize(v int64)`

SetOriginalSize sets OriginalSize field to given value.


### GetQuerierDiskCachedFiles

`func (o *ScanStats) GetQuerierDiskCachedFiles() int64`

GetQuerierDiskCachedFiles returns the QuerierDiskCachedFiles field if non-nil, zero value otherwise.

### GetQuerierDiskCachedFilesOk

`func (o *ScanStats) GetQuerierDiskCachedFilesOk() (*int64, bool)`

GetQuerierDiskCachedFilesOk returns a tuple with the QuerierDiskCachedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerierDiskCachedFiles

`func (o *ScanStats) SetQuerierDiskCachedFiles(v int64)`

SetQuerierDiskCachedFiles sets QuerierDiskCachedFiles field to given value.


### GetQuerierFiles

`func (o *ScanStats) GetQuerierFiles() int64`

GetQuerierFiles returns the QuerierFiles field if non-nil, zero value otherwise.

### GetQuerierFilesOk

`func (o *ScanStats) GetQuerierFilesOk() (*int64, bool)`

GetQuerierFilesOk returns a tuple with the QuerierFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerierFiles

`func (o *ScanStats) SetQuerierFiles(v int64)`

SetQuerierFiles sets QuerierFiles field to given value.


### GetQuerierMemoryCachedFiles

`func (o *ScanStats) GetQuerierMemoryCachedFiles() int64`

GetQuerierMemoryCachedFiles returns the QuerierMemoryCachedFiles field if non-nil, zero value otherwise.

### GetQuerierMemoryCachedFilesOk

`func (o *ScanStats) GetQuerierMemoryCachedFilesOk() (*int64, bool)`

GetQuerierMemoryCachedFilesOk returns a tuple with the QuerierMemoryCachedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerierMemoryCachedFiles

`func (o *ScanStats) SetQuerierMemoryCachedFiles(v int64)`

SetQuerierMemoryCachedFiles sets QuerierMemoryCachedFiles field to given value.


### GetRecords

`func (o *ScanStats) GetRecords() int64`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *ScanStats) GetRecordsOk() (*int64, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *ScanStats) SetRecords(v int64)`

SetRecords sets Records field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


