# InputCreateCaseStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | Name of the status | 
**Stage** | [**InputCaseStage**](InputCaseStage.md) |  | 
**Order** | Pointer to **int32** | Used to order the values in the ui | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Colour** | Pointer to **string** | Optional color to use in the ui | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 

## Methods

### NewInputCreateCaseStatus

`func NewInputCreateCaseStatus(value string, stage InputCaseStage, ) *InputCreateCaseStatus`

NewInputCreateCaseStatus instantiates a new InputCreateCaseStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputCreateCaseStatusWithDefaults

`func NewInputCreateCaseStatusWithDefaults() *InputCreateCaseStatus`

NewInputCreateCaseStatusWithDefaults instantiates a new InputCreateCaseStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *InputCreateCaseStatus) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InputCreateCaseStatus) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InputCreateCaseStatus) SetValue(v string)`

SetValue sets Value field to given value.


### GetStage

`func (o *InputCreateCaseStatus) GetStage() InputCaseStage`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *InputCreateCaseStatus) GetStageOk() (*InputCaseStage, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *InputCreateCaseStatus) SetStage(v InputCaseStage)`

SetStage sets Stage field to given value.


### GetOrder

`func (o *InputCreateCaseStatus) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *InputCreateCaseStatus) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *InputCreateCaseStatus) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *InputCreateCaseStatus) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetDescription

`func (o *InputCreateCaseStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputCreateCaseStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputCreateCaseStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InputCreateCaseStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetColour

`func (o *InputCreateCaseStatus) GetColour() string`

GetColour returns the Colour field if non-nil, zero value otherwise.

### GetColourOk

`func (o *InputCreateCaseStatus) GetColourOk() (*string, bool)`

GetColourOk returns a tuple with the Colour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColour

`func (o *InputCreateCaseStatus) SetColour(v string)`

SetColour sets Colour field to given value.

### HasColour

`func (o *InputCreateCaseStatus) HasColour() bool`

HasColour returns a boolean if a field has been set.

### GetHidden

`func (o *InputCreateCaseStatus) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *InputCreateCaseStatus) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *InputCreateCaseStatus) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *InputCreateCaseStatus) HasHidden() bool`

HasHidden returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


