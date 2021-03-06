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

// ImportEntityAllOf struct for ImportEntityAllOf
type ImportEntityAllOf struct {
	// Library name
	Library string `json:"library"`
	Functions []Function `json:"functions"`
}

// NewImportEntityAllOf instantiates a new ImportEntityAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportEntityAllOf(library string, functions []Function, ) *ImportEntityAllOf {
	this := ImportEntityAllOf{}
	this.Library = library
	this.Functions = functions
	return &this
}

// NewImportEntityAllOfWithDefaults instantiates a new ImportEntityAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportEntityAllOfWithDefaults() *ImportEntityAllOf {
	this := ImportEntityAllOf{}
	return &this
}

// GetLibrary returns the Library field value
func (o *ImportEntityAllOf) GetLibrary() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Library
}

// GetLibraryOk returns a tuple with the Library field value
// and a boolean to check if the value has been set.
func (o *ImportEntityAllOf) GetLibraryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Library, true
}

// SetLibrary sets field value
func (o *ImportEntityAllOf) SetLibrary(v string) {
	o.Library = v
}

// GetFunctions returns the Functions field value
func (o *ImportEntityAllOf) GetFunctions() []Function {
	if o == nil  {
		var ret []Function
		return ret
	}

	return o.Functions
}

// GetFunctionsOk returns a tuple with the Functions field value
// and a boolean to check if the value has been set.
func (o *ImportEntityAllOf) GetFunctionsOk() (*[]Function, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Functions, true
}

// SetFunctions sets field value
func (o *ImportEntityAllOf) SetFunctions(v []Function) {
	o.Functions = v
}

func (o ImportEntityAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["library"] = o.Library
	}
	if true {
		toSerialize["functions"] = o.Functions
	}
	return json.Marshal(toSerialize)
}

type NullableImportEntityAllOf struct {
	value *ImportEntityAllOf
	isSet bool
}

func (v NullableImportEntityAllOf) Get() *ImportEntityAllOf {
	return v.value
}

func (v *NullableImportEntityAllOf) Set(val *ImportEntityAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableImportEntityAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableImportEntityAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportEntityAllOf(val *ImportEntityAllOf) *NullableImportEntityAllOf {
	return &NullableImportEntityAllOf{value: val, isSet: true}
}

func (v NullableImportEntityAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportEntityAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


