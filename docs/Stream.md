# Stream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricsMeta** | Pointer to [**NullableMetadata**](Metadata.md) |  | [optional] 
**Name** | **string** |  | 
**Schema** | Pointer to [**[]StreamProperty**](StreamProperty.md) |  | [optional] 
**Settings** | [**StreamSettings**](StreamSettings.md) |  | 
**Stats** | [**StreamStats**](StreamStats.md) |  | 
**StorageType** | **string** |  | 
**StreamType** | [**StreamType**](StreamType.md) |  | 
**TotalFields** | **int32** |  | 
**UdsSchema** | Pointer to [**[]StreamProperty**](StreamProperty.md) |  | [optional] 

## Methods

### NewStream

`func NewStream(name string, settings StreamSettings, stats StreamStats, storageType string, streamType StreamType, totalFields int32, ) *Stream`

NewStream instantiates a new Stream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamWithDefaults

`func NewStreamWithDefaults() *Stream`

NewStreamWithDefaults instantiates a new Stream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricsMeta

`func (o *Stream) GetMetricsMeta() Metadata`

GetMetricsMeta returns the MetricsMeta field if non-nil, zero value otherwise.

### GetMetricsMetaOk

`func (o *Stream) GetMetricsMetaOk() (*Metadata, bool)`

GetMetricsMetaOk returns a tuple with the MetricsMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsMeta

`func (o *Stream) SetMetricsMeta(v Metadata)`

SetMetricsMeta sets MetricsMeta field to given value.

### HasMetricsMeta

`func (o *Stream) HasMetricsMeta() bool`

HasMetricsMeta returns a boolean if a field has been set.

### SetMetricsMetaNil

`func (o *Stream) SetMetricsMetaNil(b bool)`

 SetMetricsMetaNil sets the value for MetricsMeta to be an explicit nil

### UnsetMetricsMeta
`func (o *Stream) UnsetMetricsMeta()`

UnsetMetricsMeta ensures that no value is present for MetricsMeta, not even an explicit nil
### GetName

`func (o *Stream) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Stream) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Stream) SetName(v string)`

SetName sets Name field to given value.


### GetSchema

`func (o *Stream) GetSchema() []StreamProperty`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *Stream) GetSchemaOk() (*[]StreamProperty, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *Stream) SetSchema(v []StreamProperty)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *Stream) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetSettings

`func (o *Stream) GetSettings() StreamSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *Stream) GetSettingsOk() (*StreamSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *Stream) SetSettings(v StreamSettings)`

SetSettings sets Settings field to given value.


### GetStats

`func (o *Stream) GetStats() StreamStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Stream) GetStatsOk() (*StreamStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Stream) SetStats(v StreamStats)`

SetStats sets Stats field to given value.


### GetStorageType

`func (o *Stream) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *Stream) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *Stream) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.


### GetStreamType

`func (o *Stream) GetStreamType() StreamType`

GetStreamType returns the StreamType field if non-nil, zero value otherwise.

### GetStreamTypeOk

`func (o *Stream) GetStreamTypeOk() (*StreamType, bool)`

GetStreamTypeOk returns a tuple with the StreamType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamType

`func (o *Stream) SetStreamType(v StreamType)`

SetStreamType sets StreamType field to given value.


### GetTotalFields

`func (o *Stream) GetTotalFields() int32`

GetTotalFields returns the TotalFields field if non-nil, zero value otherwise.

### GetTotalFieldsOk

`func (o *Stream) GetTotalFieldsOk() (*int32, bool)`

GetTotalFieldsOk returns a tuple with the TotalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFields

`func (o *Stream) SetTotalFields(v int32)`

SetTotalFields sets TotalFields field to given value.


### GetUdsSchema

`func (o *Stream) GetUdsSchema() []StreamProperty`

GetUdsSchema returns the UdsSchema field if non-nil, zero value otherwise.

### GetUdsSchemaOk

`func (o *Stream) GetUdsSchemaOk() (*[]StreamProperty, bool)`

GetUdsSchemaOk returns a tuple with the UdsSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdsSchema

`func (o *Stream) SetUdsSchema(v []StreamProperty)`

SetUdsSchema sets UdsSchema field to given value.

### HasUdsSchema

`func (o *Stream) HasUdsSchema() bool`

HasUdsSchema returns a boolean if a field has been set.

### SetUdsSchemaNil

`func (o *Stream) SetUdsSchemaNil(b bool)`

 SetUdsSchemaNil sets the value for UdsSchema to be an explicit nil

### UnsetUdsSchema
`func (o *Stream) UnsetUdsSchema()`

UnsetUdsSchema ensures that no value is present for UdsSchema, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


