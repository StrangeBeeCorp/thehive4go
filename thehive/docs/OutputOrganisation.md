# OutputOrganisation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnderscoreId** | **string** |  | 
**UnderscoreType** | **string** |  | 
**UnderscoreCreatedBy** | **string** |  | 
**UnderscoreUpdatedBy** | Pointer to **string** |  | [optional] 
**UnderscoreCreatedAt** | **int64** |  | 
**UnderscoreUpdatedAt** | Pointer to **int64** |  | [optional] 
**Name** | **string** | This field can be used to reference an organisation instead of its _id | 
**Description** | **string** |  | 
**TaskRule** | **string** |  | 
**ObservableRule** | **string** |  | 
**Links** | Pointer to [**[]OrganisationLink**](OrganisationLink.md) |  | [optional] 
**Avatar** | Pointer to **string** |  | [optional] 
**Locked** | **bool** |  | 
**ExtraData** | **map[string]interface{}** |  | 

## Methods

### NewOutputOrganisation

`func NewOutputOrganisation(underscoreId string, underscoreType string, underscoreCreatedBy string, underscoreCreatedAt int64, name string, description string, taskRule string, observableRule string, locked bool, extraData map[string]interface{}, ) *OutputOrganisation`

NewOutputOrganisation instantiates a new OutputOrganisation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputOrganisationWithDefaults

`func NewOutputOrganisationWithDefaults() *OutputOrganisation`

NewOutputOrganisationWithDefaults instantiates a new OutputOrganisation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnderscoreId

`func (o *OutputOrganisation) GetUnderscoreId() string`

GetUnderscoreId returns the UnderscoreId field if non-nil, zero value otherwise.

### GetUnderscoreIdOk

`func (o *OutputOrganisation) GetUnderscoreIdOk() (*string, bool)`

GetUnderscoreIdOk returns a tuple with the UnderscoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreId

`func (o *OutputOrganisation) SetUnderscoreId(v string)`

SetUnderscoreId sets UnderscoreId field to given value.


### GetUnderscoreType

`func (o *OutputOrganisation) GetUnderscoreType() string`

GetUnderscoreType returns the UnderscoreType field if non-nil, zero value otherwise.

### GetUnderscoreTypeOk

`func (o *OutputOrganisation) GetUnderscoreTypeOk() (*string, bool)`

GetUnderscoreTypeOk returns a tuple with the UnderscoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreType

`func (o *OutputOrganisation) SetUnderscoreType(v string)`

SetUnderscoreType sets UnderscoreType field to given value.


### GetUnderscoreCreatedBy

`func (o *OutputOrganisation) GetUnderscoreCreatedBy() string`

GetUnderscoreCreatedBy returns the UnderscoreCreatedBy field if non-nil, zero value otherwise.

### GetUnderscoreCreatedByOk

`func (o *OutputOrganisation) GetUnderscoreCreatedByOk() (*string, bool)`

GetUnderscoreCreatedByOk returns a tuple with the UnderscoreCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedBy

`func (o *OutputOrganisation) SetUnderscoreCreatedBy(v string)`

SetUnderscoreCreatedBy sets UnderscoreCreatedBy field to given value.


### GetUnderscoreUpdatedBy

`func (o *OutputOrganisation) GetUnderscoreUpdatedBy() string`

GetUnderscoreUpdatedBy returns the UnderscoreUpdatedBy field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedByOk

`func (o *OutputOrganisation) GetUnderscoreUpdatedByOk() (*string, bool)`

GetUnderscoreUpdatedByOk returns a tuple with the UnderscoreUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedBy

`func (o *OutputOrganisation) SetUnderscoreUpdatedBy(v string)`

SetUnderscoreUpdatedBy sets UnderscoreUpdatedBy field to given value.

### HasUnderscoreUpdatedBy

`func (o *OutputOrganisation) HasUnderscoreUpdatedBy() bool`

HasUnderscoreUpdatedBy returns a boolean if a field has been set.

### GetUnderscoreCreatedAt

`func (o *OutputOrganisation) GetUnderscoreCreatedAt() int64`

GetUnderscoreCreatedAt returns the UnderscoreCreatedAt field if non-nil, zero value otherwise.

### GetUnderscoreCreatedAtOk

`func (o *OutputOrganisation) GetUnderscoreCreatedAtOk() (*int64, bool)`

GetUnderscoreCreatedAtOk returns a tuple with the UnderscoreCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedAt

`func (o *OutputOrganisation) SetUnderscoreCreatedAt(v int64)`

SetUnderscoreCreatedAt sets UnderscoreCreatedAt field to given value.


### GetUnderscoreUpdatedAt

`func (o *OutputOrganisation) GetUnderscoreUpdatedAt() int64`

GetUnderscoreUpdatedAt returns the UnderscoreUpdatedAt field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedAtOk

`func (o *OutputOrganisation) GetUnderscoreUpdatedAtOk() (*int64, bool)`

GetUnderscoreUpdatedAtOk returns a tuple with the UnderscoreUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedAt

`func (o *OutputOrganisation) SetUnderscoreUpdatedAt(v int64)`

SetUnderscoreUpdatedAt sets UnderscoreUpdatedAt field to given value.

### HasUnderscoreUpdatedAt

`func (o *OutputOrganisation) HasUnderscoreUpdatedAt() bool`

HasUnderscoreUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *OutputOrganisation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OutputOrganisation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OutputOrganisation) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *OutputOrganisation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OutputOrganisation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OutputOrganisation) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTaskRule

`func (o *OutputOrganisation) GetTaskRule() string`

GetTaskRule returns the TaskRule field if non-nil, zero value otherwise.

### GetTaskRuleOk

`func (o *OutputOrganisation) GetTaskRuleOk() (*string, bool)`

GetTaskRuleOk returns a tuple with the TaskRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRule

`func (o *OutputOrganisation) SetTaskRule(v string)`

SetTaskRule sets TaskRule field to given value.


### GetObservableRule

`func (o *OutputOrganisation) GetObservableRule() string`

GetObservableRule returns the ObservableRule field if non-nil, zero value otherwise.

### GetObservableRuleOk

`func (o *OutputOrganisation) GetObservableRuleOk() (*string, bool)`

GetObservableRuleOk returns a tuple with the ObservableRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservableRule

`func (o *OutputOrganisation) SetObservableRule(v string)`

SetObservableRule sets ObservableRule field to given value.


### GetLinks

`func (o *OutputOrganisation) GetLinks() []OrganisationLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OutputOrganisation) GetLinksOk() (*[]OrganisationLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OutputOrganisation) SetLinks(v []OrganisationLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OutputOrganisation) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAvatar

`func (o *OutputOrganisation) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *OutputOrganisation) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *OutputOrganisation) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *OutputOrganisation) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetLocked

`func (o *OutputOrganisation) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *OutputOrganisation) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *OutputOrganisation) SetLocked(v bool)`

SetLocked sets Locked field to given value.


### GetExtraData

`func (o *OutputOrganisation) GetExtraData() map[string]interface{}`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *OutputOrganisation) GetExtraDataOk() (*map[string]interface{}, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *OutputOrganisation) SetExtraData(v map[string]interface{})`

SetExtraData sets ExtraData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


