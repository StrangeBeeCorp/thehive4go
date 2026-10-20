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

// checks if the InputUpdateObservable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputUpdateObservable{}

// InputUpdateObservable struct for InputUpdateObservable
type InputUpdateObservable struct {
	// New observable type. Must match an existing observable type defined in TheHive. Use `POST /api/v1/query` with `listObservableType` to list available types.
	DataType *string `json:"dataType,omitempty"`
	// New description of the observable. Supports [TheHive-flavored Markdown](https://docs.strangebee.com/thehive/user-guides/thehive-flavored-markdown/). Send `null` to unset.
	Message *string `json:"message,omitempty"`
	// New [Traffic Light Protocol (TLP)](https://www.misp-project.org/taxonomies.html#_tlp) level: 0 (Clear), 1 (Green), 2 (Amber), 3 (Amber+Strict), 4 (Red).
	Tlp *int32 `json:"tlp,omitempty"`
	// New [Permissible Actions Protocol (PAP)](https://www.misp-project.org/taxonomies.html#_pap) level: 0 (Clear), 1 (Green), 2 (Amber), 3 (Red).
	Pap *int32 `json:"pap,omitempty"`
	// Replaces all current tags with this set. Unknown tags are automatically created. Use `addTags` and `removeTags` to add or remove individual tags without replacing the full set. Use `POST /api/v1/query` with `listTag` to retrieve existing custom tags, or with `listTaxonomy` and `tags` to retrieve taxonomy tags.
	Tags []string `json:"tags,omitempty"`
	// Marks the observable as an indicator of compromise (IOC), indicating it's linked to suspicious or malicious activity.
	Ioc *bool `json:"ioc,omitempty"`
	// Marks the observable as sighted, meaning it has been detected or observed in the environment.
	Sighted *bool `json:"sighted,omitempty"`
	// Date and time when the observable was last sighted. Expected format: Unix timestamp in milliseconds. Send `null` to unset.
	SightedAt *int64 `json:"sightedAt,omitempty"`
	// When `true`, excludes the observable from similarity checks used to identify related alerts and cases.
	IgnoreSimilarity *bool `json:"ignoreSimilarity,omitempty"`
	// Tags to add to the current set of observable tags. Unknown tags are automatically created. Use `POST /api/v1/query` with `listTag` to retrieve existing custom tags, or with `listTaxonomy` and `tags` to retrieve taxonomy tags.
	AddTags []string `json:"addTags,omitempty"`
	// Tags to remove from the current set of observable tags. Tags not present on the observable are silently ignored. Use `GET /api/v1/observable/{observableId}` and read the `tags` field to see the current tags on the observable.
	RemoveTags []string `json:"removeTags,omitempty"`
	// When `true`, makes the observable accessible to external users with [TheHive Portal](https://docs.strangebee.com/thehive/administration/thehive-portal/about-thehive-portal/). Platinum or One license required.
	External *bool `json:"external,omitempty"`
}

// NewInputUpdateObservable instantiates a new InputUpdateObservable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputUpdateObservable() *InputUpdateObservable {
	this := InputUpdateObservable{}
	return &this
}

// NewInputUpdateObservableWithDefaults instantiates a new InputUpdateObservable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputUpdateObservableWithDefaults() *InputUpdateObservable {
	this := InputUpdateObservable{}
	return &this
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *InputUpdateObservable) GetDataType() string {
	if o == nil || IsNil(o.DataType) {
		var ret string
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateObservable) GetDataTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DataType) {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *InputUpdateObservable) HasDataType() bool {
	if o != nil && !IsNil(o.DataType) {
		return true
	}

	return false
}

// SetDataType gets a reference to the given string and assigns it to the DataType field.
func (o *InputUpdateObservable) SetDataType(v string) {
	o.DataType = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *InputUpdateObservable) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateObservable) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InputUpdateObservable) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *InputUpdateObservable) SetMessage(v string) {
	o.Message = &v
}

// GetTlp returns the Tlp field value if set, zero value otherwise.
func (o *InputUpdateObservable) GetTlp() int32 {
	if o == nil || IsNil(o.Tlp) {
		var ret int32
		return ret
	}
	return *o.Tlp
}

// GetTlpOk returns a tuple with the Tlp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateObservable) GetTlpOk() (*int32, bool) {
	if o == nil || IsNil(o.Tlp) {
		return nil, false
	}
	return o.Tlp, true
}

