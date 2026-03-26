# InputUpdateObservableWithIds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** |  | 
**DataType** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Tlp** | Pointer to **int32** |  | [optional] 
**Pap** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Ioc** | Pointer to **bool** |  | [optional] 
**Sighted** | Pointer to **bool** |  | [optional] 
**SightedAt** | Pointer to **int64** |  | [optional] 
**IgnoreSimilarity** | Pointer to **bool** |  | [optional] 
**AddTags** | Pointer to **[]string** | Those tags will be added to the current observable | [optional] 
**RemoveTags** | Pointer to **[]string** | Those tags will be removed from the current observable | [optional] 
**External** | Pointer to **bool** |  | [optional] 

## Methods

### NewInputUpdateObservableWithIds

`func NewInputUpdateObservableWithIds(ids []string, ) *InputUpdateObservableWithIds`

NewInputUpdateObservableWithIds instantiates a new InputUpdateObservableWithIds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputUpdateObservableWithIdsWithDefaults

`func NewInputUpdateObservableWithIdsWithDefaults() *InputUpdateObservableWithIds`

NewInputUpdateObservableWithIdsWithDefaults instantiates a new InputUpdateObservableWithIds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *InputUpdateObservableWithIds) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *InputUpdateObservableWithIds) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *InputUpdateObservableWithIds) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetDataType

`func (o *InputUpdateObservableWithIds) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *InputUpdateObservableWithIds) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *InputUpdateObservableWithIds) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *InputUpdateObservableWithIds) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetMessage

`func (o *InputUpdateObservableWithIds) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InputUpdateObservableWithIds) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InputUpdateObservableWithIds) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InputUpdateObservableWithIds) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTlp

`func (o *InputUpdateObservableWithIds) GetTlp() int32`

GetTlp returns the Tlp field if non-nil, zero value otherwise.

### GetTlpOk

`func (o *InputUpdateObservableWithIds) GetTlpOk() (*int32, bool)`

GetTlpOk returns a tuple with the Tlp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlp

`func (o *InputUpdateObservableWithIds) SetTlp(v int32)`

SetTlp sets Tlp field to given value.

### HasTlp

`func (o *InputUpdateObservableWithIds) HasTlp() bool`

HasTlp returns a boolean if a field has been set.

### GetPap

`func (o *InputUpdateObservableWithIds) GetPap() int32`

GetPap returns the Pap field if non-nil, zero value otherwise.

### GetPapOk

`func (o *InputUpdateObservableWithIds) GetPapOk() (*int32, bool)`

GetPapOk returns a tuple with the Pap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPap

`func (o *InputUpdateObservableWithIds) SetPap(v int32)`

SetPap sets Pap field to given value.

### HasPap

`func (o *InputUpdateObservableWithIds) HasPap() bool`

HasPap returns a boolean if a field has been set.

### GetTags

`func (o *InputUpdateObservableWithIds) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InputUpdateObservableWithIds) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InputUpdateObservableWithIds) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InputUpdateObservableWithIds) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIoc

`func (o *InputUpdateObservableWithIds) GetIoc() bool`

GetIoc returns the Ioc field if non-nil, zero value otherwise.

### GetIocOk

`func (o *InputUpdateObservableWithIds) GetIocOk() (*bool, bool)`

GetIocOk returns a tuple with the Ioc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoc

`func (o *InputUpdateObservableWithIds) SetIoc(v bool)`

SetIoc sets Ioc field to given value.

### HasIoc

`func (o *InputUpdateObservableWithIds) HasIoc() bool`

HasIoc returns a boolean if a field has been set.

### GetSighted

`func (o *InputUpdateObservableWithIds) GetSighted() bool`

GetSighted returns the Sighted field if non-nil, zero value otherwise.

### GetSightedOk

`func (o *InputUpdateObservableWithIds) GetSightedOk() (*bool, bool)`

GetSightedOk returns a tuple with the Sighted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSighted

`func (o *InputUpdateObservableWithIds) SetSighted(v bool)`

SetSighted sets Sighted field to given value.

### HasSighted

`func (o *InputUpdateObservableWithIds) HasSighted() bool`

HasSighted returns a boolean if a field has been set.

### GetSightedAt

`func (o *InputUpdateObservableWithIds) GetSightedAt() int64`

GetSightedAt returns the SightedAt field if non-nil, zero value otherwise.

### GetSightedAtOk

`func (o *InputUpdateObservableWithIds) GetSightedAtOk() (*int64, bool)`

GetSightedAtOk returns a tuple with the SightedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSightedAt

`func (o *InputUpdateObservableWithIds) SetSightedAt(v int64)`

SetSightedAt sets SightedAt field to given value.

### HasSightedAt

`func (o *InputUpdateObservableWithIds) HasSightedAt() bool`

HasSightedAt returns a boolean if a field has been set.

### GetIgnoreSimilarity

`func (o *InputUpdateObservableWithIds) GetIgnoreSimilarity() bool`

GetIgnoreSimilarity returns the IgnoreSimilarity field if non-nil, zero value otherwise.

### GetIgnoreSimilarityOk

`func (o *InputUpdateObservableWithIds) GetIgnoreSimilarityOk() (*bool, bool)`

GetIgnoreSimilarityOk returns a tuple with the IgnoreSimilarity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSimilarity

`func (o *InputUpdateObservableWithIds) SetIgnoreSimilarity(v bool)`

SetIgnoreSimilarity sets IgnoreSimilarity field to given value.

### HasIgnoreSimilarity

`func (o *InputUpdateObservableWithIds) HasIgnoreSimilarity() bool`

HasIgnoreSimilarity returns a boolean if a field has been set.

### GetAddTags

`func (o *InputUpdateObservableWithIds) GetAddTags() []string`

GetAddTags returns the AddTags field if non-nil, zero value otherwise.

### GetAddTagsOk

`func (o *InputUpdateObservableWithIds) GetAddTagsOk() (*[]string, bool)`

GetAddTagsOk returns a tuple with the AddTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTags

`func (o *InputUpdateObservableWithIds) SetAddTags(v []string)`

SetAddTags sets AddTags field to given value.

### HasAddTags

`func (o *InputUpdateObservableWithIds) HasAddTags() bool`

HasAddTags returns a boolean if a field has been set.

### GetRemoveTags

`func (o *InputUpdateObservableWithIds) GetRemoveTags() []string`

GetRemoveTags returns the RemoveTags field if non-nil, zero value otherwise.

### GetRemoveTagsOk

`func (o *InputUpdateObservableWithIds) GetRemoveTagsOk() (*[]string, bool)`

GetRemoveTagsOk returns a tuple with the RemoveTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveTags

`func (o *InputUpdateObservableWithIds) SetRemoveTags(v []string)`

SetRemoveTags sets RemoveTags field to given value.

### HasRemoveTags

`func (o *InputUpdateObservableWithIds) HasRemoveTags() bool`

HasRemoveTags returns a boolean if a field has been set.

### GetExternal

`func (o *InputUpdateObservableWithIds) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *InputUpdateObservableWithIds) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *InputUpdateObservableWithIds) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *InputUpdateObservableWithIds) HasExternal() bool`

HasExternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


