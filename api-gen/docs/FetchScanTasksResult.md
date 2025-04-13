# FetchScanTasksResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteFiles** | Pointer to [**[]DeleteFile**](DeleteFile.md) | Delete files referenced by file scan tasks | [optional] 
**FileScanTasks** | Pointer to [**[]FileScanTask**](FileScanTask.md) |  | [optional] 
**PlanTasks** | Pointer to **[]string** |  | [optional] 

## Methods

### NewFetchScanTasksResult

`func NewFetchScanTasksResult() *FetchScanTasksResult`

NewFetchScanTasksResult instantiates a new FetchScanTasksResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchScanTasksResultWithDefaults

`func NewFetchScanTasksResultWithDefaults() *FetchScanTasksResult`

NewFetchScanTasksResultWithDefaults instantiates a new FetchScanTasksResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteFiles

`func (o *FetchScanTasksResult) GetDeleteFiles() []DeleteFile`

GetDeleteFiles returns the DeleteFiles field if non-nil, zero value otherwise.

### GetDeleteFilesOk

`func (o *FetchScanTasksResult) GetDeleteFilesOk() (*[]DeleteFile, bool)`

GetDeleteFilesOk returns a tuple with the DeleteFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteFiles

`func (o *FetchScanTasksResult) SetDeleteFiles(v []DeleteFile)`

SetDeleteFiles sets DeleteFiles field to given value.

### HasDeleteFiles

`func (o *FetchScanTasksResult) HasDeleteFiles() bool`

HasDeleteFiles returns a boolean if a field has been set.

### GetFileScanTasks

`func (o *FetchScanTasksResult) GetFileScanTasks() []FileScanTask`

GetFileScanTasks returns the FileScanTasks field if non-nil, zero value otherwise.

### GetFileScanTasksOk

`func (o *FetchScanTasksResult) GetFileScanTasksOk() (*[]FileScanTask, bool)`

GetFileScanTasksOk returns a tuple with the FileScanTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileScanTasks

`func (o *FetchScanTasksResult) SetFileScanTasks(v []FileScanTask)`

SetFileScanTasks sets FileScanTasks field to given value.

### HasFileScanTasks

`func (o *FetchScanTasksResult) HasFileScanTasks() bool`

HasFileScanTasks returns a boolean if a field has been set.

### GetPlanTasks

`func (o *FetchScanTasksResult) GetPlanTasks() []string`

GetPlanTasks returns the PlanTasks field if non-nil, zero value otherwise.

### GetPlanTasksOk

`func (o *FetchScanTasksResult) GetPlanTasksOk() (*[]string, bool)`

GetPlanTasksOk returns a tuple with the PlanTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanTasks

`func (o *FetchScanTasksResult) SetPlanTasks(v []string)`

SetPlanTasks sets PlanTasks field to given value.

### HasPlanTasks

`func (o *FetchScanTasksResult) HasPlanTasks() bool`

HasPlanTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


