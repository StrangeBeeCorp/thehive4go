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

// checks if the OutputPattern type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutputPattern{}

// OutputPattern struct for OutputPattern
type OutputPattern struct {
	// Unique identifier of the attack technique.
	UnderscoreId string `json:"_id"`
	// Entity type, always \"Pattern\".
	UnderscoreType string `json:"_type"`
	// Login of the user who created the attack technique record.
	UnderscoreCreatedBy string `json:"_createdBy"`
	// Login of the user who last updated the attack technique record.
	UnderscoreUpdatedBy *string `json:"_updatedBy,omitempty"`
	// Date and time when the attack technique record was created. Returned as a Unix timestamp in milliseconds.
	UnderscoreCreatedAt int64 `json:"_createdAt"`
	// Date and time when the attack technique record was last updated. Returned as a Unix timestamp in milliseconds.
	UnderscoreUpdatedAt *int64 `json:"_updatedAt,omitempty"`
	// MITRE ATT&CK pattern identifier.
	PatternId string `json:"patternId"`
	// Name of the attack technique.
	Name string `json:"name"`
	// Description of the attack technique.
	Description *string `json:"description,omitempty"`
	// MITRE ATT&CK tactics associated with this technique.
	Tactics []string `json:"tactics,omitempty"`
	// URL of the MITRE ATT&CK page for this technique.
	Url string `json:"url"`
	// MITRE ATT&CK object type for this entry.
	PatternType string `json:"patternType"`
	// CAPEC (Common Attack Pattern Enumeration and Classification) identifier, if applicable.
	CapecId *string `json:"capecId,omitempty"`
	// URL of the CAPEC entry for this technique, if applicable.
	CapecUrl *string `json:"capecUrl,omitempty"`
	// Whether this technique has been revoked in the MITRE ATT&CK framework.
	Revoked bool `json:"revoked"`
	// Data sources that can be used to detect this technique.
	DataSources []string `json:"dataSources,omitempty"`
	// Defenses this technique can bypass.
	DefenseBypassed []string `json:"defenseBypassed,omitempty"`
	// Guidance for detecting this technique.
	Detection *string `json:"detection,omitempty"`
	// Permissions required to execute this technique.
	PermissionsRequired []string `json:"permissionsRequired,omitempty"`
	// Target platforms for this technique.
	Platforms []string `json:"platforms,omitempty"`
	// Whether this technique can be executed remotely.
	RemoteSupport bool `json:"remoteSupport"`
	// System requirements for executing this technique.
	SystemRequirements []string `json:"systemRequirements,omitempty"`
	// Version of this technique in the MITRE ATT&CK framework.
	Version *string `json:"version,omitempty"`
	// Extra data fields populated when specific keys are requested in the `extraData` query parameter.
	ExtraData map[string]interface{} `json:"extraData"`
}

type _OutputPattern OutputPattern

// NewOutputPattern instantiates a new OutputPattern object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputPattern(underscoreId string, underscoreType string, underscoreCreatedBy string, underscoreCreatedAt int64, patternId string, name string, url string, patternType string, revoked bool, remoteSupport bool, extraData map[string]interface{}) *OutputPattern {
	this := OutputPattern{}
	this.UnderscoreId = underscoreId
	this.UnderscoreType = underscoreType
	this.UnderscoreCreatedBy = underscoreCreatedBy
	this.UnderscoreCreatedAt = underscoreCreatedAt
	this.PatternId = patternId
	this.Name = name
	this.Url = url
	this.PatternType = patternType
	this.Revoked = revoked
	this.RemoteSupport = remoteSupport
	this.ExtraData = extraData
	return &this
}

// NewOutputPatternWithDefaults instantiates a new OutputPattern object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputPatternWithDefaults() *OutputPattern {
	this := OutputPattern{}
	return &this
}

// GetUnderscoreId returns the UnderscoreId field value
func (o *OutputPattern) GetUnderscoreId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnderscoreId
}

// GetUnderscoreIdOk returns a tuple with the UnderscoreId field value
// and a boolean to check if the value has been set.
func (o *OutputPattern) GetUnderscoreIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreId, true
}

// SetUnderscoreId sets field value
func (o *OutputPattern) SetUnderscoreId(v string) {
	o.UnderscoreId = v
}

