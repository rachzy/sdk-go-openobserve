# StreamOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyBeforeFlattening** | Pointer to **bool** |  | [optional] 
**IsRemoved** | Pointer to **bool** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**Stream** | Pointer to **string** |  | [optional] 
**StreamType** | Pointer to [**StreamType**](StreamType.md) |  | [optional] 

## Methods

### NewStreamOrder

`func NewStreamOrder() *StreamOrder`

NewStreamOrder instantiates a new StreamOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamOrderWithDefaults

`func NewStreamOrderWithDefaults() *StreamOrder`

NewStreamOrderWithDefaults instantiates a new StreamOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplyBeforeFlattening

`func (o *StreamOrder) GetApplyBeforeFlattening() bool`

GetApplyBeforeFlattening returns the ApplyBeforeFlattening field if non-nil, zero value otherwise.

### GetApplyBeforeFlatteningOk

`func (o *StreamOrder) GetApplyBeforeFlatteningOk() (*bool, bool)`

GetApplyBeforeFlatteningOk returns a tuple with the ApplyBeforeFlattening field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyBeforeFlattening

`func (o *StreamOrder) SetApplyBeforeFlattening(v bool)`

SetApplyBeforeFlattening sets ApplyBeforeFlattening field to given value.

### HasApplyBeforeFlattening

`func (o *StreamOrder) HasApplyBeforeFlattening() bool`

HasApplyBeforeFlattening returns a boolean if a field has been set.

### GetIsRemoved

`func (o *StreamOrder) GetIsRemoved() bool`

GetIsRemoved returns the IsRemoved field if non-nil, zero value otherwise.

### GetIsRemovedOk

`func (o *StreamOrder) GetIsRemovedOk() (*bool, bool)`

GetIsRemovedOk returns a tuple with the IsRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemoved

`func (o *StreamOrder) SetIsRemoved(v bool)`

SetIsRemoved sets IsRemoved field to given value.

### HasIsRemoved

`func (o *StreamOrder) HasIsRemoved() bool`

HasIsRemoved returns a boolean if a field has been set.

### GetOrder

`func (o *StreamOrder) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *StreamOrder) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *StreamOrder) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *StreamOrder) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetStream

`func (o *StreamOrder) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *StreamOrder) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *StreamOrder) SetStream(v string)`

SetStream sets Stream field to given value.

### HasStream

`func (o *StreamOrder) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetStreamType

`func (o *StreamOrder) GetStreamType() StreamType`

GetStreamType returns the StreamType field if non-nil, zero value otherwise.

### GetStreamTypeOk

`func (o *StreamOrder) GetStreamTypeOk() (*StreamType, bool)`

GetStreamTypeOk returns a tuple with the StreamType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamType

`func (o *StreamOrder) SetStreamType(v StreamType)`

SetStreamType sets StreamType field to given value.

### HasStreamType

`func (o *StreamOrder) HasStreamType() bool`

HasStreamType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


