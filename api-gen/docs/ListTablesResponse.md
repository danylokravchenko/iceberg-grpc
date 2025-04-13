# ListTablesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | An opaque token that allows clients to make use of pagination for list APIs (e.g. ListTables). Clients may initiate the first paginated request by sending an empty query parameter &#x60;pageToken&#x60; to the server. Servers that support pagination should identify the &#x60;pageToken&#x60; parameter and return a &#x60;next-page-token&#x60; in the response if there are more results available.  After the initial request, the value of &#x60;next-page-token&#x60; from each response must be used as the &#x60;pageToken&#x60; parameter value for the next request. The server must return &#x60;null&#x60; value for the &#x60;next-page-token&#x60; in the last response. Servers that support pagination must return all results in a single response with the value of &#x60;next-page-token&#x60; set to &#x60;null&#x60; if the query parameter &#x60;pageToken&#x60; is not set in the request. Servers that do not support pagination should ignore the &#x60;pageToken&#x60; parameter and return all results in a single response. The &#x60;next-page-token&#x60; must be omitted from the response. Clients must interpret either &#x60;null&#x60; or missing response value of &#x60;next-page-token&#x60; as the end of the listing results. | [optional] 
**Identifiers** | Pointer to [**[]TableIdentifier**](TableIdentifier.md) |  | [optional] 

## Methods

### NewListTablesResponse

`func NewListTablesResponse() *ListTablesResponse`

NewListTablesResponse instantiates a new ListTablesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTablesResponseWithDefaults

`func NewListTablesResponseWithDefaults() *ListTablesResponse`

NewListTablesResponseWithDefaults instantiates a new ListTablesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *ListTablesResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListTablesResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListTablesResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListTablesResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetIdentifiers

`func (o *ListTablesResponse) GetIdentifiers() []TableIdentifier`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *ListTablesResponse) GetIdentifiersOk() (*[]TableIdentifier, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *ListTablesResponse) SetIdentifiers(v []TableIdentifier)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *ListTablesResponse) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


