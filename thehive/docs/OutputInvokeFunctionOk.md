# OutputInvokeFunctionOk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**DurationMillis** | **int64** | Time taken by the function | 
**Stdout** | **string** | Content of stdout | 
**Stderr** | **string** | Content of stderr | 

## Methods

### NewOutputInvokeFunctionOk

`func NewOutputInvokeFunctionOk(result interface{}, durationMillis int64, stdout string, stderr string, ) *OutputInvokeFunctionOk`

NewOutputInvokeFunctionOk instantiates a new OutputInvokeFunctionOk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputInvokeFunctionOkWithDefaults

`func NewOutputInvokeFunctionOkWithDefaults() *OutputInvokeFunctionOk`

NewOutputInvokeFunctionOkWithDefaults instantiates a new OutputInvokeFunctionOk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *OutputInvokeFunctionOk) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OutputInvokeFunctionOk) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OutputInvokeFunctionOk) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *OutputInvokeFunctionOk) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *OutputInvokeFunctionOk) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetDurationMillis

`func (o *OutputInvokeFunctionOk) GetDurationMillis() int64`

GetDurationMillis returns the DurationMillis field if non-nil, zero value otherwise.

### GetDurationMillisOk

`func (o *OutputInvokeFunctionOk) GetDurationMillisOk() (*int64, bool)`

GetDurationMillisOk returns a tuple with the DurationMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMillis

`func (o *OutputInvokeFunctionOk) SetDurationMillis(v int64)`

SetDurationMillis sets DurationMillis field to given value.


### GetStdout

`func (o *OutputInvokeFunctionOk) GetStdout() string`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *OutputInvokeFunctionOk) GetStdoutOk() (*string, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *OutputInvokeFunctionOk) SetStdout(v string)`

SetStdout sets Stdout field to given value.


### GetStderr

`func (o *OutputInvokeFunctionOk) GetStderr() string`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *OutputInvokeFunctionOk) GetStderrOk() (*string, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *OutputInvokeFunctionOk) SetStderr(v string)`

SetStderr sets Stderr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


