# InputObservableType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the observable type. Once created, this becomes a valid value for the &#x60;dataType&#x60; field when creating or updating observables. |
**IsAttachment** | Pointer to **bool** | When &#x60;true&#x60;, observables of this type require a file upload instead of a text value. | [optional]
**IsCaseSensitive** | Pointer to **bool** | When &#x60;false&#x60;, new observables of this type are automatically normalized to lowercase upon creation to ensure that values are treated as identical regardless of letter casing. Existing observables aren&#39;t affected. | [optional]

## Methods

### NewInputObservableType

`func NewInputObservableType(name string, ) *InputObservableType`

NewInputObservableType instantiates a new InputObservableType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputObservableTypeWithDefaults

`func NewInputObservableTypeWithDefaults() *InputObservableType`

NewInputObservableTypeWithDefaults instantiates a new InputObservableType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InputObservableType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputObservableType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputObservableType) SetName(v string)`

SetName sets Name field to given value.


### GetIsAttachment

`func (o *InputObservableType) GetIsAttachment() bool`

GetIsAttachment returns the IsAttachment field if non-nil, zero value otherwise.

### GetIsAttachmentOk

`func (o *InputObservableType) GetIsAttachmentOk() (*bool, bool)`

GetIsAttachmentOk returns a tuple with the IsAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAttachment

`func (o *InputObservableType) SetIsAttachment(v bool)`

SetIsAttachment sets IsAttachment field to given value.

### HasIsAttachment

`func (o *InputObservableType) HasIsAttachment() bool`

HasIsAttachment returns a boolean if a field has been set.

### GetIsCaseSensitive

`func (o *InputObservableType) GetIsCaseSensitive() bool`

GetIsCaseSensitive returns the IsCaseSensitive field if non-nil, zero value otherwise.

### GetIsCaseSensitiveOk

`func (o *InputObservableType) GetIsCaseSensitiveOk() (*bool, bool)`

GetIsCaseSensitiveOk returns a tuple with the IsCaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCaseSensitive

`func (o *InputObservableType) SetIsCaseSensitive(v bool)`

SetIsCaseSensitive sets IsCaseSensitive field to given value.

### HasIsCaseSensitive

`func (o *InputObservableType) HasIsCaseSensitive() bool`

HasIsCaseSensitive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
