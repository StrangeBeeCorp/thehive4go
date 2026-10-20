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

// checks if the OutputListView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutputListView{}

// OutputListView struct for OutputListView
type OutputListView struct {
	// Internal unique identifier of the view.
	UnderscoreId string `json:"_id"`
	// Entity type, always `ListView`.
	UnderscoreType string `json:"_type"`
	// Date and time when the view was created. Returned as a Unix timestamp in milliseconds.
	UnderscoreCreatedAt int64 `json:"_createdAt"`
	// Login of the user who created the view.
	UnderscoreCreatedBy string `json:"_createdBy"`
	// Date and time when the view was last updated. Returned as a Unix timestamp in milliseconds.
	UnderscoreUpdatedAt *int64 `json:"_updatedAt,omitempty"`
	// Login of the user who last updated the view.
	UnderscoreUpdatedBy *string `json:"_updatedBy,omitempty"`
	// Name of the view.
	Name string `json:"name"`
	// Type of list the view applies to. Determines which list the view's filters, sort order, displayed columns, and display options apply to.
	Entity string `json:"entity"`
	// Filter applied to the list. Uses the same operators as the `filter` step of the Query API.
	Filter map[string]interface{} `json:"filter"`
	// Display options saved for the view.
	ListOptions ListOptions `json:"listOptions"`
	// Sort order saved for the view. Each entry is a field name prefixed with `+` for ascending order or `-` for descending order.
	SortList []string `json:"sortList,omitempty"`
	// Optional columns displayed in the list, on top of the fixed set of columns shown for that entity type.
	ShowColumns []string `json:"showColumns,omitempty"`
	// Whether the view is shared with your entire organization. When false, the view is private and visible only to you.
	IsShared bool `json:"isShared"`
}

type _OutputListView OutputListView

// NewOutputListView instantiates a new OutputListView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputListView(underscoreId string, underscoreType string, underscoreCreatedAt int64, underscoreCreatedBy string, name string, entity string, filter map[string]interface{}, listOptions ListOptions, isShared bool) *OutputListView {
	this := OutputListView{}
	this.UnderscoreId = underscoreId
	this.UnderscoreType = underscoreType
	this.UnderscoreCreatedAt = underscoreCreatedAt
	this.UnderscoreCreatedBy = underscoreCreatedBy
	this.Name = name
	this.Entity = entity
	this.Filter = filter
	this.ListOptions = listOptions
	this.IsShared = isShared
	return &this
}

// NewOutputListViewWithDefaults instantiates a new OutputListView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputListViewWithDefaults() *OutputListView {
	this := OutputListView{}
	return &this
}

// GetUnderscoreId returns the UnderscoreId field value
func (o *OutputListView) GetUnderscoreId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnderscoreId
}

// GetUnderscoreIdOk returns a tuple with the UnderscoreId field value
// and a boolean to check if the value has been set.
func (o *OutputListView) GetUnderscoreIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreId, true
}

// SetUnderscoreId sets field value
func (o *OutputListView) SetUnderscoreId(v string) {
	o.UnderscoreId = v
}

// GetUnderscoreType returns the UnderscoreType field value
func (o *OutputListView) GetUnderscoreType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnderscoreType
}

// GetUnderscoreTypeOk returns a tuple with the UnderscoreType field value
// and a boolean to check if the value has been set.
func (o *OutputListView) GetUnderscoreTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreType, true
}

// SetUnderscoreType sets field value
func (o *OutputListView) SetUnderscoreType(v string) {
	o.UnderscoreType = v
}

// GetUnderscoreCreatedAt returns the UnderscoreCreatedAt field value
func (o *OutputListView) GetUnderscoreCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UnderscoreCreatedAt
}

// GetUnderscoreCreatedAtOk returns a tuple with the UnderscoreCreatedAt field value
// and a boolean to check if the value has been set.
func (o *OutputListView) GetUnderscoreCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreCreatedAt, true
}

// SetUnderscoreCreatedAt sets field value
func (o *OutputListView) SetUnderscoreCreatedAt(v int64) {
	o.UnderscoreCreatedAt = v
}

