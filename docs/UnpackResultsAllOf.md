# UnpackResultsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unpack submission ID | 
**Status** | [**Status**](status.md) |  | 
**Sha256** | **string** | SHA256 hash of the submitted file | 
**Time** | **float32** | Timestamp of original submission (in UTC) | 
**Private** | Pointer to **bool** | Sample is private | [optional] 
**UserSubmitted** | Pointer to **bool** | Sample was submitted by current user | [optional] 
**Results** | [**[]Result**](Result.md) | Array of results from submission, including the parent | 

## Methods

### NewUnpackResultsAllOf

`func NewUnpackResultsAllOf(id string, status Status, sha256 string, time float32, results []Result, ) *UnpackResultsAllOf`

NewUnpackResultsAllOf instantiates a new UnpackResultsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnpackResultsAllOfWithDefaults

`func NewUnpackResultsAllOfWithDefaults() *UnpackResultsAllOf`

NewUnpackResultsAllOfWithDefaults instantiates a new UnpackResultsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UnpackResultsAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnpackResultsAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnpackResultsAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *UnpackResultsAllOf) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnpackResultsAllOf) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnpackResultsAllOf) SetStatus(v Status)`

SetStatus sets Status field to given value.


### GetSha256

`func (o *UnpackResultsAllOf) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *UnpackResultsAllOf) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *UnpackResultsAllOf) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.


### GetTime

`func (o *UnpackResultsAllOf) GetTime() float32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *UnpackResultsAllOf) GetTimeOk() (*float32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *UnpackResultsAllOf) SetTime(v float32)`

SetTime sets Time field to given value.


### GetPrivate

`func (o *UnpackResultsAllOf) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *UnpackResultsAllOf) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *UnpackResultsAllOf) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *UnpackResultsAllOf) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetUserSubmitted

`func (o *UnpackResultsAllOf) GetUserSubmitted() bool`

GetUserSubmitted returns the UserSubmitted field if non-nil, zero value otherwise.

### GetUserSubmittedOk

`func (o *UnpackResultsAllOf) GetUserSubmittedOk() (*bool, bool)`

GetUserSubmittedOk returns a tuple with the UserSubmitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSubmitted

`func (o *UnpackResultsAllOf) SetUserSubmitted(v bool)`

SetUserSubmitted sets UserSubmitted field to given value.

### HasUserSubmitted

`func (o *UnpackResultsAllOf) HasUserSubmitted() bool`

HasUserSubmitted returns a boolean if a field has been set.

### GetResults

`func (o *UnpackResultsAllOf) GetResults() []Result`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *UnpackResultsAllOf) GetResultsOk() (*[]Result, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *UnpackResultsAllOf) SetResults(v []Result)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


