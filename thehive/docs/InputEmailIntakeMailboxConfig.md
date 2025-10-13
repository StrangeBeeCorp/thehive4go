# InputEmailIntakeMailboxConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | [**InputEmailIntakeProvider**](InputEmailIntakeProvider.md) |  | 
**Credential** | [**InputEmailIntakeCredential**](InputEmailIntakeCredential.md) |  | 
**Inbox** | Pointer to **string** | The folder where the emails are fetched | [optional] [default to "Inbox"]
**Archive** | Pointer to **string** | The folder where the email should be moved | [optional] 
**MarkAsRead** | Pointer to **bool** |  | [optional] [default to false]
**Kind** | Pointer to **string** | The type of the provider config (imap, api) | [optional] 

## Methods

### NewInputEmailIntakeMailboxConfig

`func NewInputEmailIntakeMailboxConfig(provider InputEmailIntakeProvider, credential InputEmailIntakeCredential, ) *InputEmailIntakeMailboxConfig`

NewInputEmailIntakeMailboxConfig instantiates a new InputEmailIntakeMailboxConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputEmailIntakeMailboxConfigWithDefaults

`func NewInputEmailIntakeMailboxConfigWithDefaults() *InputEmailIntakeMailboxConfig`

NewInputEmailIntakeMailboxConfigWithDefaults instantiates a new InputEmailIntakeMailboxConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *InputEmailIntakeMailboxConfig) GetProvider() InputEmailIntakeProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *InputEmailIntakeMailboxConfig) GetProviderOk() (*InputEmailIntakeProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *InputEmailIntakeMailboxConfig) SetProvider(v InputEmailIntakeProvider)`

SetProvider sets Provider field to given value.


### GetCredential

`func (o *InputEmailIntakeMailboxConfig) GetCredential() InputEmailIntakeCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *InputEmailIntakeMailboxConfig) GetCredentialOk() (*InputEmailIntakeCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *InputEmailIntakeMailboxConfig) SetCredential(v InputEmailIntakeCredential)`

SetCredential sets Credential field to given value.


### GetInbox

`func (o *InputEmailIntakeMailboxConfig) GetInbox() string`

GetInbox returns the Inbox field if non-nil, zero value otherwise.

### GetInboxOk

`func (o *InputEmailIntakeMailboxConfig) GetInboxOk() (*string, bool)`

GetInboxOk returns a tuple with the Inbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbox

`func (o *InputEmailIntakeMailboxConfig) SetInbox(v string)`

SetInbox sets Inbox field to given value.

### HasInbox

`func (o *InputEmailIntakeMailboxConfig) HasInbox() bool`

HasInbox returns a boolean if a field has been set.

### GetArchive

`func (o *InputEmailIntakeMailboxConfig) GetArchive() string`

GetArchive returns the Archive field if non-nil, zero value otherwise.

### GetArchiveOk

`func (o *InputEmailIntakeMailboxConfig) GetArchiveOk() (*string, bool)`

GetArchiveOk returns a tuple with the Archive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchive

`func (o *InputEmailIntakeMailboxConfig) SetArchive(v string)`

SetArchive sets Archive field to given value.

### HasArchive

`func (o *InputEmailIntakeMailboxConfig) HasArchive() bool`

HasArchive returns a boolean if a field has been set.

### GetMarkAsRead

`func (o *InputEmailIntakeMailboxConfig) GetMarkAsRead() bool`

GetMarkAsRead returns the MarkAsRead field if non-nil, zero value otherwise.

### GetMarkAsReadOk

`func (o *InputEmailIntakeMailboxConfig) GetMarkAsReadOk() (*bool, bool)`

GetMarkAsReadOk returns a tuple with the MarkAsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkAsRead

`func (o *InputEmailIntakeMailboxConfig) SetMarkAsRead(v bool)`

SetMarkAsRead sets MarkAsRead field to given value.

### HasMarkAsRead

`func (o *InputEmailIntakeMailboxConfig) HasMarkAsRead() bool`

HasMarkAsRead returns a boolean if a field has been set.

### GetKind

`func (o *InputEmailIntakeMailboxConfig) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *InputEmailIntakeMailboxConfig) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *InputEmailIntakeMailboxConfig) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *InputEmailIntakeMailboxConfig) HasKind() bool`

HasKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


