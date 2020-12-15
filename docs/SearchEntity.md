# SearchEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **float32** | Submission timestamp | [optional] 
**SubmissionId** | Pointer to **string** | Unpacking submission ID | [optional] 
**Children** | Pointer to **[]string** |  | [optional] 
**Sha256** | Pointer to **string** | Parent submission SHA256 | [optional] 
**Status** | Pointer to **string** | Unpacking status | [optional] 
**Metadata** | Pointer to **string** | Sample metadata | [optional] 

## Methods

### NewSearchEntity

`func NewSearchEntity() *SearchEntity`

NewSearchEntity instantiates a new SearchEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchEntityWithDefaults

`func NewSearchEntityWithDefaults() *SearchEntity`

NewSearchEntityWithDefaults instantiates a new SearchEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *SearchEntity) GetCreated() float32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SearchEntity) GetCreatedOk() (*float32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SearchEntity) SetCreated(v float32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SearchEntity) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetSubmissionId

`func (o *SearchEntity) GetSubmissionId() string`

GetSubmissionId returns the SubmissionId field if non-nil, zero value otherwise.

### GetSubmissionIdOk

`func (o *SearchEntity) GetSubmissionIdOk() (*string, bool)`

GetSubmissionIdOk returns a tuple with the SubmissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionId

`func (o *SearchEntity) SetSubmissionId(v string)`

SetSubmissionId sets SubmissionId field to given value.

### HasSubmissionId

`func (o *SearchEntity) HasSubmissionId() bool`

HasSubmissionId returns a boolean if a field has been set.

### GetChildren

`func (o *SearchEntity) GetChildren() []string`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *SearchEntity) GetChildrenOk() (*[]string, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *SearchEntity) SetChildren(v []string)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *SearchEntity) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetSha256

`func (o *SearchEntity) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *SearchEntity) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *SearchEntity) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *SearchEntity) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetStatus

`func (o *SearchEntity) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SearchEntity) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SearchEntity) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SearchEntity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMetadata

`func (o *SearchEntity) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SearchEntity) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SearchEntity) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SearchEntity) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


