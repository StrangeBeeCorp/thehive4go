# InputUpdateOrganisation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | New name for the organization. | [optional]
**Description** | Pointer to **string** | New description for the organization. | [optional]
**TaskRule** | Pointer to **string** | Enum: &#x60;autoShare&#x60; &#x60;manual&#x60;&lt;br&gt;New default task-sharing rule applied to new case shares with this organization: &#x60;autoShare&#x60; shares new tasks with every sharing organization automatically, &#x60;manual&#x60; requires explicit sharing per task. See [organization-level sharing rules](https://docs.strangebee.com/thehive/administration/organizations/about-organizations-sharing-rules/). | [optional] [default to "manual"]
**ObservableRule** | Pointer to **string** | Enum: &#x60;autoShare&#x60; &#x60;manual&#x60;&lt;br&gt;New default observable-sharing rule applied to new case shares with this organization: &#x60;autoShare&#x60; shares new observables with every sharing organization automatically, &#x60;manual&#x60; requires explicit sharing per observable. See [organization-level sharing rules](https://docs.strangebee.com/thehive/administration/organizations/about-organizations-sharing-rules/). | [optional] [default to "manual"]
**Locked** | Pointer to **bool** | Whether to lock or unlock the organization. | [optional]
**Avatar** | Pointer to **string** | New avatar for the organization, encoded as a Base64 string. | [optional]

## Methods

### NewInputUpdateOrganisation

`func NewInputUpdateOrganisation() *InputUpdateOrganisation`

NewInputUpdateOrganisation instantiates a new InputUpdateOrganisation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputUpdateOrganisationWithDefaults

`func NewInputUpdateOrganisationWithDefaults() *InputUpdateOrganisation`

NewInputUpdateOrganisationWithDefaults instantiates a new InputUpdateOrganisation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InputUpdateOrganisation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputUpdateOrganisation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputUpdateOrganisation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InputUpdateOrganisation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InputUpdateOrganisation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputUpdateOrganisation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputUpdateOrganisation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InputUpdateOrganisation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTaskRule

`func (o *InputUpdateOrganisation) GetTaskRule() string`

GetTaskRule returns the TaskRule field if non-nil, zero value otherwise.

### GetTaskRuleOk

`func (o *InputUpdateOrganisation) GetTaskRuleOk() (*string, bool)`

GetTaskRuleOk returns a tuple with the TaskRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRule

`func (o *InputUpdateOrganisation) SetTaskRule(v string)`

SetTaskRule sets TaskRule field to given value.

### HasTaskRule

`func (o *InputUpdateOrganisation) HasTaskRule() bool`

HasTaskRule returns a boolean if a field has been set.

### GetObservableRule

`func (o *InputUpdateOrganisation) GetObservableRule() string`

GetObservableRule returns the ObservableRule field if non-nil, zero value otherwise.

### GetObservableRuleOk

`func (o *InputUpdateOrganisation) GetObservableRuleOk() (*string, bool)`

GetObservableRuleOk returns a tuple with the ObservableRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservableRule

`func (o *InputUpdateOrganisation) SetObservableRule(v string)`

SetObservableRule sets ObservableRule field to given value.

### HasObservableRule

`func (o *InputUpdateOrganisation) HasObservableRule() bool`

HasObservableRule returns a boolean if a field has been set.

### GetLocked

`func (o *InputUpdateOrganisation) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *InputUpdateOrganisation) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *InputUpdateOrganisation) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *InputUpdateOrganisation) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetAvatar

`func (o *InputUpdateOrganisation) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *InputUpdateOrganisation) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *InputUpdateOrganisation) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *InputUpdateOrganisation) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
