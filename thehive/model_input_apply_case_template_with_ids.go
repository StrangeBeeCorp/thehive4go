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

// checks if the InputApplyCaseTemplateWithIds type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputApplyCaseTemplateWithIds{}

// InputApplyCaseTemplateWithIds struct for InputApplyCaseTemplateWithIds
type InputApplyCaseTemplateWithIds struct {
	// Identifiers of the entities to target. IDs must be prefixed with `~`. Case numbers and names are passed as plain strings without a prefix.
	Ids []string `json:"ids"`
	// Name or ID of the case template to apply. Must match an existing case template. Use `POST /api/v1/query` with `listCaseTemplate` to list available case templates.
	CaseTemplate string `json:"caseTemplate"`
	// Applies the title prefix from the case template to each case.
	UpdateTitlePrefix *bool `json:"updateTitlePrefix,omitempty"`
	// Replaces each case's description with the one from the case template.
	UpdateDescription *bool `json:"updateDescription,omitempty"`
	// Applies the tags from the case template to each case.
	UpdateTags *bool `json:"updateTags,omitempty"`
	// Applies the severity from the case template to each case.
	UpdateSeverity *bool `json:"updateSeverity,omitempty"`
	// Applies the flag from the case template to each case.
	UpdateFlag *bool `json:"updateFlag,omitempty"`
	// Applies the [Traffic Light Protocol (TLP)](https://www.misp-project.org/taxonomies.html#_tlp) level from the case template to each case.
	UpdateTlp *bool `json:"updateTlp,omitempty"`
	// Applies the [Permissible Actions Protocol (PAP)](https://www.misp-project.org/taxonomies.html#_pap) level from the case template to each case.
	UpdatePap *bool `json:"updatePap,omitempty"`
	// Applies custom field values from the case template to each case, adding them to any existing values without removing what's already set. Custom fields not defined in the template remain unchanged.
	UpdateCustomFields *bool `json:"updateCustomFields,omitempty"`
	// IDs or titles of tasks in the case template to import into each case. If empty, no tasks are imported. Must match tasks that exist in the specified case template. Use `POST /api/v1/query` with `listCaseTemplate` followed by a `tasks` operation to retrieve available tasks, their IDs, and titles.
	ImportTasks []string `json:"importTasks,omitempty"`
	// IDs of pages in the case template to import into each case. If empty, no pages are imported. Must match page IDs that exist in the specified case template. Use `POST /api/v1/query` with `listCaseTemplate` followed by a `pageTemplates` operation to retrieve available page IDs.
	ImportPages []string `json:"importPages,omitempty"`
}

type _InputApplyCaseTemplateWithIds InputApplyCaseTemplateWithIds

// NewInputApplyCaseTemplateWithIds instantiates a new InputApplyCaseTemplateWithIds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputApplyCaseTemplateWithIds(ids []string, caseTemplate string) *InputApplyCaseTemplateWithIds {
	this := InputApplyCaseTemplateWithIds{}
	this.Ids = ids
	this.CaseTemplate = caseTemplate
	var updateTitlePrefix bool = false
	this.UpdateTitlePrefix = &updateTitlePrefix
	var updateDescription bool = false
	this.UpdateDescription = &updateDescription
	var updateTags bool = false
	this.UpdateTags = &updateTags
	var updateSeverity bool = false
	this.UpdateSeverity = &updateSeverity
	var updateFlag bool = false
	this.UpdateFlag = &updateFlag
	var updateTlp bool = false
	this.UpdateTlp = &updateTlp
	var updatePap bool = false
	this.UpdatePap = &updatePap
	var updateCustomFields bool = false
	this.UpdateCustomFields = &updateCustomFields
	return &this
}

// NewInputApplyCaseTemplateWithIdsWithDefaults instantiates a new InputApplyCaseTemplateWithIds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputApplyCaseTemplateWithIdsWithDefaults() *InputApplyCaseTemplateWithIds {
	this := InputApplyCaseTemplateWithIds{}
	var updateTitlePrefix bool = false
	this.UpdateTitlePrefix = &updateTitlePrefix
	var updateDescription bool = false
	this.UpdateDescription = &updateDescription
	var updateTags bool = false
	this.UpdateTags = &updateTags
	var updateSeverity bool = false
	this.UpdateSeverity = &updateSeverity
	var updateFlag bool = false
	this.UpdateFlag = &updateFlag
	var updateTlp bool = false
	this.UpdateTlp = &updateTlp
	var updatePap bool = false
	this.UpdatePap = &updatePap
	var updateCustomFields bool = false
	this.UpdateCustomFields = &updateCustomFields
	return &this
}

