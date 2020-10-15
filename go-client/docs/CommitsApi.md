# \CommitsApi

All URIs are relative to *https://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCommit**](CommitsApi.md#GetCommit) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId} | Get commit
[**GetCommitsPaged**](CommitsApi.md#GetCommitsPaged) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/commits | Get commits



## GetCommit

> Commit GetCommit(ctx, projectKey, repositorySlug, commitId).Path(path).Execute()

Get commit



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
    commitId := "commitId_example" // string | 
    path := "path_example" // string | an optional path to filter the commit by. If supplied the details returned may not be for the specified commit. Instead, starting from the specified commit, they will be the details for the first commit affecting the specified path. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommitsApi.GetCommit(context.Background(), projectKey, repositorySlug, commitId).Path(path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.GetCommit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommit`: Commit
    fmt.Fprintf(os.Stdout, "Response from `CommitsApi.GetCommit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 
**commitId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **path** | **string** | an optional path to filter the commit by. If supplied the details returned may not be for the specified commit. Instead, starting from the specified commit, they will be the details for the first commit affecting the specified path. | 

### Return type

[**Commit**](commit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommitsPaged

> CommitsPage GetCommitsPaged(ctx, projectKey, repositorySlug).FollowRenames(followRenames).IgnoreMissing(ignoreMissing).Merges(merges).Path(path).Since(since).Until(until).WithCounts(withCounts).Limit(limit).Start(start).Execute()

Get commits



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
    followRenames := true // bool | if true, the commit history of the specified file will be followed past renames. Only valid for a path to a single file. (optional) (default to false)
    ignoreMissing := true // bool | true to ignore missing commits, false otherwise (optional) (default to false)
    merges := "merges_example" // string |   if present, controls how merge commits should be filtered. Can be either exclude, to exclude merge commits, include, to include both merge commits and non-merge commits or only, to only return merge commits. (optional)
    path := "path_example" // string | an optional path to filter commits by (optional)
    since := "since_example" // string | the commit ID or ref (exclusively) to retrieve commits after (optional)
    until := "until_example" // string | the commit ID (SHA1) or ref (inclusively) to retrieve commits before (optional)
    withCounts := true // bool | optionally include the total number of commits and total number of unique authors (optional) (default to false)
    limit := 987 // int32 |  (optional)
    start := 987 // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommitsApi.GetCommitsPaged(context.Background(), projectKey, repositorySlug).FollowRenames(followRenames).IgnoreMissing(ignoreMissing).Merges(merges).Path(path).Since(since).Until(until).WithCounts(withCounts).Limit(limit).Start(start).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.GetCommitsPaged``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommitsPaged`: CommitsPage
    fmt.Fprintf(os.Stdout, "Response from `CommitsApi.GetCommitsPaged`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommitsPagedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **followRenames** | **bool** | if true, the commit history of the specified file will be followed past renames. Only valid for a path to a single file. | [default to false]
 **ignoreMissing** | **bool** | true to ignore missing commits, false otherwise | [default to false]
 **merges** | **string** |   if present, controls how merge commits should be filtered. Can be either exclude, to exclude merge commits, include, to include both merge commits and non-merge commits or only, to only return merge commits. | 
 **path** | **string** | an optional path to filter commits by | 
 **since** | **string** | the commit ID or ref (exclusively) to retrieve commits after | 
 **until** | **string** | the commit ID (SHA1) or ref (inclusively) to retrieve commits before | 
 **withCounts** | **bool** | optionally include the total number of commits and total number of unique authors | [default to false]
 **limit** | **int32** |  | 
 **start** | **int32** |  | 

### Return type

[**CommitsPage**](commitsPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

