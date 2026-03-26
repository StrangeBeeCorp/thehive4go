# OutputListView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnderscoreId** | **string** |  | 
**UnderscoreType** | **string** |  | 
**UnderscoreCreatedAt** | **int64** |  | 
**UnderscoreCreatedBy** | **string** |  | 
**UnderscoreUpdatedAt** | Pointer to **int64** |  | [optional] 
**UnderscoreUpdatedBy** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Entity** | **string** |  | 
**Filter** | **map[string]interface{}** |  | 
**ListOptions** | [**ListOptions**](ListOptions.md) |  | 
**SortList** | Pointer to **[]string** |  | [optional] 
**ShowColumns** | Pointer to **[]string** |  | [optional] 
**IsShared** | **bool** |  | 

## Methods

### NewOutputListView

`func NewOutputListView(underscoreId string, underscoreType string, underscoreCreatedAt int64, underscoreCreatedBy string, name string, entity string, filter map[string]interface{}, listOptions ListOptions, isShared bool, ) *OutputListView`

NewOutputListView instantiates a new OutputListView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputListViewWithDefaults

`func NewOutputListViewWithDefaults() *OutputListView`

NewOutputListViewWithDefaults instantiates a new OutputListView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnderscoreId

`func (o *OutputListView) GetUnderscoreId() string`

GetUnderscoreId returns the UnderscoreId field if non-nil, zero value otherwise.

### GetUnderscoreIdOk

`func (o *OutputListView) GetUnderscoreIdOk() (*string, bool)`

GetUnderscoreIdOk returns a tuple with the UnderscoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreId

`func (o *OutputListView) SetUnderscoreId(v string)`

SetUnderscoreId sets UnderscoreId field to given value.


### GetUnderscoreType

`func (o *OutputListView) GetUnderscoreType() string`

GetUnderscoreType returns the UnderscoreType field if non-nil, zero value otherwise.

### GetUnderscoreTypeOk

`func (o *OutputListView) GetUnderscoreTypeOk() (*string, bool)`

GetUnderscoreTypeOk returns a tuple with the UnderscoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreType

`func (o *OutputListView) SetUnderscoreType(v string)`

SetUnderscoreType sets UnderscoreType field to given value.


### GetUnderscoreCreatedAt

`func (o *OutputListView) GetUnderscoreCreatedAt() int64`

GetUnderscoreCreatedAt returns the UnderscoreCreatedAt field if non-nil, zero value otherwise.

### GetUnderscoreCreatedAtOk

`func (o *OutputListView) GetUnderscoreCreatedAtOk() (*int64, bool)`

GetUnderscoreCreatedAtOk returns a tuple with the UnderscoreCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedAt

`func (o *OutputListView) SetUnderscoreCreatedAt(v int64)`

SetUnderscoreCreatedAt sets UnderscoreCreatedAt field to given value.


### GetUnderscoreCreatedBy

`func (o *OutputListView) GetUnderscoreCreatedBy() string`

GetUnderscoreCreatedBy returns the UnderscoreCreatedBy field if non-nil, zero value otherwise.

### GetUnderscoreCreatedByOk

`func (o *OutputListView) GetUnderscoreCreatedByOk() (*string, bool)`

GetUnderscoreCreatedByOk returns a tuple with the UnderscoreCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedBy

`func (o *OutputListView) SetUnderscoreCreatedBy(v string)`

SetUnderscoreCreatedBy sets UnderscoreCreatedBy field to given value.


### GetUnderscoreUpdatedAt

`func (o *OutputListView) GetUnderscoreUpdatedAt() int64`

GetUnderscoreUpdatedAt returns the UnderscoreUpdatedAt field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedAtOk

`func (o *OutputListView) GetUnderscoreUpdatedAtOk() (*int64, bool)`

GetUnderscoreUpdatedAtOk returns a tuple with the UnderscoreUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedAt

`func (o *OutputListView) SetUnderscoreUpdatedAt(v int64)`

SetUnderscoreUpdatedAt sets UnderscoreUpdatedAt field to given value.

### HasUnderscoreUpdatedAt

`func (o *OutputListView) HasUnderscoreUpdatedAt() bool`

HasUnderscoreUpdatedAt returns a boolean if a field has been set.

### GetUnderscoreUpdatedBy

`func (o *OutputListView) GetUnderscoreUpdatedBy() string`

GetUnderscoreUpdatedBy returns the UnderscoreUpdatedBy field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedByOk

`func (o *OutputListView) GetUnderscoreUpdatedByOk() (*string, bool)`

GetUnderscoreUpdatedByOk returns a tuple with the UnderscoreUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedBy

`func (o *OutputListView) SetUnderscoreUpdatedBy(v string)`

SetUnderscoreUpdatedBy sets UnderscoreUpdatedBy field to given value.

### HasUnderscoreUpdatedBy

`func (o *OutputListView) HasUnderscoreUpdatedBy() bool`

HasUnderscoreUpdatedBy returns a boolean if a field has been set.

### GetName

`func (o *OutputListView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OutputListView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OutputListView) SetName(v string)`

SetName sets Name field to given value.


### GetEntity

`func (o *OutputListView) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *OutputListView) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *OutputListView) SetEntity(v string)`

SetEntity sets Entity field to given value.


### GetFilter

`func (o *OutputListView) GetFilter() map[string]interface{}`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *OutputListView) GetFilterOk() (*map[string]interface{}, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *OutputListView) SetFilter(v map[string]interface{})`

SetFilter sets Filter field to given value.


### GetListOptions

`func (o *OutputListView) GetListOptions() ListOptions`

GetListOptions returns the ListOptions field if non-nil, zero value otherwise.

### GetListOptionsOk

`func (o *OutputListView) GetListOptionsOk() (*ListOptions, bool)`

GetListOptionsOk returns a tuple with the ListOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOptions

`func (o *OutputListView) SetListOptions(v ListOptions)`

SetListOptions sets ListOptions field to given value.


### GetSortList

`func (o *OutputListView) GetSortList() []string`

GetSortList returns the SortList field if non-nil, zero value otherwise.

### GetSortListOk

`func (o *OutputListView) GetSortListOk() (*[]string, bool)`

GetSortListOk returns a tuple with the SortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortList

`func (o *OutputListView) SetSortList(v []string)`

SetSortList sets SortList field to given value.

### HasSortList

`func (o *OutputListView) HasSortList() bool`

HasSortList returns a boolean if a field has been set.

### GetShowColumns

`func (o *OutputListView) GetShowColumns() []string`

GetShowColumns returns the ShowColumns field if non-nil, zero value otherwise.

### GetShowColumnsOk

`func (o *OutputListView) GetShowColumnsOk() (*[]string, bool)`

GetShowColumnsOk returns a tuple with the ShowColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowColumns

`func (o *OutputListView) SetShowColumns(v []string)`

SetShowColumns sets ShowColumns field to given value.

### HasShowColumns

`func (o *OutputListView) HasShowColumns() bool`

HasShowColumns returns a boolean if a field has been set.

### GetIsShared

`func (o *OutputListView) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *OutputListView) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *OutputListView) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


