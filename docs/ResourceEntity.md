# ResourceEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | Pointer to [**[]ResourceEntry**](ResourceEntry.md) |  | [optional] 
**Type** | Pointer to **string** | Resource type | [optional] 

## Methods

### NewResourceEntity

`func NewResourceEntity() *ResourceEntity`

NewResourceEntity instantiates a new ResourceEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceEntityWithDefaults

`func NewResourceEntityWithDefaults() *ResourceEntity`

NewResourceEntityWithDefaults instantiates a new ResourceEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *ResourceEntity) GetEntries() []ResourceEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *ResourceEntity) GetEntriesOk() (*[]ResourceEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *ResourceEntity) SetEntries(v []ResourceEntry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *ResourceEntity) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetType

`func (o *ResourceEntity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceEntity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceEntity) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResourceEntity) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


