# OutputAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnderscoreId** | **string** |  | 
**UnderscoreType** | **string** |  | 
**UnderscoreCreatedBy** | **string** |  | 
**UnderscoreUpdatedBy** | Pointer to **string** |  | [optional] 
**UnderscoreCreatedAt** | **int64** |  | 
**UnderscoreUpdatedAt** | Pointer to **int64** |  | [optional] 
**Type** | **string** |  | 
**Source** | **string** |  | 
**SourceRef** | **string** |  | 
**ExternalLink** | Pointer to **string** |  | [optional] 
**Title** | **string** |  | 
**Description** | **string** |  | 
**Severity** | **int32** |  | 
**SeverityLabel** | **string** |  | 
**Date** | **int64** |  | 
**Tags** | Pointer to **[]string** |  | [optional] 
**Tlp** | **int32** |  | 
**TlpLabel** | **string** |  | 
**Pap** | **int32** |  | 
**PapLabel** | **string** |  | 
**Follow** | **bool** |  | 
**CustomFields** | Pointer to [**[]OutputCustomFieldValue**](OutputCustomFieldValue.md) |  | [optional] 
**CaseTemplate** | Pointer to **string** |  | [optional] 
**ObservableCount** | **int64** |  | 
**CaseId** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**Stage** | **string** |  | 
**Assignee** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**ExtraData** | **map[string]interface{}** |  | 
**NewDate** | **int64** |  | 
**InProgressDate** | Pointer to **int64** |  | [optional] 
**ClosedDate** | Pointer to **int64** |  | [optional] 
**ImportedDate** | Pointer to **int64** |  | [optional] 
**TimeToDetect** | **int64** |  | 
**TimeToTriage** | Pointer to **int64** |  | [optional] 
**TimeToQualify** | Pointer to **int64** |  | [optional] 
**TimeToAcknowledge** | Pointer to **int64** |  | [optional] 

## Methods

### NewOutputAlert

`func NewOutputAlert(underscoreId string, underscoreType string, underscoreCreatedBy string, underscoreCreatedAt int64, type_ string, source string, sourceRef string, title string, description string, severity int32, severityLabel string, date int64, tlp int32, tlpLabel string, pap int32, papLabel string, follow bool, observableCount int64, status string, stage string, extraData map[string]interface{}, newDate int64, timeToDetect int64, ) *OutputAlert`

NewOutputAlert instantiates a new OutputAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputAlertWithDefaults

`func NewOutputAlertWithDefaults() *OutputAlert`

NewOutputAlertWithDefaults instantiates a new OutputAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnderscoreId

`func (o *OutputAlert) GetUnderscoreId() string`

GetUnderscoreId returns the UnderscoreId field if non-nil, zero value otherwise.

### GetUnderscoreIdOk

`func (o *OutputAlert) GetUnderscoreIdOk() (*string, bool)`

GetUnderscoreIdOk returns a tuple with the UnderscoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreId

`func (o *OutputAlert) SetUnderscoreId(v string)`

SetUnderscoreId sets UnderscoreId field to given value.


### GetUnderscoreType

`func (o *OutputAlert) GetUnderscoreType() string`

GetUnderscoreType returns the UnderscoreType field if non-nil, zero value otherwise.

### GetUnderscoreTypeOk

`func (o *OutputAlert) GetUnderscoreTypeOk() (*string, bool)`

GetUnderscoreTypeOk returns a tuple with the UnderscoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreType

`func (o *OutputAlert) SetUnderscoreType(v string)`

SetUnderscoreType sets UnderscoreType field to given value.


### GetUnderscoreCreatedBy

`func (o *OutputAlert) GetUnderscoreCreatedBy() string`

GetUnderscoreCreatedBy returns the UnderscoreCreatedBy field if non-nil, zero value otherwise.

### GetUnderscoreCreatedByOk

`func (o *OutputAlert) GetUnderscoreCreatedByOk() (*string, bool)`

GetUnderscoreCreatedByOk returns a tuple with the UnderscoreCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedBy

`func (o *OutputAlert) SetUnderscoreCreatedBy(v string)`

SetUnderscoreCreatedBy sets UnderscoreCreatedBy field to given value.


### GetUnderscoreUpdatedBy

`func (o *OutputAlert) GetUnderscoreUpdatedBy() string`

GetUnderscoreUpdatedBy returns the UnderscoreUpdatedBy field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedByOk

`func (o *OutputAlert) GetUnderscoreUpdatedByOk() (*string, bool)`

GetUnderscoreUpdatedByOk returns a tuple with the UnderscoreUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedBy

