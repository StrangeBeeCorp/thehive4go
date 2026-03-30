# \EmailIntakeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNewConfig**](EmailIntakeAPI.md#AddNewConfig) | **Post** /api/v1/connector/email-intake/config | 
[**AvailableProviders**](EmailIntakeAPI.md#AvailableProviders) | **Get** /api/v1/connector/email-intake/providers | 
[**DeleteConfig**](EmailIntakeAPI.md#DeleteConfig) | **Delete** /api/v1/connector/email-intake/config/{configId} | 
[**GWSetAuthorizationCode**](EmailIntakeAPI.md#GWSetAuthorizationCode) | **Post** /api/v1/connector/email-intake/config/authorizationCode | 
[**GetConfig**](EmailIntakeAPI.md#GetConfig) | **Get** /api/v1/connector/email-intake/config/{configId} | 
[**GetConfigs**](EmailIntakeAPI.md#GetConfigs) | **Get** /api/v1/connector/email-intake/configs | 
[**GetMailboxFolders**](EmailIntakeAPI.md#GetMailboxFolders) | **Post** /api/v1/connector/email-intake/folders | 
[**Sync**](EmailIntakeAPI.md#Sync) | **Post** /api/v1/connector/email-intake/sync | 
[**TestConfig**](EmailIntakeAPI.md#TestConfig) | **Post** /api/v1/connector/email-intake/config/test | 
[**UpdateConfig**](EmailIntakeAPI.md#UpdateConfig) | **Put** /api/v1/connector/email-intake/config/{configId} | 
[**UpdateConfigs**](EmailIntakeAPI.md#UpdateConfigs) | **Put** /api/v1/connector/email-intake/configs | 



## AddNewConfig

> OutputEmailIntakeConfig AddNewConfig(ctx).InputEmailIntakeConfig(inputEmailIntakeConfig).Execute()



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
	inputEmailIntakeConfig := *openapiclient.NewInputEmailIntakeConfig("Name_example", openapiclient.InputEmailIntakeMailboxConfig{InputEmailIntakeApiMailbox: openapiclient.NewInputEmailIntakeApiMailbox(*openapiclient.NewInputEmailIntakeApiProvider(openapiclient.EmailIntakeProviderName("imap")), *openapiclient.NewInputEmailIntakeCredential("Email_example"))}) // InputEmailIntakeConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailIntakeAPI.AddNewConfig(context.Background()).InputEmailIntakeConfig(inputEmailIntakeConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailIntakeAPI.AddNewConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddNewConfig`: OutputEmailIntakeConfig
	fmt.Fprintf(os.Stdout, "Response from `EmailIntakeAPI.AddNewConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddNewConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputEmailIntakeConfig** | [**InputEmailIntakeConfig**](InputEmailIntakeConfig.md) |  | 

### Return type

[**OutputEmailIntakeConfig**](OutputEmailIntakeConfig.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AvailableProviders

> []OutputRichEmailIntakeProvider AvailableProviders(ctx).Execute()





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
	resp, r, err := apiClient.EmailIntakeAPI.AvailableProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailIntakeAPI.AvailableProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AvailableProviders`: []OutputRichEmailIntakeProvider
	fmt.Fprintf(os.Stdout, "Response from `EmailIntakeAPI.AvailableProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAvailableProvidersRequest struct via the builder pattern


### Return type

[**[]OutputRichEmailIntakeProvider**](OutputRichEmailIntakeProvider.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConfig

> DeleteConfig(ctx, configId).Execute()



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
	configId := "configId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EmailIntakeAPI.DeleteConfig(context.Background(), configId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailIntakeAPI.DeleteConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfigRequest struct via the builder pattern


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


## GWSetAuthorizationCode

> interface{} GWSetAuthorizationCode(ctx).Body(body).Execute()





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
	body := interface{}(987) // interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailIntakeAPI.GWSetAuthorizationCode(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailIntakeAPI.GWSetAuthorizationCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GWSetAuthorizationCode`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `EmailIntakeAPI.GWSetAuthorizationCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGWSetAuthorizationCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **interface{}** |  | 

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


## GetConfig

> OutputEmailIntakeConfig GetConfig(ctx, configId).Execute()



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
	configId := "configId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailIntakeAPI.GetConfig(context.Background(), configId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailIntakeAPI.GetConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfig`: OutputEmailIntakeConfig
	fmt.Fprintf(os.Stdout, "Response from `EmailIntakeAPI.GetConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputEmailIntakeConfig**](OutputEmailIntakeConfig.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigs

> OutputRichEmailIntake GetConfigs(ctx).Execute()





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
	resp, r, err := apiClient.EmailIntakeAPI.GetConfigs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailIntakeAPI.GetConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfigs`: OutputRichEmailIntake
	fmt.Fprintf(os.Stdout, "Response from `EmailIntakeAPI.GetConfigs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigsRequest struct via the builder pattern


### Return type

[**OutputRichEmailIntake**](OutputRichEmailIntake.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMailboxFolders

> []string GetMailboxFolders(ctx).InputEmailIntakeConfig(inputEmailIntakeConfig).Execute()





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
	inputEmailIntakeConfig := *openapiclient.NewInputEmailIntakeConfig("Name_example", openapiclient.InputEmailIntakeMailboxConfig{InputEmailIntakeApiMailbox: openapiclient.NewInputEmailIntakeApiMailbox(*openapiclient.NewInputEmailIntakeApiProvider(openapiclient.EmailIntakeProviderName("imap")), *openapiclient.NewInputEmailIntakeCredential("Email_example"))}) // InputEmailIntakeConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailIntakeAPI.GetMailboxFolders(context.Background()).InputEmailIntakeConfig(inputEmailIntakeConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailIntakeAPI.GetMailboxFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMailboxFolders`: []string
	fmt.Fprintf(os.Stdout, "Response from `EmailIntakeAPI.GetMailboxFolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMailboxFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputEmailIntakeConfig** | [**InputEmailIntakeConfig**](InputEmailIntakeConfig.md) |  | 

### Return type

**[]string**

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Sync

> Sync(ctx).Execute()





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
	r, err := apiClient.EmailIntakeAPI.Sync(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailIntakeAPI.Sync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSyncRequest struct via the builder pattern


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


## TestConfig

> TestConfig(ctx).InputEmailIntakeConfig(inputEmailIntakeConfig).Execute()





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
	inputEmailIntakeConfig := *openapiclient.NewInputEmailIntakeConfig("Name_example", openapiclient.InputEmailIntakeMailboxConfig{InputEmailIntakeApiMailbox: openapiclient.NewInputEmailIntakeApiMailbox(*openapiclient.NewInputEmailIntakeApiProvider(openapiclient.EmailIntakeProviderName("imap")), *openapiclient.NewInputEmailIntakeCredential("Email_example"))}) // InputEmailIntakeConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EmailIntakeAPI.TestConfig(context.Background()).InputEmailIntakeConfig(inputEmailIntakeConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailIntakeAPI.TestConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputEmailIntakeConfig** | [**InputEmailIntakeConfig**](InputEmailIntakeConfig.md) |  | 

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


## UpdateConfig

> UpdateConfig(ctx, configId).InputEmailIntakeConfig(inputEmailIntakeConfig).Execute()



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
	configId := "configId_example" // string | 
	inputEmailIntakeConfig := *openapiclient.NewInputEmailIntakeConfig("Name_example", openapiclient.InputEmailIntakeMailboxConfig{InputEmailIntakeApiMailbox: openapiclient.NewInputEmailIntakeApiMailbox(*openapiclient.NewInputEmailIntakeApiProvider(openapiclient.EmailIntakeProviderName("imap")), *openapiclient.NewInputEmailIntakeCredential("Email_example"))}) // InputEmailIntakeConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EmailIntakeAPI.UpdateConfig(context.Background(), configId).InputEmailIntakeConfig(inputEmailIntakeConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailIntakeAPI.UpdateConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputEmailIntakeConfig** | [**InputEmailIntakeConfig**](InputEmailIntakeConfig.md) |  | 

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


## UpdateConfigs

> UpdateConfigs(ctx).InputEmailIntake(inputEmailIntake).Execute()





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
	inputEmailIntake := *openapiclient.NewInputEmailIntake("Interval_example") // InputEmailIntake | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EmailIntakeAPI.UpdateConfigs(context.Background()).InputEmailIntake(inputEmailIntake).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailIntakeAPI.UpdateConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputEmailIntake** | [**InputEmailIntake**](InputEmailIntake.md) |  | 

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

