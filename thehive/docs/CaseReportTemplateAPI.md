# \CaseReportTemplateAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAttachmentToCaseReportTemplate**](CaseReportTemplateAPI.md#CreateAttachmentToCaseReportTemplate) | **Post** /api/v1/caseReportTemplate/{templateId}/attachment |
[**CreateCaseReportTemplate**](CaseReportTemplateAPI.md#CreateCaseReportTemplate) | **Post** /api/v1/caseReportTemplate |
[**DeleteAttachmentForCaseReportTemplate**](CaseReportTemplateAPI.md#DeleteAttachmentForCaseReportTemplate) | **Delete** /api/v1/caseReportTemplate/{templateId}/attachment/{attachmentId} |
[**DeleteCaseReportTemplate**](CaseReportTemplateAPI.md#DeleteCaseReportTemplate) | **Delete** /api/v1/caseReportTemplate/{idOrName} |
[**DownloadAttachmentForCaseReportTemplate**](CaseReportTemplateAPI.md#DownloadAttachmentForCaseReportTemplate) | **Get** /api/v1/caseReportTemplate/{templateId}/attachment/{attachmentId}/download |
[**GetAttachmentForCaseReportTemplate**](CaseReportTemplateAPI.md#GetAttachmentForCaseReportTemplate) | **Get** /api/v1/caseReportTemplate/{templateId}/attachment/{attachmentId} |
[**GetCaseReportTemplate**](CaseReportTemplateAPI.md#GetCaseReportTemplate) | **Get** /api/v1/caseReportTemplate/{templateId} |
[**GetCaseReportTemplateOptions**](CaseReportTemplateAPI.md#GetCaseReportTemplateOptions) | **Get** /api/v1/caseReportTemplate/_info |
[**UpdateCaseReportTemplate**](CaseReportTemplateAPI.md#UpdateCaseReportTemplate) | **Patch** /api/v1/caseReportTemplate/{idOrName} |



## CreateAttachmentToCaseReportTemplate

> OutputAttachments CreateAttachmentToCaseReportTemplate(ctx, templateId).Attachments(attachments).CanRename(canRename).Execute()



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
	templateId := "~354" // string |
	attachments := []*os.File{"TODO"} // []*os.File |
	canRename := true // bool | If set to `true`, the files can be renamed if they already exist with the same name (optional)

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
**templateId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAttachmentToCaseReportTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attachments** | **[]*os.File** |  |
 **canRename** | **bool** | If set to &#x60;true&#x60;, the files can be renamed if they already exist with the same name |

### Return type

[**OutputAttachments**](OutputAttachments.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCaseReportTemplate

> OutputCaseReportTemplate CreateCaseReportTemplate(ctx).InputCreateCaseReportTemplate(inputCreateCaseReportTemplate).Execute()





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

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAttachmentForCaseReportTemplate

> DeleteAttachmentForCaseReportTemplate(ctx, templateId, attachmentId).Execute()



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
	templateId := "~354" // string |
	attachmentId := "~354" // string |

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
**templateId** | **string** |  |
**attachmentId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAttachmentForCaseReportTemplateRequest struct via the builder pattern


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


## DeleteCaseReportTemplate

> DeleteCaseReportTemplate(ctx, idOrName).Execute()



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
	idOrName := "~354" // string |

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
**idOrName** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCaseReportTemplateRequest struct via the builder pattern


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


## DownloadAttachmentForCaseReportTemplate

> *os.File DownloadAttachmentForCaseReportTemplate(ctx, templateId, attachmentId).Execute()



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
	templateId := "~354" // string |
	attachmentId := "~354" // string |

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
**templateId** | **string** |  |
**attachmentId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadAttachmentForCaseReportTemplateRequest struct via the builder pattern


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


## GetAttachmentForCaseReportTemplate

> *os.File GetAttachmentForCaseReportTemplate(ctx, templateId, attachmentId).IfNoneMatch(ifNoneMatch).Execute()



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
	templateId := "~354" // string |
	attachmentId := "~354" // string |
	ifNoneMatch := "ifNoneMatch_example" // string |  (optional)

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
**templateId** | **string** |  |
**attachmentId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentForCaseReportTemplateRequest struct via the builder pattern


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


## GetCaseReportTemplate

> OutputCaseReportTemplate GetCaseReportTemplate(ctx, templateId).Execute()



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
	templateId := "~354" // string |

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
**templateId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetCaseReportTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputCaseReportTemplate**](OutputCaseReportTemplate.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCaseReportTemplateOptions

> OutputCaseReportTemplateObjects GetCaseReportTemplateOptions(ctx).Execute()



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

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCaseReportTemplate

> UpdateCaseReportTemplate(ctx, idOrName).InputUpdateCaseReportTemplate(inputUpdateCaseReportTemplate).Execute()



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
	idOrName := "~354" // string |
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
**idOrName** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCaseReportTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputUpdateCaseReportTemplate** | [**InputUpdateCaseReportTemplate**](InputUpdateCaseReportTemplate.md) |  |

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
