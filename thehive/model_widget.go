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

// Widget - struct for Widget
type Widget struct {
	AlertList        *AlertList
	AlertTable       *AlertTable
	Comments         *Comments
	CustomFields     *CustomFields
	CustomFieldsList *CustomFieldsList
	Image            *Image
	ObservableList   *ObservableList
	ObservableTable  *ObservableTable
	Pages            *Pages
	TTPList          *TTPList
	TTPTable         *TTPTable
	TaskList         *TaskList
	TaskTable        *TaskTable
	Text             *Text
	Timeline         *Timeline
}

// AlertListAsWidget is a convenience function that returns AlertList wrapped in Widget
func AlertListAsWidget(v *AlertList) Widget {
	return Widget{
		AlertList: v,
	}
}

// AlertTableAsWidget is a convenience function that returns AlertTable wrapped in Widget
func AlertTableAsWidget(v *AlertTable) Widget {
	return Widget{
		AlertTable: v,
	}
}

// CommentsAsWidget is a convenience function that returns Comments wrapped in Widget
func CommentsAsWidget(v *Comments) Widget {
	return Widget{
		Comments: v,
	}
}

// CustomFieldsAsWidget is a convenience function that returns CustomFields wrapped in Widget
func CustomFieldsAsWidget(v *CustomFields) Widget {
	return Widget{
		CustomFields: v,
	}
}

// CustomFieldsListAsWidget is a convenience function that returns CustomFieldsList wrapped in Widget
func CustomFieldsListAsWidget(v *CustomFieldsList) Widget {
	return Widget{
		CustomFieldsList: v,
	}
}

// ImageAsWidget is a convenience function that returns Image wrapped in Widget
func ImageAsWidget(v *Image) Widget {
	return Widget{
		Image: v,
	}
}

// ObservableListAsWidget is a convenience function that returns ObservableList wrapped in Widget
func ObservableListAsWidget(v *ObservableList) Widget {
	return Widget{
		ObservableList: v,
	}
}

// ObservableTableAsWidget is a convenience function that returns ObservableTable wrapped in Widget
func ObservableTableAsWidget(v *ObservableTable) Widget {
	return Widget{
		ObservableTable: v,
	}
}

// PagesAsWidget is a convenience function that returns Pages wrapped in Widget
func PagesAsWidget(v *Pages) Widget {
	return Widget{
		Pages: v,
	}
}

// TTPListAsWidget is a convenience function that returns TTPList wrapped in Widget
func TTPListAsWidget(v *TTPList) Widget {
	return Widget{
		TTPList: v,
	}
}

// TTPTableAsWidget is a convenience function that returns TTPTable wrapped in Widget
func TTPTableAsWidget(v *TTPTable) Widget {
	return Widget{
		TTPTable: v,
	}
}

// TaskListAsWidget is a convenience function that returns TaskList wrapped in Widget
func TaskListAsWidget(v *TaskList) Widget {
	return Widget{
		TaskList: v,
	}
}

// TaskTableAsWidget is a convenience function that returns TaskTable wrapped in Widget
func TaskTableAsWidget(v *TaskTable) Widget {
	return Widget{
		TaskTable: v,
	}
}

// TextAsWidget is a convenience function that returns Text wrapped in Widget
func TextAsWidget(v *Text) Widget {
	return Widget{
		Text: v,
	}
}

