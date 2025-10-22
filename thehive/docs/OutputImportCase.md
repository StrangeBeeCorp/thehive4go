# OutputImportCase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Case** | [**OutputCase**](OutputCase.md) |  | 
**Observables** | Pointer to [**[]OutputObservable**](OutputObservable.md) |  | [optional] 
**Procedures** | Pointer to [**[]OutputProcedure**](OutputProcedure.md) |  | [optional] 
**Errors** | Pointer to **[]interface{}** | Field is present if there were some errors during import | [optional] 

## Methods

### NewOutputImportCase

`func NewOutputImportCase(case_ OutputCase, ) *OutputImportCase`

NewOutputImportCase instantiates a new OutputImportCase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputImportCaseWithDefaults

`func NewOutputImportCaseWithDefaults() *OutputImportCase`

NewOutputImportCaseWithDefaults instantiates a new OutputImportCase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCase

`func (o *OutputImportCase) GetCase() OutputCase`

GetCase returns the Case field if non-nil, zero value otherwise.

### GetCaseOk

`func (o *OutputImportCase) GetCaseOk() (*OutputCase, bool)`

GetCaseOk returns a tuple with the Case field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCase

`func (o *OutputImportCase) SetCase(v OutputCase)`

SetCase sets Case field to given value.


### GetObservables

`func (o *OutputImportCase) GetObservables() []OutputObservable`

GetObservables returns the Observables field if non-nil, zero value otherwise.

### GetObservablesOk

`func (o *OutputImportCase) GetObservablesOk() (*[]OutputObservable, bool)`

GetObservablesOk returns a tuple with the Observables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservables

`func (o *OutputImportCase) SetObservables(v []OutputObservable)`

SetObservables sets Observables field to given value.

### HasObservables

`func (o *OutputImportCase) HasObservables() bool`

HasObservables returns a boolean if a field has been set.

### GetProcedures

`func (o *OutputImportCase) GetProcedures() []OutputProcedure`

GetProcedures returns the Procedures field if non-nil, zero value otherwise.

### GetProceduresOk

`func (o *OutputImportCase) GetProceduresOk() (*[]OutputProcedure, bool)`

GetProceduresOk returns a tuple with the Procedures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcedures

`func (o *OutputImportCase) SetProcedures(v []OutputProcedure)`

SetProcedures sets Procedures field to given value.

### HasProcedures

`func (o *OutputImportCase) HasProcedures() bool`

HasProcedures returns a boolean if a field has been set.

### GetErrors

`func (o *OutputImportCase) GetErrors() []interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *OutputImportCase) GetErrorsOk() (*[]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *OutputImportCase) SetErrors(v []interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *OutputImportCase) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


