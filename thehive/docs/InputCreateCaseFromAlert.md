# InputCreateCaseFromAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Title of the case. When omitted, the alert title is used. | [optional]
**Description** | Pointer to **string** | Details of the case. Supports [TheHive-flavored Markdown](https://docs.strangebee.com/thehive/user-guides/thehive-flavored-markdown/). When omitted, the alert description is used. | [optional]
**Severity** | Pointer to **int32** | Numeric severity level of the case: 1 (Low), 2 (Medium), 3 (High), 4 (Critical). When omitted, the alert severity is used. | [optional] [default to 2]
**StartDate** | Pointer to **int64** | Start date of the incident. Unix timestamp in milliseconds. When omitted, the alert date is used. | [optional]
**EndDate** | Pointer to **int64** | End date of the incident. Unix timestamp in milliseconds. | [optional]
**Tags** | Pointer to **[]string** | Labels to categorize the case. Unknown tags are automatically created. Use &#x60;POST /api/v1/query&#x60; with &#x60;listTag&#x60; to retrieve existing custom tags, or with &#x60;listTaxonomy&#x60; and &#x60;tags&#x60; to retrieve taxonomy tags. When omitted, the alert tags are used. | [optional]
**Flag** | Pointer to **bool** | Marks the case as important for quick filtering. | [optional] [default to false]
**Tlp** | Pointer to **int32** | [Traffic Light Protocol (TLP)](https://www.misp-project.org/taxonomies.html#_tlp) level controlling how case information can be shared: 0 (Clear — no restrictions), 1 (Green — community), 2 (Amber — organization and clients), 3 (Amber+Strict — organization only), 4 (Red — named recipients only). When omitted, the alert TLP is used. | [optional] [default to 2]
**Pap** | Pointer to **int32** | [Permissible Actions Protocol (PAP)](https://www.misp-project.org/taxonomies.html#_pap) level controlling what actions can be taken with case information: 0 (Clear — no restrictions), 1 (Green — active actions allowed), 2 (Amber — passive checks only), 3 (Red — non-detectable actions only). When omitted, the alert PAP is used. | [optional] [default to 2]
**Status** | Pointer to **string** | Name of the case status to apply. Must match an existing status. Use &#x60;POST /api/v1/query&#x60; with &#x60;listCaseStatus&#x60; to list available statuses. | [optional] [default to "New"]
**Summary** | Pointer to **string** | Brief summary of the investigation findings. Supports [TheHive-flavored Markdown](https://docs.strangebee.com/thehive/user-guides/thehive-flavored-markdown/). | [optional]
**Assignee** | Pointer to **string** | Login of the user to assign the case to. Typically an email address. If a default domain is configured in the [authentication settings](https://docs.strangebee.com/thehive/administration/authentication/configure-authentication), a plain username is also accepted. Use &#x60;POST /api/v1/query&#x60; with &#x60;listUser&#x60; to retrieve available users and their logins. If omitted, the case is unassigned. | [optional]
**CustomFields** | Pointer to [**InputCreateCaseCustomFields**](InputCreateCaseCustomFields.md) |  | [optional]
**CaseTemplate** | Pointer to **string** | Name or ID of the case template to use. Applies the template&#39;s tasks, custom fields, and pages. Must match an existing case template. Use &#x60;POST /api/v1/query&#x60; with &#x60;listCaseTemplate&#x60; to list available case templates. | [optional]
**Tasks** | Pointer to [**[]InputCreateTask**](InputCreateTask.md) | Additional tasks to create in the case, on top of any tasks inherited from the alert or a case template. | [optional]
**Pages** | Pointer to [**[]InputCreatePage**](InputCreatePage.md) | Additional pages to create in the case. Pages are collaborative notes attached to the case for sharing information within the team during an investigation, such as lessons learned or meeting notes. | [optional]
**SharingParameters** | Pointer to [**[]InputShare**](InputShare.md) | Per-organization sharing configuration for this case. Overrides the [organization-level sharing rules](https://docs.strangebee.com/thehive/administration/organizations/about-organizations-sharing-rules/) for specific linked organizations. | [optional]
**TaskRule** | Pointer to **string** | Enum: &#x60;autoShare&#x60; &#x60;manual&#x60;&lt;br&gt;Task-sharing rule for the owner organization: &#x60;autoShare&#x60; shares new tasks with every sharing organization automatically, &#x60;manual&#x60; requires explicit sharing per task. Overrides the [organization-level sharing rules](https://docs.strangebee.com/thehive/administration/organizations/about-organizations-sharing-rules/). | [optional]
**ObservableRule** | Pointer to **string** | Enum: &#x60;autoShare&#x60; &#x60;manual&#x60;&lt;br&gt;Observable-sharing rule for the owner organization: &#x60;autoShare&#x60; shares new observables with every sharing organization automatically, &#x60;manual&#x60; requires explicit sharing per observable. Overrides the [organization-level sharing rules](https://docs.strangebee.com/thehive/administration/organizations/about-organizations-sharing-rules/). | [optional]

## Methods

### NewInputCreateCaseFromAlert

`func NewInputCreateCaseFromAlert() *InputCreateCaseFromAlert`

NewInputCreateCaseFromAlert instantiates a new InputCreateCaseFromAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputCreateCaseFromAlertWithDefaults

`func NewInputCreateCaseFromAlertWithDefaults() *InputCreateCaseFromAlert`

NewInputCreateCaseFromAlertWithDefaults instantiates a new InputCreateCaseFromAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *InputCreateCaseFromAlert) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InputCreateCaseFromAlert) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InputCreateCaseFromAlert) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InputCreateCaseFromAlert) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *InputCreateCaseFromAlert) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputCreateCaseFromAlert) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputCreateCaseFromAlert) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InputCreateCaseFromAlert) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSeverity

