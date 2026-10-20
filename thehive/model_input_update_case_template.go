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

// checks if the InputUpdateCaseTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputUpdateCaseTemplate{}

// InputUpdateCaseTemplate struct for InputUpdateCaseTemplate
type InputUpdateCaseTemplate struct {
	// New name for the case template, must be unique in your organization.
	Name *string `json:"name,omitempty"`
	// New display name shown to users when selecting this template.
	DisplayName *string `json:"displayName,omitempty"`
	// New prefix for case titles. Send `null` to remove the prefix.
	TitlePrefix *string `json:"titlePrefix,omitempty"`
	// New default description for cases. Supports [TheHive-flavored Markdown](https://docs.strangebee.com/thehive/user-guides/thehive-flavored-markdown/). Send `null` to clear this field.
	Description *string `json:"description,omitempty"`
	// New default severity level for cases: 1 (Low), 2 (Medium), 3 (High), 4 (Critical). Send `null` to clear this field.
	Severity *int32 `json:"severity,omitempty"`
	// Replaces the template's current tag list with this set. Unknown tags are automatically created. Use `POST /api/v1/query` with `listTag` to retrieve existing free-form tags, or with `listTaxonomy` and `tags` to retrieve taxonomy tags.
	Tags []string `json:"tags,omitempty"`
	// Marks cases as important by default.
	Flag *bool `json:"flag,omitempty"`
	// New default [Traffic Light Protocol (TLP)](https://www.misp-project.org/taxonomies.html#_tlp) level for cases: 0 (Clear), 1 (Green), 2 (Amber), 3 (Amber+Strict), 4 (Red). Send `null` to clear this field.
	Tlp *int32 `json:"tlp,omitempty"`
	// New default [Permissible Actions Protocol (PAP)](https://www.misp-project.org/taxonomies.html#_pap) level for cases: 0 (Clear), 1 (Green), 2 (Amber), 3 (Red). Send `null` to clear this field.
	Pap *int32 `json:"pap,omitempty"`
	// New default summary for cases. Supports [TheHive-flavored Markdown](https://docs.strangebee.com/thehive/user-guides/thehive-flavored-markdown/). Send `null` to clear this field.
	Summary      *string                              `json:"summary,omitempty"`
	CustomFields *InputUpdateCaseTemplateCustomFields `json:"customFields,omitempty"`
	// Replaces the template's current task list with this set. A task represents one or more actions, called task logs, that must be completed during the investigation.
	Tasks []InputCreateTask `json:"tasks,omitempty"`
}

// NewInputUpdateCaseTemplate instantiates a new InputUpdateCaseTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputUpdateCaseTemplate() *InputUpdateCaseTemplate {
	this := InputUpdateCaseTemplate{}
	return &this
}

// NewInputUpdateCaseTemplateWithDefaults instantiates a new InputUpdateCaseTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputUpdateCaseTemplateWithDefaults() *InputUpdateCaseTemplate {
	this := InputUpdateCaseTemplate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InputUpdateCaseTemplate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateCaseTemplate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InputUpdateCaseTemplate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InputUpdateCaseTemplate) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *InputUpdateCaseTemplate) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateCaseTemplate) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *InputUpdateCaseTemplate) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *InputUpdateCaseTemplate) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetTitlePrefix returns the TitlePrefix field value if set, zero value otherwise.
func (o *InputUpdateCaseTemplate) GetTitlePrefix() string {
	if o == nil || IsNil(o.TitlePrefix) {
		var ret string
		return ret
	}
	return *o.TitlePrefix
}

// GetTitlePrefixOk returns a tuple with the TitlePrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateCaseTemplate) GetTitlePrefixOk() (*string, bool) {
	if o == nil || IsNil(o.TitlePrefix) {
		return nil, false
	}
	return o.TitlePrefix, true
}

// HasTitlePrefix returns a boolean if a field has been set.
func (o *InputUpdateCaseTemplate) HasTitlePrefix() bool {
	if o != nil && !IsNil(o.TitlePrefix) {
		return true
	}

	return false
}

// SetTitlePrefix gets a reference to the given string and assigns it to the TitlePrefix field.
func (o *InputUpdateCaseTemplate) SetTitlePrefix(v string) {
	o.TitlePrefix = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InputUpdateCaseTemplate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateCaseTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InputUpdateCaseTemplate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InputUpdateCaseTemplate) SetDescription(v string) {
	o.Description = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *InputUpdateCaseTemplate) GetSeverity() int32 {
	if o == nil || IsNil(o.Severity) {
		var ret int32
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateCaseTemplate) GetSeverityOk() (*int32, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *InputUpdateCaseTemplate) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given int32 and assigns it to the Severity field.
func (o *InputUpdateCaseTemplate) SetSeverity(v int32) {
	o.Severity = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InputUpdateCaseTemplate) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateCaseTemplate) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InputUpdateCaseTemplate) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InputUpdateCaseTemplate) SetTags(v []string) {
	o.Tags = v
}

// GetFlag returns the Flag field value if set, zero value otherwise.
func (o *InputUpdateCaseTemplate) GetFlag() bool {
	if o == nil || IsNil(o.Flag) {
		var ret bool
		return ret
	}
	return *o.Flag
}

