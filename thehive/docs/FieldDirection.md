# FieldDirection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** |  | 
**Direction** | [**Direction**](Direction.md) |  | 

## Methods

### NewFieldDirection

`func NewFieldDirection(field string, direction Direction, ) *FieldDirection`

NewFieldDirection instantiates a new FieldDirection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldDirectionWithDefaults

`func NewFieldDirectionWithDefaults() *FieldDirection`

NewFieldDirectionWithDefaults instantiates a new FieldDirection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *FieldDirection) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *FieldDirection) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *FieldDirection) SetField(v string)`

SetField sets Field field to given value.


### GetDirection

`func (o *FieldDirection) GetDirection() Direction`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *FieldDirection) GetDirectionOk() (*Direction, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *FieldDirection) SetDirection(v Direction)`

SetDirection sets Direction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


