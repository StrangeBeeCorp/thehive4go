# OutputUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnderscoreId** | **string** |  | 
**UnderscoreCreatedBy** | **string** |  | 
**UnderscoreUpdatedBy** | Pointer to **string** |  | [optional] 
**UnderscoreCreatedAt** | **int32** |  | 
**UnderscoreUpdatedAt** | Pointer to **int32** |  | [optional] 
**Login** | **string** |  | 
**Name** | **string** |  | 
**Email** | Pointer to **string** |  | [optional] 
**HasKey** | **bool** |  | 
**HasPassword** | **bool** |  | 
**HasMFA** | **bool** |  | 
**Locked** | **bool** |  | 
**Profile** | **string** |  | 
**Permissions** | Pointer to **[]string** |  | [optional] 
**Organisation** | **string** |  | 
**Avatar** | Pointer to **string** |  | [optional] 
**Organisations** | Pointer to [**[]OutputOrganisationProfile**](OutputOrganisationProfile.md) |  | [optional] 
**Type** | **string** |  | 
**DefaultOrganisation** | Pointer to **string** |  | [optional] 
**ExtraData** | **map[string]interface{}** |  | 

## Methods

### NewOutputUser

`func NewOutputUser(underscoreId string, underscoreCreatedBy string, underscoreCreatedAt int32, login string, name string, hasKey bool, hasPassword bool, hasMFA bool, locked bool, profile string, organisation string, type_ string, extraData map[string]interface{}, ) *OutputUser`

NewOutputUser instantiates a new OutputUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputUserWithDefaults

`func NewOutputUserWithDefaults() *OutputUser`

NewOutputUserWithDefaults instantiates a new OutputUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnderscoreId

`func (o *OutputUser) GetUnderscoreId() string`

GetUnderscoreId returns the UnderscoreId field if non-nil, zero value otherwise.

### GetUnderscoreIdOk

`func (o *OutputUser) GetUnderscoreIdOk() (*string, bool)`

GetUnderscoreIdOk returns a tuple with the UnderscoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreId

`func (o *OutputUser) SetUnderscoreId(v string)`

SetUnderscoreId sets UnderscoreId field to given value.


### GetUnderscoreCreatedBy

`func (o *OutputUser) GetUnderscoreCreatedBy() string`

GetUnderscoreCreatedBy returns the UnderscoreCreatedBy field if non-nil, zero value otherwise.

### GetUnderscoreCreatedByOk

`func (o *OutputUser) GetUnderscoreCreatedByOk() (*string, bool)`

GetUnderscoreCreatedByOk returns a tuple with the UnderscoreCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedBy

`func (o *OutputUser) SetUnderscoreCreatedBy(v string)`

SetUnderscoreCreatedBy sets UnderscoreCreatedBy field to given value.


### GetUnderscoreUpdatedBy

`func (o *OutputUser) GetUnderscoreUpdatedBy() string`

GetUnderscoreUpdatedBy returns the UnderscoreUpdatedBy field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedByOk

`func (o *OutputUser) GetUnderscoreUpdatedByOk() (*string, bool)`

GetUnderscoreUpdatedByOk returns a tuple with the UnderscoreUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedBy

`func (o *OutputUser) SetUnderscoreUpdatedBy(v string)`

SetUnderscoreUpdatedBy sets UnderscoreUpdatedBy field to given value.

### HasUnderscoreUpdatedBy

`func (o *OutputUser) HasUnderscoreUpdatedBy() bool`

HasUnderscoreUpdatedBy returns a boolean if a field has been set.

### GetUnderscoreCreatedAt

`func (o *OutputUser) GetUnderscoreCreatedAt() int32`

GetUnderscoreCreatedAt returns the UnderscoreCreatedAt field if non-nil, zero value otherwise.

### GetUnderscoreCreatedAtOk

`func (o *OutputUser) GetUnderscoreCreatedAtOk() (*int32, bool)`

GetUnderscoreCreatedAtOk returns a tuple with the UnderscoreCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedAt

`func (o *OutputUser) SetUnderscoreCreatedAt(v int32)`

SetUnderscoreCreatedAt sets UnderscoreCreatedAt field to given value.


### GetUnderscoreUpdatedAt

`func (o *OutputUser) GetUnderscoreUpdatedAt() int32`

GetUnderscoreUpdatedAt returns the UnderscoreUpdatedAt field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedAtOk

