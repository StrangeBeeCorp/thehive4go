# Filter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | Pointer to [**[]Filter**](Filter.md) |  | [optional] 
**Any** | **interface{}** |  | 
**Between** | [**FieldFromTo**](FieldFromTo.md) |  | 
**Contains** | **string** |  | 
**EndsWith** | [**FieldValue**](FieldValue.md) |  | 
**Eq** | [**FieldValue**](FieldValue.md) |  | 
**Gt** | [**FieldValue**](FieldValue.md) |  | 
**Gte** | [**FieldValue**](FieldValue.md) |  | 
**UnderscoreId** | **string** |  | 
**In** | [**FieldValues**](FieldValues.md) |  | 
**Like** | [**FieldValue**](FieldValue.md) |  | 
**Lt** | [**FieldValue**](FieldValue.md) |  | 
**Lte** | [**FieldValue**](FieldValue.md) |  | 
**Match** | [**FieldValue**](FieldValue.md) |  | 
**Ne** | [**FieldValue**](FieldValue.md) |  | 
**Not** | [**Filter**](Filter.md) |  | 
**Or** | Pointer to [**[]Filter**](Filter.md) |  | [optional] 
**StartsWith** | [**FieldValue**](FieldValue.md) |  | 

## Methods

### NewFilter

`func NewFilter(any interface{}, between FieldFromTo, contains string, endsWith FieldValue, eq FieldValue, gt FieldValue, gte FieldValue, underscoreId string, in FieldValues, like FieldValue, lt FieldValue, lte FieldValue, match FieldValue, ne FieldValue, not Filter, startsWith FieldValue, ) *Filter`

NewFilter instantiates a new Filter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterWithDefaults

`func NewFilterWithDefaults() *Filter`

NewFilterWithDefaults instantiates a new Filter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *Filter) GetAnd() []Filter`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *Filter) GetAndOk() (*[]Filter, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *Filter) SetAnd(v []Filter)`

SetAnd sets And field to given value.

### HasAnd

`func (o *Filter) HasAnd() bool`

HasAnd returns a boolean if a field has been set.

### GetAny

`func (o *Filter) GetAny() interface{}`

GetAny returns the Any field if non-nil, zero value otherwise.

### GetAnyOk

`func (o *Filter) GetAnyOk() (*interface{}, bool)`

GetAnyOk returns a tuple with the Any field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAny

`func (o *Filter) SetAny(v interface{})`

SetAny sets Any field to given value.


### SetAnyNil

`func (o *Filter) SetAnyNil(b bool)`

 SetAnyNil sets the value for Any to be an explicit nil

### UnsetAny
`func (o *Filter) UnsetAny()`

UnsetAny ensures that no value is present for Any, not even an explicit nil
### GetBetween

`func (o *Filter) GetBetween() FieldFromTo`

GetBetween returns the Between field if non-nil, zero value otherwise.

### GetBetweenOk

`func (o *Filter) GetBetweenOk() (*FieldFromTo, bool)`

GetBetweenOk returns a tuple with the Between field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetween

`func (o *Filter) SetBetween(v FieldFromTo)`

SetBetween sets Between field to given value.


### GetContains

`func (o *Filter) GetContains() string`

GetContains returns the Contains field if non-nil, zero value otherwise.

### GetContainsOk

`func (o *Filter) GetContainsOk() (*string, bool)`

GetContainsOk returns a tuple with the Contains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContains

`func (o *Filter) SetContains(v string)`

SetContains sets Contains field to given value.


### GetEndsWith

`func (o *Filter) GetEndsWith() FieldValue`

GetEndsWith returns the EndsWith field if non-nil, zero value otherwise.

### GetEndsWithOk

`func (o *Filter) GetEndsWithOk() (*FieldValue, bool)`

GetEndsWithOk returns a tuple with the EndsWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsWith

`func (o *Filter) SetEndsWith(v FieldValue)`

SetEndsWith sets EndsWith field to given value.


### GetEq

`func (o *Filter) GetEq() FieldValue`

GetEq returns the Eq field if non-nil, zero value otherwise.

### GetEqOk

`func (o *Filter) GetEqOk() (*FieldValue, bool)`

GetEqOk returns a tuple with the Eq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEq

`func (o *Filter) SetEq(v FieldValue)`

