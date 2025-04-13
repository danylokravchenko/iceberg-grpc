# NotExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Child** | [**Expression**](Expression.md) |  | 

## Methods

### NewNotExpression

`func NewNotExpression(type_ string, child Expression, ) *NotExpression`

NewNotExpression instantiates a new NotExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotExpressionWithDefaults

`func NewNotExpressionWithDefaults() *NotExpression`

NewNotExpressionWithDefaults instantiates a new NotExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NotExpression) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotExpression) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotExpression) SetType(v string)`

SetType sets Type field to given value.


### GetChild

`func (o *NotExpression) GetChild() Expression`

GetChild returns the Child field if non-nil, zero value otherwise.

### GetChildOk

`func (o *NotExpression) GetChildOk() (*Expression, bool)`

GetChildOk returns a tuple with the Child field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChild

`func (o *NotExpression) SetChild(v Expression)`

SetChild sets Child field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


