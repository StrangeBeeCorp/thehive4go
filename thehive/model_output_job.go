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

// checks if the OutputJob type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutputJob{}

// OutputJob struct for OutputJob
type OutputJob struct {
	// Internal unique identifier of the job.
	UnderscoreId string `json:"_id"`
	// Entity type, always `case_artifact_job`.
	UnderscoreType string `json:"_type"`
	// Login of the user who launched the job.
	UnderscoreCreatedBy string `json:"_createdBy"`
	// Login of the user who last updated the job.
	UnderscoreUpdatedBy *string `json:"_updatedBy,omitempty"`
	// Date and time when the job was created. Returned as a Unix timestamp in milliseconds.
	UnderscoreCreatedAt int64 `json:"_createdAt"`
	// Date and time when the job was last updated. Returned as a Unix timestamp in milliseconds.
	UnderscoreUpdatedAt *int64 `json:"_updatedAt,omitempty"`
	// ID of the analyzer that ran.
	AnalyzerId string `json:"analyzerId"`
	// Display name of the analyzer that ran.
	AnalyzerName string `json:"analyzerName"`
	// Definition of the analyzer that ran, as returned by Cortex.
	AnalyzerDefinition string `json:"analyzerDefinition"`
	// Status of the job. One of `Waiting`, `InProgress`, `Success`, `Failure`, or `Deleted`.
	Status string `json:"status"`
	// Date and time when the job started. Returned as a Unix timestamp in milliseconds.
	StartDate int64 `json:"startDate"`
	// Date and time when the job finished. Returned as a Unix timestamp in milliseconds.
	EndDate *int64 `json:"endDate,omitempty"`
	// Cortex analysis report for the job, once the analyzer has completed.
	Report map[string]interface{} `json:"report,omitempty"`
	// Name of the Cortex server the analyzer ran on.
	CortexId string `json:"cortexId"`
	// ID of the corresponding job on the Cortex server.
	CortexJobId string `json:"cortexJobId"`
	// Internal unique identifier of the job, identical to `_id`.
	Id string `json:"id"`
	// The observable the job analyzed.
	CaseArtifact map[string]interface{} `json:"case_artifact,omitempty"`
	// Serialized JSON array of follow-up operations produced by the analyzer report.
	Operations string `json:"operations"`
}

type _OutputJob OutputJob

// NewOutputJob instantiates a new OutputJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputJob(underscoreId string, underscoreType string, underscoreCreatedBy string, underscoreCreatedAt int64, analyzerId string, analyzerName string, analyzerDefinition string, status string, startDate int64, cortexId string, cortexJobId string, id string, operations string) *OutputJob {
	this := OutputJob{}
	this.UnderscoreId = underscoreId
	this.UnderscoreType = underscoreType
	this.UnderscoreCreatedBy = underscoreCreatedBy
	this.UnderscoreCreatedAt = underscoreCreatedAt
	this.AnalyzerId = analyzerId
	this.AnalyzerName = analyzerName
	this.AnalyzerDefinition = analyzerDefinition
	this.Status = status
	this.StartDate = startDate
	this.CortexId = cortexId
	this.CortexJobId = cortexJobId
	this.Id = id
	this.Operations = operations
	return &this
}

// NewOutputJobWithDefaults instantiates a new OutputJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputJobWithDefaults() *OutputJob {
	this := OutputJob{}
	return &this
}

// GetUnderscoreId returns the UnderscoreId field value
func (o *OutputJob) GetUnderscoreId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnderscoreId
}

// GetUnderscoreIdOk returns a tuple with the UnderscoreId field value
// and a boolean to check if the value has been set.
func (o *OutputJob) GetUnderscoreIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreId, true
}

// SetUnderscoreId sets field value
func (o *OutputJob) SetUnderscoreId(v string) {
	o.UnderscoreId = v
}

// GetUnderscoreType returns the UnderscoreType field value
func (o *OutputJob) GetUnderscoreType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnderscoreType
}

// GetUnderscoreTypeOk returns a tuple with the UnderscoreType field value
// and a boolean to check if the value has been set.
func (o *OutputJob) GetUnderscoreTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreType, true
}

// SetUnderscoreType sets field value
func (o *OutputJob) SetUnderscoreType(v string) {
	o.UnderscoreType = v
}

// GetUnderscoreCreatedBy returns the UnderscoreCreatedBy field value
func (o *OutputJob) GetUnderscoreCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnderscoreCreatedBy
}

// GetUnderscoreCreatedByOk returns a tuple with the UnderscoreCreatedBy field value
// and a boolean to check if the value has been set.
func (o *OutputJob) GetUnderscoreCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreCreatedBy, true
}

// SetUnderscoreCreatedBy sets field value
func (o *OutputJob) SetUnderscoreCreatedBy(v string) {
	o.UnderscoreCreatedBy = v
}

