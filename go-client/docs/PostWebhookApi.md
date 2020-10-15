# \PostWebhookApi

All URIs are relative to *https://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePostWebhook**](PostWebhookApi.md#DeletePostWebhook) | **Delete** /rest/webhook/1.0/projects/{projectKey}/repos/{repositorySlug}/configurations/{ID} | Delete post webhook
[**GetPostWebhooks**](PostWebhookApi.md#GetPostWebhooks) | **Get** /rest/webhook/1.0/projects/{projectKey}/repos/{repositorySlug}/configurations | Get Post Webhooks



## DeletePostWebhook

> DeletePostWebhook(ctx, projectKey, repositorySlug, iD).Execute()

Delete post webhook



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
    iD := 987 // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PostWebhookApi.DeletePostWebhook(context.Background(), projectKey, repositorySlug, iD).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostWebhookApi.DeletePostWebhook``: %v\n", err)
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
**iD** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePostWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPostWebhooks

> []PostWebhook GetPostWebhooks(ctx, projectKey, repositorySlug).Execute()

Get Post Webhooks



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
    resp, r, err := api_client.PostWebhookApi.GetPostWebhooks(context.Background(), projectKey, repositorySlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostWebhookApi.GetPostWebhooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostWebhooks`: []PostWebhook
    fmt.Fprintf(os.Stdout, "Response from `PostWebhookApi.GetPostWebhooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPostWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]PostWebhook**](postWebhook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

