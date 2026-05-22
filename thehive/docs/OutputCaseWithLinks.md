# OutputCaseWithLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnderscoreId** | **string** |  |
**UnderscoreType** | **string** |  |
**UnderscoreCreatedBy** | **string** |  |
**UnderscoreUpdatedBy** | Pointer to **string** |  | [optional]
**UnderscoreCreatedAt** | **int64** |  |
**UnderscoreUpdatedAt** | Pointer to **int64** |  | [optional]
**Number** | **int32** | An incremental number to reference the case. This field can be used in the apis in place of the _id |
**Title** | **string** |  |
**Description** | **string** |  |
**Severity** | **int32** |  |
**SeverityLabel** | **string** |  |
**StartDate** | **int64** |  |
**EndDate** | Pointer to **int64** |  | [optional]
**Tags** | Pointer to **[]string** |  | [optional]
**Flag** | **bool** |  |
**Tlp** | **int32** |  |
**TlpLabel** | **string** |  |
**Pap** | **int32** |  |
**PapLabel** | **string** |  |
**Status** | **string** |  |
**Stage** | **string** | The value of the stage depends on the status of the case. Can be one of &#39;New&#39; &#39;InProgress&#39; or &#39;Closed&#39; |
**Summary** | Pointer to **string** |  | [optional]
**ImpactStatus** | Pointer to **string** |  | [optional]
**Assignee** | Pointer to **string** |  | [optional]
**Access** | [**Access**](Access.md) |  |
**CustomFields** | Pointer to [**[]OutputCustomFieldValue**](OutputCustomFieldValue.md) |  | [optional]
**UserPermissions** | Pointer to **[]string** | A list of permissions the current user has access on the case. This depends on the profile of the user and the sharing of the Case | [optional]
**ExtraData** | **map[string]interface{}** |  |
**NewDate** | **int64** |  |
**InProgressDate** | Pointer to **int64** |  | [optional]
**ClosedDate** | Pointer to **int64** |  | [optional]
**AlertDate** | Pointer to **int64** |  | [optional]
**AlertNewDate** | Pointer to **int64** |  | [optional]
**AlertInProgressDate** | Pointer to **int64** |  | [optional]
**AlertImportedDate** | Pointer to **int64** |  | [optional]
**TimeToDetect** | **int64** |  |
**TimeToTriage** | Pointer to **int64** |  | [optional]
**TimeToQualify** | Pointer to **int64** |  | [optional]
**TimeToAcknowledge** | Pointer to **int64** |  | [optional]
**TimeToResolve** | Pointer to **int64** |  | [optional]
**HandlingDuration** | Pointer to **int64** |  | [optional]
**LinkedWith** | Pointer to [**[]OutputObservable**](OutputObservable.md) |  | [optional]
**LinksCount** | **int32** |  |

## Methods

### NewOutputCaseWithLinks

`func NewOutputCaseWithLinks(underscoreId string, underscoreType string, underscoreCreatedBy string, underscoreCreatedAt int64, number int32, title string, description string, severity int32, severityLabel string, startDate int64, flag bool, tlp int32, tlpLabel string, pap int32, papLabel string, status string, stage string, access Access, extraData map[string]interface{}, newDate int64, timeToDetect int64, linksCount int32, ) *OutputCaseWithLinks`

NewOutputCaseWithLinks instantiates a new OutputCaseWithLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputCaseWithLinksWithDefaults

`func NewOutputCaseWithLinksWithDefaults() *OutputCaseWithLinks`

NewOutputCaseWithLinksWithDefaults instantiates a new OutputCaseWithLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnderscoreId

`func (o *OutputCaseWithLinks) GetUnderscoreId() string`

GetUnderscoreId returns the UnderscoreId field if non-nil, zero value otherwise.

### GetUnderscoreIdOk

`func (o *OutputCaseWithLinks) GetUnderscoreIdOk() (*string, bool)`

GetUnderscoreIdOk returns a tuple with the UnderscoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreId

`func (o *OutputCaseWithLinks) SetUnderscoreId(v string)`

