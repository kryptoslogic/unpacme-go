# PrivateFeedYaraTags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Numer of tags available | [optional] 
**Malwareid** | Pointer to **map[string]int32** | Yara tag names | [optional] 

## Methods

### NewPrivateFeedYaraTags

`func NewPrivateFeedYaraTags() *PrivateFeedYaraTags`

NewPrivateFeedYaraTags instantiates a new PrivateFeedYaraTags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateFeedYaraTagsWithDefaults

`func NewPrivateFeedYaraTagsWithDefaults() *PrivateFeedYaraTags`

NewPrivateFeedYaraTagsWithDefaults instantiates a new PrivateFeedYaraTags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PrivateFeedYaraTags) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PrivateFeedYaraTags) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PrivateFeedYaraTags) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PrivateFeedYaraTags) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetMalwareid

`func (o *PrivateFeedYaraTags) GetMalwareid() map[string]int32`

GetMalwareid returns the Malwareid field if non-nil, zero value otherwise.

### GetMalwareidOk

`func (o *PrivateFeedYaraTags) GetMalwareidOk() (*map[string]int32, bool)`

GetMalwareidOk returns a tuple with the Malwareid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMalwareid

`func (o *PrivateFeedYaraTags) SetMalwareid(v map[string]int32)`

SetMalwareid sets Malwareid field to given value.

### HasMalwareid

`func (o *PrivateFeedYaraTags) HasMalwareid() bool`

HasMalwareid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


