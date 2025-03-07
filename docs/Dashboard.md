# Dashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] 
**DashboardId** | Pointer to **string** |  | [optional] 
**Description** | **string** |  | 
**Layouts** | Pointer to [**[]Layout**](Layout.md) |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Panels** | Pointer to [**[]Panel**](Panel.md) |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Title** | **string** |  | 
**Variables** | Pointer to [**NullableVariables**](Variables.md) |  | [optional] 

## Methods

### NewDashboard

`func NewDashboard(description string, title string, ) *Dashboard`

NewDashboard instantiates a new Dashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardWithDefaults

`func NewDashboardWithDefaults() *Dashboard`

NewDashboardWithDefaults instantiates a new Dashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *Dashboard) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Dashboard) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Dashboard) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Dashboard) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDashboardId

`func (o *Dashboard) GetDashboardId() string`

GetDashboardId returns the DashboardId field if non-nil, zero value otherwise.

### GetDashboardIdOk

`func (o *Dashboard) GetDashboardIdOk() (*string, bool)`

GetDashboardIdOk returns a tuple with the DashboardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardId

`func (o *Dashboard) SetDashboardId(v string)`

SetDashboardId sets DashboardId field to given value.

### HasDashboardId

`func (o *Dashboard) HasDashboardId() bool`

HasDashboardId returns a boolean if a field has been set.

### GetDescription

`func (o *Dashboard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Dashboard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Dashboard) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLayouts

`func (o *Dashboard) GetLayouts() []Layout`

GetLayouts returns the Layouts field if non-nil, zero value otherwise.

### GetLayoutsOk

`func (o *Dashboard) GetLayoutsOk() (*[]Layout, bool)`

GetLayoutsOk returns a tuple with the Layouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayouts

`func (o *Dashboard) SetLayouts(v []Layout)`

SetLayouts sets Layouts field to given value.

### HasLayouts

`func (o *Dashboard) HasLayouts() bool`

HasLayouts returns a boolean if a field has been set.

### SetLayoutsNil

`func (o *Dashboard) SetLayoutsNil(b bool)`

 SetLayoutsNil sets the value for Layouts to be an explicit nil

### UnsetLayouts
`func (o *Dashboard) UnsetLayouts()`

UnsetLayouts ensures that no value is present for Layouts, not even an explicit nil
### GetOwner

`func (o *Dashboard) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Dashboard) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Dashboard) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Dashboard) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPanels

`func (o *Dashboard) GetPanels() []Panel`

GetPanels returns the Panels field if non-nil, zero value otherwise.

### GetPanelsOk

`func (o *Dashboard) GetPanelsOk() (*[]Panel, bool)`

GetPanelsOk returns a tuple with the Panels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanels

`func (o *Dashboard) SetPanels(v []Panel)`

SetPanels sets Panels field to given value.

### HasPanels

`func (o *Dashboard) HasPanels() bool`

HasPanels returns a boolean if a field has been set.

### GetRole

`func (o *Dashboard) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Dashboard) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Dashboard) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *Dashboard) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTitle

`func (o *Dashboard) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Dashboard) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Dashboard) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetVariables

`func (o *Dashboard) GetVariables() Variables`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *Dashboard) GetVariablesOk() (*Variables, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *Dashboard) SetVariables(v Variables)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *Dashboard) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *Dashboard) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *Dashboard) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


