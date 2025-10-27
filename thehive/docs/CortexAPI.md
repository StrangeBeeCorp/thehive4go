# \CortexAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAnAction**](CortexAPI.md#CreateAnAction) | **Post** /api/v1/connector/cortex/action |
[**CreateAnalyzerTemplate**](CortexAPI.md#CreateAnalyzerTemplate) | **Post** /api/v1/connector/cortex/analyzer/template |
[**CreateCortexJob**](CortexAPI.md#CreateCortexJob) | **Post** /api/v1/connector/cortex/job |
[**DeleteAnalyzerTemplate**](CortexAPI.md#DeleteAnalyzerTemplate) | **Delete** /api/v1/connector/cortex/analyzer/template/{analyzerTemplateId} |
[**GetActionByEntity**](CortexAPI.md#GetActionByEntity) | **Get** /api/v1/connector/cortex/action/{entityType}/{entityId} |
[**GetAnalyzerById**](CortexAPI.md#GetAnalyzerById) | **Get** /api/v1/connector/cortex/analyzer/{analyzerId} |
[**GetAnalyzerTemplateContent**](CortexAPI.md#GetAnalyzerTemplateContent) | **Get** /api/v1/connector/cortex/analyzer/template/content/{analyzerId} |
[**GetCortexJob**](CortexAPI.md#GetCortexJob) | **Get** /api/v1/connector/cortex/job/{jobId} |
[**ImportAnalyzerTemplates**](CortexAPI.md#ImportAnalyzerTemplates) | **Post** /api/v1/connector/cortex/analyzer/template/_import |
[**ListAnalyzers**](CortexAPI.md#ListAnalyzers) | **Get** /api/v1/connector/cortex/analyzer |
[**ListAnalyzersByType**](CortexAPI.md#ListAnalyzersByType) | **Get** /api/v1/connector/cortex/analyzer/type/{dataType} |
[**ListResponders**](CortexAPI.md#ListResponders) | **Get** /api/v1/connector/cortex/responder/{entityType}/{entityId} |
[**UpdateAnalyzerTemplate**](CortexAPI.md#UpdateAnalyzerTemplate) | **Patch** /api/v1/connector/cortex/analyzer/template/{analyzerTemplateId} |



## CreateAnAction

> OutputAction CreateAnAction(ctx).InputAction(inputAction).Execute()



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
	inputAction := *openapiclient.NewInputAction("ResponderId_example", "ObjectType_example", "ObjectId_example") // InputAction |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CortexAPI.CreateAnAction(context.Background()).InputAction(inputAction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CortexAPI.CreateAnAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAnAction`: OutputAction
	fmt.Fprintf(os.Stdout, "Response from `CortexAPI.CreateAnAction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAnActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputAction** | [**InputAction**](InputAction.md) |  |

### Return type

[**OutputAction**](OutputAction.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAnalyzerTemplate

> OutputAnalyzerTemplate CreateAnalyzerTemplate(ctx).InputAnalyzerTemplate(inputAnalyzerTemplate).Execute()



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
	inputAnalyzerTemplate := *openapiclient.NewInputAnalyzerTemplate("AnalyzerId_example", "Content_example") // InputAnalyzerTemplate |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CortexAPI.CreateAnalyzerTemplate(context.Background()).InputAnalyzerTemplate(inputAnalyzerTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CortexAPI.CreateAnalyzerTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAnalyzerTemplate`: OutputAnalyzerTemplate
	fmt.Fprintf(os.Stdout, "Response from `CortexAPI.CreateAnalyzerTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAnalyzerTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputAnalyzerTemplate** | [**InputAnalyzerTemplate**](InputAnalyzerTemplate.md) |  |

### Return type

[**OutputAnalyzerTemplate**](OutputAnalyzerTemplate.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCortexJob

> OutputJob CreateCortexJob(ctx).InputJob(inputJob).Execute()



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
	inputJob := *openapiclient.NewInputJob("AnalyzerId_example", "CortexId_example", "ArtifactId_example") // InputJob |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CortexAPI.CreateCortexJob(context.Background()).InputJob(inputJob).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CortexAPI.CreateCortexJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCortexJob`: OutputJob
	fmt.Fprintf(os.Stdout, "Response from `CortexAPI.CreateCortexJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCortexJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputJob** | [**InputJob**](InputJob.md) |  |

### Return type

[**OutputJob**](OutputJob.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAnalyzerTemplate

> DeleteAnalyzerTemplate(ctx, analyzerTemplateId).Execute()



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
	analyzerTemplateId := "~354" // string |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CortexAPI.DeleteAnalyzerTemplate(context.Background(), analyzerTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CortexAPI.DeleteAnalyzerTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analyzerTemplateId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnalyzerTemplateRequest struct via the builder pattern


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


## GetActionByEntity

> []OutputAction GetActionByEntity(ctx, entityType, entityId).Execute()



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
	entityType := "entityType_example" // string |
	entityId := "~354" // string |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CortexAPI.GetActionByEntity(context.Background(), entityType, entityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CortexAPI.GetActionByEntity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActionByEntity`: []OutputAction
	fmt.Fprintf(os.Stdout, "Response from `CortexAPI.GetActionByEntity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityType** | **string** |  |
**entityId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetActionByEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]OutputAction**](OutputAction.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalyzerById

> OutputWorker GetAnalyzerById(ctx, analyzerId).Execute()



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
	analyzerId := "analyzerId_example" // string |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CortexAPI.GetAnalyzerById(context.Background(), analyzerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CortexAPI.GetAnalyzerById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnalyzerById`: OutputWorker
	fmt.Fprintf(os.Stdout, "Response from `CortexAPI.GetAnalyzerById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analyzerId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyzerByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputWorker**](OutputWorker.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalyzerTemplateContent

> string GetAnalyzerTemplateContent(ctx, analyzerId).Execute()



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
	analyzerId := "~354" // string |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CortexAPI.GetAnalyzerTemplateContent(context.Background(), analyzerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CortexAPI.GetAnalyzerTemplateContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnalyzerTemplateContent`: string
	fmt.Fprintf(os.Stdout, "Response from `CortexAPI.GetAnalyzerTemplateContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analyzerId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyzerTemplateContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCortexJob

> OutputJob GetCortexJob(ctx, jobId).Execute()



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
	jobId := "~354" // string |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CortexAPI.GetCortexJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CortexAPI.GetCortexJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCortexJob`: OutputJob
	fmt.Fprintf(os.Stdout, "Response from `CortexAPI.GetCortexJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetCortexJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputJob**](OutputJob.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportAnalyzerTemplates

> map[string]bool ImportAnalyzerTemplates(ctx).Templates(templates).Execute()



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
	templates := os.NewFile(1234, "some_file") // *os.File |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CortexAPI.ImportAnalyzerTemplates(context.Background()).Templates(templates).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CortexAPI.ImportAnalyzerTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportAnalyzerTemplates`: map[string]bool
	fmt.Fprintf(os.Stdout, "Response from `CortexAPI.ImportAnalyzerTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportAnalyzerTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templates** | ***os.File** |  |

### Return type

**map[string]bool**

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAnalyzers

> []OutputWorker ListAnalyzers(ctx).Range_(range_).Execute()



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
	range_ := "range__example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CortexAPI.ListAnalyzers(context.Background()).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CortexAPI.ListAnalyzers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAnalyzers`: []OutputWorker
	fmt.Fprintf(os.Stdout, "Response from `CortexAPI.ListAnalyzers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAnalyzersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** |  |

### Return type

[**[]OutputWorker**](OutputWorker.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAnalyzersByType

> []OutputWorker ListAnalyzersByType(ctx, dataType).Execute()



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
	dataType := "dataType_example" // string |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CortexAPI.ListAnalyzersByType(context.Background(), dataType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CortexAPI.ListAnalyzersByType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAnalyzersByType`: []OutputWorker
	fmt.Fprintf(os.Stdout, "Response from `CortexAPI.ListAnalyzersByType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataType** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiListAnalyzersByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OutputWorker**](OutputWorker.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResponders

> []OutputWorker ListResponders(ctx, entityType, entityId).Execute()



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
	entityType := "entityType_example" // string |
	entityId := "~354" // string |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CortexAPI.ListResponders(context.Background(), entityType, entityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CortexAPI.ListResponders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListResponders`: []OutputWorker
	fmt.Fprintf(os.Stdout, "Response from `CortexAPI.ListResponders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityType** | **string** |  |
**entityId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiListRespondersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]OutputWorker**](OutputWorker.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAnalyzerTemplate

> OutputAnalyzerTemplate UpdateAnalyzerTemplate(ctx, analyzerTemplateId).InputAnalyzerTemplateUpdate(inputAnalyzerTemplateUpdate).Execute()



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
	analyzerTemplateId := "~354" // string |
	inputAnalyzerTemplateUpdate := *openapiclient.NewInputAnalyzerTemplateUpdate() // InputAnalyzerTemplateUpdate |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CortexAPI.UpdateAnalyzerTemplate(context.Background(), analyzerTemplateId).InputAnalyzerTemplateUpdate(inputAnalyzerTemplateUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CortexAPI.UpdateAnalyzerTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAnalyzerTemplate`: OutputAnalyzerTemplate
	fmt.Fprintf(os.Stdout, "Response from `CortexAPI.UpdateAnalyzerTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analyzerTemplateId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAnalyzerTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputAnalyzerTemplateUpdate** | [**InputAnalyzerTemplateUpdate**](InputAnalyzerTemplateUpdate.md) |  |

### Return type

[**OutputAnalyzerTemplate**](OutputAnalyzerTemplate.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