SetUnderscoreId sets UnderscoreId field to given value.


### GetUnderscoreType

`func (o *OutputCaseWithLinks) GetUnderscoreType() string`

GetUnderscoreType returns the UnderscoreType field if non-nil, zero value otherwise.

### GetUnderscoreTypeOk

`func (o *OutputCaseWithLinks) GetUnderscoreTypeOk() (*string, bool)`

GetUnderscoreTypeOk returns a tuple with the UnderscoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreType

`func (o *OutputCaseWithLinks) SetUnderscoreType(v string)`

SetUnderscoreType sets UnderscoreType field to given value.


### GetUnderscoreCreatedBy

`func (o *OutputCaseWithLinks) GetUnderscoreCreatedBy() string`

GetUnderscoreCreatedBy returns the UnderscoreCreatedBy field if non-nil, zero value otherwise.

### GetUnderscoreCreatedByOk

`func (o *OutputCaseWithLinks) GetUnderscoreCreatedByOk() (*string, bool)`

GetUnderscoreCreatedByOk returns a tuple with the UnderscoreCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedBy

`func (o *OutputCaseWithLinks) SetUnderscoreCreatedBy(v string)`

SetUnderscoreCreatedBy sets UnderscoreCreatedBy field to given value.


### GetUnderscoreUpdatedBy

`func (o *OutputCaseWithLinks) GetUnderscoreUpdatedBy() string`

GetUnderscoreUpdatedBy returns the UnderscoreUpdatedBy field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedByOk

`func (o *OutputCaseWithLinks) GetUnderscoreUpdatedByOk() (*string, bool)`

GetUnderscoreUpdatedByOk returns a tuple with the UnderscoreUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedBy

`func (o *OutputCaseWithLinks) SetUnderscoreUpdatedBy(v string)`

SetUnderscoreUpdatedBy sets UnderscoreUpdatedBy field to given value.

### HasUnderscoreUpdatedBy

`func (o *OutputCaseWithLinks) HasUnderscoreUpdatedBy() bool`

HasUnderscoreUpdatedBy returns a boolean if a field has been set.

### GetUnderscoreCreatedAt

`func (o *OutputCaseWithLinks) GetUnderscoreCreatedAt() int64`

GetUnderscoreCreatedAt returns the UnderscoreCreatedAt field if non-nil, zero value otherwise.

### GetUnderscoreCreatedAtOk

`func (o *OutputCaseWithLinks) GetUnderscoreCreatedAtOk() (*int64, bool)`

GetUnderscoreCreatedAtOk returns a tuple with the UnderscoreCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedAt

`func (o *OutputCaseWithLinks) SetUnderscoreCreatedAt(v int64)`

SetUnderscoreCreatedAt sets UnderscoreCreatedAt field to given value.


### GetUnderscoreUpdatedAt

`func (o *OutputCaseWithLinks) GetUnderscoreUpdatedAt() int64`

GetUnderscoreUpdatedAt returns the UnderscoreUpdatedAt field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedAtOk

`func (o *OutputCaseWithLinks) GetUnderscoreUpdatedAtOk() (*int64, bool)`

GetUnderscoreUpdatedAtOk returns a tuple with the UnderscoreUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedAt

`func (o *OutputCaseWithLinks) SetUnderscoreUpdatedAt(v int64)`

SetUnderscoreUpdatedAt sets UnderscoreUpdatedAt field to given value.

### HasUnderscoreUpdatedAt

`func (o *OutputCaseWithLinks) HasUnderscoreUpdatedAt() bool`

HasUnderscoreUpdatedAt returns a boolean if a field has been set.

### GetNumber

`func (o *OutputCaseWithLinks) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *OutputCaseWithLinks) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *OutputCaseWithLinks) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetTitle

`func (o *OutputCaseWithLinks) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OutputCaseWithLinks) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OutputCaseWithLinks) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *OutputCaseWithLinks) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OutputCaseWithLinks) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OutputCaseWithLinks) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSeverity

