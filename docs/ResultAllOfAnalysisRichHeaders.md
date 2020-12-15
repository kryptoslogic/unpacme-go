# ResultAllOfAnalysisRichHeaders

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checksum** | Pointer to **string** | Rich header checksum | [optional] 
**Entries** | Pointer to [**[]RichHeader**](RichHeader.md) |  | [optional] 

## Methods

### NewResultAllOfAnalysisRichHeaders

`func NewResultAllOfAnalysisRichHeaders() *ResultAllOfAnalysisRichHeaders`

NewResultAllOfAnalysisRichHeaders instantiates a new ResultAllOfAnalysisRichHeaders object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultAllOfAnalysisRichHeadersWithDefaults

`func NewResultAllOfAnalysisRichHeadersWithDefaults() *ResultAllOfAnalysisRichHeaders`

NewResultAllOfAnalysisRichHeadersWithDefaults instantiates a new ResultAllOfAnalysisRichHeaders object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecksum

`func (o *ResultAllOfAnalysisRichHeaders) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *ResultAllOfAnalysisRichHeaders) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *ResultAllOfAnalysisRichHeaders) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *ResultAllOfAnalysisRichHeaders) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetEntries

`func (o *ResultAllOfAnalysisRichHeaders) GetEntries() []RichHeader`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *ResultAllOfAnalysisRichHeaders) GetEntriesOk() (*[]RichHeader, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *ResultAllOfAnalysisRichHeaders) SetEntries(v []RichHeader)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *ResultAllOfAnalysisRichHeaders) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


