# Panel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**PanelConfig**](PanelConfig.md) |  | 
**CustomQuery** | **bool** |  | 
**Fields** | [**PanelFields**](PanelFields.md) |  | 
**Id** | **string** |  | 
**Query** | **string** |  | 
**QueryType** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewPanel

`func NewPanel(config PanelConfig, customQuery bool, fields PanelFields, id string, query string, type_ string, ) *Panel`

NewPanel instantiates a new Panel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPanelWithDefaults

`func NewPanelWithDefaults() *Panel`

NewPanelWithDefaults instantiates a new Panel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *Panel) GetConfig() PanelConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Panel) GetConfigOk() (*PanelConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Panel) SetConfig(v PanelConfig)`

SetConfig sets Config field to given value.


### GetCustomQuery

`func (o *Panel) GetCustomQuery() bool`

GetCustomQuery returns the CustomQuery field if non-nil, zero value otherwise.

### GetCustomQueryOk

`func (o *Panel) GetCustomQueryOk() (*bool, bool)`

GetCustomQueryOk returns a tuple with the CustomQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomQuery

`func (o *Panel) SetCustomQuery(v bool)`

SetCustomQuery sets CustomQuery field to given value.


### GetFields

`func (o *Panel) GetFields() PanelFields`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *Panel) GetFieldsOk() (*PanelFields, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *Panel) SetFields(v PanelFields)`

SetFields sets Fields field to given value.


### GetId

`func (o *Panel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Panel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Panel) SetId(v string)`

SetId sets Id field to given value.


### GetQuery

`func (o *Panel) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Panel) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Panel) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetQueryType

`func (o *Panel) GetQueryType() string`

GetQueryType returns the QueryType field if non-nil, zero value otherwise.

### GetQueryTypeOk

`func (o *Panel) GetQueryTypeOk() (*string, bool)`

GetQueryTypeOk returns a tuple with the QueryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryType

`func (o *Panel) SetQueryType(v string)`

SetQueryType sets QueryType field to given value.

### HasQueryType

`func (o *Panel) HasQueryType() bool`

HasQueryType returns a boolean if a field has been set.

### GetType

`func (o *Panel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Panel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Panel) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


