# InputImportCase1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaseTemplate** | Pointer to **string** | Name of the case template to apply to the created case. Must match an existing case template. Use &#x60;POST /api/v1/query&#x60; with &#x60;listCaseTemplate&#x60; to list available case templates. | [optional]
**Assignee** | Pointer to **string** | Login of the user to assign as owner of the case. Typically an email address. If a default domain is configured in the [authentication settings](https://docs.strangebee.com/thehive/administration/authentication/configure-authentication), a plain username is also accepted. | [optional]
**Tasks** | Pointer to [**[]InputCreateTask**](InputCreateTask.md) | Tasks to create in the case. A task represents one or more actions, called task logs, that must be completed during the investigation. If omitted and a case template is provided, the template&#39;s tasks are applied instead. | [optional]
**Pages** | Pointer to [**[]InputCreatePage**](InputCreatePage.md) | Pages to create in the case for structured documentation such as lessons learned, investigation reports, or meeting notes. [Case pages are temporary](https://docs.strangebee.com/thehive/user-guides/knowledge-base/about-case-pages/#permanent-vs-temporary-documentation): they&#39;re scoped to the active investigation and can&#39;t be searched. If omitted and a case template is provided, the template&#39;s pages are applied instead. | [optional]
**CustomFields** | Pointer to [**InputCreateCaseCustomFields**](InputCreateCaseCustomFields.md) |  | [optional]
**SharingParameters** | Pointer to [**[]InputShare**](InputShare.md) | Per-organization sharing configuration for this case. Overrides the [organization-level sharing rules](https://docs.strangebee.com/thehive/administration/organizations/about-organizations-sharing-rules/) for specific linked organizations. | [optional]
**TaskRule** | Pointer to **string** | Enum: &#x60;autoShare&#x60; &#x60;manual&#x60;&lt;br&gt;Task-sharing rule for the owner organization: &#x60;autoShare&#x60; shares new tasks with every sharing organization automatically, &#x60;manual&#x60; requires explicit sharing per task. Overrides the [organization-level sharing rules](https://docs.strangebee.com/thehive/administration/organizations/about-organizations-sharing-rules/). | [optional]
**ObservableRule** | Pointer to **string** | Enum: &#x60;autoShare&#x60; &#x60;manual&#x60;&lt;br&gt;Observable-sharing rule for the owner organization: &#x60;autoShare&#x60; shares new observables with every sharing organization automatically, &#x60;manual&#x60; requires explicit sharing per observable. Overrides the [organization-level sharing rules](https://docs.strangebee.com/thehive/administration/organizations/about-organizations-sharing-rules/). | [optional]

## Methods

### NewInputImportCase1

`func NewInputImportCase1() *InputImportCase1`

NewInputImportCase1 instantiates a new InputImportCase1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputImportCase1WithDefaults

`func NewInputImportCase1WithDefaults() *InputImportCase1`

NewInputImportCase1WithDefaults instantiates a new InputImportCase1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaseTemplate

`func (o *InputImportCase1) GetCaseTemplate() string`

GetCaseTemplate returns the CaseTemplate field if non-nil, zero value otherwise.

### GetCaseTemplateOk

`func (o *InputImportCase1) GetCaseTemplateOk() (*string, bool)`

GetCaseTemplateOk returns a tuple with the CaseTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseTemplate

`func (o *InputImportCase1) SetCaseTemplate(v string)`

SetCaseTemplate sets CaseTemplate field to given value.

### HasCaseTemplate

`func (o *InputImportCase1) HasCaseTemplate() bool`

HasCaseTemplate returns a boolean if a field has been set.

### GetAssignee

`func (o *InputImportCase1) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *InputImportCase1) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *InputImportCase1) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *InputImportCase1) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetTasks

`func (o *InputImportCase1) GetTasks() []InputCreateTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *InputImportCase1) GetTasksOk() (*[]InputCreateTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *InputImportCase1) SetTasks(v []InputCreateTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *InputImportCase1) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetPages

`func (o *InputImportCase1) GetPages() []InputCreatePage`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *InputImportCase1) GetPagesOk() (*[]InputCreatePage, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *InputImportCase1) SetPages(v []InputCreatePage)`

SetPages sets Pages field to given value.

### HasPages

`func (o *InputImportCase1) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetCustomFields

`func (o *InputImportCase1) GetCustomFields() InputCreateCaseCustomFields`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *InputImportCase1) GetCustomFieldsOk() (*InputCreateCaseCustomFields, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *InputImportCase1) SetCustomFields(v InputCreateCaseCustomFields)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *InputImportCase1) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetSharingParameters

`func (o *InputImportCase1) GetSharingParameters() []InputShare`

GetSharingParameters returns the SharingParameters field if non-nil, zero value otherwise.

### GetSharingParametersOk

`func (o *InputImportCase1) GetSharingParametersOk() (*[]InputShare, bool)`

GetSharingParametersOk returns a tuple with the SharingParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingParameters

`func (o *InputImportCase1) SetSharingParameters(v []InputShare)`

SetSharingParameters sets SharingParameters field to given value.

### HasSharingParameters

`func (o *InputImportCase1) HasSharingParameters() bool`

HasSharingParameters returns a boolean if a field has been set.

### GetTaskRule

`func (o *InputImportCase1) GetTaskRule() string`

GetTaskRule returns the TaskRule field if non-nil, zero value otherwise.

### GetTaskRuleOk

`func (o *InputImportCase1) GetTaskRuleOk() (*string, bool)`

GetTaskRuleOk returns a tuple with the TaskRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRule

`func (o *InputImportCase1) SetTaskRule(v string)`

SetTaskRule sets TaskRule field to given value.

### HasTaskRule

`func (o *InputImportCase1) HasTaskRule() bool`

HasTaskRule returns a boolean if a field has been set.

### GetObservableRule

`func (o *InputImportCase1) GetObservableRule() string`

GetObservableRule returns the ObservableRule field if non-nil, zero value otherwise.

### GetObservableRuleOk

`func (o *InputImportCase1) GetObservableRuleOk() (*string, bool)`

GetObservableRuleOk returns a tuple with the ObservableRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservableRule

`func (o *InputImportCase1) SetObservableRule(v string)`

SetObservableRule sets ObservableRule field to given value.

### HasObservableRule

`func (o *InputImportCase1) HasObservableRule() bool`

HasObservableRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
