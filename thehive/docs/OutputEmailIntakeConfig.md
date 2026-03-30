# OutputEmailIntakeConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Mailbox** | [**OutputEmailIntakeMailbox**](OutputEmailIntakeMailbox.md) |  | 
**Organisations** | Pointer to **[]string** |  | [optional] 
**Enabled** | **bool** |  | 
**CreatedAt** | **int64** |  | 
**AlertProperties** | [**OutputEmailIntakeAlertProperties**](OutputEmailIntakeAlertProperties.md) |  | 

## Methods

### NewOutputEmailIntakeConfig

`func NewOutputEmailIntakeConfig(id string, name string, mailbox OutputEmailIntakeMailbox, enabled bool, createdAt int64, alertProperties OutputEmailIntakeAlertProperties, ) *OutputEmailIntakeConfig`

NewOutputEmailIntakeConfig instantiates a new OutputEmailIntakeConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputEmailIntakeConfigWithDefaults

`func NewOutputEmailIntakeConfigWithDefaults() *OutputEmailIntakeConfig`

NewOutputEmailIntakeConfigWithDefaults instantiates a new OutputEmailIntakeConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OutputEmailIntakeConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutputEmailIntakeConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OutputEmailIntakeConfig) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *OutputEmailIntakeConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OutputEmailIntakeConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OutputEmailIntakeConfig) SetName(v string)`

SetName sets Name field to given value.


### GetMailbox

`func (o *OutputEmailIntakeConfig) GetMailbox() OutputEmailIntakeMailbox`

GetMailbox returns the Mailbox field if non-nil, zero value otherwise.

### GetMailboxOk

`func (o *OutputEmailIntakeConfig) GetMailboxOk() (*OutputEmailIntakeMailbox, bool)`

GetMailboxOk returns a tuple with the Mailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailbox

`func (o *OutputEmailIntakeConfig) SetMailbox(v OutputEmailIntakeMailbox)`

SetMailbox sets Mailbox field to given value.


### GetOrganisations

`func (o *OutputEmailIntakeConfig) GetOrganisations() []string`

GetOrganisations returns the Organisations field if non-nil, zero value otherwise.

### GetOrganisationsOk

`func (o *OutputEmailIntakeConfig) GetOrganisationsOk() (*[]string, bool)`

GetOrganisationsOk returns a tuple with the Organisations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisations

`func (o *OutputEmailIntakeConfig) SetOrganisations(v []string)`

SetOrganisations sets Organisations field to given value.

### HasOrganisations

`func (o *OutputEmailIntakeConfig) HasOrganisations() bool`

HasOrganisations returns a boolean if a field has been set.

### GetEnabled

`func (o *OutputEmailIntakeConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OutputEmailIntakeConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OutputEmailIntakeConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCreatedAt

`func (o *OutputEmailIntakeConfig) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OutputEmailIntakeConfig) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OutputEmailIntakeConfig) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetAlertProperties

`func (o *OutputEmailIntakeConfig) GetAlertProperties() OutputEmailIntakeAlertProperties`

GetAlertProperties returns the AlertProperties field if non-nil, zero value otherwise.

### GetAlertPropertiesOk

`func (o *OutputEmailIntakeConfig) GetAlertPropertiesOk() (*OutputEmailIntakeAlertProperties, bool)`

GetAlertPropertiesOk returns a tuple with the AlertProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertProperties

`func (o *OutputEmailIntakeConfig) SetAlertProperties(v OutputEmailIntakeAlertProperties)`

SetAlertProperties sets AlertProperties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


