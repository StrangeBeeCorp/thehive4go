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

// checks if the OutputAlert type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutputAlert{}

// OutputAlert struct for OutputAlert
type OutputAlert struct {
	// Internal unique identifier of the alert.
	UnderscoreId string `json:"_id"`
	// Entity type, always `Alert`.
	UnderscoreType string `json:"_type"`
	// Login of the user who created the alert.
	UnderscoreCreatedBy string `json:"_createdBy"`
	// Login of the user who last updated the alert.
	UnderscoreUpdatedBy *string `json:"_updatedBy,omitempty"`
	// Date and time when the alert was created. Returned as a Unix timestamp in milliseconds.
	UnderscoreCreatedAt int64 `json:"_createdAt"`
	// Date and time when the alert was last updated. Returned as a Unix timestamp in milliseconds.
	UnderscoreUpdatedAt *int64 `json:"_updatedAt,omitempty"`
	// Type of the alert.
	Type string `json:"type"`
	// Source system that generated the alert.
	Source string `json:"source"`
	// Unique reference identifier from the source system.
	SourceRef string `json:"sourceRef"`
	// External URL linking to the alert in the source system.
	ExternalLink *string `json:"externalLink,omitempty"`
	// Short title describing the alert.
	Title string `json:"title"`
	// Full description of the alert.
	Description string `json:"description"`
	// Numeric severity level of the alert: 1 (Low), 2 (Medium), 3 (High), 4 (Critical).
	Severity int32 `json:"severity"`
	// Human-readable severity label.
	SeverityLabel string `json:"severityLabel"`
	// Date and time of the event that triggered the alert, as a Unix timestamp in milliseconds.
	Date int64 `json:"date"`
	// Tags assigned to the alert.
	Tags []string `json:"tags,omitempty"`
	// Numeric Traffic Light Protocol (TLP) level: 0 (Clear), 1 (Green), 2 (Amber), 3 (Amber+Strict), 4 (Red).
	Tlp int32 `json:"tlp"`
	// Human-readable TLP label.
	TlpLabel string `json:"tlpLabel"`
	// Numeric Permissible Actions Protocol (PAP) level: 0 (Clear), 1 (Green), 2 (Amber), 3 (Red).
	Pap int32 `json:"pap"`
	// Human-readable PAP label.
	PapLabel string `json:"papLabel"`
	// Whether the alert is being followed for change notifications.
	Follow bool `json:"follow"`
	// Custom field values attached to the alert.
	CustomFields []OutputCustomFieldValue `json:"customFields,omitempty"`
	// Name of the case template used when creating a case from this alert.
	CaseTemplate *string `json:"caseTemplate,omitempty"`
	// Number of observables attached to the alert.
	ObservableCount int64 `json:"observableCount"`
	// ID of the case created from or linked to this alert.
	CaseId *string `json:"caseId,omitempty"`
	// Current status of the alert.
	Status string `json:"status"`
	// Stage derived from the alert status.
	Stage string `json:"stage"`
	// Login of the user assigned to the alert.
	Assignee *string `json:"assignee,omitempty"`
	// Brief summary of the triage findings.
	Summary *string `json:"summary,omitempty"`
	// Extra data fields populated when specific keys are requested in the `extraData` query parameter.
	ExtraData map[string]interface{} `json:"extraData"`
	// Date when the alert first entered a status linked to the *New* stage. Returned as a Unix timestamp in milliseconds.
	NewDate int64 `json:"newDate"`
	// Date when the alert first entered a status linked to the *In Progress* stage. Returned as a Unix timestamp in milliseconds.
	InProgressDate *int64 `json:"inProgressDate,omitempty"`
	// Date when the alert first entered a status linked to the *Closed* stage because no further investigation was required. Returned as a Unix timestamp in milliseconds.
	ClosedDate *int64 `json:"closedDate,omitempty"`
	// Date when the alert first entered the *Imported* status. This occurs when the alert is merged into a new or existing case for investigation. Returned as a Unix timestamp in milliseconds.
	ImportedDate *int64 `json:"importedDate,omitempty"`
	// [Time to Detect (TTD)](https://docs.strangebee.com/thehive/user-guides/key-performance-indicators/key-performance-indicators/#time-to-detect-ttd), in milliseconds.
	TimeToDetect int64 `json:"timeToDetect"`
	// [Time to Triage (TTT)](https://docs.strangebee.com/thehive/user-guides/key-performance-indicators/key-performance-indicators/#time-to-triage-ttt), in milliseconds.
	TimeToTriage *int64 `json:"timeToTriage,omitempty"`
	// [Time to Qualify (TTQ)](https://docs.strangebee.com/thehive/user-guides/key-performance-indicators/key-performance-indicators/#time-to-qualify-ttq), in milliseconds.
	TimeToQualify *int64 `json:"timeToQualify,omitempty"`
	// [Time to Acknowledge (TTA)](https://docs.strangebee.com/thehive/user-guides/key-performance-indicators/key-performance-indicators/#time-to-acknowledge-tta), in milliseconds.
	TimeToAcknowledge *int64 `json:"timeToAcknowledge,omitempty"`
}