`func (o *InputCreateCaseFromAlert) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *InputCreateCaseFromAlert) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *InputCreateCaseFromAlert) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *InputCreateCaseFromAlert) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetStartDate

`func (o *InputCreateCaseFromAlert) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InputCreateCaseFromAlert) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InputCreateCaseFromAlert) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InputCreateCaseFromAlert) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *InputCreateCaseFromAlert) GetEndDate() int64`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *InputCreateCaseFromAlert) GetEndDateOk() (*int64, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *InputCreateCaseFromAlert) SetEndDate(v int64)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *InputCreateCaseFromAlert) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetTags

`func (o *InputCreateCaseFromAlert) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InputCreateCaseFromAlert) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InputCreateCaseFromAlert) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InputCreateCaseFromAlert) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetFlag

`func (o *InputCreateCaseFromAlert) GetFlag() bool`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *InputCreateCaseFromAlert) GetFlagOk() (*bool, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *InputCreateCaseFromAlert) SetFlag(v bool)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *InputCreateCaseFromAlert) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetTlp

`func (o *InputCreateCaseFromAlert) GetTlp() int32`

GetTlp returns the Tlp field if non-nil, zero value otherwise.

### GetTlpOk

`func (o *InputCreateCaseFromAlert) GetTlpOk() (*int32, bool)`

GetTlpOk returns a tuple with the Tlp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlp

`func (o *InputCreateCaseFromAlert) SetTlp(v int32)`

SetTlp sets Tlp field to given value.

### HasTlp

`func (o *InputCreateCaseFromAlert) HasTlp() bool`

HasTlp returns a boolean if a field has been set.

### GetPap

`func (o *InputCreateCaseFromAlert) GetPap() int32`

GetPap returns the Pap field if non-nil, zero value otherwise.

### GetPapOk

`func (o *InputCreateCaseFromAlert) GetPapOk() (*int32, bool)`

GetPapOk returns a tuple with the Pap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPap

`func (o *InputCreateCaseFromAlert) SetPap(v int32)`

SetPap sets Pap field to given value.

### HasPap

`func (o *InputCreateCaseFromAlert) HasPap() bool`

HasPap returns a boolean if a field has been set.

### GetStatus

`func (o *InputCreateCaseFromAlert) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InputCreateCaseFromAlert) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InputCreateCaseFromAlert) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InputCreateCaseFromAlert) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSummary

`func (o *InputCreateCaseFromAlert) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *InputCreateCaseFromAlert) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *InputCreateCaseFromAlert) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *InputCreateCaseFromAlert) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetAssignee

