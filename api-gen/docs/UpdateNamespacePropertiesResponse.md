# UpdateNamespacePropertiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Updated** | **[]string** | List of property keys that were added or updated | 
**Removed** | **[]string** | List of properties that were removed | 
**Missing** | Pointer to **[]string** | List of properties requested for removal that were not found in the namespace&#39;s properties. Represents a partial success response. Server&#39;s do not need to implement this. | [optional] 

## Methods

### NewUpdateNamespacePropertiesResponse

`func NewUpdateNamespacePropertiesResponse(updated []string, removed []string, ) *UpdateNamespacePropertiesResponse`

NewUpdateNamespacePropertiesResponse instantiates a new UpdateNamespacePropertiesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNamespacePropertiesResponseWithDefaults

`func NewUpdateNamespacePropertiesResponseWithDefaults() *UpdateNamespacePropertiesResponse`

NewUpdateNamespacePropertiesResponseWithDefaults instantiates a new UpdateNamespacePropertiesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdated

`func (o *UpdateNamespacePropertiesResponse) GetUpdated() []string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *UpdateNamespacePropertiesResponse) GetUpdatedOk() (*[]string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *UpdateNamespacePropertiesResponse) SetUpdated(v []string)`

SetUpdated sets Updated field to given value.


### GetRemoved

`func (o *UpdateNamespacePropertiesResponse) GetRemoved() []string`

GetRemoved returns the Removed field if non-nil, zero value otherwise.

### GetRemovedOk

`func (o *UpdateNamespacePropertiesResponse) GetRemovedOk() (*[]string, bool)`

GetRemovedOk returns a tuple with the Removed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoved

`func (o *UpdateNamespacePropertiesResponse) SetRemoved(v []string)`

SetRemoved sets Removed field to given value.


### GetMissing

`func (o *UpdateNamespacePropertiesResponse) GetMissing() []string`

GetMissing returns the Missing field if non-nil, zero value otherwise.

### GetMissingOk

`func (o *UpdateNamespacePropertiesResponse) GetMissingOk() (*[]string, bool)`

GetMissingOk returns a tuple with the Missing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissing

`func (o *UpdateNamespacePropertiesResponse) SetMissing(v []string)`

SetMissing sets Missing field to given value.

### HasMissing

`func (o *UpdateNamespacePropertiesResponse) HasMissing() bool`

HasMissing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


