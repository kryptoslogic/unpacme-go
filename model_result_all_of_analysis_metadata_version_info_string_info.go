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

// ResultAllOfAnalysisMetadataVersionInfoStringInfo struct for ResultAllOfAnalysisMetadataVersionInfoStringInfo
type ResultAllOfAnalysisMetadataVersionInfoStringInfo struct {
	// Company name
	CompanyName *string `json:"CompanyName,omitempty"`
	// File description
	FileDescription *string `json:"FileDescription,omitempty"`
	// File version
	FileVersion *string `json:"FileVersion,omitempty"`
	// Legal copyright
	LegalCopyright *string `json:"LegalCopyright,omitempty"`
	// Original file name
	OriginalFilename *string `json:"OriginalFilename,omitempty"`
	// Product name
	ProductName *string `json:"ProductName,omitempty"`
	// Product version
	ProductVersion *string `json:"ProductVersion,omitempty"`
}

// NewResultAllOfAnalysisMetadataVersionInfoStringInfo instantiates a new ResultAllOfAnalysisMetadataVersionInfoStringInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultAllOfAnalysisMetadataVersionInfoStringInfo() *ResultAllOfAnalysisMetadataVersionInfoStringInfo {
	this := ResultAllOfAnalysisMetadataVersionInfoStringInfo{}
	return &this
}

// NewResultAllOfAnalysisMetadataVersionInfoStringInfoWithDefaults instantiates a new ResultAllOfAnalysisMetadataVersionInfoStringInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultAllOfAnalysisMetadataVersionInfoStringInfoWithDefaults() *ResultAllOfAnalysisMetadataVersionInfoStringInfo {
	this := ResultAllOfAnalysisMetadataVersionInfoStringInfo{}
	return &this
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *ResultAllOfAnalysisMetadataVersionInfoStringInfo) GetCompanyName() string {
	if o == nil || o.CompanyName == nil {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultAllOfAnalysisMetadataVersionInfoStringInfo) GetCompanyNameOk() (*string, bool) {
	if o == nil || o.CompanyName == nil {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *ResultAllOfAnalysisMetadataVersionInfoStringInfo) HasCompanyName() bool {
	if o != nil && o.CompanyName != nil {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *ResultAllOfAnalysisMetadataVersionInfoStringInfo) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetFileDescription returns the FileDescription field value if set, zero value otherwise.
func (o *ResultAllOfAnalysisMetadataVersionInfoStringInfo) GetFileDescription() string {
	if o == nil || o.FileDescription == nil {
		var ret string
		return ret
	}
	return *o.FileDescription
}

// GetFileDescriptionOk returns a tuple with the FileDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultAllOfAnalysisMetadataVersionInfoStringInfo) GetFileDescriptionOk() (*string, bool) {
	if o == nil || o.FileDescription == nil {
		return nil, false
	}
	return o.FileDescription, true
}

// HasFileDescription returns a boolean if a field has been set.
func (o *ResultAllOfAnalysisMetadataVersionInfoStringInfo) HasFileDescription() bool {
	if o != nil && o.FileDescription != nil {
		return true
	}

	return false
}

// SetFileDescription gets a reference to the given string and assigns it to the FileDescription field.
func (o *ResultAllOfAnalysisMetadataVersionInfoStringInfo) SetFileDescription(v string) {
	o.FileDescription = &v
}

// GetFileVersion returns the FileVersion field value if set, zero value otherwise.
func (o *ResultAllOfAnalysisMetadataVersionInfoStringInfo) GetFileVersion() string {
	if o == nil || o.FileVersion == nil {
		var ret string
		return ret
	}
	return *o.FileVersion
}

// GetFileVersionOk returns a tuple with the FileVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultAllOfAnalysisMetadataVersionInfoStringInfo) GetFileVersionOk() (*string, bool) {
	if o == nil || o.FileVersion == nil {
		return nil, false
	}
	return o.FileVersion, true
}

// HasFileVersion returns a boolean if a field has been set.
func (o *ResultAllOfAnalysisMetadataVersionInfoStringInfo) HasFileVersion() bool {
	if o != nil && o.FileVersion != nil {
		return true
	}

	return false
}

// SetFileVersion gets a reference to the given string and assigns it to the FileVersion field.
func (o *ResultAllOfAnalysisMetadataVersionInfoStringInfo) SetFileVersion(v string) {
	o.FileVersion = &v
}

// GetLegalCopyright returns the LegalCopyright field value if set, zero value otherwise.
func (o *ResultAllOfAnalysisMetadataVersionInfoStringInfo) GetLegalCopyright() string {
	if o == nil || o.LegalCopyright == nil {
		var ret string
		return ret
	}
	return *o.LegalCopyright
}

// GetLegalCopyrightOk returns a tuple with the LegalCopyright field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultAllOfAnalysisMetadataVersionInfoStringInfo) GetLegalCopyrightOk() (*string, bool) {
	if o == nil || o.LegalCopyright == nil {
		return nil, false
	}
	return o.LegalCopyright, true
}

// HasLegalCopyright returns a boolean if a field has been set.
func (o *ResultAllOfAnalysisMetadataVersionInfoStringInfo) HasLegalCopyright() bool {
	if o != nil && o.LegalCopyright != nil {
		return true
	}

	return false
}

