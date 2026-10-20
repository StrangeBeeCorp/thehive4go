# CreateAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | Username for HTTP Basic authentication. |
**Password** | **string** | Password for HTTP Basic authentication. |
**Type** | **string** | Authentication type identifier. Must be &#x60;basic&#x60;. |
**Key** | **string** | API key included in the authentication header. |
**Prefix** | Pointer to **string** | Optional prefix prepended to the API key in the header, separated by a space. | [optional]
**ClientId** | **string** | Client ID assigned by the OAuth 2.0 provider when registering TheHive. |
**ClientSecret** | **string** | Client secret issued by the OAuth 2.0 provider. |
**GrantType** | Pointer to [**GrantType**](GrantType.md) | Default value is &#39;client_credentials&#39; | [optional]
**TokenUrl** | **string** | OAuth 2.0 token endpoint used to obtain an access token. |
**Scope** | Pointer to **[]string** | Default value is empty array | [optional]
**ClientAuthenticationMethod** | [**ClientAuthenticationMethod**](ClientAuthenticationMethod.md) | Defines how client credentials are sent to the token endpoint. |
**TokenParameters** | Pointer to [**[]TokenParameter**](TokenParameter.md) | Additional parameters to include in the token request. | [optional]

## Methods

### NewCreateAuth

`func NewCreateAuth(username string, password string, type_ string, key string, clientId string, clientSecret string, tokenUrl string, clientAuthenticationMethod ClientAuthenticationMethod, ) *CreateAuth`

NewCreateAuth instantiates a new CreateAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAuthWithDefaults

`func NewCreateAuthWithDefaults() *CreateAuth`

NewCreateAuthWithDefaults instantiates a new CreateAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *CreateAuth) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateAuth) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateAuth) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *CreateAuth) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateAuth) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateAuth) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetType

`func (o *CreateAuth) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateAuth) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateAuth) SetType(v string)`

SetType sets Type field to given value.


### GetKey

`func (o *CreateAuth) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateAuth) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateAuth) SetKey(v string)`

SetKey sets Key field to given value.


### GetPrefix

`func (o *CreateAuth) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *CreateAuth) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *CreateAuth) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *CreateAuth) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetClientId

`func (o *CreateAuth) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CreateAuth) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CreateAuth) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *CreateAuth) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *CreateAuth) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *CreateAuth) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetGrantType

`func (o *CreateAuth) GetGrantType() GrantType`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *CreateAuth) GetGrantTypeOk() (*GrantType, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *CreateAuth) SetGrantType(v GrantType)`

SetGrantType sets GrantType field to given value.

### HasGrantType

`func (o *CreateAuth) HasGrantType() bool`

HasGrantType returns a boolean if a field has been set.

### GetTokenUrl

`func (o *CreateAuth) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *CreateAuth) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *CreateAuth) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.


### GetScope

`func (o *CreateAuth) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateAuth) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateAuth) SetScope(v []string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CreateAuth) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetClientAuthenticationMethod

`func (o *CreateAuth) GetClientAuthenticationMethod() ClientAuthenticationMethod`

GetClientAuthenticationMethod returns the ClientAuthenticationMethod field if non-nil, zero value otherwise.

### GetClientAuthenticationMethodOk

`func (o *CreateAuth) GetClientAuthenticationMethodOk() (*ClientAuthenticationMethod, bool)`

GetClientAuthenticationMethodOk returns a tuple with the ClientAuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAuthenticationMethod

`func (o *CreateAuth) SetClientAuthenticationMethod(v ClientAuthenticationMethod)`

SetClientAuthenticationMethod sets ClientAuthenticationMethod field to given value.


### GetTokenParameters

`func (o *CreateAuth) GetTokenParameters() []TokenParameter`

GetTokenParameters returns the TokenParameters field if non-nil, zero value otherwise.

### GetTokenParametersOk

`func (o *CreateAuth) GetTokenParametersOk() (*[]TokenParameter, bool)`

GetTokenParametersOk returns a tuple with the TokenParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenParameters

`func (o *CreateAuth) SetTokenParameters(v []TokenParameter)`

SetTokenParameters sets TokenParameters field to given value.

### HasTokenParameters

`func (o *CreateAuth) HasTokenParameters() bool`

HasTokenParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
