# CustomFieldsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Kind** | [**Kind**](Kind.md) |  | 

## Methods

### NewCustomFieldsList

`func NewCustomFieldsList(kind Kind, ) *CustomFieldsList`

NewCustomFieldsList instantiates a new CustomFieldsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldsListWithDefaults

`func NewCustomFieldsListWithDefaults() *CustomFieldsList`

NewCustomFieldsListWithDefaults instantiates a new CustomFieldsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CustomFieldsList) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CustomFieldsList) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CustomFieldsList) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CustomFieldsList) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetKind

`func (o *CustomFieldsList) GetKind() Kind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CustomFieldsList) GetKindOk() (*Kind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CustomFieldsList) SetKind(v Kind)`

SetKind sets Kind field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


