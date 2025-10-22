# OutputTimeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]OutputTimelineEvent**](OutputTimelineEvent.md) |  | [optional] 

## Methods

### NewOutputTimeline

`func NewOutputTimeline() *OutputTimeline`

NewOutputTimeline instantiates a new OutputTimeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputTimelineWithDefaults

`func NewOutputTimelineWithDefaults() *OutputTimeline`

NewOutputTimelineWithDefaults instantiates a new OutputTimeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *OutputTimeline) GetEvents() []OutputTimelineEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *OutputTimeline) GetEventsOk() (*[]OutputTimelineEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *OutputTimeline) SetEvents(v []OutputTimelineEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *OutputTimeline) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


