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

// checks if the OutputOrganisation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutputOrganisation{}

// OutputOrganisation struct for OutputOrganisation
type OutputOrganisation struct {
	// Internal unique identifier of the organization.
	UnderscoreId string `json:"_id"`
	// Entity type, always `Organisation`.
	UnderscoreType string `json:"_type"`
	// Login of the user who created the organization.
	UnderscoreCreatedBy string `json:"_createdBy"`
	// Login of the user who last updated the organization.
	UnderscoreUpdatedBy *string `json:"_updatedBy,omitempty"`
	// Date and time when the organization was created. Returned as a Unix timestamp in milliseconds.
	UnderscoreCreatedAt int64 `json:"_createdAt"`
	// Date and time when the organization was last updated. Returned as a Unix timestamp in milliseconds.
	UnderscoreUpdatedAt *int64 `json:"_updatedAt,omitempty"`
	// Unique name of the organization, which can be used instead of `_id`.
	Name string `json:"name"`
	// Description of the organization.
	Description string `json:"description"`
	// Task-sharing rule applied to new case shares with this organization. One of `autoShare`, `manual`.
	TaskRule string `json:"taskRule"`
	// Observable-sharing rule applied to new case shares with this organization. One of `autoShare`, `manual`.
	ObservableRule string `json:"observableRule"`
	// Sharing links configured for this organization.
	Links []OrganisationLink `json:"links,omitempty"`
	// Base64-encoded avatar image of the organization.
	Avatar *string `json:"avatar,omitempty"`
	// Whether the organization is locked. Members of a locked organization can't log in.
	Locked bool `json:"locked"`
	// Extra data fields populated when specific keys are requested in the `extraData` query parameter.
	ExtraData map[string]interface{} `json:"extraData"`
}

type _OutputOrganisation OutputOrganisation

// NewOutputOrganisation instantiates a new OutputOrganisation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputOrganisation(underscoreId string, underscoreType string, underscoreCreatedBy string, underscoreCreatedAt int64, name string, description string, taskRule string, observableRule string, locked bool, extraData map[string]interface{}) *OutputOrganisation {
	this := OutputOrganisation{}
	this.UnderscoreId = underscoreId
	this.UnderscoreType = underscoreType
	this.UnderscoreCreatedBy = underscoreCreatedBy
	this.UnderscoreCreatedAt = underscoreCreatedAt
	this.Name = name
	this.Description = description
	this.TaskRule = taskRule
	this.ObservableRule = observableRule
	this.Locked = locked
	this.ExtraData = extraData
	return &this
}

// NewOutputOrganisationWithDefaults instantiates a new OutputOrganisation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputOrganisationWithDefaults() *OutputOrganisation {
	this := OutputOrganisation{}
	return &this
}

// GetUnderscoreId returns the UnderscoreId field value
func (o *OutputOrganisation) GetUnderscoreId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnderscoreId
}

// GetUnderscoreIdOk returns a tuple with the UnderscoreId field value
// and a boolean to check if the value has been set.
func (o *OutputOrganisation) GetUnderscoreIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreId, true
}

// SetUnderscoreId sets field value
func (o *OutputOrganisation) SetUnderscoreId(v string) {
	o.UnderscoreId = v
}

// GetUnderscoreType returns the UnderscoreType field value
func (o *OutputOrganisation) GetUnderscoreType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnderscoreType
}

// GetUnderscoreTypeOk returns a tuple with the UnderscoreType field value
// and a boolean to check if the value has been set.
func (o *OutputOrganisation) GetUnderscoreTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreType, true
}

// SetUnderscoreType sets field value
func (o *OutputOrganisation) SetUnderscoreType(v string) {
	o.UnderscoreType = v
}

// GetUnderscoreCreatedBy returns the UnderscoreCreatedBy field value
func (o *OutputOrganisation) GetUnderscoreCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnderscoreCreatedBy
}

