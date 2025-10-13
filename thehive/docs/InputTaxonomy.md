# InputTaxonomy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | **string** |  | 
**Description** | **string** |  | 
**Version** | **int32** |  | 
**Exclusive** | Pointer to **bool** |  | [optional] 
**Predicates** | Pointer to [**[]InputPredicate**](InputPredicate.md) |  | [optional] 
**Values** | Pointer to [**[]InputValue**](InputValue.md) |  | [optional] 

## Methods

### NewInputTaxonomy

`func NewInputTaxonomy(namespace string, description string, version int32, ) *InputTaxonomy`

NewInputTaxonomy instantiates a new InputTaxonomy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputTaxonomyWithDefaults

`func NewInputTaxonomyWithDefaults() *InputTaxonomy`

NewInputTaxonomyWithDefaults instantiates a new InputTaxonomy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *InputTaxonomy) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *InputTaxonomy) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *InputTaxonomy) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetDescription

`func (o *InputTaxonomy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputTaxonomy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputTaxonomy) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetVersion

`func (o *InputTaxonomy) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InputTaxonomy) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InputTaxonomy) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetExclusive

`func (o *InputTaxonomy) GetExclusive() bool`

GetExclusive returns the Exclusive field if non-nil, zero value otherwise.

### GetExclusiveOk

`func (o *InputTaxonomy) GetExclusiveOk() (*bool, bool)`

GetExclusiveOk returns a tuple with the Exclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusive

`func (o *InputTaxonomy) SetExclusive(v bool)`

SetExclusive sets Exclusive field to given value.

### HasExclusive

`func (o *InputTaxonomy) HasExclusive() bool`

HasExclusive returns a boolean if a field has been set.

### GetPredicates

`func (o *InputTaxonomy) GetPredicates() []InputPredicate`

GetPredicates returns the Predicates field if non-nil, zero value otherwise.

### GetPredicatesOk

`func (o *InputTaxonomy) GetPredicatesOk() (*[]InputPredicate, bool)`

GetPredicatesOk returns a tuple with the Predicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicates

`func (o *InputTaxonomy) SetPredicates(v []InputPredicate)`

SetPredicates sets Predicates field to given value.

### HasPredicates

`func (o *InputTaxonomy) HasPredicates() bool`

HasPredicates returns a boolean if a field has been set.

### GetValues

`func (o *InputTaxonomy) GetValues() []InputValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *InputTaxonomy) GetValuesOk() (*[]InputValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *InputTaxonomy) SetValues(v []InputValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *InputTaxonomy) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


