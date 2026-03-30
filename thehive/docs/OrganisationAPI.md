# \OrganisationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAttachmentToOrganisation**](OrganisationAPI.md#AddAttachmentToOrganisation) | **Post** /api/v1/attachment | 
[**BulkLinkOrganisations**](OrganisationAPI.md#BulkLinkOrganisations) | **Put** /api/v1/organisation/{orgId}/links | 
[**CreateOrganisation**](OrganisationAPI.md#CreateOrganisation) | **Post** /api/v1/organisation | 
[**DeleteAttachment**](OrganisationAPI.md#DeleteAttachment) | **Delete** /api/v1/attachment/{attachmentId} | 
[**DownloadAttachment**](OrganisationAPI.md#DownloadAttachment) | **Get** /api/v1/attachment/{attachmentId}/download | 
[**GetAttachment**](OrganisationAPI.md#GetAttachment) | **Get** /api/v1/attachment/{attachmentId} | 
[**GetOrganisation**](OrganisationAPI.md#GetOrganisation) | **Get** /api/v1/organisation/{orgId} | 
[**GetOrganisationAvatar**](OrganisationAPI.md#GetOrganisationAvatar) | **Get** /api/v1/organisation/{orgId}/avatar/{fileHash} | 
[**LinkOrganisations**](OrganisationAPI.md#LinkOrganisations) | **Put** /api/v1/organisation/{orgId}/link/{otherOrgId} | 
[**ListOrganisationLinks**](OrganisationAPI.md#ListOrganisationLinks) | **Get** /api/v1/organisation/{orgId}/links | 
[**ListSharingProfiles**](OrganisationAPI.md#ListSharingProfiles) | **Get** /api/v1/sharingProfile | 
[**UnlinkOrganisations**](OrganisationAPI.md#UnlinkOrganisations) | **Delete** /api/v1/organisation/{orgId}/link/{otherOrgId} | 
[**UpdateOrganisation**](OrganisationAPI.md#UpdateOrganisation) | **Patch** /api/v1/organisation/{orgId} | 



## AddAttachmentToOrganisation

> OutputAttachments AddAttachmentToOrganisation(ctx).Attachments(attachments).CanRename(canRename).Execute()



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
	attachments := []*os.File{"TODO"} // []*os.File | 
	canRename := true // bool | If set to `true`, the files can be renamed if they already exist with the same name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganisationAPI.AddAttachmentToOrganisation(context.Background()).Attachments(attachments).CanRename(canRename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganisationAPI.AddAttachmentToOrganisation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddAttachmentToOrganisation`: OutputAttachments
	fmt.Fprintf(os.Stdout, "Response from `OrganisationAPI.AddAttachmentToOrganisation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAttachmentToOrganisationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attachments** | **[]*os.File** |  | 
 **canRename** | **bool** | If set to &#x60;true&#x60;, the files can be renamed if they already exist with the same name | 

### Return type

[**OutputAttachments**](OutputAttachments.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkLinkOrganisations

> BulkLinkOrganisations(ctx, orgId).InputOrganisationBulkLink(inputOrganisationBulkLink).Execute()



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
	orgId := "orgId_example" // string | 
	inputOrganisationBulkLink := *openapiclient.NewInputOrganisationBulkLink() // InputOrganisationBulkLink | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganisationAPI.BulkLinkOrganisations(context.Background(), orgId).InputOrganisationBulkLink(inputOrganisationBulkLink).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganisationAPI.BulkLinkOrganisations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkLinkOrganisationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputOrganisationBulkLink** | [**InputOrganisationBulkLink**](InputOrganisationBulkLink.md) |  | 

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


## CreateOrganisation

> OutputOrganisation CreateOrganisation(ctx).InputCreateOrganisation(inputCreateOrganisation).Execute()



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
	resp, r, err := apiClient.OrganisationAPI.CreateOrganisation(context.Background()).InputCreateOrganisation(inputCreateOrganisation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganisationAPI.CreateOrganisation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganisation`: OutputOrganisation
	fmt.Fprintf(os.Stdout, "Response from `OrganisationAPI.CreateOrganisation`: %v\n", resp)
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

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAttachment

> DeleteAttachment(ctx, attachmentId).Execute()



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
	attachmentId := "attachmentId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganisationAPI.DeleteAttachment(context.Background(), attachmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganisationAPI.DeleteAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAttachmentRequest struct via the builder pattern


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


## DownloadAttachment

> *os.File DownloadAttachment(ctx, attachmentId).Execute()



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
	attachmentId := "attachmentId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganisationAPI.DownloadAttachment(context.Background(), attachmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganisationAPI.DownloadAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadAttachment`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `OrganisationAPI.DownloadAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetAttachment

> *os.File GetAttachment(ctx, attachmentId).IfNoneMatch(ifNoneMatch).Execute()



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
	attachmentId := "attachmentId_example" // string | 
	ifNoneMatch := "ifNoneMatch_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganisationAPI.GetAttachment(context.Background(), attachmentId).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganisationAPI.GetAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachment`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `OrganisationAPI.GetAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifNoneMatch** | **string** |  | 

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


## GetOrganisation

> OutputOrganisation GetOrganisation(ctx, orgId).Execute()



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
	orgId := "orgId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganisationAPI.GetOrganisation(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganisationAPI.GetOrganisation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganisation`: OutputOrganisation
	fmt.Fprintf(os.Stdout, "Response from `OrganisationAPI.GetOrganisation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganisationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputOrganisation**](OutputOrganisation.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganisationAvatar

> *os.File GetOrganisationAvatar(ctx, orgId, fileHash).IfNoneMatch(ifNoneMatch).Execute()



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
	orgId := "orgId_example" // string | 
	fileHash := "fileHash_example" // string | 
	ifNoneMatch := "ifNoneMatch_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganisationAPI.GetOrganisationAvatar(context.Background(), orgId, fileHash).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganisationAPI.GetOrganisationAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganisationAvatar`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `OrganisationAPI.GetOrganisationAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**fileHash** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganisationAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifNoneMatch** | **string** |  | 

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


## LinkOrganisations

> LinkOrganisations(ctx, orgId, otherOrgId).InputOrganisationLink(inputOrganisationLink).Execute()



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
	orgId := "orgId_example" // string | 
	otherOrgId := "otherOrgId_example" // string | 
	inputOrganisationLink := *openapiclient.NewInputOrganisationLink() // InputOrganisationLink |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganisationAPI.LinkOrganisations(context.Background(), orgId, otherOrgId).InputOrganisationLink(inputOrganisationLink).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganisationAPI.LinkOrganisations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**otherOrgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLinkOrganisationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inputOrganisationLink** | [**InputOrganisationLink**](InputOrganisationLink.md) |  | 

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


## ListOrganisationLinks

> []OutputOrganisationLink ListOrganisationLinks(ctx, orgId).Execute()



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
	orgId := "orgId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganisationAPI.ListOrganisationLinks(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganisationAPI.ListOrganisationLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganisationLinks`: []OutputOrganisationLink
	fmt.Fprintf(os.Stdout, "Response from `OrganisationAPI.ListOrganisationLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganisationLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OutputOrganisationLink**](OutputOrganisationLink.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSharingProfiles

> []OutputSharingProfile ListSharingProfiles(ctx).Execute()



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
	resp, r, err := apiClient.OrganisationAPI.ListSharingProfiles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganisationAPI.ListSharingProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSharingProfiles`: []OutputSharingProfile
	fmt.Fprintf(os.Stdout, "Response from `OrganisationAPI.ListSharingProfiles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSharingProfilesRequest struct via the builder pattern


### Return type

[**[]OutputSharingProfile**](OutputSharingProfile.md)

### Authorization

[Basic](../README.md#Basic), [ApiKey](../README.md#ApiKey), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkOrganisations

> UnlinkOrganisations(ctx, orgId, otherOrgId).Execute()



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
	orgId := "orgId_example" // string | 
	otherOrgId := "otherOrgId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganisationAPI.UnlinkOrganisations(context.Background(), orgId, otherOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganisationAPI.UnlinkOrganisations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**otherOrgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlinkOrganisationsRequest struct via the builder pattern


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


## UpdateOrganisation

> UpdateOrganisation(ctx, orgId).InputUpdateOrganisation(inputUpdateOrganisation).Execute()



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
	orgId := "orgId_example" // string | 
	inputUpdateOrganisation := *openapiclient.NewInputUpdateOrganisation() // InputUpdateOrganisation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganisationAPI.UpdateOrganisation(context.Background(), orgId).InputUpdateOrganisation(inputUpdateOrganisation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganisationAPI.UpdateOrganisation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganisationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inputUpdateOrganisation** | [**InputUpdateOrganisation**](InputUpdateOrganisation.md) |  | 

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