type _OutputAlert OutputAlert

// NewOutputAlert instantiates a new OutputAlert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputAlert(underscoreId string, underscoreType string, underscoreCreatedBy string, underscoreCreatedAt int64, type_ string, source string, sourceRef string, title string, description string, severity int32, severityLabel string, date int64, tlp int32, tlpLabel string, pap int32, papLabel string, follow bool, observableCount int64, status string, stage string, extraData map[string]interface{}, newDate int64, timeToDetect int64) *OutputAlert {
	this := OutputAlert{}
	this.UnderscoreId = underscoreId
	this.UnderscoreType = underscoreType
	this.UnderscoreCreatedBy = underscoreCreatedBy
	this.UnderscoreCreatedAt = underscoreCreatedAt
	this.Type = type_
	this.Source = source
	this.SourceRef = sourceRef
	this.Title = title
	this.Description = description
	this.Severity = severity
	this.SeverityLabel = severityLabel
	this.Date = date
	this.Tlp = tlp
	this.TlpLabel = tlpLabel
	this.Pap = pap
	this.PapLabel = papLabel
	this.Follow = follow
	this.ObservableCount = observableCount
	this.Status = status
	this.Stage = stage
	this.ExtraData = extraData
	this.NewDate = newDate
	this.TimeToDetect = timeToDetect
	return &this
}

// NewOutputAlertWithDefaults instantiates a new OutputAlert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputAlertWithDefaults() *OutputAlert {
	this := OutputAlert{}
	return &this
}

// GetUnderscoreId returns the UnderscoreId field value
func (o *OutputAlert) GetUnderscoreId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnderscoreId
}

// GetUnderscoreIdOk returns a tuple with the UnderscoreId field value
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetUnderscoreIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreId, true
}

// SetUnderscoreId sets field value
func (o *OutputAlert) SetUnderscoreId(v string) {
	o.UnderscoreId = v
}

// GetUnderscoreType returns the UnderscoreType field value
func (o *OutputAlert) GetUnderscoreType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnderscoreType
}

// GetUnderscoreTypeOk returns a tuple with the UnderscoreType field value
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetUnderscoreTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreType, true
}

// SetUnderscoreType sets field value
func (o *OutputAlert) SetUnderscoreType(v string) {
	o.UnderscoreType = v
}

// GetUnderscoreCreatedBy returns the UnderscoreCreatedBy field value
func (o *OutputAlert) GetUnderscoreCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnderscoreCreatedBy
}

// GetUnderscoreCreatedByOk returns a tuple with the UnderscoreCreatedBy field value
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetUnderscoreCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreCreatedBy, true
}

// SetUnderscoreCreatedBy sets field value
func (o *OutputAlert) SetUnderscoreCreatedBy(v string) {
	o.UnderscoreCreatedBy = v
}

// GetUnderscoreUpdatedBy returns the UnderscoreUpdatedBy field value if set, zero value otherwise.
func (o *OutputAlert) GetUnderscoreUpdatedBy() string {
	if o == nil || IsNil(o.UnderscoreUpdatedBy) {
		var ret string
		return ret
	}
	return *o.UnderscoreUpdatedBy
}

