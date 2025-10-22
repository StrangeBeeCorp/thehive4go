# FieldValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** |  | 
**Values** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewFieldValues

`func NewFieldValues(field string, ) *FieldValues`

NewFieldValues instantiates a new FieldValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldValuesWithDefaults

`func NewFieldValuesWithDefaults() *FieldValues`

NewFieldValuesWithDefaults instantiates a new FieldValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *FieldValues) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *FieldValues) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *FieldValues) SetField(v string)`

SetField sets Field field to given value.


### GetValues

`func (o *FieldValues) GetValues() []interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *FieldValues) GetValuesOk() (*[]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *FieldValues) SetValues(v []interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *FieldValues) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


