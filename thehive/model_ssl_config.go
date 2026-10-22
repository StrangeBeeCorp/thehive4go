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

// checks if the SSLConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SSLConfig{}

// SSLConfig SSL/TLS configuration.
type SSLConfig struct {
	// Whether to use the default JVM SSL configuration.
	Default *bool `json:"default,omitempty"`
	// SSL/TLS protocol version to use.
	Protocol *string `json:"protocol,omitempty"`
	// Whether to check certificate revocation lists.
	CheckRevocation *bool `json:"checkRevocation,omitempty"`
	// Certificate revocation list URLs to check.
	RevocationLists []string `json:"revocationLists,omitempty"`
	// Override for the platform default cipher suites.
	EnabledCipherSuites []string `json:"enabledCipherSuites,omitempty"`
	// Override for the platform default TLS protocols.
	EnabledProtocols []string `json:"enabledProtocols,omitempty"`
	// Hostname verifier class name.
	HostnameVerifierClass *string `json:"hostnameVerifierClass,omitempty"`
	// SecureRandom implementation class name.
	SecureRandom *string `json:"secureRandom,omitempty"`
	// Trust manager configuration.
	TrustManager *Manager `json:"trustManager,omitempty"`
	// Key manager configuration.
	KeyManager          *Manager        `json:"keyManager,omitempty"`
	SslParametersConfig *SSLParameters  `json:"sslParametersConfig,omitempty"`
	Debug               *SSLDebugConfig `json:"debug,omitempty"`
	Loose               SSLLooseConfig  `json:"loose"`
}

type _SSLConfig SSLConfig

// NewSSLConfig instantiates a new SSLConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSSLConfig(loose SSLLooseConfig) *SSLConfig {
	this := SSLConfig{}
	var default_ bool = false
	this.Default = &default_
	var protocol string = "TLSv1.2"
	this.Protocol = &protocol
	var hostnameVerifierClass string = "NoopHostnameVerifier"
	this.HostnameVerifierClass = &hostnameVerifierClass
	this.Loose = loose
	return &this
}

// NewSSLConfigWithDefaults instantiates a new SSLConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSSLConfigWithDefaults() *SSLConfig {
	this := SSLConfig{}
	var default_ bool = false
	this.Default = &default_
	var protocol string = "TLSv1.2"
	this.Protocol = &protocol
	var hostnameVerifierClass string = "NoopHostnameVerifier"
	this.HostnameVerifierClass = &hostnameVerifierClass
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *SSLConfig) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSLConfig) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *SSLConfig) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *SSLConfig) SetDefault(v bool) {
	o.Default = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *SSLConfig) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSLConfig) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *SSLConfig) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *SSLConfig) SetProtocol(v string) {
	o.Protocol = &v
}

// GetCheckRevocation returns the CheckRevocation field value if set, zero value otherwise.
func (o *SSLConfig) GetCheckRevocation() bool {
	if o == nil || IsNil(o.CheckRevocation) {
		var ret bool
		return ret
	}
	return *o.CheckRevocation
}

// GetCheckRevocationOk returns a tuple with the CheckRevocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSLConfig) GetCheckRevocationOk() (*bool, bool) {
	if o == nil || IsNil(o.CheckRevocation) {
		return nil, false
	}
	return o.CheckRevocation, true
}

// HasCheckRevocation returns a boolean if a field has been set.
func (o *SSLConfig) HasCheckRevocation() bool {
	if o != nil && !IsNil(o.CheckRevocation) {
		return true
	}

	return false
}

// SetCheckRevocation gets a reference to the given bool and assigns it to the CheckRevocation field.
func (o *SSLConfig) SetCheckRevocation(v bool) {
	o.CheckRevocation = &v
}

// GetRevocationLists returns the RevocationLists field value if set, zero value otherwise.
func (o *SSLConfig) GetRevocationLists() []string {
	if o == nil || IsNil(o.RevocationLists) {
		var ret []string
		return ret
	}
	return o.RevocationLists
}

// GetRevocationListsOk returns a tuple with the RevocationLists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSLConfig) GetRevocationListsOk() ([]string, bool) {
	if o == nil || IsNil(o.RevocationLists) {
		return nil, false
	}
	return o.RevocationLists, true
}

// HasRevocationLists returns a boolean if a field has been set.
func (o *SSLConfig) HasRevocationLists() bool {
	if o != nil && !IsNil(o.RevocationLists) {
		return true
	}

	return false
}

// SetRevocationLists gets a reference to the given []string and assigns it to the RevocationLists field.
func (o *SSLConfig) SetRevocationLists(v []string) {
	o.RevocationLists = v
}

// GetEnabledCipherSuites returns the EnabledCipherSuites field value if set, zero value otherwise.
func (o *SSLConfig) GetEnabledCipherSuites() []string {
	if o == nil || IsNil(o.EnabledCipherSuites) {
		var ret []string
		return ret
	}
	return o.EnabledCipherSuites
}

