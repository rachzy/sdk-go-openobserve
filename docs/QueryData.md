# QueryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** |  | 
**MaxRecordSize** | Pointer to **NullableInt64** |  | [optional] 
**Stream** | **string** |  | 
**StreamType** | [**StreamType**](StreamType.md) |  | 

## Methods

### NewQueryData

`func NewQueryData(field string, stream string, streamType StreamType, ) *QueryData`

NewQueryData instantiates a new QueryData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryDataWithDefaults

`func NewQueryDataWithDefaults() *QueryData`

NewQueryDataWithDefaults instantiates a new QueryData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *QueryData) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *QueryData) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *QueryData) SetField(v string)`

SetField sets Field field to given value.


### GetMaxRecordSize

`func (o *QueryData) GetMaxRecordSize() int64`

GetMaxRecordSize returns the MaxRecordSize field if non-nil, zero value otherwise.

### GetMaxRecordSizeOk

`func (o *QueryData) GetMaxRecordSizeOk() (*int64, bool)`

GetMaxRecordSizeOk returns a tuple with the MaxRecordSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRecordSize

`func (o *QueryData) SetMaxRecordSize(v int64)`

SetMaxRecordSize sets MaxRecordSize field to given value.

### HasMaxRecordSize

`func (o *QueryData) HasMaxRecordSize() bool`

HasMaxRecordSize returns a boolean if a field has been set.

### SetMaxRecordSizeNil

`func (o *QueryData) SetMaxRecordSizeNil(b bool)`

 SetMaxRecordSizeNil sets the value for MaxRecordSize to be an explicit nil

### UnsetMaxRecordSize
`func (o *QueryData) UnsetMaxRecordSize()`

UnsetMaxRecordSize ensures that no value is present for MaxRecordSize, not even an explicit nil
### GetStream

`func (o *QueryData) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *QueryData) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *QueryData) SetStream(v string)`

SetStream sets Stream field to given value.


### GetStreamType

`func (o *QueryData) GetStreamType() StreamType`

GetStreamType returns the StreamType field if non-nil, zero value otherwise.

### GetStreamTypeOk

`func (o *QueryData) GetStreamTypeOk() (*StreamType, bool)`

GetStreamTypeOk returns a tuple with the StreamType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamType

`func (o *QueryData) SetStreamType(v StreamType)`

SetStreamType sets StreamType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


