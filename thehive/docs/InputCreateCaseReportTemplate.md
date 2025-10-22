# InputCreateCaseReportTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Group** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Definition** | [**CaseReportTemplateDefinition**](CaseReportTemplateDefinition.md) |  | 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewInputCreateCaseReportTemplate

`func NewInputCreateCaseReportTemplate(title string, definition CaseReportTemplateDefinition, ) *InputCreateCaseReportTemplate`

NewInputCreateCaseReportTemplate instantiates a new InputCreateCaseReportTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputCreateCaseReportTemplateWithDefaults

`func NewInputCreateCaseReportTemplateWithDefaults() *InputCreateCaseReportTemplate`

NewInputCreateCaseReportTemplateWithDefaults instantiates a new InputCreateCaseReportTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *InputCreateCaseReportTemplate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InputCreateCaseReportTemplate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InputCreateCaseReportTemplate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetGroup

`func (o *InputCreateCaseReportTemplate) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InputCreateCaseReportTemplate) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InputCreateCaseReportTemplate) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *InputCreateCaseReportTemplate) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetDescription

`func (o *InputCreateCaseReportTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputCreateCaseReportTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputCreateCaseReportTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InputCreateCaseReportTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefinition

`func (o *InputCreateCaseReportTemplate) GetDefinition() CaseReportTemplateDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *InputCreateCaseReportTemplate) GetDefinitionOk() (*CaseReportTemplateDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *InputCreateCaseReportTemplate) SetDefinition(v CaseReportTemplateDefinition)`

SetDefinition sets Definition field to given value.


### GetVersion

`func (o *InputCreateCaseReportTemplate) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InputCreateCaseReportTemplate) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InputCreateCaseReportTemplate) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InputCreateCaseReportTemplate) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


