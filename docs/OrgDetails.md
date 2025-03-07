# OrgDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserObj** | [**OrgUser**](OrgUser.md) |  | 
**Id** | **int64** |  | 
**Identifier** | **string** |  | 
**IngestThreshold** | **int64** |  | 
**Name** | **string** |  | 
**SearchThreshold** | **int64** |  | 
**Type** | **string** |  | 
**UserEmail** | **string** |  | 

## Methods

### NewOrgDetails

`func NewOrgDetails(userObj OrgUser, id int64, identifier string, ingestThreshold int64, name string, searchThreshold int64, type_ string, userEmail string, ) *OrgDetails`

NewOrgDetails instantiates a new OrgDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgDetailsWithDefaults

`func NewOrgDetailsWithDefaults() *OrgDetails`

NewOrgDetailsWithDefaults instantiates a new OrgDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserObj

`func (o *OrgDetails) GetUserObj() OrgUser`

GetUserObj returns the UserObj field if non-nil, zero value otherwise.

### GetUserObjOk

`func (o *OrgDetails) GetUserObjOk() (*OrgUser, bool)`

GetUserObjOk returns a tuple with the UserObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserObj

`func (o *OrgDetails) SetUserObj(v OrgUser)`

SetUserObj sets UserObj field to given value.


### GetId

`func (o *OrgDetails) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgDetails) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgDetails) SetId(v int64)`

SetId sets Id field to given value.


### GetIdentifier

`func (o *OrgDetails) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *OrgDetails) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *OrgDetails) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetIngestThreshold

`func (o *OrgDetails) GetIngestThreshold() int64`

GetIngestThreshold returns the IngestThreshold field if non-nil, zero value otherwise.

### GetIngestThresholdOk

`func (o *OrgDetails) GetIngestThresholdOk() (*int64, bool)`

GetIngestThresholdOk returns a tuple with the IngestThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestThreshold

`func (o *OrgDetails) SetIngestThreshold(v int64)`

SetIngestThreshold sets IngestThreshold field to given value.


### GetName

`func (o *OrgDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrgDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrgDetails) SetName(v string)`

SetName sets Name field to given value.


### GetSearchThreshold

`func (o *OrgDetails) GetSearchThreshold() int64`

GetSearchThreshold returns the SearchThreshold field if non-nil, zero value otherwise.

### GetSearchThresholdOk

`func (o *OrgDetails) GetSearchThresholdOk() (*int64, bool)`

GetSearchThresholdOk returns a tuple with the SearchThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchThreshold

`func (o *OrgDetails) SetSearchThreshold(v int64)`

SetSearchThreshold sets SearchThreshold field to given value.


### GetType

`func (o *OrgDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrgDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrgDetails) SetType(v string)`

SetType sets Type field to given value.


### GetUserEmail

`func (o *OrgDetails) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *OrgDetails) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *OrgDetails) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


