# InputUpdateTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Flag** | Pointer to **bool** |  | [optional] 
**StartDate** | Pointer to **int32** |  | [optional] 
**EndDate** | Pointer to **int32** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**DueDate** | Pointer to **int32** |  | [optional] 
**Assignee** | Pointer to **string** |  | [optional] 
**Mandatory** | Pointer to **bool** |  | [optional] 

## Methods

### NewInputUpdateTask

`func NewInputUpdateTask() *InputUpdateTask`

NewInputUpdateTask instantiates a new InputUpdateTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputUpdateTaskWithDefaults

`func NewInputUpdateTaskWithDefaults() *InputUpdateTask`

NewInputUpdateTaskWithDefaults instantiates a new InputUpdateTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *InputUpdateTask) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InputUpdateTask) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InputUpdateTask) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InputUpdateTask) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetGroup

`func (o *InputUpdateTask) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InputUpdateTask) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InputUpdateTask) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *InputUpdateTask) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetDescription

`func (o *InputUpdateTask) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputUpdateTask) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputUpdateTask) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InputUpdateTask) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *InputUpdateTask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InputUpdateTask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InputUpdateTask) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InputUpdateTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFlag

`func (o *InputUpdateTask) GetFlag() bool`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *InputUpdateTask) GetFlagOk() (*bool, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *InputUpdateTask) SetFlag(v bool)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *InputUpdateTask) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetStartDate

`func (o *InputUpdateTask) GetStartDate() int32`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InputUpdateTask) GetStartDateOk() (*int32, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InputUpdateTask) SetStartDate(v int32)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InputUpdateTask) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *InputUpdateTask) GetEndDate() int32`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *InputUpdateTask) GetEndDateOk() (*int32, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *InputUpdateTask) SetEndDate(v int32)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *InputUpdateTask) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetOrder

`func (o *InputUpdateTask) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *InputUpdateTask) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *InputUpdateTask) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *InputUpdateTask) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetDueDate

`func (o *InputUpdateTask) GetDueDate() int32`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *InputUpdateTask) GetDueDateOk() (*int32, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *InputUpdateTask) SetDueDate(v int32)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *InputUpdateTask) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetAssignee

`func (o *InputUpdateTask) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *InputUpdateTask) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *InputUpdateTask) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *InputUpdateTask) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetMandatory

`func (o *InputUpdateTask) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *InputUpdateTask) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *InputUpdateTask) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.

### HasMandatory

`func (o *InputUpdateTask) HasMandatory() bool`

HasMandatory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


