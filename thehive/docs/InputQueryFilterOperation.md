# InputQueryFilterOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Eq** | Pointer to **map[string]interface{}** |  | [optional]
**And** | Pointer to [**[]Filter**](Filter.md) |  | [optional]
**Or** | Pointer to [**[]Filter**](Filter.md) |  | [optional]
**Not** | Pointer to [**Filter**](Filter.md) |  | [optional]
**Any** | Pointer to **map[string]interface{}** |  | [optional]
**Like** | Pointer to [**FieldValue**](FieldValue.md) |  | [optional]
**Gt** | Pointer to [**FieldValue**](FieldValue.md) |  | [optional]
**Gte** | Pointer to [**FieldValue**](FieldValue.md) |  | [optional]
**Lt** | Pointer to [**FieldValue**](FieldValue.md) |  | [optional]
**Lte** | Pointer to [**FieldValue**](FieldValue.md) |  | [optional]
**Ne** | Pointer to [**FieldValue**](FieldValue.md) |  | [optional]
**In** | Pointer to [**FieldValue**](FieldValue.md) |  | [optional]
**Between** | Pointer to [**FieldValue**](FieldValue.md) |  | [optional]
**Contains** | Pointer to [**FieldValue**](FieldValue.md) |  | [optional]
**EndsWith** | Pointer to [**FieldValue**](FieldValue.md) |  | [optional]
**StartsWith** | Pointer to [**FieldValue**](FieldValue.md) |  | [optional]
**Match** | Pointer to [**FieldValue**](FieldValue.md) |  | [optional]
**UnderscoreId** | Pointer to **string** |  | [optional]
**Name** | **string** |  |

## Methods

### NewInputQueryFilterOperation

`func NewInputQueryFilterOperation(name string, ) *InputQueryFilterOperation`

NewInputQueryFilterOperation instantiates a new InputQueryFilterOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputQueryFilterOperationWithDefaults

`func NewInputQueryFilterOperationWithDefaults() *InputQueryFilterOperation`

NewInputQueryFilterOperationWithDefaults instantiates a new InputQueryFilterOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEq

`func (o *InputQueryFilterOperation) GetEq() map[string]interface{}`

GetEq returns the Eq field if non-nil, zero value otherwise.

### GetEqOk

`func (o *InputQueryFilterOperation) GetEqOk() (*map[string]interface{}, bool)`

GetEqOk returns a tuple with the Eq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEq

`func (o *InputQueryFilterOperation) SetEq(v map[string]interface{})`

SetEq sets Eq field to given value.

### HasEq

`func (o *InputQueryFilterOperation) HasEq() bool`

HasEq returns a boolean if a field has been set.

### GetAnd

`func (o *InputQueryFilterOperation) GetAnd() []Filter`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *InputQueryFilterOperation) GetAndOk() (*[]Filter, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *InputQueryFilterOperation) SetAnd(v []Filter)`

SetAnd sets And field to given value.

### HasAnd

`func (o *InputQueryFilterOperation) HasAnd() bool`

HasAnd returns a boolean if a field has been set.

### GetOr

`func (o *InputQueryFilterOperation) GetOr() []Filter`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *InputQueryFilterOperation) GetOrOk() (*[]Filter, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *InputQueryFilterOperation) SetOr(v []Filter)`

SetOr sets Or field to given value.

### HasOr

`func (o *InputQueryFilterOperation) HasOr() bool`

HasOr returns a boolean if a field has been set.

### GetNot

`func (o *InputQueryFilterOperation) GetNot() Filter`

GetNot returns the Not field if non-nil, zero value otherwise.

### GetNotOk

`func (o *InputQueryFilterOperation) GetNotOk() (*Filter, bool)`

GetNotOk returns a tuple with the Not field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNot

`func (o *InputQueryFilterOperation) SetNot(v Filter)`

SetNot sets Not field to given value.

### HasNot

`func (o *InputQueryFilterOperation) HasNot() bool`

