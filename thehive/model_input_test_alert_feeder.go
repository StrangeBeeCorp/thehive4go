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

// checks if the InputTestAlertFeeder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputTestAlertFeeder{}

// InputTestAlertFeeder struct for InputTestAlertFeeder
type InputTestAlertFeeder struct {
	// Unique name for the alert feeder.
	Name string `json:"name"`
	// Description providing additional context about the alert feeder configuration.
	Description string `json:"description"`
	// HTTP method to use when requesting data from the external system.
	Method Method `json:"method"`
	// Endpoint URL of the external system.
	Url string `json:"url"`
	// Polling interval. The minimum interval is 1 minute.
	Interval Interval `json:"interval"`
	// Optional HTTP request body to send with the request. Used when the HTTP method is `POST`.
	Body *string `json:"body,omitempty"`
	// Optional HTTP headers to include in the request.
	Headers []Header `json:"headers,omitempty"`
	// Whether the alert feeder is active and runs on schedule.
	Enabled *bool `json:"enabled,omitempty"`
	// Optional authentication configuration. Supports `none`, `basic`, `bearer`, `key`, and `oauth2`. Omitting this field or setting it to `none` means no authentication.
	Auth *UpdateAuth `json:"auth,omitempty"`
	// Optional proxy configuration for the HTTP request.
	ProxyConfig *ClientProxyWSConfigDto `json:"proxyConfig,omitempty"`
	// Maximum time to wait for the external system to respond.
	RequestTimeout *Interval `json:"requestTimeout,omitempty"`
	// Maximum response payload size in bytes that the alert feeder accepts from the external system.
	ResponseMaxSize *int64 `json:"responseMaxSize,omitempty"`
}

type _InputTestAlertFeeder InputTestAlertFeeder

// NewInputTestAlertFeeder instantiates a new InputTestAlertFeeder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputTestAlertFeeder(name string, description string, method Method, url string, interval Interval) *InputTestAlertFeeder {
	this := InputTestAlertFeeder{}
	this.Name = name
	this.Description = description
	this.Method = method
	this.Url = url
	this.Interval = interval
	var enabled bool = true
	this.Enabled = &enabled
	var responseMaxSize int64 = 10485760
	this.ResponseMaxSize = &responseMaxSize
	return &this
}

// NewInputTestAlertFeederWithDefaults instantiates a new InputTestAlertFeeder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputTestAlertFeederWithDefaults() *InputTestAlertFeeder {
	this := InputTestAlertFeeder{}
	var enabled bool = true
	this.Enabled = &enabled
	var responseMaxSize int64 = 10485760
	this.ResponseMaxSize = &responseMaxSize
	return &this
}

// GetName returns the Name field value
func (o *InputTestAlertFeeder) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InputTestAlertFeeder) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InputTestAlertFeeder) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *InputTestAlertFeeder) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *InputTestAlertFeeder) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *InputTestAlertFeeder) SetDescription(v string) {
	o.Description = v
}

// GetMethod returns the Method field value
func (o *InputTestAlertFeeder) GetMethod() Method {
	if o == nil {
		var ret Method
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *InputTestAlertFeeder) GetMethodOk() (*Method, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *InputTestAlertFeeder) SetMethod(v Method) {
	o.Method = v
}

// GetUrl returns the Url field value
func (o *InputTestAlertFeeder) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *InputTestAlertFeeder) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *InputTestAlertFeeder) SetUrl(v string) {
	o.Url = v
}

// GetInterval returns the Interval field value
func (o *InputTestAlertFeeder) GetInterval() Interval {
	if o == nil {
		var ret Interval
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *InputTestAlertFeeder) GetIntervalOk() (*Interval, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *InputTestAlertFeeder) SetInterval(v Interval) {
	o.Interval = v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *InputTestAlertFeeder) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTestAlertFeeder) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *InputTestAlertFeeder) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *InputTestAlertFeeder) SetBody(v string) {
	o.Body = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *InputTestAlertFeeder) GetHeaders() []Header {
	if o == nil || IsNil(o.Headers) {
		var ret []Header
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTestAlertFeeder) GetHeadersOk() ([]Header, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *InputTestAlertFeeder) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []Header and assigns it to the Headers field.
func (o *InputTestAlertFeeder) SetHeaders(v []Header) {
	o.Headers = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InputTestAlertFeeder) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTestAlertFeeder) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InputTestAlertFeeder) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InputTestAlertFeeder) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAuth returns the Auth field value if set, zero value otherwise.
func (o *InputTestAlertFeeder) GetAuth() UpdateAuth {
	if o == nil || IsNil(o.Auth) {
		var ret UpdateAuth
		return ret
	}
	return *o.Auth
}

// GetAuthOk returns a tuple with the Auth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTestAlertFeeder) GetAuthOk() (*UpdateAuth, bool) {
	if o == nil || IsNil(o.Auth) {
		return nil, false
	}
	return o.Auth, true
}

// HasAuth returns a boolean if a field has been set.
func (o *InputTestAlertFeeder) HasAuth() bool {
	if o != nil && !IsNil(o.Auth) {
		return true
	}

	return false
}

