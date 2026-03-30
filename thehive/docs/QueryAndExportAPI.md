# \QueryAndExportAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportAPIUnstableRoute**](QueryAndExportAPI.md#ExportAPIUnstableRoute) | **Get** /api/v1/export | 
[**GetExportFieldsUnstableRoute**](QueryAndExportAPI.md#GetExportFieldsUnstableRoute) | **Get** /api/v1/export/_fields | 
[**QueryAPI**](QueryAndExportAPI.md#QueryAPI) | **Post** /api/v1/query | 



## ExportAPIUnstableRoute

> *os.File ExportAPIUnstableRoute(ctx).Query(query).Options(options).Execute()





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
	query := "query_example" // string | 
	options := "options_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueryAndExportAPI.ExportAPIUnstableRoute(context.Background()).Query(query).Options(options).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueryAndExportAPI.ExportAPIUnstableRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportAPIUnstableRoute`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `QueryAndExportAPI.ExportAPIUnstableRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportAPIUnstableRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** |  | 
 **options** | **string** |  | 

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


## GetExportFieldsUnstableRoute

> OutputExportFieldsMapping GetExportFieldsUnstableRoute(ctx).Execute()





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
	resp, r, err := apiClient.QueryAndExportAPI.GetExportFieldsUnstableRoute(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueryAndExportAPI.GetExportFieldsUnstableRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExportFieldsUnstableRoute`: OutputExportFieldsMapping
	fmt.Fprintf(os.Stdout, "Response from `QueryAndExportAPI.GetExportFieldsUnstableRoute`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetExportFieldsUnstableRouteRequest struct via the builder pattern


### Return type

[**OutputExportFieldsMapping**](OutputExportFieldsMapping.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryAPI

> interface{} QueryAPI(ctx).InputQuery(inputQuery).Execute()





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
	inputQuery := *openapiclient.NewInputQuery() // InputQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueryAndExportAPI.QueryAPI(context.Background()).InputQuery(inputQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueryAndExportAPI.QueryAPI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryAPI`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `QueryAndExportAPI.QueryAPI`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryAPIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputQuery** | [**InputQuery**](InputQuery.md) |  | 

### Return type

**interface{}**

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

