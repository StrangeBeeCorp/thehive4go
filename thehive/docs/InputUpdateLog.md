# InputUpdateLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**IncludeInTimeline** | Pointer to **int64** |  | [optional] 

## Methods

### NewInputUpdateLog

`func NewInputUpdateLog() *InputUpdateLog`

NewInputUpdateLog instantiates a new InputUpdateLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputUpdateLogWithDefaults

`func NewInputUpdateLogWithDefaults() *InputUpdateLog`

NewInputUpdateLogWithDefaults instantiates a new InputUpdateLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *InputUpdateLog) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InputUpdateLog) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InputUpdateLog) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InputUpdateLog) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetIncludeInTimeline

`func (o *InputUpdateLog) GetIncludeInTimeline() int64`

GetIncludeInTimeline returns the IncludeInTimeline field if non-nil, zero value otherwise.

### GetIncludeInTimelineOk

`func (o *InputUpdateLog) GetIncludeInTimelineOk() (*int64, bool)`

GetIncludeInTimelineOk returns a tuple with the IncludeInTimeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInTimeline

`func (o *InputUpdateLog) SetIncludeInTimeline(v int64)`

SetIncludeInTimeline sets IncludeInTimeline field to given value.

### HasIncludeInTimeline

`func (o *InputUpdateLog) HasIncludeInTimeline() bool`

HasIncludeInTimeline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


