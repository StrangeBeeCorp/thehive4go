# InputUserOrganisation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organisation** | **string** |  | 
**Profile** | **string** |  | 
**Default** | Pointer to **bool** |  | [optional] 

## Methods

### NewInputUserOrganisation

`func NewInputUserOrganisation(organisation string, profile string, ) *InputUserOrganisation`

NewInputUserOrganisation instantiates a new InputUserOrganisation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputUserOrganisationWithDefaults

`func NewInputUserOrganisationWithDefaults() *InputUserOrganisation`

NewInputUserOrganisationWithDefaults instantiates a new InputUserOrganisation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganisation

`func (o *InputUserOrganisation) GetOrganisation() string`

GetOrganisation returns the Organisation field if non-nil, zero value otherwise.

### GetOrganisationOk

`func (o *InputUserOrganisation) GetOrganisationOk() (*string, bool)`

GetOrganisationOk returns a tuple with the Organisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisation

`func (o *InputUserOrganisation) SetOrganisation(v string)`

SetOrganisation sets Organisation field to given value.


### GetProfile

`func (o *InputUserOrganisation) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *InputUserOrganisation) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *InputUserOrganisation) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetDefault

`func (o *InputUserOrganisation) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *InputUserOrganisation) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *InputUserOrganisation) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *InputUserOrganisation) HasDefault() bool`

HasDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


