# ResponseTook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterTotal** | **int32** |  | 
**ClusterWaitQueue** | **int32** |  | 
**IdxTook** | **int32** |  | 
**Nodes** | Pointer to [**[]ResponseNodeTook**](ResponseNodeTook.md) |  | [optional] 
**Total** | **int32** |  | 
**WaitQueue** | **int32** |  | 

## Methods

### NewResponseTook

`func NewResponseTook(clusterTotal int32, clusterWaitQueue int32, idxTook int32, total int32, waitQueue int32, ) *ResponseTook`

NewResponseTook instantiates a new ResponseTook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseTookWithDefaults

`func NewResponseTookWithDefaults() *ResponseTook`

NewResponseTookWithDefaults instantiates a new ResponseTook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterTotal

`func (o *ResponseTook) GetClusterTotal() int32`

GetClusterTotal returns the ClusterTotal field if non-nil, zero value otherwise.

### GetClusterTotalOk

`func (o *ResponseTook) GetClusterTotalOk() (*int32, bool)`

GetClusterTotalOk returns a tuple with the ClusterTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterTotal

`func (o *ResponseTook) SetClusterTotal(v int32)`

SetClusterTotal sets ClusterTotal field to given value.


### GetClusterWaitQueue

`func (o *ResponseTook) GetClusterWaitQueue() int32`

GetClusterWaitQueue returns the ClusterWaitQueue field if non-nil, zero value otherwise.

### GetClusterWaitQueueOk

`func (o *ResponseTook) GetClusterWaitQueueOk() (*int32, bool)`

GetClusterWaitQueueOk returns a tuple with the ClusterWaitQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterWaitQueue

`func (o *ResponseTook) SetClusterWaitQueue(v int32)`

SetClusterWaitQueue sets ClusterWaitQueue field to given value.


### GetIdxTook

`func (o *ResponseTook) GetIdxTook() int32`

GetIdxTook returns the IdxTook field if non-nil, zero value otherwise.

### GetIdxTookOk

`func (o *ResponseTook) GetIdxTookOk() (*int32, bool)`

GetIdxTookOk returns a tuple with the IdxTook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdxTook

`func (o *ResponseTook) SetIdxTook(v int32)`

SetIdxTook sets IdxTook field to given value.


### GetNodes

`func (o *ResponseTook) GetNodes() []ResponseNodeTook`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ResponseTook) GetNodesOk() (*[]ResponseNodeTook, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ResponseTook) SetNodes(v []ResponseNodeTook)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *ResponseTook) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetTotal

`func (o *ResponseTook) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ResponseTook) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ResponseTook) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetWaitQueue

`func (o *ResponseTook) GetWaitQueue() int32`

GetWaitQueue returns the WaitQueue field if non-nil, zero value otherwise.

### GetWaitQueueOk

`func (o *ResponseTook) GetWaitQueueOk() (*int32, bool)`

GetWaitQueueOk returns a tuple with the WaitQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitQueue

`func (o *ResponseTook) SetWaitQueue(v int32)`

SetWaitQueue sets WaitQueue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


