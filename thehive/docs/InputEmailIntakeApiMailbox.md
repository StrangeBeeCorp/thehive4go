# InputEmailIntakeApiMailbox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | [**InputEmailIntakeApiProvider**](InputEmailIntakeApiProvider.md) |  | 
**Credential** | [**InputEmailIntakeCredential**](InputEmailIntakeCredential.md) |  | 
**Inbox** | Pointer to **string** | The folder where the emails are fetched | [optional] [default to "Inbox"]
**Archive** | Pointer to **string** | The folder where the email should be moved | [optional] 
**MarkAsRead** | Pointer to **bool** |  | [optional] [default to false]
**Kind** | Pointer to **string** | The type of the provider config (imap, api) | [optional] 

## Methods

### NewInputEmailIntakeApiMailbox

`func NewInputEmailIntakeApiMailbox(provider InputEmailIntakeApiProvider, credential InputEmailIntakeCredential, ) *InputEmailIntakeApiMailbox`

NewInputEmailIntakeApiMailbox instantiates a new InputEmailIntakeApiMailbox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputEmailIntakeApiMailboxWithDefaults

`func NewInputEmailIntakeApiMailboxWithDefaults() *InputEmailIntakeApiMailbox`

NewInputEmailIntakeApiMailboxWithDefaults instantiates a new InputEmailIntakeApiMailbox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *InputEmailIntakeApiMailbox) GetProvider() InputEmailIntakeApiProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *InputEmailIntakeApiMailbox) GetProviderOk() (*InputEmailIntakeApiProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *InputEmailIntakeApiMailbox) SetProvider(v InputEmailIntakeApiProvider)`

SetProvider sets Provider field to given value.


### GetCredential

`func (o *InputEmailIntakeApiMailbox) GetCredential() InputEmailIntakeCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *InputEmailIntakeApiMailbox) GetCredentialOk() (*InputEmailIntakeCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *InputEmailIntakeApiMailbox) SetCredential(v InputEmailIntakeCredential)`

SetCredential sets Credential field to given value.


### GetInbox

`func (o *InputEmailIntakeApiMailbox) GetInbox() string`

GetInbox returns the Inbox field if non-nil, zero value otherwise.

### GetInboxOk

`func (o *InputEmailIntakeApiMailbox) GetInboxOk() (*string, bool)`

GetInboxOk returns a tuple with the Inbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbox

`func (o *InputEmailIntakeApiMailbox) SetInbox(v string)`

SetInbox sets Inbox field to given value.

### HasInbox

`func (o *InputEmailIntakeApiMailbox) HasInbox() bool`

HasInbox returns a boolean if a field has been set.

### GetArchive

`func (o *InputEmailIntakeApiMailbox) GetArchive() string`

GetArchive returns the Archive field if non-nil, zero value otherwise.

### GetArchiveOk

`func (o *InputEmailIntakeApiMailbox) GetArchiveOk() (*string, bool)`

GetArchiveOk returns a tuple with the Archive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchive

`func (o *InputEmailIntakeApiMailbox) SetArchive(v string)`

SetArchive sets Archive field to given value.

### HasArchive

`func (o *InputEmailIntakeApiMailbox) HasArchive() bool`

HasArchive returns a boolean if a field has been set.

### GetMarkAsRead

`func (o *InputEmailIntakeApiMailbox) GetMarkAsRead() bool`

GetMarkAsRead returns the MarkAsRead field if non-nil, zero value otherwise.

### GetMarkAsReadOk

`func (o *InputEmailIntakeApiMailbox) GetMarkAsReadOk() (*bool, bool)`

GetMarkAsReadOk returns a tuple with the MarkAsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkAsRead

`func (o *InputEmailIntakeApiMailbox) SetMarkAsRead(v bool)`

SetMarkAsRead sets MarkAsRead field to given value.

### HasMarkAsRead

`func (o *InputEmailIntakeApiMailbox) HasMarkAsRead() bool`

HasMarkAsRead returns a boolean if a field has been set.

### GetKind

`func (o *InputEmailIntakeApiMailbox) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *InputEmailIntakeApiMailbox) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *InputEmailIntakeApiMailbox) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *InputEmailIntakeApiMailbox) HasKind() bool`

HasKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


