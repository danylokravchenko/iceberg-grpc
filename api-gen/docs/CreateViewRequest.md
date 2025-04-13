# CreateViewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Location** | Pointer to **string** |  | [optional] 
**Schema** | [**Schema**](Schema.md) |  | 
**ViewVersion** | [**ViewVersion**](ViewVersion.md) | The view version to create, will replace the schema-id sent within the view-version with the id assigned to the provided schema | 
**Properties** | **map[string]string** |  | 

## Methods

### NewCreateViewRequest

`func NewCreateViewRequest(name string, schema Schema, viewVersion ViewVersion, properties map[string]string, ) *CreateViewRequest`

NewCreateViewRequest instantiates a new CreateViewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateViewRequestWithDefaults

`func NewCreateViewRequestWithDefaults() *CreateViewRequest`

NewCreateViewRequestWithDefaults instantiates a new CreateViewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateViewRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateViewRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateViewRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLocation

`func (o *CreateViewRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CreateViewRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CreateViewRequest) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CreateViewRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetSchema

`func (o *CreateViewRequest) GetSchema() Schema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *CreateViewRequest) GetSchemaOk() (*Schema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *CreateViewRequest) SetSchema(v Schema)`

SetSchema sets Schema field to given value.


### GetViewVersion

`func (o *CreateViewRequest) GetViewVersion() ViewVersion`

GetViewVersion returns the ViewVersion field if non-nil, zero value otherwise.

### GetViewVersionOk

`func (o *CreateViewRequest) GetViewVersionOk() (*ViewVersion, bool)`

GetViewVersionOk returns a tuple with the ViewVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewVersion

`func (o *CreateViewRequest) SetViewVersion(v ViewVersion)`

SetViewVersion sets ViewVersion field to given value.


### GetProperties

`func (o *CreateViewRequest) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CreateViewRequest) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CreateViewRequest) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


