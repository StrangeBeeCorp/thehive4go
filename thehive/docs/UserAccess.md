# UserAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | **[]string** | Logins of the users who can access the case. Must include at least the case assignee and the user making the request. Must be a valid email address. Use &#x60;POST /api/v1/query&#x60; with &#x60;listUser&#x60; to retrieve available users and their email addresses. |
**Kind** | **string** | Only a specific set of internal users can see the case. Platinum or One license required. |

## Methods

### NewUserAccess

`func NewUserAccess(users []string, kind string, ) *UserAccess`

NewUserAccess instantiates a new UserAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAccessWithDefaults

`func NewUserAccessWithDefaults() *UserAccess`

NewUserAccessWithDefaults instantiates a new UserAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *UserAccess) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UserAccess) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UserAccess) SetUsers(v []string)`

SetUsers sets Users field to given value.


### GetKind

`func (o *UserAccess) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *UserAccess) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *UserAccess) SetKind(v string)`

SetKind sets Kind field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
