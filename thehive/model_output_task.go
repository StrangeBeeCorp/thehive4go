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

// checks if the OutputTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutputTask{}

// OutputTask struct for OutputTask
type OutputTask struct {
	// Internal unique identifier of the task.
	UnderscoreId string `json:"_id"`
	// Entity type, always `Task`.
	UnderscoreType string `json:"_type"`
	// Login of the user who created the task.
	UnderscoreCreatedBy string `json:"_createdBy"`
	// Login of the user who last updated the task.
	UnderscoreUpdatedBy *string `json:"_updatedBy,omitempty"`
	// Date and time when the task was created. Returned as a Unix timestamp in milliseconds.
	UnderscoreCreatedAt int64 `json:"_createdAt"`
	// Date and time when the task was last updated. Returned as a Unix timestamp in milliseconds.
	UnderscoreUpdatedAt *int64 `json:"_updatedAt,omitempty"`
	// Title of the task.
	Title string `json:"title"`
	// Group name used to organize the task within the case.
	Group string `json:"group"`
	// Details of the task.
	Description *string `json:"description,omitempty"`
	// Current status of the task. One of `Waiting`, `InProgress`, `Completed`,or `Cancel`.
	Status string `json:"status"`
	// Whether the task is flagged as important.
	Flag bool `json:"flag"`
	// Date and time when the task started. Returned as a Unix timestamp in milliseconds.
	StartDate *int64 `json:"startDate,omitempty"`
	// Date and time when the task ended. Returned as a Unix timestamp in milliseconds.
	EndDate *int64 `json:"endDate,omitempty"`
	// Login of the user assigned to the task.
	Assignee *string `json:"assignee,omitempty"`
	// Display order of the task within its group.
	Order int32 `json:"order"`
	// Due date for the task. Returned as a Unix timestamp in milliseconds.
	DueDate *int64 `json:"dueDate,omitempty"`
	// Whether the task must be completed before the case can be closed.
	Mandatory bool                   `json:"mandatory"`
	ExtraData map[string]interface{} `json:"extraData"`
	// Duration of the task from start to completion, in milliseconds.
	TimeToHandle *int64 `json:"timeToHandle,omitempty"`
}

type _OutputTask OutputTask

// NewOutputTask instantiates a new OutputTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputTask(underscoreId string, underscoreType string, underscoreCreatedBy string, underscoreCreatedAt int64, title string, group string, status string, flag bool, order int32, mandatory bool, extraData map[string]interface{}) *OutputTask {
	this := OutputTask{}
	this.UnderscoreId = underscoreId
	this.UnderscoreType = underscoreType
	this.UnderscoreCreatedBy = underscoreCreatedBy
	this.UnderscoreCreatedAt = underscoreCreatedAt
	this.Title = title
	this.Group = group
	this.Status = status
	this.Flag = flag
	this.Order = order
	this.Mandatory = mandatory
	this.ExtraData = extraData
	return &this
}

// NewOutputTaskWithDefaults instantiates a new OutputTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputTaskWithDefaults() *OutputTask {
	this := OutputTask{}
	return &this
}

// GetUnderscoreId returns the UnderscoreId field value
func (o *OutputTask) GetUnderscoreId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnderscoreId
}

// GetUnderscoreIdOk returns a tuple with the UnderscoreId field value
// and a boolean to check if the value has been set.
func (o *OutputTask) GetUnderscoreIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreId, true
}

// SetUnderscoreId sets field value
func (o *OutputTask) SetUnderscoreId(v string) {
	o.UnderscoreId = v
}

// GetUnderscoreType returns the UnderscoreType field value
func (o *OutputTask) GetUnderscoreType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnderscoreType
}

// GetUnderscoreTypeOk returns a tuple with the UnderscoreType field value
// and a boolean to check if the value has been set.
func (o *OutputTask) GetUnderscoreTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreType, true
}

// SetUnderscoreType sets field value
func (o *OutputTask) SetUnderscoreType(v string) {
	o.UnderscoreType = v
}

// GetUnderscoreCreatedBy returns the UnderscoreCreatedBy field value
func (o *OutputTask) GetUnderscoreCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnderscoreCreatedBy
}

// GetUnderscoreCreatedByOk returns a tuple with the UnderscoreCreatedBy field value
// and a boolean to check if the value has been set.
func (o *OutputTask) GetUnderscoreCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreCreatedBy, true
}

