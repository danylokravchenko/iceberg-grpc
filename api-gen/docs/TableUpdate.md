# TableUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**FormatVersion** | **int32** |  | 
**Schema** | [**Schema**](Schema.md) |  | 
**LastColumnId** | Pointer to **int32** | This optional field is **DEPRECATED for REMOVAL** since it more safe to handle this internally, and shouldn&#39;t be exposed to the clients. The highest assigned column ID for the table. This is used to ensure columns are always assigned an unused ID when evolving schemas. When omitted, it will be computed on the server side. | [optional] 
**SchemaId** | **int32** | Schema ID to set as current, or -1 to set last added schema | 
**Spec** | [**PartitionSpec**](PartitionSpec.md) |  | 
**SpecId** | **int32** | Partition spec ID to set as the default, or -1 to set last added spec | 
**SortOrder** | [**SortOrder**](SortOrder.md) |  | 
**SortOrderId** | **int32** | Sort order ID to set as the default, or -1 to set last added sort order | 
**Snapshot** | [**Snapshot**](Snapshot.md) |  | 
**RefName** | **string** |  | 
**Type** | **string** |  | 
**SnapshotId** | **int64** |  | 
**MaxRefAgeMs** | Pointer to **int64** |  | [optional] 
**MaxSnapshotAgeMs** | Pointer to **int64** |  | [optional] 
**MinSnapshotsToKeep** | Pointer to **int32** |  | [optional] 
**SnapshotIds** | **[]int64** |  | 
**Location** | **string** |  | 
**Updates** | **map[string]string** |  | 
**Removals** | **[]string** |  | 
**Statistics** | [**StatisticsFile**](StatisticsFile.md) |  | 
**SpecIds** | **[]int32** |  | 
**SchemaIds** | **[]int32** |  | 

## Methods

### NewTableUpdate

`func NewTableUpdate(action string, formatVersion int32, schema Schema, schemaId int32, spec PartitionSpec, specId int32, sortOrder SortOrder, sortOrderId int32, snapshot Snapshot, refName string, type_ string, snapshotId int64, snapshotIds []int64, location string, updates map[string]string, removals []string, statistics StatisticsFile, specIds []int32, schemaIds []int32, ) *TableUpdate`

NewTableUpdate instantiates a new TableUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableUpdateWithDefaults

`func NewTableUpdateWithDefaults() *TableUpdate`

NewTableUpdateWithDefaults instantiates a new TableUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *TableUpdate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TableUpdate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TableUpdate) SetAction(v string)`

SetAction sets Action field to given value.


### GetFormatVersion

`func (o *TableUpdate) GetFormatVersion() int32`

GetFormatVersion returns the FormatVersion field if non-nil, zero value otherwise.

### GetFormatVersionOk

`func (o *TableUpdate) GetFormatVersionOk() (*int32, bool)`

GetFormatVersionOk returns a tuple with the FormatVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatVersion

`func (o *TableUpdate) SetFormatVersion(v int32)`

SetFormatVersion sets FormatVersion field to given value.


### GetSchema

`func (o *TableUpdate) GetSchema() Schema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *TableUpdate) GetSchemaOk() (*Schema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *TableUpdate) SetSchema(v Schema)`

SetSchema sets Schema field to given value.


### GetLastColumnId

`func (o *TableUpdate) GetLastColumnId() int32`

GetLastColumnId returns the LastColumnId field if non-nil, zero value otherwise.

### GetLastColumnIdOk

`func (o *TableUpdate) GetLastColumnIdOk() (*int32, bool)`

GetLastColumnIdOk returns a tuple with the LastColumnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastColumnId

`func (o *TableUpdate) SetLastColumnId(v int32)`

SetLastColumnId sets LastColumnId field to given value.

### HasLastColumnId

`func (o *TableUpdate) HasLastColumnId() bool`

HasLastColumnId returns a boolean if a field has been set.

### GetSchemaId

`func (o *TableUpdate) GetSchemaId() int32`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *TableUpdate) GetSchemaIdOk() (*int32, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *TableUpdate) SetSchemaId(v int32)`

