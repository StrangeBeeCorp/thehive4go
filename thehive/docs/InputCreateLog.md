# InputCreateLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**StartDate** | Pointer to **int32** |  | [optional] 
**IncludeInTimeline** | Pointer to **int32** |  | [optional] 
**Attachments** | Pointer to **[]*os.File** | ignored with json content, use a multipart if you want to pass attachments | [optional] 

## Methods

### NewInputCreateLog

`func NewInputCreateLog(message string, ) *InputCreateLog`

NewInputCreateLog instantiates a new InputCreateLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputCreateLogWithDefaults

`func NewInputCreateLogWithDefaults() *InputCreateLog`

NewInputCreateLogWithDefaults instantiates a new InputCreateLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *InputCreateLog) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InputCreateLog) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InputCreateLog) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetStartDate

`func (o *InputCreateLog) GetStartDate() int32`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InputCreateLog) GetStartDateOk() (*int32, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InputCreateLog) SetStartDate(v int32)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InputCreateLog) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetIncludeInTimeline

`func (o *InputCreateLog) GetIncludeInTimeline() int32`

GetIncludeInTimeline returns the IncludeInTimeline field if non-nil, zero value otherwise.

### GetIncludeInTimelineOk

`func (o *InputCreateLog) GetIncludeInTimelineOk() (*int32, bool)`

GetIncludeInTimelineOk returns a tuple with the IncludeInTimeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInTimeline

`func (o *InputCreateLog) SetIncludeInTimeline(v int32)`

SetIncludeInTimeline sets IncludeInTimeline field to given value.

### HasIncludeInTimeline

`func (o *InputCreateLog) HasIncludeInTimeline() bool`

HasIncludeInTimeline returns a boolean if a field has been set.

### GetAttachments

`func (o *InputCreateLog) GetAttachments() []*os.File`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *InputCreateLog) GetAttachmentsOk() (*[]*os.File, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *InputCreateLog) SetAttachments(v []*os.File)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *InputCreateLog) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


