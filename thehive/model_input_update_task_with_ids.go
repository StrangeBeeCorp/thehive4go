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

// checks if the InputUpdateTaskWithIds type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputUpdateTaskWithIds{}

// InputUpdateTaskWithIds struct for InputUpdateTaskWithIds
type InputUpdateTaskWithIds struct {
	// Identifiers of the entities to target. IDs must be prefixed with `~`. Case numbers and names are passed as plain strings without a prefix.
	Ids []string `json:"ids"`
	// New title of the task.
	Title *string `json:"title,omitempty"`
	// New group of the task.
	Group *string `json:"group,omitempty"`
	// New details of the task. Supports [TheHive-flavored Markdown](https://docs.strangebee.com/thehive/user-guides/thehive-flavored-markdown/). Send `null` to clear this field.
	Description *string `json:"description,omitempty"`
	// Enum: `Waiting` `InProgress` `Completed` `Cancel`<br>New status of the task.
	Status *string `json:"status,omitempty"`
	// Marks the task as important for quick filtering.
	Flag *bool `json:"flag,omitempty"`
	// New start date of the task. Unix timestamp in milliseconds. Send `null` to clear this field.
	StartDate *int64 `json:"startDate,omitempty"`
	// New end date of the task. Unix timestamp in milliseconds.
	EndDate *int64 `json:"endDate,omitempty"`
	// New display order of the task within its group.
	Order *int32 `json:"order,omitempty"`
	// New due date of the task. Unix timestamp in milliseconds. Send `null` to clear this field.
	DueDate *int64 `json:"dueDate,omitempty"`
	// Login of the user to assign to the task. Typically an email address. If a default domain is configured in the [authentication settings](https://docs.strangebee.com/thehive/administration/authentication/configure-authentication), a plain username is also accepted. Use `POST /api/v1/query` with `listUser` to retrieve available users and their logins. Send `null` to unassign.
	Assignee *string `json:"assignee,omitempty"`
	// Whether the task must be completed before the case can be closed.
	Mandatory *bool `json:"mandatory,omitempty"`
}

type _InputUpdateTaskWithIds InputUpdateTaskWithIds

// NewInputUpdateTaskWithIds instantiates a new InputUpdateTaskWithIds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputUpdateTaskWithIds(ids []string) *InputUpdateTaskWithIds {
	this := InputUpdateTaskWithIds{}
	this.Ids = ids
	return &this
}

// NewInputUpdateTaskWithIdsWithDefaults instantiates a new InputUpdateTaskWithIds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputUpdateTaskWithIdsWithDefaults() *InputUpdateTaskWithIds {
	this := InputUpdateTaskWithIds{}
	return &this
}

// GetIds returns the Ids field value
func (o *InputUpdateTaskWithIds) GetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *InputUpdateTaskWithIds) GetIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ids, true
}

// SetIds sets field value
func (o *InputUpdateTaskWithIds) SetIds(v []string) {
	o.Ids = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *InputUpdateTaskWithIds) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateTaskWithIds) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *InputUpdateTaskWithIds) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *InputUpdateTaskWithIds) SetTitle(v string) {
	o.Title = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *InputUpdateTaskWithIds) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateTaskWithIds) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *InputUpdateTaskWithIds) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *InputUpdateTaskWithIds) SetGroup(v string) {
	o.Group = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InputUpdateTaskWithIds) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateTaskWithIds) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InputUpdateTaskWithIds) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InputUpdateTaskWithIds) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InputUpdateTaskWithIds) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateTaskWithIds) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InputUpdateTaskWithIds) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InputUpdateTaskWithIds) SetStatus(v string) {
	o.Status = &v
}

// GetFlag returns the Flag field value if set, zero value otherwise.
func (o *InputUpdateTaskWithIds) GetFlag() bool {
	if o == nil || IsNil(o.Flag) {
		var ret bool
		return ret
	}
	return *o.Flag
}

// GetFlagOk returns a tuple with the Flag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateTaskWithIds) GetFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.Flag) {
		return nil, false
	}
	return o.Flag, true
}

// HasFlag returns a boolean if a field has been set.
func (o *InputUpdateTaskWithIds) HasFlag() bool {
	if o != nil && !IsNil(o.Flag) {
		return true
	}

	return false
}

