# OutputProcedure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnderscoreId** | **string** |  |
**UnderscoreCreatedAt** | **int64** |  |
**UnderscoreCreatedBy** | **string** |  |
**UnderscoreUpdatedAt** | Pointer to **int64** |  | [optional]
**UnderscoreUpdatedBy** | Pointer to **string** |  | [optional]
**Description** | Pointer to **string** |  | [optional]
**OccurDate** | **int64** |  |
**PatternId** | Pointer to **string** |  | [optional]
**PatternName** | Pointer to **string** |  | [optional]
**Tactic** | Pointer to **string** |  | [optional]
**TacticLabel** | Pointer to **string** |  | [optional]
**ExtraData** | **map[string]interface{}** |  |

## Methods

### NewOutputProcedure

`func NewOutputProcedure(underscoreId string, underscoreCreatedAt int64, underscoreCreatedBy string, occurDate int64, extraData map[string]interface{}, ) *OutputProcedure`

NewOutputProcedure instantiates a new OutputProcedure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputProcedureWithDefaults

`func NewOutputProcedureWithDefaults() *OutputProcedure`

NewOutputProcedureWithDefaults instantiates a new OutputProcedure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnderscoreId

`func (o *OutputProcedure) GetUnderscoreId() string`

GetUnderscoreId returns the UnderscoreId field if non-nil, zero value otherwise.

### GetUnderscoreIdOk

`func (o *OutputProcedure) GetUnderscoreIdOk() (*string, bool)`

GetUnderscoreIdOk returns a tuple with the UnderscoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreId

`func (o *OutputProcedure) SetUnderscoreId(v string)`

SetUnderscoreId sets UnderscoreId field to given value.


### GetUnderscoreCreatedAt

`func (o *OutputProcedure) GetUnderscoreCreatedAt() int64`

GetUnderscoreCreatedAt returns the UnderscoreCreatedAt field if non-nil, zero value otherwise.

### GetUnderscoreCreatedAtOk

`func (o *OutputProcedure) GetUnderscoreCreatedAtOk() (*int64, bool)`

GetUnderscoreCreatedAtOk returns a tuple with the UnderscoreCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedAt

`func (o *OutputProcedure) SetUnderscoreCreatedAt(v int64)`

SetUnderscoreCreatedAt sets UnderscoreCreatedAt field to given value.


### GetUnderscoreCreatedBy

`func (o *OutputProcedure) GetUnderscoreCreatedBy() string`

GetUnderscoreCreatedBy returns the UnderscoreCreatedBy field if non-nil, zero value otherwise.

### GetUnderscoreCreatedByOk

`func (o *OutputProcedure) GetUnderscoreCreatedByOk() (*string, bool)`

GetUnderscoreCreatedByOk returns a tuple with the UnderscoreCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedBy

`func (o *OutputProcedure) SetUnderscoreCreatedBy(v string)`

SetUnderscoreCreatedBy sets UnderscoreCreatedBy field to given value.


### GetUnderscoreUpdatedAt

`func (o *OutputProcedure) GetUnderscoreUpdatedAt() int64`

GetUnderscoreUpdatedAt returns the UnderscoreUpdatedAt field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedAtOk

`func (o *OutputProcedure) GetUnderscoreUpdatedAtOk() (*int64, bool)`

GetUnderscoreUpdatedAtOk returns a tuple with the UnderscoreUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedAt

`func (o *OutputProcedure) SetUnderscoreUpdatedAt(v int64)`

SetUnderscoreUpdatedAt sets UnderscoreUpdatedAt field to given value.

### HasUnderscoreUpdatedAt

`func (o *OutputProcedure) HasUnderscoreUpdatedAt() bool`

HasUnderscoreUpdatedAt returns a boolean if a field has been set.

### GetUnderscoreUpdatedBy

`func (o *OutputProcedure) GetUnderscoreUpdatedBy() string`

GetUnderscoreUpdatedBy returns the UnderscoreUpdatedBy field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedByOk

`func (o *OutputProcedure) GetUnderscoreUpdatedByOk() (*string, bool)`

GetUnderscoreUpdatedByOk returns a tuple with the UnderscoreUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedBy

`func (o *OutputProcedure) SetUnderscoreUpdatedBy(v string)`

SetUnderscoreUpdatedBy sets UnderscoreUpdatedBy field to given value.

### HasUnderscoreUpdatedBy

`func (o *OutputProcedure) HasUnderscoreUpdatedBy() bool`

HasUnderscoreUpdatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *OutputProcedure) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OutputProcedure) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OutputProcedure) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OutputProcedure) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOccurDate

`func (o *OutputProcedure) GetOccurDate() int64`

GetOccurDate returns the OccurDate field if non-nil, zero value otherwise.

### GetOccurDateOk

`func (o *OutputProcedure) GetOccurDateOk() (*int64, bool)`

GetOccurDateOk returns a tuple with the OccurDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurDate

`func (o *OutputProcedure) SetOccurDate(v int64)`

SetOccurDate sets OccurDate field to given value.


### GetPatternId

`func (o *OutputProcedure) GetPatternId() string`

GetPatternId returns the PatternId field if non-nil, zero value otherwise.

### GetPatternIdOk

`func (o *OutputProcedure) GetPatternIdOk() (*string, bool)`

GetPatternIdOk returns a tuple with the PatternId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternId

`func (o *OutputProcedure) SetPatternId(v string)`

SetPatternId sets PatternId field to given value.

### HasPatternId

`func (o *OutputProcedure) HasPatternId() bool`

HasPatternId returns a boolean if a field has been set.

### GetPatternName

`func (o *OutputProcedure) GetPatternName() string`

GetPatternName returns the PatternName field if non-nil, zero value otherwise.

### GetPatternNameOk

`func (o *OutputProcedure) GetPatternNameOk() (*string, bool)`

GetPatternNameOk returns a tuple with the PatternName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternName

`func (o *OutputProcedure) SetPatternName(v string)`

SetPatternName sets PatternName field to given value.

### HasPatternName

`func (o *OutputProcedure) HasPatternName() bool`

HasPatternName returns a boolean if a field has been set.

### GetTactic

`func (o *OutputProcedure) GetTactic() string`

GetTactic returns the Tactic field if non-nil, zero value otherwise.

### GetTacticOk

`func (o *OutputProcedure) GetTacticOk() (*string, bool)`

GetTacticOk returns a tuple with the Tactic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTactic

`func (o *OutputProcedure) SetTactic(v string)`

SetTactic sets Tactic field to given value.

### HasTactic

`func (o *OutputProcedure) HasTactic() bool`

HasTactic returns a boolean if a field has been set.

### GetTacticLabel

`func (o *OutputProcedure) GetTacticLabel() string`

GetTacticLabel returns the TacticLabel field if non-nil, zero value otherwise.

### GetTacticLabelOk

`func (o *OutputProcedure) GetTacticLabelOk() (*string, bool)`

GetTacticLabelOk returns a tuple with the TacticLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTacticLabel

`func (o *OutputProcedure) SetTacticLabel(v string)`

SetTacticLabel sets TacticLabel field to given value.

### HasTacticLabel

`func (o *OutputProcedure) HasTacticLabel() bool`

HasTacticLabel returns a boolean if a field has been set.

### GetExtraData

`func (o *OutputProcedure) GetExtraData() map[string]interface{}`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *OutputProcedure) GetExtraDataOk() (*map[string]interface{}, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *OutputProcedure) SetExtraData(v map[string]interface{})`

SetExtraData sets ExtraData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
