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

// checks if the InputCreateAlertStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputCreateAlertStatus{}

// InputCreateAlertStatus struct for InputCreateAlertStatus
type InputCreateAlertStatus struct {
	// Display name of the alert status. Immutable after creation.
	Value string          `json:"value"`
	Stage InputAlertStage `json:"stage"`
	// Position of the status in the UI dropdown list. Lower values appear first.
	Order *int32 `json:"order,omitempty"`
	// Human-readable description of the status.
	Description *string `json:"description,omitempty"`
	// Hex color code used to display the status in the UI.
	Colour *string `json:"colour,omitempty" validate:"regexp=^#[0-9a-fA-F]{6,6}|$"`
	// When `true`, hides the status from the UI dropdown. Hidden statuses remain accessible through the API.
	Hidden *bool `json:"hidden,omitempty"`
}

type _InputCreateAlertStatus InputCreateAlertStatus

// NewInputCreateAlertStatus instantiates a new InputCreateAlertStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputCreateAlertStatus(value string, stage InputAlertStage) *InputCreateAlertStatus {
	this := InputCreateAlertStatus{}
	this.Value = value
	this.Stage = stage
	return &this
}

// NewInputCreateAlertStatusWithDefaults instantiates a new InputCreateAlertStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputCreateAlertStatusWithDefaults() *InputCreateAlertStatus {
	this := InputCreateAlertStatus{}
	return &this
}

// GetValue returns the Value field value
func (o *InputCreateAlertStatus) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *InputCreateAlertStatus) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *InputCreateAlertStatus) SetValue(v string) {
	o.Value = v
}

// GetStage returns the Stage field value
func (o *InputCreateAlertStatus) GetStage() InputAlertStage {
	if o == nil {
		var ret InputAlertStage
		return ret
	}

	return o.Stage
}

// GetStageOk returns a tuple with the Stage field value
// and a boolean to check if the value has been set.
func (o *InputCreateAlertStatus) GetStageOk() (*InputAlertStage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stage, true
}

// SetStage sets field value
func (o *InputCreateAlertStatus) SetStage(v InputAlertStage) {
	o.Stage = v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *InputCreateAlertStatus) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateAlertStatus) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *InputCreateAlertStatus) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *InputCreateAlertStatus) SetOrder(v int32) {
	o.Order = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InputCreateAlertStatus) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateAlertStatus) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InputCreateAlertStatus) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InputCreateAlertStatus) SetDescription(v string) {
	o.Description = &v
}

// GetColour returns the Colour field value if set, zero value otherwise.
func (o *InputCreateAlertStatus) GetColour() string {
	if o == nil || IsNil(o.Colour) {
		var ret string
		return ret
	}
	return *o.Colour
}

// GetColourOk returns a tuple with the Colour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateAlertStatus) GetColourOk() (*string, bool) {
	if o == nil || IsNil(o.Colour) {
		return nil, false
	}
	return o.Colour, true
}

// HasColour returns a boolean if a field has been set.
func (o *InputCreateAlertStatus) HasColour() bool {
	if o != nil && !IsNil(o.Colour) {
		return true
	}

	return false
}

// SetColour gets a reference to the given string and assigns it to the Colour field.
func (o *InputCreateAlertStatus) SetColour(v string) {
	o.Colour = &v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *InputCreateAlertStatus) GetHidden() bool {
	if o == nil || IsNil(o.Hidden) {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateAlertStatus) GetHiddenOk() (*bool, bool) {
	if o == nil || IsNil(o.Hidden) {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *InputCreateAlertStatus) HasHidden() bool {
	if o != nil && !IsNil(o.Hidden) {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *InputCreateAlertStatus) SetHidden(v bool) {
	o.Hidden = &v
}

func (o InputCreateAlertStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputCreateAlertStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	toSerialize["stage"] = o.Stage
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Colour) {
		toSerialize["colour"] = o.Colour
	}
	if !IsNil(o.Hidden) {
		toSerialize["hidden"] = o.Hidden
	}
	return toSerialize, nil
}

func (o *InputCreateAlertStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
		"stage",
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

	varInputCreateAlertStatus := _InputCreateAlertStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInputCreateAlertStatus)

	if err != nil {
		return err
	}

	*o = InputCreateAlertStatus(varInputCreateAlertStatus)

	return err
}

type NullableInputCreateAlertStatus struct {
	value *InputCreateAlertStatus
	isSet bool
}

func (v NullableInputCreateAlertStatus) Get() *InputCreateAlertStatus {
	return v.value
}

func (v *NullableInputCreateAlertStatus) Set(val *InputCreateAlertStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableInputCreateAlertStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableInputCreateAlertStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputCreateAlertStatus(val *InputCreateAlertStatus) *NullableInputCreateAlertStatus {
	return &NullableInputCreateAlertStatus{value: val, isSet: true}
}

func (v NullableInputCreateAlertStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputCreateAlertStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
