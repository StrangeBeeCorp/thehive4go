# \OrganizationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAttachmentToOrganisation**](OrganizationAPI.md#AddAttachmentToOrganisation) | **Post** /api/v1/attachment | Add files to an organization
[**BulkLinkOrganisations**](OrganizationAPI.md#BulkLinkOrganisations) | **Put** /api/v1/organisation/{orgId}/links | Replace all links for an organization
[**CreateOrganisation**](OrganizationAPI.md#CreateOrganisation) | **Post** /api/v1/organisation | Create an organization
[**DeleteAttachment**](OrganizationAPI.md#DeleteAttachment) | **Delete** /api/v1/attachment/{attachmentId} | Remove a file from an organization
[**DownloadAttachment**](OrganizationAPI.md#DownloadAttachment) | **Get** /api/v1/attachment/{attachmentId}/download | Download a file from an organization
[**GetAttachment**](OrganizationAPI.md#GetAttachment) | **Get** /api/v1/attachment/{attachmentId} | Get a file from an organization
[**GetOrganisation**](OrganizationAPI.md#GetOrganisation) | **Get** /api/v1/organisation/{orgId} | Get an organization
[**GetOrganisationAvatar**](OrganizationAPI.md#GetOrganisationAvatar) | **Get** /api/v1/organisation/{orgId}/avatar/{fileHash} | Get the avatar of an organization
[**LinkOrganisations**](OrganizationAPI.md#LinkOrganisations) | **Put** /api/v1/organisation/{orgId}/link/{otherOrgId} | Link two organizations
[**ListOrganisationLinks**](OrganizationAPI.md#ListOrganisationLinks) | **Get** /api/v1/organisation/{orgId}/links | List organization links
[**ListSharingProfiles**](OrganizationAPI.md#ListSharingProfiles) | **Get** /api/v1/sharingProfile | List sharing profiles
[**UnlinkOrganisations**](OrganizationAPI.md#UnlinkOrganisations) | **Delete** /api/v1/organisation/{orgId}/link/{otherOrgId} | Remove a link between two organizations
[**UpdateOrganisation**](OrganizationAPI.md#UpdateOrganisation) | **Patch** /api/v1/organisation/{orgId} | Update an organization



## AddAttachmentToOrganisation

> OutputAttachments AddAttachmentToOrganisation(ctx).Attachments(attachments).CanRename(canRename).Execute()

Add files to an organization



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
	attachments := []*os.File{"TODO"} // []*os.File | Files to attach. Each part is a multipart form file field.
	canRename := true // bool | When `true`, renames files that already exist with the same name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.AddAttachmentToOrganisation(context.Background()).Attachments(attachments).CanRename(canRename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.AddAttachmentToOrganisation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddAttachmentToOrganisation`: OutputAttachments
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.AddAttachmentToOrganisation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAttachmentToOrganisationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attachments** | **[]*os.File** | Files to attach. Each part is a multipart form file field. |
 **canRename** | **bool** | When &#x60;true&#x60;, renames files that already exist with the same name. |

### Return type

[**OutputAttachments**](OutputAttachments.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkLinkOrganisations

> BulkLinkOrganisations(ctx, orgId).InputOrganisationBulkLink(inputOrganisationBulkLink).Execute()

Replace all links for an organization



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
	orgId := "orgId_example" // string | The organization ID preceded by `~`, or organization name.
	inputOrganisationBulkLink := *openapiclient.NewInputOrganisationBulkLink() // InputOrganisationBulkLink |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.BulkLinkOrganisations(context.Background(), orgId).InputOrganisationBulkLink(inputOrganisationBulkLink).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.BulkLinkOrganisations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization ID preceded by &#x60;~&#x60;, or organization name. |

### Other Parameters

Other parameters are passed through a pointer to a apiBulkLinkOrganisationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputOrganisationBulkLink** | [**InputOrganisationBulkLink**](InputOrganisationBulkLink.md) |  |

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


## CreateOrganisation

> OutputOrganisation CreateOrganisation(ctx).InputCreateOrganisation(inputCreateOrganisation).Execute()

Create an organization



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
	inputCreateOrganisation := *openapiclient.NewInputCreateOrganisation("Name_example", "Description_example") // InputCreateOrganisation |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.CreateOrganisation(context.Background()).InputCreateOrganisation(inputCreateOrganisation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.CreateOrganisation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganisation`: OutputOrganisation
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.CreateOrganisation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganisationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputCreateOrganisation** | [**InputCreateOrganisation**](InputCreateOrganisation.md) |  |

### Return type

[**OutputOrganisation**](OutputOrganisation.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAttachment

> DeleteAttachment(ctx, attachmentId).Execute()

Remove a file from an organization



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
	attachmentId := "attachmentId_example" // string | The attachment ID preceded by `~`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.DeleteAttachment(context.Background(), attachmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.DeleteAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string** | The attachment ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAttachmentRequest struct via the builder pattern


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


## DownloadAttachment

> *os.File DownloadAttachment(ctx, attachmentId).Execute()

Download a file from an organization



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
	attachmentId := "attachmentId_example" // string | The attachment ID preceded by `~`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.DownloadAttachment(context.Background(), attachmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.DownloadAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadAttachment`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.DownloadAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string** | The attachment ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttachment

> *os.File GetAttachment(ctx, attachmentId).IfNoneMatch(ifNoneMatch).Execute()

Get a file from an organization



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
	attachmentId := "attachmentId_example" // string | The attachment ID preceded by `~`.
	ifNoneMatch := "ifNoneMatch_example" // string | ETag from a previous response. If it matches, returns `304 Not Modified` instead of the file content. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.GetAttachment(context.Background(), attachmentId).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachment`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string** | The attachment ID preceded by &#x60;~&#x60;. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifNoneMatch** | **string** | ETag from a previous response. If it matches, returns &#x60;304 Not Modified&#x60; instead of the file content. |

### Return type

[***os.File**](*os.File.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganisation

> OutputOrganisation GetOrganisation(ctx, orgId).Execute()

Get an organization



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
	orgId := "orgId_example" // string | The organization ID preceded by `~`, or organization name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.GetOrganisation(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetOrganisation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganisation`: OutputOrganisation
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetOrganisation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization ID preceded by &#x60;~&#x60;, or organization name. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganisationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputOrganisation**](OutputOrganisation.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganisationAvatar

> *os.File GetOrganisationAvatar(ctx, orgId, fileHash).IfNoneMatch(ifNoneMatch).Execute()

Get the avatar of an organization



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
	orgId := "orgId_example" // string | The organization ID preceded by `~`, or organization name.
	fileHash := "fileHash_example" // string | The hash of the avatar file.
	ifNoneMatch := "ifNoneMatch_example" // string | ETag from a previous response. If it matches, returns `304 Not Modified` instead of the image content. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.GetOrganisationAvatar(context.Background(), orgId, fileHash).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetOrganisationAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganisationAvatar`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetOrganisationAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization ID preceded by &#x60;~&#x60;, or organization name. |
**fileHash** | **string** | The hash of the avatar file. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganisationAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifNoneMatch** | **string** | ETag from a previous response. If it matches, returns &#x60;304 Not Modified&#x60; instead of the image content. |

### Return type

[***os.File**](*os.File.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkOrganisations

> LinkOrganisations(ctx, orgId, otherOrgId).InputOrganisationLink(inputOrganisationLink).Execute()

Link two organizations



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
	orgId := "orgId_example" // string | The organization ID preceded by `~`, or organization name.
	otherOrgId := "otherOrgId_example" // string | The ID preceded by `~`, or name of the organization to link with.
	inputOrganisationLink := *openapiclient.NewInputOrganisationLink() // InputOrganisationLink |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.LinkOrganisations(context.Background(), orgId, otherOrgId).InputOrganisationLink(inputOrganisationLink).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.LinkOrganisations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization ID preceded by &#x60;~&#x60;, or organization name. |
**otherOrgId** | **string** | The ID preceded by &#x60;~&#x60;, or name of the organization to link with. |

### Other Parameters

Other parameters are passed through a pointer to a apiLinkOrganisationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inputOrganisationLink** | [**InputOrganisationLink**](InputOrganisationLink.md) |  |

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


## ListOrganisationLinks

> []OutputOrganisationLink ListOrganisationLinks(ctx, orgId).Execute()

List organization links



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
	orgId := "orgId_example" // string | The organization ID preceded by `~`, or organization name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.ListOrganisationLinks(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.ListOrganisationLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganisationLinks`: []OutputOrganisationLink
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.ListOrganisationLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization ID preceded by &#x60;~&#x60;, or organization name. |

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganisationLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OutputOrganisationLink**](OutputOrganisationLink.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSharingProfiles

> []OutputSharingProfile ListSharingProfiles(ctx).Execute()

List sharing profiles



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
	resp, r, err := apiClient.OrganizationAPI.ListSharingProfiles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.ListSharingProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSharingProfiles`: []OutputSharingProfile
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.ListSharingProfiles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSharingProfilesRequest struct via the builder pattern


### Return type

[**[]OutputSharingProfile**](OutputSharingProfile.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkOrganisations

> UnlinkOrganisations(ctx, orgId, otherOrgId).Execute()

Remove a link between two organizations



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
	orgId := "orgId_example" // string | The organization ID preceded by `~`, or organization name.
	otherOrgId := "otherOrgId_example" // string | The ID preceded by `~`, or name of the organization to unlink.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.UnlinkOrganisations(context.Background(), orgId, otherOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.UnlinkOrganisations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization ID preceded by &#x60;~&#x60;, or organization name. |
**otherOrgId** | **string** | The ID preceded by &#x60;~&#x60;, or name of the organization to unlink. |

### Other Parameters

Other parameters are passed through a pointer to a apiUnlinkOrganisationsRequest struct via the builder pattern


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


## UpdateOrganisation

> UpdateOrganisation(ctx, orgId).InputUpdateOrganisation(inputUpdateOrganisation).Execute()

Update an organization



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
	orgId := "orgId_example" // string | The organization ID preceded by `~`, or organization name.
	inputUpdateOrganisation := *openapiclient.NewInputUpdateOrganisation() // InputUpdateOrganisation |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.UpdateOrganisation(context.Background(), orgId).InputUpdateOrganisation(inputUpdateOrganisation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.UpdateOrganisation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization ID preceded by &#x60;~&#x60;, or organization name. |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganisationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputUpdateOrganisation** | [**InputUpdateOrganisation**](InputUpdateOrganisation.md) |  |

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
