# SSLConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **bool** |  | [optional] [default to false]
**Protocol** | Pointer to **string** |  | [optional] [default to "TLSv1.2"]
**CheckRevocation** | Pointer to **bool** |  | [optional] 
**RevocationLists** | Pointer to **[]string** |  | [optional] 
**EnabledCipherSuites** | Pointer to **[]string** |  | [optional] 
**EnabledProtocols** | Pointer to **[]string** |  | [optional] 
**HostnameVerifierClass** | Pointer to **string** |  | [optional] [default to "NoopHostnameVerifier"]
**SecureRandom** | Pointer to **string** |  | [optional] 
**TrustManager** | Pointer to [**Manager**](Manager.md) |  | [optional] 
**KeyManager** | Pointer to [**Manager**](Manager.md) |  | [optional] 
**SslParametersConfig** | Pointer to [**SSLParameters**](SSLParameters.md) |  | [optional] 
**Debug** | Pointer to [**SSLDebugConfig**](SSLDebugConfig.md) |  | [optional] 
**Loose** | [**SSLLooseConfig**](SSLLooseConfig.md) |  | 

## Methods

### NewSSLConfig

`func NewSSLConfig(loose SSLLooseConfig, ) *SSLConfig`

NewSSLConfig instantiates a new SSLConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSLConfigWithDefaults

`func NewSSLConfigWithDefaults() *SSLConfig`

NewSSLConfigWithDefaults instantiates a new SSLConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *SSLConfig) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *SSLConfig) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *SSLConfig) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *SSLConfig) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetProtocol

