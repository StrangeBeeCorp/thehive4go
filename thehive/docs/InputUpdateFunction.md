# InputUpdateFunction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to [**InputFunctionMode**](InputFunctionMode.md) |  | [optional] 
**Definition** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Types** | Pointer to [**[]InputFunctionType**](InputFunctionType.md) |  | [optional] 

## Methods

### NewInputUpdateFunction

`func NewInputUpdateFunction() *InputUpdateFunction`

NewInputUpdateFunction instantiates a new InputUpdateFunction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputUpdateFunctionWithDefaults

`func NewInputUpdateFunctionWithDefaults() *InputUpdateFunction`

NewInputUpdateFunctionWithDefaults instantiates a new InputUpdateFunction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InputUpdateFunction) GetMode() InputFunctionMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InputUpdateFunction) GetModeOk() (*InputFunctionMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InputUpdateFunction) SetMode(v InputFunctionMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InputUpdateFunction) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetDefinition

`func (o *InputUpdateFunction) GetDefinition() string`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *InputUpdateFunction) GetDefinitionOk() (*string, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *InputUpdateFunction) SetDefinition(v string)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *InputUpdateFunction) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetDescription

`func (o *InputUpdateFunction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputUpdateFunction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputUpdateFunction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InputUpdateFunction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConfig

`func (o *InputUpdateFunction) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *InputUpdateFunction) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *InputUpdateFunction) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *InputUpdateFunction) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetTypes

`func (o *InputUpdateFunction) GetTypes() []InputFunctionType`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *InputUpdateFunction) GetTypesOk() (*[]InputFunctionType, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *InputUpdateFunction) SetTypes(v []InputFunctionType)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *InputUpdateFunction) HasTypes() bool`

HasTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


