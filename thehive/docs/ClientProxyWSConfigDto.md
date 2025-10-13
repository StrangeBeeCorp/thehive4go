# ClientProxyWSConfigDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeout** | Pointer to [**Timeout**](Timeout.md) |  | [optional] 
**FollowRedirects** | Pointer to **bool** |  | [optional] [default to true]
**UseProxyProperties** | Pointer to **bool** |  | [optional] [default to true]
**UserAgent** | Pointer to **string** |  | [optional] 
**CompressionEnabled** | Pointer to **bool** |  | [optional] [default to false]
**Ssl** | [**SSLConfig**](SSLConfig.md) |  | 
**MaxConnectionsPerHost** | Pointer to **int32** |  | [optional] [default to -1]
**MaxConnectionsTotal** | Pointer to **int32** |  | [optional] [default to -1]
**MaxConnectionLifetime** | Pointer to **string** |  | [optional] [default to "Duration.Inf"]
**IdleConnectionInPoolTimeout** | Pointer to **string** |  | [optional] [default to "1 minute"]
**ConnectionPoolCleanerPeriod** | Pointer to **string** |  | [optional] [default to "1 minute"]
**MaxNumberOfRedirects** | Pointer to **int32** |  | [optional] [default to 5]
**MaxRequestRetry** | Pointer to **int32** |  | [optional] [default to 5]
**DisableUrlEncoding** | Pointer to **bool** |  | [optional] [default to false]
**KeepAlive** | Pointer to **bool** |  | [optional] [default to true]
**UseLaxCookieEncoder** | Pointer to **bool** |  | [optional] [default to false]
**UseCookieStore** | Pointer to **bool** |  | [optional] [default to false]
**Proxy** | Pointer to [**ProxyConfig**](ProxyConfig.md) |  | [optional] 

## Methods

### NewClientProxyWSConfigDto

`func NewClientProxyWSConfigDto(ssl SSLConfig, ) *ClientProxyWSConfigDto`

NewClientProxyWSConfigDto instantiates a new ClientProxyWSConfigDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientProxyWSConfigDtoWithDefaults

`func NewClientProxyWSConfigDtoWithDefaults() *ClientProxyWSConfigDto`

NewClientProxyWSConfigDtoWithDefaults instantiates a new ClientProxyWSConfigDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeout

`func (o *ClientProxyWSConfigDto) GetTimeout() Timeout`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ClientProxyWSConfigDto) GetTimeoutOk() (*Timeout, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ClientProxyWSConfigDto) SetTimeout(v Timeout)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ClientProxyWSConfigDto) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetFollowRedirects

`func (o *ClientProxyWSConfigDto) GetFollowRedirects() bool`

GetFollowRedirects returns the FollowRedirects field if non-nil, zero value otherwise.

### GetFollowRedirectsOk

`func (o *ClientProxyWSConfigDto) GetFollowRedirectsOk() (*bool, bool)`

GetFollowRedirectsOk returns a tuple with the FollowRedirects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowRedirects

`func (o *ClientProxyWSConfigDto) SetFollowRedirects(v bool)`

SetFollowRedirects sets FollowRedirects field to given value.

### HasFollowRedirects

`func (o *ClientProxyWSConfigDto) HasFollowRedirects() bool`

HasFollowRedirects returns a boolean if a field has been set.

### GetUseProxyProperties

`func (o *ClientProxyWSConfigDto) GetUseProxyProperties() bool`

GetUseProxyProperties returns the UseProxyProperties field if non-nil, zero value otherwise.

### GetUseProxyPropertiesOk

`func (o *ClientProxyWSConfigDto) GetUseProxyPropertiesOk() (*bool, bool)`

GetUseProxyPropertiesOk returns a tuple with the UseProxyProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseProxyProperties

`func (o *ClientProxyWSConfigDto) SetUseProxyProperties(v bool)`

SetUseProxyProperties sets UseProxyProperties field to given value.

### HasUseProxyProperties

`func (o *ClientProxyWSConfigDto) HasUseProxyProperties() bool`

HasUseProxyProperties returns a boolean if a field has been set.

### GetUserAgent

`func (o *ClientProxyWSConfigDto) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *ClientProxyWSConfigDto) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *ClientProxyWSConfigDto) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *ClientProxyWSConfigDto) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetCompressionEnabled

`func (o *ClientProxyWSConfigDto) GetCompressionEnabled() bool`

GetCompressionEnabled returns the CompressionEnabled field if non-nil, zero value otherwise.

### GetCompressionEnabledOk

`func (o *ClientProxyWSConfigDto) GetCompressionEnabledOk() (*bool, bool)`

GetCompressionEnabledOk returns a tuple with the CompressionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionEnabled

`func (o *ClientProxyWSConfigDto) SetCompressionEnabled(v bool)`

SetCompressionEnabled sets CompressionEnabled field to given value.

### HasCompressionEnabled

`func (o *ClientProxyWSConfigDto) HasCompressionEnabled() bool`

HasCompressionEnabled returns a boolean if a field has been set.

### GetSsl

`func (o *ClientProxyWSConfigDto) GetSsl() SSLConfig`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *ClientProxyWSConfigDto) GetSslOk() (*SSLConfig, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *ClientProxyWSConfigDto) SetSsl(v SSLConfig)`

