# CommitTableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to [**TableIdentifier**](TableIdentifier.md) | Table identifier to update; must be present for CommitTransactionRequest | [optional] 
**Requirements** | [**[]TableRequirement**](TableRequirement.md) |  | 
**Updates** | [**[]TableUpdate**](TableUpdate.md) |  | 

## Methods

### NewCommitTableRequest

`func NewCommitTableRequest(requirements []TableRequirement, updates []TableUpdate, ) *CommitTableRequest`

NewCommitTableRequest instantiates a new CommitTableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitTableRequestWithDefaults

`func NewCommitTableRequestWithDefaults() *CommitTableRequest`

NewCommitTableRequestWithDefaults instantiates a new CommitTableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *CommitTableRequest) GetIdentifier() TableIdentifier`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CommitTableRequest) GetIdentifierOk() (*TableIdentifier, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CommitTableRequest) SetIdentifier(v TableIdentifier)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *CommitTableRequest) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetRequirements

`func (o *CommitTableRequest) GetRequirements() []TableRequirement`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *CommitTableRequest) GetRequirementsOk() (*[]TableRequirement, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *CommitTableRequest) SetRequirements(v []TableRequirement)`

SetRequirements sets Requirements field to given value.


### GetUpdates

`func (o *CommitTableRequest) GetUpdates() []TableUpdate`

GetUpdates returns the Updates field if non-nil, zero value otherwise.

### GetUpdatesOk

`func (o *CommitTableRequest) GetUpdatesOk() (*[]TableUpdate, bool)`

GetUpdatesOk returns a tuple with the Updates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdates

`func (o *CommitTableRequest) SetUpdates(v []TableUpdate)`

SetUpdates sets Updates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


