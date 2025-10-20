# OutputFunction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnderscoreId** | **string** |  |
**UnderscoreType** | **string** |  |
**UnderscoreCreatedBy** | **string** |  |
**UnderscoreCreatedAt** | **int64** |  |
**UnderscoreUpdatedBy** | Pointer to **string** |  | [optional]
**UnderscoreUpdatedAt** | Pointer to **int64** |  | [optional]
**Name** | **string** |  |
**Mode** | **string** |  |
**Definition** | **string** |  |
**Description** | Pointer to **string** |  | [optional]
**Config** | **map[string]interface{}** |  |
**LastSuccessDate** | Pointer to **int64** |  | [optional]
**LastSuccessDetails** | Pointer to **map[string]interface{}** |  | [optional]
**LastErrorDate** | Pointer to **int64** |  | [optional]
**LastErrorDetails** | Pointer to **map[string]interface{}** |  | [optional]
**Types** | Pointer to **[]string** |  | [optional]

## Methods

### NewOutputFunction

`func NewOutputFunction(underscoreId string, underscoreType string, underscoreCreatedBy string, underscoreCreatedAt int64, name string, mode string, definition string, config map[string]interface{}, ) *OutputFunction`

NewOutputFunction instantiates a new OutputFunction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputFunctionWithDefaults

`func NewOutputFunctionWithDefaults() *OutputFunction`

NewOutputFunctionWithDefaults instantiates a new OutputFunction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnderscoreId

`func (o *OutputFunction) GetUnderscoreId() string`

GetUnderscoreId returns the UnderscoreId field if non-nil, zero value otherwise.

### GetUnderscoreIdOk

`func (o *OutputFunction) GetUnderscoreIdOk() (*string, bool)`

GetUnderscoreIdOk returns a tuple with the UnderscoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreId

`func (o *OutputFunction) SetUnderscoreId(v string)`

SetUnderscoreId sets UnderscoreId field to given value.


### GetUnderscoreType

`func (o *OutputFunction) GetUnderscoreType() string`

GetUnderscoreType returns the UnderscoreType field if non-nil, zero value otherwise.

### GetUnderscoreTypeOk

`func (o *OutputFunction) GetUnderscoreTypeOk() (*string, bool)`

GetUnderscoreTypeOk returns a tuple with the UnderscoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreType

`func (o *OutputFunction) SetUnderscoreType(v string)`

SetUnderscoreType sets UnderscoreType field to given value.


### GetUnderscoreCreatedBy

`func (o *OutputFunction) GetUnderscoreCreatedBy() string`

GetUnderscoreCreatedBy returns the UnderscoreCreatedBy field if non-nil, zero value otherwise.

### GetUnderscoreCreatedByOk

`func (o *OutputFunction) GetUnderscoreCreatedByOk() (*string, bool)`

GetUnderscoreCreatedByOk returns a tuple with the UnderscoreCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedBy

`func (o *OutputFunction) SetUnderscoreCreatedBy(v string)`

SetUnderscoreCreatedBy sets UnderscoreCreatedBy field to given value.


### GetUnderscoreCreatedAt

`func (o *OutputFunction) GetUnderscoreCreatedAt() int64`

GetUnderscoreCreatedAt returns the UnderscoreCreatedAt field if non-nil, zero value otherwise.

### GetUnderscoreCreatedAtOk

`func (o *OutputFunction) GetUnderscoreCreatedAtOk() (*int64, bool)`

GetUnderscoreCreatedAtOk returns a tuple with the UnderscoreCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedAt

`func (o *OutputFunction) SetUnderscoreCreatedAt(v int64)`

SetUnderscoreCreatedAt sets UnderscoreCreatedAt field to given value.


### GetUnderscoreUpdatedBy

`func (o *OutputFunction) GetUnderscoreUpdatedBy() string`

GetUnderscoreUpdatedBy returns the UnderscoreUpdatedBy field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedByOk

`func (o *OutputFunction) GetUnderscoreUpdatedByOk() (*string, bool)`

GetUnderscoreUpdatedByOk returns a tuple with the UnderscoreUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedBy

`func (o *OutputFunction) SetUnderscoreUpdatedBy(v string)`

SetUnderscoreUpdatedBy sets UnderscoreUpdatedBy field to given value.

### HasUnderscoreUpdatedBy

`func (o *OutputFunction) HasUnderscoreUpdatedBy() bool`

HasUnderscoreUpdatedBy returns a boolean if a field has been set.

