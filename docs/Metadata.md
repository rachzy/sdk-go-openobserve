# Metadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Help** | **string** |  | 
**MetricFamilyName** | **string** |  | 
**MetricType** | [**MetricType**](MetricType.md) |  | 
**Unit** | **string** |  | 

## Methods

### NewMetadata

`func NewMetadata(help string, metricFamilyName string, metricType MetricType, unit string, ) *Metadata`

NewMetadata instantiates a new Metadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithDefaults

`func NewMetadataWithDefaults() *Metadata`

NewMetadataWithDefaults instantiates a new Metadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHelp

`func (o *Metadata) GetHelp() string`

GetHelp returns the Help field if non-nil, zero value otherwise.

### GetHelpOk

`func (o *Metadata) GetHelpOk() (*string, bool)`

GetHelpOk returns a tuple with the Help field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelp

`func (o *Metadata) SetHelp(v string)`

SetHelp sets Help field to given value.


### GetMetricFamilyName

`func (o *Metadata) GetMetricFamilyName() string`

GetMetricFamilyName returns the MetricFamilyName field if non-nil, zero value otherwise.

### GetMetricFamilyNameOk

`func (o *Metadata) GetMetricFamilyNameOk() (*string, bool)`

GetMetricFamilyNameOk returns a tuple with the MetricFamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricFamilyName

`func (o *Metadata) SetMetricFamilyName(v string)`

SetMetricFamilyName sets MetricFamilyName field to given value.


### GetMetricType

`func (o *Metadata) GetMetricType() MetricType`

GetMetricType returns the MetricType field if non-nil, zero value otherwise.

### GetMetricTypeOk

`func (o *Metadata) GetMetricTypeOk() (*MetricType, bool)`

GetMetricTypeOk returns a tuple with the MetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricType

`func (o *Metadata) SetMetricType(v MetricType)`

SetMetricType sets MetricType field to given value.


### GetUnit

`func (o *Metadata) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Metadata) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Metadata) SetUnit(v string)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


