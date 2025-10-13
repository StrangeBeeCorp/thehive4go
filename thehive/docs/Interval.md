# Interval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **int32** |  | 
**Unit** | [**IntervalUnit**](IntervalUnit.md) |  | 

## Methods

### NewInterval

`func NewInterval(value int32, unit IntervalUnit, ) *Interval`

NewInterval instantiates a new Interval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntervalWithDefaults

`func NewIntervalWithDefaults() *Interval`

NewIntervalWithDefaults instantiates a new Interval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *Interval) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Interval) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Interval) SetValue(v int32)`

SetValue sets Value field to given value.


### GetUnit

`func (o *Interval) GetUnit() IntervalUnit`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Interval) GetUnitOk() (*IntervalUnit, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Interval) SetUnit(v IntervalUnit)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


