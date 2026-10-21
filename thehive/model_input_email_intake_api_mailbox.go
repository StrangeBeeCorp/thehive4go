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

// checks if the InputEmailIntakeApiMailbox type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputEmailIntakeApiMailbox{}

// InputEmailIntakeApiMailbox struct for InputEmailIntakeApiMailbox
type InputEmailIntakeApiMailbox struct {
	Provider   InputEmailIntakeApiProvider `json:"provider"`
	Credential InputEmailIntakeCredential  `json:"credential"`
	// Folder from which emails are fetched. Use `POST /api/v1/connector/email-intake/folders` to list available folders.
	Inbox *string `json:"inbox,omitempty"`
	// Folder where processed emails are moved after fetching. Use `POST /api/v1/connector/email-intake/folders` to list available folders. When omitted, emails aren't moved.
	Archive *string `json:"archive,omitempty"`
	// Marks processed emails as read in the mailbox after fetching.
	MarkAsRead *bool `json:"markAsRead,omitempty"`
	// Discriminator identifying this as an API-based mailbox configuration. Must be `api`.
	Kind *string `json:"_kind,omitempty"`
}

type _InputEmailIntakeApiMailbox InputEmailIntakeApiMailbox

// NewInputEmailIntakeApiMailbox instantiates a new InputEmailIntakeApiMailbox object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputEmailIntakeApiMailbox(provider InputEmailIntakeApiProvider, credential InputEmailIntakeCredential) *InputEmailIntakeApiMailbox {
	this := InputEmailIntakeApiMailbox{}
	this.Provider = provider
	this.Credential = credential
	var inbox string = "Inbox"
	this.Inbox = &inbox
	var markAsRead bool = false
	this.MarkAsRead = &markAsRead
	return &this
}

// NewInputEmailIntakeApiMailboxWithDefaults instantiates a new InputEmailIntakeApiMailbox object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputEmailIntakeApiMailboxWithDefaults() *InputEmailIntakeApiMailbox {
	this := InputEmailIntakeApiMailbox{}
	var inbox string = "Inbox"
	this.Inbox = &inbox
	var markAsRead bool = false
	this.MarkAsRead = &markAsRead
	return &this
}

// GetProvider returns the Provider field value
func (o *InputEmailIntakeApiMailbox) GetProvider() InputEmailIntakeApiProvider {
	if o == nil {
		var ret InputEmailIntakeApiProvider
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *InputEmailIntakeApiMailbox) GetProviderOk() (*InputEmailIntakeApiProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *InputEmailIntakeApiMailbox) SetProvider(v InputEmailIntakeApiProvider) {
	o.Provider = v
}

// GetCredential returns the Credential field value
func (o *InputEmailIntakeApiMailbox) GetCredential() InputEmailIntakeCredential {
	if o == nil {
		var ret InputEmailIntakeCredential
		return ret
	}

	return o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value
// and a boolean to check if the value has been set.
func (o *InputEmailIntakeApiMailbox) GetCredentialOk() (*InputEmailIntakeCredential, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credential, true
}

// SetCredential sets field value
func (o *InputEmailIntakeApiMailbox) SetCredential(v InputEmailIntakeCredential) {
	o.Credential = v
}

// GetInbox returns the Inbox field value if set, zero value otherwise.
func (o *InputEmailIntakeApiMailbox) GetInbox() string {
	if o == nil || IsNil(o.Inbox) {
		var ret string
		return ret
	}
	return *o.Inbox
}

// GetInboxOk returns a tuple with the Inbox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputEmailIntakeApiMailbox) GetInboxOk() (*string, bool) {
	if o == nil || IsNil(o.Inbox) {
		return nil, false
	}
	return o.Inbox, true
}

// HasInbox returns a boolean if a field has been set.
func (o *InputEmailIntakeApiMailbox) HasInbox() bool {
	if o != nil && !IsNil(o.Inbox) {
		return true
	}

	return false
}

// SetInbox gets a reference to the given string and assigns it to the Inbox field.
func (o *InputEmailIntakeApiMailbox) SetInbox(v string) {
	o.Inbox = &v
}

// GetArchive returns the Archive field value if set, zero value otherwise.
func (o *InputEmailIntakeApiMailbox) GetArchive() string {
	if o == nil || IsNil(o.Archive) {
		var ret string
		return ret
	}
	return *o.Archive
}

// GetArchiveOk returns a tuple with the Archive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputEmailIntakeApiMailbox) GetArchiveOk() (*string, bool) {
	if o == nil || IsNil(o.Archive) {
		return nil, false
	}
	return o.Archive, true
}