// GetUnderscoreUpdatedByOk returns a tuple with the UnderscoreUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetUnderscoreUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UnderscoreUpdatedBy) {
		return nil, false
	}
	return o.UnderscoreUpdatedBy, true
}

// HasUnderscoreUpdatedBy returns a boolean if a field has been set.
func (o *OutputAlert) HasUnderscoreUpdatedBy() bool {
	if o != nil && !IsNil(o.UnderscoreUpdatedBy) {
		return true
	}

	return false
}

// SetUnderscoreUpdatedBy gets a reference to the given string and assigns it to the UnderscoreUpdatedBy field.
func (o *OutputAlert) SetUnderscoreUpdatedBy(v string) {
	o.UnderscoreUpdatedBy = &v
}

// GetUnderscoreCreatedAt returns the UnderscoreCreatedAt field value
func (o *OutputAlert) GetUnderscoreCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UnderscoreCreatedAt
}

// GetUnderscoreCreatedAtOk returns a tuple with the UnderscoreCreatedAt field value
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetUnderscoreCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreCreatedAt, true
}

// SetUnderscoreCreatedAt sets field value
func (o *OutputAlert) SetUnderscoreCreatedAt(v int64) {
	o.UnderscoreCreatedAt = v
}

// GetUnderscoreUpdatedAt returns the UnderscoreUpdatedAt field value if set, zero value otherwise.
func (o *OutputAlert) GetUnderscoreUpdatedAt() int64 {
	if o == nil || IsNil(o.UnderscoreUpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UnderscoreUpdatedAt
}

// GetUnderscoreUpdatedAtOk returns a tuple with the UnderscoreUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetUnderscoreUpdatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.UnderscoreUpdatedAt) {
		return nil, false
	}
	return o.UnderscoreUpdatedAt, true
}

// HasUnderscoreUpdatedAt returns a boolean if a field has been set.
func (o *OutputAlert) HasUnderscoreUpdatedAt() bool {
	if o != nil && !IsNil(o.UnderscoreUpdatedAt) {
		return true
	}

	return false
}

// SetUnderscoreUpdatedAt gets a reference to the given int64 and assigns it to the UnderscoreUpdatedAt field.
func (o *OutputAlert) SetUnderscoreUpdatedAt(v int64) {
	o.UnderscoreUpdatedAt = &v
}

// GetType returns the Type field value
func (o *OutputAlert) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OutputAlert) SetType(v string) {
	o.Type = v
}

// GetSource returns the Source field value
func (o *OutputAlert) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *OutputAlert) SetSource(v string) {
	o.Source = v
}

// GetSourceRef returns the SourceRef field value
func (o *OutputAlert) GetSourceRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceRef
}

// GetSourceRefOk returns a tuple with the SourceRef field value
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetSourceRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceRef, true
}

// SetSourceRef sets field value
func (o *OutputAlert) SetSourceRef(v string) {
	o.SourceRef = v
}

// GetExternalLink returns the ExternalLink field value if set, zero value otherwise.
func (o *OutputAlert) GetExternalLink() string {
	if o == nil || IsNil(o.ExternalLink) {
		var ret string
		return ret
	}
	return *o.ExternalLink
}

// GetExternalLinkOk returns a tuple with the ExternalLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetExternalLinkOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalLink) {
		return nil, false
	}
	return o.ExternalLink, true
}

// HasExternalLink returns a boolean if a field has been set.
func (o *OutputAlert) HasExternalLink() bool {
	if o != nil && !IsNil(o.ExternalLink) {
		return true
	}

	return false
}

// SetExternalLink gets a reference to the given string and assigns it to the ExternalLink field.
func (o *OutputAlert) SetExternalLink(v string) {
	o.ExternalLink = &v
}

// GetTitle returns the Title field value
func (o *OutputAlert) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *OutputAlert) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value
func (o *OutputAlert) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *OutputAlert) SetDescription(v string) {
	o.Description = v
}

