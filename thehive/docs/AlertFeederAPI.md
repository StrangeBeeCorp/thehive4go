# \AlertFeederAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAnAlertFeeder**](AlertFeederAPI.md#CreateAnAlertFeeder) | **Post** /api/v1/connector/alert-feeder | 
[**DeleteAnAlertFeeder**](AlertFeederAPI.md#DeleteAnAlertFeeder) | **Delete** /api/v1/connector/alert-feeder/{alertFeederName} | 
[**GetAlertFeeders**](AlertFeederAPI.md#GetAlertFeeders) | **Get** /api/v1/connector/alert-feeder | 
[**RunAnAlertFeeder**](AlertFeederAPI.md#RunAnAlertFeeder) | **Post** /api/v1/connector/alert-feeder/run/{alertFeederName} | 
[**TestAnAlertFeeder**](AlertFeederAPI.md#TestAnAlertFeeder) | **Post** /api/v1/connector/alert-feeder/test | 
[**UpdateAnAlertFeeder**](AlertFeederAPI.md#UpdateAnAlertFeeder) | **Put** /api/v1/connector/alert-feeder/{alertFeederName} | 



## CreateAnAlertFeeder

> OutputAlertFeeder CreateAnAlertFeeder(ctx).InputAlertFeeder(inputAlertFeeder).Execute()





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
	inputAlertFeeder := *openapiclient.NewInputAlertFeeder("Name_example", "Description_example", openapiclient.Method("GET"), "Url_example", *openapiclient.NewInterval(int32(123), openapiclient.IntervalUnit("Days")), "FunctionName_example") // InputAlertFeeder | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertFeederAPI.CreateAnAlertFeeder(context.Background()).InputAlertFeeder(inputAlertFeeder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertFeederAPI.CreateAnAlertFeeder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAnAlertFeeder`: OutputAlertFeeder
	fmt.Fprintf(os.Stdout, "Response from `AlertFeederAPI.CreateAnAlertFeeder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAnAlertFeederRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputAlertFeeder** | [**InputAlertFeeder**](InputAlertFeeder.md) |  | 

### Return type

[**OutputAlertFeeder**](OutputAlertFeeder.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAnAlertFeeder

> DeleteAnAlertFeeder(ctx, alertFeederName).Execute()





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
	alertFeederName := "alertFeederName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertFeederAPI.DeleteAnAlertFeeder(context.Background(), alertFeederName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertFeederAPI.DeleteAnAlertFeeder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertFeederName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnAlertFeederRequest struct via the builder pattern


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


## GetAlertFeeders

> []OutputAlertFeeder GetAlertFeeders(ctx).Execute()





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
	resp, r, err := apiClient.AlertFeederAPI.GetAlertFeeders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertFeederAPI.GetAlertFeeders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertFeeders`: []OutputAlertFeeder
	fmt.Fprintf(os.Stdout, "Response from `AlertFeederAPI.GetAlertFeeders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertFeedersRequest struct via the builder pattern


### Return type

[**[]OutputAlertFeeder**](OutputAlertFeeder.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunAnAlertFeeder

> OutputInvokeFunction RunAnAlertFeeder(ctx, alertFeederName).DryRun(dryRun).Execute()





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
	alertFeederName := "alertFeederName_example" // string | 
	dryRun := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertFeederAPI.RunAnAlertFeeder(context.Background(), alertFeederName).DryRun(dryRun).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertFeederAPI.RunAnAlertFeeder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunAnAlertFeeder`: OutputInvokeFunction
	fmt.Fprintf(os.Stdout, "Response from `AlertFeederAPI.RunAnAlertFeeder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertFeederName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunAnAlertFeederRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dryRun** | **bool** |  | [default to false]

### Return type

[**OutputInvokeFunction**](OutputInvokeFunction.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestAnAlertFeeder

> interface{} TestAnAlertFeeder(ctx).InputTestAlertFeeder(inputTestAlertFeeder).Execute()





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
	inputTestAlertFeeder := *openapiclient.NewInputTestAlertFeeder("Name_example", "Description_example", openapiclient.Method("GET"), "Url_example", *openapiclient.NewInterval(int32(123), openapiclient.IntervalUnit("Days"))) // InputTestAlertFeeder | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertFeederAPI.TestAnAlertFeeder(context.Background()).InputTestAlertFeeder(inputTestAlertFeeder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertFeederAPI.TestAnAlertFeeder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestAnAlertFeeder`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AlertFeederAPI.TestAnAlertFeeder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestAnAlertFeederRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputTestAlertFeeder** | [**InputTestAlertFeeder**](InputTestAlertFeeder.md) |  | 

### Return type

**interface{}**

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAnAlertFeeder

> OutputAlertFeeder UpdateAnAlertFeeder(ctx, alertFeederName).InputUpdateAlertFeeder(inputUpdateAlertFeeder).Execute()





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
	alertFeederName := "alertFeederName_example" // string | 
	inputUpdateAlertFeeder := *openapiclient.NewInputUpdateAlertFeeder("Description_example", openapiclient.Method("GET"), "Url_example", *openapiclient.NewInterval(int32(123), openapiclient.IntervalUnit("Days"))) // InputUpdateAlertFeeder | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertFeederAPI.UpdateAnAlertFeeder(context.Background(), alertFeederName).InputUpdateAlertFeeder(inputUpdateAlertFeeder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertFeederAPI.UpdateAnAlertFeeder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAnAlertFeeder`: OutputAlertFeeder
	fmt.Fprintf(os.Stdout, "Response from `AlertFeederAPI.UpdateAnAlertFeeder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertFeederName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAnAlertFeederRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputUpdateAlertFeeder** | [**InputUpdateAlertFeeder**](InputUpdateAlertFeeder.md) |  | 

### Return type

[**OutputAlertFeeder**](OutputAlertFeeder.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

