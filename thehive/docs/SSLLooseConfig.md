# SSLLooseConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptAnyCertificate** | Pointer to **bool** |  | [optional] [default to false]
**AllowLegacyHelloMessages** | Pointer to **bool** |  | [optional] 
**AllowUnsafeRenegotiation** | Pointer to **bool** |  | [optional] 
**DisableHostnameVerification** | Pointer to **bool** |  | [optional] [default to false]
**DisableSNI** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewSSLLooseConfig

`func NewSSLLooseConfig() *SSLLooseConfig`

NewSSLLooseConfig instantiates a new SSLLooseConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSLLooseConfigWithDefaults

`func NewSSLLooseConfigWithDefaults() *SSLLooseConfig`

NewSSLLooseConfigWithDefaults instantiates a new SSLLooseConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptAnyCertificate

`func (o *SSLLooseConfig) GetAcceptAnyCertificate() bool`

GetAcceptAnyCertificate returns the AcceptAnyCertificate field if non-nil, zero value otherwise.

### GetAcceptAnyCertificateOk

`func (o *SSLLooseConfig) GetAcceptAnyCertificateOk() (*bool, bool)`

GetAcceptAnyCertificateOk returns a tuple with the AcceptAnyCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptAnyCertificate

`func (o *SSLLooseConfig) SetAcceptAnyCertificate(v bool)`

SetAcceptAnyCertificate sets AcceptAnyCertificate field to given value.

### HasAcceptAnyCertificate

`func (o *SSLLooseConfig) HasAcceptAnyCertificate() bool`

HasAcceptAnyCertificate returns a boolean if a field has been set.

### GetAllowLegacyHelloMessages

`func (o *SSLLooseConfig) GetAllowLegacyHelloMessages() bool`

GetAllowLegacyHelloMessages returns the AllowLegacyHelloMessages field if non-nil, zero value otherwise.

### GetAllowLegacyHelloMessagesOk

`func (o *SSLLooseConfig) GetAllowLegacyHelloMessagesOk() (*bool, bool)`

GetAllowLegacyHelloMessagesOk returns a tuple with the AllowLegacyHelloMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLegacyHelloMessages

`func (o *SSLLooseConfig) SetAllowLegacyHelloMessages(v bool)`

SetAllowLegacyHelloMessages sets AllowLegacyHelloMessages field to given value.

### HasAllowLegacyHelloMessages

`func (o *SSLLooseConfig) HasAllowLegacyHelloMessages() bool`

HasAllowLegacyHelloMessages returns a boolean if a field has been set.

### GetAllowUnsafeRenegotiation

`func (o *SSLLooseConfig) GetAllowUnsafeRenegotiation() bool`

GetAllowUnsafeRenegotiation returns the AllowUnsafeRenegotiation field if non-nil, zero value otherwise.

### GetAllowUnsafeRenegotiationOk

`func (o *SSLLooseConfig) GetAllowUnsafeRenegotiationOk() (*bool, bool)`

GetAllowUnsafeRenegotiationOk returns a tuple with the AllowUnsafeRenegotiation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUnsafeRenegotiation

`func (o *SSLLooseConfig) SetAllowUnsafeRenegotiation(v bool)`

SetAllowUnsafeRenegotiation sets AllowUnsafeRenegotiation field to given value.

### HasAllowUnsafeRenegotiation

`func (o *SSLLooseConfig) HasAllowUnsafeRenegotiation() bool`

HasAllowUnsafeRenegotiation returns a boolean if a field has been set.

### GetDisableHostnameVerification

`func (o *SSLLooseConfig) GetDisableHostnameVerification() bool`

GetDisableHostnameVerification returns the DisableHostnameVerification field if non-nil, zero value otherwise.

### GetDisableHostnameVerificationOk

`func (o *SSLLooseConfig) GetDisableHostnameVerificationOk() (*bool, bool)`

GetDisableHostnameVerificationOk returns a tuple with the DisableHostnameVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableHostnameVerification

`func (o *SSLLooseConfig) SetDisableHostnameVerification(v bool)`

SetDisableHostnameVerification sets DisableHostnameVerification field to given value.

### HasDisableHostnameVerification

`func (o *SSLLooseConfig) HasDisableHostnameVerification() bool`

HasDisableHostnameVerification returns a boolean if a field has been set.

### GetDisableSNI

`func (o *SSLLooseConfig) GetDisableSNI() bool`

GetDisableSNI returns the DisableSNI field if non-nil, zero value otherwise.

### GetDisableSNIOk

`func (o *SSLLooseConfig) GetDisableSNIOk() (*bool, bool)`

GetDisableSNIOk returns a tuple with the DisableSNI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSNI

`func (o *SSLLooseConfig) SetDisableSNI(v bool)`

SetDisableSNI sets DisableSNI field to given value.

### HasDisableSNI

`func (o *SSLLooseConfig) HasDisableSNI() bool`

HasDisableSNI returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


