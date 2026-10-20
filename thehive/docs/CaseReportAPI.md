# \CaseReportAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCaseReport**](CaseReportAPI.md#DeleteCaseReport) | **Delete** /api/v1/caseReport/{reportId} | Delete a case report
[**DownloadCaseReport**](CaseReportAPI.md#DownloadCaseReport) | **Get** /api/v1/caseReport/{reportId}/download | Download a case report file
[**GenerateCaseReport**](CaseReportAPI.md#GenerateCaseReport) | **Post** /api/v1/case/{caseId}/report | Generate a case report from a template
[**RenderCaseReportTemplate**](CaseReportAPI.md#RenderCaseReportTemplate) | **Post** /api/v1/caseReport/render | Preview a case report without saving
[**RenderCaseReportTemplateGET**](CaseReportAPI.md#RenderCaseReportTemplateGET) | **Get** /api/v1/caseReport/render | Preview a case report without saving
[**UpdateCaseReport**](CaseReportAPI.md#UpdateCaseReport) | **Patch** /api/v1/caseReport/{reportId} | Replace the content of a case report
[**UploadCaseReport**](CaseReportAPI.md#UploadCaseReport) | **Post** /api/v1/case/{caseId}/report/upload | Upload a case report file
[**ViewCaseReport**](CaseReportAPI.md#ViewCaseReport) | **Get** /api/v1/caseReport/{reportId}/view | View a case report in the browser



## DeleteCaseReport

> DeleteCaseReport(ctx, reportId).Execute()

Delete a case report



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
	reportId := "reportId_example" // string | The case report ID preceded by `~`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CaseReportAPI.DeleteCaseReport(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseReportAPI.DeleteCaseReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** | The case report ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCaseReportRequest struct via the builder pattern


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


## DownloadCaseReport

> *os.File DownloadCaseReport(ctx, reportId).Execute()

Download a case report file



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
	reportId := "reportId_example" // string | The case report ID preceded by `~`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaseReportAPI.DownloadCaseReport(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseReportAPI.DownloadCaseReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadCaseReport`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `CaseReportAPI.DownloadCaseReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** | The case report ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadCaseReportRequest struct via the builder pattern


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


## GenerateCaseReport

> OutputCaseReport GenerateCaseReport(ctx, caseId).InputGenerateCaseReport(inputGenerateCaseReport).Execute()

Generate a case report from a template



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
	caseId := "caseId_example" // string | The case ID preceded by `~`, or the case number.
	inputGenerateCaseReport := *openapiclient.NewInputGenerateCaseReport("CaseReportTemplateId_example", "Format_example") // InputGenerateCaseReport |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaseReportAPI.GenerateCaseReport(context.Background(), caseId).InputGenerateCaseReport(inputGenerateCaseReport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseReportAPI.GenerateCaseReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateCaseReport`: OutputCaseReport
	fmt.Fprintf(os.Stdout, "Response from `CaseReportAPI.GenerateCaseReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** | The case ID preceded by &#x60;~&#x60;, or the case number. |

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateCaseReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputGenerateCaseReport** | [**InputGenerateCaseReport**](InputGenerateCaseReport.md) |  |

### Return type

[**OutputCaseReport**](OutputCaseReport.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenderCaseReportTemplate

> *os.File RenderCaseReportTemplate(ctx).InputRenderCaseReport(inputRenderCaseReport).Execute()

Preview a case report without saving



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
	inputRenderCaseReport := *openapiclient.NewInputRenderCaseReport("Format_example") // InputRenderCaseReport |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaseReportAPI.RenderCaseReportTemplate(context.Background()).InputRenderCaseReport(inputRenderCaseReport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseReportAPI.RenderCaseReportTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenderCaseReportTemplate`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `CaseReportAPI.RenderCaseReportTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRenderCaseReportTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputRenderCaseReport** | [**InputRenderCaseReport**](InputRenderCaseReport.md) |  |

### Return type

[***os.File**](*os.File.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenderCaseReportTemplateGET

> *os.File RenderCaseReportTemplateGET(ctx).Format(format).CaseReportTemplateId(caseReportTemplateId).CaseId(caseId).MaxElements(maxElements).Execute()

Preview a case report without saving



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
	format := "format_example" // string | Output format for the rendered report. Accepted values: `html`, `markdown`, `word`.
	caseReportTemplateId := "caseReportTemplateId_example" // string | ID of the case report template to use. Must match an existing template. Use `POST /api/v1/query` with `listCaseReportTemplate` to retrieve available templates and their IDs.
	caseId := "caseId_example" // string | ID of the case to use for data. Use `POST /api/v1/query` with `listCase` to retrieve available cases and their IDs. If omitted, TheHive uses fake data to fill the template. (optional)
	maxElements := int32(56) // int32 | Maximum number of elements displayed per widget. Doesn't affect the number of elements included when saving or downloading the report. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaseReportAPI.RenderCaseReportTemplateGET(context.Background()).Format(format).CaseReportTemplateId(caseReportTemplateId).CaseId(caseId).MaxElements(maxElements).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseReportAPI.RenderCaseReportTemplateGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenderCaseReportTemplateGET`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `CaseReportAPI.RenderCaseReportTemplateGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRenderCaseReportTemplateGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** | Output format for the rendered report. Accepted values: &#x60;html&#x60;, &#x60;markdown&#x60;, &#x60;word&#x60;. |
 **caseReportTemplateId** | **string** | ID of the case report template to use. Must match an existing template. Use &#x60;POST /api/v1/query&#x60; with &#x60;listCaseReportTemplate&#x60; to retrieve available templates and their IDs. |
 **caseId** | **string** | ID of the case to use for data. Use &#x60;POST /api/v1/query&#x60; with &#x60;listCase&#x60; to retrieve available cases and their IDs. If omitted, TheHive uses fake data to fill the template. |
 **maxElements** | **int32** | Maximum number of elements displayed per widget. Doesn&#39;t affect the number of elements included when saving or downloading the report. |

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


## UpdateCaseReport

> UpdateCaseReport(ctx, reportId).File(file).Execute()

Replace the content of a case report



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
	reportId := "reportId_example" // string | The case report ID preceded by `~`.
	file := os.NewFile(1234, "some_file") // *os.File | The new report file to replace the existing content.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CaseReportAPI.UpdateCaseReport(context.Background(), reportId).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseReportAPI.UpdateCaseReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** | The case report ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCaseReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | The new report file to replace the existing content. |

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadCaseReport

> OutputCaseReport UploadCaseReport(ctx, caseId).File(file).Execute()

Upload a case report file



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
	caseId := "caseId_example" // string | The case ID preceded by `~`, or the case number.
	file := os.NewFile(1234, "some_file") // *os.File | The report file to attach to the case.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaseReportAPI.UploadCaseReport(context.Background(), caseId).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseReportAPI.UploadCaseReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadCaseReport`: OutputCaseReport
	fmt.Fprintf(os.Stdout, "Response from `CaseReportAPI.UploadCaseReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** | The case ID preceded by &#x60;~&#x60;, or the case number. |

### Other Parameters

Other parameters are passed through a pointer to a apiUploadCaseReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | The report file to attach to the case. |

### Return type

[**OutputCaseReport**](OutputCaseReport.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewCaseReport

> *os.File ViewCaseReport(ctx, reportId).Execute()

View a case report in the browser



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
	reportId := "reportId_example" // string | The case report ID preceded by `~`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaseReportAPI.ViewCaseReport(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseReportAPI.ViewCaseReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ViewCaseReport`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `CaseReportAPI.ViewCaseReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** | The case report ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiViewCaseReportRequest struct via the builder pattern


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
