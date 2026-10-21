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

// checks if the InputUpdateAlert type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputUpdateAlert{}

// InputUpdateAlert struct for InputUpdateAlert
type InputUpdateAlert struct {
	// Type of the alert. Identifies the event category and groups related alerts.
	Type *string `json:"type,omitempty"`
	// Source system that generated the alert.
	Source *string `json:"source,omitempty"`
	// Unique reference identifier from the source system. Combined with `type` and `source`, uniquely identifies an alert.
	SourceRef *string `json:"sourceRef,omitempty"`
	// External URL linking to the alert in the source system. Send `null` to unset.
	ExternalLink *string `json:"externalLink,omitempty"`
	// New title of the alert.
	Title *string `json:"title,omitempty"`
	// New description of the alert. Supports [TheHive-flavored Markdown](https://docs.strangebee.com/thehive/user-guides/thehive-flavored-markdown/).
	Description *string `json:"description,omitempty"`
	// New severity level of the alert: 1 (Low), 2 (Medium), 3 (High), 4 (Critical).
	Severity *int32 `json:"severity,omitempty"`
	// New date and time of the event that triggered the alert, as a Unix timestamp in milliseconds.
	Date *int64 `json:"date,omitempty"`
	// Date and time of the last synchronization with the source system, as a Unix timestamp in milliseconds.
	LastSyncDate *int64 `json:"lastSyncDate,omitempty"`
	// Replaces all current tags with this set. Unknown tags are automatically created. Use `addTags` and `removeTags` to add or remove individual tags without replacing the full set. Use `POST /api/v1/query` with `listTag` to retrieve existing custom tags, or with `listTaxonomy` and `tags` to retrieve taxonomy tags.
	Tags []string `json:"tags,omitempty"`
	// New [Traffic Light Protocol (TLP)](https://www.misp-project.org/taxonomies.html#_tlp) level: 0 (Clear), 1 (Green), 2 (Amber), 3 (Amber+Strict), 4 (Red).
	Tlp *int32 `json:"tlp,omitempty"`
	// New [Permissible Actions Protocol (PAP)](https://www.misp-project.org/taxonomies.html#_pap) level: 0 (Clear), 1 (Green), 2 (Amber), 3 (Red).
	Pap *int32 `json:"pap,omitempty"`
	// Whether to follow the alert for change notifications.
	Follow       *bool                         `json:"follow,omitempty"`
	CustomFields *InputUpdateAlertCustomFields `json:"customFields,omitempty"`
	// New alert status. Must match an existing status. To close an alert, set this to a status linked to the *Closed* stage. Use `POST /api/v1/query` with `listAlertStatus` to list available statuses.
	Status *string `json:"status,omitempty"`
	// New summary of the triage findings. Send `null` to unset.
	Summary *string `json:"summary,omitempty"`
	// Login of the user to assign the alert to. Typically an email address. If a default domain is configured in the [authentication settings](https://docs.strangebee.com/thehive/administration/authentication/configure-authentication), a plain username is also accepted. Use `POST /api/v1/query` with `listUser` to retrieve available users and their logins. Send `null` to unassign.
	Assignee *string `json:"assignee,omitempty"`
	// Tags to add to the current set of alert tags. Unknown tags are automatically created. Use `POST /api/v1/query` with `listTag` to retrieve existing custom tags, or with `listTaxonomy` and `tags` to retrieve taxonomy tags.
	AddTags []string `json:"addTags,omitempty"`
	// Tags to remove from the current set of alert tags. Tags not present on the alert are silently ignored. Use `GET /api/v1/alert/{alertId}` and read the `tags` field to see the current tags on the alert.
	RemoveTags []string `json:"removeTags,omitempty"`
}

// NewInputUpdateAlert instantiates a new InputUpdateAlert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputUpdateAlert() *InputUpdateAlert {
	this := InputUpdateAlert{}
	return &this
}

// NewInputUpdateAlertWithDefaults instantiates a new InputUpdateAlert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputUpdateAlertWithDefaults() *InputUpdateAlert {
	this := InputUpdateAlert{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InputUpdateAlert) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateAlert) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InputUpdateAlert) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InputUpdateAlert) SetType(v string) {
	o.Type = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *InputUpdateAlert) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateAlert) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *InputUpdateAlert) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *InputUpdateAlert) SetSource(v string) {
	o.Source = &v
}

// GetSourceRef returns the SourceRef field value if set, zero value otherwise.
func (o *InputUpdateAlert) GetSourceRef() string {
	if o == nil || IsNil(o.SourceRef) {
		var ret string
		return ret
	}
	return *o.SourceRef
}

// GetSourceRefOk returns a tuple with the SourceRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateAlert) GetSourceRefOk() (*string, bool) {
	if o == nil || IsNil(o.SourceRef) {
		return nil, false
	}
	return o.SourceRef, true
}

// HasSourceRef returns a boolean if a field has been set.
func (o *InputUpdateAlert) HasSourceRef() bool {
	if o != nil && !IsNil(o.SourceRef) {
		return true
	}

	return false
}

