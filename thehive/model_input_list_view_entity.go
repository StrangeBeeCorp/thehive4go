/*
TheHive

 ## Get started  This guide covers everything you need to make your first successful API call to TheHive.  The publicly hosted API documentation reflects the latest version of TheHive. To browse the documentation that matches the version of TheHive you are running, open it from within the application: select the **?** button in the top right corner, then **API Docs**.  The v0 API is deprecated and isn't documented. New integrations should target v1.  ### Prerequisites  You need a running TheHive instance and an unlocked user account.  ### Authentication  Almost all endpoints require authentication.  TheHive supports two authentication methods:  #### API key (recommended)  Use your API key as a bearer token in the `Authorization` header.  Generate a new API key with `POST /api/v1/user/{userId}/key/renew`, or retrieve an existing one with `GET /api/v1/user/{userId}/key`.  If you have a [Normal user account](https://docs.strangebee.com/thehive/user-guides/organization/configure-organization/manage-user-accounts/about-user-accounts/#types), you can also manage your API key in your [user settings](https://docs.strangebee.com/thehive/user-guides/manage-user-settings/#manage-your-api-key).  curl ```bash curl -X <http_method> -H 'Authorization: Bearer <api_key>' \\   https://<thehive_host>/api/v1/<endpoint> ```  python ```python import requests  headers = {\"Authorization\": \"Bearer <api_key>\"} response = requests.<http_method>(\"https://<thehive_host>/api/v1/<endpoint>\", headers=headers) ```  #### Basic authentication  Only available for [Normal user accounts](https://docs.strangebee.com/thehive/user-guides/organization/configure-organization/manage-user-accounts/about-user-accounts/#types).  Pass your login and password using HTTP Basic authentication. The login is typically an email address. If a default domain is configured in the [authentication settings](https://docs.strangebee.com/thehive/administration/authentication/configure-authentication), a plain username is also accepted.  curl ```bash curl -X <http_method> -u <login>:<password> \\   https://<thehive_host>/api/v1/<endpoint> ```  python ```python import requests  response = requests.<http_method>(\"https://<thehive_host>/api/v1/<endpoint>\", auth=(\"<login>\", \"<password>\")) ```  The following endpoints are public and can be called without credentials:  - `POST /api/v1/login`: Authenticate with login, password, and a TOTP code if MFA is enabled. Returns a session cookie used to authenticate subsequent requests. - `GET /api/v1/logout`: Invalidate the current session. - `POST /api/v1/logout`: Invalidate the current session. Same as `GET`, provided for clients that can't send `GET` requests with a body. - `GET /api/v1/auth/local/passwordPolicy`: Retrieve the password policy configured on the server such as minimum length or complexity requirements. - `GET /api/v1/status/public`: Retrieve public server information, including the TheHive version and license status, without exposing sensitive configuration.  All other endpoints return `401 Unauthorized` when called without valid credentials.  ### Your first API call  Start with `GET /api/v1/user/current`: it requires no request body, confirms your credentials work, and returns your user ID.  curl ```bash curl -X GET -H 'Authorization: Bearer <api_key>' \\      https://<thehive_host>/api/v1/user/current ```  python ```python import requests  BASE_URL = \"https://<thehive_host>\" API_KEY = \"<api_key>\"  headers = {\"Authorization\": f\"Bearer {API_KEY}\"} response = requests.get(f\"{BASE_URL}/api/v1/user/current\", headers=headers) user = response.json() ```  ### X-Organisation header  By default, API calls target the user's default organization. The `defaultOrganisation` field returned by `GET /api/v1/user/current` tells you which one it is.  If your account belongs to several organizations and you need to target a specific one, add the `X-Organisation` header. This is also required when performing administration actions, as these must be executed in the context of the `admin` organization.  curl ```bash curl -H 'Authorization: Bearer <api_key>' \\      -H 'X-Organisation: <organization_name>' \\      https://<thehive_host>/api/v1/user/current ```  python ```python import requests  BASE_URL = \"https://<thehive_host>\" API_KEY = \"<api_key>\"  headers = {\"Authorization\": f\"Bearer {API_KEY}\", \"X-Organisation\": \"<organization_name>\"} response = requests.get(f\"{BASE_URL}/api/v1/user/current\", headers=headers) user = response.json() ```  ### Request and response format  All request and response bodies use JSON: `Content-Type: application/json`.  Dates are represented as Unix timestamps in milliseconds.  ### How the API is organized  TheHive exposes two kinds of endpoints.  #### The Query API  `POST /api/v1/query` handles all listing, filtering, sorting, and pagination needs.  A single endpoint accepts a structured body that chains operations:  ```json {   \"query\": [     { \"_name\": \"listCase\" },     { \"_name\": \"filter\", \"_eq\": { \"_field\": \"status\", \"_value\": \"Open\" } },     { \"_name\": \"sort\", \"_fields\": [{ \"_createdAt\": \"desc\" }] },     { \"_name\": \"page\", \"from\": 0, \"to\": 15, \"extraData\": [\"taskStats\", \"shareCount\"] }   ],   \"excludeFields\": [\"description\", \"summary\"] } ```  See the [**Query and Export**](https://docs.strangebee.com/thehive/api-docs/#tag/Query-and-Export) section for the full query reference.  #### Standard REST endpoints  Standard REST endpoints handle write operations and direct entity lookups.  They follow conventional HTTP conventions:  | <nobr>Operation </nobr>| Method | Example | |-----------|--------|---------| | <nobr>Create </nobr>| `POST` | `POST /api/v1/case` | | <nobr>Read </nobr>| `GET` | `GET /api/v1/case/{idOrName}` | | <nobr>Update </nobr>| `PATCH` | `PATCH /api/v1/case/{idOrName}` | | <nobr>Replace / set </nobr>| `PUT` | `PUT /api/v1/case/{caseId}/shares` | | <nobr>Delete </nobr>| `DELETE` | `DELETE /api/v1/case/{idOrName}` |  ### Permissions and capabilities  Each user account has permissions defined by their profile. TheHive checks them on every API call.  Some endpoints are only available with a Gold, Platinum, or One license. When that's the case, the required license tier is listed on the endpoint.  To check which license is active on your instance and what capabilities it includes, call `GET /api/v1/status`.  ### Request size limits  TheHive enforces the following request size limits:  | <nobr>Request type </nobr>| Default limit | Description | |---|---|---| | <nobr>Standard API requests (JSON payloads) </nobr>| 10 MB | Maximum size for request bodies processed in memory | | <nobr>File uploads and attachments </nobr>| 1 GB | Maximum size for streamed request bodies written to disk |  TheHive rejects requests that exceed these limits.  To stay within these limits: - Avoid sending excessively large JSON payloads. - Upload large files as attachments rather than embedding data directly in requests. - Implement rate limiting and additional protections at the reverse proxy level when exposing the API externally.  ### Errors  Every endpoint returns at least the following HTTP status codes:  | <nobr>Code </nobr>| Meaning | |------|---------| | <nobr>`400 Bad Request` </nobr>| Invalid request parameters or body | | <nobr>`404 Not Found` </nobr>| Resource does not exist | | <nobr>`500 Internal Server Error` </nobr>| Unexpected server-side error |  Endpoints that require authentication additionally return:  | <nobr>Code </nobr>| Meaning | |------|---------| | <nobr>`401 Unauthorized` </nobr>| Missing or invalid credentials | | <nobr>`403 Forbidden` </nobr>| Authenticated but insufficient permissions |  Other codes may appear depending on the specific operation. Refer to individual endpoint documentation for the full list.  The standard error response body is:  ```json {   \"type\": \"AuthorizationError\",   \"message\": \"You don't have permission to perform this action\" } ```  A small number of endpoints deviate from this format. These cases are noted in the relevant endpoint descriptions.

API version: v6.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package thehive

import (
	"encoding/json"
	"fmt"
)

// InputListViewEntity Type of list the view applies to. Determines which list the view's filters, sort order, displayed columns, and display options apply to.
type InputListViewEntity string

// List of InputListViewEntity
const (
	INPUTLISTVIEWENTITY_ALERT                        InputListViewEntity = "Alert"
	INPUTLISTVIEWENTITY_ALERT_OBSERVABLE             InputListViewEntity = "AlertObservable"
	INPUTLISTVIEWENTITY_ALERT_TTP                    InputListViewEntity = "AlertTTP"
	INPUTLISTVIEWENTITY_ALERT_HISTORY                InputListViewEntity = "AlertHistory"
	INPUTLISTVIEWENTITY_ALERT_SIMILAR_CASES          InputListViewEntity = "AlertSimilarCases"
	INPUTLISTVIEWENTITY_ALERT_SIMILAR_ALERTS         InputListViewEntity = "AlertSimilarAlerts"
	INPUTLISTVIEWENTITY_CASE                         InputListViewEntity = "Case"
	INPUTLISTVIEWENTITY_CASE_TASK                    InputListViewEntity = "CaseTask"
	INPUTLISTVIEWENTITY_CASE_OBSERVABLE              InputListViewEntity = "CaseObservable"
	INPUTLISTVIEWENTITY_CASE_TTP                     InputListViewEntity = "CaseTTP"
	INPUTLISTVIEWENTITY_CASE_HISTORY                 InputListViewEntity = "CaseHistory"
	INPUTLISTVIEWENTITY_CASE_SIMILAR_CASES           InputListViewEntity = "CaseSimilarCases"
	INPUTLISTVIEWENTITY_CASE_SIMILAR_ALERTS          InputListViewEntity = "CaseSimilarAlerts"
	INPUTLISTVIEWENTITY_CASE_SIMILAR_LINKED_ALERTS   InputListViewEntity = "CaseSimilarLinkedAlerts"
	INPUTLISTVIEWENTITY_DASHBOARD                    InputListViewEntity = "Dashboard"
	INPUTLISTVIEWENTITY_ORGANISATION                 InputListViewEntity = "Organisation"
	INPUTLISTVIEWENTITY_ORGANISATION_USERS           InputListViewEntity = "OrganisationUsers"
	INPUTLISTVIEWENTITY_ORGANISATION_TEMPLATE_CASE   InputListViewEntity = "OrganisationTemplateCase"
	INPUTLISTVIEWENTITY_ORGANISATION_TEMPLATE_REPORT InputListViewEntity = "OrganisationTemplateReport"
	INPUTLISTVIEWENTITY_ORGANISATION_TEMPLATE_PAGES  InputListViewEntity = "OrganisationTemplatePages"
	INPUTLISTVIEWENTITY_ORGANISATION_CUSTOM_TAGS     InputListViewEntity = "OrganisationCustomTags"
	INPUTLISTVIEWENTITY_ORGANISATION_FUNCTIONS       InputListViewEntity = "OrganisationFunctions"
	INPUTLISTVIEWENTITY_USER                         InputListViewEntity = "User"
	INPUTLISTVIEWENTITY_PROFILES                     InputListViewEntity = "Profiles"
	INPUTLISTVIEWENTITY_CUSTOM_FIELDS                InputListViewEntity = "CustomFields"
	INPUTLISTVIEWENTITY_OBSERVABLE_TYPES             InputListViewEntity = "ObservableTypes"
	INPUTLISTVIEWENTITY_TAXONOMIES                   InputListViewEntity = "Taxonomies"
	INPUTLISTVIEWENTITY_TASK                         InputListViewEntity = "Task"
	INPUTLISTVIEWENTITY_ATTACK_PATTERNS              InputListViewEntity = "AttackPatterns"
)

// All allowed values of InputListViewEntity enum
var AllowedInputListViewEntityEnumValues = []InputListViewEntity{
	"Alert",
	"AlertObservable",
	"AlertTTP",
	"AlertHistory",
	"AlertSimilarCases",
	"AlertSimilarAlerts",
	"Case",
	"CaseTask",
	"CaseObservable",
	"CaseTTP",
	"CaseHistory",
	"CaseSimilarCases",
	"CaseSimilarAlerts",
	"CaseSimilarLinkedAlerts",
	"Dashboard",
	"Organisation",
	"OrganisationUsers",
	"OrganisationTemplateCase",
	"OrganisationTemplateReport",
	"OrganisationTemplatePages",
	"OrganisationCustomTags",
	"OrganisationFunctions",
	"User",
	"Profiles",
	"CustomFields",
	"ObservableTypes",
	"Taxonomies",
	"Task",
	"AttackPatterns",
}

func (v *InputListViewEntity) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InputListViewEntity(value)
	for _, existing := range AllowedInputListViewEntityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InputListViewEntity", value)
}

// NewInputListViewEntityFromValue returns a pointer to a valid InputListViewEntity
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInputListViewEntityFromValue(v string) (*InputListViewEntity, error) {
	ev := InputListViewEntity(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InputListViewEntity: valid values are %v", v, AllowedInputListViewEntityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InputListViewEntity) IsValid() bool {
	for _, existing := range AllowedInputListViewEntityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InputListViewEntity value
func (v InputListViewEntity) Ptr() *InputListViewEntity {
	return &v
}

type NullableInputListViewEntity struct {
	value *InputListViewEntity
	isSet bool
}

func (v NullableInputListViewEntity) Get() *InputListViewEntity {
	return v.value
}

func (v *NullableInputListViewEntity) Set(val *InputListViewEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableInputListViewEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableInputListViewEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputListViewEntity(val *InputListViewEntity) *NullableInputListViewEntity {
	return &NullableInputListViewEntity{value: val, isSet: true}
}

func (v NullableInputListViewEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputListViewEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
