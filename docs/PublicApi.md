# \PublicApi

All URIs are relative to *https://api.unpac.me/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPublicFeed**](PublicApi.md#GetPublicFeed) | **Get** /public/feed | Get public feed
[**GetPublicResults**](PublicApi.md#GetPublicResults) | **Get** /public/results/{unpack_id} | Get unpack results by ID
[**GetPublicUnpackStatus**](PublicApi.md#GetPublicUnpackStatus) | **Get** /public/status/{unpack_id} | Get unpack status by ID



## GetPublicFeed

> PublicFeed GetPublicFeed(ctx).Execute()

Get public feed



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/kryptoslogic/unpacme-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicApi.GetPublicFeed(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.GetPublicFeed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPublicFeed`: PublicFeed
    fmt.Fprintf(os.Stdout, "Response from `PublicApi.GetPublicFeed`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicFeedRequest struct via the builder pattern


### Return type

[**PublicFeed**](PublicFeed.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicResults

> UnpackResults GetPublicResults(ctx, unpackId).Execute()

Get unpack results by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/kryptoslogic/unpacme-go"
)

func main() {
    unpackId := "unpackId_example" // string | ID of unpacking submission

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicApi.GetPublicResults(context.Background(), unpackId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.GetPublicResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPublicResults`: UnpackResults
    fmt.Fprintf(os.Stdout, "Response from `PublicApi.GetPublicResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**unpackId** | **string** | ID of unpacking submission | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UnpackResults**](UnpackResults.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicUnpackStatus

> UnpackStatus GetPublicUnpackStatus(ctx, unpackId).Execute()

Get unpack status by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/kryptoslogic/unpacme-go"
)

func main() {
    unpackId := "unpackId_example" // string | ID of unpacking submission

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicApi.GetPublicUnpackStatus(context.Background(), unpackId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.GetPublicUnpackStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPublicUnpackStatus`: UnpackStatus
    fmt.Fprintf(os.Stdout, "Response from `PublicApi.GetPublicUnpackStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**unpackId** | **string** | ID of unpacking submission | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicUnpackStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UnpackStatus**](UnpackStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

