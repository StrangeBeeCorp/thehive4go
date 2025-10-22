# SSLParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientAuth** | Pointer to **string** |  | [optional] [default to "Default"]
**Protocols** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSSLParameters

`func NewSSLParameters() *SSLParameters`

NewSSLParameters instantiates a new SSLParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSLParametersWithDefaults

`func NewSSLParametersWithDefaults() *SSLParameters`

NewSSLParametersWithDefaults instantiates a new SSLParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientAuth

`func (o *SSLParameters) GetClientAuth() string`

GetClientAuth returns the ClientAuth field if non-nil, zero value otherwise.

### GetClientAuthOk

`func (o *SSLParameters) GetClientAuthOk() (*string, bool)`

GetClientAuthOk returns a tuple with the ClientAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAuth

`func (o *SSLParameters) SetClientAuth(v string)`

SetClientAuth sets ClientAuth field to given value.

### HasClientAuth

`func (o *SSLParameters) HasClientAuth() bool`

HasClientAuth returns a boolean if a field has been set.

### GetProtocols

`func (o *SSLParameters) GetProtocols() []string`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *SSLParameters) GetProtocolsOk() (*[]string, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *SSLParameters) SetProtocols(v []string)`

SetProtocols sets Protocols field to given value.

### HasProtocols

`func (o *SSLParameters) HasProtocols() bool`

HasProtocols returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


