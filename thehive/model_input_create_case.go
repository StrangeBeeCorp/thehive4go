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

// checks if the InputCreateCase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputCreateCase{}

// InputCreateCase struct for InputCreateCase
type InputCreateCase struct {
	// Title of the case.
	Title string `json:"title"`
	// Details of the case. Supports [TheHive-flavored Markdown](https://docs.strangebee.com/thehive/user-guides/thehive-flavored-markdown/).
	Description string `json:"description"`
	// Numeric severity level of the case: 1 (Low), 2 (Medium), 3 (High), 4 (Critical).
	Severity *int32 `json:"severity,omitempty"`
	// Start date of the incident. Expected format: Unix timestamp in milliseconds. Defaults to the current date. Used for KPI calculations.
	StartDate *int64 `json:"startDate,omitempty"`
	// End date of the incident. Expected format: Unix timestamp in milliseconds. Used for KPI calculations.
	EndDate *int64 `json:"endDate,omitempty"`
	// Labels to categorize the case. Unknown tags are automatically created. Use `POST /api/v1/query` with `listTag` to retrieve existing custom tags, or with `listTaxonomy` and `tags` to retrieve taxonomy tags.
	Tags []string `json:"tags,omitempty"`
	// Marks the case as important for quick filtering.
	Flag *bool `json:"flag,omitempty"`
	// [Traffic Light Protocol (TLP)](https://www.misp-project.org/taxonomies.html#_tlp) level controlling how case information can be shared: 0 (Clear — no restrictions), 1 (Green — community), 2 (Amber — organization and clients), 3 (Amber+Strict — organization only), 4 (Red — named recipients only).
	Tlp *int32 `json:"tlp,omitempty"`
	// [Permissible Actions Protocol (PAP)](https://www.misp-project.org/taxonomies.html#_pap) level controlling what actions can be taken with case information: 0 (Clear — no restrictions), 1 (Green — active actions allowed), 2 (Amber — passive checks only), 3 (Red — non-detectable actions only).
	Pap *int32 `json:"pap,omitempty"`
	// Name of the case status to apply. Must match an existing status. Use `POST /api/v1/query` with `listCaseStatus` to list available statuses.
	Status *string `json:"status,omitempty"`
	// Brief summary of the investigation findings. Supports [TheHive-flavored Markdown](https://docs.strangebee.com/thehive/user-guides/thehive-flavored-markdown/).
	Summary *string `json:"summary,omitempty"`
	// Login of the user to assign the case to. Typically an email address. If a default domain is configured in the [authentication settings](https://docs.strangebee.com/thehive/administration/authentication/configure-authentication), a plain username is also accepted. Use `POST /api/v1/query` with `listUser` to retrieve available users and their logins. If omitted, the case is unassigned.
	Assignee     *string                      `json:"assignee,omitempty"`
	Access       *Access                      `json:"access,omitempty"`
	CustomFields *InputCreateCaseCustomFields `json:"customFields,omitempty"`
	// Name or ID of the case template to use. Applies the template's tasks, custom fields, and pages. Must match an existing case template. Use `POST /api/v1/query` with `listCaseTemplate` to list available case templates.
	CaseTemplate *string `json:"caseTemplate,omitempty"`
	// Tasks to create in the case. A task represents one or more actions, called task logs, that must be completed during the investigation. If omitted and a case template is provided, the template's tasks are applied instead.
	Tasks []InputCreateTask `json:"tasks,omitempty"`
	// Pages to create in the case for structured documentation such as lessons learned, investigation reports, or meeting notes. [Case pages are temporary](https://docs.strangebee.com/thehive/user-guides/knowledge-base/about-case-pages/#permanent-vs-temporary-documentation): they're scoped to the active investigation and can't be searched. If omitted and a case template is provided, the template's pages are applied instead.
	Pages []InputCreatePage `json:"pages,omitempty"`
	// Per-organization sharing configuration for this case. Overrides the [organization-level sharing rules](https://docs.strangebee.com/thehive/administration/organizations/about-organizations-sharing-rules/) for specific linked organizations.
	SharingParameters []InputShare `json:"sharingParameters,omitempty"`
	// Enum: `autoShare` `manual`<br>Task-sharing rule for the owner organization: `autoShare` shares new tasks with every sharing organization automatically, `manual` requires explicit sharing per task. Overrides the [organization-level sharing rules](https://docs.strangebee.com/thehive/administration/organizations/about-organizations-sharing-rules/).
	TaskRule *string `json:"taskRule,omitempty"`
	// Enum: `autoShare` `manual`<br>Observable-sharing rule for the owner organization: `autoShare` shares new observables with every sharing organization automatically, `manual` requires explicit sharing per observable. Overrides the [organization-level sharing rules](https://docs.strangebee.com/thehive/administration/organizations/about-organizations-sharing-rules/).
	ObservableRule *string `json:"observableRule,omitempty"`
}

