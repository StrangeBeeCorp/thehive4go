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

// checks if the InputImportCase1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputImportCase1{}

// InputImportCase1 struct for InputImportCase1
type InputImportCase1 struct {
	// Name of the case template to apply to the created case. Must match an existing case template. Use `POST /api/v1/query` with `listCaseTemplate` to list available case templates.
	CaseTemplate *string `json:"caseTemplate,omitempty"`
	// Login of the user to assign as owner of the case. Typically an email address. If a default domain is configured in the [authentication settings](https://docs.strangebee.com/thehive/administration/authentication/configure-authentication), a plain username is also accepted.
	Assignee *string `json:"assignee,omitempty"`
	// Tasks to create in the case. A task represents one or more actions, called task logs, that must be completed during the investigation. If omitted and a case template is provided, the template's tasks are applied instead.
	Tasks []InputCreateTask `json:"tasks,omitempty"`
	// Pages to create in the case for structured documentation such as lessons learned, investigation reports, or meeting notes. [Case pages are temporary](https://docs.strangebee.com/thehive/user-guides/knowledge-base/about-case-pages/#permanent-vs-temporary-documentation): they're scoped to the active investigation and can't be searched. If omitted and a case template is provided, the template's pages are applied instead.
	Pages        []InputCreatePage            `json:"pages,omitempty"`
	CustomFields *InputCreateCaseCustomFields `json:"customFields,omitempty"`
	// Per-organization sharing configuration for this case. Overrides the [organization-level sharing rules](https://docs.strangebee.com/thehive/administration/organizations/about-organizations-sharing-rules/) for specific linked organizations.
	SharingParameters []InputShare `json:"sharingParameters,omitempty"`
	// Enum: `autoShare` `manual`<br>Task-sharing rule for the owner organization: `autoShare` shares new tasks with every sharing organization automatically, `manual` requires explicit sharing per task. Overrides the [organization-level sharing rules](https://docs.strangebee.com/thehive/administration/organizations/about-organizations-sharing-rules/).
	TaskRule *string `json:"taskRule,omitempty"`
	// Enum: `autoShare` `manual`<br>Observable-sharing rule for the owner organization: `autoShare` shares new observables with every sharing organization automatically, `manual` requires explicit sharing per observable. Overrides the [organization-level sharing rules](https://docs.strangebee.com/thehive/administration/organizations/about-organizations-sharing-rules/).
	ObservableRule *string `json:"observableRule,omitempty"`
}

// NewInputImportCase1 instantiates a new InputImportCase1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputImportCase1() *InputImportCase1 {
	this := InputImportCase1{}
	return &this
}

// NewInputImportCase1WithDefaults instantiates a new InputImportCase1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputImportCase1WithDefaults() *InputImportCase1 {
	this := InputImportCase1{}
	return &this
}

// GetCaseTemplate returns the CaseTemplate field value if set, zero value otherwise.
func (o *InputImportCase1) GetCaseTemplate() string {
	if o == nil || IsNil(o.CaseTemplate) {
		var ret string
		return ret
	}
	return *o.CaseTemplate
}

// GetCaseTemplateOk returns a tuple with the CaseTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputImportCase1) GetCaseTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.CaseTemplate) {
		return nil, false
	}
	return o.CaseTemplate, true
}

// HasCaseTemplate returns a boolean if a field has been set.
func (o *InputImportCase1) HasCaseTemplate() bool {
	if o != nil && !IsNil(o.CaseTemplate) {
		return true
	}

	return false
}

