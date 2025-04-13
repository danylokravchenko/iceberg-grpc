# \CatalogAPIAPI

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelPlanning**](CatalogAPIAPI.md#CancelPlanning) | **Delete** /v1/{prefix}/namespaces/{namespace}/tables/{table}/plan/{plan-id} | Cancels scan planning for a plan-id
[**CommitTransaction**](CatalogAPIAPI.md#CommitTransaction) | **Post** /v1/{prefix}/transactions/commit | Commit updates to multiple tables in an atomic operation
[**CreateNamespace**](CatalogAPIAPI.md#CreateNamespace) | **Post** /v1/{prefix}/namespaces | Create a namespace
[**CreateTable**](CatalogAPIAPI.md#CreateTable) | **Post** /v1/{prefix}/namespaces/{namespace}/tables | Create a table in the given namespace
[**CreateView**](CatalogAPIAPI.md#CreateView) | **Post** /v1/{prefix}/namespaces/{namespace}/views | Create a view in the given namespace
[**DropNamespace**](CatalogAPIAPI.md#DropNamespace) | **Delete** /v1/{prefix}/namespaces/{namespace} | Drop a namespace from the catalog. Namespace must be empty.
[**DropTable**](CatalogAPIAPI.md#DropTable) | **Delete** /v1/{prefix}/namespaces/{namespace}/tables/{table} | Drop a table from the catalog
[**DropView**](CatalogAPIAPI.md#DropView) | **Delete** /v1/{prefix}/namespaces/{namespace}/views/{view} | Drop a view from the catalog
[**FetchPlanningResult**](CatalogAPIAPI.md#FetchPlanningResult) | **Get** /v1/{prefix}/namespaces/{namespace}/tables/{table}/plan/{plan-id} | Fetches the result of scan planning for a plan-id
[**FetchScanTasks**](CatalogAPIAPI.md#FetchScanTasks) | **Post** /v1/{prefix}/namespaces/{namespace}/tables/{table}/tasks | Fetches result tasks for a plan task
[**ListNamespaces**](CatalogAPIAPI.md#ListNamespaces) | **Get** /v1/{prefix}/namespaces | List namespaces, optionally providing a parent namespace to list underneath
[**ListTables**](CatalogAPIAPI.md#ListTables) | **Get** /v1/{prefix}/namespaces/{namespace}/tables | List all table identifiers underneath a given namespace
[**ListViews**](CatalogAPIAPI.md#ListViews) | **Get** /v1/{prefix}/namespaces/{namespace}/views | List all view identifiers underneath a given namespace
[**LoadCredentials**](CatalogAPIAPI.md#LoadCredentials) | **Get** /v1/{prefix}/namespaces/{namespace}/tables/{table}/credentials | Load vended credentials for a table from the catalog
[**LoadNamespaceMetadata**](CatalogAPIAPI.md#LoadNamespaceMetadata) | **Get** /v1/{prefix}/namespaces/{namespace} | Load the metadata properties for a namespace
[**LoadTable**](CatalogAPIAPI.md#LoadTable) | **Get** /v1/{prefix}/namespaces/{namespace}/tables/{table} | Load a table from the catalog
[**LoadView**](CatalogAPIAPI.md#LoadView) | **Get** /v1/{prefix}/namespaces/{namespace}/views/{view} | Load a view from the catalog
[**NamespaceExists**](CatalogAPIAPI.md#NamespaceExists) | **Head** /v1/{prefix}/namespaces/{namespace} | Check if a namespace exists
[**PlanTableScan**](CatalogAPIAPI.md#PlanTableScan) | **Post** /v1/{prefix}/namespaces/{namespace}/tables/{table}/plan | Submit a scan for planning
[**RegisterTable**](CatalogAPIAPI.md#RegisterTable) | **Post** /v1/{prefix}/namespaces/{namespace}/register | Register a table in the given namespace using given metadata file location
[**RenameTable**](CatalogAPIAPI.md#RenameTable) | **Post** /v1/{prefix}/tables/rename | Rename a table from its current name to a new name
[**RenameView**](CatalogAPIAPI.md#RenameView) | **Post** /v1/{prefix}/views/rename | Rename a view from its current name to a new name
[**ReplaceView**](CatalogAPIAPI.md#ReplaceView) | **Post** /v1/{prefix}/namespaces/{namespace}/views/{view} | Replace a view
[**ReportMetrics**](CatalogAPIAPI.md#ReportMetrics) | **Post** /v1/{prefix}/namespaces/{namespace}/tables/{table}/metrics | Send a metrics report to this endpoint to be processed by the backend
[**TableExists**](CatalogAPIAPI.md#TableExists) | **Head** /v1/{prefix}/namespaces/{namespace}/tables/{table} | Check if a table exists
[**UpdateProperties**](CatalogAPIAPI.md#UpdateProperties) | **Post** /v1/{prefix}/namespaces/{namespace}/properties | Set or remove properties on a namespace
[**UpdateTable**](CatalogAPIAPI.md#UpdateTable) | **Post** /v1/{prefix}/namespaces/{namespace}/tables/{table} | Commit updates to a table
[**ViewExists**](CatalogAPIAPI.md#ViewExists) | **Head** /v1/{prefix}/namespaces/{namespace}/views/{view} | Check if a view exists



## CancelPlanning

> CancelPlanning(ctx, prefix, namespace, table, planId).Execute()

Cancels scan planning for a plan-id



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	namespace := "accounting" // string | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (`0x1F`) byte.
	table := "sales" // string | A table name
	planId := "planId_example" // string | ID used to track a planning request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogAPIAPI.CancelPlanning(context.Background(), prefix, namespace, table, planId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.CancelPlanning``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 
**namespace** | **string** | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (&#x60;0x1F&#x60;) byte. | 
**table** | **string** | A table name | 
**planId** | **string** | ID used to track a planning request | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelPlanningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommitTransaction

> CommitTransaction(ctx, prefix).CommitTransactionRequest(commitTransactionRequest).Execute()

Commit updates to multiple tables in an atomic operation

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	commitTransactionRequest := *openapiclient.NewCommitTransactionRequest([]openapiclient.CommitTableRequest{*openapiclient.NewCommitTableRequest([]openapiclient.TableRequirement{*openapiclient.NewTableRequirement("Type_example")}, []openapiclient.TableUpdate{*openapiclient.NewTableUpdate("Action_example", int32(123), *openapiclient.NewSchema("Type_example", []openapiclient.StructField{*openapiclient.NewStructField(int32(123), "Name_example", openapiclient.Type{ListType: openapiclient.NewListType("Type_example", int32(123), openapiclient.Type{ListType: openapiclient.NewListType("Type_example", int32(123), openapiclient.Type{ListType: }, false)}, false)}, false)}), int32(123), *openapiclient.NewPartitionSpec([]openapiclient.PartitionField{*openapiclient.NewPartitionField(int32(123), "Name_example", "[identity, year, month, day, hour, bucket[256], truncate[16]]")}), int32(123), *openapiclient.NewSortOrder(int32(123), []openapiclient.SortField{*openapiclient.NewSortField(int32(123), "[identity, year, month, day, hour, bucket[256], truncate[16]]", openapiclient.SortDirection("asc"), openapiclient.NullOrder("nulls-first"))}), int32(123), *openapiclient.NewSnapshot(int64(123), int64(123), "ManifestList_example", *openapiclient.NewSnapshotSummary("Operation_example")), "RefName_example", "Type_example", int64(123), []int64{int64(123)}, "Location_example", map[string]string{"key": "Inner_example"}, []string{"Removals_example"}, *openapiclient.NewStatisticsFile(int64(123), "StatisticsPath_example", int64(123), int64(123), []openapiclient.BlobMetadata{*openapiclient.NewBlobMetadata("Type_example", int64(123), int64(123), []int32{int32(123)})}), []int32{int32(123)}, []int32{int32(123)}, "Uuid_example")})}) // CommitTransactionRequest | Commit updates to multiple tables in an atomic operation  A commit for a single table consists of a table identifier with requirements and updates. Requirements are assertions that will be validated before attempting to make and commit changes. For example, `assert-ref-snapshot-id` will check that a named ref's snapshot ID has a certain value. Server implementations are required to fail with a 400 status code if any unknown updates or requirements are received. Updates are changes to make to table metadata. For example, after asserting that the current main ref is at the expected snapshot, a commit may add a new child snapshot and set the ref to the new snapshot id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogAPIAPI.CommitTransaction(context.Background(), prefix).CommitTransactionRequest(commitTransactionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.CommitTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommitTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commitTransactionRequest** | [**CommitTransactionRequest**](CommitTransactionRequest.md) | Commit updates to multiple tables in an atomic operation  A commit for a single table consists of a table identifier with requirements and updates. Requirements are assertions that will be validated before attempting to make and commit changes. For example, &#x60;assert-ref-snapshot-id&#x60; will check that a named ref&#39;s snapshot ID has a certain value. Server implementations are required to fail with a 400 status code if any unknown updates or requirements are received. Updates are changes to make to table metadata. For example, after asserting that the current main ref is at the expected snapshot, a commit may add a new child snapshot and set the ref to the new snapshot id. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNamespace

> CreateNamespaceResponse CreateNamespace(ctx, prefix).CreateNamespaceRequest(createNamespaceRequest).Execute()

Create a namespace



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	createNamespaceRequest := *openapiclient.NewCreateNamespaceRequest([]string{"Namespace_example"}) // CreateNamespaceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogAPIAPI.CreateNamespace(context.Background(), prefix).CreateNamespaceRequest(createNamespaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.CreateNamespace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNamespace`: CreateNamespaceResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogAPIAPI.CreateNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNamespaceRequest** | [**CreateNamespaceRequest**](CreateNamespaceRequest.md) |  | 

### Return type

[**CreateNamespaceResponse**](CreateNamespaceResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTable

> LoadTableResult CreateTable(ctx, prefix, namespace).CreateTableRequest(createTableRequest).XIcebergAccessDelegation(xIcebergAccessDelegation).Execute()

Create a table in the given namespace



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	namespace := "accounting" // string | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (`0x1F`) byte.
	createTableRequest := *openapiclient.NewCreateTableRequest("Name_example", *openapiclient.NewSchema("Type_example", []openapiclient.StructField{*openapiclient.NewStructField(int32(123), "Name_example", openapiclient.Type{ListType: openapiclient.NewListType("Type_example", int32(123), openapiclient.Type{ListType: openapiclient.NewListType("Type_example", int32(123), openapiclient.Type{ListType: }, false)}, false)}, false)})) // CreateTableRequest | 
	xIcebergAccessDelegation := "vended-credentials,remote-signing" // string | Optional signal to the server that the client supports delegated access via a comma-separated list of access mechanisms.  The server may choose to supply access via any or none of the requested mechanisms.  Specific properties and handling for `vended-credentials` is documented in the `LoadTableResult` schema section of this spec document.  The protocol and specification for `remote-signing` is documented in  the `s3-signer-open-api.yaml` OpenApi spec in the `aws` module.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogAPIAPI.CreateTable(context.Background(), prefix, namespace).CreateTableRequest(createTableRequest).XIcebergAccessDelegation(xIcebergAccessDelegation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.CreateTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTable`: LoadTableResult
	fmt.Fprintf(os.Stdout, "Response from `CatalogAPIAPI.CreateTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 
**namespace** | **string** | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (&#x60;0x1F&#x60;) byte. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createTableRequest** | [**CreateTableRequest**](CreateTableRequest.md) |  | 
 **xIcebergAccessDelegation** | **string** | Optional signal to the server that the client supports delegated access via a comma-separated list of access mechanisms.  The server may choose to supply access via any or none of the requested mechanisms.  Specific properties and handling for &#x60;vended-credentials&#x60; is documented in the &#x60;LoadTableResult&#x60; schema section of this spec document.  The protocol and specification for &#x60;remote-signing&#x60; is documented in  the &#x60;s3-signer-open-api.yaml&#x60; OpenApi spec in the &#x60;aws&#x60; module.  | 

### Return type

[**LoadTableResult**](LoadTableResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateView

> LoadViewResult CreateView(ctx, prefix, namespace).CreateViewRequest(createViewRequest).Execute()

Create a view in the given namespace



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	namespace := "accounting" // string | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (`0x1F`) byte.
	createViewRequest := *openapiclient.NewCreateViewRequest("Name_example", *openapiclient.NewSchema("Type_example", []openapiclient.StructField{*openapiclient.NewStructField(int32(123), "Name_example", openapiclient.Type{ListType: openapiclient.NewListType("Type_example", int32(123), openapiclient.Type{ListType: openapiclient.NewListType("Type_example", int32(123), openapiclient.Type{ListType: }, false)}, false)}, false)}), *openapiclient.NewViewVersion(int32(123), int64(123), int32(123), map[string]string{"key": "Inner_example"}, []openapiclient.ViewRepresentation{openapiclient.ViewRepresentation{SQLViewRepresentation: openapiclient.NewSQLViewRepresentation("Type_example", "Sql_example", "Dialect_example")}}, []string{"DefaultNamespace_example"}), map[string]string{"key": "Inner_example"}) // CreateViewRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogAPIAPI.CreateView(context.Background(), prefix, namespace).CreateViewRequest(createViewRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.CreateView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateView`: LoadViewResult
	fmt.Fprintf(os.Stdout, "Response from `CatalogAPIAPI.CreateView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 
**namespace** | **string** | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (&#x60;0x1F&#x60;) byte. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createViewRequest** | [**CreateViewRequest**](CreateViewRequest.md) |  | 

### Return type

[**LoadViewResult**](LoadViewResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DropNamespace

> DropNamespace(ctx, prefix, namespace).Execute()

Drop a namespace from the catalog. Namespace must be empty.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	namespace := "accounting" // string | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (`0x1F`) byte.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogAPIAPI.DropNamespace(context.Background(), prefix, namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.DropNamespace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 
**namespace** | **string** | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (&#x60;0x1F&#x60;) byte. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDropNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DropTable

> DropTable(ctx, prefix, namespace, table).PurgeRequested(purgeRequested).Execute()

Drop a table from the catalog



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	namespace := "accounting" // string | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (`0x1F`) byte.
	table := "sales" // string | A table name
	purgeRequested := true // bool | Whether the user requested to purge the underlying table's data and metadata (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogAPIAPI.DropTable(context.Background(), prefix, namespace, table).PurgeRequested(purgeRequested).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.DropTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 
**namespace** | **string** | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (&#x60;0x1F&#x60;) byte. | 
**table** | **string** | A table name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDropTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **purgeRequested** | **bool** | Whether the user requested to purge the underlying table&#39;s data and metadata | [default to false]

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DropView

> DropView(ctx, prefix, namespace, view).Execute()

Drop a view from the catalog



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	namespace := "accounting" // string | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (`0x1F`) byte.
	view := "sales" // string | A view name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogAPIAPI.DropView(context.Background(), prefix, namespace, view).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.DropView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 
**namespace** | **string** | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (&#x60;0x1F&#x60;) byte. | 
**view** | **string** | A view name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDropViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPlanningResult

> FetchPlanningResult FetchPlanningResult(ctx, prefix, namespace, table, planId).Execute()

Fetches the result of scan planning for a plan-id



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	namespace := "accounting" // string | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (`0x1F`) byte.
	table := "sales" // string | A table name
	planId := "planId_example" // string | ID used to track a planning request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogAPIAPI.FetchPlanningResult(context.Background(), prefix, namespace, table, planId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.FetchPlanningResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchPlanningResult`: FetchPlanningResult
	fmt.Fprintf(os.Stdout, "Response from `CatalogAPIAPI.FetchPlanningResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 
**namespace** | **string** | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (&#x60;0x1F&#x60;) byte. | 
**table** | **string** | A table name | 
**planId** | **string** | ID used to track a planning request | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchPlanningResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**FetchPlanningResult**](FetchPlanningResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchScanTasks

> FetchScanTasksResult FetchScanTasks(ctx, prefix, namespace, table).FetchScanTasksRequest(fetchScanTasksRequest).Execute()

Fetches result tasks for a plan task



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	namespace := "accounting" // string | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (`0x1F`) byte.
	table := "sales" // string | A table name
	fetchScanTasksRequest := *openapiclient.NewFetchScanTasksRequest("PlanTask_example") // FetchScanTasksRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogAPIAPI.FetchScanTasks(context.Background(), prefix, namespace, table).FetchScanTasksRequest(fetchScanTasksRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.FetchScanTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchScanTasks`: FetchScanTasksResult
	fmt.Fprintf(os.Stdout, "Response from `CatalogAPIAPI.FetchScanTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 
**namespace** | **string** | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (&#x60;0x1F&#x60;) byte. | 
**table** | **string** | A table name | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchScanTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fetchScanTasksRequest** | [**FetchScanTasksRequest**](FetchScanTasksRequest.md) |  | 

### Return type

[**FetchScanTasksResult**](FetchScanTasksResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNamespaces

> ListNamespacesResponse ListNamespaces(ctx, prefix).PageToken(pageToken).PageSize(pageSize).Parent(parent).Execute()

List namespaces, optionally providing a parent namespace to list underneath



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	pageToken := "pageToken_example" // string |  (optional)
	pageSize := int32(56) // int32 | For servers that support pagination, this signals an upper bound of the number of results that a client will receive. For servers that do not support pagination, clients may receive results larger than the indicated `pageSize`. (optional)
	parent := "accounting%1Ftax" // string | An optional namespace, underneath which to list namespaces. If not provided or empty, all top-level namespaces should be listed. If parent is a multipart namespace, the parts must be separated by the unit separator (`0x1F`) byte. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogAPIAPI.ListNamespaces(context.Background(), prefix).PageToken(pageToken).PageSize(pageSize).Parent(parent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.ListNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNamespaces`: ListNamespacesResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogAPIAPI.ListNamespaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** |  | 
 **pageSize** | **int32** | For servers that support pagination, this signals an upper bound of the number of results that a client will receive. For servers that do not support pagination, clients may receive results larger than the indicated &#x60;pageSize&#x60;. | 
 **parent** | **string** | An optional namespace, underneath which to list namespaces. If not provided or empty, all top-level namespaces should be listed. If parent is a multipart namespace, the parts must be separated by the unit separator (&#x60;0x1F&#x60;) byte. | 

### Return type

[**ListNamespacesResponse**](ListNamespacesResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTables

> ListTablesResponse ListTables(ctx, prefix, namespace).PageToken(pageToken).PageSize(pageSize).Execute()

List all table identifiers underneath a given namespace



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	namespace := "accounting" // string | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (`0x1F`) byte.
	pageToken := "pageToken_example" // string |  (optional)
	pageSize := int32(56) // int32 | For servers that support pagination, this signals an upper bound of the number of results that a client will receive. For servers that do not support pagination, clients may receive results larger than the indicated `pageSize`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogAPIAPI.ListTables(context.Background(), prefix, namespace).PageToken(pageToken).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.ListTables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTables`: ListTablesResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogAPIAPI.ListTables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 
**namespace** | **string** | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (&#x60;0x1F&#x60;) byte. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageToken** | **string** |  | 
 **pageSize** | **int32** | For servers that support pagination, this signals an upper bound of the number of results that a client will receive. For servers that do not support pagination, clients may receive results larger than the indicated &#x60;pageSize&#x60;. | 

### Return type

[**ListTablesResponse**](ListTablesResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListViews

> ListTablesResponse ListViews(ctx, prefix, namespace).PageToken(pageToken).PageSize(pageSize).Execute()

List all view identifiers underneath a given namespace



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	namespace := "accounting" // string | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (`0x1F`) byte.
	pageToken := "pageToken_example" // string |  (optional)
	pageSize := int32(56) // int32 | For servers that support pagination, this signals an upper bound of the number of results that a client will receive. For servers that do not support pagination, clients may receive results larger than the indicated `pageSize`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogAPIAPI.ListViews(context.Background(), prefix, namespace).PageToken(pageToken).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.ListViews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListViews`: ListTablesResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogAPIAPI.ListViews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 
**namespace** | **string** | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (&#x60;0x1F&#x60;) byte. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageToken** | **string** |  | 
 **pageSize** | **int32** | For servers that support pagination, this signals an upper bound of the number of results that a client will receive. For servers that do not support pagination, clients may receive results larger than the indicated &#x60;pageSize&#x60;. | 

### Return type

[**ListTablesResponse**](ListTablesResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoadCredentials

> LoadCredentialsResponse LoadCredentials(ctx, prefix, namespace, table).Execute()

Load vended credentials for a table from the catalog



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	namespace := "accounting" // string | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (`0x1F`) byte.
	table := "sales" // string | A table name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogAPIAPI.LoadCredentials(context.Background(), prefix, namespace, table).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.LoadCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoadCredentials`: LoadCredentialsResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogAPIAPI.LoadCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 
**namespace** | **string** | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (&#x60;0x1F&#x60;) byte. | 
**table** | **string** | A table name | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**LoadCredentialsResponse**](LoadCredentialsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoadNamespaceMetadata

> GetNamespaceResponse LoadNamespaceMetadata(ctx, prefix, namespace).Execute()

Load the metadata properties for a namespace



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	namespace := "accounting" // string | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (`0x1F`) byte.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogAPIAPI.LoadNamespaceMetadata(context.Background(), prefix, namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.LoadNamespaceMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoadNamespaceMetadata`: GetNamespaceResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogAPIAPI.LoadNamespaceMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 
**namespace** | **string** | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (&#x60;0x1F&#x60;) byte. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadNamespaceMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNamespaceResponse**](GetNamespaceResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoadTable

> LoadTableResult LoadTable(ctx, prefix, namespace, table).XIcebergAccessDelegation(xIcebergAccessDelegation).IfNoneMatch(ifNoneMatch).Snapshots(snapshots).Execute()

Load a table from the catalog



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	namespace := "accounting" // string | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (`0x1F`) byte.
	table := "sales" // string | A table name
	xIcebergAccessDelegation := "vended-credentials,remote-signing" // string | Optional signal to the server that the client supports delegated access via a comma-separated list of access mechanisms.  The server may choose to supply access via any or none of the requested mechanisms.  Specific properties and handling for `vended-credentials` is documented in the `LoadTableResult` schema section of this spec document.  The protocol and specification for `remote-signing` is documented in  the `s3-signer-open-api.yaml` OpenApi spec in the `aws` module.  (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | An optional header that allows the server to return 304 (Not Modified) if the metadata is current. The content is the value of the ETag received in a CreateTableResponse or LoadTableResponse. (optional)
	snapshots := "snapshots_example" // string | The snapshots to return in the body of the metadata. Setting the value to `all` would return the full set of snapshots currently valid for the table. Setting the value to `refs` would load all snapshots referenced by branches or tags. Default if no param is provided is `all`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogAPIAPI.LoadTable(context.Background(), prefix, namespace, table).XIcebergAccessDelegation(xIcebergAccessDelegation).IfNoneMatch(ifNoneMatch).Snapshots(snapshots).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.LoadTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoadTable`: LoadTableResult
	fmt.Fprintf(os.Stdout, "Response from `CatalogAPIAPI.LoadTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 
**namespace** | **string** | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (&#x60;0x1F&#x60;) byte. | 
**table** | **string** | A table name | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIcebergAccessDelegation** | **string** | Optional signal to the server that the client supports delegated access via a comma-separated list of access mechanisms.  The server may choose to supply access via any or none of the requested mechanisms.  Specific properties and handling for &#x60;vended-credentials&#x60; is documented in the &#x60;LoadTableResult&#x60; schema section of this spec document.  The protocol and specification for &#x60;remote-signing&#x60; is documented in  the &#x60;s3-signer-open-api.yaml&#x60; OpenApi spec in the &#x60;aws&#x60; module.  | 
 **ifNoneMatch** | **string** | An optional header that allows the server to return 304 (Not Modified) if the metadata is current. The content is the value of the ETag received in a CreateTableResponse or LoadTableResponse. | 
 **snapshots** | **string** | The snapshots to return in the body of the metadata. Setting the value to &#x60;all&#x60; would return the full set of snapshots currently valid for the table. Setting the value to &#x60;refs&#x60; would load all snapshots referenced by branches or tags. Default if no param is provided is &#x60;all&#x60;. | 

### Return type

[**LoadTableResult**](LoadTableResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoadView

> LoadViewResult LoadView(ctx, prefix, namespace, view).Execute()

Load a view from the catalog



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	namespace := "accounting" // string | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (`0x1F`) byte.
	view := "sales" // string | A view name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogAPIAPI.LoadView(context.Background(), prefix, namespace, view).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.LoadView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoadView`: LoadViewResult
	fmt.Fprintf(os.Stdout, "Response from `CatalogAPIAPI.LoadView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 
**namespace** | **string** | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (&#x60;0x1F&#x60;) byte. | 
**view** | **string** | A view name | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**LoadViewResult**](LoadViewResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NamespaceExists

> NamespaceExists(ctx, prefix, namespace).Execute()

Check if a namespace exists



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	namespace := "accounting" // string | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (`0x1F`) byte.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogAPIAPI.NamespaceExists(context.Background(), prefix, namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.NamespaceExists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 
**namespace** | **string** | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (&#x60;0x1F&#x60;) byte. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNamespaceExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlanTableScan

> PlanTableScanResult PlanTableScan(ctx, prefix, namespace, table).PlanTableScanRequest(planTableScanRequest).Execute()

Submit a scan for planning



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	namespace := "accounting" // string | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (`0x1F`) byte.
	table := "sales" // string | A table name
	planTableScanRequest := *openapiclient.NewPlanTableScanRequest() // PlanTableScanRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogAPIAPI.PlanTableScan(context.Background(), prefix, namespace, table).PlanTableScanRequest(planTableScanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.PlanTableScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlanTableScan`: PlanTableScanResult
	fmt.Fprintf(os.Stdout, "Response from `CatalogAPIAPI.PlanTableScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 
**namespace** | **string** | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (&#x60;0x1F&#x60;) byte. | 
**table** | **string** | A table name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlanTableScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **planTableScanRequest** | [**PlanTableScanRequest**](PlanTableScanRequest.md) |  | 

### Return type

[**PlanTableScanResult**](PlanTableScanResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterTable

> LoadTableResult RegisterTable(ctx, prefix, namespace).RegisterTableRequest(registerTableRequest).Execute()

Register a table in the given namespace using given metadata file location



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	namespace := "accounting" // string | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (`0x1F`) byte.
	registerTableRequest := *openapiclient.NewRegisterTableRequest("Name_example", "MetadataLocation_example") // RegisterTableRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogAPIAPI.RegisterTable(context.Background(), prefix, namespace).RegisterTableRequest(registerTableRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.RegisterTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterTable`: LoadTableResult
	fmt.Fprintf(os.Stdout, "Response from `CatalogAPIAPI.RegisterTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 
**namespace** | **string** | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (&#x60;0x1F&#x60;) byte. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **registerTableRequest** | [**RegisterTableRequest**](RegisterTableRequest.md) |  | 

### Return type

[**LoadTableResult**](LoadTableResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenameTable

> RenameTable(ctx, prefix).RenameTableRequest(renameTableRequest).Execute()

Rename a table from its current name to a new name



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	renameTableRequest := *openapiclient.NewRenameTableRequest(*openapiclient.NewTableIdentifier([]string{"Namespace_example"}, "Name_example"), *openapiclient.NewTableIdentifier([]string{"Namespace_example"}, "Name_example")) // RenameTableRequest | Current table identifier to rename and new table identifier to rename to

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogAPIAPI.RenameTable(context.Background(), prefix).RenameTableRequest(renameTableRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.RenameTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenameTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **renameTableRequest** | [**RenameTableRequest**](RenameTableRequest.md) | Current table identifier to rename and new table identifier to rename to | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenameView

> RenameView(ctx, prefix).RenameTableRequest(renameTableRequest).Execute()

Rename a view from its current name to a new name



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	renameTableRequest := *openapiclient.NewRenameTableRequest(*openapiclient.NewTableIdentifier([]string{"Namespace_example"}, "Name_example"), *openapiclient.NewTableIdentifier([]string{"Namespace_example"}, "Name_example")) // RenameTableRequest | Current view identifier to rename and new view identifier to rename to

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogAPIAPI.RenameView(context.Background(), prefix).RenameTableRequest(renameTableRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.RenameView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenameViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **renameTableRequest** | [**RenameTableRequest**](RenameTableRequest.md) | Current view identifier to rename and new view identifier to rename to | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceView

> LoadViewResult ReplaceView(ctx, prefix, namespace, view).CommitViewRequest(commitViewRequest).Execute()

Replace a view



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	namespace := "accounting" // string | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (`0x1F`) byte.
	view := "sales" // string | A view name
	commitViewRequest := *openapiclient.NewCommitViewRequest([]openapiclient.ViewUpdate{*openapiclient.NewViewUpdate("Action_example", int32(123), *openapiclient.NewSchema("Type_example", []openapiclient.StructField{*openapiclient.NewStructField(int32(123), "Name_example", openapiclient.Type{ListType: openapiclient.NewListType("Type_example", int32(123), openapiclient.Type{ListType: openapiclient.NewListType("Type_example", int32(123), openapiclient.Type{ListType: }, false)}, false)}, false)}), "Location_example", map[string]string{"key": "Inner_example"}, []string{"Removals_example"}, *openapiclient.NewViewVersion(int32(123), int64(123), int32(123), map[string]string{"key": "Inner_example"}, []openapiclient.ViewRepresentation{openapiclient.ViewRepresentation{SQLViewRepresentation: openapiclient.NewSQLViewRepresentation("Type_example", "Sql_example", "Dialect_example")}}, []string{"DefaultNamespace_example"}), int32(123), "Uuid_example")}) // CommitViewRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogAPIAPI.ReplaceView(context.Background(), prefix, namespace, view).CommitViewRequest(commitViewRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.ReplaceView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceView`: LoadViewResult
	fmt.Fprintf(os.Stdout, "Response from `CatalogAPIAPI.ReplaceView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 
**namespace** | **string** | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (&#x60;0x1F&#x60;) byte. | 
**view** | **string** | A view name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **commitViewRequest** | [**CommitViewRequest**](CommitViewRequest.md) |  | 

### Return type

[**LoadViewResult**](LoadViewResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportMetrics

> ReportMetrics(ctx, prefix, namespace, table).ReportMetricsRequest(reportMetricsRequest).Execute()

Send a metrics report to this endpoint to be processed by the backend

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	namespace := "accounting" // string | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (`0x1F`) byte.
	table := "sales" // string | A table name
	reportMetricsRequest := *openapiclient.NewReportMetricsRequest("ReportType_example", "TableName_example", int64(123), openapiclient.Expression{AndOrExpression: openapiclient.NewAndOrExpression("[true, false, eq, and, or, not, in, not-in, lt, lt-eq, gt, gt-eq, not-eq, starts-with, not-starts-with, is-null, not-null, is-nan, not-nan]", openapiclient.Expression{AndOrExpression: openapiclient.NewAndOrExpression("[true, false, eq, and, or, not, in, not-in, lt, lt-eq, gt, gt-eq, not-eq, starts-with, not-starts-with, is-null, not-null, is-nan, not-nan]", openapiclient.Expression{AndOrExpression: }, openapiclient.Expression{AndOrExpression: })}, openapiclient.Expression{AndOrExpression: })}, int32(123), []int32{int32(123)}, []string{"ProjectedFieldNames_example"}, map[string]MetricResult{"key": *openapiclient.NewMetricResult("Unit_example", int64(123), "TimeUnit_example", int64(123), int64(123))}, int64(123), "Operation_example") // ReportMetricsRequest | The request containing the metrics report to be sent

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogAPIAPI.ReportMetrics(context.Background(), prefix, namespace, table).ReportMetricsRequest(reportMetricsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.ReportMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 
**namespace** | **string** | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (&#x60;0x1F&#x60;) byte. | 
**table** | **string** | A table name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reportMetricsRequest** | [**ReportMetricsRequest**](ReportMetricsRequest.md) | The request containing the metrics report to be sent | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TableExists

> TableExists(ctx, prefix, namespace, table).Execute()

Check if a table exists



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	namespace := "accounting" // string | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (`0x1F`) byte.
	table := "sales" // string | A table name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogAPIAPI.TableExists(context.Background(), prefix, namespace, table).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.TableExists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 
**namespace** | **string** | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (&#x60;0x1F&#x60;) byte. | 
**table** | **string** | A table name | 

### Other Parameters

Other parameters are passed through a pointer to a apiTableExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProperties

> UpdateNamespacePropertiesResponse UpdateProperties(ctx, prefix, namespace).UpdateNamespacePropertiesRequest(updateNamespacePropertiesRequest).Execute()

Set or remove properties on a namespace



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	namespace := "accounting" // string | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (`0x1F`) byte.
	updateNamespacePropertiesRequest := *openapiclient.NewUpdateNamespacePropertiesRequest() // UpdateNamespacePropertiesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogAPIAPI.UpdateProperties(context.Background(), prefix, namespace).UpdateNamespacePropertiesRequest(updateNamespacePropertiesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.UpdateProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProperties`: UpdateNamespacePropertiesResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogAPIAPI.UpdateProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 
**namespace** | **string** | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (&#x60;0x1F&#x60;) byte. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNamespacePropertiesRequest** | [**UpdateNamespacePropertiesRequest**](UpdateNamespacePropertiesRequest.md) |  | 

### Return type

[**UpdateNamespacePropertiesResponse**](UpdateNamespacePropertiesResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTable

> CommitTableResponse UpdateTable(ctx, prefix, namespace, table).CommitTableRequest(commitTableRequest).Execute()

Commit updates to a table



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	namespace := "accounting" // string | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (`0x1F`) byte.
	table := "sales" // string | A table name
	commitTableRequest := *openapiclient.NewCommitTableRequest([]openapiclient.TableRequirement{*openapiclient.NewTableRequirement("Type_example")}, []openapiclient.TableUpdate{*openapiclient.NewTableUpdate("Action_example", int32(123), *openapiclient.NewSchema("Type_example", []openapiclient.StructField{*openapiclient.NewStructField(int32(123), "Name_example", openapiclient.Type{ListType: openapiclient.NewListType("Type_example", int32(123), openapiclient.Type{ListType: openapiclient.NewListType("Type_example", int32(123), openapiclient.Type{ListType: }, false)}, false)}, false)}), int32(123), *openapiclient.NewPartitionSpec([]openapiclient.PartitionField{*openapiclient.NewPartitionField(int32(123), "Name_example", "[identity, year, month, day, hour, bucket[256], truncate[16]]")}), int32(123), *openapiclient.NewSortOrder(int32(123), []openapiclient.SortField{*openapiclient.NewSortField(int32(123), "[identity, year, month, day, hour, bucket[256], truncate[16]]", openapiclient.SortDirection("asc"), openapiclient.NullOrder("nulls-first"))}), int32(123), *openapiclient.NewSnapshot(int64(123), int64(123), "ManifestList_example", *openapiclient.NewSnapshotSummary("Operation_example")), "RefName_example", "Type_example", int64(123), []int64{int64(123)}, "Location_example", map[string]string{"key": "Inner_example"}, []string{"Removals_example"}, *openapiclient.NewStatisticsFile(int64(123), "StatisticsPath_example", int64(123), int64(123), []openapiclient.BlobMetadata{*openapiclient.NewBlobMetadata("Type_example", int64(123), int64(123), []int32{int32(123)})}), []int32{int32(123)}, []int32{int32(123)}, "Uuid_example")}) // CommitTableRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogAPIAPI.UpdateTable(context.Background(), prefix, namespace, table).CommitTableRequest(commitTableRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.UpdateTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTable`: CommitTableResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogAPIAPI.UpdateTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 
**namespace** | **string** | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (&#x60;0x1F&#x60;) byte. | 
**table** | **string** | A table name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **commitTableRequest** | [**CommitTableRequest**](CommitTableRequest.md) |  | 

### Return type

[**CommitTableResponse**](CommitTableResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewExists

> ViewExists(ctx, prefix, namespace, view).Execute()

Check if a view exists



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	namespace := "accounting" // string | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (`0x1F`) byte.
	view := "sales" // string | A view name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogAPIAPI.ViewExists(context.Background(), prefix, namespace, view).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.ViewExists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 
**namespace** | **string** | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (&#x60;0x1F&#x60;) byte. | 
**view** | **string** | A view name | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

