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

// checks if the OutputCase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutputCase{}

// OutputCase struct for OutputCase
type OutputCase struct {
	// Internal unique identifier of the case.
	UnderscoreId string `json:"_id"`
	// Entity type, always `Case`.
	UnderscoreType string `json:"_type"`
	// Login of the user who created the case.
	UnderscoreCreatedBy string `json:"_createdBy"`
	// Login of the user who last updated the case.
	UnderscoreUpdatedBy *string `json:"_updatedBy,omitempty"`
	// Date and time when the case was created. Returned as a Unix timestamp in milliseconds.
	UnderscoreCreatedAt int64 `json:"_createdAt"`
	// Date and time when the case was last updated. Returned as a Unix timestamp in milliseconds.
	UnderscoreUpdatedAt *int64 `json:"_updatedAt,omitempty"`
	// An incremental number to reference the case. This field can be used in the API in place of `_id`.
	Number int32 `json:"number"`
	// Title of the case.
	Title string `json:"title"`
	// Details of the case.
	Description string `json:"description"`
	// Numeric severity level: 1 (Low), 2 (Medium), 3 (High), 4 (Critical).
	Severity int32 `json:"severity"`
	// Human-readable severity label.
	SeverityLabel string `json:"severityLabel"`
	// Start date of the incident. Returned as a Unix timestamp in milliseconds.
	StartDate int64 `json:"startDate"`
	// End date of the incident. Returned as a Unix timestamp in milliseconds.
	EndDate *int64 `json:"endDate,omitempty"`
	// Tags assigned to the case.
	Tags []string `json:"tags,omitempty"`
	// Whether the case is flagged for quick filtering.
	Flag bool `json:"flag"`
	// Numeric Traffic Light Protocol (TLP) level: 0 (Clear), 1 (Green), 2 (Amber), 3 (Amber+Strict), 4 (Red).
	Tlp int32 `json:"tlp"`
	// Human-readable TLP label.
	TlpLabel string `json:"tlpLabel"`
	// Numeric Permissible Actions Protocol (PAP) level: 0 (Clear), 1 (Green), 2 (Amber), 3 (Red).
	Pap int32 `json:"pap"`
	// Human-readable PAP label.
	PapLabel string `json:"papLabel"`
	// Current status of the case.
	Status string `json:"status"`
	// Stage derived from the case status. One of *New*, *InProgress*, or *Closed*.
	Stage string `json:"stage"`
	// Brief summary of the investigation findings.
	Summary *string `json:"summary,omitempty"`
	// Verdict on whether the incident had real consequences for the organization. Only set when the case is closed with a status that requires an impact. One of `WithImpact`, `NoImpact`, or `NotApplicable`.
	ImpactStatus *string `json:"impactStatus,omitempty"`
	// Login of the user assigned to the case.
	Assignee *string `json:"assignee,omitempty"`
	// Access control settings of the case.
	Access Access `json:"access"`
	// Custom field values set on the case.
	CustomFields []OutputCustomFieldValue `json:"customFields,omitempty"`
	// Permissions the current user has on this case, based on their profile and the case sharing configuration.
	UserPermissions []string               `json:"userPermissions,omitempty"`
	ExtraData       map[string]interface{} `json:"extraData"`
	// Date when the case first entered a status linked to the *New* stage. Returned as a Unix timestamp in milliseconds.
	NewDate int64 `json:"newDate"`
	// Date when the case first entered a status linked to the *In Progress* stage. Returned as a Unix timestamp in milliseconds.
	InProgressDate *int64 `json:"inProgressDate,omitempty"`
	// Date when the case first entered a status linked to the *Closed* stage. Returned as a Unix timestamp in milliseconds.
	ClosedDate *int64 `json:"closedDate,omitempty"`
	// The date the incident occurred according to the earliest linked alert. Returned as a Unix timestamp in milliseconds.
	AlertDate *int64 `json:"alertDate,omitempty"`
	// Date when the earliest linked alert first entered the *New* stage. Returned as a Unix timestamp in milliseconds.
	AlertNewDate *int64 `json:"alertNewDate,omitempty"`
	// Date when the earliest linked alert first entered the *In Progress* stage. Returned as a Unix timestamp in milliseconds.
	AlertInProgressDate *int64 `json:"alertInProgressDate,omitempty"`
	// Date when the earliest linked alert was imported into a case. Returned as a Unix timestamp in milliseconds.
	AlertImportedDate *int64 `json:"alertImportedDate,omitempty"`
	// [Time to Detect (TTD)](https://docs.strangebee.com/thehive/user-guides/key-performance-indicators/key-performance-indicators/#time-to-detect-ttd), in milliseconds.
	TimeToDetect int64 `json:"timeToDetect"`
	// [Time to Triage (TTT)](https://docs.strangebee.com/thehive/user-guides/key-performance-indicators/key-performance-indicators/#time-to-triage-ttt), in milliseconds.
	TimeToTriage *int64 `json:"timeToTriage,omitempty"`
	// [Time to Qualify (TTQ)](https://docs.strangebee.com/thehive/user-guides/key-performance-indicators/key-performance-indicators/#time-to-qualify-ttq), in milliseconds.
	TimeToQualify *int64 `json:"timeToQualify,omitempty"`
	// [Time to Acknowledge (TTA)](https://docs.strangebee.com/thehive/user-guides/key-performance-indicators/key-performance-indicators/#time-to-acknowledge-tta), in milliseconds.
	TimeToAcknowledge *int64 `json:"timeToAcknowledge,omitempty"`
	// [Time to Resolve (TTR)](https://docs.strangebee.com/thehive/user-guides/key-performance-indicators/key-performance-indicators/#time-to-resolve-ttr), in milliseconds.
	TimeToResolve *int64 `json:"timeToResolve,omitempty"`
	// [Handling duration](https://docs.strangebee.com/thehive/user-guides/key-performance-indicators/key-performance-indicators/#time-to-resolve-ttr), in milliseconds.
	HandlingDuration *int64 `json:"handlingDuration,omitempty"`
}

