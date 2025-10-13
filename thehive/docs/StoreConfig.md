# StoreConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **string** |  | [optional] 
**FilePath** | Pointer to **string** |  | [optional] 
**IsFileOnClasspath** | Pointer to **bool** |  | [optional] [default to false]
**Password** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewStoreConfig

`func NewStoreConfig(type_ string, ) *StoreConfig`

NewStoreConfig instantiates a new StoreConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreConfigWithDefaults

`func NewStoreConfigWithDefaults() *StoreConfig`

NewStoreConfigWithDefaults instantiates a new StoreConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *StoreConfig) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *StoreConfig) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *StoreConfig) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *StoreConfig) HasData() bool`

HasData returns a boolean if a field has been set.

### GetFilePath

`func (o *StoreConfig) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *StoreConfig) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *StoreConfig) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *StoreConfig) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetIsFileOnClasspath

`func (o *StoreConfig) GetIsFileOnClasspath() bool`

GetIsFileOnClasspath returns the IsFileOnClasspath field if non-nil, zero value otherwise.

### GetIsFileOnClasspathOk

`func (o *StoreConfig) GetIsFileOnClasspathOk() (*bool, bool)`

GetIsFileOnClasspathOk returns a tuple with the IsFileOnClasspath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFileOnClasspath

`func (o *StoreConfig) SetIsFileOnClasspath(v bool)`

SetIsFileOnClasspath sets IsFileOnClasspath field to given value.

### HasIsFileOnClasspath

`func (o *StoreConfig) HasIsFileOnClasspath() bool`

HasIsFileOnClasspath returns a boolean if a field has been set.

### GetPassword

`func (o *StoreConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *StoreConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *StoreConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *StoreConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetType

`func (o *StoreConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StoreConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StoreConfig) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


