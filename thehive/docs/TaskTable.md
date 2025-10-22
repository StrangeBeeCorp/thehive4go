# TaskTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Columns** | Pointer to [**[]TaskField**](TaskField.md) |  | [optional] 
**Filter** | Pointer to [**Filter**](Filter.md) |  | [optional] 
**Sort** | Pointer to [**[]AlertListSortInner**](AlertListSortInner.md) |  | [optional] 
**MaxElements** | Pointer to **int32** |  | [optional] 
**Kind** | [**Kind**](Kind.md) |  | 

## Methods

### NewTaskTable

`func NewTaskTable(kind Kind, ) *TaskTable`

NewTaskTable instantiates a new TaskTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskTableWithDefaults

`func NewTaskTableWithDefaults() *TaskTable`

NewTaskTableWithDefaults instantiates a new TaskTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *TaskTable) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TaskTable) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TaskTable) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TaskTable) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetColumns

`func (o *TaskTable) GetColumns() []TaskField`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *TaskTable) GetColumnsOk() (*[]TaskField, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *TaskTable) SetColumns(v []TaskField)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *TaskTable) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetFilter

`func (o *TaskTable) GetFilter() Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *TaskTable) GetFilterOk() (*Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *TaskTable) SetFilter(v Filter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *TaskTable) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSort

`func (o *TaskTable) GetSort() []AlertListSortInner`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *TaskTable) GetSortOk() (*[]AlertListSortInner, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *TaskTable) SetSort(v []AlertListSortInner)`

SetSort sets Sort field to given value.

### HasSort

`func (o *TaskTable) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetMaxElements

`func (o *TaskTable) GetMaxElements() int32`

GetMaxElements returns the MaxElements field if non-nil, zero value otherwise.

### GetMaxElementsOk

`func (o *TaskTable) GetMaxElementsOk() (*int32, bool)`

GetMaxElementsOk returns a tuple with the MaxElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxElements

`func (o *TaskTable) SetMaxElements(v int32)`

SetMaxElements sets MaxElements field to given value.

### HasMaxElements

`func (o *TaskTable) HasMaxElements() bool`

HasMaxElements returns a boolean if a field has been set.

### GetKind

`func (o *TaskTable) GetKind() Kind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TaskTable) GetKindOk() (*Kind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TaskTable) SetKind(v Kind)`

SetKind sets Kind field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


