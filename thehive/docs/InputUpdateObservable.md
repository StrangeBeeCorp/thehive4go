# InputUpdateObservable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewInputUpdateObservable

`func NewInputUpdateObservable() *InputUpdateObservable`

NewInputUpdateObservable instantiates a new InputUpdateObservable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputUpdateObservableWithDefaults

`func NewInputUpdateObservableWithDefaults() *InputUpdateObservable`

NewInputUpdateObservableWithDefaults instantiates a new InputUpdateObservable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *InputUpdateObservable) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *InputUpdateObservable) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *InputUpdateObservable) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *InputUpdateObservable) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetMessage

`func (o *InputUpdateObservable) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InputUpdateObservable) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InputUpdateObservable) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InputUpdateObservable) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTlp

`func (o *InputUpdateObservable) GetTlp() int32`

GetTlp returns the Tlp field if non-nil, zero value otherwise.

### GetTlpOk

`func (o *InputUpdateObservable) GetTlpOk() (*int32, bool)`

GetTlpOk returns a tuple with the Tlp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlp

`func (o *InputUpdateObservable) SetTlp(v int32)`

SetTlp sets Tlp field to given value.

### HasTlp

`func (o *InputUpdateObservable) HasTlp() bool`

HasTlp returns a boolean if a field has been set.

### GetPap

`func (o *InputUpdateObservable) GetPap() int32`

GetPap returns the Pap field if non-nil, zero value otherwise.

### GetPapOk

`func (o *InputUpdateObservable) GetPapOk() (*int32, bool)`

GetPapOk returns a tuple with the Pap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPap

`func (o *InputUpdateObservable) SetPap(v int32)`

SetPap sets Pap field to given value.

### HasPap

`func (o *InputUpdateObservable) HasPap() bool`

HasPap returns a boolean if a field has been set.

### GetTags

`func (o *InputUpdateObservable) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InputUpdateObservable) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InputUpdateObservable) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InputUpdateObservable) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIoc

`func (o *InputUpdateObservable) GetIoc() bool`

GetIoc returns the Ioc field if non-nil, zero value otherwise.

### GetIocOk

`func (o *InputUpdateObservable) GetIocOk() (*bool, bool)`

GetIocOk returns a tuple with the Ioc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoc

`func (o *InputUpdateObservable) SetIoc(v bool)`

SetIoc sets Ioc field to given value.

### HasIoc

`func (o *InputUpdateObservable) HasIoc() bool`

HasIoc returns a boolean if a field has been set.

### GetSighted

`func (o *InputUpdateObservable) GetSighted() bool`

GetSighted returns the Sighted field if non-nil, zero value otherwise.

### GetSightedOk

`func (o *InputUpdateObservable) GetSightedOk() (*bool, bool)`

GetSightedOk returns a tuple with the Sighted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSighted

`func (o *InputUpdateObservable) SetSighted(v bool)`

SetSighted sets Sighted field to given value.

### HasSighted

`func (o *InputUpdateObservable) HasSighted() bool`

HasSighted returns a boolean if a field has been set.

### GetSightedAt

`func (o *InputUpdateObservable) GetSightedAt() int64`

GetSightedAt returns the SightedAt field if non-nil, zero value otherwise.

### GetSightedAtOk

`func (o *InputUpdateObservable) GetSightedAtOk() (*int64, bool)`

GetSightedAtOk returns a tuple with the SightedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSightedAt

`func (o *InputUpdateObservable) SetSightedAt(v int64)`

SetSightedAt sets SightedAt field to given value.

### HasSightedAt

`func (o *InputUpdateObservable) HasSightedAt() bool`

HasSightedAt returns a boolean if a field has been set.

### GetIgnoreSimilarity

`func (o *InputUpdateObservable) GetIgnoreSimilarity() bool`

GetIgnoreSimilarity returns the IgnoreSimilarity field if non-nil, zero value otherwise.

### GetIgnoreSimilarityOk

`func (o *InputUpdateObservable) GetIgnoreSimilarityOk() (*bool, bool)`

GetIgnoreSimilarityOk returns a tuple with the IgnoreSimilarity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSimilarity

`func (o *InputUpdateObservable) SetIgnoreSimilarity(v bool)`

SetIgnoreSimilarity sets IgnoreSimilarity field to given value.

### HasIgnoreSimilarity

`func (o *InputUpdateObservable) HasIgnoreSimilarity() bool`

HasIgnoreSimilarity returns a boolean if a field has been set.

### GetAddTags

`func (o *InputUpdateObservable) GetAddTags() []string`

GetAddTags returns the AddTags field if non-nil, zero value otherwise.

### GetAddTagsOk

`func (o *InputUpdateObservable) GetAddTagsOk() (*[]string, bool)`

GetAddTagsOk returns a tuple with the AddTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTags

`func (o *InputUpdateObservable) SetAddTags(v []string)`

SetAddTags sets AddTags field to given value.

### HasAddTags

`func (o *InputUpdateObservable) HasAddTags() bool`

HasAddTags returns a boolean if a field has been set.

### GetRemoveTags

`func (o *InputUpdateObservable) GetRemoveTags() []string`

GetRemoveTags returns the RemoveTags field if non-nil, zero value otherwise.

### GetRemoveTagsOk

`func (o *InputUpdateObservable) GetRemoveTagsOk() (*[]string, bool)`

GetRemoveTagsOk returns a tuple with the RemoveTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveTags

`func (o *InputUpdateObservable) SetRemoveTags(v []string)`

SetRemoveTags sets RemoveTags field to given value.

### HasRemoveTags

`func (o *InputUpdateObservable) HasRemoveTags() bool`

HasRemoveTags returns a boolean if a field has been set.

### GetExternal

`func (o *InputUpdateObservable) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *InputUpdateObservable) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *InputUpdateObservable) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *InputUpdateObservable) HasExternal() bool`

HasExternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


