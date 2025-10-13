# OutputEmailIntakeImapProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Host** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Ssl** | Pointer to **bool** |  | [optional] 
**StartTLS** | Pointer to **bool** |  | [optional] 
**CheckServerIdentity** | Pointer to **bool** |  | [optional] 
**Certificates** | Pointer to [**[]OutputEmailIntakeSSLCert**](OutputEmailIntakeSSLCert.md) |  | [optional] 

## Methods

### NewOutputEmailIntakeImapProvider

`func NewOutputEmailIntakeImapProvider(name string, ) *OutputEmailIntakeImapProvider`

NewOutputEmailIntakeImapProvider instantiates a new OutputEmailIntakeImapProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputEmailIntakeImapProviderWithDefaults

`func NewOutputEmailIntakeImapProviderWithDefaults() *OutputEmailIntakeImapProvider`

NewOutputEmailIntakeImapProviderWithDefaults instantiates a new OutputEmailIntakeImapProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OutputEmailIntakeImapProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OutputEmailIntakeImapProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OutputEmailIntakeImapProvider) SetName(v string)`

SetName sets Name field to given value.


### GetHost

`func (o *OutputEmailIntakeImapProvider) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *OutputEmailIntakeImapProvider) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *OutputEmailIntakeImapProvider) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *OutputEmailIntakeImapProvider) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetProtocol

`func (o *OutputEmailIntakeImapProvider) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *OutputEmailIntakeImapProvider) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *OutputEmailIntakeImapProvider) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *OutputEmailIntakeImapProvider) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPort

`func (o *OutputEmailIntakeImapProvider) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OutputEmailIntakeImapProvider) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OutputEmailIntakeImapProvider) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *OutputEmailIntakeImapProvider) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSsl

`func (o *OutputEmailIntakeImapProvider) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *OutputEmailIntakeImapProvider) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *OutputEmailIntakeImapProvider) SetSsl(v bool)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *OutputEmailIntakeImapProvider) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetStartTLS

`func (o *OutputEmailIntakeImapProvider) GetStartTLS() bool`

GetStartTLS returns the StartTLS field if non-nil, zero value otherwise.

### GetStartTLSOk

`func (o *OutputEmailIntakeImapProvider) GetStartTLSOk() (*bool, bool)`

GetStartTLSOk returns a tuple with the StartTLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTLS

`func (o *OutputEmailIntakeImapProvider) SetStartTLS(v bool)`

SetStartTLS sets StartTLS field to given value.

### HasStartTLS

`func (o *OutputEmailIntakeImapProvider) HasStartTLS() bool`

HasStartTLS returns a boolean if a field has been set.

### GetCheckServerIdentity

`func (o *OutputEmailIntakeImapProvider) GetCheckServerIdentity() bool`

GetCheckServerIdentity returns the CheckServerIdentity field if non-nil, zero value otherwise.

### GetCheckServerIdentityOk

`func (o *OutputEmailIntakeImapProvider) GetCheckServerIdentityOk() (*bool, bool)`

GetCheckServerIdentityOk returns a tuple with the CheckServerIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckServerIdentity

`func (o *OutputEmailIntakeImapProvider) SetCheckServerIdentity(v bool)`

SetCheckServerIdentity sets CheckServerIdentity field to given value.

### HasCheckServerIdentity

`func (o *OutputEmailIntakeImapProvider) HasCheckServerIdentity() bool`

HasCheckServerIdentity returns a boolean if a field has been set.

### GetCertificates

`func (o *OutputEmailIntakeImapProvider) GetCertificates() []OutputEmailIntakeSSLCert`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *OutputEmailIntakeImapProvider) GetCertificatesOk() (*[]OutputEmailIntakeSSLCert, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *OutputEmailIntakeImapProvider) SetCertificates(v []OutputEmailIntakeSSLCert)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *OutputEmailIntakeImapProvider) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


