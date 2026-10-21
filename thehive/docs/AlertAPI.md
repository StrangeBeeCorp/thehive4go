# \AlertAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAlertAttachment**](AlertAPI.md#AddAlertAttachment) | **Post** /api/v1/alert/{alertId}/attachments | Add attachments to an alert
[**BulkUpdateAlert**](AlertAPI.md#BulkUpdateAlert) | **Patch** /api/v1/alert/_bulk | Update multiple alerts
[**CreateAlert**](AlertAPI.md#CreateAlert) | **Post** /api/v1/alert | Create an alert
[**CreateCaseFromAlert**](AlertAPI.md#CreateCaseFromAlert) | **Post** /api/v1/alert/{alertId}/case | Create a case from an alert
[**DeleteAlert**](AlertAPI.md#DeleteAlert) | **Delete** /api/v1/alert/{alertId} | Delete an alert
[**DeleteAlertAttachment**](AlertAPI.md#DeleteAlertAttachment) | **Delete** /api/v1/alert/{alertId}/attachment/{attachmentId} | Remove an attachment from an alert
[**DeleteAlertInBulk**](AlertAPI.md#DeleteAlertInBulk) | **Post** /api/v1/alert/delete/_bulk | Delete multiple alerts
[**DownloadAlertAttachment**](AlertAPI.md#DownloadAlertAttachment) | **Get** /api/v1/alert/{alertId}/attachment/{attachmentId}/download | Download an attachment from an alert
[**FollowAlert**](AlertAPI.md#FollowAlert) | **Post** /api/v1/alert/{alertId}/follow | Follow an alert
[**GetAlert**](AlertAPI.md#GetAlert) | **Get** /api/v1/alert/{alertId} | Get an alert
[**GetAlertAttachment**](AlertAPI.md#GetAlertAttachment) | **Get** /api/v1/alert/{alertId}/attachment/{attachmentId} | Get an attachment from an alert
[**GetSimilarObservablesBetweenAlertAndAnotherAlertOrCase**](AlertAPI.md#GetSimilarObservablesBetweenAlertAndAnotherAlertOrCase) | **Get** /api/v1/alert/{alertId}/similar/{alertOrCaseId}/observables | List observables shared with another alert or case
[**ImportAlertObservablesAndProceduresInCase**](AlertAPI.md#ImportAlertObservablesAndProceduresInCase) | **Post** /api/v1/alert/{alertId}/import/{caseId} | Import alert observables and procedures into a case
[**MergeAlertWithCase**](AlertAPI.md#MergeAlertWithCase) | **Post** /api/v1/alert/{alertId}/merge/{caseId} | Merge an alert into an existing case
[**MergeBulkAlertsWithCase**](AlertAPI.md#MergeBulkAlertsWithCase) | **Post** /api/v1/alert/merge/_bulk | Merge multiple alerts into a case
[**UnfollowAlert**](AlertAPI.md#UnfollowAlert) | **Post** /api/v1/alert/{alertId}/unfollow | Unfollow an alert
[**UpdateAlert**](AlertAPI.md#UpdateAlert) | **Patch** /api/v1/alert/{alertId} | Update an alert



## AddAlertAttachment

> OutputAttachments AddAlertAttachment(ctx, alertId).Attachments(attachments).CanRename(canRename).Execute()

Add attachments to an alert



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
	alertId := "alertId_example" // string | The alert ID preceded by `~`.
	attachments := []*os.File{"TODO"} // []*os.File | Files to attach. Each part is a multipart form file field.
	canRename := true // bool | When `true`, renames files that already exist with the same name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertAPI.AddAlertAttachment(context.Background(), alertId).Attachments(attachments).CanRename(canRename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.AddAlertAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddAlertAttachment`: OutputAttachments
	fmt.Fprintf(os.Stdout, "Response from `AlertAPI.AddAlertAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string** | The alert ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiAddAlertAttachmentRequest struct via the builder pattern


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


## BulkUpdateAlert

> BulkUpdateAlert(ctx).InputUpdateAlertWithIds(inputUpdateAlertWithIds).Execute()

Update multiple alerts



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
	inputUpdateAlertWithIds := *openapiclient.NewInputUpdateAlertWithIds([]string{"Ids_example"}) // InputUpdateAlertWithIds |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertAPI.BulkUpdateAlert(context.Background()).InputUpdateAlertWithIds(inputUpdateAlertWithIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.BulkUpdateAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputUpdateAlertWithIds** | [**InputUpdateAlertWithIds**](InputUpdateAlertWithIds.md) |  |

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


## CreateAlert

> OutputAlert CreateAlert(ctx).InputCreateAlert(inputCreateAlert).Execute()

Create an alert



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
	inputCreateAlert := *openapiclient.NewInputCreateAlert("Type_example", "Source_example", "SourceRef_example", "Title_example", "Description_example") // InputCreateAlert |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertAPI.CreateAlert(context.Background()).InputCreateAlert(inputCreateAlert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.CreateAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAlert`: OutputAlert
	fmt.Fprintf(os.Stdout, "Response from `AlertAPI.CreateAlert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputCreateAlert** | [**InputCreateAlert**](InputCreateAlert.md) |  |

### Return type

[**OutputAlert**](OutputAlert.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCaseFromAlert

> OutputCase CreateCaseFromAlert(ctx, alertId).InputCreateCaseFromAlert(inputCreateCaseFromAlert).Execute()

Create a case from an alert



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
	alertId := "alertId_example" // string | The alert ID preceded by `~`.
	inputCreateCaseFromAlert := *openapiclient.NewInputCreateCaseFromAlert() // InputCreateCaseFromAlert |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertAPI.CreateCaseFromAlert(context.Background(), alertId).InputCreateCaseFromAlert(inputCreateCaseFromAlert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.CreateCaseFromAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCaseFromAlert`: OutputCase
	fmt.Fprintf(os.Stdout, "Response from `AlertAPI.CreateCaseFromAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string** | The alert ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCaseFromAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputCreateCaseFromAlert** | [**InputCreateCaseFromAlert**](InputCreateCaseFromAlert.md) |  |

### Return type

[**OutputCase**](OutputCase.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlert

> DeleteAlert(ctx, alertId).Execute()

Delete an alert



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
	alertId := "alertId_example" // string | The alert ID preceded by `~`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertAPI.DeleteAlert(context.Background(), alertId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.DeleteAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string** | The alert ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertRequest struct via the builder pattern


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


## DeleteAlertAttachment

> DeleteAlertAttachment(ctx, alertId, attachmentId).Execute()

Remove an attachment from an alert



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
	alertId := "alertId_example" // string | The alert ID preceded by `~`.
	attachmentId := "attachmentId_example" // string | The attachment ID preceded by `~`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertAPI.DeleteAlertAttachment(context.Background(), alertId, attachmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.DeleteAlertAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string** | The alert ID preceded by &#x60;~&#x60;. |
**attachmentId** | **string** | The attachment ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertAttachmentRequest struct via the builder pattern


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


## DeleteAlertInBulk

> DeleteAlertInBulk(ctx).DeleteAlertInBulkRequest(deleteAlertInBulkRequest).Execute()

Delete multiple alerts



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
	deleteAlertInBulkRequest := *openapiclient.NewDeleteAlertInBulkRequest([]string{"Ids_example"}) // DeleteAlertInBulkRequest |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertAPI.DeleteAlertInBulk(context.Background()).DeleteAlertInBulkRequest(deleteAlertInBulkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.DeleteAlertInBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertInBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteAlertInBulkRequest** | [**DeleteAlertInBulkRequest**](DeleteAlertInBulkRequest.md) |  |

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


## DownloadAlertAttachment

> *os.File DownloadAlertAttachment(ctx, alertId, attachmentId).Execute()

Download an attachment from an alert



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
	alertId := "alertId_example" // string | The alert ID preceded by `~`.
	attachmentId := "attachmentId_example" // string | The attachment ID preceded by `~`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertAPI.DownloadAlertAttachment(context.Background(), alertId, attachmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.DownloadAlertAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadAlertAttachment`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AlertAPI.DownloadAlertAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string** | The alert ID preceded by &#x60;~&#x60;. |
**attachmentId** | **string** | The attachment ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadAlertAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## FollowAlert

> FollowAlert(ctx, alertId).Execute()

Follow an alert



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
	alertId := "alertId_example" // string | The alert ID preceded by `~`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertAPI.FollowAlert(context.Background(), alertId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.FollowAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string** | The alert ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiFollowAlertRequest struct via the builder pattern


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


## GetAlert

> OutputAlert GetAlert(ctx, alertId).Execute()

Get an alert



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
	alertId := "alertId_example" // string | The alert ID preceded by `~`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertAPI.GetAlert(context.Background(), alertId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.GetAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlert`: OutputAlert
	fmt.Fprintf(os.Stdout, "Response from `AlertAPI.GetAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string** | The alert ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputAlert**](OutputAlert.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertAttachment

> *os.File GetAlertAttachment(ctx, alertId, attachmentId).IfNoneMatch(ifNoneMatch).Execute()

Get an attachment from an alert



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
	alertId := "alertId_example" // string | The alert ID preceded by `~`.
	attachmentId := "attachmentId_example" // string | The attachment ID preceded by `~`.
	ifNoneMatch := "ifNoneMatch_example" // string | ETag from a previous response. If it matches, returns `304 Not Modified` instead of the file content. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertAPI.GetAlertAttachment(context.Background(), alertId, attachmentId).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.GetAlertAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertAttachment`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AlertAPI.GetAlertAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string** | The alert ID preceded by &#x60;~&#x60;. |
**attachmentId** | **string** | The attachment ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertAttachmentRequest struct via the builder pattern


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


## GetSimilarObservablesBetweenAlertAndAnotherAlertOrCase

> []OutputObservable GetSimilarObservablesBetweenAlertAndAnotherAlertOrCase(ctx, alertId, alertOrCaseId).Execute()

List observables shared with another alert or case



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
	alertId := "alertId_example" // string | The alert ID preceded by `~`.
	alertOrCaseId := "alertOrCaseId_example" // string | The ID of the alert or case, preceded by `~`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertAPI.GetSimilarObservablesBetweenAlertAndAnotherAlertOrCase(context.Background(), alertId, alertOrCaseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.GetSimilarObservablesBetweenAlertAndAnotherAlertOrCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimilarObservablesBetweenAlertAndAnotherAlertOrCase`: []OutputObservable
	fmt.Fprintf(os.Stdout, "Response from `AlertAPI.GetSimilarObservablesBetweenAlertAndAnotherAlertOrCase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string** | The alert ID preceded by &#x60;~&#x60;. |
**alertOrCaseId** | **string** | The ID of the alert or case, preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetSimilarObservablesBetweenAlertAndAnotherAlertOrCaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]OutputObservable**](OutputObservable.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportAlertObservablesAndProceduresInCase

> OutputCase ImportAlertObservablesAndProceduresInCase(ctx, alertId, caseId).Execute()

Import alert observables and procedures into a case



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
	alertId := "alertId_example" // string | The alert ID preceded by `~`.
	caseId := "caseId_example" // string | The case ID preceded by `~`, or case number.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertAPI.ImportAlertObservablesAndProceduresInCase(context.Background(), alertId, caseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.ImportAlertObservablesAndProceduresInCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportAlertObservablesAndProceduresInCase`: OutputCase
	fmt.Fprintf(os.Stdout, "Response from `AlertAPI.ImportAlertObservablesAndProceduresInCase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string** | The alert ID preceded by &#x60;~&#x60;. |
**caseId** | **string** | The case ID preceded by &#x60;~&#x60;, or case number. |

### Other Parameters

Other parameters are passed through a pointer to a apiImportAlertObservablesAndProceduresInCaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OutputCase**](OutputCase.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MergeAlertWithCase

> OutputCase MergeAlertWithCase(ctx, alertId, caseId).Execute()

Merge an alert into an existing case



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
	alertId := "alertId_example" // string | The alert ID preceded by `~`.
	caseId := "caseId_example" // string | The case ID preceded by `~`, or case number.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertAPI.MergeAlertWithCase(context.Background(), alertId, caseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.MergeAlertWithCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MergeAlertWithCase`: OutputCase
	fmt.Fprintf(os.Stdout, "Response from `AlertAPI.MergeAlertWithCase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string** | The alert ID preceded by &#x60;~&#x60;. |
**caseId** | **string** | The case ID preceded by &#x60;~&#x60;, or case number. |

### Other Parameters

Other parameters are passed through a pointer to a apiMergeAlertWithCaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OutputCase**](OutputCase.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MergeBulkAlertsWithCase

> OutputCase MergeBulkAlertsWithCase(ctx).InputAlertsMergeWithCase(inputAlertsMergeWithCase).Execute()

Merge multiple alerts into a case



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
	inputAlertsMergeWithCase := *openapiclient.NewInputAlertsMergeWithCase("CaseId_example") // InputAlertsMergeWithCase |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertAPI.MergeBulkAlertsWithCase(context.Background()).InputAlertsMergeWithCase(inputAlertsMergeWithCase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.MergeBulkAlertsWithCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MergeBulkAlertsWithCase`: OutputCase
	fmt.Fprintf(os.Stdout, "Response from `AlertAPI.MergeBulkAlertsWithCase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMergeBulkAlertsWithCaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputAlertsMergeWithCase** | [**InputAlertsMergeWithCase**](InputAlertsMergeWithCase.md) |  |

### Return type

[**OutputCase**](OutputCase.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnfollowAlert

> UnfollowAlert(ctx, alertId).Execute()

Unfollow an alert



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
	alertId := "alertId_example" // string | The alert ID preceded by `~`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertAPI.UnfollowAlert(context.Background(), alertId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.UnfollowAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string** | The alert ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiUnfollowAlertRequest struct via the builder pattern


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


## UpdateAlert

> UpdateAlert(ctx, alertId).InputUpdateAlert(inputUpdateAlert).Execute()

Update an alert



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
	alertId := "alertId_example" // string | The alert ID preceded by `~`.
	inputUpdateAlert := *openapiclient.NewInputUpdateAlert() // InputUpdateAlert |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertAPI.UpdateAlert(context.Background(), alertId).InputUpdateAlert(inputUpdateAlert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.UpdateAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string** | The alert ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputUpdateAlert** | [**InputUpdateAlert**](InputUpdateAlert.md) |  |

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
