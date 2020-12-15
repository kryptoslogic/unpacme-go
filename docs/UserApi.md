# \UserApi

All URIs are relative to *https://api.unpac.me/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePrivateUserMalpedia**](UserApi.md#DeletePrivateUserMalpedia) | **Delete** /private/user/malpedia | Remove Malpedia authentication
[**GetPrivateUserAccess**](UserApi.md#GetPrivateUserAccess) | **Get** /private/user/access | Get user settings
[**GetPrivateUserMalpedia**](UserApi.md#GetPrivateUserMalpedia) | **Get** /private/user/malpedia | Get user Malpedia info
[**PostPrivateUserMalpedia**](UserApi.md#PostPrivateUserMalpedia) | **Post** /private/user/malpedia | Authenticate user to Malpedia



## DeletePrivateUserMalpedia

> DeletePrivateUserMalpedia(ctx).Execute()

Remove Malpedia authentication



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
    resp, r, err := api_client.UserApi.DeletePrivateUserMalpedia(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.DeletePrivateUserMalpedia``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrivateUserMalpediaRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivateUserAccess

> UserAccess GetPrivateUserAccess(ctx).Execute()

Get user settings



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
    resp, r, err := api_client.UserApi.GetPrivateUserAccess(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetPrivateUserAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrivateUserAccess`: UserAccess
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetPrivateUserAccess`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateUserAccessRequest struct via the builder pattern


### Return type

[**UserAccess**](UserAccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivateUserMalpedia

> InlineResponse2001 GetPrivateUserMalpedia(ctx).Execute()

Get user Malpedia info



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
    resp, r, err := api_client.UserApi.GetPrivateUserMalpedia(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetPrivateUserMalpedia``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrivateUserMalpedia`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetPrivateUserMalpedia`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateUserMalpediaRequest struct via the builder pattern


### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPrivateUserMalpedia

> PostPrivateUserMalpedia(ctx).InlineObject1(inlineObject1).Execute()

Authenticate user to Malpedia



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
    inlineObject1 := *openapiclient.NewInlineObject1() // InlineObject1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.PostPrivateUserMalpedia(context.Background()).InlineObject1(inlineObject1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.PostPrivateUserMalpedia``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPrivateUserMalpediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject1** | [**InlineObject1**](InlineObject1.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

