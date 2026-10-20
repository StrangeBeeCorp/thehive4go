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

// checks if the InputCreateAlert type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputCreateAlert{}

// InputCreateAlert struct for InputCreateAlert
type InputCreateAlert struct {
	// Type of the alert. Identifies the event category and groups related alerts.
	Type string `json:"type"`
	// Source system that generated the alert.
	Source string `json:"source"`
	// Unique reference identifier from the source system. Combined with `type` and `source`, uniquely identifies an alert.
	SourceRef string `json:"sourceRef"`
	// External URL linking to the alert in the source system.
	ExternalLink *string `json:"externalLink,omitempty"`
	// Short title describing the alert.
	Title string `json:"title"`
	// Full description of the alert. Supports [TheHive-flavored Markdown](https://docs.strangebee.com/thehive/user-guides/thehive-flavored-markdown/).
	Description string `json:"description"`
	// Numeric severity level of the alert: 1 (Low), 2 (Medium), 3 (High), 4 (Critical).
	Severity *int32 `json:"severity,omitempty"`
	// Date and time of the event that triggered the alert, as a Unix timestamp in milliseconds. If not provided, the current date is used.
	Date *int64 `json:"date,omitempty"`
	// Labels to categorize the alert. Unknown tags are automatically created. Use `POST /api/v1/query` with `listTag` to retrieve existing custom tags, or with `listTaxonomy` and `tags` to retrieve taxonomy tags.
	Tags []string `json:"tags,omitempty"`
	// Marks the alert as important for quick filtering.
	Flag *bool `json:"flag,omitempty"`
	// [Traffic Light Protocol (TLP)](https://www.misp-project.org/taxonomies.html#_tlp) level controlling how alert information can be shared: 0 (Clear — no restrictions), 1 (Green — community), 2 (Amber — organization and clients), 3 (Amber+Strict — organization only), 4 (Red — named recipients only).
	Tlp *int32 `json:"tlp,omitempty"`
	// [Permissible Actions Protocol (PAP)](https://www.misp-project.org/taxonomies.html#_pap) level controlling what actions can be taken with alert information: 0 (Clear — no restrictions), 1 (Green — active actions allowed), 2 (Amber — passive checks only), 3 (Red — non-detectable actions only).
	Pap          *int32                        `json:"pap,omitempty"`
	CustomFields *InputCreateAlertCustomFields `json:"customFields,omitempty"`
	// Brief summary of the triage findings. Supports [TheHive-flavored Markdown](https://docs.strangebee.com/thehive/user-guides/thehive-flavored-markdown/).
	Summary *string `json:"summary,omitempty"`
	// Name of the alert status to apply. Must match an existing status. Use `POST /api/v1/query` with `listAlertStatus` to list available statuses.
	Status *string `json:"status,omitempty"`
	// Login of the user to assign the alert to. Typically an email address. If a default domain is configured in the [authentication settings](https://docs.strangebee.com/thehive/administration/authentication/configure-authentication), a plain username is also accepted. Use `POST /api/v1/query` with `listUser` to retrieve available users and their logins. If omitted, the alert is unassigned.
	Assignee *string `json:"assignee,omitempty"`
	// Name of the case template to apply when creating a case from this alert. Applies the template's tasks, custom fields, and pages. Must match an existing case template. Use `POST /api/v1/query` with `listCaseTemplate` to list available case templates.
	CaseTemplate *string `json:"caseTemplate,omitempty"`
	// Observables to attach to the alert. Can also be added after creation with `POST /api/v1/alert/{alertId}/observable`.
	Observables []InputCreateObservable `json:"observables,omitempty"`
	// Tactics, techniques, and procedures (TTPs) to link to the alert based on the MITRE ATT&CK framework.
	Procedures []InputProcedure `json:"procedures,omitempty"`
}

type _InputCreateAlert InputCreateAlert

// NewInputCreateAlert instantiates a new InputCreateAlert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputCreateAlert(type_ string, source string, sourceRef string, title string, description string) *InputCreateAlert {
	this := InputCreateAlert{}
	this.Type = type_
	this.Source = source
	this.SourceRef = sourceRef
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

// NewInputCreateAlertWithDefaults instantiates a new InputCreateAlert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputCreateAlertWithDefaults() *InputCreateAlert {
	this := InputCreateAlert{}
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

// GetType returns the Type field value
func (o *InputCreateAlert) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InputCreateAlert) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InputCreateAlert) SetType(v string) {
	o.Type = v
}

// GetSource returns the Source field value
func (o *InputCreateAlert) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *InputCreateAlert) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *InputCreateAlert) SetSource(v string) {
	o.Source = v
}

// GetSourceRef returns the SourceRef field value
func (o *InputCreateAlert) GetSourceRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceRef
}

// GetSourceRefOk returns a tuple with the SourceRef field value
// and a boolean to check if the value has been set.
func (o *InputCreateAlert) GetSourceRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceRef, true
}

// SetSourceRef sets field value
func (o *InputCreateAlert) SetSourceRef(v string) {
	o.SourceRef = v
}

