# \BuildStatusApi

All URIs are relative to *https://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBuildStatusesPaged**](BuildStatusApi.md#GetBuildStatusesPaged) | **Get** /rest/build-status/1.0/commits/{commitHash} | Get build statuses
[**PostBuildResult**](BuildStatusApi.md#PostBuildResult) | **Post** /rest/build-status/1.0/commits/{commitHash} | Post build-result



## GetBuildStatusesPaged

> BuildStatusPage GetBuildStatusesPaged(ctx, commitHash).OrderBy(orderBy).Limit(limit).Start(start).Execute()

Get build statuses



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
    commitHash := "commitHash_example" // string | 
    orderBy := "orderBy_example" // string | how the results should be ordered. Options are NEWEST, OLDEST, STATUS. Defaults to NEWEST if not provided. (optional)
    limit := 987 // int32 |  (optional)
    start := 987 // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BuildStatusApi.GetBuildStatusesPaged(context.Background(), commitHash).OrderBy(orderBy).Limit(limit).Start(start).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildStatusApi.GetBuildStatusesPaged``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuildStatusesPaged`: BuildStatusPage
    fmt.Fprintf(os.Stdout, "Response from `BuildStatusApi.GetBuildStatusesPaged`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commitHash** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuildStatusesPagedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderBy** | **string** | how the results should be ordered. Options are NEWEST, OLDEST, STATUS. Defaults to NEWEST if not provided. | 
 **limit** | **int32** |  | 
 **start** | **int32** |  | 

### Return type

[**BuildStatusPage**](buildStatusPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBuildResult

> PostBuildResult(ctx, commitHash).BuildStatus(buildStatus).Execute()

Post build-result



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
    commitHash := "commitHash_example" // string | 
    buildStatus := openapiclient.buildStatus{State: "State_example", Key: "Key_example", Name: "Name_example", Url: "Url_example", Description: "Description_example"} // BuildStatus | Associates a build status with a commit.  The state, the key and the url are mandatory. The name and description fields are optional.  All fields (mandatory or optional) are limited to 255 characters, except for the url, which is limited to 450 characters.  Supported values for the state are SUCCESSFUL, FAILED and INPROGRESS.  The authenticated user must have LICENSED permission or higher to call this resource. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BuildStatusApi.PostBuildResult(context.Background(), commitHash).BuildStatus(buildStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildStatusApi.PostBuildResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commitHash** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBuildResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **buildStatus** | [**BuildStatus**](BuildStatus.md) | Associates a build status with a commit.  The state, the key and the url are mandatory. The name and description fields are optional.  All fields (mandatory or optional) are limited to 255 characters, except for the url, which is limited to 450 characters.  Supported values for the state are SUCCESSFUL, FAILED and INPROGRESS.  The authenticated user must have LICENSED permission or higher to call this resource. | 

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

