# ListDashboardsResponseBodyItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** |  | 
**DashboardId** | **string** |  | 
**Description** | **string** |  | 
**FolderId** | **string** |  | 
**FolderName** | **string** |  | 
**Hash** | **string** |  | 
**Owner** | **string** |  | 
**Role** | **string** |  | 
**Title** | **string** |  | 
**UpdatedAt** | **int64** |  | 
**V1** | Pointer to [**NullableV1Dashboard**](V1Dashboard.md) |  | [optional] 
**V2** | Pointer to [**NullableV2Dashboard**](V2Dashboard.md) |  | [optional] 
**V3** | Pointer to [**NullableV3Dashboard**](V3Dashboard.md) |  | [optional] 
**V4** | Pointer to [**NullableV4Dashboard**](V4Dashboard.md) |  | [optional] 
**V5** | Pointer to [**NullableV5Dashboard**](V5Dashboard.md) |  | [optional] 
**Version** | **int32** |  | 

## Methods

### NewListDashboardsResponseBodyItem

`func NewListDashboardsResponseBodyItem(created time.Time, dashboardId string, description string, folderId string, folderName string, hash string, owner string, role string, title string, updatedAt int64, version int32, ) *ListDashboardsResponseBodyItem`

NewListDashboardsResponseBodyItem instantiates a new ListDashboardsResponseBodyItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDashboardsResponseBodyItemWithDefaults

`func NewListDashboardsResponseBodyItemWithDefaults() *ListDashboardsResponseBodyItem`

NewListDashboardsResponseBodyItemWithDefaults instantiates a new ListDashboardsResponseBodyItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ListDashboardsResponseBodyItem) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ListDashboardsResponseBodyItem) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ListDashboardsResponseBodyItem) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetDashboardId

`func (o *ListDashboardsResponseBodyItem) GetDashboardId() string`

GetDashboardId returns the DashboardId field if non-nil, zero value otherwise.

### GetDashboardIdOk

`func (o *ListDashboardsResponseBodyItem) GetDashboardIdOk() (*string, bool)`

GetDashboardIdOk returns a tuple with the DashboardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardId

`func (o *ListDashboardsResponseBodyItem) SetDashboardId(v string)`

SetDashboardId sets DashboardId field to given value.


### GetDescription

`func (o *ListDashboardsResponseBodyItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListDashboardsResponseBodyItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListDashboardsResponseBodyItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFolderId

`func (o *ListDashboardsResponseBodyItem) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *ListDashboardsResponseBodyItem) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *ListDashboardsResponseBodyItem) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.


### GetFolderName

`func (o *ListDashboardsResponseBodyItem) GetFolderName() string`

GetFolderName returns the FolderName field if non-nil, zero value otherwise.

### GetFolderNameOk

`func (o *ListDashboardsResponseBodyItem) GetFolderNameOk() (*string, bool)`

GetFolderNameOk returns a tuple with the FolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderName

`func (o *ListDashboardsResponseBodyItem) SetFolderName(v string)`

SetFolderName sets FolderName field to given value.


### GetHash

`func (o *ListDashboardsResponseBodyItem) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ListDashboardsResponseBodyItem) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ListDashboardsResponseBodyItem) SetHash(v string)`

SetHash sets Hash field to given value.


### GetOwner

`func (o *ListDashboardsResponseBodyItem) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListDashboardsResponseBodyItem) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListDashboardsResponseBodyItem) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetRole

`func (o *ListDashboardsResponseBodyItem) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ListDashboardsResponseBodyItem) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ListDashboardsResponseBodyItem) SetRole(v string)`

SetRole sets Role field to given value.


### GetTitle

`func (o *ListDashboardsResponseBodyItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ListDashboardsResponseBodyItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ListDashboardsResponseBodyItem) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUpdatedAt

