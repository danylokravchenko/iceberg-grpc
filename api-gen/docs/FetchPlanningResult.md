# FetchPlanningResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteFiles** | Pointer to [**[]DeleteFile**](DeleteFile.md) | Delete files referenced by file scan tasks | [optional] 
**FileScanTasks** | Pointer to [**[]FileScanTask**](FileScanTask.md) |  | [optional] 
**PlanTasks** | Pointer to **[]string** |  | [optional] 
**Status** | [**PlanStatus**](PlanStatus.md) |  | 
**Error** | [**ErrorModel**](ErrorModel.md) |  | 

## Methods

### NewFetchPlanningResult

`func NewFetchPlanningResult(status PlanStatus, error_ ErrorModel, ) *FetchPlanningResult`

NewFetchPlanningResult instantiates a new FetchPlanningResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchPlanningResultWithDefaults

`func NewFetchPlanningResultWithDefaults() *FetchPlanningResult`

NewFetchPlanningResultWithDefaults instantiates a new FetchPlanningResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteFiles

`func (o *FetchPlanningResult) GetDeleteFiles() []DeleteFile`

GetDeleteFiles returns the DeleteFiles field if non-nil, zero value otherwise.

### GetDeleteFilesOk

`func (o *FetchPlanningResult) GetDeleteFilesOk() (*[]DeleteFile, bool)`

GetDeleteFilesOk returns a tuple with the DeleteFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteFiles

`func (o *FetchPlanningResult) SetDeleteFiles(v []DeleteFile)`

SetDeleteFiles sets DeleteFiles field to given value.

### HasDeleteFiles

`func (o *FetchPlanningResult) HasDeleteFiles() bool`

HasDeleteFiles returns a boolean if a field has been set.

### GetFileScanTasks

`func (o *FetchPlanningResult) GetFileScanTasks() []FileScanTask`

GetFileScanTasks returns the FileScanTasks field if non-nil, zero value otherwise.

### GetFileScanTasksOk

`func (o *FetchPlanningResult) GetFileScanTasksOk() (*[]FileScanTask, bool)`

GetFileScanTasksOk returns a tuple with the FileScanTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileScanTasks

`func (o *FetchPlanningResult) SetFileScanTasks(v []FileScanTask)`

SetFileScanTasks sets FileScanTasks field to given value.

### HasFileScanTasks

`func (o *FetchPlanningResult) HasFileScanTasks() bool`

HasFileScanTasks returns a boolean if a field has been set.

### GetPlanTasks

`func (o *FetchPlanningResult) GetPlanTasks() []string`

GetPlanTasks returns the PlanTasks field if non-nil, zero value otherwise.

### GetPlanTasksOk

`func (o *FetchPlanningResult) GetPlanTasksOk() (*[]string, bool)`

GetPlanTasksOk returns a tuple with the PlanTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanTasks

`func (o *FetchPlanningResult) SetPlanTasks(v []string)`

SetPlanTasks sets PlanTasks field to given value.

### HasPlanTasks

`func (o *FetchPlanningResult) HasPlanTasks() bool`

HasPlanTasks returns a boolean if a field has been set.

### GetStatus

`func (o *FetchPlanningResult) GetStatus() PlanStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FetchPlanningResult) GetStatusOk() (*PlanStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FetchPlanningResult) SetStatus(v PlanStatus)`

SetStatus sets Status field to given value.


### GetError

`func (o *FetchPlanningResult) GetError() ErrorModel`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *FetchPlanningResult) GetErrorOk() (*ErrorModel, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *FetchPlanningResult) SetError(v ErrorModel)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


