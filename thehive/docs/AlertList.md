# AlertList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Title displayed above the widget. | [optional]
**Fields** | Pointer to [**[]AlertField**](AlertField.md) | Alert fields to display for each item in the list. | [optional]
**Filter** | Pointer to [**Filter**](Filter.md) | Filter to apply to the alert list. See the [Query and Export](https://docs.strangebee.com/thehive/api-docs/#tag/Query-and-Export) section for the filter syntax. | [optional]
**Sort** | Pointer to [**[]AlertListSortInner**](AlertListSortInner.md) | Sort order for the alert list. See the [Query and Export](https://docs.strangebee.com/thehive/api-docs/#tag/Query-and-Export) section for the sort syntax. | [optional]
**MaxElements** | Pointer to **int32** | Maximum number of alerts to include in the widget. If omitted, the global &#x60;maxElements&#x60; limit applies. | [optional]
**Kind** | [**Kind**](Kind.md) |  |

## Methods

### NewAlertList

`func NewAlertList(kind Kind, ) *AlertList`

NewAlertList instantiates a new AlertList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertListWithDefaults

`func NewAlertListWithDefaults() *AlertList`

NewAlertListWithDefaults instantiates a new AlertList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *AlertList) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AlertList) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AlertList) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AlertList) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFields

`func (o *AlertList) GetFields() []AlertField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AlertList) GetFieldsOk() (*[]AlertField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AlertList) SetFields(v []AlertField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *AlertList) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFilter

`func (o *AlertList) GetFilter() Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AlertList) GetFilterOk() (*Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AlertList) SetFilter(v Filter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *AlertList) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSort

`func (o *AlertList) GetSort() []AlertListSortInner`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *AlertList) GetSortOk() (*[]AlertListSortInner, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *AlertList) SetSort(v []AlertListSortInner)`

SetSort sets Sort field to given value.

### HasSort

`func (o *AlertList) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetMaxElements

`func (o *AlertList) GetMaxElements() int32`

GetMaxElements returns the MaxElements field if non-nil, zero value otherwise.

### GetMaxElementsOk

`func (o *AlertList) GetMaxElementsOk() (*int32, bool)`

GetMaxElementsOk returns a tuple with the MaxElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxElements

`func (o *AlertList) SetMaxElements(v int32)`

SetMaxElements sets MaxElements field to given value.

### HasMaxElements

`func (o *AlertList) HasMaxElements() bool`

HasMaxElements returns a boolean if a field has been set.

### GetKind

`func (o *AlertList) GetKind() Kind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AlertList) GetKindOk() (*Kind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AlertList) SetKind(v Kind)`

SetKind sets Kind field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
