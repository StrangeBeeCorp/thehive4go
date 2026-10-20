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

// checks if the InputEmailIntakeConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputEmailIntakeConfig{}

// InputEmailIntakeConfig struct for InputEmailIntakeConfig
type InputEmailIntakeConfig struct {
	// Pre-registered configuration ID for Google Workspace OAuth 2.0 configurations. Obtain this by calling `POST /api/v1/connector/email-intake/config/authorizationCode` before saving the configuration. Omit for all other provider types.
	Id *string `json:"id,omitempty"`
	// Display name for the mailbox configuration.
	Name    string                        `json:"name"`
	Mailbox InputEmailIntakeMailboxConfig `json:"mailbox"`
	// Organizations where incoming emails from this mailbox create alerts. Must match existing organizations. Use `GET /api/v1/organisation` to list available organizations.
	Organisations []string `json:"organisations,omitempty"`
	// Enables or deactivates this specific mailbox configuration. When turned off, emails aren't fetched from this mailbox.
	Enabled *bool `json:"enabled,omitempty"`
	// Creation date of the configuration, as a Unix timestamp in milliseconds. Set automatically when the configuration is saved.
	CreatedAt       *int64                           `json:"createdAt,omitempty"`
	AlertProperties *InputEmailIntakeAlertProperties `json:"alertProperties,omitempty"`
}

type _InputEmailIntakeConfig InputEmailIntakeConfig

// NewInputEmailIntakeConfig instantiates a new InputEmailIntakeConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputEmailIntakeConfig(name string, mailbox InputEmailIntakeMailboxConfig) *InputEmailIntakeConfig {
	this := InputEmailIntakeConfig{}
	this.Name = name
	this.Mailbox = mailbox
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// NewInputEmailIntakeConfigWithDefaults instantiates a new InputEmailIntakeConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputEmailIntakeConfigWithDefaults() *InputEmailIntakeConfig {
	this := InputEmailIntakeConfig{}
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InputEmailIntakeConfig) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputEmailIntakeConfig) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InputEmailIntakeConfig) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InputEmailIntakeConfig) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *InputEmailIntakeConfig) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InputEmailIntakeConfig) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InputEmailIntakeConfig) SetName(v string) {
	o.Name = v
}

// GetMailbox returns the Mailbox field value
func (o *InputEmailIntakeConfig) GetMailbox() InputEmailIntakeMailboxConfig {
	if o == nil {
		var ret InputEmailIntakeMailboxConfig
		return ret
	}

	return o.Mailbox
}

// GetMailboxOk returns a tuple with the Mailbox field value
// and a boolean to check if the value has been set.
func (o *InputEmailIntakeConfig) GetMailboxOk() (*InputEmailIntakeMailboxConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mailbox, true
}

// SetMailbox sets field value
func (o *InputEmailIntakeConfig) SetMailbox(v InputEmailIntakeMailboxConfig) {
	o.Mailbox = v
}

// GetOrganisations returns the Organisations field value if set, zero value otherwise.
func (o *InputEmailIntakeConfig) GetOrganisations() []string {
	if o == nil || IsNil(o.Organisations) {
		var ret []string
		return ret
	}
	return o.Organisations
}

// GetOrganisationsOk returns a tuple with the Organisations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputEmailIntakeConfig) GetOrganisationsOk() ([]string, bool) {
	if o == nil || IsNil(o.Organisations) {
		return nil, false
	}
	return o.Organisations, true
}

// HasOrganisations returns a boolean if a field has been set.
func (o *InputEmailIntakeConfig) HasOrganisations() bool {
	if o != nil && !IsNil(o.Organisations) {
		return true
	}

	return false
}

// SetOrganisations gets a reference to the given []string and assigns it to the Organisations field.
func (o *InputEmailIntakeConfig) SetOrganisations(v []string) {
	o.Organisations = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InputEmailIntakeConfig) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputEmailIntakeConfig) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InputEmailIntakeConfig) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InputEmailIntakeConfig) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InputEmailIntakeConfig) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputEmailIntakeConfig) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InputEmailIntakeConfig) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *InputEmailIntakeConfig) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetAlertProperties returns the AlertProperties field value if set, zero value otherwise.
func (o *InputEmailIntakeConfig) GetAlertProperties() InputEmailIntakeAlertProperties {
	if o == nil || IsNil(o.AlertProperties) {
		var ret InputEmailIntakeAlertProperties
		return ret
	}
	return *o.AlertProperties
}

// GetAlertPropertiesOk returns a tuple with the AlertProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputEmailIntakeConfig) GetAlertPropertiesOk() (*InputEmailIntakeAlertProperties, bool) {
	if o == nil || IsNil(o.AlertProperties) {
		return nil, false
	}
	return o.AlertProperties, true
}

// HasAlertProperties returns a boolean if a field has been set.
func (o *InputEmailIntakeConfig) HasAlertProperties() bool {
	if o != nil && !IsNil(o.AlertProperties) {
		return true
	}

	return false
}

// SetAlertProperties gets a reference to the given InputEmailIntakeAlertProperties and assigns it to the AlertProperties field.
func (o *InputEmailIntakeConfig) SetAlertProperties(v InputEmailIntakeAlertProperties) {
	o.AlertProperties = &v
}

func (o InputEmailIntakeConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputEmailIntakeConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["mailbox"] = o.Mailbox
	if !IsNil(o.Organisations) {
		toSerialize["organisations"] = o.Organisations
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.AlertProperties) {
		toSerialize["alertProperties"] = o.AlertProperties
	}
	return toSerialize, nil
}

func (o *InputEmailIntakeConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"mailbox",
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

	varInputEmailIntakeConfig := _InputEmailIntakeConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInputEmailIntakeConfig)

	if err != nil {
		return err
	}

	*o = InputEmailIntakeConfig(varInputEmailIntakeConfig)

	return err
}

type NullableInputEmailIntakeConfig struct {
	value *InputEmailIntakeConfig
	isSet bool
}

func (v NullableInputEmailIntakeConfig) Get() *InputEmailIntakeConfig {
	return v.value
}

func (v *NullableInputEmailIntakeConfig) Set(val *InputEmailIntakeConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableInputEmailIntakeConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableInputEmailIntakeConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputEmailIntakeConfig(val *InputEmailIntakeConfig) *NullableInputEmailIntakeConfig {
	return &NullableInputEmailIntakeConfig{value: val, isSet: true}
}

func (v NullableInputEmailIntakeConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputEmailIntakeConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
