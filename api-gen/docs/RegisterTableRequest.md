# RegisterTableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**MetadataLocation** | **string** |  | 
**Overwrite** | Pointer to **bool** | Whether to overwrite table metadata if the table already exists | [optional] [default to false]

## Methods

### NewRegisterTableRequest

`func NewRegisterTableRequest(name string, metadataLocation string, ) *RegisterTableRequest`

NewRegisterTableRequest instantiates a new RegisterTableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterTableRequestWithDefaults

`func NewRegisterTableRequestWithDefaults() *RegisterTableRequest`

NewRegisterTableRequestWithDefaults instantiates a new RegisterTableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RegisterTableRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegisterTableRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegisterTableRequest) SetName(v string)`

SetName sets Name field to given value.


### GetMetadataLocation

`func (o *RegisterTableRequest) GetMetadataLocation() string`

GetMetadataLocation returns the MetadataLocation field if non-nil, zero value otherwise.

### GetMetadataLocationOk

`func (o *RegisterTableRequest) GetMetadataLocationOk() (*string, bool)`

GetMetadataLocationOk returns a tuple with the MetadataLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataLocation

`func (o *RegisterTableRequest) SetMetadataLocation(v string)`

SetMetadataLocation sets MetadataLocation field to given value.


### GetOverwrite

`func (o *RegisterTableRequest) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *RegisterTableRequest) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *RegisterTableRequest) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *RegisterTableRequest) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


