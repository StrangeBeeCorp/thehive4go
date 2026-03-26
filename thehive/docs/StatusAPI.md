# \StatusAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPlatformPublicStatus**](StatusAPI.md#GetPlatformPublicStatus) | **Get** /api/v1/status/public | 
[**GetPlatformStatus**](StatusAPI.md#GetPlatformStatus) | **Get** /api/v1/status | 



## GetPlatformPublicStatus

> OutputPublicStatus GetPlatformPublicStatus(ctx).Execute()



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
	resp, r, err := apiClient.StatusAPI.GetPlatformPublicStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusAPI.GetPlatformPublicStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlatformPublicStatus`: OutputPublicStatus
	fmt.Fprintf(os.Stdout, "Response from `StatusAPI.GetPlatformPublicStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformPublicStatusRequest struct via the builder pattern


### Return type

[**OutputPublicStatus**](OutputPublicStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformStatus

> OutputStatus GetPlatformStatus(ctx).Verbose(verbose).Execute()



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
	verbose := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatusAPI.GetPlatformStatus(context.Background()).Verbose(verbose).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusAPI.GetPlatformStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlatformStatus`: OutputStatus
	fmt.Fprintf(os.Stdout, "Response from `StatusAPI.GetPlatformStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verbose** | **bool** |  | [default to false]

### Return type

[**OutputStatus**](OutputStatus.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

