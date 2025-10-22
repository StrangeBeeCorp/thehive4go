# InputCreatePage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Content** | **string** |  | 
**Order** | Pointer to **int32** |  | [optional] 
**Category** | **string** |  | 

## Methods

### NewInputCreatePage

`func NewInputCreatePage(title string, content string, category string, ) *InputCreatePage`

NewInputCreatePage instantiates a new InputCreatePage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputCreatePageWithDefaults

`func NewInputCreatePageWithDefaults() *InputCreatePage`

NewInputCreatePageWithDefaults instantiates a new InputCreatePage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *InputCreatePage) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InputCreatePage) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InputCreatePage) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetContent

`func (o *InputCreatePage) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *InputCreatePage) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *InputCreatePage) SetContent(v string)`

SetContent sets Content field to given value.


### GetOrder

`func (o *InputCreatePage) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *InputCreatePage) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *InputCreatePage) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *InputCreatePage) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetCategory

`func (o *InputCreatePage) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InputCreatePage) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InputCreatePage) SetCategory(v string)`

SetCategory sets Category field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


