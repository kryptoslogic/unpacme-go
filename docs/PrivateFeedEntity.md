# PrivateFeedEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autoit** | Pointer to **bool** | AutoIt detected | [optional] 
**ChildCount** | Pointer to **int32** | Number of unpacked children | [optional] 
**Children** | Pointer to [**map[string]PrivateFeedEntityChildren**](private_feed_entity_children.md) |  | [optional] 
**Created** | Pointer to **float32** | Submission timestamp | [optional] 
**Downloader** | Pointer to **bool** | Downloader detected | [optional] 
**Id** | Pointer to **string** | Unpacking submission ID | [optional] 
**Malwareid** | Pointer to [**[]MalwareIdShort**](MalwareIdShort.md) |  | [optional] 
**SubmissionSha256** | Pointer to **string** | Parent submission SHA256 | [optional] 

## Methods

### NewPrivateFeedEntity

`func NewPrivateFeedEntity() *PrivateFeedEntity`

NewPrivateFeedEntity instantiates a new PrivateFeedEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateFeedEntityWithDefaults

`func NewPrivateFeedEntityWithDefaults() *PrivateFeedEntity`

NewPrivateFeedEntityWithDefaults instantiates a new PrivateFeedEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoit

`func (o *PrivateFeedEntity) GetAutoit() bool`

GetAutoit returns the Autoit field if non-nil, zero value otherwise.

### GetAutoitOk

`func (o *PrivateFeedEntity) GetAutoitOk() (*bool, bool)`

GetAutoitOk returns a tuple with the Autoit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoit

`func (o *PrivateFeedEntity) SetAutoit(v bool)`

SetAutoit sets Autoit field to given value.

### HasAutoit

`func (o *PrivateFeedEntity) HasAutoit() bool`

HasAutoit returns a boolean if a field has been set.

### GetChildCount

`func (o *PrivateFeedEntity) GetChildCount() int32`

GetChildCount returns the ChildCount field if non-nil, zero value otherwise.

### GetChildCountOk

`func (o *PrivateFeedEntity) GetChildCountOk() (*int32, bool)`

GetChildCountOk returns a tuple with the ChildCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildCount

`func (o *PrivateFeedEntity) SetChildCount(v int32)`

SetChildCount sets ChildCount field to given value.

### HasChildCount

`func (o *PrivateFeedEntity) HasChildCount() bool`

HasChildCount returns a boolean if a field has been set.

### GetChildren

`func (o *PrivateFeedEntity) GetChildren() map[string]PrivateFeedEntityChildren`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *PrivateFeedEntity) GetChildrenOk() (*map[string]PrivateFeedEntityChildren, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *PrivateFeedEntity) SetChildren(v map[string]PrivateFeedEntityChildren)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *PrivateFeedEntity) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetCreated

`func (o *PrivateFeedEntity) GetCreated() float32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PrivateFeedEntity) GetCreatedOk() (*float32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PrivateFeedEntity) SetCreated(v float32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PrivateFeedEntity) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDownloader

`func (o *PrivateFeedEntity) GetDownloader() bool`

GetDownloader returns the Downloader field if non-nil, zero value otherwise.

### GetDownloaderOk

`func (o *PrivateFeedEntity) GetDownloaderOk() (*bool, bool)`

GetDownloaderOk returns a tuple with the Downloader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloader

`func (o *PrivateFeedEntity) SetDownloader(v bool)`

SetDownloader sets Downloader field to given value.

### HasDownloader

`func (o *PrivateFeedEntity) HasDownloader() bool`

HasDownloader returns a boolean if a field has been set.

### GetId

`func (o *PrivateFeedEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateFeedEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateFeedEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PrivateFeedEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMalwareid

`func (o *PrivateFeedEntity) GetMalwareid() []MalwareIdShort`

GetMalwareid returns the Malwareid field if non-nil, zero value otherwise.

### GetMalwareidOk

`func (o *PrivateFeedEntity) GetMalwareidOk() (*[]MalwareIdShort, bool)`

GetMalwareidOk returns a tuple with the Malwareid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMalwareid

`func (o *PrivateFeedEntity) SetMalwareid(v []MalwareIdShort)`

SetMalwareid sets Malwareid field to given value.

### HasMalwareid

`func (o *PrivateFeedEntity) HasMalwareid() bool`

HasMalwareid returns a boolean if a field has been set.

### GetSubmissionSha256

`func (o *PrivateFeedEntity) GetSubmissionSha256() string`

GetSubmissionSha256 returns the SubmissionSha256 field if non-nil, zero value otherwise.

### GetSubmissionSha256Ok

`func (o *PrivateFeedEntity) GetSubmissionSha256Ok() (*string, bool)`

GetSubmissionSha256Ok returns a tuple with the SubmissionSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionSha256

`func (o *PrivateFeedEntity) SetSubmissionSha256(v string)`

SetSubmissionSha256 sets SubmissionSha256 field to given value.

### HasSubmissionSha256

`func (o *PrivateFeedEntity) HasSubmissionSha256() bool`

HasSubmissionSha256 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