SetSchemaId sets SchemaId field to given value.


### GetSpec

`func (o *TableUpdate) GetSpec() PartitionSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *TableUpdate) GetSpecOk() (*PartitionSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *TableUpdate) SetSpec(v PartitionSpec)`

SetSpec sets Spec field to given value.


### GetSpecId

`func (o *TableUpdate) GetSpecId() int32`

GetSpecId returns the SpecId field if non-nil, zero value otherwise.

### GetSpecIdOk

`func (o *TableUpdate) GetSpecIdOk() (*int32, bool)`

GetSpecIdOk returns a tuple with the SpecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecId

`func (o *TableUpdate) SetSpecId(v int32)`

SetSpecId sets SpecId field to given value.


### GetSortOrder

`func (o *TableUpdate) GetSortOrder() SortOrder`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *TableUpdate) GetSortOrderOk() (*SortOrder, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *TableUpdate) SetSortOrder(v SortOrder)`

SetSortOrder sets SortOrder field to given value.


### GetSortOrderId

`func (o *TableUpdate) GetSortOrderId() int32`

GetSortOrderId returns the SortOrderId field if non-nil, zero value otherwise.

### GetSortOrderIdOk

`func (o *TableUpdate) GetSortOrderIdOk() (*int32, bool)`

GetSortOrderIdOk returns a tuple with the SortOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrderId

`func (o *TableUpdate) SetSortOrderId(v int32)`

SetSortOrderId sets SortOrderId field to given value.


### GetSnapshot

`func (o *TableUpdate) GetSnapshot() Snapshot`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *TableUpdate) GetSnapshotOk() (*Snapshot, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *TableUpdate) SetSnapshot(v Snapshot)`

SetSnapshot sets Snapshot field to given value.


### GetRefName

`func (o *TableUpdate) GetRefName() string`

GetRefName returns the RefName field if non-nil, zero value otherwise.

### GetRefNameOk

`func (o *TableUpdate) GetRefNameOk() (*string, bool)`

GetRefNameOk returns a tuple with the RefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefName

`func (o *TableUpdate) SetRefName(v string)`

SetRefName sets RefName field to given value.


### GetType

`func (o *TableUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TableUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TableUpdate) SetType(v string)`

SetType sets Type field to given value.


### GetSnapshotId

