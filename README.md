# Go API client for openapi


# Introduction
Welcome to the UNPACME API! All the malware unpacking and file analysis features that you are familiar with on the [unpac.me](https://www.unpac.me/) website are available through our API. You can easily integrate our unpacker into your malware analysis pipeline and begin unpacking at scale!


# Authentication
The public UNPACME API is publicly available and can be accessed without authentication.

In order to use the private UNPACME API you must sign up for an account with UNPACME. Once you have a valid user account you can view your personal API key in your user profile. 

<SecurityDefinitions />

# Response Structure
When interacting with the UNPACME API, if the request was correctly handled, a <b>200</b> HTTP status code will be returned. The body of the response will usually be a JSON object (except for file downloads).

## Response Status Codes

Status Code  | Description | Notes
------------- | ------------- | -
200  | OK | The request was successful
400  | Bad Request | The request was somehow incorrect. This can be caused by missing arguments or arguments with wrong values.
401 | Unauthorized | The supplied credentials, if any, are not sufficient to access the resource
403 | Forbidden | The account does not have enough privileges to make the request.
404 | Not Found | The requested resource is not found
429  | Too Many Requests | The request frequency has exceeded one of the account quotas (minute, daily or monthly). Monthly quotas are reset on the 1st of the month at 00:00 UTC.
500 | Server Error | The server could not return the representation due to an internal server error


## Error Response

If an error has occurred while handling the request an error status code will be returend along with a JSON error message with the following properties.


Property  | Description
------------- | -------------
Error  | The error type
Description  | A more informative message

# Example Clients

The following clients can be used to interact with the UNPACME API directly and are provided as examples. These clients are community projects and are not maintained or developed by UNPACME. UNPACME makes no claim as to the safety of these clients, use at your own risk.

  - [UnpacMe Client](https://github.com/larsborn/UnpacMeClient) (Python)
  - [UnpacMe Library](https://github.com/R3MRUM/unpacme) (Python)

<br>


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./openapi"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.unpac.me/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*FeedApi* | [**GetPrivateFeed**](docs/FeedApi.md#getprivatefeed) | **Get** /private/feed/unpacked | Get full unpacked sample feed
*FeedApi* | [**GetPrivateFeedYaraFiltered**](docs/FeedApi.md#getprivatefeedyarafiltered) | **Get** /private/feed/unpacked/yara/{yara_rule} | Get full unpacked sample feed filtered by yara rule
*FeedApi* | [**GetPrivateFeedYaraTags**](docs/FeedApi.md#getprivatefeedyaratags) | **Get** /private/feed/unpacked/yara | Get list of yara tags in feed
*PublicApi* | [**GetPublicFeed**](docs/PublicApi.md#getpublicfeed) | **Get** /public/feed | Get public feed
*PublicApi* | [**GetPublicResults**](docs/PublicApi.md#getpublicresults) | **Get** /public/results/{unpack_id} | Get unpack results by ID
*PublicApi* | [**GetPublicUnpackStatus**](docs/PublicApi.md#getpublicunpackstatus) | **Get** /public/status/{unpack_id} | Get unpack status by ID
*UnpackingApi* | [**GetPrivateDownload**](docs/UnpackingApi.md#getprivatedownload) | **Get** /private/download/{sample_hash} | Download sample by hash
*UnpackingApi* | [**GetPrivateHistory**](docs/UnpackingApi.md#getprivatehistory) | **Get** /private/history | Get history
*UnpackingApi* | [**GetPrivateResults**](docs/UnpackingApi.md#getprivateresults) | **Get** /private/results/{unpack_id} | Get unpack results by ID
*UnpackingApi* | [**GetPrivateSearchbyHash**](docs/UnpackingApi.md#getprivatesearchbyhash) | **Get** /private/search/hash/{sample_hash} | Search for parent submission by hash
*UnpackingApi* | [**GetPrivateUnpackStatus**](docs/UnpackingApi.md#getprivateunpackstatus) | **Get** /private/status/{unpack_id} | Get unpack status by ID
*UnpackingApi* | [**PostPrivateUpload**](docs/UnpackingApi.md#postprivateupload) | **Post** /private/upload/ | Submit sample for unpacking
*UserApi* | [**DeletePrivateUserMalpedia**](docs/UserApi.md#deleteprivateusermalpedia) | **Delete** /private/user/malpedia | Remove Malpedia authentication
*UserApi* | [**GetPrivateUserAccess**](docs/UserApi.md#getprivateuseraccess) | **Get** /private/user/access | Get user settings
*UserApi* | [**GetPrivateUserMalpedia**](docs/UserApi.md#getprivateusermalpedia) | **Get** /private/user/malpedia | Get user Malpedia info
*UserApi* | [**PostPrivateUserMalpedia**](docs/UserApi.md#postprivateusermalpedia) | **Post** /private/user/malpedia | Authenticate user to Malpedia


## Documentation For Models

 - [DeepmatchEntity](docs/DeepmatchEntity.md)
 - [DeepmatchEntityAllOf](docs/DeepmatchEntityAllOf.md)
 - [DetectitEntity](docs/DetectitEntity.md)
 - [DetectitEntityAllOf](docs/DetectitEntityAllOf.md)
 - [Export](docs/Export.md)
 - [ExportAllOf](docs/ExportAllOf.md)
 - [FeedEntity](docs/FeedEntity.md)
 - [Function](docs/Function.md)
 - [FunctionAllOf](docs/FunctionAllOf.md)
 - [History](docs/History.md)
 - [HistoryEntity](docs/HistoryEntity.md)
 - [ImportEntity](docs/ImportEntity.md)
 - [ImportEntityAllOf](docs/ImportEntityAllOf.md)
 - [InlineObject](docs/InlineObject.md)
 - [InlineObject1](docs/InlineObject1.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [MalwareId](docs/MalwareId.md)
 - [MalwareIdAllOf](docs/MalwareIdAllOf.md)
 - [MalwareIdShort](docs/MalwareIdShort.md)
 - [MalwareIdShortAllOf](docs/MalwareIdShortAllOf.md)
 - [PrivateFeed](docs/PrivateFeed.md)
 - [PrivateFeedEntity](docs/PrivateFeedEntity.md)
 - [PrivateFeedEntityChildren](docs/PrivateFeedEntityChildren.md)
 - [PrivateFeedFiltered](docs/PrivateFeedFiltered.md)
 - [PrivateFeedYaraTags](docs/PrivateFeedYaraTags.md)
 - [PublicFeed](docs/PublicFeed.md)
 - [Resource](docs/Resource.md)
 - [ResourceAllOf](docs/ResourceAllOf.md)
 - [ResourceEntity](docs/ResourceEntity.md)
 - [ResourceEntry](docs/ResourceEntry.md)
 - [Result](docs/Result.md)
 - [ResultAllOf](docs/ResultAllOf.md)
 - [ResultAllOfAnalysis](docs/ResultAllOfAnalysis.md)
 - [ResultAllOfAnalysisExports](docs/ResultAllOfAnalysisExports.md)
 - [ResultAllOfAnalysisImports](docs/ResultAllOfAnalysisImports.md)
 - [ResultAllOfAnalysisMetadata](docs/ResultAllOfAnalysisMetadata.md)
 - [ResultAllOfAnalysisMetadataVersionInfo](docs/ResultAllOfAnalysisMetadataVersionInfo.md)
 - [ResultAllOfAnalysisMetadataVersionInfoStringInfo](docs/ResultAllOfAnalysisMetadataVersionInfoStringInfo.md)
 - [ResultAllOfAnalysisMetadataVersionInfoVarInfo](docs/ResultAllOfAnalysisMetadataVersionInfoVarInfo.md)
 - [ResultAllOfAnalysisRichHeaders](docs/ResultAllOfAnalysisRichHeaders.md)
 - [ResultAllOfHashes](docs/ResultAllOfHashes.md)
 - [ResultAllOfStrings](docs/ResultAllOfStrings.md)
 - [RichHeader](docs/RichHeader.md)
 - [RichHeaderAllOf](docs/RichHeaderAllOf.md)
 - [SearchEntity](docs/SearchEntity.md)
 - [SearchResults](docs/SearchResults.md)
 - [Section](docs/Section.md)
 - [SectionAllOf](docs/SectionAllOf.md)
 - [Status](docs/Status.md)
 - [UnpackResults](docs/UnpackResults.md)
 - [UnpackResultsAllOf](docs/UnpackResultsAllOf.md)
 - [UnpackStatus](docs/UnpackStatus.md)
 - [UnpackStatusAllOf](docs/UnpackStatusAllOf.md)
 - [UserAccess](docs/UserAccess.md)
 - [UserAccessAllOf](docs/UserAccessAllOf.md)


## Documentation For Authorization



### api_key

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


Kryptos Logic