# InputCreateProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ProfileType**](ProfileType.md) | Profile type: - &#x60;Organisation&#x60;: for users in non-admin organizations. - &#x60;Admin&#x60;: for users in the Admin organization. - &#x60;External&#x60;: for [TheHive Portal](https://docs.strangebee.com/thehive/administration/thehive-portal/about-thehive-portal/) users. Requires a Platinum or One license.  If omitted, TheHive infers the type from the provided permissions. | [optional]
**Name** | **string** | Name of the profile. |
**Permissions** | Pointer to **[]string** | Permissions to assign to the profile. Use &#x60;GET /api/v1/permission&#x60; to list all available permissions. | [optional]

## Methods

### NewInputCreateProfile

`func NewInputCreateProfile(name string, ) *InputCreateProfile`

NewInputCreateProfile instantiates a new InputCreateProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputCreateProfileWithDefaults

`func NewInputCreateProfileWithDefaults() *InputCreateProfile`

NewInputCreateProfileWithDefaults instantiates a new InputCreateProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputCreateProfile) GetType() ProfileType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputCreateProfile) GetTypeOk() (*ProfileType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputCreateProfile) SetType(v ProfileType)`

SetType sets Type field to given value.

### HasType

`func (o *InputCreateProfile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *InputCreateProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputCreateProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputCreateProfile) SetName(v string)`

SetName sets Name field to given value.


### GetPermissions

`func (o *InputCreateProfile) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *InputCreateProfile) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *InputCreateProfile) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *InputCreateProfile) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