// GetSeverity returns the Severity field value
func (o *OutputAlert) GetSeverity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetSeverityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *OutputAlert) SetSeverity(v int32) {
	o.Severity = v
}

// GetSeverityLabel returns the SeverityLabel field value
func (o *OutputAlert) GetSeverityLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeverityLabel
}

// GetSeverityLabelOk returns a tuple with the SeverityLabel field value
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetSeverityLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeverityLabel, true
}

// SetSeverityLabel sets field value
func (o *OutputAlert) SetSeverityLabel(v string) {
	o.SeverityLabel = v
}

// GetDate returns the Date field value
func (o *OutputAlert) GetDate() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *OutputAlert) SetDate(v int64) {
	o.Date = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *OutputAlert) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *OutputAlert) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *OutputAlert) SetTags(v []string) {
	o.Tags = v
}

// GetTlp returns the Tlp field value
func (o *OutputAlert) GetTlp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Tlp
}

// GetTlpOk returns a tuple with the Tlp field value
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetTlpOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tlp, true
}

// SetTlp sets field value
func (o *OutputAlert) SetTlp(v int32) {
	o.Tlp = v
}

// GetTlpLabel returns the TlpLabel field value
func (o *OutputAlert) GetTlpLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TlpLabel
}

// GetTlpLabelOk returns a tuple with the TlpLabel field value
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetTlpLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TlpLabel, true
}

// SetTlpLabel sets field value
func (o *OutputAlert) SetTlpLabel(v string) {
	o.TlpLabel = v
}

// GetPap returns the Pap field value
func (o *OutputAlert) GetPap() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pap
}

// GetPapOk returns a tuple with the Pap field value
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetPapOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pap, true
}

// SetPap sets field value
func (o *OutputAlert) SetPap(v int32) {
	o.Pap = v
}

// GetPapLabel returns the PapLabel field value
func (o *OutputAlert) GetPapLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PapLabel
}

// GetPapLabelOk returns a tuple with the PapLabel field value
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetPapLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PapLabel, true
}

// SetPapLabel sets field value
func (o *OutputAlert) SetPapLabel(v string) {
	o.PapLabel = v
}

// GetFollow returns the Follow field value
func (o *OutputAlert) GetFollow() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Follow
}

// GetFollowOk returns a tuple with the Follow field value
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetFollowOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Follow, true
}

// SetFollow sets field value
func (o *OutputAlert) SetFollow(v bool) {
	o.Follow = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *OutputAlert) GetCustomFields() []OutputCustomFieldValue {
	if o == nil || IsNil(o.CustomFields) {
		var ret []OutputCustomFieldValue
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetCustomFieldsOk() ([]OutputCustomFieldValue, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *OutputAlert) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given []OutputCustomFieldValue and assigns it to the CustomFields field.
func (o *OutputAlert) SetCustomFields(v []OutputCustomFieldValue) {
	o.CustomFields = v
}

// GetCaseTemplate returns the CaseTemplate field value if set, zero value otherwise.
func (o *OutputAlert) GetCaseTemplate() string {
	if o == nil || IsNil(o.CaseTemplate) {
		var ret string
		return ret
	}
	return *o.CaseTemplate
}

// GetCaseTemplateOk returns a tuple with the CaseTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetCaseTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.CaseTemplate) {
		return nil, false
	}
	return o.CaseTemplate, true
}

// HasCaseTemplate returns a boolean if a field has been set.
func (o *OutputAlert) HasCaseTemplate() bool {
	if o != nil && !IsNil(o.CaseTemplate) {
		return true
	}

	return false
}

// SetCaseTemplate gets a reference to the given string and assigns it to the CaseTemplate field.
func (o *OutputAlert) SetCaseTemplate(v string) {
	o.CaseTemplate = &v
}

// GetObservableCount returns the ObservableCount field value
func (o *OutputAlert) GetObservableCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ObservableCount
}

// GetObservableCountOk returns a tuple with the ObservableCount field value
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetObservableCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObservableCount, true
}

