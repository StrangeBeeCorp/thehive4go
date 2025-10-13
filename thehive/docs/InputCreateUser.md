# InputCreateUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | **string** | Login the user will need to enter, this will be the main identifier for the user | 
**Name** | **string** | Name of the user | 
**Email** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Profile** | **string** | Profile to assign the user in the organisation | 
**Organisation** | Pointer to **string** |  | [optional] 
**Avatar** | Pointer to ***os.File** | ignored in json body | [optional] 
**Type** | Pointer to [**InputUserType**](InputUserType.md) |  | [optional] 

## Methods

### NewInputCreateUser

`func NewInputCreateUser(login string, name string, profile string, ) *InputCreateUser`

NewInputCreateUser instantiates a new InputCreateUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputCreateUserWithDefaults

`func NewInputCreateUserWithDefaults() *InputCreateUser`

NewInputCreateUserWithDefaults instantiates a new InputCreateUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *InputCreateUser) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *InputCreateUser) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *InputCreateUser) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetName

`func (o *InputCreateUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputCreateUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputCreateUser) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *InputCreateUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InputCreateUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InputCreateUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *InputCreateUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPassword

`func (o *InputCreateUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InputCreateUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InputCreateUser) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *InputCreateUser) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetProfile

`func (o *InputCreateUser) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *InputCreateUser) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *InputCreateUser) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetOrganisation

`func (o *InputCreateUser) GetOrganisation() string`

GetOrganisation returns the Organisation field if non-nil, zero value otherwise.

### GetOrganisationOk

`func (o *InputCreateUser) GetOrganisationOk() (*string, bool)`

GetOrganisationOk returns a tuple with the Organisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisation

`func (o *InputCreateUser) SetOrganisation(v string)`

SetOrganisation sets Organisation field to given value.

### HasOrganisation

`func (o *InputCreateUser) HasOrganisation() bool`

HasOrganisation returns a boolean if a field has been set.

### GetAvatar

`func (o *InputCreateUser) GetAvatar() *os.File`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *InputCreateUser) GetAvatarOk() (**os.File, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *InputCreateUser) SetAvatar(v *os.File)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *InputCreateUser) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetType

`func (o *InputCreateUser) GetType() InputUserType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputCreateUser) GetTypeOk() (*InputUserType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputCreateUser) SetType(v InputUserType)`

SetType sets Type field to given value.

### HasType

`func (o *InputCreateUser) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