`func (o *SSLConfig) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SSLConfig) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SSLConfig) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SSLConfig) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetCheckRevocation

`func (o *SSLConfig) GetCheckRevocation() bool`

GetCheckRevocation returns the CheckRevocation field if non-nil, zero value otherwise.

### GetCheckRevocationOk

`func (o *SSLConfig) GetCheckRevocationOk() (*bool, bool)`

GetCheckRevocationOk returns a tuple with the CheckRevocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckRevocation

`func (o *SSLConfig) SetCheckRevocation(v bool)`

SetCheckRevocation sets CheckRevocation field to given value.

### HasCheckRevocation

`func (o *SSLConfig) HasCheckRevocation() bool`

HasCheckRevocation returns a boolean if a field has been set.

### GetRevocationLists

`func (o *SSLConfig) GetRevocationLists() []string`

GetRevocationLists returns the RevocationLists field if non-nil, zero value otherwise.

### GetRevocationListsOk

`func (o *SSLConfig) GetRevocationListsOk() (*[]string, bool)`

GetRevocationListsOk returns a tuple with the RevocationLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationLists

`func (o *SSLConfig) SetRevocationLists(v []string)`

SetRevocationLists sets RevocationLists field to given value.

### HasRevocationLists

`func (o *SSLConfig) HasRevocationLists() bool`

HasRevocationLists returns a boolean if a field has been set.

### GetEnabledCipherSuites

`func (o *SSLConfig) GetEnabledCipherSuites() []string`

GetEnabledCipherSuites returns the EnabledCipherSuites field if non-nil, zero value otherwise.

### GetEnabledCipherSuitesOk

`func (o *SSLConfig) GetEnabledCipherSuitesOk() (*[]string, bool)`

GetEnabledCipherSuitesOk returns a tuple with the EnabledCipherSuites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledCipherSuites

`func (o *SSLConfig) SetEnabledCipherSuites(v []string)`

SetEnabledCipherSuites sets EnabledCipherSuites field to given value.

### HasEnabledCipherSuites

`func (o *SSLConfig) HasEnabledCipherSuites() bool`

HasEnabledCipherSuites returns a boolean if a field has been set.

### GetEnabledProtocols

`func (o *SSLConfig) GetEnabledProtocols() []string`

GetEnabledProtocols returns the EnabledProtocols field if non-nil, zero value otherwise.

### GetEnabledProtocolsOk

`func (o *SSLConfig) GetEnabledProtocolsOk() (*[]string, bool)`

GetEnabledProtocolsOk returns a tuple with the EnabledProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledProtocols

`func (o *SSLConfig) SetEnabledProtocols(v []string)`

SetEnabledProtocols sets EnabledProtocols field to given value.

### HasEnabledProtocols

`func (o *SSLConfig) HasEnabledProtocols() bool`

HasEnabledProtocols returns a boolean if a field has been set.

### GetHostnameVerifierClass

`func (o *SSLConfig) GetHostnameVerifierClass() string`

GetHostnameVerifierClass returns the HostnameVerifierClass field if non-nil, zero value otherwise.

### GetHostnameVerifierClassOk

`func (o *SSLConfig) GetHostnameVerifierClassOk() (*string, bool)`

GetHostnameVerifierClassOk returns a tuple with the HostnameVerifierClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameVerifierClass

`func (o *SSLConfig) SetHostnameVerifierClass(v string)`

SetHostnameVerifierClass sets HostnameVerifierClass field to given value.

### HasHostnameVerifierClass

`func (o *SSLConfig) HasHostnameVerifierClass() bool`

HasHostnameVerifierClass returns a boolean if a field has been set.

### GetSecureRandom

`func (o *SSLConfig) GetSecureRandom() string`

GetSecureRandom returns the SecureRandom field if non-nil, zero value otherwise.

### GetSecureRandomOk

`func (o *SSLConfig) GetSecureRandomOk() (*string, bool)`

GetSecureRandomOk returns a tuple with the SecureRandom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureRandom

`func (o *SSLConfig) SetSecureRandom(v string)`

SetSecureRandom sets SecureRandom field to given value.

### HasSecureRandom

`func (o *SSLConfig) HasSecureRandom() bool`

HasSecureRandom returns a boolean if a field has been set.

### GetTrustManager

`func (o *SSLConfig) GetTrustManager() Manager`

GetTrustManager returns the TrustManager field if non-nil, zero value otherwise.

### GetTrustManagerOk

`func (o *SSLConfig) GetTrustManagerOk() (*Manager, bool)`

GetTrustManagerOk returns a tuple with the TrustManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustManager

`func (o *SSLConfig) SetTrustManager(v Manager)`

SetTrustManager sets TrustManager field to given value.

### HasTrustManager

`func (o *SSLConfig) HasTrustManager() bool`

HasTrustManager returns a boolean if a field has been set.

### GetKeyManager

`func (o *SSLConfig) GetKeyManager() Manager`

GetKeyManager returns the KeyManager field if non-nil, zero value otherwise.

### GetKeyManagerOk

`func (o *SSLConfig) GetKeyManagerOk() (*Manager, bool)`

GetKeyManagerOk returns a tuple with the KeyManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyManager

`func (o *SSLConfig) SetKeyManager(v Manager)`

SetKeyManager sets KeyManager field to given value.

### HasKeyManager

`func (o *SSLConfig) HasKeyManager() bool`

HasKeyManager returns a boolean if a field has been set.

### GetSslParametersConfig

`func (o *SSLConfig) GetSslParametersConfig() SSLParameters`

GetSslParametersConfig returns the SslParametersConfig field if non-nil, zero value otherwise.

### GetSslParametersConfigOk

`func (o *SSLConfig) GetSslParametersConfigOk() (*SSLParameters, bool)`

GetSslParametersConfigOk returns a tuple with the SslParametersConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslParametersConfig

`func (o *SSLConfig) SetSslParametersConfig(v SSLParameters)`

SetSslParametersConfig sets SslParametersConfig field to given value.

### HasSslParametersConfig

`func (o *SSLConfig) HasSslParametersConfig() bool`

HasSslParametersConfig returns a boolean if a field has been set.

### GetDebug

`func (o *SSLConfig) GetDebug() SSLDebugConfig`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *SSLConfig) GetDebugOk() (*SSLDebugConfig, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *SSLConfig) SetDebug(v SSLDebugConfig)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *SSLConfig) HasDebug() bool`

HasDebug returns a boolean if a field has been set.

### GetLoose

`func (o *SSLConfig) GetLoose() SSLLooseConfig`

GetLoose returns the Loose field if non-nil, zero value otherwise.

### GetLooseOk

`func (o *SSLConfig) GetLooseOk() (*SSLLooseConfig, bool)`

GetLooseOk returns a tuple with the Loose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoose

`func (o *SSLConfig) SetLoose(v SSLLooseConfig)`

SetLoose sets Loose field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


