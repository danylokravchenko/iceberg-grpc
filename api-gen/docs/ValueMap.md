# ValueMap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | Pointer to **[]int32** | List of integer column ids for each corresponding value | [optional] 
**Values** | Pointer to [**[]PrimitiveTypeValue**](PrimitiveTypeValue.md) | List of primitive type values, matched to &#39;keys&#39; by index | [optional] 

## Methods

### NewValueMap

`func NewValueMap() *ValueMap`

NewValueMap instantiates a new ValueMap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValueMapWithDefaults

`func NewValueMapWithDefaults() *ValueMap`

NewValueMapWithDefaults instantiates a new ValueMap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *ValueMap) GetKeys() []int32`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *ValueMap) GetKeysOk() (*[]int32, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *ValueMap) SetKeys(v []int32)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *ValueMap) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetValues

`func (o *ValueMap) GetValues() []PrimitiveTypeValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ValueMap) GetValuesOk() (*[]PrimitiveTypeValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ValueMap) SetValues(v []PrimitiveTypeValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *ValueMap) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