type _OutputCase OutputCase

// NewOutputCase instantiates a new OutputCase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputCase(underscoreId string, underscoreType string, underscoreCreatedBy string, underscoreCreatedAt int64, number int32, title string, description string, severity int32, severityLabel string, startDate int64, flag bool, tlp int32, tlpLabel string, pap int32, papLabel string, status string, stage string, access Access, extraData map[string]interface{}, newDate int64, timeToDetect int64) *OutputCase {
	this := OutputCase{}
	this.UnderscoreId = underscoreId
	this.UnderscoreType = underscoreType
	this.UnderscoreCreatedBy = underscoreCreatedBy
	this.UnderscoreCreatedAt = underscoreCreatedAt
	this.Number = number
	this.Title = title
	this.Description = description
	this.Severity = severity
	this.SeverityLabel = severityLabel
	this.StartDate = startDate
	this.Flag = flag
	this.Tlp = tlp
	this.TlpLabel = tlpLabel
	this.Pap = pap
	this.PapLabel = papLabel
	this.Status = status
	this.Stage = stage
	this.Access = access
	this.ExtraData = extraData
	this.NewDate = newDate
	this.TimeToDetect = timeToDetect
	return &this
}

// NewOutputCaseWithDefaults instantiates a new OutputCase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputCaseWithDefaults() *OutputCase {
	this := OutputCase{}
	return &this
}

// GetUnderscoreId returns the UnderscoreId field value
func (o *OutputCase) GetUnderscoreId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnderscoreId
}

// GetUnderscoreIdOk returns a tuple with the UnderscoreId field value
// and a boolean to check if the value has been set.
func (o *OutputCase) GetUnderscoreIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreId, true
}

// SetUnderscoreId sets field value
func (o *OutputCase) SetUnderscoreId(v string) {
	o.UnderscoreId = v
}

// GetUnderscoreType returns the UnderscoreType field value
func (o *OutputCase) GetUnderscoreType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnderscoreType
}

// GetUnderscoreTypeOk returns a tuple with the UnderscoreType field value
// and a boolean to check if the value has been set.
func (o *OutputCase) GetUnderscoreTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreType, true
}

// SetUnderscoreType sets field value
func (o *OutputCase) SetUnderscoreType(v string) {
	o.UnderscoreType = v
}

// GetUnderscoreCreatedBy returns the UnderscoreCreatedBy field value
func (o *OutputCase) GetUnderscoreCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnderscoreCreatedBy
}

