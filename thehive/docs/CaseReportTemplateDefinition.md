# CaseReportTemplateDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Widgets** | Pointer to [**[]Widget**](Widget.md) |  | [optional] 
**Header** | Pointer to [**HeaderWidget**](HeaderWidget.md) |  | [optional] 
**Footer** | Pointer to [**FooterWidget**](FooterWidget.md) |  | [optional] 
**DateFormat** | Pointer to **string** |  | [optional] [default to "yyyy-MM-dd"]
**DateTimeFormat** | Pointer to **string** |  | [optional] [default to "yyyy-MM-dd HH:mm"]
**I18n** | [**I18n**](I18n.md) |  | 

## Methods

### NewCaseReportTemplateDefinition

`func NewCaseReportTemplateDefinition(i18n I18n, ) *CaseReportTemplateDefinition`

NewCaseReportTemplateDefinition instantiates a new CaseReportTemplateDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaseReportTemplateDefinitionWithDefaults

`func NewCaseReportTemplateDefinitionWithDefaults() *CaseReportTemplateDefinition`

NewCaseReportTemplateDefinitionWithDefaults instantiates a new CaseReportTemplateDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWidgets

`func (o *CaseReportTemplateDefinition) GetWidgets() []Widget`

GetWidgets returns the Widgets field if non-nil, zero value otherwise.

### GetWidgetsOk

`func (o *CaseReportTemplateDefinition) GetWidgetsOk() (*[]Widget, bool)`

GetWidgetsOk returns a tuple with the Widgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgets

`func (o *CaseReportTemplateDefinition) SetWidgets(v []Widget)`

SetWidgets sets Widgets field to given value.

### HasWidgets

`func (o *CaseReportTemplateDefinition) HasWidgets() bool`

HasWidgets returns a boolean if a field has been set.

### GetHeader

`func (o *CaseReportTemplateDefinition) GetHeader() HeaderWidget`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *CaseReportTemplateDefinition) GetHeaderOk() (*HeaderWidget, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *CaseReportTemplateDefinition) SetHeader(v HeaderWidget)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *CaseReportTemplateDefinition) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetFooter

`func (o *CaseReportTemplateDefinition) GetFooter() FooterWidget`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *CaseReportTemplateDefinition) GetFooterOk() (*FooterWidget, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *CaseReportTemplateDefinition) SetFooter(v FooterWidget)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *CaseReportTemplateDefinition) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetDateFormat

`func (o *CaseReportTemplateDefinition) GetDateFormat() string`

GetDateFormat returns the DateFormat field if non-nil, zero value otherwise.

### GetDateFormatOk

`func (o *CaseReportTemplateDefinition) GetDateFormatOk() (*string, bool)`

GetDateFormatOk returns a tuple with the DateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFormat

`func (o *CaseReportTemplateDefinition) SetDateFormat(v string)`

SetDateFormat sets DateFormat field to given value.

### HasDateFormat

`func (o *CaseReportTemplateDefinition) HasDateFormat() bool`

HasDateFormat returns a boolean if a field has been set.

### GetDateTimeFormat

`func (o *CaseReportTemplateDefinition) GetDateTimeFormat() string`

GetDateTimeFormat returns the DateTimeFormat field if non-nil, zero value otherwise.

### GetDateTimeFormatOk

`func (o *CaseReportTemplateDefinition) GetDateTimeFormatOk() (*string, bool)`

GetDateTimeFormatOk returns a tuple with the DateTimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFormat

`func (o *CaseReportTemplateDefinition) SetDateTimeFormat(v string)`

SetDateTimeFormat sets DateTimeFormat field to given value.

### HasDateTimeFormat

`func (o *CaseReportTemplateDefinition) HasDateTimeFormat() bool`

HasDateTimeFormat returns a boolean if a field has been set.

### GetI18n

`func (o *CaseReportTemplateDefinition) GetI18n() I18n`

GetI18n returns the I18n field if non-nil, zero value otherwise.

### GetI18nOk

`func (o *CaseReportTemplateDefinition) GetI18nOk() (*I18n, bool)`

GetI18nOk returns a tuple with the I18n field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI18n

`func (o *CaseReportTemplateDefinition) SetI18n(v I18n)`

SetI18n sets I18n field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