// SetUnderscoreCreatedBy sets field value
func (o *OutputTask) SetUnderscoreCreatedBy(v string) {
	o.UnderscoreCreatedBy = v
}

// GetUnderscoreUpdatedBy returns the UnderscoreUpdatedBy field value if set, zero value otherwise.
func (o *OutputTask) GetUnderscoreUpdatedBy() string {
	if o == nil || IsNil(o.UnderscoreUpdatedBy) {
		var ret string
		return ret
	}
	return *o.UnderscoreUpdatedBy
}

// GetUnderscoreUpdatedByOk returns a tuple with the UnderscoreUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputTask) GetUnderscoreUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UnderscoreUpdatedBy) {
		return nil, false
	}
	return o.UnderscoreUpdatedBy, true
}

// HasUnderscoreUpdatedBy returns a boolean if a field has been set.
func (o *OutputTask) HasUnderscoreUpdatedBy() bool {
	if o != nil && !IsNil(o.UnderscoreUpdatedBy) {
		return true
	}

	return false
}

// SetUnderscoreUpdatedBy gets a reference to the given string and assigns it to the UnderscoreUpdatedBy field.
func (o *OutputTask) SetUnderscoreUpdatedBy(v string) {
	o.UnderscoreUpdatedBy = &v
}

// GetUnderscoreCreatedAt returns the UnderscoreCreatedAt field value
func (o *OutputTask) GetUnderscoreCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UnderscoreCreatedAt
}

// GetUnderscoreCreatedAtOk returns a tuple with the UnderscoreCreatedAt field value
// and a boolean to check if the value has been set.
func (o *OutputTask) GetUnderscoreCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreCreatedAt, true
}

// SetUnderscoreCreatedAt sets field value
func (o *OutputTask) SetUnderscoreCreatedAt(v int64) {
	o.UnderscoreCreatedAt = v
}

// GetUnderscoreUpdatedAt returns the UnderscoreUpdatedAt field value if set, zero value otherwise.
func (o *OutputTask) GetUnderscoreUpdatedAt() int64 {
	if o == nil || IsNil(o.UnderscoreUpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UnderscoreUpdatedAt
}

// GetUnderscoreUpdatedAtOk returns a tuple with the UnderscoreUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputTask) GetUnderscoreUpdatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.UnderscoreUpdatedAt) {
		return nil, false
	}
	return o.UnderscoreUpdatedAt, true
}

// HasUnderscoreUpdatedAt returns a boolean if a field has been set.
func (o *OutputTask) HasUnderscoreUpdatedAt() bool {
	if o != nil && !IsNil(o.UnderscoreUpdatedAt) {
		return true
	}

	return false
}

// SetUnderscoreUpdatedAt gets a reference to the given int64 and assigns it to the UnderscoreUpdatedAt field.
func (o *OutputTask) SetUnderscoreUpdatedAt(v int64) {
	o.UnderscoreUpdatedAt = &v
}

// GetTitle returns the Title field value
func (o *OutputTask) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *OutputTask) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *OutputTask) SetTitle(v string) {
	o.Title = v
}

// GetGroup returns the Group field value
func (o *OutputTask) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *OutputTask) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *OutputTask) SetGroup(v string) {
	o.Group = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OutputTask) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputTask) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OutputTask) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OutputTask) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value
func (o *OutputTask) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OutputTask) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *OutputTask) SetStatus(v string) {
	o.Status = v
}

// GetFlag returns the Flag field value
func (o *OutputTask) GetFlag() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Flag
}

// GetFlagOk returns a tuple with the Flag field value
// and a boolean to check if the value has been set.
func (o *OutputTask) GetFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Flag, true
}

