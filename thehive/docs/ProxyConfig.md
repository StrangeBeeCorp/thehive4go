# ProxyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**State** | [**State**](State.md) |  | 
**Protocol** | Pointer to **string** |  | [optional] 
**Principal** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**NtlmDomain** | Pointer to **string** |  | [optional] 
**Encoding** | Pointer to **string** |  | [optional] 
**NonProxyHosts** | Pointer to **[]string** |  | [optional] 

## Methods

### NewProxyConfig

`func NewProxyConfig(state State, ) *ProxyConfig`

NewProxyConfig instantiates a new ProxyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyConfigWithDefaults

`func NewProxyConfigWithDefaults() *ProxyConfig`

NewProxyConfigWithDefaults instantiates a new ProxyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *ProxyConfig) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ProxyConfig) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ProxyConfig) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ProxyConfig) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *ProxyConfig) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ProxyConfig) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ProxyConfig) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ProxyConfig) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetState

`func (o *ProxyConfig) GetState() State`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ProxyConfig) GetStateOk() (*State, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ProxyConfig) SetState(v State)`

SetState sets State field to given value.


### GetProtocol

`func (o *ProxyConfig) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ProxyConfig) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ProxyConfig) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ProxyConfig) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPrincipal

`func (o *ProxyConfig) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *ProxyConfig) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *ProxyConfig) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *ProxyConfig) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetPassword

`func (o *ProxyConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ProxyConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ProxyConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ProxyConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetNtlmDomain

`func (o *ProxyConfig) GetNtlmDomain() string`

GetNtlmDomain returns the NtlmDomain field if non-nil, zero value otherwise.

### GetNtlmDomainOk

`func (o *ProxyConfig) GetNtlmDomainOk() (*string, bool)`

GetNtlmDomainOk returns a tuple with the NtlmDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtlmDomain

`func (o *ProxyConfig) SetNtlmDomain(v string)`

SetNtlmDomain sets NtlmDomain field to given value.

### HasNtlmDomain

`func (o *ProxyConfig) HasNtlmDomain() bool`

HasNtlmDomain returns a boolean if a field has been set.

### GetEncoding

`func (o *ProxyConfig) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *ProxyConfig) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *ProxyConfig) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *ProxyConfig) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetNonProxyHosts

`func (o *ProxyConfig) GetNonProxyHosts() []string`

GetNonProxyHosts returns the NonProxyHosts field if non-nil, zero value otherwise.

### GetNonProxyHostsOk

`func (o *ProxyConfig) GetNonProxyHostsOk() (*[]string, bool)`

GetNonProxyHostsOk returns a tuple with the NonProxyHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonProxyHosts

`func (o *ProxyConfig) SetNonProxyHosts(v []string)`

SetNonProxyHosts sets NonProxyHosts field to given value.

### HasNonProxyHosts

`func (o *ProxyConfig) HasNonProxyHosts() bool`

HasNonProxyHosts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