// GetUnderscoreType returns the UnderscoreType field value
func (o *OutputPattern) GetUnderscoreType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnderscoreType
}

// GetUnderscoreTypeOk returns a tuple with the UnderscoreType field value
// and a boolean to check if the value has been set.
func (o *OutputPattern) GetUnderscoreTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreType, true
}

// SetUnderscoreType sets field value
func (o *OutputPattern) SetUnderscoreType(v string) {
	o.UnderscoreType = v
}

// GetUnderscoreCreatedBy returns the UnderscoreCreatedBy field value
func (o *OutputPattern) GetUnderscoreCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnderscoreCreatedBy
}

// GetUnderscoreCreatedByOk returns a tuple with the UnderscoreCreatedBy field value
// and a boolean to check if the value has been set.
func (o *OutputPattern) GetUnderscoreCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreCreatedBy, true
}

// SetUnderscoreCreatedBy sets field value
func (o *OutputPattern) SetUnderscoreCreatedBy(v string) {
	o.UnderscoreCreatedBy = v
}

// GetUnderscoreUpdatedBy returns the UnderscoreUpdatedBy field value if set, zero value otherwise.
func (o *OutputPattern) GetUnderscoreUpdatedBy() string {
	if o == nil || IsNil(o.UnderscoreUpdatedBy) {
		var ret string
		return ret
	}
	return *o.UnderscoreUpdatedBy
}

// GetUnderscoreUpdatedByOk returns a tuple with the UnderscoreUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputPattern) GetUnderscoreUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UnderscoreUpdatedBy) {
		return nil, false
	}
	return o.UnderscoreUpdatedBy, true
}

// HasUnderscoreUpdatedBy returns a boolean if a field has been set.
func (o *OutputPattern) HasUnderscoreUpdatedBy() bool {
	if o != nil && !IsNil(o.UnderscoreUpdatedBy) {
		return true
	}

	return false
}

// SetUnderscoreUpdatedBy gets a reference to the given string and assigns it to the UnderscoreUpdatedBy field.
func (o *OutputPattern) SetUnderscoreUpdatedBy(v string) {
	o.UnderscoreUpdatedBy = &v
}

// GetUnderscoreCreatedAt returns the UnderscoreCreatedAt field value
func (o *OutputPattern) GetUnderscoreCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UnderscoreCreatedAt
}

// GetUnderscoreCreatedAtOk returns a tuple with the UnderscoreCreatedAt field value
// and a boolean to check if the value has been set.
func (o *OutputPattern) GetUnderscoreCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreCreatedAt, true
}

// SetUnderscoreCreatedAt sets field value
func (o *OutputPattern) SetUnderscoreCreatedAt(v int64) {
	o.UnderscoreCreatedAt = v
}

// GetUnderscoreUpdatedAt returns the UnderscoreUpdatedAt field value if set, zero value otherwise.
func (o *OutputPattern) GetUnderscoreUpdatedAt() int64 {
	if o == nil || IsNil(o.UnderscoreUpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UnderscoreUpdatedAt
}

// GetUnderscoreUpdatedAtOk returns a tuple with the UnderscoreUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputPattern) GetUnderscoreUpdatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.UnderscoreUpdatedAt) {
		return nil, false
	}
	return o.UnderscoreUpdatedAt, true
}

// HasUnderscoreUpdatedAt returns a boolean if a field has been set.
func (o *OutputPattern) HasUnderscoreUpdatedAt() bool {
	if o != nil && !IsNil(o.UnderscoreUpdatedAt) {
		return true
	}

	return false
}

// SetUnderscoreUpdatedAt gets a reference to the given int64 and assigns it to the UnderscoreUpdatedAt field.
func (o *OutputPattern) SetUnderscoreUpdatedAt(v int64) {
	o.UnderscoreUpdatedAt = &v
}

// GetPatternId returns the PatternId field value
func (o *OutputPattern) GetPatternId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PatternId
}

// GetPatternIdOk returns a tuple with the PatternId field value
// and a boolean to check if the value has been set.
func (o *OutputPattern) GetPatternIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PatternId, true
}

