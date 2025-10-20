# OutputCaseReportTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnderscoreId** | **string** |  |
**UnderscoreType** | **string** |  |
**UnderscoreCreatedBy** | **string** |  |
**UnderscoreUpdatedBy** | Pointer to **string** |  | [optional]
**UnderscoreCreatedAt** | **int64** |  |
**UnderscoreUpdatedAt** | Pointer to **int64** |  | [optional]
**Title** | **string** |  |
**Group** | **string** |  |
**Description** | **string** |  |
**Version** | **int32** |  |
**Definition** | [**CaseReportTemplateDefinition**](CaseReportTemplateDefinition.md) |  |

## Methods

### NewOutputCaseReportTemplate

`func NewOutputCaseReportTemplate(underscoreId string, underscoreType string, underscoreCreatedBy string, underscoreCreatedAt int64, title string, group string, description string, version int32, definition CaseReportTemplateDefinition, ) *OutputCaseReportTemplate`

NewOutputCaseReportTemplate instantiates a new OutputCaseReportTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputCaseReportTemplateWithDefaults

`func NewOutputCaseReportTemplateWithDefaults() *OutputCaseReportTemplate`

NewOutputCaseReportTemplateWithDefaults instantiates a new OutputCaseReportTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnderscoreId

`func (o *OutputCaseReportTemplate) GetUnderscoreId() string`

GetUnderscoreId returns the UnderscoreId field if non-nil, zero value otherwise.

### GetUnderscoreIdOk

`func (o *OutputCaseReportTemplate) GetUnderscoreIdOk() (*string, bool)`

GetUnderscoreIdOk returns a tuple with the UnderscoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreId

`func (o *OutputCaseReportTemplate) SetUnderscoreId(v string)`

SetUnderscoreId sets UnderscoreId field to given value.


### GetUnderscoreType

`func (o *OutputCaseReportTemplate) GetUnderscoreType() string`

GetUnderscoreType returns the UnderscoreType field if non-nil, zero value otherwise.

### GetUnderscoreTypeOk

`func (o *OutputCaseReportTemplate) GetUnderscoreTypeOk() (*string, bool)`

GetUnderscoreTypeOk returns a tuple with the UnderscoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreType

`func (o *OutputCaseReportTemplate) SetUnderscoreType(v string)`

SetUnderscoreType sets UnderscoreType field to given value.


### GetUnderscoreCreatedBy

`func (o *OutputCaseReportTemplate) GetUnderscoreCreatedBy() string`

GetUnderscoreCreatedBy returns the UnderscoreCreatedBy field if non-nil, zero value otherwise.

### GetUnderscoreCreatedByOk

`func (o *OutputCaseReportTemplate) GetUnderscoreCreatedByOk() (*string, bool)`

GetUnderscoreCreatedByOk returns a tuple with the UnderscoreCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedBy

`func (o *OutputCaseReportTemplate) SetUnderscoreCreatedBy(v string)`

SetUnderscoreCreatedBy sets UnderscoreCreatedBy field to given value.


### GetUnderscoreUpdatedBy

`func (o *OutputCaseReportTemplate) GetUnderscoreUpdatedBy() string`

GetUnderscoreUpdatedBy returns the UnderscoreUpdatedBy field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedByOk

`func (o *OutputCaseReportTemplate) GetUnderscoreUpdatedByOk() (*string, bool)`

GetUnderscoreUpdatedByOk returns a tuple with the UnderscoreUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedBy

`func (o *OutputCaseReportTemplate) SetUnderscoreUpdatedBy(v string)`

SetUnderscoreUpdatedBy sets UnderscoreUpdatedBy field to given value.

### HasUnderscoreUpdatedBy

`func (o *OutputCaseReportTemplate) HasUnderscoreUpdatedBy() bool`

HasUnderscoreUpdatedBy returns a boolean if a field has been set.

### GetUnderscoreCreatedAt

`func (o *OutputCaseReportTemplate) GetUnderscoreCreatedAt() int64`

GetUnderscoreCreatedAt returns the UnderscoreCreatedAt field if non-nil, zero value otherwise.

### GetUnderscoreCreatedAtOk

`func (o *OutputCaseReportTemplate) GetUnderscoreCreatedAtOk() (*int64, bool)`

GetUnderscoreCreatedAtOk returns a tuple with the UnderscoreCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedAt

`func (o *OutputCaseReportTemplate) SetUnderscoreCreatedAt(v int64)`

SetUnderscoreCreatedAt sets UnderscoreCreatedAt field to given value.


### GetUnderscoreUpdatedAt

`func (o *OutputCaseReportTemplate) GetUnderscoreUpdatedAt() int64`

GetUnderscoreUpdatedAt returns the UnderscoreUpdatedAt field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedAtOk

`func (o *OutputCaseReportTemplate) GetUnderscoreUpdatedAtOk() (*int64, bool)`

GetUnderscoreUpdatedAtOk returns a tuple with the UnderscoreUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedAt

`func (o *OutputCaseReportTemplate) SetUnderscoreUpdatedAt(v int64)`

SetUnderscoreUpdatedAt sets UnderscoreUpdatedAt field to given value.

### HasUnderscoreUpdatedAt

`func (o *OutputCaseReportTemplate) HasUnderscoreUpdatedAt() bool`

HasUnderscoreUpdatedAt returns a boolean if a field has been set.

### GetTitle

`func (o *OutputCaseReportTemplate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OutputCaseReportTemplate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OutputCaseReportTemplate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetGroup

`func (o *OutputCaseReportTemplate) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *OutputCaseReportTemplate) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *OutputCaseReportTemplate) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetDescription

`func (o *OutputCaseReportTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OutputCaseReportTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OutputCaseReportTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetVersion

`func (o *OutputCaseReportTemplate) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OutputCaseReportTemplate) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OutputCaseReportTemplate) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetDefinition

`func (o *OutputCaseReportTemplate) GetDefinition() CaseReportTemplateDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *OutputCaseReportTemplate) GetDefinitionOk() (*CaseReportTemplateDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *OutputCaseReportTemplate) SetDefinition(v CaseReportTemplateDefinition)`

SetDefinition sets Definition field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
