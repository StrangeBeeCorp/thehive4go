# InputJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalyzerId** | **string** |  | 
**CortexId** | **string** |  | 
**ArtifactId** | **string** | &#x3D;&#x3D; observableId field | 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewInputJob

`func NewInputJob(analyzerId string, cortexId string, artifactId string, ) *InputJob`

NewInputJob instantiates a new InputJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputJobWithDefaults

`func NewInputJobWithDefaults() *InputJob`

NewInputJobWithDefaults instantiates a new InputJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyzerId

`func (o *InputJob) GetAnalyzerId() string`

GetAnalyzerId returns the AnalyzerId field if non-nil, zero value otherwise.

### GetAnalyzerIdOk

`func (o *InputJob) GetAnalyzerIdOk() (*string, bool)`

GetAnalyzerIdOk returns a tuple with the AnalyzerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzerId

`func (o *InputJob) SetAnalyzerId(v string)`

SetAnalyzerId sets AnalyzerId field to given value.


### GetCortexId

`func (o *InputJob) GetCortexId() string`

GetCortexId returns the CortexId field if non-nil, zero value otherwise.

### GetCortexIdOk

`func (o *InputJob) GetCortexIdOk() (*string, bool)`

GetCortexIdOk returns a tuple with the CortexId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCortexId

`func (o *InputJob) SetCortexId(v string)`

SetCortexId sets CortexId field to given value.


### GetArtifactId

`func (o *InputJob) GetArtifactId() string`

GetArtifactId returns the ArtifactId field if non-nil, zero value otherwise.

### GetArtifactIdOk

`func (o *InputJob) GetArtifactIdOk() (*string, bool)`

GetArtifactIdOk returns a tuple with the ArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactId

`func (o *InputJob) SetArtifactId(v string)`

SetArtifactId sets ArtifactId field to given value.


### GetParameters

`func (o *InputJob) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *InputJob) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *InputJob) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *InputJob) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


