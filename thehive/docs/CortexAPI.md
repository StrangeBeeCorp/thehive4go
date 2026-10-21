# \CortexAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountActionsScopedToACaseOrAlert**](CortexAPI.md#CountActionsScopedToACaseOrAlert) | **Get** /api/v1/responder-execution/{scope}/{rootId}/count | Count responder executions for a case or alert
[**CreateAnAction**](CortexAPI.md#CreateAnAction) | **Post** /api/v1/connector/cortex/action | Run a responder
[**CreateAnalyzerTemplate**](CortexAPI.md#CreateAnalyzerTemplate) | **Post** /api/v1/connector/cortex/analyzer/template | Create an analyzer template
[**CreateCortexJob**](CortexAPI.md#CreateCortexJob) | **Post** /api/v1/connector/cortex/job | Run an analyzer on an observable
[**CreateSeveralActions**](CortexAPI.md#CreateSeveralActions) | **Post** /api/v1/connector/cortex/actions | Run responders in bulk
[**DeleteAnalyzerTemplate**](CortexAPI.md#DeleteAnalyzerTemplate) | **Delete** /api/v1/connector/cortex/analyzer/template/{analyzerTemplateId} | Delete an analyzer template
[**GetAnalyzerById**](CortexAPI.md#GetAnalyzerById) | **Get** /api/v1/connector/cortex/analyzer/{analyzerId} | Get an analyzer
[**GetAnalyzerTemplateContent**](CortexAPI.md#GetAnalyzerTemplateContent) | **Get** /api/v1/connector/cortex/analyzer/template/content/{analyzerId} | Get an analyzer template&#39;s content
[**GetCortexJob**](CortexAPI.md#GetCortexJob) | **Get** /api/v1/connector/cortex/job/{jobId} | Get a Cortex job
[**ImportAnalyzerTemplates**](CortexAPI.md#ImportAnalyzerTemplates) | **Post** /api/v1/connector/cortex/analyzer/template/_import | Import analyzer templates in bulk
[**ListActionsScopedToACaseOrAlert**](CortexAPI.md#ListActionsScopedToACaseOrAlert) | **Get** /api/v1/responder-execution/{scope}/{rootId} | List responder executions for a case or alert
[**ListAnalyzers**](CortexAPI.md#ListAnalyzers) | **Get** /api/v1/connector/cortex/analyzer | List available analyzers
[**ListAnalyzersByType**](CortexAPI.md#ListAnalyzersByType) | **Get** /api/v1/connector/cortex/analyzer/type/{dataType} | List analyzers by observable data type
[**ListResponders**](CortexAPI.md#ListResponders) | **Get** /api/v1/connector/cortex/responder/{entityType}/{entityId} | List responders available for an entity
[**ListRespondersForEntityType**](CortexAPI.md#ListRespondersForEntityType) | **Get** /api/v1/connector/cortex/responders/{entityType} | List responders available for an entity type
[**UpdateAnalyzerTemplate**](CortexAPI.md#UpdateAnalyzerTemplate) | **Patch** /api/v1/connector/cortex/analyzer/template/{analyzerTemplateId} | Update an analyzer template



## CountActionsScopedToACaseOrAlert

> int32 CountActionsScopedToACaseOrAlert(ctx, scope, rootId).Filter(filter).Execute()

Count responder executions for a case or alert



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
	scope := openapiclient.ActionScope("case") // ActionScope | Type of the target: `case` for a case, or `alert` for an alert.
	rootId := "rootId_example" // string | The case ID preceded by `~` or the case number when `scope` is `case`, or the alert ID preceded by `~` when `scope` is `alert`.
	filter := map[string]interface{}{ ... } // map[string]interface{} | Filter expression, using the JSON syntax of the `filter` operation of `POST /api/v1/query` without its `_name` field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CortexAPI.CountActionsScopedToACaseOrAlert(context.Background(), scope, rootId).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CortexAPI.CountActionsScopedToACaseOrAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountActionsScopedToACaseOrAlert`: int32
	fmt.Fprintf(os.Stdout, "Response from `CortexAPI.CountActionsScopedToACaseOrAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scope** | [**ActionScope**](.md) | Type of the target: &#x60;case&#x60; for a case, or &#x60;alert&#x60; for an alert. |
**rootId** | **string** | The case ID preceded by &#x60;~&#x60; or the case number when &#x60;scope&#x60; is &#x60;case&#x60;, or the alert ID preceded by &#x60;~&#x60; when &#x60;scope&#x60; is &#x60;alert&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiCountActionsScopedToACaseOrAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | [**map[string]interface{}**](map[string]interface{}.md) | Filter expression, using the JSON syntax of the &#x60;filter&#x60; operation of &#x60;POST /api/v1/query&#x60; without its &#x60;_name&#x60; field. |

### Return type

**int32**

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAnAction

> OutputAction CreateAnAction(ctx).InputAction(inputAction).Execute()

Run a responder



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

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAnalyzerTemplate

> OutputAnalyzerTemplate CreateAnalyzerTemplate(ctx).InputAnalyzerTemplate(inputAnalyzerTemplate).Execute()

Create an analyzer template



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

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCortexJob

> OutputJob CreateCortexJob(ctx).InputJob(inputJob).Execute()

Run an analyzer on an observable



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

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSeveralActions

> []OutputAction CreateSeveralActions(ctx).InputAction(inputAction).Execute()

Run responders in bulk



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
	inputAction := []openapiclient.InputAction{*openapiclient.NewInputAction("ResponderId_example", "ObjectType_example", "ObjectId_example")} // []InputAction |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CortexAPI.CreateSeveralActions(context.Background()).InputAction(inputAction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CortexAPI.CreateSeveralActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSeveralActions`: []OutputAction
	fmt.Fprintf(os.Stdout, "Response from `CortexAPI.CreateSeveralActions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSeveralActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputAction** | [**[]InputAction**](InputAction.md) |  |

### Return type

[**[]OutputAction**](OutputAction.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAnalyzerTemplate

> DeleteAnalyzerTemplate(ctx, analyzerTemplateId).Execute()

Delete an analyzer template



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
	analyzerTemplateId := "analyzerTemplateId_example" // string | The analyzer template ID preceded by `~`, or the name of the analyzer it's linked to.

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
**analyzerTemplateId** | **string** | The analyzer template ID preceded by &#x60;~&#x60;, or the name of the analyzer it&#39;s linked to. |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnalyzerTemplateRequest struct via the builder pattern


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


## GetAnalyzerById

> OutputWorker GetAnalyzerById(ctx, analyzerId).Execute()

Get an analyzer



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
	analyzerId := "analyzerId_example" // string | The analyzer ID.

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
**analyzerId** | **string** | The analyzer ID. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyzerByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputWorker**](OutputWorker.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalyzerTemplateContent

> string GetAnalyzerTemplateContent(ctx, analyzerId).Execute()

Get an analyzer template's content



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
	analyzerId := "analyzerId_example" // string | The analyzer template ID preceded by `~`, or the name of the analyzer it's linked to.

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
**analyzerId** | **string** | The analyzer template ID preceded by &#x60;~&#x60;, or the name of the analyzer it&#39;s linked to. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyzerTemplateContentRequest struct via the builder pattern


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


## GetCortexJob

> OutputJob GetCortexJob(ctx, jobId).Execute()

Get a Cortex job



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
	jobId := "jobId_example" // string | The job ID preceded by `~`.

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
**jobId** | **string** | The job ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetCortexJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputJob**](OutputJob.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportAnalyzerTemplates

> map[string]bool ImportAnalyzerTemplates(ctx).Templates(templates).Execute()

Import analyzer templates in bulk



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
	templates := os.NewFile(1234, "some_file") // *os.File | ZIP archive containing one or more analyzer template files to import, such as the official templates archive available at `https://catalogs.download.strangebee.com/latest/zip/report-templates.zip?version=<thehive_version>`.

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
 **templates** | ***os.File** | ZIP archive containing one or more analyzer template files to import, such as the official templates archive available at &#x60;https://catalogs.download.strangebee.com/latest/zip/report-templates.zip?version&#x3D;&lt;thehive_version&gt;&#x60;. |

### Return type

**map[string]bool**

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListActionsScopedToACaseOrAlert

> []OutputAction ListActionsScopedToACaseOrAlert(ctx, scope, rootId).Filter(filter).Sort(sort).PageFrom(pageFrom).PageTo(pageTo).Execute()

List responder executions for a case or alert



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
	scope := openapiclient.ActionScope("case") // ActionScope | Type of the target: `case` for a case, or `alert` for an alert.
	rootId := "rootId_example" // string | The case ID preceded by `~` or the case number when `scope` is `case`, or the alert ID preceded by `~` when `scope` is `alert`.
	filter := map[string]interface{}{ ... } // map[string]interface{} | Filter expression, using the JSON syntax of the `filter` operation of `POST /api/v1/query` without its `_name` field. (optional)
	sort := []interface{}{interface{}(123)} // []interface{} | Sort criteria, using the JSON syntax of the `_fields` value of the `sort` operation of `POST /api/v1/query`. (optional)
	pageFrom := int32(56) // int32 | Index of the first result to return, 0-based. Defaults to `0`. (optional)
	pageTo := int32(56) // int32 | Index of the last result to return, exclusive. Defaults to `pageFrom` plus 30. The range between `pageFrom` and `pageTo` can't exceed 300. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CortexAPI.ListActionsScopedToACaseOrAlert(context.Background(), scope, rootId).Filter(filter).Sort(sort).PageFrom(pageFrom).PageTo(pageTo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CortexAPI.ListActionsScopedToACaseOrAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListActionsScopedToACaseOrAlert`: []OutputAction
	fmt.Fprintf(os.Stdout, "Response from `CortexAPI.ListActionsScopedToACaseOrAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scope** | [**ActionScope**](.md) | Type of the target: &#x60;case&#x60; for a case, or &#x60;alert&#x60; for an alert. |
**rootId** | **string** | The case ID preceded by &#x60;~&#x60; or the case number when &#x60;scope&#x60; is &#x60;case&#x60;, or the alert ID preceded by &#x60;~&#x60; when &#x60;scope&#x60; is &#x60;alert&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiListActionsScopedToACaseOrAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | [**map[string]interface{}**](map[string]interface{}.md) | Filter expression, using the JSON syntax of the &#x60;filter&#x60; operation of &#x60;POST /api/v1/query&#x60; without its &#x60;_name&#x60; field. |
 **sort** | **[]interface{}** | Sort criteria, using the JSON syntax of the &#x60;_fields&#x60; value of the &#x60;sort&#x60; operation of &#x60;POST /api/v1/query&#x60;. |
 **pageFrom** | **int32** | Index of the first result to return, 0-based. Defaults to &#x60;0&#x60;. |
 **pageTo** | **int32** | Index of the last result to return, exclusive. Defaults to &#x60;pageFrom&#x60; plus 30. The range between &#x60;pageFrom&#x60; and &#x60;pageTo&#x60; can&#39;t exceed 300. |

### Return type

[**[]OutputAction**](OutputAction.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAnalyzers

> []AnalyzerOutput ListAnalyzers(ctx).Range_(range_).Execute()

List available analyzers



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
	range_ := "range__example" // string | Range of analyzers to return: `all` to return every available analyzer, or Cortex's `<start>-<end>` range syntax, for example `0-25`. Omitting this parameter defaults to Cortex's `0-10` range, returning only the first 10 analyzers. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CortexAPI.ListAnalyzers(context.Background()).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CortexAPI.ListAnalyzers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAnalyzers`: []AnalyzerOutput
	fmt.Fprintf(os.Stdout, "Response from `CortexAPI.ListAnalyzers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAnalyzersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Range of analyzers to return: &#x60;all&#x60; to return every available analyzer, or Cortex&#39;s &#x60;&lt;start&gt;-&lt;end&gt;&#x60; range syntax, for example &#x60;0-25&#x60;. Omitting this parameter defaults to Cortex&#39;s &#x60;0-10&#x60; range, returning only the first 10 analyzers. |

### Return type

[**[]AnalyzerOutput**](AnalyzerOutput.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAnalyzersByType

> []OutputWorker ListAnalyzersByType(ctx, dataType).Execute()

List analyzers by observable data type



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
	dataType := "dataType_example" // string | The observable data type analyzers must support.

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
**dataType** | **string** | The observable data type analyzers must support. |

### Other Parameters

Other parameters are passed through a pointer to a apiListAnalyzersByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OutputWorker**](OutputWorker.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResponders

> []ListRespondersOutput ListResponders(ctx, entityType, entityId).Execute()

List responders available for an entity



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
	entityType := "entityType_example" // string | Type of the target: `case` for a case, `case_artifact` for an observable, `case_task` for a task, `case_task_log` for a task log, or `alert` for an alert.
	entityId := "entityId_example" // string | The ID of the entity preceded by `~`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CortexAPI.ListResponders(context.Background(), entityType, entityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CortexAPI.ListResponders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListResponders`: []ListRespondersOutput
	fmt.Fprintf(os.Stdout, "Response from `CortexAPI.ListResponders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityType** | **string** | Type of the target: &#x60;case&#x60; for a case, &#x60;case_artifact&#x60; for an observable, &#x60;case_task&#x60; for a task, &#x60;case_task_log&#x60; for a task log, or &#x60;alert&#x60; for an alert. |
**entityId** | **string** | The ID of the entity preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiListRespondersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ListRespondersOutput**](ListRespondersOutput.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRespondersForEntityType

> []OutputWorker1 ListRespondersForEntityType(ctx, entityType).Execute()

List responders available for an entity type



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
	entityType := "entityType_example" // string | Type of the target: `case` for a case, `case_artifact` for an observable, `case_task` for a task, `case_task_log` for a task log, or `alert` for an alert.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CortexAPI.ListRespondersForEntityType(context.Background(), entityType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CortexAPI.ListRespondersForEntityType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRespondersForEntityType`: []OutputWorker1
	fmt.Fprintf(os.Stdout, "Response from `CortexAPI.ListRespondersForEntityType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityType** | **string** | Type of the target: &#x60;case&#x60; for a case, &#x60;case_artifact&#x60; for an observable, &#x60;case_task&#x60; for a task, &#x60;case_task_log&#x60; for a task log, or &#x60;alert&#x60; for an alert. |

### Other Parameters

Other parameters are passed through a pointer to a apiListRespondersForEntityTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OutputWorker1**](OutputWorker1.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAnalyzerTemplate

> OutputAnalyzerTemplate UpdateAnalyzerTemplate(ctx, analyzerTemplateId).InputAnalyzerTemplateUpdate(inputAnalyzerTemplateUpdate).Execute()

Update an analyzer template



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
	analyzerTemplateId := "analyzerTemplateId_example" // string | The analyzer template ID preceded by `~`, or the name of the analyzer it's linked to.
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
**analyzerTemplateId** | **string** | The analyzer template ID preceded by &#x60;~&#x60;, or the name of the analyzer it&#39;s linked to. |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAnalyzerTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputAnalyzerTemplateUpdate** | [**InputAnalyzerTemplateUpdate**](InputAnalyzerTemplateUpdate.md) |  |

### Return type

[**OutputAnalyzerTemplate**](OutputAnalyzerTemplate.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