// GetUnderscoreCreatedByOk returns a tuple with the UnderscoreCreatedBy field value
// and a boolean to check if the value has been set.
func (o *OutputCase) GetUnderscoreCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreCreatedBy, true
}

// SetUnderscoreCreatedBy sets field value
func (o *OutputCase) SetUnderscoreCreatedBy(v string) {
	o.UnderscoreCreatedBy = v
}

// GetUnderscoreUpdatedBy returns the UnderscoreUpdatedBy field value if set, zero value otherwise.
func (o *OutputCase) GetUnderscoreUpdatedBy() string {
	if o == nil || IsNil(o.UnderscoreUpdatedBy) {
		var ret string
		return ret
	}
	return *o.UnderscoreUpdatedBy
}

// GetUnderscoreUpdatedByOk returns a tuple with the UnderscoreUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCase) GetUnderscoreUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UnderscoreUpdatedBy) {
		return nil, false
	}
	return o.UnderscoreUpdatedBy, true
}

// HasUnderscoreUpdatedBy returns a boolean if a field has been set.
func (o *OutputCase) HasUnderscoreUpdatedBy() bool {
	if o != nil && !IsNil(o.UnderscoreUpdatedBy) {
		return true
	}

	return false
}

// SetUnderscoreUpdatedBy gets a reference to the given string and assigns it to the UnderscoreUpdatedBy field.
func (o *OutputCase) SetUnderscoreUpdatedBy(v string) {
	o.UnderscoreUpdatedBy = &v
}

// GetUnderscoreCreatedAt returns the UnderscoreCreatedAt field value
func (o *OutputCase) GetUnderscoreCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UnderscoreCreatedAt
}

// GetUnderscoreCreatedAtOk returns a tuple with the UnderscoreCreatedAt field value
// and a boolean to check if the value has been set.
func (o *OutputCase) GetUnderscoreCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreCreatedAt, true
}

// SetUnderscoreCreatedAt sets field value
func (o *OutputCase) SetUnderscoreCreatedAt(v int64) {
	o.UnderscoreCreatedAt = v
}

// GetUnderscoreUpdatedAt returns the UnderscoreUpdatedAt field value if set, zero value otherwise.
func (o *OutputCase) GetUnderscoreUpdatedAt() int64 {
	if o == nil || IsNil(o.UnderscoreUpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UnderscoreUpdatedAt
}

// GetUnderscoreUpdatedAtOk returns a tuple with the UnderscoreUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCase) GetUnderscoreUpdatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.UnderscoreUpdatedAt) {
		return nil, false
	}
	return o.UnderscoreUpdatedAt, true
}

// HasUnderscoreUpdatedAt returns a boolean if a field has been set.
func (o *OutputCase) HasUnderscoreUpdatedAt() bool {
	if o != nil && !IsNil(o.UnderscoreUpdatedAt) {
		return true
	}

	return false
}

// SetUnderscoreUpdatedAt gets a reference to the given int64 and assigns it to the UnderscoreUpdatedAt field.
func (o *OutputCase) SetUnderscoreUpdatedAt(v int64) {
	o.UnderscoreUpdatedAt = &v
}

// GetNumber returns the Number field value
func (o *OutputCase) GetNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *OutputCase) GetNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *OutputCase) SetNumber(v int32) {
	o.Number = v
}

// GetTitle returns the Title field value
func (o *OutputCase) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *OutputCase) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *OutputCase) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value
func (o *OutputCase) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *OutputCase) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *OutputCase) SetDescription(v string) {
	o.Description = v
}

// GetSeverity returns the Severity field value
func (o *OutputCase) GetSeverity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *OutputCase) GetSeverityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *OutputCase) SetSeverity(v int32) {
	o.Severity = v
}

// GetSeverityLabel returns the SeverityLabel field value
func (o *OutputCase) GetSeverityLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeverityLabel
}

// GetSeverityLabelOk returns a tuple with the SeverityLabel field value
// and a boolean to check if the value has been set.
func (o *OutputCase) GetSeverityLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeverityLabel, true
}

// SetSeverityLabel sets field value
func (o *OutputCase) SetSeverityLabel(v string) {
	o.SeverityLabel = v
}

