# OutputEmailIntakeMailbox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | [**OutputEmailIntakeImapProvider**](OutputEmailIntakeImapProvider.md) |  | 
**Credential** | [**OutputEmailIntakeCredential**](OutputEmailIntakeCredential.md) |  | 
**Inbox** | **string** |  | 
**Archive** | Pointer to **string** |  | [optional] 
**MarkAsRead** | **bool** |  | 
**Kind** | **string** |  | 

## Methods

### NewOutputEmailIntakeMailbox

`func NewOutputEmailIntakeMailbox(provider OutputEmailIntakeImapProvider, credential OutputEmailIntakeCredential, inbox string, markAsRead bool, kind string, ) *OutputEmailIntakeMailbox`

NewOutputEmailIntakeMailbox instantiates a new OutputEmailIntakeMailbox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputEmailIntakeMailboxWithDefaults

`func NewOutputEmailIntakeMailboxWithDefaults() *OutputEmailIntakeMailbox`

NewOutputEmailIntakeMailboxWithDefaults instantiates a new OutputEmailIntakeMailbox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *OutputEmailIntakeMailbox) GetProvider() OutputEmailIntakeImapProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *OutputEmailIntakeMailbox) GetProviderOk() (*OutputEmailIntakeImapProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *OutputEmailIntakeMailbox) SetProvider(v OutputEmailIntakeImapProvider)`

SetProvider sets Provider field to given value.


### GetCredential

`func (o *OutputEmailIntakeMailbox) GetCredential() OutputEmailIntakeCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *OutputEmailIntakeMailbox) GetCredentialOk() (*OutputEmailIntakeCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *OutputEmailIntakeMailbox) SetCredential(v OutputEmailIntakeCredential)`

SetCredential sets Credential field to given value.


### GetInbox

`func (o *OutputEmailIntakeMailbox) GetInbox() string`

GetInbox returns the Inbox field if non-nil, zero value otherwise.

### GetInboxOk

`func (o *OutputEmailIntakeMailbox) GetInboxOk() (*string, bool)`

GetInboxOk returns a tuple with the Inbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbox

`func (o *OutputEmailIntakeMailbox) SetInbox(v string)`

SetInbox sets Inbox field to given value.


### GetArchive

`func (o *OutputEmailIntakeMailbox) GetArchive() string`

GetArchive returns the Archive field if non-nil, zero value otherwise.

### GetArchiveOk

`func (o *OutputEmailIntakeMailbox) GetArchiveOk() (*string, bool)`

GetArchiveOk returns a tuple with the Archive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchive

`func (o *OutputEmailIntakeMailbox) SetArchive(v string)`

SetArchive sets Archive field to given value.

### HasArchive

`func (o *OutputEmailIntakeMailbox) HasArchive() bool`

HasArchive returns a boolean if a field has been set.

### GetMarkAsRead

`func (o *OutputEmailIntakeMailbox) GetMarkAsRead() bool`

GetMarkAsRead returns the MarkAsRead field if non-nil, zero value otherwise.

### GetMarkAsReadOk

`func (o *OutputEmailIntakeMailbox) GetMarkAsReadOk() (*bool, bool)`

GetMarkAsReadOk returns a tuple with the MarkAsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkAsRead

`func (o *OutputEmailIntakeMailbox) SetMarkAsRead(v bool)`

SetMarkAsRead sets MarkAsRead field to given value.


### GetKind

`func (o *OutputEmailIntakeMailbox) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *OutputEmailIntakeMailbox) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *OutputEmailIntakeMailbox) SetKind(v string)`

SetKind sets Kind field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


