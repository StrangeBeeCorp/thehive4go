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

// checks if the OutputLicense type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutputLicense{}

// OutputLicense struct for OutputLicense
type OutputLicense struct {
	// Internal ID of the license.
	UnderscoreId string `json:"_id"`
	// Date the license was added to TheHive, as a Unix timestamp in milliseconds.
	UnderscoreCreatedAt int64 `json:"_createdAt"`
	// Login of the user who added the license.
	UnderscoreCreatedBy string `json:"_createdBy"`
	// Entity type. Always `License`.
	UnderscoreType string `json:"_type"`
	// License identifier from the license file.
	Id string `json:"id"`
	// Customer name as recorded in the license file.
	Customer string `json:"customer"`
	// License plan. One of `No`, `Community`, `Gold`, `Platinum`, or `One`.
	Plan string `json:"plan"`
	// License kind. One of `Regular`, `Trial`, or `Development`.
	Kind string `json:"kind"`
	// Start of the license validity period, as a Unix timestamp in milliseconds.
	ValidFrom int64 `json:"validFrom"`
	// End of the license validity period, as a Unix timestamp in milliseconds.
	ExpiresAt int64 `json:"expiresAt"`
	// Whether this is the currently active license.
	Current bool `json:"current"`
	// Capability names enabled by this license.
	Capabilities []string `json:"capabilities,omitempty"`
	// Resource quotas defined by the license, mapping capability name to its limit. A value of `-1` means unlimited.
	Quotas map[string]int32 `json:"quotas"`
}

type _OutputLicense OutputLicense

// NewOutputLicense instantiates a new OutputLicense object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputLicense(underscoreId string, underscoreCreatedAt int64, underscoreCreatedBy string, underscoreType string, id string, customer string, plan string, kind string, validFrom int64, expiresAt int64, current bool, quotas map[string]int32) *OutputLicense {
	this := OutputLicense{}
	this.UnderscoreId = underscoreId
	this.UnderscoreCreatedAt = underscoreCreatedAt
	this.UnderscoreCreatedBy = underscoreCreatedBy
	this.UnderscoreType = underscoreType
	this.Id = id
	this.Customer = customer
	this.Plan = plan
	this.Kind = kind
	this.ValidFrom = validFrom
	this.ExpiresAt = expiresAt
	this.Current = current
	this.Quotas = quotas
	return &this
}

// NewOutputLicenseWithDefaults instantiates a new OutputLicense object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputLicenseWithDefaults() *OutputLicense {
	this := OutputLicense{}
	return &this
}

// GetUnderscoreId returns the UnderscoreId field value
func (o *OutputLicense) GetUnderscoreId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnderscoreId
}

// GetUnderscoreIdOk returns a tuple with the UnderscoreId field value
// and a boolean to check if the value has been set.
func (o *OutputLicense) GetUnderscoreIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreId, true
}

// SetUnderscoreId sets field value
func (o *OutputLicense) SetUnderscoreId(v string) {
	o.UnderscoreId = v
}

// GetUnderscoreCreatedAt returns the UnderscoreCreatedAt field value
func (o *OutputLicense) GetUnderscoreCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UnderscoreCreatedAt
}

// GetUnderscoreCreatedAtOk returns a tuple with the UnderscoreCreatedAt field value
// and a boolean to check if the value has been set.
func (o *OutputLicense) GetUnderscoreCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreCreatedAt, true
}

// SetUnderscoreCreatedAt sets field value
func (o *OutputLicense) SetUnderscoreCreatedAt(v int64) {
	o.UnderscoreCreatedAt = v
}

// GetUnderscoreCreatedBy returns the UnderscoreCreatedBy field value
func (o *OutputLicense) GetUnderscoreCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnderscoreCreatedBy
}

// GetUnderscoreCreatedByOk returns a tuple with the UnderscoreCreatedBy field value
// and a boolean to check if the value has been set.
func (o *OutputLicense) GetUnderscoreCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreCreatedBy, true
}

// SetUnderscoreCreatedBy sets field value
func (o *OutputLicense) SetUnderscoreCreatedBy(v string) {
	o.UnderscoreCreatedBy = v
}

// GetUnderscoreType returns the UnderscoreType field value
func (o *OutputLicense) GetUnderscoreType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnderscoreType
}

// GetUnderscoreTypeOk returns a tuple with the UnderscoreType field value
// and a boolean to check if the value has been set.
func (o *OutputLicense) GetUnderscoreTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreType, true
}

// SetUnderscoreType sets field value
func (o *OutputLicense) SetUnderscoreType(v string) {
	o.UnderscoreType = v
}

