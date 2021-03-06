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

// Function API function
type Function struct {
	// API imported by ordinal
	ImportByOrdinal bool `json:"import_by_ordinal"`
	// API name
	Name string `json:"name"`
}

// NewFunction instantiates a new Function object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunction(importByOrdinal bool, name string, ) *Function {
	this := Function{}
	this.ImportByOrdinal = importByOrdinal
	this.Name = name
	return &this
}

// NewFunctionWithDefaults instantiates a new Function object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionWithDefaults() *Function {
	this := Function{}
	return &this
}

// GetImportByOrdinal returns the ImportByOrdinal field value
func (o *Function) GetImportByOrdinal() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.ImportByOrdinal
}

// GetImportByOrdinalOk returns a tuple with the ImportByOrdinal field value
// and a boolean to check if the value has been set.
func (o *Function) GetImportByOrdinalOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ImportByOrdinal, true
}

// SetImportByOrdinal sets field value
func (o *Function) SetImportByOrdinal(v bool) {
	o.ImportByOrdinal = v
}

// GetName returns the Name field value
func (o *Function) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Function) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Function) SetName(v string) {
	o.Name = v
}

func (o Function) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["import_by_ordinal"] = o.ImportByOrdinal
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableFunction struct {
	value *Function
	isSet bool
}

func (v NullableFunction) Get() *Function {
	return v.value
}

func (v *NullableFunction) Set(val *Function) {
	v.value = val
	v.isSet = true
}

func (v NullableFunction) IsSet() bool {
	return v.isSet
}

func (v *NullableFunction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunction(val *Function) *NullableFunction {
	return &NullableFunction{value: val, isSet: true}
}

func (v NullableFunction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


