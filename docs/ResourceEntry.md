# ResourceEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | Pointer to [**[]Resource**](Resource.md) |  | [optional] 
**Name** | Pointer to **string** | Resource name | [optional] 

## Methods

### NewResourceEntry

`func NewResourceEntry() *ResourceEntry`

NewResourceEntry instantiates a new ResourceEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceEntryWithDefaults

`func NewResourceEntryWithDefaults() *ResourceEntry`

NewResourceEntryWithDefaults instantiates a new ResourceEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *ResourceEntry) GetEntries() []Resource`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *ResourceEntry) GetEntriesOk() (*[]Resource, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *ResourceEntry) SetEntries(v []Resource)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *ResourceEntry) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetName

`func (o *ResourceEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceEntry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceEntry) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


