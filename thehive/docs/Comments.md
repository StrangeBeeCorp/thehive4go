# Comments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Filter** | Pointer to [**Filter**](Filter.md) |  | [optional] 
**Sort** | Pointer to [**[]AlertListSortInner**](AlertListSortInner.md) |  | [optional] 
**MaxElements** | Pointer to **int32** |  | [optional] 
**Kind** | [**Kind**](Kind.md) |  | 

## Methods

### NewComments

`func NewComments(kind Kind, ) *Comments`

NewComments instantiates a new Comments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentsWithDefaults

`func NewCommentsWithDefaults() *Comments`

NewCommentsWithDefaults instantiates a new Comments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *Comments) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Comments) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Comments) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Comments) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFilter

`func (o *Comments) GetFilter() Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *Comments) GetFilterOk() (*Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *Comments) SetFilter(v Filter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *Comments) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSort

`func (o *Comments) GetSort() []AlertListSortInner`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *Comments) GetSortOk() (*[]AlertListSortInner, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *Comments) SetSort(v []AlertListSortInner)`

SetSort sets Sort field to given value.

### HasSort

`func (o *Comments) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetMaxElements

`func (o *Comments) GetMaxElements() int32`

GetMaxElements returns the MaxElements field if non-nil, zero value otherwise.

### GetMaxElementsOk

`func (o *Comments) GetMaxElementsOk() (*int32, bool)`

GetMaxElementsOk returns a tuple with the MaxElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxElements

`func (o *Comments) SetMaxElements(v int32)`

SetMaxElements sets MaxElements field to given value.

### HasMaxElements

`func (o *Comments) HasMaxElements() bool`

HasMaxElements returns a boolean if a field has been set.

### GetKind

`func (o *Comments) GetKind() Kind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Comments) GetKindOk() (*Kind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Comments) SetKind(v Kind)`

SetKind sets Kind field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


