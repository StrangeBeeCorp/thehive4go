# InputEmailIntakeCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The e-mail of the user account | 
**BasicAuth** | Pointer to [**InputEmailBasicAuthConfig**](InputEmailBasicAuthConfig.md) |  | [optional] 
**OAuth2** | Pointer to [**InputEmailIntakeOAuth2**](InputEmailIntakeOAuth2.md) |  | [optional] 

## Methods

### NewInputEmailIntakeCredential

`func NewInputEmailIntakeCredential(email string, ) *InputEmailIntakeCredential`

NewInputEmailIntakeCredential instantiates a new InputEmailIntakeCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputEmailIntakeCredentialWithDefaults

`func NewInputEmailIntakeCredentialWithDefaults() *InputEmailIntakeCredential`

NewInputEmailIntakeCredentialWithDefaults instantiates a new InputEmailIntakeCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InputEmailIntakeCredential) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InputEmailIntakeCredential) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InputEmailIntakeCredential) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetBasicAuth

`func (o *InputEmailIntakeCredential) GetBasicAuth() InputEmailBasicAuthConfig`

GetBasicAuth returns the BasicAuth field if non-nil, zero value otherwise.

### GetBasicAuthOk

`func (o *InputEmailIntakeCredential) GetBasicAuthOk() (*InputEmailBasicAuthConfig, bool)`

GetBasicAuthOk returns a tuple with the BasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuth

`func (o *InputEmailIntakeCredential) SetBasicAuth(v InputEmailBasicAuthConfig)`

SetBasicAuth sets BasicAuth field to given value.

### HasBasicAuth

`func (o *InputEmailIntakeCredential) HasBasicAuth() bool`

HasBasicAuth returns a boolean if a field has been set.

### GetOAuth2

`func (o *InputEmailIntakeCredential) GetOAuth2() InputEmailIntakeOAuth2`

GetOAuth2 returns the OAuth2 field if non-nil, zero value otherwise.

### GetOAuth2Ok

`func (o *InputEmailIntakeCredential) GetOAuth2Ok() (*InputEmailIntakeOAuth2, bool)`

GetOAuth2Ok returns a tuple with the OAuth2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuth2

`func (o *InputEmailIntakeCredential) SetOAuth2(v InputEmailIntakeOAuth2)`

SetOAuth2 sets OAuth2 field to given value.

### HasOAuth2

`func (o *InputEmailIntakeCredential) HasOAuth2() bool`

HasOAuth2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


