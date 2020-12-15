# History

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **int32** | Cursor used to scroll to page of results | [optional] 
**Results** | Pointer to [**[]HistoryEntity**](HistoryEntity.md) |  | [optional] 

## Methods

### NewHistory

`func NewHistory() *History`

NewHistory instantiates a new History object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryWithDefaults

`func NewHistoryWithDefaults() *History`

NewHistoryWithDefaults instantiates a new History object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *History) GetCursor() int32`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *History) GetCursorOk() (*int32, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *History) SetCursor(v int32)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *History) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetResults

`func (o *History) GetResults() []HistoryEntity`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *History) GetResultsOk() (*[]HistoryEntity, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *History) SetResults(v []HistoryEntity)`

SetResults sets Results field to given value.

### HasResults

`func (o *History) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


