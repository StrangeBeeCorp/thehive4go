# OutputLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnderscoreId** | **string** | Internal ID of the license. |
**UnderscoreCreatedAt** | **int64** | Date the license was added to TheHive, as a Unix timestamp in milliseconds. |
**UnderscoreCreatedBy** | **string** | Login of the user who added the license. |
**UnderscoreType** | **string** | Entity type. Always &#x60;License&#x60;. |
**Id** | **string** | License identifier from the license file. |
**Customer** | **string** | Customer name as recorded in the license file. |
**Plan** | **string** | License plan. One of &#x60;No&#x60;, &#x60;Community&#x60;, &#x60;Gold&#x60;, &#x60;Platinum&#x60;, or &#x60;One&#x60;. |
**Kind** | **string** | License kind. One of &#x60;Regular&#x60;, &#x60;Trial&#x60;, or &#x60;Development&#x60;. |
**ValidFrom** | **int64** | Start of the license validity period, as a Unix timestamp in milliseconds. |
**ExpiresAt** | **int64** | End of the license validity period, as a Unix timestamp in milliseconds. |
**Current** | **bool** | Whether this is the currently active license. |
**Capabilities** | Pointer to **[]string** | Capability names enabled by this license. | [optional]
**Quotas** | **map[string]int32** | Resource quotas defined by the license, mapping capability name to its limit. A value of &#x60;-1&#x60; means unlimited. |

## Methods

### NewOutputLicense

`func NewOutputLicense(underscoreId string, underscoreCreatedAt int64, underscoreCreatedBy string, underscoreType string, id string, customer string, plan string, kind string, validFrom int64, expiresAt int64, current bool, quotas map[string]int32, ) *OutputLicense`

NewOutputLicense instantiates a new OutputLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputLicenseWithDefaults

`func NewOutputLicenseWithDefaults() *OutputLicense`

NewOutputLicenseWithDefaults instantiates a new OutputLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnderscoreId

`func (o *OutputLicense) GetUnderscoreId() string`

GetUnderscoreId returns the UnderscoreId field if non-nil, zero value otherwise.

### GetUnderscoreIdOk

`func (o *OutputLicense) GetUnderscoreIdOk() (*string, bool)`

GetUnderscoreIdOk returns a tuple with the UnderscoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreId

`func (o *OutputLicense) SetUnderscoreId(v string)`

SetUnderscoreId sets UnderscoreId field to given value.


### GetUnderscoreCreatedAt

`func (o *OutputLicense) GetUnderscoreCreatedAt() int64`

GetUnderscoreCreatedAt returns the UnderscoreCreatedAt field if non-nil, zero value otherwise.

### GetUnderscoreCreatedAtOk

`func (o *OutputLicense) GetUnderscoreCreatedAtOk() (*int64, bool)`

GetUnderscoreCreatedAtOk returns a tuple with the UnderscoreCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedAt

`func (o *OutputLicense) SetUnderscoreCreatedAt(v int64)`

SetUnderscoreCreatedAt sets UnderscoreCreatedAt field to given value.


### GetUnderscoreCreatedBy

`func (o *OutputLicense) GetUnderscoreCreatedBy() string`

GetUnderscoreCreatedBy returns the UnderscoreCreatedBy field if non-nil, zero value otherwise.

### GetUnderscoreCreatedByOk

`func (o *OutputLicense) GetUnderscoreCreatedByOk() (*string, bool)`

GetUnderscoreCreatedByOk returns a tuple with the UnderscoreCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedBy

`func (o *OutputLicense) SetUnderscoreCreatedBy(v string)`

SetUnderscoreCreatedBy sets UnderscoreCreatedBy field to given value.


### GetUnderscoreType

`func (o *OutputLicense) GetUnderscoreType() string`

GetUnderscoreType returns the UnderscoreType field if non-nil, zero value otherwise.

### GetUnderscoreTypeOk

`func (o *OutputLicense) GetUnderscoreTypeOk() (*string, bool)`

GetUnderscoreTypeOk returns a tuple with the UnderscoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreType

`func (o *OutputLicense) SetUnderscoreType(v string)`

SetUnderscoreType sets UnderscoreType field to given value.


### GetId

`func (o *OutputLicense) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutputLicense) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OutputLicense) SetId(v string)`

SetId sets Id field to given value.


### GetCustomer

`func (o *OutputLicense) GetCustomer() string`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *OutputLicense) GetCustomerOk() (*string, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *OutputLicense) SetCustomer(v string)`

SetCustomer sets Customer field to given value.


### GetPlan

`func (o *OutputLicense) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *OutputLicense) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *OutputLicense) SetPlan(v string)`

SetPlan sets Plan field to given value.


### GetKind

`func (o *OutputLicense) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *OutputLicense) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *OutputLicense) SetKind(v string)`

SetKind sets Kind field to given value.


### GetValidFrom

`func (o *OutputLicense) GetValidFrom() int64`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *OutputLicense) GetValidFromOk() (*int64, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *OutputLicense) SetValidFrom(v int64)`

SetValidFrom sets ValidFrom field to given value.


### GetExpiresAt

`func (o *OutputLicense) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *OutputLicense) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *OutputLicense) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.


### GetCurrent

`func (o *OutputLicense) GetCurrent() bool`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *OutputLicense) GetCurrentOk() (*bool, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *OutputLicense) SetCurrent(v bool)`

SetCurrent sets Current field to given value.


### GetCapabilities

`func (o *OutputLicense) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *OutputLicense) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *OutputLicense) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *OutputLicense) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetQuotas

`func (o *OutputLicense) GetQuotas() map[string]int32`

GetQuotas returns the Quotas field if non-nil, zero value otherwise.

### GetQuotasOk

`func (o *OutputLicense) GetQuotasOk() (*map[string]int32, bool)`

GetQuotasOk returns a tuple with the Quotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotas

`func (o *OutputLicense) SetQuotas(v map[string]int32)`

SetQuotas sets Quotas field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
