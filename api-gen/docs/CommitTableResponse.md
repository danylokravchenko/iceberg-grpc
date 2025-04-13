# CommitTableResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataLocation** | **string** |  | 
**Metadata** | [**TableMetadata**](TableMetadata.md) |  | 

## Methods

### NewCommitTableResponse

`func NewCommitTableResponse(metadataLocation string, metadata TableMetadata, ) *CommitTableResponse`

NewCommitTableResponse instantiates a new CommitTableResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitTableResponseWithDefaults

`func NewCommitTableResponseWithDefaults() *CommitTableResponse`

NewCommitTableResponseWithDefaults instantiates a new CommitTableResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataLocation

`func (o *CommitTableResponse) GetMetadataLocation() string`

GetMetadataLocation returns the MetadataLocation field if non-nil, zero value otherwise.

### GetMetadataLocationOk

`func (o *CommitTableResponse) GetMetadataLocationOk() (*string, bool)`

GetMetadataLocationOk returns a tuple with the MetadataLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataLocation

`func (o *CommitTableResponse) SetMetadataLocation(v string)`

SetMetadataLocation sets MetadataLocation field to given value.


### GetMetadata

`func (o *CommitTableResponse) GetMetadata() TableMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CommitTableResponse) GetMetadataOk() (*TableMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CommitTableResponse) SetMetadata(v TableMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


