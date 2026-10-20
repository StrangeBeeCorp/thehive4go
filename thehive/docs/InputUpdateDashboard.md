# InputUpdateDashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | New title for the dashboard. Leave unset to keep the current title. | [optional]
**Group** | Pointer to **string** | New group used to categorize the dashboard. Leave unset to keep the current group. | [optional]
**Description** | Pointer to **string** | New description for the dashboard. Leave unset to keep the current description. | [optional]
**Definition** | Pointer to **map[string]interface{}** | New widget layout and configuration for the dashboard, structured as rows of widgets. Leave unset to keep the current definition. This structure isn&#39;t formally documented: the simplest way to build a valid value is to edit the dashboard from TheHive web interface, then retrieve it with &#x60;GET /api/v1/dashboard/{dashboardId}&#x60; and use its &#x60;definition&#x60; value as a starting point. | [optional]
**Status** | Pointer to [**DashboardStatus**](DashboardStatus.md) | New visibility for the dashboard: &#x60;Private&#x60; restricts access to you, &#x60;Shared&#x60; makes it visible to your entire organization. Setting this to &#x60;Deleted&#x60; deletes the dashboard. Leave unset to keep the current visibility. | [optional]
**Version** | Pointer to **int32** | New version number for the dashboard definition format. Leave unset to keep the current version. | [optional]

## Methods

### NewInputUpdateDashboard

`func NewInputUpdateDashboard() *InputUpdateDashboard`

NewInputUpdateDashboard instantiates a new InputUpdateDashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputUpdateDashboardWithDefaults

`func NewInputUpdateDashboardWithDefaults() *InputUpdateDashboard`

NewInputUpdateDashboardWithDefaults instantiates a new InputUpdateDashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *InputUpdateDashboard) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InputUpdateDashboard) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InputUpdateDashboard) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InputUpdateDashboard) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetGroup

`func (o *InputUpdateDashboard) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InputUpdateDashboard) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InputUpdateDashboard) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *InputUpdateDashboard) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetDescription

`func (o *InputUpdateDashboard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputUpdateDashboard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputUpdateDashboard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InputUpdateDashboard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefinition

`func (o *InputUpdateDashboard) GetDefinition() map[string]interface{}`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *InputUpdateDashboard) GetDefinitionOk() (*map[string]interface{}, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *InputUpdateDashboard) SetDefinition(v map[string]interface{})`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *InputUpdateDashboard) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetStatus

`func (o *InputUpdateDashboard) GetStatus() DashboardStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InputUpdateDashboard) GetStatusOk() (*DashboardStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InputUpdateDashboard) SetStatus(v DashboardStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InputUpdateDashboard) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *InputUpdateDashboard) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InputUpdateDashboard) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InputUpdateDashboard) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InputUpdateDashboard) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
