# FileScanTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataFile** | [**DataFile**](DataFile.md) |  | 
**DeleteFileReferences** | Pointer to **[]int32** | A list of indices in the delete files array (0-based) | [optional] 
**ResidualFilter** | Pointer to [**Expression**](Expression.md) | An optional filter to be applied to rows in this file scan task. If the residual is not present, the client must produce the residual or use the original filter. | [optional] 

## Methods

### NewFileScanTask

`func NewFileScanTask(dataFile DataFile, ) *FileScanTask`

NewFileScanTask instantiates a new FileScanTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileScanTaskWithDefaults

`func NewFileScanTaskWithDefaults() *FileScanTask`

NewFileScanTaskWithDefaults instantiates a new FileScanTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataFile

`func (o *FileScanTask) GetDataFile() DataFile`

GetDataFile returns the DataFile field if non-nil, zero value otherwise.

### GetDataFileOk

`func (o *FileScanTask) GetDataFileOk() (*DataFile, bool)`

GetDataFileOk returns a tuple with the DataFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFile

`func (o *FileScanTask) SetDataFile(v DataFile)`

SetDataFile sets DataFile field to given value.


### GetDeleteFileReferences

`func (o *FileScanTask) GetDeleteFileReferences() []int32`

GetDeleteFileReferences returns the DeleteFileReferences field if non-nil, zero value otherwise.

### GetDeleteFileReferencesOk

`func (o *FileScanTask) GetDeleteFileReferencesOk() (*[]int32, bool)`

GetDeleteFileReferencesOk returns a tuple with the DeleteFileReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteFileReferences

`func (o *FileScanTask) SetDeleteFileReferences(v []int32)`

SetDeleteFileReferences sets DeleteFileReferences field to given value.

### HasDeleteFileReferences

`func (o *FileScanTask) HasDeleteFileReferences() bool`

HasDeleteFileReferences returns a boolean if a field has been set.

### GetResidualFilter

`func (o *FileScanTask) GetResidualFilter() Expression`

GetResidualFilter returns the ResidualFilter field if non-nil, zero value otherwise.

### GetResidualFilterOk

`func (o *FileScanTask) GetResidualFilterOk() (*Expression, bool)`

GetResidualFilterOk returns a tuple with the ResidualFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidualFilter

`func (o *FileScanTask) SetResidualFilter(v Expression)`

SetResidualFilter sets ResidualFilter field to given value.

### HasResidualFilter

`func (o *FileScanTask) HasResidualFilter() bool`

HasResidualFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


