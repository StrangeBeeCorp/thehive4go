# OutputPublicStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sso** | **bool** |  | 
**SsoProviders** | Pointer to [**[]OutputSsoProvider**](OutputSsoProvider.md) |  | [optional] 
**Version** | **string** |  | 

## Methods

### NewOutputPublicStatus

`func NewOutputPublicStatus(sso bool, version string, ) *OutputPublicStatus`

NewOutputPublicStatus instantiates a new OutputPublicStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputPublicStatusWithDefaults

`func NewOutputPublicStatusWithDefaults() *OutputPublicStatus`

NewOutputPublicStatusWithDefaults instantiates a new OutputPublicStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSso

`func (o *OutputPublicStatus) GetSso() bool`

GetSso returns the Sso field if non-nil, zero value otherwise.

### GetSsoOk

`func (o *OutputPublicStatus) GetSsoOk() (*bool, bool)`

GetSsoOk returns a tuple with the Sso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSso

`func (o *OutputPublicStatus) SetSso(v bool)`

SetSso sets Sso field to given value.


### GetSsoProviders

`func (o *OutputPublicStatus) GetSsoProviders() []OutputSsoProvider`

GetSsoProviders returns the SsoProviders field if non-nil, zero value otherwise.

### GetSsoProvidersOk

`func (o *OutputPublicStatus) GetSsoProvidersOk() (*[]OutputSsoProvider, bool)`

GetSsoProvidersOk returns a tuple with the SsoProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoProviders

`func (o *OutputPublicStatus) SetSsoProviders(v []OutputSsoProvider)`

SetSsoProviders sets SsoProviders field to given value.

### HasSsoProviders

`func (o *OutputPublicStatus) HasSsoProviders() bool`

HasSsoProviders returns a boolean if a field has been set.

### GetVersion

`func (o *OutputPublicStatus) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OutputPublicStatus) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OutputPublicStatus) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


