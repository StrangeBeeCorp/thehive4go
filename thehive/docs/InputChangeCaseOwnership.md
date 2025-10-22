# InputChangeCaseOwnership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organisation** | **string** | Id or name of the new organisation that should own the case | 
**KeepProfile** | Pointer to **string** | Set a Profile to keep an access to the case, if none is set the current organisation will no longer have access to the case | [optional] 
**TaskRule** | Pointer to **string** |  | [optional] 
**ObservableRule** | Pointer to **string** |  | [optional] 

## Methods

### NewInputChangeCaseOwnership

`func NewInputChangeCaseOwnership(organisation string, ) *InputChangeCaseOwnership`

NewInputChangeCaseOwnership instantiates a new InputChangeCaseOwnership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputChangeCaseOwnershipWithDefaults

`func NewInputChangeCaseOwnershipWithDefaults() *InputChangeCaseOwnership`

NewInputChangeCaseOwnershipWithDefaults instantiates a new InputChangeCaseOwnership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganisation

`func (o *InputChangeCaseOwnership) GetOrganisation() string`

GetOrganisation returns the Organisation field if non-nil, zero value otherwise.

### GetOrganisationOk

`func (o *InputChangeCaseOwnership) GetOrganisationOk() (*string, bool)`

GetOrganisationOk returns a tuple with the Organisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisation

`func (o *InputChangeCaseOwnership) SetOrganisation(v string)`

SetOrganisation sets Organisation field to given value.


### GetKeepProfile

`func (o *InputChangeCaseOwnership) GetKeepProfile() string`

GetKeepProfile returns the KeepProfile field if non-nil, zero value otherwise.

### GetKeepProfileOk

`func (o *InputChangeCaseOwnership) GetKeepProfileOk() (*string, bool)`

GetKeepProfileOk returns a tuple with the KeepProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepProfile

`func (o *InputChangeCaseOwnership) SetKeepProfile(v string)`

SetKeepProfile sets KeepProfile field to given value.

### HasKeepProfile

`func (o *InputChangeCaseOwnership) HasKeepProfile() bool`

HasKeepProfile returns a boolean if a field has been set.

### GetTaskRule

`func (o *InputChangeCaseOwnership) GetTaskRule() string`

GetTaskRule returns the TaskRule field if non-nil, zero value otherwise.

### GetTaskRuleOk

`func (o *InputChangeCaseOwnership) GetTaskRuleOk() (*string, bool)`

GetTaskRuleOk returns a tuple with the TaskRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRule

`func (o *InputChangeCaseOwnership) SetTaskRule(v string)`

SetTaskRule sets TaskRule field to given value.

### HasTaskRule

`func (o *InputChangeCaseOwnership) HasTaskRule() bool`

HasTaskRule returns a boolean if a field has been set.

### GetObservableRule

`func (o *InputChangeCaseOwnership) GetObservableRule() string`

GetObservableRule returns the ObservableRule field if non-nil, zero value otherwise.

### GetObservableRuleOk

`func (o *InputChangeCaseOwnership) GetObservableRuleOk() (*string, bool)`

GetObservableRuleOk returns a tuple with the ObservableRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservableRule

`func (o *InputChangeCaseOwnership) SetObservableRule(v string)`

SetObservableRule sets ObservableRule field to given value.

### HasObservableRule

`func (o *InputChangeCaseOwnership) HasObservableRule() bool`

HasObservableRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


