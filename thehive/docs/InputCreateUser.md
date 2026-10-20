# InputCreateUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | **string** | Login name for the user account, used to sign in. Must be a valid email address. |
**Name** | **string** | Display name for the user account. |
**Email** | Pointer to **string** | Email address for sending notifications. | [optional]
**Password** | Pointer to **string** | Initial password for the user account. | [optional]
**Profile** | **string** | Permission profile for the user in the organization. Must match an existing profile. Use &#x60;GET /api/v1/profile&#x60; to list available profiles. |
**Organisation** | Pointer to **string** | Organization the user belongs to. Must match an existing organization. Use &#x60;POST /api/v1/query&#x60; with &#x60;listOrganisation&#x60; to retrieve available organizations. | [optional]
**Type** | Pointer to [**UserType**](UserType.md) | Type of user account. Normal accounts can use the web interface and the API. Service accounts authenticate exclusively with an API key and can&#39;t sign in to the web interface. External accounts access TheHive through the Portal (requires a Platinum or One license). | [optional]

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

### GetType

`func (o *InputCreateUser) GetType() UserType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputCreateUser) GetTypeOk() (*UserType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputCreateUser) SetType(v UserType)`

SetType sets Type field to given value.

### HasType

`func (o *InputCreateUser) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
