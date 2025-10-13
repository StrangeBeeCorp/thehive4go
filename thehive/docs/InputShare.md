# InputShare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organisation** | **string** | Name or id of the organisation | 
**Share** | Pointer to **bool** | If true, new observables and tasks will also be shared to the organisation | [optional] [default to true]
**Profile** | Pointer to **string** | Profile used to define the permissions of the organisation members | [optional] [default to "analyst"]
**TaskRule** | Pointer to **string** | Sharing rule to apply to the tasks, can be manual or autoShare | [optional] [default to "Sharing rule applied on the case"]
**ObservableRule** | Pointer to **string** | Sharing rule to apply to the observables, can be manual or autoShare | [optional] [default to "Sharing rule applied on the case"]

## Methods

### NewInputShare

`func NewInputShare(organisation string, ) *InputShare`

NewInputShare instantiates a new InputShare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputShareWithDefaults

`func NewInputShareWithDefaults() *InputShare`

NewInputShareWithDefaults instantiates a new InputShare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganisation

`func (o *InputShare) GetOrganisation() string`

GetOrganisation returns the Organisation field if non-nil, zero value otherwise.

### GetOrganisationOk

`func (o *InputShare) GetOrganisationOk() (*string, bool)`

GetOrganisationOk returns a tuple with the Organisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisation

`func (o *InputShare) SetOrganisation(v string)`

SetOrganisation sets Organisation field to given value.


### GetShare

`func (o *InputShare) GetShare() bool`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *InputShare) GetShareOk() (*bool, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *InputShare) SetShare(v bool)`

SetShare sets Share field to given value.

### HasShare

`func (o *InputShare) HasShare() bool`

HasShare returns a boolean if a field has been set.

### GetProfile

`func (o *InputShare) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *InputShare) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *InputShare) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *InputShare) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetTaskRule

`func (o *InputShare) GetTaskRule() string`

GetTaskRule returns the TaskRule field if non-nil, zero value otherwise.

### GetTaskRuleOk

`func (o *InputShare) GetTaskRuleOk() (*string, bool)`

GetTaskRuleOk returns a tuple with the TaskRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRule

`func (o *InputShare) SetTaskRule(v string)`

SetTaskRule sets TaskRule field to given value.

### HasTaskRule

`func (o *InputShare) HasTaskRule() bool`

HasTaskRule returns a boolean if a field has been set.

### GetObservableRule

`func (o *InputShare) GetObservableRule() string`

GetObservableRule returns the ObservableRule field if non-nil, zero value otherwise.

### GetObservableRuleOk

`func (o *InputShare) GetObservableRuleOk() (*string, bool)`

GetObservableRuleOk returns a tuple with the ObservableRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservableRule

`func (o *InputShare) SetObservableRule(v string)`

SetObservableRule sets ObservableRule field to given value.

### HasObservableRule

`func (o *InputShare) HasObservableRule() bool`

HasObservableRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