// HasTlp returns a boolean if a field has been set.
func (o *InputUpdateObservable) HasTlp() bool {
	if o != nil && !IsNil(o.Tlp) {
		return true
	}

	return false
}

// SetTlp gets a reference to the given int32 and assigns it to the Tlp field.
func (o *InputUpdateObservable) SetTlp(v int32) {
	o.Tlp = &v
}

// GetPap returns the Pap field value if set, zero value otherwise.
func (o *InputUpdateObservable) GetPap() int32 {
	if o == nil || IsNil(o.Pap) {
		var ret int32
		return ret
	}
	return *o.Pap
}

// GetPapOk returns a tuple with the Pap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateObservable) GetPapOk() (*int32, bool) {
	if o == nil || IsNil(o.Pap) {
		return nil, false
	}
	return o.Pap, true
}

// HasPap returns a boolean if a field has been set.
func (o *InputUpdateObservable) HasPap() bool {
	if o != nil && !IsNil(o.Pap) {
		return true
	}

	return false
}

// SetPap gets a reference to the given int32 and assigns it to the Pap field.
func (o *InputUpdateObservable) SetPap(v int32) {
	o.Pap = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InputUpdateObservable) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateObservable) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InputUpdateObservable) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InputUpdateObservable) SetTags(v []string) {
	o.Tags = v
}

// GetIoc returns the Ioc field value if set, zero value otherwise.
func (o *InputUpdateObservable) GetIoc() bool {
	if o == nil || IsNil(o.Ioc) {
		var ret bool
		return ret
	}
	return *o.Ioc
}

// GetIocOk returns a tuple with the Ioc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateObservable) GetIocOk() (*bool, bool) {
	if o == nil || IsNil(o.Ioc) {
		return nil, false
	}
	return o.Ioc, true
}

// HasIoc returns a boolean if a field has been set.
func (o *InputUpdateObservable) HasIoc() bool {
	if o != nil && !IsNil(o.Ioc) {
		return true
	}

	return false
}

// SetIoc gets a reference to the given bool and assigns it to the Ioc field.
func (o *InputUpdateObservable) SetIoc(v bool) {
	o.Ioc = &v
}

// GetSighted returns the Sighted field value if set, zero value otherwise.
func (o *InputUpdateObservable) GetSighted() bool {
	if o == nil || IsNil(o.Sighted) {
		var ret bool
		return ret
	}
	return *o.Sighted
}

// GetSightedOk returns a tuple with the Sighted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateObservable) GetSightedOk() (*bool, bool) {
	if o == nil || IsNil(o.Sighted) {
		return nil, false
	}
	return o.Sighted, true
}

// HasSighted returns a boolean if a field has been set.
func (o *InputUpdateObservable) HasSighted() bool {
	if o != nil && !IsNil(o.Sighted) {
		return true
	}

	return false
}

// SetSighted gets a reference to the given bool and assigns it to the Sighted field.
func (o *InputUpdateObservable) SetSighted(v bool) {
	o.Sighted = &v
}

// GetSightedAt returns the SightedAt field value if set, zero value otherwise.
func (o *InputUpdateObservable) GetSightedAt() int64 {
	if o == nil || IsNil(o.SightedAt) {
		var ret int64
		return ret
	}
	return *o.SightedAt
}

// GetSightedAtOk returns a tuple with the SightedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateObservable) GetSightedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.SightedAt) {
		return nil, false
	}
	return o.SightedAt, true
}

// HasSightedAt returns a boolean if a field has been set.
func (o *InputUpdateObservable) HasSightedAt() bool {
	if o != nil && !IsNil(o.SightedAt) {
		return true
	}

	return false
}

// SetSightedAt gets a reference to the given int64 and assigns it to the SightedAt field.
func (o *InputUpdateObservable) SetSightedAt(v int64) {
	o.SightedAt = &v
}

// GetIgnoreSimilarity returns the IgnoreSimilarity field value if set, zero value otherwise.
func (o *InputUpdateObservable) GetIgnoreSimilarity() bool {
	if o == nil || IsNil(o.IgnoreSimilarity) {
		var ret bool
		return ret
	}
	return *o.IgnoreSimilarity
}

// GetIgnoreSimilarityOk returns a tuple with the IgnoreSimilarity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateObservable) GetIgnoreSimilarityOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreSimilarity) {
		return nil, false
	}
	return o.IgnoreSimilarity, true
}

// HasIgnoreSimilarity returns a boolean if a field has been set.
func (o *InputUpdateObservable) HasIgnoreSimilarity() bool {
	if o != nil && !IsNil(o.IgnoreSimilarity) {
		return true
	}

	return false
}

