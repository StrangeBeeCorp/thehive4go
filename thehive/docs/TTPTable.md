# TTPTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Columns** | Pointer to [**[]TTPField**](TTPField.md) |  | [optional] 
**Filter** | Pointer to [**Filter**](Filter.md) |  | [optional] 
**Sort** | Pointer to [**[]AlertListSortInner**](AlertListSortInner.md) |  | [optional] 
**MaxElements** | Pointer to **int32** |  | [optional] 
**Kind** | [**Kind**](Kind.md) |  | 

## Methods

### NewTTPTable

`func NewTTPTable(kind Kind, ) *TTPTable`

NewTTPTable instantiates a new TTPTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTTPTableWithDefaults

`func NewTTPTableWithDefaults() *TTPTable`

NewTTPTableWithDefaults instantiates a new TTPTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *TTPTable) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TTPTable) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TTPTable) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TTPTable) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetColumns

`func (o *TTPTable) GetColumns() []TTPField`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *TTPTable) GetColumnsOk() (*[]TTPField, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *TTPTable) SetColumns(v []TTPField)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *TTPTable) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetFilter

`func (o *TTPTable) GetFilter() Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *TTPTable) GetFilterOk() (*Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *TTPTable) SetFilter(v Filter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *TTPTable) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSort

`func (o *TTPTable) GetSort() []AlertListSortInner`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *TTPTable) GetSortOk() (*[]AlertListSortInner, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *TTPTable) SetSort(v []AlertListSortInner)`

SetSort sets Sort field to given value.

### HasSort

`func (o *TTPTable) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetMaxElements

`func (o *TTPTable) GetMaxElements() int32`

GetMaxElements returns the MaxElements field if non-nil, zero value otherwise.

### GetMaxElementsOk

`func (o *TTPTable) GetMaxElementsOk() (*int32, bool)`

GetMaxElementsOk returns a tuple with the MaxElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxElements

`func (o *TTPTable) SetMaxElements(v int32)`

SetMaxElements sets MaxElements field to given value.

### HasMaxElements

`func (o *TTPTable) HasMaxElements() bool`

HasMaxElements returns a boolean if a field has been set.

### GetKind

`func (o *TTPTable) GetKind() Kind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TTPTable) GetKindOk() (*Kind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TTPTable) SetKind(v Kind)`

SetKind sets Kind field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