`func (o *InputCreateCaseFromAlert) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *InputCreateCaseFromAlert) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *InputCreateCaseFromAlert) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *InputCreateCaseFromAlert) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetCustomFields

`func (o *InputCreateCaseFromAlert) GetCustomFields() InputCreateCaseCustomFields`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *InputCreateCaseFromAlert) GetCustomFieldsOk() (*InputCreateCaseCustomFields, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *InputCreateCaseFromAlert) SetCustomFields(v InputCreateCaseCustomFields)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *InputCreateCaseFromAlert) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCaseTemplate

`func (o *InputCreateCaseFromAlert) GetCaseTemplate() string`

GetCaseTemplate returns the CaseTemplate field if non-nil, zero value otherwise.

### GetCaseTemplateOk

`func (o *InputCreateCaseFromAlert) GetCaseTemplateOk() (*string, bool)`

GetCaseTemplateOk returns a tuple with the CaseTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseTemplate

`func (o *InputCreateCaseFromAlert) SetCaseTemplate(v string)`

SetCaseTemplate sets CaseTemplate field to given value.

### HasCaseTemplate

`func (o *InputCreateCaseFromAlert) HasCaseTemplate() bool`

HasCaseTemplate returns a boolean if a field has been set.

### GetTasks

`func (o *InputCreateCaseFromAlert) GetTasks() []InputCreateTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *InputCreateCaseFromAlert) GetTasksOk() (*[]InputCreateTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *InputCreateCaseFromAlert) SetTasks(v []InputCreateTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *InputCreateCaseFromAlert) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetPages

`func (o *InputCreateCaseFromAlert) GetPages() []InputCreatePage`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *InputCreateCaseFromAlert) GetPagesOk() (*[]InputCreatePage, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *InputCreateCaseFromAlert) SetPages(v []InputCreatePage)`

SetPages sets Pages field to given value.

### HasPages

`func (o *InputCreateCaseFromAlert) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetSharingParameters

`func (o *InputCreateCaseFromAlert) GetSharingParameters() []InputShare`

GetSharingParameters returns the SharingParameters field if non-nil, zero value otherwise.

### GetSharingParametersOk

`func (o *InputCreateCaseFromAlert) GetSharingParametersOk() (*[]InputShare, bool)`

GetSharingParametersOk returns a tuple with the SharingParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingParameters

`func (o *InputCreateCaseFromAlert) SetSharingParameters(v []InputShare)`

SetSharingParameters sets SharingParameters field to given value.

### HasSharingParameters

`func (o *InputCreateCaseFromAlert) HasSharingParameters() bool`

HasSharingParameters returns a boolean if a field has been set.

### GetTaskRule

`func (o *InputCreateCaseFromAlert) GetTaskRule() string`

GetTaskRule returns the TaskRule field if non-nil, zero value otherwise.

### GetTaskRuleOk

`func (o *InputCreateCaseFromAlert) GetTaskRuleOk() (*string, bool)`

GetTaskRuleOk returns a tuple with the TaskRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRule

`func (o *InputCreateCaseFromAlert) SetTaskRule(v string)`

SetTaskRule sets TaskRule field to given value.

### HasTaskRule

`func (o *InputCreateCaseFromAlert) HasTaskRule() bool`

HasTaskRule returns a boolean if a field has been set.

### GetObservableRule

`func (o *InputCreateCaseFromAlert) GetObservableRule() string`

GetObservableRule returns the ObservableRule field if non-nil, zero value otherwise.

### GetObservableRuleOk

`func (o *InputCreateCaseFromAlert) GetObservableRuleOk() (*string, bool)`

GetObservableRuleOk returns a tuple with the ObservableRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservableRule

`func (o *InputCreateCaseFromAlert) SetObservableRule(v string)`

SetObservableRule sets ObservableRule field to given value.

### HasObservableRule

`func (o *InputCreateCaseFromAlert) HasObservableRule() bool`

HasObservableRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
