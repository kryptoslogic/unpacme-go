# ResultAllOf

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

### NewResultAllOf

`func NewResultAllOf(hashes ResultAllOfHashes, analysis ResultAllOfAnalysis, ) *ResultAllOf`

NewResultAllOf instantiates a new ResultAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultAllOfWithDefaults

`func NewResultAllOfWithDefaults() *ResultAllOf`

NewResultAllOfWithDefaults instantiates a new ResultAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHashes

`func (o *ResultAllOf) GetHashes() ResultAllOfHashes`

GetHashes returns the Hashes field if non-nil, zero value otherwise.

### GetHashesOk

`func (o *ResultAllOf) GetHashesOk() (*ResultAllOfHashes, bool)`

GetHashesOk returns a tuple with the Hashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashes

`func (o *ResultAllOf) SetHashes(v ResultAllOfHashes)`

SetHashes sets Hashes field to given value.


### GetAnalysis

`func (o *ResultAllOf) GetAnalysis() ResultAllOfAnalysis`

GetAnalysis returns the Analysis field if non-nil, zero value otherwise.

### GetAnalysisOk

`func (o *ResultAllOf) GetAnalysisOk() (*ResultAllOfAnalysis, bool)`

GetAnalysisOk returns a tuple with the Analysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysis

`func (o *ResultAllOf) SetAnalysis(v ResultAllOfAnalysis)`

SetAnalysis sets Analysis field to given value.


### GetDeepmatch

`func (o *ResultAllOf) GetDeepmatch() []DeepmatchEntity`

GetDeepmatch returns the Deepmatch field if non-nil, zero value otherwise.

### GetDeepmatchOk

`func (o *ResultAllOf) GetDeepmatchOk() (*[]DeepmatchEntity, bool)`

GetDeepmatchOk returns a tuple with the Deepmatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeepmatch

`func (o *ResultAllOf) SetDeepmatch(v []DeepmatchEntity)`

SetDeepmatch sets Deepmatch field to given value.

### HasDeepmatch

`func (o *ResultAllOf) HasDeepmatch() bool`

HasDeepmatch returns a boolean if a field has been set.

### GetDetectit

`func (o *ResultAllOf) GetDetectit() []DetectitEntity`

GetDetectit returns the Detectit field if non-nil, zero value otherwise.

### GetDetectitOk

`func (o *ResultAllOf) GetDetectitOk() (*[]DetectitEntity, bool)`

GetDetectitOk returns a tuple with the Detectit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectit

`func (o *ResultAllOf) SetDetectit(v []DetectitEntity)`

SetDetectit sets Detectit field to given value.

### HasDetectit

`func (o *ResultAllOf) HasDetectit() bool`

HasDetectit returns a boolean if a field has been set.

### GetMalwareId

`func (o *ResultAllOf) GetMalwareId() []MalwareId`

GetMalwareId returns the MalwareId field if non-nil, zero value otherwise.

### GetMalwareIdOk

`func (o *ResultAllOf) GetMalwareIdOk() (*[]MalwareId, bool)`

GetMalwareIdOk returns a tuple with the MalwareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMalwareId

`func (o *ResultAllOf) SetMalwareId(v []MalwareId)`

SetMalwareId sets MalwareId field to given value.

### HasMalwareId

`func (o *ResultAllOf) HasMalwareId() bool`

HasMalwareId returns a boolean if a field has been set.

### GetMalwareIdRestricted

`func (o *ResultAllOf) GetMalwareIdRestricted() []MalwareId`

GetMalwareIdRestricted returns the MalwareIdRestricted field if non-nil, zero value otherwise.

### GetMalwareIdRestrictedOk

`func (o *ResultAllOf) GetMalwareIdRestrictedOk() (*[]MalwareId, bool)`

GetMalwareIdRestrictedOk returns a tuple with the MalwareIdRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMalwareIdRestricted

`func (o *ResultAllOf) SetMalwareIdRestricted(v []MalwareId)`

SetMalwareIdRestricted sets MalwareIdRestricted field to given value.

### HasMalwareIdRestricted

`func (o *ResultAllOf) HasMalwareIdRestricted() bool`

HasMalwareIdRestricted returns a boolean if a field has been set.

### GetStrings

`func (o *ResultAllOf) GetStrings() ResultAllOfStrings`

GetStrings returns the Strings field if non-nil, zero value otherwise.

### GetStringsOk

`func (o *ResultAllOf) GetStringsOk() (*ResultAllOfStrings, bool)`

GetStringsOk returns a tuple with the Strings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrings

`func (o *ResultAllOf) SetStrings(v ResultAllOfStrings)`

SetStrings sets Strings field to given value.

### HasStrings

`func (o *ResultAllOf) HasStrings() bool`

HasStrings returns a boolean if a field has been set.

### GetUrls

`func (o *ResultAllOf) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *ResultAllOf) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *ResultAllOf) SetUrls(v []string)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *ResultAllOf) HasUrls() bool`

HasUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


