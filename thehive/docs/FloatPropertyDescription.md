# FloatPropertyDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Cardinality** | [**Cardinality**](Cardinality.md) |  | 
**Aggregable** | **bool** |  | 
**IndexType** | [**BasicIndexType**](BasicIndexType.md) |  | 
**Type** | **string** |  | 

## Methods

### NewFloatPropertyDescription

`func NewFloatPropertyDescription(name string, cardinality Cardinality, aggregable bool, indexType BasicIndexType, type_ string, ) *FloatPropertyDescription`

NewFloatPropertyDescription instantiates a new FloatPropertyDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFloatPropertyDescriptionWithDefaults

`func NewFloatPropertyDescriptionWithDefaults() *FloatPropertyDescription`

NewFloatPropertyDescriptionWithDefaults instantiates a new FloatPropertyDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FloatPropertyDescription) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FloatPropertyDescription) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FloatPropertyDescription) SetName(v string)`

SetName sets Name field to given value.


### GetCardinality

`func (o *FloatPropertyDescription) GetCardinality() Cardinality`

GetCardinality returns the Cardinality field if non-nil, zero value otherwise.

### GetCardinalityOk

`func (o *FloatPropertyDescription) GetCardinalityOk() (*Cardinality, bool)`

GetCardinalityOk returns a tuple with the Cardinality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardinality

`func (o *FloatPropertyDescription) SetCardinality(v Cardinality)`

SetCardinality sets Cardinality field to given value.


### GetAggregable

`func (o *FloatPropertyDescription) GetAggregable() bool`

GetAggregable returns the Aggregable field if non-nil, zero value otherwise.

### GetAggregableOk

`func (o *FloatPropertyDescription) GetAggregableOk() (*bool, bool)`

GetAggregableOk returns a tuple with the Aggregable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregable

`func (o *FloatPropertyDescription) SetAggregable(v bool)`

SetAggregable sets Aggregable field to given value.


### GetIndexType

`func (o *FloatPropertyDescription) GetIndexType() BasicIndexType`

GetIndexType returns the IndexType field if non-nil, zero value otherwise.

### GetIndexTypeOk

`func (o *FloatPropertyDescription) GetIndexTypeOk() (*BasicIndexType, bool)`

GetIndexTypeOk returns a tuple with the IndexType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexType

`func (o *FloatPropertyDescription) SetIndexType(v BasicIndexType)`

SetIndexType sets IndexType field to given value.


### GetType

`func (o *FloatPropertyDescription) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FloatPropertyDescription) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FloatPropertyDescription) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


