# \CustomFieldAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomField**](CustomFieldAPI.md#CreateCustomField) | **Post** /api/v1/customField | Create a custom field
[**DeleteCustomField**](CustomFieldAPI.md#DeleteCustomField) | **Delete** /api/v1/customField/{customFieldId} | Delete a custom field
[**ListCustomFields**](CustomFieldAPI.md#ListCustomFields) | **Get** /api/v1/customField | List all custom fields
[**UpdateCustomField**](CustomFieldAPI.md#UpdateCustomField) | **Patch** /api/v1/customField/{customFieldId} | Update a custom field



## CreateCustomField

> OutputCustomField CreateCustomField(ctx).InputCustomField(inputCustomField).Execute()

Create a custom field



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
	inputCustomField := *openapiclient.NewInputCustomField("Name_example", "Group_example", "Description_example", openapiclient.CustomFieldType("string")) // InputCustomField |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFieldAPI.CreateCustomField(context.Background()).InputCustomField(inputCustomField).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldAPI.CreateCustomField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomField`: OutputCustomField
	fmt.Fprintf(os.Stdout, "Response from `CustomFieldAPI.CreateCustomField`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputCustomField** | [**InputCustomField**](InputCustomField.md) |  |

### Return type

[**OutputCustomField**](OutputCustomField.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomField

> DeleteCustomField(ctx, customFieldId).Force(force).Execute()

Delete a custom field



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
	customFieldId := "customFieldId_example" // string | The custom field ID preceded by `~`, or custom field name.
	force := true // bool | If `true`, deletes the custom field even when assigned to cases or alerts. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomFieldAPI.DeleteCustomField(context.Background(), customFieldId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldAPI.DeleteCustomField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customFieldId** | **string** | The custom field ID preceded by &#x60;~&#x60;, or custom field name. |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | If &#x60;true&#x60;, deletes the custom field even when assigned to cases or alerts. | [default to false]

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


## ListCustomFields

> []OutputCustomField ListCustomFields(ctx).Execute()

List all custom fields



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
	resp, r, err := apiClient.CustomFieldAPI.ListCustomFields(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldAPI.ListCustomFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCustomFields`: []OutputCustomField
	fmt.Fprintf(os.Stdout, "Response from `CustomFieldAPI.ListCustomFields`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCustomFieldsRequest struct via the builder pattern


### Return type

[**[]OutputCustomField**](OutputCustomField.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomField

> UpdateCustomField(ctx, customFieldId).InputUpdateCustomField(inputUpdateCustomField).Execute()

Update a custom field



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
	customFieldId := "customFieldId_example" // string | The custom field ID preceded by `~`, or custom field name.
	inputUpdateCustomField := *openapiclient.NewInputUpdateCustomField() // InputUpdateCustomField |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomFieldAPI.UpdateCustomField(context.Background(), customFieldId).InputUpdateCustomField(inputUpdateCustomField).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldAPI.UpdateCustomField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customFieldId** | **string** | The custom field ID preceded by &#x60;~&#x60;, or custom field name. |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputUpdateCustomField** | [**InputUpdateCustomField**](InputUpdateCustomField.md) |  |

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