SetSsl sets Ssl field to given value.


### GetMaxConnectionsPerHost

`func (o *ClientProxyWSConfigDto) GetMaxConnectionsPerHost() int32`

GetMaxConnectionsPerHost returns the MaxConnectionsPerHost field if non-nil, zero value otherwise.

### GetMaxConnectionsPerHostOk

`func (o *ClientProxyWSConfigDto) GetMaxConnectionsPerHostOk() (*int32, bool)`

GetMaxConnectionsPerHostOk returns a tuple with the MaxConnectionsPerHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnectionsPerHost

`func (o *ClientProxyWSConfigDto) SetMaxConnectionsPerHost(v int32)`

SetMaxConnectionsPerHost sets MaxConnectionsPerHost field to given value.

### HasMaxConnectionsPerHost

`func (o *ClientProxyWSConfigDto) HasMaxConnectionsPerHost() bool`

HasMaxConnectionsPerHost returns a boolean if a field has been set.

### GetMaxConnectionsTotal

`func (o *ClientProxyWSConfigDto) GetMaxConnectionsTotal() int32`

GetMaxConnectionsTotal returns the MaxConnectionsTotal field if non-nil, zero value otherwise.

### GetMaxConnectionsTotalOk

`func (o *ClientProxyWSConfigDto) GetMaxConnectionsTotalOk() (*int32, bool)`

GetMaxConnectionsTotalOk returns a tuple with the MaxConnectionsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnectionsTotal

`func (o *ClientProxyWSConfigDto) SetMaxConnectionsTotal(v int32)`

SetMaxConnectionsTotal sets MaxConnectionsTotal field to given value.

### HasMaxConnectionsTotal

`func (o *ClientProxyWSConfigDto) HasMaxConnectionsTotal() bool`

HasMaxConnectionsTotal returns a boolean if a field has been set.

### GetMaxConnectionLifetime

`func (o *ClientProxyWSConfigDto) GetMaxConnectionLifetime() string`

GetMaxConnectionLifetime returns the MaxConnectionLifetime field if non-nil, zero value otherwise.

### GetMaxConnectionLifetimeOk

`func (o *ClientProxyWSConfigDto) GetMaxConnectionLifetimeOk() (*string, bool)`

GetMaxConnectionLifetimeOk returns a tuple with the MaxConnectionLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnectionLifetime

`func (o *ClientProxyWSConfigDto) SetMaxConnectionLifetime(v string)`

SetMaxConnectionLifetime sets MaxConnectionLifetime field to given value.

### HasMaxConnectionLifetime

`func (o *ClientProxyWSConfigDto) HasMaxConnectionLifetime() bool`

HasMaxConnectionLifetime returns a boolean if a field has been set.

### GetIdleConnectionInPoolTimeout

`func (o *ClientProxyWSConfigDto) GetIdleConnectionInPoolTimeout() string`

GetIdleConnectionInPoolTimeout returns the IdleConnectionInPoolTimeout field if non-nil, zero value otherwise.

### GetIdleConnectionInPoolTimeoutOk

`func (o *ClientProxyWSConfigDto) GetIdleConnectionInPoolTimeoutOk() (*string, bool)`

GetIdleConnectionInPoolTimeoutOk returns a tuple with the IdleConnectionInPoolTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleConnectionInPoolTimeout

`func (o *ClientProxyWSConfigDto) SetIdleConnectionInPoolTimeout(v string)`

SetIdleConnectionInPoolTimeout sets IdleConnectionInPoolTimeout field to given value.

### HasIdleConnectionInPoolTimeout

`func (o *ClientProxyWSConfigDto) HasIdleConnectionInPoolTimeout() bool`

HasIdleConnectionInPoolTimeout returns a boolean if a field has been set.

### GetConnectionPoolCleanerPeriod

`func (o *ClientProxyWSConfigDto) GetConnectionPoolCleanerPeriod() string`

GetConnectionPoolCleanerPeriod returns the ConnectionPoolCleanerPeriod field if non-nil, zero value otherwise.

### GetConnectionPoolCleanerPeriodOk

`func (o *ClientProxyWSConfigDto) GetConnectionPoolCleanerPeriodOk() (*string, bool)`

GetConnectionPoolCleanerPeriodOk returns a tuple with the ConnectionPoolCleanerPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPoolCleanerPeriod

`func (o *ClientProxyWSConfigDto) SetConnectionPoolCleanerPeriod(v string)`

SetConnectionPoolCleanerPeriod sets ConnectionPoolCleanerPeriod field to given value.

### HasConnectionPoolCleanerPeriod

`func (o *ClientProxyWSConfigDto) HasConnectionPoolCleanerPeriod() bool`

HasConnectionPoolCleanerPeriod returns a boolean if a field has been set.

### GetMaxNumberOfRedirects

`func (o *ClientProxyWSConfigDto) GetMaxNumberOfRedirects() int32`

