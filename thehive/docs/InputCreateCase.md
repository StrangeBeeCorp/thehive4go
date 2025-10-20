# InputCreateCase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  |
**Description** | **string** | Description of the case, supports markdown |
**Severity** | Pointer to **int32** |  | [optional] [default to 2]
**StartDate** | Pointer to **int64** |  | [optional] [default to now()]
**EndDate** | Pointer to **int64** |  | [optional]
**Tags** | Pointer to **[]string** |  | [optional]
**Flag** | Pointer to **bool** |  | [optional] [default to false]
**Tlp** | Pointer to **int32** |  | [optional] [default to 2]
**Pap** | Pointer to **int32** |  | [optional] [default to 2]
**Status** | Pointer to **string** |  | [optional] [default to "New"]
**Summary** | Pointer to **string** |  | [optional]
**Assignee** | Pointer to **string** | User to assign the case to | [optional]
**Access** | Pointer to [**Access**](Access.md) | Access type. The default value is the organisation based access | [optional]
**CustomFields** | Pointer to [**InputCreateAlertCustomFields**](InputCreateAlertCustomFields.md) |  | [optional]
**CaseTemplate** | Pointer to **string** | Name or id of the Case Template to use | [optional]
**Tasks** | Pointer to [**[]InputCreateTask**](InputCreateTask.md) | Tasks to create. If null, tasks from the Case Template will be used | [optional]
**Pages** | Pointer to [**[]InputCreatePage**](InputCreatePage.md) |  | [optional]
**SharingParameters** | Pointer to [**[]InputShare**](InputShare.md) |  | [optional]
**TaskRule** | Pointer to **string** |  | [optional]
**ObservableRule** | Pointer to **string** |  | [optional]

## Methods

### NewInputCreateCase

`func NewInputCreateCase(title string, description string, ) *InputCreateCase`

NewInputCreateCase instantiates a new InputCreateCase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputCreateCaseWithDefaults

`func NewInputCreateCaseWithDefaults() *InputCreateCase`

NewInputCreateCaseWithDefaults instantiates a new InputCreateCase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *InputCreateCase) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InputCreateCase) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InputCreateCase) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *InputCreateCase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputCreateCase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputCreateCase) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSeverity

`func (o *InputCreateCase) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *InputCreateCase) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *InputCreateCase) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *InputCreateCase) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetStartDate

`func (o *InputCreateCase) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InputCreateCase) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InputCreateCase) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InputCreateCase) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *InputCreateCase) GetEndDate() int64`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *InputCreateCase) GetEndDateOk() (*int64, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *InputCreateCase) SetEndDate(v int64)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *InputCreateCase) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetTags

`func (o *InputCreateCase) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InputCreateCase) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InputCreateCase) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InputCreateCase) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetFlag

`func (o *InputCreateCase) GetFlag() bool`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *InputCreateCase) GetFlagOk() (*bool, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *InputCreateCase) SetFlag(v bool)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *InputCreateCase) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetTlp

`func (o *InputCreateCase) GetTlp() int32`

GetTlp returns the Tlp field if non-nil, zero value otherwise.

### GetTlpOk

`func (o *InputCreateCase) GetTlpOk() (*int32, bool)`

GetTlpOk returns a tuple with the Tlp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlp

`func (o *InputCreateCase) SetTlp(v int32)`

SetTlp sets Tlp field to given value.

### HasTlp

`func (o *InputCreateCase) HasTlp() bool`

HasTlp returns a boolean if a field has been set.

### GetPap

`func (o *InputCreateCase) GetPap() int32`

GetPap returns the Pap field if non-nil, zero value otherwise.

### GetPapOk

`func (o *InputCreateCase) GetPapOk() (*int32, bool)`

GetPapOk returns a tuple with the Pap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPap

`func (o *InputCreateCase) SetPap(v int32)`

SetPap sets Pap field to given value.

### HasPap

`func (o *InputCreateCase) HasPap() bool`

