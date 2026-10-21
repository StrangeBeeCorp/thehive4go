# InputComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Text of the comment. |
**External** | Pointer to **bool** | Whether the comment is visible to external users through [TheHive Portal](https://docs.strangebee.com/thehive/administration/thehive-portal/about-thehive-portal/). Set to &#x60;true&#x60; to share the comment with external portal users. Requires a Platinum or One license and the case to be shared through TheHive Portal. Has no effect on alert comments, which are always internal. | [optional] [default to false]

## Methods

### NewInputComment

`func NewInputComment(message string, ) *InputComment`

NewInputComment instantiates a new InputComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputCommentWithDefaults

`func NewInputCommentWithDefaults() *InputComment`

NewInputCommentWithDefaults instantiates a new InputComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *InputComment) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InputComment) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InputComment) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetExternal

`func (o *InputComment) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *InputComment) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *InputComment) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *InputComment) HasExternal() bool`

HasExternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
