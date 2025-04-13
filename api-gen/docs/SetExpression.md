# SetExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Term** | [**Term**](Term.md) |  | 
**Values** | **[]map[string]interface{}** |  | 

## Methods

### NewSetExpression

`func NewSetExpression(type_ string, term Term, values []map[string]interface{}, ) *SetExpression`

NewSetExpression instantiates a new SetExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetExpressionWithDefaults

`func NewSetExpressionWithDefaults() *SetExpression`

NewSetExpressionWithDefaults instantiates a new SetExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SetExpression) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SetExpression) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SetExpression) SetType(v string)`

SetType sets Type field to given value.


### GetTerm

`func (o *SetExpression) GetTerm() Term`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *SetExpression) GetTermOk() (*Term, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *SetExpression) SetTerm(v Term)`

SetTerm sets Term field to given value.


### GetValues

`func (o *SetExpression) GetValues() []map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *SetExpression) GetValuesOk() (*[]map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *SetExpression) SetValues(v []map[string]interface{})`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


