# Template

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** |  | [optional] 
**IsDefault** | Pointer to **NullableBool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**DestinationType**](DestinationType.md) |  | [optional] 

## Methods

### NewTemplate

`func NewTemplate() *Template`

NewTemplate instantiates a new Template object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateWithDefaults

`func NewTemplateWithDefaults() *Template`

NewTemplateWithDefaults instantiates a new Template object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *Template) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Template) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Template) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *Template) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetIsDefault

`func (o *Template) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *Template) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *Template) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *Template) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *Template) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *Template) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetName

`func (o *Template) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Template) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Template) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Template) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *Template) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Template) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Template) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Template) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *Template) GetType() DestinationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Template) GetTypeOk() (*DestinationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Template) SetType(v DestinationType)`

SetType sets Type field to given value.

### HasType

`func (o *Template) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


