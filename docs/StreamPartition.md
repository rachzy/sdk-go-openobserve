# StreamPartition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** |  | [optional] 
**Field** | **string** |  | 
**Types** | Pointer to [**StreamPartitionType**](StreamPartitionType.md) |  | [optional] 

## Methods

### NewStreamPartition

`func NewStreamPartition(field string, ) *StreamPartition`

NewStreamPartition instantiates a new StreamPartition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamPartitionWithDefaults

`func NewStreamPartitionWithDefaults() *StreamPartition`

NewStreamPartitionWithDefaults instantiates a new StreamPartition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *StreamPartition) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *StreamPartition) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *StreamPartition) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *StreamPartition) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetField

`func (o *StreamPartition) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *StreamPartition) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *StreamPartition) SetField(v string)`

SetField sets Field field to given value.


### GetTypes

`func (o *StreamPartition) GetTypes() StreamPartitionType`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *StreamPartition) GetTypesOk() (*StreamPartitionType, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *StreamPartition) SetTypes(v StreamPartitionType)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *StreamPartition) HasTypes() bool`

HasTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


