# InputApplyCaseTemplateWithIds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** |  | 
**CaseTemplate** | **string** | Id or name of the CaseTemplate to apply | 
**UpdateTitlePrefix** | Pointer to **bool** |  | [optional] [default to false]
**UpdateDescription** | Pointer to **bool** |  | [optional] [default to false]
**UpdateTags** | Pointer to **bool** |  | [optional] [default to false]
**UpdateSeverity** | Pointer to **bool** |  | [optional] [default to false]
**UpdateFlag** | Pointer to **bool** |  | [optional] [default to false]
**UpdateTlp** | Pointer to **bool** |  | [optional] [default to false]
**UpdatePap** | Pointer to **bool** |  | [optional] [default to false]
**UpdateCustomFields** | Pointer to **bool** |  | [optional] [default to false]
**ImportTasks** | Pointer to **[]string** | Define which tasks to import in the case | [optional] 
**ImportPages** | Pointer to **[]string** | Define which pages to import in the case | [optional] 

## Methods

### NewInputApplyCaseTemplateWithIds

`func NewInputApplyCaseTemplateWithIds(ids []string, caseTemplate string, ) *InputApplyCaseTemplateWithIds`

NewInputApplyCaseTemplateWithIds instantiates a new InputApplyCaseTemplateWithIds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputApplyCaseTemplateWithIdsWithDefaults

`func NewInputApplyCaseTemplateWithIdsWithDefaults() *InputApplyCaseTemplateWithIds`

NewInputApplyCaseTemplateWithIdsWithDefaults instantiates a new InputApplyCaseTemplateWithIds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *InputApplyCaseTemplateWithIds) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *InputApplyCaseTemplateWithIds) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *InputApplyCaseTemplateWithIds) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetCaseTemplate

`func (o *InputApplyCaseTemplateWithIds) GetCaseTemplate() string`

GetCaseTemplate returns the CaseTemplate field if non-nil, zero value otherwise.

### GetCaseTemplateOk

`func (o *InputApplyCaseTemplateWithIds) GetCaseTemplateOk() (*string, bool)`

GetCaseTemplateOk returns a tuple with the CaseTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseTemplate

`func (o *InputApplyCaseTemplateWithIds) SetCaseTemplate(v string)`

SetCaseTemplate sets CaseTemplate field to given value.


### GetUpdateTitlePrefix

`func (o *InputApplyCaseTemplateWithIds) GetUpdateTitlePrefix() bool`

GetUpdateTitlePrefix returns the UpdateTitlePrefix field if non-nil, zero value otherwise.

### GetUpdateTitlePrefixOk

`func (o *InputApplyCaseTemplateWithIds) GetUpdateTitlePrefixOk() (*bool, bool)`

GetUpdateTitlePrefixOk returns a tuple with the UpdateTitlePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTitlePrefix

`func (o *InputApplyCaseTemplateWithIds) SetUpdateTitlePrefix(v bool)`

SetUpdateTitlePrefix sets UpdateTitlePrefix field to given value.

### HasUpdateTitlePrefix

`func (o *InputApplyCaseTemplateWithIds) HasUpdateTitlePrefix() bool`

HasUpdateTitlePrefix returns a boolean if a field has been set.

### GetUpdateDescription

`func (o *InputApplyCaseTemplateWithIds) GetUpdateDescription() bool`

GetUpdateDescription returns the UpdateDescription field if non-nil, zero value otherwise.

### GetUpdateDescriptionOk

`func (o *InputApplyCaseTemplateWithIds) GetUpdateDescriptionOk() (*bool, bool)`

GetUpdateDescriptionOk returns a tuple with the UpdateDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDescription

`func (o *InputApplyCaseTemplateWithIds) SetUpdateDescription(v bool)`

SetUpdateDescription sets UpdateDescription field to given value.

### HasUpdateDescription

`func (o *InputApplyCaseTemplateWithIds) HasUpdateDescription() bool`

HasUpdateDescription returns a boolean if a field has been set.

### GetUpdateTags

`func (o *InputApplyCaseTemplateWithIds) GetUpdateTags() bool`

GetUpdateTags returns the UpdateTags field if non-nil, zero value otherwise.

### GetUpdateTagsOk

`func (o *InputApplyCaseTemplateWithIds) GetUpdateTagsOk() (*bool, bool)`

GetUpdateTagsOk returns a tuple with the UpdateTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTags

`func (o *InputApplyCaseTemplateWithIds) SetUpdateTags(v bool)`

SetUpdateTags sets UpdateTags field to given value.

### HasUpdateTags

`func (o *InputApplyCaseTemplateWithIds) HasUpdateTags() bool`

HasUpdateTags returns a boolean if a field has been set.

### GetUpdateSeverity

`func (o *InputApplyCaseTemplateWithIds) GetUpdateSeverity() bool`

GetUpdateSeverity returns the UpdateSeverity field if non-nil, zero value otherwise.

### GetUpdateSeverityOk

