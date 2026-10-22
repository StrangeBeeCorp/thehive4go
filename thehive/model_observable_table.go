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

// checks if the ObservableTable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObservableTable{}

// ObservableTable struct for ObservableTable
type ObservableTable struct {
	// Title displayed above the widget.
	Title *string `json:"title,omitempty"`
	// Observable fields to include as columns in the table.
	Columns []ObservableField `json:"columns,omitempty"`
	// Filter to apply to the observable list. See the [Query and Export](https://docs.strangebee.com/thehive/api-docs/#tag/Query-and-Export) section for the filter syntax.
	Filter *Filter `json:"filter,omitempty"`
	// Sort order for the observable list. See the [Query and Export](https://docs.strangebee.com/thehive/api-docs/#tag/Query-and-Export) section for the sort syntax.
	Sort []AlertListSortInner `json:"sort,omitempty"`
	// When `true`, masks the data of observables flagged as sensitive.
	ProtectData bool `json:"protectData"`
	// Maximum number of observables to include in the widget. If omitted, the global `maxElements` limit applies.
	MaxElements *int32 `json:"maxElements,omitempty"`
	Kind        Kind   `json:"_kind"`
}

type _ObservableTable ObservableTable

// NewObservableTable instantiates a new ObservableTable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObservableTable(protectData bool, kind Kind) *ObservableTable {
	this := ObservableTable{}
	this.ProtectData = protectData
	this.Kind = kind
	return &this
}

// NewObservableTableWithDefaults instantiates a new ObservableTable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObservableTableWithDefaults() *ObservableTable {
	this := ObservableTable{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ObservableTable) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservableTable) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ObservableTable) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ObservableTable) SetTitle(v string) {
	o.Title = &v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *ObservableTable) GetColumns() []ObservableField {
	if o == nil || IsNil(o.Columns) {
		var ret []ObservableField
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservableTable) GetColumnsOk() ([]ObservableField, bool) {
	if o == nil || IsNil(o.Columns) {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *ObservableTable) HasColumns() bool {
	if o != nil && !IsNil(o.Columns) {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []ObservableField and assigns it to the Columns field.
func (o *ObservableTable) SetColumns(v []ObservableField) {
	o.Columns = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *ObservableTable) GetFilter() Filter {
	if o == nil || IsNil(o.Filter) {
		var ret Filter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservableTable) GetFilterOk() (*Filter, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *ObservableTable) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given Filter and assigns it to the Filter field.
func (o *ObservableTable) SetFilter(v Filter) {
	o.Filter = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *ObservableTable) GetSort() []AlertListSortInner {
	if o == nil || IsNil(o.Sort) {
		var ret []AlertListSortInner
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservableTable) GetSortOk() ([]AlertListSortInner, bool) {
	if o == nil || IsNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *ObservableTable) HasSort() bool {
	if o != nil && !IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given []AlertListSortInner and assigns it to the Sort field.
func (o *ObservableTable) SetSort(v []AlertListSortInner) {
	o.Sort = v
}

// GetProtectData returns the ProtectData field value
func (o *ObservableTable) GetProtectData() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ProtectData
}

// GetProtectDataOk returns a tuple with the ProtectData field value
// and a boolean to check if the value has been set.
func (o *ObservableTable) GetProtectDataOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProtectData, true
}

// SetProtectData sets field value
func (o *ObservableTable) SetProtectData(v bool) {
	o.ProtectData = v
}

// GetMaxElements returns the MaxElements field value if set, zero value otherwise.
func (o *ObservableTable) GetMaxElements() int32 {
	if o == nil || IsNil(o.MaxElements) {
		var ret int32
		return ret
	}
	return *o.MaxElements
}

// GetMaxElementsOk returns a tuple with the MaxElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservableTable) GetMaxElementsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxElements) {
		return nil, false
	}
	return o.MaxElements, true
}

// HasMaxElements returns a boolean if a field has been set.
func (o *ObservableTable) HasMaxElements() bool {
	if o != nil && !IsNil(o.MaxElements) {
		return true
	}

	return false
}

// SetMaxElements gets a reference to the given int32 and assigns it to the MaxElements field.
func (o *ObservableTable) SetMaxElements(v int32) {
	o.MaxElements = &v
}

// GetKind returns the Kind field value
func (o *ObservableTable) GetKind() Kind {
	if o == nil {
		var ret Kind
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ObservableTable) GetKindOk() (*Kind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ObservableTable) SetKind(v Kind) {
	o.Kind = v
}

func (o ObservableTable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObservableTable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Columns) {
		toSerialize["columns"] = o.Columns
	}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.Sort) {
		toSerialize["sort"] = o.Sort
	}
	toSerialize["protectData"] = o.ProtectData
	if !IsNil(o.MaxElements) {
		toSerialize["maxElements"] = o.MaxElements
	}
	toSerialize["_kind"] = o.Kind
	return toSerialize, nil
}

func (o *ObservableTable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"protectData",
		"_kind",
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

	varObservableTable := _ObservableTable{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// Postprocessed by scripts/tolerant-response-decoder: unknown fields are ignored,
	// since TheHive adds fields to its responses in minor versions.
	err = decoder.Decode(&varObservableTable)

	if err != nil {
		return err
	}

	*o = ObservableTable(varObservableTable)

	return err
}

type NullableObservableTable struct {
	value *ObservableTable
	isSet bool
}

func (v NullableObservableTable) Get() *ObservableTable {
	return v.value
}

func (v *NullableObservableTable) Set(val *ObservableTable) {
	v.value = val
	v.isSet = true
}

func (v NullableObservableTable) IsSet() bool {
	return v.isSet
}

func (v *NullableObservableTable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObservableTable(val *ObservableTable) *NullableObservableTable {
	return &NullableObservableTable{value: val, isSet: true}
}

func (v NullableObservableTable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObservableTable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
