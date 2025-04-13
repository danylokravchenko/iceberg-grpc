# CompletedPlanningResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteFiles** | Pointer to [**[]DeleteFile**](DeleteFile.md) | Delete files referenced by file scan tasks | [optional] 
**FileScanTasks** | Pointer to [**[]FileScanTask**](FileScanTask.md) |  | [optional] 
**PlanTasks** | Pointer to **[]string** |  | [optional] 
**Status** | [**PlanStatus**](PlanStatus.md) |  | 

## Methods

### NewCompletedPlanningResult

`func NewCompletedPlanningResult(status PlanStatus, ) *CompletedPlanningResult`

NewCompletedPlanningResult instantiates a new CompletedPlanningResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompletedPlanningResultWithDefaults

`func NewCompletedPlanningResultWithDefaults() *CompletedPlanningResult`

NewCompletedPlanningResultWithDefaults instantiates a new CompletedPlanningResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteFiles

`func (o *CompletedPlanningResult) GetDeleteFiles() []DeleteFile`

GetDeleteFiles returns the DeleteFiles field if non-nil, zero value otherwise.

### GetDeleteFilesOk

`func (o *CompletedPlanningResult) GetDeleteFilesOk() (*[]DeleteFile, bool)`

GetDeleteFilesOk returns a tuple with the DeleteFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteFiles

`func (o *CompletedPlanningResult) SetDeleteFiles(v []DeleteFile)`

SetDeleteFiles sets DeleteFiles field to given value.

### HasDeleteFiles

`func (o *CompletedPlanningResult) HasDeleteFiles() bool`

HasDeleteFiles returns a boolean if a field has been set.

### GetFileScanTasks

`func (o *CompletedPlanningResult) GetFileScanTasks() []FileScanTask`

GetFileScanTasks returns the FileScanTasks field if non-nil, zero value otherwise.

### GetFileScanTasksOk

`func (o *CompletedPlanningResult) GetFileScanTasksOk() (*[]FileScanTask, bool)`

GetFileScanTasksOk returns a tuple with the FileScanTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileScanTasks

`func (o *CompletedPlanningResult) SetFileScanTasks(v []FileScanTask)`

SetFileScanTasks sets FileScanTasks field to given value.

### HasFileScanTasks

`func (o *CompletedPlanningResult) HasFileScanTasks() bool`

HasFileScanTasks returns a boolean if a field has been set.

### GetPlanTasks

`func (o *CompletedPlanningResult) GetPlanTasks() []string`

GetPlanTasks returns the PlanTasks field if non-nil, zero value otherwise.

### GetPlanTasksOk

`func (o *CompletedPlanningResult) GetPlanTasksOk() (*[]string, bool)`

GetPlanTasksOk returns a tuple with the PlanTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanTasks

`func (o *CompletedPlanningResult) SetPlanTasks(v []string)`

SetPlanTasks sets PlanTasks field to given value.

### HasPlanTasks

`func (o *CompletedPlanningResult) HasPlanTasks() bool`

HasPlanTasks returns a boolean if a field has been set.

### GetStatus

`func (o *CompletedPlanningResult) GetStatus() PlanStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CompletedPlanningResult) GetStatusOk() (*PlanStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CompletedPlanningResult) SetStatus(v PlanStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


