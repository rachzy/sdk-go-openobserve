# BulkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | **bool** |  | 
**Items** | [**[]map[string]BulkResponseItem**](map[string]BulkResponseItem.md) |  | 
**Took** | **int32** |  | 

## Methods

### NewBulkResponse

`func NewBulkResponse(errors bool, items []map[string]BulkResponseItem, took int32, ) *BulkResponse`

NewBulkResponse instantiates a new BulkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkResponseWithDefaults

`func NewBulkResponseWithDefaults() *BulkResponse`

NewBulkResponseWithDefaults instantiates a new BulkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *BulkResponse) GetErrors() bool`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BulkResponse) GetErrorsOk() (*bool, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BulkResponse) SetErrors(v bool)`

SetErrors sets Errors field to given value.


### GetItems

`func (o *BulkResponse) GetItems() []map[string]BulkResponseItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BulkResponse) GetItemsOk() (*[]map[string]BulkResponseItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BulkResponse) SetItems(v []map[string]BulkResponseItem)`

SetItems sets Items field to given value.


### GetTook

`func (o *BulkResponse) GetTook() int32`

GetTook returns the Took field if non-nil, zero value otherwise.

### GetTookOk

`func (o *BulkResponse) GetTookOk() (*int32, bool)`

GetTookOk returns a tuple with the Took field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTook

`func (o *BulkResponse) SetTook(v int32)`

SetTook sets Took field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


