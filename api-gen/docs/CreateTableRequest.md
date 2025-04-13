# CreateTableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Location** | Pointer to **string** |  | [optional] 
**Schema** | [**Schema**](Schema.md) |  | 
**PartitionSpec** | Pointer to [**PartitionSpec**](PartitionSpec.md) |  | [optional] 
**WriteOrder** | Pointer to [**SortOrder**](SortOrder.md) |  | [optional] 
**StageCreate** | Pointer to **bool** |  | [optional] 
**Properties** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCreateTableRequest

`func NewCreateTableRequest(name string, schema Schema, ) *CreateTableRequest`

NewCreateTableRequest instantiates a new CreateTableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTableRequestWithDefaults

`func NewCreateTableRequestWithDefaults() *CreateTableRequest`

NewCreateTableRequestWithDefaults instantiates a new CreateTableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTableRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTableRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTableRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLocation

`func (o *CreateTableRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CreateTableRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CreateTableRequest) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CreateTableRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetSchema

`func (o *CreateTableRequest) GetSchema() Schema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *CreateTableRequest) GetSchemaOk() (*Schema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *CreateTableRequest) SetSchema(v Schema)`

SetSchema sets Schema field to given value.


### GetPartitionSpec

`func (o *CreateTableRequest) GetPartitionSpec() PartitionSpec`

GetPartitionSpec returns the PartitionSpec field if non-nil, zero value otherwise.

### GetPartitionSpecOk

`func (o *CreateTableRequest) GetPartitionSpecOk() (*PartitionSpec, bool)`

GetPartitionSpecOk returns a tuple with the PartitionSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionSpec

`func (o *CreateTableRequest) SetPartitionSpec(v PartitionSpec)`

SetPartitionSpec sets PartitionSpec field to given value.

### HasPartitionSpec

`func (o *CreateTableRequest) HasPartitionSpec() bool`

HasPartitionSpec returns a boolean if a field has been set.

### GetWriteOrder

`func (o *CreateTableRequest) GetWriteOrder() SortOrder`

GetWriteOrder returns the WriteOrder field if non-nil, zero value otherwise.

### GetWriteOrderOk

`func (o *CreateTableRequest) GetWriteOrderOk() (*SortOrder, bool)`

GetWriteOrderOk returns a tuple with the WriteOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteOrder

`func (o *CreateTableRequest) SetWriteOrder(v SortOrder)`

SetWriteOrder sets WriteOrder field to given value.

### HasWriteOrder

`func (o *CreateTableRequest) HasWriteOrder() bool`

HasWriteOrder returns a boolean if a field has been set.

### GetStageCreate

`func (o *CreateTableRequest) GetStageCreate() bool`

GetStageCreate returns the StageCreate field if non-nil, zero value otherwise.

### GetStageCreateOk

`func (o *CreateTableRequest) GetStageCreateOk() (*bool, bool)`

GetStageCreateOk returns a tuple with the StageCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageCreate

`func (o *CreateTableRequest) SetStageCreate(v bool)`

SetStageCreate sets StageCreate field to given value.

### HasStageCreate

`func (o *CreateTableRequest) HasStageCreate() bool`

HasStageCreate returns a boolean if a field has been set.

### GetProperties

`func (o *CreateTableRequest) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CreateTableRequest) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CreateTableRequest) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CreateTableRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


