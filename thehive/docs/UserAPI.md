# \UserAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTemporaryAttachment**](UserAPI.md#AddTemporaryAttachment) | **Post** /api/v1/user/current/attachments | Upload temporary attachments
[**ChangeUserPassword**](UserAPI.md#ChangeUserPassword) | **Post** /api/v1/user/{userId}/password/change | Change a user account password
[**CreateUser**](UserAPI.md#CreateUser) | **Post** /api/v1/user | Create a user account
[**DeleteUser**](UserAPI.md#DeleteUser) | **Delete** /api/v1/user/{userId}/force | Delete a user account
[**GetCurrentUserInfo**](UserAPI.md#GetCurrentUserInfo) | **Get** /api/v1/user/current | Get current user info
[**GetUser**](UserAPI.md#GetUser) | **Get** /api/v1/user/{userId} | Get a user account
[**GetUserApiKey**](UserAPI.md#GetUserApiKey) | **Get** /api/v1/user/{userId}/key | Get a user API key
[**GetUserAvatar**](UserAPI.md#GetUserAvatar) | **Get** /api/v1/user/{userId}/avatar/{file} | Get a user avatar
[**LockUser**](UserAPI.md#LockUser) | **Delete** /api/v1/user/{userId} | Lock a user account
[**RemoveUserApiKey**](UserAPI.md#RemoveUserApiKey) | **Delete** /api/v1/user/{userId}/key | Revoke a user API key
[**RenewUserApiKey**](UserAPI.md#RenewUserApiKey) | **Post** /api/v1/user/{userId}/key/renew | Renew a user API key
[**SetUserOrganisations**](UserAPI.md#SetUserOrganisations) | **Put** /api/v1/user/{userId}/organisations | Set user account organizations
[**SetUserPassword**](UserAPI.md#SetUserPassword) | **Post** /api/v1/user/{userId}/password/set | Set a user account password
[**UpdateUser**](UserAPI.md#UpdateUser) | **Patch** /api/v1/user/{userId} | Update a user account



## AddTemporaryAttachment

> OutputAttachments AddTemporaryAttachment(ctx).Attachments(attachments).CanRename(canRename).Execute()

Upload temporary attachments



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StrangeBeeCorp/thehive4go/thehive"
)

func main() {
	attachments := []*os.File{"TODO"} // []*os.File | Files to attach. Each part is a multipart form file field.
	canRename := true // bool | When `true`, renames files that already exist with the same name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.AddTemporaryAttachment(context.Background()).Attachments(attachments).CanRename(canRename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.AddTemporaryAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddTemporaryAttachment`: OutputAttachments
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.AddTemporaryAttachment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddTemporaryAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attachments** | **[]*os.File** | Files to attach. Each part is a multipart form file field. |
 **canRename** | **bool** | When &#x60;true&#x60;, renames files that already exist with the same name. |

### Return type

[**OutputAttachments**](OutputAttachments.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeUserPassword

> ChangeUserPassword(ctx, userId).InputChangeUserPassword(inputChangeUserPassword).Execute()

Change a user account password



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StrangeBeeCorp/thehive4go/thehive"
)

func main() {
	userId := "userId_example" // string | The user ID preceded by `~`, or the user login.
	inputChangeUserPassword := *openapiclient.NewInputChangeUserPassword("Password_example", "CurrentPassword_example") // InputChangeUserPassword |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.ChangeUserPassword(context.Background(), userId).InputChangeUserPassword(inputChangeUserPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ChangeUserPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user ID preceded by &#x60;~&#x60;, or the user login. |

### Other Parameters

Other parameters are passed through a pointer to a apiChangeUserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputChangeUserPassword** | [**InputChangeUserPassword**](InputChangeUserPassword.md) |  |

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> OutputUser CreateUser(ctx).InputCreateUser(inputCreateUser).Execute()

Create a user account



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StrangeBeeCorp/thehive4go/thehive"
)

func main() {
	inputCreateUser := *openapiclient.NewInputCreateUser("Login_example", "Name_example", "Profile_example") // InputCreateUser |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.CreateUser(context.Background()).InputCreateUser(inputCreateUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.CreateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUser`: OutputUser
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputCreateUser** | [**InputCreateUser**](InputCreateUser.md) |  |

### Return type

[**OutputUser**](OutputUser.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser(ctx, userId).Organisation(organisation).Execute()

Delete a user account



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StrangeBeeCorp/thehive4go/thehive"
)

func main() {
	userId := "userId_example" // string | The user ID preceded by `~`, or the user login.
	organisation := "TheOrganization" // string | Name or ID of the organization to remove the user from. Use `POST /api/v1/query` with `listOrganisation` to retrieve available organizations. If omitted, the user is deleted from all organizations. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.DeleteUser(context.Background(), userId).Organisation(organisation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user ID preceded by &#x60;~&#x60;, or the user login. |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organisation** | **string** | Name or ID of the organization to remove the user from. Use &#x60;POST /api/v1/query&#x60; with &#x60;listOrganisation&#x60; to retrieve available organizations. If omitted, the user is deleted from all organizations. |

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUserInfo

> OutputUser GetCurrentUserInfo(ctx).Execute()

Get current user info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StrangeBeeCorp/thehive4go/thehive"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetCurrentUserInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetCurrentUserInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentUserInfo`: OutputUser
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetCurrentUserInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserInfoRequest struct via the builder pattern


### Return type

[**OutputUser**](OutputUser.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> OutputUser GetUser(ctx, userId).Execute()

Get a user account



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StrangeBeeCorp/thehive4go/thehive"
)

func main() {
	userId := "userId_example" // string | The user ID preceded by `~`, or the user login.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: OutputUser
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user ID preceded by &#x60;~&#x60;, or the user login. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputUser**](OutputUser.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserApiKey

> string GetUserApiKey(ctx, userId).Execute()

Get a user API key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StrangeBeeCorp/thehive4go/thehive"
)

func main() {
	userId := "userId_example" // string | The user ID preceded by `~`, or the user login.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetUserApiKey(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUserApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserApiKey`: string
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUserApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user ID preceded by &#x60;~&#x60;, or the user login. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserAvatar

> *os.File GetUserAvatar(ctx, userId, file).IfNoneMatch(ifNoneMatch).Execute()

Get a user avatar



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StrangeBeeCorp/thehive4go/thehive"
)

func main() {
	userId := "userId_example" // string | The user ID preceded by `~`, or the user login.
	file := "file_example" // string | The avatar filename returned in the user's `avatar` field.
	ifNoneMatch := "ifNoneMatch_example" // string | ETag from a previous response. If it matches, returns `304 Not Modified` instead of the file content. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetUserAvatar(context.Background(), userId, file).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUserAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserAvatar`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUserAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user ID preceded by &#x60;~&#x60;, or the user login. |
**file** | **string** | The avatar filename returned in the user&#39;s &#x60;avatar&#x60; field. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifNoneMatch** | **string** | ETag from a previous response. If it matches, returns &#x60;304 Not Modified&#x60; instead of the file content. |

### Return type

[***os.File**](*os.File.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LockUser

> LockUser(ctx, userId).Execute()

Lock a user account



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StrangeBeeCorp/thehive4go/thehive"
)

func main() {
	userId := "userId_example" // string | The user ID preceded by `~`, or the user login.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.LockUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.LockUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user ID preceded by &#x60;~&#x60;, or the user login. |

### Other Parameters

Other parameters are passed through a pointer to a apiLockUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveUserApiKey

> RemoveUserApiKey(ctx, userId).Execute()

Revoke a user API key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StrangeBeeCorp/thehive4go/thehive"
)

func main() {
	userId := "userId_example" // string | The user ID preceded by `~`, or the user login.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.RemoveUserApiKey(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.RemoveUserApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user ID preceded by &#x60;~&#x60;, or the user login. |

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenewUserApiKey

> string RenewUserApiKey(ctx, userId).Execute()

Renew a user API key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StrangeBeeCorp/thehive4go/thehive"
)

func main() {
	userId := "userId_example" // string | The user ID preceded by `~`, or the user login.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.RenewUserApiKey(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.RenewUserApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenewUserApiKey`: string
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.RenewUserApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user ID preceded by &#x60;~&#x60;, or the user login. |

### Other Parameters

Other parameters are passed through a pointer to a apiRenewUserApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetUserOrganisations

> OutputSetUserOrganisations SetUserOrganisations(ctx, userId).InputSetUserOrganisations(inputSetUserOrganisations).Execute()

Set user account organizations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StrangeBeeCorp/thehive4go/thehive"
)

func main() {
	userId := "userId_example" // string | The user ID preceded by `~`, or the user login.
	inputSetUserOrganisations := *openapiclient.NewInputSetUserOrganisations() // InputSetUserOrganisations |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.SetUserOrganisations(context.Background(), userId).InputSetUserOrganisations(inputSetUserOrganisations).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.SetUserOrganisations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetUserOrganisations`: OutputSetUserOrganisations
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.SetUserOrganisations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user ID preceded by &#x60;~&#x60;, or the user login. |

### Other Parameters

Other parameters are passed through a pointer to a apiSetUserOrganisationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputSetUserOrganisations** | [**InputSetUserOrganisations**](InputSetUserOrganisations.md) |  |

### Return type

[**OutputSetUserOrganisations**](OutputSetUserOrganisations.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetUserPassword

> SetUserPassword(ctx, userId).InputSetUserPassword(inputSetUserPassword).Execute()

Set a user account password



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StrangeBeeCorp/thehive4go/thehive"
)

func main() {
	userId := "userId_example" // string | The user ID preceded by `~`, or the user login.
	inputSetUserPassword := *openapiclient.NewInputSetUserPassword("Password_example") // InputSetUserPassword |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.SetUserPassword(context.Background(), userId).InputSetUserPassword(inputSetUserPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.SetUserPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user ID preceded by &#x60;~&#x60;, or the user login. |

### Other Parameters

Other parameters are passed through a pointer to a apiSetUserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputSetUserPassword** | [**InputSetUserPassword**](InputSetUserPassword.md) |  |

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> UpdateUser(ctx, userId).InputUpdateUser(inputUpdateUser).Execute()

Update a user account



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StrangeBeeCorp/thehive4go/thehive"
)

func main() {
	userId := "userId_example" // string | The user ID preceded by `~`, or the user login.
	inputUpdateUser := *openapiclient.NewInputUpdateUser() // InputUpdateUser |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.UpdateUser(context.Background(), userId).InputUpdateUser(inputUpdateUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user ID preceded by &#x60;~&#x60;, or the user login. |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputUpdateUser** | [**InputUpdateUser**](InputUpdateUser.md) |  |

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
