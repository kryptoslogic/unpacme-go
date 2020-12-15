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

// DetectitEntity DetectIt Entity
type DetectitEntity struct {
	// Detection name
	Name string `json:"name"`
	// Detection options
	Options string `json:"options"`
	// Detection description
	String string `json:"string"`
	// Detection type
	Type string `json:"type"`
	// DetectIt version
	Version string `json:"version"`
}

// NewDetectitEntity instantiates a new DetectitEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetectitEntity(name string, options string, string_ string, type_ string, version string, ) *DetectitEntity {
	this := DetectitEntity{}
	this.Name = name
	this.Options = options
	this.String = string_
	this.Type = type_
	this.Version = version
	return &this
}

// NewDetectitEntityWithDefaults instantiates a new DetectitEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetectitEntityWithDefaults() *DetectitEntity {
	this := DetectitEntity{}
	return &this
}

// GetName returns the Name field value
func (o *DetectitEntity) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DetectitEntity) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DetectitEntity) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value
func (o *DetectitEntity) GetOptions() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *DetectitEntity) GetOptionsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *DetectitEntity) SetOptions(v string) {
	o.Options = v
}

// GetString returns the String field value
func (o *DetectitEntity) GetString() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.String
}

// GetStringOk returns a tuple with the String field value
// and a boolean to check if the value has been set.
func (o *DetectitEntity) GetStringOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.String, true
}

// SetString sets field value
func (o *DetectitEntity) SetString(v string) {
	o.String = v
}

// GetType returns the Type field value
func (o *DetectitEntity) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DetectitEntity) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DetectitEntity) SetType(v string) {
	o.Type = v
}

// GetVersion returns the Version field value
func (o *DetectitEntity) GetVersion() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *DetectitEntity) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *DetectitEntity) SetVersion(v string) {
	o.Version = v
}

func (o DetectitEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["options"] = o.Options
	}
	if true {
		toSerialize["string"] = o.String
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableDetectitEntity struct {
	value *DetectitEntity
	isSet bool
}

func (v NullableDetectitEntity) Get() *DetectitEntity {
	return v.value
}

func (v *NullableDetectitEntity) Set(val *DetectitEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableDetectitEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableDetectitEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetectitEntity(val *DetectitEntity) *NullableDetectitEntity {
	return &NullableDetectitEntity{value: val, isSet: true}
}

func (v NullableDetectitEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetectitEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


