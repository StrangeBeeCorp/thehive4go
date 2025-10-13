# InputCustomEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **int32** |  | 
**EndDate** | Pointer to **int32** |  | [optional] 
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewInputCustomEvent

`func NewInputCustomEvent(date int32, title string, ) *InputCustomEvent`

NewInputCustomEvent instantiates a new InputCustomEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputCustomEventWithDefaults

`func NewInputCustomEventWithDefaults() *InputCustomEvent`

NewInputCustomEventWithDefaults instantiates a new InputCustomEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *InputCustomEvent) GetDate() int32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *InputCustomEvent) GetDateOk() (*int32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *InputCustomEvent) SetDate(v int32)`

SetDate sets Date field to given value.


### GetEndDate

`func (o *InputCustomEvent) GetEndDate() int32`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *InputCustomEvent) GetEndDateOk() (*int32, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *InputCustomEvent) SetEndDate(v int32)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *InputCustomEvent) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetTitle

`func (o *InputCustomEvent) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InputCustomEvent) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InputCustomEvent) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *InputCustomEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputCustomEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputCustomEvent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InputCustomEvent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


