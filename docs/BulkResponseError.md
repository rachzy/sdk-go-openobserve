# BulkResponseError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **string** |  | 
**IndexUuid** | **string** |  | 
**Reason** | **string** |  | 
**Shard** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewBulkResponseError

`func NewBulkResponseError(index string, indexUuid string, reason string, shard string, type_ string, ) *BulkResponseError`

NewBulkResponseError instantiates a new BulkResponseError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkResponseErrorWithDefaults

`func NewBulkResponseErrorWithDefaults() *BulkResponseError`

NewBulkResponseErrorWithDefaults instantiates a new BulkResponseError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *BulkResponseError) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *BulkResponseError) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *BulkResponseError) SetIndex(v string)`

SetIndex sets Index field to given value.


### GetIndexUuid

`func (o *BulkResponseError) GetIndexUuid() string`

GetIndexUuid returns the IndexUuid field if non-nil, zero value otherwise.

### GetIndexUuidOk

`func (o *BulkResponseError) GetIndexUuidOk() (*string, bool)`

GetIndexUuidOk returns a tuple with the IndexUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexUuid

`func (o *BulkResponseError) SetIndexUuid(v string)`

SetIndexUuid sets IndexUuid field to given value.


### GetReason

`func (o *BulkResponseError) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BulkResponseError) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BulkResponseError) SetReason(v string)`

SetReason sets Reason field to given value.


### GetShard

`func (o *BulkResponseError) GetShard() string`

GetShard returns the Shard field if non-nil, zero value otherwise.

### GetShardOk

`func (o *BulkResponseError) GetShardOk() (*string, bool)`

GetShardOk returns a tuple with the Shard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShard

`func (o *BulkResponseError) SetShard(v string)`

SetShard sets Shard field to given value.


### GetType

`func (o *BulkResponseError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BulkResponseError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BulkResponseError) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


