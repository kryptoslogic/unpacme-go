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

// Section PE section
type Section struct {
	// Section characteristics
	Characteristics int32 `json:"Characteristics"`
	// Section entropy
	Entropy float32 `json:"Entropy"`
	// MD5 of section data
	MD5 string `json:"MD5"`
	// Section name
	Name string `json:"Name"`
	// Hex encoded section name
	NameHex string `json:"Name_Hex"`
	// Number of line numbers
	NumberOfLinenumbers int32 `json:"NumberOfLinenumbers"`
	// Number of relocations
	NumberOfRelocations int32 `json:"NumberOfRelocations"`
	// Pointer to line numbers
	PointerToLinenumbers int32 `json:"PointerToLinenumbers"`
	// Pointer to raw data
	PointerToRawData int32 `json:"PointerToRawData"`
	// Pointer to relocations
	PointerToRelocations int32 `json:"PointerToRelocations"`
	// SHA1 of section data
	SHA1 string `json:"SHA1"`
	// SHA256 of section data
	SHA256 string `json:"SHA256"`
	// Size of raw data
	SizeOfRawData int32 `json:"SizeOfRawData"`
	// Virtual address of section
	VirtualAddress int32 `json:"VirtualAddress"`
	// Virtual size of section
	VirtualSize int32 `json:"VirtualSize"`
}

// NewSection instantiates a new Section object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSection(characteristics int32, entropy float32, mD5 string, name string, nameHex string, numberOfLinenumbers int32, numberOfRelocations int32, pointerToLinenumbers int32, pointerToRawData int32, pointerToRelocations int32, sHA1 string, sHA256 string, sizeOfRawData int32, virtualAddress int32, virtualSize int32, ) *Section {
	this := Section{}
	this.Characteristics = characteristics
	this.Entropy = entropy
	this.MD5 = mD5
	this.Name = name
	this.NameHex = nameHex
	this.NumberOfLinenumbers = numberOfLinenumbers
	this.NumberOfRelocations = numberOfRelocations
	this.PointerToLinenumbers = pointerToLinenumbers
	this.PointerToRawData = pointerToRawData
	this.PointerToRelocations = pointerToRelocations
	this.SHA1 = sHA1
	this.SHA256 = sHA256
	this.SizeOfRawData = sizeOfRawData
	this.VirtualAddress = virtualAddress
	this.VirtualSize = virtualSize
	return &this
}

// NewSectionWithDefaults instantiates a new Section object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSectionWithDefaults() *Section {
	this := Section{}
	return &this
}

// GetCharacteristics returns the Characteristics field value
func (o *Section) GetCharacteristics() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Characteristics
}

// GetCharacteristicsOk returns a tuple with the Characteristics field value
// and a boolean to check if the value has been set.
func (o *Section) GetCharacteristicsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Characteristics, true
}

// SetCharacteristics sets field value
func (o *Section) SetCharacteristics(v int32) {
	o.Characteristics = v
}

// GetEntropy returns the Entropy field value
func (o *Section) GetEntropy() float32 {
	if o == nil  {
		var ret float32
		return ret
	}

	return o.Entropy
}

// GetEntropyOk returns a tuple with the Entropy field value
// and a boolean to check if the value has been set.
func (o *Section) GetEntropyOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Entropy, true
}

// SetEntropy sets field value
func (o *Section) SetEntropy(v float32) {
	o.Entropy = v
}

// GetMD5 returns the MD5 field value
func (o *Section) GetMD5() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.MD5
}

// GetMD5Ok returns a tuple with the MD5 field value
// and a boolean to check if the value has been set.
func (o *Section) GetMD5Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MD5, true
}

// SetMD5 sets field value
func (o *Section) SetMD5(v string) {
	o.MD5 = v
}

// GetName returns the Name field value
func (o *Section) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Section) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Section) SetName(v string) {
	o.Name = v
}

// GetNameHex returns the NameHex field value
func (o *Section) GetNameHex() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.NameHex
}

// GetNameHexOk returns a tuple with the NameHex field value
// and a boolean to check if the value has been set.
func (o *Section) GetNameHexOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NameHex, true
}

// SetNameHex sets field value
func (o *Section) SetNameHex(v string) {
	o.NameHex = v
}

// GetNumberOfLinenumbers returns the NumberOfLinenumbers field value
func (o *Section) GetNumberOfLinenumbers() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.NumberOfLinenumbers
}

// GetNumberOfLinenumbersOk returns a tuple with the NumberOfLinenumbers field value
// and a boolean to check if the value has been set.
func (o *Section) GetNumberOfLinenumbersOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NumberOfLinenumbers, true
}

// SetNumberOfLinenumbers sets field value
func (o *Section) SetNumberOfLinenumbers(v int32) {
	o.NumberOfLinenumbers = v
}

// GetNumberOfRelocations returns the NumberOfRelocations field value
func (o *Section) GetNumberOfRelocations() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.NumberOfRelocations
}

// GetNumberOfRelocationsOk returns a tuple with the NumberOfRelocations field value
// and a boolean to check if the value has been set.
func (o *Section) GetNumberOfRelocationsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NumberOfRelocations, true
}

// SetNumberOfRelocations sets field value
func (o *Section) SetNumberOfRelocations(v int32) {
	o.NumberOfRelocations = v
}

