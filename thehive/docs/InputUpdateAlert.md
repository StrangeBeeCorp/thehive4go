# InputUpdateAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Type of the alert. Identifies the event category and groups related alerts. | [optional]
**Source** | Pointer to **string** | Source system that generated the alert. | [optional]
**SourceRef** | Pointer to **string** | Unique reference identifier from the source system. Combined with &#x60;type&#x60; and &#x60;source&#x60;, uniquely identifies an alert. | [optional]
**ExternalLink** | Pointer to **string** | External URL linking to the alert in the source system. Send &#x60;null&#x60; to unset. | [optional]
**Title** | Pointer to **string** | New title of the alert. | [optional]
**Description** | Pointer to **string** | New description of the alert. Supports [TheHive-flavored Markdown](https://docs.strangebee.com/thehive/user-guides/thehive-flavored-markdown/). | [optional]
**Severity** | Pointer to **int32** | New severity level of the alert: 1 (Low), 2 (Medium), 3 (High), 4 (Critical). | [optional]
**Date** | Pointer to **int64** | New date and time of the event that triggered the alert, as a Unix timestamp in milliseconds. | [optional]
**LastSyncDate** | Pointer to **int64** | Date and time of the last synchronization with the source system, as a Unix timestamp in milliseconds. | [optional]
**Tags** | Pointer to **[]string** | Replaces all current tags with this set. Unknown tags are automatically created. Use &#x60;addTags&#x60; and &#x60;removeTags&#x60; to add or remove individual tags without replacing the full set. Use &#x60;POST /api/v1/query&#x60; with &#x60;listTag&#x60; to retrieve existing custom tags, or with &#x60;listTaxonomy&#x60; and &#x60;tags&#x60; to retrieve taxonomy tags. | [optional]
**Tlp** | Pointer to **int32** | New [Traffic Light Protocol (TLP)](https://www.misp-project.org/taxonomies.html#_tlp) level: 0 (Clear), 1 (Green), 2 (Amber), 3 (Amber+Strict), 4 (Red). | [optional]
**Pap** | Pointer to **int32** | New [Permissible Actions Protocol (PAP)](https://www.misp-project.org/taxonomies.html#_pap) level: 0 (Clear), 1 (Green), 2 (Amber), 3 (Red). | [optional]
**Follow** | Pointer to **bool** | Whether to follow the alert for change notifications. | [optional]
**CustomFields** | Pointer to [**InputUpdateAlertCustomFields**](InputUpdateAlertCustomFields.md) |  | [optional]
**Status** | Pointer to **string** | New alert status. Must match an existing status. To close an alert, set this to a status linked to the *Closed* stage. Use &#x60;POST /api/v1/query&#x60; with &#x60;listAlertStatus&#x60; to list available statuses. | [optional]
**Summary** | Pointer to **string** | New summary of the triage findings. Send &#x60;null&#x60; to unset. | [optional]
**Assignee** | Pointer to **string** | Login of the user to assign the alert to. Typically an email address. If a default domain is configured in the [authentication settings](https://docs.strangebee.com/thehive/administration/authentication/configure-authentication), a plain username is also accepted. Use &#x60;POST /api/v1/query&#x60; with &#x60;listUser&#x60; to retrieve available users and their logins. Send &#x60;null&#x60; to unassign. | [optional]
**AddTags** | Pointer to **[]string** | Tags to add to the current set of alert tags. Unknown tags are automatically created. Use &#x60;POST /api/v1/query&#x60; with &#x60;listTag&#x60; to retrieve existing custom tags, or with &#x60;listTaxonomy&#x60; and &#x60;tags&#x60; to retrieve taxonomy tags. | [optional]
**RemoveTags** | Pointer to **[]string** | Tags to remove from the current set of alert tags. Tags not present on the alert are silently ignored. Use &#x60;GET /api/v1/alert/{alertId}&#x60; and read the &#x60;tags&#x60; field to see the current tags on the alert. | [optional]

## Methods

### NewInputUpdateAlert

`func NewInputUpdateAlert() *InputUpdateAlert`

NewInputUpdateAlert instantiates a new InputUpdateAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputUpdateAlertWithDefaults

`func NewInputUpdateAlertWithDefaults() *InputUpdateAlert`

NewInputUpdateAlertWithDefaults instantiates a new InputUpdateAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputUpdateAlert) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputUpdateAlert) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputUpdateAlert) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InputUpdateAlert) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSource

`func (o *InputUpdateAlert) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InputUpdateAlert) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InputUpdateAlert) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *InputUpdateAlert) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceRef

`func (o *InputUpdateAlert) GetSourceRef() string`

GetSourceRef returns the SourceRef field if non-nil, zero value otherwise.

### GetSourceRefOk

`func (o *InputUpdateAlert) GetSourceRefOk() (*string, bool)`

GetSourceRefOk returns a tuple with the SourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRef

`func (o *InputUpdateAlert) SetSourceRef(v string)`

SetSourceRef sets SourceRef field to given value.

### HasSourceRef

`func (o *InputUpdateAlert) HasSourceRef() bool`

HasSourceRef returns a boolean if a field has been set.

### GetExternalLink

`func (o *InputUpdateAlert) GetExternalLink() string`

GetExternalLink returns the ExternalLink field if non-nil, zero value otherwise.

### GetExternalLinkOk

`func (o *InputUpdateAlert) GetExternalLinkOk() (*string, bool)`

GetExternalLinkOk returns a tuple with the ExternalLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLink

`func (o *InputUpdateAlert) SetExternalLink(v string)`

SetExternalLink sets ExternalLink field to given value.

### HasExternalLink

`func (o *InputUpdateAlert) HasExternalLink() bool`

HasExternalLink returns a boolean if a field has been set.

### GetTitle

`func (o *InputUpdateAlert) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InputUpdateAlert) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InputUpdateAlert) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InputUpdateAlert) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *InputUpdateAlert) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputUpdateAlert) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputUpdateAlert) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InputUpdateAlert) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSeverity

