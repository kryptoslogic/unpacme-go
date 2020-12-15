# FeedEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autoit** | Pointer to **bool** | AutoIt detected | [optional] 
**Children** | Pointer to **int32** | Number of unpacked children | [optional] 
**Created** | Pointer to **float32** | Submission timestamp | [optional] 
**Deepmatch** | Pointer to [**[]DeepmatchEntity**](DeepmatchEntity.md) | DeepMatch matches | [optional] 
**Downloader** | Pointer to **bool** | Downloader detected | [optional] 
**Id** | Pointer to **string** | Unpacking submission ID | [optional] 
**Malwareid** | Pointer to [**[]MalwareIdShort**](MalwareIdShort.md) |  | [optional] 
**MalwareidRestricted** | Pointer to [**[]MalwareIdShort**](MalwareIdShort.md) |  | [optional] 
**Sha256** | Pointer to **string** | Parent submission SHA256 | [optional] 
**Status** | Pointer to **string** | Unpacking status | [optional] 

## Methods

### NewFeedEntity

`func NewFeedEntity() *FeedEntity`

NewFeedEntity instantiates a new FeedEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedEntityWithDefaults

`func NewFeedEntityWithDefaults() *FeedEntity`

NewFeedEntityWithDefaults instantiates a new FeedEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoit

`func (o *FeedEntity) GetAutoit() bool`

GetAutoit returns the Autoit field if non-nil, zero value otherwise.

### GetAutoitOk

`func (o *FeedEntity) GetAutoitOk() (*bool, bool)`

GetAutoitOk returns a tuple with the Autoit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoit

`func (o *FeedEntity) SetAutoit(v bool)`

SetAutoit sets Autoit field to given value.

### HasAutoit

`func (o *FeedEntity) HasAutoit() bool`

HasAutoit returns a boolean if a field has been set.

### GetChildren

`func (o *FeedEntity) GetChildren() int32`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *FeedEntity) GetChildrenOk() (*int32, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *FeedEntity) SetChildren(v int32)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *FeedEntity) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetCreated

`func (o *FeedEntity) GetCreated() float32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FeedEntity) GetCreatedOk() (*float32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FeedEntity) SetCreated(v float32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FeedEntity) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDeepmatch

`func (o *FeedEntity) GetDeepmatch() []DeepmatchEntity`

GetDeepmatch returns the Deepmatch field if non-nil, zero value otherwise.

### GetDeepmatchOk

`func (o *FeedEntity) GetDeepmatchOk() (*[]DeepmatchEntity, bool)`

GetDeepmatchOk returns a tuple with the Deepmatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeepmatch

`func (o *FeedEntity) SetDeepmatch(v []DeepmatchEntity)`

SetDeepmatch sets Deepmatch field to given value.

### HasDeepmatch

`func (o *FeedEntity) HasDeepmatch() bool`

HasDeepmatch returns a boolean if a field has been set.

### GetDownloader

`func (o *FeedEntity) GetDownloader() bool`

GetDownloader returns the Downloader field if non-nil, zero value otherwise.

### GetDownloaderOk

`func (o *FeedEntity) GetDownloaderOk() (*bool, bool)`

GetDownloaderOk returns a tuple with the Downloader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloader

`func (o *FeedEntity) SetDownloader(v bool)`

SetDownloader sets Downloader field to given value.

### HasDownloader

`func (o *FeedEntity) HasDownloader() bool`

HasDownloader returns a boolean if a field has been set.

### GetId

`func (o *FeedEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FeedEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FeedEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FeedEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMalwareid

`func (o *FeedEntity) GetMalwareid() []MalwareIdShort`

GetMalwareid returns the Malwareid field if non-nil, zero value otherwise.

### GetMalwareidOk

`func (o *FeedEntity) GetMalwareidOk() (*[]MalwareIdShort, bool)`

GetMalwareidOk returns a tuple with the Malwareid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMalwareid

`func (o *FeedEntity) SetMalwareid(v []MalwareIdShort)`

SetMalwareid sets Malwareid field to given value.

### HasMalwareid

`func (o *FeedEntity) HasMalwareid() bool`

HasMalwareid returns a boolean if a field has been set.

### GetMalwareidRestricted

`func (o *FeedEntity) GetMalwareidRestricted() []MalwareIdShort`

GetMalwareidRestricted returns the MalwareidRestricted field if non-nil, zero value otherwise.

### GetMalwareidRestrictedOk

`func (o *FeedEntity) GetMalwareidRestrictedOk() (*[]MalwareIdShort, bool)`

GetMalwareidRestrictedOk returns a tuple with the MalwareidRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMalwareidRestricted

`func (o *FeedEntity) SetMalwareidRestricted(v []MalwareIdShort)`

SetMalwareidRestricted sets MalwareidRestricted field to given value.

### HasMalwareidRestricted

`func (o *FeedEntity) HasMalwareidRestricted() bool`

HasMalwareidRestricted returns a boolean if a field has been set.

### GetSha256

`func (o *FeedEntity) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *FeedEntity) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *FeedEntity) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *FeedEntity) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetStatus

`func (o *FeedEntity) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FeedEntity) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FeedEntity) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FeedEntity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


