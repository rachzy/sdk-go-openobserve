# CancelQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSuccess** | **bool** |  | 
**TraceId** | **string** |  | 

## Methods

### NewCancelQueryResponse

`func NewCancelQueryResponse(isSuccess bool, traceId string, ) *CancelQueryResponse`

NewCancelQueryResponse instantiates a new CancelQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelQueryResponseWithDefaults

`func NewCancelQueryResponseWithDefaults() *CancelQueryResponse`

NewCancelQueryResponseWithDefaults instantiates a new CancelQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSuccess

`func (o *CancelQueryResponse) GetIsSuccess() bool`

GetIsSuccess returns the IsSuccess field if non-nil, zero value otherwise.

### GetIsSuccessOk

`func (o *CancelQueryResponse) GetIsSuccessOk() (*bool, bool)`

GetIsSuccessOk returns a tuple with the IsSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuccess

`func (o *CancelQueryResponse) SetIsSuccess(v bool)`

SetIsSuccess sets IsSuccess field to given value.


### GetTraceId

`func (o *CancelQueryResponse) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *CancelQueryResponse) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *CancelQueryResponse) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