// GetPointerToLinenumbers returns the PointerToLinenumbers field value
func (o *Section) GetPointerToLinenumbers() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.PointerToLinenumbers
}

// GetPointerToLinenumbersOk returns a tuple with the PointerToLinenumbers field value
// and a boolean to check if the value has been set.
func (o *Section) GetPointerToLinenumbersOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PointerToLinenumbers, true
}

// SetPointerToLinenumbers sets field value
func (o *Section) SetPointerToLinenumbers(v int32) {
	o.PointerToLinenumbers = v
}

// GetPointerToRawData returns the PointerToRawData field value
func (o *Section) GetPointerToRawData() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.PointerToRawData
}

// GetPointerToRawDataOk returns a tuple with the PointerToRawData field value
// and a boolean to check if the value has been set.
func (o *Section) GetPointerToRawDataOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PointerToRawData, true
}

// SetPointerToRawData sets field value
func (o *Section) SetPointerToRawData(v int32) {
	o.PointerToRawData = v
}

// GetPointerToRelocations returns the PointerToRelocations field value
func (o *Section) GetPointerToRelocations() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.PointerToRelocations
}

// GetPointerToRelocationsOk returns a tuple with the PointerToRelocations field value
// and a boolean to check if the value has been set.
func (o *Section) GetPointerToRelocationsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PointerToRelocations, true
}

// SetPointerToRelocations sets field value
func (o *Section) SetPointerToRelocations(v int32) {
	o.PointerToRelocations = v
}

// GetSHA1 returns the SHA1 field value
func (o *Section) GetSHA1() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SHA1
}

// GetSHA1Ok returns a tuple with the SHA1 field value
// and a boolean to check if the value has been set.
func (o *Section) GetSHA1Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SHA1, true
}

// SetSHA1 sets field value
func (o *Section) SetSHA1(v string) {
	o.SHA1 = v
}

// GetSHA256 returns the SHA256 field value
func (o *Section) GetSHA256() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SHA256
}

// GetSHA256Ok returns a tuple with the SHA256 field value
// and a boolean to check if the value has been set.
func (o *Section) GetSHA256Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SHA256, true
}

// SetSHA256 sets field value
func (o *Section) SetSHA256(v string) {
	o.SHA256 = v
}

// GetSizeOfRawData returns the SizeOfRawData field value
func (o *Section) GetSizeOfRawData() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.SizeOfRawData
}

// GetSizeOfRawDataOk returns a tuple with the SizeOfRawData field value
// and a boolean to check if the value has been set.
func (o *Section) GetSizeOfRawDataOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SizeOfRawData, true
}

// SetSizeOfRawData sets field value
func (o *Section) SetSizeOfRawData(v int32) {
	o.SizeOfRawData = v
}

// GetVirtualAddress returns the VirtualAddress field value
func (o *Section) GetVirtualAddress() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.VirtualAddress
}

// GetVirtualAddressOk returns a tuple with the VirtualAddress field value
// and a boolean to check if the value has been set.
func (o *Section) GetVirtualAddressOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VirtualAddress, true
}

// SetVirtualAddress sets field value
func (o *Section) SetVirtualAddress(v int32) {
	o.VirtualAddress = v
}

// GetVirtualSize returns the VirtualSize field value
func (o *Section) GetVirtualSize() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.VirtualSize
}

// GetVirtualSizeOk returns a tuple with the VirtualSize field value
// and a boolean to check if the value has been set.
func (o *Section) GetVirtualSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VirtualSize, true
}

// SetVirtualSize sets field value
func (o *Section) SetVirtualSize(v int32) {
	o.VirtualSize = v
}

func (o Section) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Characteristics"] = o.Characteristics
	}
	if true {
		toSerialize["Entropy"] = o.Entropy
	}
	if true {
		toSerialize["MD5"] = o.MD5
	}
	if true {
		toSerialize["Name"] = o.Name
	}
	if true {
		toSerialize["Name_Hex"] = o.NameHex
	}
	if true {
		toSerialize["NumberOfLinenumbers"] = o.NumberOfLinenumbers
	}
	if true {
		toSerialize["NumberOfRelocations"] = o.NumberOfRelocations
	}
	if true {
		toSerialize["PointerToLinenumbers"] = o.PointerToLinenumbers
	}
	if true {
		toSerialize["PointerToRawData"] = o.PointerToRawData
	}
	if true {
		toSerialize["PointerToRelocations"] = o.PointerToRelocations
	}
	if true {
		toSerialize["SHA1"] = o.SHA1
	}
	if true {
		toSerialize["SHA256"] = o.SHA256
	}
	if true {
		toSerialize["SizeOfRawData"] = o.SizeOfRawData
	}
	if true {
		toSerialize["VirtualAddress"] = o.VirtualAddress
	}
	if true {
		toSerialize["VirtualSize"] = o.VirtualSize
	}
	return json.Marshal(toSerialize)
}

type NullableSection struct {
	value *Section
	isSet bool
}

func (v NullableSection) Get() *Section {
	return v.value
}

func (v *NullableSection) Set(val *Section) {
	v.value = val
	v.isSet = true
}

func (v NullableSection) IsSet() bool {
	return v.isSet
}

func (v *NullableSection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSection(val *Section) *NullableSection {
	return &NullableSection{value: val, isSet: true}
}

func (v NullableSection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


