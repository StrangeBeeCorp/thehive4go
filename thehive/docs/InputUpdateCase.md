# InputUpdateCase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | New title of the case. | [optional]
**Description** | Pointer to **string** | New details of the case. Supports [TheHive-flavored Markdown](https://docs.strangebee.com/thehive/user-guides/thehive-flavored-markdown/). | [optional]
**Severity** | Pointer to **int32** | New severity level of the case: 1 (Low), 2 (Medium), 3 (High), 4 (Critical). | [optional]
**StartDate** | Pointer to **int64** | New start date of the incident. Expected format: Unix timestamp in milliseconds. | [optional]
**EndDate** | Pointer to **int64** | New end date of the incident. Expected format: Unix timestamp in milliseconds. Send &#x60;null&#x60; to unset. | [optional]
**Tags** | Pointer to **[]string** | Replaces all current tags with this set. Unknown tags are automatically created. Use &#x60;addTags&#x60; and &#x60;removeTags&#x60; to add or remove individual tags without replacing the full set. Use &#x60;POST /api/v1/query&#x60; with &#x60;listTag&#x60; to retrieve existing custom tags, or with &#x60;listTaxonomy&#x60; and &#x60;tags&#x60; to retrieve taxonomy tags. | [optional]
**Flag** | Pointer to **bool** | Marks the case as important for quick filtering. | [optional]
**Tlp** | Pointer to **int32** | New [Traffic Light Protocol (TLP)](https://www.misp-project.org/taxonomies.html#_tlp) level: 0 (Clear), 1 (Green), 2 (Amber), 3 (Amber+Strict), 4 (Red). | [optional]
**Pap** | Pointer to **int32** | New [Permissible Actions Protocol (PAP)](https://www.misp-project.org/taxonomies.html#_pap) level: 0 (Clear), 1 (Green), 2 (Amber), 3 (Red). | [optional]
**Status** | Pointer to **string** | New case status. Must match an existing status. To close a case, set this to a status linked to the *Closed* stage. Use &#x60;POST /api/v1/query&#x60; with &#x60;listCaseStatus&#x60; to list available statuses. | [optional]
**Summary** | Pointer to **string** | New investigation summary. Send &#x60;null&#x60; to unset. | [optional]
**Assignee** | Pointer to **string** | Login of the user to assign the case to. Typically an email address. If a default domain is configured in the [authentication settings](https://docs.strangebee.com/thehive/administration/authentication/configure-authentication), a plain username is also accepted. Use &#x60;POST /api/v1/query&#x60; with &#x60;listUser&#x60; to retrieve available users and their logins. Send &#x60;null&#x60; to unassign. | [optional]
**ImpactStatus** | Pointer to **string** | Enum: &#x60;NoImpact&#x60; &#x60;WithImpact&#x60; &#x60;NotApplicable&#x60;&lt;br&gt;Verdict on whether the incident had real consequences for the organization. Required when you close a case with a status whose &#x60;requiresImpact&#x60; field is &#x60;true&#x60;. TheHive rejects the update if you send it while closing a case with a status that doesn&#39;t require an impact. Use &#x60;POST /api/v1/query&#x60; with &#x60;listCaseStatus&#x60; to check which statuses require an impact. &#x60;WithImpact&#x60; means the incident caused damage or data loss. &#x60;NoImpact&#x60; means it was resolved without consequences. &#x60;NotApplicable&#x60; means the concept doesn&#39;t apply to this case. Send &#x60;null&#x60; to unset. | [optional]
**CustomFields** | Pointer to [**InputUpdateCaseCustomFields**](InputUpdateCaseCustomFields.md) |  | [optional]
**TaskRule** | Pointer to **string** | Enum: &#x60;autoShare&#x60; &#x60;manual&#x60;&lt;br&gt;New task-sharing rule for the owner organization: &#x60;autoShare&#x60; shares new tasks with every sharing organization automatically, &#x60;manual&#x60; requires explicit sharing per task. Overrides the [organization-level sharing rules](https://docs.strangebee.com/thehive/administration/organizations/about-organizations-sharing-rules/). | [optional]
**ObservableRule** | Pointer to **string** | Enum: &#x60;autoShare&#x60; &#x60;manual&#x60;&lt;br&gt;New observable-sharing rule for the owner organization: &#x60;autoShare&#x60; shares new observables with every sharing organization automatically, &#x60;manual&#x60; requires explicit sharing per observable. Overrides the [organization-level sharing rules](https://docs.strangebee.com/thehive/administration/organizations/about-organizations-sharing-rules/). | [optional]
**AddTags** | Pointer to **[]string** | Tags to add to the current set of case tags. Unknown tags are automatically created. Use &#x60;POST /api/v1/query&#x60; with &#x60;listTag&#x60; to retrieve existing custom tags, or with &#x60;listTaxonomy&#x60; and &#x60;tags&#x60; to retrieve taxonomy tags. | [optional]
**RemoveTags** | Pointer to **[]string** | Tags to remove from the current set of case tags. Tags not present on the case are silently ignored. Use &#x60;GET /api/v1/case/{idOrName}&#x60; and read the &#x60;tags&#x60; field to see the current tags on the case. | [optional]

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