// GetEnabledCipherSuitesOk returns a tuple with the EnabledCipherSuites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSLConfig) GetEnabledCipherSuitesOk() ([]string, bool) {
	if o == nil || IsNil(o.EnabledCipherSuites) {
		return nil, false
	}
	return o.EnabledCipherSuites, true
}

// HasEnabledCipherSuites returns a boolean if a field has been set.
func (o *SSLConfig) HasEnabledCipherSuites() bool {
	if o != nil && !IsNil(o.EnabledCipherSuites) {
		return true
	}

	return false
}

// SetEnabledCipherSuites gets a reference to the given []string and assigns it to the EnabledCipherSuites field.
func (o *SSLConfig) SetEnabledCipherSuites(v []string) {
	o.EnabledCipherSuites = v
}

// GetEnabledProtocols returns the EnabledProtocols field value if set, zero value otherwise.
func (o *SSLConfig) GetEnabledProtocols() []string {
	if o == nil || IsNil(o.EnabledProtocols) {
		var ret []string
		return ret
	}
	return o.EnabledProtocols
}

// GetEnabledProtocolsOk returns a tuple with the EnabledProtocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSLConfig) GetEnabledProtocolsOk() ([]string, bool) {
	if o == nil || IsNil(o.EnabledProtocols) {
		return nil, false
	}
	return o.EnabledProtocols, true
}

// HasEnabledProtocols returns a boolean if a field has been set.
func (o *SSLConfig) HasEnabledProtocols() bool {
	if o != nil && !IsNil(o.EnabledProtocols) {
		return true
	}

	return false
}

// SetEnabledProtocols gets a reference to the given []string and assigns it to the EnabledProtocols field.
func (o *SSLConfig) SetEnabledProtocols(v []string) {
	o.EnabledProtocols = v
}

// GetHostnameVerifierClass returns the HostnameVerifierClass field value if set, zero value otherwise.
func (o *SSLConfig) GetHostnameVerifierClass() string {
	if o == nil || IsNil(o.HostnameVerifierClass) {
		var ret string
		return ret
	}
	return *o.HostnameVerifierClass
}

// GetHostnameVerifierClassOk returns a tuple with the HostnameVerifierClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSLConfig) GetHostnameVerifierClassOk() (*string, bool) {
	if o == nil || IsNil(o.HostnameVerifierClass) {
		return nil, false
	}
	return o.HostnameVerifierClass, true
}

// HasHostnameVerifierClass returns a boolean if a field has been set.
func (o *SSLConfig) HasHostnameVerifierClass() bool {
	if o != nil && !IsNil(o.HostnameVerifierClass) {
		return true
	}

	return false
}

// SetHostnameVerifierClass gets a reference to the given string and assigns it to the HostnameVerifierClass field.
func (o *SSLConfig) SetHostnameVerifierClass(v string) {
	o.HostnameVerifierClass = &v
}

// GetSecureRandom returns the SecureRandom field value if set, zero value otherwise.
func (o *SSLConfig) GetSecureRandom() string {
	if o == nil || IsNil(o.SecureRandom) {
		var ret string
		return ret
	}
	return *o.SecureRandom
}

// GetSecureRandomOk returns a tuple with the SecureRandom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSLConfig) GetSecureRandomOk() (*string, bool) {
	if o == nil || IsNil(o.SecureRandom) {
		return nil, false
	}
	return o.SecureRandom, true
}

// HasSecureRandom returns a boolean if a field has been set.
func (o *SSLConfig) HasSecureRandom() bool {
	if o != nil && !IsNil(o.SecureRandom) {
		return true
	}

	return false
}

// SetSecureRandom gets a reference to the given string and assigns it to the SecureRandom field.
func (o *SSLConfig) SetSecureRandom(v string) {
	o.SecureRandom = &v
}

// GetTrustManager returns the TrustManager field value if set, zero value otherwise.
func (o *SSLConfig) GetTrustManager() Manager {
	if o == nil || IsNil(o.TrustManager) {
		var ret Manager
		return ret
	}
	return *o.TrustManager
}

// GetTrustManagerOk returns a tuple with the TrustManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSLConfig) GetTrustManagerOk() (*Manager, bool) {
	if o == nil || IsNil(o.TrustManager) {
		return nil, false
	}
	return o.TrustManager, true
}

// HasTrustManager returns a boolean if a field has been set.
func (o *SSLConfig) HasTrustManager() bool {
	if o != nil && !IsNil(o.TrustManager) {
		return true
	}

	return false
}

// SetTrustManager gets a reference to the given Manager and assigns it to the TrustManager field.
func (o *SSLConfig) SetTrustManager(v Manager) {
	o.TrustManager = &v
}

// GetKeyManager returns the KeyManager field value if set, zero value otherwise.
func (o *SSLConfig) GetKeyManager() Manager {
	if o == nil || IsNil(o.KeyManager) {
		var ret Manager
		return ret
	}
	return *o.KeyManager
}

// GetKeyManagerOk returns a tuple with the KeyManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSLConfig) GetKeyManagerOk() (*Manager, bool) {
	if o == nil || IsNil(o.KeyManager) {
		return nil, false
	}
	return o.KeyManager, true
}

