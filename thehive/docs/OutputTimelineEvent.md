# OutputTimelineEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **int64** |  |
**Kind** | **string** |  |
**Entity** | **string** |  |
**EntityId** | **string** |  |
**Details** | **map[string]interface{}** |  |
**EndDate** | Pointer to **int64** |  | [optional]

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
