# UserAccessAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** | API key | 
**FeedDownloadUrl** | Pointer to **string** | Sample feed direct download URL (only available to Feed subscribers) | [optional] 
**FeedPassword** | Pointer to **string** | Sample feed direct download password (only available to Feed subscribers) | [optional] 
**FeedUsername** | Pointer to **string** | Sample feed direct download username (only available to Feed subscribers) | [optional] 
**MonthLimit** | **int32** | Monthly upload limit | 
**MonthSubmissions** | **int32** | Number of submissions for the current month | 
**MonthSubmissionsPrivate** | Pointer to **int32** | Number of private submissions for the current month | [optional] 
**Roles** | **[]string** | Service access roles | 
**TotalSubmissions** | **int32** | Lifetime total number of submissions | 

## Methods

### NewUserAccessAllOf

`func NewUserAccessAllOf(apiKey string, monthLimit int32, monthSubmissions int32, roles []string, totalSubmissions int32, ) *UserAccessAllOf`

NewUserAccessAllOf instantiates a new UserAccessAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAccessAllOfWithDefaults

`func NewUserAccessAllOfWithDefaults() *UserAccessAllOf`

NewUserAccessAllOfWithDefaults instantiates a new UserAccessAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *UserAccessAllOf) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *UserAccessAllOf) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *UserAccessAllOf) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetFeedDownloadUrl

`func (o *UserAccessAllOf) GetFeedDownloadUrl() string`

GetFeedDownloadUrl returns the FeedDownloadUrl field if non-nil, zero value otherwise.

### GetFeedDownloadUrlOk

`func (o *UserAccessAllOf) GetFeedDownloadUrlOk() (*string, bool)`

GetFeedDownloadUrlOk returns a tuple with the FeedDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedDownloadUrl

`func (o *UserAccessAllOf) SetFeedDownloadUrl(v string)`

SetFeedDownloadUrl sets FeedDownloadUrl field to given value.

### HasFeedDownloadUrl

`func (o *UserAccessAllOf) HasFeedDownloadUrl() bool`

HasFeedDownloadUrl returns a boolean if a field has been set.

### GetFeedPassword

`func (o *UserAccessAllOf) GetFeedPassword() string`

GetFeedPassword returns the FeedPassword field if non-nil, zero value otherwise.

### GetFeedPasswordOk

`func (o *UserAccessAllOf) GetFeedPasswordOk() (*string, bool)`

GetFeedPasswordOk returns a tuple with the FeedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedPassword

`func (o *UserAccessAllOf) SetFeedPassword(v string)`

SetFeedPassword sets FeedPassword field to given value.

### HasFeedPassword

`func (o *UserAccessAllOf) HasFeedPassword() bool`

HasFeedPassword returns a boolean if a field has been set.

### GetFeedUsername

`func (o *UserAccessAllOf) GetFeedUsername() string`

GetFeedUsername returns the FeedUsername field if non-nil, zero value otherwise.

### GetFeedUsernameOk

`func (o *UserAccessAllOf) GetFeedUsernameOk() (*string, bool)`

GetFeedUsernameOk returns a tuple with the FeedUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedUsername

`func (o *UserAccessAllOf) SetFeedUsername(v string)`

SetFeedUsername sets FeedUsername field to given value.

### HasFeedUsername

`func (o *UserAccessAllOf) HasFeedUsername() bool`

HasFeedUsername returns a boolean if a field has been set.

### GetMonthLimit

`func (o *UserAccessAllOf) GetMonthLimit() int32`

GetMonthLimit returns the MonthLimit field if non-nil, zero value otherwise.

### GetMonthLimitOk

`func (o *UserAccessAllOf) GetMonthLimitOk() (*int32, bool)`

GetMonthLimitOk returns a tuple with the MonthLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthLimit

`func (o *UserAccessAllOf) SetMonthLimit(v int32)`

SetMonthLimit sets MonthLimit field to given value.


### GetMonthSubmissions

`func (o *UserAccessAllOf) GetMonthSubmissions() int32`

GetMonthSubmissions returns the MonthSubmissions field if non-nil, zero value otherwise.

### GetMonthSubmissionsOk

`func (o *UserAccessAllOf) GetMonthSubmissionsOk() (*int32, bool)`

GetMonthSubmissionsOk returns a tuple with the MonthSubmissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthSubmissions

`func (o *UserAccessAllOf) SetMonthSubmissions(v int32)`

SetMonthSubmissions sets MonthSubmissions field to given value.


### GetMonthSubmissionsPrivate

`func (o *UserAccessAllOf) GetMonthSubmissionsPrivate() int32`

GetMonthSubmissionsPrivate returns the MonthSubmissionsPrivate field if non-nil, zero value otherwise.

### GetMonthSubmissionsPrivateOk

`func (o *UserAccessAllOf) GetMonthSubmissionsPrivateOk() (*int32, bool)`

GetMonthSubmissionsPrivateOk returns a tuple with the MonthSubmissionsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthSubmissionsPrivate

`func (o *UserAccessAllOf) SetMonthSubmissionsPrivate(v int32)`

SetMonthSubmissionsPrivate sets MonthSubmissionsPrivate field to given value.

### HasMonthSubmissionsPrivate

`func (o *UserAccessAllOf) HasMonthSubmissionsPrivate() bool`

HasMonthSubmissionsPrivate returns a boolean if a field has been set.

### GetRoles

`func (o *UserAccessAllOf) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserAccessAllOf) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserAccessAllOf) SetRoles(v []string)`

SetRoles sets Roles field to given value.


### GetTotalSubmissions

`func (o *UserAccessAllOf) GetTotalSubmissions() int32`

GetTotalSubmissions returns the TotalSubmissions field if non-nil, zero value otherwise.

### GetTotalSubmissionsOk

`func (o *UserAccessAllOf) GetTotalSubmissionsOk() (*int32, bool)`

GetTotalSubmissionsOk returns a tuple with the TotalSubmissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSubmissions

`func (o *UserAccessAllOf) SetTotalSubmissions(v int32)`

SetTotalSubmissions sets TotalSubmissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