// HasArchive returns a boolean if a field has been set.
func (o *InputEmailIntakeApiMailbox) HasArchive() bool {
	if o != nil && !IsNil(o.Archive) {
		return true
	}

	return false
}

// SetArchive gets a reference to the given string and assigns it to the Archive field.
func (o *InputEmailIntakeApiMailbox) SetArchive(v string) {
	o.Archive = &v
}

// GetMarkAsRead returns the MarkAsRead field value if set, zero value otherwise.
func (o *InputEmailIntakeApiMailbox) GetMarkAsRead() bool {
	if o == nil || IsNil(o.MarkAsRead) {
		var ret bool
		return ret
	}
	return *o.MarkAsRead
}

// GetMarkAsReadOk returns a tuple with the MarkAsRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputEmailIntakeApiMailbox) GetMarkAsReadOk() (*bool, bool) {
	if o == nil || IsNil(o.MarkAsRead) {
		return nil, false
	}
	return o.MarkAsRead, true
}

// HasMarkAsRead returns a boolean if a field has been set.
func (o *InputEmailIntakeApiMailbox) HasMarkAsRead() bool {
	if o != nil && !IsNil(o.MarkAsRead) {
		return true
	}

	return false
}

// SetMarkAsRead gets a reference to the given bool and assigns it to the MarkAsRead field.
func (o *InputEmailIntakeApiMailbox) SetMarkAsRead(v bool) {
	o.MarkAsRead = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *InputEmailIntakeApiMailbox) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputEmailIntakeApiMailbox) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *InputEmailIntakeApiMailbox) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *InputEmailIntakeApiMailbox) SetKind(v string) {
	o.Kind = &v
}

func (o InputEmailIntakeApiMailbox) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputEmailIntakeApiMailbox) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["provider"] = o.Provider
	toSerialize["credential"] = o.Credential
	if !IsNil(o.Inbox) {
		toSerialize["inbox"] = o.Inbox
	}
	if !IsNil(o.Archive) {
		toSerialize["archive"] = o.Archive
	}
	if !IsNil(o.MarkAsRead) {
		toSerialize["markAsRead"] = o.MarkAsRead
	}
	if !IsNil(o.Kind) {
		toSerialize["_kind"] = o.Kind
	}
	return toSerialize, nil
}

func (o *InputEmailIntakeApiMailbox) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"provider",
		"credential",
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

	varInputEmailIntakeApiMailbox := _InputEmailIntakeApiMailbox{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInputEmailIntakeApiMailbox)

	if err != nil {
		return err
	}

	*o = InputEmailIntakeApiMailbox(varInputEmailIntakeApiMailbox)

	return err
}

type NullableInputEmailIntakeApiMailbox struct {
	value *InputEmailIntakeApiMailbox
	isSet bool
}

func (v NullableInputEmailIntakeApiMailbox) Get() *InputEmailIntakeApiMailbox {
	return v.value
}

func (v *NullableInputEmailIntakeApiMailbox) Set(val *InputEmailIntakeApiMailbox) {
	v.value = val
	v.isSet = true
}

func (v NullableInputEmailIntakeApiMailbox) IsSet() bool {
	return v.isSet
}

func (v *NullableInputEmailIntakeApiMailbox) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputEmailIntakeApiMailbox(val *InputEmailIntakeApiMailbox) *NullableInputEmailIntakeApiMailbox {
	return &NullableInputEmailIntakeApiMailbox{value: val, isSet: true}
}

func (v NullableInputEmailIntakeApiMailbox) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputEmailIntakeApiMailbox) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
