# AddSnapshotUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Snapshot** | [**Snapshot**](Snapshot.md) |  | 

## Methods

### NewAddSnapshotUpdate

`func NewAddSnapshotUpdate(snapshot Snapshot, ) *AddSnapshotUpdate`

NewAddSnapshotUpdate instantiates a new AddSnapshotUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSnapshotUpdateWithDefaults

`func NewAddSnapshotUpdateWithDefaults() *AddSnapshotUpdate`

NewAddSnapshotUpdateWithDefaults instantiates a new AddSnapshotUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AddSnapshotUpdate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AddSnapshotUpdate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AddSnapshotUpdate) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AddSnapshotUpdate) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetSnapshot

`func (o *AddSnapshotUpdate) GetSnapshot() Snapshot`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *AddSnapshotUpdate) GetSnapshotOk() (*Snapshot, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *AddSnapshotUpdate) SetSnapshot(v Snapshot)`

SetSnapshot sets Snapshot field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


