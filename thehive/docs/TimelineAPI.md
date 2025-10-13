# \TimelineAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomEvent**](TimelineAPI.md#CreateCustomEvent) | **Post** /api/v1/case/{caseId}/customEvent | 
[**DeleteACustomEvent**](TimelineAPI.md#DeleteACustomEvent) | **Delete** /api/v1/customEvent/{eventId} | 
[**UpdateACustomEvent**](TimelineAPI.md#UpdateACustomEvent) | **Patch** /api/v1/customEvent/{eventId} | 



## CreateCustomEvent

> OutputCustomEvent CreateCustomEvent(ctx, caseId).InputCustomEvent(inputCustomEvent).Execute()



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
	inputCustomEvent := *openapiclient.NewInputCustomEvent(int32(1640000000000), "Title_example") // InputCustomEvent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimelineAPI.CreateCustomEvent(context.Background(), caseId).InputCustomEvent(inputCustomEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimelineAPI.CreateCustomEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomEvent`: OutputCustomEvent
	fmt.Fprintf(os.Stdout, "Response from `TimelineAPI.CreateCustomEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputCustomEvent** | [**InputCustomEvent**](InputCustomEvent.md) |  | 

### Return type

[**OutputCustomEvent**](OutputCustomEvent.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteACustomEvent

> DeleteACustomEvent(ctx, eventId).Execute()



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
	eventId := "~354" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TimelineAPI.DeleteACustomEvent(context.Background(), eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimelineAPI.DeleteACustomEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteACustomEventRequest struct via the builder pattern


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


## UpdateACustomEvent

> UpdateACustomEvent(ctx, eventId).InputUpdateCustomEvent(inputUpdateCustomEvent).Execute()



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
	eventId := "~354" // string | 
	inputUpdateCustomEvent := *openapiclient.NewInputUpdateCustomEvent() // InputUpdateCustomEvent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TimelineAPI.UpdateACustomEvent(context.Background(), eventId).InputUpdateCustomEvent(inputUpdateCustomEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimelineAPI.UpdateACustomEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateACustomEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputUpdateCustomEvent** | [**InputUpdateCustomEvent**](InputUpdateCustomEvent.md) |  | 

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

