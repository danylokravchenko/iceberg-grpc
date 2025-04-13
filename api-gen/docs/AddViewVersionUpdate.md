# AddViewVersionUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**ViewVersion** | [**ViewVersion**](ViewVersion.md) |  | 

## Methods

### NewAddViewVersionUpdate

`func NewAddViewVersionUpdate(viewVersion ViewVersion, ) *AddViewVersionUpdate`

NewAddViewVersionUpdate instantiates a new AddViewVersionUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddViewVersionUpdateWithDefaults

`func NewAddViewVersionUpdateWithDefaults() *AddViewVersionUpdate`

NewAddViewVersionUpdateWithDefaults instantiates a new AddViewVersionUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AddViewVersionUpdate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AddViewVersionUpdate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AddViewVersionUpdate) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AddViewVersionUpdate) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetViewVersion

`func (o *AddViewVersionUpdate) GetViewVersion() ViewVersion`

GetViewVersion returns the ViewVersion field if non-nil, zero value otherwise.

### GetViewVersionOk

`func (o *AddViewVersionUpdate) GetViewVersionOk() (*ViewVersion, bool)`

GetViewVersionOk returns a tuple with the ViewVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewVersion

`func (o *AddViewVersionUpdate) SetViewVersion(v ViewVersion)`

SetViewVersion sets ViewVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


