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

// InlineResponse2001 struct for InlineResponse2001
type InlineResponse2001 struct {
	// Is user authorized to use Malpedia API
	MalpediaAuthorized *bool `json:"malpedia_authorized,omitempty"`
	// Malpedia API token
	MalpediaToken *string `json:"malpedia_token,omitempty"`
}

// NewInlineResponse2001 instantiates a new InlineResponse2001 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2001() *InlineResponse2001 {
	this := InlineResponse2001{}
	return &this
}

// NewInlineResponse2001WithDefaults instantiates a new InlineResponse2001 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2001WithDefaults() *InlineResponse2001 {
	this := InlineResponse2001{}
	return &this
}

// GetMalpediaAuthorized returns the MalpediaAuthorized field value if set, zero value otherwise.
func (o *InlineResponse2001) GetMalpediaAuthorized() bool {
	if o == nil || o.MalpediaAuthorized == nil {
		var ret bool
		return ret
	}
	return *o.MalpediaAuthorized
}

// GetMalpediaAuthorizedOk returns a tuple with the MalpediaAuthorized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetMalpediaAuthorizedOk() (*bool, bool) {
	if o == nil || o.MalpediaAuthorized == nil {
		return nil, false
	}
	return o.MalpediaAuthorized, true
}

// HasMalpediaAuthorized returns a boolean if a field has been set.
func (o *InlineResponse2001) HasMalpediaAuthorized() bool {
	if o != nil && o.MalpediaAuthorized != nil {
		return true
	}

	return false
}

// SetMalpediaAuthorized gets a reference to the given bool and assigns it to the MalpediaAuthorized field.
func (o *InlineResponse2001) SetMalpediaAuthorized(v bool) {
	o.MalpediaAuthorized = &v
}

// GetMalpediaToken returns the MalpediaToken field value if set, zero value otherwise.
func (o *InlineResponse2001) GetMalpediaToken() string {
	if o == nil || o.MalpediaToken == nil {
		var ret string
		return ret
	}
	return *o.MalpediaToken
}

// GetMalpediaTokenOk returns a tuple with the MalpediaToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetMalpediaTokenOk() (*string, bool) {
	if o == nil || o.MalpediaToken == nil {
		return nil, false
	}
	return o.MalpediaToken, true
}

// HasMalpediaToken returns a boolean if a field has been set.
func (o *InlineResponse2001) HasMalpediaToken() bool {
	if o != nil && o.MalpediaToken != nil {
		return true
	}

	return false
}

// SetMalpediaToken gets a reference to the given string and assigns it to the MalpediaToken field.
func (o *InlineResponse2001) SetMalpediaToken(v string) {
	o.MalpediaToken = &v
}

func (o InlineResponse2001) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MalpediaAuthorized != nil {
		toSerialize["malpedia_authorized"] = o.MalpediaAuthorized
	}
	if o.MalpediaToken != nil {
		toSerialize["malpedia_token"] = o.MalpediaToken
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2001 struct {
	value *InlineResponse2001
	isSet bool
}

func (v NullableInlineResponse2001) Get() *InlineResponse2001 {
	return v.value
}

func (v *NullableInlineResponse2001) Set(val *InlineResponse2001) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2001) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2001) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2001(val *InlineResponse2001) *NullableInlineResponse2001 {
	return &NullableInlineResponse2001{value: val, isSet: true}
}

func (v NullableInlineResponse2001) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2001) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