### GetUnderscoreUpdatedAt

`func (o *OutputFunction) GetUnderscoreUpdatedAt() int64`

GetUnderscoreUpdatedAt returns the UnderscoreUpdatedAt field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedAtOk

`func (o *OutputFunction) GetUnderscoreUpdatedAtOk() (*int64, bool)`

GetUnderscoreUpdatedAtOk returns a tuple with the UnderscoreUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedAt

`func (o *OutputFunction) SetUnderscoreUpdatedAt(v int64)`

SetUnderscoreUpdatedAt sets UnderscoreUpdatedAt field to given value.

### HasUnderscoreUpdatedAt

`func (o *OutputFunction) HasUnderscoreUpdatedAt() bool`

HasUnderscoreUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *OutputFunction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OutputFunction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OutputFunction) SetName(v string)`

SetName sets Name field to given value.


### GetMode

`func (o *OutputFunction) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *OutputFunction) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *OutputFunction) SetMode(v string)`

SetMode sets Mode field to given value.


### GetDefinition

`func (o *OutputFunction) GetDefinition() string`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *OutputFunction) GetDefinitionOk() (*string, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *OutputFunction) SetDefinition(v string)`

SetDefinition sets Definition field to given value.


### GetDescription

`func (o *OutputFunction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OutputFunction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OutputFunction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OutputFunction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConfig

`func (o *OutputFunction) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *OutputFunction) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *OutputFunction) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.


### GetLastSuccessDate

`func (o *OutputFunction) GetLastSuccessDate() int64`

GetLastSuccessDate returns the LastSuccessDate field if non-nil, zero value otherwise.

### GetLastSuccessDateOk

`func (o *OutputFunction) GetLastSuccessDateOk() (*int64, bool)`

GetLastSuccessDateOk returns a tuple with the LastSuccessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessDate

`func (o *OutputFunction) SetLastSuccessDate(v int64)`

SetLastSuccessDate sets LastSuccessDate field to given value.

### HasLastSuccessDate

`func (o *OutputFunction) HasLastSuccessDate() bool`

HasLastSuccessDate returns a boolean if a field has been set.

### GetLastSuccessDetails

`func (o *OutputFunction) GetLastSuccessDetails() map[string]interface{}`

GetLastSuccessDetails returns the LastSuccessDetails field if non-nil, zero value otherwise.

### GetLastSuccessDetailsOk

`func (o *OutputFunction) GetLastSuccessDetailsOk() (*map[string]interface{}, bool)`

GetLastSuccessDetailsOk returns a tuple with the LastSuccessDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessDetails

`func (o *OutputFunction) SetLastSuccessDetails(v map[string]interface{})`

SetLastSuccessDetails sets LastSuccessDetails field to given value.

### HasLastSuccessDetails

`func (o *OutputFunction) HasLastSuccessDetails() bool`

HasLastSuccessDetails returns a boolean if a field has been set.

### GetLastErrorDate

`func (o *OutputFunction) GetLastErrorDate() int64`

GetLastErrorDate returns the LastErrorDate field if non-nil, zero value otherwise.

### GetLastErrorDateOk

`func (o *OutputFunction) GetLastErrorDateOk() (*int64, bool)`

GetLastErrorDateOk returns a tuple with the LastErrorDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorDate

`func (o *OutputFunction) SetLastErrorDate(v int64)`

SetLastErrorDate sets LastErrorDate field to given value.

### HasLastErrorDate

`func (o *OutputFunction) HasLastErrorDate() bool`

HasLastErrorDate returns a boolean if a field has been set.

### GetLastErrorDetails

`func (o *OutputFunction) GetLastErrorDetails() map[string]interface{}`

GetLastErrorDetails returns the LastErrorDetails field if non-nil, zero value otherwise.

### GetLastErrorDetailsOk

`func (o *OutputFunction) GetLastErrorDetailsOk() (*map[string]interface{}, bool)`

GetLastErrorDetailsOk returns a tuple with the LastErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorDetails

`func (o *OutputFunction) SetLastErrorDetails(v map[string]interface{})`

SetLastErrorDetails sets LastErrorDetails field to given value.

### HasLastErrorDetails

`func (o *OutputFunction) HasLastErrorDetails() bool`

HasLastErrorDetails returns a boolean if a field has been set.

### GetTypes

`func (o *OutputFunction) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *OutputFunction) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *OutputFunction) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *OutputFunction) HasTypes() bool`

HasTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