GetMaxNumberOfRedirects returns the MaxNumberOfRedirects field if non-nil, zero value otherwise.

### GetMaxNumberOfRedirectsOk

`func (o *ClientProxyWSConfigDto) GetMaxNumberOfRedirectsOk() (*int32, bool)`

GetMaxNumberOfRedirectsOk returns a tuple with the MaxNumberOfRedirects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfRedirects

`func (o *ClientProxyWSConfigDto) SetMaxNumberOfRedirects(v int32)`

SetMaxNumberOfRedirects sets MaxNumberOfRedirects field to given value.

### HasMaxNumberOfRedirects

`func (o *ClientProxyWSConfigDto) HasMaxNumberOfRedirects() bool`

HasMaxNumberOfRedirects returns a boolean if a field has been set.

### GetMaxRequestRetry

`func (o *ClientProxyWSConfigDto) GetMaxRequestRetry() int32`

GetMaxRequestRetry returns the MaxRequestRetry field if non-nil, zero value otherwise.

### GetMaxRequestRetryOk

`func (o *ClientProxyWSConfigDto) GetMaxRequestRetryOk() (*int32, bool)`

GetMaxRequestRetryOk returns a tuple with the MaxRequestRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRequestRetry

`func (o *ClientProxyWSConfigDto) SetMaxRequestRetry(v int32)`

SetMaxRequestRetry sets MaxRequestRetry field to given value.

### HasMaxRequestRetry

`func (o *ClientProxyWSConfigDto) HasMaxRequestRetry() bool`

HasMaxRequestRetry returns a boolean if a field has been set.

### GetDisableUrlEncoding

`func (o *ClientProxyWSConfigDto) GetDisableUrlEncoding() bool`

GetDisableUrlEncoding returns the DisableUrlEncoding field if non-nil, zero value otherwise.

### GetDisableUrlEncodingOk

`func (o *ClientProxyWSConfigDto) GetDisableUrlEncodingOk() (*bool, bool)`

GetDisableUrlEncodingOk returns a tuple with the DisableUrlEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableUrlEncoding

`func (o *ClientProxyWSConfigDto) SetDisableUrlEncoding(v bool)`

SetDisableUrlEncoding sets DisableUrlEncoding field to given value.

### HasDisableUrlEncoding

`func (o *ClientProxyWSConfigDto) HasDisableUrlEncoding() bool`

HasDisableUrlEncoding returns a boolean if a field has been set.

### GetKeepAlive

`func (o *ClientProxyWSConfigDto) GetKeepAlive() bool`

GetKeepAlive returns the KeepAlive field if non-nil, zero value otherwise.

### GetKeepAliveOk

`func (o *ClientProxyWSConfigDto) GetKeepAliveOk() (*bool, bool)`

GetKeepAliveOk returns a tuple with the KeepAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepAlive

`func (o *ClientProxyWSConfigDto) SetKeepAlive(v bool)`

SetKeepAlive sets KeepAlive field to given value.

### HasKeepAlive

`func (o *ClientProxyWSConfigDto) HasKeepAlive() bool`

HasKeepAlive returns a boolean if a field has been set.

### GetUseLaxCookieEncoder

`func (o *ClientProxyWSConfigDto) GetUseLaxCookieEncoder() bool`

GetUseLaxCookieEncoder returns the UseLaxCookieEncoder field if non-nil, zero value otherwise.

### GetUseLaxCookieEncoderOk

`func (o *ClientProxyWSConfigDto) GetUseLaxCookieEncoderOk() (*bool, bool)`

GetUseLaxCookieEncoderOk returns a tuple with the UseLaxCookieEncoder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLaxCookieEncoder

`func (o *ClientProxyWSConfigDto) SetUseLaxCookieEncoder(v bool)`

SetUseLaxCookieEncoder sets UseLaxCookieEncoder field to given value.

### HasUseLaxCookieEncoder

`func (o *ClientProxyWSConfigDto) HasUseLaxCookieEncoder() bool`

HasUseLaxCookieEncoder returns a boolean if a field has been set.

### GetUseCookieStore

`func (o *ClientProxyWSConfigDto) GetUseCookieStore() bool`

GetUseCookieStore returns the UseCookieStore field if non-nil, zero value otherwise.

### GetUseCookieStoreOk

`func (o *ClientProxyWSConfigDto) GetUseCookieStoreOk() (*bool, bool)`

GetUseCookieStoreOk returns a tuple with the UseCookieStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCookieStore

`func (o *ClientProxyWSConfigDto) SetUseCookieStore(v bool)`

SetUseCookieStore sets UseCookieStore field to given value.

### HasUseCookieStore

`func (o *ClientProxyWSConfigDto) HasUseCookieStore() bool`

HasUseCookieStore returns a boolean if a field has been set.

### GetProxy

`func (o *ClientProxyWSConfigDto) GetProxy() ProxyConfig`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *ClientProxyWSConfigDto) GetProxyOk() (*ProxyConfig, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *ClientProxyWSConfigDto) SetProxy(v ProxyConfig)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *ClientProxyWSConfigDto) HasProxy() bool`

HasProxy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