// SetSourceRef gets a reference to the given string and assigns it to the SourceRef field.
func (o *InputUpdateAlert) SetSourceRef(v string) {
	o.SourceRef = &v
}

// GetExternalLink returns the ExternalLink field value if set, zero value otherwise.
func (o *InputUpdateAlert) GetExternalLink() string {
	if o == nil || IsNil(o.ExternalLink) {
		var ret string
		return ret
	}
	return *o.ExternalLink
}

// GetExternalLinkOk returns a tuple with the ExternalLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateAlert) GetExternalLinkOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalLink) {
		return nil, false
	}
	return o.ExternalLink, true
}

// HasExternalLink returns a boolean if a field has been set.
func (o *InputUpdateAlert) HasExternalLink() bool {
	if o != nil && !IsNil(o.ExternalLink) {
		return true
	}

	return false
}

// SetExternalLink gets a reference to the given string and assigns it to the ExternalLink field.
func (o *InputUpdateAlert) SetExternalLink(v string) {
	o.ExternalLink = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *InputUpdateAlert) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateAlert) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *InputUpdateAlert) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *InputUpdateAlert) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InputUpdateAlert) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateAlert) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InputUpdateAlert) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InputUpdateAlert) SetDescription(v string) {
	o.Description = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *InputUpdateAlert) GetSeverity() int32 {
	if o == nil || IsNil(o.Severity) {
		var ret int32
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateAlert) GetSeverityOk() (*int32, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *InputUpdateAlert) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given int32 and assigns it to the Severity field.
func (o *InputUpdateAlert) SetSeverity(v int32) {
	o.Severity = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *InputUpdateAlert) GetDate() int64 {
	if o == nil || IsNil(o.Date) {
		var ret int64
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateAlert) GetDateOk() (*int64, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *InputUpdateAlert) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given int64 and assigns it to the Date field.
func (o *InputUpdateAlert) SetDate(v int64) {
	o.Date = &v
}

// GetLastSyncDate returns the LastSyncDate field value if set, zero value otherwise.
func (o *InputUpdateAlert) GetLastSyncDate() int64 {
	if o == nil || IsNil(o.LastSyncDate) {
		var ret int64
		return ret
	}
	return *o.LastSyncDate
}

// GetLastSyncDateOk returns a tuple with the LastSyncDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateAlert) GetLastSyncDateOk() (*int64, bool) {
	if o == nil || IsNil(o.LastSyncDate) {
		return nil, false
	}
	return o.LastSyncDate, true
}

// HasLastSyncDate returns a boolean if a field has been set.
func (o *InputUpdateAlert) HasLastSyncDate() bool {
	if o != nil && !IsNil(o.LastSyncDate) {
		return true
	}

	return false
}

// SetLastSyncDate gets a reference to the given int64 and assigns it to the LastSyncDate field.
func (o *InputUpdateAlert) SetLastSyncDate(v int64) {
	o.LastSyncDate = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InputUpdateAlert) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateAlert) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InputUpdateAlert) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InputUpdateAlert) SetTags(v []string) {
	o.Tags = v
}

// GetTlp returns the Tlp field value if set, zero value otherwise.
func (o *InputUpdateAlert) GetTlp() int32 {
	if o == nil || IsNil(o.Tlp) {
		var ret int32
		return ret
	}
	return *o.Tlp
}

// GetTlpOk returns a tuple with the Tlp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateAlert) GetTlpOk() (*int32, bool) {
	if o == nil || IsNil(o.Tlp) {
		return nil, false
	}
	return o.Tlp, true
}

// HasTlp returns a boolean if a field has been set.
func (o *InputUpdateAlert) HasTlp() bool {
	if o != nil && !IsNil(o.Tlp) {
		return true
	}

	return false
}

// SetTlp gets a reference to the given int32 and assigns it to the Tlp field.
func (o *InputUpdateAlert) SetTlp(v int32) {
	o.Tlp = &v
}

// GetPap returns the Pap field value if set, zero value otherwise.
func (o *InputUpdateAlert) GetPap() int32 {
	if o == nil || IsNil(o.Pap) {
		var ret int32
		return ret
	}
	return *o.Pap
}

// GetPapOk returns a tuple with the Pap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateAlert) GetPapOk() (*int32, bool) {
	if o == nil || IsNil(o.Pap) {
		return nil, false
	}
	return o.Pap, true
}

// HasPap returns a boolean if a field has been set.
func (o *InputUpdateAlert) HasPap() bool {
	if o != nil && !IsNil(o.Pap) {
		return true
	}

	return false
}

// SetPap gets a reference to the given int32 and assigns it to the Pap field.
func (o *InputUpdateAlert) SetPap(v int32) {
	o.Pap = &v
}

// GetFollow returns the Follow field value if set, zero value otherwise.
func (o *InputUpdateAlert) GetFollow() bool {
	if o == nil || IsNil(o.Follow) {
		var ret bool
		return ret
	}
	return *o.Follow
}

// GetFollowOk returns a tuple with the Follow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateAlert) GetFollowOk() (*bool, bool) {
	if o == nil || IsNil(o.Follow) {
		return nil, false
	}
	return o.Follow, true
}

