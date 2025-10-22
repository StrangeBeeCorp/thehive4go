# OutputAlertFeeder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** |  | 
**Method** | [**Method**](Method.md) |  | 
**Url** | **string** |  | 
**Interval** | [**Interval**](Interval.md) |  | 
**Function** | [**OutputFunction**](OutputFunction.md) |  | 
**Headers** | Pointer to [**[]Header**](Header.md) |  | [optional] 
**Auth** | Pointer to [**Auth**](Auth.md) |  | [optional] 
**Body** | Pointer to **string** |  | [optional] 
**Enabled** | **bool** |  | 
**RequestTimeout** | [**Interval**](Interval.md) |  | 
**ResponseMaxSize** | **int64** |  | 
**ProxyConfig** | Pointer to [**ClientProxyWSConfigDto**](ClientProxyWSConfigDto.md) |  | [optional] 

## Methods

### NewOutputAlertFeeder

`func NewOutputAlertFeeder(name string, description string, method Method, url string, interval Interval, function OutputFunction, enabled bool, requestTimeout Interval, responseMaxSize int64, ) *OutputAlertFeeder`

NewOutputAlertFeeder instantiates a new OutputAlertFeeder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputAlertFeederWithDefaults

`func NewOutputAlertFeederWithDefaults() *OutputAlertFeeder`

NewOutputAlertFeederWithDefaults instantiates a new OutputAlertFeeder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OutputAlertFeeder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OutputAlertFeeder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OutputAlertFeeder) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *OutputAlertFeeder) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OutputAlertFeeder) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OutputAlertFeeder) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMethod

`func (o *OutputAlertFeeder) GetMethod() Method`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *OutputAlertFeeder) GetMethodOk() (*Method, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *OutputAlertFeeder) SetMethod(v Method)`

SetMethod sets Method field to given value.


### GetUrl

`func (o *OutputAlertFeeder) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OutputAlertFeeder) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OutputAlertFeeder) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetInterval

`func (o *OutputAlertFeeder) GetInterval() Interval`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *OutputAlertFeeder) GetIntervalOk() (*Interval, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *OutputAlertFeeder) SetInterval(v Interval)`

SetInterval sets Interval field to given value.


### GetFunction

`func (o *OutputAlertFeeder) GetFunction() OutputFunction`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *OutputAlertFeeder) GetFunctionOk() (*OutputFunction, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *OutputAlertFeeder) SetFunction(v OutputFunction)`

SetFunction sets Function field to given value.


### GetHeaders

`func (o *OutputAlertFeeder) GetHeaders() []Header`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *OutputAlertFeeder) GetHeadersOk() (*[]Header, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *OutputAlertFeeder) SetHeaders(v []Header)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *OutputAlertFeeder) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetAuth

`func (o *OutputAlertFeeder) GetAuth() Auth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *OutputAlertFeeder) GetAuthOk() (*Auth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *OutputAlertFeeder) SetAuth(v Auth)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *OutputAlertFeeder) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetBody

`func (o *OutputAlertFeeder) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *OutputAlertFeeder) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *OutputAlertFeeder) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *OutputAlertFeeder) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetEnabled

`func (o *OutputAlertFeeder) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OutputAlertFeeder) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OutputAlertFeeder) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetRequestTimeout

`func (o *OutputAlertFeeder) GetRequestTimeout() Interval`

GetRequestTimeout returns the RequestTimeout field if non-nil, zero value otherwise.

### GetRequestTimeoutOk

`func (o *OutputAlertFeeder) GetRequestTimeoutOk() (*Interval, bool)`

GetRequestTimeoutOk returns a tuple with the RequestTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTimeout

`func (o *OutputAlertFeeder) SetRequestTimeout(v Interval)`

SetRequestTimeout sets RequestTimeout field to given value.


### GetResponseMaxSize

`func (o *OutputAlertFeeder) GetResponseMaxSize() int64`

GetResponseMaxSize returns the ResponseMaxSize field if non-nil, zero value otherwise.

### GetResponseMaxSizeOk

`func (o *OutputAlertFeeder) GetResponseMaxSizeOk() (*int64, bool)`

GetResponseMaxSizeOk returns a tuple with the ResponseMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMaxSize

`func (o *OutputAlertFeeder) SetResponseMaxSize(v int64)`

SetResponseMaxSize sets ResponseMaxSize field to given value.


### GetProxyConfig

`func (o *OutputAlertFeeder) GetProxyConfig() ClientProxyWSConfigDto`

GetProxyConfig returns the ProxyConfig field if non-nil, zero value otherwise.

### GetProxyConfigOk

`func (o *OutputAlertFeeder) GetProxyConfigOk() (*ClientProxyWSConfigDto, bool)`

GetProxyConfigOk returns a tuple with the ProxyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyConfig

`func (o *OutputAlertFeeder) SetProxyConfig(v ClientProxyWSConfigDto)`

SetProxyConfig sets ProxyConfig field to given value.

### HasProxyConfig

`func (o *OutputAlertFeeder) HasProxyConfig() bool`

HasProxyConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


