# ImportCaseReportTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Title of the created template. If absent, the title stored in the file is used. | [optional]
**Description** | Pointer to **string** | Description of the created template. If absent, the description stored in the file is used. | [optional]
**Group** | Pointer to **string** | Group label used to categorize the template. | [optional] [default to "default"]

## Methods

### NewImportCaseReportTemplate

`func NewImportCaseReportTemplate() *ImportCaseReportTemplate`

NewImportCaseReportTemplate instantiates a new ImportCaseReportTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportCaseReportTemplateWithDefaults

`func NewImportCaseReportTemplateWithDefaults() *ImportCaseReportTemplate`

NewImportCaseReportTemplateWithDefaults instantiates a new ImportCaseReportTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ImportCaseReportTemplate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ImportCaseReportTemplate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ImportCaseReportTemplate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ImportCaseReportTemplate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *ImportCaseReportTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ImportCaseReportTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ImportCaseReportTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ImportCaseReportTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroup

`func (o *ImportCaseReportTemplate) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ImportCaseReportTemplate) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ImportCaseReportTemplate) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ImportCaseReportTemplate) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
