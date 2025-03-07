# BulkResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Index** | **string** |  | 
**PrimaryTerm** | Pointer to **NullableInt64** |  | [optional] 
**SeqNo** | Pointer to **NullableInt64** |  | [optional] 
**Shards** | Pointer to [**NullableShardResponse**](ShardResponse.md) |  | [optional] 
**Version** | Pointer to **NullableInt64** |  | [optional] 
**Error** | Pointer to [**NullableBulkResponseError**](BulkResponseError.md) |  | [optional] 
**OriginalRecord** | Pointer to **map[string]interface{}** |  | [optional] 
**Result** | Pointer to **NullableString** |  | [optional] 
**Status** | **int64** |  | 

## Methods

### NewBulkResponseItem

`func NewBulkResponseItem(id string, index string, status int64, ) *BulkResponseItem`

NewBulkResponseItem instantiates a new BulkResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkResponseItemWithDefaults

`func NewBulkResponseItemWithDefaults() *BulkResponseItem`

NewBulkResponseItemWithDefaults instantiates a new BulkResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkResponseItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkResponseItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkResponseItem) SetId(v string)`

SetId sets Id field to given value.


### GetIndex

`func (o *BulkResponseItem) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *BulkResponseItem) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *BulkResponseItem) SetIndex(v string)`

SetIndex sets Index field to given value.


### GetPrimaryTerm

`func (o *BulkResponseItem) GetPrimaryTerm() int64`

GetPrimaryTerm returns the PrimaryTerm field if non-nil, zero value otherwise.

### GetPrimaryTermOk

`func (o *BulkResponseItem) GetPrimaryTermOk() (*int64, bool)`

GetPrimaryTermOk returns a tuple with the PrimaryTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTerm

`func (o *BulkResponseItem) SetPrimaryTerm(v int64)`

SetPrimaryTerm sets PrimaryTerm field to given value.

### HasPrimaryTerm

`func (o *BulkResponseItem) HasPrimaryTerm() bool`

HasPrimaryTerm returns a boolean if a field has been set.

### SetPrimaryTermNil

`func (o *BulkResponseItem) SetPrimaryTermNil(b bool)`

 SetPrimaryTermNil sets the value for PrimaryTerm to be an explicit nil

### UnsetPrimaryTerm
`func (o *BulkResponseItem) UnsetPrimaryTerm()`

UnsetPrimaryTerm ensures that no value is present for PrimaryTerm, not even an explicit nil
### GetSeqNo

`func (o *BulkResponseItem) GetSeqNo() int64`

GetSeqNo returns the SeqNo field if non-nil, zero value otherwise.

### GetSeqNoOk

`func (o *BulkResponseItem) GetSeqNoOk() (*int64, bool)`

GetSeqNoOk returns a tuple with the SeqNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqNo

`func (o *BulkResponseItem) SetSeqNo(v int64)`

SetSeqNo sets SeqNo field to given value.

### HasSeqNo

`func (o *BulkResponseItem) HasSeqNo() bool`

HasSeqNo returns a boolean if a field has been set.

### SetSeqNoNil

`func (o *BulkResponseItem) SetSeqNoNil(b bool)`

 SetSeqNoNil sets the value for SeqNo to be an explicit nil

### UnsetSeqNo
`func (o *BulkResponseItem) UnsetSeqNo()`

UnsetSeqNo ensures that no value is present for SeqNo, not even an explicit nil
### GetShards

`func (o *BulkResponseItem) GetShards() ShardResponse`

GetShards returns the Shards field if non-nil, zero value otherwise.

### GetShardsOk

`func (o *BulkResponseItem) GetShardsOk() (*ShardResponse, bool)`

GetShardsOk returns a tuple with the Shards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShards

`func (o *BulkResponseItem) SetShards(v ShardResponse)`

SetShards sets Shards field to given value.

### HasShards

`func (o *BulkResponseItem) HasShards() bool`

HasShards returns a boolean if a field has been set.

### SetShardsNil

`func (o *BulkResponseItem) SetShardsNil(b bool)`

 SetShardsNil sets the value for Shards to be an explicit nil

### UnsetShards
`func (o *BulkResponseItem) UnsetShards()`

UnsetShards ensures that no value is present for Shards, not even an explicit nil
### GetVersion

`func (o *BulkResponseItem) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BulkResponseItem) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BulkResponseItem) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BulkResponseItem) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *BulkResponseItem) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *BulkResponseItem) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetError

`func (o *BulkResponseItem) GetError() BulkResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BulkResponseItem) GetErrorOk() (*BulkResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BulkResponseItem) SetError(v BulkResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *BulkResponseItem) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *BulkResponseItem) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *BulkResponseItem) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetOriginalRecord

`func (o *BulkResponseItem) GetOriginalRecord() map[string]interface{}`

GetOriginalRecord returns the OriginalRecord field if non-nil, zero value otherwise.

### GetOriginalRecordOk

`func (o *BulkResponseItem) GetOriginalRecordOk() (*map[string]interface{}, bool)`

GetOriginalRecordOk returns a tuple with the OriginalRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalRecord

`func (o *BulkResponseItem) SetOriginalRecord(v map[string]interface{})`

SetOriginalRecord sets OriginalRecord field to given value.

### HasOriginalRecord

`func (o *BulkResponseItem) HasOriginalRecord() bool`

HasOriginalRecord returns a boolean if a field has been set.

### GetResult

`func (o *BulkResponseItem) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *BulkResponseItem) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *BulkResponseItem) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *BulkResponseItem) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *BulkResponseItem) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *BulkResponseItem) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetStatus

`func (o *BulkResponseItem) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkResponseItem) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkResponseItem) SetStatus(v int64)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