// GetStartDate returns the StartDate field value
func (o *OutputCase) GetStartDate() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *OutputCase) GetStartDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *OutputCase) SetStartDate(v int64) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *OutputCase) GetEndDate() int64 {
	if o == nil || IsNil(o.EndDate) {
		var ret int64
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCase) GetEndDateOk() (*int64, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *OutputCase) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given int64 and assigns it to the EndDate field.
func (o *OutputCase) SetEndDate(v int64) {
	o.EndDate = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *OutputCase) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCase) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *OutputCase) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *OutputCase) SetTags(v []string) {
	o.Tags = v
}

// GetFlag returns the Flag field value
func (o *OutputCase) GetFlag() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Flag
}

// GetFlagOk returns a tuple with the Flag field value
// and a boolean to check if the value has been set.
func (o *OutputCase) GetFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Flag, true
}

// SetFlag sets field value
func (o *OutputCase) SetFlag(v bool) {
	o.Flag = v
}

// GetTlp returns the Tlp field value
func (o *OutputCase) GetTlp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Tlp
}

// GetTlpOk returns a tuple with the Tlp field value
// and a boolean to check if the value has been set.
func (o *OutputCase) GetTlpOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tlp, true
}

// SetTlp sets field value
func (o *OutputCase) SetTlp(v int32) {
	o.Tlp = v
}

// GetTlpLabel returns the TlpLabel field value
func (o *OutputCase) GetTlpLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TlpLabel
}

// GetTlpLabelOk returns a tuple with the TlpLabel field value
// and a boolean to check if the value has been set.
func (o *OutputCase) GetTlpLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TlpLabel, true
}

// SetTlpLabel sets field value
func (o *OutputCase) SetTlpLabel(v string) {
	o.TlpLabel = v
}

// GetPap returns the Pap field value
func (o *OutputCase) GetPap() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pap
}

// GetPapOk returns a tuple with the Pap field value
// and a boolean to check if the value has been set.
func (o *OutputCase) GetPapOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pap, true
}

// SetPap sets field value
func (o *OutputCase) SetPap(v int32) {
	o.Pap = v
}

// GetPapLabel returns the PapLabel field value
func (o *OutputCase) GetPapLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PapLabel
}

// GetPapLabelOk returns a tuple with the PapLabel field value
// and a boolean to check if the value has been set.
func (o *OutputCase) GetPapLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PapLabel, true
}

// SetPapLabel sets field value
func (o *OutputCase) SetPapLabel(v string) {
	o.PapLabel = v
}

// GetStatus returns the Status field value
func (o *OutputCase) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OutputCase) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *OutputCase) SetStatus(v string) {
	o.Status = v
}

// GetStage returns the Stage field value
func (o *OutputCase) GetStage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stage
}

// GetStageOk returns a tuple with the Stage field value
// and a boolean to check if the value has been set.
func (o *OutputCase) GetStageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stage, true
}

// SetStage sets field value
func (o *OutputCase) SetStage(v string) {
	o.Stage = v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *OutputCase) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCase) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *OutputCase) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *OutputCase) SetSummary(v string) {
	o.Summary = &v
}

// GetImpactStatus returns the ImpactStatus field value if set, zero value otherwise.
func (o *OutputCase) GetImpactStatus() string {
	if o == nil || IsNil(o.ImpactStatus) {
		var ret string
		return ret
	}
	return *o.ImpactStatus
}

// GetImpactStatusOk returns a tuple with the ImpactStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCase) GetImpactStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ImpactStatus) {
		return nil, false
	}
	return o.ImpactStatus, true
}

// HasImpactStatus returns a boolean if a field has been set.
func (o *OutputCase) HasImpactStatus() bool {
	if o != nil && !IsNil(o.ImpactStatus) {
		return true
	}

	return false
}

