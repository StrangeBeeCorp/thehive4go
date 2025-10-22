# InputCustomField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**DisplayName** | Pointer to **string** |  | [optional] 
**Group** | **string** |  | 
**Description** | **string** |  | 
**Type** | [**CustomFieldType**](CustomFieldType.md) |  | 
**Mandatory** | Pointer to **bool** |  | [optional] 
**Options** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewInputCustomField

`func NewInputCustomField(name string, group string, description string, type_ CustomFieldType, ) *InputCustomField`

NewInputCustomField instantiates a new InputCustomField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputCustomFieldWithDefaults

`func NewInputCustomFieldWithDefaults() *InputCustomField`

NewInputCustomFieldWithDefaults instantiates a new InputCustomField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InputCustomField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputCustomField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputCustomField) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *InputCustomField) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InputCustomField) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InputCustomField) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InputCustomField) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetGroup

`func (o *InputCustomField) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InputCustomField) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InputCustomField) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetDescription

`func (o *InputCustomField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputCustomField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputCustomField) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *InputCustomField) GetType() CustomFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputCustomField) GetTypeOk() (*CustomFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputCustomField) SetType(v CustomFieldType)`

SetType sets Type field to given value.


### GetMandatory

`func (o *InputCustomField) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *InputCustomField) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *InputCustomField) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.

### HasMandatory

`func (o *InputCustomField) HasMandatory() bool`

HasMandatory returns a boolean if a field has been set.

### GetOptions

`func (o *InputCustomField) GetOptions() []interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *InputCustomField) GetOptionsOk() (*[]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *InputCustomField) SetOptions(v []interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *InputCustomField) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


