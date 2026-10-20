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

// checks if the InputCreateCaseTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputCreateCaseTemplate{}

// InputCreateCaseTemplate struct for InputCreateCaseTemplate
type InputCreateCaseTemplate struct {
	// Name of the case template, must be unique in your organization.
	Name string `json:"name"`
	// Name shown to users when selecting this template. Defaults to the template name if not specified.
	DisplayName *string `json:"displayName,omitempty"`
	// Prefix automatically added to case titles when the template is applied. Add a hyphen after the prefix to separate it from the case title.
	TitlePrefix *string `json:"titlePrefix,omitempty"`
	// Default description for cases. Supports [TheHive-flavored Markdown](https://docs.strangebee.com/thehive/user-guides/thehive-flavored-markdown/).
	Description *string `json:"description,omitempty"`
	// Default severity level for cases: 1 (Low), 2 (Medium), 3 (High), 4 (Critical).
	Severity *int32 `json:"severity,omitempty"`
	// Default tags for cases. Unknown tags are automatically created. Use `POST /api/v1/query` with `listTag` to retrieve existing free-form tags, or with `listTaxonomy` and `tags` to retrieve taxonomy tags.
	Tags []string `json:"tags,omitempty"`
	// Marks cases as important by default.
	Flag *bool `json:"flag,omitempty"`
	// Default [Traffic Light Protocol (TLP)](https://www.misp-project.org/taxonomies.html#_tlp) level for cases controlling how case information can be shared: 0 (Clear — no restrictions), 1 (Green — community), 2 (Amber — organization and clients), 3 (Amber+Strict — organization only), 4 (Red — named recipients only).
	Tlp *int32 `json:"tlp,omitempty"`
	// Default [Permissible Actions Protocol (PAP)](https://www.misp-project.org/taxonomies.html#_pap) level for cases controlling what actions can be taken with case information: 0 (Clear — no restrictions), 1 (Green — active actions allowed), 2 (Amber — passive checks only), 3 (Red — non-detectable actions only).
	Pap *int32 `json:"pap,omitempty"`
	// Default summary for cases. Supports [TheHive-flavored Markdown](https://docs.strangebee.com/thehive/user-guides/thehive-flavored-markdown/).
	Summary *string `json:"summary,omitempty"`
	// Task templates to include in this case template. A task represents one or more actions, called task logs, that must be completed during the investigation. These tasks are automatically created on cases when the template is applied.
	Tasks []InputCreateTask `json:"tasks,omitempty"`
	// Use `pageTemplates` instead.
	PageTemplateIds []string                             `json:"pageTemplateIds,omitempty"`
	CustomFields    *InputCreateCaseTemplateCustomFields `json:"customFields,omitempty"`
	// Pages enable the creation and sharing of valuable information within a team on a case. Pages to include in this case template are automatically created on cases when the template is applied. Preferred over the deprecated `pageTemplateIds` field.
	PageTemplates []InputCreatePage `json:"pageTemplates,omitempty"`
}

type _InputCreateCaseTemplate InputCreateCaseTemplate

// NewInputCreateCaseTemplate instantiates a new InputCreateCaseTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputCreateCaseTemplate(name string) *InputCreateCaseTemplate {
	this := InputCreateCaseTemplate{}
	this.Name = name
	return &this
}

// NewInputCreateCaseTemplateWithDefaults instantiates a new InputCreateCaseTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputCreateCaseTemplateWithDefaults() *InputCreateCaseTemplate {
	this := InputCreateCaseTemplate{}
	return &this
}

// GetName returns the Name field value
func (o *InputCreateCaseTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InputCreateCaseTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InputCreateCaseTemplate) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *InputCreateCaseTemplate) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateCaseTemplate) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *InputCreateCaseTemplate) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *InputCreateCaseTemplate) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetTitlePrefix returns the TitlePrefix field value if set, zero value otherwise.
func (o *InputCreateCaseTemplate) GetTitlePrefix() string {
	if o == nil || IsNil(o.TitlePrefix) {
		var ret string
		return ret
	}
	return *o.TitlePrefix
}

// GetTitlePrefixOk returns a tuple with the TitlePrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateCaseTemplate) GetTitlePrefixOk() (*string, bool) {
	if o == nil || IsNil(o.TitlePrefix) {
		return nil, false
	}
	return o.TitlePrefix, true
}

