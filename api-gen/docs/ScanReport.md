# ScanReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TableName** | **string** |  | 
**SnapshotId** | **int64** |  | 
**Filter** | [**Expression**](Expression.md) |  | 
**SchemaId** | **int32** |  | 
**ProjectedFieldIds** | **[]int32** |  | 
**ProjectedFieldNames** | **[]string** |  | 
**Metrics** | [**map[string]MetricResult**](MetricResult.md) |  | 
**Metadata** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewScanReport

`func NewScanReport(tableName string, snapshotId int64, filter Expression, schemaId int32, projectedFieldIds []int32, projectedFieldNames []string, metrics map[string]MetricResult, ) *ScanReport`

NewScanReport instantiates a new ScanReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanReportWithDefaults

`func NewScanReportWithDefaults() *ScanReport`

NewScanReportWithDefaults instantiates a new ScanReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTableName

`func (o *ScanReport) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *ScanReport) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *ScanReport) SetTableName(v string)`

SetTableName sets TableName field to given value.


### GetSnapshotId

`func (o *ScanReport) GetSnapshotId() int64`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *ScanReport) GetSnapshotIdOk() (*int64, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *ScanReport) SetSnapshotId(v int64)`

SetSnapshotId sets SnapshotId field to given value.


### GetFilter

`func (o *ScanReport) GetFilter() Expression`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ScanReport) GetFilterOk() (*Expression, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ScanReport) SetFilter(v Expression)`

SetFilter sets Filter field to given value.


### GetSchemaId

`func (o *ScanReport) GetSchemaId() int32`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *ScanReport) GetSchemaIdOk() (*int32, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *ScanReport) SetSchemaId(v int32)`

SetSchemaId sets SchemaId field to given value.


### GetProjectedFieldIds

`func (o *ScanReport) GetProjectedFieldIds() []int32`

GetProjectedFieldIds returns the ProjectedFieldIds field if non-nil, zero value otherwise.

### GetProjectedFieldIdsOk

`func (o *ScanReport) GetProjectedFieldIdsOk() (*[]int32, bool)`

GetProjectedFieldIdsOk returns a tuple with the ProjectedFieldIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedFieldIds

`func (o *ScanReport) SetProjectedFieldIds(v []int32)`

SetProjectedFieldIds sets ProjectedFieldIds field to given value.


### GetProjectedFieldNames

`func (o *ScanReport) GetProjectedFieldNames() []string`

GetProjectedFieldNames returns the ProjectedFieldNames field if non-nil, zero value otherwise.

### GetProjectedFieldNamesOk

`func (o *ScanReport) GetProjectedFieldNamesOk() (*[]string, bool)`

GetProjectedFieldNamesOk returns a tuple with the ProjectedFieldNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedFieldNames

`func (o *ScanReport) SetProjectedFieldNames(v []string)`

SetProjectedFieldNames sets ProjectedFieldNames field to given value.


### GetMetrics

`func (o *ScanReport) GetMetrics() map[string]MetricResult`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ScanReport) GetMetricsOk() (*map[string]MetricResult, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ScanReport) SetMetrics(v map[string]MetricResult)`

SetMetrics sets Metrics field to given value.


### GetMetadata

`func (o *ScanReport) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ScanReport) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ScanReport) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ScanReport) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


