# ExportFieldDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldPath** | **string** | Path of the field as it should be referenced in the &#x60;fields&#x60; parameter of the Export API endpoint. Nested fields use a dotted notation (for example, &#x60;attachment.name&#x60;) |
**FieldType** | **string** | Logical type of the field. Possible values are &#x60;string&#x60;, &#x60;string[]&#x60;, &#x60;number&#x60;, &#x60;boolean&#x60;, &#x60;date&#x60;, &#x60;user&#x60;, &#x60;customFields&#x60;, &#x60;organisationLink[]&#x60;, and &#x60;organisationProfile[]&#x60; |
**SelectedByDefault** | **bool** | Whether the field is part of the default selection when the &#x60;fields&#x60; parameter is omitted in a CSV export |

## Methods

### NewExportFieldDescription

`func NewExportFieldDescription(fieldPath string, fieldType string, selectedByDefault bool, ) *ExportFieldDescription`

NewExportFieldDescription instantiates a new ExportFieldDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportFieldDescriptionWithDefaults

`func NewExportFieldDescriptionWithDefaults() *ExportFieldDescription`

NewExportFieldDescriptionWithDefaults instantiates a new ExportFieldDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldPath

`func (o *ExportFieldDescription) GetFieldPath() string`

GetFieldPath returns the FieldPath field if non-nil, zero value otherwise.

### GetFieldPathOk

`func (o *ExportFieldDescription) GetFieldPathOk() (*string, bool)`

GetFieldPathOk returns a tuple with the FieldPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldPath

`func (o *ExportFieldDescription) SetFieldPath(v string)`

SetFieldPath sets FieldPath field to given value.


### GetFieldType

`func (o *ExportFieldDescription) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *ExportFieldDescription) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *ExportFieldDescription) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.


### GetSelectedByDefault

`func (o *ExportFieldDescription) GetSelectedByDefault() bool`

GetSelectedByDefault returns the SelectedByDefault field if non-nil, zero value otherwise.

### GetSelectedByDefaultOk

`func (o *ExportFieldDescription) GetSelectedByDefaultOk() (*bool, bool)`

GetSelectedByDefaultOk returns a tuple with the SelectedByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedByDefault

`func (o *ExportFieldDescription) SetSelectedByDefault(v bool)`

SetSelectedByDefault sets SelectedByDefault field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
