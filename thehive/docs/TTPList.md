# TTPList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Title displayed above the widget. | [optional]
**Fields** | Pointer to [**[]TTPField**](TTPField.md) | TTP fields to display for each item in the list. | [optional]
**Filter** | Pointer to [**Filter**](Filter.md) | Filter to apply to the TTP list. See the [Query and Export](https://docs.strangebee.com/thehive/api-docs/#tag/Query-and-Export) section for the filter syntax. | [optional]
**Sort** | Pointer to [**[]AlertListSortInner**](AlertListSortInner.md) | Sort order for the TTP list. See the [Query and Export](https://docs.strangebee.com/thehive/api-docs/#tag/Query-and-Export) section for the sort syntax. | [optional]
**MaxElements** | Pointer to **int32** | Maximum number of TTPs to include in the widget. If omitted, the global &#x60;maxElements&#x60; limit applies. | [optional]
**Kind** | [**Kind**](Kind.md) |  |

## Methods

### NewTTPList

`func NewTTPList(kind Kind, ) *TTPList`

NewTTPList instantiates a new TTPList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTTPListWithDefaults

`func NewTTPListWithDefaults() *TTPList`

NewTTPListWithDefaults instantiates a new TTPList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *TTPList) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TTPList) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TTPList) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TTPList) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFields

`func (o *TTPList) GetFields() []TTPField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *TTPList) GetFieldsOk() (*[]TTPField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *TTPList) SetFields(v []TTPField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *TTPList) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFilter

`func (o *TTPList) GetFilter() Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *TTPList) GetFilterOk() (*Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *TTPList) SetFilter(v Filter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *TTPList) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSort

`func (o *TTPList) GetSort() []AlertListSortInner`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *TTPList) GetSortOk() (*[]AlertListSortInner, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *TTPList) SetSort(v []AlertListSortInner)`

SetSort sets Sort field to given value.

### HasSort

`func (o *TTPList) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetMaxElements

`func (o *TTPList) GetMaxElements() int32`

GetMaxElements returns the MaxElements field if non-nil, zero value otherwise.

### GetMaxElementsOk

`func (o *TTPList) GetMaxElementsOk() (*int32, bool)`

GetMaxElementsOk returns a tuple with the MaxElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxElements

`func (o *TTPList) SetMaxElements(v int32)`

SetMaxElements sets MaxElements field to given value.

### HasMaxElements

`func (o *TTPList) HasMaxElements() bool`

HasMaxElements returns a boolean if a field has been set.

### GetKind

`func (o *TTPList) GetKind() Kind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TTPList) GetKindOk() (*Kind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TTPList) SetKind(v Kind)`

SetKind sets Kind field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
