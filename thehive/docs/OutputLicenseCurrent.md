# OutputLicenseCurrent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** | Error message describing why the license failed validation. |
**Fallback** | [**OutputLicense**](OutputLicense.md) | Placeholder license representing an unlicensed instance. |
**NotFound** | Pointer to **bool** | Always &#x60;true&#x60; for this response variant. | [optional] [default to true]
**License** | [**OutputLicense**](OutputLicense.md) | The currently active license. |

## Methods

### NewOutputLicenseCurrent

`func NewOutputLicenseCurrent(error_ string, fallback OutputLicense, license OutputLicense, ) *OutputLicenseCurrent`

NewOutputLicenseCurrent instantiates a new OutputLicenseCurrent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputLicenseCurrentWithDefaults

`func NewOutputLicenseCurrentWithDefaults() *OutputLicenseCurrent`

NewOutputLicenseCurrentWithDefaults instantiates a new OutputLicenseCurrent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *OutputLicenseCurrent) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *OutputLicenseCurrent) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *OutputLicenseCurrent) SetError(v string)`

SetError sets Error field to given value.


### GetFallback

`func (o *OutputLicenseCurrent) GetFallback() OutputLicense`

GetFallback returns the Fallback field if non-nil, zero value otherwise.

### GetFallbackOk

`func (o *OutputLicenseCurrent) GetFallbackOk() (*OutputLicense, bool)`

GetFallbackOk returns a tuple with the Fallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallback

`func (o *OutputLicenseCurrent) SetFallback(v OutputLicense)`

SetFallback sets Fallback field to given value.


### GetNotFound

`func (o *OutputLicenseCurrent) GetNotFound() bool`

GetNotFound returns the NotFound field if non-nil, zero value otherwise.

### GetNotFoundOk

`func (o *OutputLicenseCurrent) GetNotFoundOk() (*bool, bool)`

GetNotFoundOk returns a tuple with the NotFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotFound

`func (o *OutputLicenseCurrent) SetNotFound(v bool)`

SetNotFound sets NotFound field to given value.

### HasNotFound

`func (o *OutputLicenseCurrent) HasNotFound() bool`

HasNotFound returns a boolean if a field has been set.

### GetLicense

`func (o *OutputLicenseCurrent) GetLicense() OutputLicense`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *OutputLicenseCurrent) GetLicenseOk() (*OutputLicense, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *OutputLicenseCurrent) SetLicense(v OutputLicense)`

SetLicense sets License field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
