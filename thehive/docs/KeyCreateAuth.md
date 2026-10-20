# KeyCreateAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | API key included in the authentication header. |
**Prefix** | Pointer to **string** | Optional prefix prepended to the API key in the header, separated by a space. | [optional]
**Type** | **string** | Authentication type identifier. Must be &#x60;key&#x60;. |

## Methods

### NewKeyCreateAuth

`func NewKeyCreateAuth(key string, type_ string, ) *KeyCreateAuth`

NewKeyCreateAuth instantiates a new KeyCreateAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyCreateAuthWithDefaults

`func NewKeyCreateAuthWithDefaults() *KeyCreateAuth`

NewKeyCreateAuthWithDefaults instantiates a new KeyCreateAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *KeyCreateAuth) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *KeyCreateAuth) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *KeyCreateAuth) SetKey(v string)`

SetKey sets Key field to given value.


### GetPrefix

`func (o *KeyCreateAuth) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *KeyCreateAuth) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *KeyCreateAuth) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *KeyCreateAuth) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetType

`func (o *KeyCreateAuth) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KeyCreateAuth) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KeyCreateAuth) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
