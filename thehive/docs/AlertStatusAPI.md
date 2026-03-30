# \AlertStatusAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlertStatus**](AlertStatusAPI.md#CreateAlertStatus) | **Post** /api/v1/alertStatus | 
[**DeleteAlertStatus**](AlertStatusAPI.md#DeleteAlertStatus) | **Delete** /api/v1/alertStatus/{id} | 
[**UpdateAlertStatus**](AlertStatusAPI.md#UpdateAlertStatus) | **Patch** /api/v1/alertStatus/{id} | 



## CreateAlertStatus

> OutputAlertStatus CreateAlertStatus(ctx).InputCreateAlertStatus(inputCreateAlertStatus).Execute()



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
	inputCreateAlertStatus := *openapiclient.NewInputCreateAlertStatus("Value_example", openapiclient.InputAlertStage("New")) // InputCreateAlertStatus | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertStatusAPI.CreateAlertStatus(context.Background()).InputCreateAlertStatus(inputCreateAlertStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertStatusAPI.CreateAlertStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAlertStatus`: OutputAlertStatus
	fmt.Fprintf(os.Stdout, "Response from `AlertStatusAPI.CreateAlertStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputCreateAlertStatus** | [**InputCreateAlertStatus**](InputCreateAlertStatus.md) |  | 

### Return type

[**OutputAlertStatus**](OutputAlertStatus.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlertStatus

> DeleteAlertStatus(ctx, id).Execute()



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
	r, err := apiClient.AlertStatusAPI.DeleteAlertStatus(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertStatusAPI.DeleteAlertStatus``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAlertStatusRequest struct via the builder pattern


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


## UpdateAlertStatus

> UpdateAlertStatus(ctx, id).InputUpdateAlertStatus(inputUpdateAlertStatus).Execute()



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
	inputUpdateAlertStatus := *openapiclient.NewInputUpdateAlertStatus() // InputUpdateAlertStatus | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertStatusAPI.UpdateAlertStatus(context.Background(), id).InputUpdateAlertStatus(inputUpdateAlertStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertStatusAPI.UpdateAlertStatus``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateAlertStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputUpdateAlertStatus** | [**InputUpdateAlertStatus**](InputUpdateAlertStatus.md) |  | 

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

