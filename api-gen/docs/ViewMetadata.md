# ViewMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ViewUuid** | **string** |  | 
**FormatVersion** | **int32** |  | 
**Location** | **string** |  | 
**CurrentVersionId** | **int32** |  | 
**Versions** | [**[]ViewVersion**](ViewVersion.md) |  | 
**VersionLog** | [**[]ViewHistoryEntry**](ViewHistoryEntry.md) |  | 
**Schemas** | [**[]Schema**](Schema.md) |  | 
**Properties** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewViewMetadata

`func NewViewMetadata(viewUuid string, formatVersion int32, location string, currentVersionId int32, versions []ViewVersion, versionLog []ViewHistoryEntry, schemas []Schema, ) *ViewMetadata`

NewViewMetadata instantiates a new ViewMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewMetadataWithDefaults

`func NewViewMetadataWithDefaults() *ViewMetadata`

NewViewMetadataWithDefaults instantiates a new ViewMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetViewUuid

`func (o *ViewMetadata) GetViewUuid() string`

GetViewUuid returns the ViewUuid field if non-nil, zero value otherwise.

### GetViewUuidOk

`func (o *ViewMetadata) GetViewUuidOk() (*string, bool)`

GetViewUuidOk returns a tuple with the ViewUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewUuid

`func (o *ViewMetadata) SetViewUuid(v string)`

SetViewUuid sets ViewUuid field to given value.


### GetFormatVersion

`func (o *ViewMetadata) GetFormatVersion() int32`

GetFormatVersion returns the FormatVersion field if non-nil, zero value otherwise.

### GetFormatVersionOk

`func (o *ViewMetadata) GetFormatVersionOk() (*int32, bool)`

GetFormatVersionOk returns a tuple with the FormatVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatVersion

`func (o *ViewMetadata) SetFormatVersion(v int32)`

SetFormatVersion sets FormatVersion field to given value.


### GetLocation

`func (o *ViewMetadata) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ViewMetadata) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ViewMetadata) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetCurrentVersionId

`func (o *ViewMetadata) GetCurrentVersionId() int32`

GetCurrentVersionId returns the CurrentVersionId field if non-nil, zero value otherwise.

### GetCurrentVersionIdOk

`func (o *ViewMetadata) GetCurrentVersionIdOk() (*int32, bool)`

GetCurrentVersionIdOk returns a tuple with the CurrentVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersionId

`func (o *ViewMetadata) SetCurrentVersionId(v int32)`

SetCurrentVersionId sets CurrentVersionId field to given value.


### GetVersions

`func (o *ViewMetadata) GetVersions() []ViewVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ViewMetadata) GetVersionsOk() (*[]ViewVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ViewMetadata) SetVersions(v []ViewVersion)`

SetVersions sets Versions field to given value.


### GetVersionLog

`func (o *ViewMetadata) GetVersionLog() []ViewHistoryEntry`

GetVersionLog returns the VersionLog field if non-nil, zero value otherwise.

### GetVersionLogOk

`func (o *ViewMetadata) GetVersionLogOk() (*[]ViewHistoryEntry, bool)`

GetVersionLogOk returns a tuple with the VersionLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionLog

`func (o *ViewMetadata) SetVersionLog(v []ViewHistoryEntry)`

SetVersionLog sets VersionLog field to given value.


### GetSchemas

`func (o *ViewMetadata) GetSchemas() []Schema`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ViewMetadata) GetSchemasOk() (*[]Schema, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ViewMetadata) SetSchemas(v []Schema)`

SetSchemas sets Schemas field to given value.


### GetProperties

`func (o *ViewMetadata) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ViewMetadata) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ViewMetadata) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ViewMetadata) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


