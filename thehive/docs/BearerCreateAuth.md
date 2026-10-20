# BearerCreateAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Bearer token included in the &#x60;Authorization: Bearer&#x60; header. |
**Type** | **string** | Authentication type identifier. Must be &#x60;bearer&#x60;. |

## Methods

### NewBearerCreateAuth

`func NewBearerCreateAuth(key string, type_ string, ) *BearerCreateAuth`

NewBearerCreateAuth instantiates a new BearerCreateAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBearerCreateAuthWithDefaults

`func NewBearerCreateAuthWithDefaults() *BearerCreateAuth`

NewBearerCreateAuthWithDefaults instantiates a new BearerCreateAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *BearerCreateAuth) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *BearerCreateAuth) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *BearerCreateAuth) SetKey(v string)`

SetKey sets Key field to given value.


### GetType

`func (o *BearerCreateAuth) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BearerCreateAuth) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BearerCreateAuth) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