`func (o *InputUpdateAlert) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *InputUpdateAlert) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *InputUpdateAlert) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *InputUpdateAlert) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetDate

`func (o *InputUpdateAlert) GetDate() int64`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *InputUpdateAlert) GetDateOk() (*int64, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *InputUpdateAlert) SetDate(v int64)`

SetDate sets Date field to given value.

### HasDate

`func (o *InputUpdateAlert) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetLastSyncDate

`func (o *InputUpdateAlert) GetLastSyncDate() int64`

GetLastSyncDate returns the LastSyncDate field if non-nil, zero value otherwise.

### GetLastSyncDateOk

`func (o *InputUpdateAlert) GetLastSyncDateOk() (*int64, bool)`

GetLastSyncDateOk returns a tuple with the LastSyncDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDate

`func (o *InputUpdateAlert) SetLastSyncDate(v int64)`

SetLastSyncDate sets LastSyncDate field to given value.

### HasLastSyncDate

`func (o *InputUpdateAlert) HasLastSyncDate() bool`

HasLastSyncDate returns a boolean if a field has been set.

### GetTags

`func (o *InputUpdateAlert) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InputUpdateAlert) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InputUpdateAlert) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InputUpdateAlert) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTlp

`func (o *InputUpdateAlert) GetTlp() int32`

GetTlp returns the Tlp field if non-nil, zero value otherwise.

### GetTlpOk

`func (o *InputUpdateAlert) GetTlpOk() (*int32, bool)`

GetTlpOk returns a tuple with the Tlp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlp

`func (o *InputUpdateAlert) SetTlp(v int32)`

SetTlp sets Tlp field to given value.

### HasTlp

`func (o *InputUpdateAlert) HasTlp() bool`

HasTlp returns a boolean if a field has been set.

### GetPap

`func (o *InputUpdateAlert) GetPap() int32`

GetPap returns the Pap field if non-nil, zero value otherwise.

### GetPapOk

`func (o *InputUpdateAlert) GetPapOk() (*int32, bool)`

GetPapOk returns a tuple with the Pap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPap

`func (o *InputUpdateAlert) SetPap(v int32)`

SetPap sets Pap field to given value.

### HasPap

`func (o *InputUpdateAlert) HasPap() bool`

HasPap returns a boolean if a field has been set.

### GetFollow

`func (o *InputUpdateAlert) GetFollow() bool`

GetFollow returns the Follow field if non-nil, zero value otherwise.

### GetFollowOk

`func (o *InputUpdateAlert) GetFollowOk() (*bool, bool)`

GetFollowOk returns a tuple with the Follow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollow

`func (o *InputUpdateAlert) SetFollow(v bool)`

SetFollow sets Follow field to given value.

### HasFollow

`func (o *InputUpdateAlert) HasFollow() bool`

HasFollow returns a boolean if a field has been set.

### GetCustomFields

`func (o *InputUpdateAlert) GetCustomFields() InputUpdateAlertCustomFields`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *InputUpdateAlert) GetCustomFieldsOk() (*InputUpdateAlertCustomFields, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *InputUpdateAlert) SetCustomFields(v InputUpdateAlertCustomFields)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *InputUpdateAlert) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetStatus

`func (o *InputUpdateAlert) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InputUpdateAlert) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InputUpdateAlert) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InputUpdateAlert) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSummary

`func (o *InputUpdateAlert) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *InputUpdateAlert) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *InputUpdateAlert) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *InputUpdateAlert) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetAssignee

`func (o *InputUpdateAlert) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *InputUpdateAlert) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *InputUpdateAlert) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *InputUpdateAlert) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetAddTags

`func (o *InputUpdateAlert) GetAddTags() []string`

GetAddTags returns the AddTags field if non-nil, zero value otherwise.

### GetAddTagsOk

`func (o *InputUpdateAlert) GetAddTagsOk() (*[]string, bool)`

GetAddTagsOk returns a tuple with the AddTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTags

`func (o *InputUpdateAlert) SetAddTags(v []string)`

SetAddTags sets AddTags field to given value.

### HasAddTags

`func (o *InputUpdateAlert) HasAddTags() bool`

HasAddTags returns a boolean if a field has been set.

### GetRemoveTags

`func (o *InputUpdateAlert) GetRemoveTags() []string`

GetRemoveTags returns the RemoveTags field if non-nil, zero value otherwise.

### GetRemoveTagsOk

`func (o *InputUpdateAlert) GetRemoveTagsOk() (*[]string, bool)`

GetRemoveTagsOk returns a tuple with the RemoveTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveTags

`func (o *InputUpdateAlert) SetRemoveTags(v []string)`

SetRemoveTags sets RemoveTags field to given value.

### HasRemoveTags

`func (o *InputUpdateAlert) HasRemoveTags() bool`

HasRemoveTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
