# InputCreateObservableAttachmentOneOfInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Filename of the attachment as stored in TheHive. |
**ContentType** | **string** | MIME type of the attachment. |
**Id** | **string** | ID of the attachment in TheHive. |
**External** | Pointer to **bool** | Whether the attachment is accessible to external users with [TheHive Portal](https://docs.strangebee.com/thehive/administration/thehive-portal/about-thehive-portal/). Platinum or One license required. | [optional]

## Methods

### NewInputCreateObservableAttachmentOneOfInner

`func NewInputCreateObservableAttachmentOneOfInner(name string, contentType string, id string, ) *InputCreateObservableAttachmentOneOfInner`

NewInputCreateObservableAttachmentOneOfInner instantiates a new InputCreateObservableAttachmentOneOfInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputCreateObservableAttachmentOneOfInnerWithDefaults

`func NewInputCreateObservableAttachmentOneOfInnerWithDefaults() *InputCreateObservableAttachmentOneOfInner`

NewInputCreateObservableAttachmentOneOfInnerWithDefaults instantiates a new InputCreateObservableAttachmentOneOfInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InputCreateObservableAttachmentOneOfInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputCreateObservableAttachmentOneOfInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputCreateObservableAttachmentOneOfInner) SetName(v string)`

SetName sets Name field to given value.


### GetContentType

`func (o *InputCreateObservableAttachmentOneOfInner) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *InputCreateObservableAttachmentOneOfInner) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *InputCreateObservableAttachmentOneOfInner) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetId

`func (o *InputCreateObservableAttachmentOneOfInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InputCreateObservableAttachmentOneOfInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InputCreateObservableAttachmentOneOfInner) SetId(v string)`

SetId sets Id field to given value.


### GetExternal

`func (o *InputCreateObservableAttachmentOneOfInner) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *InputCreateObservableAttachmentOneOfInner) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *InputCreateObservableAttachmentOneOfInner) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *InputCreateObservableAttachmentOneOfInner) HasExternal() bool`

HasExternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