`func (o *OutputCaseWithLinks) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *OutputCaseWithLinks) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *OutputCaseWithLinks) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.


### GetSeverityLabel

`func (o *OutputCaseWithLinks) GetSeverityLabel() string`

GetSeverityLabel returns the SeverityLabel field if non-nil, zero value otherwise.

### GetSeverityLabelOk

`func (o *OutputCaseWithLinks) GetSeverityLabelOk() (*string, bool)`

GetSeverityLabelOk returns a tuple with the SeverityLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityLabel

`func (o *OutputCaseWithLinks) SetSeverityLabel(v string)`

SetSeverityLabel sets SeverityLabel field to given value.


### GetStartDate

`func (o *OutputCaseWithLinks) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *OutputCaseWithLinks) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *OutputCaseWithLinks) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *OutputCaseWithLinks) GetEndDate() int64`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *OutputCaseWithLinks) GetEndDateOk() (*int64, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *OutputCaseWithLinks) SetEndDate(v int64)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *OutputCaseWithLinks) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetTags

`func (o *OutputCaseWithLinks) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OutputCaseWithLinks) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OutputCaseWithLinks) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OutputCaseWithLinks) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetFlag

`func (o *OutputCaseWithLinks) GetFlag() bool`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *OutputCaseWithLinks) GetFlagOk() (*bool, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *OutputCaseWithLinks) SetFlag(v bool)`

SetFlag sets Flag field to given value.


### GetTlp

`func (o *OutputCaseWithLinks) GetTlp() int32`

GetTlp returns the Tlp field if non-nil, zero value otherwise.

### GetTlpOk

`func (o *OutputCaseWithLinks) GetTlpOk() (*int32, bool)`

GetTlpOk returns a tuple with the Tlp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlp

`func (o *OutputCaseWithLinks) SetTlp(v int32)`

SetTlp sets Tlp field to given value.


### GetTlpLabel

`func (o *OutputCaseWithLinks) GetTlpLabel() string`

GetTlpLabel returns the TlpLabel field if non-nil, zero value otherwise.

### GetTlpLabelOk

`func (o *OutputCaseWithLinks) GetTlpLabelOk() (*string, bool)`

GetTlpLabelOk returns a tuple with the TlpLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlpLabel

`func (o *OutputCaseWithLinks) SetTlpLabel(v string)`

SetTlpLabel sets TlpLabel field to given value.


### GetPap

`func (o *OutputCaseWithLinks) GetPap() int32`

GetPap returns the Pap field if non-nil, zero value otherwise.

### GetPapOk

`func (o *OutputCaseWithLinks) GetPapOk() (*int32, bool)`

GetPapOk returns a tuple with the Pap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPap

`func (o *OutputCaseWithLinks) SetPap(v int32)`

SetPap sets Pap field to given value.


### GetPapLabel

`func (o *OutputCaseWithLinks) GetPapLabel() string`

GetPapLabel returns the PapLabel field if non-nil, zero value otherwise.

### GetPapLabelOk

`func (o *OutputCaseWithLinks) GetPapLabelOk() (*string, bool)`

GetPapLabelOk returns a tuple with the PapLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPapLabel

`func (o *OutputCaseWithLinks) SetPapLabel(v string)`

SetPapLabel sets PapLabel field to given value.


### GetStatus

`func (o *OutputCaseWithLinks) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OutputCaseWithLinks) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OutputCaseWithLinks) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStage

`func (o *OutputCaseWithLinks) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *OutputCaseWithLinks) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *OutputCaseWithLinks) SetStage(v string)`

SetStage sets Stage field to given value.


### GetSummary

`func (o *OutputCaseWithLinks) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *OutputCaseWithLinks) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *OutputCaseWithLinks) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *OutputCaseWithLinks) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetImpactStatus

`func (o *OutputCaseWithLinks) GetImpactStatus() string`

GetImpactStatus returns the ImpactStatus field if non-nil, zero value otherwise.

### GetImpactStatusOk

`func (o *OutputCaseWithLinks) GetImpactStatusOk() (*string, bool)`

GetImpactStatusOk returns a tuple with the ImpactStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactStatus

`func (o *OutputCaseWithLinks) SetImpactStatus(v string)`

SetImpactStatus sets ImpactStatus field to given value.

### HasImpactStatus

`func (o *OutputCaseWithLinks) HasImpactStatus() bool`

HasImpactStatus returns a boolean if a field has been set.

### GetAssignee

`func (o *OutputCaseWithLinks) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *OutputCaseWithLinks) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *OutputCaseWithLinks) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *OutputCaseWithLinks) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetAccess

