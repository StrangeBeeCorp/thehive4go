# \BrandingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBrandingAttachment**](BrandingAPI.md#DeleteBrandingAttachment) | **Delete** /api/v1/branding/assets/{kind} |
[**GetBranding**](BrandingAPI.md#GetBranding) | **Get** /api/v1/branding |
[**GetBrandingAttachment**](BrandingAPI.md#GetBrandingAttachment) | **Get** /api/v1/branding/assets/{kind} |
[**SetBranding**](BrandingAPI.md#SetBranding) | **Post** /api/v1/branding |



## DeleteBrandingAttachment

> DeleteBrandingAttachment(ctx, kind).Execute()



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
	kind := "kind_example" // string |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BrandingAPI.DeleteBrandingAttachment(context.Background(), kind).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BrandingAPI.DeleteBrandingAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kind** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBrandingAttachmentRequest struct via the builder pattern


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


## GetBranding

> OutputBranding GetBranding(ctx).Execute()



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
	resp, r, err := apiClient.BrandingAPI.GetBranding(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BrandingAPI.GetBranding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBranding`: OutputBranding
	fmt.Fprintf(os.Stdout, "Response from `BrandingAPI.GetBranding`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandingRequest struct via the builder pattern


### Return type

[**OutputBranding**](OutputBranding.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrandingAttachment

> *os.File GetBrandingAttachment(ctx, kind).IfNoneMatch(ifNoneMatch).Execute()



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
	kind := "kind_example" // string |
	ifNoneMatch := "ifNoneMatch_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BrandingAPI.GetBrandingAttachment(context.Background(), kind).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BrandingAPI.GetBrandingAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBrandingAttachment`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `BrandingAPI.GetBrandingAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kind** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandingAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifNoneMatch** | **string** |  |

### Return type

[***os.File**](*os.File.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetBranding

> OutputBranding SetBranding(ctx).Title(title).LoginLogo(loginLogo).MenuLogo(menuLogo).Favicon(favicon).Execute()



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
	title := "title_example" // string |  (optional)
	loginLogo := os.NewFile(1234, "some_file") // *os.File |  (optional)
	menuLogo := os.NewFile(1234, "some_file") // *os.File |  (optional)
	favicon := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BrandingAPI.SetBranding(context.Background()).Title(title).LoginLogo(loginLogo).MenuLogo(menuLogo).Favicon(favicon).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BrandingAPI.SetBranding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetBranding`: OutputBranding
	fmt.Fprintf(os.Stdout, "Response from `BrandingAPI.SetBranding`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetBrandingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **title** | **string** |  |
 **loginLogo** | ***os.File** |  |
 **menuLogo** | ***os.File** |  |
 **favicon** | ***os.File** |  |

### Return type

[**OutputBranding**](OutputBranding.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
