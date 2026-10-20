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

// checks if the OutputCaseReportTemplateObjects type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutputCaseReportTemplateObjects{}

// OutputCaseReportTemplateObjects struct for OutputCaseReportTemplateObjects
type OutputCaseReportTemplateObjects struct {
	// Available widget types for use in case report templates.
	Widgets []Kind `json:"widgets,omitempty"`
	// Available alert fields for use in alert widgets.
	AlertFields []AlertField `json:"alertFields,omitempty"`
	// Available observable fields for use in observable widgets.
	ObservableFields []ObservableField `json:"observableFields,omitempty"`
	// Available task fields for use in task widgets.
	TaskFields []TaskField `json:"taskFields,omitempty"`
	// Available TTP fields for use in TTP widgets.
	TtpFields []TTPField `json:"ttpFields,omitempty"`
	// Available task log fields for use in task list widgets.
	LogFields []LogField `json:"logFields,omitempty"`
	// Available event types for use in timeline widgets.
	TimelineEvents []TimelineEvents `json:"timelineEvents,omitempty"`
}

// NewOutputCaseReportTemplateObjects instantiates a new OutputCaseReportTemplateObjects object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputCaseReportTemplateObjects() *OutputCaseReportTemplateObjects {
	this := OutputCaseReportTemplateObjects{}
	return &this
}

// NewOutputCaseReportTemplateObjectsWithDefaults instantiates a new OutputCaseReportTemplateObjects object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputCaseReportTemplateObjectsWithDefaults() *OutputCaseReportTemplateObjects {
	this := OutputCaseReportTemplateObjects{}
	return &this
}

// GetWidgets returns the Widgets field value if set, zero value otherwise.
func (o *OutputCaseReportTemplateObjects) GetWidgets() []Kind {
	if o == nil || IsNil(o.Widgets) {
		var ret []Kind
		return ret
	}
	return o.Widgets
}

// GetWidgetsOk returns a tuple with the Widgets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCaseReportTemplateObjects) GetWidgetsOk() ([]Kind, bool) {
	if o == nil || IsNil(o.Widgets) {
		return nil, false
	}
	return o.Widgets, true
}

// HasWidgets returns a boolean if a field has been set.
func (o *OutputCaseReportTemplateObjects) HasWidgets() bool {
	if o != nil && !IsNil(o.Widgets) {
		return true
	}

	return false
}

// SetWidgets gets a reference to the given []Kind and assigns it to the Widgets field.
func (o *OutputCaseReportTemplateObjects) SetWidgets(v []Kind) {
	o.Widgets = v
}

// GetAlertFields returns the AlertFields field value if set, zero value otherwise.
func (o *OutputCaseReportTemplateObjects) GetAlertFields() []AlertField {
	if o == nil || IsNil(o.AlertFields) {
		var ret []AlertField
		return ret
	}
	return o.AlertFields
}

// GetAlertFieldsOk returns a tuple with the AlertFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCaseReportTemplateObjects) GetAlertFieldsOk() ([]AlertField, bool) {
	if o == nil || IsNil(o.AlertFields) {
		return nil, false
	}
	return o.AlertFields, true
}

// HasAlertFields returns a boolean if a field has been set.
func (o *OutputCaseReportTemplateObjects) HasAlertFields() bool {
	if o != nil && !IsNil(o.AlertFields) {
		return true
	}

	return false
}

// SetAlertFields gets a reference to the given []AlertField and assigns it to the AlertFields field.
func (o *OutputCaseReportTemplateObjects) SetAlertFields(v []AlertField) {
	o.AlertFields = v
}

// GetObservableFields returns the ObservableFields field value if set, zero value otherwise.
func (o *OutputCaseReportTemplateObjects) GetObservableFields() []ObservableField {
	if o == nil || IsNil(o.ObservableFields) {
		var ret []ObservableField
		return ret
	}
	return o.ObservableFields
}

// GetObservableFieldsOk returns a tuple with the ObservableFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCaseReportTemplateObjects) GetObservableFieldsOk() ([]ObservableField, bool) {
	if o == nil || IsNil(o.ObservableFields) {
		return nil, false
	}
	return o.ObservableFields, true
}

// HasObservableFields returns a boolean if a field has been set.
func (o *OutputCaseReportTemplateObjects) HasObservableFields() bool {
	if o != nil && !IsNil(o.ObservableFields) {
		return true
	}

	return false
}

// SetObservableFields gets a reference to the given []ObservableField and assigns it to the ObservableFields field.
func (o *OutputCaseReportTemplateObjects) SetObservableFields(v []ObservableField) {
	o.ObservableFields = v
}

// GetTaskFields returns the TaskFields field value if set, zero value otherwise.
func (o *OutputCaseReportTemplateObjects) GetTaskFields() []TaskField {
	if o == nil || IsNil(o.TaskFields) {
		var ret []TaskField
		return ret
	}
	return o.TaskFields
}

// GetTaskFieldsOk returns a tuple with the TaskFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCaseReportTemplateObjects) GetTaskFieldsOk() ([]TaskField, bool) {
	if o == nil || IsNil(o.TaskFields) {
		return nil, false
	}
	return o.TaskFields, true
}

