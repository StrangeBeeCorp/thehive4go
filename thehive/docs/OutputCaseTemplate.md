# OutputCaseTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnderscoreId** | **string** |  | 
**UnderscoreType** | **string** |  | 
**UnderscoreCreatedBy** | **string** |  | 
**UnderscoreUpdatedBy** | Pointer to **string** |  | [optional] 
**UnderscoreCreatedAt** | **int64** |  | 
**UnderscoreUpdatedAt** | Pointer to **int64** |  | [optional] 
**Name** | **string** |  | 
**DisplayName** | **string** |  | 
**TitlePrefix** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **int32** |  | [optional] 
**SeverityLabel** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Flag** | **bool** |  | 
**Tlp** | Pointer to **int32** |  | [optional] 
**TlpLabel** | Pointer to **string** |  | [optional] 
**Pap** | Pointer to **int32** |  | [optional] 
**PapLabel** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to [**[]OutputCustomFieldValue**](OutputCustomFieldValue.md) |  | [optional] 
**Tasks** | Pointer to [**[]OutputTask**](OutputTask.md) |  | [optional] 
**ExtraData** | **map[string]interface{}** |  | 

## Methods

### NewOutputCaseTemplate

`func NewOutputCaseTemplate(underscoreId string, underscoreType string, underscoreCreatedBy string, underscoreCreatedAt int64, name string, displayName string, flag bool, extraData map[string]interface{}, ) *OutputCaseTemplate`

NewOutputCaseTemplate instantiates a new OutputCaseTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputCaseTemplateWithDefaults

`func NewOutputCaseTemplateWithDefaults() *OutputCaseTemplate`

NewOutputCaseTemplateWithDefaults instantiates a new OutputCaseTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnderscoreId

`func (o *OutputCaseTemplate) GetUnderscoreId() string`

GetUnderscoreId returns the UnderscoreId field if non-nil, zero value otherwise.

### GetUnderscoreIdOk

`func (o *OutputCaseTemplate) GetUnderscoreIdOk() (*string, bool)`

GetUnderscoreIdOk returns a tuple with the UnderscoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreId

`func (o *OutputCaseTemplate) SetUnderscoreId(v string)`

SetUnderscoreId sets UnderscoreId field to given value.


### GetUnderscoreType

`func (o *OutputCaseTemplate) GetUnderscoreType() string`

GetUnderscoreType returns the UnderscoreType field if non-nil, zero value otherwise.

### GetUnderscoreTypeOk

`func (o *OutputCaseTemplate) GetUnderscoreTypeOk() (*string, bool)`

GetUnderscoreTypeOk returns a tuple with the UnderscoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreType

`func (o *OutputCaseTemplate) SetUnderscoreType(v string)`

SetUnderscoreType sets UnderscoreType field to given value.


### GetUnderscoreCreatedBy

`func (o *OutputCaseTemplate) GetUnderscoreCreatedBy() string`

GetUnderscoreCreatedBy returns the UnderscoreCreatedBy field if non-nil, zero value otherwise.

### GetUnderscoreCreatedByOk

`func (o *OutputCaseTemplate) GetUnderscoreCreatedByOk() (*string, bool)`

GetUnderscoreCreatedByOk returns a tuple with the UnderscoreCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedBy

`func (o *OutputCaseTemplate) SetUnderscoreCreatedBy(v string)`

SetUnderscoreCreatedBy sets UnderscoreCreatedBy field to given value.


### GetUnderscoreUpdatedBy

`func (o *OutputCaseTemplate) GetUnderscoreUpdatedBy() string`

GetUnderscoreUpdatedBy returns the UnderscoreUpdatedBy field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedByOk

`func (o *OutputCaseTemplate) GetUnderscoreUpdatedByOk() (*string, bool)`

GetUnderscoreUpdatedByOk returns a tuple with the UnderscoreUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedBy

`func (o *OutputCaseTemplate) SetUnderscoreUpdatedBy(v string)`

SetUnderscoreUpdatedBy sets UnderscoreUpdatedBy field to given value.

### HasUnderscoreUpdatedBy

`func (o *OutputCaseTemplate) HasUnderscoreUpdatedBy() bool`

HasUnderscoreUpdatedBy returns a boolean if a field has been set.

### GetUnderscoreCreatedAt

`func (o *OutputCaseTemplate) GetUnderscoreCreatedAt() int64`

GetUnderscoreCreatedAt returns the UnderscoreCreatedAt field if non-nil, zero value otherwise.

### GetUnderscoreCreatedAtOk

`func (o *OutputCaseTemplate) GetUnderscoreCreatedAtOk() (*int64, bool)`

GetUnderscoreCreatedAtOk returns a tuple with the UnderscoreCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedAt

`func (o *OutputCaseTemplate) SetUnderscoreCreatedAt(v int64)`

SetUnderscoreCreatedAt sets UnderscoreCreatedAt field to given value.


### GetUnderscoreUpdatedAt

`func (o *OutputCaseTemplate) GetUnderscoreUpdatedAt() int64`

GetUnderscoreUpdatedAt returns the UnderscoreUpdatedAt field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedAtOk

`func (o *OutputCaseTemplate) GetUnderscoreUpdatedAtOk() (*int64, bool)`

GetUnderscoreUpdatedAtOk returns a tuple with the UnderscoreUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedAt

`func (o *OutputCaseTemplate) SetUnderscoreUpdatedAt(v int64)`

SetUnderscoreUpdatedAt sets UnderscoreUpdatedAt field to given value.

### HasUnderscoreUpdatedAt