// SetIgnoreSimilarity gets a reference to the given bool and assigns it to the IgnoreSimilarity field.
func (o *InputUpdateObservable) SetIgnoreSimilarity(v bool) {
	o.IgnoreSimilarity = &v
}

// GetAddTags returns the AddTags field value if set, zero value otherwise.
func (o *InputUpdateObservable) GetAddTags() []string {
	if o == nil || IsNil(o.AddTags) {
		var ret []string
		return ret
	}
	return o.AddTags
}

// GetAddTagsOk returns a tuple with the AddTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateObservable) GetAddTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.AddTags) {
		return nil, false
	}
	return o.AddTags, true
}

// HasAddTags returns a boolean if a field has been set.
func (o *InputUpdateObservable) HasAddTags() bool {
	if o != nil && !IsNil(o.AddTags) {
		return true
	}

	return false
}

// SetAddTags gets a reference to the given []string and assigns it to the AddTags field.
func (o *InputUpdateObservable) SetAddTags(v []string) {
	o.AddTags = v
}

// GetRemoveTags returns the RemoveTags field value if set, zero value otherwise.
func (o *InputUpdateObservable) GetRemoveTags() []string {
	if o == nil || IsNil(o.RemoveTags) {
		var ret []string
		return ret
	}
	return o.RemoveTags
}

// GetRemoveTagsOk returns a tuple with the RemoveTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateObservable) GetRemoveTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.RemoveTags) {
		return nil, false
	}
	return o.RemoveTags, true
}

// HasRemoveTags returns a boolean if a field has been set.
func (o *InputUpdateObservable) HasRemoveTags() bool {
	if o != nil && !IsNil(o.RemoveTags) {
		return true
	}

	return false
}

// SetRemoveTags gets a reference to the given []string and assigns it to the RemoveTags field.
func (o *InputUpdateObservable) SetRemoveTags(v []string) {
	o.RemoveTags = v
}

// GetExternal returns the External field value if set, zero value otherwise.
func (o *InputUpdateObservable) GetExternal() bool {
	if o == nil || IsNil(o.External) {
		var ret bool
		return ret
	}
	return *o.External
}

// GetExternalOk returns a tuple with the External field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateObservable) GetExternalOk() (*bool, bool) {
	if o == nil || IsNil(o.External) {
		return nil, false
	}
	return o.External, true
}

// HasExternal returns a boolean if a field has been set.
func (o *InputUpdateObservable) HasExternal() bool {
	if o != nil && !IsNil(o.External) {
		return true
	}

	return false
}

// SetExternal gets a reference to the given bool and assigns it to the External field.
func (o *InputUpdateObservable) SetExternal(v bool) {
	o.External = &v
}

func (o InputUpdateObservable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputUpdateObservable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataType) {
		toSerialize["dataType"] = o.DataType
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Tlp) {
		toSerialize["tlp"] = o.Tlp
	}
	if !IsNil(o.Pap) {
		toSerialize["pap"] = o.Pap
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Ioc) {
		toSerialize["ioc"] = o.Ioc
	}
	if !IsNil(o.Sighted) {
		toSerialize["sighted"] = o.Sighted
	}
	if !IsNil(o.SightedAt) {
		toSerialize["sightedAt"] = o.SightedAt
	}
	if !IsNil(o.IgnoreSimilarity) {
		toSerialize["ignoreSimilarity"] = o.IgnoreSimilarity
	}
	if !IsNil(o.AddTags) {
		toSerialize["addTags"] = o.AddTags
	}
	if !IsNil(o.RemoveTags) {
		toSerialize["removeTags"] = o.RemoveTags
	}
	if !IsNil(o.External) {
		toSerialize["external"] = o.External
	}
	return toSerialize, nil
}

type NullableInputUpdateObservable struct {
	value *InputUpdateObservable
	isSet bool
}

func (v NullableInputUpdateObservable) Get() *InputUpdateObservable {
	return v.value
}

func (v *NullableInputUpdateObservable) Set(val *InputUpdateObservable) {
	v.value = val
	v.isSet = true
}

func (v NullableInputUpdateObservable) IsSet() bool {
	return v.isSet
}

func (v *NullableInputUpdateObservable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputUpdateObservable(val *InputUpdateObservable) *NullableInputUpdateObservable {
	return &NullableInputUpdateObservable{value: val, isSet: true}
}

func (v NullableInputUpdateObservable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputUpdateObservable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
