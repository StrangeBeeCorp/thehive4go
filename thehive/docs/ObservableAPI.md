# \ObservableAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateObservableInAlert**](ObservableAPI.md#CreateObservableInAlert) | **Post** /api/v1/alert/{alertId}/observable |
[**CreateObservableInCase**](ObservableAPI.md#CreateObservableInCase) | **Post** /api/v1/case/{caseId}/observable |
[**DeleteObservable**](ObservableAPI.md#DeleteObservable) | **Delete** /api/v1/observable/{observableId} |
[**DownloadAttachmentFromObservable**](ObservableAPI.md#DownloadAttachmentFromObservable) | **Get** /api/v1/observable/{observableId}/attachment/{attachmentId}/download |
[**GetObservable**](ObservableAPI.md#GetObservable) | **Get** /api/v1/observable/{observableId} |
[**UpdateBulkOfObservables**](ObservableAPI.md#UpdateBulkOfObservables) | **Patch** /api/v1/observable/_bulk |
[**UpdateObservable**](ObservableAPI.md#UpdateObservable) | **Patch** /api/v1/observable/{observableId} |



## CreateObservableInAlert

> []OutputObservable CreateObservableInAlert(ctx, alertId).InputCreateObservable(inputCreateObservable).DataType(dataType).Execute()





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
	alertId := "~354" // string |
	inputCreateObservable := *openapiclient.NewInputCreateObservable("DataType_example") // InputCreateObservable |
	dataType := "dataType_example" // string | allow to set the dataType from the query parameters (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObservableAPI.CreateObservableInAlert(context.Background(), alertId).InputCreateObservable(inputCreateObservable).DataType(dataType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservableAPI.CreateObservableInAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateObservableInAlert`: []OutputObservable
	fmt.Fprintf(os.Stdout, "Response from `ObservableAPI.CreateObservableInAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateObservableInAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputCreateObservable** | [**InputCreateObservable**](InputCreateObservable.md) |  |
 **dataType** | **string** | allow to set the dataType from the query parameters |

### Return type

[**[]OutputObservable**](OutputObservable.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateObservableInCase

> []OutputObservable CreateObservableInCase(ctx, caseId).InputCreateObservable(inputCreateObservable).DataType(dataType).Execute()





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
	inputCreateObservable := *openapiclient.NewInputCreateObservable("DataType_example") // InputCreateObservable |
	dataType := "dataType_example" // string | allow to set the dataType from the query parameters (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObservableAPI.CreateObservableInCase(context.Background(), caseId).InputCreateObservable(inputCreateObservable).DataType(dataType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservableAPI.CreateObservableInCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateObservableInCase`: []OutputObservable
	fmt.Fprintf(os.Stdout, "Response from `ObservableAPI.CreateObservableInCase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateObservableInCaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputCreateObservable** | [**InputCreateObservable**](InputCreateObservable.md) |  |
 **dataType** | **string** | allow to set the dataType from the query parameters |

### Return type

[**[]OutputObservable**](OutputObservable.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteObservable

> DeleteObservable(ctx, observableId).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ObservableAPI.DeleteObservable(context.Background(), observableId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservableAPI.DeleteObservable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**observableId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteObservableRequest struct via the builder pattern


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


## DownloadAttachmentFromObservable

> *os.File DownloadAttachmentFromObservable(ctx, observableId, attachmentId).AsZip(asZip).Execute()





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
	asZip := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObservableAPI.DownloadAttachmentFromObservable(context.Background(), observableId, attachmentId).AsZip(asZip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservableAPI.DownloadAttachmentFromObservable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadAttachmentFromObservable`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ObservableAPI.DownloadAttachmentFromObservable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**observableId** | **string** |  |
**attachmentId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadAttachmentFromObservableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **asZip** | **bool** |  | [default to false]

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


## GetObservable

> OutputObservable GetObservable(ctx, observableId).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObservableAPI.GetObservable(context.Background(), observableId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservableAPI.GetObservable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObservable`: OutputObservable
	fmt.Fprintf(os.Stdout, "Response from `ObservableAPI.GetObservable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**observableId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetObservableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputObservable**](OutputObservable.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBulkOfObservables

> UpdateBulkOfObservables(ctx).InputUpdateObservableWithIds(inputUpdateObservableWithIds).Execute()



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
	inputUpdateObservableWithIds := *openapiclient.NewInputUpdateObservableWithIds([]string{"~354"}) // InputUpdateObservableWithIds |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ObservableAPI.UpdateBulkOfObservables(context.Background()).InputUpdateObservableWithIds(inputUpdateObservableWithIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservableAPI.UpdateBulkOfObservables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBulkOfObservablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputUpdateObservableWithIds** | [**InputUpdateObservableWithIds**](InputUpdateObservableWithIds.md) |  |

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


## UpdateObservable

> UpdateObservable(ctx, observableId).InputUpdateObservable(inputUpdateObservable).Execute()



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
	inputUpdateObservable := *openapiclient.NewInputUpdateObservable() // InputUpdateObservable |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ObservableAPI.UpdateObservable(context.Background(), observableId).InputUpdateObservable(inputUpdateObservable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservableAPI.UpdateObservable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**observableId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateObservableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputUpdateObservable** | [**InputUpdateObservable**](InputUpdateObservable.md) |  |

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