// SetObservableCount sets field value
func (o *OutputAlert) SetObservableCount(v int64) {
	o.ObservableCount = v
}

// GetCaseId returns the CaseId field value if set, zero value otherwise.
func (o *OutputAlert) GetCaseId() string {
	if o == nil || IsNil(o.CaseId) {
		var ret string
		return ret
	}
	return *o.CaseId
}

// GetCaseIdOk returns a tuple with the CaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetCaseIdOk() (*string, bool) {
	if o == nil || IsNil(o.CaseId) {
		return nil, false
	}
	return o.CaseId, true
}

// HasCaseId returns a boolean if a field has been set.
func (o *OutputAlert) HasCaseId() bool {
	if o != nil && !IsNil(o.CaseId) {
		return true
	}

	return false
}

// SetCaseId gets a reference to the given string and assigns it to the CaseId field.
func (o *OutputAlert) SetCaseId(v string) {
	o.CaseId = &v
}

// GetStatus returns the Status field value
func (o *OutputAlert) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *OutputAlert) SetStatus(v string) {
	o.Status = v
}

// GetStage returns the Stage field value
func (o *OutputAlert) GetStage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stage
}

// GetStageOk returns a tuple with the Stage field value
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetStageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stage, true
}

// SetStage sets field value
func (o *OutputAlert) SetStage(v string) {
	o.Stage = v
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *OutputAlert) GetAssignee() string {
	if o == nil || IsNil(o.Assignee) {
		var ret string
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetAssigneeOk() (*string, bool) {
	if o == nil || IsNil(o.Assignee) {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *OutputAlert) HasAssignee() bool {
	if o != nil && !IsNil(o.Assignee) {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given string and assigns it to the Assignee field.
func (o *OutputAlert) SetAssignee(v string) {
	o.Assignee = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *OutputAlert) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *OutputAlert) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *OutputAlert) SetSummary(v string) {
	o.Summary = &v
}

// GetExtraData returns the ExtraData field value
func (o *OutputAlert) GetExtraData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ExtraData
}

// GetExtraDataOk returns a tuple with the ExtraData field value
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetExtraDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.ExtraData, true
}

// SetExtraData sets field value
func (o *OutputAlert) SetExtraData(v map[string]interface{}) {
	o.ExtraData = v
}

// GetNewDate returns the NewDate field value
func (o *OutputAlert) GetNewDate() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NewDate
}

// GetNewDateOk returns a tuple with the NewDate field value
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetNewDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewDate, true
}

// SetNewDate sets field value
func (o *OutputAlert) SetNewDate(v int64) {
	o.NewDate = v
}

// GetInProgressDate returns the InProgressDate field value if set, zero value otherwise.
func (o *OutputAlert) GetInProgressDate() int64 {
	if o == nil || IsNil(o.InProgressDate) {
		var ret int64
		return ret
	}
	return *o.InProgressDate
}

// GetInProgressDateOk returns a tuple with the InProgressDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetInProgressDateOk() (*int64, bool) {
	if o == nil || IsNil(o.InProgressDate) {
		return nil, false
	}
	return o.InProgressDate, true
}

// HasInProgressDate returns a boolean if a field has been set.
func (o *OutputAlert) HasInProgressDate() bool {
	if o != nil && !IsNil(o.InProgressDate) {
		return true
	}

	return false
}

// SetInProgressDate gets a reference to the given int64 and assigns it to the InProgressDate field.
func (o *OutputAlert) SetInProgressDate(v int64) {
	o.InProgressDate = &v
}

// GetClosedDate returns the ClosedDate field value if set, zero value otherwise.
func (o *OutputAlert) GetClosedDate() int64 {
	if o == nil || IsNil(o.ClosedDate) {
		var ret int64
		return ret
	}
	return *o.ClosedDate
}

// GetClosedDateOk returns a tuple with the ClosedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetClosedDateOk() (*int64, bool) {
	if o == nil || IsNil(o.ClosedDate) {
		return nil, false
	}
	return o.ClosedDate, true
}

// HasClosedDate returns a boolean if a field has been set.
func (o *OutputAlert) HasClosedDate() bool {
	if o != nil && !IsNil(o.ClosedDate) {
		return true
	}

	return false
}

