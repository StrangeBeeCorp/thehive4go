# InputUpdateAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**External** | Pointer to **bool** | Whether the attachment is accessible to external users with [TheHive Portal](https://docs.strangebee.com/thehive/administration/thehive-portal/about-thehive-portal/). Requires a Platinum or One license. If absent, the current value is unchanged. | [optional]

## Methods

### NewInputUpdateAttachment

`func NewInputUpdateAttachment() *InputUpdateAttachment`

NewInputUpdateAttachment instantiates a new InputUpdateAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputUpdateAttachmentWithDefaults

`func NewInputUpdateAttachmentWithDefaults() *InputUpdateAttachment`

NewInputUpdateAttachmentWithDefaults instantiates a new InputUpdateAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternal

`func (o *InputUpdateAttachment) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *InputUpdateAttachment) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *InputUpdateAttachment) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *InputUpdateAttachment) HasExternal() bool`

HasExternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
