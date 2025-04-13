# ViewVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionId** | **int32** |  | 
**TimestampMs** | **int64** |  | 
**SchemaId** | **int32** | Schema ID to set as current, or -1 to set last added schema | 
**Summary** | **map[string]string** |  | 
**Representations** | [**[]ViewRepresentation**](ViewRepresentation.md) |  | 
**DefaultCatalog** | Pointer to **string** |  | [optional] 
**DefaultNamespace** | **[]string** | Reference to one or more levels of a namespace | 

## Methods

### NewViewVersion

`func NewViewVersion(versionId int32, timestampMs int64, schemaId int32, summary map[string]string, representations []ViewRepresentation, defaultNamespace []string, ) *ViewVersion`

NewViewVersion instantiates a new ViewVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewVersionWithDefaults

`func NewViewVersionWithDefaults() *ViewVersion`

NewViewVersionWithDefaults instantiates a new ViewVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionId

`func (o *ViewVersion) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *ViewVersion) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *ViewVersion) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.


### GetTimestampMs

`func (o *ViewVersion) GetTimestampMs() int64`

GetTimestampMs returns the TimestampMs field if non-nil, zero value otherwise.

### GetTimestampMsOk

`func (o *ViewVersion) GetTimestampMsOk() (*int64, bool)`

GetTimestampMsOk returns a tuple with the TimestampMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampMs

`func (o *ViewVersion) SetTimestampMs(v int64)`

SetTimestampMs sets TimestampMs field to given value.


### GetSchemaId

`func (o *ViewVersion) GetSchemaId() int32`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *ViewVersion) GetSchemaIdOk() (*int32, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *ViewVersion) SetSchemaId(v int32)`

SetSchemaId sets SchemaId field to given value.


### GetSummary

`func (o *ViewVersion) GetSummary() map[string]string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ViewVersion) GetSummaryOk() (*map[string]string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ViewVersion) SetSummary(v map[string]string)`

SetSummary sets Summary field to given value.


### GetRepresentations

`func (o *ViewVersion) GetRepresentations() []ViewRepresentation`

GetRepresentations returns the Representations field if non-nil, zero value otherwise.

### GetRepresentationsOk

`func (o *ViewVersion) GetRepresentationsOk() (*[]ViewRepresentation, bool)`

GetRepresentationsOk returns a tuple with the Representations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepresentations

`func (o *ViewVersion) SetRepresentations(v []ViewRepresentation)`

SetRepresentations sets Representations field to given value.


### GetDefaultCatalog

`func (o *ViewVersion) GetDefaultCatalog() string`

GetDefaultCatalog returns the DefaultCatalog field if non-nil, zero value otherwise.

### GetDefaultCatalogOk

`func (o *ViewVersion) GetDefaultCatalogOk() (*string, bool)`

GetDefaultCatalogOk returns a tuple with the DefaultCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCatalog

`func (o *ViewVersion) SetDefaultCatalog(v string)`

SetDefaultCatalog sets DefaultCatalog field to given value.

### HasDefaultCatalog

`func (o *ViewVersion) HasDefaultCatalog() bool`

HasDefaultCatalog returns a boolean if a field has been set.

### GetDefaultNamespace

`func (o *ViewVersion) GetDefaultNamespace() []string`

GetDefaultNamespace returns the DefaultNamespace field if non-nil, zero value otherwise.

### GetDefaultNamespaceOk

`func (o *ViewVersion) GetDefaultNamespaceOk() (*[]string, bool)`

GetDefaultNamespaceOk returns a tuple with the DefaultNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNamespace

`func (o *ViewVersion) SetDefaultNamespace(v []string)`

SetDefaultNamespace sets DefaultNamespace field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


