# OutputPattern

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnderscoreId** | **string** |  | 
**UnderscoreType** | **string** |  | 
**UnderscoreCreatedBy** | **string** |  | 
**UnderscoreUpdatedBy** | Pointer to **string** |  | [optional] 
**UnderscoreCreatedAt** | **int32** |  | 
**UnderscoreUpdatedAt** | Pointer to **int32** |  | [optional] 
**PatternId** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Tactics** | Pointer to **[]string** |  | [optional] 
**Url** | **string** |  | 
**PatternType** | **string** |  | 
**CapecId** | Pointer to **string** |  | [optional] 
**CapecUrl** | Pointer to **string** |  | [optional] 
**Revoked** | **bool** |  | 
**DataSources** | Pointer to **[]string** |  | [optional] 
**DefenseBypassed** | Pointer to **[]string** |  | [optional] 
**Detection** | Pointer to **string** |  | [optional] 
**PermissionsRequired** | Pointer to **[]string** |  | [optional] 
**Platforms** | Pointer to **[]string** |  | [optional] 
**RemoteSupport** | **bool** |  | 
**SystemRequirements** | Pointer to **[]string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**ExtraData** | **map[string]interface{}** |  | 

## Methods

### NewOutputPattern

`func NewOutputPattern(underscoreId string, underscoreType string, underscoreCreatedBy string, underscoreCreatedAt int32, patternId string, name string, url string, patternType string, revoked bool, remoteSupport bool, extraData map[string]interface{}, ) *OutputPattern`

NewOutputPattern instantiates a new OutputPattern object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputPatternWithDefaults

`func NewOutputPatternWithDefaults() *OutputPattern`

NewOutputPatternWithDefaults instantiates a new OutputPattern object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnderscoreId

`func (o *OutputPattern) GetUnderscoreId() string`

GetUnderscoreId returns the UnderscoreId field if non-nil, zero value otherwise.

### GetUnderscoreIdOk

`func (o *OutputPattern) GetUnderscoreIdOk() (*string, bool)`

GetUnderscoreIdOk returns a tuple with the UnderscoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreId

`func (o *OutputPattern) SetUnderscoreId(v string)`

SetUnderscoreId sets UnderscoreId field to given value.


### GetUnderscoreType

`func (o *OutputPattern) GetUnderscoreType() string`

GetUnderscoreType returns the UnderscoreType field if non-nil, zero value otherwise.

### GetUnderscoreTypeOk

`func (o *OutputPattern) GetUnderscoreTypeOk() (*string, bool)`

GetUnderscoreTypeOk returns a tuple with the UnderscoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreType

`func (o *OutputPattern) SetUnderscoreType(v string)`

SetUnderscoreType sets UnderscoreType field to given value.


### GetUnderscoreCreatedBy

`func (o *OutputPattern) GetUnderscoreCreatedBy() string`

GetUnderscoreCreatedBy returns the UnderscoreCreatedBy field if non-nil, zero value otherwise.

### GetUnderscoreCreatedByOk

`func (o *OutputPattern) GetUnderscoreCreatedByOk() (*string, bool)`

GetUnderscoreCreatedByOk returns a tuple with the UnderscoreCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedBy

`func (o *OutputPattern) SetUnderscoreCreatedBy(v string)`

SetUnderscoreCreatedBy sets UnderscoreCreatedBy field to given value.


### GetUnderscoreUpdatedBy

`func (o *OutputPattern) GetUnderscoreUpdatedBy() string`

GetUnderscoreUpdatedBy returns the UnderscoreUpdatedBy field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedByOk

`func (o *OutputPattern) GetUnderscoreUpdatedByOk() (*string, bool)`

GetUnderscoreUpdatedByOk returns a tuple with the UnderscoreUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedBy

`func (o *OutputPattern) SetUnderscoreUpdatedBy(v string)`

SetUnderscoreUpdatedBy sets UnderscoreUpdatedBy field to given value.

### HasUnderscoreUpdatedBy

`func (o *OutputPattern) HasUnderscoreUpdatedBy() bool`

HasUnderscoreUpdatedBy returns a boolean if a field has been set.

### GetUnderscoreCreatedAt

`func (o *OutputPattern) GetUnderscoreCreatedAt() int32`

