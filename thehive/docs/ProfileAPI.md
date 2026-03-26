# \ProfileAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProfile**](ProfileAPI.md#CreateProfile) | **Post** /api/v1/profile | 
[**DeleteProfile**](ProfileAPI.md#DeleteProfile) | **Delete** /api/v1/profile/{profileId} | 
[**GetProfile**](ProfileAPI.md#GetProfile) | **Get** /api/v1/profile/{profileId} | 
[**UpdateProfile**](ProfileAPI.md#UpdateProfile) | **Patch** /api/v1/profile/{profileId} | 



## CreateProfile

> OutputProfile CreateProfile(ctx).InputCreateProfile(inputCreateProfile).Execute()



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
	inputCreateProfile := *openapiclient.NewInputCreateProfile("Name_example") // InputCreateProfile | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfileAPI.CreateProfile(context.Background()).InputCreateProfile(inputCreateProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfileAPI.CreateProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProfile`: OutputProfile
	fmt.Fprintf(os.Stdout, "Response from `ProfileAPI.CreateProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputCreateProfile** | [**InputCreateProfile**](InputCreateProfile.md) |  | 

### Return type

[**OutputProfile**](OutputProfile.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProfile

> DeleteProfile(ctx, profileId).Execute()



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
	profileId := "profileId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProfileAPI.DeleteProfile(context.Background(), profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfileAPI.DeleteProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProfileRequest struct via the builder pattern


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


## GetProfile

> OutputProfile GetProfile(ctx, profileId).Execute()



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
	profileId := "profileId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfileAPI.GetProfile(context.Background(), profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfileAPI.GetProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProfile`: OutputProfile
	fmt.Fprintf(os.Stdout, "Response from `ProfileAPI.GetProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputProfile**](OutputProfile.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProfile

> UpdateProfile(ctx, profileId).InputUpdateProfile(inputUpdateProfile).Execute()



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
	profileId := "profileId_example" // string | 
	inputUpdateProfile := *openapiclient.NewInputUpdateProfile() // InputUpdateProfile | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProfileAPI.UpdateProfile(context.Background(), profileId).InputUpdateProfile(inputUpdateProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfileAPI.UpdateProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputUpdateProfile** | [**InputUpdateProfile**](InputUpdateProfile.md) |  | 

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

