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

// UserAccessAllOf struct for UserAccessAllOf
type UserAccessAllOf struct {
	// API key
	ApiKey string `json:"api_key"`
	// Sample feed direct download URL (only available to Feed subscribers)
	FeedDownloadUrl *string `json:"feed_download_url,omitempty"`
	// Sample feed direct download password (only available to Feed subscribers)
	FeedPassword *string `json:"feed_password,omitempty"`
	// Sample feed direct download username (only available to Feed subscribers)
	FeedUsername *string `json:"feed_username,omitempty"`
	// Monthly upload limit
	MonthLimit int32 `json:"month_limit"`
	// Number of submissions for the current month
	MonthSubmissions int32 `json:"month_submissions"`
	// Number of private submissions for the current month
	MonthSubmissionsPrivate *int32 `json:"month_submissions_private,omitempty"`
	// Service access roles
	Roles []string `json:"roles"`
	// Lifetime total number of submissions
	TotalSubmissions int32 `json:"total_submissions"`
}

// NewUserAccessAllOf instantiates a new UserAccessAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAccessAllOf(apiKey string, monthLimit int32, monthSubmissions int32, roles []string, totalSubmissions int32, ) *UserAccessAllOf {
	this := UserAccessAllOf{}
	this.ApiKey = apiKey
	this.MonthLimit = monthLimit
	this.MonthSubmissions = monthSubmissions
	this.Roles = roles
	this.TotalSubmissions = totalSubmissions
	return &this
}

// NewUserAccessAllOfWithDefaults instantiates a new UserAccessAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAccessAllOfWithDefaults() *UserAccessAllOf {
	this := UserAccessAllOf{}
	return &this
}

// GetApiKey returns the ApiKey field value
func (o *UserAccessAllOf) GetApiKey() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *UserAccessAllOf) GetApiKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *UserAccessAllOf) SetApiKey(v string) {
	o.ApiKey = v
}

// GetFeedDownloadUrl returns the FeedDownloadUrl field value if set, zero value otherwise.
func (o *UserAccessAllOf) GetFeedDownloadUrl() string {
	if o == nil || o.FeedDownloadUrl == nil {
		var ret string
		return ret
	}
	return *o.FeedDownloadUrl
}

// GetFeedDownloadUrlOk returns a tuple with the FeedDownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessAllOf) GetFeedDownloadUrlOk() (*string, bool) {
	if o == nil || o.FeedDownloadUrl == nil {
		return nil, false
	}
	return o.FeedDownloadUrl, true
}

// HasFeedDownloadUrl returns a boolean if a field has been set.
func (o *UserAccessAllOf) HasFeedDownloadUrl() bool {
	if o != nil && o.FeedDownloadUrl != nil {
		return true
	}

	return false
}

// SetFeedDownloadUrl gets a reference to the given string and assigns it to the FeedDownloadUrl field.
func (o *UserAccessAllOf) SetFeedDownloadUrl(v string) {
	o.FeedDownloadUrl = &v
}

// GetFeedPassword returns the FeedPassword field value if set, zero value otherwise.
func (o *UserAccessAllOf) GetFeedPassword() string {
	if o == nil || o.FeedPassword == nil {
		var ret string
		return ret
	}
	return *o.FeedPassword
}

// GetFeedPasswordOk returns a tuple with the FeedPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessAllOf) GetFeedPasswordOk() (*string, bool) {
	if o == nil || o.FeedPassword == nil {
		return nil, false
	}
	return o.FeedPassword, true
}

// HasFeedPassword returns a boolean if a field has been set.
func (o *UserAccessAllOf) HasFeedPassword() bool {
	if o != nil && o.FeedPassword != nil {
		return true
	}

	return false
}

// SetFeedPassword gets a reference to the given string and assigns it to the FeedPassword field.
func (o *UserAccessAllOf) SetFeedPassword(v string) {
	o.FeedPassword = &v
}

// GetFeedUsername returns the FeedUsername field value if set, zero value otherwise.
func (o *UserAccessAllOf) GetFeedUsername() string {
	if o == nil || o.FeedUsername == nil {
		var ret string
		return ret
	}
	return *o.FeedUsername
}

// GetFeedUsernameOk returns a tuple with the FeedUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessAllOf) GetFeedUsernameOk() (*string, bool) {
	if o == nil || o.FeedUsername == nil {
		return nil, false
	}
	return o.FeedUsername, true
}

// HasFeedUsername returns a boolean if a field has been set.
func (o *UserAccessAllOf) HasFeedUsername() bool {
	if o != nil && o.FeedUsername != nil {
		return true
	}

	return false
}

