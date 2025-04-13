# CountMap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | Pointer to **[]int32** | List of integer column ids for each corresponding value | [optional] 
**Values** | Pointer to **[]int64** | List of Long values, matched to &#39;keys&#39; by index | [optional] 

## Methods

### NewCountMap

`func NewCountMap() *CountMap`

NewCountMap instantiates a new CountMap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountMapWithDefaults

`func NewCountMapWithDefaults() *CountMap`

NewCountMapWithDefaults instantiates a new CountMap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *CountMap) GetKeys() []int32`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *CountMap) GetKeysOk() (*[]int32, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *CountMap) SetKeys(v []int32)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *CountMap) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetValues

`func (o *CountMap) GetValues() []int64`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CountMap) GetValuesOk() (*[]int64, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CountMap) SetValues(v []int64)`

SetValues sets Values field to given value.

### HasValues

`func (o *CountMap) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