`func (o *TableUpdate) GetSnapshotId() int64`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *TableUpdate) GetSnapshotIdOk() (*int64, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *TableUpdate) SetSnapshotId(v int64)`

SetSnapshotId sets SnapshotId field to given value.


### GetMaxRefAgeMs

`func (o *TableUpdate) GetMaxRefAgeMs() int64`

GetMaxRefAgeMs returns the MaxRefAgeMs field if non-nil, zero value otherwise.

### GetMaxRefAgeMsOk

`func (o *TableUpdate) GetMaxRefAgeMsOk() (*int64, bool)`

GetMaxRefAgeMsOk returns a tuple with the MaxRefAgeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRefAgeMs

`func (o *TableUpdate) SetMaxRefAgeMs(v int64)`

SetMaxRefAgeMs sets MaxRefAgeMs field to given value.

### HasMaxRefAgeMs

`func (o *TableUpdate) HasMaxRefAgeMs() bool`

HasMaxRefAgeMs returns a boolean if a field has been set.

### GetMaxSnapshotAgeMs

`func (o *TableUpdate) GetMaxSnapshotAgeMs() int64`

GetMaxSnapshotAgeMs returns the MaxSnapshotAgeMs field if non-nil, zero value otherwise.

### GetMaxSnapshotAgeMsOk

`func (o *TableUpdate) GetMaxSnapshotAgeMsOk() (*int64, bool)`

GetMaxSnapshotAgeMsOk returns a tuple with the MaxSnapshotAgeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSnapshotAgeMs

`func (o *TableUpdate) SetMaxSnapshotAgeMs(v int64)`

SetMaxSnapshotAgeMs sets MaxSnapshotAgeMs field to given value.

### HasMaxSnapshotAgeMs

`func (o *TableUpdate) HasMaxSnapshotAgeMs() bool`

HasMaxSnapshotAgeMs returns a boolean if a field has been set.

### GetMinSnapshotsToKeep

`func (o *TableUpdate) GetMinSnapshotsToKeep() int32`

GetMinSnapshotsToKeep returns the MinSnapshotsToKeep field if non-nil, zero value otherwise.

### GetMinSnapshotsToKeepOk

`func (o *TableUpdate) GetMinSnapshotsToKeepOk() (*int32, bool)`

GetMinSnapshotsToKeepOk returns a tuple with the MinSnapshotsToKeep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSnapshotsToKeep

`func (o *TableUpdate) SetMinSnapshotsToKeep(v int32)`

SetMinSnapshotsToKeep sets MinSnapshotsToKeep field to given value.

### HasMinSnapshotsToKeep

`func (o *TableUpdate) HasMinSnapshotsToKeep() bool`

HasMinSnapshotsToKeep returns a boolean if a field has been set.

### GetSnapshotIds

`func (o *TableUpdate) GetSnapshotIds() []int64`

GetSnapshotIds returns the SnapshotIds field if non-nil, zero value otherwise.

### GetSnapshotIdsOk

`func (o *TableUpdate) GetSnapshotIdsOk() (*[]int64, bool)`

GetSnapshotIdsOk returns a tuple with the SnapshotIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotIds

`func (o *TableUpdate) SetSnapshotIds(v []int64)`

SetSnapshotIds sets SnapshotIds field to given value.


### GetLocation

`func (o *TableUpdate) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *TableUpdate) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *TableUpdate) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetUpdates

`func (o *TableUpdate) GetUpdates() map[string]string`

GetUpdates returns the Updates field if non-nil, zero value otherwise.

### GetUpdatesOk

`func (o *TableUpdate) GetUpdatesOk() (*map[string]string, bool)`

GetUpdatesOk returns a tuple with the Updates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdates

`func (o *TableUpdate) SetUpdates(v map[string]string)`

SetUpdates sets Updates field to given value.


### GetRemovals

`func (o *TableUpdate) GetRemovals() []string`

GetRemovals returns the Removals field if non-nil, zero value otherwise.

### GetRemovalsOk

`func (o *TableUpdate) GetRemovalsOk() (*[]string, bool)`

GetRemovalsOk returns a tuple with the Removals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovals

`func (o *TableUpdate) SetRemovals(v []string)`

SetRemovals sets Removals field to given value.


### GetStatistics

`func (o *TableUpdate) GetStatistics() StatisticsFile`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *TableUpdate) GetStatisticsOk() (*StatisticsFile, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *TableUpdate) SetStatistics(v StatisticsFile)`

SetStatistics sets Statistics field to given value.


### GetSpecIds

`func (o *TableUpdate) GetSpecIds() []int32`

GetSpecIds returns the SpecIds field if non-nil, zero value otherwise.

### GetSpecIdsOk

`func (o *TableUpdate) GetSpecIdsOk() (*[]int32, bool)`

GetSpecIdsOk returns a tuple with the SpecIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecIds

`func (o *TableUpdate) SetSpecIds(v []int32)`

SetSpecIds sets SpecIds field to given value.


### GetSchemaIds

`func (o *TableUpdate) GetSchemaIds() []int32`

GetSchemaIds returns the SchemaIds field if non-nil, zero value otherwise.

### GetSchemaIdsOk

`func (o *TableUpdate) GetSchemaIdsOk() (*[]int32, bool)`

GetSchemaIdsOk returns a tuple with the SchemaIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaIds

`func (o *TableUpdate) SetSchemaIds(v []int32)`

SetSchemaIds sets SchemaIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


