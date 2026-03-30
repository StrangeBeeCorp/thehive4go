# OutputJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnderscoreId** | **string** |  | 
**UnderscoreType** | **string** |  | 
**UnderscoreCreatedBy** | **string** |  | 
**UnderscoreUpdatedBy** | Pointer to **string** |  | [optional] 
**UnderscoreCreatedAt** | **int64** |  | 
**UnderscoreUpdatedAt** | Pointer to **int64** |  | [optional] 
**AnalyzerId** | **string** |  | 
**AnalyzerName** | **string** |  | 
**AnalyzerDefinition** | **string** |  | 
**Status** | **string** |  | 
**StartDate** | **int64** |  | 
**EndDate** | Pointer to **int64** |  | [optional] 
**Report** | Pointer to **map[string]interface{}** |  | [optional] 
**CortexId** | **string** |  | 
**CortexJobId** | **string** |  | 
**Id** | **string** |  | 
**CaseArtifact** | Pointer to **map[string]interface{}** |  | [optional] 
**Operations** | **string** |  | 

## Methods

### NewOutputJob

`func NewOutputJob(underscoreId string, underscoreType string, underscoreCreatedBy string, underscoreCreatedAt int64, analyzerId string, analyzerName string, analyzerDefinition string, status string, startDate int64, cortexId string, cortexJobId string, id string, operations string, ) *OutputJob`

NewOutputJob instantiates a new OutputJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputJobWithDefaults

`func NewOutputJobWithDefaults() *OutputJob`

NewOutputJobWithDefaults instantiates a new OutputJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnderscoreId

`func (o *OutputJob) GetUnderscoreId() string`

GetUnderscoreId returns the UnderscoreId field if non-nil, zero value otherwise.

### GetUnderscoreIdOk

`func (o *OutputJob) GetUnderscoreIdOk() (*string, bool)`

GetUnderscoreIdOk returns a tuple with the UnderscoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreId

`func (o *OutputJob) SetUnderscoreId(v string)`

SetUnderscoreId sets UnderscoreId field to given value.


### GetUnderscoreType

`func (o *OutputJob) GetUnderscoreType() string`

GetUnderscoreType returns the UnderscoreType field if non-nil, zero value otherwise.

### GetUnderscoreTypeOk

`func (o *OutputJob) GetUnderscoreTypeOk() (*string, bool)`

GetUnderscoreTypeOk returns a tuple with the UnderscoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreType

`func (o *OutputJob) SetUnderscoreType(v string)`

SetUnderscoreType sets UnderscoreType field to given value.


### GetUnderscoreCreatedBy

`func (o *OutputJob) GetUnderscoreCreatedBy() string`

GetUnderscoreCreatedBy returns the UnderscoreCreatedBy field if non-nil, zero value otherwise.

### GetUnderscoreCreatedByOk

`func (o *OutputJob) GetUnderscoreCreatedByOk() (*string, bool)`

GetUnderscoreCreatedByOk returns a tuple with the UnderscoreCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedBy

`func (o *OutputJob) SetUnderscoreCreatedBy(v string)`

SetUnderscoreCreatedBy sets UnderscoreCreatedBy field to given value.


### GetUnderscoreUpdatedBy

`func (o *OutputJob) GetUnderscoreUpdatedBy() string`

GetUnderscoreUpdatedBy returns the UnderscoreUpdatedBy field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedByOk

`func (o *OutputJob) GetUnderscoreUpdatedByOk() (*string, bool)`

GetUnderscoreUpdatedByOk returns a tuple with the UnderscoreUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedBy

`func (o *OutputJob) SetUnderscoreUpdatedBy(v string)`

SetUnderscoreUpdatedBy sets UnderscoreUpdatedBy field to given value.

### HasUnderscoreUpdatedBy

`func (o *OutputJob) HasUnderscoreUpdatedBy() bool`

HasUnderscoreUpdatedBy returns a boolean if a field has been set.

### GetUnderscoreCreatedAt

`func (o *OutputJob) GetUnderscoreCreatedAt() int64`

GetUnderscoreCreatedAt returns the UnderscoreCreatedAt field if non-nil, zero value otherwise.

### GetUnderscoreCreatedAtOk

`func (o *OutputJob) GetUnderscoreCreatedAtOk() (*int64, bool)`

GetUnderscoreCreatedAtOk returns a tuple with the UnderscoreCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedAt

`func (o *OutputJob) SetUnderscoreCreatedAt(v int64)`

SetUnderscoreCreatedAt sets UnderscoreCreatedAt field to given value.


### GetUnderscoreUpdatedAt

`func (o *OutputJob) GetUnderscoreUpdatedAt() int64`

GetUnderscoreUpdatedAt returns the UnderscoreUpdatedAt field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedAtOk

