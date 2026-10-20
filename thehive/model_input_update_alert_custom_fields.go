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
	"gopkg.in/validator.v2"
)

// InputUpdateAlertCustomFields -  Custom field values to update on the alert. Must match existing custom fields. Use `GET /api/v1/customField` to list available custom fields.  Several formats are accepted, with different behaviors:  Replacements  - To replace all values of a custom field with a single value — Object format. Overwrites the current values of each listed custom field. Fields not listed are left untouched. `{\"threat-type\": \"Malware\", \"detection-source\": \"EDR\"}` - To replace all values of a custom field with one or more values — Array format. Overwrites the current values of each listed custom field. Fields not listed are deleted. `[{\"name\": \"threat-type\", \"value\": \"Malware\", \"order\": 0}, {\"name\": \"threat-type\", \"value\": \"Ransomware\", \"order\": 1}, {\"name\": \"detection-source\", \"value\": \"EDR\", \"order\": 2}]` - To replace one specific value inside a custom field with one or more values — Object format with a field:value key. Only the targeted value changes. Every other value in that custom field is kept, and fields not listed are left untouched. `{\"threat-type:Malware\":\"Ransomware\"}` or `{\"threat-type:Malware\":[\"Ransomware\", \"Spyware\"]}`  Additions  To add one or more values to a custom field without touching its existing values — Object format with an empty target. The trailing colon means that no value is targeted, so the new value is appended.  `{\"threat-type:\": \"Malware\"}` or `{\"threat-type:\": [\"Malware\", \"Ransomware\"]}`  Deletions  To delete one specific value from a custom field — Object format with an empty replacement. Only the targeted value is removed. Every other value in that field is kept.  `{\"threat-type:Malware\": \"\"}`
type InputUpdateAlertCustomFields struct {
	ArrayOfInputCustomFieldValue *[]InputCustomFieldValue
	MapmapOfStringAny            *map[string]interface{}
}

// []InputCustomFieldValueAsInputUpdateAlertCustomFields is a convenience function that returns []InputCustomFieldValue wrapped in InputUpdateAlertCustomFields
func ArrayOfInputCustomFieldValueAsInputUpdateAlertCustomFields(v *[]InputCustomFieldValue) InputUpdateAlertCustomFields {
	return InputUpdateAlertCustomFields{
		ArrayOfInputCustomFieldValue: v,
	}
}

// map[string]interface{}AsInputUpdateAlertCustomFields is a convenience function that returns map[string]interface{} wrapped in InputUpdateAlertCustomFields
func MapmapOfStringAnyAsInputUpdateAlertCustomFields(v *map[string]interface{}) InputUpdateAlertCustomFields {
	return InputUpdateAlertCustomFields{
		MapmapOfStringAny: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *InputUpdateAlertCustomFields) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfInputCustomFieldValue
	err = newStrictDecoder(data).Decode(&dst.ArrayOfInputCustomFieldValue)
	if err == nil {
		jsonArrayOfInputCustomFieldValue, _ := json.Marshal(dst.ArrayOfInputCustomFieldValue)
		if string(jsonArrayOfInputCustomFieldValue) == "{}" { // empty struct
			dst.ArrayOfInputCustomFieldValue = nil
		} else {
			if err = validator.Validate(dst.ArrayOfInputCustomFieldValue); err != nil {
				dst.ArrayOfInputCustomFieldValue = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfInputCustomFieldValue = nil
	}

	// try to unmarshal data into MapmapOfStringAny
	err = newStrictDecoder(data).Decode(&dst.MapmapOfStringAny)
	if err == nil {
		jsonMapmapOfStringAny, _ := json.Marshal(dst.MapmapOfStringAny)
		if string(jsonMapmapOfStringAny) == "{}" { // empty struct
			dst.MapmapOfStringAny = nil
		} else {
			if err = validator.Validate(dst.MapmapOfStringAny); err != nil {
				dst.MapmapOfStringAny = nil
			} else {
				match++
			}
		}
	} else {
		dst.MapmapOfStringAny = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfInputCustomFieldValue = nil
		dst.MapmapOfStringAny = nil

		return fmt.Errorf("data matches more than one schema in oneOf(InputUpdateAlertCustomFields)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(InputUpdateAlertCustomFields)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src InputUpdateAlertCustomFields) MarshalJSON() ([]byte, error) {
	if src.ArrayOfInputCustomFieldValue != nil {
		return json.Marshal(&src.ArrayOfInputCustomFieldValue)
	}

	if src.MapmapOfStringAny != nil {
		return json.Marshal(&src.MapmapOfStringAny)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *InputUpdateAlertCustomFields) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfInputCustomFieldValue != nil {
		return obj.ArrayOfInputCustomFieldValue
	}

	if obj.MapmapOfStringAny != nil {
		return obj.MapmapOfStringAny
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj InputUpdateAlertCustomFields) GetActualInstanceValue() interface{} {
	if obj.ArrayOfInputCustomFieldValue != nil {
		return *obj.ArrayOfInputCustomFieldValue
	}

	if obj.MapmapOfStringAny != nil {
		return *obj.MapmapOfStringAny
	}

	// all schemas are nil
	return nil
}

type NullableInputUpdateAlertCustomFields struct {
	value *InputUpdateAlertCustomFields
	isSet bool
}

func (v NullableInputUpdateAlertCustomFields) Get() *InputUpdateAlertCustomFields {
	return v.value
}

func (v *NullableInputUpdateAlertCustomFields) Set(val *InputUpdateAlertCustomFields) {
	v.value = val
	v.isSet = true
}

func (v NullableInputUpdateAlertCustomFields) IsSet() bool {
	return v.isSet
}

func (v *NullableInputUpdateAlertCustomFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputUpdateAlertCustomFields(val *InputUpdateAlertCustomFields) *NullableInputUpdateAlertCustomFields {
	return &NullableInputUpdateAlertCustomFields{value: val, isSet: true}
}

func (v NullableInputUpdateAlertCustomFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputUpdateAlertCustomFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
