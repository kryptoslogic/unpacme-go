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

// ExportAllOf struct for ExportAllOf
type ExportAllOf struct {
	// Export address
	Address int32 `json:"address"`
	// Export name
	Name string `json:"name"`
	// Export ordinal
	Ordinal int32 `json:"ordinal"`
}

// NewExportAllOf instantiates a new ExportAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportAllOf(address int32, name string, ordinal int32, ) *ExportAllOf {
	this := ExportAllOf{}
	this.Address = address
	this.Name = name
	this.Ordinal = ordinal
	return &this
}

// NewExportAllOfWithDefaults instantiates a new ExportAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportAllOfWithDefaults() *ExportAllOf {
	this := ExportAllOf{}
	return &this
}

// GetAddress returns the Address field value
func (o *ExportAllOf) GetAddress() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *ExportAllOf) GetAddressOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *ExportAllOf) SetAddress(v int32) {
	o.Address = v
}

// GetName returns the Name field value
func (o *ExportAllOf) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExportAllOf) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExportAllOf) SetName(v string) {
	o.Name = v
}

// GetOrdinal returns the Ordinal field value
func (o *ExportAllOf) GetOrdinal() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Ordinal
}

// GetOrdinalOk returns a tuple with the Ordinal field value
// and a boolean to check if the value has been set.
func (o *ExportAllOf) GetOrdinalOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Ordinal, true
}

// SetOrdinal sets field value
func (o *ExportAllOf) SetOrdinal(v int32) {
	o.Ordinal = v
}

func (o ExportAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["ordinal"] = o.Ordinal
	}
	return json.Marshal(toSerialize)
}

type NullableExportAllOf struct {
	value *ExportAllOf
	isSet bool
}

func (v NullableExportAllOf) Get() *ExportAllOf {
	return v.value
}

func (v *NullableExportAllOf) Set(val *ExportAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableExportAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableExportAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportAllOf(val *ExportAllOf) *NullableExportAllOf {
	return &NullableExportAllOf{value: val, isSet: true}
}

func (v NullableExportAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

