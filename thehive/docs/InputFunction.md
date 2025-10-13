# InputFunction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the function, should be unique in an organisation | 
**Mode** | Pointer to [**InputFunctionMode**](InputFunctionMode.md) |  | [optional] 
**Definition** | **string** | Code of the function | 
**Description** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Types** | Pointer to [**[]InputFunctionType**](InputFunctionType.md) |  | [optional] 

## Methods

### NewInputFunction

`func NewInputFunction(name string, definition string, ) *InputFunction`

NewInputFunction instantiates a new InputFunction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputFunctionWithDefaults

`func NewInputFunctionWithDefaults() *InputFunction`

NewInputFunctionWithDefaults instantiates a new InputFunction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InputFunction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputFunction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputFunction) SetName(v string)`

SetName sets Name field to given value.


### GetMode

`func (o *InputFunction) GetMode() InputFunctionMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InputFunction) GetModeOk() (*InputFunctionMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InputFunction) SetMode(v InputFunctionMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InputFunction) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetDefinition

`func (o *InputFunction) GetDefinition() string`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *InputFunction) GetDefinitionOk() (*string, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *InputFunction) SetDefinition(v string)`

SetDefinition sets Definition field to given value.


### GetDescription

`func (o *InputFunction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputFunction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputFunction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InputFunction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConfig

`func (o *InputFunction) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *InputFunction) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *InputFunction) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *InputFunction) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetTypes

`func (o *InputFunction) GetTypes() []InputFunctionType`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *InputFunction) GetTypesOk() (*[]InputFunctionType, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *InputFunction) SetTypes(v []InputFunctionType)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *InputFunction) HasTypes() bool`

HasTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


