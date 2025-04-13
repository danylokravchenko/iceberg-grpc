# CompletedPlanningWithIDResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteFiles** | Pointer to [**[]DeleteFile**](DeleteFile.md) | Delete files referenced by file scan tasks | [optional] 
**FileScanTasks** | Pointer to [**[]FileScanTask**](FileScanTask.md) |  | [optional] 
**PlanTasks** | Pointer to **[]string** |  | [optional] 
**Status** | [**PlanStatus**](PlanStatus.md) |  | 
**PlanId** | Pointer to **string** | ID used to track a planning request | [optional] 

## Methods

### NewCompletedPlanningWithIDResult

`func NewCompletedPlanningWithIDResult(status PlanStatus, ) *CompletedPlanningWithIDResult`

NewCompletedPlanningWithIDResult instantiates a new CompletedPlanningWithIDResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompletedPlanningWithIDResultWithDefaults

`func NewCompletedPlanningWithIDResultWithDefaults() *CompletedPlanningWithIDResult`

NewCompletedPlanningWithIDResultWithDefaults instantiates a new CompletedPlanningWithIDResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteFiles

`func (o *CompletedPlanningWithIDResult) GetDeleteFiles() []DeleteFile`

GetDeleteFiles returns the DeleteFiles field if non-nil, zero value otherwise.

### GetDeleteFilesOk

`func (o *CompletedPlanningWithIDResult) GetDeleteFilesOk() (*[]DeleteFile, bool)`

GetDeleteFilesOk returns a tuple with the DeleteFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteFiles

`func (o *CompletedPlanningWithIDResult) SetDeleteFiles(v []DeleteFile)`

SetDeleteFiles sets DeleteFiles field to given value.

### HasDeleteFiles

`func (o *CompletedPlanningWithIDResult) HasDeleteFiles() bool`

HasDeleteFiles returns a boolean if a field has been set.

### GetFileScanTasks

`func (o *CompletedPlanningWithIDResult) GetFileScanTasks() []FileScanTask`

GetFileScanTasks returns the FileScanTasks field if non-nil, zero value otherwise.

### GetFileScanTasksOk

`func (o *CompletedPlanningWithIDResult) GetFileScanTasksOk() (*[]FileScanTask, bool)`

GetFileScanTasksOk returns a tuple with the FileScanTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileScanTasks

`func (o *CompletedPlanningWithIDResult) SetFileScanTasks(v []FileScanTask)`

SetFileScanTasks sets FileScanTasks field to given value.

### HasFileScanTasks

`func (o *CompletedPlanningWithIDResult) HasFileScanTasks() bool`

HasFileScanTasks returns a boolean if a field has been set.

### GetPlanTasks

`func (o *CompletedPlanningWithIDResult) GetPlanTasks() []string`

GetPlanTasks returns the PlanTasks field if non-nil, zero value otherwise.

### GetPlanTasksOk

`func (o *CompletedPlanningWithIDResult) GetPlanTasksOk() (*[]string, bool)`

GetPlanTasksOk returns a tuple with the PlanTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanTasks

`func (o *CompletedPlanningWithIDResult) SetPlanTasks(v []string)`

SetPlanTasks sets PlanTasks field to given value.

### HasPlanTasks

`func (o *CompletedPlanningWithIDResult) HasPlanTasks() bool`

HasPlanTasks returns a boolean if a field has been set.

### GetStatus

`func (o *CompletedPlanningWithIDResult) GetStatus() PlanStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CompletedPlanningWithIDResult) GetStatusOk() (*PlanStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CompletedPlanningWithIDResult) SetStatus(v PlanStatus)`

SetStatus sets Status field to given value.


### GetPlanId

`func (o *CompletedPlanningWithIDResult) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *CompletedPlanningWithIDResult) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *CompletedPlanningWithIDResult) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *CompletedPlanningWithIDResult) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


