# InputAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponderId** | **string** | ID of the responder to run. Must match a responder available for the target case, observable, task, task log, or alert. Use &#x60;GET /api/v1/connector/cortex/responder/{entityType}/{entityId}&#x60; to list responders available for that entity. |
**CortexId** | Pointer to **string** | Name of the Cortex server to run the responder on, when the organization is connected to several servers. Check the &#x60;cortexIds&#x60; field returned by &#x60;GET /api/v1/connector/cortex/responder/{entityType}/{entityId}&#x60; for the servers where &#x60;responderId&#x60; is available. Omit to let TheHive pick the first server that offers the responder. | [optional]
**ObjectType** | **string** | Enum: &#x60;case&#x60; &#x60;case_artifact&#x60; &#x60;case_task&#x60; &#x60;case_task_log&#x60; &#x60;alert&#x60;&lt;br&gt;Type of entity the responder runs on: &#x60;case&#x60; for a case, &#x60;case_artifact&#x60; for an observable, &#x60;case_task&#x60; for a task, &#x60;case_task_log&#x60; for a task log, or &#x60;alert&#x60; for an alert. |
**ObjectId** | **string** | ID of the case, observable, task, task log, or alert identified by &#x60;objectType&#x60;. Accepts the entity&#39;s ID preceded by &#x60;~&#x60;, or another identifier accepted by that entity type, such as the case number for cases. Use &#x60;POST /api/v1/query&#x60; with &#x60;listCase&#x60;, &#x60;listObservable&#x60;, &#x60;listTask&#x60;, &#x60;listLog&#x60;, or &#x60;listAlert&#x60; to retrieve available IDs. |
**Parameters** | Pointer to **map[string]interface{}** | Additional parameters passed to the responder. Available keys depend on &#x60;responderId&#x60; and are defined by that responder&#39;s own configuration on the Cortex server. See [How to Create a Responder](https://docs.strangebee.com/cortex/api/how-to-create-a-responder/#input) for the expected format. | [optional]
**Tlp** | Pointer to **int32** | Enum: &#x60;0&#x60; &#x60;1&#x60; &#x60;2&#x60; &#x60;3&#x60; &#x60;4&#x60;&lt;br&gt;Numeric [Traffic Light Protocol (TLP)](https://www.misp-project.org/taxonomies.html#_tlp) level of the target entity: 0 (Clear), 1 (Green), 2 (Amber), 3 (Amber+Strict), 4 (Red). Checked against the responder&#39;s maximum TLP: the responder runs only if its own maximum TLP is greater than or equal to this value. | [optional]

## Methods

### NewInputAction

`func NewInputAction(responderId string, objectType string, objectId string, ) *InputAction`

NewInputAction instantiates a new InputAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputActionWithDefaults

`func NewInputActionWithDefaults() *InputAction`

NewInputActionWithDefaults instantiates a new InputAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponderId

`func (o *InputAction) GetResponderId() string`

GetResponderId returns the ResponderId field if non-nil, zero value otherwise.

### GetResponderIdOk

`func (o *InputAction) GetResponderIdOk() (*string, bool)`

GetResponderIdOk returns a tuple with the ResponderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponderId

`func (o *InputAction) SetResponderId(v string)`

SetResponderId sets ResponderId field to given value.


### GetCortexId

`func (o *InputAction) GetCortexId() string`

GetCortexId returns the CortexId field if non-nil, zero value otherwise.

### GetCortexIdOk

`func (o *InputAction) GetCortexIdOk() (*string, bool)`

GetCortexIdOk returns a tuple with the CortexId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCortexId

`func (o *InputAction) SetCortexId(v string)`

SetCortexId sets CortexId field to given value.

### HasCortexId

`func (o *InputAction) HasCortexId() bool`

HasCortexId returns a boolean if a field has been set.

### GetObjectType

`func (o *InputAction) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *InputAction) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *InputAction) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectId

`func (o *InputAction) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *InputAction) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *InputAction) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetParameters

`func (o *InputAction) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *InputAction) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *InputAction) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *InputAction) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetTlp

`func (o *InputAction) GetTlp() int32`

GetTlp returns the Tlp field if non-nil, zero value otherwise.

### GetTlpOk

`func (o *InputAction) GetTlpOk() (*int32, bool)`

GetTlpOk returns a tuple with the Tlp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlp

`func (o *InputAction) SetTlp(v int32)`

SetTlp sets Tlp field to given value.

### HasTlp

`func (o *InputAction) HasTlp() bool`

HasTlp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
