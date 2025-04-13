# MetricResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | **string** |  | 
**Value** | **int64** |  | 
**TimeUnit** | **string** |  | 
**Count** | **int64** |  | 
**TotalDuration** | **int64** |  | 

## Methods

### NewMetricResult

`func NewMetricResult(unit string, value int64, timeUnit string, count int64, totalDuration int64, ) *MetricResult`

NewMetricResult instantiates a new MetricResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricResultWithDefaults

`func NewMetricResultWithDefaults() *MetricResult`

NewMetricResultWithDefaults instantiates a new MetricResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *MetricResult) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricResult) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricResult) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetValue

`func (o *MetricResult) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MetricResult) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MetricResult) SetValue(v int64)`

SetValue sets Value field to given value.


### GetTimeUnit

`func (o *MetricResult) GetTimeUnit() string`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *MetricResult) GetTimeUnitOk() (*string, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *MetricResult) SetTimeUnit(v string)`

SetTimeUnit sets TimeUnit field to given value.


### GetCount

`func (o *MetricResult) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *MetricResult) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *MetricResult) SetCount(v int64)`

SetCount sets Count field to given value.


### GetTotalDuration

`func (o *MetricResult) GetTotalDuration() int64`

GetTotalDuration returns the TotalDuration field if non-nil, zero value otherwise.

### GetTotalDurationOk

`func (o *MetricResult) GetTotalDurationOk() (*int64, bool)`

GetTotalDurationOk returns a tuple with the TotalDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDuration

`func (o *MetricResult) SetTotalDuration(v int64)`

SetTotalDuration sets TotalDuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


