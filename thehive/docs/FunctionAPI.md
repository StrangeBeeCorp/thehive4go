# \FunctionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFunction**](FunctionAPI.md#CreateFunction) | **Post** /api/v1/function | Create a function
[**DeleteFunction**](FunctionAPI.md#DeleteFunction) | **Delete** /api/v1/function/{functionId} | Delete a function
[**GetDocumentationForTheContextObject**](FunctionAPI.md#GetDocumentationForTheContextObject) | **Get** /api/v1/function/_context/documentation | Get context object documentation
[**GetFunction**](FunctionAPI.md#GetFunction) | **Get** /api/v1/function/{functionId} | Get a function
[**InvokeFunction**](FunctionAPI.md#InvokeFunction) | **Post** /api/v1/function/{function} | Invoke a function
[**InvokeFunctionOnAnObject**](FunctionAPI.md#InvokeFunctionOnAnObject) | **Post** /api/v1/function/{function}/{objectType}/{objectIdOrName} | Invoke a function on a case or alert
[**TestFunction**](FunctionAPI.md#TestFunction) | **Post** /api/v1/function/_test | Test a function
[**UpdateFunction**](FunctionAPI.md#UpdateFunction) | **Patch** /api/v1/function/{functionId} | Update a function



## CreateFunction

> OutputFunction CreateFunction(ctx).InputFunction(inputFunction).Execute()

Create a function



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

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFunction

> DeleteFunction(ctx, functionId).Execute()

Delete a function



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
	functionId := "functionId_example" // string | The function ID preceded by `~`, or the function name.

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
**functionId** | **string** | The function ID preceded by &#x60;~&#x60;, or the function name. |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFunctionRequest struct via the builder pattern


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


## GetDocumentationForTheContextObject

> OutputContextDocumentation GetDocumentationForTheContextObject(ctx).Execute()

Get context object documentation



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

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunction

> OutputFunction GetFunction(ctx, functionId).Execute()

Get a function



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
	functionId := "functionId_example" // string | The function ID preceded by `~`, or the function name.

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
**functionId** | **string** | The function ID preceded by &#x60;~&#x60;, or the function name. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputFunction**](OutputFunction.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvokeFunction

> OutputInvokeFunctionOk InvokeFunction(ctx, function).DryRun(dryRun).Body(body).Execute()

Invoke a function



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
	function := "function_example" // string | The function ID preceded by `~`, or the function name.
	dryRun := true // bool | Set to `true` to simulate the execution without persisting any changes. (optional) (default to false)
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
**function** | **string** | The function ID preceded by &#x60;~&#x60;, or the function name. |

### Other Parameters

Other parameters are passed through a pointer to a apiInvokeFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dryRun** | **bool** | Set to &#x60;true&#x60; to simulate the execution without persisting any changes. | [default to false]
 **body** | **interface{}** |  |

### Return type

[**OutputInvokeFunctionOk**](OutputInvokeFunctionOk.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvokeFunctionOnAnObject

> InvokeFunctionOnAnObject200Response InvokeFunctionOnAnObject(ctx, function, objectType, objectIdOrName).DryRun(dryRun).Sync(sync).Execute()

Invoke a function on a case or alert



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
	function := "function_example" // string | The function ID preceded by `~`, or the function name.
	objectType := "objectType_example" // string | Type of the object to pass as input. Accepted values: `case`, `alert`.
	objectIdOrName := "objectIdOrName_example" // string | The object ID preceded by `~`, or a type-specific identifier: the case number for cases, or the alert reference (`type;source;sourceRef`) for alerts.
	dryRun := true // bool | Set to `true` to simulate the execution without persisting any changes. (optional) (default to false)
	sync := true // bool | Set to `true` to wait for the function to complete and return its result. Defaults to `false` to run the function in the background and return immediately. (optional) (default to false)

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
**function** | **string** | The function ID preceded by &#x60;~&#x60;, or the function name. |
**objectType** | **string** | Type of the object to pass as input. Accepted values: &#x60;case&#x60;, &#x60;alert&#x60;. |
**objectIdOrName** | **string** | The object ID preceded by &#x60;~&#x60;, or a type-specific identifier: the case number for cases, or the alert reference (&#x60;type;source;sourceRef&#x60;) for alerts. |

### Other Parameters

Other parameters are passed through a pointer to a apiInvokeFunctionOnAnObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dryRun** | **bool** | Set to &#x60;true&#x60; to simulate the execution without persisting any changes. | [default to false]
 **sync** | **bool** | Set to &#x60;true&#x60; to wait for the function to complete and return its result. Defaults to &#x60;false&#x60; to run the function in the background and return immediately. | [default to false]

### Return type

[**InvokeFunctionOnAnObject200Response**](InvokeFunctionOnAnObject200Response.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestFunction

> OutputInvokeFunctionOk TestFunction(ctx).InputTestFunction(inputTestFunction).DryRun(dryRun).Execute()

Test a function



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
	dryRun := true // bool | Set to `true` to simulate the execution without persisting any changes. (optional) (default to false)

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
 **dryRun** | **bool** | Set to &#x60;true&#x60; to simulate the execution without persisting any changes. | [default to false]

### Return type

[**OutputInvokeFunctionOk**](OutputInvokeFunctionOk.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFunction

> UpdateFunction(ctx, functionId).InputUpdateFunction(inputUpdateFunction).Execute()

Update a function



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
	functionId := "functionId_example" // string | The function ID preceded by `~`, or the function name.
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
**functionId** | **string** | The function ID preceded by &#x60;~&#x60;, or the function name. |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputUpdateFunction** | [**InputUpdateFunction**](InputUpdateFunction.md) |  |

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
