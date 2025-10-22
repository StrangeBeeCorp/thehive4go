# EntityDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** |  | 
**Path** | **string** |  | 
**InitialQuery** | **string** |  | 
**Attributes** | [**[]PropertyDescription**](PropertyDescription.md) |  | 

## Methods

### NewEntityDescription

`func NewEntityDescription(label string, path string, initialQuery string, attributes []PropertyDescription, ) *EntityDescription`

NewEntityDescription instantiates a new EntityDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityDescriptionWithDefaults

`func NewEntityDescriptionWithDefaults() *EntityDescription`

NewEntityDescriptionWithDefaults instantiates a new EntityDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *EntityDescription) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *EntityDescription) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *EntityDescription) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetPath

`func (o *EntityDescription) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *EntityDescription) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *EntityDescription) SetPath(v string)`

SetPath sets Path field to given value.


### GetInitialQuery

`func (o *EntityDescription) GetInitialQuery() string`

GetInitialQuery returns the InitialQuery field if non-nil, zero value otherwise.

### GetInitialQueryOk

`func (o *EntityDescription) GetInitialQueryOk() (*string, bool)`

GetInitialQueryOk returns a tuple with the InitialQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialQuery

`func (o *EntityDescription) SetInitialQuery(v string)`

SetInitialQuery sets InitialQuery field to given value.


### GetAttributes

`func (o *EntityDescription) GetAttributes() []PropertyDescription`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EntityDescription) GetAttributesOk() (*[]PropertyDescription, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EntityDescription) SetAttributes(v []PropertyDescription)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


