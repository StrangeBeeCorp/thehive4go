# InputImportCase1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaseTemplate** | Pointer to **string** |  | [optional] 
**Assignee** | Pointer to **string** |  | [optional] 
**Tasks** | Pointer to [**[]InputCreateTask**](InputCreateTask.md) |  | [optional] 
**Pages** | Pointer to [**[]InputCreatePage**](InputCreatePage.md) |  | [optional] 
**CustomFields** | Pointer to [**InputCreateAlertCustomFields**](InputCreateAlertCustomFields.md) |  | [optional] 
**SharingParameters** | Pointer to [**[]InputShare**](InputShare.md) |  | [optional] 
**TaskRule** | Pointer to **string** |  | [optional] 
**ObservableRule** | Pointer to **string** |  | [optional] 

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

`func (o *InputImportCase1) GetCustomFields() InputCreateAlertCustomFields`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *InputImportCase1) GetCustomFieldsOk() (*InputCreateAlertCustomFields, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *InputImportCase1) SetCustomFields(v InputCreateAlertCustomFields)`

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