type _InputCreateCase InputCreateCase

// NewInputCreateCase instantiates a new InputCreateCase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputCreateCase(title string, description string) *InputCreateCase {
	this := InputCreateCase{}
	this.Title = title
	this.Description = description
	var severity int32 = 2
	this.Severity = &severity
	var flag bool = false
	this.Flag = &flag
	var tlp int32 = 2
	this.Tlp = &tlp
	var pap int32 = 2
	this.Pap = &pap
	var status string = "New"
	this.Status = &status
	return &this
}

// NewInputCreateCaseWithDefaults instantiates a new InputCreateCase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputCreateCaseWithDefaults() *InputCreateCase {
	this := InputCreateCase{}
	var severity int32 = 2
	this.Severity = &severity
	var flag bool = false
	this.Flag = &flag
	var tlp int32 = 2
	this.Tlp = &tlp
	var pap int32 = 2
	this.Pap = &pap
	var status string = "New"
	this.Status = &status
	return &this
}

// GetTitle returns the Title field value
func (o *InputCreateCase) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *InputCreateCase) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *InputCreateCase) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value
func (o *InputCreateCase) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *InputCreateCase) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *InputCreateCase) SetDescription(v string) {
	o.Description = v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *InputCreateCase) GetSeverity() int32 {
	if o == nil || IsNil(o.Severity) {
		var ret int32
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateCase) GetSeverityOk() (*int32, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *InputCreateCase) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given int32 and assigns it to the Severity field.
func (o *InputCreateCase) SetSeverity(v int32) {
	o.Severity = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *InputCreateCase) GetStartDate() int64 {
	if o == nil || IsNil(o.StartDate) {
		var ret int64
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateCase) GetStartDateOk() (*int64, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *InputCreateCase) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given int64 and assigns it to the StartDate field.
func (o *InputCreateCase) SetStartDate(v int64) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *InputCreateCase) GetEndDate() int64 {
	if o == nil || IsNil(o.EndDate) {
		var ret int64
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateCase) GetEndDateOk() (*int64, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *InputCreateCase) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given int64 and assigns it to the EndDate field.
func (o *InputCreateCase) SetEndDate(v int64) {
	o.EndDate = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InputCreateCase) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateCase) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InputCreateCase) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InputCreateCase) SetTags(v []string) {
	o.Tags = v
}

// GetFlag returns the Flag field value if set, zero value otherwise.
func (o *InputCreateCase) GetFlag() bool {
	if o == nil || IsNil(o.Flag) {
		var ret bool
		return ret
	}
	return *o.Flag
}

// GetFlagOk returns a tuple with the Flag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateCase) GetFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.Flag) {
		return nil, false
	}
	return o.Flag, true
}

// HasFlag returns a boolean if a field has been set.
func (o *InputCreateCase) HasFlag() bool {
	if o != nil && !IsNil(o.Flag) {
		return true
	}

	return false
}

// SetFlag gets a reference to the given bool and assigns it to the Flag field.
func (o *InputCreateCase) SetFlag(v bool) {
	o.Flag = &v
}

// GetTlp returns the Tlp field value if set, zero value otherwise.
func (o *InputCreateCase) GetTlp() int32 {
	if o == nil || IsNil(o.Tlp) {
		var ret int32
		return ret
	}
	return *o.Tlp
}

// GetTlpOk returns a tuple with the Tlp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateCase) GetTlpOk() (*int32, bool) {
	if o == nil || IsNil(o.Tlp) {
		return nil, false
	}
	return o.Tlp, true
}

// HasTlp returns a boolean if a field has been set.
func (o *InputCreateCase) HasTlp() bool {
	if o != nil && !IsNil(o.Tlp) {
		return true
	}

	return false
}

// SetTlp gets a reference to the given int32 and assigns it to the Tlp field.
func (o *InputCreateCase) SetTlp(v int32) {
	o.Tlp = &v
}

