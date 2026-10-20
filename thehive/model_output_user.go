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

// checks if the OutputUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutputUser{}

// OutputUser struct for OutputUser
type OutputUser struct {
	// Internal unique identifier of the user.
	UnderscoreId string `json:"_id"`
	// Login of the user who created the user account.
	UnderscoreCreatedBy string `json:"_createdBy"`
	// Login of the user who last updated the user account.
	UnderscoreUpdatedBy *string `json:"_updatedBy,omitempty"`
	// Date and time when the user account was created. Returned as a Unix timestamp in milliseconds.
	UnderscoreCreatedAt int64 `json:"_createdAt"`
	// Date and time when the user account was last updated. Returned as a Unix timestamp in milliseconds.
	UnderscoreUpdatedAt *int64 `json:"_updatedAt,omitempty"`
	// Login name for the user account. Always a valid email address.
	Login string `json:"login"`
	// Display name for the user account.
	Name string `json:"name"`
	// Email address of the user for notifications.
	Email *string `json:"email,omitempty"`
	// Whether the user account has an API key.
	HasKey bool `json:"hasKey"`
	// Whether the user account has a password set.
	HasPassword bool `json:"hasPassword"`
	// Whether multi-factor authentication (MFA) is enabled for the user account.
	HasMFA bool `json:"hasMFA"`
	// Whether the user account is locked. A locked account can't sign in or use the API.
	Locked bool `json:"locked"`
	// Permission profile of the user in the current organization.
	Profile string `json:"profile"`
	// Set of permissions granted to the user in the current organization.
	Permissions []string `json:"permissions,omitempty"`
	// Name of the organization in the context of which the user was retrieved.
	Organisation string `json:"organisation"`
	// Relative path to the user's avatar image.
	Avatar *string `json:"avatar,omitempty"`
	// Organizations the user belongs to, with the profile and avatar used in each one.
	Organisations []OutputOrganisationProfile `json:"organisations,omitempty"`
	// Type of user account.
	Type UserType `json:"type"`
	// Default organization used when the user signs in.
	DefaultOrganisation *string `json:"defaultOrganisation,omitempty"`
	// Extra data fields populated when specific keys are requested in the `extraData` query parameter.
	ExtraData map[string]interface{} `json:"extraData"`
}

type _OutputUser OutputUser

// NewOutputUser instantiates a new OutputUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputUser(underscoreId string, underscoreCreatedBy string, underscoreCreatedAt int64, login string, name string, hasKey bool, hasPassword bool, hasMFA bool, locked bool, profile string, organisation string, type_ UserType, extraData map[string]interface{}) *OutputUser {
	this := OutputUser{}
	this.UnderscoreId = underscoreId
	this.UnderscoreCreatedBy = underscoreCreatedBy
	this.UnderscoreCreatedAt = underscoreCreatedAt
	this.Login = login
	this.Name = name
	this.HasKey = hasKey
	this.HasPassword = hasPassword
	this.HasMFA = hasMFA
	this.Locked = locked
	this.Profile = profile
	this.Organisation = organisation
	this.Type = type_
	this.ExtraData = extraData
	return &this
}

// NewOutputUserWithDefaults instantiates a new OutputUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputUserWithDefaults() *OutputUser {
	this := OutputUser{}
	return &this
}

// GetUnderscoreId returns the UnderscoreId field value
func (o *OutputUser) GetUnderscoreId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnderscoreId
}

// GetUnderscoreIdOk returns a tuple with the UnderscoreId field value
// and a boolean to check if the value has been set.
func (o *OutputUser) GetUnderscoreIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreId, true
}

// SetUnderscoreId sets field value
func (o *OutputUser) SetUnderscoreId(v string) {
	o.UnderscoreId = v
}

// GetUnderscoreCreatedBy returns the UnderscoreCreatedBy field value
func (o *OutputUser) GetUnderscoreCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnderscoreCreatedBy
}

// GetUnderscoreCreatedByOk returns a tuple with the UnderscoreCreatedBy field value
// and a boolean to check if the value has been set.
func (o *OutputUser) GetUnderscoreCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreCreatedBy, true
}

// SetUnderscoreCreatedBy sets field value
func (o *OutputUser) SetUnderscoreCreatedBy(v string) {
	o.UnderscoreCreatedBy = v
}

// GetUnderscoreUpdatedBy returns the UnderscoreUpdatedBy field value if set, zero value otherwise.
func (o *OutputUser) GetUnderscoreUpdatedBy() string {
	if o == nil || IsNil(o.UnderscoreUpdatedBy) {
		var ret string
		return ret
	}
	return *o.UnderscoreUpdatedBy
}

// GetUnderscoreUpdatedByOk returns a tuple with the UnderscoreUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputUser) GetUnderscoreUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UnderscoreUpdatedBy) {
		return nil, false
	}
	return o.UnderscoreUpdatedBy, true
}

// HasUnderscoreUpdatedBy returns a boolean if a field has been set.
func (o *OutputUser) HasUnderscoreUpdatedBy() bool {
	if o != nil && !IsNil(o.UnderscoreUpdatedBy) {
		return true
	}

	return false
}

