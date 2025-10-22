# InputAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponderId** | **string** |  | 
**CortexId** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | 
**ObjectId** | **string** |  | 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**Tlp** | Pointer to **int32** |  | [optional] 

## Methods

### NewInputAction

`func NewInputAction(responderId string, objectType string, objectId string, ) *InputAction`

NewInputAction instantiates a new InputAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputActionWithDefaults

`func NewInputActionWithDefaults() *InputAction`

NewInputActionWithDefaults instantiates a new InputAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponderId

`func (o *InputAction) GetResponderId() string`

GetResponderId returns the ResponderId field if non-nil, zero value otherwise.

### GetResponderIdOk

`func (o *InputAction) GetResponderIdOk() (*string, bool)`

GetResponderIdOk returns a tuple with the ResponderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponderId

`func (o *InputAction) SetResponderId(v string)`

SetResponderId sets ResponderId field to given value.


### GetCortexId

`func (o *InputAction) GetCortexId() string`

GetCortexId returns the CortexId field if non-nil, zero value otherwise.

### GetCortexIdOk

`func (o *InputAction) GetCortexIdOk() (*string, bool)`

GetCortexIdOk returns a tuple with the CortexId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCortexId

`func (o *InputAction) SetCortexId(v string)`

SetCortexId sets CortexId field to given value.

### HasCortexId

`func (o *InputAction) HasCortexId() bool`

HasCortexId returns a boolean if a field has been set.

### GetObjectType

`func (o *InputAction) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *InputAction) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *InputAction) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectId

`func (o *InputAction) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *InputAction) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *InputAction) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetParameters

`func (o *InputAction) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *InputAction) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *InputAction) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *InputAction) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetTlp

`func (o *InputAction) GetTlp() int32`

GetTlp returns the Tlp field if non-nil, zero value otherwise.

### GetTlpOk

`func (o *InputAction) GetTlpOk() (*int32, bool)`

GetTlpOk returns a tuple with the Tlp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlp

`func (o *InputAction) SetTlp(v int32)`

SetTlp sets Tlp field to given value.

### HasTlp

`func (o *InputAction) HasTlp() bool`

HasTlp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


