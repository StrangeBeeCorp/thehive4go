# AlertTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Columns** | Pointer to [**[]AlertField**](AlertField.md) |  | [optional] 
**Filter** | Pointer to [**Filter**](Filter.md) |  | [optional] 
**Sort** | Pointer to [**[]AlertListSortInner**](AlertListSortInner.md) |  | [optional] 
**MaxElements** | Pointer to **int32** |  | [optional] 
**Kind** | [**Kind**](Kind.md) |  | 

## Methods

### NewAlertTable

`func NewAlertTable(kind Kind, ) *AlertTable`

NewAlertTable instantiates a new AlertTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertTableWithDefaults

`func NewAlertTableWithDefaults() *AlertTable`

NewAlertTableWithDefaults instantiates a new AlertTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *AlertTable) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AlertTable) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AlertTable) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AlertTable) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetColumns

`func (o *AlertTable) GetColumns() []AlertField`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *AlertTable) GetColumnsOk() (*[]AlertField, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *AlertTable) SetColumns(v []AlertField)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *AlertTable) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetFilter

`func (o *AlertTable) GetFilter() Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AlertTable) GetFilterOk() (*Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AlertTable) SetFilter(v Filter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *AlertTable) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSort

`func (o *AlertTable) GetSort() []AlertListSortInner`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *AlertTable) GetSortOk() (*[]AlertListSortInner, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *AlertTable) SetSort(v []AlertListSortInner)`

SetSort sets Sort field to given value.

### HasSort

`func (o *AlertTable) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetMaxElements

`func (o *AlertTable) GetMaxElements() int32`

GetMaxElements returns the MaxElements field if non-nil, zero value otherwise.

### GetMaxElementsOk

`func (o *AlertTable) GetMaxElementsOk() (*int32, bool)`

GetMaxElementsOk returns a tuple with the MaxElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxElements

`func (o *AlertTable) SetMaxElements(v int32)`

SetMaxElements sets MaxElements field to given value.

### HasMaxElements

`func (o *AlertTable) HasMaxElements() bool`

HasMaxElements returns a boolean if a field has been set.

### GetKind

`func (o *AlertTable) GetKind() Kind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AlertTable) GetKindOk() (*Kind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AlertTable) SetKind(v Kind)`

SetKind sets Kind field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


