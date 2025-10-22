# UrlPropertyDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Cardinality** | [**Cardinality**](Cardinality.md) |  | 
**Aggregable** | **bool** |  | 
**IndexType** | [**AnyIndexType**](AnyIndexType.md) |  | 
**Type** | **string** |  | 

## Methods

### NewUrlPropertyDescription

`func NewUrlPropertyDescription(name string, cardinality Cardinality, aggregable bool, indexType AnyIndexType, type_ string, ) *UrlPropertyDescription`

NewUrlPropertyDescription instantiates a new UrlPropertyDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUrlPropertyDescriptionWithDefaults

`func NewUrlPropertyDescriptionWithDefaults() *UrlPropertyDescription`

NewUrlPropertyDescriptionWithDefaults instantiates a new UrlPropertyDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UrlPropertyDescription) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UrlPropertyDescription) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UrlPropertyDescription) SetName(v string)`

SetName sets Name field to given value.


### GetCardinality

`func (o *UrlPropertyDescription) GetCardinality() Cardinality`

GetCardinality returns the Cardinality field if non-nil, zero value otherwise.

### GetCardinalityOk

`func (o *UrlPropertyDescription) GetCardinalityOk() (*Cardinality, bool)`

GetCardinalityOk returns a tuple with the Cardinality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardinality

`func (o *UrlPropertyDescription) SetCardinality(v Cardinality)`

SetCardinality sets Cardinality field to given value.


### GetAggregable

`func (o *UrlPropertyDescription) GetAggregable() bool`

GetAggregable returns the Aggregable field if non-nil, zero value otherwise.

### GetAggregableOk

`func (o *UrlPropertyDescription) GetAggregableOk() (*bool, bool)`

GetAggregableOk returns a tuple with the Aggregable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregable

`func (o *UrlPropertyDescription) SetAggregable(v bool)`

SetAggregable sets Aggregable field to given value.


### GetIndexType

`func (o *UrlPropertyDescription) GetIndexType() AnyIndexType`

GetIndexType returns the IndexType field if non-nil, zero value otherwise.

### GetIndexTypeOk

`func (o *UrlPropertyDescription) GetIndexTypeOk() (*AnyIndexType, bool)`

GetIndexTypeOk returns a tuple with the IndexType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexType

`func (o *UrlPropertyDescription) SetIndexType(v AnyIndexType)`

SetIndexType sets IndexType field to given value.


### GetType

`func (o *UrlPropertyDescription) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UrlPropertyDescription) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UrlPropertyDescription) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


