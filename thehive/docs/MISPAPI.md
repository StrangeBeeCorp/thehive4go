# \MISPAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportCaseToMISP**](MISPAPI.md#ExportCaseToMISP) | **Post** /api/v1/connector/misp/export/{caseId}/{mispName} |
[**GetMISPStatus**](MISPAPI.md#GetMISPStatus) | **Get** /api/v1/connector/misp/status |
[**ImportCaseFromMISP**](MISPAPI.md#ImportCaseFromMISP) | **Post** /api/v1/connector/misp/case/import |
[**SyncWithMISPServers**](MISPAPI.md#SyncWithMISPServers) | **Get** /api/v1/connector/misp/_syncAlerts |



## ExportCaseToMISP

> ExportCaseToMISP(ctx, caseId, mispName).Execute()



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
	caseId := "~354" // string |
	mispName := "mispName_example" // string |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MISPAPI.ExportCaseToMISP(context.Background(), caseId, mispName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MISPAPI.ExportCaseToMISP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  |
**mispName** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiExportCaseToMISPRequest struct via the builder pattern


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


## GetMISPStatus

> map[string]interface{} GetMISPStatus(ctx).Execute()



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
	resp, r, err := apiClient.MISPAPI.GetMISPStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MISPAPI.GetMISPStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMISPStatus`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MISPAPI.GetMISPStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMISPStatusRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportCaseFromMISP

> OutputCase ImportCaseFromMISP(ctx).Json(json).File(file).Execute()



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
	json := *openapiclient.NewInputImportCase1() // InputImportCase1 |
	file := os.NewFile(1234, "some_file") // *os.File |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MISPAPI.ImportCaseFromMISP(context.Background()).Json(json).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MISPAPI.ImportCaseFromMISP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportCaseFromMISP`: OutputCase
	fmt.Fprintf(os.Stdout, "Response from `MISPAPI.ImportCaseFromMISP`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportCaseFromMISPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **json** | [**InputImportCase1**](InputImportCase1.md) |  |
 **file** | ***os.File** |  |

### Return type

[**OutputCase**](OutputCase.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncWithMISPServers

> SyncWithMISPServers(ctx).Execute()



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
	r, err := apiClient.MISPAPI.SyncWithMISPServers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MISPAPI.SyncWithMISPServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSyncWithMISPServersRequest struct via the builder pattern


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