// GetUnderscoreCreatedBy returns the UnderscoreCreatedBy field value
func (o *OutputListView) GetUnderscoreCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnderscoreCreatedBy
}

// GetUnderscoreCreatedByOk returns a tuple with the UnderscoreCreatedBy field value
// and a boolean to check if the value has been set.
func (o *OutputListView) GetUnderscoreCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreCreatedBy, true
}

// SetUnderscoreCreatedBy sets field value
func (o *OutputListView) SetUnderscoreCreatedBy(v string) {
	o.UnderscoreCreatedBy = v
}

// GetUnderscoreUpdatedAt returns the UnderscoreUpdatedAt field value if set, zero value otherwise.
func (o *OutputListView) GetUnderscoreUpdatedAt() int64 {
	if o == nil || IsNil(o.UnderscoreUpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UnderscoreUpdatedAt
}

// GetUnderscoreUpdatedAtOk returns a tuple with the UnderscoreUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputListView) GetUnderscoreUpdatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.UnderscoreUpdatedAt) {
		return nil, false
	}
	return o.UnderscoreUpdatedAt, true
}

// HasUnderscoreUpdatedAt returns a boolean if a field has been set.
func (o *OutputListView) HasUnderscoreUpdatedAt() bool {
	if o != nil && !IsNil(o.UnderscoreUpdatedAt) {
		return true
	}

	return false
}

// SetUnderscoreUpdatedAt gets a reference to the given int64 and assigns it to the UnderscoreUpdatedAt field.
func (o *OutputListView) SetUnderscoreUpdatedAt(v int64) {
	o.UnderscoreUpdatedAt = &v
}

// GetUnderscoreUpdatedBy returns the UnderscoreUpdatedBy field value if set, zero value otherwise.
func (o *OutputListView) GetUnderscoreUpdatedBy() string {
	if o == nil || IsNil(o.UnderscoreUpdatedBy) {
		var ret string
		return ret
	}
	return *o.UnderscoreUpdatedBy
}

// GetUnderscoreUpdatedByOk returns a tuple with the UnderscoreUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputListView) GetUnderscoreUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UnderscoreUpdatedBy) {
		return nil, false
	}
	return o.UnderscoreUpdatedBy, true
}

// HasUnderscoreUpdatedBy returns a boolean if a field has been set.
func (o *OutputListView) HasUnderscoreUpdatedBy() bool {
	if o != nil && !IsNil(o.UnderscoreUpdatedBy) {
		return true
	}

	return false
}

// SetUnderscoreUpdatedBy gets a reference to the given string and assigns it to the UnderscoreUpdatedBy field.
func (o *OutputListView) SetUnderscoreUpdatedBy(v string) {
	o.UnderscoreUpdatedBy = &v
}

// GetName returns the Name field value
func (o *OutputListView) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OutputListView) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OutputListView) SetName(v string) {
	o.Name = v
}

// GetEntity returns the Entity field value
func (o *OutputListView) GetEntity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Entity
}

// GetEntityOk returns a tuple with the Entity field value
// and a boolean to check if the value has been set.
func (o *OutputListView) GetEntityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entity, true
}

// SetEntity sets field value
func (o *OutputListView) SetEntity(v string) {
	o.Entity = v
}

// GetFilter returns the Filter field value
func (o *OutputListView) GetFilter() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *OutputListView) GetFilterOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Filter, true
}

// SetFilter sets field value
func (o *OutputListView) SetFilter(v map[string]interface{}) {
	o.Filter = v
}

// GetListOptions returns the ListOptions field value
func (o *OutputListView) GetListOptions() ListOptions {
	if o == nil {
		var ret ListOptions
		return ret
	}

	return o.ListOptions
}

// GetListOptionsOk returns a tuple with the ListOptions field value
// and a boolean to check if the value has been set.
func (o *OutputListView) GetListOptionsOk() (*ListOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListOptions, true
}

// SetListOptions sets field value
func (o *OutputListView) SetListOptions(v ListOptions) {
	o.ListOptions = v
}

