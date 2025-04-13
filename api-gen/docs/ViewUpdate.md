# ViewUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**FormatVersion** | **int32** |  | 
**Schema** | [**Schema**](Schema.md) |  | 
**LastColumnId** | Pointer to **int32** | This optional field is **DEPRECATED for REMOVAL** since it more safe to handle this internally, and shouldn&#39;t be exposed to the clients. The highest assigned column ID for the table. This is used to ensure columns are always assigned an unused ID when evolving schemas. When omitted, it will be computed on the server side. | [optional] 
**Location** | **string** |  | 
**Updates** | **map[string]string** |  | 
**Removals** | **[]string** |  | 
**ViewVersion** | [**ViewVersion**](ViewVersion.md) |  | 
**ViewVersionId** | **int32** | The view version id to set as current, or -1 to set last added view version id | 

## Methods

### NewViewUpdate

`func NewViewUpdate(action string, formatVersion int32, schema Schema, location string, updates map[string]string, removals []string, viewVersion ViewVersion, viewVersionId int32, ) *ViewUpdate`

NewViewUpdate instantiates a new ViewUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewUpdateWithDefaults

`func NewViewUpdateWithDefaults() *ViewUpdate`

NewViewUpdateWithDefaults instantiates a new ViewUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ViewUpdate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ViewUpdate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ViewUpdate) SetAction(v string)`

SetAction sets Action field to given value.


### GetFormatVersion

`func (o *ViewUpdate) GetFormatVersion() int32`

GetFormatVersion returns the FormatVersion field if non-nil, zero value otherwise.

### GetFormatVersionOk

`func (o *ViewUpdate) GetFormatVersionOk() (*int32, bool)`

GetFormatVersionOk returns a tuple with the FormatVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatVersion

`func (o *ViewUpdate) SetFormatVersion(v int32)`

SetFormatVersion sets FormatVersion field to given value.


### GetSchema

`func (o *ViewUpdate) GetSchema() Schema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *ViewUpdate) GetSchemaOk() (*Schema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *ViewUpdate) SetSchema(v Schema)`

SetSchema sets Schema field to given value.


### GetLastColumnId

`func (o *ViewUpdate) GetLastColumnId() int32`

GetLastColumnId returns the LastColumnId field if non-nil, zero value otherwise.

### GetLastColumnIdOk

`func (o *ViewUpdate) GetLastColumnIdOk() (*int32, bool)`

GetLastColumnIdOk returns a tuple with the LastColumnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastColumnId

`func (o *ViewUpdate) SetLastColumnId(v int32)`

SetLastColumnId sets LastColumnId field to given value.

### HasLastColumnId

`func (o *ViewUpdate) HasLastColumnId() bool`

HasLastColumnId returns a boolean if a field has been set.

### GetLocation

`func (o *ViewUpdate) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ViewUpdate) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ViewUpdate) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetUpdates

`func (o *ViewUpdate) GetUpdates() map[string]string`

GetUpdates returns the Updates field if non-nil, zero value otherwise.

### GetUpdatesOk

`func (o *ViewUpdate) GetUpdatesOk() (*map[string]string, bool)`

GetUpdatesOk returns a tuple with the Updates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdates

`func (o *ViewUpdate) SetUpdates(v map[string]string)`

SetUpdates sets Updates field to given value.


### GetRemovals

`func (o *ViewUpdate) GetRemovals() []string`

GetRemovals returns the Removals field if non-nil, zero value otherwise.

### GetRemovalsOk

`func (o *ViewUpdate) GetRemovalsOk() (*[]string, bool)`

GetRemovalsOk returns a tuple with the Removals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovals

`func (o *ViewUpdate) SetRemovals(v []string)`

SetRemovals sets Removals field to given value.


### GetViewVersion

`func (o *ViewUpdate) GetViewVersion() ViewVersion`

GetViewVersion returns the ViewVersion field if non-nil, zero value otherwise.

### GetViewVersionOk

`func (o *ViewUpdate) GetViewVersionOk() (*ViewVersion, bool)`

GetViewVersionOk returns a tuple with the ViewVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewVersion

`func (o *ViewUpdate) SetViewVersion(v ViewVersion)`

SetViewVersion sets ViewVersion field to given value.


### GetViewVersionId

`func (o *ViewUpdate) GetViewVersionId() int32`

GetViewVersionId returns the ViewVersionId field if non-nil, zero value otherwise.

### GetViewVersionIdOk

`func (o *ViewUpdate) GetViewVersionIdOk() (*int32, bool)`

GetViewVersionIdOk returns a tuple with the ViewVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewVersionId

`func (o *ViewUpdate) SetViewVersionId(v int32)`

SetViewVersionId sets ViewVersionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


