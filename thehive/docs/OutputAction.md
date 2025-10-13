# OutputAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnderscoreId** | **string** |  | 
**UnderscoreType** | **string** |  | 
**UnderscoreCreatedBy** | **string** |  | 
**UnderscoreUpdatedBy** | Pointer to **string** |  | [optional] 
**UnderscoreCreatedAt** | **time.Time** |  | 
**UnderscoreUpdatedAt** | Pointer to **time.Time** |  | [optional] 
**ResponderId** | **string** |  | 
**ResponderName** | Pointer to **string** |  | [optional] 
**ResponderDefinition** | Pointer to **string** |  | [optional] 
**CortexId** | Pointer to **string** |  | [optional] 
**CortexJobId** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | 
**ObjectId** | **string** |  | 
**Status** | **string** |  | 
**StartDate** | **time.Time** |  | 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Operations** | **string** |  | 
**Report** | **string** |  | 

## Methods

### NewOutputAction

`func NewOutputAction(underscoreId string, underscoreType string, underscoreCreatedBy string, underscoreCreatedAt time.Time, responderId string, objectType string, objectId string, status string, startDate time.Time, operations string, report string, ) *OutputAction`

NewOutputAction instantiates a new OutputAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputActionWithDefaults

`func NewOutputActionWithDefaults() *OutputAction`

NewOutputActionWithDefaults instantiates a new OutputAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnderscoreId

`func (o *OutputAction) GetUnderscoreId() string`

GetUnderscoreId returns the UnderscoreId field if non-nil, zero value otherwise.

### GetUnderscoreIdOk

`func (o *OutputAction) GetUnderscoreIdOk() (*string, bool)`

GetUnderscoreIdOk returns a tuple with the UnderscoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreId

`func (o *OutputAction) SetUnderscoreId(v string)`

SetUnderscoreId sets UnderscoreId field to given value.


### GetUnderscoreType

`func (o *OutputAction) GetUnderscoreType() string`

GetUnderscoreType returns the UnderscoreType field if non-nil, zero value otherwise.

### GetUnderscoreTypeOk

`func (o *OutputAction) GetUnderscoreTypeOk() (*string, bool)`

GetUnderscoreTypeOk returns a tuple with the UnderscoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreType

`func (o *OutputAction) SetUnderscoreType(v string)`

SetUnderscoreType sets UnderscoreType field to given value.


### GetUnderscoreCreatedBy

`func (o *OutputAction) GetUnderscoreCreatedBy() string`

GetUnderscoreCreatedBy returns the UnderscoreCreatedBy field if non-nil, zero value otherwise.

### GetUnderscoreCreatedByOk

`func (o *OutputAction) GetUnderscoreCreatedByOk() (*string, bool)`

GetUnderscoreCreatedByOk returns a tuple with the UnderscoreCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedBy

`func (o *OutputAction) SetUnderscoreCreatedBy(v string)`

SetUnderscoreCreatedBy sets UnderscoreCreatedBy field to given value.


### GetUnderscoreUpdatedBy

`func (o *OutputAction) GetUnderscoreUpdatedBy() string`

GetUnderscoreUpdatedBy returns the UnderscoreUpdatedBy field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedByOk

`func (o *OutputAction) GetUnderscoreUpdatedByOk() (*string, bool)`

GetUnderscoreUpdatedByOk returns a tuple with the UnderscoreUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedBy

`func (o *OutputAction) SetUnderscoreUpdatedBy(v string)`

SetUnderscoreUpdatedBy sets UnderscoreUpdatedBy field to given value.

### HasUnderscoreUpdatedBy

`func (o *OutputAction) HasUnderscoreUpdatedBy() bool`

HasUnderscoreUpdatedBy returns a boolean if a field has been set.

### GetUnderscoreCreatedAt

`func (o *OutputAction) GetUnderscoreCreatedAt() time.Time`

GetUnderscoreCreatedAt returns the UnderscoreCreatedAt field if non-nil, zero value otherwise.

### GetUnderscoreCreatedAtOk

`func (o *OutputAction) GetUnderscoreCreatedAtOk() (*time.Time, bool)`

GetUnderscoreCreatedAtOk returns a tuple with the UnderscoreCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedAt

`func (o *OutputAction) SetUnderscoreCreatedAt(v time.Time)`

SetUnderscoreCreatedAt sets UnderscoreCreatedAt field to given value.


