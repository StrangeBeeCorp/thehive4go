# InputUpdateCustomField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**CustomFieldType**](CustomFieldType.md) |  | [optional] 
**Options** | Pointer to **[]interface{}** |  | [optional] 
**Mandatory** | Pointer to **bool** |  | [optional] 

## Methods

### NewInputUpdateCustomField

`func NewInputUpdateCustomField() *InputUpdateCustomField`

NewInputUpdateCustomField instantiates a new InputUpdateCustomField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputUpdateCustomFieldWithDefaults

`func NewInputUpdateCustomFieldWithDefaults() *InputUpdateCustomField`

NewInputUpdateCustomFieldWithDefaults instantiates a new InputUpdateCustomField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *InputUpdateCustomField) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InputUpdateCustomField) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InputUpdateCustomField) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InputUpdateCustomField) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetGroup

`func (o *InputUpdateCustomField) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InputUpdateCustomField) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InputUpdateCustomField) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *InputUpdateCustomField) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetDescription

`func (o *InputUpdateCustomField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputUpdateCustomField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputUpdateCustomField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InputUpdateCustomField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *InputUpdateCustomField) GetType() CustomFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputUpdateCustomField) GetTypeOk() (*CustomFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputUpdateCustomField) SetType(v CustomFieldType)`

SetType sets Type field to given value.

### HasType

`func (o *InputUpdateCustomField) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOptions

`func (o *InputUpdateCustomField) GetOptions() []interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *InputUpdateCustomField) GetOptionsOk() (*[]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *InputUpdateCustomField) SetOptions(v []interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *InputUpdateCustomField) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetMandatory

`func (o *InputUpdateCustomField) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *InputUpdateCustomField) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *InputUpdateCustomField) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.

### HasMandatory

`func (o *InputUpdateCustomField) HasMandatory() bool`

HasMandatory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


