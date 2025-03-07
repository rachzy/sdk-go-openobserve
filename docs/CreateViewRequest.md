# CreateViewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **interface{}** | Base64 encoded string, containing all the data for a given view. This data is expected to be versioned so that the frontend can deserialize as required. | 
**ViewName** | **string** | User-readable name of the view, doesn&#39;t need to be unique. | 

## Methods

### NewCreateViewRequest

`func NewCreateViewRequest(data interface{}, viewName string, ) *CreateViewRequest`

NewCreateViewRequest instantiates a new CreateViewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateViewRequestWithDefaults

`func NewCreateViewRequestWithDefaults() *CreateViewRequest`

NewCreateViewRequestWithDefaults instantiates a new CreateViewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateViewRequest) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateViewRequest) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateViewRequest) SetData(v interface{})`

SetData sets Data field to given value.


### SetDataNil

`func (o *CreateViewRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *CreateViewRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetViewName

`func (o *CreateViewRequest) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *CreateViewRequest) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *CreateViewRequest) SetViewName(v string)`

SetViewName sets ViewName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


