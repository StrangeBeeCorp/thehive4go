# Manager

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **string** |  | [optional] 
**Stores** | Pointer to [**[]StoreConfig**](StoreConfig.md) |  | [optional] 

## Methods

### NewManager

`func NewManager() *Manager`

NewManager instantiates a new Manager object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagerWithDefaults

`func NewManagerWithDefaults() *Manager`

NewManagerWithDefaults instantiates a new Manager object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *Manager) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *Manager) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *Manager) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *Manager) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetStores

`func (o *Manager) GetStores() []StoreConfig`

GetStores returns the Stores field if non-nil, zero value otherwise.

### GetStoresOk

`func (o *Manager) GetStoresOk() (*[]StoreConfig, bool)`

GetStoresOk returns a tuple with the Stores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStores

`func (o *Manager) SetStores(v []StoreConfig)`

SetStores sets Stores field to given value.

### HasStores

`func (o *Manager) HasStores() bool`

HasStores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


