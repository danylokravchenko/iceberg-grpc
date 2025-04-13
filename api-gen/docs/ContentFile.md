# ContentFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** |  | 
**FilePath** | **string** |  | 
**FileFormat** | [**FileFormat**](FileFormat.md) |  | 
**SpecId** | **int32** |  | 
**Partition** | [**[]PrimitiveTypeValue**](PrimitiveTypeValue.md) | A list of partition field values ordered based on the fields of the partition spec specified by the &#x60;spec-id&#x60; | 
**FileSizeInBytes** | **int64** | Total file size in bytes | 
**RecordCount** | **int64** | Number of records in the file | 
**KeyMetadata** | Pointer to **string** | Encryption key metadata blob | [optional] 
**SplitOffsets** | Pointer to **[]int64** | List of splittable offsets | [optional] 
**SortOrderId** | Pointer to **int32** |  | [optional] 

## Methods

### NewContentFile

`func NewContentFile(content string, filePath string, fileFormat FileFormat, specId int32, partition []PrimitiveTypeValue, fileSizeInBytes int64, recordCount int64, ) *ContentFile`

NewContentFile instantiates a new ContentFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentFileWithDefaults

`func NewContentFileWithDefaults() *ContentFile`

NewContentFileWithDefaults instantiates a new ContentFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *ContentFile) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ContentFile) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ContentFile) SetContent(v string)`

SetContent sets Content field to given value.


### GetFilePath

`func (o *ContentFile) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *ContentFile) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *ContentFile) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.


### GetFileFormat

`func (o *ContentFile) GetFileFormat() FileFormat`

GetFileFormat returns the FileFormat field if non-nil, zero value otherwise.

### GetFileFormatOk

`func (o *ContentFile) GetFileFormatOk() (*FileFormat, bool)`

GetFileFormatOk returns a tuple with the FileFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFormat

`func (o *ContentFile) SetFileFormat(v FileFormat)`

SetFileFormat sets FileFormat field to given value.


### GetSpecId

`func (o *ContentFile) GetSpecId() int32`

GetSpecId returns the SpecId field if non-nil, zero value otherwise.

### GetSpecIdOk

`func (o *ContentFile) GetSpecIdOk() (*int32, bool)`

GetSpecIdOk returns a tuple with the SpecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecId

`func (o *ContentFile) SetSpecId(v int32)`

SetSpecId sets SpecId field to given value.


### GetPartition

`func (o *ContentFile) GetPartition() []PrimitiveTypeValue`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *ContentFile) GetPartitionOk() (*[]PrimitiveTypeValue, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *ContentFile) SetPartition(v []PrimitiveTypeValue)`

SetPartition sets Partition field to given value.


### GetFileSizeInBytes

`func (o *ContentFile) GetFileSizeInBytes() int64`

GetFileSizeInBytes returns the FileSizeInBytes field if non-nil, zero value otherwise.

### GetFileSizeInBytesOk

`func (o *ContentFile) GetFileSizeInBytesOk() (*int64, bool)`

GetFileSizeInBytesOk returns a tuple with the FileSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSizeInBytes

`func (o *ContentFile) SetFileSizeInBytes(v int64)`

SetFileSizeInBytes sets FileSizeInBytes field to given value.


### GetRecordCount

`func (o *ContentFile) GetRecordCount() int64`

GetRecordCount returns the RecordCount field if non-nil, zero value otherwise.

### GetRecordCountOk

`func (o *ContentFile) GetRecordCountOk() (*int64, bool)`

GetRecordCountOk returns a tuple with the RecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordCount

`func (o *ContentFile) SetRecordCount(v int64)`

SetRecordCount sets RecordCount field to given value.


### GetKeyMetadata

`func (o *ContentFile) GetKeyMetadata() string`

GetKeyMetadata returns the KeyMetadata field if non-nil, zero value otherwise.

### GetKeyMetadataOk

`func (o *ContentFile) GetKeyMetadataOk() (*string, bool)`

GetKeyMetadataOk returns a tuple with the KeyMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyMetadata

`func (o *ContentFile) SetKeyMetadata(v string)`

SetKeyMetadata sets KeyMetadata field to given value.

### HasKeyMetadata

`func (o *ContentFile) HasKeyMetadata() bool`

HasKeyMetadata returns a boolean if a field has been set.

### GetSplitOffsets

`func (o *ContentFile) GetSplitOffsets() []int64`

GetSplitOffsets returns the SplitOffsets field if non-nil, zero value otherwise.

### GetSplitOffsetsOk

`func (o *ContentFile) GetSplitOffsetsOk() (*[]int64, bool)`

GetSplitOffsetsOk returns a tuple with the SplitOffsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitOffsets

`func (o *ContentFile) SetSplitOffsets(v []int64)`

SetSplitOffsets sets SplitOffsets field to given value.

### HasSplitOffsets

`func (o *ContentFile) HasSplitOffsets() bool`

HasSplitOffsets returns a boolean if a field has been set.

### GetSortOrderId

`func (o *ContentFile) GetSortOrderId() int32`

GetSortOrderId returns the SortOrderId field if non-nil, zero value otherwise.

### GetSortOrderIdOk

`func (o *ContentFile) GetSortOrderIdOk() (*int32, bool)`

GetSortOrderIdOk returns a tuple with the SortOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrderId

`func (o *ContentFile) SetSortOrderId(v int32)`

SetSortOrderId sets SortOrderId field to given value.

### HasSortOrderId

`func (o *ContentFile) HasSortOrderId() bool`

HasSortOrderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