SetEq sets Eq field to given value.


### GetGt

`func (o *Filter) GetGt() FieldValue`

GetGt returns the Gt field if non-nil, zero value otherwise.

### GetGtOk

`func (o *Filter) GetGtOk() (*FieldValue, bool)`

GetGtOk returns a tuple with the Gt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGt

`func (o *Filter) SetGt(v FieldValue)`

SetGt sets Gt field to given value.


### GetGte

`func (o *Filter) GetGte() FieldValue`

GetGte returns the Gte field if non-nil, zero value otherwise.

### GetGteOk

`func (o *Filter) GetGteOk() (*FieldValue, bool)`

GetGteOk returns a tuple with the Gte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGte

`func (o *Filter) SetGte(v FieldValue)`

SetGte sets Gte field to given value.


### GetUnderscoreId

`func (o *Filter) GetUnderscoreId() string`

GetUnderscoreId returns the UnderscoreId field if non-nil, zero value otherwise.

### GetUnderscoreIdOk

`func (o *Filter) GetUnderscoreIdOk() (*string, bool)`

GetUnderscoreIdOk returns a tuple with the UnderscoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreId

`func (o *Filter) SetUnderscoreId(v string)`

SetUnderscoreId sets UnderscoreId field to given value.


### GetIn

`func (o *Filter) GetIn() FieldValues`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *Filter) GetInOk() (*FieldValues, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *Filter) SetIn(v FieldValues)`

SetIn sets In field to given value.


### GetLike

`func (o *Filter) GetLike() FieldValue`

GetLike returns the Like field if non-nil, zero value otherwise.

### GetLikeOk

`func (o *Filter) GetLikeOk() (*FieldValue, bool)`

GetLikeOk returns a tuple with the Like field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLike

`func (o *Filter) SetLike(v FieldValue)`

SetLike sets Like field to given value.


### GetLt

`func (o *Filter) GetLt() FieldValue`

GetLt returns the Lt field if non-nil, zero value otherwise.

### GetLtOk

`func (o *Filter) GetLtOk() (*FieldValue, bool)`

GetLtOk returns a tuple with the Lt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLt

`func (o *Filter) SetLt(v FieldValue)`

SetLt sets Lt field to given value.


### GetLte

`func (o *Filter) GetLte() FieldValue`

GetLte returns the Lte field if non-nil, zero value otherwise.

### GetLteOk

`func (o *Filter) GetLteOk() (*FieldValue, bool)`

GetLteOk returns a tuple with the Lte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLte

`func (o *Filter) SetLte(v FieldValue)`

SetLte sets Lte field to given value.


### GetMatch

`func (o *Filter) GetMatch() FieldValue`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *Filter) GetMatchOk() (*FieldValue, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *Filter) SetMatch(v FieldValue)`

SetMatch sets Match field to given value.


### GetNe

`func (o *Filter) GetNe() FieldValue`

GetNe returns the Ne field if non-nil, zero value otherwise.

### GetNeOk

`func (o *Filter) GetNeOk() (*FieldValue, bool)`

GetNeOk returns a tuple with the Ne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNe

`func (o *Filter) SetNe(v FieldValue)`

SetNe sets Ne field to given value.


### GetNot

`func (o *Filter) GetNot() Filter`

GetNot returns the Not field if non-nil, zero value otherwise.

### GetNotOk

`func (o *Filter) GetNotOk() (*Filter, bool)`

GetNotOk returns a tuple with the Not field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNot

`func (o *Filter) SetNot(v Filter)`

SetNot sets Not field to given value.


### GetOr

`func (o *Filter) GetOr() []Filter`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *Filter) GetOrOk() (*[]Filter, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *Filter) SetOr(v []Filter)`

SetOr sets Or field to given value.

### HasOr

`func (o *Filter) HasOr() bool`

HasOr returns a boolean if a field has been set.

### GetStartsWith

`func (o *Filter) GetStartsWith() FieldValue`

GetStartsWith returns the StartsWith field if non-nil, zero value otherwise.

### GetStartsWithOk

`func (o *Filter) GetStartsWithOk() (*FieldValue, bool)`

GetStartsWithOk returns a tuple with the StartsWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsWith

`func (o *Filter) SetStartsWith(v FieldValue)`

SetStartsWith sets StartsWith field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


