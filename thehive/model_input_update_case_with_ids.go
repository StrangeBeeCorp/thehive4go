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

// checks if the InputUpdateCaseWithIds type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputUpdateCaseWithIds{}

// InputUpdateCaseWithIds struct for InputUpdateCaseWithIds
type InputUpdateCaseWithIds struct {
	// Identifiers of the entities to target. IDs must be prefixed with `~`. Case numbers and names are passed as plain strings without a prefix.
	Ids []string `json:"ids"`
	// New title of the case.
	Title *string `json:"title,omitempty"`
	// New details of the case. Supports [TheHive-flavored Markdown](https://docs.strangebee.com/thehive/user-guides/thehive-flavored-markdown/).
	Description *string `json:"description,omitempty"`
	// New severity level of the case: 1 (Low), 2 (Medium), 3 (High), 4 (Critical).
	Severity *int32 `json:"severity,omitempty"`
	// New start date of the incident. Expected format: Unix timestamp in milliseconds.
	StartDate *int64 `json:"startDate,omitempty"`
	// New end date of the incident. Expected format: Unix timestamp in milliseconds. Send `null` to unset.
	EndDate *int64 `json:"endDate,omitempty"`
	// Replaces all current tags with this set. Unknown tags are automatically created. Use `addTags` and `removeTags` to add or remove individual tags without replacing the full set. Use `POST /api/v1/query` with `listTag` to retrieve existing custom tags, or with `listTaxonomy` and `tags` to retrieve taxonomy tags.
	Tags []string `json:"tags,omitempty"`
	// Marks the case as important for quick filtering.
	Flag *bool `json:"flag,omitempty"`
	// New [Traffic Light Protocol (TLP)](https://www.misp-project.org/taxonomies.html#_tlp) level: 0 (Clear), 1 (Green), 2 (Amber), 3 (Amber+Strict), 4 (Red).
	Tlp *int32 `json:"tlp,omitempty"`
	// New [Permissible Actions Protocol (PAP)](https://www.misp-project.org/taxonomies.html#_pap) level: 0 (Clear), 1 (Green), 2 (Amber), 3 (Red).
	Pap *int32 `json:"pap,omitempty"`
	// New case status. Must match an existing status. To close a case, set this to a status linked to the *Closed* stage. Use `POST /api/v1/query` with `listCaseStatus` to list available statuses.
	Status *string `json:"status,omitempty"`
	// New investigation summary. Send `null` to unset.
	Summary *string `json:"summary,omitempty"`
	// Login of the user to assign the case to. Typically an email address. If a default domain is configured in the [authentication settings](https://docs.strangebee.com/thehive/administration/authentication/configure-authentication), a plain username is also accepted. Use `POST /api/v1/query` with `listUser` to retrieve available users and their logins. Send `null` to unassign.
	Assignee *string `json:"assignee,omitempty"`
	// Enum: `NoImpact` `WithImpact` `NotApplicable`<br>Verdict on whether the incident had real consequences for the organization. Required when you close a case with a status whose `requiresImpact` field is `true`. TheHive rejects the update if you send it while closing a case with a status that doesn't require an impact. Use `POST /api/v1/query` with `listCaseStatus` to check which statuses require an impact. `WithImpact` means the incident caused damage or data loss. `NoImpact` means it was resolved without consequences. `NotApplicable` means the concept doesn't apply to this case. Send `null` to unset.
	ImpactStatus *string                      `json:"impactStatus,omitempty"`
	CustomFields *InputUpdateCaseCustomFields `json:"customFields,omitempty"`
	// Enum: `autoShare` `manual`<br>New task-sharing rule for the owner organization: `autoShare` shares new tasks with every sharing organization automatically, `manual` requires explicit sharing per task. Overrides the [organization-level sharing rules](https://docs.strangebee.com/thehive/administration/organizations/about-organizations-sharing-rules/).
	TaskRule *string `json:"taskRule,omitempty"`
	// Enum: `autoShare` `manual`<br>New observable-sharing rule for the owner organization: `autoShare` shares new observables with every sharing organization automatically, `manual` requires explicit sharing per observable. Overrides the [organization-level sharing rules](https://docs.strangebee.com/thehive/administration/organizations/about-organizations-sharing-rules/).
	ObservableRule *string `json:"observableRule,omitempty"`
	// Tags to add to the current set of case tags. Unknown tags are automatically created. Use `POST /api/v1/query` with `listTag` to retrieve existing custom tags, or with `listTaxonomy` and `tags` to retrieve taxonomy tags.
	AddTags []string `json:"addTags,omitempty"`
	// Tags to remove from the current set of case tags. Tags not present on the case are silently ignored. Use `GET /api/v1/case/{idOrName}` and read the `tags` field to see the current tags on the case.
	RemoveTags []string `json:"removeTags,omitempty"`
}

