# SetCurrentSchemaUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**SchemaId** | **int32** | Schema ID to set as current, or -1 to set last added schema | 

## Methods

### NewSetCurrentSchemaUpdate

`func NewSetCurrentSchemaUpdate(schemaId int32, ) *SetCurrentSchemaUpdate`

NewSetCurrentSchemaUpdate instantiates a new SetCurrentSchemaUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetCurrentSchemaUpdateWithDefaults

`func NewSetCurrentSchemaUpdateWithDefaults() *SetCurrentSchemaUpdate`

NewSetCurrentSchemaUpdateWithDefaults instantiates a new SetCurrentSchemaUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *SetCurrentSchemaUpdate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SetCurrentSchemaUpdate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SetCurrentSchemaUpdate) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SetCurrentSchemaUpdate) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetSchemaId

`func (o *SetCurrentSchemaUpdate) GetSchemaId() int32`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *SetCurrentSchemaUpdate) GetSchemaIdOk() (*int32, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *SetCurrentSchemaUpdate) SetSchemaId(v int32)`

SetSchemaId sets SchemaId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


