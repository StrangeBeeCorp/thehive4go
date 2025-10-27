# \TaskLogAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAttachmentsToTaskLog**](TaskLogAPI.md#AddAttachmentsToTaskLog) | **Post** /api/v1/log/{logId}/attachments |
[**CreateTaskLog**](TaskLogAPI.md#CreateTaskLog) | **Post** /api/v1/task/{taskId}/log |
[**DeleteAttachmentFromTaskLog**](TaskLogAPI.md#DeleteAttachmentFromTaskLog) | **Delete** /api/v1/log/{logId}/attachments/{attachmentId} |
[**DeleteTaskLog**](TaskLogAPI.md#DeleteTaskLog) | **Delete** /api/v1/log/{logId} |
[**DownloadAttachmentFromLog**](TaskLogAPI.md#DownloadAttachmentFromLog) | **Get** /api/v1/log/{logId}/attachment/{attachmentId}/download |
[**GetAttachmentFromLog**](TaskLogAPI.md#GetAttachmentFromLog) | **Get** /api/v1/log/{caseId}/attachment/{attachmentId} |
[**GetAttachmentFromObservable**](TaskLogAPI.md#GetAttachmentFromObservable) | **Get** /api/v1/observable/{observableId}/attachment/{attachmentId} |
[**UpdateTaskLog**](TaskLogAPI.md#UpdateTaskLog) | **Patch** /api/v1/log/{logId} |



## AddAttachmentsToTaskLog

> AddAttachmentsToTaskLog(ctx, logId).Attachments(attachments).Execute()



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
	logId := "~354" // string |
	attachments := []*os.File{"TODO"} // []*os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaskLogAPI.AddAttachmentsToTaskLog(context.Background(), logId).Attachments(attachments).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskLogAPI.AddAttachmentsToTaskLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiAddAttachmentsToTaskLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attachments** | **[]*os.File** |  |

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTaskLog

> OutputLog CreateTaskLog(ctx, taskId).InputCreateLog(inputCreateLog).Execute()



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
	taskId := "~354" // string |
	inputCreateLog := *openapiclient.NewInputCreateLog("Message_example") // InputCreateLog |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskLogAPI.CreateTaskLog(context.Background(), taskId).InputCreateLog(inputCreateLog).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskLogAPI.CreateTaskLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTaskLog`: OutputLog
	fmt.Fprintf(os.Stdout, "Response from `TaskLogAPI.CreateTaskLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaskLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputCreateLog** | [**InputCreateLog**](InputCreateLog.md) |  |

### Return type

[**OutputLog**](OutputLog.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAttachmentFromTaskLog

> DeleteAttachmentFromTaskLog(ctx, logId, attachmentId).Execute()



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
	logId := "~354" // string |
	attachmentId := "~354" // string |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaskLogAPI.DeleteAttachmentFromTaskLog(context.Background(), logId, attachmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskLogAPI.DeleteAttachmentFromTaskLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logId** | **string** |  |
**attachmentId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAttachmentFromTaskLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTaskLog

> DeleteTaskLog(ctx, logId).Execute()



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
	logId := "~354" // string |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaskLogAPI.DeleteTaskLog(context.Background(), logId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskLogAPI.DeleteTaskLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTaskLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadAttachmentFromLog

> *os.File DownloadAttachmentFromLog(ctx, logId, attachmentId).Execute()



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
	logId := "~354" // string |
	attachmentId := "~354" // string |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskLogAPI.DownloadAttachmentFromLog(context.Background(), logId, attachmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskLogAPI.DownloadAttachmentFromLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadAttachmentFromLog`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `TaskLogAPI.DownloadAttachmentFromLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logId** | **string** |  |
**attachmentId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadAttachmentFromLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttachmentFromLog

> *os.File GetAttachmentFromLog(ctx, caseId, attachmentId).IfNoneMatch(ifNoneMatch).Execute()



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
	caseId := "~354" // string |
	attachmentId := "~354" // string |
	ifNoneMatch := "ifNoneMatch_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskLogAPI.GetAttachmentFromLog(context.Background(), caseId, attachmentId).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskLogAPI.GetAttachmentFromLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachmentFromLog`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `TaskLogAPI.GetAttachmentFromLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  |
**attachmentId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentFromLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifNoneMatch** | **string** |  |

### Return type

[***os.File**](*os.File.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttachmentFromObservable

> *os.File GetAttachmentFromObservable(ctx, observableId, attachmentId).IfNoneMatch(ifNoneMatch).Execute()



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
	observableId := "~354" // string |
	attachmentId := "~354" // string |
	ifNoneMatch := "ifNoneMatch_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskLogAPI.GetAttachmentFromObservable(context.Background(), observableId, attachmentId).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskLogAPI.GetAttachmentFromObservable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachmentFromObservable`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `TaskLogAPI.GetAttachmentFromObservable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**observableId** | **string** |  |
**attachmentId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentFromObservableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifNoneMatch** | **string** |  |

### Return type

[***os.File**](*os.File.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTaskLog

> UpdateTaskLog(ctx, logId).InputUpdateLog(inputUpdateLog).Execute()



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
	logId := "~354" // string |
	inputUpdateLog := *openapiclient.NewInputUpdateLog() // InputUpdateLog |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaskLogAPI.UpdateTaskLog(context.Background(), logId).InputUpdateLog(inputUpdateLog).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskLogAPI.UpdateTaskLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTaskLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputUpdateLog** | [**InputUpdateLog**](InputUpdateLog.md) |  |

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
