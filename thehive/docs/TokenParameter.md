# TokenParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Additional parameter key to include in the OAuth 2.0 token request. |
**Value** | **string** | Additional parameter value. |

## Methods

### NewTokenParameter

`func NewTokenParameter(key string, value string, ) *TokenParameter`

NewTokenParameter instantiates a new TokenParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenParameterWithDefaults

`func NewTokenParameterWithDefaults() *TokenParameter`

NewTokenParameterWithDefaults instantiates a new TokenParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *TokenParameter) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TokenParameter) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TokenParameter) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *TokenParameter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TokenParameter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TokenParameter) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
