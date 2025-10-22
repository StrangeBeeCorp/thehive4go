# OkWithErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Imported** | Pointer to [**[]Imported**](Imported.md) |  | [optional] 
**Errors** | Pointer to [**[]Error**](Error.md) |  | [optional] 

## Methods

### NewOkWithErrors

`func NewOkWithErrors() *OkWithErrors`

NewOkWithErrors instantiates a new OkWithErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOkWithErrorsWithDefaults

`func NewOkWithErrorsWithDefaults() *OkWithErrors`

NewOkWithErrorsWithDefaults instantiates a new OkWithErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImported

`func (o *OkWithErrors) GetImported() []Imported`

GetImported returns the Imported field if non-nil, zero value otherwise.

### GetImportedOk

`func (o *OkWithErrors) GetImportedOk() (*[]Imported, bool)`

GetImportedOk returns a tuple with the Imported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImported

`func (o *OkWithErrors) SetImported(v []Imported)`

SetImported sets Imported field to given value.

### HasImported

`func (o *OkWithErrors) HasImported() bool`

HasImported returns a boolean if a field has been set.

### GetErrors

`func (o *OkWithErrors) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *OkWithErrors) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *OkWithErrors) SetErrors(v []Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *OkWithErrors) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