// SetImpactStatus gets a reference to the given string and assigns it to the ImpactStatus field.
func (o *OutputCase) SetImpactStatus(v string) {
	o.ImpactStatus = &v
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *OutputCase) GetAssignee() string {
	if o == nil || IsNil(o.Assignee) {
		var ret string
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCase) GetAssigneeOk() (*string, bool) {
	if o == nil || IsNil(o.Assignee) {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *OutputCase) HasAssignee() bool {
	if o != nil && !IsNil(o.Assignee) {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given string and assigns it to the Assignee field.
func (o *OutputCase) SetAssignee(v string) {
	o.Assignee = &v
}

// GetAccess returns the Access field value
func (o *OutputCase) GetAccess() Access {
	if o == nil {
		var ret Access
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *OutputCase) GetAccessOk() (*Access, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *OutputCase) SetAccess(v Access) {
	o.Access = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *OutputCase) GetCustomFields() []OutputCustomFieldValue {
	if o == nil || IsNil(o.CustomFields) {
		var ret []OutputCustomFieldValue
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCase) GetCustomFieldsOk() ([]OutputCustomFieldValue, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *OutputCase) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given []OutputCustomFieldValue and assigns it to the CustomFields field.
func (o *OutputCase) SetCustomFields(v []OutputCustomFieldValue) {
	o.CustomFields = v
}

// GetUserPermissions returns the UserPermissions field value if set, zero value otherwise.
func (o *OutputCase) GetUserPermissions() []string {
	if o == nil || IsNil(o.UserPermissions) {
		var ret []string
		return ret
	}
	return o.UserPermissions
}

// GetUserPermissionsOk returns a tuple with the UserPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCase) GetUserPermissionsOk() ([]string, bool) {
	if o == nil || IsNil(o.UserPermissions) {
		return nil, false
	}
	return o.UserPermissions, true
}

// HasUserPermissions returns a boolean if a field has been set.
func (o *OutputCase) HasUserPermissions() bool {
	if o != nil && !IsNil(o.UserPermissions) {
		return true
	}

	return false
}

// SetUserPermissions gets a reference to the given []string and assigns it to the UserPermissions field.
func (o *OutputCase) SetUserPermissions(v []string) {
	o.UserPermissions = v
}

// GetExtraData returns the ExtraData field value
func (o *OutputCase) GetExtraData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ExtraData
}

// GetExtraDataOk returns a tuple with the ExtraData field value
// and a boolean to check if the value has been set.
func (o *OutputCase) GetExtraDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.ExtraData, true
}

// SetExtraData sets field value
func (o *OutputCase) SetExtraData(v map[string]interface{}) {
	o.ExtraData = v
}

// GetNewDate returns the NewDate field value
func (o *OutputCase) GetNewDate() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NewDate
}

// GetNewDateOk returns a tuple with the NewDate field value
// and a boolean to check if the value has been set.
func (o *OutputCase) GetNewDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewDate, true
}

// SetNewDate sets field value
func (o *OutputCase) SetNewDate(v int64) {
	o.NewDate = v
}

// GetInProgressDate returns the InProgressDate field value if set, zero value otherwise.
func (o *OutputCase) GetInProgressDate() int64 {
	if o == nil || IsNil(o.InProgressDate) {
		var ret int64
		return ret
	}
	return *o.InProgressDate
}

// GetInProgressDateOk returns a tuple with the InProgressDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCase) GetInProgressDateOk() (*int64, bool) {
	if o == nil || IsNil(o.InProgressDate) {
		return nil, false
	}
	return o.InProgressDate, true
}

// HasInProgressDate returns a boolean if a field has been set.
func (o *OutputCase) HasInProgressDate() bool {
	if o != nil && !IsNil(o.InProgressDate) {
		return true
	}

	return false
}

// SetInProgressDate gets a reference to the given int64 and assigns it to the InProgressDate field.
func (o *OutputCase) SetInProgressDate(v int64) {
	o.InProgressDate = &v
}

// GetClosedDate returns the ClosedDate field value if set, zero value otherwise.
func (o *OutputCase) GetClosedDate() int64 {
	if o == nil || IsNil(o.ClosedDate) {
		var ret int64
		return ret
	}
	return *o.ClosedDate
}

// GetClosedDateOk returns a tuple with the ClosedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCase) GetClosedDateOk() (*int64, bool) {
	if o == nil || IsNil(o.ClosedDate) {
		return nil, false
	}
	return o.ClosedDate, true
}

// HasClosedDate returns a boolean if a field has been set.
func (o *OutputCase) HasClosedDate() bool {
	if o != nil && !IsNil(o.ClosedDate) {
		return true
	}

	return false
}

