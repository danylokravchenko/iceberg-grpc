# FailedPlanningResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | [**ErrorModel**](ErrorModel.md) |  | 
**Status** | [**PlanStatus**](PlanStatus.md) |  | 

## Methods

### NewFailedPlanningResult

`func NewFailedPlanningResult(error_ ErrorModel, status PlanStatus, ) *FailedPlanningResult`

NewFailedPlanningResult instantiates a new FailedPlanningResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailedPlanningResultWithDefaults

`func NewFailedPlanningResultWithDefaults() *FailedPlanningResult`

NewFailedPlanningResultWithDefaults instantiates a new FailedPlanningResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *FailedPlanningResult) GetError() ErrorModel`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *FailedPlanningResult) GetErrorOk() (*ErrorModel, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *FailedPlanningResult) SetError(v ErrorModel)`

SetError sets Error field to given value.


### GetStatus

`func (o *FailedPlanningResult) GetStatus() PlanStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FailedPlanningResult) GetStatusOk() (*PlanStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FailedPlanningResult) SetStatus(v PlanStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


