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

// ResourceEntry struct for ResourceEntry
type ResourceEntry struct {
	Entries *[]Resource `json:"entries,omitempty"`
	// Resource name
	Name *string `json:"name,omitempty"`
}

// NewResourceEntry instantiates a new ResourceEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceEntry() *ResourceEntry {
	this := ResourceEntry{}
	return &this
}

// NewResourceEntryWithDefaults instantiates a new ResourceEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceEntryWithDefaults() *ResourceEntry {
	this := ResourceEntry{}
	return &this
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *ResourceEntry) GetEntries() []Resource {
	if o == nil || o.Entries == nil {
		var ret []Resource
		return ret
	}
	return *o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceEntry) GetEntriesOk() (*[]Resource, bool) {
	if o == nil || o.Entries == nil {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *ResourceEntry) HasEntries() bool {
	if o != nil && o.Entries != nil {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []Resource and assigns it to the Entries field.
func (o *ResourceEntry) SetEntries(v []Resource) {
	o.Entries = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResourceEntry) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceEntry) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResourceEntry) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ResourceEntry) SetName(v string) {
	o.Name = &v
}

func (o ResourceEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Entries != nil {
		toSerialize["entries"] = o.Entries
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableResourceEntry struct {
	value *ResourceEntry
	isSet bool
}

func (v NullableResourceEntry) Get() *ResourceEntry {
	return v.value
}

func (v *NullableResourceEntry) Set(val *ResourceEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceEntry(val *ResourceEntry) *NullableResourceEntry {
	return &NullableResourceEntry{value: val, isSet: true}
}

func (v NullableResourceEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


