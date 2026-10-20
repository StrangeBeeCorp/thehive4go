# \AuditAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFlowOfAudit**](AuditAPI.md#GetFlowOfAudit) | **Get** /api/v1/flow | Get the audit trail of an object



## GetFlowOfAudit

> []OutputStreamAudit GetFlowOfAudit(ctx).RootId(rootId).Count(count).Execute()

Get the audit trail of an object



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
	rootId := "rootId_example" // string | Case ID to scope the audit trail to a single case and its linked objects. Unlike other case parameters, this doesn't accept the case number. Must match an existing case. Use `POST /api/v1/query` with `listCase` to retrieve available cases and their IDs. Omit this parameter, or set it to `any`, to retrieve the most recent audit entries across all object types visible to you. (optional)
	count := int32(56) // int32 | Maximum number of audit entries to return. (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditAPI.GetFlowOfAudit(context.Background()).RootId(rootId).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditAPI.GetFlowOfAudit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlowOfAudit`: []OutputStreamAudit
	fmt.Fprintf(os.Stdout, "Response from `AuditAPI.GetFlowOfAudit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFlowOfAuditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rootId** | **string** | Case ID to scope the audit trail to a single case and its linked objects. Unlike other case parameters, this doesn&#39;t accept the case number. Must match an existing case. Use &#x60;POST /api/v1/query&#x60; with &#x60;listCase&#x60; to retrieve available cases and their IDs. Omit this parameter, or set it to &#x60;any&#x60;, to retrieve the most recent audit entries across all object types visible to you. |
 **count** | **int32** | Maximum number of audit entries to return. | [default to 10]

### Return type

[**[]OutputStreamAudit**](OutputStreamAudit.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
