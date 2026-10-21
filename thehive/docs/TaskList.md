# TaskList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Title displayed above the widget. | [optional]
**Fields** | Pointer to [**[]TaskField**](TaskField.md) | Task fields to display for each item in the list. | [optional]
**LogColumns** | Pointer to [**[]LogField**](LogField.md) | Task log fields to display under each task when &#x60;withTaskLogs&#x60; is &#x60;true&#x60;. | [optional]
**WithTaskLogs** | Pointer to **bool** | When &#x60;true&#x60;, includes task logs under each task entry. | [optional] [default to true]
**Filter** | Pointer to [**Filter**](Filter.md) | Filter to apply to the task list. See the [Query and Export](https://docs.strangebee.com/thehive/api-docs/#tag/Query-and-Export) section for the filter syntax. | [optional]
**Sort** | Pointer to [**[]AlertListSortInner**](AlertListSortInner.md) | Sort order for the task list. See the [Query and Export](https://docs.strangebee.com/thehive/api-docs/#tag/Query-and-Export) section for the sort syntax. | [optional]
**MaxElements** | Pointer to **int32** | Maximum number of tasks to include in the widget. If omitted, the global &#x60;maxElements&#x60; limit applies. | [optional]
**Kind** | [**Kind**](Kind.md) |  |

## Methods

### NewTaskList

`func NewTaskList(kind Kind, ) *TaskList`

NewTaskList instantiates a new TaskList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskListWithDefaults

`func NewTaskListWithDefaults() *TaskList`

NewTaskListWithDefaults instantiates a new TaskList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *TaskList) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TaskList) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TaskList) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TaskList) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFields

`func (o *TaskList) GetFields() []TaskField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *TaskList) GetFieldsOk() (*[]TaskField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *TaskList) SetFields(v []TaskField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *TaskList) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetLogColumns

`func (o *TaskList) GetLogColumns() []LogField`

GetLogColumns returns the LogColumns field if non-nil, zero value otherwise.

### GetLogColumnsOk

`func (o *TaskList) GetLogColumnsOk() (*[]LogField, bool)`

GetLogColumnsOk returns a tuple with the LogColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogColumns

`func (o *TaskList) SetLogColumns(v []LogField)`

SetLogColumns sets LogColumns field to given value.

### HasLogColumns

`func (o *TaskList) HasLogColumns() bool`

HasLogColumns returns a boolean if a field has been set.

### GetWithTaskLogs

`func (o *TaskList) GetWithTaskLogs() bool`

GetWithTaskLogs returns the WithTaskLogs field if non-nil, zero value otherwise.

### GetWithTaskLogsOk

`func (o *TaskList) GetWithTaskLogsOk() (*bool, bool)`

GetWithTaskLogsOk returns a tuple with the WithTaskLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithTaskLogs

`func (o *TaskList) SetWithTaskLogs(v bool)`

SetWithTaskLogs sets WithTaskLogs field to given value.

### HasWithTaskLogs

`func (o *TaskList) HasWithTaskLogs() bool`

HasWithTaskLogs returns a boolean if a field has been set.

### GetFilter

`func (o *TaskList) GetFilter() Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *TaskList) GetFilterOk() (*Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *TaskList) SetFilter(v Filter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *TaskList) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSort

`func (o *TaskList) GetSort() []AlertListSortInner`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *TaskList) GetSortOk() (*[]AlertListSortInner, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *TaskList) SetSort(v []AlertListSortInner)`

SetSort sets Sort field to given value.

### HasSort

`func (o *TaskList) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetMaxElements

`func (o *TaskList) GetMaxElements() int32`

GetMaxElements returns the MaxElements field if non-nil, zero value otherwise.

### GetMaxElementsOk

`func (o *TaskList) GetMaxElementsOk() (*int32, bool)`

GetMaxElementsOk returns a tuple with the MaxElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxElements

`func (o *TaskList) SetMaxElements(v int32)`

SetMaxElements sets MaxElements field to given value.

### HasMaxElements

`func (o *TaskList) HasMaxElements() bool`

HasMaxElements returns a boolean if a field has been set.

### GetKind

`func (o *TaskList) GetKind() Kind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TaskList) GetKindOk() (*Kind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TaskList) SetKind(v Kind)`

SetKind sets Kind field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
