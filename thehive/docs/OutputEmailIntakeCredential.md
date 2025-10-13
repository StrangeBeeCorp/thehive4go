# OutputEmailIntakeCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**BasicAuth** | Pointer to [**OutputEmailIntakeBasicAuth**](OutputEmailIntakeBasicAuth.md) |  | [optional] 
**OAuth2** | Pointer to [**OutputEmailIntakeOAuth2**](OutputEmailIntakeOAuth2.md) |  | [optional] 

## Methods

### NewOutputEmailIntakeCredential

`func NewOutputEmailIntakeCredential(email string, ) *OutputEmailIntakeCredential`

NewOutputEmailIntakeCredential instantiates a new OutputEmailIntakeCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputEmailIntakeCredentialWithDefaults

`func NewOutputEmailIntakeCredentialWithDefaults() *OutputEmailIntakeCredential`

NewOutputEmailIntakeCredentialWithDefaults instantiates a new OutputEmailIntakeCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *OutputEmailIntakeCredential) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OutputEmailIntakeCredential) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OutputEmailIntakeCredential) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetBasicAuth

`func (o *OutputEmailIntakeCredential) GetBasicAuth() OutputEmailIntakeBasicAuth`

GetBasicAuth returns the BasicAuth field if non-nil, zero value otherwise.

### GetBasicAuthOk

`func (o *OutputEmailIntakeCredential) GetBasicAuthOk() (*OutputEmailIntakeBasicAuth, bool)`

GetBasicAuthOk returns a tuple with the BasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuth

`func (o *OutputEmailIntakeCredential) SetBasicAuth(v OutputEmailIntakeBasicAuth)`

SetBasicAuth sets BasicAuth field to given value.

### HasBasicAuth

`func (o *OutputEmailIntakeCredential) HasBasicAuth() bool`

HasBasicAuth returns a boolean if a field has been set.

### GetOAuth2

`func (o *OutputEmailIntakeCredential) GetOAuth2() OutputEmailIntakeOAuth2`

GetOAuth2 returns the OAuth2 field if non-nil, zero value otherwise.

### GetOAuth2Ok

`func (o *OutputEmailIntakeCredential) GetOAuth2Ok() (*OutputEmailIntakeOAuth2, bool)`

GetOAuth2Ok returns a tuple with the OAuth2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuth2

`func (o *OutputEmailIntakeCredential) SetOAuth2(v OutputEmailIntakeOAuth2)`

SetOAuth2 sets OAuth2 field to given value.

### HasOAuth2

`func (o *OutputEmailIntakeCredential) HasOAuth2() bool`

HasOAuth2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