// SetPatternId sets field value
func (o *OutputPattern) SetPatternId(v string) {
	o.PatternId = v
}

// GetName returns the Name field value
func (o *OutputPattern) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OutputPattern) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OutputPattern) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OutputPattern) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputPattern) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OutputPattern) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OutputPattern) SetDescription(v string) {
	o.Description = &v
}

// GetTactics returns the Tactics field value if set, zero value otherwise.
func (o *OutputPattern) GetTactics() []string {
	if o == nil || IsNil(o.Tactics) {
		var ret []string
		return ret
	}
	return o.Tactics
}

// GetTacticsOk returns a tuple with the Tactics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputPattern) GetTacticsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tactics) {
		return nil, false
	}
	return o.Tactics, true
}

// HasTactics returns a boolean if a field has been set.
func (o *OutputPattern) HasTactics() bool {
	if o != nil && !IsNil(o.Tactics) {
		return true
	}

	return false
}

// SetTactics gets a reference to the given []string and assigns it to the Tactics field.
func (o *OutputPattern) SetTactics(v []string) {
	o.Tactics = v
}

// GetUrl returns the Url field value
func (o *OutputPattern) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *OutputPattern) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *OutputPattern) SetUrl(v string) {
	o.Url = v
}

// GetPatternType returns the PatternType field value
func (o *OutputPattern) GetPatternType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PatternType
}

// GetPatternTypeOk returns a tuple with the PatternType field value
// and a boolean to check if the value has been set.
func (o *OutputPattern) GetPatternTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PatternType, true
}

// SetPatternType sets field value
func (o *OutputPattern) SetPatternType(v string) {
	o.PatternType = v
}

// GetCapecId returns the CapecId field value if set, zero value otherwise.
func (o *OutputPattern) GetCapecId() string {
	if o == nil || IsNil(o.CapecId) {
		var ret string
		return ret
	}
	return *o.CapecId
}

// GetCapecIdOk returns a tuple with the CapecId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputPattern) GetCapecIdOk() (*string, bool) {
	if o == nil || IsNil(o.CapecId) {
		return nil, false
	}
	return o.CapecId, true
}

// HasCapecId returns a boolean if a field has been set.
func (o *OutputPattern) HasCapecId() bool {
	if o != nil && !IsNil(o.CapecId) {
		return true
	}

	return false
}

// SetCapecId gets a reference to the given string and assigns it to the CapecId field.
func (o *OutputPattern) SetCapecId(v string) {
	o.CapecId = &v
}

// GetCapecUrl returns the CapecUrl field value if set, zero value otherwise.
func (o *OutputPattern) GetCapecUrl() string {
	if o == nil || IsNil(o.CapecUrl) {
		var ret string
		return ret
	}
	return *o.CapecUrl
}

// GetCapecUrlOk returns a tuple with the CapecUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputPattern) GetCapecUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CapecUrl) {
		return nil, false
	}
	return o.CapecUrl, true
}

// HasCapecUrl returns a boolean if a field has been set.
func (o *OutputPattern) HasCapecUrl() bool {
	if o != nil && !IsNil(o.CapecUrl) {
		return true
	}

	return false
}

// SetCapecUrl gets a reference to the given string and assigns it to the CapecUrl field.
func (o *OutputPattern) SetCapecUrl(v string) {
	o.CapecUrl = &v
}

// GetRevoked returns the Revoked field value
func (o *OutputPattern) GetRevoked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Revoked
}

// GetRevokedOk returns a tuple with the Revoked field value
// and a boolean to check if the value has been set.
func (o *OutputPattern) GetRevokedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revoked, true
}

// SetRevoked sets field value
func (o *OutputPattern) SetRevoked(v bool) {
	o.Revoked = v
}

// GetDataSources returns the DataSources field value if set, zero value otherwise.
func (o *OutputPattern) GetDataSources() []string {
	if o == nil || IsNil(o.DataSources) {
		var ret []string
		return ret
	}
	return o.DataSources
}

// GetDataSourcesOk returns a tuple with the DataSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputPattern) GetDataSourcesOk() ([]string, bool) {
	if o == nil || IsNil(o.DataSources) {
		return nil, false
	}
	return o.DataSources, true
}

