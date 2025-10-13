# EnumerationPropertyDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Cardinality** | [**Cardinality**](Cardinality.md) |  | 
**Values** | [**EnumerationPropertyDescriptionValues**](EnumerationPropertyDescriptionValues.md) |  | 
**Labels** | **[]string** |  | 
**Aggregable** | **bool** |  | 
**IndexType** | [**BasicIndexType**](BasicIndexType.md) |  | 
**Type** | **string** |  | 

## Methods

### NewEnumerationPropertyDescription

`func NewEnumerationPropertyDescription(name string, cardinality Cardinality, values EnumerationPropertyDescriptionValues, labels []string, aggregable bool, indexType BasicIndexType, type_ string, ) *EnumerationPropertyDescription`

NewEnumerationPropertyDescription instantiates a new EnumerationPropertyDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnumerationPropertyDescriptionWithDefaults

`func NewEnumerationPropertyDescriptionWithDefaults() *EnumerationPropertyDescription`

NewEnumerationPropertyDescriptionWithDefaults instantiates a new EnumerationPropertyDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnumerationPropertyDescription) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnumerationPropertyDescription) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnumerationPropertyDescription) SetName(v string)`

SetName sets Name field to given value.


### GetCardinality

`func (o *EnumerationPropertyDescription) GetCardinality() Cardinality`

GetCardinality returns the Cardinality field if non-nil, zero value otherwise.

### GetCardinalityOk

`func (o *EnumerationPropertyDescription) GetCardinalityOk() (*Cardinality, bool)`

GetCardinalityOk returns a tuple with the Cardinality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardinality

`func (o *EnumerationPropertyDescription) SetCardinality(v Cardinality)`

SetCardinality sets Cardinality field to given value.


### GetValues

`func (o *EnumerationPropertyDescription) GetValues() EnumerationPropertyDescriptionValues`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *EnumerationPropertyDescription) GetValuesOk() (*EnumerationPropertyDescriptionValues, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *EnumerationPropertyDescription) SetValues(v EnumerationPropertyDescriptionValues)`

SetValues sets Values field to given value.


### GetLabels

`func (o *EnumerationPropertyDescription) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *EnumerationPropertyDescription) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *EnumerationPropertyDescription) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetAggregable

`func (o *EnumerationPropertyDescription) GetAggregable() bool`

GetAggregable returns the Aggregable field if non-nil, zero value otherwise.

### GetAggregableOk

`func (o *EnumerationPropertyDescription) GetAggregableOk() (*bool, bool)`

GetAggregableOk returns a tuple with the Aggregable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregable

`func (o *EnumerationPropertyDescription) SetAggregable(v bool)`

SetAggregable sets Aggregable field to given value.


### GetIndexType

`func (o *EnumerationPropertyDescription) GetIndexType() BasicIndexType`

GetIndexType returns the IndexType field if non-nil, zero value otherwise.

### GetIndexTypeOk

`func (o *EnumerationPropertyDescription) GetIndexTypeOk() (*BasicIndexType, bool)`

GetIndexTypeOk returns a tuple with the IndexType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexType

`func (o *EnumerationPropertyDescription) SetIndexType(v BasicIndexType)`

SetIndexType sets IndexType field to given value.


### GetType

`func (o *EnumerationPropertyDescription) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnumerationPropertyDescription) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnumerationPropertyDescription) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