// GetUnderscoreUpdatedBy returns the UnderscoreUpdatedBy field value if set, zero value otherwise.
func (o *OutputJob) GetUnderscoreUpdatedBy() string {
	if o == nil || IsNil(o.UnderscoreUpdatedBy) {
		var ret string
		return ret
	}
	return *o.UnderscoreUpdatedBy
}

// GetUnderscoreUpdatedByOk returns a tuple with the UnderscoreUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputJob) GetUnderscoreUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UnderscoreUpdatedBy) {
		return nil, false
	}
	return o.UnderscoreUpdatedBy, true
}

// HasUnderscoreUpdatedBy returns a boolean if a field has been set.
func (o *OutputJob) HasUnderscoreUpdatedBy() bool {
	if o != nil && !IsNil(o.UnderscoreUpdatedBy) {
		return true
	}

	return false
}

// SetUnderscoreUpdatedBy gets a reference to the given string and assigns it to the UnderscoreUpdatedBy field.
func (o *OutputJob) SetUnderscoreUpdatedBy(v string) {
	o.UnderscoreUpdatedBy = &v
}

// GetUnderscoreCreatedAt returns the UnderscoreCreatedAt field value
func (o *OutputJob) GetUnderscoreCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UnderscoreCreatedAt
}

// GetUnderscoreCreatedAtOk returns a tuple with the UnderscoreCreatedAt field value
// and a boolean to check if the value has been set.
func (o *OutputJob) GetUnderscoreCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreCreatedAt, true
}

// SetUnderscoreCreatedAt sets field value
func (o *OutputJob) SetUnderscoreCreatedAt(v int64) {
	o.UnderscoreCreatedAt = v
}

// GetUnderscoreUpdatedAt returns the UnderscoreUpdatedAt field value if set, zero value otherwise.
func (o *OutputJob) GetUnderscoreUpdatedAt() int64 {
	if o == nil || IsNil(o.UnderscoreUpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UnderscoreUpdatedAt
}

// GetUnderscoreUpdatedAtOk returns a tuple with the UnderscoreUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputJob) GetUnderscoreUpdatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.UnderscoreUpdatedAt) {
		return nil, false
	}
	return o.UnderscoreUpdatedAt, true
}

// HasUnderscoreUpdatedAt returns a boolean if a field has been set.
func (o *OutputJob) HasUnderscoreUpdatedAt() bool {
	if o != nil && !IsNil(o.UnderscoreUpdatedAt) {
		return true
	}

	return false
}

// SetUnderscoreUpdatedAt gets a reference to the given int64 and assigns it to the UnderscoreUpdatedAt field.
func (o *OutputJob) SetUnderscoreUpdatedAt(v int64) {
	o.UnderscoreUpdatedAt = &v
}

// GetAnalyzerId returns the AnalyzerId field value
func (o *OutputJob) GetAnalyzerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AnalyzerId
}

// GetAnalyzerIdOk returns a tuple with the AnalyzerId field value
// and a boolean to check if the value has been set.
func (o *OutputJob) GetAnalyzerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnalyzerId, true
}

// SetAnalyzerId sets field value
func (o *OutputJob) SetAnalyzerId(v string) {
	o.AnalyzerId = v
}

// GetAnalyzerName returns the AnalyzerName field value
func (o *OutputJob) GetAnalyzerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AnalyzerName
}

// GetAnalyzerNameOk returns a tuple with the AnalyzerName field value
// and a boolean to check if the value has been set.
func (o *OutputJob) GetAnalyzerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnalyzerName, true
}

// SetAnalyzerName sets field value
func (o *OutputJob) SetAnalyzerName(v string) {
	o.AnalyzerName = v
}

// GetAnalyzerDefinition returns the AnalyzerDefinition field value
func (o *OutputJob) GetAnalyzerDefinition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AnalyzerDefinition
}

// GetAnalyzerDefinitionOk returns a tuple with the AnalyzerDefinition field value
// and a boolean to check if the value has been set.
func (o *OutputJob) GetAnalyzerDefinitionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnalyzerDefinition, true
}

// SetAnalyzerDefinition sets field value
func (o *OutputJob) SetAnalyzerDefinition(v string) {
	o.AnalyzerDefinition = v
}

// GetStatus returns the Status field value
func (o *OutputJob) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OutputJob) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *OutputJob) SetStatus(v string) {
	o.Status = v
}

