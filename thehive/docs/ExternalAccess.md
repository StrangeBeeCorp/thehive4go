# ExternalAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | **[]string** | Selected users&#39; logins. Must contain at least one external user or &#39;*&#39; (all external users of the organisation). | 
**Kind** | **string** |  | 

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


