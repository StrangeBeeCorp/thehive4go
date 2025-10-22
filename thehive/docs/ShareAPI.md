# \ShareAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSharesByShareId**](ShareAPI.md#DeleteSharesByShareId) | **Delete** /api/v1/case/shares | 
[**ListSharesOfCase**](ShareAPI.md#ListSharesOfCase) | **Get** /api/v1/case/{caseId}/shares | 
[**ListSharesOfObservable**](ShareAPI.md#ListSharesOfObservable) | **Get** /api/v1/observable/{observableId}/shares | 
[**ListSharesOfTask**](ShareAPI.md#ListSharesOfTask) | **Get** /api/v1/task/{taskId}/shares | 
[**RemoveAShareFromCase**](ShareAPI.md#RemoveAShareFromCase) | **Delete** /api/v1/case/share/{shareId} | 
[**SetSharesForCase**](ShareAPI.md#SetSharesForCase) | **Put** /api/v1/case/{caseId}/shares | 
[**ShareACase**](ShareAPI.md#ShareACase) | **Post** /api/v1/case/{caseId}/shares | 
[**ShareATask**](ShareAPI.md#ShareATask) | **Post** /api/v1/task/{taskId}/shares | 
[**ShareAnObservable**](ShareAPI.md#ShareAnObservable) | **Post** /api/v1/observable/{observableId}/shares | 
[**UnshareACase**](ShareAPI.md#UnshareACase) | **Delete** /api/v1/case/{caseId}/shares | 
[**UnshareATask**](ShareAPI.md#UnshareATask) | **Delete** /api/v1/task/{taskId}/shares | 
[**UnshareAnObservable**](ShareAPI.md#UnshareAnObservable) | **Delete** /api/v1/observable/{observableId}/shares | 
[**UpdateAShare**](ShareAPI.md#UpdateAShare) | **Patch** /api/v1/case/share/{shareId} | 



## DeleteSharesByShareId

> DeleteSharesByShareId(ctx).DeleteAlertInBulkRequest(deleteAlertInBulkRequest).Execute()



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
	r, err := apiClient.ShareAPI.DeleteSharesByShareId(context.Background()).DeleteAlertInBulkRequest(deleteAlertInBulkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShareAPI.DeleteSharesByShareId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSharesByShareIdRequest struct via the builder pattern


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


## ListSharesOfCase

> []OutputShare ListSharesOfCase(ctx, caseId).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShareAPI.ListSharesOfCase(context.Background(), caseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShareAPI.ListSharesOfCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSharesOfCase`: []OutputShare
	fmt.Fprintf(os.Stdout, "Response from `ShareAPI.ListSharesOfCase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSharesOfCaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OutputShare**](OutputShare.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSharesOfObservable

> []OutputShare ListSharesOfObservable(ctx, observableId).Execute()



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
	observableId := "~354" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShareAPI.ListSharesOfObservable(context.Background(), observableId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShareAPI.ListSharesOfObservable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSharesOfObservable`: []OutputShare
	fmt.Fprintf(os.Stdout, "Response from `ShareAPI.ListSharesOfObservable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**observableId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSharesOfObservableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OutputShare**](OutputShare.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSharesOfTask

> []OutputShare ListSharesOfTask(ctx, taskId).Execute()



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
	taskId := "~354" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShareAPI.ListSharesOfTask(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShareAPI.ListSharesOfTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSharesOfTask`: []OutputShare
	fmt.Fprintf(os.Stdout, "Response from `ShareAPI.ListSharesOfTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSharesOfTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OutputShare**](OutputShare.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAShareFromCase

> RemoveAShareFromCase(ctx, shareId).Execute()



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
	shareId := "~354" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ShareAPI.RemoveAShareFromCase(context.Background(), shareId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShareAPI.RemoveAShareFromCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shareId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAShareFromCaseRequest struct via the builder pattern


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


## SetSharesForCase

> []OutputShare SetSharesForCase(ctx, caseId).InputCreateShares(inputCreateShares).Execute()





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
	inputCreateShares := *openapiclient.NewInputCreateShares() // InputCreateShares | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShareAPI.SetSharesForCase(context.Background(), caseId).InputCreateShares(inputCreateShares).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShareAPI.SetSharesForCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetSharesForCase`: []OutputShare
	fmt.Fprintf(os.Stdout, "Response from `ShareAPI.SetSharesForCase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSharesForCaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputCreateShares** | [**InputCreateShares**](InputCreateShares.md) |  | 

### Return type

[**[]OutputShare**](OutputShare.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShareACase

> []OutputShare ShareACase(ctx, caseId).InputCreateShares(inputCreateShares).Execute()





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
	inputCreateShares := *openapiclient.NewInputCreateShares() // InputCreateShares | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShareAPI.ShareACase(context.Background(), caseId).InputCreateShares(inputCreateShares).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShareAPI.ShareACase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShareACase`: []OutputShare
	fmt.Fprintf(os.Stdout, "Response from `ShareAPI.ShareACase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShareACaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputCreateShares** | [**InputCreateShares**](InputCreateShares.md) |  | 

### Return type

[**[]OutputShare**](OutputShare.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShareATask

> ShareATask(ctx, taskId).InputCreateShare(inputCreateShare).Execute()





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
	taskId := "~354" // string | 
	inputCreateShare := *openapiclient.NewInputCreateShare() // InputCreateShare | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ShareAPI.ShareATask(context.Background(), taskId).InputCreateShare(inputCreateShare).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShareAPI.ShareATask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShareATaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputCreateShare** | [**InputCreateShare**](InputCreateShare.md) |  | 

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


## ShareAnObservable

> ShareAnObservable(ctx, observableId).InputCreateShare(inputCreateShare).Execute()





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
	observableId := "~354" // string | 
	inputCreateShare := *openapiclient.NewInputCreateShare() // InputCreateShare | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ShareAPI.ShareAnObservable(context.Background(), observableId).InputCreateShare(inputCreateShare).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShareAPI.ShareAnObservable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**observableId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShareAnObservableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputCreateShare** | [**InputCreateShare**](InputCreateShare.md) |  | 

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


## UnshareACase

> UnshareACase(ctx, caseId).InputRemoveShares(inputRemoveShares).Execute()



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
	inputRemoveShares := *openapiclient.NewInputRemoveShares() // InputRemoveShares | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ShareAPI.UnshareACase(context.Background(), caseId).InputRemoveShares(inputRemoveShares).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShareAPI.UnshareACase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnshareACaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputRemoveShares** | [**InputRemoveShares**](InputRemoveShares.md) |  | 

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


## UnshareATask

> UnshareATask(ctx, taskId).InputRemoveShares(inputRemoveShares).Execute()



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
	taskId := "~354" // string | 
	inputRemoveShares := *openapiclient.NewInputRemoveShares() // InputRemoveShares | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ShareAPI.UnshareATask(context.Background(), taskId).InputRemoveShares(inputRemoveShares).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShareAPI.UnshareATask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnshareATaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputRemoveShares** | [**InputRemoveShares**](InputRemoveShares.md) |  | 

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


## UnshareAnObservable

> UnshareAnObservable(ctx, observableId).InputRemoveShares(inputRemoveShares).Execute()



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
	observableId := "~354" // string | 
	inputRemoveShares := *openapiclient.NewInputRemoveShares() // InputRemoveShares | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ShareAPI.UnshareAnObservable(context.Background(), observableId).InputRemoveShares(inputRemoveShares).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShareAPI.UnshareAnObservable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**observableId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnshareAnObservableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputRemoveShares** | [**InputRemoveShares**](InputRemoveShares.md) |  | 

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


## UpdateAShare

> UpdateAShare(ctx, shareId).InputUpdateShare(inputUpdateShare).Execute()



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
	shareId := "~354" // string | 
	inputUpdateShare := *openapiclient.NewInputUpdateShare("Profile_example") // InputUpdateShare | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ShareAPI.UpdateAShare(context.Background(), shareId).InputUpdateShare(inputUpdateShare).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShareAPI.UpdateAShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shareId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputUpdateShare** | [**InputUpdateShare**](InputUpdateShare.md) |  | 

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

