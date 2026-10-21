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

// InputCreateObservableAttachment - Attachment for file-type observables. Required when the observable type has `isAttachment` set to `true`.  Four formats are accepted: - String — name of a multipart part containing the file to upload: `\"malware_sample\"` - String array — names of multiple multipart parts: `[\"malware_sample\", \"config_file\"]` - Object — reference to an existing attachment stored in TheHive: `{\"name\": \"malware_sample.exe\", \"contentType\": \"application/octet-stream\", \"id\": \"~128458762\"}` - Object array — references to multiple existing attachments: `[{\"name\": \"malware_sample.exe\", \"contentType\": \"application/octet-stream\", \"id\": \"~128458762\"}]`
type InputCreateObservableAttachment struct {
	InputAttachment                                  *InputAttachment
	ArrayOfInputCreateObservableAttachmentOneOfInner *[]InputCreateObservableAttachmentOneOfInner
	ArrayOfString                                    *[]string
	String                                           *string
}

// InputAttachmentAsInputCreateObservableAttachment is a convenience function that returns InputAttachment wrapped in InputCreateObservableAttachment
func InputAttachmentAsInputCreateObservableAttachment(v *InputAttachment) InputCreateObservableAttachment {
	return InputCreateObservableAttachment{
		InputAttachment: v,
	}
}

// []InputCreateObservableAttachmentOneOfInnerAsInputCreateObservableAttachment is a convenience function that returns []InputCreateObservableAttachmentOneOfInner wrapped in InputCreateObservableAttachment
func ArrayOfInputCreateObservableAttachmentOneOfInnerAsInputCreateObservableAttachment(v *[]InputCreateObservableAttachmentOneOfInner) InputCreateObservableAttachment {
	return InputCreateObservableAttachment{
		ArrayOfInputCreateObservableAttachmentOneOfInner: v,
	}
}

// []stringAsInputCreateObservableAttachment is a convenience function that returns []string wrapped in InputCreateObservableAttachment
func ArrayOfStringAsInputCreateObservableAttachment(v *[]string) InputCreateObservableAttachment {
	return InputCreateObservableAttachment{
		ArrayOfString: v,
	}
}

// stringAsInputCreateObservableAttachment is a convenience function that returns string wrapped in InputCreateObservableAttachment
func StringAsInputCreateObservableAttachment(v *string) InputCreateObservableAttachment {
	return InputCreateObservableAttachment{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *InputCreateObservableAttachment) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into InputAttachment
	err = newStrictDecoder(data).Decode(&dst.InputAttachment)
	if err == nil {
		jsonInputAttachment, _ := json.Marshal(dst.InputAttachment)
		if string(jsonInputAttachment) == "{}" { // empty struct
			dst.InputAttachment = nil
		} else {
			if err = validator.Validate(dst.InputAttachment); err != nil {
				dst.InputAttachment = nil
			} else {
				match++
			}
		}
	} else {
		dst.InputAttachment = nil
	}

	// try to unmarshal data into ArrayOfInputCreateObservableAttachmentOneOfInner
	err = newStrictDecoder(data).Decode(&dst.ArrayOfInputCreateObservableAttachmentOneOfInner)
	if err == nil {
		jsonArrayOfInputCreateObservableAttachmentOneOfInner, _ := json.Marshal(dst.ArrayOfInputCreateObservableAttachmentOneOfInner)
		if string(jsonArrayOfInputCreateObservableAttachmentOneOfInner) == "{}" { // empty struct
			dst.ArrayOfInputCreateObservableAttachmentOneOfInner = nil
		} else {
			if err = validator.Validate(dst.ArrayOfInputCreateObservableAttachmentOneOfInner); err != nil {
				dst.ArrayOfInputCreateObservableAttachmentOneOfInner = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfInputCreateObservableAttachmentOneOfInner = nil
	}

	// try to unmarshal data into ArrayOfString
	err = newStrictDecoder(data).Decode(&dst.ArrayOfString)
	if err == nil {
		jsonArrayOfString, _ := json.Marshal(dst.ArrayOfString)
		if string(jsonArrayOfString) == "{}" { // empty struct
			dst.ArrayOfString = nil
		} else {
			if err = validator.Validate(dst.ArrayOfString); err != nil {
				dst.ArrayOfString = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfString = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			if err = validator.Validate(dst.String); err != nil {
				dst.String = nil
			} else {
				match++
			}
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.InputAttachment = nil
		dst.ArrayOfInputCreateObservableAttachmentOneOfInner = nil
		dst.ArrayOfString = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(InputCreateObservableAttachment)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(InputCreateObservableAttachment)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src InputCreateObservableAttachment) MarshalJSON() ([]byte, error) {
	if src.InputAttachment != nil {
		return json.Marshal(&src.InputAttachment)
	}

	if src.ArrayOfInputCreateObservableAttachmentOneOfInner != nil {
		return json.Marshal(&src.ArrayOfInputCreateObservableAttachmentOneOfInner)
	}

	if src.ArrayOfString != nil {
		return json.Marshal(&src.ArrayOfString)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *InputCreateObservableAttachment) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.InputAttachment != nil {
		return obj.InputAttachment
	}

	if obj.ArrayOfInputCreateObservableAttachmentOneOfInner != nil {
		return obj.ArrayOfInputCreateObservableAttachmentOneOfInner
	}

	if obj.ArrayOfString != nil {
		return obj.ArrayOfString
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj InputCreateObservableAttachment) GetActualInstanceValue() interface{} {
	if obj.InputAttachment != nil {
		return *obj.InputAttachment
	}

	if obj.ArrayOfInputCreateObservableAttachmentOneOfInner != nil {
		return *obj.ArrayOfInputCreateObservableAttachmentOneOfInner
	}

	if obj.ArrayOfString != nil {
		return *obj.ArrayOfString
	}

	if obj.String != nil {
		return *obj.String
	}

	// all schemas are nil
	return nil
}

type NullableInputCreateObservableAttachment struct {
	value *InputCreateObservableAttachment
	isSet bool
}

func (v NullableInputCreateObservableAttachment) Get() *InputCreateObservableAttachment {
	return v.value
}

func (v *NullableInputCreateObservableAttachment) Set(val *InputCreateObservableAttachment) {
	v.value = val
	v.isSet = true
}

func (v NullableInputCreateObservableAttachment) IsSet() bool {
	return v.isSet
}

func (v *NullableInputCreateObservableAttachment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputCreateObservableAttachment(val *InputCreateObservableAttachment) *NullableInputCreateObservableAttachment {
	return &NullableInputCreateObservableAttachment{value: val, isSet: true}
}

func (v NullableInputCreateObservableAttachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputCreateObservableAttachment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