// GetUnderscoreCreatedByOk returns a tuple with the UnderscoreCreatedBy field value
// and a boolean to check if the value has been set.
func (o *OutputOrganisation) GetUnderscoreCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreCreatedBy, true
}

// SetUnderscoreCreatedBy sets field value
func (o *OutputOrganisation) SetUnderscoreCreatedBy(v string) {
	o.UnderscoreCreatedBy = v
}

// GetUnderscoreUpdatedBy returns the UnderscoreUpdatedBy field value if set, zero value otherwise.
func (o *OutputOrganisation) GetUnderscoreUpdatedBy() string {
	if o == nil || IsNil(o.UnderscoreUpdatedBy) {
		var ret string
		return ret
	}
	return *o.UnderscoreUpdatedBy
}

// GetUnderscoreUpdatedByOk returns a tuple with the UnderscoreUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputOrganisation) GetUnderscoreUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UnderscoreUpdatedBy) {
		return nil, false
	}
	return o.UnderscoreUpdatedBy, true
}

// HasUnderscoreUpdatedBy returns a boolean if a field has been set.
func (o *OutputOrganisation) HasUnderscoreUpdatedBy() bool {
	if o != nil && !IsNil(o.UnderscoreUpdatedBy) {
		return true
	}

	return false
}

// SetUnderscoreUpdatedBy gets a reference to the given string and assigns it to the UnderscoreUpdatedBy field.
func (o *OutputOrganisation) SetUnderscoreUpdatedBy(v string) {
	o.UnderscoreUpdatedBy = &v
}

// GetUnderscoreCreatedAt returns the UnderscoreCreatedAt field value
func (o *OutputOrganisation) GetUnderscoreCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UnderscoreCreatedAt
}

// GetUnderscoreCreatedAtOk returns a tuple with the UnderscoreCreatedAt field value
// and a boolean to check if the value has been set.
func (o *OutputOrganisation) GetUnderscoreCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreCreatedAt, true
}

// SetUnderscoreCreatedAt sets field value
func (o *OutputOrganisation) SetUnderscoreCreatedAt(v int64) {
	o.UnderscoreCreatedAt = v
}

// GetUnderscoreUpdatedAt returns the UnderscoreUpdatedAt field value if set, zero value otherwise.
func (o *OutputOrganisation) GetUnderscoreUpdatedAt() int64 {
	if o == nil || IsNil(o.UnderscoreUpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UnderscoreUpdatedAt
}

// GetUnderscoreUpdatedAtOk returns a tuple with the UnderscoreUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputOrganisation) GetUnderscoreUpdatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.UnderscoreUpdatedAt) {
		return nil, false
	}
	return o.UnderscoreUpdatedAt, true
}

// HasUnderscoreUpdatedAt returns a boolean if a field has been set.
func (o *OutputOrganisation) HasUnderscoreUpdatedAt() bool {
	if o != nil && !IsNil(o.UnderscoreUpdatedAt) {
		return true
	}

	return false
}

// SetUnderscoreUpdatedAt gets a reference to the given int64 and assigns it to the UnderscoreUpdatedAt field.
func (o *OutputOrganisation) SetUnderscoreUpdatedAt(v int64) {
	o.UnderscoreUpdatedAt = &v
}

// GetName returns the Name field value
func (o *OutputOrganisation) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OutputOrganisation) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OutputOrganisation) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *OutputOrganisation) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *OutputOrganisation) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *OutputOrganisation) SetDescription(v string) {
	o.Description = v
}

// GetTaskRule returns the TaskRule field value
func (o *OutputOrganisation) GetTaskRule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaskRule
}

// GetTaskRuleOk returns a tuple with the TaskRule field value
// and a boolean to check if the value has been set.
func (o *OutputOrganisation) GetTaskRuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskRule, true
}

// SetTaskRule sets field value
func (o *OutputOrganisation) SetTaskRule(v string) {
	o.TaskRule = v
}

// GetObservableRule returns the ObservableRule field value
func (o *OutputOrganisation) GetObservableRule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObservableRule
}

