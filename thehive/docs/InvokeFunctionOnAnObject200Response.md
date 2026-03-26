# InvokeFunctionOnAnObject200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**DurationMillis** | **int64** | Time taken by the function | 
**Stdout** | **string** | Content of stdout | 
**Stderr** | **string** | Content of stderr | 

## Methods

### NewInvokeFunctionOnAnObject200Response

`func NewInvokeFunctionOnAnObject200Response(result interface{}, durationMillis int64, stdout string, stderr string, ) *InvokeFunctionOnAnObject200Response`

NewInvokeFunctionOnAnObject200Response instantiates a new InvokeFunctionOnAnObject200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvokeFunctionOnAnObject200ResponseWithDefaults

`func NewInvokeFunctionOnAnObject200ResponseWithDefaults() *InvokeFunctionOnAnObject200Response`

NewInvokeFunctionOnAnObject200ResponseWithDefaults instantiates a new InvokeFunctionOnAnObject200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *InvokeFunctionOnAnObject200Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *InvokeFunctionOnAnObject200Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *InvokeFunctionOnAnObject200Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *InvokeFunctionOnAnObject200Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *InvokeFunctionOnAnObject200Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetDurationMillis

`func (o *InvokeFunctionOnAnObject200Response) GetDurationMillis() int64`

GetDurationMillis returns the DurationMillis field if non-nil, zero value otherwise.

### GetDurationMillisOk

`func (o *InvokeFunctionOnAnObject200Response) GetDurationMillisOk() (*int64, bool)`

GetDurationMillisOk returns a tuple with the DurationMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMillis

`func (o *InvokeFunctionOnAnObject200Response) SetDurationMillis(v int64)`

SetDurationMillis sets DurationMillis field to given value.


### GetStdout

`func (o *InvokeFunctionOnAnObject200Response) GetStdout() string`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *InvokeFunctionOnAnObject200Response) GetStdoutOk() (*string, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *InvokeFunctionOnAnObject200Response) SetStdout(v string)`

SetStdout sets Stdout field to given value.


### GetStderr

`func (o *InvokeFunctionOnAnObject200Response) GetStderr() string`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *InvokeFunctionOnAnObject200Response) GetStderrOk() (*string, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *InvokeFunctionOnAnObject200Response) SetStderr(v string)`

SetStderr sets Stderr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


