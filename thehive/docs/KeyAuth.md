# KeyAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Prefix** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewKeyAuth

`func NewKeyAuth(key string, type_ string, ) *KeyAuth`

NewKeyAuth instantiates a new KeyAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyAuthWithDefaults

`func NewKeyAuthWithDefaults() *KeyAuth`

NewKeyAuthWithDefaults instantiates a new KeyAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *KeyAuth) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *KeyAuth) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *KeyAuth) SetKey(v string)`

SetKey sets Key field to given value.


### GetPrefix

`func (o *KeyAuth) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *KeyAuth) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *KeyAuth) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *KeyAuth) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetType

`func (o *KeyAuth) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KeyAuth) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KeyAuth) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