// GetIds returns the Ids field value
func (o *InputApplyCaseTemplateWithIds) GetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *InputApplyCaseTemplateWithIds) GetIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ids, true
}

// SetIds sets field value
func (o *InputApplyCaseTemplateWithIds) SetIds(v []string) {
	o.Ids = v
}

// GetCaseTemplate returns the CaseTemplate field value
func (o *InputApplyCaseTemplateWithIds) GetCaseTemplate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CaseTemplate
}

// GetCaseTemplateOk returns a tuple with the CaseTemplate field value
// and a boolean to check if the value has been set.
func (o *InputApplyCaseTemplateWithIds) GetCaseTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CaseTemplate, true
}

// SetCaseTemplate sets field value
func (o *InputApplyCaseTemplateWithIds) SetCaseTemplate(v string) {
	o.CaseTemplate = v
}

// GetUpdateTitlePrefix returns the UpdateTitlePrefix field value if set, zero value otherwise.
func (o *InputApplyCaseTemplateWithIds) GetUpdateTitlePrefix() bool {
	if o == nil || IsNil(o.UpdateTitlePrefix) {
		var ret bool
		return ret
	}
	return *o.UpdateTitlePrefix
}

// GetUpdateTitlePrefixOk returns a tuple with the UpdateTitlePrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputApplyCaseTemplateWithIds) GetUpdateTitlePrefixOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdateTitlePrefix) {
		return nil, false
	}
	return o.UpdateTitlePrefix, true
}

// HasUpdateTitlePrefix returns a boolean if a field has been set.
func (o *InputApplyCaseTemplateWithIds) HasUpdateTitlePrefix() bool {
	if o != nil && !IsNil(o.UpdateTitlePrefix) {
		return true
	}

	return false
}

// SetUpdateTitlePrefix gets a reference to the given bool and assigns it to the UpdateTitlePrefix field.
func (o *InputApplyCaseTemplateWithIds) SetUpdateTitlePrefix(v bool) {
	o.UpdateTitlePrefix = &v
}

// GetUpdateDescription returns the UpdateDescription field value if set, zero value otherwise.
func (o *InputApplyCaseTemplateWithIds) GetUpdateDescription() bool {
	if o == nil || IsNil(o.UpdateDescription) {
		var ret bool
		return ret
	}
	return *o.UpdateDescription
}

// GetUpdateDescriptionOk returns a tuple with the UpdateDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputApplyCaseTemplateWithIds) GetUpdateDescriptionOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdateDescription) {
		return nil, false
	}
	return o.UpdateDescription, true
}

// HasUpdateDescription returns a boolean if a field has been set.
func (o *InputApplyCaseTemplateWithIds) HasUpdateDescription() bool {
	if o != nil && !IsNil(o.UpdateDescription) {
		return true
	}

	return false
}

// SetUpdateDescription gets a reference to the given bool and assigns it to the UpdateDescription field.
func (o *InputApplyCaseTemplateWithIds) SetUpdateDescription(v bool) {
	o.UpdateDescription = &v
}

// GetUpdateTags returns the UpdateTags field value if set, zero value otherwise.
func (o *InputApplyCaseTemplateWithIds) GetUpdateTags() bool {
	if o == nil || IsNil(o.UpdateTags) {
		var ret bool
		return ret
	}
	return *o.UpdateTags
}

// GetUpdateTagsOk returns a tuple with the UpdateTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputApplyCaseTemplateWithIds) GetUpdateTagsOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdateTags) {
		return nil, false
	}
	return o.UpdateTags, true
}

// HasUpdateTags returns a boolean if a field has been set.
func (o *InputApplyCaseTemplateWithIds) HasUpdateTags() bool {
	if o != nil && !IsNil(o.UpdateTags) {
		return true
	}

	return false
}

// SetUpdateTags gets a reference to the given bool and assigns it to the UpdateTags field.
func (o *InputApplyCaseTemplateWithIds) SetUpdateTags(v bool) {
	o.UpdateTags = &v
}

// GetUpdateSeverity returns the UpdateSeverity field value if set, zero value otherwise.
func (o *InputApplyCaseTemplateWithIds) GetUpdateSeverity() bool {
	if o == nil || IsNil(o.UpdateSeverity) {
		var ret bool
		return ret
	}
	return *o.UpdateSeverity
}

// GetUpdateSeverityOk returns a tuple with the UpdateSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputApplyCaseTemplateWithIds) GetUpdateSeverityOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdateSeverity) {
		return nil, false
	}
	return o.UpdateSeverity, true
}

