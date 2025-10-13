# InputRenderCaseReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | **string** |  | 
**Definition** | Pointer to [**CaseReportTemplateDefinition**](CaseReportTemplateDefinition.md) | Either provide &#x60;definition&#x60; or &#x60;caseReportTemplateId&#x60; | [optional] 
**CaseReportTemplateId** | Pointer to **string** | Either provide &#x60;definition&#x60; or &#x60;caseReportTemplateId&#x60; | [optional] 
**CaseId** | Pointer to **string** | The _id of the entity or its &#39;name&#39; (depends of the entity) | [optional] 
**MaxElements** | Pointer to **int32** | Limit the number of elements in tables and lists, useful for a preview | [optional] 

## Methods

### NewInputRenderCaseReport

`func NewInputRenderCaseReport(format string, ) *InputRenderCaseReport`

NewInputRenderCaseReport instantiates a new InputRenderCaseReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputRenderCaseReportWithDefaults

`func NewInputRenderCaseReportWithDefaults() *InputRenderCaseReport`

NewInputRenderCaseReportWithDefaults instantiates a new InputRenderCaseReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *InputRenderCaseReport) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *InputRenderCaseReport) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *InputRenderCaseReport) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetDefinition

`func (o *InputRenderCaseReport) GetDefinition() CaseReportTemplateDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *InputRenderCaseReport) GetDefinitionOk() (*CaseReportTemplateDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *InputRenderCaseReport) SetDefinition(v CaseReportTemplateDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *InputRenderCaseReport) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetCaseReportTemplateId

`func (o *InputRenderCaseReport) GetCaseReportTemplateId() string`

GetCaseReportTemplateId returns the CaseReportTemplateId field if non-nil, zero value otherwise.

### GetCaseReportTemplateIdOk

`func (o *InputRenderCaseReport) GetCaseReportTemplateIdOk() (*string, bool)`

GetCaseReportTemplateIdOk returns a tuple with the CaseReportTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseReportTemplateId

`func (o *InputRenderCaseReport) SetCaseReportTemplateId(v string)`

SetCaseReportTemplateId sets CaseReportTemplateId field to given value.

### HasCaseReportTemplateId

`func (o *InputRenderCaseReport) HasCaseReportTemplateId() bool`

HasCaseReportTemplateId returns a boolean if a field has been set.

### GetCaseId

`func (o *InputRenderCaseReport) GetCaseId() string`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *InputRenderCaseReport) GetCaseIdOk() (*string, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *InputRenderCaseReport) SetCaseId(v string)`

SetCaseId sets CaseId field to given value.

### HasCaseId

`func (o *InputRenderCaseReport) HasCaseId() bool`

HasCaseId returns a boolean if a field has been set.

### GetMaxElements

`func (o *InputRenderCaseReport) GetMaxElements() int32`

GetMaxElements returns the MaxElements field if non-nil, zero value otherwise.

### GetMaxElementsOk

`func (o *InputRenderCaseReport) GetMaxElementsOk() (*int32, bool)`

GetMaxElementsOk returns a tuple with the MaxElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxElements

`func (o *InputRenderCaseReport) SetMaxElements(v int32)`

SetMaxElements sets MaxElements field to given value.

### HasMaxElements

`func (o *InputRenderCaseReport) HasMaxElements() bool`

HasMaxElements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