// GetStartDate returns the StartDate field value
func (o *OutputJob) GetStartDate() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *OutputJob) GetStartDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *OutputJob) SetStartDate(v int64) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *OutputJob) GetEndDate() int64 {
	if o == nil || IsNil(o.EndDate) {
		var ret int64
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputJob) GetEndDateOk() (*int64, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *OutputJob) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given int64 and assigns it to the EndDate field.
func (o *OutputJob) SetEndDate(v int64) {
	o.EndDate = &v
}

// GetReport returns the Report field value if set, zero value otherwise.
func (o *OutputJob) GetReport() map[string]interface{} {
	if o == nil || IsNil(o.Report) {
		var ret map[string]interface{}
		return ret
	}
	return o.Report
}

// GetReportOk returns a tuple with the Report field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputJob) GetReportOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Report) {
		return map[string]interface{}{}, false
	}
	return o.Report, true
}

// HasReport returns a boolean if a field has been set.
func (o *OutputJob) HasReport() bool {
	if o != nil && !IsNil(o.Report) {
		return true
	}

	return false
}

// SetReport gets a reference to the given map[string]interface{} and assigns it to the Report field.
func (o *OutputJob) SetReport(v map[string]interface{}) {
	o.Report = v
}

// GetCortexId returns the CortexId field value
func (o *OutputJob) GetCortexId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CortexId
}

// GetCortexIdOk returns a tuple with the CortexId field value
// and a boolean to check if the value has been set.
func (o *OutputJob) GetCortexIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CortexId, true
}

// SetCortexId sets field value
func (o *OutputJob) SetCortexId(v string) {
	o.CortexId = v
}

// GetCortexJobId returns the CortexJobId field value
func (o *OutputJob) GetCortexJobId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CortexJobId
}

// GetCortexJobIdOk returns a tuple with the CortexJobId field value
// and a boolean to check if the value has been set.
func (o *OutputJob) GetCortexJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CortexJobId, true
}

// SetCortexJobId sets field value
func (o *OutputJob) SetCortexJobId(v string) {
	o.CortexJobId = v
}

// GetId returns the Id field value
func (o *OutputJob) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OutputJob) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OutputJob) SetId(v string) {
	o.Id = v
}

// GetCaseArtifact returns the CaseArtifact field value if set, zero value otherwise.
func (o *OutputJob) GetCaseArtifact() map[string]interface{} {
	if o == nil || IsNil(o.CaseArtifact) {
		var ret map[string]interface{}
		return ret
	}
	return o.CaseArtifact
}

// GetCaseArtifactOk returns a tuple with the CaseArtifact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputJob) GetCaseArtifactOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CaseArtifact) {
		return map[string]interface{}{}, false
	}
	return o.CaseArtifact, true
}

// HasCaseArtifact returns a boolean if a field has been set.
func (o *OutputJob) HasCaseArtifact() bool {
	if o != nil && !IsNil(o.CaseArtifact) {
		return true
	}

	return false
}

// SetCaseArtifact gets a reference to the given map[string]interface{} and assigns it to the CaseArtifact field.
func (o *OutputJob) SetCaseArtifact(v map[string]interface{}) {
	o.CaseArtifact = v
}

// GetOperations returns the Operations field value
func (o *OutputJob) GetOperations() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value
// and a boolean to check if the value has been set.
func (o *OutputJob) GetOperationsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operations, true
}

// SetOperations sets field value
func (o *OutputJob) SetOperations(v string) {
	o.Operations = v
}

func (o OutputJob) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutputJob) ToMap() (map[string]interface{}, error) {
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
	toSerialize["analyzerId"] = o.AnalyzerId
	toSerialize["analyzerName"] = o.AnalyzerName
	toSerialize["analyzerDefinition"] = o.AnalyzerDefinition
	toSerialize["status"] = o.Status
	toSerialize["startDate"] = o.StartDate
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.Report) {
		toSerialize["report"] = o.Report
	}
	toSerialize["cortexId"] = o.CortexId
	toSerialize["cortexJobId"] = o.CortexJobId
	toSerialize["id"] = o.Id
	if !IsNil(o.CaseArtifact) {
		toSerialize["case_artifact"] = o.CaseArtifact
	}
	toSerialize["operations"] = o.Operations
	return toSerialize, nil
}

func (o *OutputJob) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"_type",
		"_createdBy",
		"_createdAt",
		"analyzerId",
		"analyzerName",
		"analyzerDefinition",
		"status",
		"startDate",
		"cortexId",
		"cortexJobId",
		"id",
		"operations",
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

	varOutputJob := _OutputJob{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// Postprocessed by scripts/tolerant-response-decoder: unknown fields are ignored,
	// since TheHive adds fields to its responses in minor versions.
	err = decoder.Decode(&varOutputJob)

	if err != nil {
		return err
	}

	*o = OutputJob(varOutputJob)

	return err
}

type NullableOutputJob struct {
	value *OutputJob
	isSet bool
}

func (v NullableOutputJob) Get() *OutputJob {
	return v.value
}

func (v *NullableOutputJob) Set(val *OutputJob) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputJob) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputJob(val *OutputJob) *NullableOutputJob {
	return &NullableOutputJob{value: val, isSet: true}
}

func (v NullableOutputJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
