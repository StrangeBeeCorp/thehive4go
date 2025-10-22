# InputEmailIntakeImapMailbox

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

### NewInputEmailIntakeImapMailbox

`func NewInputEmailIntakeImapMailbox(provider InputEmailIntakeProvider, credential InputEmailIntakeCredential, ) *InputEmailIntakeImapMailbox`

NewInputEmailIntakeImapMailbox instantiates a new InputEmailIntakeImapMailbox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputEmailIntakeImapMailboxWithDefaults

`func NewInputEmailIntakeImapMailboxWithDefaults() *InputEmailIntakeImapMailbox`

NewInputEmailIntakeImapMailboxWithDefaults instantiates a new InputEmailIntakeImapMailbox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *InputEmailIntakeImapMailbox) GetProvider() InputEmailIntakeProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *InputEmailIntakeImapMailbox) GetProviderOk() (*InputEmailIntakeProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *InputEmailIntakeImapMailbox) SetProvider(v InputEmailIntakeProvider)`

SetProvider sets Provider field to given value.


### GetCredential

`func (o *InputEmailIntakeImapMailbox) GetCredential() InputEmailIntakeCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *InputEmailIntakeImapMailbox) GetCredentialOk() (*InputEmailIntakeCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *InputEmailIntakeImapMailbox) SetCredential(v InputEmailIntakeCredential)`

SetCredential sets Credential field to given value.


### GetInbox

`func (o *InputEmailIntakeImapMailbox) GetInbox() string`

GetInbox returns the Inbox field if non-nil, zero value otherwise.

### GetInboxOk

`func (o *InputEmailIntakeImapMailbox) GetInboxOk() (*string, bool)`

GetInboxOk returns a tuple with the Inbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbox

`func (o *InputEmailIntakeImapMailbox) SetInbox(v string)`

SetInbox sets Inbox field to given value.

### HasInbox

`func (o *InputEmailIntakeImapMailbox) HasInbox() bool`

HasInbox returns a boolean if a field has been set.

### GetArchive

`func (o *InputEmailIntakeImapMailbox) GetArchive() string`

GetArchive returns the Archive field if non-nil, zero value otherwise.

### GetArchiveOk

`func (o *InputEmailIntakeImapMailbox) GetArchiveOk() (*string, bool)`

GetArchiveOk returns a tuple with the Archive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchive

`func (o *InputEmailIntakeImapMailbox) SetArchive(v string)`

SetArchive sets Archive field to given value.

### HasArchive

`func (o *InputEmailIntakeImapMailbox) HasArchive() bool`

HasArchive returns a boolean if a field has been set.

### GetMarkAsRead

`func (o *InputEmailIntakeImapMailbox) GetMarkAsRead() bool`

GetMarkAsRead returns the MarkAsRead field if non-nil, zero value otherwise.

### GetMarkAsReadOk

`func (o *InputEmailIntakeImapMailbox) GetMarkAsReadOk() (*bool, bool)`

GetMarkAsReadOk returns a tuple with the MarkAsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkAsRead

`func (o *InputEmailIntakeImapMailbox) SetMarkAsRead(v bool)`

SetMarkAsRead sets MarkAsRead field to given value.

### HasMarkAsRead

`func (o *InputEmailIntakeImapMailbox) HasMarkAsRead() bool`

HasMarkAsRead returns a boolean if a field has been set.

### GetKind

`func (o *InputEmailIntakeImapMailbox) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *InputEmailIntakeImapMailbox) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *InputEmailIntakeImapMailbox) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *InputEmailIntakeImapMailbox) HasKind() bool`

HasKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