// GetExternalLink returns the ExternalLink field value if set, zero value otherwise.
func (o *InputCreateAlert) GetExternalLink() string {
	if o == nil || IsNil(o.ExternalLink) {
		var ret string
		return ret
	}
	return *o.ExternalLink
}

// GetExternalLinkOk returns a tuple with the ExternalLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateAlert) GetExternalLinkOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalLink) {
		return nil, false
	}
	return o.ExternalLink, true
}

// HasExternalLink returns a boolean if a field has been set.
func (o *InputCreateAlert) HasExternalLink() bool {
	if o != nil && !IsNil(o.ExternalLink) {
		return true
	}

	return false
}

// SetExternalLink gets a reference to the given string and assigns it to the ExternalLink field.
func (o *InputCreateAlert) SetExternalLink(v string) {
	o.ExternalLink = &v
}

// GetTitle returns the Title field value
func (o *InputCreateAlert) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *InputCreateAlert) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *InputCreateAlert) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value
func (o *InputCreateAlert) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *InputCreateAlert) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *InputCreateAlert) SetDescription(v string) {
	o.Description = v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *InputCreateAlert) GetSeverity() int32 {
	if o == nil || IsNil(o.Severity) {
		var ret int32
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateAlert) GetSeverityOk() (*int32, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *InputCreateAlert) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given int32 and assigns it to the Severity field.
func (o *InputCreateAlert) SetSeverity(v int32) {
	o.Severity = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *InputCreateAlert) GetDate() int64 {
	if o == nil || IsNil(o.Date) {
		var ret int64
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateAlert) GetDateOk() (*int64, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *InputCreateAlert) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given int64 and assigns it to the Date field.
func (o *InputCreateAlert) SetDate(v int64) {
	o.Date = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InputCreateAlert) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateAlert) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InputCreateAlert) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InputCreateAlert) SetTags(v []string) {
	o.Tags = v
}

// GetFlag returns the Flag field value if set, zero value otherwise.
func (o *InputCreateAlert) GetFlag() bool {
	if o == nil || IsNil(o.Flag) {
		var ret bool
		return ret
	}
	return *o.Flag
}

// GetFlagOk returns a tuple with the Flag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateAlert) GetFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.Flag) {
		return nil, false
	}
	return o.Flag, true
}

// HasFlag returns a boolean if a field has been set.
func (o *InputCreateAlert) HasFlag() bool {
	if o != nil && !IsNil(o.Flag) {
		return true
	}

	return false
}

// SetFlag gets a reference to the given bool and assigns it to the Flag field.
func (o *InputCreateAlert) SetFlag(v bool) {
	o.Flag = &v
}

// GetTlp returns the Tlp field value if set, zero value otherwise.
func (o *InputCreateAlert) GetTlp() int32 {
	if o == nil || IsNil(o.Tlp) {
		var ret int32
		return ret
	}
	return *o.Tlp
}

// GetTlpOk returns a tuple with the Tlp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateAlert) GetTlpOk() (*int32, bool) {
	if o == nil || IsNil(o.Tlp) {
		return nil, false
	}
	return o.Tlp, true
}

// HasTlp returns a boolean if a field has been set.
func (o *InputCreateAlert) HasTlp() bool {
	if o != nil && !IsNil(o.Tlp) {
		return true
	}

	return false
}

// SetTlp gets a reference to the given int32 and assigns it to the Tlp field.
func (o *InputCreateAlert) SetTlp(v int32) {
	o.Tlp = &v
}

// GetPap returns the Pap field value if set, zero value otherwise.
func (o *InputCreateAlert) GetPap() int32 {
	if o == nil || IsNil(o.Pap) {
		var ret int32
		return ret
	}
	return *o.Pap
}

// GetPapOk returns a tuple with the Pap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateAlert) GetPapOk() (*int32, bool) {
	if o == nil || IsNil(o.Pap) {
		return nil, false
	}
	return o.Pap, true
}

// HasPap returns a boolean if a field has been set.
func (o *InputCreateAlert) HasPap() bool {
	if o != nil && !IsNil(o.Pap) {
		return true
	}

	return false
}

