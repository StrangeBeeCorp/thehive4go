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

// checks if the InputEmailIntakeOAuth2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputEmailIntakeOAuth2{}

// InputEmailIntakeOAuth2 OAuth 2.0 credentials. Required for Google Workspace, Microsoft 365, and Microsoft Graph API providers. Requires a Platinum or One license.
type InputEmailIntakeOAuth2 struct {
	// OAuth 2.0 client ID from the identity provider.
	ClientId string `json:"clientId"`
	// Tenant ID from Microsoft Entra ID. Required for Microsoft 365 and Microsoft Graph API providers.
	TenantId *string `json:"tenantId,omitempty"`
	// OAuth 2.0 client secret from the identity provider.
	Secret string `json:"secret"`
	// OAuth 2.0 authority URL. Defaults to the Microsoft identity platform endpoint when using Microsoft providers.
	Authority *string `json:"authority,omitempty"`
	// OAuth 2.0 permission scopes required by the application.
	Scopes []string `json:"scopes,omitempty"`
	// Redirect URI registered in the OAuth 2.0 application. Required for authorization code flows such as Google Workspace.
	RedirectUri *string `json:"redirectUri,omitempty"`
	// Authorization code from the OAuth 2.0 consent flow. Used only for Google Workspace configurations. Obtain this by calling `POST /api/v1/connector/email-intake/config/authorizationCode`.
	AuthorizationCode *string `json:"authorizationCode,omitempty"`
}

type _InputEmailIntakeOAuth2 InputEmailIntakeOAuth2

// NewInputEmailIntakeOAuth2 instantiates a new InputEmailIntakeOAuth2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputEmailIntakeOAuth2(clientId string, secret string) *InputEmailIntakeOAuth2 {
	this := InputEmailIntakeOAuth2{}
	this.ClientId = clientId
	this.Secret = secret
	return &this
}

// NewInputEmailIntakeOAuth2WithDefaults instantiates a new InputEmailIntakeOAuth2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputEmailIntakeOAuth2WithDefaults() *InputEmailIntakeOAuth2 {
	this := InputEmailIntakeOAuth2{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *InputEmailIntakeOAuth2) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *InputEmailIntakeOAuth2) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *InputEmailIntakeOAuth2) SetClientId(v string) {
	o.ClientId = v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *InputEmailIntakeOAuth2) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputEmailIntakeOAuth2) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *InputEmailIntakeOAuth2) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *InputEmailIntakeOAuth2) SetTenantId(v string) {
	o.TenantId = &v
}

// GetSecret returns the Secret field value
func (o *InputEmailIntakeOAuth2) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *InputEmailIntakeOAuth2) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *InputEmailIntakeOAuth2) SetSecret(v string) {
	o.Secret = v
}

// GetAuthority returns the Authority field value if set, zero value otherwise.
func (o *InputEmailIntakeOAuth2) GetAuthority() string {
	if o == nil || IsNil(o.Authority) {
		var ret string
		return ret
	}
	return *o.Authority
}

// GetAuthorityOk returns a tuple with the Authority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputEmailIntakeOAuth2) GetAuthorityOk() (*string, bool) {
	if o == nil || IsNil(o.Authority) {
		return nil, false
	}
	return o.Authority, true
}

// HasAuthority returns a boolean if a field has been set.
func (o *InputEmailIntakeOAuth2) HasAuthority() bool {
	if o != nil && !IsNil(o.Authority) {
		return true
	}

	return false
}

// SetAuthority gets a reference to the given string and assigns it to the Authority field.
func (o *InputEmailIntakeOAuth2) SetAuthority(v string) {
	o.Authority = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *InputEmailIntakeOAuth2) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputEmailIntakeOAuth2) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *InputEmailIntakeOAuth2) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *InputEmailIntakeOAuth2) SetScopes(v []string) {
	o.Scopes = v
}

// GetRedirectUri returns the RedirectUri field value if set, zero value otherwise.
func (o *InputEmailIntakeOAuth2) GetRedirectUri() string {
	if o == nil || IsNil(o.RedirectUri) {
		var ret string
		return ret
	}
	return *o.RedirectUri
}

// GetRedirectUriOk returns a tuple with the RedirectUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputEmailIntakeOAuth2) GetRedirectUriOk() (*string, bool) {
	if o == nil || IsNil(o.RedirectUri) {
		return nil, false
	}
	return o.RedirectUri, true
}

// HasRedirectUri returns a boolean if a field has been set.
func (o *InputEmailIntakeOAuth2) HasRedirectUri() bool {
	if o != nil && !IsNil(o.RedirectUri) {
		return true
	}

	return false
}

// SetRedirectUri gets a reference to the given string and assigns it to the RedirectUri field.
func (o *InputEmailIntakeOAuth2) SetRedirectUri(v string) {
	o.RedirectUri = &v
}

// GetAuthorizationCode returns the AuthorizationCode field value if set, zero value otherwise.
func (o *InputEmailIntakeOAuth2) GetAuthorizationCode() string {
	if o == nil || IsNil(o.AuthorizationCode) {
		var ret string
		return ret
	}
	return *o.AuthorizationCode
}

// GetAuthorizationCodeOk returns a tuple with the AuthorizationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputEmailIntakeOAuth2) GetAuthorizationCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorizationCode) {
		return nil, false
	}
	return o.AuthorizationCode, true
}

// HasAuthorizationCode returns a boolean if a field has been set.
func (o *InputEmailIntakeOAuth2) HasAuthorizationCode() bool {
	if o != nil && !IsNil(o.AuthorizationCode) {
		return true
	}

	return false
}

// SetAuthorizationCode gets a reference to the given string and assigns it to the AuthorizationCode field.
func (o *InputEmailIntakeOAuth2) SetAuthorizationCode(v string) {
	o.AuthorizationCode = &v
}

func (o InputEmailIntakeOAuth2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputEmailIntakeOAuth2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clientId"] = o.ClientId
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	toSerialize["secret"] = o.Secret
	if !IsNil(o.Authority) {
		toSerialize["authority"] = o.Authority
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.RedirectUri) {
		toSerialize["redirectUri"] = o.RedirectUri
	}
	if !IsNil(o.AuthorizationCode) {
		toSerialize["authorizationCode"] = o.AuthorizationCode
	}
	return toSerialize, nil
}

func (o *InputEmailIntakeOAuth2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"clientId",
		"secret",
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

	varInputEmailIntakeOAuth2 := _InputEmailIntakeOAuth2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInputEmailIntakeOAuth2)

	if err != nil {
		return err
	}

	*o = InputEmailIntakeOAuth2(varInputEmailIntakeOAuth2)

	return err
}

type NullableInputEmailIntakeOAuth2 struct {
	value *InputEmailIntakeOAuth2
	isSet bool
}

func (v NullableInputEmailIntakeOAuth2) Get() *InputEmailIntakeOAuth2 {
	return v.value
}

func (v *NullableInputEmailIntakeOAuth2) Set(val *InputEmailIntakeOAuth2) {
	v.value = val
	v.isSet = true
}

func (v NullableInputEmailIntakeOAuth2) IsSet() bool {
	return v.isSet
}

func (v *NullableInputEmailIntakeOAuth2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputEmailIntakeOAuth2(val *InputEmailIntakeOAuth2) *NullableInputEmailIntakeOAuth2 {
	return &NullableInputEmailIntakeOAuth2{value: val, isSet: true}
}

func (v NullableInputEmailIntakeOAuth2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputEmailIntakeOAuth2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
