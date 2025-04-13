# SetPartitionStatisticsUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**PartitionStatistics** | [**PartitionStatisticsFile**](PartitionStatisticsFile.md) |  | 

## Methods

### NewSetPartitionStatisticsUpdate

`func NewSetPartitionStatisticsUpdate(partitionStatistics PartitionStatisticsFile, ) *SetPartitionStatisticsUpdate`

NewSetPartitionStatisticsUpdate instantiates a new SetPartitionStatisticsUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetPartitionStatisticsUpdateWithDefaults

`func NewSetPartitionStatisticsUpdateWithDefaults() *SetPartitionStatisticsUpdate`

NewSetPartitionStatisticsUpdateWithDefaults instantiates a new SetPartitionStatisticsUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *SetPartitionStatisticsUpdate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SetPartitionStatisticsUpdate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SetPartitionStatisticsUpdate) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SetPartitionStatisticsUpdate) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetPartitionStatistics

`func (o *SetPartitionStatisticsUpdate) GetPartitionStatistics() PartitionStatisticsFile`

GetPartitionStatistics returns the PartitionStatistics field if non-nil, zero value otherwise.

### GetPartitionStatisticsOk

`func (o *SetPartitionStatisticsUpdate) GetPartitionStatisticsOk() (*PartitionStatisticsFile, bool)`

GetPartitionStatisticsOk returns a tuple with the PartitionStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionStatistics

`func (o *SetPartitionStatisticsUpdate) SetPartitionStatistics(v PartitionStatisticsFile)`

SetPartitionStatistics sets PartitionStatistics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


