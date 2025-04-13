# SetCurrentViewVersionUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**ViewVersionId** | **int32** | The view version id to set as current, or -1 to set last added view version id | 

## Methods

### NewSetCurrentViewVersionUpdate

`func NewSetCurrentViewVersionUpdate(viewVersionId int32, ) *SetCurrentViewVersionUpdate`

NewSetCurrentViewVersionUpdate instantiates a new SetCurrentViewVersionUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetCurrentViewVersionUpdateWithDefaults

`func NewSetCurrentViewVersionUpdateWithDefaults() *SetCurrentViewVersionUpdate`

NewSetCurrentViewVersionUpdateWithDefaults instantiates a new SetCurrentViewVersionUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *SetCurrentViewVersionUpdate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SetCurrentViewVersionUpdate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SetCurrentViewVersionUpdate) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SetCurrentViewVersionUpdate) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetViewVersionId

`func (o *SetCurrentViewVersionUpdate) GetViewVersionId() int32`

GetViewVersionId returns the ViewVersionId field if non-nil, zero value otherwise.

### GetViewVersionIdOk

`func (o *SetCurrentViewVersionUpdate) GetViewVersionIdOk() (*int32, bool)`

GetViewVersionIdOk returns a tuple with the ViewVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewVersionId

`func (o *SetCurrentViewVersionUpdate) SetViewVersionId(v int32)`

SetViewVersionId sets ViewVersionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