// SetUnderscoreUpdatedBy gets a reference to the given string and assigns it to the UnderscoreUpdatedBy field.
func (o *OutputUser) SetUnderscoreUpdatedBy(v string) {
	o.UnderscoreUpdatedBy = &v
}

// GetUnderscoreCreatedAt returns the UnderscoreCreatedAt field value
func (o *OutputUser) GetUnderscoreCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UnderscoreCreatedAt
}

// GetUnderscoreCreatedAtOk returns a tuple with the UnderscoreCreatedAt field value
// and a boolean to check if the value has been set.
func (o *OutputUser) GetUnderscoreCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderscoreCreatedAt, true
}

// SetUnderscoreCreatedAt sets field value
func (o *OutputUser) SetUnderscoreCreatedAt(v int64) {
	o.UnderscoreCreatedAt = v
}

// GetUnderscoreUpdatedAt returns the UnderscoreUpdatedAt field value if set, zero value otherwise.
func (o *OutputUser) GetUnderscoreUpdatedAt() int64 {
	if o == nil || IsNil(o.UnderscoreUpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UnderscoreUpdatedAt
}

// GetUnderscoreUpdatedAtOk returns a tuple with the UnderscoreUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputUser) GetUnderscoreUpdatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.UnderscoreUpdatedAt) {
		return nil, false
	}
	return o.UnderscoreUpdatedAt, true
}

// HasUnderscoreUpdatedAt returns a boolean if a field has been set.
func (o *OutputUser) HasUnderscoreUpdatedAt() bool {
	if o != nil && !IsNil(o.UnderscoreUpdatedAt) {
		return true
	}

	return false
}

// SetUnderscoreUpdatedAt gets a reference to the given int64 and assigns it to the UnderscoreUpdatedAt field.
func (o *OutputUser) SetUnderscoreUpdatedAt(v int64) {
	o.UnderscoreUpdatedAt = &v
}

// GetLogin returns the Login field value
func (o *OutputUser) GetLogin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Login
}

// GetLoginOk returns a tuple with the Login field value
// and a boolean to check if the value has been set.
func (o *OutputUser) GetLoginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Login, true
}

// SetLogin sets field value
func (o *OutputUser) SetLogin(v string) {
	o.Login = v
}

// GetName returns the Name field value
func (o *OutputUser) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OutputUser) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OutputUser) SetName(v string) {
	o.Name = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *OutputUser) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputUser) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *OutputUser) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *OutputUser) SetEmail(v string) {
	o.Email = &v
}

// GetHasKey returns the HasKey field value
func (o *OutputUser) GetHasKey() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasKey
}

// GetHasKeyOk returns a tuple with the HasKey field value
// and a boolean to check if the value has been set.
func (o *OutputUser) GetHasKeyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasKey, true
}

// SetHasKey sets field value
func (o *OutputUser) SetHasKey(v bool) {
	o.HasKey = v
}

// GetHasPassword returns the HasPassword field value
func (o *OutputUser) GetHasPassword() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasPassword
}

// GetHasPasswordOk returns a tuple with the HasPassword field value
// and a boolean to check if the value has been set.
func (o *OutputUser) GetHasPasswordOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasPassword, true
}

// SetHasPassword sets field value
func (o *OutputUser) SetHasPassword(v bool) {
	o.HasPassword = v
}

// GetHasMFA returns the HasMFA field value
func (o *OutputUser) GetHasMFA() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMFA
}

// GetHasMFAOk returns a tuple with the HasMFA field value
// and a boolean to check if the value has been set.
func (o *OutputUser) GetHasMFAOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMFA, true
}

// SetHasMFA sets field value
func (o *OutputUser) SetHasMFA(v bool) {
	o.HasMFA = v
}

// GetLocked returns the Locked field value
func (o *OutputUser) GetLocked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Locked
}

// GetLockedOk returns a tuple with the Locked field value
// and a boolean to check if the value has been set.
func (o *OutputUser) GetLockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locked, true
}

// SetLocked sets field value
func (o *OutputUser) SetLocked(v bool) {
	o.Locked = v
}

// GetProfile returns the Profile field value
func (o *OutputUser) GetProfile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *OutputUser) GetProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *OutputUser) SetProfile(v string) {
	o.Profile = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *OutputUser) GetPermissions() []string {
	if o == nil || IsNil(o.Permissions) {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputUser) GetPermissionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *OutputUser) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *OutputUser) SetPermissions(v []string) {
	o.Permissions = v
}

// GetOrganisation returns the Organisation field value
func (o *OutputUser) GetOrganisation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Organisation
}

// GetOrganisationOk returns a tuple with the Organisation field value
// and a boolean to check if the value has been set.
func (o *OutputUser) GetOrganisationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Organisation, true
}

