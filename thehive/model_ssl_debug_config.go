/*
TheHive

 ## Get started  This guide covers everything you need to make your first successful API call to TheHive.  The publicly hosted API documentation reflects the latest version of TheHive. To browse the documentation that matches the version of TheHive you are running, open it from within the application: select the **?** button in the top right corner, then **API Docs**.  The v0 API is deprecated and isn't documented. New integrations should target v1.  ### Prerequisites  You need a running TheHive instance and an unlocked user account.  ### Authentication  Almost all endpoints require authentication.  TheHive supports two authentication methods:  #### API key (recommended)  Use your API key as a bearer token in the `Authorization` header.  Generate a new API key with `POST /api/v1/user/{userId}/key/renew`, or retrieve an existing one with `GET /api/v1/user/{userId}/key`.  If you have a [Normal user account](https://docs.strangebee.com/thehive/user-guides/organization/configure-organization/manage-user-accounts/about-user-accounts/#types), you can also manage your API key in your [user settings](https://docs.strangebee.com/thehive/user-guides/manage-user-settings/#manage-your-api-key).  curl ```bash curl -X <http_method> -H 'Authorization: Bearer <api_key>' \\   https://<thehive_host>/api/v1/<endpoint> ```  python ```python import requests  headers = {\"Authorization\": \"Bearer <api_key>\"} response = requests.<http_method>(\"https://<thehive_host>/api/v1/<endpoint>\", headers=headers) ```  #### Basic authentication  Only available for [Normal user accounts](https://docs.strangebee.com/thehive/user-guides/organization/configure-organization/manage-user-accounts/about-user-accounts/#types).  Pass your login and password using HTTP Basic authentication. The login is typically an email address. If a default domain is configured in the [authentication settings](https://docs.strangebee.com/thehive/administration/authentication/configure-authentication), a plain username is also accepted.  curl ```bash curl -X <http_method> -u <login>:<password> \\   https://<thehive_host>/api/v1/<endpoint> ```  python ```python import requests  response = requests.<http_method>(\"https://<thehive_host>/api/v1/<endpoint>\", auth=(\"<login>\", \"<password>\")) ```  The following endpoints are public and can be called without credentials:  - `POST /api/v1/login`: Authenticate with login, password, and a TOTP code if MFA is enabled. Returns a session cookie used to authenticate subsequent requests. - `GET /api/v1/logout`: Invalidate the current session. - `POST /api/v1/logout`: Invalidate the current session. Same as `GET`, provided for clients that can't send `GET` requests with a body. - `GET /api/v1/auth/local/passwordPolicy`: Retrieve the password policy configured on the server such as minimum length or complexity requirements. - `GET /api/v1/status/public`: Retrieve public server information, including the TheHive version and license status, without exposing sensitive configuration.  All other endpoints return `401 Unauthorized` when called without valid credentials.  ### Your first API call  Start with `GET /api/v1/user/current`: it requires no request body, confirms your credentials work, and returns your user ID.  curl ```bash curl -X GET -H 'Authorization: Bearer <api_key>' \\      https://<thehive_host>/api/v1/user/current ```  python ```python import requests  BASE_URL = \"https://<thehive_host>\" API_KEY = \"<api_key>\"  headers = {\"Authorization\": f\"Bearer {API_KEY}\"} response = requests.get(f\"{BASE_URL}/api/v1/user/current\", headers=headers) user = response.json() ```  ### X-Organisation header  By default, API calls target the user's default organization. The `defaultOrganisation` field returned by `GET /api/v1/user/current` tells you which one it is.  If your account belongs to several organizations and you need to target a specific one, add the `X-Organisation` header. This is also required when performing administration actions, as these must be executed in the context of the `admin` organization.  curl ```bash curl -H 'Authorization: Bearer <api_key>' \\      -H 'X-Organisation: <organization_name>' \\      https://<thehive_host>/api/v1/user/current ```  python ```python import requests  BASE_URL = \"https://<thehive_host>\" API_KEY = \"<api_key>\"  headers = {\"Authorization\": f\"Bearer {API_KEY}\", \"X-Organisation\": \"<organization_name>\"} response = requests.get(f\"{BASE_URL}/api/v1/user/current\", headers=headers) user = response.json() ```  ### Request and response format  All request and response bodies use JSON: `Content-Type: application/json`.  Dates are represented as Unix timestamps in milliseconds.  ### How the API is organized  TheHive exposes two kinds of endpoints.  #### The Query API  `POST /api/v1/query` handles all listing, filtering, sorting, and pagination needs.  A single endpoint accepts a structured body that chains operations:  ```json {   \"query\": [     { \"_name\": \"listCase\" },     { \"_name\": \"filter\", \"_eq\": { \"_field\": \"status\", \"_value\": \"Open\" } },     { \"_name\": \"sort\", \"_fields\": [{ \"_createdAt\": \"desc\" }] },     { \"_name\": \"page\", \"from\": 0, \"to\": 15, \"extraData\": [\"taskStats\", \"shareCount\"] }   ],   \"excludeFields\": [\"description\", \"summary\"] } ```  See the [**Query and Export**](https://docs.strangebee.com/thehive/api-docs/#tag/Query-and-Export) section for the full query reference.  #### Standard REST endpoints  Standard REST endpoints handle write operations and direct entity lookups.  They follow conventional HTTP conventions:  | <nobr>Operation </nobr>| Method | Example | |-----------|--------|---------| | <nobr>Create </nobr>| `POST` | `POST /api/v1/case` | | <nobr>Read </nobr>| `GET` | `GET /api/v1/case/{idOrName}` | | <nobr>Update </nobr>| `PATCH` | `PATCH /api/v1/case/{idOrName}` | | <nobr>Replace / set </nobr>| `PUT` | `PUT /api/v1/case/{caseId}/shares` | | <nobr>Delete </nobr>| `DELETE` | `DELETE /api/v1/case/{idOrName}` |  ### Permissions and capabilities  Each user account has permissions defined by their profile. TheHive checks them on every API call.  Some endpoints are only available with a Gold, Platinum, or One license. When that's the case, the required license tier is listed on the endpoint.  To check which license is active on your instance and what capabilities it includes, call `GET /api/v1/status`.  ### Request size limits  TheHive enforces the following request size limits:  | <nobr>Request type </nobr>| Default limit | Description | |---|---|---| | <nobr>Standard API requests (JSON payloads) </nobr>| 10 MB | Maximum size for request bodies processed in memory | | <nobr>File uploads and attachments </nobr>| 1 GB | Maximum size for streamed request bodies written to disk |  TheHive rejects requests that exceed these limits.  To stay within these limits: - Avoid sending excessively large JSON payloads. - Upload large files as attachments rather than embedding data directly in requests. - Implement rate limiting and additional protections at the reverse proxy level when exposing the API externally.  ### Errors  Every endpoint returns at least the following HTTP status codes:  | <nobr>Code </nobr>| Meaning | |------|---------| | <nobr>`400 Bad Request` </nobr>| Invalid request parameters or body | | <nobr>`404 Not Found` </nobr>| Resource does not exist | | <nobr>`500 Internal Server Error` </nobr>| Unexpected server-side error |  Endpoints that require authentication additionally return:  | <nobr>Code </nobr>| Meaning | |------|---------| | <nobr>`401 Unauthorized` </nobr>| Missing or invalid credentials | | <nobr>`403 Forbidden` </nobr>| Authenticated but insufficient permissions |  Other codes may appear depending on the specific operation. Refer to individual endpoint documentation for the full list.  The standard error response body is:  ```json {   \"type\": \"AuthorizationError\",   \"message\": \"You don't have permission to perform this action\" } ```  A small number of endpoints deviate from this format. These cases are noted in the relevant endpoint descriptions.

API version: v6.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package thehive

import (
	"encoding/json"
)

// checks if the SSLDebugConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SSLDebugConfig{}

// SSLDebugConfig SSL debug logging configuration.
type SSLDebugConfig struct {
	// Enable tracing of SSLContext, SSLEngine, SSLSocketFactory, and key and trust managers.
	All *bool `json:"all,omitempty"`
	// Enable tracing of key manager operations.
	Keymanager *bool `json:"keymanager,omitempty"`
	// Enable tracing of SSLEngine and SSLSocketFactory operations.
	Ssl *bool `json:"ssl,omitempty"`
	// Enable tracing of SSLContext operations.
	Sslctx *bool `json:"sslctx,omitempty"`
	// Enable tracing of trust manager operations.
	Trustmanager *bool `json:"trustmanager,omitempty"`
}

// NewSSLDebugConfig instantiates a new SSLDebugConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSSLDebugConfig() *SSLDebugConfig {
	this := SSLDebugConfig{}
	var all bool = false
	this.All = &all
	var keymanager bool = false
	this.Keymanager = &keymanager
	var ssl bool = false
	this.Ssl = &ssl
	var sslctx bool = false
	this.Sslctx = &sslctx
	var trustmanager bool = false
	this.Trustmanager = &trustmanager
	return &this
}

// NewSSLDebugConfigWithDefaults instantiates a new SSLDebugConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSSLDebugConfigWithDefaults() *SSLDebugConfig {
	this := SSLDebugConfig{}
	var all bool = false
	this.All = &all
	var keymanager bool = false
	this.Keymanager = &keymanager
	var ssl bool = false
	this.Ssl = &ssl
	var sslctx bool = false
	this.Sslctx = &sslctx
	var trustmanager bool = false
	this.Trustmanager = &trustmanager
	return &this
}

// GetAll returns the All field value if set, zero value otherwise.
func (o *SSLDebugConfig) GetAll() bool {
	if o == nil || IsNil(o.All) {
		var ret bool
		return ret
	}
	return *o.All
}

// GetAllOk returns a tuple with the All field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSLDebugConfig) GetAllOk() (*bool, bool) {
	if o == nil || IsNil(o.All) {
		return nil, false
	}
	return o.All, true
}

// HasAll returns a boolean if a field has been set.
func (o *SSLDebugConfig) HasAll() bool {
	if o != nil && !IsNil(o.All) {
		return true
	}

	return false
}

// SetAll gets a reference to the given bool and assigns it to the All field.
func (o *SSLDebugConfig) SetAll(v bool) {
	o.All = &v
}

// GetKeymanager returns the Keymanager field value if set, zero value otherwise.
func (o *SSLDebugConfig) GetKeymanager() bool {
	if o == nil || IsNil(o.Keymanager) {
		var ret bool
		return ret
	}
	return *o.Keymanager
}

// GetKeymanagerOk returns a tuple with the Keymanager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSLDebugConfig) GetKeymanagerOk() (*bool, bool) {
	if o == nil || IsNil(o.Keymanager) {
		return nil, false
	}
	return o.Keymanager, true
}

// HasKeymanager returns a boolean if a field has been set.
func (o *SSLDebugConfig) HasKeymanager() bool {
	if o != nil && !IsNil(o.Keymanager) {
		return true
	}

	return false
}

// SetKeymanager gets a reference to the given bool and assigns it to the Keymanager field.
func (o *SSLDebugConfig) SetKeymanager(v bool) {
	o.Keymanager = &v
}

// GetSsl returns the Ssl field value if set, zero value otherwise.
func (o *SSLDebugConfig) GetSsl() bool {
	if o == nil || IsNil(o.Ssl) {
		var ret bool
		return ret
	}
	return *o.Ssl
}

// GetSslOk returns a tuple with the Ssl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSLDebugConfig) GetSslOk() (*bool, bool) {
	if o == nil || IsNil(o.Ssl) {
		return nil, false
	}
	return o.Ssl, true
}

// HasSsl returns a boolean if a field has been set.
func (o *SSLDebugConfig) HasSsl() bool {
	if o != nil && !IsNil(o.Ssl) {
		return true
	}

	return false
}

// SetSsl gets a reference to the given bool and assigns it to the Ssl field.
func (o *SSLDebugConfig) SetSsl(v bool) {
	o.Ssl = &v
}

// GetSslctx returns the Sslctx field value if set, zero value otherwise.
func (o *SSLDebugConfig) GetSslctx() bool {
	if o == nil || IsNil(o.Sslctx) {
		var ret bool
		return ret
	}
	return *o.Sslctx
}

// GetSslctxOk returns a tuple with the Sslctx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSLDebugConfig) GetSslctxOk() (*bool, bool) {
	if o == nil || IsNil(o.Sslctx) {
		return nil, false
	}
	return o.Sslctx, true
}

// HasSslctx returns a boolean if a field has been set.
func (o *SSLDebugConfig) HasSslctx() bool {
	if o != nil && !IsNil(o.Sslctx) {
		return true
	}

	return false
}

// SetSslctx gets a reference to the given bool and assigns it to the Sslctx field.
func (o *SSLDebugConfig) SetSslctx(v bool) {
	o.Sslctx = &v
}

// GetTrustmanager returns the Trustmanager field value if set, zero value otherwise.
func (o *SSLDebugConfig) GetTrustmanager() bool {
	if o == nil || IsNil(o.Trustmanager) {
		var ret bool
		return ret
	}
	return *o.Trustmanager
}

// GetTrustmanagerOk returns a tuple with the Trustmanager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSLDebugConfig) GetTrustmanagerOk() (*bool, bool) {
	if o == nil || IsNil(o.Trustmanager) {
		return nil, false
	}
	return o.Trustmanager, true
}

// HasTrustmanager returns a boolean if a field has been set.
func (o *SSLDebugConfig) HasTrustmanager() bool {
	if o != nil && !IsNil(o.Trustmanager) {
		return true
	}

	return false
}

// SetTrustmanager gets a reference to the given bool and assigns it to the Trustmanager field.
func (o *SSLDebugConfig) SetTrustmanager(v bool) {
	o.Trustmanager = &v
}

func (o SSLDebugConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SSLDebugConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.All) {
		toSerialize["all"] = o.All
	}
	if !IsNil(o.Keymanager) {
		toSerialize["keymanager"] = o.Keymanager
	}
	if !IsNil(o.Ssl) {
		toSerialize["ssl"] = o.Ssl
	}
	if !IsNil(o.Sslctx) {
		toSerialize["sslctx"] = o.Sslctx
	}
	if !IsNil(o.Trustmanager) {
		toSerialize["trustmanager"] = o.Trustmanager
	}
	return toSerialize, nil
}

type NullableSSLDebugConfig struct {
	value *SSLDebugConfig
	isSet bool
}

func (v NullableSSLDebugConfig) Get() *SSLDebugConfig {
	return v.value
}

func (v *NullableSSLDebugConfig) Set(val *SSLDebugConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSSLDebugConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSSLDebugConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSSLDebugConfig(val *SSLDebugConfig) *NullableSSLDebugConfig {
	return &NullableSSLDebugConfig{value: val, isSet: true}
}

func (v NullableSSLDebugConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSSLDebugConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
