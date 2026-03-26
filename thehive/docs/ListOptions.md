# ListOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemsPerPage** | **int32** |  | 
**ListAsGroup** | **bool** |  | 
**AutoRefresh** | **bool** |  | 
**StatsIsOpen** | **bool** |  | 

## Methods

### NewListOptions

`func NewListOptions(itemsPerPage int32, listAsGroup bool, autoRefresh bool, statsIsOpen bool, ) *ListOptions`

NewListOptions instantiates a new ListOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOptionsWithDefaults

`func NewListOptionsWithDefaults() *ListOptions`

NewListOptionsWithDefaults instantiates a new ListOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemsPerPage

`func (o *ListOptions) GetItemsPerPage() int32`

GetItemsPerPage returns the ItemsPerPage field if non-nil, zero value otherwise.

### GetItemsPerPageOk

`func (o *ListOptions) GetItemsPerPageOk() (*int32, bool)`

GetItemsPerPageOk returns a tuple with the ItemsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerPage

`func (o *ListOptions) SetItemsPerPage(v int32)`

SetItemsPerPage sets ItemsPerPage field to given value.


### GetListAsGroup

`func (o *ListOptions) GetListAsGroup() bool`

GetListAsGroup returns the ListAsGroup field if non-nil, zero value otherwise.

### GetListAsGroupOk

`func (o *ListOptions) GetListAsGroupOk() (*bool, bool)`

GetListAsGroupOk returns a tuple with the ListAsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListAsGroup

`func (o *ListOptions) SetListAsGroup(v bool)`

SetListAsGroup sets ListAsGroup field to given value.


### GetAutoRefresh

`func (o *ListOptions) GetAutoRefresh() bool`

GetAutoRefresh returns the AutoRefresh field if non-nil, zero value otherwise.

### GetAutoRefreshOk

`func (o *ListOptions) GetAutoRefreshOk() (*bool, bool)`

GetAutoRefreshOk returns a tuple with the AutoRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRefresh

`func (o *ListOptions) SetAutoRefresh(v bool)`

SetAutoRefresh sets AutoRefresh field to given value.


### GetStatsIsOpen

`func (o *ListOptions) GetStatsIsOpen() bool`

GetStatsIsOpen returns the StatsIsOpen field if non-nil, zero value otherwise.

### GetStatsIsOpenOk

`func (o *ListOptions) GetStatsIsOpenOk() (*bool, bool)`

GetStatsIsOpenOk returns a tuple with the StatsIsOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsIsOpen

`func (o *ListOptions) SetStatsIsOpen(v bool)`

SetStatsIsOpen sets StatsIsOpen field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


