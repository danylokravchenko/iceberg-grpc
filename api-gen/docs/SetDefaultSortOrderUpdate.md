# SetDefaultSortOrderUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**SortOrderId** | **int32** | Sort order ID to set as the default, or -1 to set last added sort order | 

## Methods

### NewSetDefaultSortOrderUpdate

`func NewSetDefaultSortOrderUpdate(sortOrderId int32, ) *SetDefaultSortOrderUpdate`

NewSetDefaultSortOrderUpdate instantiates a new SetDefaultSortOrderUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetDefaultSortOrderUpdateWithDefaults

`func NewSetDefaultSortOrderUpdateWithDefaults() *SetDefaultSortOrderUpdate`

NewSetDefaultSortOrderUpdateWithDefaults instantiates a new SetDefaultSortOrderUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *SetDefaultSortOrderUpdate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SetDefaultSortOrderUpdate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SetDefaultSortOrderUpdate) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SetDefaultSortOrderUpdate) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetSortOrderId

`func (o *SetDefaultSortOrderUpdate) GetSortOrderId() int32`

GetSortOrderId returns the SortOrderId field if non-nil, zero value otherwise.

### GetSortOrderIdOk

`func (o *SetDefaultSortOrderUpdate) GetSortOrderIdOk() (*int32, bool)`

GetSortOrderIdOk returns a tuple with the SortOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrderId

`func (o *SetDefaultSortOrderUpdate) SetSortOrderId(v int32)`

SetSortOrderId sets SortOrderId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


