# InputEmailIntake

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | To disable the whole module | [optional] [default to true]
**Interval** | **string** |  | 

## Methods

### NewInputEmailIntake

`func NewInputEmailIntake(interval string, ) *InputEmailIntake`

NewInputEmailIntake instantiates a new InputEmailIntake object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputEmailIntakeWithDefaults

`func NewInputEmailIntakeWithDefaults() *InputEmailIntake`

NewInputEmailIntakeWithDefaults instantiates a new InputEmailIntake object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InputEmailIntake) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InputEmailIntake) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InputEmailIntake) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InputEmailIntake) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetInterval

`func (o *InputEmailIntake) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *InputEmailIntake) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *InputEmailIntake) SetInterval(v string)`

SetInterval sets Interval field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