// SetFlag gets a reference to the given bool and assigns it to the Flag field.
func (o *InputUpdateTaskWithIds) SetFlag(v bool) {
	o.Flag = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *InputUpdateTaskWithIds) GetStartDate() int64 {
	if o == nil || IsNil(o.StartDate) {
		var ret int64
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateTaskWithIds) GetStartDateOk() (*int64, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *InputUpdateTaskWithIds) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given int64 and assigns it to the StartDate field.
func (o *InputUpdateTaskWithIds) SetStartDate(v int64) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *InputUpdateTaskWithIds) GetEndDate() int64 {
	if o == nil || IsNil(o.EndDate) {
		var ret int64
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateTaskWithIds) GetEndDateOk() (*int64, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *InputUpdateTaskWithIds) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given int64 and assigns it to the EndDate field.
func (o *InputUpdateTaskWithIds) SetEndDate(v int64) {
	o.EndDate = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *InputUpdateTaskWithIds) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateTaskWithIds) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *InputUpdateTaskWithIds) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *InputUpdateTaskWithIds) SetOrder(v int32) {
	o.Order = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *InputUpdateTaskWithIds) GetDueDate() int64 {
	if o == nil || IsNil(o.DueDate) {
		var ret int64
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateTaskWithIds) GetDueDateOk() (*int64, bool) {
	if o == nil || IsNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *InputUpdateTaskWithIds) HasDueDate() bool {
	if o != nil && !IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given int64 and assigns it to the DueDate field.
func (o *InputUpdateTaskWithIds) SetDueDate(v int64) {
	o.DueDate = &v
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *InputUpdateTaskWithIds) GetAssignee() string {
	if o == nil || IsNil(o.Assignee) {
		var ret string
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateTaskWithIds) GetAssigneeOk() (*string, bool) {
	if o == nil || IsNil(o.Assignee) {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *InputUpdateTaskWithIds) HasAssignee() bool {
	if o != nil && !IsNil(o.Assignee) {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given string and assigns it to the Assignee field.
func (o *InputUpdateTaskWithIds) SetAssignee(v string) {
	o.Assignee = &v
}

// GetMandatory returns the Mandatory field value if set, zero value otherwise.
func (o *InputUpdateTaskWithIds) GetMandatory() bool {
	if o == nil || IsNil(o.Mandatory) {
		var ret bool
		return ret
	}
	return *o.Mandatory
}

// GetMandatoryOk returns a tuple with the Mandatory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateTaskWithIds) GetMandatoryOk() (*bool, bool) {
	if o == nil || IsNil(o.Mandatory) {
		return nil, false
	}
	return o.Mandatory, true
}

// HasMandatory returns a boolean if a field has been set.
func (o *InputUpdateTaskWithIds) HasMandatory() bool {
	if o != nil && !IsNil(o.Mandatory) {
		return true
	}

	return false
}

// SetMandatory gets a reference to the given bool and assigns it to the Mandatory field.
func (o *InputUpdateTaskWithIds) SetMandatory(v bool) {
	o.Mandatory = &v
}

func (o InputUpdateTaskWithIds) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputUpdateTaskWithIds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ids"] = o.Ids
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Flag) {
		toSerialize["flag"] = o.Flag
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.DueDate) {
		toSerialize["dueDate"] = o.DueDate
	}
	if !IsNil(o.Assignee) {
		toSerialize["assignee"] = o.Assignee
	}
	if !IsNil(o.Mandatory) {
		toSerialize["mandatory"] = o.Mandatory
	}
	return toSerialize, nil
}

func (o *InputUpdateTaskWithIds) UnmarshalJSON(data []byte) (err error) {
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

	varInputUpdateTaskWithIds := _InputUpdateTaskWithIds{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInputUpdateTaskWithIds)

	if err != nil {
		return err
	}

	*o = InputUpdateTaskWithIds(varInputUpdateTaskWithIds)

	return err
}

type NullableInputUpdateTaskWithIds struct {
	value *InputUpdateTaskWithIds
	isSet bool
}

func (v NullableInputUpdateTaskWithIds) Get() *InputUpdateTaskWithIds {
	return v.value
}

func (v *NullableInputUpdateTaskWithIds) Set(val *InputUpdateTaskWithIds) {
	v.value = val
	v.isSet = true
}

func (v NullableInputUpdateTaskWithIds) IsSet() bool {
	return v.isSet
}

func (v *NullableInputUpdateTaskWithIds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputUpdateTaskWithIds(val *InputUpdateTaskWithIds) *NullableInputUpdateTaskWithIds {
	return &NullableInputUpdateTaskWithIds{value: val, isSet: true}
}

func (v NullableInputUpdateTaskWithIds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputUpdateTaskWithIds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
