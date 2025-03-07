# OrganizationSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableWebsocketSearch** | Pointer to **bool** |  | [optional] 
**MinAutoRefreshInterval** | Pointer to **int32** |  | [optional] 
**ScrapeInterval** | Pointer to **int32** | Ideally this should be the same as prometheus-scrape-interval (in seconds). | [optional] 
**SpanIdFieldName** | Pointer to **string** |  | [optional] 
**ToggleIngestionLogs** | Pointer to **bool** |  | [optional] 
**TraceIdFieldName** | Pointer to **string** |  | [optional] 

## Methods

### NewOrganizationSetting

`func NewOrganizationSetting() *OrganizationSetting`

NewOrganizationSetting instantiates a new OrganizationSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationSettingWithDefaults

`func NewOrganizationSettingWithDefaults() *OrganizationSetting`

NewOrganizationSettingWithDefaults instantiates a new OrganizationSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableWebsocketSearch

`func (o *OrganizationSetting) GetEnableWebsocketSearch() bool`

GetEnableWebsocketSearch returns the EnableWebsocketSearch field if non-nil, zero value otherwise.

### GetEnableWebsocketSearchOk

`func (o *OrganizationSetting) GetEnableWebsocketSearchOk() (*bool, bool)`

GetEnableWebsocketSearchOk returns a tuple with the EnableWebsocketSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableWebsocketSearch

`func (o *OrganizationSetting) SetEnableWebsocketSearch(v bool)`

SetEnableWebsocketSearch sets EnableWebsocketSearch field to given value.

### HasEnableWebsocketSearch

`func (o *OrganizationSetting) HasEnableWebsocketSearch() bool`

HasEnableWebsocketSearch returns a boolean if a field has been set.

### GetMinAutoRefreshInterval

`func (o *OrganizationSetting) GetMinAutoRefreshInterval() int32`

GetMinAutoRefreshInterval returns the MinAutoRefreshInterval field if non-nil, zero value otherwise.

### GetMinAutoRefreshIntervalOk

`func (o *OrganizationSetting) GetMinAutoRefreshIntervalOk() (*int32, bool)`

GetMinAutoRefreshIntervalOk returns a tuple with the MinAutoRefreshInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAutoRefreshInterval

`func (o *OrganizationSetting) SetMinAutoRefreshInterval(v int32)`

SetMinAutoRefreshInterval sets MinAutoRefreshInterval field to given value.

### HasMinAutoRefreshInterval

`func (o *OrganizationSetting) HasMinAutoRefreshInterval() bool`

HasMinAutoRefreshInterval returns a boolean if a field has been set.

### GetScrapeInterval

`func (o *OrganizationSetting) GetScrapeInterval() int32`

GetScrapeInterval returns the ScrapeInterval field if non-nil, zero value otherwise.

### GetScrapeIntervalOk

`func (o *OrganizationSetting) GetScrapeIntervalOk() (*int32, bool)`

GetScrapeIntervalOk returns a tuple with the ScrapeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrapeInterval

`func (o *OrganizationSetting) SetScrapeInterval(v int32)`

SetScrapeInterval sets ScrapeInterval field to given value.

### HasScrapeInterval

`func (o *OrganizationSetting) HasScrapeInterval() bool`

HasScrapeInterval returns a boolean if a field has been set.

### GetSpanIdFieldName

`func (o *OrganizationSetting) GetSpanIdFieldName() string`

GetSpanIdFieldName returns the SpanIdFieldName field if non-nil, zero value otherwise.

### GetSpanIdFieldNameOk

`func (o *OrganizationSetting) GetSpanIdFieldNameOk() (*string, bool)`

GetSpanIdFieldNameOk returns a tuple with the SpanIdFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanIdFieldName

`func (o *OrganizationSetting) SetSpanIdFieldName(v string)`

SetSpanIdFieldName sets SpanIdFieldName field to given value.

### HasSpanIdFieldName

`func (o *OrganizationSetting) HasSpanIdFieldName() bool`

HasSpanIdFieldName returns a boolean if a field has been set.

### GetToggleIngestionLogs

`func (o *OrganizationSetting) GetToggleIngestionLogs() bool`

GetToggleIngestionLogs returns the ToggleIngestionLogs field if non-nil, zero value otherwise.

### GetToggleIngestionLogsOk

`func (o *OrganizationSetting) GetToggleIngestionLogsOk() (*bool, bool)`

GetToggleIngestionLogsOk returns a tuple with the ToggleIngestionLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToggleIngestionLogs

`func (o *OrganizationSetting) SetToggleIngestionLogs(v bool)`

SetToggleIngestionLogs sets ToggleIngestionLogs field to given value.

### HasToggleIngestionLogs

`func (o *OrganizationSetting) HasToggleIngestionLogs() bool`

HasToggleIngestionLogs returns a boolean if a field has been set.

### GetTraceIdFieldName

`func (o *OrganizationSetting) GetTraceIdFieldName() string`

GetTraceIdFieldName returns the TraceIdFieldName field if non-nil, zero value otherwise.

### GetTraceIdFieldNameOk

`func (o *OrganizationSetting) GetTraceIdFieldNameOk() (*string, bool)`

GetTraceIdFieldNameOk returns a tuple with the TraceIdFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceIdFieldName

`func (o *OrganizationSetting) SetTraceIdFieldName(v string)`

SetTraceIdFieldName sets TraceIdFieldName field to given value.

### HasTraceIdFieldName

`func (o *OrganizationSetting) HasTraceIdFieldName() bool`

HasTraceIdFieldName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


