# InputCustomFieldValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of a CustomField | 
**Value** | **interface{}** |  | 
**Order** | Pointer to **int32** |  | [optional] 

## Methods

### NewInputCustomFieldValue

`func NewInputCustomFieldValue(name string, value interface{}, ) *InputCustomFieldValue`

NewInputCustomFieldValue instantiates a new InputCustomFieldValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputCustomFieldValueWithDefaults

`func NewInputCustomFieldValueWithDefaults() *InputCustomFieldValue`

NewInputCustomFieldValueWithDefaults instantiates a new InputCustomFieldValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InputCustomFieldValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputCustomFieldValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputCustomFieldValue) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *InputCustomFieldValue) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InputCustomFieldValue) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InputCustomFieldValue) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *InputCustomFieldValue) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *InputCustomFieldValue) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetOrder

`func (o *InputCustomFieldValue) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *InputCustomFieldValue) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *InputCustomFieldValue) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *InputCustomFieldValue) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


