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

// checks if the OAuth2OutputAuth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2OutputAuth{}

// OAuth2OutputAuth struct for OAuth2OutputAuth
type OAuth2OutputAuth struct {
	// Client ID assigned by the OAuth 2.0 provider.
	ClientId  string     `json:"clientId"`
	GrantType *GrantType `json:"grantType,omitempty"`
	// OAuth 2.0 token endpoint.
	TokenUrl string `json:"tokenUrl"`
	// Scopes requested from the authorization server.
	Scope []string `json:"scope,omitempty"`
	// Method used to send client credentials to the token endpoint.
	ClientAuthenticationMethod ClientAuthenticationMethod `json:"clientAuthenticationMethod"`
	// Additional parameters included in the token request.
	TokenParameters []TokenParameter `json:"tokenParameters,omitempty"`
	// Authentication type identifier.
	Type string `json:"type"`
}

type _OAuth2OutputAuth OAuth2OutputAuth

// NewOAuth2OutputAuth instantiates a new OAuth2OutputAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2OutputAuth(clientId string, tokenUrl string, clientAuthenticationMethod ClientAuthenticationMethod, type_ string) *OAuth2OutputAuth {
	this := OAuth2OutputAuth{}
	this.ClientId = clientId
	this.TokenUrl = tokenUrl
	this.ClientAuthenticationMethod = clientAuthenticationMethod
	this.Type = type_
	return &this
}

// NewOAuth2OutputAuthWithDefaults instantiates a new OAuth2OutputAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2OutputAuthWithDefaults() *OAuth2OutputAuth {
	this := OAuth2OutputAuth{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *OAuth2OutputAuth) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *OAuth2OutputAuth) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *OAuth2OutputAuth) SetClientId(v string) {
	o.ClientId = v
}

// GetGrantType returns the GrantType field value if set, zero value otherwise.
func (o *OAuth2OutputAuth) GetGrantType() GrantType {
	if o == nil || IsNil(o.GrantType) {
		var ret GrantType
		return ret
	}
	return *o.GrantType
}

// GetGrantTypeOk returns a tuple with the GrantType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2OutputAuth) GetGrantTypeOk() (*GrantType, bool) {
	if o == nil || IsNil(o.GrantType) {
		return nil, false
	}
	return o.GrantType, true
}

// HasGrantType returns a boolean if a field has been set.
func (o *OAuth2OutputAuth) HasGrantType() bool {
	if o != nil && !IsNil(o.GrantType) {
		return true
	}

	return false
}

// SetGrantType gets a reference to the given GrantType and assigns it to the GrantType field.
func (o *OAuth2OutputAuth) SetGrantType(v GrantType) {
	o.GrantType = &v
}

// GetTokenUrl returns the TokenUrl field value
func (o *OAuth2OutputAuth) GetTokenUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenUrl
}

// GetTokenUrlOk returns a tuple with the TokenUrl field value
// and a boolean to check if the value has been set.
func (o *OAuth2OutputAuth) GetTokenUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenUrl, true
}

// SetTokenUrl sets field value
func (o *OAuth2OutputAuth) SetTokenUrl(v string) {
	o.TokenUrl = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *OAuth2OutputAuth) GetScope() []string {
	if o == nil || IsNil(o.Scope) {
		var ret []string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2OutputAuth) GetScopeOk() ([]string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *OAuth2OutputAuth) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given []string and assigns it to the Scope field.
func (o *OAuth2OutputAuth) SetScope(v []string) {
	o.Scope = v
}

// GetClientAuthenticationMethod returns the ClientAuthenticationMethod field value
func (o *OAuth2OutputAuth) GetClientAuthenticationMethod() ClientAuthenticationMethod {
	if o == nil {
		var ret ClientAuthenticationMethod
		return ret
	}

	return o.ClientAuthenticationMethod
}

// GetClientAuthenticationMethodOk returns a tuple with the ClientAuthenticationMethod field value
// and a boolean to check if the value has been set.
func (o *OAuth2OutputAuth) GetClientAuthenticationMethodOk() (*ClientAuthenticationMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientAuthenticationMethod, true
}

// SetClientAuthenticationMethod sets field value
func (o *OAuth2OutputAuth) SetClientAuthenticationMethod(v ClientAuthenticationMethod) {
	o.ClientAuthenticationMethod = v
}

// GetTokenParameters returns the TokenParameters field value if set, zero value otherwise.
func (o *OAuth2OutputAuth) GetTokenParameters() []TokenParameter {
	if o == nil || IsNil(o.TokenParameters) {
		var ret []TokenParameter
		return ret
	}
	return o.TokenParameters
}

// GetTokenParametersOk returns a tuple with the TokenParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2OutputAuth) GetTokenParametersOk() ([]TokenParameter, bool) {
	if o == nil || IsNil(o.TokenParameters) {
		return nil, false
	}
	return o.TokenParameters, true
}

// HasTokenParameters returns a boolean if a field has been set.
func (o *OAuth2OutputAuth) HasTokenParameters() bool {
	if o != nil && !IsNil(o.TokenParameters) {
		return true
	}

	return false
}

// SetTokenParameters gets a reference to the given []TokenParameter and assigns it to the TokenParameters field.
func (o *OAuth2OutputAuth) SetTokenParameters(v []TokenParameter) {
	o.TokenParameters = v
}

// GetType returns the Type field value
func (o *OAuth2OutputAuth) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OAuth2OutputAuth) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OAuth2OutputAuth) SetType(v string) {
	o.Type = v
}

func (o OAuth2OutputAuth) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2OutputAuth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clientId"] = o.ClientId
	if !IsNil(o.GrantType) {
		toSerialize["grantType"] = o.GrantType
	}
	toSerialize["tokenUrl"] = o.TokenUrl
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	toSerialize["clientAuthenticationMethod"] = o.ClientAuthenticationMethod
	if !IsNil(o.TokenParameters) {
		toSerialize["tokenParameters"] = o.TokenParameters
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *OAuth2OutputAuth) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"clientId",
		"tokenUrl",
		"clientAuthenticationMethod",
		"type",
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

	varOAuth2OutputAuth := _OAuth2OutputAuth{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// Postprocessed by scripts/tolerant-response-decoder: unknown fields are ignored,
	// since TheHive adds fields to its responses in minor versions.
	err = decoder.Decode(&varOAuth2OutputAuth)

	if err != nil {
		return err
	}

	*o = OAuth2OutputAuth(varOAuth2OutputAuth)

	return err
}

type NullableOAuth2OutputAuth struct {
	value *OAuth2OutputAuth
	isSet bool
}

func (v NullableOAuth2OutputAuth) Get() *OAuth2OutputAuth {
	return v.value
}

func (v *NullableOAuth2OutputAuth) Set(val *OAuth2OutputAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2OutputAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2OutputAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2OutputAuth(val *OAuth2OutputAuth) *NullableOAuth2OutputAuth {
	return &NullableOAuth2OutputAuth{value: val, isSet: true}
}

func (v NullableOAuth2OutputAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2OutputAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
