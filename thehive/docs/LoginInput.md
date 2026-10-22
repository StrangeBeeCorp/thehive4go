# LoginInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | **string** | Login used to sign in to TheHive. Typically an email address. If a default domain is configured in the [authentication settings](https://docs.strangebee.com/thehive/administration/authentication/configure-authentication), a plain username is also accepted. |
**Password** | **string** | Password associated with &#x60;user&#x60;. |
**Organisation** | Pointer to **string** | ID or name of the organization to log into. Must be an organization you belong to that isn&#39;t locked, or login fails. If omitted, TheHive uses your default organization when it isn&#39;t locked, otherwise an organization you belong to. | [optional]
**Code** | Pointer to **string** | Current code from your authenticator app. Required if multifactor authentication is enabled for &#x60;user&#x60;. | [optional]

## Methods

### NewLoginInput

`func NewLoginInput(user string, password string, ) *LoginInput`

NewLoginInput instantiates a new LoginInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginInputWithDefaults

`func NewLoginInputWithDefaults() *LoginInput`

NewLoginInputWithDefaults instantiates a new LoginInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *LoginInput) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LoginInput) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LoginInput) SetUser(v string)`

SetUser sets User field to given value.


### GetPassword

`func (o *LoginInput) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LoginInput) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LoginInput) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetOrganisation

`func (o *LoginInput) GetOrganisation() string`

GetOrganisation returns the Organisation field if non-nil, zero value otherwise.

### GetOrganisationOk

`func (o *LoginInput) GetOrganisationOk() (*string, bool)`

GetOrganisationOk returns a tuple with the Organisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisation

`func (o *LoginInput) SetOrganisation(v string)`

SetOrganisation sets Organisation field to given value.

### HasOrganisation

`func (o *LoginInput) HasOrganisation() bool`

HasOrganisation returns a boolean if a field has been set.

### GetCode

`func (o *LoginInput) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *LoginInput) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *LoginInput) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *LoginInput) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