// HasUpdateSeverity returns a boolean if a field has been set.
func (o *InputApplyCaseTemplateWithIds) HasUpdateSeverity() bool {
	if o != nil && !IsNil(o.UpdateSeverity) {
		return true
	}

	return false
}

// SetUpdateSeverity gets a reference to the given bool and assigns it to the UpdateSeverity field.
func (o *InputApplyCaseTemplateWithIds) SetUpdateSeverity(v bool) {
	o.UpdateSeverity = &v
}

// GetUpdateFlag returns the UpdateFlag field value if set, zero value otherwise.
func (o *InputApplyCaseTemplateWithIds) GetUpdateFlag() bool {
	if o == nil || IsNil(o.UpdateFlag) {
		var ret bool
		return ret
	}
	return *o.UpdateFlag
}

// GetUpdateFlagOk returns a tuple with the UpdateFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputApplyCaseTemplateWithIds) GetUpdateFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdateFlag) {
		return nil, false
	}
	return o.UpdateFlag, true
}

// HasUpdateFlag returns a boolean if a field has been set.
func (o *InputApplyCaseTemplateWithIds) HasUpdateFlag() bool {
	if o != nil && !IsNil(o.UpdateFlag) {
		return true
	}

	return false
}

// SetUpdateFlag gets a reference to the given bool and assigns it to the UpdateFlag field.
func (o *InputApplyCaseTemplateWithIds) SetUpdateFlag(v bool) {
	o.UpdateFlag = &v
}

// GetUpdateTlp returns the UpdateTlp field value if set, zero value otherwise.
func (o *InputApplyCaseTemplateWithIds) GetUpdateTlp() bool {
	if o == nil || IsNil(o.UpdateTlp) {
		var ret bool
		return ret
	}
	return *o.UpdateTlp
}

// GetUpdateTlpOk returns a tuple with the UpdateTlp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputApplyCaseTemplateWithIds) GetUpdateTlpOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdateTlp) {
		return nil, false
	}
	return o.UpdateTlp, true
}

// HasUpdateTlp returns a boolean if a field has been set.
func (o *InputApplyCaseTemplateWithIds) HasUpdateTlp() bool {
	if o != nil && !IsNil(o.UpdateTlp) {
		return true
	}

	return false
}

// SetUpdateTlp gets a reference to the given bool and assigns it to the UpdateTlp field.
func (o *InputApplyCaseTemplateWithIds) SetUpdateTlp(v bool) {
	o.UpdateTlp = &v
}

// GetUpdatePap returns the UpdatePap field value if set, zero value otherwise.
func (o *InputApplyCaseTemplateWithIds) GetUpdatePap() bool {
	if o == nil || IsNil(o.UpdatePap) {
		var ret bool
		return ret
	}
	return *o.UpdatePap
}

// GetUpdatePapOk returns a tuple with the UpdatePap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputApplyCaseTemplateWithIds) GetUpdatePapOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdatePap) {
		return nil, false
	}
	return o.UpdatePap, true
}

// HasUpdatePap returns a boolean if a field has been set.
func (o *InputApplyCaseTemplateWithIds) HasUpdatePap() bool {
	if o != nil && !IsNil(o.UpdatePap) {
		return true
	}

	return false
}

// SetUpdatePap gets a reference to the given bool and assigns it to the UpdatePap field.
func (o *InputApplyCaseTemplateWithIds) SetUpdatePap(v bool) {
	o.UpdatePap = &v
}

// GetUpdateCustomFields returns the UpdateCustomFields field value if set, zero value otherwise.
func (o *InputApplyCaseTemplateWithIds) GetUpdateCustomFields() bool {
	if o == nil || IsNil(o.UpdateCustomFields) {
		var ret bool
		return ret
	}
	return *o.UpdateCustomFields
}

// GetUpdateCustomFieldsOk returns a tuple with the UpdateCustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputApplyCaseTemplateWithIds) GetUpdateCustomFieldsOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdateCustomFields) {
		return nil, false
	}
	return o.UpdateCustomFields, true
}

// HasUpdateCustomFields returns a boolean if a field has been set.
func (o *InputApplyCaseTemplateWithIds) HasUpdateCustomFields() bool {
	if o != nil && !IsNil(o.UpdateCustomFields) {
		return true
	}

	return false
}

// SetUpdateCustomFields gets a reference to the given bool and assigns it to the UpdateCustomFields field.
func (o *InputApplyCaseTemplateWithIds) SetUpdateCustomFields(v bool) {
	o.UpdateCustomFields = &v
}

