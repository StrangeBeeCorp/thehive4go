# OutputInvokeFunction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorType** | **string** |  | 
**ErrorMessage** | **string** |  | 
**Stdout** | **string** | Content of stdout | 
**Stderr** | **string** | Content of stderr | 
**Result** | **interface{}** |  | 
**DurationMillis** | **int64** | Time taken by the function | 

## Methods

### NewOutputInvokeFunction

`func NewOutputInvokeFunction(errorType string, errorMessage string, stdout string, stderr string, result interface{}, durationMillis int64, ) *OutputInvokeFunction`

NewOutputInvokeFunction instantiates a new OutputInvokeFunction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputInvokeFunctionWithDefaults

`func NewOutputInvokeFunctionWithDefaults() *OutputInvokeFunction`

NewOutputInvokeFunctionWithDefaults instantiates a new OutputInvokeFunction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorType

`func (o *OutputInvokeFunction) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *OutputInvokeFunction) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *OutputInvokeFunction) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.


### GetErrorMessage

`func (o *OutputInvokeFunction) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *OutputInvokeFunction) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *OutputInvokeFunction) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### GetStdout

`func (o *OutputInvokeFunction) GetStdout() string`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *OutputInvokeFunction) GetStdoutOk() (*string, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *OutputInvokeFunction) SetStdout(v string)`

SetStdout sets Stdout field to given value.


### GetStderr

`func (o *OutputInvokeFunction) GetStderr() string`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *OutputInvokeFunction) GetStderrOk() (*string, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *OutputInvokeFunction) SetStderr(v string)`

SetStderr sets Stderr field to given value.


### GetResult

`func (o *OutputInvokeFunction) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OutputInvokeFunction) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OutputInvokeFunction) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *OutputInvokeFunction) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *OutputInvokeFunction) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetDurationMillis

`func (o *OutputInvokeFunction) GetDurationMillis() int64`

GetDurationMillis returns the DurationMillis field if non-nil, zero value otherwise.

### GetDurationMillisOk

`func (o *OutputInvokeFunction) GetDurationMillisOk() (*int64, bool)`

GetDurationMillisOk returns a tuple with the DurationMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMillis

`func (o *OutputInvokeFunction) SetDurationMillis(v int64)`

SetDurationMillis sets DurationMillis field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


