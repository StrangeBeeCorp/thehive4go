# Pages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Filter** | Pointer to [**Filter**](Filter.md) |  | [optional] 
**Kind** | [**Kind**](Kind.md) |  | 

## Methods

### NewPages

`func NewPages(kind Kind, ) *Pages`

NewPages instantiates a new Pages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagesWithDefaults

`func NewPagesWithDefaults() *Pages`

NewPagesWithDefaults instantiates a new Pages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *Pages) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Pages) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Pages) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Pages) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFilter

`func (o *Pages) GetFilter() Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *Pages) GetFilterOk() (*Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *Pages) SetFilter(v Filter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *Pages) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetKind

`func (o *Pages) GetKind() Kind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Pages) GetKindOk() (*Kind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Pages) SetKind(v Kind)`

SetKind sets Kind field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


