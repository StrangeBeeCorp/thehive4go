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

// checks if the OutputComment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutputComment{}

// OutputComment struct for OutputComment
type OutputComment struct {
	// Unique identifier of the comment.
	UnderscoreId string `json:"_id"`
	// Entity type, always `Comment`.
	UnderscoreType string `json:"_type"`
	// Login of the user who created the comment.
	CreatedBy string `json:"createdBy"`
	// Date and time when the comment was created. Returned as a Unix timestamp in milliseconds.
	CreatedAt int64 `json:"createdAt"`
	// Date and time when the comment was last updated. Returned as a Unix timestamp in milliseconds.
	UpdatedAt *int64 `json:"updatedAt,omitempty"`
	// Login of the user who last updated the comment.
	UpdatedBy *string `json:"updatedBy,omitempty"`
	// Text of the comment.
	Message string `json:"message"`
	// Whether the comment has been edited since it was created.
	IsEdited  bool                   `json:"isEdited"`
	ExtraData map[string]interface{} `json:"extraData"`
	// Whether the comment is visible to external users through TheHive Portal.
	External *bool `json:"external,omitempty"`
}

type _OutputComment OutputComment

// NewOutputComment instantiates a new OutputComment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputComment(underscoreId string, underscoreType string, createdBy string, createdAt int64, message string, isEdited bool, extraData map[string]interface{}) *OutputComment {
	this := OutputComment{}
	this.UnderscoreId = underscoreId
	this.UnderscoreType = underscoreType
	this.CreatedBy = createdBy
	this.CreatedAt = createdAt
	this.Message = message
	this.IsEdited = isEdited
	this.ExtraData = extraData
	return &this
}

// NewOutputCommentWithDefaults instantiates a new OutputComment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputCommentWithDefaults() *OutputComment {
	this := OutputComment{}
	return &this
}

// GetUnderscoreId returns the UnderscoreId field value
func (o *OutputComment) GetUnderscoreId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnderscoreId
}

// GetUnderscoreIdOk returns a tuple with the UnderscoreId field value
// and a boolean to check if the value has been set.
func (o *OutputComment) GetUnderscoreIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreId, true
}

// SetUnderscoreId sets field value
func (o *OutputComment) SetUnderscoreId(v string) {
	o.UnderscoreId = v
}

// GetUnderscoreType returns the UnderscoreType field value
func (o *OutputComment) GetUnderscoreType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnderscoreType
}

// GetUnderscoreTypeOk returns a tuple with the UnderscoreType field value
// and a boolean to check if the value has been set.
func (o *OutputComment) GetUnderscoreTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreType, true
}

// SetUnderscoreType sets field value
func (o *OutputComment) SetUnderscoreType(v string) {
	o.UnderscoreType = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *OutputComment) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *OutputComment) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *OutputComment) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *OutputComment) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OutputComment) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *OutputComment) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *OutputComment) GetUpdatedAt() int64 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputComment) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *OutputComment) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *OutputComment) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *OutputComment) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputComment) GetUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *OutputComment) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *OutputComment) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// GetMessage returns the Message field value
func (o *OutputComment) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *OutputComment) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *OutputComment) SetMessage(v string) {
	o.Message = v
}

// GetIsEdited returns the IsEdited field value
func (o *OutputComment) GetIsEdited() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEdited
}

// GetIsEditedOk returns a tuple with the IsEdited field value
// and a boolean to check if the value has been set.
func (o *OutputComment) GetIsEditedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEdited, true
}

// SetIsEdited sets field value
func (o *OutputComment) SetIsEdited(v bool) {
	o.IsEdited = v
}

// GetExtraData returns the ExtraData field value
func (o *OutputComment) GetExtraData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ExtraData
}

// GetExtraDataOk returns a tuple with the ExtraData field value
// and a boolean to check if the value has been set.
func (o *OutputComment) GetExtraDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.ExtraData, true
}

// SetExtraData sets field value
func (o *OutputComment) SetExtraData(v map[string]interface{}) {
	o.ExtraData = v
}

// GetExternal returns the External field value if set, zero value otherwise.
func (o *OutputComment) GetExternal() bool {
	if o == nil || IsNil(o.External) {
		var ret bool
		return ret
	}
	return *o.External
}

// GetExternalOk returns a tuple with the External field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputComment) GetExternalOk() (*bool, bool) {
	if o == nil || IsNil(o.External) {
		return nil, false
	}
	return o.External, true
}

// HasExternal returns a boolean if a field has been set.
func (o *OutputComment) HasExternal() bool {
	if o != nil && !IsNil(o.External) {
		return true
	}

	return false
}

// SetExternal gets a reference to the given bool and assigns it to the External field.
func (o *OutputComment) SetExternal(v bool) {
	o.External = &v
}

func (o OutputComment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutputComment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.UnderscoreId
	toSerialize["_type"] = o.UnderscoreType
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["createdAt"] = o.CreatedAt
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	toSerialize["message"] = o.Message
	toSerialize["isEdited"] = o.IsEdited
	toSerialize["extraData"] = o.ExtraData
	if !IsNil(o.External) {
		toSerialize["external"] = o.External
	}
	return toSerialize, nil
}

func (o *OutputComment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"_type",
		"createdBy",
		"createdAt",
		"message",
		"isEdited",
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

	varOutputComment := _OutputComment{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// Postprocessed by scripts/tolerant-response-decoder: unknown fields are ignored,
	// since TheHive adds fields to its responses in minor versions.
	err = decoder.Decode(&varOutputComment)

	if err != nil {
		return err
	}

	*o = OutputComment(varOutputComment)

	return err
}

type NullableOutputComment struct {
	value *OutputComment
	isSet bool
}

func (v NullableOutputComment) Get() *OutputComment {
	return v.value
}

func (v *NullableOutputComment) Set(val *OutputComment) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputComment) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputComment(val *OutputComment) *NullableOutputComment {
	return &NullableOutputComment{value: val, isSet: true}
}

func (v NullableOutputComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