// HasTitlePrefix returns a boolean if a field has been set.
func (o *InputCreateCaseTemplate) HasTitlePrefix() bool {
	if o != nil && !IsNil(o.TitlePrefix) {
		return true
	}

	return false
}

// SetTitlePrefix gets a reference to the given string and assigns it to the TitlePrefix field.
func (o *InputCreateCaseTemplate) SetTitlePrefix(v string) {
	o.TitlePrefix = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InputCreateCaseTemplate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateCaseTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InputCreateCaseTemplate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InputCreateCaseTemplate) SetDescription(v string) {
	o.Description = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *InputCreateCaseTemplate) GetSeverity() int32 {
	if o == nil || IsNil(o.Severity) {
		var ret int32
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateCaseTemplate) GetSeverityOk() (*int32, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *InputCreateCaseTemplate) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given int32 and assigns it to the Severity field.
func (o *InputCreateCaseTemplate) SetSeverity(v int32) {
	o.Severity = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InputCreateCaseTemplate) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateCaseTemplate) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InputCreateCaseTemplate) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InputCreateCaseTemplate) SetTags(v []string) {
	o.Tags = v
}

// GetFlag returns the Flag field value if set, zero value otherwise.
func (o *InputCreateCaseTemplate) GetFlag() bool {
	if o == nil || IsNil(o.Flag) {
		var ret bool
		return ret
	}
	return *o.Flag
}

// GetFlagOk returns a tuple with the Flag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateCaseTemplate) GetFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.Flag) {
		return nil, false
	}
	return o.Flag, true
}

// HasFlag returns a boolean if a field has been set.
func (o *InputCreateCaseTemplate) HasFlag() bool {
	if o != nil && !IsNil(o.Flag) {
		return true
	}

	return false
}

// SetFlag gets a reference to the given bool and assigns it to the Flag field.
func (o *InputCreateCaseTemplate) SetFlag(v bool) {
	o.Flag = &v
}

// GetTlp returns the Tlp field value if set, zero value otherwise.
func (o *InputCreateCaseTemplate) GetTlp() int32 {
	if o == nil || IsNil(o.Tlp) {
		var ret int32
		return ret
	}
	return *o.Tlp
}

// GetTlpOk returns a tuple with the Tlp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateCaseTemplate) GetTlpOk() (*int32, bool) {
	if o == nil || IsNil(o.Tlp) {
		return nil, false
	}
	return o.Tlp, true
}

// HasTlp returns a boolean if a field has been set.
func (o *InputCreateCaseTemplate) HasTlp() bool {
	if o != nil && !IsNil(o.Tlp) {
		return true
	}

	return false
}

// SetTlp gets a reference to the given int32 and assigns it to the Tlp field.
func (o *InputCreateCaseTemplate) SetTlp(v int32) {
	o.Tlp = &v
}

// GetPap returns the Pap field value if set, zero value otherwise.
func (o *InputCreateCaseTemplate) GetPap() int32 {
	if o == nil || IsNil(o.Pap) {
		var ret int32
		return ret
	}
	return *o.Pap
}

// GetPapOk returns a tuple with the Pap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateCaseTemplate) GetPapOk() (*int32, bool) {
	if o == nil || IsNil(o.Pap) {
		return nil, false
	}
	return o.Pap, true
}

// HasPap returns a boolean if a field has been set.
func (o *InputCreateCaseTemplate) HasPap() bool {
	if o != nil && !IsNil(o.Pap) {
		return true
	}

	return false
}

