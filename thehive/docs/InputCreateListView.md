# InputCreateListView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Entity** | [**InputListViewEntity**](InputListViewEntity.md) |  | 
**Filter** | **map[string]interface{}** |  | 
**ListOptions** | [**ListOptions**](ListOptions.md) |  | 
**SortList** | Pointer to **[]string** |  | [optional] 
**ShowColumns** | Pointer to **[]string** |  | [optional] 
**IsShared** | **bool** |  | 

## Methods

### NewInputCreateListView

`func NewInputCreateListView(name string, entity InputListViewEntity, filter map[string]interface{}, listOptions ListOptions, isShared bool, ) *InputCreateListView`

NewInputCreateListView instantiates a new InputCreateListView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputCreateListViewWithDefaults

`func NewInputCreateListViewWithDefaults() *InputCreateListView`

NewInputCreateListViewWithDefaults instantiates a new InputCreateListView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InputCreateListView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputCreateListView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputCreateListView) SetName(v string)`

SetName sets Name field to given value.


### GetEntity

`func (o *InputCreateListView) GetEntity() InputListViewEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *InputCreateListView) GetEntityOk() (*InputListViewEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *InputCreateListView) SetEntity(v InputListViewEntity)`

SetEntity sets Entity field to given value.


### GetFilter

`func (o *InputCreateListView) GetFilter() map[string]interface{}`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *InputCreateListView) GetFilterOk() (*map[string]interface{}, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *InputCreateListView) SetFilter(v map[string]interface{})`

SetFilter sets Filter field to given value.


### GetListOptions

`func (o *InputCreateListView) GetListOptions() ListOptions`

GetListOptions returns the ListOptions field if non-nil, zero value otherwise.

### GetListOptionsOk

`func (o *InputCreateListView) GetListOptionsOk() (*ListOptions, bool)`

GetListOptionsOk returns a tuple with the ListOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOptions

`func (o *InputCreateListView) SetListOptions(v ListOptions)`

SetListOptions sets ListOptions field to given value.


### GetSortList

`func (o *InputCreateListView) GetSortList() []string`

GetSortList returns the SortList field if non-nil, zero value otherwise.

### GetSortListOk

`func (o *InputCreateListView) GetSortListOk() (*[]string, bool)`

GetSortListOk returns a tuple with the SortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortList

`func (o *InputCreateListView) SetSortList(v []string)`

SetSortList sets SortList field to given value.

### HasSortList

`func (o *InputCreateListView) HasSortList() bool`

HasSortList returns a boolean if a field has been set.

### GetShowColumns

`func (o *InputCreateListView) GetShowColumns() []string`

GetShowColumns returns the ShowColumns field if non-nil, zero value otherwise.

### GetShowColumnsOk

`func (o *InputCreateListView) GetShowColumnsOk() (*[]string, bool)`

GetShowColumnsOk returns a tuple with the ShowColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowColumns

`func (o *InputCreateListView) SetShowColumns(v []string)`

SetShowColumns sets ShowColumns field to given value.

### HasShowColumns

`func (o *InputCreateListView) HasShowColumns() bool`

HasShowColumns returns a boolean if a field has been set.

### GetIsShared

`func (o *InputCreateListView) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *InputCreateListView) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *InputCreateListView) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