GetUnderscoreCreatedAt returns the UnderscoreCreatedAt field if non-nil, zero value otherwise.

### GetUnderscoreCreatedAtOk

`func (o *OutputPattern) GetUnderscoreCreatedAtOk() (*int32, bool)`

GetUnderscoreCreatedAtOk returns a tuple with the UnderscoreCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreCreatedAt

`func (o *OutputPattern) SetUnderscoreCreatedAt(v int32)`

SetUnderscoreCreatedAt sets UnderscoreCreatedAt field to given value.


### GetUnderscoreUpdatedAt

`func (o *OutputPattern) GetUnderscoreUpdatedAt() int32`

GetUnderscoreUpdatedAt returns the UnderscoreUpdatedAt field if non-nil, zero value otherwise.

### GetUnderscoreUpdatedAtOk

`func (o *OutputPattern) GetUnderscoreUpdatedAtOk() (*int32, bool)`

GetUnderscoreUpdatedAtOk returns a tuple with the UnderscoreUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderscoreUpdatedAt

`func (o *OutputPattern) SetUnderscoreUpdatedAt(v int32)`

SetUnderscoreUpdatedAt sets UnderscoreUpdatedAt field to given value.

### HasUnderscoreUpdatedAt

`func (o *OutputPattern) HasUnderscoreUpdatedAt() bool`

HasUnderscoreUpdatedAt returns a boolean if a field has been set.

### GetPatternId

`func (o *OutputPattern) GetPatternId() string`

GetPatternId returns the PatternId field if non-nil, zero value otherwise.

### GetPatternIdOk

`func (o *OutputPattern) GetPatternIdOk() (*string, bool)`

GetPatternIdOk returns a tuple with the PatternId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternId

`func (o *OutputPattern) SetPatternId(v string)`

SetPatternId sets PatternId field to given value.


### GetName

`func (o *OutputPattern) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OutputPattern) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OutputPattern) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *OutputPattern) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OutputPattern) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OutputPattern) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OutputPattern) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTactics

`func (o *OutputPattern) GetTactics() []string`

GetTactics returns the Tactics field if non-nil, zero value otherwise.

### GetTacticsOk

`func (o *OutputPattern) GetTacticsOk() (*[]string, bool)`

GetTacticsOk returns a tuple with the Tactics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTactics

`func (o *OutputPattern) SetTactics(v []string)`

SetTactics sets Tactics field to given value.

### HasTactics

`func (o *OutputPattern) HasTactics() bool`

HasTactics returns a boolean if a field has been set.

### GetUrl

`func (o *OutputPattern) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OutputPattern) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OutputPattern) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetPatternType

`func (o *OutputPattern) GetPatternType() string`

GetPatternType returns the PatternType field if non-nil, zero value otherwise.

### GetPatternTypeOk

`func (o *OutputPattern) GetPatternTypeOk() (*string, bool)`

GetPatternTypeOk returns a tuple with the PatternType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternType

`func (o *OutputPattern) SetPatternType(v string)`

SetPatternType sets PatternType field to given value.


### GetCapecId

`func (o *OutputPattern) GetCapecId() string`

GetCapecId returns the CapecId field if non-nil, zero value otherwise.

### GetCapecIdOk

`func (o *OutputPattern) GetCapecIdOk() (*string, bool)`

GetCapecIdOk returns a tuple with the CapecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapecId

`func (o *OutputPattern) SetCapecId(v string)`

SetCapecId sets CapecId field to given value.

### HasCapecId

`func (o *OutputPattern) HasCapecId() bool`

HasCapecId returns a boolean if a field has been set.

### GetCapecUrl

`func (o *OutputPattern) GetCapecUrl() string`

GetCapecUrl returns the CapecUrl field if non-nil, zero value otherwise.

### GetCapecUrlOk

`func (o *OutputPattern) GetCapecUrlOk() (*string, bool)`

GetCapecUrlOk returns a tuple with the CapecUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapecUrl

`func (o *OutputPattern) SetCapecUrl(v string)`

SetCapecUrl sets CapecUrl field to given value.

### HasCapecUrl

`func (o *OutputPattern) HasCapecUrl() bool`

HasCapecUrl returns a boolean if a field has been set.

### GetRevoked

