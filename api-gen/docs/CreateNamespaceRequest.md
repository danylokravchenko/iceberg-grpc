# CreateNamespaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | **[]string** | Reference to one or more levels of a namespace | 
**Properties** | Pointer to **map[string]string** | Configured string to string map of properties for the namespace | [optional] [default to {}]

## Methods

### NewCreateNamespaceRequest

`func NewCreateNamespaceRequest(namespace []string, ) *CreateNamespaceRequest`

NewCreateNamespaceRequest instantiates a new CreateNamespaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNamespaceRequestWithDefaults

`func NewCreateNamespaceRequestWithDefaults() *CreateNamespaceRequest`

NewCreateNamespaceRequestWithDefaults instantiates a new CreateNamespaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *CreateNamespaceRequest) GetNamespace() []string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CreateNamespaceRequest) GetNamespaceOk() (*[]string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CreateNamespaceRequest) SetNamespace(v []string)`

SetNamespace sets Namespace field to given value.


### GetProperties

`func (o *CreateNamespaceRequest) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CreateNamespaceRequest) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CreateNamespaceRequest) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CreateNamespaceRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


