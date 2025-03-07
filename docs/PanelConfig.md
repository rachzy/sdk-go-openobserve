# PanelConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**LegendsPosition** | Pointer to **NullableString** |  | [optional] 
**PromqlLegend** | Pointer to **NullableString** |  | [optional] 
**ShowLegends** | **bool** |  | 
**Title** | **string** |  | 
**Unit** | Pointer to **NullableString** |  | [optional] 
**UnitCustom** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPanelConfig

`func NewPanelConfig(description string, showLegends bool, title string, ) *PanelConfig`

NewPanelConfig instantiates a new PanelConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPanelConfigWithDefaults

`func NewPanelConfigWithDefaults() *PanelConfig`

NewPanelConfigWithDefaults instantiates a new PanelConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PanelConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PanelConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PanelConfig) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLegendsPosition

`func (o *PanelConfig) GetLegendsPosition() string`

GetLegendsPosition returns the LegendsPosition field if non-nil, zero value otherwise.

### GetLegendsPositionOk

`func (o *PanelConfig) GetLegendsPositionOk() (*string, bool)`

GetLegendsPositionOk returns a tuple with the LegendsPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegendsPosition

`func (o *PanelConfig) SetLegendsPosition(v string)`

SetLegendsPosition sets LegendsPosition field to given value.

### HasLegendsPosition

`func (o *PanelConfig) HasLegendsPosition() bool`

HasLegendsPosition returns a boolean if a field has been set.

### SetLegendsPositionNil

`func (o *PanelConfig) SetLegendsPositionNil(b bool)`

 SetLegendsPositionNil sets the value for LegendsPosition to be an explicit nil

### UnsetLegendsPosition
`func (o *PanelConfig) UnsetLegendsPosition()`

UnsetLegendsPosition ensures that no value is present for LegendsPosition, not even an explicit nil
### GetPromqlLegend

`func (o *PanelConfig) GetPromqlLegend() string`

GetPromqlLegend returns the PromqlLegend field if non-nil, zero value otherwise.

### GetPromqlLegendOk

`func (o *PanelConfig) GetPromqlLegendOk() (*string, bool)`

GetPromqlLegendOk returns a tuple with the PromqlLegend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromqlLegend

`func (o *PanelConfig) SetPromqlLegend(v string)`

SetPromqlLegend sets PromqlLegend field to given value.

### HasPromqlLegend

`func (o *PanelConfig) HasPromqlLegend() bool`

HasPromqlLegend returns a boolean if a field has been set.

### SetPromqlLegendNil

`func (o *PanelConfig) SetPromqlLegendNil(b bool)`

 SetPromqlLegendNil sets the value for PromqlLegend to be an explicit nil

### UnsetPromqlLegend
`func (o *PanelConfig) UnsetPromqlLegend()`

UnsetPromqlLegend ensures that no value is present for PromqlLegend, not even an explicit nil
### GetShowLegends

`func (o *PanelConfig) GetShowLegends() bool`

GetShowLegends returns the ShowLegends field if non-nil, zero value otherwise.

### GetShowLegendsOk

`func (o *PanelConfig) GetShowLegendsOk() (*bool, bool)`

GetShowLegendsOk returns a tuple with the ShowLegends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLegends

`func (o *PanelConfig) SetShowLegends(v bool)`

SetShowLegends sets ShowLegends field to given value.


### GetTitle

`func (o *PanelConfig) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PanelConfig) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PanelConfig) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUnit

`func (o *PanelConfig) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *PanelConfig) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *PanelConfig) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *PanelConfig) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *PanelConfig) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *PanelConfig) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetUnitCustom

`func (o *PanelConfig) GetUnitCustom() string`

GetUnitCustom returns the UnitCustom field if non-nil, zero value otherwise.

### GetUnitCustomOk

`func (o *PanelConfig) GetUnitCustomOk() (*string, bool)`

GetUnitCustomOk returns a tuple with the UnitCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitCustom

`func (o *PanelConfig) SetUnitCustom(v string)`

SetUnitCustom sets UnitCustom field to given value.

### HasUnitCustom

`func (o *PanelConfig) HasUnitCustom() bool`

HasUnitCustom returns a boolean if a field has been set.

### SetUnitCustomNil

`func (o *PanelConfig) SetUnitCustomNil(b bool)`

 SetUnitCustomNil sets the value for UnitCustom to be an explicit nil

### UnsetUnitCustom
`func (o *PanelConfig) UnsetUnitCustom()`

UnsetUnitCustom ensures that no value is present for UnitCustom, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


