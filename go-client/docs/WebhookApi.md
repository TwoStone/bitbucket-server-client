# \WebhookApi

All URIs are relative to *https://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhook**](WebhookApi.md#CreateWebhook) | **Post** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks | Create webhook
[**DeleteWebhook**](WebhookApi.md#DeleteWebhook) | **Delete** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks/{webhookId} | Delete Webhook
[**GetWebhook**](WebhookApi.md#GetWebhook) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks/{webhookId} | Get Webhook
[**GetWebhooksPaged**](WebhookApi.md#GetWebhooksPaged) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks | Get webhooks
[**UpdateWebhook**](WebhookApi.md#UpdateWebhook) | **Put** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks/{webhookId} | Update webhook



## CreateWebhook

> Webhook CreateWebhook(ctx, projectKey, repositorySlug).Webhook(webhook).Execute()

Create webhook



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
    webhook := openapiclient.webhook{Id: 123, Name: "Name_example", CreateDate: int64(123), UpdatedDate: int64(123), Events: []WebhookEvent{openapiclient.webhookEvent{}), Configuration: openapiclient.webhook_configuration{Secret: "Secret_example"}, Url: "Url_example"} // Webhook |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhookApi.CreateWebhook(context.Background(), projectKey, repositorySlug).Webhook(webhook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.CreateWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWebhook`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhookApi.CreateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **webhook** | [**Webhook**](Webhook.md) |  | 

### Return type

[**Webhook**](webhook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhook

> DeleteWebhook(ctx, projectKey, repositorySlug, webhookId).Execute()

Delete Webhook



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
    webhookId := 987 // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhookApi.DeleteWebhook(context.Background(), projectKey, repositorySlug, webhookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.DeleteWebhook``: %v\n", err)
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
**webhookId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhook

> Webhook GetWebhook(ctx, projectKey, repositorySlug, webhookId).Execute()

Get Webhook



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
    webhookId := 987 // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhookApi.GetWebhook(context.Background(), projectKey, repositorySlug, webhookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.GetWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhook`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhookApi.GetWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 
**webhookId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Webhook**](webhook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhooksPaged

> WebhooksPage GetWebhooksPaged(ctx, projectKey, repositorySlug).Start(start).Limit(limit).Event(event).Execute()

Get webhooks



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
    start := 987 // int32 |  (optional)
    limit := 987 // int32 |  (optional)
    event := "event_example" // string | list of {@link com.atlassian.webhooks.WebhookEvent} ids to filter for (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhookApi.GetWebhooksPaged(context.Background(), projectKey, repositorySlug).Start(start).Limit(limit).Event(event).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.GetWebhooksPaged``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhooksPaged`: WebhooksPage
    fmt.Fprintf(os.Stdout, "Response from `WebhookApi.GetWebhooksPaged`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhooksPagedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int32** |  | 
 **limit** | **int32** |  | 
 **event** | **string** | list of {@link com.atlassian.webhooks.WebhookEvent} ids to filter for | 

### Return type

[**WebhooksPage**](webhooksPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhook

> Webhook UpdateWebhook(ctx, projectKey, repositorySlug, webhookId).Webhook(webhook).Execute()

Update webhook



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
    webhookId := 987 // int32 | 
    webhook := openapiclient.webhook{Id: 123, Name: "Name_example", CreateDate: int64(123), UpdatedDate: int64(123), Events: []WebhookEvent{openapiclient.webhookEvent{}), Configuration: openapiclient.webhook_configuration{Secret: "Secret_example"}, Url: "Url_example"} // Webhook |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhookApi.UpdateWebhook(context.Background(), projectKey, repositorySlug, webhookId).Webhook(webhook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.UpdateWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebhook`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhookApi.UpdateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 
**webhookId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **webhook** | [**Webhook**](Webhook.md) |  | 

### Return type

[**Webhook**](webhook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

