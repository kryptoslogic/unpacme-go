# \UnpackingApi

All URIs are relative to *https://api.unpac.me/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPrivateDownload**](UnpackingApi.md#GetPrivateDownload) | **Get** /private/download/{sample_hash} | Download sample by hash
[**GetPrivateHistory**](UnpackingApi.md#GetPrivateHistory) | **Get** /private/history | Get history
[**GetPrivateResults**](UnpackingApi.md#GetPrivateResults) | **Get** /private/results/{unpack_id} | Get unpack results by ID
[**GetPrivateSearchbyHash**](UnpackingApi.md#GetPrivateSearchbyHash) | **Get** /private/search/hash/{sample_hash} | Search for parent submission by hash
[**GetPrivateUnpackStatus**](UnpackingApi.md#GetPrivateUnpackStatus) | **Get** /private/status/{unpack_id} | Get unpack status by ID
[**PostPrivateUpload**](UnpackingApi.md#PostPrivateUpload) | **Post** /private/upload/ | Submit sample for unpacking



## GetPrivateDownload

> *os.File GetPrivateDownload(ctx, sampleHash).Execute()

Download sample by hash



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
    sampleHash := "sampleHash_example" // string | SHA256 hash of sample to download

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UnpackingApi.GetPrivateDownload(context.Background(), sampleHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnpackingApi.GetPrivateDownload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrivateDownload`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `UnpackingApi.GetPrivateDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sampleHash** | **string** | SHA256 hash of sample to download | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivateHistory

> History GetPrivateHistory(ctx).Cursor(cursor).Execute()

Get history



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
    cursor := int32(56) // int32 | Scroll history to cursor (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UnpackingApi.GetPrivateHistory(context.Background()).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnpackingApi.GetPrivateHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrivateHistory`: History
    fmt.Fprintf(os.Stdout, "Response from `UnpackingApi.GetPrivateHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **int32** | Scroll history to cursor | 

### Return type

[**History**](History.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivateResults

> UnpackResults GetPrivateResults(ctx, unpackId).Execute()

Get unpack results by ID



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
    unpackId := "unpackId_example" // string | ID of unpacking submission

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UnpackingApi.GetPrivateResults(context.Background(), unpackId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnpackingApi.GetPrivateResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrivateResults`: UnpackResults
    fmt.Fprintf(os.Stdout, "Response from `UnpackingApi.GetPrivateResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**unpackId** | **string** | ID of unpacking submission | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UnpackResults**](UnpackResults.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivateSearchbyHash

> SearchResults GetPrivateSearchbyHash(ctx, sampleHash).Execute()

Search for parent submission by hash



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
    sampleHash := "sampleHash_example" // string | SHA256 hash of parent sample

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UnpackingApi.GetPrivateSearchbyHash(context.Background(), sampleHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnpackingApi.GetPrivateSearchbyHash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrivateSearchbyHash`: SearchResults
    fmt.Fprintf(os.Stdout, "Response from `UnpackingApi.GetPrivateSearchbyHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sampleHash** | **string** | SHA256 hash of parent sample | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateSearchbyHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SearchResults**](SearchResults.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivateUnpackStatus

> UnpackStatus GetPrivateUnpackStatus(ctx, unpackId).Execute()

Get unpack status by ID



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
    unpackId := "unpackId_example" // string | ID of unpacking submission

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UnpackingApi.GetPrivateUnpackStatus(context.Background(), unpackId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnpackingApi.GetPrivateUnpackStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrivateUnpackStatus`: UnpackStatus
    fmt.Fprintf(os.Stdout, "Response from `UnpackingApi.GetPrivateUnpackStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**unpackId** | **string** | ID of unpacking submission | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateUnpackStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UnpackStatus**](UnpackStatus.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPrivateUpload

> InlineResponse200 PostPrivateUpload(ctx).Private(private).File(file).Execute()

Submit sample for unpacking



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
    private := true // bool | Mark sample as private (only available to PRO users) (optional) (default to false)
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UnpackingApi.PostPrivateUpload(context.Background()).Private(private).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnpackingApi.PostPrivateUpload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPrivateUpload`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `UnpackingApi.PostPrivateUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPrivateUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **private** | **bool** | Mark sample as private (only available to PRO users) | [default to false]
 **file** | ***os.File** |  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: multipart/form-data:
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

