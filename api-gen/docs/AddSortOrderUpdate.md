# AddSortOrderUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**SortOrder** | [**SortOrder**](SortOrder.md) |  | 

## Methods

### NewAddSortOrderUpdate

`func NewAddSortOrderUpdate(sortOrder SortOrder, ) *AddSortOrderUpdate`

NewAddSortOrderUpdate instantiates a new AddSortOrderUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSortOrderUpdateWithDefaults

`func NewAddSortOrderUpdateWithDefaults() *AddSortOrderUpdate`

NewAddSortOrderUpdateWithDefaults instantiates a new AddSortOrderUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AddSortOrderUpdate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AddSortOrderUpdate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AddSortOrderUpdate) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AddSortOrderUpdate) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetSortOrder

`func (o *AddSortOrderUpdate) GetSortOrder() SortOrder`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *AddSortOrderUpdate) GetSortOrderOk() (*SortOrder, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *AddSortOrderUpdate) SetSortOrder(v SortOrder)`

SetSortOrder sets SortOrder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


