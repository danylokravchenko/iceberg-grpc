# RemovePartitionStatisticsUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**SnapshotId** | **int64** |  | 

## Methods

### NewRemovePartitionStatisticsUpdate

`func NewRemovePartitionStatisticsUpdate(snapshotId int64, ) *RemovePartitionStatisticsUpdate`

NewRemovePartitionStatisticsUpdate instantiates a new RemovePartitionStatisticsUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemovePartitionStatisticsUpdateWithDefaults

`func NewRemovePartitionStatisticsUpdateWithDefaults() *RemovePartitionStatisticsUpdate`

NewRemovePartitionStatisticsUpdateWithDefaults instantiates a new RemovePartitionStatisticsUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *RemovePartitionStatisticsUpdate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RemovePartitionStatisticsUpdate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RemovePartitionStatisticsUpdate) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *RemovePartitionStatisticsUpdate) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetSnapshotId

`func (o *RemovePartitionStatisticsUpdate) GetSnapshotId() int64`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *RemovePartitionStatisticsUpdate) GetSnapshotIdOk() (*int64, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *RemovePartitionStatisticsUpdate) SetSnapshotId(v int64)`

SetSnapshotId sets SnapshotId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