// SetPap gets a reference to the given int32 and assigns it to the Pap field.
func (o *InputCreateAlert) SetPap(v int32) {
	o.Pap = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *InputCreateAlert) GetCustomFields() InputCreateAlertCustomFields {
	if o == nil || IsNil(o.CustomFields) {
		var ret InputCreateAlertCustomFields
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateAlert) GetCustomFieldsOk() (*InputCreateAlertCustomFields, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *InputCreateAlert) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given InputCreateAlertCustomFields and assigns it to the CustomFields field.
func (o *InputCreateAlert) SetCustomFields(v InputCreateAlertCustomFields) {
	o.CustomFields = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *InputCreateAlert) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateAlert) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *InputCreateAlert) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *InputCreateAlert) SetSummary(v string) {
	o.Summary = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InputCreateAlert) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateAlert) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InputCreateAlert) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InputCreateAlert) SetStatus(v string) {
	o.Status = &v
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *InputCreateAlert) GetAssignee() string {
	if o == nil || IsNil(o.Assignee) {
		var ret string
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateAlert) GetAssigneeOk() (*string, bool) {
	if o == nil || IsNil(o.Assignee) {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *InputCreateAlert) HasAssignee() bool {
	if o != nil && !IsNil(o.Assignee) {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given string and assigns it to the Assignee field.
func (o *InputCreateAlert) SetAssignee(v string) {
	o.Assignee = &v
}

// GetCaseTemplate returns the CaseTemplate field value if set, zero value otherwise.
func (o *InputCreateAlert) GetCaseTemplate() string {
	if o == nil || IsNil(o.CaseTemplate) {
		var ret string
		return ret
	}
	return *o.CaseTemplate
}

// GetCaseTemplateOk returns a tuple with the CaseTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateAlert) GetCaseTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.CaseTemplate) {
		return nil, false
	}
	return o.CaseTemplate, true
}

// HasCaseTemplate returns a boolean if a field has been set.
func (o *InputCreateAlert) HasCaseTemplate() bool {
	if o != nil && !IsNil(o.CaseTemplate) {
		return true
	}

	return false
}

// SetCaseTemplate gets a reference to the given string and assigns it to the CaseTemplate field.
func (o *InputCreateAlert) SetCaseTemplate(v string) {
	o.CaseTemplate = &v
}

// GetObservables returns the Observables field value if set, zero value otherwise.
func (o *InputCreateAlert) GetObservables() []InputCreateObservable {
	if o == nil || IsNil(o.Observables) {
		var ret []InputCreateObservable
		return ret
	}
	return o.Observables
}

// GetObservablesOk returns a tuple with the Observables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateAlert) GetObservablesOk() ([]InputCreateObservable, bool) {
	if o == nil || IsNil(o.Observables) {
		return nil, false
	}
	return o.Observables, true
}

// HasObservables returns a boolean if a field has been set.
func (o *InputCreateAlert) HasObservables() bool {
	if o != nil && !IsNil(o.Observables) {
		return true
	}

	return false
}

// SetObservables gets a reference to the given []InputCreateObservable and assigns it to the Observables field.
func (o *InputCreateAlert) SetObservables(v []InputCreateObservable) {
	o.Observables = v
}

// GetProcedures returns the Procedures field value if set, zero value otherwise.
func (o *InputCreateAlert) GetProcedures() []InputProcedure {
	if o == nil || IsNil(o.Procedures) {
		var ret []InputProcedure
		return ret
	}
	return o.Procedures
}

// GetProceduresOk returns a tuple with the Procedures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateAlert) GetProceduresOk() ([]InputProcedure, bool) {
	if o == nil || IsNil(o.Procedures) {
		return nil, false
	}
	return o.Procedures, true
}

// HasProcedures returns a boolean if a field has been set.
func (o *InputCreateAlert) HasProcedures() bool {
	if o != nil && !IsNil(o.Procedures) {
		return true
	}

	return false
}

// SetProcedures gets a reference to the given []InputProcedure and assigns it to the Procedures field.
func (o *InputCreateAlert) SetProcedures(v []InputProcedure) {
	o.Procedures = v
}

func (o InputCreateAlert) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputCreateAlert) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["source"] = o.Source
	toSerialize["sourceRef"] = o.SourceRef
	if !IsNil(o.ExternalLink) {
		toSerialize["externalLink"] = o.ExternalLink
	}
	toSerialize["title"] = o.Title
	toSerialize["description"] = o.Description
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
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
	if !IsNil(o.CustomFields) {
		toSerialize["customFields"] = o.CustomFields
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Assignee) {
		toSerialize["assignee"] = o.Assignee
	}
	if !IsNil(o.CaseTemplate) {
		toSerialize["caseTemplate"] = o.CaseTemplate
	}
	if !IsNil(o.Observables) {
		toSerialize["observables"] = o.Observables
	}
	if !IsNil(o.Procedures) {
		toSerialize["procedures"] = o.Procedures
	}
	return toSerialize, nil
}

func (o *InputCreateAlert) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"source",
		"sourceRef",
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

	varInputCreateAlert := _InputCreateAlert{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInputCreateAlert)

	if err != nil {
		return err
	}

	*o = InputCreateAlert(varInputCreateAlert)

	return err
}

type NullableInputCreateAlert struct {
	value *InputCreateAlert
	isSet bool
}

func (v NullableInputCreateAlert) Get() *InputCreateAlert {
	return v.value
}

func (v *NullableInputCreateAlert) Set(val *InputCreateAlert) {
	v.value = val
	v.isSet = true
}

func (v NullableInputCreateAlert) IsSet() bool {
	return v.isSet
}

func (v *NullableInputCreateAlert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputCreateAlert(val *InputCreateAlert) *NullableInputCreateAlert {
	return &NullableInputCreateAlert{value: val, isSet: true}
}

func (v NullableInputCreateAlert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputCreateAlert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
