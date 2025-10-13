# InputTestFunction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the function | [optional] 
**Definition** | **string** | Code of the function | 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Input** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewInputTestFunction

`func NewInputTestFunction(definition string, ) *InputTestFunction`

NewInputTestFunction instantiates a new InputTestFunction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputTestFunctionWithDefaults

`func NewInputTestFunctionWithDefaults() *InputTestFunction`

NewInputTestFunctionWithDefaults instantiates a new InputTestFunction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InputTestFunction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputTestFunction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputTestFunction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InputTestFunction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefinition

`func (o *InputTestFunction) GetDefinition() string`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *InputTestFunction) GetDefinitionOk() (*string, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *InputTestFunction) SetDefinition(v string)`

SetDefinition sets Definition field to given value.


### GetConfig

`func (o *InputTestFunction) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *InputTestFunction) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *InputTestFunction) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *InputTestFunction) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetInput

`func (o *InputTestFunction) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *InputTestFunction) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *InputTestFunction) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *InputTestFunction) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *InputTestFunction) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *InputTestFunction) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


