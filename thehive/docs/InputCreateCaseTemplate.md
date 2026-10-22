# InputCreateCaseTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the case template, must be unique in your organization. |
**DisplayName** | Pointer to **string** | Name shown to users when selecting this template. Defaults to the template name if not specified. | [optional]
**TitlePrefix** | Pointer to **string** | Prefix automatically added to case titles when the template is applied. Add a hyphen after the prefix to separate it from the case title. | [optional]
**Description** | Pointer to **string** | Default description for cases. Supports [TheHive-flavored Markdown](https://docs.strangebee.com/thehive/user-guides/thehive-flavored-markdown/). | [optional]
**Severity** | Pointer to **int32** | Default severity level for cases: 1 (Low), 2 (Medium), 3 (High), 4 (Critical). | [optional]
**Tags** | Pointer to **[]string** | Default tags for cases. Unknown tags are automatically created. Use &#x60;POST /api/v1/query&#x60; with &#x60;listTag&#x60; to retrieve existing free-form tags, or with &#x60;listTaxonomy&#x60; and &#x60;tags&#x60; to retrieve taxonomy tags. | [optional]
**Flag** | Pointer to **bool** | Marks cases as important by default. | [optional]
**Tlp** | Pointer to **int32** | Default [Traffic Light Protocol (TLP)](https://www.misp-project.org/taxonomies.html#_tlp) level for cases controlling how case information can be shared: 0 (Clear — no restrictions), 1 (Green — community), 2 (Amber — organization and clients), 3 (Amber+Strict — organization only), 4 (Red — named recipients only). | [optional]
**Pap** | Pointer to **int32** | Default [Permissible Actions Protocol (PAP)](https://www.misp-project.org/taxonomies.html#_pap) level for cases controlling what actions can be taken with case information: 0 (Clear — no restrictions), 1 (Green — active actions allowed), 2 (Amber — passive checks only), 3 (Red — non-detectable actions only). | [optional]
**Summary** | Pointer to **string** | Default summary for cases. Supports [TheHive-flavored Markdown](https://docs.strangebee.com/thehive/user-guides/thehive-flavored-markdown/). | [optional]
**Tasks** | Pointer to [**[]InputCreateTask**](InputCreateTask.md) | Task templates to include in this case template. A task represents one or more actions, called task logs, that must be completed during the investigation. These tasks are automatically created on cases when the template is applied. | [optional]
**PageTemplateIds** | Pointer to **[]string** | Use &#x60;pageTemplates&#x60; instead. | [optional]
**CustomFields** | Pointer to [**InputCreateCaseTemplateCustomFields**](InputCreateCaseTemplateCustomFields.md) |  | [optional]
**PageTemplates** | Pointer to [**[]InputCreatePage**](InputCreatePage.md) | Pages enable the creation and sharing of valuable information within a team on a case. Pages to include in this case template are automatically created on cases when the template is applied. Preferred over the deprecated &#x60;pageTemplateIds&#x60; field. | [optional]

## Methods

### NewInputCreateCaseTemplate

`func NewInputCreateCaseTemplate(name string, ) *InputCreateCaseTemplate`

NewInputCreateCaseTemplate instantiates a new InputCreateCaseTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputCreateCaseTemplateWithDefaults

`func NewInputCreateCaseTemplateWithDefaults() *InputCreateCaseTemplate`

NewInputCreateCaseTemplateWithDefaults instantiates a new InputCreateCaseTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InputCreateCaseTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputCreateCaseTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputCreateCaseTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *InputCreateCaseTemplate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InputCreateCaseTemplate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InputCreateCaseTemplate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InputCreateCaseTemplate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetTitlePrefix

`func (o *InputCreateCaseTemplate) GetTitlePrefix() string`

GetTitlePrefix returns the TitlePrefix field if non-nil, zero value otherwise.

### GetTitlePrefixOk

`func (o *InputCreateCaseTemplate) GetTitlePrefixOk() (*string, bool)`

GetTitlePrefixOk returns a tuple with the TitlePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitlePrefix

`func (o *InputCreateCaseTemplate) SetTitlePrefix(v string)`

SetTitlePrefix sets TitlePrefix field to given value.

### HasTitlePrefix

`func (o *InputCreateCaseTemplate) HasTitlePrefix() bool`

HasTitlePrefix returns a boolean if a field has been set.

### GetDescription

`func (o *InputCreateCaseTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputCreateCaseTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputCreateCaseTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InputCreateCaseTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSeverity

`func (o *InputCreateCaseTemplate) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *InputCreateCaseTemplate) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *InputCreateCaseTemplate) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *InputCreateCaseTemplate) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetTags

`func (o *InputCreateCaseTemplate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InputCreateCaseTemplate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InputCreateCaseTemplate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InputCreateCaseTemplate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetFlag

`func (o *InputCreateCaseTemplate) GetFlag() bool`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *InputCreateCaseTemplate) GetFlagOk() (*bool, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *InputCreateCaseTemplate) SetFlag(v bool)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *InputCreateCaseTemplate) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetTlp

`func (o *InputCreateCaseTemplate) GetTlp() int32`

GetTlp returns the Tlp field if non-nil, zero value otherwise.

### GetTlpOk

`func (o *InputCreateCaseTemplate) GetTlpOk() (*int32, bool)`

GetTlpOk returns a tuple with the Tlp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlp

`func (o *InputCreateCaseTemplate) SetTlp(v int32)`

SetTlp sets Tlp field to given value.

### HasTlp

`func (o *InputCreateCaseTemplate) HasTlp() bool`

HasTlp returns a boolean if a field has been set.

### GetPap

`func (o *InputCreateCaseTemplate) GetPap() int32`

GetPap returns the Pap field if non-nil, zero value otherwise.

### GetPapOk

`func (o *InputCreateCaseTemplate) GetPapOk() (*int32, bool)`

GetPapOk returns a tuple with the Pap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPap

`func (o *InputCreateCaseTemplate) SetPap(v int32)`

SetPap sets Pap field to given value.

### HasPap

`func (o *InputCreateCaseTemplate) HasPap() bool`

HasPap returns a boolean if a field has been set.

### GetSummary

`func (o *InputCreateCaseTemplate) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *InputCreateCaseTemplate) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *InputCreateCaseTemplate) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *InputCreateCaseTemplate) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTasks

