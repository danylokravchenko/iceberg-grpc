# DataFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** |  | 
**ColumnSizes** | Pointer to [**CountMap**](CountMap.md) | Map of column id to total count, including null and NaN | [optional] 
**ValueCounts** | Pointer to [**CountMap**](CountMap.md) | Map of column id to null value count | [optional] 
**NullValueCounts** | Pointer to [**CountMap**](CountMap.md) | Map of column id to null value count | [optional] 
**NanValueCounts** | Pointer to [**CountMap**](CountMap.md) | Map of column id to number of NaN values in the column | [optional] 
**LowerBounds** | Pointer to [**ValueMap**](ValueMap.md) | Map of column id to lower bound primitive type values | [optional] 
**UpperBounds** | Pointer to [**ValueMap**](ValueMap.md) | Map of column id to upper bound primitive type values | [optional] 

## Methods

### NewDataFile

`func NewDataFile(content string, ) *DataFile`

NewDataFile instantiates a new DataFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataFileWithDefaults

`func NewDataFileWithDefaults() *DataFile`

NewDataFileWithDefaults instantiates a new DataFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *DataFile) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *DataFile) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *DataFile) SetContent(v string)`

SetContent sets Content field to given value.


### GetColumnSizes

`func (o *DataFile) GetColumnSizes() CountMap`

GetColumnSizes returns the ColumnSizes field if non-nil, zero value otherwise.

### GetColumnSizesOk

`func (o *DataFile) GetColumnSizesOk() (*CountMap, bool)`

GetColumnSizesOk returns a tuple with the ColumnSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnSizes

`func (o *DataFile) SetColumnSizes(v CountMap)`

SetColumnSizes sets ColumnSizes field to given value.

### HasColumnSizes

`func (o *DataFile) HasColumnSizes() bool`

HasColumnSizes returns a boolean if a field has been set.

### GetValueCounts

`func (o *DataFile) GetValueCounts() CountMap`

GetValueCounts returns the ValueCounts field if non-nil, zero value otherwise.

### GetValueCountsOk

`func (o *DataFile) GetValueCountsOk() (*CountMap, bool)`

GetValueCountsOk returns a tuple with the ValueCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueCounts

`func (o *DataFile) SetValueCounts(v CountMap)`

SetValueCounts sets ValueCounts field to given value.

### HasValueCounts

`func (o *DataFile) HasValueCounts() bool`

HasValueCounts returns a boolean if a field has been set.

### GetNullValueCounts

`func (o *DataFile) GetNullValueCounts() CountMap`

GetNullValueCounts returns the NullValueCounts field if non-nil, zero value otherwise.

### GetNullValueCountsOk

`func (o *DataFile) GetNullValueCountsOk() (*CountMap, bool)`

GetNullValueCountsOk returns a tuple with the NullValueCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullValueCounts

`func (o *DataFile) SetNullValueCounts(v CountMap)`

SetNullValueCounts sets NullValueCounts field to given value.

### HasNullValueCounts

`func (o *DataFile) HasNullValueCounts() bool`

HasNullValueCounts returns a boolean if a field has been set.

### GetNanValueCounts

`func (o *DataFile) GetNanValueCounts() CountMap`

GetNanValueCounts returns the NanValueCounts field if non-nil, zero value otherwise.

### GetNanValueCountsOk

`func (o *DataFile) GetNanValueCountsOk() (*CountMap, bool)`

GetNanValueCountsOk returns a tuple with the NanValueCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNanValueCounts

`func (o *DataFile) SetNanValueCounts(v CountMap)`

SetNanValueCounts sets NanValueCounts field to given value.

### HasNanValueCounts

`func (o *DataFile) HasNanValueCounts() bool`

HasNanValueCounts returns a boolean if a field has been set.

### GetLowerBounds

`func (o *DataFile) GetLowerBounds() ValueMap`

GetLowerBounds returns the LowerBounds field if non-nil, zero value otherwise.

### GetLowerBoundsOk

`func (o *DataFile) GetLowerBoundsOk() (*ValueMap, bool)`

GetLowerBoundsOk returns a tuple with the LowerBounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBounds

`func (o *DataFile) SetLowerBounds(v ValueMap)`

SetLowerBounds sets LowerBounds field to given value.

### HasLowerBounds

`func (o *DataFile) HasLowerBounds() bool`

HasLowerBounds returns a boolean if a field has been set.

### GetUpperBounds

`func (o *DataFile) GetUpperBounds() ValueMap`

GetUpperBounds returns the UpperBounds field if non-nil, zero value otherwise.

### GetUpperBoundsOk

`func (o *DataFile) GetUpperBoundsOk() (*ValueMap, bool)`

GetUpperBoundsOk returns a tuple with the UpperBounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBounds

`func (o *DataFile) SetUpperBounds(v ValueMap)`

SetUpperBounds sets UpperBounds field to given value.

### HasUpperBounds

`func (o *DataFile) HasUpperBounds() bool`

HasUpperBounds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


