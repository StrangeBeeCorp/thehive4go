# \FunctionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFunction**](FunctionAPI.md#CreateFunction) | **Post** /api/v1/function | 
[**DeleteFunction**](FunctionAPI.md#DeleteFunction) | **Delete** /api/v1/function/{functionId} | 
[**GetDocumentationForTheContextObject**](FunctionAPI.md#GetDocumentationForTheContextObject) | **Get** /api/v1/function/_context/documentation | 
[**GetFunction**](FunctionAPI.md#GetFunction) | **Get** /api/v1/function/{functionId} | 
[**InvokeFunction**](FunctionAPI.md#InvokeFunction) | **Post** /api/v1/function/{function} | 
[**InvokeFunctionOnAnObject**](FunctionAPI.md#InvokeFunctionOnAnObject) | **Post** /api/v1/function/{function}/{objectType}/{objectIdOrName} | 
[**TestFunction**](FunctionAPI.md#TestFunction) | **Post** /api/v1/function/_test | 
[**UpdateFunction**](FunctionAPI.md#UpdateFunction) | **Patch** /api/v1/function/{functionId} | 



## CreateFunction

> OutputFunction CreateFunction(ctx).InputFunction(inputFunction).Execute()





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
	inputFunction := *openapiclient.NewInputFunction("Name_example", "Definition_example") // InputFunction | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionAPI.CreateFunction(context.Background()).InputFunction(inputFunction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionAPI.CreateFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFunction`: OutputFunction
	fmt.Fprintf(os.Stdout, "Response from `FunctionAPI.CreateFunction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputFunction** | [**InputFunction**](InputFunction.md) |  | 

### Return type

[**OutputFunction**](OutputFunction.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFunction

> DeleteFunction(ctx, functionId).Execute()



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
	functionId := "functionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FunctionAPI.DeleteFunction(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionAPI.DeleteFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFunctionRequest struct via the builder pattern


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


## GetDocumentationForTheContextObject

> OutputContextDocumentation GetDocumentationForTheContextObject(ctx).Execute()



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
	resp, r, err := apiClient.FunctionAPI.GetDocumentationForTheContextObject(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionAPI.GetDocumentationForTheContextObject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDocumentationForTheContextObject`: OutputContextDocumentation
	fmt.Fprintf(os.Stdout, "Response from `FunctionAPI.GetDocumentationForTheContextObject`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentationForTheContextObjectRequest struct via the builder pattern


### Return type

[**OutputContextDocumentation**](OutputContextDocumentation.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunction

> OutputFunction GetFunction(ctx, functionId).Execute()



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
	functionId := "functionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionAPI.GetFunction(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionAPI.GetFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunction`: OutputFunction
	fmt.Fprintf(os.Stdout, "Response from `FunctionAPI.GetFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputFunction**](OutputFunction.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvokeFunction

> OutputInvokeFunctionOk InvokeFunction(ctx, function).DryRun(dryRun).Body(body).Execute()



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
	function := "function_example" // string | 
	dryRun := true // bool |  (optional) (default to false)
	body := interface{}(987) // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionAPI.InvokeFunction(context.Background(), function).DryRun(dryRun).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionAPI.InvokeFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvokeFunction`: OutputInvokeFunctionOk
	fmt.Fprintf(os.Stdout, "Response from `FunctionAPI.InvokeFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**function** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvokeFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dryRun** | **bool** |  | [default to false]
 **body** | **interface{}** |  | 

### Return type

[**OutputInvokeFunctionOk**](OutputInvokeFunctionOk.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvokeFunctionOnAnObject

> InvokeFunctionOnAnObject200Response InvokeFunctionOnAnObject(ctx, function, objectType, objectIdOrName).DryRun(dryRun).Sync(sync).Execute()





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
	function := "function_example" // string | 
	objectType := "objectType_example" // string | 
	objectIdOrName := "objectIdOrName_example" // string | 
	dryRun := true // bool |  (optional) (default to false)
	sync := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionAPI.InvokeFunctionOnAnObject(context.Background(), function, objectType, objectIdOrName).DryRun(dryRun).Sync(sync).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionAPI.InvokeFunctionOnAnObject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvokeFunctionOnAnObject`: InvokeFunctionOnAnObject200Response
	fmt.Fprintf(os.Stdout, "Response from `FunctionAPI.InvokeFunctionOnAnObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**function** | **string** |  | 
**objectType** | **string** |  | 
**objectIdOrName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvokeFunctionOnAnObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dryRun** | **bool** |  | [default to false]
 **sync** | **bool** |  | [default to false]

### Return type

[**InvokeFunctionOnAnObject200Response**](InvokeFunctionOnAnObject200Response.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestFunction

> OutputInvokeFunctionOk TestFunction(ctx).InputTestFunction(inputTestFunction).DryRun(dryRun).Execute()





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
	inputTestFunction := *openapiclient.NewInputTestFunction("Definition_example") // InputTestFunction | 
	dryRun := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionAPI.TestFunction(context.Background()).InputTestFunction(inputTestFunction).DryRun(dryRun).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionAPI.TestFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestFunction`: OutputInvokeFunctionOk
	fmt.Fprintf(os.Stdout, "Response from `FunctionAPI.TestFunction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputTestFunction** | [**InputTestFunction**](InputTestFunction.md) |  | 
 **dryRun** | **bool** |  | [default to false]

### Return type

[**OutputInvokeFunctionOk**](OutputInvokeFunctionOk.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFunction

> UpdateFunction(ctx, functionId).InputUpdateFunction(inputUpdateFunction).Execute()



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
	functionId := "functionId_example" // string | 
	inputUpdateFunction := *openapiclient.NewInputUpdateFunction() // InputUpdateFunction | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FunctionAPI.UpdateFunction(context.Background(), functionId).InputUpdateFunction(inputUpdateFunction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionAPI.UpdateFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputUpdateFunction** | [**InputUpdateFunction**](InputUpdateFunction.md) |  | 

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

