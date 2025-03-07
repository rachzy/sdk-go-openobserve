# SearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | Pointer to **[]string** |  | [optional] 
**Encoding** | Pointer to [**RequestEncoding**](RequestEncoding.md) |  | [optional] 
**Query** | [**SearchQuery**](SearchQuery.md) |  | 
**Regions** | Pointer to **[]string** |  | [optional] 
**SearchEventContext** | Pointer to [**NullableSearchEventContext**](SearchEventContext.md) |  | [optional] 
**SearchType** | Pointer to [**NullableSearchEventType**](SearchEventType.md) |  | [optional] 
**Timeout** | Pointer to **int64** |  | [optional] 
**UseCache** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewSearchRequest

`func NewSearchRequest(query SearchQuery, ) *SearchRequest`

NewSearchRequest instantiates a new SearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchRequestWithDefaults

`func NewSearchRequestWithDefaults() *SearchRequest`

NewSearchRequestWithDefaults instantiates a new SearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusters

`func (o *SearchRequest) GetClusters() []string`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *SearchRequest) GetClustersOk() (*[]string, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *SearchRequest) SetClusters(v []string)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *SearchRequest) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetEncoding

`func (o *SearchRequest) GetEncoding() RequestEncoding`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *SearchRequest) GetEncodingOk() (*RequestEncoding, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *SearchRequest) SetEncoding(v RequestEncoding)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *SearchRequest) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetQuery

`func (o *SearchRequest) GetQuery() SearchQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SearchRequest) GetQueryOk() (*SearchQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SearchRequest) SetQuery(v SearchQuery)`

SetQuery sets Query field to given value.


### GetRegions

`func (o *SearchRequest) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *SearchRequest) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *SearchRequest) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *SearchRequest) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetSearchEventContext

`func (o *SearchRequest) GetSearchEventContext() SearchEventContext`

GetSearchEventContext returns the SearchEventContext field if non-nil, zero value otherwise.

### GetSearchEventContextOk

`func (o *SearchRequest) GetSearchEventContextOk() (*SearchEventContext, bool)`

GetSearchEventContextOk returns a tuple with the SearchEventContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchEventContext

`func (o *SearchRequest) SetSearchEventContext(v SearchEventContext)`

SetSearchEventContext sets SearchEventContext field to given value.

### HasSearchEventContext

`func (o *SearchRequest) HasSearchEventContext() bool`

HasSearchEventContext returns a boolean if a field has been set.

### SetSearchEventContextNil

`func (o *SearchRequest) SetSearchEventContextNil(b bool)`

 SetSearchEventContextNil sets the value for SearchEventContext to be an explicit nil

### UnsetSearchEventContext
`func (o *SearchRequest) UnsetSearchEventContext()`

UnsetSearchEventContext ensures that no value is present for SearchEventContext, not even an explicit nil
### GetSearchType

`func (o *SearchRequest) GetSearchType() SearchEventType`

GetSearchType returns the SearchType field if non-nil, zero value otherwise.

### GetSearchTypeOk

`func (o *SearchRequest) GetSearchTypeOk() (*SearchEventType, bool)`

GetSearchTypeOk returns a tuple with the SearchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchType

`func (o *SearchRequest) SetSearchType(v SearchEventType)`

SetSearchType sets SearchType field to given value.

### HasSearchType

`func (o *SearchRequest) HasSearchType() bool`

HasSearchType returns a boolean if a field has been set.

### SetSearchTypeNil

`func (o *SearchRequest) SetSearchTypeNil(b bool)`

 SetSearchTypeNil sets the value for SearchType to be an explicit nil

### UnsetSearchType
`func (o *SearchRequest) UnsetSearchType()`

UnsetSearchType ensures that no value is present for SearchType, not even an explicit nil
### GetTimeout

`func (o *SearchRequest) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SearchRequest) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *SearchRequest) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *SearchRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUseCache

`func (o *SearchRequest) GetUseCache() bool`

GetUseCache returns the UseCache field if non-nil, zero value otherwise.

### GetUseCacheOk

`func (o *SearchRequest) GetUseCacheOk() (*bool, bool)`

GetUseCacheOk returns a tuple with the UseCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCache

`func (o *SearchRequest) SetUseCache(v bool)`

SetUseCache sets UseCache field to given value.

### HasUseCache

`func (o *SearchRequest) HasUseCache() bool`

HasUseCache returns a boolean if a field has been set.

### SetUseCacheNil

`func (o *SearchRequest) SetUseCacheNil(b bool)`

 SetUseCacheNil sets the value for UseCache to be an explicit nil

### UnsetUseCache
`func (o *SearchRequest) UnsetUseCache()`

UnsetUseCache ensures that no value is present for UseCache, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


