# OutputEmailIntake

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**Interval** | **string** |  | 
**Configs** | Pointer to [**[]OutputEmailIntakeConfig**](OutputEmailIntakeConfig.md) |  | [optional] 

## Methods

### NewOutputEmailIntake

`func NewOutputEmailIntake(enabled bool, interval string, ) *OutputEmailIntake`

NewOutputEmailIntake instantiates a new OutputEmailIntake object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputEmailIntakeWithDefaults

`func NewOutputEmailIntakeWithDefaults() *OutputEmailIntake`

NewOutputEmailIntakeWithDefaults instantiates a new OutputEmailIntake object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *OutputEmailIntake) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OutputEmailIntake) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OutputEmailIntake) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInterval

`func (o *OutputEmailIntake) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *OutputEmailIntake) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *OutputEmailIntake) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetConfigs

`func (o *OutputEmailIntake) GetConfigs() []OutputEmailIntakeConfig`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *OutputEmailIntake) GetConfigsOk() (*[]OutputEmailIntakeConfig, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *OutputEmailIntake) SetConfigs(v []OutputEmailIntakeConfig)`

SetConfigs sets Configs field to given value.

### HasConfigs

`func (o *OutputEmailIntake) HasConfigs() bool`

HasConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


