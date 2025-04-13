# RemoveSchemasUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**SchemaIds** | **[]int32** |  | 

## Methods

### NewRemoveSchemasUpdate

`func NewRemoveSchemasUpdate(schemaIds []int32, ) *RemoveSchemasUpdate`

NewRemoveSchemasUpdate instantiates a new RemoveSchemasUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveSchemasUpdateWithDefaults

`func NewRemoveSchemasUpdateWithDefaults() *RemoveSchemasUpdate`

NewRemoveSchemasUpdateWithDefaults instantiates a new RemoveSchemasUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *RemoveSchemasUpdate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RemoveSchemasUpdate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RemoveSchemasUpdate) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *RemoveSchemasUpdate) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetSchemaIds

`func (o *RemoveSchemasUpdate) GetSchemaIds() []int32`

GetSchemaIds returns the SchemaIds field if non-nil, zero value otherwise.

### GetSchemaIdsOk

`func (o *RemoveSchemasUpdate) GetSchemaIdsOk() (*[]int32, bool)`

GetSchemaIdsOk returns a tuple with the SchemaIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaIds

`func (o *RemoveSchemasUpdate) SetSchemaIds(v []int32)`

SetSchemaIds sets SchemaIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


