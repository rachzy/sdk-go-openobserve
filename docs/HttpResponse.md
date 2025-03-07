# HttpResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** |  | 
**ErrorDetail** | Pointer to **NullableString** |  | [optional] 
**Message** | **string** |  | 
**TraceId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewHttpResponse

`func NewHttpResponse(code int32, message string, ) *HttpResponse`

NewHttpResponse instantiates a new HttpResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpResponseWithDefaults

`func NewHttpResponseWithDefaults() *HttpResponse`

NewHttpResponseWithDefaults instantiates a new HttpResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *HttpResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *HttpResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *HttpResponse) SetCode(v int32)`

SetCode sets Code field to given value.


### GetErrorDetail

`func (o *HttpResponse) GetErrorDetail() string`

GetErrorDetail returns the ErrorDetail field if non-nil, zero value otherwise.

### GetErrorDetailOk

`func (o *HttpResponse) GetErrorDetailOk() (*string, bool)`

GetErrorDetailOk returns a tuple with the ErrorDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetail

`func (o *HttpResponse) SetErrorDetail(v string)`

SetErrorDetail sets ErrorDetail field to given value.

### HasErrorDetail

`func (o *HttpResponse) HasErrorDetail() bool`

HasErrorDetail returns a boolean if a field has been set.

### SetErrorDetailNil

`func (o *HttpResponse) SetErrorDetailNil(b bool)`

 SetErrorDetailNil sets the value for ErrorDetail to be an explicit nil

### UnsetErrorDetail
`func (o *HttpResponse) UnsetErrorDetail()`

UnsetErrorDetail ensures that no value is present for ErrorDetail, not even an explicit nil
### GetMessage

`func (o *HttpResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HttpResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HttpResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTraceId

`func (o *HttpResponse) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *HttpResponse) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *HttpResponse) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *HttpResponse) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### SetTraceIdNil

`func (o *HttpResponse) SetTraceIdNil(b bool)`

 SetTraceIdNil sets the value for TraceId to be an explicit nil

### UnsetTraceId
`func (o *HttpResponse) UnsetTraceId()`

UnsetTraceId ensures that no value is present for TraceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


