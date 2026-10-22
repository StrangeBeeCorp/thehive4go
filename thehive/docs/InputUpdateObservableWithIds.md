# InputUpdateObservableWithIds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** | Identifiers of the entities to target. IDs must be prefixed with &#x60;~&#x60;. Case numbers and names are passed as plain strings without a prefix. |
**DataType** | Pointer to **string** | New observable type. Must match an existing observable type defined in TheHive. Use &#x60;POST /api/v1/query&#x60; with &#x60;listObservableType&#x60; to list available types. | [optional]
**Message** | Pointer to **string** | New description of the observable. Supports [TheHive-flavored Markdown](https://docs.strangebee.com/thehive/user-guides/thehive-flavored-markdown/). Send &#x60;null&#x60; to unset. | [optional]
**Tlp** | Pointer to **int32** | New [Traffic Light Protocol (TLP)](https://www.misp-project.org/taxonomies.html#_tlp) level: 0 (Clear), 1 (Green), 2 (Amber), 3 (Amber+Strict), 4 (Red). | [optional]
**Pap** | Pointer to **int32** | New [Permissible Actions Protocol (PAP)](https://www.misp-project.org/taxonomies.html#_pap) level: 0 (Clear), 1 (Green), 2 (Amber), 3 (Red). | [optional]
**Tags** | Pointer to **[]string** | Replaces all current tags with this set. Unknown tags are automatically created. Use &#x60;addTags&#x60; and &#x60;removeTags&#x60; to add or remove individual tags without replacing the full set. Use &#x60;POST /api/v1/query&#x60; with &#x60;listTag&#x60; to retrieve existing custom tags, or with &#x60;listTaxonomy&#x60; and &#x60;tags&#x60; to retrieve taxonomy tags. | [optional]
**Ioc** | Pointer to **bool** | Marks the observable as an indicator of compromise (IOC), indicating it&#39;s linked to suspicious or malicious activity. | [optional]
**Sighted** | Pointer to **bool** | Marks the observable as sighted, meaning it has been detected or observed in the environment. | [optional]
**SightedAt** | Pointer to **int64** | Date and time when the observable was last sighted. Expected format: Unix timestamp in milliseconds. Send &#x60;null&#x60; to unset. | [optional]
**IgnoreSimilarity** | Pointer to **bool** | When &#x60;true&#x60;, excludes the observable from similarity checks used to identify related alerts and cases. | [optional]
**AddTags** | Pointer to **[]string** | Tags to add to the current set of observable tags. Unknown tags are automatically created. Use &#x60;POST /api/v1/query&#x60; with &#x60;listTag&#x60; to retrieve existing custom tags, or with &#x60;listTaxonomy&#x60; and &#x60;tags&#x60; to retrieve taxonomy tags. | [optional]
**RemoveTags** | Pointer to **[]string** | Tags to remove from the current set of observable tags. Tags not present on the observable are silently ignored. Use &#x60;GET /api/v1/observable/{observableId}&#x60; and read the &#x60;tags&#x60; field to see the current tags on the observable. | [optional]
**External** | Pointer to **bool** | When &#x60;true&#x60;, makes the observable accessible to external users with [TheHive Portal](https://docs.strangebee.com/thehive/administration/thehive-portal/about-thehive-portal/). Platinum or One license required. | [optional]

## Methods

### NewInputUpdateObservableWithIds

`func NewInputUpdateObservableWithIds(ids []string, ) *InputUpdateObservableWithIds`

NewInputUpdateObservableWithIds instantiates a new InputUpdateObservableWithIds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputUpdateObservableWithIdsWithDefaults

`func NewInputUpdateObservableWithIdsWithDefaults() *InputUpdateObservableWithIds`

NewInputUpdateObservableWithIdsWithDefaults instantiates a new InputUpdateObservableWithIds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *InputUpdateObservableWithIds) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *InputUpdateObservableWithIds) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *InputUpdateObservableWithIds) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetDataType

`func (o *InputUpdateObservableWithIds) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *InputUpdateObservableWithIds) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *InputUpdateObservableWithIds) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *InputUpdateObservableWithIds) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetMessage

`func (o *InputUpdateObservableWithIds) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InputUpdateObservableWithIds) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InputUpdateObservableWithIds) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InputUpdateObservableWithIds) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTlp

`func (o *InputUpdateObservableWithIds) GetTlp() int32`

GetTlp returns the Tlp field if non-nil, zero value otherwise.

### GetTlpOk

`func (o *InputUpdateObservableWithIds) GetTlpOk() (*int32, bool)`

GetTlpOk returns a tuple with the Tlp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlp

`func (o *InputUpdateObservableWithIds) SetTlp(v int32)`

SetTlp sets Tlp field to given value.

### HasTlp

`func (o *InputUpdateObservableWithIds) HasTlp() bool`

HasTlp returns a boolean if a field has been set.

### GetPap

`func (o *InputUpdateObservableWithIds) GetPap() int32`

