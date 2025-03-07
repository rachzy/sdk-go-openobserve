# Transform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Function** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**NumArgs** | Pointer to **int32** |  | [optional] 
**Params** | Pointer to **string** |  | [optional] 
**Streams** | Pointer to [**[]StreamOrder**](StreamOrder.md) |  | [optional] 
**TransType** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewTransform

`func NewTransform(function string, ) *Transform`

NewTransform instantiates a new Transform object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransformWithDefaults

`func NewTransformWithDefaults() *Transform`

NewTransformWithDefaults instantiates a new Transform object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunction

`func (o *Transform) GetFunction() string`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *Transform) GetFunctionOk() (*string, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *Transform) SetFunction(v string)`

SetFunction sets Function field to given value.


### GetName

`func (o *Transform) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Transform) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Transform) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Transform) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumArgs

`func (o *Transform) GetNumArgs() int32`

GetNumArgs returns the NumArgs field if non-nil, zero value otherwise.

### GetNumArgsOk

`func (o *Transform) GetNumArgsOk() (*int32, bool)`

GetNumArgsOk returns a tuple with the NumArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumArgs

`func (o *Transform) SetNumArgs(v int32)`

SetNumArgs sets NumArgs field to given value.

### HasNumArgs

`func (o *Transform) HasNumArgs() bool`

HasNumArgs returns a boolean if a field has been set.

### GetParams

`func (o *Transform) GetParams() string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *Transform) GetParamsOk() (*string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *Transform) SetParams(v string)`

SetParams sets Params field to given value.

### HasParams

`func (o *Transform) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetStreams

`func (o *Transform) GetStreams() []StreamOrder`

GetStreams returns the Streams field if non-nil, zero value otherwise.

### GetStreamsOk

`func (o *Transform) GetStreamsOk() (*[]StreamOrder, bool)`

GetStreamsOk returns a tuple with the Streams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreams

`func (o *Transform) SetStreams(v []StreamOrder)`

SetStreams sets Streams field to given value.

### HasStreams

`func (o *Transform) HasStreams() bool`

HasStreams returns a boolean if a field has been set.

### SetStreamsNil

`func (o *Transform) SetStreamsNil(b bool)`

 SetStreamsNil sets the value for Streams to be an explicit nil

### UnsetStreams
`func (o *Transform) UnsetStreams()`

UnsetStreams ensures that no value is present for Streams, not even an explicit nil
### GetTransType

`func (o *Transform) GetTransType() int32`

GetTransType returns the TransType field if non-nil, zero value otherwise.

### GetTransTypeOk

`func (o *Transform) GetTransTypeOk() (*int32, bool)`

GetTransTypeOk returns a tuple with the TransType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransType

`func (o *Transform) SetTransType(v int32)`

SetTransType sets TransType field to given value.

### HasTransType

`func (o *Transform) HasTransType() bool`

HasTransType returns a boolean if a field has been set.

### SetTransTypeNil

`func (o *Transform) SetTransTypeNil(b bool)`

 SetTransTypeNil sets the value for TransType to be an explicit nil

### UnsetTransType
`func (o *Transform) UnsetTransType()`

UnsetTransType ensures that no value is present for TransType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


