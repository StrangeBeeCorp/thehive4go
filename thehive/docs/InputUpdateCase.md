# InputUpdateCase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **int32** |  | [optional] 
**StartDate** | Pointer to **int64** |  | [optional] 
**EndDate** | Pointer to **int64** |  | [optional] 
**Tags** | Pointer to **[]string** | Set the case tags to this array | [optional] 
**Flag** | Pointer to **bool** |  | [optional] 
**Tlp** | Pointer to **int32** |  | [optional] 
**Pap** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Assignee** | Pointer to **string** |  | [optional] 
**ImpactStatus** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to [**InputUpdateCaseCustomFields**](InputUpdateCaseCustomFields.md) |  | [optional] 
**TaskRule** | Pointer to **string** |  | [optional] 
**ObservableRule** | Pointer to **string** |  | [optional] 
**AddTags** | Pointer to **[]string** | Those tags will be added to the current case | [optional] 
**RemoveTags** | Pointer to **[]string** | Those tags will be removed from the current case | [optional] 

## Methods

### NewInputUpdateCase

`func NewInputUpdateCase() *InputUpdateCase`

NewInputUpdateCase instantiates a new InputUpdateCase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputUpdateCaseWithDefaults

`func NewInputUpdateCaseWithDefaults() *InputUpdateCase`

NewInputUpdateCaseWithDefaults instantiates a new InputUpdateCase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *InputUpdateCase) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InputUpdateCase) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InputUpdateCase) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InputUpdateCase) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *InputUpdateCase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputUpdateCase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputUpdateCase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InputUpdateCase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSeverity

`func (o *InputUpdateCase) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *InputUpdateCase) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *InputUpdateCase) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *InputUpdateCase) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetStartDate

`func (o *InputUpdateCase) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InputUpdateCase) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InputUpdateCase) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InputUpdateCase) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *InputUpdateCase) GetEndDate() int64`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *InputUpdateCase) GetEndDateOk() (*int64, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *InputUpdateCase) SetEndDate(v int64)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *InputUpdateCase) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetTags

`func (o *InputUpdateCase) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InputUpdateCase) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InputUpdateCase) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InputUpdateCase) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetFlag

`func (o *InputUpdateCase) GetFlag() bool`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *InputUpdateCase) GetFlagOk() (*bool, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *InputUpdateCase) SetFlag(v bool)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *InputUpdateCase) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetTlp

`func (o *InputUpdateCase) GetTlp() int32`

GetTlp returns the Tlp field if non-nil, zero value otherwise.

### GetTlpOk

`func (o *InputUpdateCase) GetTlpOk() (*int32, bool)`

GetTlpOk returns a tuple with the Tlp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlp

`func (o *InputUpdateCase) SetTlp(v int32)`

SetTlp sets Tlp field to given value.

### HasTlp

`func (o *InputUpdateCase) HasTlp() bool`

HasTlp returns a boolean if a field has been set.

### GetPap

`func (o *InputUpdateCase) GetPap() int32`

GetPap returns the Pap field if non-nil, zero value otherwise.

### GetPapOk

`func (o *InputUpdateCase) GetPapOk() (*int32, bool)`

GetPapOk returns a tuple with the Pap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPap

`func (o *InputUpdateCase) SetPap(v int32)`

SetPap sets Pap field to given value.

### HasPap

`func (o *InputUpdateCase) HasPap() bool`

HasPap returns a boolean if a field has been set.

### GetStatus

`func (o *InputUpdateCase) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InputUpdateCase) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InputUpdateCase) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InputUpdateCase) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSummary

`func (o *InputUpdateCase) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *InputUpdateCase) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *InputUpdateCase) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *InputUpdateCase) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetAssignee

`func (o *InputUpdateCase) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *InputUpdateCase) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *InputUpdateCase) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *InputUpdateCase) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetImpactStatus

`func (o *InputUpdateCase) GetImpactStatus() string`

GetImpactStatus returns the ImpactStatus field if non-nil, zero value otherwise.

### GetImpactStatusOk

`func (o *InputUpdateCase) GetImpactStatusOk() (*string, bool)`

GetImpactStatusOk returns a tuple with the ImpactStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactStatus

`func (o *InputUpdateCase) SetImpactStatus(v string)`

SetImpactStatus sets ImpactStatus field to given value.

### HasImpactStatus

`func (o *InputUpdateCase) HasImpactStatus() bool`

HasImpactStatus returns a boolean if a field has been set.

### GetCustomFields

`func (o *InputUpdateCase) GetCustomFields() InputUpdateCaseCustomFields`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *InputUpdateCase) GetCustomFieldsOk() (*InputUpdateCaseCustomFields, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *InputUpdateCase) SetCustomFields(v InputUpdateCaseCustomFields)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *InputUpdateCase) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetTaskRule

`func (o *InputUpdateCase) GetTaskRule() string`

GetTaskRule returns the TaskRule field if non-nil, zero value otherwise.

### GetTaskRuleOk

`func (o *InputUpdateCase) GetTaskRuleOk() (*string, bool)`

GetTaskRuleOk returns a tuple with the TaskRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRule

`func (o *InputUpdateCase) SetTaskRule(v string)`

SetTaskRule sets TaskRule field to given value.

### HasTaskRule

`func (o *InputUpdateCase) HasTaskRule() bool`

HasTaskRule returns a boolean if a field has been set.

### GetObservableRule

`func (o *InputUpdateCase) GetObservableRule() string`

GetObservableRule returns the ObservableRule field if non-nil, zero value otherwise.

### GetObservableRuleOk

`func (o *InputUpdateCase) GetObservableRuleOk() (*string, bool)`

GetObservableRuleOk returns a tuple with the ObservableRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservableRule

`func (o *InputUpdateCase) SetObservableRule(v string)`

SetObservableRule sets ObservableRule field to given value.

### HasObservableRule

`func (o *InputUpdateCase) HasObservableRule() bool`

HasObservableRule returns a boolean if a field has been set.

### GetAddTags

`func (o *InputUpdateCase) GetAddTags() []string`

GetAddTags returns the AddTags field if non-nil, zero value otherwise.

### GetAddTagsOk

`func (o *InputUpdateCase) GetAddTagsOk() (*[]string, bool)`

GetAddTagsOk returns a tuple with the AddTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTags

`func (o *InputUpdateCase) SetAddTags(v []string)`

SetAddTags sets AddTags field to given value.

### HasAddTags

`func (o *InputUpdateCase) HasAddTags() bool`

HasAddTags returns a boolean if a field has been set.

### GetRemoveTags

`func (o *InputUpdateCase) GetRemoveTags() []string`

GetRemoveTags returns the RemoveTags field if non-nil, zero value otherwise.

### GetRemoveTagsOk

`func (o *InputUpdateCase) GetRemoveTagsOk() (*[]string, bool)`

GetRemoveTagsOk returns a tuple with the RemoveTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveTags

`func (o *InputUpdateCase) SetRemoveTags(v []string)`

SetRemoveTags sets RemoveTags field to given value.

### HasRemoveTags

`func (o *InputUpdateCase) HasRemoveTags() bool`

HasRemoveTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


