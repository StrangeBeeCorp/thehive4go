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

// checks if the InputPredicate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputPredicate{}

// InputPredicate A predicate is a tag label within the taxonomy namespace, forming a tag of the form `namespace:predicate`.
type InputPredicate struct {
	// Identifier string for this predicate, used in the tag label.
	Value string `json:"value"`
	// Full display label for this predicate, shown instead of the identifier where supported.
	Expanded *string `json:"expanded,omitempty"`
	// Whether only one entry from this predicate can be selected at a time.
	Exclusive *bool `json:"exclusive,omitempty"`
	// Description of what this predicate represents.
	Description *string `json:"description,omitempty"`
	// Hex color code for the predicate tag's visual appearance. Must be a six-digit hex color preceded by `#`.
	Colour *string `json:"colour,omitempty" validate:"regexp=^#[0-9a-fA-F]{6,6}|$"`
}

type _InputPredicate InputPredicate

// NewInputPredicate instantiates a new InputPredicate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputPredicate(value string) *InputPredicate {
	this := InputPredicate{}
	this.Value = value
	return &this
}

// NewInputPredicateWithDefaults instantiates a new InputPredicate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputPredicateWithDefaults() *InputPredicate {
	this := InputPredicate{}
	return &this
}

// GetValue returns the Value field value
func (o *InputPredicate) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *InputPredicate) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *InputPredicate) SetValue(v string) {
	o.Value = v
}

// GetExpanded returns the Expanded field value if set, zero value otherwise.
func (o *InputPredicate) GetExpanded() string {
	if o == nil || IsNil(o.Expanded) {
		var ret string
		return ret
	}
	return *o.Expanded
}

// GetExpandedOk returns a tuple with the Expanded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputPredicate) GetExpandedOk() (*string, bool) {
	if o == nil || IsNil(o.Expanded) {
		return nil, false
	}
	return o.Expanded, true
}

// HasExpanded returns a boolean if a field has been set.
func (o *InputPredicate) HasExpanded() bool {
	if o != nil && !IsNil(o.Expanded) {
		return true
	}

	return false
}

// SetExpanded gets a reference to the given string and assigns it to the Expanded field.
func (o *InputPredicate) SetExpanded(v string) {
	o.Expanded = &v
}

// GetExclusive returns the Exclusive field value if set, zero value otherwise.
func (o *InputPredicate) GetExclusive() bool {
	if o == nil || IsNil(o.Exclusive) {
		var ret bool
		return ret
	}
	return *o.Exclusive
}

// GetExclusiveOk returns a tuple with the Exclusive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputPredicate) GetExclusiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Exclusive) {
		return nil, false
	}
	return o.Exclusive, true
}

// HasExclusive returns a boolean if a field has been set.
func (o *InputPredicate) HasExclusive() bool {
	if o != nil && !IsNil(o.Exclusive) {
		return true
	}

	return false
}

// SetExclusive gets a reference to the given bool and assigns it to the Exclusive field.
func (o *InputPredicate) SetExclusive(v bool) {
	o.Exclusive = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InputPredicate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputPredicate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InputPredicate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InputPredicate) SetDescription(v string) {
	o.Description = &v
}

// GetColour returns the Colour field value if set, zero value otherwise.
func (o *InputPredicate) GetColour() string {
	if o == nil || IsNil(o.Colour) {
		var ret string
		return ret
	}
	return *o.Colour
}

// GetColourOk returns a tuple with the Colour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputPredicate) GetColourOk() (*string, bool) {
	if o == nil || IsNil(o.Colour) {
		return nil, false
	}
	return o.Colour, true
}

// HasColour returns a boolean if a field has been set.
func (o *InputPredicate) HasColour() bool {
	if o != nil && !IsNil(o.Colour) {
		return true
	}

	return false
}

// SetColour gets a reference to the given string and assigns it to the Colour field.
func (o *InputPredicate) SetColour(v string) {
	o.Colour = &v
}

func (o InputPredicate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputPredicate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	if !IsNil(o.Expanded) {
		toSerialize["expanded"] = o.Expanded
	}
	if !IsNil(o.Exclusive) {
		toSerialize["exclusive"] = o.Exclusive
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Colour) {
		toSerialize["colour"] = o.Colour
	}
	return toSerialize, nil
}

func (o *InputPredicate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
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

	varInputPredicate := _InputPredicate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInputPredicate)

	if err != nil {
		return err
	}

	*o = InputPredicate(varInputPredicate)

	return err
}

type NullableInputPredicate struct {
	value *InputPredicate
	isSet bool
}

func (v NullableInputPredicate) Get() *InputPredicate {
	return v.value
}

func (v *NullableInputPredicate) Set(val *InputPredicate) {
	v.value = val
	v.isSet = true
}

func (v NullableInputPredicate) IsSet() bool {
	return v.isSet
}

func (v *NullableInputPredicate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputPredicate(val *InputPredicate) *NullableInputPredicate {
	return &NullableInputPredicate{value: val, isSet: true}
}

func (v NullableInputPredicate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputPredicate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
