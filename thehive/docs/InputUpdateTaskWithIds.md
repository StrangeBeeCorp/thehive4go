# InputUpdateTaskWithIds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** |  |
**Title** | Pointer to **string** |  | [optional]
**Group** | Pointer to **string** |  | [optional]
**Description** | Pointer to **string** |  | [optional]
**Status** | Pointer to **string** |  | [optional]
**Flag** | Pointer to **bool** |  | [optional]
**StartDate** | Pointer to **int32** |  | [optional]
**EndDate** | Pointer to **int64** |  | [optional]
**Order** | Pointer to **int32** |  | [optional]
**DueDate** | Pointer to **int32** |  | [optional]
**Assignee** | Pointer to **string** |  | [optional]
**Mandatory** | Pointer to **bool** |  | [optional]

## Methods

### NewInputUpdateTaskWithIds

`func NewInputUpdateTaskWithIds(ids []string, ) *InputUpdateTaskWithIds`

NewInputUpdateTaskWithIds instantiates a new InputUpdateTaskWithIds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputUpdateTaskWithIdsWithDefaults

`func NewInputUpdateTaskWithIdsWithDefaults() *InputUpdateTaskWithIds`

NewInputUpdateTaskWithIdsWithDefaults instantiates a new InputUpdateTaskWithIds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *InputUpdateTaskWithIds) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *InputUpdateTaskWithIds) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *InputUpdateTaskWithIds) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetTitle

`func (o *InputUpdateTaskWithIds) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InputUpdateTaskWithIds) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InputUpdateTaskWithIds) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InputUpdateTaskWithIds) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetGroup

`func (o *InputUpdateTaskWithIds) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InputUpdateTaskWithIds) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InputUpdateTaskWithIds) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *InputUpdateTaskWithIds) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetDescription

`func (o *InputUpdateTaskWithIds) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputUpdateTaskWithIds) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputUpdateTaskWithIds) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InputUpdateTaskWithIds) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *InputUpdateTaskWithIds) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InputUpdateTaskWithIds) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InputUpdateTaskWithIds) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InputUpdateTaskWithIds) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFlag

`func (o *InputUpdateTaskWithIds) GetFlag() bool`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *InputUpdateTaskWithIds) GetFlagOk() (*bool, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *InputUpdateTaskWithIds) SetFlag(v bool)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *InputUpdateTaskWithIds) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetStartDate

`func (o *InputUpdateTaskWithIds) GetStartDate() int32`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InputUpdateTaskWithIds) GetStartDateOk() (*int32, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InputUpdateTaskWithIds) SetStartDate(v int32)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InputUpdateTaskWithIds) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *InputUpdateTaskWithIds) GetEndDate() int64`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *InputUpdateTaskWithIds) GetEndDateOk() (*int64, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *InputUpdateTaskWithIds) SetEndDate(v int64)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *InputUpdateTaskWithIds) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetOrder

`func (o *InputUpdateTaskWithIds) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *InputUpdateTaskWithIds) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *InputUpdateTaskWithIds) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *InputUpdateTaskWithIds) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetDueDate

`func (o *InputUpdateTaskWithIds) GetDueDate() int32`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *InputUpdateTaskWithIds) GetDueDateOk() (*int32, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *InputUpdateTaskWithIds) SetDueDate(v int32)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *InputUpdateTaskWithIds) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetAssignee

`func (o *InputUpdateTaskWithIds) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *InputUpdateTaskWithIds) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *InputUpdateTaskWithIds) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *InputUpdateTaskWithIds) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetMandatory

`func (o *InputUpdateTaskWithIds) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *InputUpdateTaskWithIds) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *InputUpdateTaskWithIds) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.

### HasMandatory

`func (o *InputUpdateTaskWithIds) HasMandatory() bool`

HasMandatory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
