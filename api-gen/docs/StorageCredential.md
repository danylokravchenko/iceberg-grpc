# StorageCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | **string** | Indicates a storage location prefix where the credential is relevant. Clients should choose the most specific prefix (by selecting the longest prefix) if several credentials of the same type are available. | 
**Config** | **map[string]string** |  | 

## Methods

### NewStorageCredential

`func NewStorageCredential(prefix string, config map[string]string, ) *StorageCredential`

NewStorageCredential instantiates a new StorageCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageCredentialWithDefaults

`func NewStorageCredentialWithDefaults() *StorageCredential`

NewStorageCredentialWithDefaults instantiates a new StorageCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *StorageCredential) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *StorageCredential) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *StorageCredential) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetConfig

`func (o *StorageCredential) GetConfig() map[string]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *StorageCredential) GetConfigOk() (*map[string]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *StorageCredential) SetConfig(v map[string]string)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


