# InputCreateObservable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** |  | 
**Data** | Pointer to [**InputObservableData**](InputObservableData.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **int64** | The &#x60;startDate&#x60; field for &#x60;InputCreateObservable&#x60; object is mapped on &#x60;_createdDate&#x60; field. This field should not be used. | [optional] 
**Attachment** | Pointer to [**InputCreateObservableAttachment**](InputCreateObservableAttachment.md) |  | [optional] 
**Tlp** | Pointer to **int32** |  | [optional] [default to 2]
**Pap** | Pointer to **int32** |  | [optional] [default to 2]
**Tags** | Pointer to **[]string** |  | [optional] 
**Ioc** | Pointer to **bool** |  | [optional] [default to false]
**Sighted** | Pointer to **bool** |  | [optional] [default to false]
**SightedAt** | Pointer to **int64** |  | [optional] 
**IgnoreSimilarity** | Pointer to **bool** |  | [optional] [default to false]
**IsZip** | Pointer to **bool** | If set to true, the file is unzipped using the &#x60;zipPassword&#x60; and each file in the zip is treated as an observable | [optional] 
**ZipPassword** | Pointer to **string** |  | [optional] 

## Methods

### NewInputCreateObservable

`func NewInputCreateObservable(dataType string, ) *InputCreateObservable`

NewInputCreateObservable instantiates a new InputCreateObservable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputCreateObservableWithDefaults

`func NewInputCreateObservableWithDefaults() *InputCreateObservable`

NewInputCreateObservableWithDefaults instantiates a new InputCreateObservable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *InputCreateObservable) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *InputCreateObservable) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *InputCreateObservable) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetData

`func (o *InputCreateObservable) GetData() InputObservableData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InputCreateObservable) GetDataOk() (*InputObservableData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InputCreateObservable) SetData(v InputObservableData)`

SetData sets Data field to given value.

### HasData

`func (o *InputCreateObservable) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *InputCreateObservable) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InputCreateObservable) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InputCreateObservable) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InputCreateObservable) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStartDate

`func (o *InputCreateObservable) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InputCreateObservable) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InputCreateObservable) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InputCreateObservable) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetAttachment

`func (o *InputCreateObservable) GetAttachment() InputCreateObservableAttachment`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *InputCreateObservable) GetAttachmentOk() (*InputCreateObservableAttachment, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *InputCreateObservable) SetAttachment(v InputCreateObservableAttachment)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *InputCreateObservable) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.

### GetTlp

`func (o *InputCreateObservable) GetTlp() int32`

GetTlp returns the Tlp field if non-nil, zero value otherwise.

### GetTlpOk

`func (o *InputCreateObservable) GetTlpOk() (*int32, bool)`

GetTlpOk returns a tuple with the Tlp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlp

`func (o *InputCreateObservable) SetTlp(v int32)`

SetTlp sets Tlp field to given value.

### HasTlp

`func (o *InputCreateObservable) HasTlp() bool`

HasTlp returns a boolean if a field has been set.

### GetPap

`func (o *InputCreateObservable) GetPap() int32`

GetPap returns the Pap field if non-nil, zero value otherwise.

### GetPapOk

`func (o *InputCreateObservable) GetPapOk() (*int32, bool)`

GetPapOk returns a tuple with the Pap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPap

`func (o *InputCreateObservable) SetPap(v int32)`

SetPap sets Pap field to given value.

### HasPap

`func (o *InputCreateObservable) HasPap() bool`

HasPap returns a boolean if a field has been set.

### GetTags

`func (o *InputCreateObservable) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InputCreateObservable) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InputCreateObservable) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InputCreateObservable) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIoc

`func (o *InputCreateObservable) GetIoc() bool`

GetIoc returns the Ioc field if non-nil, zero value otherwise.

### GetIocOk

`func (o *InputCreateObservable) GetIocOk() (*bool, bool)`

GetIocOk returns a tuple with the Ioc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoc

`func (o *InputCreateObservable) SetIoc(v bool)`

SetIoc sets Ioc field to given value.

### HasIoc

`func (o *InputCreateObservable) HasIoc() bool`

HasIoc returns a boolean if a field has been set.

### GetSighted

`func (o *InputCreateObservable) GetSighted() bool`

GetSighted returns the Sighted field if non-nil, zero value otherwise.

### GetSightedOk

`func (o *InputCreateObservable) GetSightedOk() (*bool, bool)`

GetSightedOk returns a tuple with the Sighted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSighted

`func (o *InputCreateObservable) SetSighted(v bool)`

SetSighted sets Sighted field to given value.

### HasSighted

`func (o *InputCreateObservable) HasSighted() bool`

HasSighted returns a boolean if a field has been set.

### GetSightedAt

`func (o *InputCreateObservable) GetSightedAt() int64`

GetSightedAt returns the SightedAt field if non-nil, zero value otherwise.

### GetSightedAtOk

`func (o *InputCreateObservable) GetSightedAtOk() (*int64, bool)`

GetSightedAtOk returns a tuple with the SightedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSightedAt

`func (o *InputCreateObservable) SetSightedAt(v int64)`

SetSightedAt sets SightedAt field to given value.

### HasSightedAt

`func (o *InputCreateObservable) HasSightedAt() bool`

HasSightedAt returns a boolean if a field has been set.

### GetIgnoreSimilarity

`func (o *InputCreateObservable) GetIgnoreSimilarity() bool`

GetIgnoreSimilarity returns the IgnoreSimilarity field if non-nil, zero value otherwise.

### GetIgnoreSimilarityOk

`func (o *InputCreateObservable) GetIgnoreSimilarityOk() (*bool, bool)`

GetIgnoreSimilarityOk returns a tuple with the IgnoreSimilarity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSimilarity

`func (o *InputCreateObservable) SetIgnoreSimilarity(v bool)`

SetIgnoreSimilarity sets IgnoreSimilarity field to given value.

### HasIgnoreSimilarity

`func (o *InputCreateObservable) HasIgnoreSimilarity() bool`

HasIgnoreSimilarity returns a boolean if a field has been set.

### GetIsZip

`func (o *InputCreateObservable) GetIsZip() bool`

GetIsZip returns the IsZip field if non-nil, zero value otherwise.

### GetIsZipOk

`func (o *InputCreateObservable) GetIsZipOk() (*bool, bool)`

GetIsZipOk returns a tuple with the IsZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsZip

`func (o *InputCreateObservable) SetIsZip(v bool)`

SetIsZip sets IsZip field to given value.

### HasIsZip

`func (o *InputCreateObservable) HasIsZip() bool`

HasIsZip returns a boolean if a field has been set.

### GetZipPassword

`func (o *InputCreateObservable) GetZipPassword() string`

GetZipPassword returns the ZipPassword field if non-nil, zero value otherwise.

### GetZipPasswordOk

`func (o *InputCreateObservable) GetZipPasswordOk() (*string, bool)`

GetZipPasswordOk returns a tuple with the ZipPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipPassword

`func (o *InputCreateObservable) SetZipPassword(v string)`

SetZipPassword sets ZipPassword field to given value.

### HasZipPassword

`func (o *InputCreateObservable) HasZipPassword() bool`

HasZipPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


