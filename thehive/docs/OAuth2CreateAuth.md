# OAuth2CreateAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | Client ID assigned by the OAuth 2.0 provider when registering TheHive. |
**ClientSecret** | **string** | Client secret issued by the OAuth 2.0 provider. |
**GrantType** | Pointer to [**GrantType**](GrantType.md) | Default value is &#39;client_credentials&#39; | [optional]
**TokenUrl** | **string** | OAuth 2.0 token endpoint used to obtain an access token. |
**Scope** | Pointer to **[]string** | Default value is empty array | [optional]
**ClientAuthenticationMethod** | [**ClientAuthenticationMethod**](ClientAuthenticationMethod.md) | Defines how client credentials are sent to the token endpoint. |
**TokenParameters** | Pointer to [**[]TokenParameter**](TokenParameter.md) | Additional parameters to include in the token request. | [optional]
**Type** | **string** | Authentication type identifier. Must be &#x60;oauth2&#x60;. |

## Methods

### NewOAuth2CreateAuth

`func NewOAuth2CreateAuth(clientId string, clientSecret string, tokenUrl string, clientAuthenticationMethod ClientAuthenticationMethod, type_ string, ) *OAuth2CreateAuth`

NewOAuth2CreateAuth instantiates a new OAuth2CreateAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2CreateAuthWithDefaults

`func NewOAuth2CreateAuthWithDefaults() *OAuth2CreateAuth`

NewOAuth2CreateAuthWithDefaults instantiates a new OAuth2CreateAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *OAuth2CreateAuth) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuth2CreateAuth) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuth2CreateAuth) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *OAuth2CreateAuth) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OAuth2CreateAuth) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OAuth2CreateAuth) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetGrantType

`func (o *OAuth2CreateAuth) GetGrantType() GrantType`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *OAuth2CreateAuth) GetGrantTypeOk() (*GrantType, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *OAuth2CreateAuth) SetGrantType(v GrantType)`

SetGrantType sets GrantType field to given value.

### HasGrantType

`func (o *OAuth2CreateAuth) HasGrantType() bool`

HasGrantType returns a boolean if a field has been set.

### GetTokenUrl

`func (o *OAuth2CreateAuth) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *OAuth2CreateAuth) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *OAuth2CreateAuth) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.


### GetScope

`func (o *OAuth2CreateAuth) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *OAuth2CreateAuth) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *OAuth2CreateAuth) SetScope(v []string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *OAuth2CreateAuth) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetClientAuthenticationMethod

`func (o *OAuth2CreateAuth) GetClientAuthenticationMethod() ClientAuthenticationMethod`

GetClientAuthenticationMethod returns the ClientAuthenticationMethod field if non-nil, zero value otherwise.

### GetClientAuthenticationMethodOk

`func (o *OAuth2CreateAuth) GetClientAuthenticationMethodOk() (*ClientAuthenticationMethod, bool)`

GetClientAuthenticationMethodOk returns a tuple with the ClientAuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAuthenticationMethod

`func (o *OAuth2CreateAuth) SetClientAuthenticationMethod(v ClientAuthenticationMethod)`

SetClientAuthenticationMethod sets ClientAuthenticationMethod field to given value.


### GetTokenParameters

`func (o *OAuth2CreateAuth) GetTokenParameters() []TokenParameter`

GetTokenParameters returns the TokenParameters field if non-nil, zero value otherwise.

### GetTokenParametersOk

`func (o *OAuth2CreateAuth) GetTokenParametersOk() (*[]TokenParameter, bool)`

GetTokenParametersOk returns a tuple with the TokenParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenParameters

`func (o *OAuth2CreateAuth) SetTokenParameters(v []TokenParameter)`

SetTokenParameters sets TokenParameters field to given value.

### HasTokenParameters

`func (o *OAuth2CreateAuth) HasTokenParameters() bool`

HasTokenParameters returns a boolean if a field has been set.

### GetType

`func (o *OAuth2CreateAuth) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OAuth2CreateAuth) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OAuth2CreateAuth) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
