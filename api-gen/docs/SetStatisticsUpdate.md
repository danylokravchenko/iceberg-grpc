# SetStatisticsUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**SnapshotId** | Pointer to **int64** | This optional field is **DEPRECATED for REMOVAL** since it contains redundant information. Clients should use the &#x60;statistics.snapshot-id&#x60; field instead. | [optional] 
**Statistics** | [**StatisticsFile**](StatisticsFile.md) |  | 

## Methods

### NewSetStatisticsUpdate

`func NewSetStatisticsUpdate(statistics StatisticsFile, ) *SetStatisticsUpdate`

NewSetStatisticsUpdate instantiates a new SetStatisticsUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetStatisticsUpdateWithDefaults

`func NewSetStatisticsUpdateWithDefaults() *SetStatisticsUpdate`

NewSetStatisticsUpdateWithDefaults instantiates a new SetStatisticsUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *SetStatisticsUpdate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SetStatisticsUpdate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SetStatisticsUpdate) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SetStatisticsUpdate) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetSnapshotId

`func (o *SetStatisticsUpdate) GetSnapshotId() int64`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *SetStatisticsUpdate) GetSnapshotIdOk() (*int64, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *SetStatisticsUpdate) SetSnapshotId(v int64)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *SetStatisticsUpdate) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetStatistics

`func (o *SetStatisticsUpdate) GetStatistics() StatisticsFile`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *SetStatisticsUpdate) GetStatisticsOk() (*StatisticsFile, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *SetStatisticsUpdate) SetStatistics(v StatisticsFile)`

SetStatistics sets Statistics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