GetPap returns the Pap field if non-nil, zero value otherwise.

### GetPapOk

`func (o *InputUpdateObservableWithIds) GetPapOk() (*int32, bool)`

GetPapOk returns a tuple with the Pap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPap

`func (o *InputUpdateObservableWithIds) SetPap(v int32)`

SetPap sets Pap field to given value.

### HasPap

`func (o *InputUpdateObservableWithIds) HasPap() bool`

HasPap returns a boolean if a field has been set.

### GetTags

`func (o *InputUpdateObservableWithIds) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InputUpdateObservableWithIds) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InputUpdateObservableWithIds) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InputUpdateObservableWithIds) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIoc

`func (o *InputUpdateObservableWithIds) GetIoc() bool`

GetIoc returns the Ioc field if non-nil, zero value otherwise.

### GetIocOk

`func (o *InputUpdateObservableWithIds) GetIocOk() (*bool, bool)`

GetIocOk returns a tuple with the Ioc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoc

`func (o *InputUpdateObservableWithIds) SetIoc(v bool)`

SetIoc sets Ioc field to given value.

### HasIoc

`func (o *InputUpdateObservableWithIds) HasIoc() bool`

HasIoc returns a boolean if a field has been set.

### GetSighted

`func (o *InputUpdateObservableWithIds) GetSighted() bool`

GetSighted returns the Sighted field if non-nil, zero value otherwise.

### GetSightedOk

`func (o *InputUpdateObservableWithIds) GetSightedOk() (*bool, bool)`

GetSightedOk returns a tuple with the Sighted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSighted

`func (o *InputUpdateObservableWithIds) SetSighted(v bool)`

SetSighted sets Sighted field to given value.

### HasSighted

`func (o *InputUpdateObservableWithIds) HasSighted() bool`

HasSighted returns a boolean if a field has been set.

### GetSightedAt

`func (o *InputUpdateObservableWithIds) GetSightedAt() int64`

GetSightedAt returns the SightedAt field if non-nil, zero value otherwise.

### GetSightedAtOk

`func (o *InputUpdateObservableWithIds) GetSightedAtOk() (*int64, bool)`

GetSightedAtOk returns a tuple with the SightedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSightedAt

`func (o *InputUpdateObservableWithIds) SetSightedAt(v int64)`

SetSightedAt sets SightedAt field to given value.

### HasSightedAt

`func (o *InputUpdateObservableWithIds) HasSightedAt() bool`

HasSightedAt returns a boolean if a field has been set.

### GetIgnoreSimilarity

`func (o *InputUpdateObservableWithIds) GetIgnoreSimilarity() bool`

GetIgnoreSimilarity returns the IgnoreSimilarity field if non-nil, zero value otherwise.

### GetIgnoreSimilarityOk

`func (o *InputUpdateObservableWithIds) GetIgnoreSimilarityOk() (*bool, bool)`

GetIgnoreSimilarityOk returns a tuple with the IgnoreSimilarity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSimilarity

`func (o *InputUpdateObservableWithIds) SetIgnoreSimilarity(v bool)`

SetIgnoreSimilarity sets IgnoreSimilarity field to given value.

### HasIgnoreSimilarity

`func (o *InputUpdateObservableWithIds) HasIgnoreSimilarity() bool`

HasIgnoreSimilarity returns a boolean if a field has been set.

### GetAddTags

`func (o *InputUpdateObservableWithIds) GetAddTags() []string`

GetAddTags returns the AddTags field if non-nil, zero value otherwise.

### GetAddTagsOk

`func (o *InputUpdateObservableWithIds) GetAddTagsOk() (*[]string, bool)`

GetAddTagsOk returns a tuple with the AddTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTags

`func (o *InputUpdateObservableWithIds) SetAddTags(v []string)`

SetAddTags sets AddTags field to given value.

### HasAddTags

`func (o *InputUpdateObservableWithIds) HasAddTags() bool`

HasAddTags returns a boolean if a field has been set.

### GetRemoveTags

`func (o *InputUpdateObservableWithIds) GetRemoveTags() []string`

GetRemoveTags returns the RemoveTags field if non-nil, zero value otherwise.

### GetRemoveTagsOk

`func (o *InputUpdateObservableWithIds) GetRemoveTagsOk() (*[]string, bool)`

GetRemoveTagsOk returns a tuple with the RemoveTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveTags

`func (o *InputUpdateObservableWithIds) SetRemoveTags(v []string)`

SetRemoveTags sets RemoveTags field to given value.

### HasRemoveTags

`func (o *InputUpdateObservableWithIds) HasRemoveTags() bool`

HasRemoveTags returns a boolean if a field has been set.

### GetExternal

`func (o *InputUpdateObservableWithIds) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *InputUpdateObservableWithIds) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *InputUpdateObservableWithIds) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *InputUpdateObservableWithIds) HasExternal() bool`

HasExternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