// SetCaseTemplate gets a reference to the given string and assigns it to the CaseTemplate field.
func (o *InputImportCase1) SetCaseTemplate(v string) {
	o.CaseTemplate = &v
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *InputImportCase1) GetAssignee() string {
	if o == nil || IsNil(o.Assignee) {
		var ret string
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputImportCase1) GetAssigneeOk() (*string, bool) {
	if o == nil || IsNil(o.Assignee) {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *InputImportCase1) HasAssignee() bool {
	if o != nil && !IsNil(o.Assignee) {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given string and assigns it to the Assignee field.
func (o *InputImportCase1) SetAssignee(v string) {
	o.Assignee = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *InputImportCase1) GetTasks() []InputCreateTask {
	if o == nil || IsNil(o.Tasks) {
		var ret []InputCreateTask
		return ret
	}
	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputImportCase1) GetTasksOk() ([]InputCreateTask, bool) {
	if o == nil || IsNil(o.Tasks) {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *InputImportCase1) HasTasks() bool {
	if o != nil && !IsNil(o.Tasks) {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []InputCreateTask and assigns it to the Tasks field.
func (o *InputImportCase1) SetTasks(v []InputCreateTask) {
	o.Tasks = v
}

// GetPages returns the Pages field value if set, zero value otherwise.
func (o *InputImportCase1) GetPages() []InputCreatePage {
	if o == nil || IsNil(o.Pages) {
		var ret []InputCreatePage
		return ret
	}
	return o.Pages
}

// GetPagesOk returns a tuple with the Pages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputImportCase1) GetPagesOk() ([]InputCreatePage, bool) {
	if o == nil || IsNil(o.Pages) {
		return nil, false
	}
	return o.Pages, true
}

// HasPages returns a boolean if a field has been set.
func (o *InputImportCase1) HasPages() bool {
	if o != nil && !IsNil(o.Pages) {
		return true
	}

	return false
}

// SetPages gets a reference to the given []InputCreatePage and assigns it to the Pages field.
func (o *InputImportCase1) SetPages(v []InputCreatePage) {
	o.Pages = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *InputImportCase1) GetCustomFields() InputCreateCaseCustomFields {
	if o == nil || IsNil(o.CustomFields) {
		var ret InputCreateCaseCustomFields
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputImportCase1) GetCustomFieldsOk() (*InputCreateCaseCustomFields, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *InputImportCase1) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given InputCreateCaseCustomFields and assigns it to the CustomFields field.
func (o *InputImportCase1) SetCustomFields(v InputCreateCaseCustomFields) {
	o.CustomFields = &v
}

// GetSharingParameters returns the SharingParameters field value if set, zero value otherwise.
func (o *InputImportCase1) GetSharingParameters() []InputShare {
	if o == nil || IsNil(o.SharingParameters) {
		var ret []InputShare
		return ret
	}
	return o.SharingParameters
}

// GetSharingParametersOk returns a tuple with the SharingParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputImportCase1) GetSharingParametersOk() ([]InputShare, bool) {
	if o == nil || IsNil(o.SharingParameters) {
		return nil, false
	}
	return o.SharingParameters, true
}

// HasSharingParameters returns a boolean if a field has been set.
func (o *InputImportCase1) HasSharingParameters() bool {
	if o != nil && !IsNil(o.SharingParameters) {
		return true
	}

	return false
}

// SetSharingParameters gets a reference to the given []InputShare and assigns it to the SharingParameters field.
func (o *InputImportCase1) SetSharingParameters(v []InputShare) {
	o.SharingParameters = v
}

// GetTaskRule returns the TaskRule field value if set, zero value otherwise.
func (o *InputImportCase1) GetTaskRule() string {
	if o == nil || IsNil(o.TaskRule) {
		var ret string
		return ret
	}
	return *o.TaskRule
}

// GetTaskRuleOk returns a tuple with the TaskRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputImportCase1) GetTaskRuleOk() (*string, bool) {
	if o == nil || IsNil(o.TaskRule) {
		return nil, false
	}
	return o.TaskRule, true
}

// HasTaskRule returns a boolean if a field has been set.
func (o *InputImportCase1) HasTaskRule() bool {
	if o != nil && !IsNil(o.TaskRule) {
		return true
	}

	return false
}

// SetTaskRule gets a reference to the given string and assigns it to the TaskRule field.
func (o *InputImportCase1) SetTaskRule(v string) {
	o.TaskRule = &v
}

// GetObservableRule returns the ObservableRule field value if set, zero value otherwise.
func (o *InputImportCase1) GetObservableRule() string {
	if o == nil || IsNil(o.ObservableRule) {
		var ret string
		return ret
	}
	return *o.ObservableRule
}

// GetObservableRuleOk returns a tuple with the ObservableRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputImportCase1) GetObservableRuleOk() (*string, bool) {
	if o == nil || IsNil(o.ObservableRule) {
		return nil, false
	}
	return o.ObservableRule, true
}

// HasObservableRule returns a boolean if a field has been set.
func (o *InputImportCase1) HasObservableRule() bool {
	if o != nil && !IsNil(o.ObservableRule) {
		return true
	}

	return false
}

// SetObservableRule gets a reference to the given string and assigns it to the ObservableRule field.
func (o *InputImportCase1) SetObservableRule(v string) {
	o.ObservableRule = &v
}

func (o InputImportCase1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputImportCase1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CaseTemplate) {
		toSerialize["caseTemplate"] = o.CaseTemplate
	}
	if !IsNil(o.Assignee) {
		toSerialize["assignee"] = o.Assignee
	}
	if !IsNil(o.Tasks) {
		toSerialize["tasks"] = o.Tasks
	}
	if !IsNil(o.Pages) {
		toSerialize["pages"] = o.Pages
	}
	if !IsNil(o.CustomFields) {
		toSerialize["customFields"] = o.CustomFields
	}
	if !IsNil(o.SharingParameters) {
		toSerialize["sharingParameters"] = o.SharingParameters
	}
	if !IsNil(o.TaskRule) {
		toSerialize["taskRule"] = o.TaskRule
	}
	if !IsNil(o.ObservableRule) {
		toSerialize["observableRule"] = o.ObservableRule
	}
	return toSerialize, nil
}

type NullableInputImportCase1 struct {
	value *InputImportCase1
	isSet bool
}

func (v NullableInputImportCase1) Get() *InputImportCase1 {
	return v.value
}

func (v *NullableInputImportCase1) Set(val *InputImportCase1) {
	v.value = val
	v.isSet = true
}

func (v NullableInputImportCase1) IsSet() bool {
	return v.isSet
}

func (v *NullableInputImportCase1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputImportCase1(val *InputImportCase1) *NullableInputImportCase1 {
	return &NullableInputImportCase1{value: val, isSet: true}
}

func (v NullableInputImportCase1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputImportCase1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
