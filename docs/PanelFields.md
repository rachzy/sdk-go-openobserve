# PanelFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | [**[]PanelFilter**](PanelFilter.md) |  | 
**Stream** | **string** |  | 
**StreamType** | [**StreamType**](StreamType.md) |  | 
**X** | [**[]AxisItem**](AxisItem.md) |  | 
**Y** | [**[]AxisItem**](AxisItem.md) |  | 

## Methods

### NewPanelFields

`func NewPanelFields(filter []PanelFilter, stream string, streamType StreamType, x []AxisItem, y []AxisItem, ) *PanelFields`

NewPanelFields instantiates a new PanelFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPanelFieldsWithDefaults

`func NewPanelFieldsWithDefaults() *PanelFields`

NewPanelFieldsWithDefaults instantiates a new PanelFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *PanelFields) GetFilter() []PanelFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *PanelFields) GetFilterOk() (*[]PanelFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *PanelFields) SetFilter(v []PanelFilter)`

SetFilter sets Filter field to given value.


### GetStream

`func (o *PanelFields) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *PanelFields) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *PanelFields) SetStream(v string)`

SetStream sets Stream field to given value.


### GetStreamType

`func (o *PanelFields) GetStreamType() StreamType`

GetStreamType returns the StreamType field if non-nil, zero value otherwise.

### GetStreamTypeOk

`func (o *PanelFields) GetStreamTypeOk() (*StreamType, bool)`

GetStreamTypeOk returns a tuple with the StreamType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamType

`func (o *PanelFields) SetStreamType(v StreamType)`

SetStreamType sets StreamType field to given value.


### GetX

`func (o *PanelFields) GetX() []AxisItem`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *PanelFields) GetXOk() (*[]AxisItem, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *PanelFields) SetX(v []AxisItem)`

SetX sets X field to given value.


### GetY

`func (o *PanelFields) GetY() []AxisItem`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *PanelFields) GetYOk() (*[]AxisItem, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *PanelFields) SetY(v []AxisItem)`

SetY sets Y field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