// GetImportTasks returns the ImportTasks field value if set, zero value otherwise.
func (o *InputApplyCaseTemplateWithIds) GetImportTasks() []string {
	if o == nil || IsNil(o.ImportTasks) {
		var ret []string
		return ret
	}
	return o.ImportTasks
}

// GetImportTasksOk returns a tuple with the ImportTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputApplyCaseTemplateWithIds) GetImportTasksOk() ([]string, bool) {
	if o == nil || IsNil(o.ImportTasks) {
		return nil, false
	}
	return o.ImportTasks, true
}

// HasImportTasks returns a boolean if a field has been set.
func (o *InputApplyCaseTemplateWithIds) HasImportTasks() bool {
	if o != nil && !IsNil(o.ImportTasks) {
		return true
	}

	return false
}

// SetImportTasks gets a reference to the given []string and assigns it to the ImportTasks field.
func (o *InputApplyCaseTemplateWithIds) SetImportTasks(v []string) {
	o.ImportTasks = v
}

// GetImportPages returns the ImportPages field value if set, zero value otherwise.
func (o *InputApplyCaseTemplateWithIds) GetImportPages() []string {
	if o == nil || IsNil(o.ImportPages) {
		var ret []string
		return ret
	}
	return o.ImportPages
}

// GetImportPagesOk returns a tuple with the ImportPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputApplyCaseTemplateWithIds) GetImportPagesOk() ([]string, bool) {
	if o == nil || IsNil(o.ImportPages) {
		return nil, false
	}
	return o.ImportPages, true
}

// HasImportPages returns a boolean if a field has been set.
func (o *InputApplyCaseTemplateWithIds) HasImportPages() bool {
	if o != nil && !IsNil(o.ImportPages) {
		return true
	}

	return false
}

// SetImportPages gets a reference to the given []string and assigns it to the ImportPages field.
func (o *InputApplyCaseTemplateWithIds) SetImportPages(v []string) {
	o.ImportPages = v
}

func (o InputApplyCaseTemplateWithIds) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputApplyCaseTemplateWithIds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ids"] = o.Ids
	toSerialize["caseTemplate"] = o.CaseTemplate
	if !IsNil(o.UpdateTitlePrefix) {
		toSerialize["updateTitlePrefix"] = o.UpdateTitlePrefix
	}
	if !IsNil(o.UpdateDescription) {
		toSerialize["updateDescription"] = o.UpdateDescription
	}
	if !IsNil(o.UpdateTags) {
		toSerialize["updateTags"] = o.UpdateTags
	}
	if !IsNil(o.UpdateSeverity) {
		toSerialize["updateSeverity"] = o.UpdateSeverity
	}
	if !IsNil(o.UpdateFlag) {
		toSerialize["updateFlag"] = o.UpdateFlag
	}
	if !IsNil(o.UpdateTlp) {
		toSerialize["updateTlp"] = o.UpdateTlp
	}
	if !IsNil(o.UpdatePap) {
		toSerialize["updatePap"] = o.UpdatePap
	}
	if !IsNil(o.UpdateCustomFields) {
		toSerialize["updateCustomFields"] = o.UpdateCustomFields
	}
	if !IsNil(o.ImportTasks) {
		toSerialize["importTasks"] = o.ImportTasks
	}
	if !IsNil(o.ImportPages) {
		toSerialize["importPages"] = o.ImportPages
	}
	return toSerialize, nil
}

func (o *InputApplyCaseTemplateWithIds) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ids",
		"caseTemplate",
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

	varInputApplyCaseTemplateWithIds := _InputApplyCaseTemplateWithIds{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInputApplyCaseTemplateWithIds)

	if err != nil {
		return err
	}

	*o = InputApplyCaseTemplateWithIds(varInputApplyCaseTemplateWithIds)

	return err
}

type NullableInputApplyCaseTemplateWithIds struct {
	value *InputApplyCaseTemplateWithIds
	isSet bool
}

func (v NullableInputApplyCaseTemplateWithIds) Get() *InputApplyCaseTemplateWithIds {
	return v.value
}

func (v *NullableInputApplyCaseTemplateWithIds) Set(val *InputApplyCaseTemplateWithIds) {
	v.value = val
	v.isSet = true
}

func (v NullableInputApplyCaseTemplateWithIds) IsSet() bool {
	return v.isSet
}

func (v *NullableInputApplyCaseTemplateWithIds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputApplyCaseTemplateWithIds(val *InputApplyCaseTemplateWithIds) *NullableInputApplyCaseTemplateWithIds {
	return &NullableInputApplyCaseTemplateWithIds{value: val, isSet: true}
}

func (v NullableInputApplyCaseTemplateWithIds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputApplyCaseTemplateWithIds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