// SetFlag sets field value
func (o *OutputTask) SetFlag(v bool) {
	o.Flag = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *OutputTask) GetStartDate() int64 {
	if o == nil || IsNil(o.StartDate) {
		var ret int64
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputTask) GetStartDateOk() (*int64, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *OutputTask) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given int64 and assigns it to the StartDate field.
func (o *OutputTask) SetStartDate(v int64) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *OutputTask) GetEndDate() int64 {
	if o == nil || IsNil(o.EndDate) {
		var ret int64
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputTask) GetEndDateOk() (*int64, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *OutputTask) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given int64 and assigns it to the EndDate field.
func (o *OutputTask) SetEndDate(v int64) {
	o.EndDate = &v
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *OutputTask) GetAssignee() string {
	if o == nil || IsNil(o.Assignee) {
		var ret string
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputTask) GetAssigneeOk() (*string, bool) {
	if o == nil || IsNil(o.Assignee) {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *OutputTask) HasAssignee() bool {
	if o != nil && !IsNil(o.Assignee) {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given string and assigns it to the Assignee field.
func (o *OutputTask) SetAssignee(v string) {
	o.Assignee = &v
}

// GetOrder returns the Order field value
func (o *OutputTask) GetOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *OutputTask) GetOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *OutputTask) SetOrder(v int32) {
	o.Order = v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *OutputTask) GetDueDate() int64 {
	if o == nil || IsNil(o.DueDate) {
		var ret int64
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputTask) GetDueDateOk() (*int64, bool) {
	if o == nil || IsNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *OutputTask) HasDueDate() bool {
	if o != nil && !IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given int64 and assigns it to the DueDate field.
func (o *OutputTask) SetDueDate(v int64) {
	o.DueDate = &v
}

// GetMandatory returns the Mandatory field value
func (o *OutputTask) GetMandatory() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Mandatory
}

// GetMandatoryOk returns a tuple with the Mandatory field value
// and a boolean to check if the value has been set.
func (o *OutputTask) GetMandatoryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mandatory, true
}

// SetMandatory sets field value
func (o *OutputTask) SetMandatory(v bool) {
	o.Mandatory = v
}

// GetExtraData returns the ExtraData field value
func (o *OutputTask) GetExtraData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ExtraData
}

// GetExtraDataOk returns a tuple with the ExtraData field value
// and a boolean to check if the value has been set.
func (o *OutputTask) GetExtraDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.ExtraData, true
}

// SetExtraData sets field value
func (o *OutputTask) SetExtraData(v map[string]interface{}) {
	o.ExtraData = v
}

// GetTimeToHandle returns the TimeToHandle field value if set, zero value otherwise.
func (o *OutputTask) GetTimeToHandle() int64 {
	if o == nil || IsNil(o.TimeToHandle) {
		var ret int64
		return ret
	}
	return *o.TimeToHandle
}

// GetTimeToHandleOk returns a tuple with the TimeToHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputTask) GetTimeToHandleOk() (*int64, bool) {
	if o == nil || IsNil(o.TimeToHandle) {
		return nil, false
	}
	return o.TimeToHandle, true
}

// HasTimeToHandle returns a boolean if a field has been set.
func (o *OutputTask) HasTimeToHandle() bool {
	if o != nil && !IsNil(o.TimeToHandle) {
		return true
	}

	return false
}

// SetTimeToHandle gets a reference to the given int64 and assigns it to the TimeToHandle field.
func (o *OutputTask) SetTimeToHandle(v int64) {
	o.TimeToHandle = &v
}

func (o OutputTask) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutputTask) ToMap() (map[string]interface{}, error) {
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
	toSerialize["title"] = o.Title
	toSerialize["group"] = o.Group
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["status"] = o.Status
	toSerialize["flag"] = o.Flag
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.Assignee) {
		toSerialize["assignee"] = o.Assignee
	}
	toSerialize["order"] = o.Order
	if !IsNil(o.DueDate) {
		toSerialize["dueDate"] = o.DueDate
	}
	toSerialize["mandatory"] = o.Mandatory
	toSerialize["extraData"] = o.ExtraData
	if !IsNil(o.TimeToHandle) {
		toSerialize["timeToHandle"] = o.TimeToHandle
	}
	return toSerialize, nil
}

func (o *OutputTask) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"_type",
		"_createdBy",
		"_createdAt",
		"title",
		"group",
		"status",
		"flag",
		"order",
		"mandatory",
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

	varOutputTask := _OutputTask{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// Postprocessed by scripts/tolerant-response-decoder: unknown fields are ignored,
	// since TheHive adds fields to its responses in minor versions.
	err = decoder.Decode(&varOutputTask)

	if err != nil {
		return err
	}

	*o = OutputTask(varOutputTask)

	return err
}

type NullableOutputTask struct {
	value *OutputTask
	isSet bool
}

func (v NullableOutputTask) Get() *OutputTask {
	return v.value
}

func (v *NullableOutputTask) Set(val *OutputTask) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputTask) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputTask(val *OutputTask) *NullableOutputTask {
	return &NullableOutputTask{value: val, isSet: true}
}

func (v NullableOutputTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