type _InputUpdateCaseWithIds InputUpdateCaseWithIds

// NewInputUpdateCaseWithIds instantiates a new InputUpdateCaseWithIds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputUpdateCaseWithIds(ids []string) *InputUpdateCaseWithIds {
	this := InputUpdateCaseWithIds{}
	this.Ids = ids
	return &this
}

// NewInputUpdateCaseWithIdsWithDefaults instantiates a new InputUpdateCaseWithIds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputUpdateCaseWithIdsWithDefaults() *InputUpdateCaseWithIds {
	this := InputUpdateCaseWithIds{}
	return &this
}

// GetIds returns the Ids field value
func (o *InputUpdateCaseWithIds) GetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *InputUpdateCaseWithIds) GetIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ids, true
}

// SetIds sets field value
func (o *InputUpdateCaseWithIds) SetIds(v []string) {
	o.Ids = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *InputUpdateCaseWithIds) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateCaseWithIds) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *InputUpdateCaseWithIds) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *InputUpdateCaseWithIds) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InputUpdateCaseWithIds) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateCaseWithIds) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InputUpdateCaseWithIds) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InputUpdateCaseWithIds) SetDescription(v string) {
	o.Description = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *InputUpdateCaseWithIds) GetSeverity() int32 {
	if o == nil || IsNil(o.Severity) {
		var ret int32
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateCaseWithIds) GetSeverityOk() (*int32, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *InputUpdateCaseWithIds) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given int32 and assigns it to the Severity field.
func (o *InputUpdateCaseWithIds) SetSeverity(v int32) {
	o.Severity = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *InputUpdateCaseWithIds) GetStartDate() int64 {
	if o == nil || IsNil(o.StartDate) {
		var ret int64
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateCaseWithIds) GetStartDateOk() (*int64, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *InputUpdateCaseWithIds) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given int64 and assigns it to the StartDate field.
func (o *InputUpdateCaseWithIds) SetStartDate(v int64) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *InputUpdateCaseWithIds) GetEndDate() int64 {
	if o == nil || IsNil(o.EndDate) {
		var ret int64
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateCaseWithIds) GetEndDateOk() (*int64, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *InputUpdateCaseWithIds) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given int64 and assigns it to the EndDate field.
func (o *InputUpdateCaseWithIds) SetEndDate(v int64) {
	o.EndDate = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InputUpdateCaseWithIds) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateCaseWithIds) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InputUpdateCaseWithIds) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InputUpdateCaseWithIds) SetTags(v []string) {
	o.Tags = v
}

// GetFlag returns the Flag field value if set, zero value otherwise.
func (o *InputUpdateCaseWithIds) GetFlag() bool {
	if o == nil || IsNil(o.Flag) {
		var ret bool
		return ret
	}
	return *o.Flag
}

// GetFlagOk returns a tuple with the Flag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateCaseWithIds) GetFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.Flag) {
		return nil, false
	}
	return o.Flag, true
}

// HasFlag returns a boolean if a field has been set.
func (o *InputUpdateCaseWithIds) HasFlag() bool {
	if o != nil && !IsNil(o.Flag) {
		return true
	}

	return false
}

// SetFlag gets a reference to the given bool and assigns it to the Flag field.
func (o *InputUpdateCaseWithIds) SetFlag(v bool) {
	o.Flag = &v
}

// GetTlp returns the Tlp field value if set, zero value otherwise.
func (o *InputUpdateCaseWithIds) GetTlp() int32 {
	if o == nil || IsNil(o.Tlp) {
		var ret int32
		return ret
	}
	return *o.Tlp
}

// GetTlpOk returns a tuple with the Tlp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateCaseWithIds) GetTlpOk() (*int32, bool) {
	if o == nil || IsNil(o.Tlp) {
		return nil, false
	}
	return o.Tlp, true
}

// HasTlp returns a boolean if a field has been set.
func (o *InputUpdateCaseWithIds) HasTlp() bool {
	if o != nil && !IsNil(o.Tlp) {
		return true
	}

	return false
}

