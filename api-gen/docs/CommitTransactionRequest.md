# CommitTransactionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TableChanges** | [**[]CommitTableRequest**](CommitTableRequest.md) |  | 

## Methods

### NewCommitTransactionRequest

`func NewCommitTransactionRequest(tableChanges []CommitTableRequest, ) *CommitTransactionRequest`

NewCommitTransactionRequest instantiates a new CommitTransactionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitTransactionRequestWithDefaults

`func NewCommitTransactionRequestWithDefaults() *CommitTransactionRequest`

NewCommitTransactionRequestWithDefaults instantiates a new CommitTransactionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTableChanges

`func (o *CommitTransactionRequest) GetTableChanges() []CommitTableRequest`

GetTableChanges returns the TableChanges field if non-nil, zero value otherwise.

### GetTableChangesOk

`func (o *CommitTransactionRequest) GetTableChangesOk() (*[]CommitTableRequest, bool)`

GetTableChangesOk returns a tuple with the TableChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableChanges

`func (o *CommitTransactionRequest) SetTableChanges(v []CommitTableRequest)`

SetTableChanges sets TableChanges field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


