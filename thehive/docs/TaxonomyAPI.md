# \TaxonomyAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateTaxonomy**](TaxonomyAPI.md#ActivateTaxonomy) | **Put** /api/v1/taxonomy/{taxonomyId}/activate | Activate a taxonomy
[**CreateTaxonomy**](TaxonomyAPI.md#CreateTaxonomy) | **Post** /api/v1/taxonomy | Create a taxonomy
[**DeactivateTaxonomy**](TaxonomyAPI.md#DeactivateTaxonomy) | **Put** /api/v1/taxonomy/{taxonomyId}/deactivate | Deactivate a taxonomy
[**DeleteTaxonomy**](TaxonomyAPI.md#DeleteTaxonomy) | **Delete** /api/v1/taxonomy/{taxonomyId} | Delete a taxonomy
[**GetTaxonomy**](TaxonomyAPI.md#GetTaxonomy) | **Get** /api/v1/taxonomy/{taxonomyId} | Get a taxonomy
[**ImportFromZipFile**](TaxonomyAPI.md#ImportFromZipFile) | **Post** /api/v1/taxonomy/import-zip | Import taxonomies from a ZIP archive



## ActivateTaxonomy

> ActivateTaxonomy(ctx, taxonomyId).Execute()

Activate a taxonomy



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
	taxonomyId := "taxonomyId_example" // string | The taxonomy ID preceded by `~`, or namespace name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaxonomyAPI.ActivateTaxonomy(context.Background(), taxonomyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyAPI.ActivateTaxonomy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxonomyId** | **string** | The taxonomy ID preceded by &#x60;~&#x60;, or namespace name. |

### Other Parameters

Other parameters are passed through a pointer to a apiActivateTaxonomyRequest struct via the builder pattern


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


## CreateTaxonomy

> OutputTaxonomy CreateTaxonomy(ctx).InputTaxonomy(inputTaxonomy).Execute()

Create a taxonomy



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
	inputTaxonomy := *openapiclient.NewInputTaxonomy("Namespace_example", "Description_example", int32(123)) // InputTaxonomy |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxonomyAPI.CreateTaxonomy(context.Background()).InputTaxonomy(inputTaxonomy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyAPI.CreateTaxonomy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTaxonomy`: OutputTaxonomy
	fmt.Fprintf(os.Stdout, "Response from `TaxonomyAPI.CreateTaxonomy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaxonomyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputTaxonomy** | [**InputTaxonomy**](InputTaxonomy.md) |  |

### Return type

[**OutputTaxonomy**](OutputTaxonomy.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateTaxonomy

> DeactivateTaxonomy(ctx, taxonomyId).Execute()

Deactivate a taxonomy



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
	taxonomyId := "taxonomyId_example" // string | The taxonomy ID preceded by `~`, or namespace name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaxonomyAPI.DeactivateTaxonomy(context.Background(), taxonomyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyAPI.DeactivateTaxonomy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxonomyId** | **string** | The taxonomy ID preceded by &#x60;~&#x60;, or namespace name. |

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateTaxonomyRequest struct via the builder pattern


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


## DeleteTaxonomy

> DeleteTaxonomy(ctx, taxonomyId).Execute()

Delete a taxonomy



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
	taxonomyId := "taxonomyId_example" // string | The taxonomy ID preceded by `~`, or namespace name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaxonomyAPI.DeleteTaxonomy(context.Background(), taxonomyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyAPI.DeleteTaxonomy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxonomyId** | **string** | The taxonomy ID preceded by &#x60;~&#x60;, or namespace name. |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTaxonomyRequest struct via the builder pattern


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


## GetTaxonomy

> OutputTaxonomy GetTaxonomy(ctx, taxonomyId).Execute()

Get a taxonomy



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
	taxonomyId := "taxonomyId_example" // string | The taxonomy ID preceded by `~`, or namespace name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxonomyAPI.GetTaxonomy(context.Background(), taxonomyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyAPI.GetTaxonomy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaxonomy`: OutputTaxonomy
	fmt.Fprintf(os.Stdout, "Response from `TaxonomyAPI.GetTaxonomy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxonomyId** | **string** | The taxonomy ID preceded by &#x60;~&#x60;, or namespace name. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxonomyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputTaxonomy**](OutputTaxonomy.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportFromZipFile

> Ok ImportFromZipFile(ctx).File(file).Execute()

Import taxonomies from a ZIP archive



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
	file := os.NewFile(1234, "some_file") // *os.File | ZIP archive containing at least one `machinetag.json` file following the MISP taxonomy JSON schema.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxonomyAPI.ImportFromZipFile(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyAPI.ImportFromZipFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportFromZipFile`: Ok
	fmt.Fprintf(os.Stdout, "Response from `TaxonomyAPI.ImportFromZipFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportFromZipFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | ZIP archive containing at least one &#x60;machinetag.json&#x60; file following the MISP taxonomy JSON schema. |

### Return type

[**Ok**](Ok.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
