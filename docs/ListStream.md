# ListStream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**List** | [**[]Stream**](Stream.md) |  | 
**Total** | **int32** |  | 

## Methods

### NewListStream

`func NewListStream(list []Stream, total int32, ) *ListStream`

NewListStream instantiates a new ListStream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListStreamWithDefaults

`func NewListStreamWithDefaults() *ListStream`

NewListStreamWithDefaults instantiates a new ListStream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetList

`func (o *ListStream) GetList() []Stream`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *ListStream) GetListOk() (*[]Stream, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *ListStream) SetList(v []Stream)`

SetList sets List field to given value.


### GetTotal

`func (o *ListStream) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListStream) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListStream) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