// SetClosedDate gets a reference to the given int64 and assigns it to the ClosedDate field.
func (o *OutputCase) SetClosedDate(v int64) {
	o.ClosedDate = &v
}

// GetAlertDate returns the AlertDate field value if set, zero value otherwise.
func (o *OutputCase) GetAlertDate() int64 {
	if o == nil || IsNil(o.AlertDate) {
		var ret int64
		return ret
	}
	return *o.AlertDate
}

// GetAlertDateOk returns a tuple with the AlertDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCase) GetAlertDateOk() (*int64, bool) {
	if o == nil || IsNil(o.AlertDate) {
		return nil, false
	}
	return o.AlertDate, true
}

// HasAlertDate returns a boolean if a field has been set.
func (o *OutputCase) HasAlertDate() bool {
	if o != nil && !IsNil(o.AlertDate) {
		return true
	}

	return false
}

// SetAlertDate gets a reference to the given int64 and assigns it to the AlertDate field.
func (o *OutputCase) SetAlertDate(v int64) {
	o.AlertDate = &v
}

// GetAlertNewDate returns the AlertNewDate field value if set, zero value otherwise.
func (o *OutputCase) GetAlertNewDate() int64 {
	if o == nil || IsNil(o.AlertNewDate) {
		var ret int64
		return ret
	}
	return *o.AlertNewDate
}

// GetAlertNewDateOk returns a tuple with the AlertNewDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCase) GetAlertNewDateOk() (*int64, bool) {
	if o == nil || IsNil(o.AlertNewDate) {
		return nil, false
	}
	return o.AlertNewDate, true
}

// HasAlertNewDate returns a boolean if a field has been set.
func (o *OutputCase) HasAlertNewDate() bool {
	if o != nil && !IsNil(o.AlertNewDate) {
		return true
	}

	return false
}

// SetAlertNewDate gets a reference to the given int64 and assigns it to the AlertNewDate field.
func (o *OutputCase) SetAlertNewDate(v int64) {
	o.AlertNewDate = &v
}

// GetAlertInProgressDate returns the AlertInProgressDate field value if set, zero value otherwise.
func (o *OutputCase) GetAlertInProgressDate() int64 {
	if o == nil || IsNil(o.AlertInProgressDate) {
		var ret int64
		return ret
	}
	return *o.AlertInProgressDate
}

// GetAlertInProgressDateOk returns a tuple with the AlertInProgressDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCase) GetAlertInProgressDateOk() (*int64, bool) {
	if o == nil || IsNil(o.AlertInProgressDate) {
		return nil, false
	}
	return o.AlertInProgressDate, true
}

// HasAlertInProgressDate returns a boolean if a field has been set.
func (o *OutputCase) HasAlertInProgressDate() bool {
	if o != nil && !IsNil(o.AlertInProgressDate) {
		return true
	}

	return false
}

// SetAlertInProgressDate gets a reference to the given int64 and assigns it to the AlertInProgressDate field.
func (o *OutputCase) SetAlertInProgressDate(v int64) {
	o.AlertInProgressDate = &v
}

// GetAlertImportedDate returns the AlertImportedDate field value if set, zero value otherwise.
func (o *OutputCase) GetAlertImportedDate() int64 {
	if o == nil || IsNil(o.AlertImportedDate) {
		var ret int64
		return ret
	}
	return *o.AlertImportedDate
}

// GetAlertImportedDateOk returns a tuple with the AlertImportedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCase) GetAlertImportedDateOk() (*int64, bool) {
	if o == nil || IsNil(o.AlertImportedDate) {
		return nil, false
	}
	return o.AlertImportedDate, true
}

// HasAlertImportedDate returns a boolean if a field has been set.
func (o *OutputCase) HasAlertImportedDate() bool {
	if o != nil && !IsNil(o.AlertImportedDate) {
		return true
	}

	return false
}

// SetAlertImportedDate gets a reference to the given int64 and assigns it to the AlertImportedDate field.
func (o *OutputCase) SetAlertImportedDate(v int64) {
	o.AlertImportedDate = &v
}

// GetTimeToDetect returns the TimeToDetect field value
func (o *OutputCase) GetTimeToDetect() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TimeToDetect
}

