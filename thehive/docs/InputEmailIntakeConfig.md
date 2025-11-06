# InputEmailIntakeConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional]
**Name** | **string** |  |
**Mailbox** | [**InputEmailIntakeMailboxConfig**](InputEmailIntakeMailboxConfig.md) |  |
**Organisations** | Pointer to **[]string** | At least one organisation | [optional]
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**CreatedAt** | Pointer to **string** |  | [optional]
**AlertProperties** | Pointer to [**InputEmailIntakeAlertProperties**](InputEmailIntakeAlertProperties.md) |  | [optional]

## Methods

### NewInputEmailIntakeConfig

`func NewInputEmailIntakeConfig(name string, mailbox InputEmailIntakeMailboxConfig, ) *InputEmailIntakeConfig`

NewInputEmailIntakeConfig instantiates a new InputEmailIntakeConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputEmailIntakeConfigWithDefaults

`func NewInputEmailIntakeConfigWithDefaults() *InputEmailIntakeConfig`

NewInputEmailIntakeConfigWithDefaults instantiates a new InputEmailIntakeConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InputEmailIntakeConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InputEmailIntakeConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InputEmailIntakeConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InputEmailIntakeConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InputEmailIntakeConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputEmailIntakeConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputEmailIntakeConfig) SetName(v string)`

SetName sets Name field to given value.


### GetMailbox

`func (o *InputEmailIntakeConfig) GetMailbox() InputEmailIntakeMailboxConfig`

GetMailbox returns the Mailbox field if non-nil, zero value otherwise.

### GetMailboxOk

`func (o *InputEmailIntakeConfig) GetMailboxOk() (*InputEmailIntakeMailboxConfig, bool)`

GetMailboxOk returns a tuple with the Mailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailbox

`func (o *InputEmailIntakeConfig) SetMailbox(v InputEmailIntakeMailboxConfig)`

SetMailbox sets Mailbox field to given value.


### GetOrganisations

`func (o *InputEmailIntakeConfig) GetOrganisations() []string`

GetOrganisations returns the Organisations field if non-nil, zero value otherwise.

### GetOrganisationsOk

`func (o *InputEmailIntakeConfig) GetOrganisationsOk() (*[]string, bool)`

GetOrganisationsOk returns a tuple with the Organisations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisations

`func (o *InputEmailIntakeConfig) SetOrganisations(v []string)`

SetOrganisations sets Organisations field to given value.

### HasOrganisations

`func (o *InputEmailIntakeConfig) HasOrganisations() bool`

HasOrganisations returns a boolean if a field has been set.

### GetEnabled

`func (o *InputEmailIntakeConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InputEmailIntakeConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InputEmailIntakeConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InputEmailIntakeConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InputEmailIntakeConfig) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InputEmailIntakeConfig) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InputEmailIntakeConfig) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InputEmailIntakeConfig) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetAlertProperties

`func (o *InputEmailIntakeConfig) GetAlertProperties() InputEmailIntakeAlertProperties`

GetAlertProperties returns the AlertProperties field if non-nil, zero value otherwise.

### GetAlertPropertiesOk

`func (o *InputEmailIntakeConfig) GetAlertPropertiesOk() (*InputEmailIntakeAlertProperties, bool)`

GetAlertPropertiesOk returns a tuple with the AlertProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertProperties

`func (o *InputEmailIntakeConfig) SetAlertProperties(v InputEmailIntakeAlertProperties)`

SetAlertProperties sets AlertProperties field to given value.

### HasAlertProperties

`func (o *InputEmailIntakeConfig) HasAlertProperties() bool`

HasAlertProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
