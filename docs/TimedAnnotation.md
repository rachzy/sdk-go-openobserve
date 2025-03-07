# TimedAnnotation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnnotationId** | Pointer to **NullableString** |  | [optional] 
**EndTime** | Pointer to **NullableInt64** |  | [optional] 
**Panels** | **[]string** |  | 
**StartTime** | **int64** |  | 
**Tags** | **[]string** |  | 
**Text** | Pointer to **NullableString** |  | [optional] 
**Title** | **string** |  | 

## Methods

### NewTimedAnnotation

`func NewTimedAnnotation(panels []string, startTime int64, tags []string, title string, ) *TimedAnnotation`

NewTimedAnnotation instantiates a new TimedAnnotation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimedAnnotationWithDefaults

`func NewTimedAnnotationWithDefaults() *TimedAnnotation`

NewTimedAnnotationWithDefaults instantiates a new TimedAnnotation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotationId

`func (o *TimedAnnotation) GetAnnotationId() string`

GetAnnotationId returns the AnnotationId field if non-nil, zero value otherwise.

### GetAnnotationIdOk

`func (o *TimedAnnotation) GetAnnotationIdOk() (*string, bool)`

GetAnnotationIdOk returns a tuple with the AnnotationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationId

`func (o *TimedAnnotation) SetAnnotationId(v string)`

SetAnnotationId sets AnnotationId field to given value.

### HasAnnotationId

`func (o *TimedAnnotation) HasAnnotationId() bool`

HasAnnotationId returns a boolean if a field has been set.

### SetAnnotationIdNil

`func (o *TimedAnnotation) SetAnnotationIdNil(b bool)`

 SetAnnotationIdNil sets the value for AnnotationId to be an explicit nil

### UnsetAnnotationId
`func (o *TimedAnnotation) UnsetAnnotationId()`

UnsetAnnotationId ensures that no value is present for AnnotationId, not even an explicit nil
### GetEndTime

`func (o *TimedAnnotation) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TimedAnnotation) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TimedAnnotation) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TimedAnnotation) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *TimedAnnotation) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *TimedAnnotation) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetPanels

`func (o *TimedAnnotation) GetPanels() []string`

GetPanels returns the Panels field if non-nil, zero value otherwise.

### GetPanelsOk

`func (o *TimedAnnotation) GetPanelsOk() (*[]string, bool)`

GetPanelsOk returns a tuple with the Panels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanels

`func (o *TimedAnnotation) SetPanels(v []string)`

SetPanels sets Panels field to given value.


### GetStartTime

`func (o *TimedAnnotation) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TimedAnnotation) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TimedAnnotation) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.


### GetTags

`func (o *TimedAnnotation) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TimedAnnotation) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TimedAnnotation) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetText

`func (o *TimedAnnotation) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TimedAnnotation) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TimedAnnotation) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *TimedAnnotation) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *TimedAnnotation) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *TimedAnnotation) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetTitle

`func (o *TimedAnnotation) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TimedAnnotation) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TimedAnnotation) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


