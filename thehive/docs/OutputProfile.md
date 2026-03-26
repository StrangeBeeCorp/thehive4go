# OutputProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnderscoreId** | **string** |  | 
**UnderscoreType** | **string** |  | 
**UnderscoreCreatedBy** | **string** |  | 
**UnderscoreUpdatedBy** | Pointer to **string** |  | [optional] 
**UnderscoreCreatedAt** | **int64** |  | 
**UnderscoreUpdatedAt** | Pointer to **int64** |  | [optional] 
**Type** | [**ProfileType**](ProfileType.md) |  | 
**Name** | **string** |  | 
**Permissions** | Pointer to **[]string** |  | [optional] 
**Editable** | **bool** |  | 
**ForAdmin** | **bool** |  | 
**ForOrg** | **bool** |  | 
**ForExternal** | **bool** |  | 
**ConsumesLicense** | **bool** |  | 

## Methods

### NewOutputProfile

`func NewOutputProfile(underscoreId string, underscoreType string, underscoreCreatedBy string, underscoreCreatedAt int64, type_ ProfileType, name string, editable bool, forAdmin bool, forOrg bool, forExternal bool, consumesLicense bool, ) *OutputProfile`

NewOutputProfile instantiates a new OutputProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputProfileWithDefaults

`func NewOutputProfileWithDefaults() *OutputProfile`

NewOutputProfileWithDefaults instantiates a new OutputProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnderscoreId

`func (o *OutputProfile) GetUnderscoreId() string`

GetUnderscoreId returns the UnderscoreId field if non-nil, zero value otherwise.

### GetUnderscoreIdOk

`func (o *OutputProfile) GetUnderscoreIdOk() (*string, bool)`

GetUnderscoreIdOk returns a tuple with the UnderscoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreId

`func (o *OutputProfile) SetUnderscoreId(v string)`

SetUnderscoreId sets UnderscoreId field to given value.


### GetUnderscoreType

`func (o *OutputProfile) GetUnderscoreType() string`

GetUnderscoreType returns the UnderscoreType field if non-nil, zero value otherwise.

### GetUnderscoreTypeOk

`func (o *OutputProfile) GetUnderscoreTypeOk() (*string, bool)`

GetUnderscoreTypeOk returns a tuple with the UnderscoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreType

`func (o *OutputProfile) SetUnderscoreType(v string)`

SetUnderscoreType sets UnderscoreType field to given value.


### GetUnderscoreCreatedBy

`func (o *OutputProfile) GetUnderscoreCreatedBy() string`

GetUnderscoreCreatedBy returns the UnderscoreCreatedBy field if non-nil, zero value otherwise.

### GetUnderscoreCreatedByOk

`func (o *OutputProfile) GetUnderscoreCreatedByOk() (*string, bool)`

GetUnderscoreCreatedByOk returns a tuple with the UnderscoreCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedBy

`func (o *OutputProfile) SetUnderscoreCreatedBy(v string)`

SetUnderscoreCreatedBy sets UnderscoreCreatedBy field to given value.


### GetUnderscoreUpdatedBy

`func (o *OutputProfile) GetUnderscoreUpdatedBy() string`

GetUnderscoreUpdatedBy returns the UnderscoreUpdatedBy field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedByOk

`func (o *OutputProfile) GetUnderscoreUpdatedByOk() (*string, bool)`

GetUnderscoreUpdatedByOk returns a tuple with the UnderscoreUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedBy

`func (o *OutputProfile) SetUnderscoreUpdatedBy(v string)`

SetUnderscoreUpdatedBy sets UnderscoreUpdatedBy field to given value.

### HasUnderscoreUpdatedBy

`func (o *OutputProfile) HasUnderscoreUpdatedBy() bool`

HasUnderscoreUpdatedBy returns a boolean if a field has been set.

### GetUnderscoreCreatedAt

`func (o *OutputProfile) GetUnderscoreCreatedAt() int64`

GetUnderscoreCreatedAt returns the UnderscoreCreatedAt field if non-nil, zero value otherwise.

### GetUnderscoreCreatedAtOk

`func (o *OutputProfile) GetUnderscoreCreatedAtOk() (*int64, bool)`

GetUnderscoreCreatedAtOk returns a tuple with the UnderscoreCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedAt

`func (o *OutputProfile) SetUnderscoreCreatedAt(v int64)`

SetUnderscoreCreatedAt sets UnderscoreCreatedAt field to given value.


### GetUnderscoreUpdatedAt

`func (o *OutputProfile) GetUnderscoreUpdatedAt() int64`

GetUnderscoreUpdatedAt returns the UnderscoreUpdatedAt field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedAtOk

`func (o *OutputProfile) GetUnderscoreUpdatedAtOk() (*int64, bool)`

GetUnderscoreUpdatedAtOk returns a tuple with the UnderscoreUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedAt

`func (o *OutputProfile) SetUnderscoreUpdatedAt(v int64)`

SetUnderscoreUpdatedAt sets UnderscoreUpdatedAt field to given value.

### HasUnderscoreUpdatedAt

`func (o *OutputProfile) HasUnderscoreUpdatedAt() bool`

HasUnderscoreUpdatedAt returns a boolean if a field has been set.

### GetType

`func (o *OutputProfile) GetType() ProfileType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OutputProfile) GetTypeOk() (*ProfileType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OutputProfile) SetType(v ProfileType)`

SetType sets Type field to given value.


### GetName

`func (o *OutputProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OutputProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OutputProfile) SetName(v string)`

SetName sets Name field to given value.


### GetPermissions

`func (o *OutputProfile) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *OutputProfile) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *OutputProfile) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *OutputProfile) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetEditable

`func (o *OutputProfile) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *OutputProfile) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *OutputProfile) SetEditable(v bool)`

SetEditable sets Editable field to given value.


### GetForAdmin

`func (o *OutputProfile) GetForAdmin() bool`

GetForAdmin returns the ForAdmin field if non-nil, zero value otherwise.

### GetForAdminOk

`func (o *OutputProfile) GetForAdminOk() (*bool, bool)`

GetForAdminOk returns a tuple with the ForAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForAdmin

`func (o *OutputProfile) SetForAdmin(v bool)`

SetForAdmin sets ForAdmin field to given value.


### GetForOrg

`func (o *OutputProfile) GetForOrg() bool`

GetForOrg returns the ForOrg field if non-nil, zero value otherwise.

### GetForOrgOk

`func (o *OutputProfile) GetForOrgOk() (*bool, bool)`

GetForOrgOk returns a tuple with the ForOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForOrg

`func (o *OutputProfile) SetForOrg(v bool)`

SetForOrg sets ForOrg field to given value.


### GetForExternal

`func (o *OutputProfile) GetForExternal() bool`

GetForExternal returns the ForExternal field if non-nil, zero value otherwise.

### GetForExternalOk

`func (o *OutputProfile) GetForExternalOk() (*bool, bool)`

GetForExternalOk returns a tuple with the ForExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForExternal

`func (o *OutputProfile) SetForExternal(v bool)`

SetForExternal sets ForExternal field to given value.


### GetConsumesLicense

`func (o *OutputProfile) GetConsumesLicense() bool`

GetConsumesLicense returns the ConsumesLicense field if non-nil, zero value otherwise.

### GetConsumesLicenseOk

`func (o *OutputProfile) GetConsumesLicenseOk() (*bool, bool)`

GetConsumesLicenseOk returns a tuple with the ConsumesLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumesLicense

`func (o *OutputProfile) SetConsumesLicense(v bool)`

SetConsumesLicense sets ConsumesLicense field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


