# InputUpdateAlertWithIds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** |  | 
**Type** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**SourceRef** | Pointer to **string** |  | [optional] 
**ExternalLink** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **int32** |  | [optional] 
**Date** | Pointer to **int64** |  | [optional] 
**LastSyncDate** | Pointer to **int64** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Tlp** | Pointer to **int32** |  | [optional] 
**Pap** | Pointer to **int32** |  | [optional] 
**Follow** | Pointer to **bool** |  | [optional] 
**CustomFields** | Pointer to [**InputCreateAlertCustomFields**](InputCreateAlertCustomFields.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Assignee** | Pointer to **string** |  | [optional] 
**AddTags** | Pointer to **[]string** | Those tags will be added to the current alert | [optional] 
**RemoveTags** | Pointer to **[]string** | Those tags will be removed from the current alert | [optional] 

## Methods

### NewInputUpdateAlertWithIds

`func NewInputUpdateAlertWithIds(ids []string, ) *InputUpdateAlertWithIds`

NewInputUpdateAlertWithIds instantiates a new InputUpdateAlertWithIds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputUpdateAlertWithIdsWithDefaults

`func NewInputUpdateAlertWithIdsWithDefaults() *InputUpdateAlertWithIds`

NewInputUpdateAlertWithIdsWithDefaults instantiates a new InputUpdateAlertWithIds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *InputUpdateAlertWithIds) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *InputUpdateAlertWithIds) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *InputUpdateAlertWithIds) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetType

`func (o *InputUpdateAlertWithIds) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputUpdateAlertWithIds) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputUpdateAlertWithIds) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InputUpdateAlertWithIds) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSource

`func (o *InputUpdateAlertWithIds) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InputUpdateAlertWithIds) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InputUpdateAlertWithIds) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *InputUpdateAlertWithIds) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceRef

`func (o *InputUpdateAlertWithIds) GetSourceRef() string`

GetSourceRef returns the SourceRef field if non-nil, zero value otherwise.

### GetSourceRefOk

`func (o *InputUpdateAlertWithIds) GetSourceRefOk() (*string, bool)`

GetSourceRefOk returns a tuple with the SourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRef

`func (o *InputUpdateAlertWithIds) SetSourceRef(v string)`

SetSourceRef sets SourceRef field to given value.

### HasSourceRef

`func (o *InputUpdateAlertWithIds) HasSourceRef() bool`

HasSourceRef returns a boolean if a field has been set.

### GetExternalLink

`func (o *InputUpdateAlertWithIds) GetExternalLink() string`

GetExternalLink returns the ExternalLink field if non-nil, zero value otherwise.

### GetExternalLinkOk

`func (o *InputUpdateAlertWithIds) GetExternalLinkOk() (*string, bool)`

GetExternalLinkOk returns a tuple with the ExternalLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLink

`func (o *InputUpdateAlertWithIds) SetExternalLink(v string)`

SetExternalLink sets ExternalLink field to given value.

### HasExternalLink

`func (o *InputUpdateAlertWithIds) HasExternalLink() bool`

HasExternalLink returns a boolean if a field has been set.

### GetTitle

`func (o *InputUpdateAlertWithIds) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InputUpdateAlertWithIds) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InputUpdateAlertWithIds) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InputUpdateAlertWithIds) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *InputUpdateAlertWithIds) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputUpdateAlertWithIds) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputUpdateAlertWithIds) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InputUpdateAlertWithIds) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSeverity

`func (o *InputUpdateAlertWithIds) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *InputUpdateAlertWithIds) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *InputUpdateAlertWithIds) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *InputUpdateAlertWithIds) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetDate

`func (o *InputUpdateAlertWithIds) GetDate() int64`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *InputUpdateAlertWithIds) GetDateOk() (*int64, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *InputUpdateAlertWithIds) SetDate(v int64)`

SetDate sets Date field to given value.

### HasDate

`func (o *InputUpdateAlertWithIds) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetLastSyncDate

`func (o *InputUpdateAlertWithIds) GetLastSyncDate() int64`

GetLastSyncDate returns the LastSyncDate field if non-nil, zero value otherwise.

### GetLastSyncDateOk

`func (o *InputUpdateAlertWithIds) GetLastSyncDateOk() (*int64, bool)`

GetLastSyncDateOk returns a tuple with the LastSyncDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDate

`func (o *InputUpdateAlertWithIds) SetLastSyncDate(v int64)`