// SetTlp gets a reference to the given int32 and assigns it to the Tlp field.
func (o *InputUpdateCaseWithIds) SetTlp(v int32) {
	o.Tlp = &v
}

// GetPap returns the Pap field value if set, zero value otherwise.
func (o *InputUpdateCaseWithIds) GetPap() int32 {
	if o == nil || IsNil(o.Pap) {
		var ret int32
		return ret
	}
	return *o.Pap
}

// GetPapOk returns a tuple with the Pap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateCaseWithIds) GetPapOk() (*int32, bool) {
	if o == nil || IsNil(o.Pap) {
		return nil, false
	}
	return o.Pap, true
}

// HasPap returns a boolean if a field has been set.
func (o *InputUpdateCaseWithIds) HasPap() bool {
	if o != nil && !IsNil(o.Pap) {
		return true
	}

	return false
}

// SetPap gets a reference to the given int32 and assigns it to the Pap field.
func (o *InputUpdateCaseWithIds) SetPap(v int32) {
	o.Pap = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InputUpdateCaseWithIds) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateCaseWithIds) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InputUpdateCaseWithIds) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InputUpdateCaseWithIds) SetStatus(v string) {
	o.Status = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *InputUpdateCaseWithIds) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateCaseWithIds) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *InputUpdateCaseWithIds) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *InputUpdateCaseWithIds) SetSummary(v string) {
	o.Summary = &v
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *InputUpdateCaseWithIds) GetAssignee() string {
	if o == nil || IsNil(o.Assignee) {
		var ret string
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateCaseWithIds) GetAssigneeOk() (*string, bool) {
	if o == nil || IsNil(o.Assignee) {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *InputUpdateCaseWithIds) HasAssignee() bool {
	if o != nil && !IsNil(o.Assignee) {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given string and assigns it to the Assignee field.
func (o *InputUpdateCaseWithIds) SetAssignee(v string) {
	o.Assignee = &v
}

// GetImpactStatus returns the ImpactStatus field value if set, zero value otherwise.
func (o *InputUpdateCaseWithIds) GetImpactStatus() string {
	if o == nil || IsNil(o.ImpactStatus) {
		var ret string
		return ret
	}
	return *o.ImpactStatus
}

// GetImpactStatusOk returns a tuple with the ImpactStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateCaseWithIds) GetImpactStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ImpactStatus) {
		return nil, false
	}
	return o.ImpactStatus, true
}

// HasImpactStatus returns a boolean if a field has been set.
func (o *InputUpdateCaseWithIds) HasImpactStatus() bool {
	if o != nil && !IsNil(o.ImpactStatus) {
		return true
	}

	return false
}

// SetImpactStatus gets a reference to the given string and assigns it to the ImpactStatus field.
func (o *InputUpdateCaseWithIds) SetImpactStatus(v string) {
	o.ImpactStatus = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *InputUpdateCaseWithIds) GetCustomFields() InputUpdateCaseCustomFields {
	if o == nil || IsNil(o.CustomFields) {
		var ret InputUpdateCaseCustomFields
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateCaseWithIds) GetCustomFieldsOk() (*InputUpdateCaseCustomFields, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *InputUpdateCaseWithIds) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given InputUpdateCaseCustomFields and assigns it to the CustomFields field.
func (o *InputUpdateCaseWithIds) SetCustomFields(v InputUpdateCaseCustomFields) {
	o.CustomFields = &v
}

// GetTaskRule returns the TaskRule field value if set, zero value otherwise.
func (o *InputUpdateCaseWithIds) GetTaskRule() string {
	if o == nil || IsNil(o.TaskRule) {
		var ret string
		return ret
	}
	return *o.TaskRule
}

// GetTaskRuleOk returns a tuple with the TaskRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateCaseWithIds) GetTaskRuleOk() (*string, bool) {
	if o == nil || IsNil(o.TaskRule) {
		return nil, false
	}
	return o.TaskRule, true
}

// HasTaskRule returns a boolean if a field has been set.
func (o *InputUpdateCaseWithIds) HasTaskRule() bool {
	if o != nil && !IsNil(o.TaskRule) {
		return true
	}

	return false
}

// SetTaskRule gets a reference to the given string and assigns it to the TaskRule field.
func (o *InputUpdateCaseWithIds) SetTaskRule(v string) {
	o.TaskRule = &v
}

// GetObservableRule returns the ObservableRule field value if set, zero value otherwise.
func (o *InputUpdateCaseWithIds) GetObservableRule() string {
	if o == nil || IsNil(o.ObservableRule) {
		var ret string
		return ret
	}
	return *o.ObservableRule
}

// GetObservableRuleOk returns a tuple with the ObservableRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateCaseWithIds) GetObservableRuleOk() (*string, bool) {
	if o == nil || IsNil(o.ObservableRule) {
		return nil, false
	}
	return o.ObservableRule, true
}

// HasObservableRule returns a boolean if a field has been set.
func (o *InputUpdateCaseWithIds) HasObservableRule() bool {
	if o != nil && !IsNil(o.ObservableRule) {
		return true
	}

	return false
}

// SetObservableRule gets a reference to the given string and assigns it to the ObservableRule field.
func (o *InputUpdateCaseWithIds) SetObservableRule(v string) {
	o.ObservableRule = &v
}

// GetAddTags returns the AddTags field value if set, zero value otherwise.
func (o *InputUpdateCaseWithIds) GetAddTags() []string {
	if o == nil || IsNil(o.AddTags) {
		var ret []string
		return ret
	}
	return o.AddTags
}

// GetAddTagsOk returns a tuple with the AddTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateCaseWithIds) GetAddTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.AddTags) {
		return nil, false
	}
	return o.AddTags, true
}

