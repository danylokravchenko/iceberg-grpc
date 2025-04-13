# GetNamespaceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | **[]string** | Reference to one or more levels of a namespace | 
**Properties** | Pointer to **map[string]string** | Properties stored on the namespace, if supported by the server. If the server does not support namespace properties, it should return null for this field. If namespace properties are supported, but none are set, it should return an empty object. | [optional] [default to {}]

## Methods

### NewGetNamespaceResponse

`func NewGetNamespaceResponse(namespace []string, ) *GetNamespaceResponse`

NewGetNamespaceResponse instantiates a new GetNamespaceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNamespaceResponseWithDefaults

`func NewGetNamespaceResponseWithDefaults() *GetNamespaceResponse`

NewGetNamespaceResponseWithDefaults instantiates a new GetNamespaceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *GetNamespaceResponse) GetNamespace() []string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *GetNamespaceResponse) GetNamespaceOk() (*[]string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *GetNamespaceResponse) SetNamespace(v []string)`

SetNamespace sets Namespace field to given value.


### GetProperties

`func (o *GetNamespaceResponse) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *GetNamespaceResponse) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *GetNamespaceResponse) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *GetNamespaceResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