SetLastSyncDate sets LastSyncDate field to given value.

### HasLastSyncDate

`func (o *InputUpdateAlertWithIds) HasLastSyncDate() bool`

HasLastSyncDate returns a boolean if a field has been set.

### GetTags

`func (o *InputUpdateAlertWithIds) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InputUpdateAlertWithIds) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InputUpdateAlertWithIds) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InputUpdateAlertWithIds) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTlp

`func (o *InputUpdateAlertWithIds) GetTlp() int32`

GetTlp returns the Tlp field if non-nil, zero value otherwise.

### GetTlpOk

`func (o *InputUpdateAlertWithIds) GetTlpOk() (*int32, bool)`

GetTlpOk returns a tuple with the Tlp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlp

`func (o *InputUpdateAlertWithIds) SetTlp(v int32)`

SetTlp sets Tlp field to given value.

### HasTlp

`func (o *InputUpdateAlertWithIds) HasTlp() bool`

HasTlp returns a boolean if a field has been set.

### GetPap

`func (o *InputUpdateAlertWithIds) GetPap() int32`

GetPap returns the Pap field if non-nil, zero value otherwise.

### GetPapOk

`func (o *InputUpdateAlertWithIds) GetPapOk() (*int32, bool)`

GetPapOk returns a tuple with the Pap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPap

`func (o *InputUpdateAlertWithIds) SetPap(v int32)`

SetPap sets Pap field to given value.

### HasPap

`func (o *InputUpdateAlertWithIds) HasPap() bool`

HasPap returns a boolean if a field has been set.

### GetFollow

`func (o *InputUpdateAlertWithIds) GetFollow() bool`

GetFollow returns the Follow field if non-nil, zero value otherwise.

### GetFollowOk

`func (o *InputUpdateAlertWithIds) GetFollowOk() (*bool, bool)`

GetFollowOk returns a tuple with the Follow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollow

`func (o *InputUpdateAlertWithIds) SetFollow(v bool)`

SetFollow sets Follow field to given value.

### HasFollow

`func (o *InputUpdateAlertWithIds) HasFollow() bool`

HasFollow returns a boolean if a field has been set.

### GetCustomFields

`func (o *InputUpdateAlertWithIds) GetCustomFields() InputCreateAlertCustomFields`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *InputUpdateAlertWithIds) GetCustomFieldsOk() (*InputCreateAlertCustomFields, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *InputUpdateAlertWithIds) SetCustomFields(v InputCreateAlertCustomFields)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *InputUpdateAlertWithIds) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetStatus

`func (o *InputUpdateAlertWithIds) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InputUpdateAlertWithIds) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InputUpdateAlertWithIds) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InputUpdateAlertWithIds) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSummary

`func (o *InputUpdateAlertWithIds) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *InputUpdateAlertWithIds) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *InputUpdateAlertWithIds) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *InputUpdateAlertWithIds) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetAssignee

`func (o *InputUpdateAlertWithIds) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *InputUpdateAlertWithIds) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *InputUpdateAlertWithIds) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *InputUpdateAlertWithIds) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetAddTags

`func (o *InputUpdateAlertWithIds) GetAddTags() []string`

GetAddTags returns the AddTags field if non-nil, zero value otherwise.

### GetAddTagsOk

`func (o *InputUpdateAlertWithIds) GetAddTagsOk() (*[]string, bool)`

GetAddTagsOk returns a tuple with the AddTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTags

`func (o *InputUpdateAlertWithIds) SetAddTags(v []string)`

SetAddTags sets AddTags field to given value.

### HasAddTags

`func (o *InputUpdateAlertWithIds) HasAddTags() bool`

HasAddTags returns a boolean if a field has been set.

### GetRemoveTags

`func (o *InputUpdateAlertWithIds) GetRemoveTags() []string`

GetRemoveTags returns the RemoveTags field if non-nil, zero value otherwise.

### GetRemoveTagsOk

`func (o *InputUpdateAlertWithIds) GetRemoveTagsOk() (*[]string, bool)`

GetRemoveTagsOk returns a tuple with the RemoveTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveTags

`func (o *InputUpdateAlertWithIds) SetRemoveTags(v []string)`

SetRemoveTags sets RemoveTags field to given value.

### HasRemoveTags

`func (o *InputUpdateAlertWithIds) HasRemoveTags() bool`

HasRemoveTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


