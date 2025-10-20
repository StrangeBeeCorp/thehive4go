# OutputShare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnderscoreId** | **string** |  |
**UnderscoreType** | **string** |  |
**UnderscoreCreatedBy** | **string** |  |
**UnderscoreUpdatedBy** | Pointer to **string** |  | [optional]
**UnderscoreCreatedAt** | **int64** |  |
**UnderscoreUpdatedAt** | Pointer to **int64** |  | [optional]
**CaseId** | **string** |  |
**ProfileName** | **string** |  |
**OrganisationName** | **string** |  |
**Owner** | **bool** |  |
**TaskRule** | **string** |  |
**ObservableRule** | **string** |  |

## Methods

### NewOutputShare

`func NewOutputShare(underscoreId string, underscoreType string, underscoreCreatedBy string, underscoreCreatedAt int64, caseId string, profileName string, organisationName string, owner bool, taskRule string, observableRule string, ) *OutputShare`

NewOutputShare instantiates a new OutputShare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputShareWithDefaults

`func NewOutputShareWithDefaults() *OutputShare`

NewOutputShareWithDefaults instantiates a new OutputShare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnderscoreId

`func (o *OutputShare) GetUnderscoreId() string`

GetUnderscoreId returns the UnderscoreId field if non-nil, zero value otherwise.

### GetUnderscoreIdOk

`func (o *OutputShare) GetUnderscoreIdOk() (*string, bool)`

GetUnderscoreIdOk returns a tuple with the UnderscoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreId

`func (o *OutputShare) SetUnderscoreId(v string)`

SetUnderscoreId sets UnderscoreId field to given value.


### GetUnderscoreType

`func (o *OutputShare) GetUnderscoreType() string`

GetUnderscoreType returns the UnderscoreType field if non-nil, zero value otherwise.

### GetUnderscoreTypeOk

`func (o *OutputShare) GetUnderscoreTypeOk() (*string, bool)`

GetUnderscoreTypeOk returns a tuple with the UnderscoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreType

`func (o *OutputShare) SetUnderscoreType(v string)`

SetUnderscoreType sets UnderscoreType field to given value.


### GetUnderscoreCreatedBy

`func (o *OutputShare) GetUnderscoreCreatedBy() string`

GetUnderscoreCreatedBy returns the UnderscoreCreatedBy field if non-nil, zero value otherwise.

### GetUnderscoreCreatedByOk

`func (o *OutputShare) GetUnderscoreCreatedByOk() (*string, bool)`

GetUnderscoreCreatedByOk returns a tuple with the UnderscoreCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedBy

`func (o *OutputShare) SetUnderscoreCreatedBy(v string)`

SetUnderscoreCreatedBy sets UnderscoreCreatedBy field to given value.


### GetUnderscoreUpdatedBy

`func (o *OutputShare) GetUnderscoreUpdatedBy() string`

GetUnderscoreUpdatedBy returns the UnderscoreUpdatedBy field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedByOk

`func (o *OutputShare) GetUnderscoreUpdatedByOk() (*string, bool)`

GetUnderscoreUpdatedByOk returns a tuple with the UnderscoreUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedBy

`func (o *OutputShare) SetUnderscoreUpdatedBy(v string)`

SetUnderscoreUpdatedBy sets UnderscoreUpdatedBy field to given value.

### HasUnderscoreUpdatedBy

`func (o *OutputShare) HasUnderscoreUpdatedBy() bool`

HasUnderscoreUpdatedBy returns a boolean if a field has been set.

### GetUnderscoreCreatedAt

`func (o *OutputShare) GetUnderscoreCreatedAt() int64`

GetUnderscoreCreatedAt returns the UnderscoreCreatedAt field if non-nil, zero value otherwise.

### GetUnderscoreCreatedAtOk

`func (o *OutputShare) GetUnderscoreCreatedAtOk() (*int64, bool)`

GetUnderscoreCreatedAtOk returns a tuple with the UnderscoreCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedAt

`func (o *OutputShare) SetUnderscoreCreatedAt(v int64)`

SetUnderscoreCreatedAt sets UnderscoreCreatedAt field to given value.


### GetUnderscoreUpdatedAt

`func (o *OutputShare) GetUnderscoreUpdatedAt() int64`

GetUnderscoreUpdatedAt returns the UnderscoreUpdatedAt field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedAtOk

`func (o *OutputShare) GetUnderscoreUpdatedAtOk() (*int64, bool)`

GetUnderscoreUpdatedAtOk returns a tuple with the UnderscoreUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedAt

`func (o *OutputShare) SetUnderscoreUpdatedAt(v int64)`

SetUnderscoreUpdatedAt sets UnderscoreUpdatedAt field to given value.

### HasUnderscoreUpdatedAt

`func (o *OutputShare) HasUnderscoreUpdatedAt() bool`

HasUnderscoreUpdatedAt returns a boolean if a field has been set.

### GetCaseId

`func (o *OutputShare) GetCaseId() string`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *OutputShare) GetCaseIdOk() (*string, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *OutputShare) SetCaseId(v string)`

SetCaseId sets CaseId field to given value.


### GetProfileName

`func (o *OutputShare) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *OutputShare) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *OutputShare) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.


### GetOrganisationName

`func (o *OutputShare) GetOrganisationName() string`

GetOrganisationName returns the OrganisationName field if non-nil, zero value otherwise.

### GetOrganisationNameOk

`func (o *OutputShare) GetOrganisationNameOk() (*string, bool)`

GetOrganisationNameOk returns a tuple with the OrganisationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisationName

`func (o *OutputShare) SetOrganisationName(v string)`

SetOrganisationName sets OrganisationName field to given value.


### GetOwner

`func (o *OutputShare) GetOwner() bool`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *OutputShare) GetOwnerOk() (*bool, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *OutputShare) SetOwner(v bool)`

SetOwner sets Owner field to given value.


### GetTaskRule

`func (o *OutputShare) GetTaskRule() string`

GetTaskRule returns the TaskRule field if non-nil, zero value otherwise.

### GetTaskRuleOk

`func (o *OutputShare) GetTaskRuleOk() (*string, bool)`

GetTaskRuleOk returns a tuple with the TaskRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRule

`func (o *OutputShare) SetTaskRule(v string)`

SetTaskRule sets TaskRule field to given value.


### GetObservableRule

`func (o *OutputShare) GetObservableRule() string`

GetObservableRule returns the ObservableRule field if non-nil, zero value otherwise.

### GetObservableRuleOk

`func (o *OutputShare) GetObservableRuleOk() (*string, bool)`

GetObservableRuleOk returns a tuple with the ObservableRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservableRule

`func (o *OutputShare) SetObservableRule(v string)`

SetObservableRule sets ObservableRule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
