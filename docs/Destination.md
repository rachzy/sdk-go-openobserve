# Destination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsRegion** | Pointer to **NullableString** |  | [optional] 
**Emails** | Pointer to **[]string** | Required when &#x60;destination_type&#x60; is &#x60;Email&#x60; | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 
**Method** | Pointer to [**MetaDestHTTPType**](MetaDestHTTPType.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SkipTlsVerify** | Pointer to **bool** |  | [optional] 
**SnsTopicArn** | Pointer to **NullableString** |  | [optional] 
**Template** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**DestinationType**](DestinationType.md) |  | [optional] 
**Url** | Pointer to **string** | Required for &#x60;Http&#x60; destination_type | [optional] 

## Methods

### NewDestination

`func NewDestination() *Destination`

NewDestination instantiates a new Destination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationWithDefaults

`func NewDestinationWithDefaults() *Destination`

NewDestinationWithDefaults instantiates a new Destination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsRegion

`func (o *Destination) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *Destination) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *Destination) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *Destination) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### SetAwsRegionNil

`func (o *Destination) SetAwsRegionNil(b bool)`

 SetAwsRegionNil sets the value for AwsRegion to be an explicit nil

### UnsetAwsRegion
`func (o *Destination) UnsetAwsRegion()`

UnsetAwsRegion ensures that no value is present for AwsRegion, not even an explicit nil
### GetEmails

`func (o *Destination) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *Destination) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *Destination) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *Destination) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetHeaders

`func (o *Destination) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *Destination) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *Destination) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *Destination) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *Destination) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *Destination) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetMethod

`func (o *Destination) GetMethod() MetaDestHTTPType`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *Destination) GetMethodOk() (*MetaDestHTTPType, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *Destination) SetMethod(v MetaDestHTTPType)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *Destination) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetName

`func (o *Destination) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Destination) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Destination) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Destination) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSkipTlsVerify

`func (o *Destination) GetSkipTlsVerify() bool`

GetSkipTlsVerify returns the SkipTlsVerify field if non-nil, zero value otherwise.

### GetSkipTlsVerifyOk

`func (o *Destination) GetSkipTlsVerifyOk() (*bool, bool)`

GetSkipTlsVerifyOk returns a tuple with the SkipTlsVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipTlsVerify

`func (o *Destination) SetSkipTlsVerify(v bool)`

SetSkipTlsVerify sets SkipTlsVerify field to given value.

### HasSkipTlsVerify

`func (o *Destination) HasSkipTlsVerify() bool`

HasSkipTlsVerify returns a boolean if a field has been set.

### GetSnsTopicArn

`func (o *Destination) GetSnsTopicArn() string`

GetSnsTopicArn returns the SnsTopicArn field if non-nil, zero value otherwise.

### GetSnsTopicArnOk

`func (o *Destination) GetSnsTopicArnOk() (*string, bool)`

GetSnsTopicArnOk returns a tuple with the SnsTopicArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnsTopicArn

`func (o *Destination) SetSnsTopicArn(v string)`

SetSnsTopicArn sets SnsTopicArn field to given value.

### HasSnsTopicArn

`func (o *Destination) HasSnsTopicArn() bool`

HasSnsTopicArn returns a boolean if a field has been set.

### SetSnsTopicArnNil

`func (o *Destination) SetSnsTopicArnNil(b bool)`

 SetSnsTopicArnNil sets the value for SnsTopicArn to be an explicit nil

### UnsetSnsTopicArn
`func (o *Destination) UnsetSnsTopicArn()`

UnsetSnsTopicArn ensures that no value is present for SnsTopicArn, not even an explicit nil
### GetTemplate

`func (o *Destination) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *Destination) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *Destination) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *Destination) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### SetTemplateNil

`func (o *Destination) SetTemplateNil(b bool)`

 SetTemplateNil sets the value for Template to be an explicit nil

### UnsetTemplate
`func (o *Destination) UnsetTemplate()`

UnsetTemplate ensures that no value is present for Template, not even an explicit nil
### GetType

`func (o *Destination) GetType() DestinationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Destination) GetTypeOk() (*DestinationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Destination) SetType(v DestinationType)`

SetType sets Type field to given value.

### HasType

`func (o *Destination) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *Destination) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Destination) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Destination) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Destination) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


