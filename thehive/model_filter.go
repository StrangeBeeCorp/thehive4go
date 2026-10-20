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

// Filter - struct for Filter
type Filter struct {
	And        *And
	Any        *Any
	Between    *Between
	Contains   *Contains
	EndsWith   *EndsWith
	Eq         *Eq
	Gt         *Gt
	Gte        *Gte
	Id         *Id
	In         *In
	Like       *Like
	Lt         *Lt
	Lte        *Lte
	Match      *Match
	Ne         *Ne
	Not        *Not
	Or         *Or
	StartsWith *StartsWith
}

// AndAsFilter is a convenience function that returns And wrapped in Filter
func AndAsFilter(v *And) Filter {
	return Filter{
		And: v,
	}
}

// AnyAsFilter is a convenience function that returns Any wrapped in Filter
func AnyAsFilter(v *Any) Filter {
	return Filter{
		Any: v,
	}
}

// BetweenAsFilter is a convenience function that returns Between wrapped in Filter
func BetweenAsFilter(v *Between) Filter {
	return Filter{
		Between: v,
	}
}

// ContainsAsFilter is a convenience function that returns Contains wrapped in Filter
func ContainsAsFilter(v *Contains) Filter {
	return Filter{
		Contains: v,
	}
}

// EndsWithAsFilter is a convenience function that returns EndsWith wrapped in Filter
func EndsWithAsFilter(v *EndsWith) Filter {
	return Filter{
		EndsWith: v,
	}
}

// EqAsFilter is a convenience function that returns Eq wrapped in Filter
func EqAsFilter(v *Eq) Filter {
	return Filter{
		Eq: v,
	}
}

// GtAsFilter is a convenience function that returns Gt wrapped in Filter
func GtAsFilter(v *Gt) Filter {
	return Filter{
		Gt: v,
	}
}

// GteAsFilter is a convenience function that returns Gte wrapped in Filter
func GteAsFilter(v *Gte) Filter {
	return Filter{
		Gte: v,
	}
}

// IdAsFilter is a convenience function that returns Id wrapped in Filter
func IdAsFilter(v *Id) Filter {
	return Filter{
		Id: v,
	}
}

// InAsFilter is a convenience function that returns In wrapped in Filter
func InAsFilter(v *In) Filter {
	return Filter{
		In: v,
	}
}

// LikeAsFilter is a convenience function that returns Like wrapped in Filter
func LikeAsFilter(v *Like) Filter {
	return Filter{
		Like: v,
	}
}

// LtAsFilter is a convenience function that returns Lt wrapped in Filter
func LtAsFilter(v *Lt) Filter {
	return Filter{
		Lt: v,
	}
}

// LteAsFilter is a convenience function that returns Lte wrapped in Filter
func LteAsFilter(v *Lte) Filter {
	return Filter{
		Lte: v,
	}
}

// MatchAsFilter is a convenience function that returns Match wrapped in Filter
func MatchAsFilter(v *Match) Filter {
	return Filter{
		Match: v,
	}
}

// NeAsFilter is a convenience function that returns Ne wrapped in Filter
func NeAsFilter(v *Ne) Filter {
	return Filter{
		Ne: v,
	}
}

// NotAsFilter is a convenience function that returns Not wrapped in Filter
func NotAsFilter(v *Not) Filter {
	return Filter{
		Not: v,
	}
}

// OrAsFilter is a convenience function that returns Or wrapped in Filter
func OrAsFilter(v *Or) Filter {
	return Filter{
		Or: v,
	}
}

