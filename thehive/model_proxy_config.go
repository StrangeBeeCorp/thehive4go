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

// checks if the ProxyConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProxyConfig{}

// ProxyConfig Proxy server configuration. If not set, no proxy is used.
type ProxyConfig struct {
	// Hostname or IP address of the proxy server.
	Host *string `json:"host,omitempty"`
	// Port number of the proxy server.
	Port  *int32 `json:"port,omitempty"`
	State State  `json:"state"`
	// Proxy protocol. Accepts `http` or `https`.
	Protocol *string `json:"protocol,omitempty"`
	// Username to authenticate with the proxy server.
	Principal *string `json:"principal,omitempty"`
	// Password to authenticate with the proxy server.
	Password *string `json:"password,omitempty"`
	// NTLM domain for proxy authentication.
	NtlmDomain *string `json:"ntlmDomain,omitempty"`
	// Character encoding for proxy authentication.
	Encoding *string `json:"encoding,omitempty"`
	// List of hostnames or patterns that bypass the proxy.
	NonProxyHosts []string `json:"nonProxyHosts,omitempty"`
}

type _ProxyConfig ProxyConfig

// NewProxyConfig instantiates a new ProxyConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxyConfig(state State) *ProxyConfig {
	this := ProxyConfig{}
	this.State = state
	return &this
}

// NewProxyConfigWithDefaults instantiates a new ProxyConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxyConfigWithDefaults() *ProxyConfig {
	this := ProxyConfig{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *ProxyConfig) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfig) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *ProxyConfig) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *ProxyConfig) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ProxyConfig) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfig) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ProxyConfig) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *ProxyConfig) SetPort(v int32) {
	o.Port = &v
}

// GetState returns the State field value
func (o *ProxyConfig) GetState() State {
	if o == nil {
		var ret State
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ProxyConfig) GetStateOk() (*State, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ProxyConfig) SetState(v State) {
	o.State = v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *ProxyConfig) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfig) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *ProxyConfig) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *ProxyConfig) SetProtocol(v string) {
	o.Protocol = &v
}

// GetPrincipal returns the Principal field value if set, zero value otherwise.
func (o *ProxyConfig) GetPrincipal() string {
	if o == nil || IsNil(o.Principal) {
		var ret string
		return ret
	}
	return *o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfig) GetPrincipalOk() (*string, bool) {
	if o == nil || IsNil(o.Principal) {
		return nil, false
	}
	return o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *ProxyConfig) HasPrincipal() bool {
	if o != nil && !IsNil(o.Principal) {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given string and assigns it to the Principal field.
func (o *ProxyConfig) SetPrincipal(v string) {
	o.Principal = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ProxyConfig) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfig) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ProxyConfig) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ProxyConfig) SetPassword(v string) {
	o.Password = &v
}

// GetNtlmDomain returns the NtlmDomain field value if set, zero value otherwise.
func (o *ProxyConfig) GetNtlmDomain() string {
	if o == nil || IsNil(o.NtlmDomain) {
		var ret string
		return ret
	}
	return *o.NtlmDomain
}

// GetNtlmDomainOk returns a tuple with the NtlmDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfig) GetNtlmDomainOk() (*string, bool) {
	if o == nil || IsNil(o.NtlmDomain) {
		return nil, false
	}
	return o.NtlmDomain, true
}

// HasNtlmDomain returns a boolean if a field has been set.
func (o *ProxyConfig) HasNtlmDomain() bool {
	if o != nil && !IsNil(o.NtlmDomain) {
		return true
	}

	return false
}

// SetNtlmDomain gets a reference to the given string and assigns it to the NtlmDomain field.
func (o *ProxyConfig) SetNtlmDomain(v string) {
	o.NtlmDomain = &v
}

// GetEncoding returns the Encoding field value if set, zero value otherwise.
func (o *ProxyConfig) GetEncoding() string {
	if o == nil || IsNil(o.Encoding) {
		var ret string
		return ret
	}
	return *o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfig) GetEncodingOk() (*string, bool) {
	if o == nil || IsNil(o.Encoding) {
		return nil, false
	}
	return o.Encoding, true
}

// HasEncoding returns a boolean if a field has been set.
func (o *ProxyConfig) HasEncoding() bool {
	if o != nil && !IsNil(o.Encoding) {
		return true
	}

	return false
}

// SetEncoding gets a reference to the given string and assigns it to the Encoding field.
func (o *ProxyConfig) SetEncoding(v string) {
	o.Encoding = &v
}

// GetNonProxyHosts returns the NonProxyHosts field value if set, zero value otherwise.
func (o *ProxyConfig) GetNonProxyHosts() []string {
	if o == nil || IsNil(o.NonProxyHosts) {
		var ret []string
		return ret
	}
	return o.NonProxyHosts
}

// GetNonProxyHostsOk returns a tuple with the NonProxyHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfig) GetNonProxyHostsOk() ([]string, bool) {
	if o == nil || IsNil(o.NonProxyHosts) {
		return nil, false
	}
	return o.NonProxyHosts, true
}

// HasNonProxyHosts returns a boolean if a field has been set.
func (o *ProxyConfig) HasNonProxyHosts() bool {
	if o != nil && !IsNil(o.NonProxyHosts) {
		return true
	}

	return false
}

// SetNonProxyHosts gets a reference to the given []string and assigns it to the NonProxyHosts field.
func (o *ProxyConfig) SetNonProxyHosts(v []string) {
	o.NonProxyHosts = v
}

func (o ProxyConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProxyConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	toSerialize["state"] = o.State
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.Principal) {
		toSerialize["principal"] = o.Principal
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.NtlmDomain) {
		toSerialize["ntlmDomain"] = o.NtlmDomain
	}
	if !IsNil(o.Encoding) {
		toSerialize["encoding"] = o.Encoding
	}
	if !IsNil(o.NonProxyHosts) {
		toSerialize["nonProxyHosts"] = o.NonProxyHosts
	}
	return toSerialize, nil
}

func (o *ProxyConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"state",
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

	varProxyConfig := _ProxyConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// Postprocessed by scripts/tolerant-response-decoder: unknown fields are ignored,
	// since TheHive adds fields to its responses in minor versions.
	err = decoder.Decode(&varProxyConfig)

	if err != nil {
		return err
	}

	*o = ProxyConfig(varProxyConfig)

	return err
}

type NullableProxyConfig struct {
	value *ProxyConfig
	isSet bool
}

func (v NullableProxyConfig) Get() *ProxyConfig {
	return v.value
}

func (v *NullableProxyConfig) Set(val *ProxyConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableProxyConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableProxyConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxyConfig(val *ProxyConfig) *NullableProxyConfig {
	return &NullableProxyConfig{value: val, isSet: true}
}

func (v NullableProxyConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxyConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
