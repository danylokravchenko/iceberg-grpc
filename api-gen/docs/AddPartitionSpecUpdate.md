# AddPartitionSpecUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Spec** | [**PartitionSpec**](PartitionSpec.md) |  | 

## Methods

### NewAddPartitionSpecUpdate

`func NewAddPartitionSpecUpdate(spec PartitionSpec, ) *AddPartitionSpecUpdate`

NewAddPartitionSpecUpdate instantiates a new AddPartitionSpecUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPartitionSpecUpdateWithDefaults

`func NewAddPartitionSpecUpdateWithDefaults() *AddPartitionSpecUpdate`

NewAddPartitionSpecUpdateWithDefaults instantiates a new AddPartitionSpecUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AddPartitionSpecUpdate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AddPartitionSpecUpdate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AddPartitionSpecUpdate) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AddPartitionSpecUpdate) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetSpec

`func (o *AddPartitionSpecUpdate) GetSpec() PartitionSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *AddPartitionSpecUpdate) GetSpecOk() (*PartitionSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *AddPartitionSpecUpdate) SetSpec(v PartitionSpec)`

SetSpec sets Spec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