// GetPap returns the Pap field value if set, zero value otherwise.
func (o *InputCreateCase) GetPap() int32 {
	if o == nil || IsNil(o.Pap) {
		var ret int32
		return ret
	}
	return *o.Pap
}

// GetPapOk returns a tuple with the Pap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateCase) GetPapOk() (*int32, bool) {
	if o == nil || IsNil(o.Pap) {
		return nil, false
	}
	return o.Pap, true
}

// HasPap returns a boolean if a field has been set.
func (o *InputCreateCase) HasPap() bool {
	if o != nil && !IsNil(o.Pap) {
		return true
	}

	return false
}

// SetPap gets a reference to the given int32 and assigns it to the Pap field.
func (o *InputCreateCase) SetPap(v int32) {
	o.Pap = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InputCreateCase) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateCase) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InputCreateCase) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InputCreateCase) SetStatus(v string) {
	o.Status = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *InputCreateCase) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateCase) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *InputCreateCase) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *InputCreateCase) SetSummary(v string) {
	o.Summary = &v
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *InputCreateCase) GetAssignee() string {
	if o == nil || IsNil(o.Assignee) {
		var ret string
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateCase) GetAssigneeOk() (*string, bool) {
	if o == nil || IsNil(o.Assignee) {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *InputCreateCase) HasAssignee() bool {
	if o != nil && !IsNil(o.Assignee) {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given string and assigns it to the Assignee field.
func (o *InputCreateCase) SetAssignee(v string) {
	o.Assignee = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *InputCreateCase) GetAccess() Access {
	if o == nil || IsNil(o.Access) {
		var ret Access
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateCase) GetAccessOk() (*Access, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *InputCreateCase) HasAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given Access and assigns it to the Access field.
func (o *InputCreateCase) SetAccess(v Access) {
	o.Access = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *InputCreateCase) GetCustomFields() InputCreateCaseCustomFields {
	if o == nil || IsNil(o.CustomFields) {
		var ret InputCreateCaseCustomFields
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateCase) GetCustomFieldsOk() (*InputCreateCaseCustomFields, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *InputCreateCase) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given InputCreateCaseCustomFields and assigns it to the CustomFields field.
func (o *InputCreateCase) SetCustomFields(v InputCreateCaseCustomFields) {
	o.CustomFields = &v
}

// GetCaseTemplate returns the CaseTemplate field value if set, zero value otherwise.
func (o *InputCreateCase) GetCaseTemplate() string {
	if o == nil || IsNil(o.CaseTemplate) {
		var ret string
		return ret
	}
	return *o.CaseTemplate
}

// GetCaseTemplateOk returns a tuple with the CaseTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateCase) GetCaseTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.CaseTemplate) {
		return nil, false
	}
	return o.CaseTemplate, true
}

// HasCaseTemplate returns a boolean if a field has been set.
func (o *InputCreateCase) HasCaseTemplate() bool {
	if o != nil && !IsNil(o.CaseTemplate) {
		return true
	}

	return false
}

// SetCaseTemplate gets a reference to the given string and assigns it to the CaseTemplate field.
func (o *InputCreateCase) SetCaseTemplate(v string) {
	o.CaseTemplate = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *InputCreateCase) GetTasks() []InputCreateTask {
	if o == nil || IsNil(o.Tasks) {
		var ret []InputCreateTask
		return ret
	}
	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateCase) GetTasksOk() ([]InputCreateTask, bool) {
	if o == nil || IsNil(o.Tasks) {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *InputCreateCase) HasTasks() bool {
	if o != nil && !IsNil(o.Tasks) {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []InputCreateTask and assigns it to the Tasks field.
func (o *InputCreateCase) SetTasks(v []InputCreateTask) {
	o.Tasks = v
}

// GetPages returns the Pages field value if set, zero value otherwise.
func (o *InputCreateCase) GetPages() []InputCreatePage {
	if o == nil || IsNil(o.Pages) {
		var ret []InputCreatePage
		return ret
	}
	return o.Pages
}

// GetPagesOk returns a tuple with the Pages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateCase) GetPagesOk() ([]InputCreatePage, bool) {
	if o == nil || IsNil(o.Pages) {
		return nil, false
	}
	return o.Pages, true
}

// HasPages returns a boolean if a field has been set.
func (o *InputCreateCase) HasPages() bool {
	if o != nil && !IsNil(o.Pages) {
		return true
	}

	return false
}

// SetPages gets a reference to the given []InputCreatePage and assigns it to the Pages field.
func (o *InputCreateCase) SetPages(v []InputCreatePage) {
	o.Pages = v
}

// GetSharingParameters returns the SharingParameters field value if set, zero value otherwise.
func (o *InputCreateCase) GetSharingParameters() []InputShare {
	if o == nil || IsNil(o.SharingParameters) {
		var ret []InputShare
		return ret
	}
	return o.SharingParameters
}

// GetSharingParametersOk returns a tuple with the SharingParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateCase) GetSharingParametersOk() ([]InputShare, bool) {
	if o == nil || IsNil(o.SharingParameters) {
		return nil, false
	}
	return o.SharingParameters, true
}

// HasSharingParameters returns a boolean if a field has been set.
func (o *InputCreateCase) HasSharingParameters() bool {
	if o != nil && !IsNil(o.SharingParameters) {
		return true
	}

	return false
}

// SetSharingParameters gets a reference to the given []InputShare and assigns it to the SharingParameters field.
func (o *InputCreateCase) SetSharingParameters(v []InputShare) {
	o.SharingParameters = v
}

// GetTaskRule returns the TaskRule field value if set, zero value otherwise.
func (o *InputCreateCase) GetTaskRule() string {
	if o == nil || IsNil(o.TaskRule) {
		var ret string
		return ret
	}
	return *o.TaskRule
}

// GetTaskRuleOk returns a tuple with the TaskRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateCase) GetTaskRuleOk() (*string, bool) {
	if o == nil || IsNil(o.TaskRule) {
		return nil, false
	}
	return o.TaskRule, true
}

// HasTaskRule returns a boolean if a field has been set.
func (o *InputCreateCase) HasTaskRule() bool {
	if o != nil && !IsNil(o.TaskRule) {
		return true
	}

	return false
}

// SetTaskRule gets a reference to the given string and assigns it to the TaskRule field.
func (o *InputCreateCase) SetTaskRule(v string) {
	o.TaskRule = &v
}

// GetObservableRule returns the ObservableRule field value if set, zero value otherwise.
func (o *InputCreateCase) GetObservableRule() string {
	if o == nil || IsNil(o.ObservableRule) {
		var ret string
		return ret
	}
	return *o.ObservableRule
}

// GetObservableRuleOk returns a tuple with the ObservableRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateCase) GetObservableRuleOk() (*string, bool) {
	if o == nil || IsNil(o.ObservableRule) {
		return nil, false
	}
	return o.ObservableRule, true
}

// HasObservableRule returns a boolean if a field has been set.
func (o *InputCreateCase) HasObservableRule() bool {
	if o != nil && !IsNil(o.ObservableRule) {
		return true
	}

	return false
}

// SetObservableRule gets a reference to the given string and assigns it to the ObservableRule field.
func (o *InputCreateCase) SetObservableRule(v string) {
	o.ObservableRule = &v
}

func (o InputCreateCase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputCreateCase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	toSerialize["description"] = o.Description
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
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
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !IsNil(o.Assignee) {
		toSerialize["assignee"] = o.Assignee
	}
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	if !IsNil(o.CustomFields) {
		toSerialize["customFields"] = o.CustomFields
	}
	if !IsNil(o.CaseTemplate) {
		toSerialize["caseTemplate"] = o.CaseTemplate
	}
	if !IsNil(o.Tasks) {
		toSerialize["tasks"] = o.Tasks
	}
	if !IsNil(o.Pages) {
		toSerialize["pages"] = o.Pages
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

func (o *InputCreateCase) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
		"description",
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

	varInputCreateCase := _InputCreateCase{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInputCreateCase)

	if err != nil {
		return err
	}

	*o = InputCreateCase(varInputCreateCase)

	return err
}

type NullableInputCreateCase struct {
	value *InputCreateCase
	isSet bool
}

func (v NullableInputCreateCase) Get() *InputCreateCase {
	return v.value
}

func (v *NullableInputCreateCase) Set(val *InputCreateCase) {
	v.value = val
	v.isSet = true
}

func (v NullableInputCreateCase) IsSet() bool {
	return v.isSet
}

func (v *NullableInputCreateCase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputCreateCase(val *InputCreateCase) *NullableInputCreateCase {
	return &NullableInputCreateCase{value: val, isSet: true}
}

func (v NullableInputCreateCase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputCreateCase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
