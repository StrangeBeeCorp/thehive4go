# ObservableList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Fields** | Pointer to [**[]ObservableField**](ObservableField.md) |  | [optional] 
**Filter** | Pointer to [**Filter**](Filter.md) |  | [optional] 
**Sort** | Pointer to [**[]AlertListSortInner**](AlertListSortInner.md) |  | [optional] 
**ProtectData** | **bool** |  | 
**MaxElements** | Pointer to **int32** |  | [optional] 
**Kind** | [**Kind**](Kind.md) |  | 

## Methods

### NewObservableList

`func NewObservableList(protectData bool, kind Kind, ) *ObservableList`

NewObservableList instantiates a new ObservableList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObservableListWithDefaults

`func NewObservableListWithDefaults() *ObservableList`

NewObservableListWithDefaults instantiates a new ObservableList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ObservableList) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ObservableList) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ObservableList) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ObservableList) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFields

`func (o *ObservableList) GetFields() []ObservableField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ObservableList) GetFieldsOk() (*[]ObservableField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ObservableList) SetFields(v []ObservableField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ObservableList) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFilter

`func (o *ObservableList) GetFilter() Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ObservableList) GetFilterOk() (*Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ObservableList) SetFilter(v Filter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ObservableList) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSort

`func (o *ObservableList) GetSort() []AlertListSortInner`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *ObservableList) GetSortOk() (*[]AlertListSortInner, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *ObservableList) SetSort(v []AlertListSortInner)`

SetSort sets Sort field to given value.

### HasSort

`func (o *ObservableList) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetProtectData

`func (o *ObservableList) GetProtectData() bool`

GetProtectData returns the ProtectData field if non-nil, zero value otherwise.

### GetProtectDataOk

`func (o *ObservableList) GetProtectDataOk() (*bool, bool)`

GetProtectDataOk returns a tuple with the ProtectData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectData

`func (o *ObservableList) SetProtectData(v bool)`

SetProtectData sets ProtectData field to given value.


### GetMaxElements

`func (o *ObservableList) GetMaxElements() int32`

GetMaxElements returns the MaxElements field if non-nil, zero value otherwise.

### GetMaxElementsOk

`func (o *ObservableList) GetMaxElementsOk() (*int32, bool)`

GetMaxElementsOk returns a tuple with the MaxElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxElements

`func (o *ObservableList) SetMaxElements(v int32)`

SetMaxElements sets MaxElements field to given value.

### HasMaxElements

`func (o *ObservableList) HasMaxElements() bool`

HasMaxElements returns a boolean if a field has been set.

### GetKind

`func (o *ObservableList) GetKind() Kind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ObservableList) GetKindOk() (*Kind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ObservableList) SetKind(v Kind)`

SetKind sets Kind field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


