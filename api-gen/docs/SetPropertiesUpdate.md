# SetPropertiesUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Updates** | **map[string]string** |  | 

## Methods

### NewSetPropertiesUpdate

`func NewSetPropertiesUpdate(updates map[string]string, ) *SetPropertiesUpdate`

NewSetPropertiesUpdate instantiates a new SetPropertiesUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetPropertiesUpdateWithDefaults

`func NewSetPropertiesUpdateWithDefaults() *SetPropertiesUpdate`

NewSetPropertiesUpdateWithDefaults instantiates a new SetPropertiesUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *SetPropertiesUpdate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SetPropertiesUpdate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SetPropertiesUpdate) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SetPropertiesUpdate) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetUpdates

`func (o *SetPropertiesUpdate) GetUpdates() map[string]string`

GetUpdates returns the Updates field if non-nil, zero value otherwise.

### GetUpdatesOk

`func (o *SetPropertiesUpdate) GetUpdatesOk() (*map[string]string, bool)`

GetUpdatesOk returns a tuple with the Updates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdates

`func (o *SetPropertiesUpdate) SetUpdates(v map[string]string)`

SetUpdates sets Updates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