HasNot returns a boolean if a field has been set.

### GetAny

`func (o *InputQueryFilterOperation) GetAny() map[string]interface{}`

GetAny returns the Any field if non-nil, zero value otherwise.

### GetAnyOk

`func (o *InputQueryFilterOperation) GetAnyOk() (*map[string]interface{}, bool)`

GetAnyOk returns a tuple with the Any field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAny

`func (o *InputQueryFilterOperation) SetAny(v map[string]interface{})`

SetAny sets Any field to given value.

### HasAny

`func (o *InputQueryFilterOperation) HasAny() bool`

HasAny returns a boolean if a field has been set.

### GetLike

`func (o *InputQueryFilterOperation) GetLike() FieldValue`

GetLike returns the Like field if non-nil, zero value otherwise.

### GetLikeOk

`func (o *InputQueryFilterOperation) GetLikeOk() (*FieldValue, bool)`

GetLikeOk returns a tuple with the Like field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLike

`func (o *InputQueryFilterOperation) SetLike(v FieldValue)`

SetLike sets Like field to given value.

### HasLike

`func (o *InputQueryFilterOperation) HasLike() bool`

HasLike returns a boolean if a field has been set.

### GetGt

`func (o *InputQueryFilterOperation) GetGt() FieldValue`

GetGt returns the Gt field if non-nil, zero value otherwise.

### GetGtOk

`func (o *InputQueryFilterOperation) GetGtOk() (*FieldValue, bool)`

GetGtOk returns a tuple with the Gt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGt

`func (o *InputQueryFilterOperation) SetGt(v FieldValue)`

SetGt sets Gt field to given value.

### HasGt

`func (o *InputQueryFilterOperation) HasGt() bool`

HasGt returns a boolean if a field has been set.

### GetGte

`func (o *InputQueryFilterOperation) GetGte() FieldValue`

GetGte returns the Gte field if non-nil, zero value otherwise.

### GetGteOk

`func (o *InputQueryFilterOperation) GetGteOk() (*FieldValue, bool)`

GetGteOk returns a tuple with the Gte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGte

`func (o *InputQueryFilterOperation) SetGte(v FieldValue)`

SetGte sets Gte field to given value.

### HasGte

`func (o *InputQueryFilterOperation) HasGte() bool`

HasGte returns a boolean if a field has been set.

### GetLt

`func (o *InputQueryFilterOperation) GetLt() FieldValue`

GetLt returns the Lt field if non-nil, zero value otherwise.

### GetLtOk

`func (o *InputQueryFilterOperation) GetLtOk() (*FieldValue, bool)`

GetLtOk returns a tuple with the Lt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLt

`func (o *InputQueryFilterOperation) SetLt(v FieldValue)`

SetLt sets Lt field to given value.

### HasLt

`func (o *InputQueryFilterOperation) HasLt() bool`

HasLt returns a boolean if a field has been set.

### GetLte

`func (o *InputQueryFilterOperation) GetLte() FieldValue`

GetLte returns the Lte field if non-nil, zero value otherwise.

### GetLteOk

`func (o *InputQueryFilterOperation) GetLteOk() (*FieldValue, bool)`

GetLteOk returns a tuple with the Lte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLte

`func (o *InputQueryFilterOperation) SetLte(v FieldValue)`

SetLte sets Lte field to given value.

### HasLte

`func (o *InputQueryFilterOperation) HasLte() bool`

HasLte returns a boolean if a field has been set.

### GetNe

`func (o *InputQueryFilterOperation) GetNe() FieldValue`

GetNe returns the Ne field if non-nil, zero value otherwise.

### GetNeOk

`func (o *InputQueryFilterOperation) GetNeOk() (*FieldValue, bool)`

GetNeOk returns a tuple with the Ne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNe

`func (o *InputQueryFilterOperation) SetNe(v FieldValue)`

SetNe sets Ne field to given value.

### HasNe

`func (o *InputQueryFilterOperation) HasNe() bool`