// StartsWithAsFilter is a convenience function that returns StartsWith wrapped in Filter
func StartsWithAsFilter(v *StartsWith) Filter {
	return Filter{
		StartsWith: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *Filter) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into And
	err = newStrictDecoder(data).Decode(&dst.And)
	if err == nil {
		jsonAnd, _ := json.Marshal(dst.And)
		if string(jsonAnd) == "{}" { // empty struct
			dst.And = nil
		} else {
			if err = validator.Validate(dst.And); err != nil {
				dst.And = nil
			} else {
				match++
			}
		}
	} else {
		dst.And = nil
	}

	// try to unmarshal data into Any
	err = newStrictDecoder(data).Decode(&dst.Any)
	if err == nil {
		jsonAny, _ := json.Marshal(dst.Any)
		if string(jsonAny) == "{}" { // empty struct
			dst.Any = nil
		} else {
			if err = validator.Validate(dst.Any); err != nil {
				dst.Any = nil
			} else {
				match++
			}
		}
	} else {
		dst.Any = nil
	}

	// try to unmarshal data into Between
	err = newStrictDecoder(data).Decode(&dst.Between)
	if err == nil {
		jsonBetween, _ := json.Marshal(dst.Between)
		if string(jsonBetween) == "{}" { // empty struct
			dst.Between = nil
		} else {
			if err = validator.Validate(dst.Between); err != nil {
				dst.Between = nil
			} else {
				match++
			}
		}
	} else {
		dst.Between = nil
	}

	// try to unmarshal data into Contains
	err = newStrictDecoder(data).Decode(&dst.Contains)
	if err == nil {
		jsonContains, _ := json.Marshal(dst.Contains)
		if string(jsonContains) == "{}" { // empty struct
			dst.Contains = nil
		} else {
			if err = validator.Validate(dst.Contains); err != nil {
				dst.Contains = nil
			} else {
				match++
			}
		}
	} else {
		dst.Contains = nil
	}

	// try to unmarshal data into EndsWith
	err = newStrictDecoder(data).Decode(&dst.EndsWith)
	if err == nil {
		jsonEndsWith, _ := json.Marshal(dst.EndsWith)
		if string(jsonEndsWith) == "{}" { // empty struct
			dst.EndsWith = nil
		} else {
			if err = validator.Validate(dst.EndsWith); err != nil {
				dst.EndsWith = nil
			} else {
				match++
			}
		}
	} else {
		dst.EndsWith = nil
	}

	// try to unmarshal data into Eq
	err = newStrictDecoder(data).Decode(&dst.Eq)
	if err == nil {
		jsonEq, _ := json.Marshal(dst.Eq)
		if string(jsonEq) == "{}" { // empty struct
			dst.Eq = nil
		} else {
			if err = validator.Validate(dst.Eq); err != nil {
				dst.Eq = nil
			} else {
				match++
			}
		}
	} else {
		dst.Eq = nil
	}

	// try to unmarshal data into Gt
	err = newStrictDecoder(data).Decode(&dst.Gt)
	if err == nil {
		jsonGt, _ := json.Marshal(dst.Gt)
		if string(jsonGt) == "{}" { // empty struct
			dst.Gt = nil
		} else {
			if err = validator.Validate(dst.Gt); err != nil {
				dst.Gt = nil
			} else {
				match++
			}
		}
	} else {
		dst.Gt = nil
	}

	// try to unmarshal data into Gte
	err = newStrictDecoder(data).Decode(&dst.Gte)
	if err == nil {
		jsonGte, _ := json.Marshal(dst.Gte)
		if string(jsonGte) == "{}" { // empty struct
			dst.Gte = nil
		} else {
			if err = validator.Validate(dst.Gte); err != nil {
				dst.Gte = nil
			} else {
				match++
			}
		}
	} else {
		dst.Gte = nil
	}

	// try to unmarshal data into Id
	err = newStrictDecoder(data).Decode(&dst.Id)
	if err == nil {
		jsonId, _ := json.Marshal(dst.Id)
		if string(jsonId) == "{}" { // empty struct
			dst.Id = nil
		} else {
			if err = validator.Validate(dst.Id); err != nil {
				dst.Id = nil
			} else {
				match++
			}
		}
	} else {
		dst.Id = nil
	}

	// try to unmarshal data into In
	err = newStrictDecoder(data).Decode(&dst.In)
	if err == nil {
		jsonIn, _ := json.Marshal(dst.In)
		if string(jsonIn) == "{}" { // empty struct
			dst.In = nil
		} else {
			if err = validator.Validate(dst.In); err != nil {
				dst.In = nil
			} else {
				match++
			}
		}
	} else {
		dst.In = nil
	}

	// try to unmarshal data into Like
	err = newStrictDecoder(data).Decode(&dst.Like)
	if err == nil {
		jsonLike, _ := json.Marshal(dst.Like)
		if string(jsonLike) == "{}" { // empty struct
			dst.Like = nil
		} else {
			if err = validator.Validate(dst.Like); err != nil {
				dst.Like = nil
			} else {
				match++
			}
		}
	} else {
		dst.Like = nil
	}

	// try to unmarshal data into Lt
	err = newStrictDecoder(data).Decode(&dst.Lt)
	if err == nil {
		jsonLt, _ := json.Marshal(dst.Lt)
		if string(jsonLt) == "{}" { // empty struct
			dst.Lt = nil
		} else {
			if err = validator.Validate(dst.Lt); err != nil {
				dst.Lt = nil
			} else {
				match++
			}
		}
	} else {
		dst.Lt = nil
	}

	// try to unmarshal data into Lte
	err = newStrictDecoder(data).Decode(&dst.Lte)
	if err == nil {
		jsonLte, _ := json.Marshal(dst.Lte)
		if string(jsonLte) == "{}" { // empty struct
			dst.Lte = nil
		} else {
			if err = validator.Validate(dst.Lte); err != nil {
				dst.Lte = nil
			} else {
				match++
			}
		}
	} else {
		dst.Lte = nil
	}

	// try to unmarshal data into Match
	err = newStrictDecoder(data).Decode(&dst.Match)
	if err == nil {
		jsonMatch, _ := json.Marshal(dst.Match)
		if string(jsonMatch) == "{}" { // empty struct
			dst.Match = nil
		} else {
			if err = validator.Validate(dst.Match); err != nil {
				dst.Match = nil
			} else {
				match++
			}
		}
	} else {
		dst.Match = nil
	}

	// try to unmarshal data into Ne
	err = newStrictDecoder(data).Decode(&dst.Ne)
	if err == nil {
		jsonNe, _ := json.Marshal(dst.Ne)
		if string(jsonNe) == "{}" { // empty struct
			dst.Ne = nil
		} else {
			if err = validator.Validate(dst.Ne); err != nil {
				dst.Ne = nil
			} else {
				match++
			}
		}
	} else {
		dst.Ne = nil
	}

	// try to unmarshal data into Not
	err = newStrictDecoder(data).Decode(&dst.Not)
	if err == nil {
		jsonNot, _ := json.Marshal(dst.Not)
		if string(jsonNot) == "{}" { // empty struct
			dst.Not = nil
		} else {
			if err = validator.Validate(dst.Not); err != nil {
				dst.Not = nil
			} else {
				match++
			}
		}
	} else {
		dst.Not = nil
	}

	// try to unmarshal data into Or
	err = newStrictDecoder(data).Decode(&dst.Or)
	if err == nil {
		jsonOr, _ := json.Marshal(dst.Or)
		if string(jsonOr) == "{}" { // empty struct
			dst.Or = nil
		} else {
			if err = validator.Validate(dst.Or); err != nil {
				dst.Or = nil
			} else {
				match++
			}
		}
	} else {
		dst.Or = nil
	}

	// try to unmarshal data into StartsWith
	err = newStrictDecoder(data).Decode(&dst.StartsWith)
	if err == nil {
		jsonStartsWith, _ := json.Marshal(dst.StartsWith)
		if string(jsonStartsWith) == "{}" { // empty struct
			dst.StartsWith = nil
		} else {
			if err = validator.Validate(dst.StartsWith); err != nil {
				dst.StartsWith = nil
			} else {
				match++
			}
		}
	} else {
		dst.StartsWith = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.And = nil
		dst.Any = nil
		dst.Between = nil
		dst.Contains = nil
		dst.EndsWith = nil
		dst.Eq = nil
		dst.Gt = nil
		dst.Gte = nil
		dst.Id = nil
		dst.In = nil
		dst.Like = nil
		dst.Lt = nil
		dst.Lte = nil
		dst.Match = nil
		dst.Ne = nil
		dst.Not = nil
		dst.Or = nil
		dst.StartsWith = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Filter)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Filter)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Filter) MarshalJSON() ([]byte, error) {
	if src.And != nil {
		return json.Marshal(&src.And)
	}

	if src.Any != nil {
		return json.Marshal(&src.Any)
	}

	if src.Between != nil {
		return json.Marshal(&src.Between)
	}

	if src.Contains != nil {
		return json.Marshal(&src.Contains)
	}

	if src.EndsWith != nil {
		return json.Marshal(&src.EndsWith)
	}

	if src.Eq != nil {
		return json.Marshal(&src.Eq)
	}

	if src.Gt != nil {
		return json.Marshal(&src.Gt)
	}

	if src.Gte != nil {
		return json.Marshal(&src.Gte)
	}

	if src.Id != nil {
		return json.Marshal(&src.Id)
	}

	if src.In != nil {
		return json.Marshal(&src.In)
	}

	if src.Like != nil {
		return json.Marshal(&src.Like)
	}

	if src.Lt != nil {
		return json.Marshal(&src.Lt)
	}

	if src.Lte != nil {
		return json.Marshal(&src.Lte)
	}

	if src.Match != nil {
		return json.Marshal(&src.Match)
	}

	if src.Ne != nil {
		return json.Marshal(&src.Ne)
	}

	if src.Not != nil {
		return json.Marshal(&src.Not)
	}

	if src.Or != nil {
		return json.Marshal(&src.Or)
	}

	if src.StartsWith != nil {
		return json.Marshal(&src.StartsWith)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Filter) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.And != nil {
		return obj.And
	}

	if obj.Any != nil {
		return obj.Any
	}

	if obj.Between != nil {
		return obj.Between
	}

	if obj.Contains != nil {
		return obj.Contains
	}

	if obj.EndsWith != nil {
		return obj.EndsWith
	}

	if obj.Eq != nil {
		return obj.Eq
	}

	if obj.Gt != nil {
		return obj.Gt
	}

	if obj.Gte != nil {
		return obj.Gte
	}

	if obj.Id != nil {
		return obj.Id
	}

	if obj.In != nil {
		return obj.In
	}

	if obj.Like != nil {
		return obj.Like
	}

	if obj.Lt != nil {
		return obj.Lt
	}

	if obj.Lte != nil {
		return obj.Lte
	}

	if obj.Match != nil {
		return obj.Match
	}

	if obj.Ne != nil {
		return obj.Ne
	}

	if obj.Not != nil {
		return obj.Not
	}

	if obj.Or != nil {
		return obj.Or
	}

	if obj.StartsWith != nil {
		return obj.StartsWith
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj Filter) GetActualInstanceValue() interface{} {
	if obj.And != nil {
		return *obj.And
	}

	if obj.Any != nil {
		return *obj.Any
	}

	if obj.Between != nil {
		return *obj.Between
	}

	if obj.Contains != nil {
		return *obj.Contains
	}

	if obj.EndsWith != nil {
		return *obj.EndsWith
	}

	if obj.Eq != nil {
		return *obj.Eq
	}

	if obj.Gt != nil {
		return *obj.Gt
	}

	if obj.Gte != nil {
		return *obj.Gte
	}

	if obj.Id != nil {
		return *obj.Id
	}

	if obj.In != nil {
		return *obj.In
	}

	if obj.Like != nil {
		return *obj.Like
	}

	if obj.Lt != nil {
		return *obj.Lt
	}

	if obj.Lte != nil {
		return *obj.Lte
	}

	if obj.Match != nil {
		return *obj.Match
	}

	if obj.Ne != nil {
		return *obj.Ne
	}

	if obj.Not != nil {
		return *obj.Not
	}

	if obj.Or != nil {
		return *obj.Or
	}

	if obj.StartsWith != nil {
		return *obj.StartsWith
	}

	// all schemas are nil
	return nil
}

type NullableFilter struct {
	value *Filter
	isSet bool
}

func (v NullableFilter) Get() *Filter {
	return v.value
}

func (v *NullableFilter) Set(val *Filter) {
	v.value = val
	v.isSet = true
}

func (v NullableFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilter(val *Filter) *NullableFilter {
	return &NullableFilter{value: val, isSet: true}
}

func (v NullableFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