// HasAddTags returns a boolean if a field has been set.
func (o *InputUpdateCaseWithIds) HasAddTags() bool {
	if o != nil && !IsNil(o.AddTags) {
		return true
	}

	return false
}

// SetAddTags gets a reference to the given []string and assigns it to the AddTags field.
func (o *InputUpdateCaseWithIds) SetAddTags(v []string) {
	o.AddTags = v
}

// GetRemoveTags returns the RemoveTags field value if set, zero value otherwise.
func (o *InputUpdateCaseWithIds) GetRemoveTags() []string {
	if o == nil || IsNil(o.RemoveTags) {
		var ret []string
		return ret
	}
	return o.RemoveTags
}

// GetRemoveTagsOk returns a tuple with the RemoveTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateCaseWithIds) GetRemoveTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.RemoveTags) {
		return nil, false
	}
	return o.RemoveTags, true
}

// HasRemoveTags returns a boolean if a field has been set.
func (o *InputUpdateCaseWithIds) HasRemoveTags() bool {
	if o != nil && !IsNil(o.RemoveTags) {
		return true
	}

	return false
}

// SetRemoveTags gets a reference to the given []string and assigns it to the RemoveTags field.
func (o *InputUpdateCaseWithIds) SetRemoveTags(v []string) {
	o.RemoveTags = v
}

func (o InputUpdateCaseWithIds) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputUpdateCaseWithIds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ids"] = o.Ids
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
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
	if !IsNil(o.ImpactStatus) {
		toSerialize["impactStatus"] = o.ImpactStatus
	}
	if !IsNil(o.CustomFields) {
		toSerialize["customFields"] = o.CustomFields
	}
	if !IsNil(o.TaskRule) {
		toSerialize["taskRule"] = o.TaskRule
	}
	if !IsNil(o.ObservableRule) {
		toSerialize["observableRule"] = o.ObservableRule
	}
	if !IsNil(o.AddTags) {
		toSerialize["addTags"] = o.AddTags
	}
	if !IsNil(o.RemoveTags) {
		toSerialize["removeTags"] = o.RemoveTags
	}
	return toSerialize, nil
}

func (o *InputUpdateCaseWithIds) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ids",
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

	varInputUpdateCaseWithIds := _InputUpdateCaseWithIds{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInputUpdateCaseWithIds)

	if err != nil {
		return err
	}

	*o = InputUpdateCaseWithIds(varInputUpdateCaseWithIds)

	return err
}

type NullableInputUpdateCaseWithIds struct {
	value *InputUpdateCaseWithIds
	isSet bool
}

func (v NullableInputUpdateCaseWithIds) Get() *InputUpdateCaseWithIds {
	return v.value
}

func (v *NullableInputUpdateCaseWithIds) Set(val *InputUpdateCaseWithIds) {
	v.value = val
	v.isSet = true
}

func (v NullableInputUpdateCaseWithIds) IsSet() bool {
	return v.isSet
}

func (v *NullableInputUpdateCaseWithIds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputUpdateCaseWithIds(val *InputUpdateCaseWithIds) *NullableInputUpdateCaseWithIds {
	return &NullableInputUpdateCaseWithIds{value: val, isSet: true}
}

func (v NullableInputUpdateCaseWithIds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputUpdateCaseWithIds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
