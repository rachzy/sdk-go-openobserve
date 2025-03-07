# RecordStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Failed** | **int32** |  | 
**Successful** | **int32** |  | 

## Methods

### NewRecordStatus

`func NewRecordStatus(failed int32, successful int32, ) *RecordStatus`

NewRecordStatus instantiates a new RecordStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordStatusWithDefaults

`func NewRecordStatusWithDefaults() *RecordStatus`

NewRecordStatusWithDefaults instantiates a new RecordStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *RecordStatus) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RecordStatus) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RecordStatus) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *RecordStatus) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFailed

`func (o *RecordStatus) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *RecordStatus) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *RecordStatus) SetFailed(v int32)`

SetFailed sets Failed field to given value.


### GetSuccessful

`func (o *RecordStatus) GetSuccessful() int32`

GetSuccessful returns the Successful field if non-nil, zero value otherwise.

### GetSuccessfulOk

`func (o *RecordStatus) GetSuccessfulOk() (*int32, bool)`

GetSuccessfulOk returns a tuple with the Successful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessful

`func (o *RecordStatus) SetSuccessful(v int32)`

SetSuccessful sets Successful field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


