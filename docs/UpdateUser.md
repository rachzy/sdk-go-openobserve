# UpdateUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangePassword** | Pointer to **bool** |  | [optional] 
**FirstName** | Pointer to **NullableString** |  | [optional] 
**LastName** | Pointer to **NullableString** |  | [optional] 
**NewPassword** | Pointer to **NullableString** |  | [optional] 
**OldPassword** | Pointer to **NullableString** |  | [optional] 
**Role** | Pointer to [**NullableUserRole**](UserRole.md) |  | [optional] 
**Token** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateUser

`func NewUpdateUser() *UpdateUser`

NewUpdateUser instantiates a new UpdateUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserWithDefaults

`func NewUpdateUserWithDefaults() *UpdateUser`

NewUpdateUserWithDefaults instantiates a new UpdateUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangePassword

`func (o *UpdateUser) GetChangePassword() bool`

GetChangePassword returns the ChangePassword field if non-nil, zero value otherwise.

### GetChangePasswordOk

`func (o *UpdateUser) GetChangePasswordOk() (*bool, bool)`

GetChangePasswordOk returns a tuple with the ChangePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangePassword

`func (o *UpdateUser) SetChangePassword(v bool)`

SetChangePassword sets ChangePassword field to given value.

### HasChangePassword

`func (o *UpdateUser) HasChangePassword() bool`

HasChangePassword returns a boolean if a field has been set.

### GetFirstName

`func (o *UpdateUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UpdateUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UpdateUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UpdateUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *UpdateUser) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *UpdateUser) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *UpdateUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UpdateUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UpdateUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UpdateUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *UpdateUser) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *UpdateUser) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetNewPassword

`func (o *UpdateUser) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *UpdateUser) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *UpdateUser) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *UpdateUser) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.

### SetNewPasswordNil

`func (o *UpdateUser) SetNewPasswordNil(b bool)`

 SetNewPasswordNil sets the value for NewPassword to be an explicit nil

### UnsetNewPassword
`func (o *UpdateUser) UnsetNewPassword()`

UnsetNewPassword ensures that no value is present for NewPassword, not even an explicit nil
### GetOldPassword

`func (o *UpdateUser) GetOldPassword() string`

GetOldPassword returns the OldPassword field if non-nil, zero value otherwise.

### GetOldPasswordOk

`func (o *UpdateUser) GetOldPasswordOk() (*string, bool)`

GetOldPasswordOk returns a tuple with the OldPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPassword

`func (o *UpdateUser) SetOldPassword(v string)`

SetOldPassword sets OldPassword field to given value.

### HasOldPassword

`func (o *UpdateUser) HasOldPassword() bool`

HasOldPassword returns a boolean if a field has been set.

### SetOldPasswordNil

`func (o *UpdateUser) SetOldPasswordNil(b bool)`

 SetOldPasswordNil sets the value for OldPassword to be an explicit nil

### UnsetOldPassword
`func (o *UpdateUser) UnsetOldPassword()`

UnsetOldPassword ensures that no value is present for OldPassword, not even an explicit nil
### GetRole

`func (o *UpdateUser) GetRole() UserRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UpdateUser) GetRoleOk() (*UserRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UpdateUser) SetRole(v UserRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *UpdateUser) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *UpdateUser) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *UpdateUser) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetToken

`func (o *UpdateUser) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateUser) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateUser) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateUser) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *UpdateUser) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *UpdateUser) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