`func (o *InputApplyCaseTemplateWithIds) GetUpdateSeverityOk() (*bool, bool)`

GetUpdateSeverityOk returns a tuple with the UpdateSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSeverity

`func (o *InputApplyCaseTemplateWithIds) SetUpdateSeverity(v bool)`

SetUpdateSeverity sets UpdateSeverity field to given value.

### HasUpdateSeverity

`func (o *InputApplyCaseTemplateWithIds) HasUpdateSeverity() bool`

HasUpdateSeverity returns a boolean if a field has been set.

### GetUpdateFlag

`func (o *InputApplyCaseTemplateWithIds) GetUpdateFlag() bool`

GetUpdateFlag returns the UpdateFlag field if non-nil, zero value otherwise.

### GetUpdateFlagOk

`func (o *InputApplyCaseTemplateWithIds) GetUpdateFlagOk() (*bool, bool)`

GetUpdateFlagOk returns a tuple with the UpdateFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateFlag

`func (o *InputApplyCaseTemplateWithIds) SetUpdateFlag(v bool)`

SetUpdateFlag sets UpdateFlag field to given value.

### HasUpdateFlag

`func (o *InputApplyCaseTemplateWithIds) HasUpdateFlag() bool`

HasUpdateFlag returns a boolean if a field has been set.

### GetUpdateTlp

`func (o *InputApplyCaseTemplateWithIds) GetUpdateTlp() bool`

GetUpdateTlp returns the UpdateTlp field if non-nil, zero value otherwise.

### GetUpdateTlpOk

`func (o *InputApplyCaseTemplateWithIds) GetUpdateTlpOk() (*bool, bool)`

GetUpdateTlpOk returns a tuple with the UpdateTlp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTlp

`func (o *InputApplyCaseTemplateWithIds) SetUpdateTlp(v bool)`

SetUpdateTlp sets UpdateTlp field to given value.

### HasUpdateTlp

`func (o *InputApplyCaseTemplateWithIds) HasUpdateTlp() bool`

HasUpdateTlp returns a boolean if a field has been set.

### GetUpdatePap

`func (o *InputApplyCaseTemplateWithIds) GetUpdatePap() bool`

GetUpdatePap returns the UpdatePap field if non-nil, zero value otherwise.

### GetUpdatePapOk

`func (o *InputApplyCaseTemplateWithIds) GetUpdatePapOk() (*bool, bool)`

GetUpdatePapOk returns a tuple with the UpdatePap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatePap

`func (o *InputApplyCaseTemplateWithIds) SetUpdatePap(v bool)`

SetUpdatePap sets UpdatePap field to given value.

### HasUpdatePap

`func (o *InputApplyCaseTemplateWithIds) HasUpdatePap() bool`

HasUpdatePap returns a boolean if a field has been set.

### GetUpdateCustomFields

`func (o *InputApplyCaseTemplateWithIds) GetUpdateCustomFields() bool`

GetUpdateCustomFields returns the UpdateCustomFields field if non-nil, zero value otherwise.

### GetUpdateCustomFieldsOk

`func (o *InputApplyCaseTemplateWithIds) GetUpdateCustomFieldsOk() (*bool, bool)`

GetUpdateCustomFieldsOk returns a tuple with the UpdateCustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateCustomFields

`func (o *InputApplyCaseTemplateWithIds) SetUpdateCustomFields(v bool)`

SetUpdateCustomFields sets UpdateCustomFields field to given value.

### HasUpdateCustomFields

`func (o *InputApplyCaseTemplateWithIds) HasUpdateCustomFields() bool`

HasUpdateCustomFields returns a boolean if a field has been set.

### GetImportTasks

`func (o *InputApplyCaseTemplateWithIds) GetImportTasks() []string`

GetImportTasks returns the ImportTasks field if non-nil, zero value otherwise.

### GetImportTasksOk

`func (o *InputApplyCaseTemplateWithIds) GetImportTasksOk() (*[]string, bool)`

GetImportTasksOk returns a tuple with the ImportTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportTasks

`func (o *InputApplyCaseTemplateWithIds) SetImportTasks(v []string)`

SetImportTasks sets ImportTasks field to given value.

### HasImportTasks

`func (o *InputApplyCaseTemplateWithIds) HasImportTasks() bool`

HasImportTasks returns a boolean if a field has been set.

### GetImportPages

`func (o *InputApplyCaseTemplateWithIds) GetImportPages() []string`

GetImportPages returns the ImportPages field if non-nil, zero value otherwise.

### GetImportPagesOk

`func (o *InputApplyCaseTemplateWithIds) GetImportPagesOk() (*[]string, bool)`

GetImportPagesOk returns a tuple with the ImportPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportPages

`func (o *InputApplyCaseTemplateWithIds) SetImportPages(v []string)`

SetImportPages sets ImportPages field to given value.

### HasImportPages

`func (o *InputApplyCaseTemplateWithIds) HasImportPages() bool`

HasImportPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


