# OrgSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alerts** | [**AlertSummary**](AlertSummary.md) |  | 
**Pipelines** | [**PipelineSummary**](PipelineSummary.md) |  | 
**Streams** | [**StreamSummary**](StreamSummary.md) |  | 
**TotalDashboards** | **int64** |  | 
**TotalFunctions** | **int64** |  | 

## Methods

### NewOrgSummary

`func NewOrgSummary(alerts AlertSummary, pipelines PipelineSummary, streams StreamSummary, totalDashboards int64, totalFunctions int64, ) *OrgSummary`

NewOrgSummary instantiates a new OrgSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgSummaryWithDefaults

`func NewOrgSummaryWithDefaults() *OrgSummary`

NewOrgSummaryWithDefaults instantiates a new OrgSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlerts

`func (o *OrgSummary) GetAlerts() AlertSummary`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *OrgSummary) GetAlertsOk() (*AlertSummary, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *OrgSummary) SetAlerts(v AlertSummary)`

SetAlerts sets Alerts field to given value.


### GetPipelines

`func (o *OrgSummary) GetPipelines() PipelineSummary`

GetPipelines returns the Pipelines field if non-nil, zero value otherwise.

### GetPipelinesOk

`func (o *OrgSummary) GetPipelinesOk() (*PipelineSummary, bool)`

GetPipelinesOk returns a tuple with the Pipelines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelines

`func (o *OrgSummary) SetPipelines(v PipelineSummary)`

SetPipelines sets Pipelines field to given value.


### GetStreams

`func (o *OrgSummary) GetStreams() StreamSummary`

GetStreams returns the Streams field if non-nil, zero value otherwise.

### GetStreamsOk

`func (o *OrgSummary) GetStreamsOk() (*StreamSummary, bool)`

GetStreamsOk returns a tuple with the Streams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreams

`func (o *OrgSummary) SetStreams(v StreamSummary)`

SetStreams sets Streams field to given value.


### GetTotalDashboards

`func (o *OrgSummary) GetTotalDashboards() int64`

GetTotalDashboards returns the TotalDashboards field if non-nil, zero value otherwise.

### GetTotalDashboardsOk

`func (o *OrgSummary) GetTotalDashboardsOk() (*int64, bool)`

GetTotalDashboardsOk returns a tuple with the TotalDashboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDashboards

`func (o *OrgSummary) SetTotalDashboards(v int64)`

SetTotalDashboards sets TotalDashboards field to given value.


### GetTotalFunctions

`func (o *OrgSummary) GetTotalFunctions() int64`

GetTotalFunctions returns the TotalFunctions field if non-nil, zero value otherwise.

### GetTotalFunctionsOk

`func (o *OrgSummary) GetTotalFunctionsOk() (*int64, bool)`

GetTotalFunctionsOk returns a tuple with the TotalFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFunctions

`func (o *OrgSummary) SetTotalFunctions(v int64)`

SetTotalFunctions sets TotalFunctions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


