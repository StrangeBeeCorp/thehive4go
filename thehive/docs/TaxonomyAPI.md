# \TaxonomyAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateTaxonomy**](TaxonomyAPI.md#ActivateTaxonomy) | **Put** /api/v1/taxonomy/{taxonomyId}/activate | 
[**CreateTaxonomy**](TaxonomyAPI.md#CreateTaxonomy) | **Post** /api/v1/taxonomy | 
[**DeactivateTaxonomy**](TaxonomyAPI.md#DeactivateTaxonomy) | **Put** /api/v1/taxonomy/{taxonomyId}/deactivate | 
[**DeleteTaxonomy**](TaxonomyAPI.md#DeleteTaxonomy) | **Delete** /api/v1/taxonomy/{taxonomyId} | 
[**GetTaxonomy**](TaxonomyAPI.md#GetTaxonomy) | **Get** /api/v1/taxonomy/{taxonomyId} | 
[**ImportFromZipFile**](TaxonomyAPI.md#ImportFromZipFile) | **Post** /api/v1/taxonomy/import-zip | 



## ActivateTaxonomy

> ActivateTaxonomy(ctx, taxonomyId).Execute()



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
	taxonomyId := "~354" // string | 

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
**taxonomyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateTaxonomyRequest struct via the builder pattern


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


## CreateTaxonomy

> OutputTaxonomy CreateTaxonomy(ctx).InputTaxonomy(inputTaxonomy).Execute()



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

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateTaxonomy

> DeactivateTaxonomy(ctx, taxonomyId).Execute()



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
	taxonomyId := "~354" // string | 

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
**taxonomyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateTaxonomyRequest struct via the builder pattern


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


## DeleteTaxonomy

> DeleteTaxonomy(ctx, taxonomyId).Execute()



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
	taxonomyId := "~354" // string | 

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
**taxonomyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTaxonomyRequest struct via the builder pattern


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


## GetTaxonomy

> OutputTaxonomy GetTaxonomy(ctx, taxonomyId).Execute()



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
	taxonomyId := "~354" // string | 

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
**taxonomyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxonomyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputTaxonomy**](OutputTaxonomy.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportFromZipFile

> Ok ImportFromZipFile(ctx).File(file).Execute()



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
	file := os.NewFile(1234, "some_file") // *os.File | 

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
 **file** | ***os.File** |  | 

### Return type

[**Ok**](Ok.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

