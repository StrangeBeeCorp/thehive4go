# ListRespondersOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the responder once enabled within the organization. |
**Name** | **string** | Display name of the responder. |
**Version** | **string** | Version of the responder. |
**Description** | **string** | Description of what the responder does. |
**DataTypeList** | Pointer to **[]string** | Data types the responder supports. | [optional]
**CortexIds** | Pointer to **[]string** | Names of the Cortex servers on which the responder is available. | [optional]

## Methods

### NewListRespondersOutput

`func NewListRespondersOutput(id string, name string, version string, description string, ) *ListRespondersOutput`

NewListRespondersOutput instantiates a new ListRespondersOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRespondersOutputWithDefaults

`func NewListRespondersOutputWithDefaults() *ListRespondersOutput`

NewListRespondersOutputWithDefaults instantiates a new ListRespondersOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListRespondersOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListRespondersOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListRespondersOutput) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ListRespondersOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListRespondersOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListRespondersOutput) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *ListRespondersOutput) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListRespondersOutput) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListRespondersOutput) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetDescription

`func (o *ListRespondersOutput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListRespondersOutput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListRespondersOutput) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDataTypeList

`func (o *ListRespondersOutput) GetDataTypeList() []string`

GetDataTypeList returns the DataTypeList field if non-nil, zero value otherwise.

### GetDataTypeListOk

`func (o *ListRespondersOutput) GetDataTypeListOk() (*[]string, bool)`

GetDataTypeListOk returns a tuple with the DataTypeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypeList

`func (o *ListRespondersOutput) SetDataTypeList(v []string)`

SetDataTypeList sets DataTypeList field to given value.

### HasDataTypeList

`func (o *ListRespondersOutput) HasDataTypeList() bool`

HasDataTypeList returns a boolean if a field has been set.

### GetCortexIds

`func (o *ListRespondersOutput) GetCortexIds() []string`

GetCortexIds returns the CortexIds field if non-nil, zero value otherwise.

### GetCortexIdsOk

`func (o *ListRespondersOutput) GetCortexIdsOk() (*[]string, bool)`

GetCortexIdsOk returns a tuple with the CortexIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCortexIds

`func (o *ListRespondersOutput) SetCortexIds(v []string)`

SetCortexIds sets CortexIds field to given value.

### HasCortexIds

`func (o *ListRespondersOutput) HasCortexIds() bool`

HasCortexIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
