# FetchScanTasksRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanTask** | **string** | An opaque string provided by the REST server that represents a unit of work to produce file scan tasks for scan planning. This allows clients to fetch tasks across multiple requests to accommodate large result sets. | 

## Methods

### NewFetchScanTasksRequest

`func NewFetchScanTasksRequest(planTask string, ) *FetchScanTasksRequest`

NewFetchScanTasksRequest instantiates a new FetchScanTasksRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchScanTasksRequestWithDefaults

`func NewFetchScanTasksRequestWithDefaults() *FetchScanTasksRequest`

NewFetchScanTasksRequestWithDefaults instantiates a new FetchScanTasksRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanTask

`func (o *FetchScanTasksRequest) GetPlanTask() string`

GetPlanTask returns the PlanTask field if non-nil, zero value otherwise.

### GetPlanTaskOk

`func (o *FetchScanTasksRequest) GetPlanTaskOk() (*string, bool)`

GetPlanTaskOk returns a tuple with the PlanTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanTask

`func (o *FetchScanTasksRequest) SetPlanTask(v string)`

SetPlanTask sets PlanTask field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


