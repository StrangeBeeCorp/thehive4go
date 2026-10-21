# InputProcedure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PatternId** | **string** | Technique ID from a loaded catalog, for example a MITRE ATT&amp;CK technique ID. Must match an existing technique. Use &#x60;POST /api/v1/query&#x60; with &#x60;listPattern&#x60; to retrieve available techniques and their IDs. |
**OccurDate** | **int64** | Date and time when the technique occurred during the incident. Expected format: Unix timestamp in milliseconds. |
**Tactic** | Pointer to **string** | Short name of the tactic the technique belongs to. Required when the technique belongs to more than one tactic. Use &#x60;POST /api/v1/query&#x60; with &#x60;listPattern&#x60; to check which tactics a technique belongs to. | [optional]
**Description** | Pointer to **string** | Description of how the technique was carried out in this incident, referred to as the procedure. Supports [TheHive-flavored Markdown](https://docs.strangebee.com/thehive/user-guides/thehive-flavored-markdown/). | [optional]

## Methods

### NewInputProcedure

`func NewInputProcedure(patternId string, occurDate int64, ) *InputProcedure`

NewInputProcedure instantiates a new InputProcedure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputProcedureWithDefaults

`func NewInputProcedureWithDefaults() *InputProcedure`

NewInputProcedureWithDefaults instantiates a new InputProcedure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPatternId

`func (o *InputProcedure) GetPatternId() string`

GetPatternId returns the PatternId field if non-nil, zero value otherwise.

### GetPatternIdOk

`func (o *InputProcedure) GetPatternIdOk() (*string, bool)`

GetPatternIdOk returns a tuple with the PatternId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternId

`func (o *InputProcedure) SetPatternId(v string)`

SetPatternId sets PatternId field to given value.


### GetOccurDate

`func (o *InputProcedure) GetOccurDate() int64`

GetOccurDate returns the OccurDate field if non-nil, zero value otherwise.

### GetOccurDateOk

`func (o *InputProcedure) GetOccurDateOk() (*int64, bool)`

GetOccurDateOk returns a tuple with the OccurDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurDate

`func (o *InputProcedure) SetOccurDate(v int64)`

SetOccurDate sets OccurDate field to given value.


### GetTactic

`func (o *InputProcedure) GetTactic() string`

GetTactic returns the Tactic field if non-nil, zero value otherwise.

### GetTacticOk

`func (o *InputProcedure) GetTacticOk() (*string, bool)`

GetTacticOk returns a tuple with the Tactic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTactic

`func (o *InputProcedure) SetTactic(v string)`

SetTactic sets Tactic field to given value.

### HasTactic

`func (o *InputProcedure) HasTactic() bool`

HasTactic returns a boolean if a field has been set.

### GetDescription

`func (o *InputProcedure) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputProcedure) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputProcedure) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InputProcedure) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
