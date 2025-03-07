# IngestionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** |  | 
**Error** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**[]StreamStatus**](StreamStatus.md) |  | [optional] 

## Methods

### NewIngestionResponse

`func NewIngestionResponse(code int32, ) *IngestionResponse`

NewIngestionResponse instantiates a new IngestionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestionResponseWithDefaults

`func NewIngestionResponseWithDefaults() *IngestionResponse`

NewIngestionResponseWithDefaults instantiates a new IngestionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *IngestionResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *IngestionResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *IngestionResponse) SetCode(v int32)`

SetCode sets Code field to given value.


### GetError

`func (o *IngestionResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *IngestionResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *IngestionResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *IngestionResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *IngestionResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *IngestionResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetStatus

`func (o *IngestionResponse) GetStatus() []StreamStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IngestionResponse) GetStatusOk() (*[]StreamStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IngestionResponse) SetStatus(v []StreamStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IngestionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


