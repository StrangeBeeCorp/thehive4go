# InputValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Predicate** | **string** |  | 
**Entry** | Pointer to [**[]InputEntry**](InputEntry.md) |  | [optional] 

## Methods

### NewInputValue

`func NewInputValue(predicate string, ) *InputValue`

NewInputValue instantiates a new InputValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputValueWithDefaults

`func NewInputValueWithDefaults() *InputValue`

NewInputValueWithDefaults instantiates a new InputValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPredicate

`func (o *InputValue) GetPredicate() string`

GetPredicate returns the Predicate field if non-nil, zero value otherwise.

### GetPredicateOk

`func (o *InputValue) GetPredicateOk() (*string, bool)`

GetPredicateOk returns a tuple with the Predicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicate

`func (o *InputValue) SetPredicate(v string)`

SetPredicate sets Predicate field to given value.


### GetEntry

`func (o *InputValue) GetEntry() []InputEntry`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *InputValue) GetEntryOk() (*[]InputEntry, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *InputValue) SetEntry(v []InputEntry)`

SetEntry sets Entry field to given value.

### HasEntry

`func (o *InputValue) HasEntry() bool`

HasEntry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