// GetFlagOk returns a tuple with the Flag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateCaseTemplate) GetFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.Flag) {
		return nil, false
	}
	return o.Flag, true
}

// HasFlag returns a boolean if a field has been set.
func (o *InputUpdateCaseTemplate) HasFlag() bool {
	if o != nil && !IsNil(o.Flag) {
		return true
	}

	return false
}

// SetFlag gets a reference to the given bool and assigns it to the Flag field.
func (o *InputUpdateCaseTemplate) SetFlag(v bool) {
	o.Flag = &v
}

// GetTlp returns the Tlp field value if set, zero value otherwise.
func (o *InputUpdateCaseTemplate) GetTlp() int32 {
	if o == nil || IsNil(o.Tlp) {
		var ret int32
		return ret
	}
	return *o.Tlp
}

// GetTlpOk returns a tuple with the Tlp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateCaseTemplate) GetTlpOk() (*int32, bool) {
	if o == nil || IsNil(o.Tlp) {
		return nil, false
	}
	return o.Tlp, true
}

// HasTlp returns a boolean if a field has been set.
func (o *InputUpdateCaseTemplate) HasTlp() bool {
	if o != nil && !IsNil(o.Tlp) {
		return true
	}

	return false
}

// SetTlp gets a reference to the given int32 and assigns it to the Tlp field.
func (o *InputUpdateCaseTemplate) SetTlp(v int32) {
	o.Tlp = &v
}

// GetPap returns the Pap field value if set, zero value otherwise.
func (o *InputUpdateCaseTemplate) GetPap() int32 {
	if o == nil || IsNil(o.Pap) {
		var ret int32
		return ret
	}
	return *o.Pap
}

// GetPapOk returns a tuple with the Pap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateCaseTemplate) GetPapOk() (*int32, bool) {
	if o == nil || IsNil(o.Pap) {
		return nil, false
	}
	return o.Pap, true
}

// HasPap returns a boolean if a field has been set.
func (o *InputUpdateCaseTemplate) HasPap() bool {
	if o != nil && !IsNil(o.Pap) {
		return true
	}

	return false
}

// SetPap gets a reference to the given int32 and assigns it to the Pap field.
func (o *InputUpdateCaseTemplate) SetPap(v int32) {
	o.Pap = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *InputUpdateCaseTemplate) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateCaseTemplate) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *InputUpdateCaseTemplate) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *InputUpdateCaseTemplate) SetSummary(v string) {
	o.Summary = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *InputUpdateCaseTemplate) GetCustomFields() InputUpdateCaseTemplateCustomFields {
	if o == nil || IsNil(o.CustomFields) {
		var ret InputUpdateCaseTemplateCustomFields
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateCaseTemplate) GetCustomFieldsOk() (*InputUpdateCaseTemplateCustomFields, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *InputUpdateCaseTemplate) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given InputUpdateCaseTemplateCustomFields and assigns it to the CustomFields field.
func (o *InputUpdateCaseTemplate) SetCustomFields(v InputUpdateCaseTemplateCustomFields) {
	o.CustomFields = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *InputUpdateCaseTemplate) GetTasks() []InputCreateTask {
	if o == nil || IsNil(o.Tasks) {
		var ret []InputCreateTask
		return ret
	}
	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateCaseTemplate) GetTasksOk() ([]InputCreateTask, bool) {
	if o == nil || IsNil(o.Tasks) {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *InputUpdateCaseTemplate) HasTasks() bool {
	if o != nil && !IsNil(o.Tasks) {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []InputCreateTask and assigns it to the Tasks field.
func (o *InputUpdateCaseTemplate) SetTasks(v []InputCreateTask) {
	o.Tasks = v
}

func (o InputUpdateCaseTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputUpdateCaseTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.TitlePrefix) {
		toSerialize["titlePrefix"] = o.TitlePrefix
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Flag) {
		toSerialize["flag"] = o.Flag
	}
	if !IsNil(o.Tlp) {
		toSerialize["tlp"] = o.Tlp
	}
	if !IsNil(o.Pap) {
		toSerialize["pap"] = o.Pap
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !IsNil(o.CustomFields) {
		toSerialize["customFields"] = o.CustomFields
	}
	if !IsNil(o.Tasks) {
		toSerialize["tasks"] = o.Tasks
	}
	return toSerialize, nil
}

type NullableInputUpdateCaseTemplate struct {
	value *InputUpdateCaseTemplate
	isSet bool
}

func (v NullableInputUpdateCaseTemplate) Get() *InputUpdateCaseTemplate {
	return v.value
}

func (v *NullableInputUpdateCaseTemplate) Set(val *InputUpdateCaseTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableInputUpdateCaseTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableInputUpdateCaseTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputUpdateCaseTemplate(val *InputUpdateCaseTemplate) *NullableInputUpdateCaseTemplate {
	return &NullableInputUpdateCaseTemplate{value: val, isSet: true}
}

func (v NullableInputUpdateCaseTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputUpdateCaseTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
