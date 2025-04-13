# AddSchemaUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Schema** | [**Schema**](Schema.md) |  | 
**LastColumnId** | Pointer to **int32** | This optional field is **DEPRECATED for REMOVAL** since it more safe to handle this internally, and shouldn&#39;t be exposed to the clients. The highest assigned column ID for the table. This is used to ensure columns are always assigned an unused ID when evolving schemas. When omitted, it will be computed on the server side. | [optional] 

## Methods

### NewAddSchemaUpdate

`func NewAddSchemaUpdate(schema Schema, ) *AddSchemaUpdate`

NewAddSchemaUpdate instantiates a new AddSchemaUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSchemaUpdateWithDefaults

`func NewAddSchemaUpdateWithDefaults() *AddSchemaUpdate`

NewAddSchemaUpdateWithDefaults instantiates a new AddSchemaUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AddSchemaUpdate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AddSchemaUpdate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AddSchemaUpdate) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AddSchemaUpdate) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetSchema

`func (o *AddSchemaUpdate) GetSchema() Schema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *AddSchemaUpdate) GetSchemaOk() (*Schema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *AddSchemaUpdate) SetSchema(v Schema)`

SetSchema sets Schema field to given value.


### GetLastColumnId

`func (o *AddSchemaUpdate) GetLastColumnId() int32`

GetLastColumnId returns the LastColumnId field if non-nil, zero value otherwise.

### GetLastColumnIdOk

`func (o *AddSchemaUpdate) GetLastColumnIdOk() (*int32, bool)`

GetLastColumnIdOk returns a tuple with the LastColumnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastColumnId

`func (o *AddSchemaUpdate) SetLastColumnId(v int32)`

SetLastColumnId sets LastColumnId field to given value.

### HasLastColumnId

`func (o *AddSchemaUpdate) HasLastColumnId() bool`

HasLastColumnId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