`func (o *OutputUser) GetUnderscoreUpdatedAtOk() (*int32, bool)`

GetUnderscoreUpdatedAtOk returns a tuple with the UnderscoreUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedAt

`func (o *OutputUser) SetUnderscoreUpdatedAt(v int32)`

SetUnderscoreUpdatedAt sets UnderscoreUpdatedAt field to given value.

### HasUnderscoreUpdatedAt

`func (o *OutputUser) HasUnderscoreUpdatedAt() bool`

HasUnderscoreUpdatedAt returns a boolean if a field has been set.

### GetLogin

`func (o *OutputUser) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *OutputUser) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *OutputUser) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetName

`func (o *OutputUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OutputUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OutputUser) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *OutputUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OutputUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OutputUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OutputUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetHasKey

`func (o *OutputUser) GetHasKey() bool`

GetHasKey returns the HasKey field if non-nil, zero value otherwise.

### GetHasKeyOk

`func (o *OutputUser) GetHasKeyOk() (*bool, bool)`

GetHasKeyOk returns a tuple with the HasKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasKey

`func (o *OutputUser) SetHasKey(v bool)`

SetHasKey sets HasKey field to given value.


### GetHasPassword

`func (o *OutputUser) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *OutputUser) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *OutputUser) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.


### GetHasMFA

`func (o *OutputUser) GetHasMFA() bool`

GetHasMFA returns the HasMFA field if non-nil, zero value otherwise.

### GetHasMFAOk

`func (o *OutputUser) GetHasMFAOk() (*bool, bool)`

GetHasMFAOk returns a tuple with the HasMFA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMFA

`func (o *OutputUser) SetHasMFA(v bool)`

SetHasMFA sets HasMFA field to given value.


### GetLocked

`func (o *OutputUser) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *OutputUser) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *OutputUser) SetLocked(v bool)`

SetLocked sets Locked field to given value.


### GetProfile

`func (o *OutputUser) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *OutputUser) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *OutputUser) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetPermissions

`func (o *OutputUser) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *OutputUser) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *OutputUser) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *OutputUser) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetOrganisation

`func (o *OutputUser) GetOrganisation() string`

GetOrganisation returns the Organisation field if non-nil, zero value otherwise.

### GetOrganisationOk

`func (o *OutputUser) GetOrganisationOk() (*string, bool)`

GetOrganisationOk returns a tuple with the Organisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisation

`func (o *OutputUser) SetOrganisation(v string)`

SetOrganisation sets Organisation field to given value.


### GetAvatar

`func (o *OutputUser) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *OutputUser) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *OutputUser) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *OutputUser) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetOrganisations

`func (o *OutputUser) GetOrganisations() []OutputOrganisationProfile`

GetOrganisations returns the Organisations field if non-nil, zero value otherwise.

### GetOrganisationsOk

`func (o *OutputUser) GetOrganisationsOk() (*[]OutputOrganisationProfile, bool)`

GetOrganisationsOk returns a tuple with the Organisations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisations

`func (o *OutputUser) SetOrganisations(v []OutputOrganisationProfile)`

SetOrganisations sets Organisations field to given value.

### HasOrganisations

`func (o *OutputUser) HasOrganisations() bool`

HasOrganisations returns a boolean if a field has been set.

### GetType

`func (o *OutputUser) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OutputUser) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OutputUser) SetType(v string)`

SetType sets Type field to given value.


### GetDefaultOrganisation

`func (o *OutputUser) GetDefaultOrganisation() string`

GetDefaultOrganisation returns the DefaultOrganisation field if non-nil, zero value otherwise.

### GetDefaultOrganisationOk

`func (o *OutputUser) GetDefaultOrganisationOk() (*string, bool)`

GetDefaultOrganisationOk returns a tuple with the DefaultOrganisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOrganisation

`func (o *OutputUser) SetDefaultOrganisation(v string)`

SetDefaultOrganisation sets DefaultOrganisation field to given value.

### HasDefaultOrganisation

`func (o *OutputUser) HasDefaultOrganisation() bool`

HasDefaultOrganisation returns a boolean if a field has been set.

### GetExtraData

`func (o *OutputUser) GetExtraData() map[string]interface{}`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *OutputUser) GetExtraDataOk() (*map[string]interface{}, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *OutputUser) SetExtraData(v map[string]interface{})`

SetExtraData sets ExtraData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