// HasTaskFields returns a boolean if a field has been set.
func (o *OutputCaseReportTemplateObjects) HasTaskFields() bool {
	if o != nil && !IsNil(o.TaskFields) {
		return true
	}

	return false
}

// SetTaskFields gets a reference to the given []TaskField and assigns it to the TaskFields field.
func (o *OutputCaseReportTemplateObjects) SetTaskFields(v []TaskField) {
	o.TaskFields = v
}

// GetTtpFields returns the TtpFields field value if set, zero value otherwise.
func (o *OutputCaseReportTemplateObjects) GetTtpFields() []TTPField {
	if o == nil || IsNil(o.TtpFields) {
		var ret []TTPField
		return ret
	}
	return o.TtpFields
}

// GetTtpFieldsOk returns a tuple with the TtpFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCaseReportTemplateObjects) GetTtpFieldsOk() ([]TTPField, bool) {
	if o == nil || IsNil(o.TtpFields) {
		return nil, false
	}
	return o.TtpFields, true
}

// HasTtpFields returns a boolean if a field has been set.
func (o *OutputCaseReportTemplateObjects) HasTtpFields() bool {
	if o != nil && !IsNil(o.TtpFields) {
		return true
	}

	return false
}

// SetTtpFields gets a reference to the given []TTPField and assigns it to the TtpFields field.
func (o *OutputCaseReportTemplateObjects) SetTtpFields(v []TTPField) {
	o.TtpFields = v
}

// GetLogFields returns the LogFields field value if set, zero value otherwise.
func (o *OutputCaseReportTemplateObjects) GetLogFields() []LogField {
	if o == nil || IsNil(o.LogFields) {
		var ret []LogField
		return ret
	}
	return o.LogFields
}

// GetLogFieldsOk returns a tuple with the LogFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCaseReportTemplateObjects) GetLogFieldsOk() ([]LogField, bool) {
	if o == nil || IsNil(o.LogFields) {
		return nil, false
	}
	return o.LogFields, true
}

// HasLogFields returns a boolean if a field has been set.
func (o *OutputCaseReportTemplateObjects) HasLogFields() bool {
	if o != nil && !IsNil(o.LogFields) {
		return true
	}

	return false
}

// SetLogFields gets a reference to the given []LogField and assigns it to the LogFields field.
func (o *OutputCaseReportTemplateObjects) SetLogFields(v []LogField) {
	o.LogFields = v
}

// GetTimelineEvents returns the TimelineEvents field value if set, zero value otherwise.
func (o *OutputCaseReportTemplateObjects) GetTimelineEvents() []TimelineEvents {
	if o == nil || IsNil(o.TimelineEvents) {
		var ret []TimelineEvents
		return ret
	}
	return o.TimelineEvents
}

// GetTimelineEventsOk returns a tuple with the TimelineEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputCaseReportTemplateObjects) GetTimelineEventsOk() ([]TimelineEvents, bool) {
	if o == nil || IsNil(o.TimelineEvents) {
		return nil, false
	}
	return o.TimelineEvents, true
}

// HasTimelineEvents returns a boolean if a field has been set.
func (o *OutputCaseReportTemplateObjects) HasTimelineEvents() bool {
	if o != nil && !IsNil(o.TimelineEvents) {
		return true
	}

	return false
}

// SetTimelineEvents gets a reference to the given []TimelineEvents and assigns it to the TimelineEvents field.
func (o *OutputCaseReportTemplateObjects) SetTimelineEvents(v []TimelineEvents) {
	o.TimelineEvents = v
}

func (o OutputCaseReportTemplateObjects) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutputCaseReportTemplateObjects) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Widgets) {
		toSerialize["widgets"] = o.Widgets
	}
	if !IsNil(o.AlertFields) {
		toSerialize["alertFields"] = o.AlertFields
	}
	if !IsNil(o.ObservableFields) {
		toSerialize["observableFields"] = o.ObservableFields
	}
	if !IsNil(o.TaskFields) {
		toSerialize["taskFields"] = o.TaskFields
	}
	if !IsNil(o.TtpFields) {
		toSerialize["ttpFields"] = o.TtpFields
	}
	if !IsNil(o.LogFields) {
		toSerialize["logFields"] = o.LogFields
	}
	if !IsNil(o.TimelineEvents) {
		toSerialize["timelineEvents"] = o.TimelineEvents
	}
	return toSerialize, nil
}

type NullableOutputCaseReportTemplateObjects struct {
	value *OutputCaseReportTemplateObjects
	isSet bool
}

func (v NullableOutputCaseReportTemplateObjects) Get() *OutputCaseReportTemplateObjects {
	return v.value
}

func (v *NullableOutputCaseReportTemplateObjects) Set(val *OutputCaseReportTemplateObjects) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputCaseReportTemplateObjects) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputCaseReportTemplateObjects) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputCaseReportTemplateObjects(val *OutputCaseReportTemplateObjects) *NullableOutputCaseReportTemplateObjects {
	return &NullableOutputCaseReportTemplateObjects{value: val, isSet: true}
}

func (v NullableOutputCaseReportTemplateObjects) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputCaseReportTemplateObjects) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
