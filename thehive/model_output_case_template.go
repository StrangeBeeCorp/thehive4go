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

// checks if the OutputCaseTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutputCaseTemplate{}

// OutputCaseTemplate struct for OutputCaseTemplate
type OutputCaseTemplate struct {
	// Internal unique identifier of the case template.
	UnderscoreId string `json:"_id"`
	// Entity type, always `caseTemplate`.
	UnderscoreType string `json:"_type"`
	// Login of the user who created the case template.
	UnderscoreCreatedBy string `json:"_createdBy"`
	// Login of the user who last updated the case template.
	UnderscoreUpdatedBy *string `json:"_updatedBy,omitempty"`
	// Date and time when the case template was created. Returned as a Unix timestamp in milliseconds.
	UnderscoreCreatedAt int64 `json:"_createdAt"`
	// Date and time when the case template was last updated. Returned as a Unix timestamp in milliseconds.
	UnderscoreUpdatedAt *int64 `json:"_updatedAt,omitempty"`
	// Unique name of the case template within the organization.
	Name string `json:"name"`
	// Name shown to users when selecting this template.
	DisplayName string `json:"displayName"`
	// Prefix automatically added to case titles when the template is applied.
	TitlePrefix *string `json:"titlePrefix,omitempty"`
	// Default description for cases.
	Description *string `json:"description,omitempty"`
	// Default severity level for cases.
	Severity *int32 `json:"severity,omitempty"`
	// Human-readable label for the severity level.
	SeverityLabel *string `json:"severityLabel,omitempty"`
	// Default tags for cases.
	Tags []string `json:"tags,omitempty"`
	// Whether cases are flagged as important by default.
	Flag bool `json:"flag"`
	// Default [Traffic Light Protocol (TLP)](https://www.misp-project.org/taxonomies.html#_tlp) level for cases.
	Tlp *int32 `json:"tlp,omitempty"`
	// Human-readable label for the TLP level.
	TlpLabel *string `json:"tlpLabel,omitempty"`
	// Default [Permissible Actions Protocol (PAP)](https://www.misp-project.org/taxonomies.html#_pap) level for cases.
	Pap *int32 `json:"pap,omitempty"`
	// Human-readable label for the PAP level.
	PapLabel *string `json:"papLabel,omitempty"`
	// Default summary for cases.
	Summary      *string                  `json:"summary,omitempty"`
	CustomFields []OutputCustomFieldValue `json:"customFields,omitempty"`
	Tasks        []OutputTask             `json:"tasks,omitempty"`
	// Extra data fields populated when specific keys are requested in the `extraData` query parameter.
	ExtraData map[string]interface{} `json:"extraData"`
}

type _OutputCaseTemplate OutputCaseTemplate

// NewOutputCaseTemplate instantiates a new OutputCaseTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputCaseTemplate(underscoreId string, underscoreType string, underscoreCreatedBy string, underscoreCreatedAt int64, name string, displayName string, flag bool, extraData map[string]interface{}) *OutputCaseTemplate {
	this := OutputCaseTemplate{}
	this.UnderscoreId = underscoreId
	this.UnderscoreType = underscoreType
	this.UnderscoreCreatedBy = underscoreCreatedBy
	this.UnderscoreCreatedAt = underscoreCreatedAt
	this.Name = name
	this.DisplayName = displayName
	this.Flag = flag
	this.ExtraData = extraData
	return &this
}

// NewOutputCaseTemplateWithDefaults instantiates a new OutputCaseTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputCaseTemplateWithDefaults() *OutputCaseTemplate {
	this := OutputCaseTemplate{}
	return &this
}

// GetUnderscoreId returns the UnderscoreId field value
func (o *OutputCaseTemplate) GetUnderscoreId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnderscoreId
}

// GetUnderscoreIdOk returns a tuple with the UnderscoreId field value
// and a boolean to check if the value has been set.
func (o *OutputCaseTemplate) GetUnderscoreIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreId, true
}

// SetUnderscoreId sets field value
func (o *OutputCaseTemplate) SetUnderscoreId(v string) {
	o.UnderscoreId = v
}

// GetUnderscoreType returns the UnderscoreType field value
func (o *OutputCaseTemplate) GetUnderscoreType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnderscoreType
}

// GetUnderscoreTypeOk returns a tuple with the UnderscoreType field value
// and a boolean to check if the value has been set.
func (o *OutputCaseTemplate) GetUnderscoreTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreType, true
}

// SetUnderscoreType sets field value
func (o *OutputCaseTemplate) SetUnderscoreType(v string) {
	o.UnderscoreType = v
}