// SetOrganisation sets field value
func (o *OutputUser) SetOrganisation(v string) {
	o.Organisation = v
}

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *OutputUser) GetAvatar() string {
	if o == nil || IsNil(o.Avatar) {
		var ret string
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputUser) GetAvatarOk() (*string, bool) {
	if o == nil || IsNil(o.Avatar) {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *OutputUser) HasAvatar() bool {
	if o != nil && !IsNil(o.Avatar) {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given string and assigns it to the Avatar field.
func (o *OutputUser) SetAvatar(v string) {
	o.Avatar = &v
}

// GetOrganisations returns the Organisations field value if set, zero value otherwise.
func (o *OutputUser) GetOrganisations() []OutputOrganisationProfile {
	if o == nil || IsNil(o.Organisations) {
		var ret []OutputOrganisationProfile
		return ret
	}
	return o.Organisations
}

// GetOrganisationsOk returns a tuple with the Organisations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputUser) GetOrganisationsOk() ([]OutputOrganisationProfile, bool) {
	if o == nil || IsNil(o.Organisations) {
		return nil, false
	}
	return o.Organisations, true
}

// HasOrganisations returns a boolean if a field has been set.
func (o *OutputUser) HasOrganisations() bool {
	if o != nil && !IsNil(o.Organisations) {
		return true
	}

	return false
}

// SetOrganisations gets a reference to the given []OutputOrganisationProfile and assigns it to the Organisations field.
func (o *OutputUser) SetOrganisations(v []OutputOrganisationProfile) {
	o.Organisations = v
}

// GetType returns the Type field value
func (o *OutputUser) GetType() UserType {
	if o == nil {
		var ret UserType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OutputUser) GetTypeOk() (*UserType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OutputUser) SetType(v UserType) {
	o.Type = v
}

// GetDefaultOrganisation returns the DefaultOrganisation field value if set, zero value otherwise.
func (o *OutputUser) GetDefaultOrganisation() string {
	if o == nil || IsNil(o.DefaultOrganisation) {
		var ret string
		return ret
	}
	return *o.DefaultOrganisation
}

// GetDefaultOrganisationOk returns a tuple with the DefaultOrganisation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputUser) GetDefaultOrganisationOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultOrganisation) {
		return nil, false
	}
	return o.DefaultOrganisation, true
}

// HasDefaultOrganisation returns a boolean if a field has been set.
func (o *OutputUser) HasDefaultOrganisation() bool {
	if o != nil && !IsNil(o.DefaultOrganisation) {
		return true
	}

	return false
}

// SetDefaultOrganisation gets a reference to the given string and assigns it to the DefaultOrganisation field.
func (o *OutputUser) SetDefaultOrganisation(v string) {
	o.DefaultOrganisation = &v
}

// GetExtraData returns the ExtraData field value
func (o *OutputUser) GetExtraData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ExtraData
}

// GetExtraDataOk returns a tuple with the ExtraData field value
// and a boolean to check if the value has been set.
func (o *OutputUser) GetExtraDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.ExtraData, true
}

// SetExtraData sets field value
func (o *OutputUser) SetExtraData(v map[string]interface{}) {
	o.ExtraData = v
}

func (o OutputUser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutputUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.UnderscoreId
	toSerialize["_createdBy"] = o.UnderscoreCreatedBy
	if !IsNil(o.UnderscoreUpdatedBy) {
		toSerialize["_updatedBy"] = o.UnderscoreUpdatedBy
	}
	toSerialize["_createdAt"] = o.UnderscoreCreatedAt
	if !IsNil(o.UnderscoreUpdatedAt) {
		toSerialize["_updatedAt"] = o.UnderscoreUpdatedAt
	}
	toSerialize["login"] = o.Login
	toSerialize["name"] = o.Name
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	toSerialize["hasKey"] = o.HasKey
	toSerialize["hasPassword"] = o.HasPassword
	toSerialize["hasMFA"] = o.HasMFA
	toSerialize["locked"] = o.Locked
	toSerialize["profile"] = o.Profile
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	toSerialize["organisation"] = o.Organisation
	if !IsNil(o.Avatar) {
		toSerialize["avatar"] = o.Avatar
	}
	if !IsNil(o.Organisations) {
		toSerialize["organisations"] = o.Organisations
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.DefaultOrganisation) {
		toSerialize["defaultOrganisation"] = o.DefaultOrganisation
	}
	toSerialize["extraData"] = o.ExtraData
	return toSerialize, nil
}

func (o *OutputUser) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"_createdBy",
		"_createdAt",
		"login",
		"name",
		"hasKey",
		"hasPassword",
		"hasMFA",
		"locked",
		"profile",
		"organisation",
		"type",
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

	varOutputUser := _OutputUser{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// Postprocessed by scripts/tolerant-response-decoder: unknown fields are ignored,
	// since TheHive adds fields to its responses in minor versions.
	err = decoder.Decode(&varOutputUser)

	if err != nil {
		return err
	}

	*o = OutputUser(varOutputUser)

	return err
}

type NullableOutputUser struct {
	value *OutputUser
	isSet bool
}

func (v NullableOutputUser) Get() *OutputUser {
	return v.value
}

func (v *NullableOutputUser) Set(val *OutputUser) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputUser) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputUser(val *OutputUser) *NullableOutputUser {
	return &NullableOutputUser{value: val, isSet: true}
}

func (v NullableOutputUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
