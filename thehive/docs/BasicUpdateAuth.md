# BasicUpdateAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | Username for HTTP Basic authentication. |
**Password** | **string** | Password for HTTP Basic authentication. |
**Type** | **string** | Authentication type identifier. Must be &#x60;basic&#x60;. |

## Methods

### NewBasicUpdateAuth

`func NewBasicUpdateAuth(username string, password string, type_ string, ) *BasicUpdateAuth`

NewBasicUpdateAuth instantiates a new BasicUpdateAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicUpdateAuthWithDefaults

`func NewBasicUpdateAuthWithDefaults() *BasicUpdateAuth`

NewBasicUpdateAuthWithDefaults instantiates a new BasicUpdateAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *BasicUpdateAuth) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *BasicUpdateAuth) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *BasicUpdateAuth) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *BasicUpdateAuth) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *BasicUpdateAuth) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *BasicUpdateAuth) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetType

`func (o *BasicUpdateAuth) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BasicUpdateAuth) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BasicUpdateAuth) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
