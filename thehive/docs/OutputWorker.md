# OutputWorker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Version** | **string** |  | 
**Description** | **string** |  | 
**DataTypeList** | Pointer to **[]string** |  | [optional] 
**CortexIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewOutputWorker

`func NewOutputWorker(id string, name string, version string, description string, ) *OutputWorker`

NewOutputWorker instantiates a new OutputWorker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputWorkerWithDefaults

`func NewOutputWorkerWithDefaults() *OutputWorker`

NewOutputWorkerWithDefaults instantiates a new OutputWorker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OutputWorker) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutputWorker) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OutputWorker) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *OutputWorker) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OutputWorker) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OutputWorker) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *OutputWorker) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OutputWorker) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OutputWorker) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetDescription

`func (o *OutputWorker) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OutputWorker) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OutputWorker) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDataTypeList

`func (o *OutputWorker) GetDataTypeList() []string`

GetDataTypeList returns the DataTypeList field if non-nil, zero value otherwise.

### GetDataTypeListOk

`func (o *OutputWorker) GetDataTypeListOk() (*[]string, bool)`

GetDataTypeListOk returns a tuple with the DataTypeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypeList

`func (o *OutputWorker) SetDataTypeList(v []string)`

SetDataTypeList sets DataTypeList field to given value.

### HasDataTypeList

`func (o *OutputWorker) HasDataTypeList() bool`

HasDataTypeList returns a boolean if a field has been set.

### GetCortexIds

`func (o *OutputWorker) GetCortexIds() []string`

GetCortexIds returns the CortexIds field if non-nil, zero value otherwise.

### GetCortexIdsOk

`func (o *OutputWorker) GetCortexIdsOk() (*[]string, bool)`

GetCortexIdsOk returns a tuple with the CortexIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCortexIds

`func (o *OutputWorker) SetCortexIds(v []string)`

SetCortexIds sets CortexIds field to given value.

### HasCortexIds

`func (o *OutputWorker) HasCortexIds() bool`

HasCortexIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