`func (o *ListDashboardsResponseBodyItem) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListDashboardsResponseBodyItem) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListDashboardsResponseBodyItem) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetV1

`func (o *ListDashboardsResponseBodyItem) GetV1() V1Dashboard`

GetV1 returns the V1 field if non-nil, zero value otherwise.

### GetV1Ok

`func (o *ListDashboardsResponseBodyItem) GetV1Ok() (*V1Dashboard, bool)`

GetV1Ok returns a tuple with the V1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV1

`func (o *ListDashboardsResponseBodyItem) SetV1(v V1Dashboard)`

SetV1 sets V1 field to given value.

### HasV1

`func (o *ListDashboardsResponseBodyItem) HasV1() bool`

HasV1 returns a boolean if a field has been set.

### SetV1Nil

`func (o *ListDashboardsResponseBodyItem) SetV1Nil(b bool)`

 SetV1Nil sets the value for V1 to be an explicit nil

### UnsetV1
`func (o *ListDashboardsResponseBodyItem) UnsetV1()`

UnsetV1 ensures that no value is present for V1, not even an explicit nil
### GetV2

`func (o *ListDashboardsResponseBodyItem) GetV2() V2Dashboard`

GetV2 returns the V2 field if non-nil, zero value otherwise.

### GetV2Ok

`func (o *ListDashboardsResponseBodyItem) GetV2Ok() (*V2Dashboard, bool)`

GetV2Ok returns a tuple with the V2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV2

`func (o *ListDashboardsResponseBodyItem) SetV2(v V2Dashboard)`

SetV2 sets V2 field to given value.

### HasV2

`func (o *ListDashboardsResponseBodyItem) HasV2() bool`

HasV2 returns a boolean if a field has been set.

### SetV2Nil

`func (o *ListDashboardsResponseBodyItem) SetV2Nil(b bool)`

 SetV2Nil sets the value for V2 to be an explicit nil

### UnsetV2
`func (o *ListDashboardsResponseBodyItem) UnsetV2()`

UnsetV2 ensures that no value is present for V2, not even an explicit nil
### GetV3

`func (o *ListDashboardsResponseBodyItem) GetV3() V3Dashboard`

GetV3 returns the V3 field if non-nil, zero value otherwise.

### GetV3Ok

`func (o *ListDashboardsResponseBodyItem) GetV3Ok() (*V3Dashboard, bool)`

GetV3Ok returns a tuple with the V3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV3

`func (o *ListDashboardsResponseBodyItem) SetV3(v V3Dashboard)`

SetV3 sets V3 field to given value.

### HasV3

`func (o *ListDashboardsResponseBodyItem) HasV3() bool`

HasV3 returns a boolean if a field has been set.

### SetV3Nil

`func (o *ListDashboardsResponseBodyItem) SetV3Nil(b bool)`

 SetV3Nil sets the value for V3 to be an explicit nil

### UnsetV3
`func (o *ListDashboardsResponseBodyItem) UnsetV3()`

UnsetV3 ensures that no value is present for V3, not even an explicit nil
### GetV4

`func (o *ListDashboardsResponseBodyItem) GetV4() V4Dashboard`

GetV4 returns the V4 field if non-nil, zero value otherwise.

### GetV4Ok

`func (o *ListDashboardsResponseBodyItem) GetV4Ok() (*V4Dashboard, bool)`

GetV4Ok returns a tuple with the V4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV4

`func (o *ListDashboardsResponseBodyItem) SetV4(v V4Dashboard)`

SetV4 sets V4 field to given value.

### HasV4

`func (o *ListDashboardsResponseBodyItem) HasV4() bool`

HasV4 returns a boolean if a field has been set.

### SetV4Nil

`func (o *ListDashboardsResponseBodyItem) SetV4Nil(b bool)`

 SetV4Nil sets the value for V4 to be an explicit nil

### UnsetV4
`func (o *ListDashboardsResponseBodyItem) UnsetV4()`

UnsetV4 ensures that no value is present for V4, not even an explicit nil
### GetV5

`func (o *ListDashboardsResponseBodyItem) GetV5() V5Dashboard`

GetV5 returns the V5 field if non-nil, zero value otherwise.

### GetV5Ok

`func (o *ListDashboardsResponseBodyItem) GetV5Ok() (*V5Dashboard, bool)`

GetV5Ok returns a tuple with the V5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV5

`func (o *ListDashboardsResponseBodyItem) SetV5(v V5Dashboard)`

SetV5 sets V5 field to given value.

### HasV5

`func (o *ListDashboardsResponseBodyItem) HasV5() bool`

HasV5 returns a boolean if a field has been set.

### SetV5Nil

`func (o *ListDashboardsResponseBodyItem) SetV5Nil(b bool)`

 SetV5Nil sets the value for V5 to be an explicit nil

### UnsetV5
`func (o *ListDashboardsResponseBodyItem) UnsetV5()`

UnsetV5 ensures that no value is present for V5, not even an explicit nil
### GetVersion

`func (o *ListDashboardsResponseBodyItem) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListDashboardsResponseBodyItem) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListDashboardsResponseBodyItem) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


