# \PageAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAPage**](PageAPI.md#CreateAPage) | **Post** /api/v1/page | Create a Knowledge Base page
[**CreateAPageInACase**](PageAPI.md#CreateAPageInACase) | **Post** /api/v1/case/{caseId}/page | Create a case page
[**DeleteAPage**](PageAPI.md#DeleteAPage) | **Delete** /api/v1/page/{pageId} | Delete a Knowledge Base page
[**DeleteAPageInACase**](PageAPI.md#DeleteAPageInACase) | **Delete** /api/v1/case/{caseId}/page/{pageId} | Delete a case page
[**UpdateAPage**](PageAPI.md#UpdateAPage) | **Patch** /api/v1/page/{pageId} | Update a Knowledge Base page
[**UpdateAPageInACase**](PageAPI.md#UpdateAPageInACase) | **Patch** /api/v1/case/{caseId}/page/{pageId} | Update a case page



## CreateAPage

> OutputPage CreateAPage(ctx).InputCreatePage(inputCreatePage).Execute()

Create a Knowledge Base page



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
	resp, r, err := apiClient.PageAPI.CreateAPage(context.Background()).InputCreatePage(inputCreatePage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PageAPI.CreateAPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAPage`: OutputPage
	fmt.Fprintf(os.Stdout, "Response from `PageAPI.CreateAPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputCreatePage** | [**InputCreatePage**](InputCreatePage.md) |  |

### Return type

[**OutputPage**](OutputPage.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAPageInACase

> OutputPage CreateAPageInACase(ctx, caseId).InputCreatePage(inputCreatePage).Execute()

Create a case page



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
	caseId := "caseId_example" // string | The case ID preceded by `~`, or the case number.
	inputCreatePage := *openapiclient.NewInputCreatePage("Title_example", "Content_example", "Category_example") // InputCreatePage |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PageAPI.CreateAPageInACase(context.Background(), caseId).InputCreatePage(inputCreatePage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PageAPI.CreateAPageInACase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAPageInACase`: OutputPage
	fmt.Fprintf(os.Stdout, "Response from `PageAPI.CreateAPageInACase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** | The case ID preceded by &#x60;~&#x60;, or the case number. |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAPageInACaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputCreatePage** | [**InputCreatePage**](InputCreatePage.md) |  |

### Return type

[**OutputPage**](OutputPage.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAPage

> DeleteAPage(ctx, pageId).Execute()

Delete a Knowledge Base page



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
	pageId := "pageId_example" // string | The page ID preceded by `~`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PageAPI.DeleteAPage(context.Background(), pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PageAPI.DeleteAPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | The page ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAPageRequest struct via the builder pattern


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


## DeleteAPageInACase

> DeleteAPageInACase(ctx, caseId, pageId).Execute()

Delete a case page



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
	caseId := "caseId_example" // string | The case ID preceded by `~`, or the case number.
	pageId := "pageId_example" // string | The page ID preceded by `~`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PageAPI.DeleteAPageInACase(context.Background(), caseId, pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PageAPI.DeleteAPageInACase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** | The case ID preceded by &#x60;~&#x60;, or the case number. |
**pageId** | **string** | The page ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAPageInACaseRequest struct via the builder pattern


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


## UpdateAPage

> UpdateAPage(ctx, pageId).InputUpdatePage(inputUpdatePage).Execute()

Update a Knowledge Base page



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
	pageId := "pageId_example" // string | The page ID preceded by `~`.
	inputUpdatePage := *openapiclient.NewInputUpdatePage() // InputUpdatePage |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PageAPI.UpdateAPage(context.Background(), pageId).InputUpdatePage(inputUpdatePage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PageAPI.UpdateAPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | The page ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputUpdatePage** | [**InputUpdatePage**](InputUpdatePage.md) |  |

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


## UpdateAPageInACase

> UpdateAPageInACase(ctx, caseId, pageId).InputUpdatePage(inputUpdatePage).Execute()

Update a case page



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
	caseId := "caseId_example" // string | The case ID preceded by `~`, or the case number.
	pageId := "pageId_example" // string | The page ID preceded by `~`.
	inputUpdatePage := *openapiclient.NewInputUpdatePage() // InputUpdatePage |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PageAPI.UpdateAPageInACase(context.Background(), caseId, pageId).InputUpdatePage(inputUpdatePage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PageAPI.UpdateAPageInACase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** | The case ID preceded by &#x60;~&#x60;, or the case number. |
**pageId** | **string** | The page ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAPageInACaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inputUpdatePage** | [**InputUpdatePage**](InputUpdatePage.md) |  |

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