`func (o *OutputJob) GetUnderscoreUpdatedAtOk() (*int64, bool)`

GetUnderscoreUpdatedAtOk returns a tuple with the UnderscoreUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedAt

`func (o *OutputJob) SetUnderscoreUpdatedAt(v int64)`

SetUnderscoreUpdatedAt sets UnderscoreUpdatedAt field to given value.

### HasUnderscoreUpdatedAt

`func (o *OutputJob) HasUnderscoreUpdatedAt() bool`

HasUnderscoreUpdatedAt returns a boolean if a field has been set.

### GetAnalyzerId

`func (o *OutputJob) GetAnalyzerId() string`

GetAnalyzerId returns the AnalyzerId field if non-nil, zero value otherwise.

### GetAnalyzerIdOk

`func (o *OutputJob) GetAnalyzerIdOk() (*string, bool)`

GetAnalyzerIdOk returns a tuple with the AnalyzerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzerId

`func (o *OutputJob) SetAnalyzerId(v string)`

SetAnalyzerId sets AnalyzerId field to given value.


### GetAnalyzerName

`func (o *OutputJob) GetAnalyzerName() string`

GetAnalyzerName returns the AnalyzerName field if non-nil, zero value otherwise.

### GetAnalyzerNameOk

`func (o *OutputJob) GetAnalyzerNameOk() (*string, bool)`

GetAnalyzerNameOk returns a tuple with the AnalyzerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzerName

`func (o *OutputJob) SetAnalyzerName(v string)`

SetAnalyzerName sets AnalyzerName field to given value.


### GetAnalyzerDefinition

`func (o *OutputJob) GetAnalyzerDefinition() string`

GetAnalyzerDefinition returns the AnalyzerDefinition field if non-nil, zero value otherwise.

### GetAnalyzerDefinitionOk

`func (o *OutputJob) GetAnalyzerDefinitionOk() (*string, bool)`

GetAnalyzerDefinitionOk returns a tuple with the AnalyzerDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzerDefinition

`func (o *OutputJob) SetAnalyzerDefinition(v string)`

SetAnalyzerDefinition sets AnalyzerDefinition field to given value.


### GetStatus

`func (o *OutputJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OutputJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OutputJob) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStartDate

`func (o *OutputJob) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *OutputJob) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *OutputJob) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *OutputJob) GetEndDate() int64`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *OutputJob) GetEndDateOk() (*int64, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *OutputJob) SetEndDate(v int64)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *OutputJob) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetReport

`func (o *OutputJob) GetReport() map[string]interface{}`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *OutputJob) GetReportOk() (*map[string]interface{}, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *OutputJob) SetReport(v map[string]interface{})`

SetReport sets Report field to given value.

### HasReport

`func (o *OutputJob) HasReport() bool`

HasReport returns a boolean if a field has been set.

### GetCortexId

`func (o *OutputJob) GetCortexId() string`

GetCortexId returns the CortexId field if non-nil, zero value otherwise.

### GetCortexIdOk

`func (o *OutputJob) GetCortexIdOk() (*string, bool)`

GetCortexIdOk returns a tuple with the CortexId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCortexId

`func (o *OutputJob) SetCortexId(v string)`

SetCortexId sets CortexId field to given value.


### GetCortexJobId

`func (o *OutputJob) GetCortexJobId() string`

GetCortexJobId returns the CortexJobId field if non-nil, zero value otherwise.

### GetCortexJobIdOk

`func (o *OutputJob) GetCortexJobIdOk() (*string, bool)`

GetCortexJobIdOk returns a tuple with the CortexJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCortexJobId

`func (o *OutputJob) SetCortexJobId(v string)`

SetCortexJobId sets CortexJobId field to given value.


### GetId

`func (o *OutputJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutputJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OutputJob) SetId(v string)`

SetId sets Id field to given value.


### GetCaseArtifact

`func (o *OutputJob) GetCaseArtifact() map[string]interface{}`

GetCaseArtifact returns the CaseArtifact field if non-nil, zero value otherwise.

### GetCaseArtifactOk

`func (o *OutputJob) GetCaseArtifactOk() (*map[string]interface{}, bool)`

GetCaseArtifactOk returns a tuple with the CaseArtifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseArtifact

`func (o *OutputJob) SetCaseArtifact(v map[string]interface{})`

SetCaseArtifact sets CaseArtifact field to given value.

### HasCaseArtifact

`func (o *OutputJob) HasCaseArtifact() bool`

HasCaseArtifact returns a boolean if a field has been set.

### GetOperations

`func (o *OutputJob) GetOperations() string`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *OutputJob) GetOperationsOk() (*string, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *OutputJob) SetOperations(v string)`

SetOperations sets Operations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


