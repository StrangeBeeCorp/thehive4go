# OutputAlertStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnderscoreId** | **string** |  |
**UnderscoreType** | **string** |  |
**UnderscoreUpdatedAt** | Pointer to **int64** |  | [optional]
**UnderscoreUpdatedBy** | Pointer to **string** |  | [optional]
**UnderscoreCreatedAt** | **int64** |  |
**UnderscoreCreatedBy** | **string** |  |
**Value** | **string** |  |
**Stage** | **string** |  |
**Order** | Pointer to **int32** |  | [optional]
**Description** | Pointer to **string** |  | [optional]
**Colour** | Pointer to **string** |  | [optional]
**Hidden** | **bool** |  |
**ExtraData** | **map[string]interface{}** |  |

## Methods

### NewOutputAlertStatus

`func NewOutputAlertStatus(underscoreId string, underscoreType string, underscoreCreatedAt int64, underscoreCreatedBy string, value string, stage string, hidden bool, extraData map[string]interface{}, ) *OutputAlertStatus`

NewOutputAlertStatus instantiates a new OutputAlertStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputAlertStatusWithDefaults

`func NewOutputAlertStatusWithDefaults() *OutputAlertStatus`

NewOutputAlertStatusWithDefaults instantiates a new OutputAlertStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnderscoreId

`func (o *OutputAlertStatus) GetUnderscoreId() string`

GetUnderscoreId returns the UnderscoreId field if non-nil, zero value otherwise.

### GetUnderscoreIdOk

`func (o *OutputAlertStatus) GetUnderscoreIdOk() (*string, bool)`

GetUnderscoreIdOk returns a tuple with the UnderscoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreId

`func (o *OutputAlertStatus) SetUnderscoreId(v string)`

SetUnderscoreId sets UnderscoreId field to given value.


### GetUnderscoreType

`func (o *OutputAlertStatus) GetUnderscoreType() string`

GetUnderscoreType returns the UnderscoreType field if non-nil, zero value otherwise.

### GetUnderscoreTypeOk

`func (o *OutputAlertStatus) GetUnderscoreTypeOk() (*string, bool)`

GetUnderscoreTypeOk returns a tuple with the UnderscoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreType

`func (o *OutputAlertStatus) SetUnderscoreType(v string)`

SetUnderscoreType sets UnderscoreType field to given value.


### GetUnderscoreUpdatedAt

`func (o *OutputAlertStatus) GetUnderscoreUpdatedAt() int64`

GetUnderscoreUpdatedAt returns the UnderscoreUpdatedAt field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedAtOk

`func (o *OutputAlertStatus) GetUnderscoreUpdatedAtOk() (*int64, bool)`

GetUnderscoreUpdatedAtOk returns a tuple with the UnderscoreUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedAt

`func (o *OutputAlertStatus) SetUnderscoreUpdatedAt(v int64)`

SetUnderscoreUpdatedAt sets UnderscoreUpdatedAt field to given value.

### HasUnderscoreUpdatedAt

`func (o *OutputAlertStatus) HasUnderscoreUpdatedAt() bool`

HasUnderscoreUpdatedAt returns a boolean if a field has been set.

### GetUnderscoreUpdatedBy

`func (o *OutputAlertStatus) GetUnderscoreUpdatedBy() string`

GetUnderscoreUpdatedBy returns the UnderscoreUpdatedBy field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedByOk

`func (o *OutputAlertStatus) GetUnderscoreUpdatedByOk() (*string, bool)`

GetUnderscoreUpdatedByOk returns a tuple with the UnderscoreUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedBy

`func (o *OutputAlertStatus) SetUnderscoreUpdatedBy(v string)`

SetUnderscoreUpdatedBy sets UnderscoreUpdatedBy field to given value.

### HasUnderscoreUpdatedBy

`func (o *OutputAlertStatus) HasUnderscoreUpdatedBy() bool`

HasUnderscoreUpdatedBy returns a boolean if a field has been set.

### GetUnderscoreCreatedAt

`func (o *OutputAlertStatus) GetUnderscoreCreatedAt() int64`

GetUnderscoreCreatedAt returns the UnderscoreCreatedAt field if non-nil, zero value otherwise.

### GetUnderscoreCreatedAtOk

`func (o *OutputAlertStatus) GetUnderscoreCreatedAtOk() (*int64, bool)`

GetUnderscoreCreatedAtOk returns a tuple with the UnderscoreCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedAt

`func (o *OutputAlertStatus) SetUnderscoreCreatedAt(v int64)`

SetUnderscoreCreatedAt sets UnderscoreCreatedAt field to given value.


### GetUnderscoreCreatedBy

`func (o *OutputAlertStatus) GetUnderscoreCreatedBy() string`

GetUnderscoreCreatedBy returns the UnderscoreCreatedBy field if non-nil, zero value otherwise.

### GetUnderscoreCreatedByOk

`func (o *OutputAlertStatus) GetUnderscoreCreatedByOk() (*string, bool)`

GetUnderscoreCreatedByOk returns a tuple with the UnderscoreCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedBy

`func (o *OutputAlertStatus) SetUnderscoreCreatedBy(v string)`

SetUnderscoreCreatedBy sets UnderscoreCreatedBy field to given value.


### GetValue

`func (o *OutputAlertStatus) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OutputAlertStatus) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OutputAlertStatus) SetValue(v string)`

SetValue sets Value field to given value.


### GetStage

`func (o *OutputAlertStatus) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *OutputAlertStatus) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *OutputAlertStatus) SetStage(v string)`

SetStage sets Stage field to given value.


### GetOrder

`func (o *OutputAlertStatus) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *OutputAlertStatus) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *OutputAlertStatus) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *OutputAlertStatus) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetDescription

`func (o *OutputAlertStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OutputAlertStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OutputAlertStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OutputAlertStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetColour

`func (o *OutputAlertStatus) GetColour() string`

GetColour returns the Colour field if non-nil, zero value otherwise.

### GetColourOk

`func (o *OutputAlertStatus) GetColourOk() (*string, bool)`

GetColourOk returns a tuple with the Colour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColour

`func (o *OutputAlertStatus) SetColour(v string)`

SetColour sets Colour field to given value.

### HasColour

`func (o *OutputAlertStatus) HasColour() bool`

HasColour returns a boolean if a field has been set.

### GetHidden

`func (o *OutputAlertStatus) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *OutputAlertStatus) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *OutputAlertStatus) SetHidden(v bool)`

SetHidden sets Hidden field to given value.


### GetExtraData

`func (o *OutputAlertStatus) GetExtraData() map[string]interface{}`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *OutputAlertStatus) GetExtraDataOk() (*map[string]interface{}, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *OutputAlertStatus) SetExtraData(v map[string]interface{})`

SetExtraData sets ExtraData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
