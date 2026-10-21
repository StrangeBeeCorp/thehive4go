/*
TheHive

 ## Get started  This guide covers everything you need to make your first successful API call to TheHive.  The publicly hosted API documentation reflects the latest version of TheHive. To browse the documentation that matches the version of TheHive you are running, open it from within the application: select the **?** button in the top right corner, then **API Docs**.  The v0 API is deprecated and isn't documented. New integrations should target v1.  ### Prerequisites  You need a running TheHive instance and an unlocked user account.  ### Authentication  Almost all endpoints require authentication.  TheHive supports two authentication methods:  #### API key (recommended)  Use your API key as a bearer token in the `Authorization` header.  Generate a new API key with `POST /api/v1/user/{userId}/key/renew`, or retrieve an existing one with `GET /api/v1/user/{userId}/key`.  If you have a [Normal user account](https://docs.strangebee.com/thehive/user-guides/organization/configure-organization/manage-user-accounts/about-user-accounts/#types), you can also manage your API key in your [user settings](https://docs.strangebee.com/thehive/user-guides/manage-user-settings/#manage-your-api-key).  curl ```bash curl -X <http_method> -H 'Authorization: Bearer <api_key>' \\   https://<thehive_host>/api/v1/<endpoint> ```  python ```python import requests  headers = {\"Authorization\": \"Bearer <api_key>\"} response = requests.<http_method>(\"https://<thehive_host>/api/v1/<endpoint>\", headers=headers) ```  #### Basic authentication  Only available for [Normal user accounts](https://docs.strangebee.com/thehive/user-guides/organization/configure-organization/manage-user-accounts/about-user-accounts/#types).  Pass your login and password using HTTP Basic authentication. The login is typically an email address. If a default domain is configured in the [authentication settings](https://docs.strangebee.com/thehive/administration/authentication/configure-authentication), a plain username is also accepted.  curl ```bash curl -X <http_method> -u <login>:<password> \\   https://<thehive_host>/api/v1/<endpoint> ```  python ```python import requests  response = requests.<http_method>(\"https://<thehive_host>/api/v1/<endpoint>\", auth=(\"<login>\", \"<password>\")) ```  The following endpoints are public and can be called without credentials:  - `POST /api/v1/login`: Authenticate with login, password, and a TOTP code if MFA is enabled. Returns a session cookie used to authenticate subsequent requests. - `GET /api/v1/logout`: Invalidate the current session. - `POST /api/v1/logout`: Invalidate the current session. Same as `GET`, provided for clients that can't send `GET` requests with a body. - `GET /api/v1/auth/local/passwordPolicy`: Retrieve the password policy configured on the server such as minimum length or complexity requirements. - `GET /api/v1/status/public`: Retrieve public server information, including the TheHive version and license status, without exposing sensitive configuration.  All other endpoints return `401 Unauthorized` when called without valid credentials.  ### Your first API call  Start with `GET /api/v1/user/current`: it requires no request body, confirms your credentials work, and returns your user ID.  curl ```bash curl -X GET -H 'Authorization: Bearer <api_key>' \\      https://<thehive_host>/api/v1/user/current ```  python ```python import requests  BASE_URL = \"https://<thehive_host>\" API_KEY = \"<api_key>\"  headers = {\"Authorization\": f\"Bearer {API_KEY}\"} response = requests.get(f\"{BASE_URL}/api/v1/user/current\", headers=headers) user = response.json() ```  ### X-Organisation header  By default, API calls target the user's default organization. The `defaultOrganisation` field returned by `GET /api/v1/user/current` tells you which one it is.  If your account belongs to several organizations and you need to target a specific one, add the `X-Organisation` header. This is also required when performing administration actions, as these must be executed in the context of the `admin` organization.  curl ```bash curl -H 'Authorization: Bearer <api_key>' \\      -H 'X-Organisation: <organization_name>' \\      https://<thehive_host>/api/v1/user/current ```  python ```python import requests  BASE_URL = \"https://<thehive_host>\" API_KEY = \"<api_key>\"  headers = {\"Authorization\": f\"Bearer {API_KEY}\", \"X-Organisation\": \"<organization_name>\"} response = requests.get(f\"{BASE_URL}/api/v1/user/current\", headers=headers) user = response.json() ```  ### Request and response format  All request and response bodies use JSON: `Content-Type: application/json`.  Dates are represented as Unix timestamps in milliseconds.  ### How the API is organized  TheHive exposes two kinds of endpoints.  #### The Query API  `POST /api/v1/query` handles all listing, filtering, sorting, and pagination needs.  A single endpoint accepts a structured body that chains operations:  ```json {   \"query\": [     { \"_name\": \"listCase\" },     { \"_name\": \"filter\", \"_eq\": { \"_field\": \"status\", \"_value\": \"Open\" } },     { \"_name\": \"sort\", \"_fields\": [{ \"_createdAt\": \"desc\" }] },     { \"_name\": \"page\", \"from\": 0, \"to\": 15, \"extraData\": [\"taskStats\", \"shareCount\"] }   ],   \"excludeFields\": [\"description\", \"summary\"] } ```  See the [**Query and Export**](https://docs.strangebee.com/thehive/api-docs/#tag/Query-and-Export) section for the full query reference.  #### Standard REST endpoints  Standard REST endpoints handle write operations and direct entity lookups.  They follow conventional HTTP conventions:  | <nobr>Operation </nobr>| Method | Example | |-----------|--------|---------| | <nobr>Create </nobr>| `POST` | `POST /api/v1/case` | | <nobr>Read </nobr>| `GET` | `GET /api/v1/case/{idOrName}` | | <nobr>Update </nobr>| `PATCH` | `PATCH /api/v1/case/{idOrName}` | | <nobr>Replace / set </nobr>| `PUT` | `PUT /api/v1/case/{caseId}/shares` | | <nobr>Delete </nobr>| `DELETE` | `DELETE /api/v1/case/{idOrName}` |  ### Permissions and capabilities  Each user account has permissions defined by their profile. TheHive checks them on every API call.  Some endpoints are only available with a Gold, Platinum, or One license. When that's the case, the required license tier is listed on the endpoint.  To check which license is active on your instance and what capabilities it includes, call `GET /api/v1/status`.  ### Request size limits  TheHive enforces the following request size limits:  | <nobr>Request type </nobr>| Default limit | Description | |---|---|---| | <nobr>Standard API requests (JSON payloads) </nobr>| 10 MB | Maximum size for request bodies processed in memory | | <nobr>File uploads and attachments </nobr>| 1 GB | Maximum size for streamed request bodies written to disk |  TheHive rejects requests that exceed these limits.  To stay within these limits: - Avoid sending excessively large JSON payloads. - Upload large files as attachments rather than embedding data directly in requests. - Implement rate limiting and additional protections at the reverse proxy level when exposing the API externally.  ### Errors  Every endpoint returns at least the following HTTP status codes:  | <nobr>Code </nobr>| Meaning | |------|---------| | <nobr>`400 Bad Request` </nobr>| Invalid request parameters or body | | <nobr>`404 Not Found` </nobr>| Resource does not exist | | <nobr>`500 Internal Server Error` </nobr>| Unexpected server-side error |  Endpoints that require authentication additionally return:  | <nobr>Code </nobr>| Meaning | |------|---------| | <nobr>`401 Unauthorized` </nobr>| Missing or invalid credentials | | <nobr>`403 Forbidden` </nobr>| Authenticated but insufficient permissions |  Other codes may appear depending on the specific operation. Refer to individual endpoint documentation for the full list.  The standard error response body is:  ```json {   \"type\": \"AuthorizationError\",   \"message\": \"You don't have permission to perform this action\" } ```  A small number of endpoints deviate from this format. These cases are noted in the relevant endpoint descriptions.

API version: v6.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package thehive

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"os"
)

// QueryAndExportAPIService QueryAndExportAPI service
type QueryAndExportAPIService service

type ApiExportAPIUnstableRouteRequest struct {
	ctx        context.Context
	ApiService *QueryAndExportAPIService
	query      *string
	options    *string
}

func (r ApiExportAPIUnstableRouteRequest) Query(query string) ApiExportAPIUnstableRouteRequest {
	r.query = &query
	return r
}

func (r ApiExportAPIUnstableRouteRequest) Options(options string) ApiExportAPIUnstableRouteRequest {
	r.options = &options
	return r
}

func (r ApiExportAPIUnstableRouteRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.ExportAPIUnstableRouteExecute(r)
}

/*
ExportAPIUnstableRoute Method for ExportAPIUnstableRoute

**Warning**: This is an unstable route. Breaking changes could happen in future releases.

Export the results of a query as a downloadable file in the specified format. The file is streamed back with a `Content-Disposition: attachment` header. When the underlying query produces a total count, the total number of exported elements is also returned in the `X-Total` response header.

The `query` parameter is a stringified JSON value that conforms to the format accepted by the **Query API** endpoint.

The `options` parameter is a stringified JSON value with the following structure:

- Common fields:
  - `format`: `(string)` accepted values are `"csv"`, `"json"`, `"txt"`, and `"misp"`
  - `fileName`: `(optional, string)` file name without extension (default: `<YYYY-MM-DD>_export`)
  - `model`: `(string)` accepted values are `"Case"`, `"Alert"`, `"User"`, `"Organisation"`, `"Procedure"`, `"Task"`, and `"Observable"`
  - `protectData`: `(optional, boolean)` protect the observable `data` field (default: `true`)

- For CSV exports:
  - `fields`: `(optional, array of strings)` list of field paths that replaces the default exported fields. The fields available for each model can be retrieved through the Get export fields endpoint. Two special paths are also accepted: `customFields` (or `customFieldValues`) expands to all custom fields, and `customFields.<name>` (or `customFieldValues.<name>`) targets a single custom field by name
  - `delimiter`: `(optional, character)` default: `,`
  - `quoteChar`: `(optional, character)` default: `"`
  - `escapeChar`: `(optional, character)` default: `\`

- JSON, TXT, and MISP exports take no specific fields

The `"txt"` and `"misp"` formats are only available for the `"Observable"` model.

Output formats:

- `csv`: the first row contains the field paths as headers. Each subsequent row contains the values for one entity.
- `json`: a single JSON array containing one object per entity.
- `txt`: one observable per line, terminated by `\r\n`. Value-based observables produce `<data>`. File-attachment observables produce `<filename>|<sha256>`.
- `misp`: one observable per line, terminated by `\r\n`. Value-based observables produce `*<dataType>\t:<data>`. File-attachment observables produce `*filename|sha256\t:<filename>|<sha256>`.

**Breaking changes in TheHive 5.3:**

- The observable `data` field is protected by default.
- The `model` field is now mandatory for all formats.
- When the `fields` parameter is set for the `"csv"` format, those fields replace the default exported fields instead of being added to them.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiExportAPIUnstableRouteRequest
*/
func (a *QueryAndExportAPIService) ExportAPIUnstableRoute(ctx context.Context) ApiExportAPIUnstableRouteRequest {
	return ApiExportAPIUnstableRouteRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return *os.File
func (a *QueryAndExportAPIService) ExportAPIUnstableRouteExecute(r ApiExportAPIUnstableRouteRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QueryAndExportAPIService.ExportAPIUnstableRoute")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/export"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.query == nil {
		return localVarReturnValue, nil, reportError("query is required and must be specified")
	}
	if r.options == nil {
		return localVarReturnValue, nil, reportError("options is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "query", r.query, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "options", r.options, "form", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/octet-stream", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v QueryAPI400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v QueryAPI400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v QueryAPI400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v QueryAPI400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v QueryAPI400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetExportFieldsUnstableRouteRequest struct {
	ctx        context.Context
	ApiService *QueryAndExportAPIService
}

func (r ApiGetExportFieldsUnstableRouteRequest) Execute() (*OutputExportFieldsMapping, *http.Response, error) {
	return r.ApiService.GetExportFieldsUnstableRouteExecute(r)
}

/*
GetExportFieldsUnstableRoute Method for GetExportFieldsUnstableRoute

**Warning**: This is an unstable route. Breaking changes could happen in future releases.

Returns the list of fields available for export, grouped by model. Each entry exposes the field path, its type, and whether it is part of the default selection. Use this endpoint to discover the field paths accepted by the `fields` parameter of the **Export API** endpoint for CSV exports.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetExportFieldsUnstableRouteRequest
*/
func (a *QueryAndExportAPIService) GetExportFieldsUnstableRoute(ctx context.Context) ApiGetExportFieldsUnstableRouteRequest {
	return ApiGetExportFieldsUnstableRouteRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OutputExportFieldsMapping
func (a *QueryAndExportAPIService) GetExportFieldsUnstableRouteExecute(r ApiGetExportFieldsUnstableRouteRequest) (*OutputExportFieldsMapping, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OutputExportFieldsMapping
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QueryAndExportAPIService.GetExportFieldsUnstableRoute")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/export/_fields"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v QueryAPI400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v QueryAPI400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v QueryAPI400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v QueryAPI400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v QueryAPI400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiQueryAPIRequest struct {
	ctx        context.Context
	ApiService *QueryAndExportAPIService
	inputQuery *InputQuery
}

func (r ApiQueryAPIRequest) InputQuery(inputQuery InputQuery) ApiQueryAPIRequest {
	r.inputQuery = &inputQuery
	return r
}

func (r ApiQueryAPIRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.QueryAPIExecute(r)
}

/*
QueryAPI Method for QueryAPI

<details>
<summary><strong>Primary operations by entity</strong></summary>

| <nobr>Entity </nobr>| <nobr>Operation </nobr>| Description |
|---|---|---|
| <nobr>Any entity </nobr>| <nobr>`listAny` </nobr>| Lists all alerts, cases, observables, tasks, and task logs accessible to the current user, in a single result set |
| <nobr>Any entity </nobr>| <nobr>`getAny` </nobr>| Gets any entity by ID across all entity types accessible to the current user |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`listCase` </nobr>| Lists all cases accessible to the current user |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`getCase` </nobr>| Gets a single case by ID or case number |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`countCase` *(deprecated)* </nobr>| Counts cases accessible to the current user. Replaced by `[{"_name": "listCase"}, {"_name": "count"}]`. Returns a value directly. Doesn't support sort, filter, pagination, extraData, or output |
| <nobr>[Alert](https://docs.strangebee.com/thehive/user-guides/analyst-corner/alerts/about-alerts/) </nobr>| <nobr>`listAlert` </nobr>| Lists all alerts accessible to the current user |
| <nobr>[Alert](https://docs.strangebee.com/thehive/user-guides/analyst-corner/alerts/about-alerts/) </nobr>| <nobr>`getAlert` </nobr>| Gets a single alert by ID or alert reference (`type;source;sourceRef`) |
| <nobr>[Alert](https://docs.strangebee.com/thehive/user-guides/analyst-corner/alerts/about-alerts/) </nobr>| <nobr>`countAlert` *(deprecated)* </nobr>| Counts alerts accessible to the current user. Replaced by `[{"_name": "listAlert"}, {"_name": "count"}]`. Returns a value directly. Doesn't support sort, filter, pagination, extraData, or output |
| <nobr>[Alert](https://docs.strangebee.com/thehive/user-guides/analyst-corner/alerts/about-alerts/) </nobr>| <nobr>`countUnreadAlert` *(deprecated)* </nobr>| Counts alerts with stage `New`. Replaced by a `listAlert`+`filter`+`count` chain. Returns a value directly. Doesn't support sort, filter, pagination, extraData, or output |
| <nobr>[Alert](https://docs.strangebee.com/thehive/user-guides/analyst-corner/alerts/about-alerts/) </nobr>| <nobr>`countImportedAlert` *(deprecated)* </nobr>| Counts alerts with [stage `Imported`](https://docs.strangebee.com/thehive/administration/status/about-statuses/#attributes). Replaced by a `listAlert`+`filter`+`count` chain. Returns a value directly. Doesn't support sort, filter, pagination, extraData, or output |
| <nobr>[Alert](https://docs.strangebee.com/thehive/user-guides/analyst-corner/alerts/about-alerts/) </nobr>| <nobr>`countRelatedAlert` *(deprecated)* </nobr>| Counts alerts related to a given case. Replaced by a `listAlert`+`filter`+`count` chain. Returns a value directly. Doesn't support sort, filter, pagination, extraData, or output |
| <nobr>[Task](https://docs.strangebee.com/thehive/user-guides/analyst-corner/tasks/about-tasks/) </nobr>| <nobr>`listTask` </nobr>| Lists all tasks accessible to the current user |
| <nobr>[Task](https://docs.strangebee.com/thehive/user-guides/analyst-corner/tasks/about-tasks/) </nobr>| <nobr>`getTask` </nobr>| Gets a single task by ID |
| <nobr>[Task](https://docs.strangebee.com/thehive/user-guides/analyst-corner/tasks/about-tasks/) </nobr>| <nobr>`listCaseTask` </nobr>| Lists tasks linked to cases (excludes case-template tasks) |
| <nobr>[Task](https://docs.strangebee.com/thehive/user-guides/analyst-corner/tasks/about-tasks/) </nobr>| <nobr>`waitingTasks` </nobr>| Lists tasks with status `Waiting` linked to cases |
| <nobr>[Task](https://docs.strangebee.com/thehive/user-guides/analyst-corner/tasks/about-tasks/) </nobr>| <nobr>`myTasks` </nobr>| Lists tasks assigned to the current user |
| <nobr>[Task](https://docs.strangebee.com/thehive/user-guides/analyst-corner/tasks/about-tasks/) </nobr>| <nobr>`waitingTask` *(deprecated)* </nobr>| Alias for `waitingTasks` |
| <nobr>[Task](https://docs.strangebee.com/thehive/user-guides/analyst-corner/tasks/about-tasks/) </nobr>| <nobr>`countTask` *(deprecated)* </nobr>| Counts non-cancelled tasks for a given case. Replaced by a `listTask`+`filter`+`count` chain. Returns a value directly. Doesn't support sort, filter, pagination, extraData, or output |
| <nobr>[Observable](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/observables/about-observables/) </nobr>| <nobr>`listObservable` </nobr>| Lists all observables accessible to the current user |
| <nobr>[Observable](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/observables/about-observables/) </nobr>| <nobr>`getObservable` </nobr>| Gets a single observable by ID |
| <nobr>[Observable](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/observables/about-observables/) </nobr>| <nobr>`countCaseObservable` *(deprecated)* </nobr>| Counts observables of a given case. Replaced by a `listObservable`+`filter`+`count` chain. Returns a value directly. Doesn't support sort, filter, pagination, extraData, or output |
| <nobr>[Observable](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/observables/about-observables/) </nobr>| <nobr>`countAlertObservable` *(deprecated)* </nobr>| Counts observables of a given alert. Replaced by a `listObservable`+`filter`+`count` chain. Returns a value directly. Doesn't support sort, filter, pagination, extraData, or output |
| <nobr>[Log](https://docs.strangebee.com/thehive/user-guides/analyst-corner/tasks/about-task-logs/) </nobr>| <nobr>`listLog` </nobr>| Lists all task logs accessible to the current user |
| <nobr>[Log](https://docs.strangebee.com/thehive/user-guides/analyst-corner/tasks/about-task-logs/) </nobr>| <nobr>`getLog` </nobr>| Gets a single task log by ID |
| <nobr>Comment </nobr>| <nobr>`listComment` </nobr>| Lists all comments accessible to the current user |
| <nobr>Comment </nobr>| <nobr>`getComment` </nobr>| Gets a single comment by ID |
| <nobr>[Share](https://docs.strangebee.com/thehive/administration/organizations/about-organizations-sharing-rules/) </nobr>| <nobr>`listShare` </nobr>| Lists all sharing configurations. |
| <nobr>[Share](https://docs.strangebee.com/thehive/administration/organizations/about-organizations-sharing-rules/) </nobr>| <nobr>`getShare` </nobr>| Gets a single sharing configuration by ID |
| <nobr>[Procedure](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/ttps/about-ttps/) </nobr>| <nobr>`listProcedure` </nobr>| Lists all procedures (TTPs) accessible to the current user |
| <nobr>[Procedure](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/ttps/about-ttps/) </nobr>| <nobr>`getProcedure` </nobr>| Gets a single procedure by ID |
| <nobr>Page </nobr>| <nobr>`listPage` </nobr>| Lists all pages accessible to the current user ([Knowledge Base pages](https://docs.strangebee.com/thehive/user-guides/knowledge-base/about-knowledge-base/) and [case pages](https://docs.strangebee.com/thehive/user-guides/knowledge-base/about-case-pages/)) |
| <nobr>Page </nobr>| <nobr>`listOrganisationPage` </nobr>| Lists the [Knowledge Base pages](https://docs.strangebee.com/thehive/user-guides/knowledge-base/about-knowledge-base/) of the current organization |
| <nobr>Page </nobr>| <nobr>`getPage` </nobr>| Gets a single page by ID ([Knowledge Base page](https://docs.strangebee.com/thehive/user-guides/knowledge-base/about-knowledge-base/) or [case page](https://docs.strangebee.com/thehive/user-guides/knowledge-base/about-case-pages/)) |
| <nobr>[Case template](https://docs.strangebee.com/thehive/user-guides/organization/configure-organization/manage-templates/case-templates/about-case-templates/) </nobr>| <nobr>`listCaseTemplate` </nobr>| Lists all case templates |
| <nobr>[Case template](https://docs.strangebee.com/thehive/user-guides/organization/configure-organization/manage-templates/case-templates/about-case-templates/) </nobr>| <nobr>`getCaseTemplate` </nobr>| Gets a single case template by ID or name |
| <nobr>[Page template](https://docs.strangebee.com/thehive/user-guides/organization/configure-organization/manage-templates/case-page-templates/about-case-page-templates/) </nobr>| <nobr>`listPageTemplate` </nobr>| Lists all page templates |
| <nobr>[Page template](https://docs.strangebee.com/thehive/user-guides/organization/configure-organization/manage-templates/case-page-templates/about-case-page-templates/) </nobr>| <nobr>`getPageTemplate` </nobr>| Gets a single page template by ID |
| <nobr>[Organization](https://docs.strangebee.com/thehive/administration/organizations/about-organizations/) </nobr>| <nobr>`listOrganisation` </nobr>| Lists all organizations |
| <nobr>[Organization](https://docs.strangebee.com/thehive/administration/organizations/about-organizations/) </nobr>| <nobr>`getOrganisation` </nobr>| Gets a single organization by ID or name |
| <nobr>[User](https://docs.strangebee.com/thehive/user-guides/organization/configure-organization/manage-user-accounts/about-user-accounts/) </nobr>| <nobr>`listUser` </nobr>| Lists all users |
| <nobr>[User](https://docs.strangebee.com/thehive/user-guides/organization/configure-organization/manage-user-accounts/about-user-accounts/) </nobr>| <nobr>`getUser` </nobr>| Gets a single user by ID or login |
| <nobr>[User](https://docs.strangebee.com/thehive/user-guides/organization/configure-organization/manage-user-accounts/about-user-accounts/) </nobr>| <nobr>`currentUser` </nobr>| Returns the currently authenticated user |
| <nobr>[User](https://docs.strangebee.com/thehive/user-guides/organization/configure-organization/manage-user-accounts/about-user-accounts/) </nobr>| <nobr>`listVisibleUsers` </nobr>| Lists users from all organizations visible to the current user |
| <nobr>[Custom field](https://docs.strangebee.com/thehive/administration/custom-fields/about-custom-fields/) </nobr>| <nobr>`listCustomField` </nobr>| Lists all custom fields |
| <nobr>[Custom field](https://docs.strangebee.com/thehive/administration/custom-fields/about-custom-fields/) </nobr>| <nobr>`getCustomField` </nobr>| Gets a single custom field by ID or name |
| <nobr>[Alert status](https://docs.strangebee.com/thehive/administration/status/about-statuses/) </nobr>| <nobr>`listAlertStatus` </nobr>| Lists all alert statuses |
| <nobr>[Alert status](https://docs.strangebee.com/thehive/administration/status/about-statuses/) </nobr>| <nobr>`getAlertStatus` </nobr>| Gets a single alert status by ID or value |
| <nobr>[Case status](https://docs.strangebee.com/thehive/administration/status/about-statuses/) </nobr>| <nobr>`listCaseStatus` </nobr>| Lists all case statuses |
| <nobr>[Case status](https://docs.strangebee.com/thehive/administration/status/about-statuses/) </nobr>| <nobr>`getCaseStatus` </nobr>| Gets a single case status by ID or value |
| <nobr>[Observable type](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/observables/about-observables/#type) </nobr>| <nobr>`listObservableType` </nobr>| Lists all observable types |
| <nobr>[Observable type](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/observables/about-observables/#type) </nobr>| <nobr>`getObservableType` </nobr>| Gets a single observable type by ID or name |
| <nobr>[Pattern](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/ttps/about-ttps/) </nobr>| <nobr>`listPattern` </nobr>| Lists all attack techniques accessible to the current user |
| <nobr>[Pattern](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/ttps/about-ttps/) </nobr>| <nobr>`getPattern` </nobr>| Gets a single attack technique by ID |
| <nobr>[Taxonomy](https://docs.strangebee.com/thehive/administration/taxonomies/about-taxonomies/) </nobr>| <nobr>`listTaxonomy` </nobr>| Lists all taxonomies accessible to the current user |
| <nobr>[Taxonomy](https://docs.strangebee.com/thehive/administration/taxonomies/about-taxonomies/) </nobr>| <nobr>`getTaxonomy` </nobr>| Gets a single taxonomy by ID or namespace |
| <nobr>[Tag](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/tags/about-tags/) </nobr>| <nobr>`listTag` </nobr>| Lists all custom tags of the current organization (excludes taxonomy-imported tags) |
| <nobr>[Tag](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/tags/about-tags/) </nobr>| <nobr>`getTag` </nobr>| Gets a single tag by ID |
| <nobr>[Tag](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/tags/about-tags/) </nobr>| <nobr>`freetags` </nobr>| Alias for `listTag` |
| <nobr>[Tag](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/tags/about-tags/) </nobr>| <nobr>`tagAutoComplete` </nobr>| Autocompletes tag suggestions. Takes a `TagHint` parameter (`namespace`, `predicate`, `value`, `freeTag`, `limit`). Doesn't support sort, filter, pagination, or extraData |
| <nobr>[Tag](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/tags/about-tags/) </nobr>| <nobr>`countFreetags` </nobr>| Counts custom tags of the current organization. Returns a value directly. Doesn't support sort, filter, pagination, extraData, or output |
| <nobr>[Catalog of patterns](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/ttps/about-ttps/) </nobr>| <nobr>`listCatalog` </nobr>| Lists all attack technique catalogs |
| <nobr>[Catalog of patterns](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/ttps/about-ttps/) </nobr>| <nobr>`getCatalog` </nobr>| Gets a single attack technique catalog by ID or name |
| <nobr>[Profile](https://docs.strangebee.com/thehive/administration/profiles/about-profiles/) </nobr>| <nobr>`listProfile` </nobr>| Lists all profiles |
| <nobr>[Profile](https://docs.strangebee.com/thehive/administration/profiles/about-profiles/) </nobr>| <nobr>`getProfile` </nobr>| Gets a single profile by ID or name |
| <nobr>[Dashboard](https://docs.strangebee.com/thehive/user-guides/analyst-corner/dashboard/about-dashboards/) </nobr>| <nobr>`listDashboard` </nobr>| Lists all dashboards accessible to the current user |
| <nobr>[Dashboard](https://docs.strangebee.com/thehive/user-guides/analyst-corner/dashboard/about-dashboards/) </nobr>| <nobr>`getDashboard` </nobr>| Gets a single dashboard by ID |
| <nobr>[Custom event](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/case-timelines/about-case-timelines/) </nobr>| <nobr>`listCustomEvent` </nobr>| Lists all custom events in timelines accessible to the current user |
| <nobr>[Custom event](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/case-timelines/about-case-timelines/) </nobr>| <nobr>`getCustomEvent` </nobr>| Gets a single custom event from timelines by ID |
| <nobr>[Case report](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/case-reports/about-case-reports/) </nobr>| <nobr>`listCaseReport` </nobr>| Lists all case reports accessible to the current user |
| <nobr>[Case report](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/case-reports/about-case-reports/) </nobr>| <nobr>`getCaseReport` </nobr>| Gets a single case report by ID |
| <nobr>[Case report template](https://docs.strangebee.com/thehive/user-guides/organization/configure-organization/manage-templates/case-report-templates/about-case-report-templates/) </nobr>| <nobr>`listCaseReportTemplate` </nobr>| Lists all case report templates |
| <nobr>[Case report template](https://docs.strangebee.com/thehive/user-guides/organization/configure-organization/manage-templates/case-report-templates/about-case-report-templates/) </nobr>| <nobr>`getCaseReportTemplate` </nobr>| Gets a single case report template by ID |
| <nobr>[Audit](https://docs.strangebee.com/thehive/user-guides/organization/about-audit-logs/) </nobr>| <nobr>`getAudit` </nobr>| Gets a single audit log by ID |
| <nobr>[Audit](https://docs.strangebee.com/thehive/user-guides/organization/about-audit-logs/) </nobr>| <nobr>`listAuditFromObject` </nobr>| Lists audit logs for a given object. Takes an `id` parameter: the object's ID |
| <nobr>Action (Cortex) </nobr>| <nobr>`listAction` </nobr>| Lists all Cortex responder actions accessible to the current user |
| <nobr>Action (Cortex) </nobr>| <nobr>`getAction` </nobr>| Gets a single Cortex responder action by ID |
| <nobr>[Analyzer template (Cortex)](https://docs.strangebee.com/thehive/administration/analyzer-templates/about-analyzer-templates/) </nobr>| <nobr>`listAnalyzerTemplate` </nobr>| Lists all analyzer report templates |
| <nobr>[Analyzer template (Cortex)](https://docs.strangebee.com/thehive/administration/analyzer-templates/about-analyzer-templates/) </nobr>| <nobr>`getReportTemplate` </nobr>| Gets a single analyzer report template by ID |
| <nobr>Job (Cortex) </nobr>| <nobr>`listJob` </nobr>| Lists all Cortex analysis jobs accessible to the current user |
| <nobr>Job (Cortex) </nobr>| <nobr>`getJob` </nobr>| Gets a single Cortex analysis job by ID |
| <nobr>[Function](https://docs.strangebee.com/thehive/user-guides/organization/configure-organization/manage-functions/about-functions/) </nobr>| <nobr>`listFunction` </nobr>| Lists all functions accessible to the current user. Requires a Platinum or One license |
| <nobr>[Function](https://docs.strangebee.com/thehive/user-guides/organization/configure-organization/manage-functions/about-functions/) </nobr>| <nobr>`getFunction` </nobr>| Gets a single function by ID or name. Requires a Platinum or One license |

</details>

<details>
<summary><strong>Related object operations by entity</strong></summary>

| <nobr>Source entity </nobr>| <nobr>Operation </nobr>| Returns |
|---|---|---|
| <nobr>Any entity </nobr>| <nobr>`count` </nobr>| Number of entities in the current result set. Recommended replacement for the deprecated `countXxx` operations |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`tasks` </nobr>| Tasks of the case |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`taskGroups` </nobr>| Distinct task group names in the case. Returns a value directly. Doesn't support sort, filter, pagination, extraData, or output |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`observables` </nobr>| Observables of the case |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`alerts` </nobr>| Alerts merged into the case |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`shares` </nobr>| Sharing configurations of the case |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`organisations` </nobr>| Organizations the case is shared with |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`procedures` </nobr>| Procedures (TTPs) attached to the case |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`pages` </nobr>| Pages attached to the case |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`attachments` </nobr>| Attachments of the case |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`comments` </nobr>| Comments on the case |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`allComments` </nobr>| Comments on the case and on its merged alerts |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`assignableUsers` </nobr>| Users who can be assigned to the case |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`similarCasesLight` </nobr>| Cases sharing observables with this case, with the matching observables grouped by data type. Doesn't support sort, filter, pagination, extraData, or output |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`similarCaseLightCount` </nobr>| Number of cases sharing observables with this case. Returns a value directly. Doesn't support sort, filter, pagination, extraData, or output |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`similarAlertsLight` </nobr>| Alerts sharing observables with this case, with the matching observables grouped by data type. Doesn't support sort, filter, pagination, extraData, or output |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`similarAlertLightCount` </nobr>| Number of alerts sharing observables with this case. Returns a value directly. Doesn't support sort, filter, pagination, extraData, or output |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`similarCases` *(deprecated)* </nobr>| Use `similarCasesLight` for improved performance. Doesn't support sort, filter, pagination, extraData, or output |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`similarCaseCount` *(deprecated)* </nobr>| Use `similarCaseLightCount` for improved performance. Returns a value directly. Doesn't support sort, filter, pagination, extraData, or output |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`similarAlerts` *(deprecated)* </nobr>| Use `similarAlertsLight` for improved performance. Doesn't support sort, filter, pagination, extraData, or output |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`similarAlertCount` *(deprecated)* </nobr>| Use `similarAlertLightCount` for improved performance. Returns a value directly. Doesn't support sort, filter, pagination, extraData, or output |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`linkedCases` *(deprecated)* </nobr>| Alias for `similarCases`. Use `similarCasesLight` instead. Doesn't support sort, filter, pagination, extraData, or output |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`linkedCaseCount` *(deprecated)* </nobr>| Use `similarCaseLightCount` for improved performance. Returns a value directly. Doesn't support sort, filter, pagination, extraData, or output |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`customEvents` </nobr>| Custom events attached to the case in the timeline |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`caseReports` </nobr>| Case reports of the case |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`timeline` </nobr>| Activity timeline of the case |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`actions` </nobr>| Cortex responder executions run on the case |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`external` </nobr>| Restricts the result set to external-access cases ([TheHive Portal](https://docs.strangebee.com/thehive/administration/thehive-portal/about-thehive-portal/)). Requires a Platinum or One license. |
| <nobr>[Alert](https://docs.strangebee.com/thehive/user-guides/analyst-corner/alerts/about-alerts/) </nobr>| <nobr>`case` </nobr>| Case the alert was merged into |
| <nobr>[Alert](https://docs.strangebee.com/thehive/user-guides/analyst-corner/alerts/about-alerts/) </nobr>| <nobr>`observables` </nobr>| Observables of the alert |
| <nobr>[Alert](https://docs.strangebee.com/thehive/user-guides/analyst-corner/alerts/about-alerts/) </nobr>| <nobr>`procedures` </nobr>| Procedures (TTPs) attached to the alert |
| <nobr>[Alert](https://docs.strangebee.com/thehive/user-guides/analyst-corner/alerts/about-alerts/) </nobr>| <nobr>`comments` </nobr>| Comments on the alert |
| <nobr>[Alert](https://docs.strangebee.com/thehive/user-guides/analyst-corner/alerts/about-alerts/) </nobr>| <nobr>`attachments` </nobr>| Attachments of the alert |
| <nobr>[Alert](https://docs.strangebee.com/thehive/user-guides/analyst-corner/alerts/about-alerts/) </nobr>| <nobr>`assignableUsers` </nobr>| Users who can be assigned to the alert |
| <nobr>[Alert](https://docs.strangebee.com/thehive/user-guides/analyst-corner/alerts/about-alerts/) </nobr>| <nobr>`similarCasesLight` </nobr>| Cases sharing observables with this alert, with the matching observables grouped by data type. Doesn't support sort, filter, pagination, extraData, or output |
| <nobr>[Alert](https://docs.strangebee.com/thehive/user-guides/analyst-corner/alerts/about-alerts/) </nobr>| <nobr>`similarCaseLightCount` </nobr>| Number of cases sharing observables with this alert. Returns a value directly. Doesn't support sort, filter, pagination, extraData, or output |
| <nobr>[Alert](https://docs.strangebee.com/thehive/user-guides/analyst-corner/alerts/about-alerts/) </nobr>| <nobr>`similarAlertsLight` </nobr>| Alerts sharing observables with this alert, with the matching observables grouped by data type. Doesn't support sort, filter, pagination, extraData, or output |
| <nobr>[Alert](https://docs.strangebee.com/thehive/user-guides/analyst-corner/alerts/about-alerts/) </nobr>| <nobr>`similarAlertLightCount` </nobr>| Number of alerts sharing observables with this alert. Returns a value directly. Doesn't support sort, filter, pagination, extraData, or output |
| <nobr>[Alert](https://docs.strangebee.com/thehive/user-guides/analyst-corner/alerts/about-alerts/) </nobr>| <nobr>`similarCases` *(deprecated)* </nobr>| Use `similarCasesLight` for improved performance. Doesn't support sort, filter, pagination, extraData, or output |
| <nobr>[Alert](https://docs.strangebee.com/thehive/user-guides/analyst-corner/alerts/about-alerts/) </nobr>| <nobr>`similarCaseCount` *(deprecated)* </nobr>| Use `similarCaseLightCount` for improved performance. Returns a value directly. Doesn't support sort, filter, pagination, extraData, or output |
| <nobr>[Alert](https://docs.strangebee.com/thehive/user-guides/analyst-corner/alerts/about-alerts/) </nobr>| <nobr>`similarAlerts` *(deprecated)* </nobr>| Use `similarAlertsLight` for improved performance. Doesn't support sort, filter, pagination, extraData, or output |
| <nobr>[Alert](https://docs.strangebee.com/thehive/user-guides/analyst-corner/alerts/about-alerts/) </nobr>| <nobr>`similarAlertCount` *(deprecated)* </nobr>| Use `similarAlertLightCount` for improved performance. Returns a value directly. Doesn't support sort, filter, pagination, extraData, or output |
| <nobr>[Alert](https://docs.strangebee.com/thehive/user-guides/analyst-corner/alerts/about-alerts/) </nobr>| <nobr>`actions` </nobr>| Cortex responder executions run on the alert |
| <nobr>[Task](https://docs.strangebee.com/thehive/user-guides/analyst-corner/tasks/about-tasks/) </nobr>| <nobr>`case` </nobr>| Case the task belongs to |
| <nobr>[Task](https://docs.strangebee.com/thehive/user-guides/analyst-corner/tasks/about-tasks/) </nobr>| <nobr>`caseTemplate` </nobr>| Case template the task originates from |
| <nobr>[Task](https://docs.strangebee.com/thehive/user-guides/analyst-corner/tasks/about-tasks/) </nobr>| <nobr>`logs` </nobr>| Task logs of the task |
| <nobr>[Task](https://docs.strangebee.com/thehive/user-guides/analyst-corner/tasks/about-tasks/) </nobr>| <nobr>`shares` </nobr>| Sharing configurations of the task |
| <nobr>[Task](https://docs.strangebee.com/thehive/user-guides/analyst-corner/tasks/about-tasks/) </nobr>| <nobr>`organisations` </nobr>| Organizations the task is shared with |
| <nobr>[Task](https://docs.strangebee.com/thehive/user-guides/analyst-corner/tasks/about-tasks/) </nobr>| <nobr>`assignableUsers` </nobr>| Users who can be assigned to the task |
| <nobr>[Task](https://docs.strangebee.com/thehive/user-guides/analyst-corner/tasks/about-tasks/) </nobr>| <nobr>`inCase` </nobr>| Restricts the result set to tasks linked to a case (excludes case-template tasks) |
| <nobr>[Task](https://docs.strangebee.com/thehive/user-guides/analyst-corner/tasks/about-tasks/) </nobr>| <nobr>`actions` </nobr>| Cortex responder executions run on the task |
| <nobr>[Observable](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/observables/about-observables/) </nobr>| <nobr>`case` </nobr>| Case the observable belongs to |
| <nobr>[Observable](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/observables/about-observables/) </nobr>| <nobr>`alert` </nobr>| Alert the observable belongs to |
| <nobr>[Observable](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/observables/about-observables/) </nobr>| <nobr>`shares` </nobr>| Sharing configurations of the observable |
| <nobr>[Observable](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/observables/about-observables/) </nobr>| <nobr>`organisations` </nobr>| Organizations the observable is shared with |
| <nobr>[Observable](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/observables/about-observables/) </nobr>| <nobr>`similar` </nobr>| Observables similar to this one |
| <nobr>[Observable](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/observables/about-observables/) </nobr>| <nobr>`attachments` </nobr>| Attachments of the observable |
| <nobr>[Observable](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/observables/about-observables/) </nobr>| <nobr>`jobs` </nobr>| Cortex analysis jobs run on the observable |
| <nobr>[Observable](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/observables/about-observables/) </nobr>| <nobr>`actions` </nobr>| Cortex responder executions run on the observable |
| <nobr>[Observable](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/observables/about-observables/) </nobr>| <nobr>`fromCase` </nobr>| Restricts the result set to observables attached to a case |
| <nobr>[Observable](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/observables/about-observables/) </nobr>| <nobr>`fromAlert` </nobr>| Restricts the result set to observables attached to an alert |
| <nobr>[Observable](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/observables/about-observables/) </nobr>| <nobr>`fromJobReport` </nobr>| Restricts the result set to observables produced by a Cortex analyzer report |
| <nobr>[Log](https://docs.strangebee.com/thehive/user-guides/analyst-corner/tasks/about-task-logs/) </nobr>| <nobr>`attachments` </nobr>| Attachments of the log |
| <nobr>[Log](https://docs.strangebee.com/thehive/user-guides/analyst-corner/tasks/about-task-logs/) </nobr>| <nobr>`actions` </nobr>| Cortex responder executions run on the task log |
| <nobr>[Share](https://docs.strangebee.com/thehive/administration/organizations/about-organizations-sharing-rules/) </nobr>| <nobr>`case` </nobr>| Case associated with the share |
| <nobr>[Share](https://docs.strangebee.com/thehive/administration/organizations/about-organizations-sharing-rules/) </nobr>| <nobr>`observables` </nobr>| Observables associated with the share |
| <nobr>[Share](https://docs.strangebee.com/thehive/administration/organizations/about-organizations-sharing-rules/) </nobr>| <nobr>`tasks` </nobr>| Tasks associated with the share |
| <nobr>[Share](https://docs.strangebee.com/thehive/administration/organizations/about-organizations-sharing-rules/) </nobr>| <nobr>`organisation` </nobr>| Organization associated with the share |
| <nobr>[Case template](https://docs.strangebee.com/thehive/user-guides/organization/configure-organization/manage-templates/case-templates/about-case-templates/) </nobr>| <nobr>`tasks` </nobr>| Task templates of the case template |
| <nobr>[Case template](https://docs.strangebee.com/thehive/user-guides/organization/configure-organization/manage-templates/case-templates/about-case-templates/) </nobr>| <nobr>`pageTemplates` </nobr>| Page templates of the case template |
| <nobr>[Page template](https://docs.strangebee.com/thehive/user-guides/organization/configure-organization/manage-templates/case-page-templates/about-case-page-templates/) </nobr>| <nobr>`caseTemplates` </nobr>| Case templates that include the page template |
| <nobr>[Page template](https://docs.strangebee.com/thehive/user-guides/organization/configure-organization/manage-templates/case-page-templates/about-case-page-templates/) </nobr>| <nobr>`headerPage` *(experimental)* </nobr>| Page template header info, without the content field. Accepts `from`, `to`, and `extraData` pagination parameters |
| <nobr>[Organization](https://docs.strangebee.com/thehive/administration/organizations/about-organizations/) </nobr>| <nobr>`users` </nobr>| Users in the organization |
| <nobr>[Organization](https://docs.strangebee.com/thehive/administration/organizations/about-organizations/) </nobr>| <nobr>`caseTemplates` </nobr>| Case templates of the organization |
| <nobr>[Organization](https://docs.strangebee.com/thehive/administration/organizations/about-organizations/) </nobr>| <nobr>`alerts` </nobr>| Alerts of the organization |
| <nobr>[Organization](https://docs.strangebee.com/thehive/administration/organizations/about-organizations/) </nobr>| <nobr>`attachments` </nobr>| Attachments of the organization |
| <nobr>[Organization](https://docs.strangebee.com/thehive/administration/organizations/about-organizations/) </nobr>| <nobr>`visible` </nobr>| Restricts the result set to organizations visible to the current user |
| <nobr>[User](https://docs.strangebee.com/thehive/user-guides/organization/configure-organization/manage-user-accounts/about-user-accounts/) </nobr>| <nobr>`tasks` </nobr>| Tasks assigned to the user |
| <nobr>[User](https://docs.strangebee.com/thehive/user-guides/organization/configure-organization/manage-user-accounts/about-user-accounts/) </nobr>| <nobr>`cases` </nobr>| Cases assigned to the user |
| <nobr>[User](https://docs.strangebee.com/thehive/user-guides/organization/configure-organization/manage-user-accounts/about-user-accounts/) </nobr>| <nobr>`alerts` </nobr>| Alerts assigned to the user |
| <nobr>[Tag](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/tags/about-tags/) </nobr>| <nobr>`freetags` </nobr>| Restricts the result set to custom tags of the current organization |
| <nobr>[Tag](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/tags/about-tags/) </nobr>| <nobr>`text` </nobr>| Display name of each tag. Doesn't support sort, filter, pagination, or extraData |
| <nobr>[Catalog of patterns](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/ttps/about-ttps/) </nobr>| <nobr>`patterns` </nobr>| Techniques in the attack catalog |
| <nobr>[Catalog of patterns](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/ttps/about-ttps/) </nobr>| <nobr>`tactics` </nobr>| Tactics in the attack catalog. Supports sort, filter, count, and aggregation. Doesn't support pagination, extraData, or output |
| <nobr>[Taxonomy](https://docs.strangebee.com/thehive/administration/taxonomies/about-taxonomies/) </nobr>| <nobr>`tags` </nobr>| Tags belonging to the taxonomy |
| <nobr>[Case report template](https://docs.strangebee.com/thehive/user-guides/organization/configure-organization/manage-templates/case-report-templates/about-case-report-templates/) </nobr>| <nobr>`attachments` </nobr>| Files uploaded to the template for use in `Image` widgets |
| <nobr>Job (Cortex) </nobr>| <nobr>`observable` </nobr>| Observable the job analyzed |
| <nobr>Job (Cortex) </nobr>| <nobr>`reportObservables` </nobr>| Observables produced by the analyzer report |

</details>

<details>
<summary><strong>Available extraData fields by entity</strong></summary>

| <nobr>Entity </nobr>| <nobr>extraData value </nobr>| Description |
|---|---|---|
| <nobr>Any entity </nobr>| <nobr>`total` </nobr>| Total number of matching results across all pages. Returned in the `X-Total` response header |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`observableStats` </nobr>| Statistics about observables in the case |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`taskStats` </nobr>| Statistics about tasks in the case |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`alerts` </nobr>| Alerts linked to the case |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`alertCount` </nobr>| Number of alerts linked to the case |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`isOwner` </nobr>| Whether the requesting organization owns the case |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`shareCount` </nobr>| Number of organizations the case is shared with |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`permissions` </nobr>| Permissions the current user has on the case |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`actionRequired` </nobr>| Whether an action is required from the current user |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`procedureCount` </nobr>| Number of procedures (TTPs) attached to the case |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`attachmentCount` </nobr>| Number of attachments in the case |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`contributors` </nobr>| Users who have contributed to the case |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`computed.handlingDuration` </nobr>| Time elapsed from case creation to close (in milliseconds) |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`computed.handlingDurationInSeconds` </nobr>| Handling duration in seconds |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`computed.handlingDurationInMinutes` </nobr>| Handling duration in minutes |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`computed.handlingDurationInHours` </nobr>| Handling duration in hours |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`computed.handlingDurationInDays` </nobr>| Handling duration in days |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`status` </nobr>| Status object with stage and color |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`owningOrganisation` </nobr>| Organization that owns the case |
| <nobr>[Case](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/about-cases/) </nobr>| <nobr>`links` </nobr>| Cases linked to this case |
| <nobr>[Alert](https://docs.strangebee.com/thehive/user-guides/analyst-corner/alerts/about-alerts/) </nobr>| <nobr>`importDate` </nobr>| Date the alert was merged into a case |
| <nobr>[Alert](https://docs.strangebee.com/thehive/user-guides/analyst-corner/alerts/about-alerts/) </nobr>| <nobr>`caseNumber` *(deprecated)* </nobr>| Number of the case the alert was merged into |
| <nobr>[Alert](https://docs.strangebee.com/thehive/user-guides/analyst-corner/alerts/about-alerts/) </nobr>| <nobr>`relatedCase` </nobr>| Case the alert was merged into |
| <nobr>[Alert](https://docs.strangebee.com/thehive/user-guides/analyst-corner/alerts/about-alerts/) </nobr>| <nobr>`status` </nobr>| Status object with stage and color |
| <nobr>[Alert](https://docs.strangebee.com/thehive/user-guides/analyst-corner/alerts/about-alerts/) </nobr>| <nobr>`procedureCount` </nobr>| Number of procedures (TTPs) attached to the alert |
| <nobr>[Task](https://docs.strangebee.com/thehive/user-guides/analyst-corner/tasks/about-tasks/) </nobr>| <nobr>`case` </nobr>| Case the task belongs to |
| <nobr>[Task](https://docs.strangebee.com/thehive/user-guides/analyst-corner/tasks/about-tasks/) </nobr>| <nobr>`caseId` </nobr>| ID of the case the task belongs to |
| <nobr>[Task](https://docs.strangebee.com/thehive/user-guides/analyst-corner/tasks/about-tasks/) </nobr>| <nobr>`caseTemplate` </nobr>| Case template the task originates from |
| <nobr>[Task](https://docs.strangebee.com/thehive/user-guides/analyst-corner/tasks/about-tasks/) </nobr>| <nobr>`caseTemplateId` </nobr>| ID of the case template |
| <nobr>[Task](https://docs.strangebee.com/thehive/user-guides/analyst-corner/tasks/about-tasks/) </nobr>| <nobr>`isOwner` </nobr>| Whether the requesting organization owns the task |
| <nobr>[Task](https://docs.strangebee.com/thehive/user-guides/analyst-corner/tasks/about-tasks/) </nobr>| <nobr>`shareCount` </nobr>| Number of organizations the task is shared with |
| <nobr>[Task](https://docs.strangebee.com/thehive/user-guides/analyst-corner/tasks/about-tasks/) </nobr>| <nobr>`actionRequired` </nobr>| Whether an action is required on the task |
| <nobr>[Task](https://docs.strangebee.com/thehive/user-guides/analyst-corner/tasks/about-tasks/) </nobr>| <nobr>`actionRequiredMap` </nobr>| Map of organizations requiring action on the task |
| <nobr>[Task](https://docs.strangebee.com/thehive/user-guides/analyst-corner/tasks/about-tasks/) </nobr>| <nobr>`logCount` </nobr>| Number of logs in the task |
| <nobr>[Observable](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/observables/about-observables/) </nobr>| <nobr>`seen` *(deprecated)* </nobr>| Whether the observable has been seen in other cases |
| <nobr>[Observable](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/observables/about-observables/) </nobr>| <nobr>`seenSummary` </nobr>| Summary of cases where the observable appears |
| <nobr>[Observable](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/observables/about-observables/) </nobr>| <nobr>`shares` </nobr>| Organizations the observable is shared with |
| <nobr>[Observable](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/observables/about-observables/) </nobr>| <nobr>`links` </nobr>| Observables linked to this one |
| <nobr>[Observable](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/observables/about-observables/) </nobr>| <nobr>`permissions` </nobr>| Permissions the current user has on the observable |
| <nobr>[Observable](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/observables/about-observables/) </nobr>| <nobr>`isOwner` </nobr>| Whether the requesting organization owns the observable |
| <nobr>[Observable](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/observables/about-observables/) </nobr>| <nobr>`shareCount` </nobr>| Number of organizations the observable is shared with |
| <nobr>[Log](https://docs.strangebee.com/thehive/user-guides/analyst-corner/tasks/about-task-logs/) </nobr>| <nobr>`case` </nobr>| Case the task log belongs to |
| <nobr>[Log](https://docs.strangebee.com/thehive/user-guides/analyst-corner/tasks/about-task-logs/) </nobr>| <nobr>`task` </nobr>| Task the log belongs to |
| <nobr>[Log](https://docs.strangebee.com/thehive/user-guides/analyst-corner/tasks/about-task-logs/) </nobr>| <nobr>`taskId` </nobr>| ID of the task the log belongs to |
| <nobr>[Log](https://docs.strangebee.com/thehive/user-guides/analyst-corner/tasks/about-task-logs/) </nobr>| <nobr>`actionCount` </nobr>| Number of actions on the task log |
| <nobr>Comment </nobr>| <nobr>`links` </nobr>| Case or alert the comment belongs to |
| <nobr>[Procedure](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/ttps/about-ttps/) </nobr>| <nobr>`pattern` </nobr>| Attack techniques associated with the procedure |
| <nobr>[Procedure](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/ttps/about-ttps/) </nobr>| <nobr>`tactic` </nobr>| Attack tactics associated with the procedure |
| <nobr>[Procedure](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/ttps/about-ttps/) </nobr>| <nobr>`patternParent` </nobr>| Parent attack technique in the hierarchy |
| <nobr>[Procedure](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/ttps/about-ttps/) </nobr>| <nobr>`patternTactics` </nobr>| All tactics associated with the attack technique |
| <nobr>[Procedure](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/ttps/about-ttps/) </nobr>| <nobr>`patternCatalog` </nobr>| Catalog the attack technique belongs to |
| <nobr>[Organization](https://docs.strangebee.com/thehive/administration/organizations/about-organizations/) </nobr>| <nobr>`userCount` </nobr>| Number of users in the organization |
| <nobr>[User](https://docs.strangebee.com/thehive/user-guides/organization/configure-organization/manage-user-accounts/about-user-accounts/) </nobr>| <nobr>`lockout` </nobr>| Information about authentication lockout for the user |
| <nobr>[User](https://docs.strangebee.com/thehive/user-guides/organization/configure-organization/manage-user-accounts/about-user-accounts/) </nobr>| <nobr>`organisations` </nobr>| Organizations the user belongs to |
| <nobr>[Pattern](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/ttps/about-ttps/) </nobr>| <nobr>`parent` </nobr>| Parent attack technique in the hierarchy |
| <nobr>[Pattern](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/ttps/about-ttps/) </nobr>| <nobr>`children` </nobr>| Child attack techniques in the hierarchy |
| <nobr>[Pattern](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/ttps/about-ttps/) </nobr>| <nobr>`tactics` </nobr>| Tactics associated with the attack technique |
| <nobr>[Pattern](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/ttps/about-ttps/) </nobr>| <nobr>`catalog` </nobr>| Catalog the attack technique belongs to |
| <nobr>[Catalog of patterns](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/ttps/about-ttps/) </nobr>| <nobr>`tactics` </nobr>| Tactics defined in the attack catalog |
| <nobr>[Tag](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/tags/about-tags/) </nobr>| <nobr>`usage` </nobr>| Usage statistics for the tag |
| <nobr>[Taxonomy](https://docs.strangebee.com/thehive/administration/taxonomies/about-taxonomies/) </nobr>| <nobr>`enabled` </nobr>| Whether the taxonomy is enabled in the current organization |
| <nobr>[Custom field](https://docs.strangebee.com/thehive/administration/custom-fields/about-custom-fields/) </nobr>| <nobr>`deletable` </nobr>| Whether the custom field can be deleted |
| <nobr>[Alert status](https://docs.strangebee.com/thehive/administration/status/about-statuses/) </nobr>| <nobr>`canDelete` </nobr>| Whether the alert status can be deleted |
| <nobr>[Case status](https://docs.strangebee.com/thehive/administration/status/about-statuses/) </nobr>| <nobr>`canDelete` </nobr>| Whether the case status can be deleted |
| <nobr>[Case template](https://docs.strangebee.com/thehive/user-guides/organization/configure-organization/manage-templates/case-templates/about-case-templates/) </nobr>| <nobr>`pageTemplateHeaders` </nobr>| Headers of the page templates linked to the case template |
| <nobr>[Page template](https://docs.strangebee.com/thehive/user-guides/organization/configure-organization/manage-templates/case-page-templates/about-case-page-templates/) </nobr>| <nobr>`caseTemplateCount` </nobr>| Number of case templates that include this page template |
| <nobr>Job (Cortex) </nobr>| <nobr>`links` </nobr>| Case or alert the Cortex job's observable belongs to |
| <nobr>Job (Cortex) </nobr>| <nobr>`report` </nobr>| Cortex analyzer report content |
| <nobr>[Attachment](https://docs.strangebee.com/thehive/user-guides/analyst-corner/cases/attachments/about-attachments/) </nobr>| <nobr>`task` </nobr>| Task log the attachment belongs to (when the attachment is a log attachment) |

</details>

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiQueryAPIRequest
*/
func (a *QueryAndExportAPIService) QueryAPI(ctx context.Context) ApiQueryAPIRequest {
	return ApiQueryAPIRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return interface{}
func (a *QueryAndExportAPIService) QueryAPIExecute(r ApiQueryAPIRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QueryAndExportAPIService.QueryAPI")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.inputQuery == nil {
		return localVarReturnValue, nil, reportError("inputQuery is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.inputQuery
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v QueryAPI400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v QueryAPI400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v QueryAPI400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v QueryAPI400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v QueryAPI400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
