# Timeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Events** | Pointer to [**[]TimelineEvents**](TimelineEvents.md) |  | [optional] 
**WithCustomEventsDescription** | **bool** |  | 
**Kind** | [**Kind**](Kind.md) |  | 

## Methods

### NewTimeline

`func NewTimeline(withCustomEventsDescription bool, kind Kind, ) *Timeline`

NewTimeline instantiates a new Timeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineWithDefaults

`func NewTimelineWithDefaults() *Timeline`

NewTimelineWithDefaults instantiates a new Timeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *Timeline) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Timeline) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Timeline) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Timeline) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetEvents

`func (o *Timeline) GetEvents() []TimelineEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Timeline) GetEventsOk() (*[]TimelineEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Timeline) SetEvents(v []TimelineEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *Timeline) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetWithCustomEventsDescription

`func (o *Timeline) GetWithCustomEventsDescription() bool`

GetWithCustomEventsDescription returns the WithCustomEventsDescription field if non-nil, zero value otherwise.

### GetWithCustomEventsDescriptionOk

`func (o *Timeline) GetWithCustomEventsDescriptionOk() (*bool, bool)`

GetWithCustomEventsDescriptionOk returns a tuple with the WithCustomEventsDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithCustomEventsDescription

`func (o *Timeline) SetWithCustomEventsDescription(v bool)`

SetWithCustomEventsDescription sets WithCustomEventsDescription field to given value.


### GetKind

`func (o *Timeline) GetKind() Kind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Timeline) GetKindOk() (*Kind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Timeline) SetKind(v Kind)`

SetKind sets Kind field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