// SetFeedUsername gets a reference to the given string and assigns it to the FeedUsername field.
func (o *UserAccessAllOf) SetFeedUsername(v string) {
	o.FeedUsername = &v
}

// GetMonthLimit returns the MonthLimit field value
func (o *UserAccessAllOf) GetMonthLimit() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.MonthLimit
}

// GetMonthLimitOk returns a tuple with the MonthLimit field value
// and a boolean to check if the value has been set.
func (o *UserAccessAllOf) GetMonthLimitOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MonthLimit, true
}

// SetMonthLimit sets field value
func (o *UserAccessAllOf) SetMonthLimit(v int32) {
	o.MonthLimit = v
}

// GetMonthSubmissions returns the MonthSubmissions field value
func (o *UserAccessAllOf) GetMonthSubmissions() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.MonthSubmissions
}

// GetMonthSubmissionsOk returns a tuple with the MonthSubmissions field value
// and a boolean to check if the value has been set.
func (o *UserAccessAllOf) GetMonthSubmissionsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MonthSubmissions, true
}

// SetMonthSubmissions sets field value
func (o *UserAccessAllOf) SetMonthSubmissions(v int32) {
	o.MonthSubmissions = v
}

// GetMonthSubmissionsPrivate returns the MonthSubmissionsPrivate field value if set, zero value otherwise.
func (o *UserAccessAllOf) GetMonthSubmissionsPrivate() int32 {
	if o == nil || o.MonthSubmissionsPrivate == nil {
		var ret int32
		return ret
	}
	return *o.MonthSubmissionsPrivate
}

// GetMonthSubmissionsPrivateOk returns a tuple with the MonthSubmissionsPrivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessAllOf) GetMonthSubmissionsPrivateOk() (*int32, bool) {
	if o == nil || o.MonthSubmissionsPrivate == nil {
		return nil, false
	}
	return o.MonthSubmissionsPrivate, true
}

// HasMonthSubmissionsPrivate returns a boolean if a field has been set.
func (o *UserAccessAllOf) HasMonthSubmissionsPrivate() bool {
	if o != nil && o.MonthSubmissionsPrivate != nil {
		return true
	}

	return false
}

// SetMonthSubmissionsPrivate gets a reference to the given int32 and assigns it to the MonthSubmissionsPrivate field.
func (o *UserAccessAllOf) SetMonthSubmissionsPrivate(v int32) {
	o.MonthSubmissionsPrivate = &v
}

// GetRoles returns the Roles field value
func (o *UserAccessAllOf) GetRoles() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *UserAccessAllOf) GetRolesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Roles, true
}

// SetRoles sets field value
func (o *UserAccessAllOf) SetRoles(v []string) {
	o.Roles = v
}

// GetTotalSubmissions returns the TotalSubmissions field value
func (o *UserAccessAllOf) GetTotalSubmissions() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.TotalSubmissions
}

// GetTotalSubmissionsOk returns a tuple with the TotalSubmissions field value
// and a boolean to check if the value has been set.
func (o *UserAccessAllOf) GetTotalSubmissionsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalSubmissions, true
}

// SetTotalSubmissions sets field value
func (o *UserAccessAllOf) SetTotalSubmissions(v int32) {
	o.TotalSubmissions = v
}

func (o UserAccessAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["api_key"] = o.ApiKey
	}
	if o.FeedDownloadUrl != nil {
		toSerialize["feed_download_url"] = o.FeedDownloadUrl
	}
	if o.FeedPassword != nil {
		toSerialize["feed_password"] = o.FeedPassword
	}
	if o.FeedUsername != nil {
		toSerialize["feed_username"] = o.FeedUsername
	}
	if true {
		toSerialize["month_limit"] = o.MonthLimit
	}
	if true {
		toSerialize["month_submissions"] = o.MonthSubmissions
	}
	if o.MonthSubmissionsPrivate != nil {
		toSerialize["month_submissions_private"] = o.MonthSubmissionsPrivate
	}
	if true {
		toSerialize["roles"] = o.Roles
	}
	if true {
		toSerialize["total_submissions"] = o.TotalSubmissions
	}
	return json.Marshal(toSerialize)
}

type NullableUserAccessAllOf struct {
	value *UserAccessAllOf
	isSet bool
}

func (v NullableUserAccessAllOf) Get() *UserAccessAllOf {
	return v.value
}

func (v *NullableUserAccessAllOf) Set(val *UserAccessAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAccessAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAccessAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAccessAllOf(val *UserAccessAllOf) *NullableUserAccessAllOf {
	return &NullableUserAccessAllOf{value: val, isSet: true}
}

func (v NullableUserAccessAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAccessAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


