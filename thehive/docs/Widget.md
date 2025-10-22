# Widget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Fields** | Pointer to [**[]TaskField**](TaskField.md) |  | [optional] 
**Filter** | Pointer to [**Filter**](Filter.md) |  | [optional] 
**Sort** | Pointer to [**[]AlertListSortInner**](AlertListSortInner.md) |  | [optional] 
**MaxElements** | Pointer to **int32** |  | [optional] 
**Kind** | [**Kind**](Kind.md) |  | 
**Columns** | Pointer to [**[]TaskField**](TaskField.md) |  | [optional] 
**AttachmentId** | **string** |  | 
**ProtectData** | **bool** |  | 
**LogColumns** | Pointer to [**[]LogField**](LogField.md) |  | [optional] 
**WithTaskLogs** | Pointer to **bool** |  | [optional] [default to true]
**Template** | **string** |  | 
**Events** | Pointer to [**[]TimelineEvents**](TimelineEvents.md) |  | [optional] 
**WithCustomEventsDescription** | **bool** |  | 

## Methods

### NewWidget

`func NewWidget(kind Kind, attachmentId string, protectData bool, template string, withCustomEventsDescription bool, ) *Widget`

NewWidget instantiates a new Widget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetWithDefaults

`func NewWidgetWithDefaults() *Widget`

NewWidgetWithDefaults instantiates a new Widget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *Widget) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Widget) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Widget) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Widget) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFields

`func (o *Widget) GetFields() []TaskField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *Widget) GetFieldsOk() (*[]TaskField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *Widget) SetFields(v []TaskField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *Widget) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFilter

`func (o *Widget) GetFilter() Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *Widget) GetFilterOk() (*Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *Widget) SetFilter(v Filter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *Widget) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSort

`func (o *Widget) GetSort() []AlertListSortInner`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *Widget) GetSortOk() (*[]AlertListSortInner, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *Widget) SetSort(v []AlertListSortInner)`

SetSort sets Sort field to given value.

### HasSort

`func (o *Widget) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetMaxElements

`func (o *Widget) GetMaxElements() int32`

GetMaxElements returns the MaxElements field if non-nil, zero value otherwise.

### GetMaxElementsOk

`func (o *Widget) GetMaxElementsOk() (*int32, bool)`

GetMaxElementsOk returns a tuple with the MaxElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxElements

`func (o *Widget) SetMaxElements(v int32)`

SetMaxElements sets MaxElements field to given value.

### HasMaxElements

`func (o *Widget) HasMaxElements() bool`

HasMaxElements returns a boolean if a field has been set.

### GetKind

`func (o *Widget) GetKind() Kind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Widget) GetKindOk() (*Kind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Widget) SetKind(v Kind)`

SetKind sets Kind field to given value.


### GetColumns

`func (o *Widget) GetColumns() []TaskField`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *Widget) GetColumnsOk() (*[]TaskField, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *Widget) SetColumns(v []TaskField)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *Widget) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetAttachmentId

`func (o *Widget) GetAttachmentId() string`

GetAttachmentId returns the AttachmentId field if non-nil, zero value otherwise.

### GetAttachmentIdOk

`func (o *Widget) GetAttachmentIdOk() (*string, bool)`

GetAttachmentIdOk returns a tuple with the AttachmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentId

`func (o *Widget) SetAttachmentId(v string)`

SetAttachmentId sets AttachmentId field to given value.


### GetProtectData

`func (o *Widget) GetProtectData() bool`

GetProtectData returns the ProtectData field if non-nil, zero value otherwise.

### GetProtectDataOk

`func (o *Widget) GetProtectDataOk() (*bool, bool)`

GetProtectDataOk returns a tuple with the ProtectData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectData

`func (o *Widget) SetProtectData(v bool)`

SetProtectData sets ProtectData field to given value.


### GetLogColumns

`func (o *Widget) GetLogColumns() []LogField`

GetLogColumns returns the LogColumns field if non-nil, zero value otherwise.

### GetLogColumnsOk

`func (o *Widget) GetLogColumnsOk() (*[]LogField, bool)`

GetLogColumnsOk returns a tuple with the LogColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogColumns

`func (o *Widget) SetLogColumns(v []LogField)`

SetLogColumns sets LogColumns field to given value.

### HasLogColumns

`func (o *Widget) HasLogColumns() bool`

HasLogColumns returns a boolean if a field has been set.

### GetWithTaskLogs

`func (o *Widget) GetWithTaskLogs() bool`

GetWithTaskLogs returns the WithTaskLogs field if non-nil, zero value otherwise.

### GetWithTaskLogsOk

`func (o *Widget) GetWithTaskLogsOk() (*bool, bool)`

GetWithTaskLogsOk returns a tuple with the WithTaskLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithTaskLogs

`func (o *Widget) SetWithTaskLogs(v bool)`

SetWithTaskLogs sets WithTaskLogs field to given value.

### HasWithTaskLogs

`func (o *Widget) HasWithTaskLogs() bool`

HasWithTaskLogs returns a boolean if a field has been set.

### GetTemplate

`func (o *Widget) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *Widget) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *Widget) SetTemplate(v string)`

SetTemplate sets Template field to given value.


### GetEvents

`func (o *Widget) GetEvents() []TimelineEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Widget) GetEventsOk() (*[]TimelineEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Widget) SetEvents(v []TimelineEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *Widget) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetWithCustomEventsDescription

`func (o *Widget) GetWithCustomEventsDescription() bool`

GetWithCustomEventsDescription returns the WithCustomEventsDescription field if non-nil, zero value otherwise.

### GetWithCustomEventsDescriptionOk

`func (o *Widget) GetWithCustomEventsDescriptionOk() (*bool, bool)`

GetWithCustomEventsDescriptionOk returns a tuple with the WithCustomEventsDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithCustomEventsDescription

`func (o *Widget) SetWithCustomEventsDescription(v bool)`

SetWithCustomEventsDescription sets WithCustomEventsDescription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


