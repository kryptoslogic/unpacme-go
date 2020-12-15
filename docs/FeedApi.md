# \FeedApi

All URIs are relative to *https://api.unpac.me/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPrivateFeed**](FeedApi.md#GetPrivateFeed) | **Get** /private/feed/unpacked | Get full unpacked sample feed
[**GetPrivateFeedYaraFiltered**](FeedApi.md#GetPrivateFeedYaraFiltered) | **Get** /private/feed/unpacked/yara/{yara_rule} | Get full unpacked sample feed filtered by yara rule
[**GetPrivateFeedYaraTags**](FeedApi.md#GetPrivateFeedYaraTags) | **Get** /private/feed/unpacked/yara | Get list of yara tags in feed



## GetPrivateFeed

> PrivateFeed GetPrivateFeed(ctx).Cursor(cursor).Execute()

Get full unpacked sample feed



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
    cursor := int32(56) // int32 | Scroll feed to cursor (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeedApi.GetPrivateFeed(context.Background()).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeedApi.GetPrivateFeed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrivateFeed`: PrivateFeed
    fmt.Fprintf(os.Stdout, "Response from `FeedApi.GetPrivateFeed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **int32** | Scroll feed to cursor | 

### Return type

[**PrivateFeed**](PrivateFeed.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivateFeedYaraFiltered

> PrivateFeedFiltered GetPrivateFeedYaraFiltered(ctx, yaraRule).Execute()

Get full unpacked sample feed filtered by yara rule



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
    yaraRule := "yaraRule_example" // string | Yara rule name used to filter feed

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeedApi.GetPrivateFeedYaraFiltered(context.Background(), yaraRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeedApi.GetPrivateFeedYaraFiltered``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrivateFeedYaraFiltered`: PrivateFeedFiltered
    fmt.Fprintf(os.Stdout, "Response from `FeedApi.GetPrivateFeedYaraFiltered`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**yaraRule** | **string** | Yara rule name used to filter feed | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateFeedYaraFilteredRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrivateFeedFiltered**](PrivateFeedFiltered.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivateFeedYaraTags

> PrivateFeedYaraTags GetPrivateFeedYaraTags(ctx).Execute()

Get list of yara tags in feed



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeedApi.GetPrivateFeedYaraTags(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeedApi.GetPrivateFeedYaraTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrivateFeedYaraTags`: PrivateFeedYaraTags
    fmt.Fprintf(os.Stdout, "Response from `FeedApi.GetPrivateFeedYaraTags`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateFeedYaraTagsRequest struct via the builder pattern


### Return type

[**PrivateFeedYaraTags**](PrivateFeedYaraTags.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

