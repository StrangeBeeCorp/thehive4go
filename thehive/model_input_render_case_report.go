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

// checks if the InputRenderCaseReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputRenderCaseReport{}

// InputRenderCaseReport struct for InputRenderCaseReport
type InputRenderCaseReport struct {
	// Output format for the rendered report.
	Format string `json:"format"`
	// Inline template definition used when rendering without a saved case report template. A definition is an ordered list of widgets—building blocks that each render one section of the report, such as a `Text` widget for free-form Mustache content or an `ObservableList` widget for the case observables. Either `definition` or `caseReportTemplateId` is required.
	Definition *CaseReportTemplateDefinition `json:"definition,omitempty"`
	// ID of the saved case report template to use. Must match an existing template. Use `POST /api/v1/query` with `listCaseReportTemplate` to retrieve available templates and their IDs. Either `caseReportTemplateId` or `definition` is required.
	CaseReportTemplateId *string `json:"caseReportTemplateId,omitempty"`
	// ID of the case to use for data. Use `POST /api/v1/query` with `listCase` to retrieve available cases and their IDs. If omitted, TheHive uses fake data to fill the template.
	CaseId *string `json:"caseId,omitempty"`
	// Maximum number of elements displayed per widget in the preview. Doesn't affect the number of elements included when saving or downloading the report.
	MaxElements *int32 `json:"maxElements,omitempty"`
}

type _InputRenderCaseReport InputRenderCaseReport

// NewInputRenderCaseReport instantiates a new InputRenderCaseReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputRenderCaseReport(format string) *InputRenderCaseReport {
	this := InputRenderCaseReport{}
	this.Format = format
	return &this
}

// NewInputRenderCaseReportWithDefaults instantiates a new InputRenderCaseReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputRenderCaseReportWithDefaults() *InputRenderCaseReport {
	this := InputRenderCaseReport{}
	return &this
}

// GetFormat returns the Format field value
func (o *InputRenderCaseReport) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *InputRenderCaseReport) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *InputRenderCaseReport) SetFormat(v string) {
	o.Format = v
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *InputRenderCaseReport) GetDefinition() CaseReportTemplateDefinition {
	if o == nil || IsNil(o.Definition) {
		var ret CaseReportTemplateDefinition
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputRenderCaseReport) GetDefinitionOk() (*CaseReportTemplateDefinition, bool) {
	if o == nil || IsNil(o.Definition) {
		return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *InputRenderCaseReport) HasDefinition() bool {
	if o != nil && !IsNil(o.Definition) {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given CaseReportTemplateDefinition and assigns it to the Definition field.
func (o *InputRenderCaseReport) SetDefinition(v CaseReportTemplateDefinition) {
	o.Definition = &v
}

// GetCaseReportTemplateId returns the CaseReportTemplateId field value if set, zero value otherwise.
func (o *InputRenderCaseReport) GetCaseReportTemplateId() string {
	if o == nil || IsNil(o.CaseReportTemplateId) {
		var ret string
		return ret
	}
	return *o.CaseReportTemplateId
}

// GetCaseReportTemplateIdOk returns a tuple with the CaseReportTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputRenderCaseReport) GetCaseReportTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.CaseReportTemplateId) {
		return nil, false
	}
	return o.CaseReportTemplateId, true
}

// HasCaseReportTemplateId returns a boolean if a field has been set.
func (o *InputRenderCaseReport) HasCaseReportTemplateId() bool {
	if o != nil && !IsNil(o.CaseReportTemplateId) {
		return true
	}

	return false
}

// SetCaseReportTemplateId gets a reference to the given string and assigns it to the CaseReportTemplateId field.
func (o *InputRenderCaseReport) SetCaseReportTemplateId(v string) {
	o.CaseReportTemplateId = &v
}

// GetCaseId returns the CaseId field value if set, zero value otherwise.
func (o *InputRenderCaseReport) GetCaseId() string {
	if o == nil || IsNil(o.CaseId) {
		var ret string
		return ret
	}
	return *o.CaseId
}

// GetCaseIdOk returns a tuple with the CaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputRenderCaseReport) GetCaseIdOk() (*string, bool) {
	if o == nil || IsNil(o.CaseId) {
		return nil, false
	}
	return o.CaseId, true
}

// HasCaseId returns a boolean if a field has been set.
func (o *InputRenderCaseReport) HasCaseId() bool {
	if o != nil && !IsNil(o.CaseId) {
		return true
	}

	return false
}

// SetCaseId gets a reference to the given string and assigns it to the CaseId field.
func (o *InputRenderCaseReport) SetCaseId(v string) {
	o.CaseId = &v
}

// GetMaxElements returns the MaxElements field value if set, zero value otherwise.
func (o *InputRenderCaseReport) GetMaxElements() int32 {
	if o == nil || IsNil(o.MaxElements) {
		var ret int32
		return ret
	}
	return *o.MaxElements
}

// GetMaxElementsOk returns a tuple with the MaxElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputRenderCaseReport) GetMaxElementsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxElements) {
		return nil, false
	}
	return o.MaxElements, true
}

// HasMaxElements returns a boolean if a field has been set.
func (o *InputRenderCaseReport) HasMaxElements() bool {
	if o != nil && !IsNil(o.MaxElements) {
		return true
	}

	return false
}

// SetMaxElements gets a reference to the given int32 and assigns it to the MaxElements field.
func (o *InputRenderCaseReport) SetMaxElements(v int32) {
	o.MaxElements = &v
}

func (o InputRenderCaseReport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputRenderCaseReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["format"] = o.Format
	if !IsNil(o.Definition) {
		toSerialize["definition"] = o.Definition
	}
	if !IsNil(o.CaseReportTemplateId) {
		toSerialize["caseReportTemplateId"] = o.CaseReportTemplateId
	}
	if !IsNil(o.CaseId) {
		toSerialize["caseId"] = o.CaseId
	}
	if !IsNil(o.MaxElements) {
		toSerialize["maxElements"] = o.MaxElements
	}
	return toSerialize, nil
}

func (o *InputRenderCaseReport) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"format",
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

	varInputRenderCaseReport := _InputRenderCaseReport{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInputRenderCaseReport)

	if err != nil {
		return err
	}

	*o = InputRenderCaseReport(varInputRenderCaseReport)

	return err
}

type NullableInputRenderCaseReport struct {
	value *InputRenderCaseReport
	isSet bool
}

func (v NullableInputRenderCaseReport) Get() *InputRenderCaseReport {
	return v.value
}

func (v *NullableInputRenderCaseReport) Set(val *InputRenderCaseReport) {
	v.value = val
	v.isSet = true
}

func (v NullableInputRenderCaseReport) IsSet() bool {
	return v.isSet
}

func (v *NullableInputRenderCaseReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputRenderCaseReport(val *InputRenderCaseReport) *NullableInputRenderCaseReport {
	return &NullableInputRenderCaseReport{value: val, isSet: true}
}

func (v NullableInputRenderCaseReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputRenderCaseReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
