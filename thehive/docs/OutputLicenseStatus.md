# OutputLicenseStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the license. |
**Customer** | **string** | Name of the customer this license is issued to. |
**Instance** | **string** | Identifier of the TheHive instance this license is bound to. |
**Plan** | **string** | License plan: &#x60;No&#x60;, &#x60;Community&#x60;, &#x60;Gold&#x60;, &#x60;Platinum&#x60;, or &#x60;One&#x60;. |
**Kind** | **string** | License kind: &#x60;Regular&#x60;, &#x60;Trial&#x60;, or &#x60;Development&#x60;. |
**ValidFrom** | **int64** | Date the license became valid. Returned as a Unix timestamp in milliseconds. |
**ExpiresAt** | **int64** | Date the license expires. Returned as a Unix timestamp in milliseconds. |
**Capabilities** | Pointer to **[]string** | Capabilities unlocked by the license plan. | [optional]
**IsValid** | **bool** | Whether the license is currently valid. |
**Error** | Pointer to **string** | Reason the license isn&#39;t valid. Absent when the license is valid. | [optional]
**Quotas** | [**map[string]OutputLicenseQuota**](OutputLicenseQuota.md) | Usage and limit for each licensed quota, keyed by capability name. |

## Methods

### NewOutputLicenseStatus

`func NewOutputLicenseStatus(id string, customer string, instance string, plan string, kind string, validFrom int64, expiresAt int64, isValid bool, quotas map[string]OutputLicenseQuota, ) *OutputLicenseStatus`

NewOutputLicenseStatus instantiates a new OutputLicenseStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputLicenseStatusWithDefaults

`func NewOutputLicenseStatusWithDefaults() *OutputLicenseStatus`

NewOutputLicenseStatusWithDefaults instantiates a new OutputLicenseStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OutputLicenseStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutputLicenseStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OutputLicenseStatus) SetId(v string)`

SetId sets Id field to given value.


### GetCustomer

`func (o *OutputLicenseStatus) GetCustomer() string`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *OutputLicenseStatus) GetCustomerOk() (*string, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *OutputLicenseStatus) SetCustomer(v string)`

SetCustomer sets Customer field to given value.


### GetInstance

`func (o *OutputLicenseStatus) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *OutputLicenseStatus) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *OutputLicenseStatus) SetInstance(v string)`

SetInstance sets Instance field to given value.


### GetPlan

`func (o *OutputLicenseStatus) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *OutputLicenseStatus) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *OutputLicenseStatus) SetPlan(v string)`

SetPlan sets Plan field to given value.


### GetKind

`func (o *OutputLicenseStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *OutputLicenseStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *OutputLicenseStatus) SetKind(v string)`

SetKind sets Kind field to given value.


### GetValidFrom

`func (o *OutputLicenseStatus) GetValidFrom() int64`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *OutputLicenseStatus) GetValidFromOk() (*int64, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *OutputLicenseStatus) SetValidFrom(v int64)`

SetValidFrom sets ValidFrom field to given value.


### GetExpiresAt

`func (o *OutputLicenseStatus) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *OutputLicenseStatus) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *OutputLicenseStatus) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.


### GetCapabilities

`func (o *OutputLicenseStatus) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *OutputLicenseStatus) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *OutputLicenseStatus) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *OutputLicenseStatus) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetIsValid

`func (o *OutputLicenseStatus) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *OutputLicenseStatus) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *OutputLicenseStatus) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.


### GetError

`func (o *OutputLicenseStatus) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *OutputLicenseStatus) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *OutputLicenseStatus) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *OutputLicenseStatus) HasError() bool`

HasError returns a boolean if a field has been set.

### GetQuotas

`func (o *OutputLicenseStatus) GetQuotas() map[string]OutputLicenseQuota`

GetQuotas returns the Quotas field if non-nil, zero value otherwise.

### GetQuotasOk

`func (o *OutputLicenseStatus) GetQuotasOk() (*map[string]OutputLicenseQuota, bool)`

GetQuotasOk returns a tuple with the Quotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotas

`func (o *OutputLicenseStatus) SetQuotas(v map[string]OutputLicenseQuota)`

SetQuotas sets Quotas field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
