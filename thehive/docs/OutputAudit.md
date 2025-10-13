# OutputAudit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnderscoreId** | **string** |  | 
**UnderscoreType** | **string** |  | 
**UnderscoreCreatedBy** | **string** |  | 
**UnderscoreUpdatedBy** | Pointer to **string** |  | [optional] 
**UnderscoreCreatedAt** | **int32** |  | 
**UnderscoreUpdatedAt** | Pointer to **int32** |  | [optional] 
**Action** | **string** |  | 
**RequestId** | **string** |  | 
**RootId** | **string** |  | 
**Details** | **map[string]interface{}** |  | 
**ObjectId** | Pointer to **string** |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **map[string]interface{}** |  | [optional] 
**Context** | Pointer to **map[string]interface{}** |  | [optional] 
**MainAction** | **bool** |  | 

## Methods

### NewOutputAudit

`func NewOutputAudit(underscoreId string, underscoreType string, underscoreCreatedBy string, underscoreCreatedAt int32, action string, requestId string, rootId string, details map[string]interface{}, mainAction bool, ) *OutputAudit`

NewOutputAudit instantiates a new OutputAudit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputAuditWithDefaults

`func NewOutputAuditWithDefaults() *OutputAudit`

NewOutputAuditWithDefaults instantiates a new OutputAudit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnderscoreId

`func (o *OutputAudit) GetUnderscoreId() string`

GetUnderscoreId returns the UnderscoreId field if non-nil, zero value otherwise.

### GetUnderscoreIdOk

`func (o *OutputAudit) GetUnderscoreIdOk() (*string, bool)`

GetUnderscoreIdOk returns a tuple with the UnderscoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreId

`func (o *OutputAudit) SetUnderscoreId(v string)`

SetUnderscoreId sets UnderscoreId field to given value.


### GetUnderscoreType

`func (o *OutputAudit) GetUnderscoreType() string`

GetUnderscoreType returns the UnderscoreType field if non-nil, zero value otherwise.

### GetUnderscoreTypeOk

`func (o *OutputAudit) GetUnderscoreTypeOk() (*string, bool)`

GetUnderscoreTypeOk returns a tuple with the UnderscoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreType

`func (o *OutputAudit) SetUnderscoreType(v string)`

SetUnderscoreType sets UnderscoreType field to given value.


### GetUnderscoreCreatedBy

`func (o *OutputAudit) GetUnderscoreCreatedBy() string`

GetUnderscoreCreatedBy returns the UnderscoreCreatedBy field if non-nil, zero value otherwise.

### GetUnderscoreCreatedByOk

`func (o *OutputAudit) GetUnderscoreCreatedByOk() (*string, bool)`

GetUnderscoreCreatedByOk returns a tuple with the UnderscoreCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedBy

`func (o *OutputAudit) SetUnderscoreCreatedBy(v string)`

SetUnderscoreCreatedBy sets UnderscoreCreatedBy field to given value.


### GetUnderscoreUpdatedBy

`func (o *OutputAudit) GetUnderscoreUpdatedBy() string`

GetUnderscoreUpdatedBy returns the UnderscoreUpdatedBy field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedByOk

`func (o *OutputAudit) GetUnderscoreUpdatedByOk() (*string, bool)`

GetUnderscoreUpdatedByOk returns a tuple with the UnderscoreUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedBy

`func (o *OutputAudit) SetUnderscoreUpdatedBy(v string)`

SetUnderscoreUpdatedBy sets UnderscoreUpdatedBy field to given value.

### HasUnderscoreUpdatedBy

`func (o *OutputAudit) HasUnderscoreUpdatedBy() bool`

HasUnderscoreUpdatedBy returns a boolean if a field has been set.

### GetUnderscoreCreatedAt

`func (o *OutputAudit) GetUnderscoreCreatedAt() int32`

GetUnderscoreCreatedAt returns the UnderscoreCreatedAt field if non-nil, zero value otherwise.

### GetUnderscoreCreatedAtOk

`func (o *OutputAudit) GetUnderscoreCreatedAtOk() (*int32, bool)`

GetUnderscoreCreatedAtOk returns a tuple with the UnderscoreCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedAt

`func (o *OutputAudit) SetUnderscoreCreatedAt(v int32)`

SetUnderscoreCreatedAt sets UnderscoreCreatedAt field to given value.


### GetUnderscoreUpdatedAt

`func (o *OutputAudit) GetUnderscoreUpdatedAt() int32`

GetUnderscoreUpdatedAt returns the UnderscoreUpdatedAt field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedAtOk

`func (o *OutputAudit) GetUnderscoreUpdatedAtOk() (*int32, bool)`

GetUnderscoreUpdatedAtOk returns a tuple with the UnderscoreUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedAt

`func (o *OutputAudit) SetUnderscoreUpdatedAt(v int32)`

SetUnderscoreUpdatedAt sets UnderscoreUpdatedAt field to given value.

### HasUnderscoreUpdatedAt

`func (o *OutputAudit) HasUnderscoreUpdatedAt() bool`

HasUnderscoreUpdatedAt returns a boolean if a field has been set.

### GetAction

`func (o *OutputAudit) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *OutputAudit) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *OutputAudit) SetAction(v string)`

SetAction sets Action field to given value.


### GetRequestId

`func (o *OutputAudit) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *OutputAudit) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *OutputAudit) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetRootId

`func (o *OutputAudit) GetRootId() string`

GetRootId returns the RootId field if non-nil, zero value otherwise.

### GetRootIdOk

`func (o *OutputAudit) GetRootIdOk() (*string, bool)`

GetRootIdOk returns a tuple with the RootId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootId

`func (o *OutputAudit) SetRootId(v string)`

SetRootId sets RootId field to given value.


### GetDetails

`func (o *OutputAudit) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *OutputAudit) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *OutputAudit) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.


### GetObjectId

`func (o *OutputAudit) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *OutputAudit) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *OutputAudit) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *OutputAudit) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetObjectType

`func (o *OutputAudit) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OutputAudit) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OutputAudit) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *OutputAudit) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObject

`func (o *OutputAudit) GetObject() map[string]interface{}`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *OutputAudit) GetObjectOk() (*map[string]interface{}, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *OutputAudit) SetObject(v map[string]interface{})`

SetObject sets Object field to given value.

### HasObject

`func (o *OutputAudit) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetContext

`func (o *OutputAudit) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *OutputAudit) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *OutputAudit) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *OutputAudit) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetMainAction

`func (o *OutputAudit) GetMainAction() bool`

GetMainAction returns the MainAction field if non-nil, zero value otherwise.

### GetMainActionOk

`func (o *OutputAudit) GetMainActionOk() (*bool, bool)`

GetMainActionOk returns a tuple with the MainAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainAction

`func (o *OutputAudit) SetMainAction(v bool)`

SetMainAction sets MainAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


