# UnpackStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unpack submission ID | 
**Status** | [**Status**](status.md) |  | 

## Methods

### NewUnpackStatus

`func NewUnpackStatus(id string, status Status, ) *UnpackStatus`

NewUnpackStatus instantiates a new UnpackStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnpackStatusWithDefaults

`func NewUnpackStatusWithDefaults() *UnpackStatus`

NewUnpackStatusWithDefaults instantiates a new UnpackStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UnpackStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnpackStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnpackStatus) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *UnpackStatus) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnpackStatus) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnpackStatus) SetStatus(v Status)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


