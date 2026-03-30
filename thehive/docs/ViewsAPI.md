# \ViewsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateViews**](ViewsAPI.md#CreateViews) | **Post** /api/v1/views | 
[**DeleteViews**](ViewsAPI.md#DeleteViews) | **Delete** /api/v1/views/{viewsId} | 
[**GetViews**](ViewsAPI.md#GetViews) | **Get** /api/v1/views/{viewsId} | 
[**UpdateViews**](ViewsAPI.md#UpdateViews) | **Patch** /api/v1/views/{viewsId} | 



## CreateViews

> OutputListView CreateViews(ctx).InputCreateListView(inputCreateListView).Execute()





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
	inputCreateListView := *openapiclient.NewInputCreateListView("Name_example", openapiclient.InputListViewEntity("Alert"), map[string]interface{}(123), *openapiclient.NewListOptions(int32(123), false, false, false), false) // InputCreateListView | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewsAPI.CreateViews(context.Background()).InputCreateListView(inputCreateListView).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewsAPI.CreateViews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateViews`: OutputListView
	fmt.Fprintf(os.Stdout, "Response from `ViewsAPI.CreateViews`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputCreateListView** | [**InputCreateListView**](InputCreateListView.md) |  | 

### Return type

[**OutputListView**](OutputListView.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteViews

> DeleteViews(ctx, viewsId).Execute()



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
	viewsId := "viewsId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ViewsAPI.DeleteViews(context.Background(), viewsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewsAPI.DeleteViews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**viewsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteViewsRequest struct via the builder pattern


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


## GetViews

> OutputListView GetViews(ctx, viewsId).Execute()



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
	viewsId := "viewsId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewsAPI.GetViews(context.Background(), viewsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewsAPI.GetViews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetViews`: OutputListView
	fmt.Fprintf(os.Stdout, "Response from `ViewsAPI.GetViews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**viewsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputListView**](OutputListView.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateViews

> UpdateViews(ctx, viewsId).InputUpdateListView(inputUpdateListView).Execute()





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
	viewsId := "viewsId_example" // string | 
	inputUpdateListView := *openapiclient.NewInputUpdateListView() // InputUpdateListView | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ViewsAPI.UpdateViews(context.Background(), viewsId).InputUpdateListView(inputUpdateListView).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewsAPI.UpdateViews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**viewsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputUpdateListView** | [**InputUpdateListView**](InputUpdateListView.md) |  | 

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

