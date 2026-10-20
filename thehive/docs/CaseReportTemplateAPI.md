# \CaseReportTemplateAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAttachmentToCaseReportTemplate**](CaseReportTemplateAPI.md#CreateAttachmentToCaseReportTemplate) | **Post** /api/v1/caseReportTemplate/{templateId}/attachment | Add an attachment to a case report template
[**CreateCaseReportTemplate**](CaseReportTemplateAPI.md#CreateCaseReportTemplate) | **Post** /api/v1/caseReportTemplate | Create a case report template
[**DeleteAttachmentForCaseReportTemplate**](CaseReportTemplateAPI.md#DeleteAttachmentForCaseReportTemplate) | **Delete** /api/v1/caseReportTemplate/{templateId}/attachment/{attachmentId} | Remove an attachment from a case report template
[**DeleteCaseReportTemplate**](CaseReportTemplateAPI.md#DeleteCaseReportTemplate) | **Delete** /api/v1/caseReportTemplate/{idOrName} | Delete a case report template
[**DownloadAttachmentForCaseReportTemplate**](CaseReportTemplateAPI.md#DownloadAttachmentForCaseReportTemplate) | **Get** /api/v1/caseReportTemplate/{templateId}/attachment/{attachmentId}/download | Download an attachment from a case report template
[**ExportCaseReportTemplateAsArchive**](CaseReportTemplateAPI.md#ExportCaseReportTemplateAsArchive) | **Get** /api/v1/caseReportTemplate/{caseReportTemplateId}/export | Export a case report template as a file
[**GetAttachmentForCaseReportTemplate**](CaseReportTemplateAPI.md#GetAttachmentForCaseReportTemplate) | **Get** /api/v1/caseReportTemplate/{templateId}/attachment/{attachmentId} | Get an attachment from a case report template
[**GetCaseReportTemplate**](CaseReportTemplateAPI.md#GetCaseReportTemplate) | **Get** /api/v1/caseReportTemplate/{templateId} | Get a case report template
[**GetCaseReportTemplateOptions**](CaseReportTemplateAPI.md#GetCaseReportTemplateOptions) | **Get** /api/v1/caseReportTemplate/_info | List available widget types and field options
[**ImportCaseReportTemplateFromFile**](CaseReportTemplateAPI.md#ImportCaseReportTemplateFromFile) | **Post** /api/v1/caseReportTemplate/import | Import a case report template from a file
[**UpdateCaseReportTemplate**](CaseReportTemplateAPI.md#UpdateCaseReportTemplate) | **Patch** /api/v1/caseReportTemplate/{idOrName} | Update a case report template



## CreateAttachmentToCaseReportTemplate

> OutputAttachments CreateAttachmentToCaseReportTemplate(ctx, templateId).Attachments(attachments).CanRename(canRename).Execute()

Add an attachment to a case report template



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
	templateId := "templateId_example" // string | The case report template ID preceded by `~`.
	attachments := []*os.File{"TODO"} // []*os.File | Files to attach. Each part is a multipart form file field.
	canRename := true // bool | When `true`, renames files that already exist with the same name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaseReportTemplateAPI.CreateAttachmentToCaseReportTemplate(context.Background(), templateId).Attachments(attachments).CanRename(canRename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseReportTemplateAPI.CreateAttachmentToCaseReportTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAttachmentToCaseReportTemplate`: OutputAttachments
	fmt.Fprintf(os.Stdout, "Response from `CaseReportTemplateAPI.CreateAttachmentToCaseReportTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The case report template ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAttachmentToCaseReportTemplateRequest struct via the builder pattern


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


## CreateCaseReportTemplate

> OutputCaseReportTemplate CreateCaseReportTemplate(ctx).InputCreateCaseReportTemplate(inputCreateCaseReportTemplate).Execute()

Create a case report template



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
	inputCreateCaseReportTemplate := *openapiclient.NewInputCreateCaseReportTemplate("Title_example", *openapiclient.NewCaseReportTemplateDefinition(*openapiclient.NewI18n())) // InputCreateCaseReportTemplate |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaseReportTemplateAPI.CreateCaseReportTemplate(context.Background()).InputCreateCaseReportTemplate(inputCreateCaseReportTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseReportTemplateAPI.CreateCaseReportTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCaseReportTemplate`: OutputCaseReportTemplate
	fmt.Fprintf(os.Stdout, "Response from `CaseReportTemplateAPI.CreateCaseReportTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCaseReportTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputCreateCaseReportTemplate** | [**InputCreateCaseReportTemplate**](InputCreateCaseReportTemplate.md) |  |

### Return type

[**OutputCaseReportTemplate**](OutputCaseReportTemplate.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAttachmentForCaseReportTemplate

> DeleteAttachmentForCaseReportTemplate(ctx, templateId, attachmentId).Execute()

Remove an attachment from a case report template



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
	templateId := "templateId_example" // string | The case report template ID preceded by `~`.
	attachmentId := "attachmentId_example" // string | The attachment ID preceded by `~`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CaseReportTemplateAPI.DeleteAttachmentForCaseReportTemplate(context.Background(), templateId, attachmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseReportTemplateAPI.DeleteAttachmentForCaseReportTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The case report template ID preceded by &#x60;~&#x60;. |
**attachmentId** | **string** | The attachment ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAttachmentForCaseReportTemplateRequest struct via the builder pattern


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


## DeleteCaseReportTemplate

> DeleteCaseReportTemplate(ctx, idOrName).Execute()

Delete a case report template



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
	idOrName := "idOrName_example" // string | The case report template ID preceded by `~`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CaseReportTemplateAPI.DeleteCaseReportTemplate(context.Background(), idOrName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseReportTemplateAPI.DeleteCaseReportTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrName** | **string** | The case report template ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCaseReportTemplateRequest struct via the builder pattern


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


## DownloadAttachmentForCaseReportTemplate

> *os.File DownloadAttachmentForCaseReportTemplate(ctx, templateId, attachmentId).Execute()

Download an attachment from a case report template



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
	templateId := "templateId_example" // string | The case report template ID preceded by `~`.
	attachmentId := "attachmentId_example" // string | The attachment ID preceded by `~`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaseReportTemplateAPI.DownloadAttachmentForCaseReportTemplate(context.Background(), templateId, attachmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseReportTemplateAPI.DownloadAttachmentForCaseReportTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadAttachmentForCaseReportTemplate`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `CaseReportTemplateAPI.DownloadAttachmentForCaseReportTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The case report template ID preceded by &#x60;~&#x60;. |
**attachmentId** | **string** | The attachment ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadAttachmentForCaseReportTemplateRequest struct via the builder pattern


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


## ExportCaseReportTemplateAsArchive

> *os.File ExportCaseReportTemplateAsArchive(ctx, caseReportTemplateId).Execute()

Export a case report template as a file



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
	caseReportTemplateId := "caseReportTemplateId_example" // string | The case report template ID preceded by `~`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaseReportTemplateAPI.ExportCaseReportTemplateAsArchive(context.Background(), caseReportTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseReportTemplateAPI.ExportCaseReportTemplateAsArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportCaseReportTemplateAsArchive`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `CaseReportTemplateAPI.ExportCaseReportTemplateAsArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseReportTemplateId** | **string** | The case report template ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiExportCaseReportTemplateAsArchiveRequest struct via the builder pattern


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


## GetAttachmentForCaseReportTemplate

> *os.File GetAttachmentForCaseReportTemplate(ctx, templateId, attachmentId).IfNoneMatch(ifNoneMatch).Execute()

Get an attachment from a case report template



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
	templateId := "templateId_example" // string | The case report template ID preceded by `~`.
	attachmentId := "attachmentId_example" // string | The attachment ID preceded by `~`.
	ifNoneMatch := "ifNoneMatch_example" // string | ETag from a previous response. If it matches, returns `304 Not Modified` instead of the file content. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaseReportTemplateAPI.GetAttachmentForCaseReportTemplate(context.Background(), templateId, attachmentId).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseReportTemplateAPI.GetAttachmentForCaseReportTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachmentForCaseReportTemplate`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `CaseReportTemplateAPI.GetAttachmentForCaseReportTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The case report template ID preceded by &#x60;~&#x60;. |
**attachmentId** | **string** | The attachment ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentForCaseReportTemplateRequest struct via the builder pattern


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


## GetCaseReportTemplate

> OutputCaseReportTemplate GetCaseReportTemplate(ctx, templateId).Execute()

Get a case report template



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
	templateId := "templateId_example" // string | The case report template ID preceded by `~`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaseReportTemplateAPI.GetCaseReportTemplate(context.Background(), templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseReportTemplateAPI.GetCaseReportTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCaseReportTemplate`: OutputCaseReportTemplate
	fmt.Fprintf(os.Stdout, "Response from `CaseReportTemplateAPI.GetCaseReportTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The case report template ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetCaseReportTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputCaseReportTemplate**](OutputCaseReportTemplate.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCaseReportTemplateOptions

> OutputCaseReportTemplateObjects GetCaseReportTemplateOptions(ctx).Execute()

List available widget types and field options



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
	resp, r, err := apiClient.CaseReportTemplateAPI.GetCaseReportTemplateOptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseReportTemplateAPI.GetCaseReportTemplateOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCaseReportTemplateOptions`: OutputCaseReportTemplateObjects
	fmt.Fprintf(os.Stdout, "Response from `CaseReportTemplateAPI.GetCaseReportTemplateOptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCaseReportTemplateOptionsRequest struct via the builder pattern


### Return type

[**OutputCaseReportTemplateObjects**](OutputCaseReportTemplateObjects.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportCaseReportTemplateFromFile

> OutputCaseReportTemplate ImportCaseReportTemplateFromFile(ctx).Json(json).File(file).Execute()

Import a case report template from a file



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
	json := *openapiclient.NewImportCaseReportTemplate() // ImportCaseReportTemplate |
	file := os.NewFile(1234, "some_file") // *os.File |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaseReportTemplateAPI.ImportCaseReportTemplateFromFile(context.Background()).Json(json).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseReportTemplateAPI.ImportCaseReportTemplateFromFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportCaseReportTemplateFromFile`: OutputCaseReportTemplate
	fmt.Fprintf(os.Stdout, "Response from `CaseReportTemplateAPI.ImportCaseReportTemplateFromFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportCaseReportTemplateFromFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **json** | [**ImportCaseReportTemplate**](ImportCaseReportTemplate.md) |  |
 **file** | ***os.File** |  |

### Return type

[**OutputCaseReportTemplate**](OutputCaseReportTemplate.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCaseReportTemplate

> UpdateCaseReportTemplate(ctx, idOrName).InputUpdateCaseReportTemplate(inputUpdateCaseReportTemplate).Execute()

Update a case report template



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
	idOrName := "idOrName_example" // string | The case report template ID preceded by `~`.
	inputUpdateCaseReportTemplate := *openapiclient.NewInputUpdateCaseReportTemplate() // InputUpdateCaseReportTemplate |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CaseReportTemplateAPI.UpdateCaseReportTemplate(context.Background(), idOrName).InputUpdateCaseReportTemplate(inputUpdateCaseReportTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseReportTemplateAPI.UpdateCaseReportTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrName** | **string** | The case report template ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCaseReportTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputUpdateCaseReportTemplate** | [**InputUpdateCaseReportTemplate**](InputUpdateCaseReportTemplate.md) |  |

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
