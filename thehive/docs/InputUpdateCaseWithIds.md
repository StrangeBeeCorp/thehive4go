# InputUpdateCaseWithIds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** |  | 
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

### NewInputUpdateCaseWithIds

`func NewInputUpdateCaseWithIds(ids []string, ) *InputUpdateCaseWithIds`

NewInputUpdateCaseWithIds instantiates a new InputUpdateCaseWithIds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputUpdateCaseWithIdsWithDefaults

`func NewInputUpdateCaseWithIdsWithDefaults() *InputUpdateCaseWithIds`

NewInputUpdateCaseWithIdsWithDefaults instantiates a new InputUpdateCaseWithIds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *InputUpdateCaseWithIds) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *InputUpdateCaseWithIds) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *InputUpdateCaseWithIds) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetTitle

`func (o *InputUpdateCaseWithIds) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InputUpdateCaseWithIds) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InputUpdateCaseWithIds) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InputUpdateCaseWithIds) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *InputUpdateCaseWithIds) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputUpdateCaseWithIds) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputUpdateCaseWithIds) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InputUpdateCaseWithIds) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSeverity

`func (o *InputUpdateCaseWithIds) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *InputUpdateCaseWithIds) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *InputUpdateCaseWithIds) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *InputUpdateCaseWithIds) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetStartDate

`func (o *InputUpdateCaseWithIds) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InputUpdateCaseWithIds) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InputUpdateCaseWithIds) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InputUpdateCaseWithIds) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *InputUpdateCaseWithIds) GetEndDate() int64`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *InputUpdateCaseWithIds) GetEndDateOk() (*int64, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *InputUpdateCaseWithIds) SetEndDate(v int64)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *InputUpdateCaseWithIds) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetTags

`func (o *InputUpdateCaseWithIds) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InputUpdateCaseWithIds) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InputUpdateCaseWithIds) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InputUpdateCaseWithIds) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetFlag

`func (o *InputUpdateCaseWithIds) GetFlag() bool`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *InputUpdateCaseWithIds) GetFlagOk() (*bool, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *InputUpdateCaseWithIds) SetFlag(v bool)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *InputUpdateCaseWithIds) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetTlp

`func (o *InputUpdateCaseWithIds) GetTlp() int32`

GetTlp returns the Tlp field if non-nil, zero value otherwise.

### GetTlpOk

`func (o *InputUpdateCaseWithIds) GetTlpOk() (*int32, bool)`

GetTlpOk returns a tuple with the Tlp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlp

`func (o *InputUpdateCaseWithIds) SetTlp(v int32)`

SetTlp sets Tlp field to given value.

### HasTlp

`func (o *InputUpdateCaseWithIds) HasTlp() bool`

HasTlp returns a boolean if a field has been set.

### GetPap

`func (o *InputUpdateCaseWithIds) GetPap() int32`

GetPap returns the Pap field if non-nil, zero value otherwise.

### GetPapOk

`func (o *InputUpdateCaseWithIds) GetPapOk() (*int32, bool)`

GetPapOk returns a tuple with the Pap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPap

`func (o *InputUpdateCaseWithIds) SetPap(v int32)`

SetPap sets Pap field to given value.

### HasPap

`func (o *InputUpdateCaseWithIds) HasPap() bool`

HasPap returns a boolean if a field has been set.

### GetStatus

`func (o *InputUpdateCaseWithIds) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InputUpdateCaseWithIds) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InputUpdateCaseWithIds) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InputUpdateCaseWithIds) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSummary

`func (o *InputUpdateCaseWithIds) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *InputUpdateCaseWithIds) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *InputUpdateCaseWithIds) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *InputUpdateCaseWithIds) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetAssignee

`func (o *InputUpdateCaseWithIds) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *InputUpdateCaseWithIds) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *InputUpdateCaseWithIds) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *InputUpdateCaseWithIds) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetImpactStatus

`func (o *InputUpdateCaseWithIds) GetImpactStatus() string`

GetImpactStatus returns the ImpactStatus field if non-nil, zero value otherwise.

### GetImpactStatusOk

`func (o *InputUpdateCaseWithIds) GetImpactStatusOk() (*string, bool)`

GetImpactStatusOk returns a tuple with the ImpactStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactStatus

`func (o *InputUpdateCaseWithIds) SetImpactStatus(v string)`

SetImpactStatus sets ImpactStatus field to given value.

### HasImpactStatus

`func (o *InputUpdateCaseWithIds) HasImpactStatus() bool`

HasImpactStatus returns a boolean if a field has been set.

### GetCustomFields

`func (o *InputUpdateCaseWithIds) GetCustomFields() InputUpdateCaseCustomFields`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *InputUpdateCaseWithIds) GetCustomFieldsOk() (*InputUpdateCaseCustomFields, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *InputUpdateCaseWithIds) SetCustomFields(v InputUpdateCaseCustomFields)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *InputUpdateCaseWithIds) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetTaskRule

`func (o *InputUpdateCaseWithIds) GetTaskRule() string`

GetTaskRule returns the TaskRule field if non-nil, zero value otherwise.

### GetTaskRuleOk

`func (o *InputUpdateCaseWithIds) GetTaskRuleOk() (*string, bool)`

GetTaskRuleOk returns a tuple with the TaskRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRule

`func (o *InputUpdateCaseWithIds) SetTaskRule(v string)`

SetTaskRule sets TaskRule field to given value.

### HasTaskRule

`func (o *InputUpdateCaseWithIds) HasTaskRule() bool`

HasTaskRule returns a boolean if a field has been set.

### GetObservableRule

`func (o *InputUpdateCaseWithIds) GetObservableRule() string`

GetObservableRule returns the ObservableRule field if non-nil, zero value otherwise.

### GetObservableRuleOk

`func (o *InputUpdateCaseWithIds) GetObservableRuleOk() (*string, bool)`

GetObservableRuleOk returns a tuple with the ObservableRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservableRule

`func (o *InputUpdateCaseWithIds) SetObservableRule(v string)`

SetObservableRule sets ObservableRule field to given value.

### HasObservableRule

`func (o *InputUpdateCaseWithIds) HasObservableRule() bool`

HasObservableRule returns a boolean if a field has been set.

### GetAddTags

`func (o *InputUpdateCaseWithIds) GetAddTags() []string`

GetAddTags returns the AddTags field if non-nil, zero value otherwise.

### GetAddTagsOk

`func (o *InputUpdateCaseWithIds) GetAddTagsOk() (*[]string, bool)`

GetAddTagsOk returns a tuple with the AddTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTags

`func (o *InputUpdateCaseWithIds) SetAddTags(v []string)`

SetAddTags sets AddTags field to given value.

### HasAddTags

`func (o *InputUpdateCaseWithIds) HasAddTags() bool`

HasAddTags returns a boolean if a field has been set.

### GetRemoveTags

`func (o *InputUpdateCaseWithIds) GetRemoveTags() []string`

GetRemoveTags returns the RemoveTags field if non-nil, zero value otherwise.

### GetRemoveTagsOk

`func (o *InputUpdateCaseWithIds) GetRemoveTagsOk() (*[]string, bool)`

GetRemoveTagsOk returns a tuple with the RemoveTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveTags

`func (o *InputUpdateCaseWithIds) SetRemoveTags(v []string)`

SetRemoveTags sets RemoveTags field to given value.

### HasRemoveTags

`func (o *InputUpdateCaseWithIds) HasRemoveTags() bool`

HasRemoveTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