// GetTimeToDetectOk returns a tuple with the TimeToDetect field value
// and a boolean to check if the value has been set.
func (o *OutputCase) GetTimeToDetectOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeToDetect, true
}

// SetTimeToDetect sets field value
func (o *OutputCase) SetTimeToDetect(v int64) {
	o.TimeToDetect = v
}

// GetTimeToTriage returns the TimeToTriage field value if set, zero value otherwise.
func (o *OutputCase) GetTimeToTriage() int64 {
	if o == nil || IsNil(o.TimeToTriage) {
		var ret int64
		return ret
	}
	return *o.TimeToTriage
}

// GetTimeToTriageOk returns a tuple with the TimeToTriage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCase) GetTimeToTriageOk() (*int64, bool) {
	if o == nil || IsNil(o.TimeToTriage) {
		return nil, false
	}
	return o.TimeToTriage, true
}

// HasTimeToTriage returns a boolean if a field has been set.
func (o *OutputCase) HasTimeToTriage() bool {
	if o != nil && !IsNil(o.TimeToTriage) {
		return true
	}

	return false
}

// SetTimeToTriage gets a reference to the given int64 and assigns it to the TimeToTriage field.
func (o *OutputCase) SetTimeToTriage(v int64) {
	o.TimeToTriage = &v
}

// GetTimeToQualify returns the TimeToQualify field value if set, zero value otherwise.
func (o *OutputCase) GetTimeToQualify() int64 {
	if o == nil || IsNil(o.TimeToQualify) {
		var ret int64
		return ret
	}
	return *o.TimeToQualify
}

// GetTimeToQualifyOk returns a tuple with the TimeToQualify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCase) GetTimeToQualifyOk() (*int64, bool) {
	if o == nil || IsNil(o.TimeToQualify) {
		return nil, false
	}
	return o.TimeToQualify, true
}

// HasTimeToQualify returns a boolean if a field has been set.
func (o *OutputCase) HasTimeToQualify() bool {
	if o != nil && !IsNil(o.TimeToQualify) {
		return true
	}

	return false
}

// SetTimeToQualify gets a reference to the given int64 and assigns it to the TimeToQualify field.
func (o *OutputCase) SetTimeToQualify(v int64) {
	o.TimeToQualify = &v
}

// GetTimeToAcknowledge returns the TimeToAcknowledge field value if set, zero value otherwise.
func (o *OutputCase) GetTimeToAcknowledge() int64 {
	if o == nil || IsNil(o.TimeToAcknowledge) {
		var ret int64
		return ret
	}
	return *o.TimeToAcknowledge
}

// GetTimeToAcknowledgeOk returns a tuple with the TimeToAcknowledge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCase) GetTimeToAcknowledgeOk() (*int64, bool) {
	if o == nil || IsNil(o.TimeToAcknowledge) {
		return nil, false
	}
	return o.TimeToAcknowledge, true
}

// HasTimeToAcknowledge returns a boolean if a field has been set.
func (o *OutputCase) HasTimeToAcknowledge() bool {
	if o != nil && !IsNil(o.TimeToAcknowledge) {
		return true
	}

	return false
}

// SetTimeToAcknowledge gets a reference to the given int64 and assigns it to the TimeToAcknowledge field.
func (o *OutputCase) SetTimeToAcknowledge(v int64) {
	o.TimeToAcknowledge = &v
}

// GetTimeToResolve returns the TimeToResolve field value if set, zero value otherwise.
func (o *OutputCase) GetTimeToResolve() int64 {
	if o == nil || IsNil(o.TimeToResolve) {
		var ret int64
		return ret
	}
	return *o.TimeToResolve
}

// GetTimeToResolveOk returns a tuple with the TimeToResolve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCase) GetTimeToResolveOk() (*int64, bool) {
	if o == nil || IsNil(o.TimeToResolve) {
		return nil, false
	}
	return o.TimeToResolve, true
}

// HasTimeToResolve returns a boolean if a field has been set.
func (o *OutputCase) HasTimeToResolve() bool {
	if o != nil && !IsNil(o.TimeToResolve) {
		return true
	}

	return false
}

// SetTimeToResolve gets a reference to the given int64 and assigns it to the TimeToResolve field.
func (o *OutputCase) SetTimeToResolve(v int64) {
	o.TimeToResolve = &v
}

