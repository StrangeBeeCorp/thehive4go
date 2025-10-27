# \PageTemplateAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAPageTemplate**](PageTemplateAPI.md#CreateAPageTemplate) | **Post** /api/v1/pageTemplate |
[**DeleteAPageTemplate**](PageTemplateAPI.md#DeleteAPageTemplate) | **Delete** /api/v1/pageTemplate/{pageTemplateId} |
[**UpdateAPageTemplate**](PageTemplateAPI.md#UpdateAPageTemplate) | **Patch** /api/v1/pageTemplate/{pageTemplateId} |



## CreateAPageTemplate

> OutputPage CreateAPageTemplate(ctx).InputCreatePage(inputCreatePage).Execute()



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
	inputCreatePage := *openapiclient.NewInputCreatePage("Title_example", "Content_example", "Category_example") // InputCreatePage |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PageTemplateAPI.CreateAPageTemplate(context.Background()).InputCreatePage(inputCreatePage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PageTemplateAPI.CreateAPageTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAPageTemplate`: OutputPage
	fmt.Fprintf(os.Stdout, "Response from `PageTemplateAPI.CreateAPageTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAPageTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputCreatePage** | [**InputCreatePage**](InputCreatePage.md) |  |

### Return type

[**OutputPage**](OutputPage.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAPageTemplate

> DeleteAPageTemplate(ctx, pageTemplateId).Execute()



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
	pageTemplateId := "~354" // string |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PageTemplateAPI.DeleteAPageTemplate(context.Background(), pageTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PageTemplateAPI.DeleteAPageTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageTemplateId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAPageTemplateRequest struct via the builder pattern


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


## UpdateAPageTemplate

> UpdateAPageTemplate(ctx, pageTemplateId).InputUpdatePage(inputUpdatePage).Execute()



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
	pageTemplateId := "~354" // string |
	inputUpdatePage := *openapiclient.NewInputUpdatePage() // InputUpdatePage |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PageTemplateAPI.UpdateAPageTemplate(context.Background(), pageTemplateId).InputUpdatePage(inputUpdatePage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PageTemplateAPI.UpdateAPageTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageTemplateId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAPageTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputUpdatePage** | [**InputUpdatePage**](InputUpdatePage.md) |  |

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
