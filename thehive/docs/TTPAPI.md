# \TTPAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProcedureForAlert**](TTPAPI.md#CreateProcedureForAlert) | **Post** /api/v1/alert/{alertId}/procedure |
[**CreateProcedureForCase**](TTPAPI.md#CreateProcedureForCase) | **Post** /api/v1/case/{caseId}/procedure |
[**CreateSeveralProceduresForAlert**](TTPAPI.md#CreateSeveralProceduresForAlert) | **Post** /api/v1/alert/{alertId}/procedures |
[**CreateSeveralProceduresForCase**](TTPAPI.md#CreateSeveralProceduresForCase) | **Post** /api/v1/case/{caseId}/procedures |
[**DeleteProcedure**](TTPAPI.md#DeleteProcedure) | **Delete** /api/v1/procedure/{procedureId} |
[**DeleteProceduresInBulk**](TTPAPI.md#DeleteProceduresInBulk) | **Post** /api/v1/procedure/delete/_bulk |
[**UpdateProcedure**](TTPAPI.md#UpdateProcedure) | **Patch** /api/v1/procedure/{procedureId} |



## CreateProcedureForAlert

> OutputProcedure CreateProcedureForAlert(ctx, alertId).InputProcedure(inputProcedure).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StrangeBee/TheHive4Go/thehive"
)

func main() {
	alertId := "~354" // string |
	inputProcedure := *openapiclient.NewInputProcedure("PatternId_example", int64(1640000000000)) // InputProcedure |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TTPAPI.CreateProcedureForAlert(context.Background(), alertId).InputProcedure(inputProcedure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TTPAPI.CreateProcedureForAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProcedureForAlert`: OutputProcedure
	fmt.Fprintf(os.Stdout, "Response from `TTPAPI.CreateProcedureForAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProcedureForAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputProcedure** | [**InputProcedure**](InputProcedure.md) |  |

### Return type

[**OutputProcedure**](OutputProcedure.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProcedureForCase

> OutputProcedure CreateProcedureForCase(ctx, caseId).InputProcedure(inputProcedure).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StrangeBee/TheHive4Go/thehive"
)

func main() {
	caseId := "~354" // string |
	inputProcedure := *openapiclient.NewInputProcedure("PatternId_example", int64(1640000000000)) // InputProcedure |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TTPAPI.CreateProcedureForCase(context.Background(), caseId).InputProcedure(inputProcedure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TTPAPI.CreateProcedureForCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProcedureForCase`: OutputProcedure
	fmt.Fprintf(os.Stdout, "Response from `TTPAPI.CreateProcedureForCase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProcedureForCaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputProcedure** | [**InputProcedure**](InputProcedure.md) |  |

### Return type

[**OutputProcedure**](OutputProcedure.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSeveralProceduresForAlert

> []OutputProcedure CreateSeveralProceduresForAlert(ctx, alertId).InputBulkProcedure(inputBulkProcedure).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StrangeBee/TheHive4Go/thehive"
)

func main() {
	alertId := "~354" // string |
	inputBulkProcedure := *openapiclient.NewInputBulkProcedure() // InputBulkProcedure |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TTPAPI.CreateSeveralProceduresForAlert(context.Background(), alertId).InputBulkProcedure(inputBulkProcedure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TTPAPI.CreateSeveralProceduresForAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSeveralProceduresForAlert`: []OutputProcedure
	fmt.Fprintf(os.Stdout, "Response from `TTPAPI.CreateSeveralProceduresForAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSeveralProceduresForAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputBulkProcedure** | [**InputBulkProcedure**](InputBulkProcedure.md) |  |

### Return type

[**[]OutputProcedure**](OutputProcedure.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSeveralProceduresForCase

> []OutputProcedure CreateSeveralProceduresForCase(ctx, caseId).InputBulkProcedure(inputBulkProcedure).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StrangeBee/TheHive4Go/thehive"
)

func main() {
	caseId := "~354" // string |
	inputBulkProcedure := *openapiclient.NewInputBulkProcedure() // InputBulkProcedure |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TTPAPI.CreateSeveralProceduresForCase(context.Background(), caseId).InputBulkProcedure(inputBulkProcedure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TTPAPI.CreateSeveralProceduresForCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSeveralProceduresForCase`: []OutputProcedure
	fmt.Fprintf(os.Stdout, "Response from `TTPAPI.CreateSeveralProceduresForCase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSeveralProceduresForCaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputBulkProcedure** | [**InputBulkProcedure**](InputBulkProcedure.md) |  |

### Return type

[**[]OutputProcedure**](OutputProcedure.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProcedure

> DeleteProcedure(ctx, procedureId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StrangeBee/TheHive4Go/thehive"
)

func main() {
	procedureId := "~354" // string |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TTPAPI.DeleteProcedure(context.Background(), procedureId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TTPAPI.DeleteProcedure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**procedureId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProcedureRequest struct via the builder pattern


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


## DeleteProceduresInBulk

> DeleteProceduresInBulk(ctx).DeleteAlertInBulkRequest(deleteAlertInBulkRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StrangeBee/TheHive4Go/thehive"
)

func main() {
	deleteAlertInBulkRequest := *openapiclient.NewDeleteAlertInBulkRequest([]string{"~354"}) // DeleteAlertInBulkRequest |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TTPAPI.DeleteProceduresInBulk(context.Background()).DeleteAlertInBulkRequest(deleteAlertInBulkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TTPAPI.DeleteProceduresInBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProceduresInBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteAlertInBulkRequest** | [**DeleteAlertInBulkRequest**](DeleteAlertInBulkRequest.md) |  |

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


## UpdateProcedure

> UpdateProcedure(ctx, procedureId).InputUpdateProcedure(inputUpdateProcedure).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StrangeBee/TheHive4Go/thehive"
)

func main() {
	procedureId := "~354" // string |
	inputUpdateProcedure := *openapiclient.NewInputUpdateProcedure() // InputUpdateProcedure |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TTPAPI.UpdateProcedure(context.Background(), procedureId).InputUpdateProcedure(inputUpdateProcedure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TTPAPI.UpdateProcedure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**procedureId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProcedureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputUpdateProcedure** | [**InputUpdateProcedure**](InputUpdateProcedure.md) |  |

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
