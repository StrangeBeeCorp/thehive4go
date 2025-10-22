# OutputContextDocumentationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Kind** | [**OutputContextDocumentationItemKind**](OutputContextDocumentationItemKind.md) |  | 
**Args** | Pointer to **[]string** |  | [optional] 

## Methods

### NewOutputContextDocumentationItem

`func NewOutputContextDocumentationItem(name string, kind OutputContextDocumentationItemKind, ) *OutputContextDocumentationItem`

NewOutputContextDocumentationItem instantiates a new OutputContextDocumentationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputContextDocumentationItemWithDefaults

`func NewOutputContextDocumentationItemWithDefaults() *OutputContextDocumentationItem`

NewOutputContextDocumentationItemWithDefaults instantiates a new OutputContextDocumentationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OutputContextDocumentationItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OutputContextDocumentationItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OutputContextDocumentationItem) SetName(v string)`

SetName sets Name field to given value.


### GetKind

`func (o *OutputContextDocumentationItem) GetKind() OutputContextDocumentationItemKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *OutputContextDocumentationItem) GetKindOk() (*OutputContextDocumentationItemKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *OutputContextDocumentationItem) SetKind(v OutputContextDocumentationItemKind)`

SetKind sets Kind field to given value.


### GetArgs

`func (o *OutputContextDocumentationItem) GetArgs() []string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *OutputContextDocumentationItem) GetArgsOk() (*[]string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *OutputContextDocumentationItem) SetArgs(v []string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *OutputContextDocumentationItem) HasArgs() bool`

HasArgs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


