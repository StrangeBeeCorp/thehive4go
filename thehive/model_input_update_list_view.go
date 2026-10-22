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

// checks if the InputUpdateListView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputUpdateListView{}

// InputUpdateListView struct for InputUpdateListView
type InputUpdateListView struct {
	// New name for the view. Leave unset to keep the current name.
	Name *string `json:"name,omitempty"`
	// Filter applied to the list. Uses the same operators as the [`filter` step of the Query API](https://docs.strangebee.com/thehive/api-docs/#tag/Query-and-Export). Leave unset to keep the current filter.
	Filter map[string]interface{} `json:"filter,omitempty"`
	// New display options for the view. Leave unset to keep the current options.
	ListOptions *ListOptions `json:"listOptions,omitempty"`
	// New sort order for the view. Each entry is a field name prefixed with `+` for ascending order or `-` for descending order. Leave unset to keep the current sort order.
	SortList []string `json:"sortList,omitempty"`
	// New set of optional columns displayed in the list, on top of the fixed set of columns shown for that entity type. The available optional columns vary by entity: check the corresponding list in TheHive web interface. Leave unset to keep the current columns.
	ShowColumns []string `json:"showColumns,omitempty"`
	// Whether to share the view with your entire organization, or make it private. Leave unset to keep the current visibility.  Every member of the organization can use, modify, and delete a shared view.
	IsShared *bool `json:"isShared,omitempty"`
}

// NewInputUpdateListView instantiates a new InputUpdateListView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputUpdateListView() *InputUpdateListView {
	this := InputUpdateListView{}
	return &this
}

// NewInputUpdateListViewWithDefaults instantiates a new InputUpdateListView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputUpdateListViewWithDefaults() *InputUpdateListView {
	this := InputUpdateListView{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InputUpdateListView) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateListView) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InputUpdateListView) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InputUpdateListView) SetName(v string) {
	o.Name = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *InputUpdateListView) GetFilter() map[string]interface{} {
	if o == nil || IsNil(o.Filter) {
		var ret map[string]interface{}
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateListView) GetFilterOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Filter) {
		return map[string]interface{}{}, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *InputUpdateListView) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given map[string]interface{} and assigns it to the Filter field.
func (o *InputUpdateListView) SetFilter(v map[string]interface{}) {
	o.Filter = v
}

// GetListOptions returns the ListOptions field value if set, zero value otherwise.
func (o *InputUpdateListView) GetListOptions() ListOptions {
	if o == nil || IsNil(o.ListOptions) {
		var ret ListOptions
		return ret
	}
	return *o.ListOptions
}

// GetListOptionsOk returns a tuple with the ListOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateListView) GetListOptionsOk() (*ListOptions, bool) {
	if o == nil || IsNil(o.ListOptions) {
		return nil, false
	}
	return o.ListOptions, true
}

// HasListOptions returns a boolean if a field has been set.
func (o *InputUpdateListView) HasListOptions() bool {
	if o != nil && !IsNil(o.ListOptions) {
		return true
	}

	return false
}

// SetListOptions gets a reference to the given ListOptions and assigns it to the ListOptions field.
func (o *InputUpdateListView) SetListOptions(v ListOptions) {
	o.ListOptions = &v
}

// GetSortList returns the SortList field value if set, zero value otherwise.
func (o *InputUpdateListView) GetSortList() []string {
	if o == nil || IsNil(o.SortList) {
		var ret []string
		return ret
	}
	return o.SortList
}

// GetSortListOk returns a tuple with the SortList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateListView) GetSortListOk() ([]string, bool) {
	if o == nil || IsNil(o.SortList) {
		return nil, false
	}
	return o.SortList, true
}

// HasSortList returns a boolean if a field has been set.
func (o *InputUpdateListView) HasSortList() bool {
	if o != nil && !IsNil(o.SortList) {
		return true
	}

	return false
}

// SetSortList gets a reference to the given []string and assigns it to the SortList field.
func (o *InputUpdateListView) SetSortList(v []string) {
	o.SortList = v
}

// GetShowColumns returns the ShowColumns field value if set, zero value otherwise.
func (o *InputUpdateListView) GetShowColumns() []string {
	if o == nil || IsNil(o.ShowColumns) {
		var ret []string
		return ret
	}
	return o.ShowColumns
}

// GetShowColumnsOk returns a tuple with the ShowColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateListView) GetShowColumnsOk() ([]string, bool) {
	if o == nil || IsNil(o.ShowColumns) {
		return nil, false
	}
	return o.ShowColumns, true
}

// HasShowColumns returns a boolean if a field has been set.
func (o *InputUpdateListView) HasShowColumns() bool {
	if o != nil && !IsNil(o.ShowColumns) {
		return true
	}

	return false
}

// SetShowColumns gets a reference to the given []string and assigns it to the ShowColumns field.
func (o *InputUpdateListView) SetShowColumns(v []string) {
	o.ShowColumns = v
}

// GetIsShared returns the IsShared field value if set, zero value otherwise.
func (o *InputUpdateListView) GetIsShared() bool {
	if o == nil || IsNil(o.IsShared) {
		var ret bool
		return ret
	}
	return *o.IsShared
}

// GetIsSharedOk returns a tuple with the IsShared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateListView) GetIsSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsShared) {
		return nil, false
	}
	return o.IsShared, true
}

// HasIsShared returns a boolean if a field has been set.
func (o *InputUpdateListView) HasIsShared() bool {
	if o != nil && !IsNil(o.IsShared) {
		return true
	}

	return false
}

// SetIsShared gets a reference to the given bool and assigns it to the IsShared field.
func (o *InputUpdateListView) SetIsShared(v bool) {
	o.IsShared = &v
}

func (o InputUpdateListView) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputUpdateListView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.ListOptions) {
		toSerialize["listOptions"] = o.ListOptions
	}
	if !IsNil(o.SortList) {
		toSerialize["sortList"] = o.SortList
	}
	if !IsNil(o.ShowColumns) {
		toSerialize["showColumns"] = o.ShowColumns
	}
	if !IsNil(o.IsShared) {
		toSerialize["isShared"] = o.IsShared
	}
	return toSerialize, nil
}

type NullableInputUpdateListView struct {
	value *InputUpdateListView
	isSet bool
}

func (v NullableInputUpdateListView) Get() *InputUpdateListView {
	return v.value
}

func (v *NullableInputUpdateListView) Set(val *InputUpdateListView) {
	v.value = val
	v.isSet = true
}

func (v NullableInputUpdateListView) IsSet() bool {
	return v.isSet
}

func (v *NullableInputUpdateListView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputUpdateListView(val *InputUpdateListView) *NullableInputUpdateListView {
	return &NullableInputUpdateListView{value: val, isSet: true}
}

func (v NullableInputUpdateListView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputUpdateListView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
