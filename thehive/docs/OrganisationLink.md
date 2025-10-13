# OrganisationLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToOrganisation** | **string** | The _id of the entity or its &#39;name&#39; (depends of the entity) | 
**Avatar** | Pointer to **string** |  | [optional] 
**LinkType** | **string** |  | 
**OtherLinkType** | **string** |  | 

## Methods

### NewOrganisationLink

`func NewOrganisationLink(toOrganisation string, linkType string, otherLinkType string, ) *OrganisationLink`

NewOrganisationLink instantiates a new OrganisationLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganisationLinkWithDefaults

`func NewOrganisationLinkWithDefaults() *OrganisationLink`

NewOrganisationLinkWithDefaults instantiates a new OrganisationLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToOrganisation

`func (o *OrganisationLink) GetToOrganisation() string`

GetToOrganisation returns the ToOrganisation field if non-nil, zero value otherwise.

### GetToOrganisationOk

`func (o *OrganisationLink) GetToOrganisationOk() (*string, bool)`

GetToOrganisationOk returns a tuple with the ToOrganisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToOrganisation

`func (o *OrganisationLink) SetToOrganisation(v string)`

SetToOrganisation sets ToOrganisation field to given value.


### GetAvatar

`func (o *OrganisationLink) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *OrganisationLink) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *OrganisationLink) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *OrganisationLink) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetLinkType

`func (o *OrganisationLink) GetLinkType() string`

GetLinkType returns the LinkType field if non-nil, zero value otherwise.

### GetLinkTypeOk

`func (o *OrganisationLink) GetLinkTypeOk() (*string, bool)`

GetLinkTypeOk returns a tuple with the LinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkType

`func (o *OrganisationLink) SetLinkType(v string)`

SetLinkType sets LinkType field to given value.


### GetOtherLinkType

`func (o *OrganisationLink) GetOtherLinkType() string`

GetOtherLinkType returns the OtherLinkType field if non-nil, zero value otherwise.

### GetOtherLinkTypeOk

`func (o *OrganisationLink) GetOtherLinkTypeOk() (*string, bool)`

GetOtherLinkTypeOk returns a tuple with the OtherLinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherLinkType

`func (o *OrganisationLink) SetOtherLinkType(v string)`

SetOtherLinkType sets OtherLinkType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


