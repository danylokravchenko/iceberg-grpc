# AsyncPlanningResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**PlanStatus**](PlanStatus.md) |  | 
**PlanId** | Pointer to **string** | ID used to track a planning request | [optional] 

## Methods

### NewAsyncPlanningResult

`func NewAsyncPlanningResult(status PlanStatus, ) *AsyncPlanningResult`

NewAsyncPlanningResult instantiates a new AsyncPlanningResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncPlanningResultWithDefaults

`func NewAsyncPlanningResultWithDefaults() *AsyncPlanningResult`

NewAsyncPlanningResultWithDefaults instantiates a new AsyncPlanningResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AsyncPlanningResult) GetStatus() PlanStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AsyncPlanningResult) GetStatusOk() (*PlanStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AsyncPlanningResult) SetStatus(v PlanStatus)`

SetStatus sets Status field to given value.


### GetPlanId

`func (o *AsyncPlanningResult) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *AsyncPlanningResult) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *AsyncPlanningResult) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *AsyncPlanningResult) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