// SetLegalCopyright gets a reference to the given string and assigns it to the LegalCopyright field.
func (o *ResultAllOfAnalysisMetadataVersionInfoStringInfo) SetLegalCopyright(v string) {
	o.LegalCopyright = &v
}

// GetOriginalFilename returns the OriginalFilename field value if set, zero value otherwise.
func (o *ResultAllOfAnalysisMetadataVersionInfoStringInfo) GetOriginalFilename() string {
	if o == nil || o.OriginalFilename == nil {
		var ret string
		return ret
	}
	return *o.OriginalFilename
}

// GetOriginalFilenameOk returns a tuple with the OriginalFilename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultAllOfAnalysisMetadataVersionInfoStringInfo) GetOriginalFilenameOk() (*string, bool) {
	if o == nil || o.OriginalFilename == nil {
		return nil, false
	}
	return o.OriginalFilename, true
}

// HasOriginalFilename returns a boolean if a field has been set.
func (o *ResultAllOfAnalysisMetadataVersionInfoStringInfo) HasOriginalFilename() bool {
	if o != nil && o.OriginalFilename != nil {
		return true
	}

	return false
}

// SetOriginalFilename gets a reference to the given string and assigns it to the OriginalFilename field.
func (o *ResultAllOfAnalysisMetadataVersionInfoStringInfo) SetOriginalFilename(v string) {
	o.OriginalFilename = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *ResultAllOfAnalysisMetadataVersionInfoStringInfo) GetProductName() string {
	if o == nil || o.ProductName == nil {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultAllOfAnalysisMetadataVersionInfoStringInfo) GetProductNameOk() (*string, bool) {
	if o == nil || o.ProductName == nil {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *ResultAllOfAnalysisMetadataVersionInfoStringInfo) HasProductName() bool {
	if o != nil && o.ProductName != nil {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *ResultAllOfAnalysisMetadataVersionInfoStringInfo) SetProductName(v string) {
	o.ProductName = &v
}

// GetProductVersion returns the ProductVersion field value if set, zero value otherwise.
func (o *ResultAllOfAnalysisMetadataVersionInfoStringInfo) GetProductVersion() string {
	if o == nil || o.ProductVersion == nil {
		var ret string
		return ret
	}
	return *o.ProductVersion
}

// GetProductVersionOk returns a tuple with the ProductVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultAllOfAnalysisMetadataVersionInfoStringInfo) GetProductVersionOk() (*string, bool) {
	if o == nil || o.ProductVersion == nil {
		return nil, false
	}
	return o.ProductVersion, true
}

// HasProductVersion returns a boolean if a field has been set.
func (o *ResultAllOfAnalysisMetadataVersionInfoStringInfo) HasProductVersion() bool {
	if o != nil && o.ProductVersion != nil {
		return true
	}

	return false
}

// SetProductVersion gets a reference to the given string and assigns it to the ProductVersion field.
func (o *ResultAllOfAnalysisMetadataVersionInfoStringInfo) SetProductVersion(v string) {
	o.ProductVersion = &v
}

func (o ResultAllOfAnalysisMetadataVersionInfoStringInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CompanyName != nil {
		toSerialize["CompanyName"] = o.CompanyName
	}
	if o.FileDescription != nil {
		toSerialize["FileDescription"] = o.FileDescription
	}
	if o.FileVersion != nil {
		toSerialize["FileVersion"] = o.FileVersion
	}
	if o.LegalCopyright != nil {
		toSerialize["LegalCopyright"] = o.LegalCopyright
	}
	if o.OriginalFilename != nil {
		toSerialize["OriginalFilename"] = o.OriginalFilename
	}
	if o.ProductName != nil {
		toSerialize["ProductName"] = o.ProductName
	}
	if o.ProductVersion != nil {
		toSerialize["ProductVersion"] = o.ProductVersion
	}
	return json.Marshal(toSerialize)
}

type NullableResultAllOfAnalysisMetadataVersionInfoStringInfo struct {
	value *ResultAllOfAnalysisMetadataVersionInfoStringInfo
	isSet bool
}

func (v NullableResultAllOfAnalysisMetadataVersionInfoStringInfo) Get() *ResultAllOfAnalysisMetadataVersionInfoStringInfo {
	return v.value
}

func (v *NullableResultAllOfAnalysisMetadataVersionInfoStringInfo) Set(val *ResultAllOfAnalysisMetadataVersionInfoStringInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableResultAllOfAnalysisMetadataVersionInfoStringInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableResultAllOfAnalysisMetadataVersionInfoStringInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultAllOfAnalysisMetadataVersionInfoStringInfo(val *ResultAllOfAnalysisMetadataVersionInfoStringInfo) *NullableResultAllOfAnalysisMetadataVersionInfoStringInfo {
	return &NullableResultAllOfAnalysisMetadataVersionInfoStringInfo{value: val, isSet: true}
}

func (v NullableResultAllOfAnalysisMetadataVersionInfoStringInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultAllOfAnalysisMetadataVersionInfoStringInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


