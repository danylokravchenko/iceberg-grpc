# SetSnapshotRefUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**RefName** | **string** |  | 
**Type** | **string** |  | 
**SnapshotId** | **int64** |  | 
**MaxRefAgeMs** | Pointer to **int64** |  | [optional] 
**MaxSnapshotAgeMs** | Pointer to **int64** |  | [optional] 
**MinSnapshotsToKeep** | Pointer to **int32** |  | [optional] 

## Methods

### NewSetSnapshotRefUpdate

`func NewSetSnapshotRefUpdate(refName string, type_ string, snapshotId int64, ) *SetSnapshotRefUpdate`

NewSetSnapshotRefUpdate instantiates a new SetSnapshotRefUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetSnapshotRefUpdateWithDefaults

`func NewSetSnapshotRefUpdateWithDefaults() *SetSnapshotRefUpdate`

NewSetSnapshotRefUpdateWithDefaults instantiates a new SetSnapshotRefUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *SetSnapshotRefUpdate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SetSnapshotRefUpdate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SetSnapshotRefUpdate) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SetSnapshotRefUpdate) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetRefName

`func (o *SetSnapshotRefUpdate) GetRefName() string`

GetRefName returns the RefName field if non-nil, zero value otherwise.

### GetRefNameOk

`func (o *SetSnapshotRefUpdate) GetRefNameOk() (*string, bool)`

GetRefNameOk returns a tuple with the RefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefName

`func (o *SetSnapshotRefUpdate) SetRefName(v string)`

SetRefName sets RefName field to given value.


### GetType

`func (o *SetSnapshotRefUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SetSnapshotRefUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SetSnapshotRefUpdate) SetType(v string)`

SetType sets Type field to given value.


### GetSnapshotId

`func (o *SetSnapshotRefUpdate) GetSnapshotId() int64`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *SetSnapshotRefUpdate) GetSnapshotIdOk() (*int64, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *SetSnapshotRefUpdate) SetSnapshotId(v int64)`

SetSnapshotId sets SnapshotId field to given value.


### GetMaxRefAgeMs

`func (o *SetSnapshotRefUpdate) GetMaxRefAgeMs() int64`

GetMaxRefAgeMs returns the MaxRefAgeMs field if non-nil, zero value otherwise.

### GetMaxRefAgeMsOk

`func (o *SetSnapshotRefUpdate) GetMaxRefAgeMsOk() (*int64, bool)`

GetMaxRefAgeMsOk returns a tuple with the MaxRefAgeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRefAgeMs

`func (o *SetSnapshotRefUpdate) SetMaxRefAgeMs(v int64)`

SetMaxRefAgeMs sets MaxRefAgeMs field to given value.

### HasMaxRefAgeMs

`func (o *SetSnapshotRefUpdate) HasMaxRefAgeMs() bool`

HasMaxRefAgeMs returns a boolean if a field has been set.

### GetMaxSnapshotAgeMs

`func (o *SetSnapshotRefUpdate) GetMaxSnapshotAgeMs() int64`

GetMaxSnapshotAgeMs returns the MaxSnapshotAgeMs field if non-nil, zero value otherwise.

### GetMaxSnapshotAgeMsOk

`func (o *SetSnapshotRefUpdate) GetMaxSnapshotAgeMsOk() (*int64, bool)`

GetMaxSnapshotAgeMsOk returns a tuple with the MaxSnapshotAgeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSnapshotAgeMs

`func (o *SetSnapshotRefUpdate) SetMaxSnapshotAgeMs(v int64)`

SetMaxSnapshotAgeMs sets MaxSnapshotAgeMs field to given value.

### HasMaxSnapshotAgeMs

`func (o *SetSnapshotRefUpdate) HasMaxSnapshotAgeMs() bool`

HasMaxSnapshotAgeMs returns a boolean if a field has been set.

### GetMinSnapshotsToKeep

`func (o *SetSnapshotRefUpdate) GetMinSnapshotsToKeep() int32`

GetMinSnapshotsToKeep returns the MinSnapshotsToKeep field if non-nil, zero value otherwise.

### GetMinSnapshotsToKeepOk

`func (o *SetSnapshotRefUpdate) GetMinSnapshotsToKeepOk() (*int32, bool)`

GetMinSnapshotsToKeepOk returns a tuple with the MinSnapshotsToKeep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSnapshotsToKeep

`func (o *SetSnapshotRefUpdate) SetMinSnapshotsToKeep(v int32)`

SetMinSnapshotsToKeep sets MinSnapshotsToKeep field to given value.

### HasMinSnapshotsToKeep

`func (o *SetSnapshotRefUpdate) HasMinSnapshotsToKeep() bool`

HasMinSnapshotsToKeep returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


