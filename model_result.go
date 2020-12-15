/*
 * UnpacMe
 *
 *  # Introduction Welcome to the UNPACME API! All the malware unpacking and file analysis features that you are familiar with on the [unpac.me](https://www.unpac.me/) website are available through our API. You can easily integrate our unpacker into your malware analysis pipeline and begin unpacking at scale!   # Authentication The public UNPACME API is publicly available and can be accessed without authentication.  In order to use the private UNPACME API you must sign up for an account with UNPACME. Once you have a valid user account you can view your personal API key in your user profile.   <SecurityDefinitions />  # Response Structure When interacting with the UNPACME API, if the request was correctly handled, a <b>200</b> HTTP status code will be returned. The body of the response will usually be a JSON object (except for file downloads).  ## Response Status Codes  Status Code  | Description | Notes ------------- | ------------- | - 200  | OK | The request was successful 400  | Bad Request | The request was somehow incorrect. This can be caused by missing arguments or arguments with wrong values. 401 | Unauthorized | The supplied credentials, if any, are not sufficient to access the resource 403 | Forbidden | The account does not have enough privileges to make the request. 404 | Not Found | The requested resource is not found 429  | Too Many Requests | The request frequency has exceeded one of the account quotas (minute, daily or monthly). Monthly quotas are reset on the 1st of the month at 00:00 UTC. 500 | Server Error | The server could not return the representation due to an internal server error   ## Error Response  If an error has occurred while handling the request an error status code will be returend along with a JSON error message with the following properties.   Property  | Description ------------- | ------------- Error  | The error type Description  | A more informative message  # Example Clients  The following clients can be used to interact with the UNPACME API directly and are provided as examples. These clients are community projects and are not maintained or developed by UNPACME. UNPACME makes no claim as to the safety of these clients, use at your own risk.    - [UnpacMe Client](https://github.com/larsborn/UnpacMeClient) (Python)   - [UnpacMe Library](https://github.com/R3MRUM/unpacme) (Python)  <br> 
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Result Unpack result
type Result struct {
	Hashes ResultAllOfHashes `json:"hashes"`
	Analysis ResultAllOfAnalysis `json:"analysis"`
	Deepmatch *[]DeepmatchEntity `json:"deepmatch,omitempty"`
	Detectit *[]DetectitEntity `json:"detectit,omitempty"`
	MalwareId *[]MalwareId `json:"malware_id,omitempty"`
	MalwareIdRestricted *[]MalwareId `json:"malware_id_restricted,omitempty"`
	Strings *ResultAllOfStrings `json:"strings,omitempty"`
	Urls *[]string `json:"urls,omitempty"`
}

// NewResult instantiates a new Result object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResult(hashes ResultAllOfHashes, analysis ResultAllOfAnalysis, ) *Result {
	this := Result{}
	this.Hashes = hashes
	this.Analysis = analysis
	return &this
}

// NewResultWithDefaults instantiates a new Result object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultWithDefaults() *Result {
	this := Result{}
	return &this
}

// GetHashes returns the Hashes field value
func (o *Result) GetHashes() ResultAllOfHashes {
	if o == nil  {
		var ret ResultAllOfHashes
		return ret
	}

	return o.Hashes
}

// GetHashesOk returns a tuple with the Hashes field value
// and a boolean to check if the value has been set.
func (o *Result) GetHashesOk() (*ResultAllOfHashes, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Hashes, true
}

// SetHashes sets field value
func (o *Result) SetHashes(v ResultAllOfHashes) {
	o.Hashes = v
}

// GetAnalysis returns the Analysis field value
func (o *Result) GetAnalysis() ResultAllOfAnalysis {
	if o == nil  {
		var ret ResultAllOfAnalysis
		return ret
	}

	return o.Analysis
}

// GetAnalysisOk returns a tuple with the Analysis field value
// and a boolean to check if the value has been set.
func (o *Result) GetAnalysisOk() (*ResultAllOfAnalysis, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Analysis, true
}

// SetAnalysis sets field value
func (o *Result) SetAnalysis(v ResultAllOfAnalysis) {
	o.Analysis = v
}

// GetDeepmatch returns the Deepmatch field value if set, zero value otherwise.
func (o *Result) GetDeepmatch() []DeepmatchEntity {
	if o == nil || o.Deepmatch == nil {
		var ret []DeepmatchEntity
		return ret
	}
	return *o.Deepmatch
}

// GetDeepmatchOk returns a tuple with the Deepmatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Result) GetDeepmatchOk() (*[]DeepmatchEntity, bool) {
	if o == nil || o.Deepmatch == nil {
		return nil, false
	}
	return o.Deepmatch, true
}

// HasDeepmatch returns a boolean if a field has been set.
func (o *Result) HasDeepmatch() bool {
	if o != nil && o.Deepmatch != nil {
		return true
	}

	return false
}

// SetDeepmatch gets a reference to the given []DeepmatchEntity and assigns it to the Deepmatch field.
func (o *Result) SetDeepmatch(v []DeepmatchEntity) {
	o.Deepmatch = &v
}

// GetDetectit returns the Detectit field value if set, zero value otherwise.
func (o *Result) GetDetectit() []DetectitEntity {
	if o == nil || o.Detectit == nil {
		var ret []DetectitEntity
		return ret
	}
	return *o.Detectit
}

// GetDetectitOk returns a tuple with the Detectit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Result) GetDetectitOk() (*[]DetectitEntity, bool) {
	if o == nil || o.Detectit == nil {
		return nil, false
	}
	return o.Detectit, true
}

// HasDetectit returns a boolean if a field has been set.
func (o *Result) HasDetectit() bool {
	if o != nil && o.Detectit != nil {
		return true
	}

	return false
}

// SetDetectit gets a reference to the given []DetectitEntity and assigns it to the Detectit field.
func (o *Result) SetDetectit(v []DetectitEntity) {
	o.Detectit = &v
}

// GetMalwareId returns the MalwareId field value if set, zero value otherwise.
func (o *Result) GetMalwareId() []MalwareId {
	if o == nil || o.MalwareId == nil {
		var ret []MalwareId
		return ret
	}
	return *o.MalwareId
}

// GetMalwareIdOk returns a tuple with the MalwareId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Result) GetMalwareIdOk() (*[]MalwareId, bool) {
	if o == nil || o.MalwareId == nil {
		return nil, false
	}
	return o.MalwareId, true
}

// HasMalwareId returns a boolean if a field has been set.
func (o *Result) HasMalwareId() bool {
	if o != nil && o.MalwareId != nil {
		return true
	}

	return false
}

// SetMalwareId gets a reference to the given []MalwareId and assigns it to the MalwareId field.
func (o *Result) SetMalwareId(v []MalwareId) {
	o.MalwareId = &v
}

// GetMalwareIdRestricted returns the MalwareIdRestricted field value if set, zero value otherwise.
func (o *Result) GetMalwareIdRestricted() []MalwareId {
	if o == nil || o.MalwareIdRestricted == nil {
		var ret []MalwareId
		return ret
	}
	return *o.MalwareIdRestricted
}

// GetMalwareIdRestrictedOk returns a tuple with the MalwareIdRestricted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Result) GetMalwareIdRestrictedOk() (*[]MalwareId, bool) {
	if o == nil || o.MalwareIdRestricted == nil {
		return nil, false
	}
	return o.MalwareIdRestricted, true
}

// HasMalwareIdRestricted returns a boolean if a field has been set.
func (o *Result) HasMalwareIdRestricted() bool {
	if o != nil && o.MalwareIdRestricted != nil {
		return true
	}

	return false
}

// SetMalwareIdRestricted gets a reference to the given []MalwareId and assigns it to the MalwareIdRestricted field.
func (o *Result) SetMalwareIdRestricted(v []MalwareId) {
	o.MalwareIdRestricted = &v
}

// GetStrings returns the Strings field value if set, zero value otherwise.
func (o *Result) GetStrings() ResultAllOfStrings {
	if o == nil || o.Strings == nil {
		var ret ResultAllOfStrings
		return ret
	}
	return *o.Strings
}

// GetStringsOk returns a tuple with the Strings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Result) GetStringsOk() (*ResultAllOfStrings, bool) {
	if o == nil || o.Strings == nil {
		return nil, false
	}
	return o.Strings, true
}

// HasStrings returns a boolean if a field has been set.
func (o *Result) HasStrings() bool {
	if o != nil && o.Strings != nil {
		return true
	}

	return false
}

// SetStrings gets a reference to the given ResultAllOfStrings and assigns it to the Strings field.
func (o *Result) SetStrings(v ResultAllOfStrings) {
	o.Strings = &v
}

// GetUrls returns the Urls field value if set, zero value otherwise.
func (o *Result) GetUrls() []string {
	if o == nil || o.Urls == nil {
		var ret []string
		return ret
	}
	return *o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Result) GetUrlsOk() (*[]string, bool) {
	if o == nil || o.Urls == nil {
		return nil, false
	}
	return o.Urls, true
}

// HasUrls returns a boolean if a field has been set.
func (o *Result) HasUrls() bool {
	if o != nil && o.Urls != nil {
		return true
	}

	return false
}

// SetUrls gets a reference to the given []string and assigns it to the Urls field.
func (o *Result) SetUrls(v []string) {
	o.Urls = &v
}

func (o Result) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hashes"] = o.Hashes
	}
	if true {
		toSerialize["analysis"] = o.Analysis
	}
	if o.Deepmatch != nil {
		toSerialize["deepmatch"] = o.Deepmatch
	}
	if o.Detectit != nil {
		toSerialize["detectit"] = o.Detectit
	}
	if o.MalwareId != nil {
		toSerialize["malware_id"] = o.MalwareId
	}
	if o.MalwareIdRestricted != nil {
		toSerialize["malware_id_restricted"] = o.MalwareIdRestricted
	}
	if o.Strings != nil {
		toSerialize["strings"] = o.Strings
	}
	if o.Urls != nil {
		toSerialize["urls"] = o.Urls
	}
	return json.Marshal(toSerialize)
}

type NullableResult struct {
	value *Result
	isSet bool
}

func (v NullableResult) Get() *Result {
	return v.value
}

func (v *NullableResult) Set(val *Result) {
	v.value = val
	v.isSet = true
}

func (v NullableResult) IsSet() bool {
	return v.isSet
}

func (v *NullableResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResult(val *Result) *NullableResult {
	return &NullableResult{value: val, isSet: true}
}

func (v NullableResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

