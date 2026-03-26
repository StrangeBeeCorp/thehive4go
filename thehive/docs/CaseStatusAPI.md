# \CaseStatusAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCaseStatus**](CaseStatusAPI.md#CreateCaseStatus) | **Post** /api/v1/caseStatus | 
[**DeleteCaseStatus**](CaseStatusAPI.md#DeleteCaseStatus) | **Delete** /api/v1/caseStatus/{id} | 
[**UpdateCaseStatus**](CaseStatusAPI.md#UpdateCaseStatus) | **Patch** /api/v1/caseStatus/{id} | 



## CreateCaseStatus

> OutputCaseStatus CreateCaseStatus(ctx).InputCreateCaseStatus(inputCreateCaseStatus).Execute()



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
	inputCreateCaseStatus := *openapiclient.NewInputCreateCaseStatus("Value_example", openapiclient.InputCaseStage("New")) // InputCreateCaseStatus | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaseStatusAPI.CreateCaseStatus(context.Background()).InputCreateCaseStatus(inputCreateCaseStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseStatusAPI.CreateCaseStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCaseStatus`: OutputCaseStatus
	fmt.Fprintf(os.Stdout, "Response from `CaseStatusAPI.CreateCaseStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCaseStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputCreateCaseStatus** | [**InputCreateCaseStatus**](InputCreateCaseStatus.md) |  | 

### Return type

[**OutputCaseStatus**](OutputCaseStatus.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCaseStatus

> DeleteCaseStatus(ctx, id).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CaseStatusAPI.DeleteCaseStatus(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseStatusAPI.DeleteCaseStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCaseStatusRequest struct via the builder pattern


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


## UpdateCaseStatus

> UpdateCaseStatus(ctx, id).InputUpdateCaseStatus(inputUpdateCaseStatus).Execute()



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
	id := "id_example" // string | 
	inputUpdateCaseStatus := *openapiclient.NewInputUpdateCaseStatus() // InputUpdateCaseStatus | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CaseStatusAPI.UpdateCaseStatus(context.Background(), id).InputUpdateCaseStatus(inputUpdateCaseStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseStatusAPI.UpdateCaseStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCaseStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputUpdateCaseStatus** | [**InputUpdateCaseStatus**](InputUpdateCaseStatus.md) |  | 

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

