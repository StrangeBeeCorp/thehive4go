# \ObservableTypeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateObservableType**](ObservableTypeAPI.md#CreateObservableType) | **Post** /api/v1/observable/type | Create an observable type
[**DeleteObservableType**](ObservableTypeAPI.md#DeleteObservableType) | **Delete** /api/v1/observable/type/{typeId} | Delete an observable type
[**GetObservableType**](ObservableTypeAPI.md#GetObservableType) | **Get** /api/v1/observable/type/{typeId} | Get an observable type
[**UpdateObservableType**](ObservableTypeAPI.md#UpdateObservableType) | **Patch** /api/v1/observable/type/{typeId} | Update the case sensitivity of an observable type



## CreateObservableType

> OutputObservableType CreateObservableType(ctx).InputObservableType(inputObservableType).Execute()

Create an observable type



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
	inputObservableType := *openapiclient.NewInputObservableType("Name_example") // InputObservableType |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObservableTypeAPI.CreateObservableType(context.Background()).InputObservableType(inputObservableType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservableTypeAPI.CreateObservableType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateObservableType`: OutputObservableType
	fmt.Fprintf(os.Stdout, "Response from `ObservableTypeAPI.CreateObservableType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateObservableTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputObservableType** | [**InputObservableType**](InputObservableType.md) |  |

### Return type

[**OutputObservableType**](OutputObservableType.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteObservableType

> DeleteObservableType(ctx, typeId).Execute()

Delete an observable type



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
	typeId := "typeId_example" // string | The observable type ID preceded by `~`, or the observable type name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ObservableTypeAPI.DeleteObservableType(context.Background(), typeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservableTypeAPI.DeleteObservableType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeId** | **string** | The observable type ID preceded by &#x60;~&#x60;, or the observable type name. |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteObservableTypeRequest struct via the builder pattern


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


## GetObservableType

> OutputObservableType GetObservableType(ctx, typeId).Execute()

Get an observable type



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
	typeId := "typeId_example" // string | The observable type ID preceded by `~`, or the observable type name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObservableTypeAPI.GetObservableType(context.Background(), typeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservableTypeAPI.GetObservableType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObservableType`: OutputObservableType
	fmt.Fprintf(os.Stdout, "Response from `ObservableTypeAPI.GetObservableType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeId** | **string** | The observable type ID preceded by &#x60;~&#x60;, or the observable type name. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetObservableTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputObservableType**](OutputObservableType.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateObservableType

> UpdateObservableType(ctx, typeId).InputUpdateObservableType(inputUpdateObservableType).Execute()

Update the case sensitivity of an observable type



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
	typeId := "typeId_example" // string | The observable type ID preceded by `~`, or the observable type name.
	inputUpdateObservableType := *openapiclient.NewInputUpdateObservableType() // InputUpdateObservableType |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ObservableTypeAPI.UpdateObservableType(context.Background(), typeId).InputUpdateObservableType(inputUpdateObservableType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservableTypeAPI.UpdateObservableType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeId** | **string** | The observable type ID preceded by &#x60;~&#x60;, or the observable type name. |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateObservableTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputUpdateObservableType** | [**InputUpdateObservableType**](InputUpdateObservableType.md) |  |

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