// GetUnderscoreCreatedBy returns the UnderscoreCreatedBy field value
func (o *OutputCaseTemplate) GetUnderscoreCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnderscoreCreatedBy
}

// GetUnderscoreCreatedByOk returns a tuple with the UnderscoreCreatedBy field value
// and a boolean to check if the value has been set.
func (o *OutputCaseTemplate) GetUnderscoreCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreCreatedBy, true
}

// SetUnderscoreCreatedBy sets field value
func (o *OutputCaseTemplate) SetUnderscoreCreatedBy(v string) {
	o.UnderscoreCreatedBy = v
}

// GetUnderscoreUpdatedBy returns the UnderscoreUpdatedBy field value if set, zero value otherwise.
func (o *OutputCaseTemplate) GetUnderscoreUpdatedBy() string {
	if o == nil || IsNil(o.UnderscoreUpdatedBy) {
		var ret string
		return ret
	}
	return *o.UnderscoreUpdatedBy
}

// GetUnderscoreUpdatedByOk returns a tuple with the UnderscoreUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCaseTemplate) GetUnderscoreUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UnderscoreUpdatedBy) {
		return nil, false
	}
	return o.UnderscoreUpdatedBy, true
}

// HasUnderscoreUpdatedBy returns a boolean if a field has been set.
func (o *OutputCaseTemplate) HasUnderscoreUpdatedBy() bool {
	if o != nil && !IsNil(o.UnderscoreUpdatedBy) {
		return true
	}

	return false
}

// SetUnderscoreUpdatedBy gets a reference to the given string and assigns it to the UnderscoreUpdatedBy field.
func (o *OutputCaseTemplate) SetUnderscoreUpdatedBy(v string) {
	o.UnderscoreUpdatedBy = &v
}

// GetUnderscoreCreatedAt returns the UnderscoreCreatedAt field value
func (o *OutputCaseTemplate) GetUnderscoreCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UnderscoreCreatedAt
}

// GetUnderscoreCreatedAtOk returns a tuple with the UnderscoreCreatedAt field value
// and a boolean to check if the value has been set.
func (o *OutputCaseTemplate) GetUnderscoreCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreCreatedAt, true
}

// SetUnderscoreCreatedAt sets field value
func (o *OutputCaseTemplate) SetUnderscoreCreatedAt(v int64) {
	o.UnderscoreCreatedAt = v
}

// GetUnderscoreUpdatedAt returns the UnderscoreUpdatedAt field value if set, zero value otherwise.
func (o *OutputCaseTemplate) GetUnderscoreUpdatedAt() int64 {
	if o == nil || IsNil(o.UnderscoreUpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UnderscoreUpdatedAt
}

// GetUnderscoreUpdatedAtOk returns a tuple with the UnderscoreUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCaseTemplate) GetUnderscoreUpdatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.UnderscoreUpdatedAt) {
		return nil, false
	}
	return o.UnderscoreUpdatedAt, true
}

// HasUnderscoreUpdatedAt returns a boolean if a field has been set.
func (o *OutputCaseTemplate) HasUnderscoreUpdatedAt() bool {
	if o != nil && !IsNil(o.UnderscoreUpdatedAt) {
		return true
	}

	return false
}

// SetUnderscoreUpdatedAt gets a reference to the given int64 and assigns it to the UnderscoreUpdatedAt field.
func (o *OutputCaseTemplate) SetUnderscoreUpdatedAt(v int64) {
	o.UnderscoreUpdatedAt = &v
}

// GetName returns the Name field value
func (o *OutputCaseTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OutputCaseTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OutputCaseTemplate) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value
func (o *OutputCaseTemplate) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *OutputCaseTemplate) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *OutputCaseTemplate) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetTitlePrefix returns the TitlePrefix field value if set, zero value otherwise.
func (o *OutputCaseTemplate) GetTitlePrefix() string {
	if o == nil || IsNil(o.TitlePrefix) {
		var ret string
		return ret
	}
	return *o.TitlePrefix
}

// GetTitlePrefixOk returns a tuple with the TitlePrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCaseTemplate) GetTitlePrefixOk() (*string, bool) {
	if o == nil || IsNil(o.TitlePrefix) {
		return nil, false
	}
	return o.TitlePrefix, true
}

// HasTitlePrefix returns a boolean if a field has been set.
func (o *OutputCaseTemplate) HasTitlePrefix() bool {
	if o != nil && !IsNil(o.TitlePrefix) {
		return true
	}

	return false
}