`func (o *OutputCaseWithLinks) GetAccess() Access`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *OutputCaseWithLinks) GetAccessOk() (*Access, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *OutputCaseWithLinks) SetAccess(v Access)`

SetAccess sets Access field to given value.


### GetCustomFields

`func (o *OutputCaseWithLinks) GetCustomFields() []OutputCustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *OutputCaseWithLinks) GetCustomFieldsOk() (*[]OutputCustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *OutputCaseWithLinks) SetCustomFields(v []OutputCustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *OutputCaseWithLinks) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetUserPermissions

`func (o *OutputCaseWithLinks) GetUserPermissions() []string`

GetUserPermissions returns the UserPermissions field if non-nil, zero value otherwise.

### GetUserPermissionsOk

`func (o *OutputCaseWithLinks) GetUserPermissionsOk() (*[]string, bool)`

GetUserPermissionsOk returns a tuple with the UserPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPermissions

`func (o *OutputCaseWithLinks) SetUserPermissions(v []string)`

SetUserPermissions sets UserPermissions field to given value.

### HasUserPermissions

`func (o *OutputCaseWithLinks) HasUserPermissions() bool`

HasUserPermissions returns a boolean if a field has been set.

### GetExtraData

`func (o *OutputCaseWithLinks) GetExtraData() map[string]interface{}`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *OutputCaseWithLinks) GetExtraDataOk() (*map[string]interface{}, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *OutputCaseWithLinks) SetExtraData(v map[string]interface{})`

SetExtraData sets ExtraData field to given value.


### GetNewDate

`func (o *OutputCaseWithLinks) GetNewDate() int64`

GetNewDate returns the NewDate field if non-nil, zero value otherwise.

### GetNewDateOk

`func (o *OutputCaseWithLinks) GetNewDateOk() (*int64, bool)`

GetNewDateOk returns a tuple with the NewDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewDate

`func (o *OutputCaseWithLinks) SetNewDate(v int64)`

SetNewDate sets NewDate field to given value.


### GetInProgressDate

`func (o *OutputCaseWithLinks) GetInProgressDate() int64`

GetInProgressDate returns the InProgressDate field if non-nil, zero value otherwise.

### GetInProgressDateOk

`func (o *OutputCaseWithLinks) GetInProgressDateOk() (*int64, bool)`

GetInProgressDateOk returns a tuple with the InProgressDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgressDate

`func (o *OutputCaseWithLinks) SetInProgressDate(v int64)`

SetInProgressDate sets InProgressDate field to given value.

### HasInProgressDate

`func (o *OutputCaseWithLinks) HasInProgressDate() bool`

HasInProgressDate returns a boolean if a field has been set.

### GetClosedDate

`func (o *OutputCaseWithLinks) GetClosedDate() int64`

GetClosedDate returns the ClosedDate field if non-nil, zero value otherwise.

### GetClosedDateOk

`func (o *OutputCaseWithLinks) GetClosedDateOk() (*int64, bool)`

GetClosedDateOk returns a tuple with the ClosedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedDate

`func (o *OutputCaseWithLinks) SetClosedDate(v int64)`

SetClosedDate sets ClosedDate field to given value.

### HasClosedDate

`func (o *OutputCaseWithLinks) HasClosedDate() bool`

HasClosedDate returns a boolean if a field has been set.

### GetAlertDate

`func (o *OutputCaseWithLinks) GetAlertDate() int64`

GetAlertDate returns the AlertDate field if non-nil, zero value otherwise.

### GetAlertDateOk

`func (o *OutputCaseWithLinks) GetAlertDateOk() (*int64, bool)`

GetAlertDateOk returns a tuple with the AlertDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertDate

`func (o *OutputCaseWithLinks) SetAlertDate(v int64)`

SetAlertDate sets AlertDate field to given value.

### HasAlertDate

`func (o *OutputCaseWithLinks) HasAlertDate() bool`

HasAlertDate returns a boolean if a field has been set.

### GetAlertNewDate

`func (o *OutputCaseWithLinks) GetAlertNewDate() int64`

GetAlertNewDate returns the AlertNewDate field if non-nil, zero value otherwise.

### GetAlertNewDateOk

`func (o *OutputCaseWithLinks) GetAlertNewDateOk() (*int64, bool)`

GetAlertNewDateOk returns a tuple with the AlertNewDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertNewDate

`func (o *OutputCaseWithLinks) SetAlertNewDate(v int64)`

SetAlertNewDate sets AlertNewDate field to given value.

### HasAlertNewDate

`func (o *OutputCaseWithLinks) HasAlertNewDate() bool`

HasAlertNewDate returns a boolean if a field has been set.

### GetAlertInProgressDate

`func (o *OutputCaseWithLinks) GetAlertInProgressDate() int64`

GetAlertInProgressDate returns the AlertInProgressDate field if non-nil, zero value otherwise.

### GetAlertInProgressDateOk

`func (o *OutputCaseWithLinks) GetAlertInProgressDateOk() (*int64, bool)`

GetAlertInProgressDateOk returns a tuple with the AlertInProgressDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertInProgressDate

`func (o *OutputCaseWithLinks) SetAlertInProgressDate(v int64)`

SetAlertInProgressDate sets AlertInProgressDate field to given value.

### HasAlertInProgressDate

`func (o *OutputCaseWithLinks) HasAlertInProgressDate() bool`

HasAlertInProgressDate returns a boolean if a field has been set.

### GetAlertImportedDate

`func (o *OutputCaseWithLinks) GetAlertImportedDate() int64`

GetAlertImportedDate returns the AlertImportedDate field if non-nil, zero value otherwise.

### GetAlertImportedDateOk

`func (o *OutputCaseWithLinks) GetAlertImportedDateOk() (*int64, bool)`

GetAlertImportedDateOk returns a tuple with the AlertImportedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertImportedDate

`func (o *OutputCaseWithLinks) SetAlertImportedDate(v int64)`

SetAlertImportedDate sets AlertImportedDate field to given value.

### HasAlertImportedDate

`func (o *OutputCaseWithLinks) HasAlertImportedDate() bool`

HasAlertImportedDate returns a boolean if a field has been set.

### GetTimeToDetect

`func (o *OutputCaseWithLinks) GetTimeToDetect() int64`

GetTimeToDetect returns the TimeToDetect field if non-nil, zero value otherwise.

### GetTimeToDetectOk

`func (o *OutputCaseWithLinks) GetTimeToDetectOk() (*int64, bool)`

GetTimeToDetectOk returns a tuple with the TimeToDetect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToDetect

`func (o *OutputCaseWithLinks) SetTimeToDetect(v int64)`

SetTimeToDetect sets TimeToDetect field to given value.


### GetTimeToTriage

`func (o *OutputCaseWithLinks) GetTimeToTriage() int64`

GetTimeToTriage returns the TimeToTriage field if non-nil, zero value otherwise.

### GetTimeToTriageOk

`func (o *OutputCaseWithLinks) GetTimeToTriageOk() (*int64, bool)`

GetTimeToTriageOk returns a tuple with the TimeToTriage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToTriage

`func (o *OutputCaseWithLinks) SetTimeToTriage(v int64)`

SetTimeToTriage sets TimeToTriage field to given value.

### HasTimeToTriage

`func (o *OutputCaseWithLinks) HasTimeToTriage() bool`

HasTimeToTriage returns a boolean if a field has been set.

### GetTimeToQualify

`func (o *OutputCaseWithLinks) GetTimeToQualify() int64`

GetTimeToQualify returns the TimeToQualify field if non-nil, zero value otherwise.

### GetTimeToQualifyOk

`func (o *OutputCaseWithLinks) GetTimeToQualifyOk() (*int64, bool)`

GetTimeToQualifyOk returns a tuple with the TimeToQualify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToQualify

`func (o *OutputCaseWithLinks) SetTimeToQualify(v int64)`

SetTimeToQualify sets TimeToQualify field to given value.

### HasTimeToQualify

`func (o *OutputCaseWithLinks) HasTimeToQualify() bool`

HasTimeToQualify returns a boolean if a field has been set.

### GetTimeToAcknowledge

`func (o *OutputCaseWithLinks) GetTimeToAcknowledge() int64`

GetTimeToAcknowledge returns the TimeToAcknowledge field if non-nil, zero value otherwise.

### GetTimeToAcknowledgeOk

`func (o *OutputCaseWithLinks) GetTimeToAcknowledgeOk() (*int64, bool)`

GetTimeToAcknowledgeOk returns a tuple with the TimeToAcknowledge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToAcknowledge

`func (o *OutputCaseWithLinks) SetTimeToAcknowledge(v int64)`

SetTimeToAcknowledge sets TimeToAcknowledge field to given value.

### HasTimeToAcknowledge

`func (o *OutputCaseWithLinks) HasTimeToAcknowledge() bool`

HasTimeToAcknowledge returns a boolean if a field has been set.

### GetTimeToResolve

`func (o *OutputCaseWithLinks) GetTimeToResolve() int64`

GetTimeToResolve returns the TimeToResolve field if non-nil, zero value otherwise.

### GetTimeToResolveOk

`func (o *OutputCaseWithLinks) GetTimeToResolveOk() (*int64, bool)`

GetTimeToResolveOk returns a tuple with the TimeToResolve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToResolve

`func (o *OutputCaseWithLinks) SetTimeToResolve(v int64)`

SetTimeToResolve sets TimeToResolve field to given value.

### HasTimeToResolve

`func (o *OutputCaseWithLinks) HasTimeToResolve() bool`

HasTimeToResolve returns a boolean if a field has been set.

### GetHandlingDuration

`func (o *OutputCaseWithLinks) GetHandlingDuration() int64`

GetHandlingDuration returns the HandlingDuration field if non-nil, zero value otherwise.

### GetHandlingDurationOk

`func (o *OutputCaseWithLinks) GetHandlingDurationOk() (*int64, bool)`

GetHandlingDurationOk returns a tuple with the HandlingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlingDuration

`func (o *OutputCaseWithLinks) SetHandlingDuration(v int64)`

SetHandlingDuration sets HandlingDuration field to given value.

### HasHandlingDuration

`func (o *OutputCaseWithLinks) HasHandlingDuration() bool`

HasHandlingDuration returns a boolean if a field has been set.

### GetLinkedWith

`func (o *OutputCaseWithLinks) GetLinkedWith() []OutputObservable`

GetLinkedWith returns the LinkedWith field if non-nil, zero value otherwise.

### GetLinkedWithOk

`func (o *OutputCaseWithLinks) GetLinkedWithOk() (*[]OutputObservable, bool)`

GetLinkedWithOk returns a tuple with the LinkedWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedWith

`func (o *OutputCaseWithLinks) SetLinkedWith(v []OutputObservable)`

SetLinkedWith sets LinkedWith field to given value.

### HasLinkedWith

`func (o *OutputCaseWithLinks) HasLinkedWith() bool`

HasLinkedWith returns a boolean if a field has been set.

### GetLinksCount

`func (o *OutputCaseWithLinks) GetLinksCount() int32`

GetLinksCount returns the LinksCount field if non-nil, zero value otherwise.

### GetLinksCountOk

`func (o *OutputCaseWithLinks) GetLinksCountOk() (*int32, bool)`

GetLinksCountOk returns a tuple with the LinksCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinksCount

`func (o *OutputCaseWithLinks) SetLinksCount(v int32)`

SetLinksCount sets LinksCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
