# OAuth2OutputAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | Client ID assigned by the OAuth 2.0 provider. |
**GrantType** | Pointer to [**GrantType**](GrantType.md) |  | [optional]
**TokenUrl** | **string** | OAuth 2.0 token endpoint. |
**Scope** | Pointer to **[]string** | Scopes requested from the authorization server. | [optional]
**ClientAuthenticationMethod** | [**ClientAuthenticationMethod**](ClientAuthenticationMethod.md) | Method used to send client credentials to the token endpoint. |
**TokenParameters** | Pointer to [**[]TokenParameter**](TokenParameter.md) | Additional parameters included in the token request. | [optional]
**Type** | **string** | Authentication type identifier. |

## Methods

### NewOAuth2OutputAuth

`func NewOAuth2OutputAuth(clientId string, tokenUrl string, clientAuthenticationMethod ClientAuthenticationMethod, type_ string, ) *OAuth2OutputAuth`

NewOAuth2OutputAuth instantiates a new OAuth2OutputAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2OutputAuthWithDefaults

`func NewOAuth2OutputAuthWithDefaults() *OAuth2OutputAuth`

NewOAuth2OutputAuthWithDefaults instantiates a new OAuth2OutputAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *OAuth2OutputAuth) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuth2OutputAuth) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuth2OutputAuth) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetGrantType

`func (o *OAuth2OutputAuth) GetGrantType() GrantType`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *OAuth2OutputAuth) GetGrantTypeOk() (*GrantType, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *OAuth2OutputAuth) SetGrantType(v GrantType)`

SetGrantType sets GrantType field to given value.

### HasGrantType

`func (o *OAuth2OutputAuth) HasGrantType() bool`

HasGrantType returns a boolean if a field has been set.

### GetTokenUrl

`func (o *OAuth2OutputAuth) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *OAuth2OutputAuth) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *OAuth2OutputAuth) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.


### GetScope

`func (o *OAuth2OutputAuth) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *OAuth2OutputAuth) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *OAuth2OutputAuth) SetScope(v []string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *OAuth2OutputAuth) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetClientAuthenticationMethod

`func (o *OAuth2OutputAuth) GetClientAuthenticationMethod() ClientAuthenticationMethod`

GetClientAuthenticationMethod returns the ClientAuthenticationMethod field if non-nil, zero value otherwise.

### GetClientAuthenticationMethodOk

`func (o *OAuth2OutputAuth) GetClientAuthenticationMethodOk() (*ClientAuthenticationMethod, bool)`

GetClientAuthenticationMethodOk returns a tuple with the ClientAuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAuthenticationMethod

`func (o *OAuth2OutputAuth) SetClientAuthenticationMethod(v ClientAuthenticationMethod)`

SetClientAuthenticationMethod sets ClientAuthenticationMethod field to given value.


### GetTokenParameters

`func (o *OAuth2OutputAuth) GetTokenParameters() []TokenParameter`

GetTokenParameters returns the TokenParameters field if non-nil, zero value otherwise.

### GetTokenParametersOk

`func (o *OAuth2OutputAuth) GetTokenParametersOk() (*[]TokenParameter, bool)`

GetTokenParametersOk returns a tuple with the TokenParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenParameters

`func (o *OAuth2OutputAuth) SetTokenParameters(v []TokenParameter)`

SetTokenParameters sets TokenParameters field to given value.

### HasTokenParameters

`func (o *OAuth2OutputAuth) HasTokenParameters() bool`

HasTokenParameters returns a boolean if a field has been set.

### GetType

`func (o *OAuth2OutputAuth) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OAuth2OutputAuth) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OAuth2OutputAuth) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
