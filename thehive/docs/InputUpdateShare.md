# InputUpdateShare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | **string** | Profile that defines the permissions of the organization members on this case. Must match an existing profile. Use &#x60;POST /api/v1/query&#x60; with &#x60;listProfile&#x60; to retrieve available permission profiles. Overrides the [organization-level sharing rules](https://docs.strangebee.com/thehive/administration/organizations/about-organizations-sharing-rules/). |

## Methods

### NewInputUpdateShare

`func NewInputUpdateShare(profile string, ) *InputUpdateShare`

NewInputUpdateShare instantiates a new InputUpdateShare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputUpdateShareWithDefaults

`func NewInputUpdateShareWithDefaults() *InputUpdateShare`

NewInputUpdateShareWithDefaults instantiates a new InputUpdateShare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *InputUpdateShare) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *InputUpdateShare) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *InputUpdateShare) SetProfile(v string)`

SetProfile sets Profile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
