# OutputExportFieldsMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldsByModel** | [**map[string][]ExportFieldDescription**](array.md) | Map of supported model name (&#x60;Case&#x60;, &#x60;Alert&#x60;, &#x60;User&#x60;, &#x60;Organisation&#x60;, &#x60;Procedure&#x60;, &#x60;Task&#x60;, &#x60;Observable&#x60;) to the list of fields available for export on that model |

## Methods

### NewOutputExportFieldsMapping

`func NewOutputExportFieldsMapping(fieldsByModel map[string][]ExportFieldDescription, ) *OutputExportFieldsMapping`

NewOutputExportFieldsMapping instantiates a new OutputExportFieldsMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputExportFieldsMappingWithDefaults

`func NewOutputExportFieldsMappingWithDefaults() *OutputExportFieldsMapping`

NewOutputExportFieldsMappingWithDefaults instantiates a new OutputExportFieldsMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldsByModel

`func (o *OutputExportFieldsMapping) GetFieldsByModel() map[string][]ExportFieldDescription`

GetFieldsByModel returns the FieldsByModel field if non-nil, zero value otherwise.

### GetFieldsByModelOk

`func (o *OutputExportFieldsMapping) GetFieldsByModelOk() (*map[string][]ExportFieldDescription, bool)`

GetFieldsByModelOk returns a tuple with the FieldsByModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsByModel

`func (o *OutputExportFieldsMapping) SetFieldsByModel(v map[string][]ExportFieldDescription)`

SetFieldsByModel sets FieldsByModel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
