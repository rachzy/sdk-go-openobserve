# SyslogRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**StreamName** | Pointer to **string** |  | [optional] 
**Subnets** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSyslogRoute

`func NewSyslogRoute() *SyslogRoute`

NewSyslogRoute instantiates a new SyslogRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyslogRouteWithDefaults

`func NewSyslogRouteWithDefaults() *SyslogRoute`

NewSyslogRouteWithDefaults instantiates a new SyslogRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SyslogRoute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyslogRoute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyslogRoute) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SyslogRoute) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrgId

`func (o *SyslogRoute) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *SyslogRoute) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *SyslogRoute) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *SyslogRoute) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetStreamName

`func (o *SyslogRoute) GetStreamName() string`

GetStreamName returns the StreamName field if non-nil, zero value otherwise.

### GetStreamNameOk

`func (o *SyslogRoute) GetStreamNameOk() (*string, bool)`

GetStreamNameOk returns a tuple with the StreamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamName

`func (o *SyslogRoute) SetStreamName(v string)`

SetStreamName sets StreamName field to given value.

### HasStreamName

`func (o *SyslogRoute) HasStreamName() bool`

HasStreamName returns a boolean if a field has been set.

### GetSubnets

`func (o *SyslogRoute) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *SyslogRoute) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *SyslogRoute) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *SyslogRoute) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