// HasKeyManager returns a boolean if a field has been set.
func (o *SSLConfig) HasKeyManager() bool {
	if o != nil && !IsNil(o.KeyManager) {
		return true
	}

	return false
}

// SetKeyManager gets a reference to the given Manager and assigns it to the KeyManager field.
func (o *SSLConfig) SetKeyManager(v Manager) {
	o.KeyManager = &v
}

// GetSslParametersConfig returns the SslParametersConfig field value if set, zero value otherwise.
func (o *SSLConfig) GetSslParametersConfig() SSLParameters {
	if o == nil || IsNil(o.SslParametersConfig) {
		var ret SSLParameters
		return ret
	}
	return *o.SslParametersConfig
}

// GetSslParametersConfigOk returns a tuple with the SslParametersConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSLConfig) GetSslParametersConfigOk() (*SSLParameters, bool) {
	if o == nil || IsNil(o.SslParametersConfig) {
		return nil, false
	}
	return o.SslParametersConfig, true
}

// HasSslParametersConfig returns a boolean if a field has been set.
func (o *SSLConfig) HasSslParametersConfig() bool {
	if o != nil && !IsNil(o.SslParametersConfig) {
		return true
	}

	return false
}

// SetSslParametersConfig gets a reference to the given SSLParameters and assigns it to the SslParametersConfig field.
func (o *SSLConfig) SetSslParametersConfig(v SSLParameters) {
	o.SslParametersConfig = &v
}

// GetDebug returns the Debug field value if set, zero value otherwise.
func (o *SSLConfig) GetDebug() SSLDebugConfig {
	if o == nil || IsNil(o.Debug) {
		var ret SSLDebugConfig
		return ret
	}
	return *o.Debug
}

// GetDebugOk returns a tuple with the Debug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSLConfig) GetDebugOk() (*SSLDebugConfig, bool) {
	if o == nil || IsNil(o.Debug) {
		return nil, false
	}
	return o.Debug, true
}

// HasDebug returns a boolean if a field has been set.
func (o *SSLConfig) HasDebug() bool {
	if o != nil && !IsNil(o.Debug) {
		return true
	}

	return false
}

// SetDebug gets a reference to the given SSLDebugConfig and assigns it to the Debug field.
func (o *SSLConfig) SetDebug(v SSLDebugConfig) {
	o.Debug = &v
}

// GetLoose returns the Loose field value
func (o *SSLConfig) GetLoose() SSLLooseConfig {
	if o == nil {
		var ret SSLLooseConfig
		return ret
	}

	return o.Loose
}

// GetLooseOk returns a tuple with the Loose field value
// and a boolean to check if the value has been set.
func (o *SSLConfig) GetLooseOk() (*SSLLooseConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Loose, true
}

// SetLoose sets field value
func (o *SSLConfig) SetLoose(v SSLLooseConfig) {
	o.Loose = v
}

func (o SSLConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SSLConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.CheckRevocation) {
		toSerialize["checkRevocation"] = o.CheckRevocation
	}
	if !IsNil(o.RevocationLists) {
		toSerialize["revocationLists"] = o.RevocationLists
	}
	if !IsNil(o.EnabledCipherSuites) {
		toSerialize["enabledCipherSuites"] = o.EnabledCipherSuites
	}
	if !IsNil(o.EnabledProtocols) {
		toSerialize["enabledProtocols"] = o.EnabledProtocols
	}
	if !IsNil(o.HostnameVerifierClass) {
		toSerialize["hostnameVerifierClass"] = o.HostnameVerifierClass
	}
	if !IsNil(o.SecureRandom) {
		toSerialize["secureRandom"] = o.SecureRandom
	}
	if !IsNil(o.TrustManager) {
		toSerialize["trustManager"] = o.TrustManager
	}
	if !IsNil(o.KeyManager) {
		toSerialize["keyManager"] = o.KeyManager
	}
	if !IsNil(o.SslParametersConfig) {
		toSerialize["sslParametersConfig"] = o.SslParametersConfig
	}
	if !IsNil(o.Debug) {
		toSerialize["debug"] = o.Debug
	}
	toSerialize["loose"] = o.Loose
	return toSerialize, nil
}

func (o *SSLConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"loose",
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

	varSSLConfig := _SSLConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// Postprocessed by scripts/tolerant-response-decoder: unknown fields are ignored,
	// since TheHive adds fields to its responses in minor versions.
	err = decoder.Decode(&varSSLConfig)

	if err != nil {
		return err
	}

	*o = SSLConfig(varSSLConfig)

	return err
}

type NullableSSLConfig struct {
	value *SSLConfig
	isSet bool
}

func (v NullableSSLConfig) Get() *SSLConfig {
	return v.value
}

func (v *NullableSSLConfig) Set(val *SSLConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSSLConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSSLConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSSLConfig(val *SSLConfig) *NullableSSLConfig {
	return &NullableSSLConfig{value: val, isSet: true}
}

func (v NullableSSLConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSSLConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
