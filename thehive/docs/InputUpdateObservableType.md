# InputUpdateObservableType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsCaseSensitive** | Pointer to **bool** | When &#x60;false&#x60;, new observables of this type are automatically normalized to lowercase upon creation to ensure that values are treated as identical regardless of letter casing. Existing observables aren&#39;t affected. | [optional]

## Methods

### NewInputUpdateObservableType

`func NewInputUpdateObservableType() *InputUpdateObservableType`

NewInputUpdateObservableType instantiates a new InputUpdateObservableType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputUpdateObservableTypeWithDefaults

`func NewInputUpdateObservableTypeWithDefaults() *InputUpdateObservableType`

NewInputUpdateObservableTypeWithDefaults instantiates a new InputUpdateObservableType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsCaseSensitive

`func (o *InputUpdateObservableType) GetIsCaseSensitive() bool`

GetIsCaseSensitive returns the IsCaseSensitive field if non-nil, zero value otherwise.

### GetIsCaseSensitiveOk

`func (o *InputUpdateObservableType) GetIsCaseSensitiveOk() (*bool, bool)`

GetIsCaseSensitiveOk returns a tuple with the IsCaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCaseSensitive

`func (o *InputUpdateObservableType) SetIsCaseSensitive(v bool)`

SetIsCaseSensitive sets IsCaseSensitive field to given value.

### HasIsCaseSensitive

`func (o *InputUpdateObservableType) HasIsCaseSensitive() bool`

HasIsCaseSensitive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
