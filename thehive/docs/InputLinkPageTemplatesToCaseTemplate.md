# InputLinkPageTemplatesToCaseTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageTemplateIds** | Pointer to **[]string** | IDs of the page templates to link to the case template. Must match existing page templates. Use &#x60;POST /api/v1/query&#x60; with &#x60;listPageTemplate&#x60; to retrieve available page templates and their IDs. Pass an empty array to remove all links. | [optional]

## Methods

### NewInputLinkPageTemplatesToCaseTemplate

`func NewInputLinkPageTemplatesToCaseTemplate() *InputLinkPageTemplatesToCaseTemplate`

NewInputLinkPageTemplatesToCaseTemplate instantiates a new InputLinkPageTemplatesToCaseTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputLinkPageTemplatesToCaseTemplateWithDefaults

`func NewInputLinkPageTemplatesToCaseTemplateWithDefaults() *InputLinkPageTemplatesToCaseTemplate`

NewInputLinkPageTemplatesToCaseTemplateWithDefaults instantiates a new InputLinkPageTemplatesToCaseTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageTemplateIds

`func (o *InputLinkPageTemplatesToCaseTemplate) GetPageTemplateIds() []string`

GetPageTemplateIds returns the PageTemplateIds field if non-nil, zero value otherwise.

### GetPageTemplateIdsOk

`func (o *InputLinkPageTemplatesToCaseTemplate) GetPageTemplateIdsOk() (*[]string, bool)`

GetPageTemplateIdsOk returns a tuple with the PageTemplateIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageTemplateIds

`func (o *InputLinkPageTemplatesToCaseTemplate) SetPageTemplateIds(v []string)`

SetPageTemplateIds sets PageTemplateIds field to given value.

### HasPageTemplateIds

`func (o *InputLinkPageTemplatesToCaseTemplate) HasPageTemplateIds() bool`

HasPageTemplateIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
