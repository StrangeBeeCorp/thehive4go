# OutputObservable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnderscoreId** | **string** |  | 
**UnderscoreType** | **string** |  | 
**UnderscoreCreatedBy** | **string** |  | 
**UnderscoreUpdatedBy** | Pointer to **string** |  | [optional] 
**UnderscoreCreatedAt** | **int64** |  | 
**UnderscoreUpdatedAt** | Pointer to **int64** |  | [optional] 
**DataType** | **string** |  | 
**Data** | Pointer to **string** |  | [optional] 
**StartDate** | **int64** |  | 
**Attachment** | Pointer to [**OutputAttachment**](OutputAttachment.md) |  | [optional] 
**Tlp** | **int32** |  | 
**TlpLabel** | **string** |  | 
**Pap** | **int32** |  | 
**PapLabel** | **string** |  | 
**Tags** | Pointer to **[]string** |  | [optional] 
**Ioc** | **bool** |  | 
**Sighted** | **bool** |  | 
**SightedAt** | Pointer to **int64** |  | [optional] 
**Reports** | **map[string]interface{}** |  | 
**Message** | Pointer to **string** |  | [optional] 
**ExtraData** | **map[string]interface{}** |  | 
**IgnoreSimilarity** | **bool** |  | 
**External** | **bool** |  | 

## Methods

### NewOutputObservable

`func NewOutputObservable(underscoreId string, underscoreType string, underscoreCreatedBy string, underscoreCreatedAt int64, dataType string, startDate int64, tlp int32, tlpLabel string, pap int32, papLabel string, ioc bool, sighted bool, reports map[string]interface{}, extraData map[string]interface{}, ignoreSimilarity bool, external bool, ) *OutputObservable`

NewOutputObservable instantiates a new OutputObservable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputObservableWithDefaults

`func NewOutputObservableWithDefaults() *OutputObservable`

NewOutputObservableWithDefaults instantiates a new OutputObservable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnderscoreId

`func (o *OutputObservable) GetUnderscoreId() string`

GetUnderscoreId returns the UnderscoreId field if non-nil, zero value otherwise.

### GetUnderscoreIdOk

`func (o *OutputObservable) GetUnderscoreIdOk() (*string, bool)`

GetUnderscoreIdOk returns a tuple with the UnderscoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreId

`func (o *OutputObservable) SetUnderscoreId(v string)`

SetUnderscoreId sets UnderscoreId field to given value.


### GetUnderscoreType

`func (o *OutputObservable) GetUnderscoreType() string`

GetUnderscoreType returns the UnderscoreType field if non-nil, zero value otherwise.

### GetUnderscoreTypeOk

`func (o *OutputObservable) GetUnderscoreTypeOk() (*string, bool)`

GetUnderscoreTypeOk returns a tuple with the UnderscoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreType

`func (o *OutputObservable) SetUnderscoreType(v string)`

SetUnderscoreType sets UnderscoreType field to given value.


### GetUnderscoreCreatedBy

`func (o *OutputObservable) GetUnderscoreCreatedBy() string`

GetUnderscoreCreatedBy returns the UnderscoreCreatedBy field if non-nil, zero value otherwise.

### GetUnderscoreCreatedByOk

`func (o *OutputObservable) GetUnderscoreCreatedByOk() (*string, bool)`

GetUnderscoreCreatedByOk returns a tuple with the UnderscoreCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedBy

`func (o *OutputObservable) SetUnderscoreCreatedBy(v string)`

SetUnderscoreCreatedBy sets UnderscoreCreatedBy field to given value.


### GetUnderscoreUpdatedBy

`func (o *OutputObservable) GetUnderscoreUpdatedBy() string`

GetUnderscoreUpdatedBy returns the UnderscoreUpdatedBy field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedByOk

`func (o *OutputObservable) GetUnderscoreUpdatedByOk() (*string, bool)`

GetUnderscoreUpdatedByOk returns a tuple with the UnderscoreUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedBy

`func (o *OutputObservable) SetUnderscoreUpdatedBy(v string)`

SetUnderscoreUpdatedBy sets UnderscoreUpdatedBy field to given value.

### HasUnderscoreUpdatedBy

`func (o *OutputObservable) HasUnderscoreUpdatedBy() bool`

HasUnderscoreUpdatedBy returns a boolean if a field has been set.

### GetUnderscoreCreatedAt

`func (o *OutputObservable) GetUnderscoreCreatedAt() int64`

GetUnderscoreCreatedAt returns the UnderscoreCreatedAt field if non-nil, zero value otherwise.

### GetUnderscoreCreatedAtOk

`func (o *OutputObservable) GetUnderscoreCreatedAtOk() (*int64, bool)`

GetUnderscoreCreatedAtOk returns a tuple with the UnderscoreCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedAt

`func (o *OutputObservable) SetUnderscoreCreatedAt(v int64)`

SetUnderscoreCreatedAt sets UnderscoreCreatedAt field to given value.


### GetUnderscoreUpdatedAt

`func (o *OutputObservable) GetUnderscoreUpdatedAt() int64`

GetUnderscoreUpdatedAt returns the UnderscoreUpdatedAt field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedAtOk

`func (o *OutputObservable) GetUnderscoreUpdatedAtOk() (*int64, bool)`

GetUnderscoreUpdatedAtOk returns a tuple with the UnderscoreUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedAt

`func (o *OutputObservable) SetUnderscoreUpdatedAt(v int64)`

SetUnderscoreUpdatedAt sets UnderscoreUpdatedAt field to given value.

