# PermissionDesc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Label** | **string** | The user-friendly name of the permission | 
**ConsumesLicense** | **bool** | &#x60;true&#x60; if the permission requires a licence, &#x60;false&#x60; otherwise | 
**Scope** | Pointer to **[]string** | The scope of the permission. Can be &#x60;admin&#x60;, &#x60;organisation&#x60; or both | [optional] 

## Methods

### NewPermissionDesc

`func NewPermissionDesc(name string, label string, consumesLicense bool, ) *PermissionDesc`

NewPermissionDesc instantiates a new PermissionDesc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionDescWithDefaults

`func NewPermissionDescWithDefaults() *PermissionDesc`

NewPermissionDescWithDefaults instantiates a new PermissionDesc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PermissionDesc) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PermissionDesc) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PermissionDesc) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *PermissionDesc) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PermissionDesc) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PermissionDesc) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetConsumesLicense

`func (o *PermissionDesc) GetConsumesLicense() bool`

GetConsumesLicense returns the ConsumesLicense field if non-nil, zero value otherwise.

### GetConsumesLicenseOk

`func (o *PermissionDesc) GetConsumesLicenseOk() (*bool, bool)`

GetConsumesLicenseOk returns a tuple with the ConsumesLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumesLicense

`func (o *PermissionDesc) SetConsumesLicense(v bool)`

SetConsumesLicense sets ConsumesLicense field to given value.


### GetScope

`func (o *PermissionDesc) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *PermissionDesc) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *PermissionDesc) SetScope(v []string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *PermissionDesc) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


