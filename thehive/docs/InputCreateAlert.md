# InputCreateAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Source** | **string** |  | 
**SourceRef** | **string** |  | 
**ExternalLink** | Pointer to **string** |  | [optional] 
**Title** | **string** |  | 
**Description** | **string** |  | 
**Severity** | Pointer to **int32** |  | [optional] 
**Date** | Pointer to **int32** |  | [optional] [default to now()]
**Tags** | Pointer to **[]string** |  | [optional] 
**Flag** | Pointer to **bool** |  | [optional] 
**Tlp** | Pointer to **int32** |  | [optional] 
**Pap** | Pointer to **int32** |  | [optional] 
**CustomFields** | Pointer to [**InputCreateAlertCustomFields**](InputCreateAlertCustomFields.md) |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Assignee** | Pointer to **string** | User to assign the alert to | [optional] 
**CaseTemplate** | Pointer to **string** |  | [optional] 
**Observables** | Pointer to [**[]InputCreateObservable**](InputCreateObservable.md) |  | [optional] 
**Procedures** | Pointer to [**[]InputProcedure**](InputProcedure.md) | List of procedures (TTPs) to link the alert to | [optional] 

## Methods

### NewInputCreateAlert

`func NewInputCreateAlert(type_ string, source string, sourceRef string, title string, description string, ) *InputCreateAlert`

NewInputCreateAlert instantiates a new InputCreateAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputCreateAlertWithDefaults

`func NewInputCreateAlertWithDefaults() *InputCreateAlert`

NewInputCreateAlertWithDefaults instantiates a new InputCreateAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputCreateAlert) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputCreateAlert) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputCreateAlert) SetType(v string)`

SetType sets Type field to given value.


### GetSource

`func (o *InputCreateAlert) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InputCreateAlert) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InputCreateAlert) SetSource(v string)`

SetSource sets Source field to given value.


### GetSourceRef

`func (o *InputCreateAlert) GetSourceRef() string`

GetSourceRef returns the SourceRef field if non-nil, zero value otherwise.

### GetSourceRefOk

`func (o *InputCreateAlert) GetSourceRefOk() (*string, bool)`

GetSourceRefOk returns a tuple with the SourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRef

`func (o *InputCreateAlert) SetSourceRef(v string)`

SetSourceRef sets SourceRef field to given value.


### GetExternalLink

`func (o *InputCreateAlert) GetExternalLink() string`

GetExternalLink returns the ExternalLink field if non-nil, zero value otherwise.

### GetExternalLinkOk

`func (o *InputCreateAlert) GetExternalLinkOk() (*string, bool)`

GetExternalLinkOk returns a tuple with the ExternalLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLink

`func (o *InputCreateAlert) SetExternalLink(v string)`

SetExternalLink sets ExternalLink field to given value.

### HasExternalLink

`func (o *InputCreateAlert) HasExternalLink() bool`

HasExternalLink returns a boolean if a field has been set.

### GetTitle

`func (o *InputCreateAlert) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InputCreateAlert) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InputCreateAlert) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *InputCreateAlert) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputCreateAlert) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputCreateAlert) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSeverity

`func (o *InputCreateAlert) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *InputCreateAlert) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *InputCreateAlert) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *InputCreateAlert) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetDate

`func (o *InputCreateAlert) GetDate() int32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *InputCreateAlert) GetDateOk() (*int32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *InputCreateAlert) SetDate(v int32)`

SetDate sets Date field to given value.

### HasDate

`func (o *InputCreateAlert) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetTags

`func (o *InputCreateAlert) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InputCreateAlert) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InputCreateAlert) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InputCreateAlert) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetFlag

`func (o *InputCreateAlert) GetFlag() bool`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *InputCreateAlert) GetFlagOk() (*bool, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *InputCreateAlert) SetFlag(v bool)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *InputCreateAlert) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetTlp

`func (o *InputCreateAlert) GetTlp() int32`

GetTlp returns the Tlp field if non-nil, zero value otherwise.

### GetTlpOk

`func (o *InputCreateAlert) GetTlpOk() (*int32, bool)`

GetTlpOk returns a tuple with the Tlp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlp

`func (o *InputCreateAlert) SetTlp(v int32)`

SetTlp sets Tlp field to given value.

### HasTlp

`func (o *InputCreateAlert) HasTlp() bool`

HasTlp returns a boolean if a field has been set.

### GetPap

`func (o *InputCreateAlert) GetPap() int32`

GetPap returns the Pap field if non-nil, zero value otherwise.

### GetPapOk

`func (o *InputCreateAlert) GetPapOk() (*int32, bool)`

GetPapOk returns a tuple with the Pap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPap

`func (o *InputCreateAlert) SetPap(v int32)`

SetPap sets Pap field to given value.

### HasPap

`func (o *InputCreateAlert) HasPap() bool`

HasPap returns a boolean if a field has been set.

### GetCustomFields

`func (o *InputCreateAlert) GetCustomFields() InputCreateAlertCustomFields`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *InputCreateAlert) GetCustomFieldsOk() (*InputCreateAlertCustomFields, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *InputCreateAlert) SetCustomFields(v InputCreateAlertCustomFields)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *InputCreateAlert) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetSummary

`func (o *InputCreateAlert) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *InputCreateAlert) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *InputCreateAlert) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *InputCreateAlert) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetStatus

`func (o *InputCreateAlert) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InputCreateAlert) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InputCreateAlert) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InputCreateAlert) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAssignee

`func (o *InputCreateAlert) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *InputCreateAlert) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *InputCreateAlert) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *InputCreateAlert) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetCaseTemplate

`func (o *InputCreateAlert) GetCaseTemplate() string`

GetCaseTemplate returns the CaseTemplate field if non-nil, zero value otherwise.

### GetCaseTemplateOk

`func (o *InputCreateAlert) GetCaseTemplateOk() (*string, bool)`

GetCaseTemplateOk returns a tuple with the CaseTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseTemplate

`func (o *InputCreateAlert) SetCaseTemplate(v string)`

SetCaseTemplate sets CaseTemplate field to given value.

### HasCaseTemplate

`func (o *InputCreateAlert) HasCaseTemplate() bool`

HasCaseTemplate returns a boolean if a field has been set.

### GetObservables

`func (o *InputCreateAlert) GetObservables() []InputCreateObservable`

GetObservables returns the Observables field if non-nil, zero value otherwise.

### GetObservablesOk

`func (o *InputCreateAlert) GetObservablesOk() (*[]InputCreateObservable, bool)`

GetObservablesOk returns a tuple with the Observables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservables

`func (o *InputCreateAlert) SetObservables(v []InputCreateObservable)`

SetObservables sets Observables field to given value.

### HasObservables

`func (o *InputCreateAlert) HasObservables() bool`

HasObservables returns a boolean if a field has been set.

### GetProcedures

`func (o *InputCreateAlert) GetProcedures() []InputProcedure`

GetProcedures returns the Procedures field if non-nil, zero value otherwise.

### GetProceduresOk

`func (o *InputCreateAlert) GetProceduresOk() (*[]InputProcedure, bool)`

GetProceduresOk returns a tuple with the Procedures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcedures

`func (o *InputCreateAlert) SetProcedures(v []InputProcedure)`

SetProcedures sets Procedures field to given value.

### HasProcedures

`func (o *InputCreateAlert) HasProcedures() bool`

HasProcedures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


