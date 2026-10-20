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

// checks if the InputCreateObservable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputCreateObservable{}

// InputCreateObservable struct for InputCreateObservable
type InputCreateObservable struct {
	// The observable type. Must match an existing observable type defined in TheHive. When set, this field takes priority over the `dataType` query parameter. Use `POST /api/v1/query` with `listObservableType` to list available types.
	DataType string                     `json:"dataType"`
	Data     *InputCreateObservableData `json:"data,omitempty"`
	// Note or additional context about the observable. Supports [TheHive-flavored Markdown](https://docs.strangebee.com/thehive/user-guides/thehive-flavored-markdown/).
	Message *string `json:"message,omitempty"`
	// Deprecated. Maps to the `_createdAt` field. Don't use this field.
	StartDate  *int64                           `json:"startDate,omitempty"`
	Attachment *InputCreateObservableAttachment `json:"attachment,omitempty"`
	// [Traffic Light Protocol (TLP)](https://www.misp-project.org/taxonomies.html#_tlp) level controlling how the observable's information can be shared: 0 (Clear — no restrictions), 1 (Green — community), 2 (Amber — organization and clients), 3 (Amber+Strict — organization only), 4 (Red — named recipients only).
	Tlp *int32 `json:"tlp,omitempty"`
	// [Permissible Actions Protocol (PAP)](https://www.misp-project.org/taxonomies.html#_pap) level controlling what actions can be taken with the observable data: 0 (Clear — no restrictions), 1 (Green — active actions allowed), 2 (Amber — passive checks only), 3 (Red — non-detectable actions only).
	Pap *int32 `json:"pap,omitempty"`
	// Labels to categorize the observable. Unknown tags are automatically created. Use `POST /api/v1/query` with `listTag` to retrieve existing custom tags, or with `listTaxonomy` and `tags` to retrieve taxonomy tags.
	Tags []string `json:"tags,omitempty"`
	// Marks the observable as an indicator of compromise (IOC), indicating it's linked to suspicious or malicious activity.
	Ioc *bool `json:"ioc,omitempty"`
	// Marks the observable as sighted, meaning it has been detected or observed in the environment.
	Sighted *bool `json:"sighted,omitempty"`
	// Date and time when the observable was sighted. Expected format: Unix timestamp in milliseconds.
	SightedAt *int64 `json:"sightedAt,omitempty"`
	// When `true`, excludes the observable from similarity checks used to identify related alerts and cases. Useful for contextual observables such as the company domain name that aren't indicators of a threat.
	IgnoreSimilarity *bool `json:"ignoreSimilarity,omitempty"`
	// When `true`, TheHive unpacks the zip archive using `zipPassword` (if provided) and creates one observable per file it contains.
	IsZip *bool `json:"isZip,omitempty"`
	// Password to decrypt the zip archive when `isZip` is `true`.
	ZipPassword *string `json:"zipPassword,omitempty"`
}

type _InputCreateObservable InputCreateObservable

// NewInputCreateObservable instantiates a new InputCreateObservable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputCreateObservable(dataType string) *InputCreateObservable {
	this := InputCreateObservable{}
	this.DataType = dataType
	var tlp int32 = 2
	this.Tlp = &tlp
	var pap int32 = 2
	this.Pap = &pap
	var ioc bool = false
	this.Ioc = &ioc
	var sighted bool = false
	this.Sighted = &sighted
	var ignoreSimilarity bool = false
	this.IgnoreSimilarity = &ignoreSimilarity
	var isZip bool = false
	this.IsZip = &isZip
	return &this
}

// NewInputCreateObservableWithDefaults instantiates a new InputCreateObservable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputCreateObservableWithDefaults() *InputCreateObservable {
	this := InputCreateObservable{}
	var tlp int32 = 2
	this.Tlp = &tlp
	var pap int32 = 2
	this.Pap = &pap
	var ioc bool = false
	this.Ioc = &ioc
	var sighted bool = false
	this.Sighted = &sighted
	var ignoreSimilarity bool = false
	this.IgnoreSimilarity = &ignoreSimilarity
	var isZip bool = false
	this.IsZip = &isZip
	return &this
}

// GetDataType returns the DataType field value
func (o *InputCreateObservable) GetDataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value
// and a boolean to check if the value has been set.
func (o *InputCreateObservable) GetDataTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataType, true
}

