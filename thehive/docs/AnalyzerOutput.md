# AnalyzerOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the analyzer once enabled within the organization. |
**Name** | **string** | Display name of the analyzer. |
**Version** | **string** | Version of the analyzer. |
**Description** | **string** | Description of what the analyzer does. |
**DataTypeList** | Pointer to **[]string** | Observable data types the analyzer supports. | [optional]
**CortexIds** | Pointer to **[]string** | Names of the Cortex servers on which the analyzer is available. | [optional]

## Methods

### NewAnalyzerOutput

`func NewAnalyzerOutput(id string, name string, version string, description string, ) *AnalyzerOutput`

NewAnalyzerOutput instantiates a new AnalyzerOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyzerOutputWithDefaults

`func NewAnalyzerOutputWithDefaults() *AnalyzerOutput`

NewAnalyzerOutputWithDefaults instantiates a new AnalyzerOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AnalyzerOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AnalyzerOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AnalyzerOutput) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AnalyzerOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AnalyzerOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AnalyzerOutput) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *AnalyzerOutput) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AnalyzerOutput) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AnalyzerOutput) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetDescription

`func (o *AnalyzerOutput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AnalyzerOutput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AnalyzerOutput) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDataTypeList

`func (o *AnalyzerOutput) GetDataTypeList() []string`

GetDataTypeList returns the DataTypeList field if non-nil, zero value otherwise.

### GetDataTypeListOk

`func (o *AnalyzerOutput) GetDataTypeListOk() (*[]string, bool)`

GetDataTypeListOk returns a tuple with the DataTypeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypeList

`func (o *AnalyzerOutput) SetDataTypeList(v []string)`

SetDataTypeList sets DataTypeList field to given value.

### HasDataTypeList

`func (o *AnalyzerOutput) HasDataTypeList() bool`

HasDataTypeList returns a boolean if a field has been set.

### GetCortexIds

`func (o *AnalyzerOutput) GetCortexIds() []string`

GetCortexIds returns the CortexIds field if non-nil, zero value otherwise.

### GetCortexIdsOk

`func (o *AnalyzerOutput) GetCortexIdsOk() (*[]string, bool)`

GetCortexIdsOk returns a tuple with the CortexIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCortexIds

`func (o *AnalyzerOutput) SetCortexIds(v []string)`

SetCortexIds sets CortexIds field to given value.

### HasCortexIds

`func (o *AnalyzerOutput) HasCortexIds() bool`

HasCortexIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
