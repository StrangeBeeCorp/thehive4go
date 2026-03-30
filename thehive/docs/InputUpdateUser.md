# InputUpdateUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Organisation** | Pointer to **string** | When updating the profile, use this field to determine the organisation to update | [optional] 
**Profile** | Pointer to **string** | Requires admin permission to update this field | [optional] 
**Locked** | Pointer to **bool** | Requires admin permission to update this field | [optional] 
**Avatar** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**DefaultOrganisation** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**UserType**](UserType.md) | Normal or Service. A service user cannot use the ui and needs to provide an api key for authentication | [optional] 

## Methods

### NewInputUpdateUser

`func NewInputUpdateUser() *InputUpdateUser`

NewInputUpdateUser instantiates a new InputUpdateUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputUpdateUserWithDefaults

`func NewInputUpdateUserWithDefaults() *InputUpdateUser`

NewInputUpdateUserWithDefaults instantiates a new InputUpdateUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InputUpdateUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputUpdateUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputUpdateUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InputUpdateUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganisation

`func (o *InputUpdateUser) GetOrganisation() string`

GetOrganisation returns the Organisation field if non-nil, zero value otherwise.

### GetOrganisationOk

`func (o *InputUpdateUser) GetOrganisationOk() (*string, bool)`

GetOrganisationOk returns a tuple with the Organisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisation

`func (o *InputUpdateUser) SetOrganisation(v string)`

SetOrganisation sets Organisation field to given value.

### HasOrganisation

`func (o *InputUpdateUser) HasOrganisation() bool`

HasOrganisation returns a boolean if a field has been set.

### GetProfile

`func (o *InputUpdateUser) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *InputUpdateUser) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *InputUpdateUser) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *InputUpdateUser) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetLocked

`func (o *InputUpdateUser) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *InputUpdateUser) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *InputUpdateUser) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *InputUpdateUser) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetAvatar

`func (o *InputUpdateUser) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *InputUpdateUser) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *InputUpdateUser) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *InputUpdateUser) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetEmail

`func (o *InputUpdateUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InputUpdateUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InputUpdateUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *InputUpdateUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDefaultOrganisation

`func (o *InputUpdateUser) GetDefaultOrganisation() string`

GetDefaultOrganisation returns the DefaultOrganisation field if non-nil, zero value otherwise.

### GetDefaultOrganisationOk

`func (o *InputUpdateUser) GetDefaultOrganisationOk() (*string, bool)`

GetDefaultOrganisationOk returns a tuple with the DefaultOrganisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOrganisation

`func (o *InputUpdateUser) SetDefaultOrganisation(v string)`

SetDefaultOrganisation sets DefaultOrganisation field to given value.

### HasDefaultOrganisation

`func (o *InputUpdateUser) HasDefaultOrganisation() bool`

HasDefaultOrganisation returns a boolean if a field has been set.

### GetType

`func (o *InputUpdateUser) GetType() UserType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputUpdateUser) GetTypeOk() (*UserType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputUpdateUser) SetType(v UserType)`

SetType sets Type field to given value.

### HasType

`func (o *InputUpdateUser) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


