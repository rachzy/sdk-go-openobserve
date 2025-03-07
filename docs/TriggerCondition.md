# TriggerCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cron** | Pointer to **string** |  | [optional] 
**Frequency** | Pointer to **int64** |  | [optional] 
**FrequencyType** | Pointer to [**FrequencyType**](FrequencyType.md) |  | [optional] 
**Operator** | Pointer to [**Operator**](Operator.md) |  | [optional] 
**Period** | **int64** |  | 
**Silence** | Pointer to **int64** |  | [optional] 
**Threshold** | Pointer to **int64** |  | [optional] 
**Timezone** | Pointer to **NullableString** |  | [optional] 
**ToleranceInSecs** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewTriggerCondition

`func NewTriggerCondition(period int64, ) *TriggerCondition`

NewTriggerCondition instantiates a new TriggerCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerConditionWithDefaults

`func NewTriggerConditionWithDefaults() *TriggerCondition`

NewTriggerConditionWithDefaults instantiates a new TriggerCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCron

`func (o *TriggerCondition) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *TriggerCondition) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *TriggerCondition) SetCron(v string)`

SetCron sets Cron field to given value.

### HasCron

`func (o *TriggerCondition) HasCron() bool`

HasCron returns a boolean if a field has been set.

### GetFrequency

`func (o *TriggerCondition) GetFrequency() int64`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *TriggerCondition) GetFrequencyOk() (*int64, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *TriggerCondition) SetFrequency(v int64)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *TriggerCondition) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetFrequencyType

`func (o *TriggerCondition) GetFrequencyType() FrequencyType`

GetFrequencyType returns the FrequencyType field if non-nil, zero value otherwise.

### GetFrequencyTypeOk

`func (o *TriggerCondition) GetFrequencyTypeOk() (*FrequencyType, bool)`

GetFrequencyTypeOk returns a tuple with the FrequencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyType

`func (o *TriggerCondition) SetFrequencyType(v FrequencyType)`

SetFrequencyType sets FrequencyType field to given value.

### HasFrequencyType

`func (o *TriggerCondition) HasFrequencyType() bool`

HasFrequencyType returns a boolean if a field has been set.

### GetOperator

`func (o *TriggerCondition) GetOperator() Operator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *TriggerCondition) GetOperatorOk() (*Operator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *TriggerCondition) SetOperator(v Operator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *TriggerCondition) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetPeriod

`func (o *TriggerCondition) GetPeriod() int64`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *TriggerCondition) GetPeriodOk() (*int64, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *TriggerCondition) SetPeriod(v int64)`

SetPeriod sets Period field to given value.


### GetSilence

`func (o *TriggerCondition) GetSilence() int64`

GetSilence returns the Silence field if non-nil, zero value otherwise.

### GetSilenceOk

`func (o *TriggerCondition) GetSilenceOk() (*int64, bool)`

GetSilenceOk returns a tuple with the Silence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSilence

`func (o *TriggerCondition) SetSilence(v int64)`

SetSilence sets Silence field to given value.

### HasSilence

`func (o *TriggerCondition) HasSilence() bool`

HasSilence returns a boolean if a field has been set.

### GetThreshold

`func (o *TriggerCondition) GetThreshold() int64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *TriggerCondition) GetThresholdOk() (*int64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *TriggerCondition) SetThreshold(v int64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *TriggerCondition) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetTimezone

`func (o *TriggerCondition) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *TriggerCondition) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *TriggerCondition) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *TriggerCondition) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *TriggerCondition) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *TriggerCondition) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetToleranceInSecs

`func (o *TriggerCondition) GetToleranceInSecs() int64`

GetToleranceInSecs returns the ToleranceInSecs field if non-nil, zero value otherwise.

### GetToleranceInSecsOk

`func (o *TriggerCondition) GetToleranceInSecsOk() (*int64, bool)`

GetToleranceInSecsOk returns a tuple with the ToleranceInSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToleranceInSecs

`func (o *TriggerCondition) SetToleranceInSecs(v int64)`

SetToleranceInSecs sets ToleranceInSecs field to given value.

### HasToleranceInSecs

`func (o *TriggerCondition) HasToleranceInSecs() bool`

HasToleranceInSecs returns a boolean if a field has been set.

### SetToleranceInSecsNil

`func (o *TriggerCondition) SetToleranceInSecsNil(b bool)`

 SetToleranceInSecsNil sets the value for ToleranceInSecs to be an explicit nil

### UnsetToleranceInSecs
`func (o *TriggerCondition) UnsetToleranceInSecs()`

UnsetToleranceInSecs ensures that no value is present for ToleranceInSecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


