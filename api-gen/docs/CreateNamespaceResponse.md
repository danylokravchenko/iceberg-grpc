# CreateNamespaceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | **[]string** | Reference to one or more levels of a namespace | 
**Properties** | Pointer to **map[string]string** | Properties stored on the namespace, if supported by the server. | [optional] [default to {}]

## Methods

### NewCreateNamespaceResponse

`func NewCreateNamespaceResponse(namespace []string, ) *CreateNamespaceResponse`

NewCreateNamespaceResponse instantiates a new CreateNamespaceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNamespaceResponseWithDefaults

`func NewCreateNamespaceResponseWithDefaults() *CreateNamespaceResponse`

NewCreateNamespaceResponseWithDefaults instantiates a new CreateNamespaceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *CreateNamespaceResponse) GetNamespace() []string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CreateNamespaceResponse) GetNamespaceOk() (*[]string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CreateNamespaceResponse) SetNamespace(v []string)`

SetNamespace sets Namespace field to given value.


### GetProperties

`func (o *CreateNamespaceResponse) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CreateNamespaceResponse) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CreateNamespaceResponse) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CreateNamespaceResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


