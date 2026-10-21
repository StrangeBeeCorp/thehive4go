# InputUpdateCaseStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to **int32** | New position in the UI dropdown list. Set to &#x60;null&#x60; to remove the custom ordering. | [optional]
**Description** | Pointer to **string** | New description. Set to &#x60;null&#x60; to remove the existing description. | [optional]
**Colour** | Pointer to **string** | New hex color code. Set to &#x60;null&#x60; to remove the custom color. | [optional]
**Hidden** | Pointer to **bool** | When &#x60;true&#x60;, hides the status from the UI dropdown. Hidden statuses remain accessible through the API. | [optional]
**RequiresImpact** | Pointer to **bool** | When &#x60;true&#x60;, closing a case with this status requires an &#x60;impactStatus&#x60;. Only a status in the *Closed* stage can require an impact, and you can&#39;t turn it off for the predefined &#x60;TruePositive&#x60; status. | [optional]

## Methods

### NewInputUpdateCaseStatus

`func NewInputUpdateCaseStatus() *InputUpdateCaseStatus`

NewInputUpdateCaseStatus instantiates a new InputUpdateCaseStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputUpdateCaseStatusWithDefaults

`func NewInputUpdateCaseStatusWithDefaults() *InputUpdateCaseStatus`

NewInputUpdateCaseStatusWithDefaults instantiates a new InputUpdateCaseStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *InputUpdateCaseStatus) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *InputUpdateCaseStatus) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *InputUpdateCaseStatus) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *InputUpdateCaseStatus) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetDescription

`func (o *InputUpdateCaseStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputUpdateCaseStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputUpdateCaseStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InputUpdateCaseStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetColour

`func (o *InputUpdateCaseStatus) GetColour() string`

GetColour returns the Colour field if non-nil, zero value otherwise.

### GetColourOk

`func (o *InputUpdateCaseStatus) GetColourOk() (*string, bool)`

GetColourOk returns a tuple with the Colour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColour

`func (o *InputUpdateCaseStatus) SetColour(v string)`

SetColour sets Colour field to given value.

### HasColour

`func (o *InputUpdateCaseStatus) HasColour() bool`

HasColour returns a boolean if a field has been set.

### GetHidden

`func (o *InputUpdateCaseStatus) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *InputUpdateCaseStatus) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *InputUpdateCaseStatus) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *InputUpdateCaseStatus) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetRequiresImpact

`func (o *InputUpdateCaseStatus) GetRequiresImpact() bool`

GetRequiresImpact returns the RequiresImpact field if non-nil, zero value otherwise.

### GetRequiresImpactOk

`func (o *InputUpdateCaseStatus) GetRequiresImpactOk() (*bool, bool)`

GetRequiresImpactOk returns a tuple with the RequiresImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresImpact

`func (o *InputUpdateCaseStatus) SetRequiresImpact(v bool)`

SetRequiresImpact sets RequiresImpact field to given value.

### HasRequiresImpact

`func (o *InputUpdateCaseStatus) HasRequiresImpact() bool`

HasRequiresImpact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