`func (o *OutputCaseTemplate) HasUnderscoreUpdatedAt() bool`

HasUnderscoreUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *OutputCaseTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OutputCaseTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OutputCaseTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *OutputCaseTemplate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *OutputCaseTemplate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *OutputCaseTemplate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetTitlePrefix

`func (o *OutputCaseTemplate) GetTitlePrefix() string`

GetTitlePrefix returns the TitlePrefix field if non-nil, zero value otherwise.

### GetTitlePrefixOk

`func (o *OutputCaseTemplate) GetTitlePrefixOk() (*string, bool)`

GetTitlePrefixOk returns a tuple with the TitlePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitlePrefix

`func (o *OutputCaseTemplate) SetTitlePrefix(v string)`

SetTitlePrefix sets TitlePrefix field to given value.

### HasTitlePrefix

`func (o *OutputCaseTemplate) HasTitlePrefix() bool`

HasTitlePrefix returns a boolean if a field has been set.

### GetDescription

`func (o *OutputCaseTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OutputCaseTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OutputCaseTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OutputCaseTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSeverity

`func (o *OutputCaseTemplate) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *OutputCaseTemplate) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *OutputCaseTemplate) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *OutputCaseTemplate) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSeverityLabel

`func (o *OutputCaseTemplate) GetSeverityLabel() string`

GetSeverityLabel returns the SeverityLabel field if non-nil, zero value otherwise.

### GetSeverityLabelOk

`func (o *OutputCaseTemplate) GetSeverityLabelOk() (*string, bool)`

GetSeverityLabelOk returns a tuple with the SeverityLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityLabel

`func (o *OutputCaseTemplate) SetSeverityLabel(v string)`

SetSeverityLabel sets SeverityLabel field to given value.

### HasSeverityLabel

`func (o *OutputCaseTemplate) HasSeverityLabel() bool`

HasSeverityLabel returns a boolean if a field has been set.

### GetTags

`func (o *OutputCaseTemplate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OutputCaseTemplate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OutputCaseTemplate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OutputCaseTemplate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetFlag

`func (o *OutputCaseTemplate) GetFlag() bool`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *OutputCaseTemplate) GetFlagOk() (*bool, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *OutputCaseTemplate) SetFlag(v bool)`

SetFlag sets Flag field to given value.


### GetTlp

`func (o *OutputCaseTemplate) GetTlp() int32`

GetTlp returns the Tlp field if non-nil, zero value otherwise.

### GetTlpOk

`func (o *OutputCaseTemplate) GetTlpOk() (*int32, bool)`

GetTlpOk returns a tuple with the Tlp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlp

`func (o *OutputCaseTemplate) SetTlp(v int32)`

SetTlp sets Tlp field to given value.

### HasTlp

`func (o *OutputCaseTemplate) HasTlp() bool`

HasTlp returns a boolean if a field has been set.

### GetTlpLabel

`func (o *OutputCaseTemplate) GetTlpLabel() string`

GetTlpLabel returns the TlpLabel field if non-nil, zero value otherwise.

### GetTlpLabelOk

`func (o *OutputCaseTemplate) GetTlpLabelOk() (*string, bool)`

GetTlpLabelOk returns a tuple with the TlpLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlpLabel

`func (o *OutputCaseTemplate) SetTlpLabel(v string)`

SetTlpLabel sets TlpLabel field to given value.

### HasTlpLabel

`func (o *OutputCaseTemplate) HasTlpLabel() bool`

HasTlpLabel returns a boolean if a field has been set.

### GetPap

`func (o *OutputCaseTemplate) GetPap() int32`

GetPap returns the Pap field if non-nil, zero value otherwise.

### GetPapOk

`func (o *OutputCaseTemplate) GetPapOk() (*int32, bool)`

GetPapOk returns a tuple with the Pap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPap

`func (o *OutputCaseTemplate) SetPap(v int32)`

SetPap sets Pap field to given value.

### HasPap

`func (o *OutputCaseTemplate) HasPap() bool`

HasPap returns a boolean if a field has been set.

### GetPapLabel

`func (o *OutputCaseTemplate) GetPapLabel() string`

GetPapLabel returns the PapLabel field if non-nil, zero value otherwise.

### GetPapLabelOk

`func (o *OutputCaseTemplate) GetPapLabelOk() (*string, bool)`

GetPapLabelOk returns a tuple with the PapLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPapLabel

`func (o *OutputCaseTemplate) SetPapLabel(v string)`

SetPapLabel sets PapLabel field to given value.

### HasPapLabel

`func (o *OutputCaseTemplate) HasPapLabel() bool`

HasPapLabel returns a boolean if a field has been set.

### GetSummary

`func (o *OutputCaseTemplate) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *OutputCaseTemplate) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *OutputCaseTemplate) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *OutputCaseTemplate) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetCustomFields

`func (o *OutputCaseTemplate) GetCustomFields() []OutputCustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *OutputCaseTemplate) GetCustomFieldsOk() (*[]OutputCustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *OutputCaseTemplate) SetCustomFields(v []OutputCustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *OutputCaseTemplate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetTasks

`func (o *OutputCaseTemplate) GetTasks() []OutputTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *OutputCaseTemplate) GetTasksOk() (*[]OutputTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *OutputCaseTemplate) SetTasks(v []OutputTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *OutputCaseTemplate) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetExtraData

`func (o *OutputCaseTemplate) GetExtraData() map[string]interface{}`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *OutputCaseTemplate) GetExtraDataOk() (*map[string]interface{}, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *OutputCaseTemplate) SetExtraData(v map[string]interface{})`

SetExtraData sets ExtraData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


