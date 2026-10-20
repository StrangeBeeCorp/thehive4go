# OutputTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnderscoreId** | **string** | Internal unique identifier of the task. |
**UnderscoreType** | **string** | Entity type, always &#x60;Task&#x60;. |
**UnderscoreCreatedBy** | **string** | Login of the user who created the task. |
**UnderscoreUpdatedBy** | Pointer to **string** | Login of the user who last updated the task. | [optional]
**UnderscoreCreatedAt** | **int64** | Date and time when the task was created. Returned as a Unix timestamp in milliseconds. |
**UnderscoreUpdatedAt** | Pointer to **int64** | Date and time when the task was last updated. Returned as a Unix timestamp in milliseconds. | [optional]
**Title** | **string** | Title of the task. |
**Group** | **string** | Group name used to organize the task within the case. |
**Description** | Pointer to **string** | Details of the task. | [optional]
**Status** | **string** | Current status of the task. One of &#x60;Waiting&#x60;, &#x60;InProgress&#x60;, &#x60;Completed&#x60;,or &#x60;Cancel&#x60;. |
**Flag** | **bool** | Whether the task is flagged as important. |
**StartDate** | Pointer to **int64** | Date and time when the task started. Returned as a Unix timestamp in milliseconds. | [optional]
**EndDate** | Pointer to **int64** | Date and time when the task ended. Returned as a Unix timestamp in milliseconds. | [optional]
**Assignee** | Pointer to **string** | Login of the user assigned to the task. | [optional]
**Order** | **int32** | Display order of the task within its group. |
**DueDate** | Pointer to **int64** | Due date for the task. Returned as a Unix timestamp in milliseconds. | [optional]
**Mandatory** | **bool** | Whether the task must be completed before the case can be closed. |
**ExtraData** | **map[string]interface{}** |  |
**TimeToHandle** | Pointer to **int64** | Duration of the task from start to completion, in milliseconds. | [optional]

## Methods

### NewOutputTask

`func NewOutputTask(underscoreId string, underscoreType string, underscoreCreatedBy string, underscoreCreatedAt int64, title string, group string, status string, flag bool, order int32, mandatory bool, extraData map[string]interface{}, ) *OutputTask`

NewOutputTask instantiates a new OutputTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputTaskWithDefaults

`func NewOutputTaskWithDefaults() *OutputTask`

NewOutputTaskWithDefaults instantiates a new OutputTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnderscoreId

`func (o *OutputTask) GetUnderscoreId() string`

GetUnderscoreId returns the UnderscoreId field if non-nil, zero value otherwise.

### GetUnderscoreIdOk

`func (o *OutputTask) GetUnderscoreIdOk() (*string, bool)`

GetUnderscoreIdOk returns a tuple with the UnderscoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreId

`func (o *OutputTask) SetUnderscoreId(v string)`

SetUnderscoreId sets UnderscoreId field to given value.


### GetUnderscoreType

`func (o *OutputTask) GetUnderscoreType() string`

GetUnderscoreType returns the UnderscoreType field if non-nil, zero value otherwise.

### GetUnderscoreTypeOk

`func (o *OutputTask) GetUnderscoreTypeOk() (*string, bool)`

GetUnderscoreTypeOk returns a tuple with the UnderscoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreType

`func (o *OutputTask) SetUnderscoreType(v string)`

SetUnderscoreType sets UnderscoreType field to given value.


### GetUnderscoreCreatedBy

`func (o *OutputTask) GetUnderscoreCreatedBy() string`

GetUnderscoreCreatedBy returns the UnderscoreCreatedBy field if non-nil, zero value otherwise.

### GetUnderscoreCreatedByOk

`func (o *OutputTask) GetUnderscoreCreatedByOk() (*string, bool)`

GetUnderscoreCreatedByOk returns a tuple with the UnderscoreCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedBy

`func (o *OutputTask) SetUnderscoreCreatedBy(v string)`

SetUnderscoreCreatedBy sets UnderscoreCreatedBy field to given value.


### GetUnderscoreUpdatedBy

`func (o *OutputTask) GetUnderscoreUpdatedBy() string`

GetUnderscoreUpdatedBy returns the UnderscoreUpdatedBy field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedByOk

`func (o *OutputTask) GetUnderscoreUpdatedByOk() (*string, bool)`

GetUnderscoreUpdatedByOk returns a tuple with the UnderscoreUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedBy

`func (o *OutputTask) SetUnderscoreUpdatedBy(v string)`

SetUnderscoreUpdatedBy sets UnderscoreUpdatedBy field to given value.

### HasUnderscoreUpdatedBy

`func (o *OutputTask) HasUnderscoreUpdatedBy() bool`

HasUnderscoreUpdatedBy returns a boolean if a field has been set.

### GetUnderscoreCreatedAt

`func (o *OutputTask) GetUnderscoreCreatedAt() int64`

GetUnderscoreCreatedAt returns the UnderscoreCreatedAt field if non-nil, zero value otherwise.

### GetUnderscoreCreatedAtOk

`func (o *OutputTask) GetUnderscoreCreatedAtOk() (*int64, bool)`

GetUnderscoreCreatedAtOk returns a tuple with the UnderscoreCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedAt

`func (o *OutputTask) SetUnderscoreCreatedAt(v int64)`

SetUnderscoreCreatedAt sets UnderscoreCreatedAt field to given value.


### GetUnderscoreUpdatedAt

`func (o *OutputTask) GetUnderscoreUpdatedAt() int64`

GetUnderscoreUpdatedAt returns the UnderscoreUpdatedAt field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedAtOk

`func (o *OutputTask) GetUnderscoreUpdatedAtOk() (*int64, bool)`

GetUnderscoreUpdatedAtOk returns a tuple with the UnderscoreUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedAt

`func (o *OutputTask) SetUnderscoreUpdatedAt(v int64)`

SetUnderscoreUpdatedAt sets UnderscoreUpdatedAt field to given value.

### HasUnderscoreUpdatedAt

`func (o *OutputTask) HasUnderscoreUpdatedAt() bool`

HasUnderscoreUpdatedAt returns a boolean if a field has been set.

### GetTitle

`func (o *OutputTask) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OutputTask) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OutputTask) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetGroup

`func (o *OutputTask) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *OutputTask) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *OutputTask) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetDescription

`func (o *OutputTask) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OutputTask) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OutputTask) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OutputTask) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *OutputTask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OutputTask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OutputTask) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetFlag

`func (o *OutputTask) GetFlag() bool`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *OutputTask) GetFlagOk() (*bool, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *OutputTask) SetFlag(v bool)`

SetFlag sets Flag field to given value.


### GetStartDate

`func (o *OutputTask) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *OutputTask) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *OutputTask) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *OutputTask) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *OutputTask) GetEndDate() int64`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *OutputTask) GetEndDateOk() (*int64, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *OutputTask) SetEndDate(v int64)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *OutputTask) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetAssignee

`func (o *OutputTask) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *OutputTask) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *OutputTask) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *OutputTask) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetOrder

`func (o *OutputTask) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *OutputTask) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *OutputTask) SetOrder(v int32)`

SetOrder sets Order field to given value.


### GetDueDate

`func (o *OutputTask) GetDueDate() int64`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *OutputTask) GetDueDateOk() (*int64, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *OutputTask) SetDueDate(v int64)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *OutputTask) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetMandatory

`func (o *OutputTask) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *OutputTask) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *OutputTask) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.


### GetExtraData

`func (o *OutputTask) GetExtraData() map[string]interface{}`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *OutputTask) GetExtraDataOk() (*map[string]interface{}, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *OutputTask) SetExtraData(v map[string]interface{})`

SetExtraData sets ExtraData field to given value.


### GetTimeToHandle

`func (o *OutputTask) GetTimeToHandle() int64`

GetTimeToHandle returns the TimeToHandle field if non-nil, zero value otherwise.

### GetTimeToHandleOk

`func (o *OutputTask) GetTimeToHandleOk() (*int64, bool)`

GetTimeToHandleOk returns a tuple with the TimeToHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToHandle

`func (o *OutputTask) SetTimeToHandle(v int64)`

SetTimeToHandle sets TimeToHandle field to given value.

### HasTimeToHandle

`func (o *OutputTask) HasTimeToHandle() bool`

HasTimeToHandle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
