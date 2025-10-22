# OutputStreamAudit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Base** | [**OutputAudit**](OutputAudit.md) |  | 
**Summary** | **map[string]map[string]int32** |  | 

## Methods

### NewOutputStreamAudit

`func NewOutputStreamAudit(base OutputAudit, summary map[string]map[string]int32, ) *OutputStreamAudit`

NewOutputStreamAudit instantiates a new OutputStreamAudit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputStreamAuditWithDefaults

`func NewOutputStreamAuditWithDefaults() *OutputStreamAudit`

NewOutputStreamAuditWithDefaults instantiates a new OutputStreamAudit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBase

`func (o *OutputStreamAudit) GetBase() OutputAudit`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *OutputStreamAudit) GetBaseOk() (*OutputAudit, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *OutputStreamAudit) SetBase(v OutputAudit)`

SetBase sets Base field to given value.


### GetSummary

`func (o *OutputStreamAudit) GetSummary() map[string]map[string]int32`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *OutputStreamAudit) GetSummaryOk() (*map[string]map[string]int32, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *OutputStreamAudit) SetSummary(v map[string]map[string]int32)`

SetSummary sets Summary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


