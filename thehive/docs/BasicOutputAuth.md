# BasicOutputAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | Username for HTTP Basic authentication. |
**Password** | **string** | Password for HTTP Basic authentication. |
**Type** | **string** | Authentication type identifier. |

## Methods

### NewBasicOutputAuth

`func NewBasicOutputAuth(username string, password string, type_ string, ) *BasicOutputAuth`

NewBasicOutputAuth instantiates a new BasicOutputAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicOutputAuthWithDefaults

`func NewBasicOutputAuthWithDefaults() *BasicOutputAuth`

NewBasicOutputAuthWithDefaults instantiates a new BasicOutputAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *BasicOutputAuth) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *BasicOutputAuth) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *BasicOutputAuth) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *BasicOutputAuth) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *BasicOutputAuth) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *BasicOutputAuth) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetType

`func (o *BasicOutputAuth) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BasicOutputAuth) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BasicOutputAuth) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