`func (o *OutputPattern) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *OutputPattern) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *OutputPattern) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.


### GetDataSources

`func (o *OutputPattern) GetDataSources() []string`

GetDataSources returns the DataSources field if non-nil, zero value otherwise.

### GetDataSourcesOk

`func (o *OutputPattern) GetDataSourcesOk() (*[]string, bool)`

GetDataSourcesOk returns a tuple with the DataSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSources

`func (o *OutputPattern) SetDataSources(v []string)`

SetDataSources sets DataSources field to given value.

### HasDataSources

`func (o *OutputPattern) HasDataSources() bool`

HasDataSources returns a boolean if a field has been set.

### GetDefenseBypassed

`func (o *OutputPattern) GetDefenseBypassed() []string`

GetDefenseBypassed returns the DefenseBypassed field if non-nil, zero value otherwise.

### GetDefenseBypassedOk

`func (o *OutputPattern) GetDefenseBypassedOk() (*[]string, bool)`

GetDefenseBypassedOk returns a tuple with the DefenseBypassed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefenseBypassed

`func (o *OutputPattern) SetDefenseBypassed(v []string)`

SetDefenseBypassed sets DefenseBypassed field to given value.

### HasDefenseBypassed

`func (o *OutputPattern) HasDefenseBypassed() bool`

HasDefenseBypassed returns a boolean if a field has been set.

### GetDetection

`func (o *OutputPattern) GetDetection() string`

GetDetection returns the Detection field if non-nil, zero value otherwise.

### GetDetectionOk

`func (o *OutputPattern) GetDetectionOk() (*string, bool)`

GetDetectionOk returns a tuple with the Detection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetection

`func (o *OutputPattern) SetDetection(v string)`

SetDetection sets Detection field to given value.

### HasDetection

`func (o *OutputPattern) HasDetection() bool`

HasDetection returns a boolean if a field has been set.

### GetPermissionsRequired

`func (o *OutputPattern) GetPermissionsRequired() []string`

GetPermissionsRequired returns the PermissionsRequired field if non-nil, zero value otherwise.

### GetPermissionsRequiredOk

`func (o *OutputPattern) GetPermissionsRequiredOk() (*[]string, bool)`

GetPermissionsRequiredOk returns a tuple with the PermissionsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionsRequired

`func (o *OutputPattern) SetPermissionsRequired(v []string)`

SetPermissionsRequired sets PermissionsRequired field to given value.

### HasPermissionsRequired

`func (o *OutputPattern) HasPermissionsRequired() bool`

HasPermissionsRequired returns a boolean if a field has been set.

### GetPlatforms

`func (o *OutputPattern) GetPlatforms() []string`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *OutputPattern) GetPlatformsOk() (*[]string, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *OutputPattern) SetPlatforms(v []string)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *OutputPattern) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetRemoteSupport

`func (o *OutputPattern) GetRemoteSupport() bool`

GetRemoteSupport returns the RemoteSupport field if non-nil, zero value otherwise.

### GetRemoteSupportOk

`func (o *OutputPattern) GetRemoteSupportOk() (*bool, bool)`

GetRemoteSupportOk returns a tuple with the RemoteSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSupport

`func (o *OutputPattern) SetRemoteSupport(v bool)`

SetRemoteSupport sets RemoteSupport field to given value.


### GetSystemRequirements

`func (o *OutputPattern) GetSystemRequirements() []string`

GetSystemRequirements returns the SystemRequirements field if non-nil, zero value otherwise.

### GetSystemRequirementsOk

`func (o *OutputPattern) GetSystemRequirementsOk() (*[]string, bool)`

GetSystemRequirementsOk returns a tuple with the SystemRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemRequirements

`func (o *OutputPattern) SetSystemRequirements(v []string)`

SetSystemRequirements sets SystemRequirements field to given value.

### HasSystemRequirements

`func (o *OutputPattern) HasSystemRequirements() bool`

HasSystemRequirements returns a boolean if a field has been set.

### GetVersion

`func (o *OutputPattern) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OutputPattern) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OutputPattern) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OutputPattern) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetExtraData

`func (o *OutputPattern) GetExtraData() map[string]interface{}`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *OutputPattern) GetExtraDataOk() (*map[string]interface{}, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *OutputPattern) SetExtraData(v map[string]interface{})`

SetExtraData sets ExtraData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


