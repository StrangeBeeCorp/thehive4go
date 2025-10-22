# InputImportCase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** | Password to read the archive | 
**SharingParameters** | Pointer to [**[]InputShare**](InputShare.md) |  | [optional] 
**TaskRule** | Pointer to **string** |  | [optional] 
**ObservableRule** | Pointer to **string** |  | [optional] 

## Methods

### NewInputImportCase

`func NewInputImportCase(password string, ) *InputImportCase`

NewInputImportCase instantiates a new InputImportCase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputImportCaseWithDefaults

`func NewInputImportCaseWithDefaults() *InputImportCase`

NewInputImportCaseWithDefaults instantiates a new InputImportCase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *InputImportCase) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InputImportCase) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InputImportCase) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetSharingParameters

`func (o *InputImportCase) GetSharingParameters() []InputShare`

GetSharingParameters returns the SharingParameters field if non-nil, zero value otherwise.

### GetSharingParametersOk

`func (o *InputImportCase) GetSharingParametersOk() (*[]InputShare, bool)`

GetSharingParametersOk returns a tuple with the SharingParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingParameters

`func (o *InputImportCase) SetSharingParameters(v []InputShare)`

SetSharingParameters sets SharingParameters field to given value.

### HasSharingParameters

`func (o *InputImportCase) HasSharingParameters() bool`

HasSharingParameters returns a boolean if a field has been set.

### GetTaskRule

`func (o *InputImportCase) GetTaskRule() string`

GetTaskRule returns the TaskRule field if non-nil, zero value otherwise.

### GetTaskRuleOk

`func (o *InputImportCase) GetTaskRuleOk() (*string, bool)`

GetTaskRuleOk returns a tuple with the TaskRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRule

`func (o *InputImportCase) SetTaskRule(v string)`

SetTaskRule sets TaskRule field to given value.

### HasTaskRule

`func (o *InputImportCase) HasTaskRule() bool`

HasTaskRule returns a boolean if a field has been set.

### GetObservableRule

`func (o *InputImportCase) GetObservableRule() string`

GetObservableRule returns the ObservableRule field if non-nil, zero value otherwise.

### GetObservableRuleOk

`func (o *InputImportCase) GetObservableRuleOk() (*string, bool)`

GetObservableRuleOk returns a tuple with the ObservableRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservableRule

`func (o *InputImportCase) SetObservableRule(v string)`

SetObservableRule sets ObservableRule field to given value.

### HasObservableRule

`func (o *InputImportCase) HasObservableRule() bool`

HasObservableRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


