# PrivateFeedFiltered

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **int32** | Cursor used to scroll to page of feed results | [optional] 
**SubmissionCount** | Pointer to **interface{}** | Number of submissions in feed page | [optional] 
**TagType** | Pointer to **string** | Type of tag used to filter feed | [optional] 
**TagValue** | Pointer to **string** | Tag label | [optional] 
**Submissions** | Pointer to [**[]PrivateFeedEntity**](PrivateFeedEntity.md) |  | [optional] 

## Methods

### NewPrivateFeedFiltered

`func NewPrivateFeedFiltered() *PrivateFeedFiltered`

NewPrivateFeedFiltered instantiates a new PrivateFeedFiltered object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateFeedFilteredWithDefaults

`func NewPrivateFeedFilteredWithDefaults() *PrivateFeedFiltered`

NewPrivateFeedFilteredWithDefaults instantiates a new PrivateFeedFiltered object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *PrivateFeedFiltered) GetCursor() int32`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *PrivateFeedFiltered) GetCursorOk() (*int32, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *PrivateFeedFiltered) SetCursor(v int32)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *PrivateFeedFiltered) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetSubmissionCount

`func (o *PrivateFeedFiltered) GetSubmissionCount() interface{}`

GetSubmissionCount returns the SubmissionCount field if non-nil, zero value otherwise.

### GetSubmissionCountOk

`func (o *PrivateFeedFiltered) GetSubmissionCountOk() (*interface{}, bool)`

GetSubmissionCountOk returns a tuple with the SubmissionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionCount

`func (o *PrivateFeedFiltered) SetSubmissionCount(v interface{})`

SetSubmissionCount sets SubmissionCount field to given value.

### HasSubmissionCount

`func (o *PrivateFeedFiltered) HasSubmissionCount() bool`

HasSubmissionCount returns a boolean if a field has been set.

### SetSubmissionCountNil

`func (o *PrivateFeedFiltered) SetSubmissionCountNil(b bool)`

 SetSubmissionCountNil sets the value for SubmissionCount to be an explicit nil

### UnsetSubmissionCount
`func (o *PrivateFeedFiltered) UnsetSubmissionCount()`

UnsetSubmissionCount ensures that no value is present for SubmissionCount, not even an explicit nil
### GetTagType

`func (o *PrivateFeedFiltered) GetTagType() string`

GetTagType returns the TagType field if non-nil, zero value otherwise.

### GetTagTypeOk

`func (o *PrivateFeedFiltered) GetTagTypeOk() (*string, bool)`

GetTagTypeOk returns a tuple with the TagType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagType

`func (o *PrivateFeedFiltered) SetTagType(v string)`

SetTagType sets TagType field to given value.

### HasTagType

`func (o *PrivateFeedFiltered) HasTagType() bool`

HasTagType returns a boolean if a field has been set.

### GetTagValue

`func (o *PrivateFeedFiltered) GetTagValue() string`

GetTagValue returns the TagValue field if non-nil, zero value otherwise.

### GetTagValueOk

`func (o *PrivateFeedFiltered) GetTagValueOk() (*string, bool)`

GetTagValueOk returns a tuple with the TagValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagValue

`func (o *PrivateFeedFiltered) SetTagValue(v string)`

SetTagValue sets TagValue field to given value.

### HasTagValue

`func (o *PrivateFeedFiltered) HasTagValue() bool`

HasTagValue returns a boolean if a field has been set.

### GetSubmissions

`func (o *PrivateFeedFiltered) GetSubmissions() []PrivateFeedEntity`

GetSubmissions returns the Submissions field if non-nil, zero value otherwise.

### GetSubmissionsOk

`func (o *PrivateFeedFiltered) GetSubmissionsOk() (*[]PrivateFeedEntity, bool)`

GetSubmissionsOk returns a tuple with the Submissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissions

`func (o *PrivateFeedFiltered) SetSubmissions(v []PrivateFeedEntity)`

SetSubmissions sets Submissions field to given value.

### HasSubmissions

`func (o *PrivateFeedFiltered) HasSubmissions() bool`

HasSubmissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


