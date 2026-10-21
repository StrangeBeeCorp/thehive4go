/*
TheHive

 ## Get started  This guide covers everything you need to make your first successful API call to TheHive.  The publicly hosted API documentation reflects the latest version of TheHive. To browse the documentation that matches the version of TheHive you are running, open it from within the application: select the **?** button in the top right corner, then **API Docs**.  The v0 API is deprecated and isn't documented. New integrations should target v1.  ### Prerequisites  You need a running TheHive instance and an unlocked user account.  ### Authentication  Almost all endpoints require authentication.  TheHive supports two authentication methods:  #### API key (recommended)  Use your API key as a bearer token in the `Authorization` header.  Generate a new API key with `POST /api/v1/user/{userId}/key/renew`, or retrieve an existing one with `GET /api/v1/user/{userId}/key`.  If you have a [Normal user account](https://docs.strangebee.com/thehive/user-guides/organization/configure-organization/manage-user-accounts/about-user-accounts/#types), you can also manage your API key in your [user settings](https://docs.strangebee.com/thehive/user-guides/manage-user-settings/#manage-your-api-key).  curl ```bash curl -X <http_method> -H 'Authorization: Bearer <api_key>' \\   https://<thehive_host>/api/v1/<endpoint> ```  python ```python import requests  headers = {\"Authorization\": \"Bearer <api_key>\"} response = requests.<http_method>(\"https://<thehive_host>/api/v1/<endpoint>\", headers=headers) ```  #### Basic authentication  Only available for [Normal user accounts](https://docs.strangebee.com/thehive/user-guides/organization/configure-organization/manage-user-accounts/about-user-accounts/#types).  Pass your login and password using HTTP Basic authentication. The login is typically an email address. If a default domain is configured in the [authentication settings](https://docs.strangebee.com/thehive/administration/authentication/configure-authentication), a plain username is also accepted.  curl ```bash curl -X <http_method> -u <login>:<password> \\   https://<thehive_host>/api/v1/<endpoint> ```  python ```python import requests  response = requests.<http_method>(\"https://<thehive_host>/api/v1/<endpoint>\", auth=(\"<login>\", \"<password>\")) ```  The following endpoints are public and can be called without credentials:  - `POST /api/v1/login`: Authenticate with login, password, and a TOTP code if MFA is enabled. Returns a session cookie used to authenticate subsequent requests. - `GET /api/v1/logout`: Invalidate the current session. - `POST /api/v1/logout`: Invalidate the current session. Same as `GET`, provided for clients that can't send `GET` requests with a body. - `GET /api/v1/auth/local/passwordPolicy`: Retrieve the password policy configured on the server such as minimum length or complexity requirements. - `GET /api/v1/status/public`: Retrieve public server information, including the TheHive version and license status, without exposing sensitive configuration.  All other endpoints return `401 Unauthorized` when called without valid credentials.  ### Your first API call  Start with `GET /api/v1/user/current`: it requires no request body, confirms your credentials work, and returns your user ID.  curl ```bash curl -X GET -H 'Authorization: Bearer <api_key>' \\      https://<thehive_host>/api/v1/user/current ```  python ```python import requests  BASE_URL = \"https://<thehive_host>\" API_KEY = \"<api_key>\"  headers = {\"Authorization\": f\"Bearer {API_KEY}\"} response = requests.get(f\"{BASE_URL}/api/v1/user/current\", headers=headers) user = response.json() ```  ### X-Organisation header  By default, API calls target the user's default organization. The `defaultOrganisation` field returned by `GET /api/v1/user/current` tells you which one it is.  If your account belongs to several organizations and you need to target a specific one, add the `X-Organisation` header. This is also required when performing administration actions, as these must be executed in the context of the `admin` organization.  curl ```bash curl -H 'Authorization: Bearer <api_key>' \\      -H 'X-Organisation: <organization_name>' \\      https://<thehive_host>/api/v1/user/current ```  python ```python import requests  BASE_URL = \"https://<thehive_host>\" API_KEY = \"<api_key>\"  headers = {\"Authorization\": f\"Bearer {API_KEY}\", \"X-Organisation\": \"<organization_name>\"} response = requests.get(f\"{BASE_URL}/api/v1/user/current\", headers=headers) user = response.json() ```  ### Request and response format  All request and response bodies use JSON: `Content-Type: application/json`.  Dates are represented as Unix timestamps in milliseconds.  ### How the API is organized  TheHive exposes two kinds of endpoints.  #### The Query API  `POST /api/v1/query` handles all listing, filtering, sorting, and pagination needs.  A single endpoint accepts a structured body that chains operations:  ```json {   \"query\": [     { \"_name\": \"listCase\" },     { \"_name\": \"filter\", \"_eq\": { \"_field\": \"status\", \"_value\": \"Open\" } },     { \"_name\": \"sort\", \"_fields\": [{ \"_createdAt\": \"desc\" }] },     { \"_name\": \"page\", \"from\": 0, \"to\": 15, \"extraData\": [\"taskStats\", \"shareCount\"] }   ],   \"excludeFields\": [\"description\", \"summary\"] } ```  See the [**Query and Export**](https://docs.strangebee.com/thehive/api-docs/#tag/Query-and-Export) section for the full query reference.  #### Standard REST endpoints  Standard REST endpoints handle write operations and direct entity lookups.  They follow conventional HTTP conventions:  | <nobr>Operation </nobr>| Method | Example | |-----------|--------|---------| | <nobr>Create </nobr>| `POST` | `POST /api/v1/case` | | <nobr>Read </nobr>| `GET` | `GET /api/v1/case/{idOrName}` | | <nobr>Update </nobr>| `PATCH` | `PATCH /api/v1/case/{idOrName}` | | <nobr>Replace / set </nobr>| `PUT` | `PUT /api/v1/case/{caseId}/shares` | | <nobr>Delete </nobr>| `DELETE` | `DELETE /api/v1/case/{idOrName}` |  ### Permissions and capabilities  Each user account has permissions defined by their profile. TheHive checks them on every API call.  Some endpoints are only available with a Gold, Platinum, or One license. When that's the case, the required license tier is listed on the endpoint.  To check which license is active on your instance and what capabilities it includes, call `GET /api/v1/status`.  ### Request size limits  TheHive enforces the following request size limits:  | <nobr>Request type </nobr>| Default limit | Description | |---|---|---| | <nobr>Standard API requests (JSON payloads) </nobr>| 10 MB | Maximum size for request bodies processed in memory | | <nobr>File uploads and attachments </nobr>| 1 GB | Maximum size for streamed request bodies written to disk |  TheHive rejects requests that exceed these limits.  To stay within these limits: - Avoid sending excessively large JSON payloads. - Upload large files as attachments rather than embedding data directly in requests. - Implement rate limiting and additional protections at the reverse proxy level when exposing the API externally.  ### Errors  Every endpoint returns at least the following HTTP status codes:  | <nobr>Code </nobr>| Meaning | |------|---------| | <nobr>`400 Bad Request` </nobr>| Invalid request parameters or body | | <nobr>`404 Not Found` </nobr>| Resource does not exist | | <nobr>`500 Internal Server Error` </nobr>| Unexpected server-side error |  Endpoints that require authentication additionally return:  | <nobr>Code </nobr>| Meaning | |------|---------| | <nobr>`401 Unauthorized` </nobr>| Missing or invalid credentials | | <nobr>`403 Forbidden` </nobr>| Authenticated but insufficient permissions |  Other codes may appear depending on the specific operation. Refer to individual endpoint documentation for the full list.  The standard error response body is:  ```json {   \"type\": \"AuthorizationError\",   \"message\": \"You don't have permission to perform this action\" } ```  A small number of endpoints deviate from this format. These cases are noted in the relevant endpoint descriptions.

API version: v6.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package thehive

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ClientProxyWSConfigDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientProxyWSConfigDto{}

// ClientProxyWSConfigDto struct for ClientProxyWSConfigDto
type ClientProxyWSConfigDto struct {
	Timeout *Timeout `json:"timeout,omitempty"`
	// Whether to follow HTTP redirects automatically.
	FollowRedirects *bool `json:"followRedirects,omitempty"`
	// Whether to use system proxy properties.
	UseProxyProperties *bool `json:"useProxyProperties,omitempty"`
	// Custom User-Agent header value to send with requests.
	UserAgent *string `json:"userAgent,omitempty"`
	// Whether to enable HTTP response compression.
	CompressionEnabled *bool     `json:"compressionEnabled,omitempty"`
	Ssl                SSLConfig `json:"ssl"`
	// Maximum number of connections per host. Use -1 for no limit.
	MaxConnectionsPerHost *int32 `json:"maxConnectionsPerHost,omitempty"`
	// Maximum total number of connections. Use -1 for no limit.
	MaxConnectionsTotal *int32 `json:"maxConnectionsTotal,omitempty"`
	// Maximum lifetime of a connection in the pool.
	MaxConnectionLifetime *string `json:"maxConnectionLifetime,omitempty"`
	// How long an idle connection is kept in the pool before being discarded.
	IdleConnectionInPoolTimeout *string `json:"idleConnectionInPoolTimeout,omitempty"`
	// How often the connection pool cleaner runs.
	ConnectionPoolCleanerPeriod *string `json:"connectionPoolCleanerPeriod,omitempty"`
	// Maximum number of redirects to follow before failing.
	MaxNumberOfRedirects *int32 `json:"maxNumberOfRedirects,omitempty"`
	// Maximum number of times a failed request is retried.
	MaxRequestRetry *int32 `json:"maxRequestRetry,omitempty"`
	// Whether to skip URL encoding of query parameters.
	DisableUrlEncoding *bool `json:"disableUrlEncoding,omitempty"`
	// Whether to use HTTP keep-alive connections.
	KeepAlive *bool `json:"keepAlive,omitempty"`
	// Whether to use a lax cookie encoder that accepts non-standard cookie syntax.
	UseLaxCookieEncoder *bool `json:"useLaxCookieEncoder,omitempty"`
	// Whether to use a cookie store to persist cookies across requests.
	UseCookieStore *bool       `json:"useCookieStore,omitempty"`
	Proxy          ProxyConfig `json:"proxy"`
}

type _ClientProxyWSConfigDto ClientProxyWSConfigDto

// NewClientProxyWSConfigDto instantiates a new ClientProxyWSConfigDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientProxyWSConfigDto(ssl SSLConfig, proxy ProxyConfig) *ClientProxyWSConfigDto {
	this := ClientProxyWSConfigDto{}
	var followRedirects bool = true
	this.FollowRedirects = &followRedirects
	var useProxyProperties bool = true
	this.UseProxyProperties = &useProxyProperties
	var compressionEnabled bool = false
	this.CompressionEnabled = &compressionEnabled
	this.Ssl = ssl
	var maxConnectionsPerHost int32 = -1
	this.MaxConnectionsPerHost = &maxConnectionsPerHost
	var maxConnectionsTotal int32 = -1
	this.MaxConnectionsTotal = &maxConnectionsTotal
	var maxConnectionLifetime string = "Duration.Inf"
	this.MaxConnectionLifetime = &maxConnectionLifetime
	var idleConnectionInPoolTimeout string = "1 minute"
	this.IdleConnectionInPoolTimeout = &idleConnectionInPoolTimeout
	var connectionPoolCleanerPeriod string = "1 minute"
	this.ConnectionPoolCleanerPeriod = &connectionPoolCleanerPeriod
	var maxNumberOfRedirects int32 = 5
	this.MaxNumberOfRedirects = &maxNumberOfRedirects
	var maxRequestRetry int32 = 5
	this.MaxRequestRetry = &maxRequestRetry
	var disableUrlEncoding bool = false
	this.DisableUrlEncoding = &disableUrlEncoding
	var keepAlive bool = true
	this.KeepAlive = &keepAlive
	var useLaxCookieEncoder bool = false
	this.UseLaxCookieEncoder = &useLaxCookieEncoder
	var useCookieStore bool = false
	this.UseCookieStore = &useCookieStore
	this.Proxy = proxy
	return &this
}

// NewClientProxyWSConfigDtoWithDefaults instantiates a new ClientProxyWSConfigDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientProxyWSConfigDtoWithDefaults() *ClientProxyWSConfigDto {
	this := ClientProxyWSConfigDto{}
	var followRedirects bool = true
	this.FollowRedirects = &followRedirects
	var useProxyProperties bool = true
	this.UseProxyProperties = &useProxyProperties
	var compressionEnabled bool = false
	this.CompressionEnabled = &compressionEnabled
	var maxConnectionsPerHost int32 = -1
	this.MaxConnectionsPerHost = &maxConnectionsPerHost
	var maxConnectionsTotal int32 = -1
	this.MaxConnectionsTotal = &maxConnectionsTotal
	var maxConnectionLifetime string = "Duration.Inf"
	this.MaxConnectionLifetime = &maxConnectionLifetime
	var idleConnectionInPoolTimeout string = "1 minute"
	this.IdleConnectionInPoolTimeout = &idleConnectionInPoolTimeout
	var connectionPoolCleanerPeriod string = "1 minute"
	this.ConnectionPoolCleanerPeriod = &connectionPoolCleanerPeriod
	var maxNumberOfRedirects int32 = 5
	this.MaxNumberOfRedirects = &maxNumberOfRedirects
	var maxRequestRetry int32 = 5
	this.MaxRequestRetry = &maxRequestRetry
	var disableUrlEncoding bool = false
	this.DisableUrlEncoding = &disableUrlEncoding
	var keepAlive bool = true
	this.KeepAlive = &keepAlive
	var useLaxCookieEncoder bool = false
	this.UseLaxCookieEncoder = &useLaxCookieEncoder
	var useCookieStore bool = false
	this.UseCookieStore = &useCookieStore
	return &this
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *ClientProxyWSConfigDto) GetTimeout() Timeout {
	if o == nil || IsNil(o.Timeout) {
		var ret Timeout
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProxyWSConfigDto) GetTimeoutOk() (*Timeout, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *ClientProxyWSConfigDto) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given Timeout and assigns it to the Timeout field.
func (o *ClientProxyWSConfigDto) SetTimeout(v Timeout) {
	o.Timeout = &v
}

// GetFollowRedirects returns the FollowRedirects field value if set, zero value otherwise.
func (o *ClientProxyWSConfigDto) GetFollowRedirects() bool {
	if o == nil || IsNil(o.FollowRedirects) {
		var ret bool
		return ret
	}
	return *o.FollowRedirects
}

// GetFollowRedirectsOk returns a tuple with the FollowRedirects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProxyWSConfigDto) GetFollowRedirectsOk() (*bool, bool) {
	if o == nil || IsNil(o.FollowRedirects) {
		return nil, false
	}
	return o.FollowRedirects, true
}

// HasFollowRedirects returns a boolean if a field has been set.
func (o *ClientProxyWSConfigDto) HasFollowRedirects() bool {
	if o != nil && !IsNil(o.FollowRedirects) {
		return true
	}

	return false
}

// SetFollowRedirects gets a reference to the given bool and assigns it to the FollowRedirects field.
func (o *ClientProxyWSConfigDto) SetFollowRedirects(v bool) {
	o.FollowRedirects = &v
}

// GetUseProxyProperties returns the UseProxyProperties field value if set, zero value otherwise.
func (o *ClientProxyWSConfigDto) GetUseProxyProperties() bool {
	if o == nil || IsNil(o.UseProxyProperties) {
		var ret bool
		return ret
	}
	return *o.UseProxyProperties
}

// GetUseProxyPropertiesOk returns a tuple with the UseProxyProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProxyWSConfigDto) GetUseProxyPropertiesOk() (*bool, bool) {
	if o == nil || IsNil(o.UseProxyProperties) {
		return nil, false
	}
	return o.UseProxyProperties, true
}

// HasUseProxyProperties returns a boolean if a field has been set.
func (o *ClientProxyWSConfigDto) HasUseProxyProperties() bool {
	if o != nil && !IsNil(o.UseProxyProperties) {
		return true
	}

	return false
}

// SetUseProxyProperties gets a reference to the given bool and assigns it to the UseProxyProperties field.
func (o *ClientProxyWSConfigDto) SetUseProxyProperties(v bool) {
	o.UseProxyProperties = &v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise.
func (o *ClientProxyWSConfigDto) GetUserAgent() string {
	if o == nil || IsNil(o.UserAgent) {
		var ret string
		return ret
	}
	return *o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProxyWSConfigDto) GetUserAgentOk() (*string, bool) {
	if o == nil || IsNil(o.UserAgent) {
		return nil, false
	}
	return o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *ClientProxyWSConfigDto) HasUserAgent() bool {
	if o != nil && !IsNil(o.UserAgent) {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given string and assigns it to the UserAgent field.
func (o *ClientProxyWSConfigDto) SetUserAgent(v string) {
	o.UserAgent = &v
}

// GetCompressionEnabled returns the CompressionEnabled field value if set, zero value otherwise.
func (o *ClientProxyWSConfigDto) GetCompressionEnabled() bool {
	if o == nil || IsNil(o.CompressionEnabled) {
		var ret bool
		return ret
	}
	return *o.CompressionEnabled
}

// GetCompressionEnabledOk returns a tuple with the CompressionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProxyWSConfigDto) GetCompressionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CompressionEnabled) {
		return nil, false
	}
	return o.CompressionEnabled, true
}

// HasCompressionEnabled returns a boolean if a field has been set.
func (o *ClientProxyWSConfigDto) HasCompressionEnabled() bool {
	if o != nil && !IsNil(o.CompressionEnabled) {
		return true
	}

	return false
}

// SetCompressionEnabled gets a reference to the given bool and assigns it to the CompressionEnabled field.
func (o *ClientProxyWSConfigDto) SetCompressionEnabled(v bool) {
	o.CompressionEnabled = &v
}

// GetSsl returns the Ssl field value
func (o *ClientProxyWSConfigDto) GetSsl() SSLConfig {
	if o == nil {
		var ret SSLConfig
		return ret
	}

	return o.Ssl
}

// GetSslOk returns a tuple with the Ssl field value
// and a boolean to check if the value has been set.
func (o *ClientProxyWSConfigDto) GetSslOk() (*SSLConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ssl, true
}

// SetSsl sets field value
func (o *ClientProxyWSConfigDto) SetSsl(v SSLConfig) {
	o.Ssl = v
}

// GetMaxConnectionsPerHost returns the MaxConnectionsPerHost field value if set, zero value otherwise.
func (o *ClientProxyWSConfigDto) GetMaxConnectionsPerHost() int32 {
	if o == nil || IsNil(o.MaxConnectionsPerHost) {
		var ret int32
		return ret
	}
	return *o.MaxConnectionsPerHost
}

// GetMaxConnectionsPerHostOk returns a tuple with the MaxConnectionsPerHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProxyWSConfigDto) GetMaxConnectionsPerHostOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxConnectionsPerHost) {
		return nil, false
	}
	return o.MaxConnectionsPerHost, true
}

// HasMaxConnectionsPerHost returns a boolean if a field has been set.
func (o *ClientProxyWSConfigDto) HasMaxConnectionsPerHost() bool {
	if o != nil && !IsNil(o.MaxConnectionsPerHost) {
		return true
	}

	return false
}

// SetMaxConnectionsPerHost gets a reference to the given int32 and assigns it to the MaxConnectionsPerHost field.
func (o *ClientProxyWSConfigDto) SetMaxConnectionsPerHost(v int32) {
	o.MaxConnectionsPerHost = &v
}

// GetMaxConnectionsTotal returns the MaxConnectionsTotal field value if set, zero value otherwise.
func (o *ClientProxyWSConfigDto) GetMaxConnectionsTotal() int32 {
	if o == nil || IsNil(o.MaxConnectionsTotal) {
		var ret int32
		return ret
	}
	return *o.MaxConnectionsTotal
}

// GetMaxConnectionsTotalOk returns a tuple with the MaxConnectionsTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProxyWSConfigDto) GetMaxConnectionsTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxConnectionsTotal) {
		return nil, false
	}
	return o.MaxConnectionsTotal, true
}

// HasMaxConnectionsTotal returns a boolean if a field has been set.
func (o *ClientProxyWSConfigDto) HasMaxConnectionsTotal() bool {
	if o != nil && !IsNil(o.MaxConnectionsTotal) {
		return true
	}

	return false
}

// SetMaxConnectionsTotal gets a reference to the given int32 and assigns it to the MaxConnectionsTotal field.
func (o *ClientProxyWSConfigDto) SetMaxConnectionsTotal(v int32) {
	o.MaxConnectionsTotal = &v
}

// GetMaxConnectionLifetime returns the MaxConnectionLifetime field value if set, zero value otherwise.
func (o *ClientProxyWSConfigDto) GetMaxConnectionLifetime() string {
	if o == nil || IsNil(o.MaxConnectionLifetime) {
		var ret string
		return ret
	}
	return *o.MaxConnectionLifetime
}

// GetMaxConnectionLifetimeOk returns a tuple with the MaxConnectionLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProxyWSConfigDto) GetMaxConnectionLifetimeOk() (*string, bool) {
	if o == nil || IsNil(o.MaxConnectionLifetime) {
		return nil, false
	}
	return o.MaxConnectionLifetime, true
}

// HasMaxConnectionLifetime returns a boolean if a field has been set.
func (o *ClientProxyWSConfigDto) HasMaxConnectionLifetime() bool {
	if o != nil && !IsNil(o.MaxConnectionLifetime) {
		return true
	}

	return false
}

// SetMaxConnectionLifetime gets a reference to the given string and assigns it to the MaxConnectionLifetime field.
func (o *ClientProxyWSConfigDto) SetMaxConnectionLifetime(v string) {
	o.MaxConnectionLifetime = &v
}

// GetIdleConnectionInPoolTimeout returns the IdleConnectionInPoolTimeout field value if set, zero value otherwise.
func (o *ClientProxyWSConfigDto) GetIdleConnectionInPoolTimeout() string {
	if o == nil || IsNil(o.IdleConnectionInPoolTimeout) {
		var ret string
		return ret
	}
	return *o.IdleConnectionInPoolTimeout
}

// GetIdleConnectionInPoolTimeoutOk returns a tuple with the IdleConnectionInPoolTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProxyWSConfigDto) GetIdleConnectionInPoolTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.IdleConnectionInPoolTimeout) {
		return nil, false
	}
	return o.IdleConnectionInPoolTimeout, true
}

// HasIdleConnectionInPoolTimeout returns a boolean if a field has been set.
func (o *ClientProxyWSConfigDto) HasIdleConnectionInPoolTimeout() bool {
	if o != nil && !IsNil(o.IdleConnectionInPoolTimeout) {
		return true
	}

	return false
}

// SetIdleConnectionInPoolTimeout gets a reference to the given string and assigns it to the IdleConnectionInPoolTimeout field.
func (o *ClientProxyWSConfigDto) SetIdleConnectionInPoolTimeout(v string) {
	o.IdleConnectionInPoolTimeout = &v
}

// GetConnectionPoolCleanerPeriod returns the ConnectionPoolCleanerPeriod field value if set, zero value otherwise.
func (o *ClientProxyWSConfigDto) GetConnectionPoolCleanerPeriod() string {
	if o == nil || IsNil(o.ConnectionPoolCleanerPeriod) {
		var ret string
		return ret
	}
	return *o.ConnectionPoolCleanerPeriod
}

// GetConnectionPoolCleanerPeriodOk returns a tuple with the ConnectionPoolCleanerPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProxyWSConfigDto) GetConnectionPoolCleanerPeriodOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionPoolCleanerPeriod) {
		return nil, false
	}
	return o.ConnectionPoolCleanerPeriod, true
}

// HasConnectionPoolCleanerPeriod returns a boolean if a field has been set.
func (o *ClientProxyWSConfigDto) HasConnectionPoolCleanerPeriod() bool {
	if o != nil && !IsNil(o.ConnectionPoolCleanerPeriod) {
		return true
	}

	return false
}

// SetConnectionPoolCleanerPeriod gets a reference to the given string and assigns it to the ConnectionPoolCleanerPeriod field.
func (o *ClientProxyWSConfigDto) SetConnectionPoolCleanerPeriod(v string) {
	o.ConnectionPoolCleanerPeriod = &v
}

// GetMaxNumberOfRedirects returns the MaxNumberOfRedirects field value if set, zero value otherwise.
func (o *ClientProxyWSConfigDto) GetMaxNumberOfRedirects() int32 {
	if o == nil || IsNil(o.MaxNumberOfRedirects) {
		var ret int32
		return ret
	}
	return *o.MaxNumberOfRedirects
}

// GetMaxNumberOfRedirectsOk returns a tuple with the MaxNumberOfRedirects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProxyWSConfigDto) GetMaxNumberOfRedirectsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxNumberOfRedirects) {
		return nil, false
	}
	return o.MaxNumberOfRedirects, true
}

// HasMaxNumberOfRedirects returns a boolean if a field has been set.
func (o *ClientProxyWSConfigDto) HasMaxNumberOfRedirects() bool {
	if o != nil && !IsNil(o.MaxNumberOfRedirects) {
		return true
	}

	return false
}

// SetMaxNumberOfRedirects gets a reference to the given int32 and assigns it to the MaxNumberOfRedirects field.
func (o *ClientProxyWSConfigDto) SetMaxNumberOfRedirects(v int32) {
	o.MaxNumberOfRedirects = &v
}

// GetMaxRequestRetry returns the MaxRequestRetry field value if set, zero value otherwise.
func (o *ClientProxyWSConfigDto) GetMaxRequestRetry() int32 {
	if o == nil || IsNil(o.MaxRequestRetry) {
		var ret int32
		return ret
	}
	return *o.MaxRequestRetry
}

// GetMaxRequestRetryOk returns a tuple with the MaxRequestRetry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProxyWSConfigDto) GetMaxRequestRetryOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxRequestRetry) {
		return nil, false
	}
	return o.MaxRequestRetry, true
}

// HasMaxRequestRetry returns a boolean if a field has been set.
func (o *ClientProxyWSConfigDto) HasMaxRequestRetry() bool {
	if o != nil && !IsNil(o.MaxRequestRetry) {
		return true
	}

	return false
}

// SetMaxRequestRetry gets a reference to the given int32 and assigns it to the MaxRequestRetry field.
func (o *ClientProxyWSConfigDto) SetMaxRequestRetry(v int32) {
	o.MaxRequestRetry = &v
}

// GetDisableUrlEncoding returns the DisableUrlEncoding field value if set, zero value otherwise.
func (o *ClientProxyWSConfigDto) GetDisableUrlEncoding() bool {
	if o == nil || IsNil(o.DisableUrlEncoding) {
		var ret bool
		return ret
	}
	return *o.DisableUrlEncoding
}

// GetDisableUrlEncodingOk returns a tuple with the DisableUrlEncoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProxyWSConfigDto) GetDisableUrlEncodingOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableUrlEncoding) {
		return nil, false
	}
	return o.DisableUrlEncoding, true
}

// HasDisableUrlEncoding returns a boolean if a field has been set.
func (o *ClientProxyWSConfigDto) HasDisableUrlEncoding() bool {
	if o != nil && !IsNil(o.DisableUrlEncoding) {
		return true
	}

	return false
}

// SetDisableUrlEncoding gets a reference to the given bool and assigns it to the DisableUrlEncoding field.
func (o *ClientProxyWSConfigDto) SetDisableUrlEncoding(v bool) {
	o.DisableUrlEncoding = &v
}

// GetKeepAlive returns the KeepAlive field value if set, zero value otherwise.
func (o *ClientProxyWSConfigDto) GetKeepAlive() bool {
	if o == nil || IsNil(o.KeepAlive) {
		var ret bool
		return ret
	}
	return *o.KeepAlive
}

// GetKeepAliveOk returns a tuple with the KeepAlive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProxyWSConfigDto) GetKeepAliveOk() (*bool, bool) {
	if o == nil || IsNil(o.KeepAlive) {
		return nil, false
	}
	return o.KeepAlive, true
}

// HasKeepAlive returns a boolean if a field has been set.
func (o *ClientProxyWSConfigDto) HasKeepAlive() bool {
	if o != nil && !IsNil(o.KeepAlive) {
		return true
	}

	return false
}

// SetKeepAlive gets a reference to the given bool and assigns it to the KeepAlive field.
func (o *ClientProxyWSConfigDto) SetKeepAlive(v bool) {
	o.KeepAlive = &v
}

// GetUseLaxCookieEncoder returns the UseLaxCookieEncoder field value if set, zero value otherwise.
func (o *ClientProxyWSConfigDto) GetUseLaxCookieEncoder() bool {
	if o == nil || IsNil(o.UseLaxCookieEncoder) {
		var ret bool
		return ret
	}
	return *o.UseLaxCookieEncoder
}

// GetUseLaxCookieEncoderOk returns a tuple with the UseLaxCookieEncoder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProxyWSConfigDto) GetUseLaxCookieEncoderOk() (*bool, bool) {
	if o == nil || IsNil(o.UseLaxCookieEncoder) {
		return nil, false
	}
	return o.UseLaxCookieEncoder, true
}

// HasUseLaxCookieEncoder returns a boolean if a field has been set.
func (o *ClientProxyWSConfigDto) HasUseLaxCookieEncoder() bool {
	if o != nil && !IsNil(o.UseLaxCookieEncoder) {
		return true
	}

	return false
}

// SetUseLaxCookieEncoder gets a reference to the given bool and assigns it to the UseLaxCookieEncoder field.
func (o *ClientProxyWSConfigDto) SetUseLaxCookieEncoder(v bool) {
	o.UseLaxCookieEncoder = &v
}

// GetUseCookieStore returns the UseCookieStore field value if set, zero value otherwise.
func (o *ClientProxyWSConfigDto) GetUseCookieStore() bool {
	if o == nil || IsNil(o.UseCookieStore) {
		var ret bool
		return ret
	}
	return *o.UseCookieStore
}

// GetUseCookieStoreOk returns a tuple with the UseCookieStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProxyWSConfigDto) GetUseCookieStoreOk() (*bool, bool) {
	if o == nil || IsNil(o.UseCookieStore) {
		return nil, false
	}
	return o.UseCookieStore, true
}

// HasUseCookieStore returns a boolean if a field has been set.
func (o *ClientProxyWSConfigDto) HasUseCookieStore() bool {
	if o != nil && !IsNil(o.UseCookieStore) {
		return true
	}

	return false
}

// SetUseCookieStore gets a reference to the given bool and assigns it to the UseCookieStore field.
func (o *ClientProxyWSConfigDto) SetUseCookieStore(v bool) {
	o.UseCookieStore = &v
}

// GetProxy returns the Proxy field value
func (o *ClientProxyWSConfigDto) GetProxy() ProxyConfig {
	if o == nil {
		var ret ProxyConfig
		return ret
	}

	return o.Proxy
}

// GetProxyOk returns a tuple with the Proxy field value
// and a boolean to check if the value has been set.
func (o *ClientProxyWSConfigDto) GetProxyOk() (*ProxyConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Proxy, true
}

// SetProxy sets field value
func (o *ClientProxyWSConfigDto) SetProxy(v ProxyConfig) {
	o.Proxy = v
}

func (o ClientProxyWSConfigDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientProxyWSConfigDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	if !IsNil(o.FollowRedirects) {
		toSerialize["followRedirects"] = o.FollowRedirects
	}
	if !IsNil(o.UseProxyProperties) {
		toSerialize["useProxyProperties"] = o.UseProxyProperties
	}
	if !IsNil(o.UserAgent) {
		toSerialize["userAgent"] = o.UserAgent
	}
	if !IsNil(o.CompressionEnabled) {
		toSerialize["compressionEnabled"] = o.CompressionEnabled
	}
	toSerialize["ssl"] = o.Ssl
	if !IsNil(o.MaxConnectionsPerHost) {
		toSerialize["maxConnectionsPerHost"] = o.MaxConnectionsPerHost
	}
	if !IsNil(o.MaxConnectionsTotal) {
		toSerialize["maxConnectionsTotal"] = o.MaxConnectionsTotal
	}
	if !IsNil(o.MaxConnectionLifetime) {
		toSerialize["maxConnectionLifetime"] = o.MaxConnectionLifetime
	}
	if !IsNil(o.IdleConnectionInPoolTimeout) {
		toSerialize["idleConnectionInPoolTimeout"] = o.IdleConnectionInPoolTimeout
	}
	if !IsNil(o.ConnectionPoolCleanerPeriod) {
		toSerialize["connectionPoolCleanerPeriod"] = o.ConnectionPoolCleanerPeriod
	}
	if !IsNil(o.MaxNumberOfRedirects) {
		toSerialize["maxNumberOfRedirects"] = o.MaxNumberOfRedirects
	}
	if !IsNil(o.MaxRequestRetry) {
		toSerialize["maxRequestRetry"] = o.MaxRequestRetry
	}
	if !IsNil(o.DisableUrlEncoding) {
		toSerialize["disableUrlEncoding"] = o.DisableUrlEncoding
	}
	if !IsNil(o.KeepAlive) {
		toSerialize["keepAlive"] = o.KeepAlive
	}
	if !IsNil(o.UseLaxCookieEncoder) {
		toSerialize["useLaxCookieEncoder"] = o.UseLaxCookieEncoder
	}
	if !IsNil(o.UseCookieStore) {
		toSerialize["useCookieStore"] = o.UseCookieStore
	}
	toSerialize["proxy"] = o.Proxy
	return toSerialize, nil
}

func (o *ClientProxyWSConfigDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ssl",
		"proxy",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varClientProxyWSConfigDto := _ClientProxyWSConfigDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// Postprocessed by scripts/tolerant-response-decoder: unknown fields are ignored,
	// since TheHive adds fields to its responses in minor versions.
	err = decoder.Decode(&varClientProxyWSConfigDto)

	if err != nil {
		return err
	}

	*o = ClientProxyWSConfigDto(varClientProxyWSConfigDto)

	return err
}

type NullableClientProxyWSConfigDto struct {
	value *ClientProxyWSConfigDto
	isSet bool
}

func (v NullableClientProxyWSConfigDto) Get() *ClientProxyWSConfigDto {
	return v.value
}

func (v *NullableClientProxyWSConfigDto) Set(val *ClientProxyWSConfigDto) {
	v.value = val
	v.isSet = true
}

func (v NullableClientProxyWSConfigDto) IsSet() bool {
	return v.isSet
}

func (v *NullableClientProxyWSConfigDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientProxyWSConfigDto(val *ClientProxyWSConfigDto) *NullableClientProxyWSConfigDto {
	return &NullableClientProxyWSConfigDto{value: val, isSet: true}
}

func (v NullableClientProxyWSConfigDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientProxyWSConfigDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
