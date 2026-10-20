# InputUpdateCaseTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | New name for the case template, must be unique in your organization. | [optional]
**DisplayName** | Pointer to **string** | New display name shown to users when selecting this template. | [optional]
**TitlePrefix** | Pointer to **string** | New prefix for case titles. Send &#x60;null&#x60; to remove the prefix. | [optional]
**Description** | Pointer to **string** | New default description for cases. Supports [TheHive-flavored Markdown](https://docs.strangebee.com/thehive/user-guides/thehive-flavored-markdown/). Send &#x60;null&#x60; to clear this field. | [optional]
**Severity** | Pointer to **int32** | New default severity level for cases: 1 (Low), 2 (Medium), 3 (High), 4 (Critical). Send &#x60;null&#x60; to clear this field. | [optional]
**Tags** | Pointer to **[]string** | Replaces the template&#39;s current tag list with this set. Unknown tags are automatically created. Use &#x60;POST /api/v1/query&#x60; with &#x60;listTag&#x60; to retrieve existing free-form tags, or with &#x60;listTaxonomy&#x60; and &#x60;tags&#x60; to retrieve taxonomy tags. | [optional]
**Flag** | Pointer to **bool** | Marks cases as important by default. | [optional]
**Tlp** | Pointer to **int32** | New default [Traffic Light Protocol (TLP)](https://www.misp-project.org/taxonomies.html#_tlp) level for cases: 0 (Clear), 1 (Green), 2 (Amber), 3 (Amber+Strict), 4 (Red). Send &#x60;null&#x60; to clear this field. | [optional]
**Pap** | Pointer to **int32** | New default [Permissible Actions Protocol (PAP)](https://www.misp-project.org/taxonomies.html#_pap) level for cases: 0 (Clear), 1 (Green), 2 (Amber), 3 (Red). Send &#x60;null&#x60; to clear this field. | [optional]
**Summary** | Pointer to **string** | New default summary for cases. Supports [TheHive-flavored Markdown](https://docs.strangebee.com/thehive/user-guides/thehive-flavored-markdown/). Send &#x60;null&#x60; to clear this field. | [optional]
**CustomFields** | Pointer to [**InputUpdateCaseTemplateCustomFields**](InputUpdateCaseTemplateCustomFields.md) |  | [optional]
**Tasks** | Pointer to [**[]InputCreateTask**](InputCreateTask.md) | Replaces the template&#39;s current task list with this set. A task represents one or more actions, called task logs, that must be completed during the investigation. | [optional]

## Methods

### NewInputUpdateCaseTemplate

`func NewInputUpdateCaseTemplate() *InputUpdateCaseTemplate`

NewInputUpdateCaseTemplate instantiates a new InputUpdateCaseTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputUpdateCaseTemplateWithDefaults

`func NewInputUpdateCaseTemplateWithDefaults() *InputUpdateCaseTemplate`

NewInputUpdateCaseTemplateWithDefaults instantiates a new InputUpdateCaseTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InputUpdateCaseTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputUpdateCaseTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputUpdateCaseTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InputUpdateCaseTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *InputUpdateCaseTemplate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InputUpdateCaseTemplate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InputUpdateCaseTemplate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InputUpdateCaseTemplate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetTitlePrefix

`func (o *InputUpdateCaseTemplate) GetTitlePrefix() string`

GetTitlePrefix returns the TitlePrefix field if non-nil, zero value otherwise.

### GetTitlePrefixOk

`func (o *InputUpdateCaseTemplate) GetTitlePrefixOk() (*string, bool)`

GetTitlePrefixOk returns a tuple with the TitlePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitlePrefix

`func (o *InputUpdateCaseTemplate) SetTitlePrefix(v string)`

SetTitlePrefix sets TitlePrefix field to given value.

### HasTitlePrefix

`func (o *InputUpdateCaseTemplate) HasTitlePrefix() bool`

HasTitlePrefix returns a boolean if a field has been set.

### GetDescription

`func (o *InputUpdateCaseTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputUpdateCaseTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputUpdateCaseTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InputUpdateCaseTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSeverity

`func (o *InputUpdateCaseTemplate) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *InputUpdateCaseTemplate) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *InputUpdateCaseTemplate) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *InputUpdateCaseTemplate) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetTags

`func (o *InputUpdateCaseTemplate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InputUpdateCaseTemplate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InputUpdateCaseTemplate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InputUpdateCaseTemplate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetFlag

`func (o *InputUpdateCaseTemplate) GetFlag() bool`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *InputUpdateCaseTemplate) GetFlagOk() (*bool, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *InputUpdateCaseTemplate) SetFlag(v bool)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *InputUpdateCaseTemplate) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetTlp

`func (o *InputUpdateCaseTemplate) GetTlp() int32`

GetTlp returns the Tlp field if non-nil, zero value otherwise.

### GetTlpOk

`func (o *InputUpdateCaseTemplate) GetTlpOk() (*int32, bool)`

GetTlpOk returns a tuple with the Tlp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlp

`func (o *InputUpdateCaseTemplate) SetTlp(v int32)`

SetTlp sets Tlp field to given value.

### HasTlp

`func (o *InputUpdateCaseTemplate) HasTlp() bool`

HasTlp returns a boolean if a field has been set.

### GetPap

`func (o *InputUpdateCaseTemplate) GetPap() int32`

GetPap returns the Pap field if non-nil, zero value otherwise.

### GetPapOk

`func (o *InputUpdateCaseTemplate) GetPapOk() (*int32, bool)`

GetPapOk returns a tuple with the Pap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPap

`func (o *InputUpdateCaseTemplate) SetPap(v int32)`

SetPap sets Pap field to given value.

### HasPap

`func (o *InputUpdateCaseTemplate) HasPap() bool`

HasPap returns a boolean if a field has been set.

### GetSummary

`func (o *InputUpdateCaseTemplate) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *InputUpdateCaseTemplate) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *InputUpdateCaseTemplate) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *InputUpdateCaseTemplate) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetCustomFields

`func (o *InputUpdateCaseTemplate) GetCustomFields() InputUpdateCaseTemplateCustomFields`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *InputUpdateCaseTemplate) GetCustomFieldsOk() (*InputUpdateCaseTemplateCustomFields, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *InputUpdateCaseTemplate) SetCustomFields(v InputUpdateCaseTemplateCustomFields)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *InputUpdateCaseTemplate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetTasks

`func (o *InputUpdateCaseTemplate) GetTasks() []InputCreateTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *InputUpdateCaseTemplate) GetTasksOk() (*[]InputCreateTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *InputUpdateCaseTemplate) SetTasks(v []InputCreateTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *InputUpdateCaseTemplate) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
