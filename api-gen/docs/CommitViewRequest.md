# CommitViewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to [**TableIdentifier**](TableIdentifier.md) | View identifier to update | [optional] 
**Requirements** | Pointer to [**[]ViewRequirement**](ViewRequirement.md) |  | [optional] 
**Updates** | [**[]ViewUpdate**](ViewUpdate.md) |  | 

## Methods

### NewCommitViewRequest

`func NewCommitViewRequest(updates []ViewUpdate, ) *CommitViewRequest`

NewCommitViewRequest instantiates a new CommitViewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitViewRequestWithDefaults

`func NewCommitViewRequestWithDefaults() *CommitViewRequest`

NewCommitViewRequestWithDefaults instantiates a new CommitViewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *CommitViewRequest) GetIdentifier() TableIdentifier`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CommitViewRequest) GetIdentifierOk() (*TableIdentifier, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CommitViewRequest) SetIdentifier(v TableIdentifier)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *CommitViewRequest) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetRequirements

`func (o *CommitViewRequest) GetRequirements() []ViewRequirement`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *CommitViewRequest) GetRequirementsOk() (*[]ViewRequirement, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *CommitViewRequest) SetRequirements(v []ViewRequirement)`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *CommitViewRequest) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.

### GetUpdates

`func (o *CommitViewRequest) GetUpdates() []ViewUpdate`

GetUpdates returns the Updates field if non-nil, zero value otherwise.

### GetUpdatesOk

`func (o *CommitViewRequest) GetUpdatesOk() (*[]ViewUpdate, bool)`

GetUpdatesOk returns a tuple with the Updates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdates

`func (o *CommitViewRequest) SetUpdates(v []ViewUpdate)`

SetUpdates sets Updates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


