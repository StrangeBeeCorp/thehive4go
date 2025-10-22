# \CaseReportAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCaseReport**](CaseReportAPI.md#DeleteCaseReport) | **Delete** /api/v1/caseReport/{reportId} | 
[**DownloadCaseReport**](CaseReportAPI.md#DownloadCaseReport) | **Get** /api/v1/caseReport/{reportId}/download | 
[**GenerateCaseReport**](CaseReportAPI.md#GenerateCaseReport) | **Post** /api/v1/case/{caseId}/report | 
[**ListSupportedFormats**](CaseReportAPI.md#ListSupportedFormats) | **Get** /api/v1/caseReport/formats | 
[**RenderCaseReportTemplate**](CaseReportAPI.md#RenderCaseReportTemplate) | **Post** /api/v1/caseReport/render | 
[**RenderCaseReportTemplateGET**](CaseReportAPI.md#RenderCaseReportTemplateGET) | **Get** /api/v1/caseReport/render | 
[**UpdateCaseReport**](CaseReportAPI.md#UpdateCaseReport) | **Patch** /api/v1/caseReport/{reportId} | 
[**UploadCaseReport**](CaseReportAPI.md#UploadCaseReport) | **Post** /api/v1/case/{caseId}/report/upload | 
[**ViewCaseReport**](CaseReportAPI.md#ViewCaseReport) | **Get** /api/v1/caseReport/{reportId}/view | 



## DeleteCaseReport

> DeleteCaseReport(ctx, reportId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StrangeBee/TheHive4Go/thehive"
)

func main() {
	reportId := "~354" // string | 

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
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCaseReportRequest struct via the builder pattern


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


## DownloadCaseReport

> *os.File DownloadCaseReport(ctx, reportId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StrangeBee/TheHive4Go/thehive"
)

func main() {
	reportId := "~354" // string | 

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
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadCaseReportRequest struct via the builder pattern


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


## GenerateCaseReport

> OutputCaseReport GenerateCaseReport(ctx, caseId).InputGenerateCaseReport(inputGenerateCaseReport).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StrangeBee/TheHive4Go/thehive"
)

func main() {
	caseId := "~354" // string | 
	inputGenerateCaseReport := *openapiclient.NewInputGenerateCaseReport("~354", "Format_example") // InputGenerateCaseReport | 

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
**caseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateCaseReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputGenerateCaseReport** | [**InputGenerateCaseReport**](InputGenerateCaseReport.md) |  | 

### Return type

[**OutputCaseReport**](OutputCaseReport.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSupportedFormats

> OutputCaseReportFormats ListSupportedFormats(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StrangeBee/TheHive4Go/thehive"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaseReportAPI.ListSupportedFormats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseReportAPI.ListSupportedFormats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSupportedFormats`: OutputCaseReportFormats
	fmt.Fprintf(os.Stdout, "Response from `CaseReportAPI.ListSupportedFormats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSupportedFormatsRequest struct via the builder pattern


### Return type

[**OutputCaseReportFormats**](OutputCaseReportFormats.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenderCaseReportTemplate

> *os.File RenderCaseReportTemplate(ctx).InputRenderCaseReport(inputRenderCaseReport).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StrangeBee/TheHive4Go/thehive"
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

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenderCaseReportTemplateGET

> *os.File RenderCaseReportTemplateGET(ctx).Format(format).CaseReportTemplateId(caseReportTemplateId).CaseId(caseId).MaxElements(maxElements).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StrangeBee/TheHive4Go/thehive"
)

func main() {
	format := "format_example" // string | 
	caseReportTemplateId := "~354" // string | 
	caseId := "~354" // string |  (optional)
	maxElements := int32(56) // int32 |  (optional)

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
 **format** | **string** |  | 
 **caseReportTemplateId** | **string** |  | 
 **caseId** | **string** |  | 
 **maxElements** | **int32** |  | 

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


## UpdateCaseReport

> UpdateCaseReport(ctx, reportId).File(file).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StrangeBee/TheHive4Go/thehive"
)

func main() {
	reportId := "~354" // string | 
	file := os.NewFile(1234, "some_file") // *os.File | 

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
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCaseReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** |  | 

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


## UploadCaseReport

> OutputCaseReport UploadCaseReport(ctx, caseId).File(file).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StrangeBee/TheHive4Go/thehive"
)

func main() {
	caseId := "~354" // string | 
	file := os.NewFile(1234, "some_file") // *os.File | 

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
**caseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadCaseReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** |  | 

### Return type

[**OutputCaseReport**](OutputCaseReport.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewCaseReport

> *os.File ViewCaseReport(ctx, reportId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StrangeBee/TheHive4Go/thehive"
)

func main() {
	reportId := "~354" // string | 

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
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewCaseReportRequest struct via the builder pattern


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

