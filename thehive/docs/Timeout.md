# Timeout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connection** | **string** |  | 
**Idle** | **string** |  | 
**Request** | **string** |  | 

## Methods

### NewTimeout

`func NewTimeout(connection string, idle string, request string, ) *Timeout`

NewTimeout instantiates a new Timeout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeoutWithDefaults

`func NewTimeoutWithDefaults() *Timeout`

NewTimeoutWithDefaults instantiates a new Timeout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnection

`func (o *Timeout) GetConnection() string`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *Timeout) GetConnectionOk() (*string, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *Timeout) SetConnection(v string)`

SetConnection sets Connection field to given value.


### GetIdle

`func (o *Timeout) GetIdle() string`

GetIdle returns the Idle field if non-nil, zero value otherwise.

### GetIdleOk

`func (o *Timeout) GetIdleOk() (*string, bool)`

GetIdleOk returns a tuple with the Idle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdle

`func (o *Timeout) SetIdle(v string)`

SetIdle sets Idle field to given value.


### GetRequest

`func (o *Timeout) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *Timeout) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *Timeout) SetRequest(v string)`

SetRequest sets Request field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


