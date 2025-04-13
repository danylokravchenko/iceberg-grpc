# ViewHistoryEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionId** | **int32** |  | 
**TimestampMs** | **int64** |  | 

## Methods

### NewViewHistoryEntry

`func NewViewHistoryEntry(versionId int32, timestampMs int64, ) *ViewHistoryEntry`

NewViewHistoryEntry instantiates a new ViewHistoryEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewHistoryEntryWithDefaults

`func NewViewHistoryEntryWithDefaults() *ViewHistoryEntry`

NewViewHistoryEntryWithDefaults instantiates a new ViewHistoryEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionId

`func (o *ViewHistoryEntry) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *ViewHistoryEntry) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *ViewHistoryEntry) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.


### GetTimestampMs

`func (o *ViewHistoryEntry) GetTimestampMs() int64`

GetTimestampMs returns the TimestampMs field if non-nil, zero value otherwise.

### GetTimestampMsOk

`func (o *ViewHistoryEntry) GetTimestampMsOk() (*int64, bool)`

GetTimestampMsOk returns a tuple with the TimestampMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampMs

`func (o *ViewHistoryEntry) SetTimestampMs(v int64)`

SetTimestampMs sets TimestampMs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


