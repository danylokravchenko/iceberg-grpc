# PlanTableScanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotId** | Pointer to **int64** | Identifier for the snapshot to scan in a point-in-time scan | [optional] 
**Select** | Pointer to **[]string** | List of selected schema fields | [optional] 
**Filter** | Pointer to [**Expression**](Expression.md) | Expression used to filter the table data | [optional] 
**CaseSensitive** | Pointer to **bool** | Enables case sensitive field matching for filter and select | [optional] [default to true]
**UseSnapshotSchema** | Pointer to **bool** | Whether to use the schema at the time the snapshot was written. When time travelling, the snapshot schema should be used (true). When scanning a branch, the table schema should be used (false). | [optional] [default to false]
**StartSnapshotId** | Pointer to **int64** | Starting snapshot ID for an incremental scan (exclusive) | [optional] 
**EndSnapshotId** | Pointer to **int64** | Ending snapshot ID for an incremental scan (inclusive). Required when start-snapshot-id is specified. | [optional] 
**StatsFields** | Pointer to **[]string** | List of fields for which the service should send column stats. | [optional] 

## Methods

### NewPlanTableScanRequest

`func NewPlanTableScanRequest() *PlanTableScanRequest`

NewPlanTableScanRequest instantiates a new PlanTableScanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanTableScanRequestWithDefaults

`func NewPlanTableScanRequestWithDefaults() *PlanTableScanRequest`

NewPlanTableScanRequestWithDefaults instantiates a new PlanTableScanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshotId

`func (o *PlanTableScanRequest) GetSnapshotId() int64`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *PlanTableScanRequest) GetSnapshotIdOk() (*int64, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *PlanTableScanRequest) SetSnapshotId(v int64)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *PlanTableScanRequest) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetSelect

`func (o *PlanTableScanRequest) GetSelect() []string`

GetSelect returns the Select field if non-nil, zero value otherwise.

### GetSelectOk

`func (o *PlanTableScanRequest) GetSelectOk() (*[]string, bool)`

GetSelectOk returns a tuple with the Select field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelect

`func (o *PlanTableScanRequest) SetSelect(v []string)`

SetSelect sets Select field to given value.

### HasSelect

`func (o *PlanTableScanRequest) HasSelect() bool`

HasSelect returns a boolean if a field has been set.

### GetFilter

`func (o *PlanTableScanRequest) GetFilter() Expression`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *PlanTableScanRequest) GetFilterOk() (*Expression, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *PlanTableScanRequest) SetFilter(v Expression)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *PlanTableScanRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetCaseSensitive

`func (o *PlanTableScanRequest) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *PlanTableScanRequest) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *PlanTableScanRequest) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *PlanTableScanRequest) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.

### GetUseSnapshotSchema

`func (o *PlanTableScanRequest) GetUseSnapshotSchema() bool`

GetUseSnapshotSchema returns the UseSnapshotSchema field if non-nil, zero value otherwise.

### GetUseSnapshotSchemaOk

`func (o *PlanTableScanRequest) GetUseSnapshotSchemaOk() (*bool, bool)`

GetUseSnapshotSchemaOk returns a tuple with the UseSnapshotSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSnapshotSchema

`func (o *PlanTableScanRequest) SetUseSnapshotSchema(v bool)`

SetUseSnapshotSchema sets UseSnapshotSchema field to given value.

### HasUseSnapshotSchema

`func (o *PlanTableScanRequest) HasUseSnapshotSchema() bool`

HasUseSnapshotSchema returns a boolean if a field has been set.

### GetStartSnapshotId

`func (o *PlanTableScanRequest) GetStartSnapshotId() int64`

GetStartSnapshotId returns the StartSnapshotId field if non-nil, zero value otherwise.

### GetStartSnapshotIdOk

`func (o *PlanTableScanRequest) GetStartSnapshotIdOk() (*int64, bool)`

GetStartSnapshotIdOk returns a tuple with the StartSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartSnapshotId

`func (o *PlanTableScanRequest) SetStartSnapshotId(v int64)`

SetStartSnapshotId sets StartSnapshotId field to given value.

### HasStartSnapshotId

`func (o *PlanTableScanRequest) HasStartSnapshotId() bool`

HasStartSnapshotId returns a boolean if a field has been set.

### GetEndSnapshotId

`func (o *PlanTableScanRequest) GetEndSnapshotId() int64`

GetEndSnapshotId returns the EndSnapshotId field if non-nil, zero value otherwise.

### GetEndSnapshotIdOk

`func (o *PlanTableScanRequest) GetEndSnapshotIdOk() (*int64, bool)`

GetEndSnapshotIdOk returns a tuple with the EndSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndSnapshotId

`func (o *PlanTableScanRequest) SetEndSnapshotId(v int64)`

SetEndSnapshotId sets EndSnapshotId field to given value.

### HasEndSnapshotId

`func (o *PlanTableScanRequest) HasEndSnapshotId() bool`

HasEndSnapshotId returns a boolean if a field has been set.

### GetStatsFields

`func (o *PlanTableScanRequest) GetStatsFields() []string`

GetStatsFields returns the StatsFields field if non-nil, zero value otherwise.

### GetStatsFieldsOk

`func (o *PlanTableScanRequest) GetStatsFieldsOk() (*[]string, bool)`

GetStatsFieldsOk returns a tuple with the StatsFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsFields

`func (o *PlanTableScanRequest) SetStatsFields(v []string)`

SetStatsFields sets StatsFields field to given value.

### HasStatsFields

`func (o *PlanTableScanRequest) HasStatsFields() bool`

HasStatsFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


