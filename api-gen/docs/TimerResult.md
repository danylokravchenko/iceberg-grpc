# TimerResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeUnit** | **string** |  | 
**Count** | **int64** |  | 
**TotalDuration** | **int64** |  | 

## Methods

### NewTimerResult

`func NewTimerResult(timeUnit string, count int64, totalDuration int64, ) *TimerResult`

NewTimerResult instantiates a new TimerResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimerResultWithDefaults

`func NewTimerResultWithDefaults() *TimerResult`

NewTimerResultWithDefaults instantiates a new TimerResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeUnit

`func (o *TimerResult) GetTimeUnit() string`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *TimerResult) GetTimeUnitOk() (*string, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *TimerResult) SetTimeUnit(v string)`

SetTimeUnit sets TimeUnit field to given value.


### GetCount

`func (o *TimerResult) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TimerResult) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TimerResult) SetCount(v int64)`

SetCount sets Count field to given value.


### GetTotalDuration

`func (o *TimerResult) GetTotalDuration() int64`

GetTotalDuration returns the TotalDuration field if non-nil, zero value otherwise.

### GetTotalDurationOk

`func (o *TimerResult) GetTotalDurationOk() (*int64, bool)`

GetTotalDurationOk returns a tuple with the TotalDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDuration

`func (o *TimerResult) SetTotalDuration(v int64)`

SetTotalDuration sets TotalDuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


