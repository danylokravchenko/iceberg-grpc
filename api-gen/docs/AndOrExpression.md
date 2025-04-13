# AndOrExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Left** | [**Expression**](Expression.md) |  | 
**Right** | [**Expression**](Expression.md) |  | 

## Methods

### NewAndOrExpression

`func NewAndOrExpression(type_ string, left Expression, right Expression, ) *AndOrExpression`

NewAndOrExpression instantiates a new AndOrExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAndOrExpressionWithDefaults

`func NewAndOrExpressionWithDefaults() *AndOrExpression`

NewAndOrExpressionWithDefaults instantiates a new AndOrExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AndOrExpression) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AndOrExpression) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AndOrExpression) SetType(v string)`

SetType sets Type field to given value.


### GetLeft

`func (o *AndOrExpression) GetLeft() Expression`

GetLeft returns the Left field if non-nil, zero value otherwise.

### GetLeftOk

`func (o *AndOrExpression) GetLeftOk() (*Expression, bool)`

GetLeftOk returns a tuple with the Left field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeft

`func (o *AndOrExpression) SetLeft(v Expression)`

SetLeft sets Left field to given value.


### GetRight

`func (o *AndOrExpression) GetRight() Expression`

GetRight returns the Right field if non-nil, zero value otherwise.

### GetRightOk

`func (o *AndOrExpression) GetRightOk() (*Expression, bool)`

GetRightOk returns a tuple with the Right field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRight

`func (o *AndOrExpression) SetRight(v Expression)`

SetRight sets Right field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