HasNe returns a boolean if a field has been set.

### GetIn

`func (o *InputQueryFilterOperation) GetIn() FieldValue`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *InputQueryFilterOperation) GetInOk() (*FieldValue, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *InputQueryFilterOperation) SetIn(v FieldValue)`

SetIn sets In field to given value.

### HasIn

`func (o *InputQueryFilterOperation) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetBetween

`func (o *InputQueryFilterOperation) GetBetween() FieldValue`

GetBetween returns the Between field if non-nil, zero value otherwise.

### GetBetweenOk

`func (o *InputQueryFilterOperation) GetBetweenOk() (*FieldValue, bool)`

GetBetweenOk returns a tuple with the Between field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetween

`func (o *InputQueryFilterOperation) SetBetween(v FieldValue)`

SetBetween sets Between field to given value.

### HasBetween

`func (o *InputQueryFilterOperation) HasBetween() bool`

HasBetween returns a boolean if a field has been set.

### GetContains

`func (o *InputQueryFilterOperation) GetContains() FieldValue`

GetContains returns the Contains field if non-nil, zero value otherwise.

### GetContainsOk

`func (o *InputQueryFilterOperation) GetContainsOk() (*FieldValue, bool)`

GetContainsOk returns a tuple with the Contains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContains

`func (o *InputQueryFilterOperation) SetContains(v FieldValue)`

SetContains sets Contains field to given value.

### HasContains

`func (o *InputQueryFilterOperation) HasContains() bool`

HasContains returns a boolean if a field has been set.

### GetEndsWith

`func (o *InputQueryFilterOperation) GetEndsWith() FieldValue`

GetEndsWith returns the EndsWith field if non-nil, zero value otherwise.

### GetEndsWithOk

`func (o *InputQueryFilterOperation) GetEndsWithOk() (*FieldValue, bool)`

GetEndsWithOk returns a tuple with the EndsWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsWith

`func (o *InputQueryFilterOperation) SetEndsWith(v FieldValue)`

SetEndsWith sets EndsWith field to given value.

### HasEndsWith

`func (o *InputQueryFilterOperation) HasEndsWith() bool`

HasEndsWith returns a boolean if a field has been set.

### GetStartsWith

`func (o *InputQueryFilterOperation) GetStartsWith() FieldValue`

GetStartsWith returns the StartsWith field if non-nil, zero value otherwise.

### GetStartsWithOk

`func (o *InputQueryFilterOperation) GetStartsWithOk() (*FieldValue, bool)`

GetStartsWithOk returns a tuple with the StartsWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsWith

`func (o *InputQueryFilterOperation) SetStartsWith(v FieldValue)`

SetStartsWith sets StartsWith field to given value.

### HasStartsWith

`func (o *InputQueryFilterOperation) HasStartsWith() bool`

HasStartsWith returns a boolean if a field has been set.

### GetMatch

`func (o *InputQueryFilterOperation) GetMatch() FieldValue`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *InputQueryFilterOperation) GetMatchOk() (*FieldValue, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *InputQueryFilterOperation) SetMatch(v FieldValue)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *InputQueryFilterOperation) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### GetUnderscoreId

`func (o *InputQueryFilterOperation) GetUnderscoreId() string`

GetUnderscoreId returns the UnderscoreId field if non-nil, zero value otherwise.

### GetUnderscoreIdOk

`func (o *InputQueryFilterOperation) GetUnderscoreIdOk() (*string, bool)`

GetUnderscoreIdOk returns a tuple with the UnderscoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreId

`func (o *InputQueryFilterOperation) SetUnderscoreId(v string)`

SetUnderscoreId sets UnderscoreId field to given value.

### HasUnderscoreId

`func (o *InputQueryFilterOperation) HasUnderscoreId() bool`

HasUnderscoreId returns a boolean if a field has been set.

### GetName

`func (o *InputQueryFilterOperation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputQueryFilterOperation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputQueryFilterOperation) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
