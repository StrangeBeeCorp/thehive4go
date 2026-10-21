# InputCreateTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Title of the task. |
**Group** | Pointer to **string** | Group name used to organize tasks within the case. | [optional]
**Description** | Pointer to **string** | Details of the task. Supports [TheHive-flavored Markdown](https://docs.strangebee.com/thehive/user-guides/thehive-flavored-markdown/). | [optional]
**Status** | Pointer to **string** | Enum: &#x60;Waiting&#x60; &#x60;InProgress&#x60; &#x60;Completed&#x60; &#x60;Cancel&#x60;&lt;br&gt;Initial status of the task. | [optional]
**Flag** | Pointer to **bool** | Marks the task as important for quick filtering. | [optional]
**StartDate** | Pointer to **int64** | Date and time when the task started. Unix timestamp in milliseconds. | [optional]
**EndDate** | Pointer to **int64** | Date and time when the task ended. Unix timestamp in milliseconds. | [optional]
**Order** | Pointer to **int32** | Display order of the task within its group. | [optional]
**DueDate** | Pointer to **int64** | Due date for the task. Unix timestamp in milliseconds. | [optional]
**Assignee** | Pointer to **string** | Email address of the user assigned to the task. Use &#x60;POST /api/v1/query&#x60; with &#x60;listUser&#x60; to retrieve available users and their logins. | [optional]
**Mandatory** | Pointer to **bool** | Whether the task must be completed before the case can be closed. | [optional]

## Methods

### NewInputCreateTask

`func NewInputCreateTask(title string, ) *InputCreateTask`

NewInputCreateTask instantiates a new InputCreateTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputCreateTaskWithDefaults

`func NewInputCreateTaskWithDefaults() *InputCreateTask`

NewInputCreateTaskWithDefaults instantiates a new InputCreateTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *InputCreateTask) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InputCreateTask) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InputCreateTask) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetGroup

`func (o *InputCreateTask) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InputCreateTask) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InputCreateTask) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *InputCreateTask) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetDescription

`func (o *InputCreateTask) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputCreateTask) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputCreateTask) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InputCreateTask) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *InputCreateTask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InputCreateTask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InputCreateTask) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InputCreateTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFlag

`func (o *InputCreateTask) GetFlag() bool`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *InputCreateTask) GetFlagOk() (*bool, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *InputCreateTask) SetFlag(v bool)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *InputCreateTask) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetStartDate

`func (o *InputCreateTask) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InputCreateTask) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InputCreateTask) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InputCreateTask) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *InputCreateTask) GetEndDate() int64`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *InputCreateTask) GetEndDateOk() (*int64, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *InputCreateTask) SetEndDate(v int64)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *InputCreateTask) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetOrder

`func (o *InputCreateTask) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *InputCreateTask) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *InputCreateTask) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *InputCreateTask) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetDueDate

`func (o *InputCreateTask) GetDueDate() int64`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *InputCreateTask) GetDueDateOk() (*int64, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *InputCreateTask) SetDueDate(v int64)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *InputCreateTask) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetAssignee

`func (o *InputCreateTask) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *InputCreateTask) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *InputCreateTask) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *InputCreateTask) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetMandatory

`func (o *InputCreateTask) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *InputCreateTask) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *InputCreateTask) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.

### HasMandatory

`func (o *InputCreateTask) HasMandatory() bool`

HasMandatory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