// SetClosedDate gets a reference to the given int64 and assigns it to the ClosedDate field.
func (o *OutputAlert) SetClosedDate(v int64) {
	o.ClosedDate = &v
}

// GetImportedDate returns the ImportedDate field value if set, zero value otherwise.
func (o *OutputAlert) GetImportedDate() int64 {
	if o == nil || IsNil(o.ImportedDate) {
		var ret int64
		return ret
	}
	return *o.ImportedDate
}

// GetImportedDateOk returns a tuple with the ImportedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetImportedDateOk() (*int64, bool) {
	if o == nil || IsNil(o.ImportedDate) {
		return nil, false
	}
	return o.ImportedDate, true
}

// HasImportedDate returns a boolean if a field has been set.
func (o *OutputAlert) HasImportedDate() bool {
	if o != nil && !IsNil(o.ImportedDate) {
		return true
	}

	return false
}

// SetImportedDate gets a reference to the given int64 and assigns it to the ImportedDate field.
func (o *OutputAlert) SetImportedDate(v int64) {
	o.ImportedDate = &v
}

// GetTimeToDetect returns the TimeToDetect field value
func (o *OutputAlert) GetTimeToDetect() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TimeToDetect
}

// GetTimeToDetectOk returns a tuple with the TimeToDetect field value
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetTimeToDetectOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeToDetect, true
}

// SetTimeToDetect sets field value
func (o *OutputAlert) SetTimeToDetect(v int64) {
	o.TimeToDetect = v
}

// GetTimeToTriage returns the TimeToTriage field value if set, zero value otherwise.
func (o *OutputAlert) GetTimeToTriage() int64 {
	if o == nil || IsNil(o.TimeToTriage) {
		var ret int64
		return ret
	}
	return *o.TimeToTriage
}

// GetTimeToTriageOk returns a tuple with the TimeToTriage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetTimeToTriageOk() (*int64, bool) {
	if o == nil || IsNil(o.TimeToTriage) {
		return nil, false
	}
	return o.TimeToTriage, true
}

// HasTimeToTriage returns a boolean if a field has been set.
func (o *OutputAlert) HasTimeToTriage() bool {
	if o != nil && !IsNil(o.TimeToTriage) {
		return true
	}

	return false
}

// SetTimeToTriage gets a reference to the given int64 and assigns it to the TimeToTriage field.
func (o *OutputAlert) SetTimeToTriage(v int64) {
	o.TimeToTriage = &v
}

// GetTimeToQualify returns the TimeToQualify field value if set, zero value otherwise.
func (o *OutputAlert) GetTimeToQualify() int64 {
	if o == nil || IsNil(o.TimeToQualify) {
		var ret int64
		return ret
	}
	return *o.TimeToQualify
}

// GetTimeToQualifyOk returns a tuple with the TimeToQualify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetTimeToQualifyOk() (*int64, bool) {
	if o == nil || IsNil(o.TimeToQualify) {
		return nil, false
	}
	return o.TimeToQualify, true
}

// HasTimeToQualify returns a boolean if a field has been set.
func (o *OutputAlert) HasTimeToQualify() bool {
	if o != nil && !IsNil(o.TimeToQualify) {
		return true
	}

	return false
}

// SetTimeToQualify gets a reference to the given int64 and assigns it to the TimeToQualify field.
func (o *OutputAlert) SetTimeToQualify(v int64) {
	o.TimeToQualify = &v
}

// GetTimeToAcknowledge returns the TimeToAcknowledge field value if set, zero value otherwise.
func (o *OutputAlert) GetTimeToAcknowledge() int64 {
	if o == nil || IsNil(o.TimeToAcknowledge) {
		var ret int64
		return ret
	}
	return *o.TimeToAcknowledge
}

// GetTimeToAcknowledgeOk returns a tuple with the TimeToAcknowledge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputAlert) GetTimeToAcknowledgeOk() (*int64, bool) {
	if o == nil || IsNil(o.TimeToAcknowledge) {
		return nil, false
	}
	return o.TimeToAcknowledge, true
}

