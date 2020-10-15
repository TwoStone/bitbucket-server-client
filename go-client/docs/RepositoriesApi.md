# \RepositoriesApi

All URIs are relative to *https://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BrowseRepositoryPaged**](RepositoriesApi.md#BrowseRepositoryPaged) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/browse | browseRepository
[**BrowseRepositoryPathPaged**](RepositoriesApi.md#BrowseRepositoryPathPaged) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/browse/{path} | browseRepositoryPath
[**CreateRepository**](RepositoriesApi.md#CreateRepository) | **Post** /rest/api/1.0/projects/{projectKey}/repos | Create repository
[**GetConfiguredDefaultBranch**](RepositoriesApi.md#GetConfiguredDefaultBranch) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/default-branch | Get default branch
[**GetRepositoriesPaged**](RepositoriesApi.md#GetRepositoriesPaged) | **Get** /rest/api/1.0/projects/{projectKey}/repos | Get Repositories
[**GetRepository**](RepositoriesApi.md#GetRepository) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug} | Get Repository
[**SearchRepositoriesPaged**](RepositoriesApi.md#SearchRepositoriesPaged) | **Get** /rest/api/1.0/repos | Search repositories



## BrowseRepositoryPaged

> Directory BrowseRepositoryPaged(ctx, projectKey, repositorySlug).At(at).Start(start).Limit(limit).Execute()

browseRepository



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
    at := "at_example" // string | the commit ID or ref to retrieve the content for. (optional)
    start := 987 // int32 |  (optional)
    limit := 987 // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.BrowseRepositoryPaged(context.Background(), projectKey, repositorySlug).At(at).Start(start).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.BrowseRepositoryPaged``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BrowseRepositoryPaged`: Directory
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.BrowseRepositoryPaged`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBrowseRepositoryPagedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **at** | **string** | the commit ID or ref to retrieve the content for. | 
 **start** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**Directory**](directory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrowseRepositoryPathPaged

> FileOrDirectory BrowseRepositoryPathPaged(ctx, projectKey, repositorySlug, path).At(at).Start(start).Limit(limit).Execute()

browseRepositoryPath



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
    path := "path_example" // string | 
    at := "at_example" // string | the commit ID or ref to retrieve the content for. (optional)
    start := 987 // int32 |  (optional)
    limit := 987 // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.BrowseRepositoryPathPaged(context.Background(), projectKey, repositorySlug, path).At(at).Start(start).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.BrowseRepositoryPathPaged``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BrowseRepositoryPathPaged`: FileOrDirectory
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.BrowseRepositoryPathPaged`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBrowseRepositoryPathPagedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **at** | **string** | the commit ID or ref to retrieve the content for. | 
 **start** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**FileOrDirectory**](fileOrDirectory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRepository

> Repository CreateRepository(ctx, projectKey).CreateRepository(createRepository).Execute()

Create repository



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
    createRepository := openapiclient.createRepository{Name: "Name_example", ScmId: "ScmId_example", Forkable: false, DefaultBranch: "DefaultBranch_example", Public: false, Description: "Description_example"} // CreateRepository |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.CreateRepository(context.Background(), projectKey).CreateRepository(createRepository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.CreateRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRepository`: Repository
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.CreateRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createRepository** | [**CreateRepository**](CreateRepository.md) |  | 

### Return type

[**Repository**](repository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfiguredDefaultBranch

> DefaultBranch GetConfiguredDefaultBranch(ctx, projectKey, repositorySlug).Execute()

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
    resp, r, err := api_client.RepositoriesApi.GetConfiguredDefaultBranch(context.Background(), projectKey, repositorySlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.GetConfiguredDefaultBranch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfiguredDefaultBranch`: DefaultBranch
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.GetConfiguredDefaultBranch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfiguredDefaultBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DefaultBranch**](defaultBranch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositoriesPaged

> RepositoriesPage GetRepositoriesPaged(ctx, projectKey).Limit(limit).Start(start).Execute()

Get Repositories



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
    limit := 987 // int32 |  (optional)
    start := 987 // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.GetRepositoriesPaged(context.Background(), projectKey).Limit(limit).Start(start).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.GetRepositoriesPaged``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRepositoriesPaged`: RepositoriesPage
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.GetRepositoriesPaged`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoriesPagedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **start** | **int32** |  | 

### Return type

[**RepositoriesPage**](repositoriesPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository

> Repository GetRepository(ctx, projectKey, repositorySlug).Execute()

Get Repository



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
    resp, r, err := api_client.RepositoriesApi.GetRepository(context.Background(), projectKey, repositorySlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.GetRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRepository`: Repository
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.GetRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Repository**](repository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchRepositoriesPaged

> RepositoriesPage SearchRepositoriesPaged(ctx).Name(name).Projectname(projectname).Permission(permission).State(state).Visibility(visibility).Start(start).Limit(limit).Execute()

Search repositories



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
    name := "name_example" // string | (optional) if specified, this will limit the resulting repository list to ones whose name matches this parameter's value. The match will be done case-insensitive and any leading and/or trailing whitespace characters on the name parameter will be stripped. (optional)
    projectname := "projectname_example" // string | (optional) if specified, this will limit the resulting repository list to ones whose project's name matches this parameter's value. The match will be done case-insensitive and any leading and/or trailing whitespace characters on the projectname parameter will be stripped. (optional)
    permission := "permission_example" // string | (optional) if specified, it must be a valid repository permission level name and will limit the resulting repository list to ones that the requesting user has the specified permission level to. If not specified, the default implicit 'read' permission level will be assumed. The currently supported explicit permission values are REPO_READ, REPO_WRITE and REPO_ADMIN. (optional)
    state := "state_example" // string | (optional) if specified, it must be a valid repository state name and will limit the resulting repository list to ones that are in the specified state. The currently supported explicit state values are AVAILABLE, INITIALISING and INITIALISATION_FAILED. Available since 5.13 (optional)
    visibility := "visibility_example" // string | (optional) if specified, this will limit the resulting repository list based on the repositories visibility. Valid values are public or private. (optional)
    start := 987 // int32 |  (optional)
    limit := 987 // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.SearchRepositoriesPaged(context.Background(), ).Name(name).Projectname(projectname).Permission(permission).State(state).Visibility(visibility).Start(start).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.SearchRepositoriesPaged``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchRepositoriesPaged`: RepositoriesPage
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.SearchRepositoriesPaged`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRepositoriesPagedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | (optional) if specified, this will limit the resulting repository list to ones whose name matches this parameter&#39;s value. The match will be done case-insensitive and any leading and/or trailing whitespace characters on the name parameter will be stripped. | 
 **projectname** | **string** | (optional) if specified, this will limit the resulting repository list to ones whose project&#39;s name matches this parameter&#39;s value. The match will be done case-insensitive and any leading and/or trailing whitespace characters on the projectname parameter will be stripped. | 
 **permission** | **string** | (optional) if specified, it must be a valid repository permission level name and will limit the resulting repository list to ones that the requesting user has the specified permission level to. If not specified, the default implicit &#39;read&#39; permission level will be assumed. The currently supported explicit permission values are REPO_READ, REPO_WRITE and REPO_ADMIN. | 
 **state** | **string** | (optional) if specified, it must be a valid repository state name and will limit the resulting repository list to ones that are in the specified state. The currently supported explicit state values are AVAILABLE, INITIALISING and INITIALISATION_FAILED. Available since 5.13 | 
 **visibility** | **string** | (optional) if specified, this will limit the resulting repository list based on the repositories visibility. Valid values are public or private. | 
 **start** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**RepositoriesPage**](repositoriesPage.md)

### Authorization

[usernamePassword](../README.md#usernamePassword)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