HasPap returns a boolean if a field has been set.

### GetStatus

`func (o *InputCreateCase) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InputCreateCase) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InputCreateCase) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InputCreateCase) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSummary

`func (o *InputCreateCase) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *InputCreateCase) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *InputCreateCase) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *InputCreateCase) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetAssignee

`func (o *InputCreateCase) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *InputCreateCase) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *InputCreateCase) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *InputCreateCase) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetAccess

`func (o *InputCreateCase) GetAccess() Access`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *InputCreateCase) GetAccessOk() (*Access, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *InputCreateCase) SetAccess(v Access)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *InputCreateCase) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetCustomFields

`func (o *InputCreateCase) GetCustomFields() InputCreateAlertCustomFields`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *InputCreateCase) GetCustomFieldsOk() (*InputCreateAlertCustomFields, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *InputCreateCase) SetCustomFields(v InputCreateAlertCustomFields)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *InputCreateCase) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCaseTemplate

`func (o *InputCreateCase) GetCaseTemplate() string`

GetCaseTemplate returns the CaseTemplate field if non-nil, zero value otherwise.

### GetCaseTemplateOk

`func (o *InputCreateCase) GetCaseTemplateOk() (*string, bool)`

GetCaseTemplateOk returns a tuple with the CaseTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseTemplate

`func (o *InputCreateCase) SetCaseTemplate(v string)`

SetCaseTemplate sets CaseTemplate field to given value.

### HasCaseTemplate

`func (o *InputCreateCase) HasCaseTemplate() bool`

HasCaseTemplate returns a boolean if a field has been set.

### GetTasks

`func (o *InputCreateCase) GetTasks() []InputCreateTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *InputCreateCase) GetTasksOk() (*[]InputCreateTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *InputCreateCase) SetTasks(v []InputCreateTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *InputCreateCase) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetPages

`func (o *InputCreateCase) GetPages() []InputCreatePage`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *InputCreateCase) GetPagesOk() (*[]InputCreatePage, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *InputCreateCase) SetPages(v []InputCreatePage)`

SetPages sets Pages field to given value.

### HasPages

`func (o *InputCreateCase) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetSharingParameters

`func (o *InputCreateCase) GetSharingParameters() []InputShare`

GetSharingParameters returns the SharingParameters field if non-nil, zero value otherwise.

### GetSharingParametersOk

`func (o *InputCreateCase) GetSharingParametersOk() (*[]InputShare, bool)`

GetSharingParametersOk returns a tuple with the SharingParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingParameters

`func (o *InputCreateCase) SetSharingParameters(v []InputShare)`

SetSharingParameters sets SharingParameters field to given value.

### HasSharingParameters

`func (o *InputCreateCase) HasSharingParameters() bool`

HasSharingParameters returns a boolean if a field has been set.

### GetTaskRule

`func (o *InputCreateCase) GetTaskRule() string`

GetTaskRule returns the TaskRule field if non-nil, zero value otherwise.

### GetTaskRuleOk

`func (o *InputCreateCase) GetTaskRuleOk() (*string, bool)`

GetTaskRuleOk returns a tuple with the TaskRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRule

`func (o *InputCreateCase) SetTaskRule(v string)`

SetTaskRule sets TaskRule field to given value.

### HasTaskRule

`func (o *InputCreateCase) HasTaskRule() bool`

HasTaskRule returns a boolean if a field has been set.

### GetObservableRule

`func (o *InputCreateCase) GetObservableRule() string`

GetObservableRule returns the ObservableRule field if non-nil, zero value otherwise.

### GetObservableRuleOk

`func (o *InputCreateCase) GetObservableRuleOk() (*string, bool)`

GetObservableRuleOk returns a tuple with the ObservableRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservableRule

`func (o *InputCreateCase) SetObservableRule(v string)`

SetObservableRule sets ObservableRule field to given value.

### HasObservableRule

`func (o *InputCreateCase) HasObservableRule() bool`

HasObservableRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
