# InputCreateDashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Title of the dashboard. |
**Group** | Pointer to **string** | Group used to categorize the dashboard on the dashboards list. | [optional] [default to "default"]
**Description** | **string** | Description of the dashboard. |
**Status** | [**DashboardStatus**](DashboardStatus.md) | Visibility of the dashboard: &#x60;Private&#x60; restricts access to you, &#x60;Shared&#x60; makes it visible to your entire organization. &#x60;Deleted&#x60; has no effect on creation and behaves like &#x60;Private&#x60;. |
**Definition** | **map[string]interface{}** | Widget layout and configuration for the dashboard, structured as rows of widgets. Each widget defines a chart type, the entity it queries, and its display options. This structure isn&#39;t formally documented: the simplest way to build a valid value is to create a dashboard from TheHive web interface, then retrieve it with &#x60;GET /api/v1/dashboard/{dashboardId}&#x60; and use its &#x60;definition&#x60; value as a starting point. |
**Version** | Pointer to **int32** | Version of the dashboard definition format. The current format is &#x60;1&#x60;. Older values indicate a legacy definition that must be converted before use. | [optional] [default to 1]

## Methods

### NewInputCreateDashboard

`func NewInputCreateDashboard(title string, description string, status DashboardStatus, definition map[string]interface{}, ) *InputCreateDashboard`

NewInputCreateDashboard instantiates a new InputCreateDashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputCreateDashboardWithDefaults

`func NewInputCreateDashboardWithDefaults() *InputCreateDashboard`

NewInputCreateDashboardWithDefaults instantiates a new InputCreateDashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *InputCreateDashboard) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InputCreateDashboard) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InputCreateDashboard) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetGroup

`func (o *InputCreateDashboard) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InputCreateDashboard) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InputCreateDashboard) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *InputCreateDashboard) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetDescription

`func (o *InputCreateDashboard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputCreateDashboard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputCreateDashboard) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetStatus

`func (o *InputCreateDashboard) GetStatus() DashboardStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InputCreateDashboard) GetStatusOk() (*DashboardStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InputCreateDashboard) SetStatus(v DashboardStatus)`

SetStatus sets Status field to given value.


### GetDefinition

`func (o *InputCreateDashboard) GetDefinition() map[string]interface{}`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *InputCreateDashboard) GetDefinitionOk() (*map[string]interface{}, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *InputCreateDashboard) SetDefinition(v map[string]interface{})`

SetDefinition sets Definition field to given value.


### GetVersion

`func (o *InputCreateDashboard) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InputCreateDashboard) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InputCreateDashboard) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InputCreateDashboard) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
