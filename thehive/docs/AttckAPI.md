# \AttckAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCatalogOfTTP**](AttckAPI.md#CreateCatalogOfTTP) | **Post** /api/v1/catalog | Create an attack technique catalog
[**DeleteACatalogOfTTP**](AttckAPI.md#DeleteACatalogOfTTP) | **Delete** /api/v1/catalog/{catalogId} | Delete an attack technique catalog
[**DeletePattern**](AttckAPI.md#DeletePattern) | **Delete** /api/v1/pattern/{patternId} | Delete an attack technique
[**GetCasePatterns**](AttckAPI.md#GetCasePatterns) | **Get** /api/v1/pattern/case/{caseId} | List attack techniques linked to a case
[**GetPattern**](AttckAPI.md#GetPattern) | **Get** /api/v1/pattern/{patternId} | Get an attack technique
[**ImportMITREAttckFile**](AttckAPI.md#ImportMITREAttckFile) | **Post** /api/v1/pattern/import/attack | Import MITRE ATT&amp;CK techniques
[**UpdateCatalogOfTTP**](AttckAPI.md#UpdateCatalogOfTTP) | **Patch** /api/v1/catalog/{catalogId} | Update an attack technique catalog



## CreateCatalogOfTTP

> OutputCatalogOfPattern CreateCatalogOfTTP(ctx).InputCatalogOfPattern(inputCatalogOfPattern).Execute()

Create an attack technique catalog



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
	inputCatalogOfPattern := *openapiclient.NewInputCatalogOfPattern("Name_example") // InputCatalogOfPattern |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttckAPI.CreateCatalogOfTTP(context.Background()).InputCatalogOfPattern(inputCatalogOfPattern).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttckAPI.CreateCatalogOfTTP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCatalogOfTTP`: OutputCatalogOfPattern
	fmt.Fprintf(os.Stdout, "Response from `AttckAPI.CreateCatalogOfTTP`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCatalogOfTTPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputCatalogOfPattern** | [**InputCatalogOfPattern**](InputCatalogOfPattern.md) |  |

### Return type

[**OutputCatalogOfPattern**](OutputCatalogOfPattern.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteACatalogOfTTP

> DeleteACatalogOfTTP(ctx, catalogId).Execute()

Delete an attack technique catalog



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
	catalogId := "catalogId_example" // string | The catalog ID preceded by `~`, or the catalog name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AttckAPI.DeleteACatalogOfTTP(context.Background(), catalogId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttckAPI.DeleteACatalogOfTTP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catalogId** | **string** | The catalog ID preceded by &#x60;~&#x60;, or the catalog name. |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteACatalogOfTTPRequest struct via the builder pattern


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


## DeletePattern

> DeletePattern(ctx, patternId).Execute()

Delete an attack technique



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
	patternId := "patternId_example" // string | The attack technique ID preceded by `~`, or its MITRE identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AttckAPI.DeletePattern(context.Background(), patternId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttckAPI.DeletePattern``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patternId** | **string** | The attack technique ID preceded by &#x60;~&#x60;, or its MITRE identifier. |

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePatternRequest struct via the builder pattern


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


## GetCasePatterns

> []OutputPattern GetCasePatterns(ctx, caseId).Execute()

List attack techniques linked to a case



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
	caseId := "caseId_example" // string | The case ID preceded by `~`, or case number.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttckAPI.GetCasePatterns(context.Background(), caseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttckAPI.GetCasePatterns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCasePatterns`: []OutputPattern
	fmt.Fprintf(os.Stdout, "Response from `AttckAPI.GetCasePatterns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** | The case ID preceded by &#x60;~&#x60;, or case number. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetCasePatternsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OutputPattern**](OutputPattern.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPattern

> OutputPattern GetPattern(ctx, patternId).Execute()

Get an attack technique



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
	patternId := "patternId_example" // string | The attack technique ID preceded by `~`, or its MITRE identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttckAPI.GetPattern(context.Background(), patternId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttckAPI.GetPattern``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPattern`: OutputPattern
	fmt.Fprintf(os.Stdout, "Response from `AttckAPI.GetPattern`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patternId** | **string** | The attack technique ID preceded by &#x60;~&#x60;, or its MITRE identifier. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetPatternRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputPattern**](OutputPattern.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportMITREAttckFile

> map[string]interface{} ImportMITREAttckFile(ctx).InputPatternImportMitre(inputPatternImportMitre).Execute()

Import MITRE ATT&CK techniques



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
	inputPatternImportMitre := *openapiclient.NewInputPatternImportMitre("Catalog_example") // InputPatternImportMitre |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttckAPI.ImportMITREAttckFile(context.Background()).InputPatternImportMitre(inputPatternImportMitre).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttckAPI.ImportMITREAttckFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportMITREAttckFile`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AttckAPI.ImportMITREAttckFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportMITREAttckFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputPatternImportMitre** | [**InputPatternImportMitre**](InputPatternImportMitre.md) |  |

### Return type

**map[string]interface{}**

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCatalogOfTTP

> UpdateCatalogOfTTP(ctx, catalogId).InputUpdateCatalogOfPattern(inputUpdateCatalogOfPattern).Execute()

Update an attack technique catalog



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
	catalogId := "catalogId_example" // string | The catalog ID preceded by `~`, or the catalog name.
	inputUpdateCatalogOfPattern := *openapiclient.NewInputUpdateCatalogOfPattern() // InputUpdateCatalogOfPattern |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AttckAPI.UpdateCatalogOfTTP(context.Background(), catalogId).InputUpdateCatalogOfPattern(inputUpdateCatalogOfPattern).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttckAPI.UpdateCatalogOfTTP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catalogId** | **string** | The catalog ID preceded by &#x60;~&#x60;, or the catalog name. |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCatalogOfTTPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputUpdateCatalogOfPattern** | [**InputUpdateCatalogOfPattern**](InputUpdateCatalogOfPattern.md) |  |

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
