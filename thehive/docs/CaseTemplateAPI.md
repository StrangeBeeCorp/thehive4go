# \CaseTemplateAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCaseTemplate**](CaseTemplateAPI.md#CreateCaseTemplate) | **Post** /api/v1/caseTemplate | 
[**DeleteCaseTemplate**](CaseTemplateAPI.md#DeleteCaseTemplate) | **Delete** /api/v1/caseTemplate/{caseTemplateNameOrId} | 
[**GetCaseTemplate**](CaseTemplateAPI.md#GetCaseTemplate) | **Get** /api/v1/caseTemplate/{caseTemplateNameOrId} | 
[**LinkPageTemplatesToACaseTemplate**](CaseTemplateAPI.md#LinkPageTemplatesToACaseTemplate) | **Put** /api/v1/caseTemplate/{caseTemplateNameOrId}/pageTemplate/link | 
[**UpdateCaseTemplate**](CaseTemplateAPI.md#UpdateCaseTemplate) | **Patch** /api/v1/caseTemplate/{caseTemplateNameOrId} | 



## CreateCaseTemplate

> OutputCaseTemplate CreateCaseTemplate(ctx).InputCreateCaseTemplate(inputCreateCaseTemplate).Execute()



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
	inputCreateCaseTemplate := *openapiclient.NewInputCreateCaseTemplate("Name_example") // InputCreateCaseTemplate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaseTemplateAPI.CreateCaseTemplate(context.Background()).InputCreateCaseTemplate(inputCreateCaseTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseTemplateAPI.CreateCaseTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCaseTemplate`: OutputCaseTemplate
	fmt.Fprintf(os.Stdout, "Response from `CaseTemplateAPI.CreateCaseTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCaseTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputCreateCaseTemplate** | [**InputCreateCaseTemplate**](InputCreateCaseTemplate.md) |  | 

### Return type

[**OutputCaseTemplate**](OutputCaseTemplate.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCaseTemplate

> DeleteCaseTemplate(ctx, caseTemplateNameOrId).Execute()



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
	caseTemplateNameOrId := "caseTemplateNameOrId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CaseTemplateAPI.DeleteCaseTemplate(context.Background(), caseTemplateNameOrId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseTemplateAPI.DeleteCaseTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseTemplateNameOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCaseTemplateRequest struct via the builder pattern


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


## GetCaseTemplate

> OutputCaseTemplate GetCaseTemplate(ctx, caseTemplateNameOrId).Execute()



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
	caseTemplateNameOrId := "caseTemplateNameOrId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaseTemplateAPI.GetCaseTemplate(context.Background(), caseTemplateNameOrId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseTemplateAPI.GetCaseTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCaseTemplate`: OutputCaseTemplate
	fmt.Fprintf(os.Stdout, "Response from `CaseTemplateAPI.GetCaseTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseTemplateNameOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCaseTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputCaseTemplate**](OutputCaseTemplate.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkPageTemplatesToACaseTemplate

> LinkPageTemplatesToACaseTemplate(ctx, caseTemplateNameOrId).InputLinkPageTemplatesToCaseTemplate(inputLinkPageTemplatesToCaseTemplate).Execute()



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
	caseTemplateNameOrId := "caseTemplateNameOrId_example" // string | 
	inputLinkPageTemplatesToCaseTemplate := *openapiclient.NewInputLinkPageTemplatesToCaseTemplate() // InputLinkPageTemplatesToCaseTemplate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CaseTemplateAPI.LinkPageTemplatesToACaseTemplate(context.Background(), caseTemplateNameOrId).InputLinkPageTemplatesToCaseTemplate(inputLinkPageTemplatesToCaseTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseTemplateAPI.LinkPageTemplatesToACaseTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseTemplateNameOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLinkPageTemplatesToACaseTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputLinkPageTemplatesToCaseTemplate** | [**InputLinkPageTemplatesToCaseTemplate**](InputLinkPageTemplatesToCaseTemplate.md) |  | 

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


## UpdateCaseTemplate

> UpdateCaseTemplate(ctx, caseTemplateNameOrId).InputUpdateCaseTemplate(inputUpdateCaseTemplate).Execute()



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
	caseTemplateNameOrId := "caseTemplateNameOrId_example" // string | 
	inputUpdateCaseTemplate := *openapiclient.NewInputUpdateCaseTemplate() // InputUpdateCaseTemplate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CaseTemplateAPI.UpdateCaseTemplate(context.Background(), caseTemplateNameOrId).InputUpdateCaseTemplate(inputUpdateCaseTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseTemplateAPI.UpdateCaseTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseTemplateNameOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCaseTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputUpdateCaseTemplate** | [**InputUpdateCaseTemplate**](InputUpdateCaseTemplate.md) |  | 

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

