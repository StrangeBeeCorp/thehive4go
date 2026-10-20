# OutputTimelineEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **int64** | Start date and time of the event. Returned as a Unix timestamp in milliseconds. |
**Kind** | **string** | The event type. Possible values: &#x60;case.start&#x60;, &#x60;case.created&#x60;, &#x60;case.new&#x60;, &#x60;case.inProgress&#x60;, &#x60;case.closed&#x60;, &#x60;case.end&#x60;, &#x60;alert.occurred&#x60;, &#x60;procedure.occurred&#x60;, &#x60;observable.sighted&#x60;, &#x60;task&#x60;, &#x60;log.created&#x60;, &#x60;custom&#x60;. |
**Entity** | **string** | The entity type involved in the event. Possible values: &#x60;Case&#x60;, &#x60;Alert&#x60;, &#x60;Procedure&#x60;, &#x60;Observable&#x60;, &#x60;Task&#x60;, &#x60;Log&#x60;, &#x60;CustomEvent&#x60;. |
**EntityId** | **string** | Unique identifier of the entity involved in the event. |
**Details** | **map[string]interface{}** | Event-specific details. Content varies by &#x60;kind&#x60;. For entity events, holds the full JSON representation of the related entity under a key matching the entity type (for example, &#x60;{\&quot;task\&quot;: {...}}&#x60;). Empty for case lifecycle events. |
**EndDate** | Pointer to **int64** | End date and time of range events such as tasks. &#x60;null&#x60; for point-in-time events. Returned as a Unix timestamp in milliseconds. | [optional]

## Methods

### NewOutputTimelineEvent

`func NewOutputTimelineEvent(date int64, kind string, entity string, entityId string, details map[string]interface{}, ) *OutputTimelineEvent`

NewOutputTimelineEvent instantiates a new OutputTimelineEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputTimelineEventWithDefaults

`func NewOutputTimelineEventWithDefaults() *OutputTimelineEvent`

NewOutputTimelineEventWithDefaults instantiates a new OutputTimelineEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *OutputTimelineEvent) GetDate() int64`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *OutputTimelineEvent) GetDateOk() (*int64, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *OutputTimelineEvent) SetDate(v int64)`

SetDate sets Date field to given value.


### GetKind

`func (o *OutputTimelineEvent) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *OutputTimelineEvent) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *OutputTimelineEvent) SetKind(v string)`

SetKind sets Kind field to given value.


### GetEntity

`func (o *OutputTimelineEvent) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *OutputTimelineEvent) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *OutputTimelineEvent) SetEntity(v string)`

SetEntity sets Entity field to given value.


### GetEntityId

`func (o *OutputTimelineEvent) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *OutputTimelineEvent) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *OutputTimelineEvent) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetDetails

`func (o *OutputTimelineEvent) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *OutputTimelineEvent) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *OutputTimelineEvent) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.


### GetEndDate

`func (o *OutputTimelineEvent) GetEndDate() int64`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *OutputTimelineEvent) GetEndDateOk() (*int64, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *OutputTimelineEvent) SetEndDate(v int64)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *OutputTimelineEvent) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
