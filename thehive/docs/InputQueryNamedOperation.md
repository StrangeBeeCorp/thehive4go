# InputQueryNamedOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  |
**ExtraData** | Pointer to **[]string** |  | [optional]
**From** | **int32** |  |
**To** | **int32** |  |
**Fields** | Pointer to **[]map[string]interface{}** |  | [optional]

## Methods

### NewInputQueryNamedOperation

`func NewInputQueryNamedOperation(name string, from int32, to int32, ) *InputQueryNamedOperation`

NewInputQueryNamedOperation instantiates a new InputQueryNamedOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputQueryNamedOperationWithDefaults

`func NewInputQueryNamedOperationWithDefaults() *InputQueryNamedOperation`

NewInputQueryNamedOperationWithDefaults instantiates a new InputQueryNamedOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InputQueryNamedOperation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputQueryNamedOperation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputQueryNamedOperation) SetName(v string)`

SetName sets Name field to given value.


### GetExtraData

`func (o *InputQueryNamedOperation) GetExtraData() []string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *InputQueryNamedOperation) GetExtraDataOk() (*[]string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *InputQueryNamedOperation) SetExtraData(v []string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *InputQueryNamedOperation) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### GetFrom

`func (o *InputQueryNamedOperation) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *InputQueryNamedOperation) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *InputQueryNamedOperation) SetFrom(v int32)`

SetFrom sets From field to given value.


### GetTo

`func (o *InputQueryNamedOperation) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *InputQueryNamedOperation) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *InputQueryNamedOperation) SetTo(v int32)`

SetTo sets To field to given value.


### GetFields

`func (o *InputQueryNamedOperation) GetFields() []map[string]interface{}`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *InputQueryNamedOperation) GetFieldsOk() (*[]map[string]interface{}, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *InputQueryNamedOperation) SetFields(v []map[string]interface{})`

SetFields sets Fields field to given value.

### HasFields

`func (o *InputQueryNamedOperation) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
