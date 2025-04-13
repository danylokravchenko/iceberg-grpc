# RenameTableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**TableIdentifier**](TableIdentifier.md) |  | 
**Destination** | [**TableIdentifier**](TableIdentifier.md) |  | 

## Methods

### NewRenameTableRequest

`func NewRenameTableRequest(source TableIdentifier, destination TableIdentifier, ) *RenameTableRequest`

NewRenameTableRequest instantiates a new RenameTableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenameTableRequestWithDefaults

`func NewRenameTableRequestWithDefaults() *RenameTableRequest`

NewRenameTableRequestWithDefaults instantiates a new RenameTableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *RenameTableRequest) GetSource() TableIdentifier`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RenameTableRequest) GetSourceOk() (*TableIdentifier, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RenameTableRequest) SetSource(v TableIdentifier)`

SetSource sets Source field to given value.


### GetDestination

`func (o *RenameTableRequest) GetDestination() TableIdentifier`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *RenameTableRequest) GetDestinationOk() (*TableIdentifier, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *RenameTableRequest) SetDestination(v TableIdentifier)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