// HasTimeToAcknowledge returns a boolean if a field has been set.
func (o *OutputAlert) HasTimeToAcknowledge() bool {
	if o != nil && !IsNil(o.TimeToAcknowledge) {
		return true
	}

	return false
}

// SetTimeToAcknowledge gets a reference to the given int64 and assigns it to the TimeToAcknowledge field.
func (o *OutputAlert) SetTimeToAcknowledge(v int64) {
	o.TimeToAcknowledge = &v
}

func (o OutputAlert) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutputAlert) ToMap() (map[string]interface{}, error) {
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
	toSerialize["type"] = o.Type
	toSerialize["source"] = o.Source
	toSerialize["sourceRef"] = o.SourceRef
	if !IsNil(o.ExternalLink) {
		toSerialize["externalLink"] = o.ExternalLink
	}
	toSerialize["title"] = o.Title
	toSerialize["description"] = o.Description
	toSerialize["severity"] = o.Severity
	toSerialize["severityLabel"] = o.SeverityLabel
	toSerialize["date"] = o.Date
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["tlp"] = o.Tlp
	toSerialize["tlpLabel"] = o.TlpLabel
	toSerialize["pap"] = o.Pap
	toSerialize["papLabel"] = o.PapLabel
	toSerialize["follow"] = o.Follow
	if !IsNil(o.CustomFields) {
		toSerialize["customFields"] = o.CustomFields
	}
	if !IsNil(o.CaseTemplate) {
		toSerialize["caseTemplate"] = o.CaseTemplate
	}
	toSerialize["observableCount"] = o.ObservableCount
	if !IsNil(o.CaseId) {
		toSerialize["caseId"] = o.CaseId
	}
	toSerialize["status"] = o.Status
	toSerialize["stage"] = o.Stage
	if !IsNil(o.Assignee) {
		toSerialize["assignee"] = o.Assignee
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	toSerialize["extraData"] = o.ExtraData
	toSerialize["newDate"] = o.NewDate
	if !IsNil(o.InProgressDate) {
		toSerialize["inProgressDate"] = o.InProgressDate
	}
	if !IsNil(o.ClosedDate) {
		toSerialize["closedDate"] = o.ClosedDate
	}
	if !IsNil(o.ImportedDate) {
		toSerialize["importedDate"] = o.ImportedDate
	}
	toSerialize["timeToDetect"] = o.TimeToDetect
	if !IsNil(o.TimeToTriage) {
		toSerialize["timeToTriage"] = o.TimeToTriage
	}
	if !IsNil(o.TimeToQualify) {
		toSerialize["timeToQualify"] = o.TimeToQualify
	}
	if !IsNil(o.TimeToAcknowledge) {
		toSerialize["timeToAcknowledge"] = o.TimeToAcknowledge
	}
	return toSerialize, nil
}

func (o *OutputAlert) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"_type",
		"_createdBy",
		"_createdAt",
		"type",
		"source",
		"sourceRef",
		"title",
		"description",
		"severity",
		"severityLabel",
		"date",
		"tlp",
		"tlpLabel",
		"pap",
		"papLabel",
		"follow",
		"observableCount",
		"status",
		"stage",
		"extraData",
		"newDate",
		"timeToDetect",
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

	varOutputAlert := _OutputAlert{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// Postprocessed by scripts/tolerant-response-decoder: unknown fields are ignored,
	// since TheHive adds fields to its responses in minor versions.
	err = decoder.Decode(&varOutputAlert)

	if err != nil {
		return err
	}

	*o = OutputAlert(varOutputAlert)

	return err
}

type NullableOutputAlert struct {
	value *OutputAlert
	isSet bool
}

func (v NullableOutputAlert) Get() *OutputAlert {
	return v.value
}

func (v *NullableOutputAlert) Set(val *OutputAlert) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputAlert) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputAlert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputAlert(val *OutputAlert) *NullableOutputAlert {
	return &NullableOutputAlert{value: val, isSet: true}
}

func (v NullableOutputAlert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputAlert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
