# CommitReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TableName** | **string** |  | 
**SnapshotId** | **int64** |  | 
**SequenceNumber** | **int64** |  | 
**Operation** | **string** |  | 
**Metrics** | [**map[string]MetricResult**](MetricResult.md) |  | 
**Metadata** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCommitReport

`func NewCommitReport(tableName string, snapshotId int64, sequenceNumber int64, operation string, metrics map[string]MetricResult, ) *CommitReport`

NewCommitReport instantiates a new CommitReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitReportWithDefaults

`func NewCommitReportWithDefaults() *CommitReport`

NewCommitReportWithDefaults instantiates a new CommitReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTableName

`func (o *CommitReport) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *CommitReport) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *CommitReport) SetTableName(v string)`

SetTableName sets TableName field to given value.


### GetSnapshotId

`func (o *CommitReport) GetSnapshotId() int64`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *CommitReport) GetSnapshotIdOk() (*int64, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *CommitReport) SetSnapshotId(v int64)`

SetSnapshotId sets SnapshotId field to given value.


### GetSequenceNumber

`func (o *CommitReport) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *CommitReport) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *CommitReport) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.


### GetOperation

`func (o *CommitReport) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *CommitReport) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *CommitReport) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetMetrics

`func (o *CommitReport) GetMetrics() map[string]MetricResult`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *CommitReport) GetMetricsOk() (*map[string]MetricResult, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *CommitReport) SetMetrics(v map[string]MetricResult)`

SetMetrics sets Metrics field to given value.


### GetMetadata

`func (o *CommitReport) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CommitReport) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CommitReport) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CommitReport) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


