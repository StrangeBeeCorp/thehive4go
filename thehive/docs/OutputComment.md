# OutputComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnderscoreId** | **string** |  | 
**UnderscoreType** | **string** |  | 
**CreatedBy** | **string** |  | 
**CreatedAt** | **int64** |  | 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**Message** | **string** |  | 
**IsEdited** | **bool** |  | 
**ExtraData** | **map[string]interface{}** |  | 
**External** | **bool** |  | 

## Methods

### NewOutputComment

`func NewOutputComment(underscoreId string, underscoreType string, createdBy string, createdAt int64, message string, isEdited bool, extraData map[string]interface{}, external bool, ) *OutputComment`

NewOutputComment instantiates a new OutputComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputCommentWithDefaults

`func NewOutputCommentWithDefaults() *OutputComment`

NewOutputCommentWithDefaults instantiates a new OutputComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnderscoreId

`func (o *OutputComment) GetUnderscoreId() string`

GetUnderscoreId returns the UnderscoreId field if non-nil, zero value otherwise.

### GetUnderscoreIdOk

`func (o *OutputComment) GetUnderscoreIdOk() (*string, bool)`

GetUnderscoreIdOk returns a tuple with the UnderscoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreId

`func (o *OutputComment) SetUnderscoreId(v string)`

SetUnderscoreId sets UnderscoreId field to given value.


### GetUnderscoreType

`func (o *OutputComment) GetUnderscoreType() string`

GetUnderscoreType returns the UnderscoreType field if non-nil, zero value otherwise.

### GetUnderscoreTypeOk

`func (o *OutputComment) GetUnderscoreTypeOk() (*string, bool)`

GetUnderscoreTypeOk returns a tuple with the UnderscoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreType

`func (o *OutputComment) SetUnderscoreType(v string)`

SetUnderscoreType sets UnderscoreType field to given value.


### GetCreatedBy

`func (o *OutputComment) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *OutputComment) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *OutputComment) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAt

`func (o *OutputComment) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OutputComment) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OutputComment) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *OutputComment) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OutputComment) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OutputComment) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OutputComment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *OutputComment) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *OutputComment) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *OutputComment) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *OutputComment) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetMessage

`func (o *OutputComment) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OutputComment) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OutputComment) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetIsEdited

`func (o *OutputComment) GetIsEdited() bool`

GetIsEdited returns the IsEdited field if non-nil, zero value otherwise.

### GetIsEditedOk

`func (o *OutputComment) GetIsEditedOk() (*bool, bool)`

GetIsEditedOk returns a tuple with the IsEdited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEdited

`func (o *OutputComment) SetIsEdited(v bool)`

SetIsEdited sets IsEdited field to given value.


### GetExtraData

`func (o *OutputComment) GetExtraData() map[string]interface{}`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *OutputComment) GetExtraDataOk() (*map[string]interface{}, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *OutputComment) SetExtraData(v map[string]interface{})`

SetExtraData sets ExtraData field to given value.


### GetExternal

`func (o *OutputComment) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *OutputComment) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *OutputComment) SetExternal(v bool)`

SetExternal sets External field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


