# Expression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Left** | [**Expression**](Expression.md) |  | 
**Right** | [**Expression**](Expression.md) |  | 
**Child** | [**Expression**](Expression.md) |  | 
**Term** | [**Term**](Term.md) |  | 
**Values** | **[]map[string]interface{}** |  | 
**Value** | **map[string]interface{}** |  | 

## Methods

### NewExpression

`func NewExpression(type_ string, left Expression, right Expression, child Expression, term Term, values []map[string]interface{}, value map[string]interface{}, ) *Expression`

NewExpression instantiates a new Expression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpressionWithDefaults

`func NewExpressionWithDefaults() *Expression`

NewExpressionWithDefaults instantiates a new Expression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Expression) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Expression) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Expression) SetType(v string)`

SetType sets Type field to given value.


### GetLeft

`func (o *Expression) GetLeft() Expression`

GetLeft returns the Left field if non-nil, zero value otherwise.

### GetLeftOk

`func (o *Expression) GetLeftOk() (*Expression, bool)`

GetLeftOk returns a tuple with the Left field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeft

`func (o *Expression) SetLeft(v Expression)`

SetLeft sets Left field to given value.


### GetRight

`func (o *Expression) GetRight() Expression`

GetRight returns the Right field if non-nil, zero value otherwise.

### GetRightOk

`func (o *Expression) GetRightOk() (*Expression, bool)`

GetRightOk returns a tuple with the Right field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRight

`func (o *Expression) SetRight(v Expression)`

SetRight sets Right field to given value.


### GetChild

`func (o *Expression) GetChild() Expression`

GetChild returns the Child field if non-nil, zero value otherwise.

### GetChildOk

`func (o *Expression) GetChildOk() (*Expression, bool)`

GetChildOk returns a tuple with the Child field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChild

`func (o *Expression) SetChild(v Expression)`

SetChild sets Child field to given value.


### GetTerm

`func (o *Expression) GetTerm() Term`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *Expression) GetTermOk() (*Term, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *Expression) SetTerm(v Term)`

SetTerm sets Term field to given value.


### GetValues

`func (o *Expression) GetValues() []map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Expression) GetValuesOk() (*[]map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Expression) SetValues(v []map[string]interface{})`

SetValues sets Values field to given value.


### GetValue

`func (o *Expression) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Expression) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Expression) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


