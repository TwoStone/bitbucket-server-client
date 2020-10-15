# \BranchesApi

All URIs are relative to *https://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBranch**](BranchesApi.md#CreateBranch) | **Post** /rest/branch-utils/1.0/projects/{projectKey}/repos/{repositorySlug}/branches | Delete branch
[**DeleteBranch**](BranchesApi.md#DeleteBranch) | **Delete** /rest/branch-utils/1.0/projects/{projectKey}/repos/{repositorySlug}/branches | 
[**GetBranchesPaged**](BranchesApi.md#GetBranchesPaged) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/branches | Get Branches
[**GetDefaultBranch**](BranchesApi.md#GetDefaultBranch) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/branches/default | Get default branch



## CreateBranch

> Branch CreateBranch(ctx, projectKey, repositorySlug).CreateBranch(createBranch).Execute()

Delete branch



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectKey := "projectKey_example" // string | 
    repositorySlug := "repositorySlug_example" // string | 
    createBranch := openapiclient.createBranch{Name: "Name_example", StartPoint: "StartPoint_example", Message: "Message_example"} // CreateBranch |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BranchesApi.CreateBranch(context.Background(), projectKey, repositorySlug).CreateBranch(createBranch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BranchesApi.CreateBranch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBranch`: Branch
    fmt.Fprintf(os.Stdout, "Response from `BranchesApi.CreateBranch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createBranch** | [**CreateBranch**](CreateBranch.md) |  | 

### Return type

[**Branch**](branch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBranch

> DeleteBranch(ctx, projectKey, repositorySlug).DeleteBranch(deleteBranch).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectKey := "projectKey_example" // string | 
    repositorySlug := "repositorySlug_example" // string | 
    deleteBranch := openapiclient.deleteBranch{Name: "Name_example", DryRun: false, EndPoint: "EndPoint_example"} // DeleteBranch |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BranchesApi.DeleteBranch(context.Background(), projectKey, repositorySlug).DeleteBranch(deleteBranch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BranchesApi.DeleteBranch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleteBranch** | [**DeleteBranch**](DeleteBranch.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBranchesPaged

> BranchesPage GetBranchesPaged(ctx, projectKey, repositorySlug).Base(base).Details(details).FilterText(filterText).OrderBy(orderBy).BoostMatches(boostMatches).Start(start).Limit(limit).Execute()

Get Branches



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectKey := "projectKey_example" // string | 
    repositorySlug := "repositorySlug_example" // string | 
    base := "base_example" // string | base branch or tag to compare each branch to (for the metadata providers that uses that information) (optional)
    details := true // bool | whether to retrieve plugin-provided metadata about each branch (optional)
    filterText := "filterText_example" // string | the text to match on (optional)
    orderBy := "orderBy_example" // string | ordering of refs either ALPHABETICAL (by name) or MODIFICATION (last updated) (optional)
    boostMatches := true // bool | controls whether exact and prefix matches will be boosted to the top (optional)
    start := 987 // int32 |  (optional)
    limit := 987 // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BranchesApi.GetBranchesPaged(context.Background(), projectKey, repositorySlug).Base(base).Details(details).FilterText(filterText).OrderBy(orderBy).BoostMatches(boostMatches).Start(start).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BranchesApi.GetBranchesPaged``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBranchesPaged`: BranchesPage
    fmt.Fprintf(os.Stdout, "Response from `BranchesApi.GetBranchesPaged`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBranchesPagedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **base** | **string** | base branch or tag to compare each branch to (for the metadata providers that uses that information) | 
 **details** | **bool** | whether to retrieve plugin-provided metadata about each branch | 
 **filterText** | **string** | the text to match on | 
 **orderBy** | **string** | ordering of refs either ALPHABETICAL (by name) or MODIFICATION (last updated) | 
 **boostMatches** | **bool** | controls whether exact and prefix matches will be boosted to the top | 
 **start** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**BranchesPage**](branchesPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultBranch

> Branch GetDefaultBranch(ctx, projectKey, repositorySlug).Execute()

Get default branch



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectKey := "projectKey_example" // string | 
    repositorySlug := "repositorySlug_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BranchesApi.GetDefaultBranch(context.Background(), projectKey, repositorySlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BranchesApi.GetDefaultBranch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultBranch`: Branch
    fmt.Fprintf(os.Stdout, "Response from `BranchesApi.GetDefaultBranch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Branch**](branch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

