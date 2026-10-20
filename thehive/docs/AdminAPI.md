# \AdminAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateLogLevel**](AdminAPI.md#UpdateLogLevel) | **Put** /api/v1/admin/log/set/{packageName}/{level} | Set the log level for a package



## UpdateLogLevel

> UpdateLogLevel(ctx, packageName, level).Execute()

Set the log level for a package



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
	packageName := "packageName_example" // string | The logger name.
	level := openapiclient.InputLogLevel("ALL") // InputLogLevel | `ALL` and `TRACE` log the most detailed information. `DEBUG` logs debugging information. `INFO` logs general information. `WARN` logs warnings and errors. `ERROR` logs errors only. `OFF` turns off logging entirely.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdminAPI.UpdateLogLevel(context.Background(), packageName, level).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.UpdateLogLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageName** | **string** | The logger name. |
**level** | [**InputLogLevel**](.md) | &#x60;ALL&#x60; and &#x60;TRACE&#x60; log the most detailed information. &#x60;DEBUG&#x60; logs debugging information. &#x60;INFO&#x60; logs general information. &#x60;WARN&#x60; logs warnings and errors. &#x60;ERROR&#x60; logs errors only. &#x60;OFF&#x60; turns off logging entirely. |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogLevelRequest struct via the builder pattern


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
