# Result

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hashes** | [**ResultAllOfHashes**](result_allOf_hashes.md) |  | 
**Analysis** | [**ResultAllOfAnalysis**](result_allOf_analysis.md) |  | 
**Deepmatch** | Pointer to [**[]DeepmatchEntity**](DeepmatchEntity.md) |  | [optional] 
**Detectit** | Pointer to [**[]DetectitEntity**](DetectitEntity.md) |  | [optional] 
**MalwareId** | Pointer to [**[]MalwareId**](MalwareId.md) |  | [optional] 
**MalwareIdRestricted** | Pointer to [**[]MalwareId**](MalwareId.md) |  | [optional] 
**Strings** | Pointer to [**ResultAllOfStrings**](result_allOf_strings.md) |  | [optional] 
**Urls** | Pointer to **[]string** |  | [optional] 

## Methods

### NewResult

`func NewResult(hashes ResultAllOfHashes, analysis ResultAllOfAnalysis, ) *Result`

NewResult instantiates a new Result object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultWithDefaults

`func NewResultWithDefaults() *Result`

NewResultWithDefaults instantiates a new Result object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHashes

`func (o *Result) GetHashes() ResultAllOfHashes`

GetHashes returns the Hashes field if non-nil, zero value otherwise.

### GetHashesOk

`func (o *Result) GetHashesOk() (*ResultAllOfHashes, bool)`

GetHashesOk returns a tuple with the Hashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashes

`func (o *Result) SetHashes(v ResultAllOfHashes)`

SetHashes sets Hashes field to given value.


### GetAnalysis

`func (o *Result) GetAnalysis() ResultAllOfAnalysis`

GetAnalysis returns the Analysis field if non-nil, zero value otherwise.

### GetAnalysisOk

`func (o *Result) GetAnalysisOk() (*ResultAllOfAnalysis, bool)`

GetAnalysisOk returns a tuple with the Analysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysis

`func (o *Result) SetAnalysis(v ResultAllOfAnalysis)`

SetAnalysis sets Analysis field to given value.


### GetDeepmatch

`func (o *Result) GetDeepmatch() []DeepmatchEntity`

GetDeepmatch returns the Deepmatch field if non-nil, zero value otherwise.

### GetDeepmatchOk

`func (o *Result) GetDeepmatchOk() (*[]DeepmatchEntity, bool)`

GetDeepmatchOk returns a tuple with the Deepmatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeepmatch

`func (o *Result) SetDeepmatch(v []DeepmatchEntity)`

SetDeepmatch sets Deepmatch field to given value.

### HasDeepmatch

`func (o *Result) HasDeepmatch() bool`

HasDeepmatch returns a boolean if a field has been set.

### GetDetectit

`func (o *Result) GetDetectit() []DetectitEntity`

GetDetectit returns the Detectit field if non-nil, zero value otherwise.

### GetDetectitOk

`func (o *Result) GetDetectitOk() (*[]DetectitEntity, bool)`

GetDetectitOk returns a tuple with the Detectit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectit

`func (o *Result) SetDetectit(v []DetectitEntity)`

SetDetectit sets Detectit field to given value.

### HasDetectit

`func (o *Result) HasDetectit() bool`

HasDetectit returns a boolean if a field has been set.

### GetMalwareId

`func (o *Result) GetMalwareId() []MalwareId`

GetMalwareId returns the MalwareId field if non-nil, zero value otherwise.

### GetMalwareIdOk

`func (o *Result) GetMalwareIdOk() (*[]MalwareId, bool)`

GetMalwareIdOk returns a tuple with the MalwareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMalwareId

`func (o *Result) SetMalwareId(v []MalwareId)`

SetMalwareId sets MalwareId field to given value.

### HasMalwareId

`func (o *Result) HasMalwareId() bool`

HasMalwareId returns a boolean if a field has been set.

### GetMalwareIdRestricted

`func (o *Result) GetMalwareIdRestricted() []MalwareId`

GetMalwareIdRestricted returns the MalwareIdRestricted field if non-nil, zero value otherwise.

### GetMalwareIdRestrictedOk

`func (o *Result) GetMalwareIdRestrictedOk() (*[]MalwareId, bool)`

GetMalwareIdRestrictedOk returns a tuple with the MalwareIdRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMalwareIdRestricted

`func (o *Result) SetMalwareIdRestricted(v []MalwareId)`

SetMalwareIdRestricted sets MalwareIdRestricted field to given value.

### HasMalwareIdRestricted

`func (o *Result) HasMalwareIdRestricted() bool`

HasMalwareIdRestricted returns a boolean if a field has been set.

### GetStrings

`func (o *Result) GetStrings() ResultAllOfStrings`

GetStrings returns the Strings field if non-nil, zero value otherwise.

### GetStringsOk

`func (o *Result) GetStringsOk() (*ResultAllOfStrings, bool)`

GetStringsOk returns a tuple with the Strings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrings

`func (o *Result) SetStrings(v ResultAllOfStrings)`

SetStrings sets Strings field to given value.

### HasStrings

`func (o *Result) HasStrings() bool`

HasStrings returns a boolean if a field has been set.

### GetUrls

`func (o *Result) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *Result) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *Result) SetUrls(v []string)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *Result) HasUrls() bool`

HasUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