// HasDataSources returns a boolean if a field has been set.
func (o *OutputPattern) HasDataSources() bool {
	if o != nil && !IsNil(o.DataSources) {
		return true
	}

	return false
}

// SetDataSources gets a reference to the given []string and assigns it to the DataSources field.
func (o *OutputPattern) SetDataSources(v []string) {
	o.DataSources = v
}

// GetDefenseBypassed returns the DefenseBypassed field value if set, zero value otherwise.
func (o *OutputPattern) GetDefenseBypassed() []string {
	if o == nil || IsNil(o.DefenseBypassed) {
		var ret []string
		return ret
	}
	return o.DefenseBypassed
}

// GetDefenseBypassedOk returns a tuple with the DefenseBypassed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputPattern) GetDefenseBypassedOk() ([]string, bool) {
	if o == nil || IsNil(o.DefenseBypassed) {
		return nil, false
	}
	return o.DefenseBypassed, true
}

// HasDefenseBypassed returns a boolean if a field has been set.
func (o *OutputPattern) HasDefenseBypassed() bool {
	if o != nil && !IsNil(o.DefenseBypassed) {
		return true
	}

	return false
}

// SetDefenseBypassed gets a reference to the given []string and assigns it to the DefenseBypassed field.
func (o *OutputPattern) SetDefenseBypassed(v []string) {
	o.DefenseBypassed = v
}

// GetDetection returns the Detection field value if set, zero value otherwise.
func (o *OutputPattern) GetDetection() string {
	if o == nil || IsNil(o.Detection) {
		var ret string
		return ret
	}
	return *o.Detection
}

// GetDetectionOk returns a tuple with the Detection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputPattern) GetDetectionOk() (*string, bool) {
	if o == nil || IsNil(o.Detection) {
		return nil, false
	}
	return o.Detection, true
}

// HasDetection returns a boolean if a field has been set.
func (o *OutputPattern) HasDetection() bool {
	if o != nil && !IsNil(o.Detection) {
		return true
	}

	return false
}

// SetDetection gets a reference to the given string and assigns it to the Detection field.
func (o *OutputPattern) SetDetection(v string) {
	o.Detection = &v
}

// GetPermissionsRequired returns the PermissionsRequired field value if set, zero value otherwise.
func (o *OutputPattern) GetPermissionsRequired() []string {
	if o == nil || IsNil(o.PermissionsRequired) {
		var ret []string
		return ret
	}
	return o.PermissionsRequired
}

// GetPermissionsRequiredOk returns a tuple with the PermissionsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputPattern) GetPermissionsRequiredOk() ([]string, bool) {
	if o == nil || IsNil(o.PermissionsRequired) {
		return nil, false
	}
	return o.PermissionsRequired, true
}

// HasPermissionsRequired returns a boolean if a field has been set.
func (o *OutputPattern) HasPermissionsRequired() bool {
	if o != nil && !IsNil(o.PermissionsRequired) {
		return true
	}

	return false
}

// SetPermissionsRequired gets a reference to the given []string and assigns it to the PermissionsRequired field.
func (o *OutputPattern) SetPermissionsRequired(v []string) {
	o.PermissionsRequired = v
}

// GetPlatforms returns the Platforms field value if set, zero value otherwise.
func (o *OutputPattern) GetPlatforms() []string {
	if o == nil || IsNil(o.Platforms) {
		var ret []string
		return ret
	}
	return o.Platforms
}

// GetPlatformsOk returns a tuple with the Platforms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputPattern) GetPlatformsOk() ([]string, bool) {
	if o == nil || IsNil(o.Platforms) {
		return nil, false
	}
	return o.Platforms, true
}

// HasPlatforms returns a boolean if a field has been set.
func (o *OutputPattern) HasPlatforms() bool {
	if o != nil && !IsNil(o.Platforms) {
		return true
	}

	return false
}

// SetPlatforms gets a reference to the given []string and assigns it to the Platforms field.
func (o *OutputPattern) SetPlatforms(v []string) {
	o.Platforms = v
}

// GetRemoteSupport returns the RemoteSupport field value
func (o *OutputPattern) GetRemoteSupport() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RemoteSupport
}

