# SetDefaultSpecUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**SpecId** | **int32** | Partition spec ID to set as the default, or -1 to set last added spec | 

## Methods

### NewSetDefaultSpecUpdate

`func NewSetDefaultSpecUpdate(specId int32, ) *SetDefaultSpecUpdate`

NewSetDefaultSpecUpdate instantiates a new SetDefaultSpecUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetDefaultSpecUpdateWithDefaults

`func NewSetDefaultSpecUpdateWithDefaults() *SetDefaultSpecUpdate`

NewSetDefaultSpecUpdateWithDefaults instantiates a new SetDefaultSpecUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *SetDefaultSpecUpdate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SetDefaultSpecUpdate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SetDefaultSpecUpdate) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SetDefaultSpecUpdate) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetSpecId

`func (o *SetDefaultSpecUpdate) GetSpecId() int32`

GetSpecId returns the SpecId field if non-nil, zero value otherwise.

### GetSpecIdOk

`func (o *SetDefaultSpecUpdate) GetSpecIdOk() (*int32, bool)`

GetSpecIdOk returns a tuple with the SpecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecId

`func (o *SetDefaultSpecUpdate) SetSpecId(v int32)`

SetSpecId sets SpecId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