// GetObservableRuleOk returns a tuple with the ObservableRule field value
// and a boolean to check if the value has been set.
func (o *OutputOrganisation) GetObservableRuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObservableRule, true
}

// SetObservableRule sets field value
func (o *OutputOrganisation) SetObservableRule(v string) {
	o.ObservableRule = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OutputOrganisation) GetLinks() []OrganisationLink {
	if o == nil || IsNil(o.Links) {
		var ret []OrganisationLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputOrganisation) GetLinksOk() ([]OrganisationLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OutputOrganisation) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []OrganisationLink and assigns it to the Links field.
func (o *OutputOrganisation) SetLinks(v []OrganisationLink) {
	o.Links = v
}

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *OutputOrganisation) GetAvatar() string {
	if o == nil || IsNil(o.Avatar) {
		var ret string
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputOrganisation) GetAvatarOk() (*string, bool) {
	if o == nil || IsNil(o.Avatar) {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *OutputOrganisation) HasAvatar() bool {
	if o != nil && !IsNil(o.Avatar) {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given string and assigns it to the Avatar field.
func (o *OutputOrganisation) SetAvatar(v string) {
	o.Avatar = &v
}

// GetLocked returns the Locked field value
func (o *OutputOrganisation) GetLocked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Locked
}

// GetLockedOk returns a tuple with the Locked field value
// and a boolean to check if the value has been set.
func (o *OutputOrganisation) GetLockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locked, true
}

// SetLocked sets field value
func (o *OutputOrganisation) SetLocked(v bool) {
	o.Locked = v
}

// GetExtraData returns the ExtraData field value
func (o *OutputOrganisation) GetExtraData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ExtraData
}

// GetExtraDataOk returns a tuple with the ExtraData field value
// and a boolean to check if the value has been set.
func (o *OutputOrganisation) GetExtraDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.ExtraData, true
}

// SetExtraData sets field value
func (o *OutputOrganisation) SetExtraData(v map[string]interface{}) {
	o.ExtraData = v
}

func (o OutputOrganisation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutputOrganisation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.UnderscoreId
	toSerialize["_type"] = o.UnderscoreType
	toSerialize["_createdBy"] = o.UnderscoreCreatedBy
	if !IsNil(o.UnderscoreUpdatedBy) {
		toSerialize["_updatedBy"] = o.UnderscoreUpdatedBy
	}
	toSerialize["_createdAt"] = o.UnderscoreCreatedAt
	if !IsNil(o.UnderscoreUpdatedAt) {
		toSerialize["_updatedAt"] = o.UnderscoreUpdatedAt
	}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["taskRule"] = o.TaskRule
	toSerialize["observableRule"] = o.ObservableRule
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Avatar) {
		toSerialize["avatar"] = o.Avatar
	}
	toSerialize["locked"] = o.Locked
	toSerialize["extraData"] = o.ExtraData
	return toSerialize, nil
}

func (o *OutputOrganisation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"_type",
		"_createdBy",
		"_createdAt",
		"name",
		"description",
		"taskRule",
		"observableRule",
		"locked",
		"extraData",
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

	varOutputOrganisation := _OutputOrganisation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// Postprocessed by scripts/tolerant-response-decoder: unknown fields are ignored,
	// since TheHive adds fields to its responses in minor versions.
	err = decoder.Decode(&varOutputOrganisation)

	if err != nil {
		return err
	}

	*o = OutputOrganisation(varOutputOrganisation)

	return err
}

type NullableOutputOrganisation struct {
	value *OutputOrganisation
	isSet bool
}

func (v NullableOutputOrganisation) Get() *OutputOrganisation {
	return v.value
}

func (v *NullableOutputOrganisation) Set(val *OutputOrganisation) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputOrganisation) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputOrganisation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputOrganisation(val *OutputOrganisation) *NullableOutputOrganisation {
	return &NullableOutputOrganisation{value: val, isSet: true}
}

func (v NullableOutputOrganisation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputOrganisation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