// GetHandlingDuration returns the HandlingDuration field value if set, zero value otherwise.
func (o *OutputCase) GetHandlingDuration() int64 {
	if o == nil || IsNil(o.HandlingDuration) {
		var ret int64
		return ret
	}
	return *o.HandlingDuration
}

// GetHandlingDurationOk returns a tuple with the HandlingDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCase) GetHandlingDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.HandlingDuration) {
		return nil, false
	}
	return o.HandlingDuration, true
}

// HasHandlingDuration returns a boolean if a field has been set.
func (o *OutputCase) HasHandlingDuration() bool {
	if o != nil && !IsNil(o.HandlingDuration) {
		return true
	}

	return false
}

// SetHandlingDuration gets a reference to the given int64 and assigns it to the HandlingDuration field.
func (o *OutputCase) SetHandlingDuration(v int64) {
	o.HandlingDuration = &v
}

func (o OutputCase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutputCase) ToMap() (map[string]interface{}, error) {
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
	toSerialize["number"] = o.Number
	toSerialize["title"] = o.Title
	toSerialize["description"] = o.Description
	toSerialize["severity"] = o.Severity
	toSerialize["severityLabel"] = o.SeverityLabel
	toSerialize["startDate"] = o.StartDate
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["flag"] = o.Flag
	toSerialize["tlp"] = o.Tlp
	toSerialize["tlpLabel"] = o.TlpLabel
	toSerialize["pap"] = o.Pap
	toSerialize["papLabel"] = o.PapLabel
	toSerialize["status"] = o.Status
	toSerialize["stage"] = o.Stage
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !IsNil(o.ImpactStatus) {
		toSerialize["impactStatus"] = o.ImpactStatus
	}
	if !IsNil(o.Assignee) {
		toSerialize["assignee"] = o.Assignee
	}
	toSerialize["access"] = o.Access
	if !IsNil(o.CustomFields) {
		toSerialize["customFields"] = o.CustomFields
	}
	if !IsNil(o.UserPermissions) {
		toSerialize["userPermissions"] = o.UserPermissions
	}
	toSerialize["extraData"] = o.ExtraData
	toSerialize["newDate"] = o.NewDate
	if !IsNil(o.InProgressDate) {
		toSerialize["inProgressDate"] = o.InProgressDate
	}
	if !IsNil(o.ClosedDate) {
		toSerialize["closedDate"] = o.ClosedDate
	}
	if !IsNil(o.AlertDate) {
		toSerialize["alertDate"] = o.AlertDate
	}
	if !IsNil(o.AlertNewDate) {
		toSerialize["alertNewDate"] = o.AlertNewDate
	}
	if !IsNil(o.AlertInProgressDate) {
		toSerialize["alertInProgressDate"] = o.AlertInProgressDate
	}
	if !IsNil(o.AlertImportedDate) {
		toSerialize["alertImportedDate"] = o.AlertImportedDate
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
	if !IsNil(o.TimeToResolve) {
		toSerialize["timeToResolve"] = o.TimeToResolve
	}
	if !IsNil(o.HandlingDuration) {
		toSerialize["handlingDuration"] = o.HandlingDuration
	}
	return toSerialize, nil
}

func (o *OutputCase) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"_type",
		"_createdBy",
		"_createdAt",
		"number",
		"title",
		"description",
		"severity",
		"severityLabel",
		"startDate",
		"flag",
		"tlp",
		"tlpLabel",
		"pap",
		"papLabel",
		"status",
		"stage",
		"access",
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

	varOutputCase := _OutputCase{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// Postprocessed by scripts/tolerant-response-decoder: unknown fields are ignored,
	// since TheHive adds fields to its responses in minor versions.
	err = decoder.Decode(&varOutputCase)

	if err != nil {
		return err
	}

	*o = OutputCase(varOutputCase)

	return err
}

type NullableOutputCase struct {
	value *OutputCase
	isSet bool
}

func (v NullableOutputCase) Get() *OutputCase {
	return v.value
}

func (v *NullableOutputCase) Set(val *OutputCase) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputCase) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputCase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputCase(val *OutputCase) *NullableOutputCase {
	return &NullableOutputCase{value: val, isSet: true}
}

func (v NullableOutputCase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputCase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
