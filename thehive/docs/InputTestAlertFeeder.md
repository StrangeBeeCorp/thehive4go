# InputTestAlertFeeder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique name for the alert feeder. |
**Description** | **string** | Description providing additional context about the alert feeder configuration. |
**Method** | [**Method**](Method.md) | HTTP method to use when requesting data from the external system. |
**Url** | **string** | Endpoint URL of the external system. |
**Interval** | [**Interval**](Interval.md) | Polling interval. The minimum interval is 1 minute. |
**Body** | Pointer to **string** | Optional HTTP request body to send with the request. Used when the HTTP method is &#x60;POST&#x60;. | [optional]
**Headers** | Pointer to [**[]Header**](Header.md) | Optional HTTP headers to include in the request. | [optional]
**Enabled** | Pointer to **bool** | Whether the alert feeder is active and runs on schedule. | [optional] [default to true]
**Auth** | Pointer to [**UpdateAuth**](UpdateAuth.md) | Optional authentication configuration. Supports &#x60;none&#x60;, &#x60;basic&#x60;, &#x60;bearer&#x60;, &#x60;key&#x60;, and &#x60;oauth2&#x60;. Omitting this field or setting it to &#x60;none&#x60; means no authentication. | [optional]
**ProxyConfig** | Pointer to [**ClientProxyWSConfigDto**](ClientProxyWSConfigDto.md) | Optional proxy configuration for the HTTP request. | [optional]
**RequestTimeout** | Pointer to [**Interval**](Interval.md) | Maximum time to wait for the external system to respond. | [optional]
**ResponseMaxSize** | Pointer to **int64** | Maximum response payload size in bytes that the alert feeder accepts from the external system. | [optional] [default to 10485760]

## Methods

### NewInputTestAlertFeeder

`func NewInputTestAlertFeeder(name string, description string, method Method, url string, interval Interval, ) *InputTestAlertFeeder`

NewInputTestAlertFeeder instantiates a new InputTestAlertFeeder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputTestAlertFeederWithDefaults

`func NewInputTestAlertFeederWithDefaults() *InputTestAlertFeeder`

NewInputTestAlertFeederWithDefaults instantiates a new InputTestAlertFeeder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InputTestAlertFeeder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputTestAlertFeeder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputTestAlertFeeder) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *InputTestAlertFeeder) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputTestAlertFeeder) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputTestAlertFeeder) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMethod

`func (o *InputTestAlertFeeder) GetMethod() Method`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *InputTestAlertFeeder) GetMethodOk() (*Method, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *InputTestAlertFeeder) SetMethod(v Method)`

SetMethod sets Method field to given value.


### GetUrl

`func (o *InputTestAlertFeeder) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InputTestAlertFeeder) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InputTestAlertFeeder) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetInterval

`func (o *InputTestAlertFeeder) GetInterval() Interval`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *InputTestAlertFeeder) GetIntervalOk() (*Interval, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *InputTestAlertFeeder) SetInterval(v Interval)`

SetInterval sets Interval field to given value.


### GetBody

`func (o *InputTestAlertFeeder) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *InputTestAlertFeeder) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *InputTestAlertFeeder) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *InputTestAlertFeeder) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetHeaders

`func (o *InputTestAlertFeeder) GetHeaders() []Header`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *InputTestAlertFeeder) GetHeadersOk() (*[]Header, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *InputTestAlertFeeder) SetHeaders(v []Header)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *InputTestAlertFeeder) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetEnabled

`func (o *InputTestAlertFeeder) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InputTestAlertFeeder) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InputTestAlertFeeder) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InputTestAlertFeeder) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAuth

`func (o *InputTestAlertFeeder) GetAuth() UpdateAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *InputTestAlertFeeder) GetAuthOk() (*UpdateAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *InputTestAlertFeeder) SetAuth(v UpdateAuth)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *InputTestAlertFeeder) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetProxyConfig

`func (o *InputTestAlertFeeder) GetProxyConfig() ClientProxyWSConfigDto`

GetProxyConfig returns the ProxyConfig field if non-nil, zero value otherwise.

### GetProxyConfigOk

`func (o *InputTestAlertFeeder) GetProxyConfigOk() (*ClientProxyWSConfigDto, bool)`

GetProxyConfigOk returns a tuple with the ProxyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyConfig

`func (o *InputTestAlertFeeder) SetProxyConfig(v ClientProxyWSConfigDto)`

SetProxyConfig sets ProxyConfig field to given value.

### HasProxyConfig

`func (o *InputTestAlertFeeder) HasProxyConfig() bool`

HasProxyConfig returns a boolean if a field has been set.

### GetRequestTimeout

`func (o *InputTestAlertFeeder) GetRequestTimeout() Interval`

GetRequestTimeout returns the RequestTimeout field if non-nil, zero value otherwise.

### GetRequestTimeoutOk

`func (o *InputTestAlertFeeder) GetRequestTimeoutOk() (*Interval, bool)`

GetRequestTimeoutOk returns a tuple with the RequestTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTimeout

`func (o *InputTestAlertFeeder) SetRequestTimeout(v Interval)`

SetRequestTimeout sets RequestTimeout field to given value.

### HasRequestTimeout

`func (o *InputTestAlertFeeder) HasRequestTimeout() bool`

HasRequestTimeout returns a boolean if a field has been set.

### GetResponseMaxSize

`func (o *InputTestAlertFeeder) GetResponseMaxSize() int64`

GetResponseMaxSize returns the ResponseMaxSize field if non-nil, zero value otherwise.

### GetResponseMaxSizeOk

`func (o *InputTestAlertFeeder) GetResponseMaxSizeOk() (*int64, bool)`

GetResponseMaxSizeOk returns a tuple with the ResponseMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMaxSize

`func (o *InputTestAlertFeeder) SetResponseMaxSize(v int64)`

SetResponseMaxSize sets ResponseMaxSize field to given value.

### HasResponseMaxSize

`func (o *InputTestAlertFeeder) HasResponseMaxSize() bool`

HasResponseMaxSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