// TimelineAsWidget is a convenience function that returns Timeline wrapped in Widget
func TimelineAsWidget(v *Timeline) Widget {
	return Widget{
		Timeline: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *Widget) UnmarshalJSON(data []byte) error {
	// Postprocessed by scripts/fix-oneof-decoder: dispatch by the OpenAPI
	// discriminator instead of naive structural matching, which fails when
	// two variants generate to byte-identical Go structs.
	var disc struct {
		Kind string `json:"_kind"`
	}
	if err := json.Unmarshal(data, &disc); err != nil {
		return fmt.Errorf("oneOf(Widget): cannot read discriminator _kind: %w", err)
	}
	switch disc.Kind {
	case "AlertList":
		dst.AlertList = new(AlertList)
		return newStrictDecoder(data).Decode(dst.AlertList)
	case "AlertTable":
		dst.AlertTable = new(AlertTable)
		return newStrictDecoder(data).Decode(dst.AlertTable)
	case "Comments":
		dst.Comments = new(Comments)
		return newStrictDecoder(data).Decode(dst.Comments)
	case "CustomFields":
		dst.CustomFields = new(CustomFields)
		return newStrictDecoder(data).Decode(dst.CustomFields)
	case "CustomFieldsList":
		dst.CustomFieldsList = new(CustomFieldsList)
		return newStrictDecoder(data).Decode(dst.CustomFieldsList)
	case "Image":
		dst.Image = new(Image)
		return newStrictDecoder(data).Decode(dst.Image)
	case "ObservableList":
		dst.ObservableList = new(ObservableList)
		return newStrictDecoder(data).Decode(dst.ObservableList)
	case "ObservableTable":
		dst.ObservableTable = new(ObservableTable)
		return newStrictDecoder(data).Decode(dst.ObservableTable)
	case "Pages":
		dst.Pages = new(Pages)
		return newStrictDecoder(data).Decode(dst.Pages)
	case "TTPList":
		dst.TTPList = new(TTPList)
		return newStrictDecoder(data).Decode(dst.TTPList)
	case "TTPTable":
		dst.TTPTable = new(TTPTable)
		return newStrictDecoder(data).Decode(dst.TTPTable)
	case "TaskList":
		dst.TaskList = new(TaskList)
		return newStrictDecoder(data).Decode(dst.TaskList)
	case "TaskTable":
		dst.TaskTable = new(TaskTable)
		return newStrictDecoder(data).Decode(dst.TaskTable)
	case "Text":
		dst.Text = new(Text)
		return newStrictDecoder(data).Decode(dst.Text)
	case "Timeline":
		dst.Timeline = new(Timeline)
		return newStrictDecoder(data).Decode(dst.Timeline)
	default:
		return fmt.Errorf("oneOf(Widget): unknown _kind value %q", disc.Kind)
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Widget) MarshalJSON() ([]byte, error) {
	if src.AlertList != nil {
		return json.Marshal(&src.AlertList)
	}

	if src.AlertTable != nil {
		return json.Marshal(&src.AlertTable)
	}

	if src.Comments != nil {
		return json.Marshal(&src.Comments)
	}

	if src.CustomFields != nil {
		return json.Marshal(&src.CustomFields)
	}

	if src.CustomFieldsList != nil {
		return json.Marshal(&src.CustomFieldsList)
	}

	if src.Image != nil {
		return json.Marshal(&src.Image)
	}

	if src.ObservableList != nil {
		return json.Marshal(&src.ObservableList)
	}

	if src.ObservableTable != nil {
		return json.Marshal(&src.ObservableTable)
	}

	if src.Pages != nil {
		return json.Marshal(&src.Pages)
	}

	if src.TTPList != nil {
		return json.Marshal(&src.TTPList)
	}

	if src.TTPTable != nil {
		return json.Marshal(&src.TTPTable)
	}

	if src.TaskList != nil {
		return json.Marshal(&src.TaskList)
	}

	if src.TaskTable != nil {
		return json.Marshal(&src.TaskTable)
	}

	if src.Text != nil {
		return json.Marshal(&src.Text)
	}

	if src.Timeline != nil {
		return json.Marshal(&src.Timeline)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Widget) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AlertList != nil {
		return obj.AlertList
	}

	if obj.AlertTable != nil {
		return obj.AlertTable
	}

	if obj.Comments != nil {
		return obj.Comments
	}

	if obj.CustomFields != nil {
		return obj.CustomFields
	}

	if obj.CustomFieldsList != nil {
		return obj.CustomFieldsList
	}

	if obj.Image != nil {
		return obj.Image
	}

	if obj.ObservableList != nil {
		return obj.ObservableList
	}

	if obj.ObservableTable != nil {
		return obj.ObservableTable
	}

	if obj.Pages != nil {
		return obj.Pages
	}

	if obj.TTPList != nil {
		return obj.TTPList
	}

	if obj.TTPTable != nil {
		return obj.TTPTable
	}

	if obj.TaskList != nil {
		return obj.TaskList
	}

	if obj.TaskTable != nil {
		return obj.TaskTable
	}

	if obj.Text != nil {
		return obj.Text
	}

	if obj.Timeline != nil {
		return obj.Timeline
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj Widget) GetActualInstanceValue() interface{} {
	if obj.AlertList != nil {
		return *obj.AlertList
	}

	if obj.AlertTable != nil {
		return *obj.AlertTable
	}

	if obj.Comments != nil {
		return *obj.Comments
	}

	if obj.CustomFields != nil {
		return *obj.CustomFields
	}

	if obj.CustomFieldsList != nil {
		return *obj.CustomFieldsList
	}

	if obj.Image != nil {
		return *obj.Image
	}

	if obj.ObservableList != nil {
		return *obj.ObservableList
	}

	if obj.ObservableTable != nil {
		return *obj.ObservableTable
	}

	if obj.Pages != nil {
		return *obj.Pages
	}

	if obj.TTPList != nil {
		return *obj.TTPList
	}

	if obj.TTPTable != nil {
		return *obj.TTPTable
	}

	if obj.TaskList != nil {
		return *obj.TaskList
	}

	if obj.TaskTable != nil {
		return *obj.TaskTable
	}

	if obj.Text != nil {
		return *obj.Text
	}

	if obj.Timeline != nil {
		return *obj.Timeline
	}

	// all schemas are nil
	return nil
}

type NullableWidget struct {
	value *Widget
	isSet bool
}

func (v NullableWidget) Get() *Widget {
	return v.value
}

func (v *NullableWidget) Set(val *Widget) {
	v.value = val
	v.isSet = true
}

func (v NullableWidget) IsSet() bool {
	return v.isSet
}

func (v *NullableWidget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidget(val *Widget) *NullableWidget {
	return &NullableWidget{value: val, isSet: true}
}

func (v NullableWidget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