// SetDataType sets field value
func (o *InputCreateObservable) SetDataType(v string) {
	o.DataType = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InputCreateObservable) GetData() InputCreateObservableData {
	if o == nil || IsNil(o.Data) {
		var ret InputCreateObservableData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateObservable) GetDataOk() (*InputCreateObservableData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InputCreateObservable) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given InputCreateObservableData and assigns it to the Data field.
func (o *InputCreateObservable) SetData(v InputCreateObservableData) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *InputCreateObservable) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateObservable) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InputCreateObservable) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *InputCreateObservable) SetMessage(v string) {
	o.Message = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *InputCreateObservable) GetStartDate() int64 {
	if o == nil || IsNil(o.StartDate) {
		var ret int64
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateObservable) GetStartDateOk() (*int64, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *InputCreateObservable) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given int64 and assigns it to the StartDate field.
func (o *InputCreateObservable) SetStartDate(v int64) {
	o.StartDate = &v
}

// GetAttachment returns the Attachment field value if set, zero value otherwise.
func (o *InputCreateObservable) GetAttachment() InputCreateObservableAttachment {
	if o == nil || IsNil(o.Attachment) {
		var ret InputCreateObservableAttachment
		return ret
	}
	return *o.Attachment
}

// GetAttachmentOk returns a tuple with the Attachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateObservable) GetAttachmentOk() (*InputCreateObservableAttachment, bool) {
	if o == nil || IsNil(o.Attachment) {
		return nil, false
	}
	return o.Attachment, true
}

// HasAttachment returns a boolean if a field has been set.
func (o *InputCreateObservable) HasAttachment() bool {
	if o != nil && !IsNil(o.Attachment) {
		return true
	}

	return false
}

// SetAttachment gets a reference to the given InputCreateObservableAttachment and assigns it to the Attachment field.
func (o *InputCreateObservable) SetAttachment(v InputCreateObservableAttachment) {
	o.Attachment = &v
}

// GetTlp returns the Tlp field value if set, zero value otherwise.
func (o *InputCreateObservable) GetTlp() int32 {
	if o == nil || IsNil(o.Tlp) {
		var ret int32
		return ret
	}
	return *o.Tlp
}

// GetTlpOk returns a tuple with the Tlp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateObservable) GetTlpOk() (*int32, bool) {
	if o == nil || IsNil(o.Tlp) {
		return nil, false
	}
	return o.Tlp, true
}

// HasTlp returns a boolean if a field has been set.
func (o *InputCreateObservable) HasTlp() bool {
	if o != nil && !IsNil(o.Tlp) {
		return true
	}

	return false
}

// SetTlp gets a reference to the given int32 and assigns it to the Tlp field.
func (o *InputCreateObservable) SetTlp(v int32) {
	o.Tlp = &v
}

// GetPap returns the Pap field value if set, zero value otherwise.
func (o *InputCreateObservable) GetPap() int32 {
	if o == nil || IsNil(o.Pap) {
		var ret int32
		return ret
	}
	return *o.Pap
}

// GetPapOk returns a tuple with the Pap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateObservable) GetPapOk() (*int32, bool) {
	if o == nil || IsNil(o.Pap) {
		return nil, false
	}
	return o.Pap, true
}

// HasPap returns a boolean if a field has been set.
func (o *InputCreateObservable) HasPap() bool {
	if o != nil && !IsNil(o.Pap) {
		return true
	}

	return false
}

// SetPap gets a reference to the given int32 and assigns it to the Pap field.
func (o *InputCreateObservable) SetPap(v int32) {
	o.Pap = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InputCreateObservable) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateObservable) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InputCreateObservable) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InputCreateObservable) SetTags(v []string) {
	o.Tags = v
}

// GetIoc returns the Ioc field value if set, zero value otherwise.
func (o *InputCreateObservable) GetIoc() bool {
	if o == nil || IsNil(o.Ioc) {
		var ret bool
		return ret
	}
	return *o.Ioc
}

// GetIocOk returns a tuple with the Ioc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateObservable) GetIocOk() (*bool, bool) {
	if o == nil || IsNil(o.Ioc) {
		return nil, false
	}
	return o.Ioc, true
}

// HasIoc returns a boolean if a field has been set.
func (o *InputCreateObservable) HasIoc() bool {
	if o != nil && !IsNil(o.Ioc) {
		return true
	}

	return false
}

// SetIoc gets a reference to the given bool and assigns it to the Ioc field.
func (o *InputCreateObservable) SetIoc(v bool) {
	o.Ioc = &v
}

// GetSighted returns the Sighted field value if set, zero value otherwise.
func (o *InputCreateObservable) GetSighted() bool {
	if o == nil || IsNil(o.Sighted) {
		var ret bool
		return ret
	}
	return *o.Sighted
}

// GetSightedOk returns a tuple with the Sighted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateObservable) GetSightedOk() (*bool, bool) {
	if o == nil || IsNil(o.Sighted) {
		return nil, false
	}
	return o.Sighted, true
}

// HasSighted returns a boolean if a field has been set.
func (o *InputCreateObservable) HasSighted() bool {
	if o != nil && !IsNil(o.Sighted) {
		return true
	}

	return false
}

// SetSighted gets a reference to the given bool and assigns it to the Sighted field.
func (o *InputCreateObservable) SetSighted(v bool) {
	o.Sighted = &v
}

// GetSightedAt returns the SightedAt field value if set, zero value otherwise.
func (o *InputCreateObservable) GetSightedAt() int64 {
	if o == nil || IsNil(o.SightedAt) {
		var ret int64
		return ret
	}
	return *o.SightedAt
}

// GetSightedAtOk returns a tuple with the SightedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateObservable) GetSightedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.SightedAt) {
		return nil, false
	}
	return o.SightedAt, true
}