// SetAuth gets a reference to the given UpdateAuth and assigns it to the Auth field.
func (o *InputTestAlertFeeder) SetAuth(v UpdateAuth) {
	o.Auth = &v
}

// GetProxyConfig returns the ProxyConfig field value if set, zero value otherwise.
func (o *InputTestAlertFeeder) GetProxyConfig() ClientProxyWSConfigDto {
	if o == nil || IsNil(o.ProxyConfig) {
		var ret ClientProxyWSConfigDto
		return ret
	}
	return *o.ProxyConfig
}

// GetProxyConfigOk returns a tuple with the ProxyConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTestAlertFeeder) GetProxyConfigOk() (*ClientProxyWSConfigDto, bool) {
	if o == nil || IsNil(o.ProxyConfig) {
		return nil, false
	}
	return o.ProxyConfig, true
}

// HasProxyConfig returns a boolean if a field has been set.
func (o *InputTestAlertFeeder) HasProxyConfig() bool {
	if o != nil && !IsNil(o.ProxyConfig) {
		return true
	}

	return false
}

// SetProxyConfig gets a reference to the given ClientProxyWSConfigDto and assigns it to the ProxyConfig field.
func (o *InputTestAlertFeeder) SetProxyConfig(v ClientProxyWSConfigDto) {
	o.ProxyConfig = &v
}

// GetRequestTimeout returns the RequestTimeout field value if set, zero value otherwise.
func (o *InputTestAlertFeeder) GetRequestTimeout() Interval {
	if o == nil || IsNil(o.RequestTimeout) {
		var ret Interval
		return ret
	}
	return *o.RequestTimeout
}

// GetRequestTimeoutOk returns a tuple with the RequestTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTestAlertFeeder) GetRequestTimeoutOk() (*Interval, bool) {
	if o == nil || IsNil(o.RequestTimeout) {
		return nil, false
	}
	return o.RequestTimeout, true
}

// HasRequestTimeout returns a boolean if a field has been set.
func (o *InputTestAlertFeeder) HasRequestTimeout() bool {
	if o != nil && !IsNil(o.RequestTimeout) {
		return true
	}

	return false
}

// SetRequestTimeout gets a reference to the given Interval and assigns it to the RequestTimeout field.
func (o *InputTestAlertFeeder) SetRequestTimeout(v Interval) {
	o.RequestTimeout = &v
}

// GetResponseMaxSize returns the ResponseMaxSize field value if set, zero value otherwise.
func (o *InputTestAlertFeeder) GetResponseMaxSize() int64 {
	if o == nil || IsNil(o.ResponseMaxSize) {
		var ret int64
		return ret
	}
	return *o.ResponseMaxSize
}

// GetResponseMaxSizeOk returns a tuple with the ResponseMaxSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTestAlertFeeder) GetResponseMaxSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.ResponseMaxSize) {
		return nil, false
	}
	return o.ResponseMaxSize, true
}

// HasResponseMaxSize returns a boolean if a field has been set.
func (o *InputTestAlertFeeder) HasResponseMaxSize() bool {
	if o != nil && !IsNil(o.ResponseMaxSize) {
		return true
	}

	return false
}

// SetResponseMaxSize gets a reference to the given int64 and assigns it to the ResponseMaxSize field.
func (o *InputTestAlertFeeder) SetResponseMaxSize(v int64) {
	o.ResponseMaxSize = &v
}

func (o InputTestAlertFeeder) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputTestAlertFeeder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["method"] = o.Method
	toSerialize["url"] = o.Url
	toSerialize["interval"] = o.Interval
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Auth) {
		toSerialize["auth"] = o.Auth
	}
	if !IsNil(o.ProxyConfig) {
		toSerialize["proxyConfig"] = o.ProxyConfig
	}
	if !IsNil(o.RequestTimeout) {
		toSerialize["requestTimeout"] = o.RequestTimeout
	}
	if !IsNil(o.ResponseMaxSize) {
		toSerialize["responseMaxSize"] = o.ResponseMaxSize
	}
	return toSerialize, nil
}

func (o *InputTestAlertFeeder) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"description",
		"method",
		"url",
		"interval",
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

	varInputTestAlertFeeder := _InputTestAlertFeeder{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInputTestAlertFeeder)

	if err != nil {
		return err
	}

	*o = InputTestAlertFeeder(varInputTestAlertFeeder)

	return err
}

type NullableInputTestAlertFeeder struct {
	value *InputTestAlertFeeder
	isSet bool
}

func (v NullableInputTestAlertFeeder) Get() *InputTestAlertFeeder {
	return v.value
}

func (v *NullableInputTestAlertFeeder) Set(val *InputTestAlertFeeder) {
	v.value = val
	v.isSet = true
}

func (v NullableInputTestAlertFeeder) IsSet() bool {
	return v.isSet
}

func (v *NullableInputTestAlertFeeder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputTestAlertFeeder(val *InputTestAlertFeeder) *NullableInputTestAlertFeeder {
	return &NullableInputTestAlertFeeder{value: val, isSet: true}
}

func (v NullableInputTestAlertFeeder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputTestAlertFeeder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