`func (o *InputCreateCaseTemplate) GetTasks() []InputCreateTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *InputCreateCaseTemplate) GetTasksOk() (*[]InputCreateTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *InputCreateCaseTemplate) SetTasks(v []InputCreateTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *InputCreateCaseTemplate) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetPageTemplateIds

`func (o *InputCreateCaseTemplate) GetPageTemplateIds() []string`

GetPageTemplateIds returns the PageTemplateIds field if non-nil, zero value otherwise.

### GetPageTemplateIdsOk

`func (o *InputCreateCaseTemplate) GetPageTemplateIdsOk() (*[]string, bool)`

GetPageTemplateIdsOk returns a tuple with the PageTemplateIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageTemplateIds

`func (o *InputCreateCaseTemplate) SetPageTemplateIds(v []string)`

SetPageTemplateIds sets PageTemplateIds field to given value.

### HasPageTemplateIds

`func (o *InputCreateCaseTemplate) HasPageTemplateIds() bool`

HasPageTemplateIds returns a boolean if a field has been set.

### GetCustomFields

`func (o *InputCreateCaseTemplate) GetCustomFields() InputCreateCaseTemplateCustomFields`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *InputCreateCaseTemplate) GetCustomFieldsOk() (*InputCreateCaseTemplateCustomFields, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *InputCreateCaseTemplate) SetCustomFields(v InputCreateCaseTemplateCustomFields)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *InputCreateCaseTemplate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetPageTemplates

`func (o *InputCreateCaseTemplate) GetPageTemplates() []InputCreatePage`

GetPageTemplates returns the PageTemplates field if non-nil, zero value otherwise.

### GetPageTemplatesOk

`func (o *InputCreateCaseTemplate) GetPageTemplatesOk() (*[]InputCreatePage, bool)`

GetPageTemplatesOk returns a tuple with the PageTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageTemplates

`func (o *InputCreateCaseTemplate) SetPageTemplates(v []InputCreatePage)`

SetPageTemplates sets PageTemplates field to given value.

### HasPageTemplates

`func (o *InputCreateCaseTemplate) HasPageTemplates() bool`

HasPageTemplates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
