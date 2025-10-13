# OutputOrganisationProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganisationId** | **string** |  | 
**Organisation** | **string** |  | 
**Profile** | **string** |  | 
**Avatar** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**[]OrganisationLink**](OrganisationLink.md) |  | [optional] 

## Methods

### NewOutputOrganisationProfile

`func NewOutputOrganisationProfile(organisationId string, organisation string, profile string, ) *OutputOrganisationProfile`

NewOutputOrganisationProfile instantiates a new OutputOrganisationProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputOrganisationProfileWithDefaults

`func NewOutputOrganisationProfileWithDefaults() *OutputOrganisationProfile`

NewOutputOrganisationProfileWithDefaults instantiates a new OutputOrganisationProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganisationId

`func (o *OutputOrganisationProfile) GetOrganisationId() string`

GetOrganisationId returns the OrganisationId field if non-nil, zero value otherwise.

### GetOrganisationIdOk

`func (o *OutputOrganisationProfile) GetOrganisationIdOk() (*string, bool)`

GetOrganisationIdOk returns a tuple with the OrganisationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisationId

`func (o *OutputOrganisationProfile) SetOrganisationId(v string)`

SetOrganisationId sets OrganisationId field to given value.


### GetOrganisation

`func (o *OutputOrganisationProfile) GetOrganisation() string`

GetOrganisation returns the Organisation field if non-nil, zero value otherwise.

### GetOrganisationOk

`func (o *OutputOrganisationProfile) GetOrganisationOk() (*string, bool)`

GetOrganisationOk returns a tuple with the Organisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisation

`func (o *OutputOrganisationProfile) SetOrganisation(v string)`

SetOrganisation sets Organisation field to given value.


### GetProfile

`func (o *OutputOrganisationProfile) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *OutputOrganisationProfile) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *OutputOrganisationProfile) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetAvatar

`func (o *OutputOrganisationProfile) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *OutputOrganisationProfile) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *OutputOrganisationProfile) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *OutputOrganisationProfile) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetLinks

`func (o *OutputOrganisationProfile) GetLinks() []OrganisationLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OutputOrganisationProfile) GetLinksOk() (*[]OrganisationLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OutputOrganisationProfile) SetLinks(v []OrganisationLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OutputOrganisationProfile) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