// SetTitlePrefix gets a reference to the given string and assigns it to the TitlePrefix field.
func (o *OutputCaseTemplate) SetTitlePrefix(v string) {
	o.TitlePrefix = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OutputCaseTemplate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCaseTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OutputCaseTemplate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OutputCaseTemplate) SetDescription(v string) {
	o.Description = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *OutputCaseTemplate) GetSeverity() int32 {
	if o == nil || IsNil(o.Severity) {
		var ret int32
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCaseTemplate) GetSeverityOk() (*int32, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *OutputCaseTemplate) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given int32 and assigns it to the Severity field.
func (o *OutputCaseTemplate) SetSeverity(v int32) {
	o.Severity = &v
}

// GetSeverityLabel returns the SeverityLabel field value if set, zero value otherwise.
func (o *OutputCaseTemplate) GetSeverityLabel() string {
	if o == nil || IsNil(o.SeverityLabel) {
		var ret string
		return ret
	}
	return *o.SeverityLabel
}

// GetSeverityLabelOk returns a tuple with the SeverityLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCaseTemplate) GetSeverityLabelOk() (*string, bool) {
	if o == nil || IsNil(o.SeverityLabel) {
		return nil, false
	}
	return o.SeverityLabel, true
}

// HasSeverityLabel returns a boolean if a field has been set.
func (o *OutputCaseTemplate) HasSeverityLabel() bool {
	if o != nil && !IsNil(o.SeverityLabel) {
		return true
	}

	return false
}

// SetSeverityLabel gets a reference to the given string and assigns it to the SeverityLabel field.
func (o *OutputCaseTemplate) SetSeverityLabel(v string) {
	o.SeverityLabel = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *OutputCaseTemplate) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCaseTemplate) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *OutputCaseTemplate) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *OutputCaseTemplate) SetTags(v []string) {
	o.Tags = v
}

// GetFlag returns the Flag field value
func (o *OutputCaseTemplate) GetFlag() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Flag
}

// GetFlagOk returns a tuple with the Flag field value
// and a boolean to check if the value has been set.
func (o *OutputCaseTemplate) GetFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Flag, true
}

// SetFlag sets field value
func (o *OutputCaseTemplate) SetFlag(v bool) {
	o.Flag = v
}

// GetTlp returns the Tlp field value if set, zero value otherwise.
func (o *OutputCaseTemplate) GetTlp() int32 {
	if o == nil || IsNil(o.Tlp) {
		var ret int32
		return ret
	}
	return *o.Tlp
}

// GetTlpOk returns a tuple with the Tlp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCaseTemplate) GetTlpOk() (*int32, bool) {
	if o == nil || IsNil(o.Tlp) {
		return nil, false
	}
	return o.Tlp, true
}

// HasTlp returns a boolean if a field has been set.
func (o *OutputCaseTemplate) HasTlp() bool {
	if o != nil && !IsNil(o.Tlp) {
		return true
	}

	return false
}

// SetTlp gets a reference to the given int32 and assigns it to the Tlp field.
func (o *OutputCaseTemplate) SetTlp(v int32) {
	o.Tlp = &v
}

// GetTlpLabel returns the TlpLabel field value if set, zero value otherwise.
func (o *OutputCaseTemplate) GetTlpLabel() string {
	if o == nil || IsNil(o.TlpLabel) {
		var ret string
		return ret
	}
	return *o.TlpLabel
}

// GetTlpLabelOk returns a tuple with the TlpLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCaseTemplate) GetTlpLabelOk() (*string, bool) {
	if o == nil || IsNil(o.TlpLabel) {
		return nil, false
	}
	return o.TlpLabel, true
}

// HasTlpLabel returns a boolean if a field has been set.
func (o *OutputCaseTemplate) HasTlpLabel() bool {
	if o != nil && !IsNil(o.TlpLabel) {
		return true
	}

	return false
}

// SetTlpLabel gets a reference to the given string and assigns it to the TlpLabel field.
func (o *OutputCaseTemplate) SetTlpLabel(v string) {
	o.TlpLabel = &v
}

// GetPap returns the Pap field value if set, zero value otherwise.
func (o *OutputCaseTemplate) GetPap() int32 {
	if o == nil || IsNil(o.Pap) {
		var ret int32
		return ret
	}
	return *o.Pap
}

// GetPapOk returns a tuple with the Pap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCaseTemplate) GetPapOk() (*int32, bool) {
	if o == nil || IsNil(o.Pap) {
		return nil, false
	}
	return o.Pap, true
}

// HasPap returns a boolean if a field has been set.
func (o *OutputCaseTemplate) HasPap() bool {
	if o != nil && !IsNil(o.Pap) {
		return true
	}

	return false
}

// SetPap gets a reference to the given int32 and assigns it to the Pap field.
func (o *OutputCaseTemplate) SetPap(v int32) {
	o.Pap = &v
}