// GetId returns the Id field value
func (o *OutputLicense) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OutputLicense) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OutputLicense) SetId(v string) {
	o.Id = v
}

// GetCustomer returns the Customer field value
func (o *OutputLicense) GetCustomer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value
// and a boolean to check if the value has been set.
func (o *OutputLicense) GetCustomerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Customer, true
}

// SetCustomer sets field value
func (o *OutputLicense) SetCustomer(v string) {
	o.Customer = v
}

// GetPlan returns the Plan field value
func (o *OutputLicense) GetPlan() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Plan
}

// GetPlanOk returns a tuple with the Plan field value
// and a boolean to check if the value has been set.
func (o *OutputLicense) GetPlanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plan, true
}

// SetPlan sets field value
func (o *OutputLicense) SetPlan(v string) {
	o.Plan = v
}

// GetKind returns the Kind field value
func (o *OutputLicense) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *OutputLicense) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *OutputLicense) SetKind(v string) {
	o.Kind = v
}

// GetValidFrom returns the ValidFrom field value
func (o *OutputLicense) GetValidFrom() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value
// and a boolean to check if the value has been set.
func (o *OutputLicense) GetValidFromOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidFrom, true
}

// SetValidFrom sets field value
func (o *OutputLicense) SetValidFrom(v int64) {
	o.ValidFrom = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *OutputLicense) GetExpiresAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *OutputLicense) GetExpiresAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *OutputLicense) SetExpiresAt(v int64) {
	o.ExpiresAt = v
}

// GetCurrent returns the Current field value
func (o *OutputLicense) GetCurrent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Current
}

// GetCurrentOk returns a tuple with the Current field value
// and a boolean to check if the value has been set.
func (o *OutputLicense) GetCurrentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Current, true
}

// SetCurrent sets field value
func (o *OutputLicense) SetCurrent(v bool) {
	o.Current = v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *OutputLicense) GetCapabilities() []string {
	if o == nil || IsNil(o.Capabilities) {
		var ret []string
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputLicense) GetCapabilitiesOk() ([]string, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *OutputLicense) HasCapabilities() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []string and assigns it to the Capabilities field.
func (o *OutputLicense) SetCapabilities(v []string) {
	o.Capabilities = v
}

// GetQuotas returns the Quotas field value
func (o *OutputLicense) GetQuotas() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.Quotas
}

// GetQuotasOk returns a tuple with the Quotas field value
// and a boolean to check if the value has been set.
func (o *OutputLicense) GetQuotasOk() (map[string]int32, bool) {
	if o == nil {
		return map[string]int32{}, false
	}
	return o.Quotas, true
}

// SetQuotas sets field value
func (o *OutputLicense) SetQuotas(v map[string]int32) {
	o.Quotas = v
}

func (o OutputLicense) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutputLicense) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.UnderscoreId
	toSerialize["_createdAt"] = o.UnderscoreCreatedAt
	toSerialize["_createdBy"] = o.UnderscoreCreatedBy
	toSerialize["_type"] = o.UnderscoreType
	toSerialize["id"] = o.Id
	toSerialize["customer"] = o.Customer
	toSerialize["plan"] = o.Plan
	toSerialize["kind"] = o.Kind
	toSerialize["validFrom"] = o.ValidFrom
	toSerialize["expiresAt"] = o.ExpiresAt
	toSerialize["current"] = o.Current
	if !IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	toSerialize["quotas"] = o.Quotas
	return toSerialize, nil
}

func (o *OutputLicense) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"_createdAt",
		"_createdBy",
		"_type",
		"id",
		"customer",
		"plan",
		"kind",
		"validFrom",
		"expiresAt",
		"current",
		"quotas",
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

	varOutputLicense := _OutputLicense{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// Postprocessed by scripts/tolerant-response-decoder: unknown fields are ignored,
	// since TheHive adds fields to its responses in minor versions.
	err = decoder.Decode(&varOutputLicense)

	if err != nil {
		return err
	}

	*o = OutputLicense(varOutputLicense)

	return err
}

type NullableOutputLicense struct {
	value *OutputLicense
	isSet bool
}

func (v NullableOutputLicense) Get() *OutputLicense {
	return v.value
}

func (v *NullableOutputLicense) Set(val *OutputLicense) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputLicense) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputLicense) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputLicense(val *OutputLicense) *NullableOutputLicense {
	return &NullableOutputLicense{value: val, isSet: true}
}

func (v NullableOutputLicense) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputLicense) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
