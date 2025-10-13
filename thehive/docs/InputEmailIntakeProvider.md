# InputEmailIntakeProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | [**EmailIntakeProviderName**](EmailIntakeProviderName.md) | Select an existing provider or define a custom one | 
**Host** | Pointer to **string** | The name of the &#39;imap&#39; server | [optional] 
**Protocol** | Pointer to **string** |  | [optional] [default to "imap"]
**Port** | Pointer to **int32** |  | [optional] [default to 993]
**Ssl** | Pointer to **bool** |  | [optional] [default to true]
**StartTLS** | Pointer to **bool** |  | [optional] [default to false]
**CheckServerIdentity** | Pointer to **bool** |  | [optional] [default to true]
**Certificates** | Pointer to [**[]InputEmailIntakeSSLCert**](InputEmailIntakeSSLCert.md) | *.pem files | [optional] 

## Methods

### NewInputEmailIntakeProvider

`func NewInputEmailIntakeProvider(name EmailIntakeProviderName, ) *InputEmailIntakeProvider`

NewInputEmailIntakeProvider instantiates a new InputEmailIntakeProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputEmailIntakeProviderWithDefaults

`func NewInputEmailIntakeProviderWithDefaults() *InputEmailIntakeProvider`

NewInputEmailIntakeProviderWithDefaults instantiates a new InputEmailIntakeProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InputEmailIntakeProvider) GetName() EmailIntakeProviderName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputEmailIntakeProvider) GetNameOk() (*EmailIntakeProviderName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputEmailIntakeProvider) SetName(v EmailIntakeProviderName)`

SetName sets Name field to given value.


### GetHost

`func (o *InputEmailIntakeProvider) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *InputEmailIntakeProvider) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *InputEmailIntakeProvider) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *InputEmailIntakeProvider) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetProtocol

`func (o *InputEmailIntakeProvider) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InputEmailIntakeProvider) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InputEmailIntakeProvider) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InputEmailIntakeProvider) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPort

`func (o *InputEmailIntakeProvider) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *InputEmailIntakeProvider) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *InputEmailIntakeProvider) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *InputEmailIntakeProvider) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSsl

`func (o *InputEmailIntakeProvider) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *InputEmailIntakeProvider) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *InputEmailIntakeProvider) SetSsl(v bool)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *InputEmailIntakeProvider) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetStartTLS

`func (o *InputEmailIntakeProvider) GetStartTLS() bool`

GetStartTLS returns the StartTLS field if non-nil, zero value otherwise.

### GetStartTLSOk

`func (o *InputEmailIntakeProvider) GetStartTLSOk() (*bool, bool)`

GetStartTLSOk returns a tuple with the StartTLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTLS

`func (o *InputEmailIntakeProvider) SetStartTLS(v bool)`

SetStartTLS sets StartTLS field to given value.

### HasStartTLS

`func (o *InputEmailIntakeProvider) HasStartTLS() bool`

HasStartTLS returns a boolean if a field has been set.

### GetCheckServerIdentity

`func (o *InputEmailIntakeProvider) GetCheckServerIdentity() bool`

GetCheckServerIdentity returns the CheckServerIdentity field if non-nil, zero value otherwise.

### GetCheckServerIdentityOk

`func (o *InputEmailIntakeProvider) GetCheckServerIdentityOk() (*bool, bool)`

GetCheckServerIdentityOk returns a tuple with the CheckServerIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckServerIdentity

`func (o *InputEmailIntakeProvider) SetCheckServerIdentity(v bool)`

SetCheckServerIdentity sets CheckServerIdentity field to given value.

### HasCheckServerIdentity

`func (o *InputEmailIntakeProvider) HasCheckServerIdentity() bool`

HasCheckServerIdentity returns a boolean if a field has been set.

### GetCertificates

`func (o *InputEmailIntakeProvider) GetCertificates() []InputEmailIntakeSSLCert`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *InputEmailIntakeProvider) GetCertificatesOk() (*[]InputEmailIntakeSSLCert, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *InputEmailIntakeProvider) SetCertificates(v []InputEmailIntakeSSLCert)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *InputEmailIntakeProvider) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


