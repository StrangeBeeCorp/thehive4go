# \CommentAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCommentInAlert**](CommentAPI.md#CreateCommentInAlert) | **Post** /api/v1/alert/{alertId}/comment | Add a comment to an alert
[**CreateCommentInCase**](CommentAPI.md#CreateCommentInCase) | **Post** /api/v1/case/{caseId}/comment | Add a comment to a case
[**DeleteComment**](CommentAPI.md#DeleteComment) | **Delete** /api/v1/comment/{commentId} | Delete a comment
[**UpdateComment**](CommentAPI.md#UpdateComment) | **Patch** /api/v1/comment/{commentId} | Update a comment



## CreateCommentInAlert

> OutputComment CreateCommentInAlert(ctx, alertId).InputComment(inputComment).Execute()

Add a comment to an alert



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
	alertId := "alertId_example" // string | The alert ID preceded by `~`.
	inputComment := *openapiclient.NewInputComment("Message_example") // InputComment |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentAPI.CreateCommentInAlert(context.Background(), alertId).InputComment(inputComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentAPI.CreateCommentInAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCommentInAlert`: OutputComment
	fmt.Fprintf(os.Stdout, "Response from `CommentAPI.CreateCommentInAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string** | The alert ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCommentInAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputComment** | [**InputComment**](InputComment.md) |  |

### Return type

[**OutputComment**](OutputComment.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCommentInCase

> OutputComment CreateCommentInCase(ctx, caseId).InputComment(inputComment).Execute()

Add a comment to a case



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
	caseId := "caseId_example" // string | The case ID preceded by `~`, or the case number.
	inputComment := *openapiclient.NewInputComment("Message_example") // InputComment |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentAPI.CreateCommentInCase(context.Background(), caseId).InputComment(inputComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentAPI.CreateCommentInCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCommentInCase`: OutputComment
	fmt.Fprintf(os.Stdout, "Response from `CommentAPI.CreateCommentInCase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** | The case ID preceded by &#x60;~&#x60;, or the case number. |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCommentInCaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputComment** | [**InputComment**](InputComment.md) |  |

### Return type

[**OutputComment**](OutputComment.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteComment

> DeleteComment(ctx, commentId).Execute()

Delete a comment



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
	commentId := "commentId_example" // string | The comment ID preceded by `~`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommentAPI.DeleteComment(context.Background(), commentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentAPI.DeleteComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** | The comment ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommentRequest struct via the builder pattern


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


## UpdateComment

> UpdateComment(ctx, commentId).InputComment(inputComment).Execute()

Update a comment



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
	commentId := "commentId_example" // string | The comment ID preceded by `~`.
	inputComment := *openapiclient.NewInputComment("Message_example") // InputComment |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommentAPI.UpdateComment(context.Background(), commentId).InputComment(inputComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentAPI.UpdateComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** | The comment ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputComment** | [**InputComment**](InputComment.md) |  |

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
