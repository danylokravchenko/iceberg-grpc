# RemovePartitionSpecsUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**SpecIds** | **[]int32** |  | 

## Methods

### NewRemovePartitionSpecsUpdate

`func NewRemovePartitionSpecsUpdate(specIds []int32, ) *RemovePartitionSpecsUpdate`

NewRemovePartitionSpecsUpdate instantiates a new RemovePartitionSpecsUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemovePartitionSpecsUpdateWithDefaults

`func NewRemovePartitionSpecsUpdateWithDefaults() *RemovePartitionSpecsUpdate`

NewRemovePartitionSpecsUpdateWithDefaults instantiates a new RemovePartitionSpecsUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *RemovePartitionSpecsUpdate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RemovePartitionSpecsUpdate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RemovePartitionSpecsUpdate) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *RemovePartitionSpecsUpdate) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetSpecIds

`func (o *RemovePartitionSpecsUpdate) GetSpecIds() []int32`

GetSpecIds returns the SpecIds field if non-nil, zero value otherwise.

### GetSpecIdsOk

`func (o *RemovePartitionSpecsUpdate) GetSpecIdsOk() (*[]int32, bool)`

GetSpecIdsOk returns a tuple with the SpecIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecIds

`func (o *RemovePartitionSpecsUpdate) SetSpecIds(v []int32)`

SetSpecIds sets SpecIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


