# InputQueryNamedOperation

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
**ExtraData** | Pointer to **[]string** |  | [optional]
**From** | **int32** |  |
**To** | **int32** |  |
**Fields** | Pointer to **[]map[string]interface{}** |  | [optional]

## Methods

### NewInputQueryNamedOperation

`func NewInputQueryNamedOperation(name string, from int32, to int32, ) *InputQueryNamedOperation`

NewInputQueryNamedOperation instantiates a new InputQueryNamedOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputQueryNamedOperationWithDefaults

`func NewInputQueryNamedOperationWithDefaults() *InputQueryNamedOperation`

NewInputQueryNamedOperationWithDefaults instantiates a new InputQueryNamedOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEq

`func (o *InputQueryNamedOperation) GetEq() map[string]interface{}`

GetEq returns the Eq field if non-nil, zero value otherwise.

### GetEqOk

`func (o *InputQueryNamedOperation) GetEqOk() (*map[string]interface{}, bool)`

GetEqOk returns a tuple with the Eq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEq

`func (o *InputQueryNamedOperation) SetEq(v map[string]interface{})`

SetEq sets Eq field to given value.

### HasEq

`func (o *InputQueryNamedOperation) HasEq() bool`

HasEq returns a boolean if a field has been set.

### GetAnd

`func (o *InputQueryNamedOperation) GetAnd() []Filter`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *InputQueryNamedOperation) GetAndOk() (*[]Filter, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *InputQueryNamedOperation) SetAnd(v []Filter)`

SetAnd sets And field to given value.

### HasAnd

`func (o *InputQueryNamedOperation) HasAnd() bool`

HasAnd returns a boolean if a field has been set.

### GetOr

`func (o *InputQueryNamedOperation) GetOr() []Filter`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *InputQueryNamedOperation) GetOrOk() (*[]Filter, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *InputQueryNamedOperation) SetOr(v []Filter)`

SetOr sets Or field to given value.

### HasOr

`func (o *InputQueryNamedOperation) HasOr() bool`

HasOr returns a boolean if a field has been set.

### GetNot

`func (o *InputQueryNamedOperation) GetNot() Filter`

GetNot returns the Not field if non-nil, zero value otherwise.

### GetNotOk

`func (o *InputQueryNamedOperation) GetNotOk() (*Filter, bool)`

GetNotOk returns a tuple with the Not field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNot

`func (o *InputQueryNamedOperation) SetNot(v Filter)`

SetNot sets Not field to given value.

### HasNot

`func (o *InputQueryNamedOperation) HasNot() bool`

HasNot returns a boolean if a field has been set.

### GetAny

`func (o *InputQueryNamedOperation) GetAny() map[string]interface{}`

GetAny returns the Any field if non-nil, zero value otherwise.

### GetAnyOk

`func (o *InputQueryNamedOperation) GetAnyOk() (*map[string]interface{}, bool)`

GetAnyOk returns a tuple with the Any field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAny

`func (o *InputQueryNamedOperation) SetAny(v map[string]interface{})`

SetAny sets Any field to given value.

### HasAny

`func (o *InputQueryNamedOperation) HasAny() bool`

HasAny returns a boolean if a field has been set.

### GetLike

`func (o *InputQueryNamedOperation) GetLike() FieldValue`

GetLike returns the Like field if non-nil, zero value otherwise.

### GetLikeOk

`func (o *InputQueryNamedOperation) GetLikeOk() (*FieldValue, bool)`

GetLikeOk returns a tuple with the Like field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLike

`func (o *InputQueryNamedOperation) SetLike(v FieldValue)`

SetLike sets Like field to given value.

### HasLike

`func (o *InputQueryNamedOperation) HasLike() bool`

HasLike returns a boolean if a field has been set.

### GetGt

`func (o *InputQueryNamedOperation) GetGt() FieldValue`

GetGt returns the Gt field if non-nil, zero value otherwise.

### GetGtOk

`func (o *InputQueryNamedOperation) GetGtOk() (*FieldValue, bool)`

GetGtOk returns a tuple with the Gt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGt

`func (o *InputQueryNamedOperation) SetGt(v FieldValue)`

SetGt sets Gt field to given value.

### HasGt

`func (o *InputQueryNamedOperation) HasGt() bool`

HasGt returns a boolean if a field has been set.

### GetGte

`func (o *InputQueryNamedOperation) GetGte() FieldValue`

GetGte returns the Gte field if non-nil, zero value otherwise.

### GetGteOk

`func (o *InputQueryNamedOperation) GetGteOk() (*FieldValue, bool)`

GetGteOk returns a tuple with the Gte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGte

`func (o *InputQueryNamedOperation) SetGte(v FieldValue)`

SetGte sets Gte field to given value.

### HasGte

`func (o *InputQueryNamedOperation) HasGte() bool`

HasGte returns a boolean if a field has been set.

### GetLt

`func (o *InputQueryNamedOperation) GetLt() FieldValue`

GetLt returns the Lt field if non-nil, zero value otherwise.

### GetLtOk

`func (o *InputQueryNamedOperation) GetLtOk() (*FieldValue, bool)`

GetLtOk returns a tuple with the Lt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLt

`func (o *InputQueryNamedOperation) SetLt(v FieldValue)`

SetLt sets Lt field to given value.

### HasLt

`func (o *InputQueryNamedOperation) HasLt() bool`

HasLt returns a boolean if a field has been set.

### GetLte

`func (o *InputQueryNamedOperation) GetLte() FieldValue`

GetLte returns the Lte field if non-nil, zero value otherwise.

### GetLteOk

`func (o *InputQueryNamedOperation) GetLteOk() (*FieldValue, bool)`

GetLteOk returns a tuple with the Lte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLte

`func (o *InputQueryNamedOperation) SetLte(v FieldValue)`

SetLte sets Lte field to given value.

### HasLte

`func (o *InputQueryNamedOperation) HasLte() bool`

HasLte returns a boolean if a field has been set.

### GetNe

`func (o *InputQueryNamedOperation) GetNe() FieldValue`

GetNe returns the Ne field if non-nil, zero value otherwise.

### GetNeOk

`func (o *InputQueryNamedOperation) GetNeOk() (*FieldValue, bool)`

GetNeOk returns a tuple with the Ne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNe

`func (o *InputQueryNamedOperation) SetNe(v FieldValue)`

SetNe sets Ne field to given value.

### HasNe

`func (o *InputQueryNamedOperation) HasNe() bool`

HasNe returns a boolean if a field has been set.

### GetIn

`func (o *InputQueryNamedOperation) GetIn() FieldValue`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *InputQueryNamedOperation) GetInOk() (*FieldValue, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *InputQueryNamedOperation) SetIn(v FieldValue)`

SetIn sets In field to given value.

### HasIn

`func (o *InputQueryNamedOperation) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetBetween

`func (o *InputQueryNamedOperation) GetBetween() FieldValue`

GetBetween returns the Between field if non-nil, zero value otherwise.

### GetBetweenOk

`func (o *InputQueryNamedOperation) GetBetweenOk() (*FieldValue, bool)`

GetBetweenOk returns a tuple with the Between field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetween

`func (o *InputQueryNamedOperation) SetBetween(v FieldValue)`

SetBetween sets Between field to given value.

### HasBetween

`func (o *InputQueryNamedOperation) HasBetween() bool`

HasBetween returns a boolean if a field has been set.

### GetContains

`func (o *InputQueryNamedOperation) GetContains() FieldValue`

GetContains returns the Contains field if non-nil, zero value otherwise.

### GetContainsOk

`func (o *InputQueryNamedOperation) GetContainsOk() (*FieldValue, bool)`

GetContainsOk returns a tuple with the Contains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContains

`func (o *InputQueryNamedOperation) SetContains(v FieldValue)`

SetContains sets Contains field to given value.

### HasContains

`func (o *InputQueryNamedOperation) HasContains() bool`

HasContains returns a boolean if a field has been set.

### GetEndsWith

`func (o *InputQueryNamedOperation) GetEndsWith() FieldValue`

GetEndsWith returns the EndsWith field if non-nil, zero value otherwise.

### GetEndsWithOk

`func (o *InputQueryNamedOperation) GetEndsWithOk() (*FieldValue, bool)`

GetEndsWithOk returns a tuple with the EndsWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsWith

`func (o *InputQueryNamedOperation) SetEndsWith(v FieldValue)`

SetEndsWith sets EndsWith field to given value.

### HasEndsWith

`func (o *InputQueryNamedOperation) HasEndsWith() bool`

HasEndsWith returns a boolean if a field has been set.

### GetStartsWith

`func (o *InputQueryNamedOperation) GetStartsWith() FieldValue`

GetStartsWith returns the StartsWith field if non-nil, zero value otherwise.

### GetStartsWithOk

`func (o *InputQueryNamedOperation) GetStartsWithOk() (*FieldValue, bool)`

GetStartsWithOk returns a tuple with the StartsWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsWith

`func (o *InputQueryNamedOperation) SetStartsWith(v FieldValue)`

SetStartsWith sets StartsWith field to given value.

### HasStartsWith

`func (o *InputQueryNamedOperation) HasStartsWith() bool`

HasStartsWith returns a boolean if a field has been set.

### GetMatch

`func (o *InputQueryNamedOperation) GetMatch() FieldValue`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *InputQueryNamedOperation) GetMatchOk() (*FieldValue, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *InputQueryNamedOperation) SetMatch(v FieldValue)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *InputQueryNamedOperation) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### GetUnderscoreId

`func (o *InputQueryNamedOperation) GetUnderscoreId() string`

GetUnderscoreId returns the UnderscoreId field if non-nil, zero value otherwise.

### GetUnderscoreIdOk

`func (o *InputQueryNamedOperation) GetUnderscoreIdOk() (*string, bool)`

GetUnderscoreIdOk returns a tuple with the UnderscoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreId

`func (o *InputQueryNamedOperation) SetUnderscoreId(v string)`

SetUnderscoreId sets UnderscoreId field to given value.

### HasUnderscoreId

`func (o *InputQueryNamedOperation) HasUnderscoreId() bool`

HasUnderscoreId returns a boolean if a field has been set.

### GetName

`func (o *InputQueryNamedOperation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputQueryNamedOperation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputQueryNamedOperation) SetName(v string)`

SetName sets Name field to given value.


### GetExtraData

`func (o *InputQueryNamedOperation) GetExtraData() []string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *InputQueryNamedOperation) GetExtraDataOk() (*[]string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *InputQueryNamedOperation) SetExtraData(v []string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *InputQueryNamedOperation) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### GetFrom

`func (o *InputQueryNamedOperation) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *InputQueryNamedOperation) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *InputQueryNamedOperation) SetFrom(v int32)`

SetFrom sets From field to given value.


### GetTo

`func (o *InputQueryNamedOperation) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *InputQueryNamedOperation) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *InputQueryNamedOperation) SetTo(v int32)`

SetTo sets To field to given value.


### GetFields

`func (o *InputQueryNamedOperation) GetFields() []map[string]interface{}`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *InputQueryNamedOperation) GetFieldsOk() (*[]map[string]interface{}, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *InputQueryNamedOperation) SetFields(v []map[string]interface{})`

SetFields sets Fields field to given value.

### HasFields

`func (o *InputQueryNamedOperation) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
