# \PullRequestsApi

All URIs are relative to *https://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePullRequest**](PullRequestsApi.md#CreatePullRequest) | **Post** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests | Create Pull Request
[**DeletePullRequest**](PullRequestsApi.md#DeletePullRequest) | **Delete** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId} | Delete pull request
[**GetDiff**](PullRequestsApi.md#GetDiff) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}.diff | Get PR Diff
[**GetPullRequest**](PullRequestsApi.md#GetPullRequest) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId} | Get pull request
[**GetPullRequestsPaged**](PullRequestsApi.md#GetPullRequestsPaged) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests | Get Pull Request Page
[**GetRestApi10ProjectsProjectKeyReposRepositorySlugPullRequestsPullRequestIdPatch**](PullRequestsApi.md#GetRestApi10ProjectsProjectKeyReposRepositorySlugPullRequestsPullRequestIdPatch) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}.patch | Get PR Patch
[**UpdatePullRequest**](PullRequestsApi.md#UpdatePullRequest) | **Put** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId} | Update pull request



## CreatePullRequest

> PullRequest CreatePullRequest(ctx, projectKey, repositorySlug).PullRequest(pullRequest).Execute()

Create Pull Request



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
    pullRequest := openapiclient.pullRequest{Id: 123, Version: 123, Title: "Title_example", Description: "Description_example", State: openapiclient.pullRequestState{}, Open: false, Closed: false, CreatedDate: int64(123), UpdatedDate: int64(123), FromRef: openapiclient.repositoryRef{Id: "Id_example", Repository: openapiclient.repositoryRef_repository{Slug: "Slug_example", Name: "Name_example", Project: openapiclient.repositoryRef_repository_project{Key: "Key_example"}}}, ToRef: openapiclient.repositoryRef{Id: "Id_example", Repository: openapiclient.repositoryRef_repository{Slug: "Slug_example", Name: "Name_example", Project: openapiclient.repositoryRef_repository_project{Key: "Key_example"}}}, Locked: false, Author: openapiclient.userRole{User: openapiclient.user{Id: 123, Name: "Name_example", EmailAddress: "EmailAddress_example", DisplayName: "DisplayName_example", Active: false, Slug: "Slug_example", Type: "Type_example"}, Role: "Role_example", Approved: false, Status: "Status_example"}, Reviewers: []UserRole{openapiclient.userRole{User: openapiclient.user{Id: 123, Name: "Name_example", EmailAddress: "EmailAddress_example", DisplayName: "DisplayName_example", Active: false, Slug: "Slug_example", Type: "Type_example"}, Role: "Role_example", Approved: false, Status: "Status_example"}), Participants: []UserRole{), Links: openapiclient.project_links{Self: []Link{openapiclient.link{Href: "Href_example", Name: "Name_example"})}} // PullRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullRequestsApi.CreatePullRequest(context.Background(), projectKey, repositorySlug).PullRequest(pullRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsApi.CreatePullRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePullRequest`: PullRequest
    fmt.Fprintf(os.Stdout, "Response from `PullRequestsApi.CreatePullRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePullRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pullRequest** | [**PullRequest**](PullRequest.md) |  | 

### Return type

[**PullRequest**](pullRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePullRequest

> DeletePullRequest(ctx, projectKey, repositorySlug, pullRequestId).PullRequestDelete(pullRequestDelete).Execute()

Delete pull request



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
    pullRequestId := 987 // int32 | 
    pullRequestDelete := openapiclient.pullRequestDelete{Version: 123} // PullRequestDelete |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullRequestsApi.DeletePullRequest(context.Background(), projectKey, repositorySlug, pullRequestId).PullRequestDelete(pullRequestDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsApi.DeletePullRequest``: %v\n", err)
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
**pullRequestId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePullRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pullRequestDelete** | [**PullRequestDelete**](PullRequestDelete.md) |  | 

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


## GetDiff

> string GetDiff(ctx, projectKey, repositorySlug, pullRequestId).ContextLines(contextLines).Whitespaces(whitespaces).Execute()

Get PR Diff



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
    pullRequestId := 987 // int32 | 
    contextLines := 987 // int32 | the number of context lines to include around added/removed lines in the diff (optional) (default to -1)
    whitespaces := "whitespaces_example" // string | optional whitespace flag which can be set to ignore-all (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullRequestsApi.GetDiff(context.Background(), projectKey, repositorySlug, pullRequestId).ContextLines(contextLines).Whitespaces(whitespaces).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsApi.GetDiff``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDiff`: string
    fmt.Fprintf(os.Stdout, "Response from `PullRequestsApi.GetDiff`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 
**pullRequestId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **contextLines** | **int32** | the number of context lines to include around added/removed lines in the diff | [default to -1]
 **whitespaces** | **string** | optional whitespace flag which can be set to ignore-all | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPullRequest

> PullRequest GetPullRequest(ctx, projectKey, repositorySlug, pullRequestId).Execute()

Get pull request



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
    pullRequestId := 987 // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullRequestsApi.GetPullRequest(context.Background(), projectKey, repositorySlug, pullRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsApi.GetPullRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPullRequest`: PullRequest
    fmt.Fprintf(os.Stdout, "Response from `PullRequestsApi.GetPullRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 
**pullRequestId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPullRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PullRequest**](pullRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPullRequestsPaged

> PullRequestsPage GetPullRequestsPaged(ctx, projectKey, repositorySlug).Limit(limit).Start(start).State(state).Order(order).At(at).Direction(direction).FilterText(filterText).Execute()

Get Pull Request Page



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
    limit := 987 // int32 |  (optional)
    start := 987 // int32 |  (optional)
    state := "state_example" // string | (optional, defaults to OPEN). Supply ALL to return pull request in any state. If a state is supplied only pull requests in the specified state will be returned. Either OPEN, DECLINED or MERGED (optional)
    order := "order_example" // string | (optional, defaults to NEWEST) the order to return pull requests in, either OLDEST (as in: \"oldest first\") or NEWEST. (optional)
    at := "at_example" // string | (optional) a fully-qualified branch ID to find pull requests to or from, such as refs/heads/master (optional)
    direction := "direction_example" // string | (optional, defaults to INCOMING) the direction relative to the specified repository. Either INCOMING or OUTGOING. (optional)
    filterText := "filterText_example" // string | (optional) If specified, only pull requests where the title or description contains the supplied string will be returned. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullRequestsApi.GetPullRequestsPaged(context.Background(), projectKey, repositorySlug).Limit(limit).Start(start).State(state).Order(order).At(at).Direction(direction).FilterText(filterText).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsApi.GetPullRequestsPaged``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPullRequestsPaged`: PullRequestsPage
    fmt.Fprintf(os.Stdout, "Response from `PullRequestsApi.GetPullRequestsPaged`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPullRequestsPagedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** |  | 
 **start** | **int32** |  | 
 **state** | **string** | (optional, defaults to OPEN). Supply ALL to return pull request in any state. If a state is supplied only pull requests in the specified state will be returned. Either OPEN, DECLINED or MERGED | 
 **order** | **string** | (optional, defaults to NEWEST) the order to return pull requests in, either OLDEST (as in: \&quot;oldest first\&quot;) or NEWEST. | 
 **at** | **string** | (optional) a fully-qualified branch ID to find pull requests to or from, such as refs/heads/master | 
 **direction** | **string** | (optional, defaults to INCOMING) the direction relative to the specified repository. Either INCOMING or OUTGOING. | 
 **filterText** | **string** | (optional) If specified, only pull requests where the title or description contains the supplied string will be returned. | 

### Return type

[**PullRequestsPage**](pullRequestsPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRestApi10ProjectsProjectKeyReposRepositorySlugPullRequestsPullRequestIdPatch

> string GetRestApi10ProjectsProjectKeyReposRepositorySlugPullRequestsPullRequestIdPatch(ctx, projectKey, repositorySlug, pullRequestId).Execute()

Get PR Patch



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
    pullRequestId := "pullRequestId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullRequestsApi.GetRestApi10ProjectsProjectKeyReposRepositorySlugPullRequestsPullRequestIdPatch(context.Background(), projectKey, repositorySlug, pullRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsApi.GetRestApi10ProjectsProjectKeyReposRepositorySlugPullRequestsPullRequestIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRestApi10ProjectsProjectKeyReposRepositorySlugPullRequestsPullRequestIdPatch`: string
    fmt.Fprintf(os.Stdout, "Response from `PullRequestsApi.GetRestApi10ProjectsProjectKeyReposRepositorySlugPullRequestsPullRequestIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 
**pullRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRestApi10ProjectsProjectKeyReposRepositorySlugPullRequestsPullRequestIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePullRequest

> PullRequest UpdatePullRequest(ctx, projectKey, repositorySlug, pullRequestId).PullRequestUpdate(pullRequestUpdate).Execute()

Update pull request



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
    pullRequestId := 987 // int32 | 
    pullRequestUpdate := openapiclient.pullRequestUpdate{Id: 123, Version: 123, Title: "Title_example", Description: "Description_example", Reviewers: []UserRole{), ToRef: } // PullRequestUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PullRequestsApi.UpdatePullRequest(context.Background(), projectKey, repositorySlug, pullRequestId).PullRequestUpdate(pullRequestUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsApi.UpdatePullRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePullRequest`: PullRequest
    fmt.Fprintf(os.Stdout, "Response from `PullRequestsApi.UpdatePullRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 
**pullRequestId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePullRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pullRequestUpdate** | [**PullRequestUpdate**](PullRequestUpdate.md) |  | 

### Return type

[**PullRequest**](pullRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

