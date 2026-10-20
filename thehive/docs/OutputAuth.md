# OutputAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | Username for HTTP Basic authentication. |
**Password** | **string** | Password for HTTP Basic authentication. |
**Type** | **string** | Authentication type identifier. |
**Key** | **string** | API key included in the authentication header. |
**Prefix** | Pointer to **string** | Optional prefix prepended to the API key in the header, separated by a space. | [optional]
**ClientId** | **string** | Client ID assigned by the OAuth 2.0 provider. |
**GrantType** | Pointer to [**GrantType**](GrantType.md) |  | [optional]
**TokenUrl** | **string** | OAuth 2.0 token endpoint. |
**Scope** | Pointer to **[]string** | Scopes requested from the authorization server. | [optional]
**ClientAuthenticationMethod** | [**ClientAuthenticationMethod**](ClientAuthenticationMethod.md) | Method used to send client credentials to the token endpoint. |
**TokenParameters** | Pointer to [**[]TokenParameter**](TokenParameter.md) | Additional parameters included in the token request. | [optional]

## Methods

### NewOutputAuth

`func NewOutputAuth(username string, password string, type_ string, key string, clientId string, tokenUrl string, clientAuthenticationMethod ClientAuthenticationMethod, ) *OutputAuth`

NewOutputAuth instantiates a new OutputAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputAuthWithDefaults

`func NewOutputAuthWithDefaults() *OutputAuth`

NewOutputAuthWithDefaults instantiates a new OutputAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *OutputAuth) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *OutputAuth) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *OutputAuth) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *OutputAuth) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *OutputAuth) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *OutputAuth) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetType

`func (o *OutputAuth) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OutputAuth) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OutputAuth) SetType(v string)`

SetType sets Type field to given value.


### GetKey

`func (o *OutputAuth) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *OutputAuth) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *OutputAuth) SetKey(v string)`

SetKey sets Key field to given value.


### GetPrefix

`func (o *OutputAuth) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *OutputAuth) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *OutputAuth) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *OutputAuth) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetClientId

`func (o *OutputAuth) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OutputAuth) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OutputAuth) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetGrantType

`func (o *OutputAuth) GetGrantType() GrantType`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *OutputAuth) GetGrantTypeOk() (*GrantType, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *OutputAuth) SetGrantType(v GrantType)`

SetGrantType sets GrantType field to given value.

### HasGrantType

`func (o *OutputAuth) HasGrantType() bool`

HasGrantType returns a boolean if a field has been set.

### GetTokenUrl

`func (o *OutputAuth) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *OutputAuth) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *OutputAuth) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.


### GetScope

`func (o *OutputAuth) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *OutputAuth) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *OutputAuth) SetScope(v []string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *OutputAuth) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetClientAuthenticationMethod

`func (o *OutputAuth) GetClientAuthenticationMethod() ClientAuthenticationMethod`

GetClientAuthenticationMethod returns the ClientAuthenticationMethod field if non-nil, zero value otherwise.

### GetClientAuthenticationMethodOk

`func (o *OutputAuth) GetClientAuthenticationMethodOk() (*ClientAuthenticationMethod, bool)`

GetClientAuthenticationMethodOk returns a tuple with the ClientAuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAuthenticationMethod

`func (o *OutputAuth) SetClientAuthenticationMethod(v ClientAuthenticationMethod)`

SetClientAuthenticationMethod sets ClientAuthenticationMethod field to given value.


### GetTokenParameters

`func (o *OutputAuth) GetTokenParameters() []TokenParameter`

GetTokenParameters returns the TokenParameters field if non-nil, zero value otherwise.

### GetTokenParametersOk

`func (o *OutputAuth) GetTokenParametersOk() (*[]TokenParameter, bool)`

GetTokenParametersOk returns a tuple with the TokenParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenParameters

`func (o *OutputAuth) SetTokenParameters(v []TokenParameter)`

SetTokenParameters sets TokenParameters field to given value.

### HasTokenParameters

`func (o *OutputAuth) HasTokenParameters() bool`

HasTokenParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
