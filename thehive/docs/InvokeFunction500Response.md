# InvokeFunction500Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorType** | **string** | Error type encountered by the server | 
**ErrorMessage** | **string** | Precise error that led to execution failure | 
**Stdout** | **string** | Content of stdout | 
**Stderr** | **string** | Content of stderr | 
**Type** | **string** |  | 
**Message** | **string** |  | 

## Methods

### NewInvokeFunction500Response

`func NewInvokeFunction500Response(errorType string, errorMessage string, stdout string, stderr string, type_ string, message string, ) *InvokeFunction500Response`

NewInvokeFunction500Response instantiates a new InvokeFunction500Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvokeFunction500ResponseWithDefaults

`func NewInvokeFunction500ResponseWithDefaults() *InvokeFunction500Response`

NewInvokeFunction500ResponseWithDefaults instantiates a new InvokeFunction500Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorType

`func (o *InvokeFunction500Response) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *InvokeFunction500Response) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *InvokeFunction500Response) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.


### GetErrorMessage

`func (o *InvokeFunction500Response) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *InvokeFunction500Response) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *InvokeFunction500Response) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### GetStdout

`func (o *InvokeFunction500Response) GetStdout() string`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *InvokeFunction500Response) GetStdoutOk() (*string, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *InvokeFunction500Response) SetStdout(v string)`

SetStdout sets Stdout field to given value.


### GetStderr

`func (o *InvokeFunction500Response) GetStderr() string`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *InvokeFunction500Response) GetStderrOk() (*string, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *InvokeFunction500Response) SetStderr(v string)`

SetStderr sets Stderr field to given value.


### GetType

`func (o *InvokeFunction500Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvokeFunction500Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvokeFunction500Response) SetType(v string)`

SetType sets Type field to given value.


### GetMessage

`func (o *InvokeFunction500Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InvokeFunction500Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InvokeFunction500Response) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