### HasUnderscoreUpdatedAt

`func (o *OutputObservable) HasUnderscoreUpdatedAt() bool`

HasUnderscoreUpdatedAt returns a boolean if a field has been set.

### GetDataType

`func (o *OutputObservable) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *OutputObservable) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *OutputObservable) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetData

`func (o *OutputObservable) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OutputObservable) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OutputObservable) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *OutputObservable) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStartDate

`func (o *OutputObservable) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *OutputObservable) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *OutputObservable) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.


### GetAttachment

`func (o *OutputObservable) GetAttachment() OutputAttachment`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *OutputObservable) GetAttachmentOk() (*OutputAttachment, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *OutputObservable) SetAttachment(v OutputAttachment)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *OutputObservable) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.

### GetTlp

`func (o *OutputObservable) GetTlp() int32`

GetTlp returns the Tlp field if non-nil, zero value otherwise.

### GetTlpOk

`func (o *OutputObservable) GetTlpOk() (*int32, bool)`

GetTlpOk returns a tuple with the Tlp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlp

`func (o *OutputObservable) SetTlp(v int32)`

SetTlp sets Tlp field to given value.


### GetTlpLabel

`func (o *OutputObservable) GetTlpLabel() string`

GetTlpLabel returns the TlpLabel field if non-nil, zero value otherwise.

### GetTlpLabelOk

`func (o *OutputObservable) GetTlpLabelOk() (*string, bool)`

GetTlpLabelOk returns a tuple with the TlpLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlpLabel

`func (o *OutputObservable) SetTlpLabel(v string)`

SetTlpLabel sets TlpLabel field to given value.


### GetPap

`func (o *OutputObservable) GetPap() int32`

GetPap returns the Pap field if non-nil, zero value otherwise.

### GetPapOk

`func (o *OutputObservable) GetPapOk() (*int32, bool)`

GetPapOk returns a tuple with the Pap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPap

`func (o *OutputObservable) SetPap(v int32)`

SetPap sets Pap field to given value.


### GetPapLabel

`func (o *OutputObservable) GetPapLabel() string`

GetPapLabel returns the PapLabel field if non-nil, zero value otherwise.

### GetPapLabelOk

`func (o *OutputObservable) GetPapLabelOk() (*string, bool)`

GetPapLabelOk returns a tuple with the PapLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPapLabel

`func (o *OutputObservable) SetPapLabel(v string)`

SetPapLabel sets PapLabel field to given value.


### GetTags

`func (o *OutputObservable) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OutputObservable) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OutputObservable) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OutputObservable) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIoc

`func (o *OutputObservable) GetIoc() bool`

GetIoc returns the Ioc field if non-nil, zero value otherwise.

### GetIocOk

`func (o *OutputObservable) GetIocOk() (*bool, bool)`

GetIocOk returns a tuple with the Ioc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoc

`func (o *OutputObservable) SetIoc(v bool)`

SetIoc sets Ioc field to given value.


### GetSighted

`func (o *OutputObservable) GetSighted() bool`

GetSighted returns the Sighted field if non-nil, zero value otherwise.

### GetSightedOk

`func (o *OutputObservable) GetSightedOk() (*bool, bool)`

GetSightedOk returns a tuple with the Sighted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSighted

`func (o *OutputObservable) SetSighted(v bool)`

SetSighted sets Sighted field to given value.


### GetSightedAt

`func (o *OutputObservable) GetSightedAt() int64`

GetSightedAt returns the SightedAt field if non-nil, zero value otherwise.

### GetSightedAtOk

`func (o *OutputObservable) GetSightedAtOk() (*int64, bool)`

GetSightedAtOk returns a tuple with the SightedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSightedAt

`func (o *OutputObservable) SetSightedAt(v int64)`

SetSightedAt sets SightedAt field to given value.

### HasSightedAt

`func (o *OutputObservable) HasSightedAt() bool`

HasSightedAt returns a boolean if a field has been set.

### GetReports

`func (o *OutputObservable) GetReports() map[string]interface{}`

GetReports returns the Reports field if non-nil, zero value otherwise.

### GetReportsOk

`func (o *OutputObservable) GetReportsOk() (*map[string]interface{}, bool)`

GetReportsOk returns a tuple with the Reports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReports

`func (o *OutputObservable) SetReports(v map[string]interface{})`

SetReports sets Reports field to given value.


### GetMessage

`func (o *OutputObservable) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OutputObservable) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OutputObservable) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OutputObservable) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetExtraData

`func (o *OutputObservable) GetExtraData() map[string]interface{}`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *OutputObservable) GetExtraDataOk() (*map[string]interface{}, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *OutputObservable) SetExtraData(v map[string]interface{})`

SetExtraData sets ExtraData field to given value.


### GetIgnoreSimilarity

`func (o *OutputObservable) GetIgnoreSimilarity() bool`

GetIgnoreSimilarity returns the IgnoreSimilarity field if non-nil, zero value otherwise.

### GetIgnoreSimilarityOk

`func (o *OutputObservable) GetIgnoreSimilarityOk() (*bool, bool)`

GetIgnoreSimilarityOk returns a tuple with the IgnoreSimilarity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSimilarity

`func (o *OutputObservable) SetIgnoreSimilarity(v bool)`

SetIgnoreSimilarity sets IgnoreSimilarity field to given value.


### GetExternal

`func (o *OutputObservable) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *OutputObservable) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *OutputObservable) SetExternal(v bool)`

SetExternal sets External field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


