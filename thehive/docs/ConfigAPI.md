# \ConfigAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserConfigurationItem**](ConfigAPI.md#GetUserConfigurationItem) | **Get** /api/v1/config/user/{path} | Get a configuration item
[**ListUserConfigurationItems**](ConfigAPI.md#ListUserConfigurationItems) | **Get** /api/v1/config/user | List your configuration items
[**SetUserConfigurationItem**](ConfigAPI.md#SetUserConfigurationItem) | **Put** /api/v1/config/user/{path} | Set a configuration item



## GetUserConfigurationItem

> OutputConfig GetUserConfigurationItem(ctx, path).Execute()

Get a configuration item



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
	path := "path_example" // string | Key of the configuration item.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigAPI.GetUserConfigurationItem(context.Background(), path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigAPI.GetUserConfigurationItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserConfigurationItem`: OutputConfig
	fmt.Fprintf(os.Stdout, "Response from `ConfigAPI.GetUserConfigurationItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | Key of the configuration item. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserConfigurationItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputConfig**](OutputConfig.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserConfigurationItems

> interface{} ListUserConfigurationItems(ctx).Path(path).Execute()

List your configuration items



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
	path := "path_example" // string | Key of a single configuration item. Omit it to return your full configuration. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigAPI.ListUserConfigurationItems(context.Background()).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigAPI.ListUserConfigurationItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserConfigurationItems`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConfigAPI.ListUserConfigurationItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUserConfigurationItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Key of a single configuration item. Omit it to return your full configuration. |

### Return type

**interface{}**

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetUserConfigurationItem

> OutputConfig SetUserConfigurationItem(ctx, path).InputConfig(inputConfig).Execute()

Set a configuration item



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
	path := "path_example" // string | Key of the configuration item.
	inputConfig := *openapiclient.NewInputConfig(interface{}(123)) // InputConfig |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigAPI.SetUserConfigurationItem(context.Background(), path).InputConfig(inputConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigAPI.SetUserConfigurationItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetUserConfigurationItem`: OutputConfig
	fmt.Fprintf(os.Stdout, "Response from `ConfigAPI.SetUserConfigurationItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | Key of the configuration item. |

### Other Parameters

Other parameters are passed through a pointer to a apiSetUserConfigurationItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputConfig** | [**InputConfig**](InputConfig.md) |  |

### Return type

[**OutputConfig**](OutputConfig.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
