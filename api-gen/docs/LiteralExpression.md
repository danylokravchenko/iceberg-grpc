# LiteralExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Term** | [**Term**](Term.md) |  | 
**Value** | **map[string]interface{}** |  | 

## Methods

### NewLiteralExpression

`func NewLiteralExpression(type_ string, term Term, value map[string]interface{}, ) *LiteralExpression`

NewLiteralExpression instantiates a new LiteralExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiteralExpressionWithDefaults

`func NewLiteralExpressionWithDefaults() *LiteralExpression`

NewLiteralExpressionWithDefaults instantiates a new LiteralExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LiteralExpression) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LiteralExpression) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LiteralExpression) SetType(v string)`

SetType sets Type field to given value.


### GetTerm

`func (o *LiteralExpression) GetTerm() Term`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *LiteralExpression) GetTermOk() (*Term, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *LiteralExpression) SetTerm(v Term)`

SetTerm sets Term field to given value.


### GetValue

`func (o *LiteralExpression) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LiteralExpression) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LiteralExpression) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


