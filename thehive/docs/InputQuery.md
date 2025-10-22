# InputQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to [**[]InputQueryNamedOperation**](InputQueryNamedOperation.md) |  | [optional] 
**ExcludeFields** | Pointer to **[]string** |  | [optional] 

## Methods

### NewInputQuery

`func NewInputQuery() *InputQuery`

NewInputQuery instantiates a new InputQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputQueryWithDefaults

`func NewInputQueryWithDefaults() *InputQuery`

NewInputQueryWithDefaults instantiates a new InputQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *InputQuery) GetQuery() []InputQueryNamedOperation`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *InputQuery) GetQueryOk() (*[]InputQueryNamedOperation, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *InputQuery) SetQuery(v []InputQueryNamedOperation)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *InputQuery) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetExcludeFields

`func (o *InputQuery) GetExcludeFields() []string`

GetExcludeFields returns the ExcludeFields field if non-nil, zero value otherwise.

### GetExcludeFieldsOk

`func (o *InputQuery) GetExcludeFieldsOk() (*[]string, bool)`

GetExcludeFieldsOk returns a tuple with the ExcludeFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFields

`func (o *InputQuery) SetExcludeFields(v []string)`

SetExcludeFields sets ExcludeFields field to given value.

### HasExcludeFields

`func (o *InputQuery) HasExcludeFields() bool`

HasExcludeFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