// GetRemoteSupportOk returns a tuple with the RemoteSupport field value
// and a boolean to check if the value has been set.
func (o *OutputPattern) GetRemoteSupportOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoteSupport, true
}

// SetRemoteSupport sets field value
func (o *OutputPattern) SetRemoteSupport(v bool) {
	o.RemoteSupport = v
}

// GetSystemRequirements returns the SystemRequirements field value if set, zero value otherwise.
func (o *OutputPattern) GetSystemRequirements() []string {
	if o == nil || IsNil(o.SystemRequirements) {
		var ret []string
		return ret
	}
	return o.SystemRequirements
}

// GetSystemRequirementsOk returns a tuple with the SystemRequirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputPattern) GetSystemRequirementsOk() ([]string, bool) {
	if o == nil || IsNil(o.SystemRequirements) {
		return nil, false
	}
	return o.SystemRequirements, true
}

// HasSystemRequirements returns a boolean if a field has been set.
func (o *OutputPattern) HasSystemRequirements() bool {
	if o != nil && !IsNil(o.SystemRequirements) {
		return true
	}

	return false
}

// SetSystemRequirements gets a reference to the given []string and assigns it to the SystemRequirements field.
func (o *OutputPattern) SetSystemRequirements(v []string) {
	o.SystemRequirements = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *OutputPattern) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputPattern) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *OutputPattern) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *OutputPattern) SetVersion(v string) {
	o.Version = &v
}

// GetExtraData returns the ExtraData field value
func (o *OutputPattern) GetExtraData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ExtraData
}

// GetExtraDataOk returns a tuple with the ExtraData field value
// and a boolean to check if the value has been set.
func (o *OutputPattern) GetExtraDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.ExtraData, true
}

// SetExtraData sets field value
func (o *OutputPattern) SetExtraData(v map[string]interface{}) {
	o.ExtraData = v
}

func (o OutputPattern) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutputPattern) ToMap() (map[string]interface{}, error) {
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
	toSerialize["patternId"] = o.PatternId
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Tactics) {
		toSerialize["tactics"] = o.Tactics
	}
	toSerialize["url"] = o.Url
	toSerialize["patternType"] = o.PatternType
	if !IsNil(o.CapecId) {
		toSerialize["capecId"] = o.CapecId
	}
	if !IsNil(o.CapecUrl) {
		toSerialize["capecUrl"] = o.CapecUrl
	}
	toSerialize["revoked"] = o.Revoked
	if !IsNil(o.DataSources) {
		toSerialize["dataSources"] = o.DataSources
	}
	if !IsNil(o.DefenseBypassed) {
		toSerialize["defenseBypassed"] = o.DefenseBypassed
	}
	if !IsNil(o.Detection) {
		toSerialize["detection"] = o.Detection
	}
	if !IsNil(o.PermissionsRequired) {
		toSerialize["permissionsRequired"] = o.PermissionsRequired
	}
	if !IsNil(o.Platforms) {
		toSerialize["platforms"] = o.Platforms
	}
	toSerialize["remoteSupport"] = o.RemoteSupport
	if !IsNil(o.SystemRequirements) {
		toSerialize["systemRequirements"] = o.SystemRequirements
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	toSerialize["extraData"] = o.ExtraData
	return toSerialize, nil
}

func (o *OutputPattern) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"_type",
		"_createdBy",
		"_createdAt",
		"patternId",
		"name",
		"url",
		"patternType",
		"revoked",
		"remoteSupport",
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

	varOutputPattern := _OutputPattern{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// Postprocessed by scripts/tolerant-response-decoder: unknown fields are ignored,
	// since TheHive adds fields to its responses in minor versions.
	err = decoder.Decode(&varOutputPattern)

	if err != nil {
		return err
	}

	*o = OutputPattern(varOutputPattern)

	return err
}

type NullableOutputPattern struct {
	value *OutputPattern
	isSet bool
}

func (v NullableOutputPattern) Get() *OutputPattern {
	return v.value
}

func (v *NullableOutputPattern) Set(val *OutputPattern) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputPattern) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputPattern) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputPattern(val *OutputPattern) *NullableOutputPattern {
	return &NullableOutputPattern{value: val, isSet: true}
}

func (v NullableOutputPattern) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputPattern) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
