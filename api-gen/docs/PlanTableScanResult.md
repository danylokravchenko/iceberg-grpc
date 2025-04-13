# PlanTableScanResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteFiles** | Pointer to [**[]DeleteFile**](DeleteFile.md) | Delete files referenced by file scan tasks | [optional] 
**FileScanTasks** | Pointer to [**[]FileScanTask**](FileScanTask.md) |  | [optional] 
**PlanTasks** | Pointer to **[]string** |  | [optional] 
**Status** | [**PlanStatus**](PlanStatus.md) |  | 
**PlanId** | Pointer to **string** | ID used to track a planning request | [optional] 
**Error** | [**ErrorModel**](ErrorModel.md) |  | 

## Methods

### NewPlanTableScanResult

`func NewPlanTableScanResult(status PlanStatus, error_ ErrorModel, ) *PlanTableScanResult`

NewPlanTableScanResult instantiates a new PlanTableScanResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanTableScanResultWithDefaults

`func NewPlanTableScanResultWithDefaults() *PlanTableScanResult`

NewPlanTableScanResultWithDefaults instantiates a new PlanTableScanResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteFiles

`func (o *PlanTableScanResult) GetDeleteFiles() []DeleteFile`

GetDeleteFiles returns the DeleteFiles field if non-nil, zero value otherwise.

### GetDeleteFilesOk

`func (o *PlanTableScanResult) GetDeleteFilesOk() (*[]DeleteFile, bool)`

GetDeleteFilesOk returns a tuple with the DeleteFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteFiles

`func (o *PlanTableScanResult) SetDeleteFiles(v []DeleteFile)`

SetDeleteFiles sets DeleteFiles field to given value.

### HasDeleteFiles

`func (o *PlanTableScanResult) HasDeleteFiles() bool`

HasDeleteFiles returns a boolean if a field has been set.

### GetFileScanTasks

`func (o *PlanTableScanResult) GetFileScanTasks() []FileScanTask`

GetFileScanTasks returns the FileScanTasks field if non-nil, zero value otherwise.

### GetFileScanTasksOk

`func (o *PlanTableScanResult) GetFileScanTasksOk() (*[]FileScanTask, bool)`

GetFileScanTasksOk returns a tuple with the FileScanTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileScanTasks

`func (o *PlanTableScanResult) SetFileScanTasks(v []FileScanTask)`

SetFileScanTasks sets FileScanTasks field to given value.

### HasFileScanTasks

`func (o *PlanTableScanResult) HasFileScanTasks() bool`

HasFileScanTasks returns a boolean if a field has been set.

### GetPlanTasks

`func (o *PlanTableScanResult) GetPlanTasks() []string`

GetPlanTasks returns the PlanTasks field if non-nil, zero value otherwise.

### GetPlanTasksOk

`func (o *PlanTableScanResult) GetPlanTasksOk() (*[]string, bool)`

GetPlanTasksOk returns a tuple with the PlanTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanTasks

`func (o *PlanTableScanResult) SetPlanTasks(v []string)`

SetPlanTasks sets PlanTasks field to given value.

### HasPlanTasks

`func (o *PlanTableScanResult) HasPlanTasks() bool`

HasPlanTasks returns a boolean if a field has been set.

### GetStatus

`func (o *PlanTableScanResult) GetStatus() PlanStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlanTableScanResult) GetStatusOk() (*PlanStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PlanTableScanResult) SetStatus(v PlanStatus)`

SetStatus sets Status field to given value.


### GetPlanId

`func (o *PlanTableScanResult) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *PlanTableScanResult) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *PlanTableScanResult) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *PlanTableScanResult) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetError

`func (o *PlanTableScanResult) GetError() ErrorModel`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PlanTableScanResult) GetErrorOk() (*ErrorModel, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PlanTableScanResult) SetError(v ErrorModel)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


