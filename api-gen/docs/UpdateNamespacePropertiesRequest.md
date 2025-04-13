# UpdateNamespacePropertiesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Removals** | Pointer to **[]string** |  | [optional] 
**Updates** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewUpdateNamespacePropertiesRequest

`func NewUpdateNamespacePropertiesRequest() *UpdateNamespacePropertiesRequest`

NewUpdateNamespacePropertiesRequest instantiates a new UpdateNamespacePropertiesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNamespacePropertiesRequestWithDefaults

`func NewUpdateNamespacePropertiesRequestWithDefaults() *UpdateNamespacePropertiesRequest`

NewUpdateNamespacePropertiesRequestWithDefaults instantiates a new UpdateNamespacePropertiesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemovals

`func (o *UpdateNamespacePropertiesRequest) GetRemovals() []string`

GetRemovals returns the Removals field if non-nil, zero value otherwise.

### GetRemovalsOk

`func (o *UpdateNamespacePropertiesRequest) GetRemovalsOk() (*[]string, bool)`

GetRemovalsOk returns a tuple with the Removals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovals

`func (o *UpdateNamespacePropertiesRequest) SetRemovals(v []string)`

SetRemovals sets Removals field to given value.

### HasRemovals

`func (o *UpdateNamespacePropertiesRequest) HasRemovals() bool`

HasRemovals returns a boolean if a field has been set.

### GetUpdates

`func (o *UpdateNamespacePropertiesRequest) GetUpdates() map[string]string`

GetUpdates returns the Updates field if non-nil, zero value otherwise.

### GetUpdatesOk

`func (o *UpdateNamespacePropertiesRequest) GetUpdatesOk() (*map[string]string, bool)`

GetUpdatesOk returns a tuple with the Updates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdates

`func (o *UpdateNamespacePropertiesRequest) SetUpdates(v map[string]string)`

SetUpdates sets Updates field to given value.

### HasUpdates

`func (o *UpdateNamespacePropertiesRequest) HasUpdates() bool`

HasUpdates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


