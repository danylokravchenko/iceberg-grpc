# ReportMetricsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportType** | **string** |  | 
**TableName** | **string** |  | 
**SnapshotId** | **int64** |  | 
**Filter** | [**Expression**](Expression.md) |  | 
**SchemaId** | **int32** |  | 
**ProjectedFieldIds** | **[]int32** |  | 
**ProjectedFieldNames** | **[]string** |  | 
**Metrics** | [**map[string]MetricResult**](MetricResult.md) |  | 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**SequenceNumber** | **int64** |  | 
**Operation** | **string** |  | 

## Methods

### NewReportMetricsRequest

`func NewReportMetricsRequest(reportType string, tableName string, snapshotId int64, filter Expression, schemaId int32, projectedFieldIds []int32, projectedFieldNames []string, metrics map[string]MetricResult, sequenceNumber int64, operation string, ) *ReportMetricsRequest`

NewReportMetricsRequest instantiates a new ReportMetricsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportMetricsRequestWithDefaults

`func NewReportMetricsRequestWithDefaults() *ReportMetricsRequest`

NewReportMetricsRequestWithDefaults instantiates a new ReportMetricsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportType

`func (o *ReportMetricsRequest) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *ReportMetricsRequest) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *ReportMetricsRequest) SetReportType(v string)`

SetReportType sets ReportType field to given value.


### GetTableName

`func (o *ReportMetricsRequest) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *ReportMetricsRequest) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *ReportMetricsRequest) SetTableName(v string)`

SetTableName sets TableName field to given value.


### GetSnapshotId

`func (o *ReportMetricsRequest) GetSnapshotId() int64`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *ReportMetricsRequest) GetSnapshotIdOk() (*int64, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *ReportMetricsRequest) SetSnapshotId(v int64)`

SetSnapshotId sets SnapshotId field to given value.


### GetFilter

`func (o *ReportMetricsRequest) GetFilter() Expression`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ReportMetricsRequest) GetFilterOk() (*Expression, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ReportMetricsRequest) SetFilter(v Expression)`

SetFilter sets Filter field to given value.


### GetSchemaId

`func (o *ReportMetricsRequest) GetSchemaId() int32`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *ReportMetricsRequest) GetSchemaIdOk() (*int32, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *ReportMetricsRequest) SetSchemaId(v int32)`

SetSchemaId sets SchemaId field to given value.


### GetProjectedFieldIds

`func (o *ReportMetricsRequest) GetProjectedFieldIds() []int32`

GetProjectedFieldIds returns the ProjectedFieldIds field if non-nil, zero value otherwise.

### GetProjectedFieldIdsOk

`func (o *ReportMetricsRequest) GetProjectedFieldIdsOk() (*[]int32, bool)`

GetProjectedFieldIdsOk returns a tuple with the ProjectedFieldIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedFieldIds

`func (o *ReportMetricsRequest) SetProjectedFieldIds(v []int32)`

SetProjectedFieldIds sets ProjectedFieldIds field to given value.


### GetProjectedFieldNames

`func (o *ReportMetricsRequest) GetProjectedFieldNames() []string`

GetProjectedFieldNames returns the ProjectedFieldNames field if non-nil, zero value otherwise.

### GetProjectedFieldNamesOk

`func (o *ReportMetricsRequest) GetProjectedFieldNamesOk() (*[]string, bool)`

GetProjectedFieldNamesOk returns a tuple with the ProjectedFieldNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedFieldNames

`func (o *ReportMetricsRequest) SetProjectedFieldNames(v []string)`

SetProjectedFieldNames sets ProjectedFieldNames field to given value.


### GetMetrics

`func (o *ReportMetricsRequest) GetMetrics() map[string]MetricResult`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ReportMetricsRequest) GetMetricsOk() (*map[string]MetricResult, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ReportMetricsRequest) SetMetrics(v map[string]MetricResult)`

SetMetrics sets Metrics field to given value.


### GetMetadata

`func (o *ReportMetricsRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ReportMetricsRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ReportMetricsRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ReportMetricsRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *ReportMetricsRequest) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *ReportMetricsRequest) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *ReportMetricsRequest) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.


### GetOperation

`func (o *ReportMetricsRequest) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ReportMetricsRequest) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ReportMetricsRequest) SetOperation(v string)`

SetOperation sets Operation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