// SetPap gets a reference to the given int32 and assigns it to the Pap field.
func (o *InputCreateCaseTemplate) SetPap(v int32) {
	o.Pap = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *InputCreateCaseTemplate) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateCaseTemplate) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *InputCreateCaseTemplate) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *InputCreateCaseTemplate) SetSummary(v string) {
	o.Summary = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *InputCreateCaseTemplate) GetTasks() []InputCreateTask {
	if o == nil || IsNil(o.Tasks) {
		var ret []InputCreateTask
		return ret
	}
	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateCaseTemplate) GetTasksOk() ([]InputCreateTask, bool) {
	if o == nil || IsNil(o.Tasks) {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *InputCreateCaseTemplate) HasTasks() bool {
	if o != nil && !IsNil(o.Tasks) {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []InputCreateTask and assigns it to the Tasks field.
func (o *InputCreateCaseTemplate) SetTasks(v []InputCreateTask) {
	o.Tasks = v
}

// GetPageTemplateIds returns the PageTemplateIds field value if set, zero value otherwise.
func (o *InputCreateCaseTemplate) GetPageTemplateIds() []string {
	if o == nil || IsNil(o.PageTemplateIds) {
		var ret []string
		return ret
	}
	return o.PageTemplateIds
}

// GetPageTemplateIdsOk returns a tuple with the PageTemplateIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateCaseTemplate) GetPageTemplateIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.PageTemplateIds) {
		return nil, false
	}
	return o.PageTemplateIds, true
}

// HasPageTemplateIds returns a boolean if a field has been set.
func (o *InputCreateCaseTemplate) HasPageTemplateIds() bool {
	if o != nil && !IsNil(o.PageTemplateIds) {
		return true
	}

	return false
}

// SetPageTemplateIds gets a reference to the given []string and assigns it to the PageTemplateIds field.
func (o *InputCreateCaseTemplate) SetPageTemplateIds(v []string) {
	o.PageTemplateIds = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *InputCreateCaseTemplate) GetCustomFields() InputCreateCaseTemplateCustomFields {
	if o == nil || IsNil(o.CustomFields) {
		var ret InputCreateCaseTemplateCustomFields
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateCaseTemplate) GetCustomFieldsOk() (*InputCreateCaseTemplateCustomFields, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *InputCreateCaseTemplate) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given InputCreateCaseTemplateCustomFields and assigns it to the CustomFields field.
func (o *InputCreateCaseTemplate) SetCustomFields(v InputCreateCaseTemplateCustomFields) {
	o.CustomFields = &v
}

// GetPageTemplates returns the PageTemplates field value if set, zero value otherwise.
func (o *InputCreateCaseTemplate) GetPageTemplates() []InputCreatePage {
	if o == nil || IsNil(o.PageTemplates) {
		var ret []InputCreatePage
		return ret
	}
	return o.PageTemplates
}

// GetPageTemplatesOk returns a tuple with the PageTemplates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateCaseTemplate) GetPageTemplatesOk() ([]InputCreatePage, bool) {
	if o == nil || IsNil(o.PageTemplates) {
		return nil, false
	}
	return o.PageTemplates, true
}

// HasPageTemplates returns a boolean if a field has been set.
func (o *InputCreateCaseTemplate) HasPageTemplates() bool {
	if o != nil && !IsNil(o.PageTemplates) {
		return true
	}

	return false
}

// SetPageTemplates gets a reference to the given []InputCreatePage and assigns it to the PageTemplates field.
func (o *InputCreateCaseTemplate) SetPageTemplates(v []InputCreatePage) {
	o.PageTemplates = v
}

func (o InputCreateCaseTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputCreateCaseTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
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
	if !IsNil(o.Tasks) {
		toSerialize["tasks"] = o.Tasks
	}
	if !IsNil(o.PageTemplateIds) {
		toSerialize["pageTemplateIds"] = o.PageTemplateIds
	}
	if !IsNil(o.CustomFields) {
		toSerialize["customFields"] = o.CustomFields
	}
	if !IsNil(o.PageTemplates) {
		toSerialize["pageTemplates"] = o.PageTemplates
	}
	return toSerialize, nil
}

func (o *InputCreateCaseTemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varInputCreateCaseTemplate := _InputCreateCaseTemplate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInputCreateCaseTemplate)

	if err != nil {
		return err
	}

	*o = InputCreateCaseTemplate(varInputCreateCaseTemplate)

	return err
}

type NullableInputCreateCaseTemplate struct {
	value *InputCreateCaseTemplate
	isSet bool
}

func (v NullableInputCreateCaseTemplate) Get() *InputCreateCaseTemplate {
	return v.value
}

func (v *NullableInputCreateCaseTemplate) Set(val *InputCreateCaseTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableInputCreateCaseTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableInputCreateCaseTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputCreateCaseTemplate(val *InputCreateCaseTemplate) *NullableInputCreateCaseTemplate {
	return &NullableInputCreateCaseTemplate{value: val, isSet: true}
}

func (v NullableInputCreateCaseTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputCreateCaseTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
