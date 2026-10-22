# ExternalAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | **[]string** | Logins of the external users who can access the case. Must be a valid email address. Use &#x60;POST /api/v1/query&#x60; with &#x60;listUser&#x60; to retrieve available users and their email addresses. |
**Kind** | **string** | Only a specific set of external users can access the case through [TheHive Portal](https://docs.strangebee.com/thehive/administration/thehive-portal/about-thehive-portal/). Internal users can still access the case through the main TheHive interface. Platinum or One license required. |

## Methods

### NewExternalAccess

`func NewExternalAccess(users []string, kind string, ) *ExternalAccess`

NewExternalAccess instantiates a new ExternalAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalAccessWithDefaults

`func NewExternalAccessWithDefaults() *ExternalAccess`

NewExternalAccessWithDefaults instantiates a new ExternalAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *ExternalAccess) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ExternalAccess) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ExternalAccess) SetUsers(v []string)`

SetUsers sets Users field to given value.


### GetKind

`func (o *ExternalAccess) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ExternalAccess) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ExternalAccess) SetKind(v string)`

SetKind sets Kind field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