// GetSortList returns the SortList field value if set, zero value otherwise.
func (o *OutputListView) GetSortList() []string {
	if o == nil || IsNil(o.SortList) {
		var ret []string
		return ret
	}
	return o.SortList
}

// GetSortListOk returns a tuple with the SortList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputListView) GetSortListOk() ([]string, bool) {
	if o == nil || IsNil(o.SortList) {
		return nil, false
	}
	return o.SortList, true
}

// HasSortList returns a boolean if a field has been set.
func (o *OutputListView) HasSortList() bool {
	if o != nil && !IsNil(o.SortList) {
		return true
	}

	return false
}

// SetSortList gets a reference to the given []string and assigns it to the SortList field.
func (o *OutputListView) SetSortList(v []string) {
	o.SortList = v
}

// GetShowColumns returns the ShowColumns field value if set, zero value otherwise.
func (o *OutputListView) GetShowColumns() []string {
	if o == nil || IsNil(o.ShowColumns) {
		var ret []string
		return ret
	}
	return o.ShowColumns
}

// GetShowColumnsOk returns a tuple with the ShowColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputListView) GetShowColumnsOk() ([]string, bool) {
	if o == nil || IsNil(o.ShowColumns) {
		return nil, false
	}
	return o.ShowColumns, true
}

// HasShowColumns returns a boolean if a field has been set.
func (o *OutputListView) HasShowColumns() bool {
	if o != nil && !IsNil(o.ShowColumns) {
		return true
	}

	return false
}

// SetShowColumns gets a reference to the given []string and assigns it to the ShowColumns field.
func (o *OutputListView) SetShowColumns(v []string) {
	o.ShowColumns = v
}

// GetIsShared returns the IsShared field value
func (o *OutputListView) GetIsShared() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsShared
}

// GetIsSharedOk returns a tuple with the IsShared field value
// and a boolean to check if the value has been set.
func (o *OutputListView) GetIsSharedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsShared, true
}

// SetIsShared sets field value
func (o *OutputListView) SetIsShared(v bool) {
	o.IsShared = v
}

func (o OutputListView) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutputListView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.UnderscoreId
	toSerialize["_type"] = o.UnderscoreType
	toSerialize["_createdAt"] = o.UnderscoreCreatedAt
	toSerialize["_createdBy"] = o.UnderscoreCreatedBy
	if !IsNil(o.UnderscoreUpdatedAt) {
		toSerialize["_updatedAt"] = o.UnderscoreUpdatedAt
	}
	if !IsNil(o.UnderscoreUpdatedBy) {
		toSerialize["_updatedBy"] = o.UnderscoreUpdatedBy
	}
	toSerialize["name"] = o.Name
	toSerialize["entity"] = o.Entity
	toSerialize["filter"] = o.Filter
	toSerialize["listOptions"] = o.ListOptions
	if !IsNil(o.SortList) {
		toSerialize["sortList"] = o.SortList
	}
	if !IsNil(o.ShowColumns) {
		toSerialize["showColumns"] = o.ShowColumns
	}
	toSerialize["isShared"] = o.IsShared
	return toSerialize, nil
}

func (o *OutputListView) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"_type",
		"_createdAt",
		"_createdBy",
		"name",
		"entity",
		"filter",
		"listOptions",
		"isShared",
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

	varOutputListView := _OutputListView{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// Postprocessed by scripts/tolerant-response-decoder: unknown fields are ignored,
	// since TheHive adds fields to its responses in minor versions.
	err = decoder.Decode(&varOutputListView)

	if err != nil {
		return err
	}

	*o = OutputListView(varOutputListView)

	return err
}

type NullableOutputListView struct {
	value *OutputListView
	isSet bool
}

func (v NullableOutputListView) Get() *OutputListView {
	return v.value
}

func (v *NullableOutputListView) Set(val *OutputListView) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputListView) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputListView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputListView(val *OutputListView) *NullableOutputListView {
	return &NullableOutputListView{value: val, isSet: true}
}

func (v NullableOutputListView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputListView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
