# OutputStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** |  | 
**GitDescription** | **string** |  | 
**Connectors** | **map[string]map[string]interface{}** |  | 
**Config** | **map[string]interface{}** |  | 
**License** | [**OutputLicenseStatus**](OutputLicenseStatus.md) |  | 
**Cluster** | Pointer to **map[string]interface{}** |  | [optional] 
**SchemaStatus** | Pointer to **[]map[string]interface{}** | Filled when verbose | [optional] 
**Features** | Pointer to **[]string** |  | [optional] 

## Methods

### NewOutputStatus

`func NewOutputStatus(version string, gitDescription string, connectors map[string]map[string]interface{}, config map[string]interface{}, license OutputLicenseStatus, ) *OutputStatus`

NewOutputStatus instantiates a new OutputStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputStatusWithDefaults

`func NewOutputStatusWithDefaults() *OutputStatus`

NewOutputStatusWithDefaults instantiates a new OutputStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *OutputStatus) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OutputStatus) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OutputStatus) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetGitDescription

`func (o *OutputStatus) GetGitDescription() string`

GetGitDescription returns the GitDescription field if non-nil, zero value otherwise.

### GetGitDescriptionOk

`func (o *OutputStatus) GetGitDescriptionOk() (*string, bool)`

GetGitDescriptionOk returns a tuple with the GitDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitDescription

`func (o *OutputStatus) SetGitDescription(v string)`

SetGitDescription sets GitDescription field to given value.


### GetConnectors

`func (o *OutputStatus) GetConnectors() map[string]map[string]interface{}`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *OutputStatus) GetConnectorsOk() (*map[string]map[string]interface{}, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *OutputStatus) SetConnectors(v map[string]map[string]interface{})`

SetConnectors sets Connectors field to given value.


### GetConfig

`func (o *OutputStatus) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *OutputStatus) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *OutputStatus) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.


### GetLicense

`func (o *OutputStatus) GetLicense() OutputLicenseStatus`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *OutputStatus) GetLicenseOk() (*OutputLicenseStatus, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *OutputStatus) SetLicense(v OutputLicenseStatus)`

SetLicense sets License field to given value.


### GetCluster

`func (o *OutputStatus) GetCluster() map[string]interface{}`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *OutputStatus) GetClusterOk() (*map[string]interface{}, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *OutputStatus) SetCluster(v map[string]interface{})`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *OutputStatus) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetSchemaStatus

`func (o *OutputStatus) GetSchemaStatus() []map[string]interface{}`

GetSchemaStatus returns the SchemaStatus field if non-nil, zero value otherwise.

### GetSchemaStatusOk

`func (o *OutputStatus) GetSchemaStatusOk() (*[]map[string]interface{}, bool)`

GetSchemaStatusOk returns a tuple with the SchemaStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaStatus

`func (o *OutputStatus) SetSchemaStatus(v []map[string]interface{})`

SetSchemaStatus sets SchemaStatus field to given value.

### HasSchemaStatus

`func (o *OutputStatus) HasSchemaStatus() bool`

HasSchemaStatus returns a boolean if a field has been set.

### GetFeatures

`func (o *OutputStatus) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *OutputStatus) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *OutputStatus) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *OutputStatus) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


