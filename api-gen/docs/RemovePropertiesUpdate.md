# RemovePropertiesUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Removals** | **[]string** |  | 

## Methods

### NewRemovePropertiesUpdate

`func NewRemovePropertiesUpdate(removals []string, ) *RemovePropertiesUpdate`

NewRemovePropertiesUpdate instantiates a new RemovePropertiesUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemovePropertiesUpdateWithDefaults

`func NewRemovePropertiesUpdateWithDefaults() *RemovePropertiesUpdate`

NewRemovePropertiesUpdateWithDefaults instantiates a new RemovePropertiesUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *RemovePropertiesUpdate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RemovePropertiesUpdate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RemovePropertiesUpdate) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *RemovePropertiesUpdate) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetRemovals

`func (o *RemovePropertiesUpdate) GetRemovals() []string`

GetRemovals returns the Removals field if non-nil, zero value otherwise.

### GetRemovalsOk

`func (o *RemovePropertiesUpdate) GetRemovalsOk() (*[]string, bool)`

GetRemovalsOk returns a tuple with the Removals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovals

`func (o *RemovePropertiesUpdate) SetRemovals(v []string)`

SetRemovals sets Removals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


