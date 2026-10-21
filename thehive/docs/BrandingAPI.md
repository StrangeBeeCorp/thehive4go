# \BrandingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBrandingAttachment**](BrandingAPI.md#DeleteBrandingAttachment) | **Delete** /api/v1/branding/assets/{kind} | Delete a branding asset
[**GetBranding**](BrandingAPI.md#GetBranding) | **Get** /api/v1/branding | Get branding
[**GetBrandingAttachment**](BrandingAPI.md#GetBrandingAttachment) | **Get** /api/v1/branding/assets/{kind} | Get a branding asset
[**SetBranding**](BrandingAPI.md#SetBranding) | **Post** /api/v1/branding | Set branding



## DeleteBrandingAttachment

> DeleteBrandingAttachment(ctx, kind).Execute()

Delete a branding asset



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
	kind := "kind_example" // string | Enum: `loginLogo` `menuLogo` `favicon`<br>Branding asset to delete.

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
**kind** | **string** | Enum: &#x60;loginLogo&#x60; &#x60;menuLogo&#x60; &#x60;favicon&#x60;&lt;br&gt;Branding asset to delete. |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBrandingAttachmentRequest struct via the builder pattern


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


## GetBranding

> OutputBranding GetBranding(ctx).Execute()

Get branding



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

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrandingAttachment

> *os.File GetBrandingAttachment(ctx, kind).IfNoneMatch(ifNoneMatch).Execute()

Get a branding asset



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
	kind := "kind_example" // string | Enum: `loginLogo` `menuLogo` `favicon`<br>Branding asset to retrieve.
	ifNoneMatch := "ifNoneMatch_example" // string | ETag from a previous response. If it matches, returns `304 Not Modified` instead of the file content. (optional)

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
**kind** | **string** | Enum: &#x60;loginLogo&#x60; &#x60;menuLogo&#x60; &#x60;favicon&#x60;&lt;br&gt;Branding asset to retrieve. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandingAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifNoneMatch** | **string** | ETag from a previous response. If it matches, returns &#x60;304 Not Modified&#x60; instead of the file content. |

### Return type

[***os.File**](*os.File.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetBranding

> OutputBranding SetBranding(ctx).Title(title).LoginLogo(loginLogo).MenuLogo(menuLogo).Favicon(favicon).Execute()

Set branding



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
	title := "title_example" // string | Title shown in the browser tab instead of TheHive. (optional)
	loginLogo := os.NewFile(1234, "some_file") // *os.File | Logo displayed on the login page. Accepts PNG or JPEG files, recommended size 200 x 200 pixels. (optional)
	menuLogo := os.NewFile(1234, "some_file") // *os.File | Logo displayed in the navigation bar. Accepts PNG or JPEG files, recommended size 84 x 84 pixels. (optional)
	favicon := os.NewFile(1234, "some_file") // *os.File | Icon displayed in the browser tab. Accepts PNG or JPEG files, recommended size 42 x 42 pixels. (optional)

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
 **title** | **string** | Title shown in the browser tab instead of TheHive. |
 **loginLogo** | ***os.File** | Logo displayed on the login page. Accepts PNG or JPEG files, recommended size 200 x 200 pixels. |
 **menuLogo** | ***os.File** | Logo displayed in the navigation bar. Accepts PNG or JPEG files, recommended size 84 x 84 pixels. |
 **favicon** | ***os.File** | Icon displayed in the browser tab. Accepts PNG or JPEG files, recommended size 42 x 42 pixels. |

### Return type

[**OutputBranding**](OutputBranding.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
