# UnaryExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Term** | [**Term**](Term.md) |  | 
**Value** | **map[string]interface{}** |  | 

## Methods

### NewUnaryExpression

`func NewUnaryExpression(type_ string, term Term, value map[string]interface{}, ) *UnaryExpression`

NewUnaryExpression instantiates a new UnaryExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnaryExpressionWithDefaults

`func NewUnaryExpressionWithDefaults() *UnaryExpression`

NewUnaryExpressionWithDefaults instantiates a new UnaryExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UnaryExpression) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnaryExpression) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnaryExpression) SetType(v string)`

SetType sets Type field to given value.


### GetTerm

`func (o *UnaryExpression) GetTerm() Term`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *UnaryExpression) GetTermOk() (*Term, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *UnaryExpression) SetTerm(v Term)`

SetTerm sets Term field to given value.


### GetValue

`func (o *UnaryExpression) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UnaryExpression) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UnaryExpression) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


