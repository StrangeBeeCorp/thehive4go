# KeyOutputAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | API key included in the authentication header. |
**Prefix** | Pointer to **string** | Optional prefix prepended to the API key in the header, separated by a space. | [optional]
**Type** | **string** | Authentication type identifier. |

## Methods

### NewKeyOutputAuth

`func NewKeyOutputAuth(key string, type_ string, ) *KeyOutputAuth`

NewKeyOutputAuth instantiates a new KeyOutputAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyOutputAuthWithDefaults

`func NewKeyOutputAuthWithDefaults() *KeyOutputAuth`

NewKeyOutputAuthWithDefaults instantiates a new KeyOutputAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *KeyOutputAuth) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *KeyOutputAuth) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *KeyOutputAuth) SetKey(v string)`

SetKey sets Key field to given value.


### GetPrefix

`func (o *KeyOutputAuth) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *KeyOutputAuth) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *KeyOutputAuth) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *KeyOutputAuth) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetType

`func (o *KeyOutputAuth) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KeyOutputAuth) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KeyOutputAuth) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