### GetUnderscoreUpdatedAt

`func (o *OutputAction) GetUnderscoreUpdatedAt() time.Time`

GetUnderscoreUpdatedAt returns the UnderscoreUpdatedAt field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedAtOk

`func (o *OutputAction) GetUnderscoreUpdatedAtOk() (*time.Time, bool)`

GetUnderscoreUpdatedAtOk returns a tuple with the UnderscoreUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedAt

`func (o *OutputAction) SetUnderscoreUpdatedAt(v time.Time)`

SetUnderscoreUpdatedAt sets UnderscoreUpdatedAt field to given value.

### HasUnderscoreUpdatedAt

`func (o *OutputAction) HasUnderscoreUpdatedAt() bool`

HasUnderscoreUpdatedAt returns a boolean if a field has been set.

### GetResponderId

`func (o *OutputAction) GetResponderId() string`

GetResponderId returns the ResponderId field if non-nil, zero value otherwise.

### GetResponderIdOk

`func (o *OutputAction) GetResponderIdOk() (*string, bool)`

GetResponderIdOk returns a tuple with the ResponderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponderId

`func (o *OutputAction) SetResponderId(v string)`

SetResponderId sets ResponderId field to given value.


### GetResponderName

`func (o *OutputAction) GetResponderName() string`

GetResponderName returns the ResponderName field if non-nil, zero value otherwise.

### GetResponderNameOk

`func (o *OutputAction) GetResponderNameOk() (*string, bool)`

GetResponderNameOk returns a tuple with the ResponderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponderName

`func (o *OutputAction) SetResponderName(v string)`

SetResponderName sets ResponderName field to given value.

### HasResponderName

`func (o *OutputAction) HasResponderName() bool`

HasResponderName returns a boolean if a field has been set.

### GetResponderDefinition

`func (o *OutputAction) GetResponderDefinition() string`

GetResponderDefinition returns the ResponderDefinition field if non-nil, zero value otherwise.

### GetResponderDefinitionOk

`func (o *OutputAction) GetResponderDefinitionOk() (*string, bool)`

GetResponderDefinitionOk returns a tuple with the ResponderDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponderDefinition

`func (o *OutputAction) SetResponderDefinition(v string)`

SetResponderDefinition sets ResponderDefinition field to given value.

### HasResponderDefinition

`func (o *OutputAction) HasResponderDefinition() bool`

HasResponderDefinition returns a boolean if a field has been set.

### GetCortexId

`func (o *OutputAction) GetCortexId() string`

GetCortexId returns the CortexId field if non-nil, zero value otherwise.

### GetCortexIdOk

`func (o *OutputAction) GetCortexIdOk() (*string, bool)`

GetCortexIdOk returns a tuple with the CortexId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCortexId

`func (o *OutputAction) SetCortexId(v string)`

SetCortexId sets CortexId field to given value.

### HasCortexId

`func (o *OutputAction) HasCortexId() bool`

HasCortexId returns a boolean if a field has been set.

### GetCortexJobId

`func (o *OutputAction) GetCortexJobId() string`

GetCortexJobId returns the CortexJobId field if non-nil, zero value otherwise.

### GetCortexJobIdOk

`func (o *OutputAction) GetCortexJobIdOk() (*string, bool)`

GetCortexJobIdOk returns a tuple with the CortexJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCortexJobId

`func (o *OutputAction) SetCortexJobId(v string)`

SetCortexJobId sets CortexJobId field to given value.

### HasCortexJobId

`func (o *OutputAction) HasCortexJobId() bool`

HasCortexJobId returns a boolean if a field has been set.

### GetObjectType

`func (o *OutputAction) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OutputAction) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OutputAction) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectId

`func (o *OutputAction) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *OutputAction) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *OutputAction) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetStatus

`func (o *OutputAction) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OutputAction) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OutputAction) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStartDate

`func (o *OutputAction) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *OutputAction) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *OutputAction) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *OutputAction) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *OutputAction) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *OutputAction) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *OutputAction) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetOperations

`func (o *OutputAction) GetOperations() string`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *OutputAction) GetOperationsOk() (*string, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *OutputAction) SetOperations(v string)`

SetOperations sets Operations field to given value.


### GetReport

`func (o *OutputAction) GetReport() string`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *OutputAction) GetReportOk() (*string, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *OutputAction) SetReport(v string)`

SetReport sets Report field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