// HasSightedAt returns a boolean if a field has been set.
func (o *InputCreateObservable) HasSightedAt() bool {
	if o != nil && !IsNil(o.SightedAt) {
		return true
	}

	return false
}

// SetSightedAt gets a reference to the given int64 and assigns it to the SightedAt field.
func (o *InputCreateObservable) SetSightedAt(v int64) {
	o.SightedAt = &v
}

// GetIgnoreSimilarity returns the IgnoreSimilarity field value if set, zero value otherwise.
func (o *InputCreateObservable) GetIgnoreSimilarity() bool {
	if o == nil || IsNil(o.IgnoreSimilarity) {
		var ret bool
		return ret
	}
	return *o.IgnoreSimilarity
}

// GetIgnoreSimilarityOk returns a tuple with the IgnoreSimilarity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateObservable) GetIgnoreSimilarityOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreSimilarity) {
		return nil, false
	}
	return o.IgnoreSimilarity, true
}

// HasIgnoreSimilarity returns a boolean if a field has been set.
func (o *InputCreateObservable) HasIgnoreSimilarity() bool {
	if o != nil && !IsNil(o.IgnoreSimilarity) {
		return true
	}

	return false
}

// SetIgnoreSimilarity gets a reference to the given bool and assigns it to the IgnoreSimilarity field.
func (o *InputCreateObservable) SetIgnoreSimilarity(v bool) {
	o.IgnoreSimilarity = &v
}

// GetIsZip returns the IsZip field value if set, zero value otherwise.
func (o *InputCreateObservable) GetIsZip() bool {
	if o == nil || IsNil(o.IsZip) {
		var ret bool
		return ret
	}
	return *o.IsZip
}

// GetIsZipOk returns a tuple with the IsZip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateObservable) GetIsZipOk() (*bool, bool) {
	if o == nil || IsNil(o.IsZip) {
		return nil, false
	}
	return o.IsZip, true
}

// HasIsZip returns a boolean if a field has been set.
func (o *InputCreateObservable) HasIsZip() bool {
	if o != nil && !IsNil(o.IsZip) {
		return true
	}

	return false
}

// SetIsZip gets a reference to the given bool and assigns it to the IsZip field.
func (o *InputCreateObservable) SetIsZip(v bool) {
	o.IsZip = &v
}

// GetZipPassword returns the ZipPassword field value if set, zero value otherwise.
func (o *InputCreateObservable) GetZipPassword() string {
	if o == nil || IsNil(o.ZipPassword) {
		var ret string
		return ret
	}
	return *o.ZipPassword
}

// GetZipPasswordOk returns a tuple with the ZipPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputCreateObservable) GetZipPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.ZipPassword) {
		return nil, false
	}
	return o.ZipPassword, true
}

// HasZipPassword returns a boolean if a field has been set.
func (o *InputCreateObservable) HasZipPassword() bool {
	if o != nil && !IsNil(o.ZipPassword) {
		return true
	}

	return false
}

// SetZipPassword gets a reference to the given string and assigns it to the ZipPassword field.
func (o *InputCreateObservable) SetZipPassword(v string) {
	o.ZipPassword = &v
}

func (o InputCreateObservable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputCreateObservable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dataType"] = o.DataType
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.Attachment) {
		toSerialize["attachment"] = o.Attachment
	}
	if !IsNil(o.Tlp) {
		toSerialize["tlp"] = o.Tlp
	}
	if !IsNil(o.Pap) {
		toSerialize["pap"] = o.Pap
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Ioc) {
		toSerialize["ioc"] = o.Ioc
	}
	if !IsNil(o.Sighted) {
		toSerialize["sighted"] = o.Sighted
	}
	if !IsNil(o.SightedAt) {
		toSerialize["sightedAt"] = o.SightedAt
	}
	if !IsNil(o.IgnoreSimilarity) {
		toSerialize["ignoreSimilarity"] = o.IgnoreSimilarity
	}
	if !IsNil(o.IsZip) {
		toSerialize["isZip"] = o.IsZip
	}
	if !IsNil(o.ZipPassword) {
		toSerialize["zipPassword"] = o.ZipPassword
	}
	return toSerialize, nil
}

func (o *InputCreateObservable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dataType",
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

	varInputCreateObservable := _InputCreateObservable{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInputCreateObservable)

	if err != nil {
		return err
	}

	*o = InputCreateObservable(varInputCreateObservable)

	return err
}

type NullableInputCreateObservable struct {
	value *InputCreateObservable
	isSet bool
}

func (v NullableInputCreateObservable) Get() *InputCreateObservable {
	return v.value
}

func (v *NullableInputCreateObservable) Set(val *InputCreateObservable) {
	v.value = val
	v.isSet = true
}

func (v NullableInputCreateObservable) IsSet() bool {
	return v.isSet
}

func (v *NullableInputCreateObservable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputCreateObservable(val *InputCreateObservable) *NullableInputCreateObservable {
	return &NullableInputCreateObservable{value: val, isSet: true}
}

func (v NullableInputCreateObservable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputCreateObservable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
