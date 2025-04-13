# PositionDeleteFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** |  | 
**ContentOffset** | Pointer to **int64** | Offset within the delete file of delete content | [optional] 
**ContentSizeInBytes** | Pointer to **int64** | Length, in bytes, of the delete content; required if content-offset is present | [optional] 

## Methods

### NewPositionDeleteFile

`func NewPositionDeleteFile(content string, ) *PositionDeleteFile`

NewPositionDeleteFile instantiates a new PositionDeleteFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPositionDeleteFileWithDefaults

`func NewPositionDeleteFileWithDefaults() *PositionDeleteFile`

NewPositionDeleteFileWithDefaults instantiates a new PositionDeleteFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *PositionDeleteFile) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PositionDeleteFile) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PositionDeleteFile) SetContent(v string)`

SetContent sets Content field to given value.


### GetContentOffset

`func (o *PositionDeleteFile) GetContentOffset() int64`

GetContentOffset returns the ContentOffset field if non-nil, zero value otherwise.

### GetContentOffsetOk

`func (o *PositionDeleteFile) GetContentOffsetOk() (*int64, bool)`

GetContentOffsetOk returns a tuple with the ContentOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentOffset

`func (o *PositionDeleteFile) SetContentOffset(v int64)`

SetContentOffset sets ContentOffset field to given value.

### HasContentOffset

`func (o *PositionDeleteFile) HasContentOffset() bool`

HasContentOffset returns a boolean if a field has been set.

### GetContentSizeInBytes

`func (o *PositionDeleteFile) GetContentSizeInBytes() int64`

GetContentSizeInBytes returns the ContentSizeInBytes field if non-nil, zero value otherwise.

### GetContentSizeInBytesOk

`func (o *PositionDeleteFile) GetContentSizeInBytesOk() (*int64, bool)`

GetContentSizeInBytesOk returns a tuple with the ContentSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSizeInBytes

`func (o *PositionDeleteFile) SetContentSizeInBytes(v int64)`

SetContentSizeInBytes sets ContentSizeInBytes field to given value.

### HasContentSizeInBytes

`func (o *PositionDeleteFile) HasContentSizeInBytes() bool`

HasContentSizeInBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


