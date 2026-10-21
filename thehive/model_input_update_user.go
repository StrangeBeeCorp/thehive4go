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

// checks if the InputUpdateUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputUpdateUser{}

// InputUpdateUser struct for InputUpdateUser
type InputUpdateUser struct {
	// New display name for the user account.
	Name *string `json:"name,omitempty"`
	// Organization context when updating the profile. Must match an existing organization the user belongs to. Use `POST /api/v1/query` with `listUser` and `extraData: [\"organisations\"]` to retrieve the organizations the user belongs to.
	Organisation *string `json:"organisation,omitempty"`
	// New permission profile for the user in the specified organization. Must match an existing profile. Use `GET /api/v1/profile` to list available profiles.
	Profile *string `json:"profile,omitempty"`
	// Locks or unlocks the user account. A locked account can't sign in or use the API. Requires `manageUser` permission.
	Locked *bool `json:"locked,omitempty"`
	// Avatar image as a raw Base64-encoded string. Always stored as JPEG regardless of the source format. Set to `null` or an empty string to remove the current avatar.
	Avatar *string `json:"avatar,omitempty"`
	// Email address for sending notifications. Set to `null` to remove the current address.
	Email *string `json:"email,omitempty" validate:"regexp=^(([^()\\\\[\\\\]\\\\\\\\.,;:\\\\s@\\"]+(\\\\.[^()\\\\[\\\\]\\\\\\\\.,;:\\\\s@\\"]+)*)|(\\".+\\"))@((\\\\[[0-9]{1,3}\\\\.[0-9]{1,3}\\\\.[0-9]{1,3}\\\\.[0-9]{1,3}])|(([a-zA-Z\\\\-0-9 -퟿豈-﷏ﷰ-￯]+\\\\.)+[a-zA-Z -퟿豈-﷏ﷰ-￯]{2,}))$"`
	// New default organization for the user account. Must match an organization the user already belongs to. Use `POST /api/v1/query` with `listUser` and `extraData: [\"organisations\"]` to retrieve the organizations the user belongs to.
	DefaultOrganisation *string `json:"defaultOrganisation,omitempty"`
	// Type of user account. Normal accounts can use the web interface and the API. Service accounts authenticate exclusively with an API key and can't sign in to the web interface. External accounts access TheHive through the Portal (requires a Platinum or One license). A user's type can't be changed to or from `External` after creation.
	Type *UserType `json:"type,omitempty"`
}

// NewInputUpdateUser instantiates a new InputUpdateUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputUpdateUser() *InputUpdateUser {
	this := InputUpdateUser{}
	return &this
}

// NewInputUpdateUserWithDefaults instantiates a new InputUpdateUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputUpdateUserWithDefaults() *InputUpdateUser {
	this := InputUpdateUser{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InputUpdateUser) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateUser) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InputUpdateUser) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InputUpdateUser) SetName(v string) {
	o.Name = &v
}

// GetOrganisation returns the Organisation field value if set, zero value otherwise.
func (o *InputUpdateUser) GetOrganisation() string {
	if o == nil || IsNil(o.Organisation) {
		var ret string
		return ret
	}
	return *o.Organisation
}

// GetOrganisationOk returns a tuple with the Organisation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateUser) GetOrganisationOk() (*string, bool) {
	if o == nil || IsNil(o.Organisation) {
		return nil, false
	}
	return o.Organisation, true
}

// HasOrganisation returns a boolean if a field has been set.
func (o *InputUpdateUser) HasOrganisation() bool {
	if o != nil && !IsNil(o.Organisation) {
		return true
	}

	return false
}

// SetOrganisation gets a reference to the given string and assigns it to the Organisation field.
func (o *InputUpdateUser) SetOrganisation(v string) {
	o.Organisation = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *InputUpdateUser) GetProfile() string {
	if o == nil || IsNil(o.Profile) {
		var ret string
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateUser) GetProfileOk() (*string, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *InputUpdateUser) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given string and assigns it to the Profile field.
func (o *InputUpdateUser) SetProfile(v string) {
	o.Profile = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *InputUpdateUser) GetLocked() bool {
	if o == nil || IsNil(o.Locked) {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateUser) GetLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *InputUpdateUser) HasLocked() bool {
	if o != nil && !IsNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *InputUpdateUser) SetLocked(v bool) {
	o.Locked = &v
}

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *InputUpdateUser) GetAvatar() string {
	if o == nil || IsNil(o.Avatar) {
		var ret string
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateUser) GetAvatarOk() (*string, bool) {
	if o == nil || IsNil(o.Avatar) {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *InputUpdateUser) HasAvatar() bool {
	if o != nil && !IsNil(o.Avatar) {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given string and assigns it to the Avatar field.
func (o *InputUpdateUser) SetAvatar(v string) {
	o.Avatar = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *InputUpdateUser) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateUser) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *InputUpdateUser) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *InputUpdateUser) SetEmail(v string) {
	o.Email = &v
}

// GetDefaultOrganisation returns the DefaultOrganisation field value if set, zero value otherwise.
func (o *InputUpdateUser) GetDefaultOrganisation() string {
	if o == nil || IsNil(o.DefaultOrganisation) {
		var ret string
		return ret
	}
	return *o.DefaultOrganisation
}

// GetDefaultOrganisationOk returns a tuple with the DefaultOrganisation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateUser) GetDefaultOrganisationOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultOrganisation) {
		return nil, false
	}
	return o.DefaultOrganisation, true
}

// HasDefaultOrganisation returns a boolean if a field has been set.
func (o *InputUpdateUser) HasDefaultOrganisation() bool {
	if o != nil && !IsNil(o.DefaultOrganisation) {
		return true
	}

	return false
}

// SetDefaultOrganisation gets a reference to the given string and assigns it to the DefaultOrganisation field.
func (o *InputUpdateUser) SetDefaultOrganisation(v string) {
	o.DefaultOrganisation = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InputUpdateUser) GetType() UserType {
	if o == nil || IsNil(o.Type) {
		var ret UserType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputUpdateUser) GetTypeOk() (*UserType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InputUpdateUser) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given UserType and assigns it to the Type field.
func (o *InputUpdateUser) SetType(v UserType) {
	o.Type = &v
}

func (o InputUpdateUser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputUpdateUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Organisation) {
		toSerialize["organisation"] = o.Organisation
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.Locked) {
		toSerialize["locked"] = o.Locked
	}
	if !IsNil(o.Avatar) {
		toSerialize["avatar"] = o.Avatar
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.DefaultOrganisation) {
		toSerialize["defaultOrganisation"] = o.DefaultOrganisation
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableInputUpdateUser struct {
	value *InputUpdateUser
	isSet bool
}

func (v NullableInputUpdateUser) Get() *InputUpdateUser {
	return v.value
}

func (v *NullableInputUpdateUser) Set(val *InputUpdateUser) {
	v.value = val
	v.isSet = true
}

func (v NullableInputUpdateUser) IsSet() bool {
	return v.isSet
}

func (v *NullableInputUpdateUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputUpdateUser(val *InputUpdateUser) *NullableInputUpdateUser {
	return &NullableInputUpdateUser{value: val, isSet: true}
}

func (v NullableInputUpdateUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputUpdateUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
