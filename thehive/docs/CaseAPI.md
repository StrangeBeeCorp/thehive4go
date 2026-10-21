# \CaseAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAttachmentToCase**](CaseAPI.md#AddAttachmentToCase) | **Post** /api/v1/case/{caseId}/attachments | Add attachments to a case
[**AddLinkWithAnExternalURL**](CaseAPI.md#AddLinkWithAnExternalURL) | **Post** /api/v1/case/{caseId}/link/external/add | Link a case to an external URL
[**AddLinkWithAnotherCase**](CaseAPI.md#AddLinkWithAnotherCase) | **Post** /api/v1/case/{caseId}/link/case/add | Link a case to another case
[**ApplyCaseTemplateOnExistingCases**](CaseAPI.md#ApplyCaseTemplateOnExistingCases) | **Post** /api/v1/case/_bulk/caseTemplate | Apply a case template to existing cases
[**BulkManageCaseAccess**](CaseAPI.md#BulkManageCaseAccess) | **Post** /api/v1/case/_bulk/access | Set the access mode of multiple cases
[**BulkUpdateCase**](CaseAPI.md#BulkUpdateCase) | **Patch** /api/v1/case/_bulk | Update multiple cases
[**ChangeCaseOwningOrganisation**](CaseAPI.md#ChangeCaseOwningOrganisation) | **Post** /api/v1/case/{caseId}/owner | Change the owner organization of a case
[**CreateCase**](CaseAPI.md#CreateCase) | **Post** /api/v1/case | Create a case
[**DeleteACustomField**](CaseAPI.md#DeleteACustomField) | **Delete** /api/v1/case/customField/{cfId} | Remove a custom field from a case
[**DeleteCase**](CaseAPI.md#DeleteCase) | **Delete** /api/v1/case/{idOrName} | Delete a case
[**DeleteCaseAttachment**](CaseAPI.md#DeleteCaseAttachment) | **Delete** /api/v1/case/{caseId}/attachment/{attachmentId} | Remove an attachment from a case
[**DeleteLinkWithAnAnotherCase**](CaseAPI.md#DeleteLinkWithAnAnotherCase) | **Post** /api/v1/case/{caseId}/link/case/remove | Remove a link to another case
[**DeleteLinkWithAnExternalURL**](CaseAPI.md#DeleteLinkWithAnExternalURL) | **Post** /api/v1/case/{caseId}/link/external/remove | Remove a link to an external URL
[**DownloadAttachmentFromCase**](CaseAPI.md#DownloadAttachmentFromCase) | **Get** /api/v1/case/{caseId}/attachment/{attachmentId}/download | Download an attachment from a case
[**ExportCaseAsArchive**](CaseAPI.md#ExportCaseAsArchive) | **Get** /api/v1/case/{caseId}/export | Export a case as an archive
[**GetAllLinkTypes**](CaseAPI.md#GetAllLinkTypes) | **Get** /api/v1/case/link/types | List all case link types
[**GetAttachmentFromCase**](CaseAPI.md#GetAttachmentFromCase) | **Get** /api/v1/case/{caseId}/attachment/{attachmentId} | Get an attachment from a case
[**GetCase**](CaseAPI.md#GetCase) | **Get** /api/v1/case/{idOrName} | Get a case
[**GetCaseTimeline**](CaseAPI.md#GetCaseTimeline) | **Get** /api/v1/case/{caseId}/timeline | Get the timeline of a case
[**GetSimilarCases**](CaseAPI.md#GetSimilarCases) | **Get** /api/v1/case/{caseId}/links | List similar cases
[**GetSimilarObservablesBetweenACaseAndAnotherCaseOrAlert**](CaseAPI.md#GetSimilarObservablesBetweenACaseAndAnotherCaseOrAlert) | **Get** /api/v1/case/{caseId}/similar/{alertOrCaseId}/observables | List observables shared with another case or alert
[**ImportCaseFromFile**](CaseAPI.md#ImportCaseFromFile) | **Post** /api/v1/case/import | Create a case from an archive
[**ManageCaseAccess**](CaseAPI.md#ManageCaseAccess) | **Post** /api/v1/case/{caseId}/access | Set the access mode of a case
[**MergeCases**](CaseAPI.md#MergeCases) | **Post** /api/v1/case/_merge/{ids} | Merge cases into a new case
[**MergeSimilarObservablesOfThisCase**](CaseAPI.md#MergeSimilarObservablesOfThisCase) | **Post** /api/v1/case/{caseId}/observable/_merge | Deduplicate observables in a case
[**UnlinkAlertFromCase**](CaseAPI.md#UnlinkAlertFromCase) | **Delete** /api/v1/case/{caseId}/alert/{alertId} | Unlink an alert from a case
[**UpdateAttachment**](CaseAPI.md#UpdateAttachment) | **Patch** /api/v1/case/{caseId}/attachment/{attachmentId} | Update a case attachment
[**UpdateCase**](CaseAPI.md#UpdateCase) | **Patch** /api/v1/case/{idOrName} | Update a case



## AddAttachmentToCase

> OutputAttachments AddAttachmentToCase(ctx, caseId).Attachments(attachments).CanRename(canRename).Execute()

Add attachments to a case



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
	attachments := []*os.File{"TODO"} // []*os.File | Files to attach. Each part is a multipart form file field.
	canRename := true // bool | When `true`, renames files that already exist with the same name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaseAPI.AddAttachmentToCase(context.Background(), caseId).Attachments(attachments).CanRename(canRename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseAPI.AddAttachmentToCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddAttachmentToCase`: OutputAttachments
	fmt.Fprintf(os.Stdout, "Response from `CaseAPI.AddAttachmentToCase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** | The case ID preceded by &#x60;~&#x60;, or the case number. |

### Other Parameters

Other parameters are passed through a pointer to a apiAddAttachmentToCaseRequest struct via the builder pattern


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


## AddLinkWithAnExternalURL

> AddLinkWithAnExternalURL(ctx, caseId).InputExternalLink(inputExternalLink).Execute()

Link a case to an external URL



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
	inputExternalLink := *openapiclient.NewInputExternalLink("Type_example", "Url_example") // InputExternalLink |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CaseAPI.AddLinkWithAnExternalURL(context.Background(), caseId).InputExternalLink(inputExternalLink).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseAPI.AddLinkWithAnExternalURL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** | The case ID preceded by &#x60;~&#x60;, or the case number. |

### Other Parameters

Other parameters are passed through a pointer to a apiAddLinkWithAnExternalURLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputExternalLink** | [**InputExternalLink**](InputExternalLink.md) |  |

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


## AddLinkWithAnotherCase

> AddLinkWithAnotherCase(ctx, caseId).InputCaseLink(inputCaseLink).Execute()

Link a case to another case



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
	inputCaseLink := *openapiclient.NewInputCaseLink("Type_example", "CaseId_example") // InputCaseLink |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CaseAPI.AddLinkWithAnotherCase(context.Background(), caseId).InputCaseLink(inputCaseLink).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseAPI.AddLinkWithAnotherCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** | The case ID preceded by &#x60;~&#x60;, or the case number. |

### Other Parameters

Other parameters are passed through a pointer to a apiAddLinkWithAnotherCaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputCaseLink** | [**InputCaseLink**](InputCaseLink.md) |  |

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


## ApplyCaseTemplateOnExistingCases

> ApplyCaseTemplateOnExistingCases(ctx).InputApplyCaseTemplateWithIds(inputApplyCaseTemplateWithIds).Execute()

Apply a case template to existing cases



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
	inputApplyCaseTemplateWithIds := *openapiclient.NewInputApplyCaseTemplateWithIds([]string{"Ids_example"}, "CaseTemplate_example") // InputApplyCaseTemplateWithIds |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CaseAPI.ApplyCaseTemplateOnExistingCases(context.Background()).InputApplyCaseTemplateWithIds(inputApplyCaseTemplateWithIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseAPI.ApplyCaseTemplateOnExistingCases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplyCaseTemplateOnExistingCasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputApplyCaseTemplateWithIds** | [**InputApplyCaseTemplateWithIds**](InputApplyCaseTemplateWithIds.md) |  |

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


## BulkManageCaseAccess

> BulkManageCaseAccess(ctx).InputManageCaseAccessWithIds(inputManageCaseAccessWithIds).Execute()

Set the access mode of multiple cases



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
	inputManageCaseAccessWithIds := *openapiclient.NewInputManageCaseAccessWithIds([]string{"Ids_example"}, openapiclient.Access{AllExternalAccess: openapiclient.NewAllExternalAccess("Kind_example")}) // InputManageCaseAccessWithIds |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CaseAPI.BulkManageCaseAccess(context.Background()).InputManageCaseAccessWithIds(inputManageCaseAccessWithIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseAPI.BulkManageCaseAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkManageCaseAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputManageCaseAccessWithIds** | [**InputManageCaseAccessWithIds**](InputManageCaseAccessWithIds.md) |  |

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


## BulkUpdateCase

> BulkUpdateCase(ctx).InputUpdateCaseWithIds(inputUpdateCaseWithIds).Execute()

Update multiple cases



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
	inputUpdateCaseWithIds := *openapiclient.NewInputUpdateCaseWithIds([]string{"Ids_example"}) // InputUpdateCaseWithIds |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CaseAPI.BulkUpdateCase(context.Background()).InputUpdateCaseWithIds(inputUpdateCaseWithIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseAPI.BulkUpdateCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateCaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputUpdateCaseWithIds** | [**InputUpdateCaseWithIds**](InputUpdateCaseWithIds.md) |  |

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


## ChangeCaseOwningOrganisation

> ChangeCaseOwningOrganisation(ctx, caseId).InputChangeCaseOwnership(inputChangeCaseOwnership).Execute()

Change the owner organization of a case



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
	inputChangeCaseOwnership := *openapiclient.NewInputChangeCaseOwnership("Organisation_example") // InputChangeCaseOwnership |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CaseAPI.ChangeCaseOwningOrganisation(context.Background(), caseId).InputChangeCaseOwnership(inputChangeCaseOwnership).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseAPI.ChangeCaseOwningOrganisation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** | The case ID preceded by &#x60;~&#x60;, or the case number. |

### Other Parameters

Other parameters are passed through a pointer to a apiChangeCaseOwningOrganisationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputChangeCaseOwnership** | [**InputChangeCaseOwnership**](InputChangeCaseOwnership.md) |  |

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


## CreateCase

> OutputCase CreateCase(ctx).InputCreateCase(inputCreateCase).Execute()

Create a case



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
	inputCreateCase := *openapiclient.NewInputCreateCase("Title_example", "Description_example") // InputCreateCase |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaseAPI.CreateCase(context.Background()).InputCreateCase(inputCreateCase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseAPI.CreateCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCase`: OutputCase
	fmt.Fprintf(os.Stdout, "Response from `CaseAPI.CreateCase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputCreateCase** | [**InputCreateCase**](InputCreateCase.md) |  |

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


## DeleteACustomField

> DeleteACustomField(ctx, cfId).Execute()

Remove a custom field from a case



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
	cfId := "cfId_example" // string | The custom field ID preceded by `~`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CaseAPI.DeleteACustomField(context.Background(), cfId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseAPI.DeleteACustomField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cfId** | **string** | The custom field ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteACustomFieldRequest struct via the builder pattern


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


## DeleteCase

> DeleteCase(ctx, idOrName).Execute()

Delete a case



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
	idOrName := "idOrName_example" // string | The case ID preceded by `~`, or the case number.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CaseAPI.DeleteCase(context.Background(), idOrName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseAPI.DeleteCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrName** | **string** | The case ID preceded by &#x60;~&#x60;, or the case number. |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCaseRequest struct via the builder pattern


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


## DeleteCaseAttachment

> DeleteCaseAttachment(ctx, caseId, attachmentId).Execute()

Remove an attachment from a case



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
	attachmentId := "attachmentId_example" // string | The attachment ID preceded by `~`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CaseAPI.DeleteCaseAttachment(context.Background(), caseId, attachmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseAPI.DeleteCaseAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** | The case ID preceded by &#x60;~&#x60;, or the case number. |
**attachmentId** | **string** | The attachment ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCaseAttachmentRequest struct via the builder pattern


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


## DeleteLinkWithAnAnotherCase

> DeleteLinkWithAnAnotherCase(ctx, caseId).InputCaseLink(inputCaseLink).Execute()

Remove a link to another case



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
	inputCaseLink := *openapiclient.NewInputCaseLink("Type_example", "CaseId_example") // InputCaseLink |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CaseAPI.DeleteLinkWithAnAnotherCase(context.Background(), caseId).InputCaseLink(inputCaseLink).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseAPI.DeleteLinkWithAnAnotherCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** | The case ID preceded by &#x60;~&#x60;, or the case number. |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLinkWithAnAnotherCaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputCaseLink** | [**InputCaseLink**](InputCaseLink.md) |  |

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


## DeleteLinkWithAnExternalURL

> DeleteLinkWithAnExternalURL(ctx, caseId).InputExternalLink(inputExternalLink).Execute()

Remove a link to an external URL



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
	inputExternalLink := *openapiclient.NewInputExternalLink("Type_example", "Url_example") // InputExternalLink |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CaseAPI.DeleteLinkWithAnExternalURL(context.Background(), caseId).InputExternalLink(inputExternalLink).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseAPI.DeleteLinkWithAnExternalURL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** | The case ID preceded by &#x60;~&#x60;, or the case number. |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLinkWithAnExternalURLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputExternalLink** | [**InputExternalLink**](InputExternalLink.md) |  |

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


## DownloadAttachmentFromCase

> *os.File DownloadAttachmentFromCase(ctx, caseId, attachmentId).Execute()

Download an attachment from a case



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
	attachmentId := "attachmentId_example" // string | The attachment ID preceded by `~`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaseAPI.DownloadAttachmentFromCase(context.Background(), caseId, attachmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseAPI.DownloadAttachmentFromCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadAttachmentFromCase`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `CaseAPI.DownloadAttachmentFromCase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** | The case ID preceded by &#x60;~&#x60;, or the case number. |
**attachmentId** | **string** | The attachment ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadAttachmentFromCaseRequest struct via the builder pattern


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


## ExportCaseAsArchive

> *os.File ExportCaseAsArchive(ctx, caseId).Password(password).Execute()

Export a case as an archive



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
	password := "password_example" // string | Password to encrypt the archive. Required when restoring it via `POST /api/v1/case/import`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaseAPI.ExportCaseAsArchive(context.Background(), caseId).Password(password).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseAPI.ExportCaseAsArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportCaseAsArchive`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `CaseAPI.ExportCaseAsArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** | The case ID preceded by &#x60;~&#x60;, or the case number. |

### Other Parameters

Other parameters are passed through a pointer to a apiExportCaseAsArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **password** | **string** | Password to encrypt the archive. Required when restoring it via &#x60;POST /api/v1/case/import&#x60;. |

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


## GetAllLinkTypes

> []string GetAllLinkTypes(ctx).Execute()

List all case link types



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
	resp, r, err := apiClient.CaseAPI.GetAllLinkTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseAPI.GetAllLinkTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllLinkTypes`: []string
	fmt.Fprintf(os.Stdout, "Response from `CaseAPI.GetAllLinkTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllLinkTypesRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttachmentFromCase

> *os.File GetAttachmentFromCase(ctx, caseId, attachmentId).IfNoneMatch(ifNoneMatch).Execute()

Get an attachment from a case



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
	attachmentId := "attachmentId_example" // string | The attachment ID preceded by `~`.
	ifNoneMatch := "ifNoneMatch_example" // string | ETag from a previous response. If it matches, returns `304 Not Modified` instead of the file content. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaseAPI.GetAttachmentFromCase(context.Background(), caseId, attachmentId).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseAPI.GetAttachmentFromCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachmentFromCase`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `CaseAPI.GetAttachmentFromCase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** | The case ID preceded by &#x60;~&#x60;, or the case number. |
**attachmentId** | **string** | The attachment ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentFromCaseRequest struct via the builder pattern


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


## GetCase

> OutputCase GetCase(ctx, idOrName).Execute()

Get a case



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
	idOrName := "idOrName_example" // string | The case ID preceded by `~`, or the case number.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaseAPI.GetCase(context.Background(), idOrName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseAPI.GetCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCase`: OutputCase
	fmt.Fprintf(os.Stdout, "Response from `CaseAPI.GetCase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrName** | **string** | The case ID preceded by &#x60;~&#x60;, or the case number. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetCaseRequest struct via the builder pattern


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


## GetCaseTimeline

> OutputTimeline GetCaseTimeline(ctx, caseId).Execute()

Get the timeline of a case



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaseAPI.GetCaseTimeline(context.Background(), caseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseAPI.GetCaseTimeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCaseTimeline`: OutputTimeline
	fmt.Fprintf(os.Stdout, "Response from `CaseAPI.GetCaseTimeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** | The case ID preceded by &#x60;~&#x60;, or the case number. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetCaseTimelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputTimeline**](OutputTimeline.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimilarCases

> []OutputCaseWithLinks GetSimilarCases(ctx, caseId).Execute()

List similar cases



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaseAPI.GetSimilarCases(context.Background(), caseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseAPI.GetSimilarCases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimilarCases`: []OutputCaseWithLinks
	fmt.Fprintf(os.Stdout, "Response from `CaseAPI.GetSimilarCases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** | The case ID preceded by &#x60;~&#x60;, or the case number. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetSimilarCasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OutputCaseWithLinks**](OutputCaseWithLinks.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimilarObservablesBetweenACaseAndAnotherCaseOrAlert

> []OutputObservable GetSimilarObservablesBetweenACaseAndAnotherCaseOrAlert(ctx, caseId, alertOrCaseId).Execute()

List observables shared with another case or alert



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
	alertOrCaseId := "alertOrCaseId_example" // string | The ID of the alert or case, preceded by `~`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaseAPI.GetSimilarObservablesBetweenACaseAndAnotherCaseOrAlert(context.Background(), caseId, alertOrCaseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseAPI.GetSimilarObservablesBetweenACaseAndAnotherCaseOrAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimilarObservablesBetweenACaseAndAnotherCaseOrAlert`: []OutputObservable
	fmt.Fprintf(os.Stdout, "Response from `CaseAPI.GetSimilarObservablesBetweenACaseAndAnotherCaseOrAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** | The case ID preceded by &#x60;~&#x60;, or the case number. |
**alertOrCaseId** | **string** | The ID of the alert or case, preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetSimilarObservablesBetweenACaseAndAnotherCaseOrAlertRequest struct via the builder pattern


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


## ImportCaseFromFile

> OutputImportCase ImportCaseFromFile(ctx).Json(json).File(file).Execute()

Create a case from an archive



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
	json := *openapiclient.NewInputImportCase("Password_example") // InputImportCase |
	file := os.NewFile(1234, "some_file") // *os.File |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaseAPI.ImportCaseFromFile(context.Background()).Json(json).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseAPI.ImportCaseFromFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportCaseFromFile`: OutputImportCase
	fmt.Fprintf(os.Stdout, "Response from `CaseAPI.ImportCaseFromFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportCaseFromFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **json** | [**InputImportCase**](InputImportCase.md) |  |
 **file** | ***os.File** |  |

### Return type

[**OutputImportCase**](OutputImportCase.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageCaseAccess

> ManageCaseAccess(ctx, caseId).InputManageCaseAccess(inputManageCaseAccess).Execute()

Set the access mode of a case



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
	inputManageCaseAccess := *openapiclient.NewInputManageCaseAccess(openapiclient.Access{AllExternalAccess: openapiclient.NewAllExternalAccess("Kind_example")}) // InputManageCaseAccess |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CaseAPI.ManageCaseAccess(context.Background(), caseId).InputManageCaseAccess(inputManageCaseAccess).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseAPI.ManageCaseAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** | The case ID preceded by &#x60;~&#x60;, or the case number. |

### Other Parameters

Other parameters are passed through a pointer to a apiManageCaseAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputManageCaseAccess** | [**InputManageCaseAccess**](InputManageCaseAccess.md) |  |

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


## MergeCases

> OutputCase MergeCases(ctx, ids).Execute()

Merge cases into a new case



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
	ids := "ids_example" // string | List of case IDs preceded by `~`, or case numbers.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaseAPI.MergeCases(context.Background(), ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseAPI.MergeCases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MergeCases`: OutputCase
	fmt.Fprintf(os.Stdout, "Response from `CaseAPI.MergeCases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ids** | **string** | List of case IDs preceded by &#x60;~&#x60;, or case numbers. |

### Other Parameters

Other parameters are passed through a pointer to a apiMergeCasesRequest struct via the builder pattern


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


## MergeSimilarObservablesOfThisCase

> OutputMergeCases MergeSimilarObservablesOfThisCase(ctx, caseId).Execute()

Deduplicate observables in a case



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaseAPI.MergeSimilarObservablesOfThisCase(context.Background(), caseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseAPI.MergeSimilarObservablesOfThisCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MergeSimilarObservablesOfThisCase`: OutputMergeCases
	fmt.Fprintf(os.Stdout, "Response from `CaseAPI.MergeSimilarObservablesOfThisCase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** | The case ID preceded by &#x60;~&#x60;, or the case number. |

### Other Parameters

Other parameters are passed through a pointer to a apiMergeSimilarObservablesOfThisCaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputMergeCases**](OutputMergeCases.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkAlertFromCase

> UnlinkAlertFromCase(ctx, caseId, alertId).Execute()

Unlink an alert from a case



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
	alertId := "alertId_example" // string | The alert ID preceded by `~`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CaseAPI.UnlinkAlertFromCase(context.Background(), caseId, alertId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseAPI.UnlinkAlertFromCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** | The case ID preceded by &#x60;~&#x60;, or the case number. |
**alertId** | **string** | The alert ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiUnlinkAlertFromCaseRequest struct via the builder pattern


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


## UpdateAttachment

> UpdateAttachment(ctx, caseId, attachmentId).InputUpdateAttachment(inputUpdateAttachment).Execute()

Update a case attachment



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
	attachmentId := "attachmentId_example" // string | The attachment ID preceded by `~`.
	inputUpdateAttachment := *openapiclient.NewInputUpdateAttachment() // InputUpdateAttachment |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CaseAPI.UpdateAttachment(context.Background(), caseId, attachmentId).InputUpdateAttachment(inputUpdateAttachment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseAPI.UpdateAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** | The case ID preceded by &#x60;~&#x60;, or the case number. |
**attachmentId** | **string** | The attachment ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inputUpdateAttachment** | [**InputUpdateAttachment**](InputUpdateAttachment.md) |  |

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


## UpdateCase

> UpdateCase(ctx, idOrName).InputUpdateCase(inputUpdateCase).Execute()

Update a case



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
	idOrName := "idOrName_example" // string | The case ID preceded by `~`, or the case number.
	inputUpdateCase := *openapiclient.NewInputUpdateCase() // InputUpdateCase |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CaseAPI.UpdateCase(context.Background(), idOrName).InputUpdateCase(inputUpdateCase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseAPI.UpdateCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrName** | **string** | The case ID preceded by &#x60;~&#x60;, or the case number. |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputUpdateCase** | [**InputUpdateCase**](InputUpdateCase.md) |  |

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
