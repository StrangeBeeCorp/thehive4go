# ObservableTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Columns** | Pointer to [**[]ObservableField**](ObservableField.md) |  | [optional] 
**Filter** | Pointer to [**Filter**](Filter.md) |  | [optional] 
**Sort** | Pointer to [**[]AlertListSortInner**](AlertListSortInner.md) |  | [optional] 
**ProtectData** | **bool** |  | 
**MaxElements** | Pointer to **int32** |  | [optional] 
**Kind** | [**Kind**](Kind.md) |  | 

## Methods

### NewObservableTable

`func NewObservableTable(protectData bool, kind Kind, ) *ObservableTable`

NewObservableTable instantiates a new ObservableTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObservableTableWithDefaults

`func NewObservableTableWithDefaults() *ObservableTable`

NewObservableTableWithDefaults instantiates a new ObservableTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ObservableTable) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ObservableTable) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ObservableTable) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ObservableTable) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetColumns

`func (o *ObservableTable) GetColumns() []ObservableField`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *ObservableTable) GetColumnsOk() (*[]ObservableField, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *ObservableTable) SetColumns(v []ObservableField)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *ObservableTable) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetFilter

`func (o *ObservableTable) GetFilter() Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ObservableTable) GetFilterOk() (*Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ObservableTable) SetFilter(v Filter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ObservableTable) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSort

`func (o *ObservableTable) GetSort() []AlertListSortInner`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *ObservableTable) GetSortOk() (*[]AlertListSortInner, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *ObservableTable) SetSort(v []AlertListSortInner)`

SetSort sets Sort field to given value.

### HasSort

`func (o *ObservableTable) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetProtectData

`func (o *ObservableTable) GetProtectData() bool`

GetProtectData returns the ProtectData field if non-nil, zero value otherwise.

### GetProtectDataOk

`func (o *ObservableTable) GetProtectDataOk() (*bool, bool)`

GetProtectDataOk returns a tuple with the ProtectData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectData

`func (o *ObservableTable) SetProtectData(v bool)`

SetProtectData sets ProtectData field to given value.


### GetMaxElements

`func (o *ObservableTable) GetMaxElements() int32`

GetMaxElements returns the MaxElements field if non-nil, zero value otherwise.

### GetMaxElementsOk

`func (o *ObservableTable) GetMaxElementsOk() (*int32, bool)`

GetMaxElementsOk returns a tuple with the MaxElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxElements

`func (o *ObservableTable) SetMaxElements(v int32)`

SetMaxElements sets MaxElements field to given value.

### HasMaxElements

`func (o *ObservableTable) HasMaxElements() bool`

HasMaxElements returns a boolean if a field has been set.

### GetKind

`func (o *ObservableTable) GetKind() Kind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ObservableTable) GetKindOk() (*Kind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ObservableTable) SetKind(v Kind)`

SetKind sets Kind field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


