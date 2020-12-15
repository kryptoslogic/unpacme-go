# PrivateFeed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **int32** | Cursor used to scroll to page of feed results | [optional] 
**SubmissionCount** | Pointer to **interface{}** | Number of submissions in feed page | [optional] 
**Submissions** | Pointer to [**[]PrivateFeedEntity**](PrivateFeedEntity.md) |  | [optional] 

## Methods

### NewPrivateFeed

`func NewPrivateFeed() *PrivateFeed`

NewPrivateFeed instantiates a new PrivateFeed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateFeedWithDefaults

`func NewPrivateFeedWithDefaults() *PrivateFeed`

NewPrivateFeedWithDefaults instantiates a new PrivateFeed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *PrivateFeed) GetCursor() int32`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *PrivateFeed) GetCursorOk() (*int32, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *PrivateFeed) SetCursor(v int32)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *PrivateFeed) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetSubmissionCount

`func (o *PrivateFeed) GetSubmissionCount() interface{}`

GetSubmissionCount returns the SubmissionCount field if non-nil, zero value otherwise.

### GetSubmissionCountOk

`func (o *PrivateFeed) GetSubmissionCountOk() (*interface{}, bool)`

GetSubmissionCountOk returns a tuple with the SubmissionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionCount

`func (o *PrivateFeed) SetSubmissionCount(v interface{})`

SetSubmissionCount sets SubmissionCount field to given value.

### HasSubmissionCount

`func (o *PrivateFeed) HasSubmissionCount() bool`

HasSubmissionCount returns a boolean if a field has been set.

### SetSubmissionCountNil

`func (o *PrivateFeed) SetSubmissionCountNil(b bool)`

 SetSubmissionCountNil sets the value for SubmissionCount to be an explicit nil

### UnsetSubmissionCount
`func (o *PrivateFeed) UnsetSubmissionCount()`

UnsetSubmissionCount ensures that no value is present for SubmissionCount, not even an explicit nil
### GetSubmissions

`func (o *PrivateFeed) GetSubmissions() []PrivateFeedEntity`

GetSubmissions returns the Submissions field if non-nil, zero value otherwise.

### GetSubmissionsOk

`func (o *PrivateFeed) GetSubmissionsOk() (*[]PrivateFeedEntity, bool)`

GetSubmissionsOk returns a tuple with the Submissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissions

`func (o *PrivateFeed) SetSubmissions(v []PrivateFeedEntity)`

SetSubmissions sets Submissions field to given value.

### HasSubmissions

`func (o *PrivateFeed) HasSubmissions() bool`

HasSubmissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


