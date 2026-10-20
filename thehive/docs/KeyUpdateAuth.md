# KeyUpdateAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | API key included in the authentication header. |
**Prefix** | Pointer to **string** | Optional prefix prepended to the API key in the header, separated by a space. | [optional]
**Type** | **string** | Authentication type identifier. Must be &#x60;key&#x60;. |

## Methods

### NewKeyUpdateAuth

`func NewKeyUpdateAuth(key string, type_ string, ) *KeyUpdateAuth`

NewKeyUpdateAuth instantiates a new KeyUpdateAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyUpdateAuthWithDefaults

`func NewKeyUpdateAuthWithDefaults() *KeyUpdateAuth`

NewKeyUpdateAuthWithDefaults instantiates a new KeyUpdateAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *KeyUpdateAuth) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *KeyUpdateAuth) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *KeyUpdateAuth) SetKey(v string)`

SetKey sets Key field to given value.


### GetPrefix

`func (o *KeyUpdateAuth) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *KeyUpdateAuth) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *KeyUpdateAuth) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *KeyUpdateAuth) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetType

`func (o *KeyUpdateAuth) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KeyUpdateAuth) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KeyUpdateAuth) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
