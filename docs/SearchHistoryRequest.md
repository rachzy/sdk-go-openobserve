# SearchHistoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | **int64** |  | 
**OrgId** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**StartTime** | **int64** |  | 
**StreamName** | Pointer to **NullableString** |  | [optional] 
**StreamType** | Pointer to **NullableString** |  | [optional] 
**TraceId** | Pointer to **NullableString** |  | [optional] 
**UserEmail** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSearchHistoryRequest

`func NewSearchHistoryRequest(endTime int64, startTime int64, ) *SearchHistoryRequest`

NewSearchHistoryRequest instantiates a new SearchHistoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchHistoryRequestWithDefaults

`func NewSearchHistoryRequestWithDefaults() *SearchHistoryRequest`

NewSearchHistoryRequestWithDefaults instantiates a new SearchHistoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *SearchHistoryRequest) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SearchHistoryRequest) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SearchHistoryRequest) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.


### GetOrgId

`func (o *SearchHistoryRequest) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *SearchHistoryRequest) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *SearchHistoryRequest) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *SearchHistoryRequest) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *SearchHistoryRequest) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *SearchHistoryRequest) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
### GetSize

`func (o *SearchHistoryRequest) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SearchHistoryRequest) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SearchHistoryRequest) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *SearchHistoryRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStartTime

`func (o *SearchHistoryRequest) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SearchHistoryRequest) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SearchHistoryRequest) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.


### GetStreamName

`func (o *SearchHistoryRequest) GetStreamName() string`

GetStreamName returns the StreamName field if non-nil, zero value otherwise.

### GetStreamNameOk

`func (o *SearchHistoryRequest) GetStreamNameOk() (*string, bool)`

GetStreamNameOk returns a tuple with the StreamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamName

`func (o *SearchHistoryRequest) SetStreamName(v string)`

SetStreamName sets StreamName field to given value.

### HasStreamName

`func (o *SearchHistoryRequest) HasStreamName() bool`

HasStreamName returns a boolean if a field has been set.

### SetStreamNameNil

`func (o *SearchHistoryRequest) SetStreamNameNil(b bool)`

 SetStreamNameNil sets the value for StreamName to be an explicit nil

### UnsetStreamName
`func (o *SearchHistoryRequest) UnsetStreamName()`

UnsetStreamName ensures that no value is present for StreamName, not even an explicit nil
### GetStreamType

`func (o *SearchHistoryRequest) GetStreamType() string`

GetStreamType returns the StreamType field if non-nil, zero value otherwise.

### GetStreamTypeOk

`func (o *SearchHistoryRequest) GetStreamTypeOk() (*string, bool)`

GetStreamTypeOk returns a tuple with the StreamType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamType

`func (o *SearchHistoryRequest) SetStreamType(v string)`

SetStreamType sets StreamType field to given value.

### HasStreamType

`func (o *SearchHistoryRequest) HasStreamType() bool`

HasStreamType returns a boolean if a field has been set.

### SetStreamTypeNil

`func (o *SearchHistoryRequest) SetStreamTypeNil(b bool)`

 SetStreamTypeNil sets the value for StreamType to be an explicit nil

### UnsetStreamType
`func (o *SearchHistoryRequest) UnsetStreamType()`

UnsetStreamType ensures that no value is present for StreamType, not even an explicit nil
### GetTraceId

`func (o *SearchHistoryRequest) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *SearchHistoryRequest) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *SearchHistoryRequest) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *SearchHistoryRequest) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### SetTraceIdNil

`func (o *SearchHistoryRequest) SetTraceIdNil(b bool)`

 SetTraceIdNil sets the value for TraceId to be an explicit nil

### UnsetTraceId
`func (o *SearchHistoryRequest) UnsetTraceId()`

UnsetTraceId ensures that no value is present for TraceId, not even an explicit nil
### GetUserEmail

`func (o *SearchHistoryRequest) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *SearchHistoryRequest) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *SearchHistoryRequest) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *SearchHistoryRequest) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### SetUserEmailNil

`func (o *SearchHistoryRequest) SetUserEmailNil(b bool)`

 SetUserEmailNil sets the value for UserEmail to be an explicit nil

### UnsetUserEmail
`func (o *SearchHistoryRequest) UnsetUserEmail()`

UnsetUserEmail ensures that no value is present for UserEmail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