// GetPapLabel returns the PapLabel field value if set, zero value otherwise.
func (o *OutputCaseTemplate) GetPapLabel() string {
	if o == nil || IsNil(o.PapLabel) {
		var ret string
		return ret
	}
	return *o.PapLabel
}

// GetPapLabelOk returns a tuple with the PapLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCaseTemplate) GetPapLabelOk() (*string, bool) {
	if o == nil || IsNil(o.PapLabel) {
		return nil, false
	}
	return o.PapLabel, true
}

// HasPapLabel returns a boolean if a field has been set.
func (o *OutputCaseTemplate) HasPapLabel() bool {
	if o != nil && !IsNil(o.PapLabel) {
		return true
	}

	return false
}

// SetPapLabel gets a reference to the given string and assigns it to the PapLabel field.
func (o *OutputCaseTemplate) SetPapLabel(v string) {
	o.PapLabel = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *OutputCaseTemplate) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCaseTemplate) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *OutputCaseTemplate) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *OutputCaseTemplate) SetSummary(v string) {
	o.Summary = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *OutputCaseTemplate) GetCustomFields() []OutputCustomFieldValue {
	if o == nil || IsNil(o.CustomFields) {
		var ret []OutputCustomFieldValue
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCaseTemplate) GetCustomFieldsOk() ([]OutputCustomFieldValue, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *OutputCaseTemplate) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given []OutputCustomFieldValue and assigns it to the CustomFields field.
func (o *OutputCaseTemplate) SetCustomFields(v []OutputCustomFieldValue) {
	o.CustomFields = v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *OutputCaseTemplate) GetTasks() []OutputTask {
	if o == nil || IsNil(o.Tasks) {
		var ret []OutputTask
		return ret
	}
	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCaseTemplate) GetTasksOk() ([]OutputTask, bool) {
	if o == nil || IsNil(o.Tasks) {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *OutputCaseTemplate) HasTasks() bool {
	if o != nil && !IsNil(o.Tasks) {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []OutputTask and assigns it to the Tasks field.
func (o *OutputCaseTemplate) SetTasks(v []OutputTask) {
	o.Tasks = v
}

// GetExtraData returns the ExtraData field value
func (o *OutputCaseTemplate) GetExtraData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ExtraData
}

// GetExtraDataOk returns a tuple with the ExtraData field value
// and a boolean to check if the value has been set.
func (o *OutputCaseTemplate) GetExtraDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.ExtraData, true
}

// SetExtraData sets field value
func (o *OutputCaseTemplate) SetExtraData(v map[string]interface{}) {
	o.ExtraData = v
}

func (o OutputCaseTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutputCaseTemplate) ToMap() (map[string]interface{}, error) {
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
	toSerialize["name"] = o.Name
	toSerialize["displayName"] = o.DisplayName
	if !IsNil(o.TitlePrefix) {
		toSerialize["titlePrefix"] = o.TitlePrefix
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.SeverityLabel) {
		toSerialize["severityLabel"] = o.SeverityLabel
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["flag"] = o.Flag
	if !IsNil(o.Tlp) {
		toSerialize["tlp"] = o.Tlp
	}
	if !IsNil(o.TlpLabel) {
		toSerialize["tlpLabel"] = o.TlpLabel
	}
	if !IsNil(o.Pap) {
		toSerialize["pap"] = o.Pap
	}
	if !IsNil(o.PapLabel) {
		toSerialize["papLabel"] = o.PapLabel
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
	toSerialize["extraData"] = o.ExtraData
	return toSerialize, nil
}

func (o *OutputCaseTemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"_type",
		"_createdBy",
		"_createdAt",
		"name",
		"displayName",
		"flag",
		"extraData",
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

	varOutputCaseTemplate := _OutputCaseTemplate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// Postprocessed by scripts/tolerant-response-decoder: unknown fields are ignored,
	// since TheHive adds fields to its responses in minor versions.
	err = decoder.Decode(&varOutputCaseTemplate)

	if err != nil {
		return err
	}

	*o = OutputCaseTemplate(varOutputCaseTemplate)

	return err
}

type NullableOutputCaseTemplate struct {
	value *OutputCaseTemplate
	isSet bool
}

func (v NullableOutputCaseTemplate) Get() *OutputCaseTemplate {
	return v.value
}

func (v *NullableOutputCaseTemplate) Set(val *OutputCaseTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputCaseTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputCaseTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputCaseTemplate(val *OutputCaseTemplate) *NullableOutputCaseTemplate {
	return &NullableOutputCaseTemplate{value: val, isSet: true}
}

func (v NullableOutputCaseTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputCaseTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