`func (o *OutputAlert) SetUnderscoreUpdatedBy(v string)`

SetUnderscoreUpdatedBy sets UnderscoreUpdatedBy field to given value.

### HasUnderscoreUpdatedBy

`func (o *OutputAlert) HasUnderscoreUpdatedBy() bool`

HasUnderscoreUpdatedBy returns a boolean if a field has been set.

### GetUnderscoreCreatedAt

`func (o *OutputAlert) GetUnderscoreCreatedAt() int64`

GetUnderscoreCreatedAt returns the UnderscoreCreatedAt field if non-nil, zero value otherwise.

### GetUnderscoreCreatedAtOk

`func (o *OutputAlert) GetUnderscoreCreatedAtOk() (*int64, bool)`

GetUnderscoreCreatedAtOk returns a tuple with the UnderscoreCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedAt

`func (o *OutputAlert) SetUnderscoreCreatedAt(v int64)`

SetUnderscoreCreatedAt sets UnderscoreCreatedAt field to given value.


### GetUnderscoreUpdatedAt

`func (o *OutputAlert) GetUnderscoreUpdatedAt() int64`

GetUnderscoreUpdatedAt returns the UnderscoreUpdatedAt field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedAtOk

`func (o *OutputAlert) GetUnderscoreUpdatedAtOk() (*int64, bool)`

GetUnderscoreUpdatedAtOk returns a tuple with the UnderscoreUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedAt

`func (o *OutputAlert) SetUnderscoreUpdatedAt(v int64)`

SetUnderscoreUpdatedAt sets UnderscoreUpdatedAt field to given value.

### HasUnderscoreUpdatedAt

`func (o *OutputAlert) HasUnderscoreUpdatedAt() bool`

HasUnderscoreUpdatedAt returns a boolean if a field has been set.

### GetType

`func (o *OutputAlert) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OutputAlert) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OutputAlert) SetType(v string)`

SetType sets Type field to given value.


### GetSource

`func (o *OutputAlert) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *OutputAlert) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *OutputAlert) SetSource(v string)`

SetSource sets Source field to given value.


### GetSourceRef

`func (o *OutputAlert) GetSourceRef() string`

GetSourceRef returns the SourceRef field if non-nil, zero value otherwise.

### GetSourceRefOk

`func (o *OutputAlert) GetSourceRefOk() (*string, bool)`

GetSourceRefOk returns a tuple with the SourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRef

`func (o *OutputAlert) SetSourceRef(v string)`

SetSourceRef sets SourceRef field to given value.


### GetExternalLink

`func (o *OutputAlert) GetExternalLink() string`

GetExternalLink returns the ExternalLink field if non-nil, zero value otherwise.

### GetExternalLinkOk

`func (o *OutputAlert) GetExternalLinkOk() (*string, bool)`

GetExternalLinkOk returns a tuple with the ExternalLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLink

`func (o *OutputAlert) SetExternalLink(v string)`

SetExternalLink sets ExternalLink field to given value.

### HasExternalLink

`func (o *OutputAlert) HasExternalLink() bool`

HasExternalLink returns a boolean if a field has been set.

### GetTitle

`func (o *OutputAlert) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OutputAlert) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OutputAlert) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *OutputAlert) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OutputAlert) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OutputAlert) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSeverity

`func (o *OutputAlert) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *OutputAlert) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *OutputAlert) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.


### GetSeverityLabel

`func (o *OutputAlert) GetSeverityLabel() string`

GetSeverityLabel returns the SeverityLabel field if non-nil, zero value otherwise.

### GetSeverityLabelOk

`func (o *OutputAlert) GetSeverityLabelOk() (*string, bool)`

GetSeverityLabelOk returns a tuple with the SeverityLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityLabel

`func (o *OutputAlert) SetSeverityLabel(v string)`

SetSeverityLabel sets SeverityLabel field to given value.


### GetDate

`func (o *OutputAlert) GetDate() int64`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *OutputAlert) GetDateOk() (*int64, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *OutputAlert) SetDate(v int64)`

SetDate sets Date field to given value.


### GetTags

`func (o *OutputAlert) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OutputAlert) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OutputAlert) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OutputAlert) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTlp

`func (o *OutputAlert) GetTlp() int32`

GetTlp returns the Tlp field if non-nil, zero value otherwise.

### GetTlpOk

`func (o *OutputAlert) GetTlpOk() (*int32, bool)`

GetTlpOk returns a tuple with the Tlp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlp

`func (o *OutputAlert) SetTlp(v int32)`

SetTlp sets Tlp field to given value.


### GetTlpLabel

`func (o *OutputAlert) GetTlpLabel() string`

GetTlpLabel returns the TlpLabel field if non-nil, zero value otherwise.

### GetTlpLabelOk

`func (o *OutputAlert) GetTlpLabelOk() (*string, bool)`

GetTlpLabelOk returns a tuple with the TlpLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlpLabel

`func (o *OutputAlert) SetTlpLabel(v string)`

SetTlpLabel sets TlpLabel field to given value.


### GetPap

`func (o *OutputAlert) GetPap() int32`

GetPap returns the Pap field if non-nil, zero value otherwise.

### GetPapOk

`func (o *OutputAlert) GetPapOk() (*int32, bool)`

GetPapOk returns a tuple with the Pap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPap

`func (o *OutputAlert) SetPap(v int32)`

SetPap sets Pap field to given value.


### GetPapLabel

`func (o *OutputAlert) GetPapLabel() string`

GetPapLabel returns the PapLabel field if non-nil, zero value otherwise.

### GetPapLabelOk

`func (o *OutputAlert) GetPapLabelOk() (*string, bool)`

GetPapLabelOk returns a tuple with the PapLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPapLabel

`func (o *OutputAlert) SetPapLabel(v string)`

SetPapLabel sets PapLabel field to given value.


### GetFollow

`func (o *OutputAlert) GetFollow() bool`

GetFollow returns the Follow field if non-nil, zero value otherwise.

### GetFollowOk

`func (o *OutputAlert) GetFollowOk() (*bool, bool)`

GetFollowOk returns a tuple with the Follow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollow

`func (o *OutputAlert) SetFollow(v bool)`

SetFollow sets Follow field to given value.


### GetCustomFields

`func (o *OutputAlert) GetCustomFields() []OutputCustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *OutputAlert) GetCustomFieldsOk() (*[]OutputCustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *OutputAlert) SetCustomFields(v []OutputCustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *OutputAlert) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCaseTemplate

`func (o *OutputAlert) GetCaseTemplate() string`

GetCaseTemplate returns the CaseTemplate field if non-nil, zero value otherwise.

### GetCaseTemplateOk

`func (o *OutputAlert) GetCaseTemplateOk() (*string, bool)`

GetCaseTemplateOk returns a tuple with the CaseTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseTemplate

`func (o *OutputAlert) SetCaseTemplate(v string)`

SetCaseTemplate sets CaseTemplate field to given value.

### HasCaseTemplate

`func (o *OutputAlert) HasCaseTemplate() bool`

HasCaseTemplate returns a boolean if a field has been set.

### GetObservableCount

`func (o *OutputAlert) GetObservableCount() int64`

GetObservableCount returns the ObservableCount field if non-nil, zero value otherwise.

### GetObservableCountOk

`func (o *OutputAlert) GetObservableCountOk() (*int64, bool)`

GetObservableCountOk returns a tuple with the ObservableCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservableCount

`func (o *OutputAlert) SetObservableCount(v int64)`

SetObservableCount sets ObservableCount field to given value.


### GetCaseId

`func (o *OutputAlert) GetCaseId() string`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *OutputAlert) GetCaseIdOk() (*string, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *OutputAlert) SetCaseId(v string)`

SetCaseId sets CaseId field to given value.

### HasCaseId

`func (o *OutputAlert) HasCaseId() bool`

HasCaseId returns a boolean if a field has been set.

### GetStatus

`func (o *OutputAlert) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OutputAlert) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OutputAlert) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStage

`func (o *OutputAlert) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *OutputAlert) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *OutputAlert) SetStage(v string)`

SetStage sets Stage field to given value.


### GetAssignee

`func (o *OutputAlert) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *OutputAlert) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *OutputAlert) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *OutputAlert) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetSummary

`func (o *OutputAlert) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *OutputAlert) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *OutputAlert) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *OutputAlert) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetExtraData

`func (o *OutputAlert) GetExtraData() map[string]interface{}`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *OutputAlert) GetExtraDataOk() (*map[string]interface{}, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *OutputAlert) SetExtraData(v map[string]interface{})`

SetExtraData sets ExtraData field to given value.


### GetNewDate

`func (o *OutputAlert) GetNewDate() int64`

GetNewDate returns the NewDate field if non-nil, zero value otherwise.

### GetNewDateOk

`func (o *OutputAlert) GetNewDateOk() (*int64, bool)`

GetNewDateOk returns a tuple with the NewDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewDate

`func (o *OutputAlert) SetNewDate(v int64)`

SetNewDate sets NewDate field to given value.


### GetInProgressDate

`func (o *OutputAlert) GetInProgressDate() int64`

GetInProgressDate returns the InProgressDate field if non-nil, zero value otherwise.

### GetInProgressDateOk

`func (o *OutputAlert) GetInProgressDateOk() (*int64, bool)`

GetInProgressDateOk returns a tuple with the InProgressDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgressDate

`func (o *OutputAlert) SetInProgressDate(v int64)`

SetInProgressDate sets InProgressDate field to given value.

### HasInProgressDate

`func (o *OutputAlert) HasInProgressDate() bool`

HasInProgressDate returns a boolean if a field has been set.

### GetClosedDate

`func (o *OutputAlert) GetClosedDate() int64`

GetClosedDate returns the ClosedDate field if non-nil, zero value otherwise.

### GetClosedDateOk

`func (o *OutputAlert) GetClosedDateOk() (*int64, bool)`

GetClosedDateOk returns a tuple with the ClosedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedDate

`func (o *OutputAlert) SetClosedDate(v int64)`

SetClosedDate sets ClosedDate field to given value.

### HasClosedDate

`func (o *OutputAlert) HasClosedDate() bool`

HasClosedDate returns a boolean if a field has been set.

### GetImportedDate

`func (o *OutputAlert) GetImportedDate() int64`

GetImportedDate returns the ImportedDate field if non-nil, zero value otherwise.

### GetImportedDateOk

`func (o *OutputAlert) GetImportedDateOk() (*int64, bool)`

GetImportedDateOk returns a tuple with the ImportedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedDate

`func (o *OutputAlert) SetImportedDate(v int64)`

SetImportedDate sets ImportedDate field to given value.

### HasImportedDate

`func (o *OutputAlert) HasImportedDate() bool`

HasImportedDate returns a boolean if a field has been set.

### GetTimeToDetect

`func (o *OutputAlert) GetTimeToDetect() int64`

GetTimeToDetect returns the TimeToDetect field if non-nil, zero value otherwise.

### GetTimeToDetectOk

`func (o *OutputAlert) GetTimeToDetectOk() (*int64, bool)`

GetTimeToDetectOk returns a tuple with the TimeToDetect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToDetect

`func (o *OutputAlert) SetTimeToDetect(v int64)`

SetTimeToDetect sets TimeToDetect field to given value.


### GetTimeToTriage

`func (o *OutputAlert) GetTimeToTriage() int64`

GetTimeToTriage returns the TimeToTriage field if non-nil, zero value otherwise.

### GetTimeToTriageOk

`func (o *OutputAlert) GetTimeToTriageOk() (*int64, bool)`

GetTimeToTriageOk returns a tuple with the TimeToTriage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToTriage

`func (o *OutputAlert) SetTimeToTriage(v int64)`

SetTimeToTriage sets TimeToTriage field to given value.

### HasTimeToTriage

`func (o *OutputAlert) HasTimeToTriage() bool`

HasTimeToTriage returns a boolean if a field has been set.

### GetTimeToQualify

`func (o *OutputAlert) GetTimeToQualify() int64`

GetTimeToQualify returns the TimeToQualify field if non-nil, zero value otherwise.

### GetTimeToQualifyOk

`func (o *OutputAlert) GetTimeToQualifyOk() (*int64, bool)`

GetTimeToQualifyOk returns a tuple with the TimeToQualify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToQualify

`func (o *OutputAlert) SetTimeToQualify(v int64)`

SetTimeToQualify sets TimeToQualify field to given value.

### HasTimeToQualify

`func (o *OutputAlert) HasTimeToQualify() bool`

HasTimeToQualify returns a boolean if a field has been set.

### GetTimeToAcknowledge

`func (o *OutputAlert) GetTimeToAcknowledge() int64`

GetTimeToAcknowledge returns the TimeToAcknowledge field if non-nil, zero value otherwise.

### GetTimeToAcknowledgeOk

`func (o *OutputAlert) GetTimeToAcknowledgeOk() (*int64, bool)`

GetTimeToAcknowledgeOk returns a tuple with the TimeToAcknowledge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToAcknowledge

`func (o *OutputAlert) SetTimeToAcknowledge(v int64)`

SetTimeToAcknowledge sets TimeToAcknowledge field to given value.

### HasTimeToAcknowledge

`func (o *OutputAlert) HasTimeToAcknowledge() bool`

HasTimeToAcknowledge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


