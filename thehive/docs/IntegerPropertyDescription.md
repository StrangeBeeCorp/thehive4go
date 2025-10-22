# IntegerPropertyDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Cardinality** | [**Cardinality**](Cardinality.md) |  | 
**Aggregable** | **bool** |  | 
**IndexType** | [**BasicIndexType**](BasicIndexType.md) |  | 
**Type** | **string** |  | 

## Methods

### NewIntegerPropertyDescription

`func NewIntegerPropertyDescription(name string, cardinality Cardinality, aggregable bool, indexType BasicIndexType, type_ string, ) *IntegerPropertyDescription`

NewIntegerPropertyDescription instantiates a new IntegerPropertyDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegerPropertyDescriptionWithDefaults

`func NewIntegerPropertyDescriptionWithDefaults() *IntegerPropertyDescription`

NewIntegerPropertyDescriptionWithDefaults instantiates a new IntegerPropertyDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IntegerPropertyDescription) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegerPropertyDescription) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegerPropertyDescription) SetName(v string)`

SetName sets Name field to given value.


### GetCardinality

`func (o *IntegerPropertyDescription) GetCardinality() Cardinality`

GetCardinality returns the Cardinality field if non-nil, zero value otherwise.

### GetCardinalityOk

`func (o *IntegerPropertyDescription) GetCardinalityOk() (*Cardinality, bool)`

GetCardinalityOk returns a tuple with the Cardinality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardinality

`func (o *IntegerPropertyDescription) SetCardinality(v Cardinality)`

SetCardinality sets Cardinality field to given value.


### GetAggregable

`func (o *IntegerPropertyDescription) GetAggregable() bool`

GetAggregable returns the Aggregable field if non-nil, zero value otherwise.

### GetAggregableOk

`func (o *IntegerPropertyDescription) GetAggregableOk() (*bool, bool)`

GetAggregableOk returns a tuple with the Aggregable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregable

`func (o *IntegerPropertyDescription) SetAggregable(v bool)`

SetAggregable sets Aggregable field to given value.


### GetIndexType

`func (o *IntegerPropertyDescription) GetIndexType() BasicIndexType`

GetIndexType returns the IndexType field if non-nil, zero value otherwise.

### GetIndexTypeOk

`func (o *IntegerPropertyDescription) GetIndexTypeOk() (*BasicIndexType, bool)`

GetIndexTypeOk returns a tuple with the IndexType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexType

`func (o *IntegerPropertyDescription) SetIndexType(v BasicIndexType)`

SetIndexType sets IndexType field to given value.


### GetType

`func (o *IntegerPropertyDescription) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegerPropertyDescription) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegerPropertyDescription) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