// HasFollow returns a boolean if a field has been set.
func (o *InputUpdateAlert) HasFollow() bool {
	if o != nil && !IsNil(o.Follow) {
		return true
	}

	return false
}

// SetFollow gets a reference to the given bool and assigns it to the Follow field.
func (o *InputUpdateAlert) SetFollow(v bool) {
	o.Follow = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *InputUpdateAlert) GetCustomFields() InputUpdateAlertCustomFields {
	if o == nil || IsNil(o.CustomFields) {
		var ret InputUpdateAlertCustomFields
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateAlert) GetCustomFieldsOk() (*InputUpdateAlertCustomFields, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *InputUpdateAlert) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given InputUpdateAlertCustomFields and assigns it to the CustomFields field.
func (o *InputUpdateAlert) SetCustomFields(v InputUpdateAlertCustomFields) {
	o.CustomFields = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InputUpdateAlert) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateAlert) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InputUpdateAlert) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InputUpdateAlert) SetStatus(v string) {
	o.Status = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *InputUpdateAlert) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateAlert) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *InputUpdateAlert) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *InputUpdateAlert) SetSummary(v string) {
	o.Summary = &v
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *InputUpdateAlert) GetAssignee() string {
	if o == nil || IsNil(o.Assignee) {
		var ret string
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateAlert) GetAssigneeOk() (*string, bool) {
	if o == nil || IsNil(o.Assignee) {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *InputUpdateAlert) HasAssignee() bool {
	if o != nil && !IsNil(o.Assignee) {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given string and assigns it to the Assignee field.
func (o *InputUpdateAlert) SetAssignee(v string) {
	o.Assignee = &v
}

// GetAddTags returns the AddTags field value if set, zero value otherwise.
func (o *InputUpdateAlert) GetAddTags() []string {
	if o == nil || IsNil(o.AddTags) {
		var ret []string
		return ret
	}
	return o.AddTags
}

// GetAddTagsOk returns a tuple with the AddTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateAlert) GetAddTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.AddTags) {
		return nil, false
	}
	return o.AddTags, true
}

// HasAddTags returns a boolean if a field has been set.
func (o *InputUpdateAlert) HasAddTags() bool {
	if o != nil && !IsNil(o.AddTags) {
		return true
	}

	return false
}

// SetAddTags gets a reference to the given []string and assigns it to the AddTags field.
func (o *InputUpdateAlert) SetAddTags(v []string) {
	o.AddTags = v
}

// GetRemoveTags returns the RemoveTags field value if set, zero value otherwise.
func (o *InputUpdateAlert) GetRemoveTags() []string {
	if o == nil || IsNil(o.RemoveTags) {
		var ret []string
		return ret
	}
	return o.RemoveTags
}

// GetRemoveTagsOk returns a tuple with the RemoveTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateAlert) GetRemoveTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.RemoveTags) {
		return nil, false
	}
	return o.RemoveTags, true
}

// HasRemoveTags returns a boolean if a field has been set.
func (o *InputUpdateAlert) HasRemoveTags() bool {
	if o != nil && !IsNil(o.RemoveTags) {
		return true
	}

	return false
}

// SetRemoveTags gets a reference to the given []string and assigns it to the RemoveTags field.
func (o *InputUpdateAlert) SetRemoveTags(v []string) {
	o.RemoveTags = v
}

func (o InputUpdateAlert) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputUpdateAlert) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.SourceRef) {
		toSerialize["sourceRef"] = o.SourceRef
	}
	if !IsNil(o.ExternalLink) {
		toSerialize["externalLink"] = o.ExternalLink
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.LastSyncDate) {
		toSerialize["lastSyncDate"] = o.LastSyncDate
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Tlp) {
		toSerialize["tlp"] = o.Tlp
	}
	if !IsNil(o.Pap) {
		toSerialize["pap"] = o.Pap
	}
	if !IsNil(o.Follow) {
		toSerialize["follow"] = o.Follow
	}
	if !IsNil(o.CustomFields) {
		toSerialize["customFields"] = o.CustomFields
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
	if !IsNil(o.AddTags) {
		toSerialize["addTags"] = o.AddTags
	}
	if !IsNil(o.RemoveTags) {
		toSerialize["removeTags"] = o.RemoveTags
	}
	return toSerialize, nil
}

type NullableInputUpdateAlert struct {
	value *InputUpdateAlert
	isSet bool
}

func (v NullableInputUpdateAlert) Get() *InputUpdateAlert {
	return v.value
}

func (v *NullableInputUpdateAlert) Set(val *InputUpdateAlert) {
	v.value = val
	v.isSet = true
}

func (v NullableInputUpdateAlert) IsSet() bool {
	return v.isSet
}

func (v *NullableInputUpdateAlert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputUpdateAlert(val *InputUpdateAlert) *NullableInputUpdateAlert {
	return &NullableInputUpdateAlert{value: val, isSet: true}
}

func (v NullableInputUpdateAlert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputUpdateAlert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
