# RemoveSnapshotsUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**SnapshotIds** | **[]int64** |  | 

## Methods

### NewRemoveSnapshotsUpdate

`func NewRemoveSnapshotsUpdate(snapshotIds []int64, ) *RemoveSnapshotsUpdate`

NewRemoveSnapshotsUpdate instantiates a new RemoveSnapshotsUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveSnapshotsUpdateWithDefaults

`func NewRemoveSnapshotsUpdateWithDefaults() *RemoveSnapshotsUpdate`

NewRemoveSnapshotsUpdateWithDefaults instantiates a new RemoveSnapshotsUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *RemoveSnapshotsUpdate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RemoveSnapshotsUpdate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RemoveSnapshotsUpdate) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *RemoveSnapshotsUpdate) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetSnapshotIds

`func (o *RemoveSnapshotsUpdate) GetSnapshotIds() []int64`

GetSnapshotIds returns the SnapshotIds field if non-nil, zero value otherwise.

### GetSnapshotIdsOk

`func (o *RemoveSnapshotsUpdate) GetSnapshotIdsOk() (*[]int64, bool)`

GetSnapshotIdsOk returns a tuple with the SnapshotIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotIds

`func (o *RemoveSnapshotsUpdate) SetSnapshotIds(v []int64)`

SetSnapshotIds sets SnapshotIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


