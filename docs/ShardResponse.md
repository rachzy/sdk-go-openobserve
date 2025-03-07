# ShardResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Failed** | **int64** |  | 
**Successful** | **int64** |  | 
**Total** | **int64** |  | 

## Methods

### NewShardResponse

`func NewShardResponse(failed int64, successful int64, total int64, ) *ShardResponse`

NewShardResponse instantiates a new ShardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShardResponseWithDefaults

`func NewShardResponseWithDefaults() *ShardResponse`

NewShardResponseWithDefaults instantiates a new ShardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailed

`func (o *ShardResponse) GetFailed() int64`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ShardResponse) GetFailedOk() (*int64, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ShardResponse) SetFailed(v int64)`

SetFailed sets Failed field to given value.


### GetSuccessful

`func (o *ShardResponse) GetSuccessful() int64`

GetSuccessful returns the Successful field if non-nil, zero value otherwise.

### GetSuccessfulOk

`func (o *ShardResponse) GetSuccessfulOk() (*int64, bool)`

GetSuccessfulOk returns a tuple with the Successful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessful

`func (o *ShardResponse) SetSuccessful(v int64)`

SetSuccessful sets Successful field to given value.


### GetTotal

`func (o *ShardResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ShardResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ShardResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


