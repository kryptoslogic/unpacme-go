# Resource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Codepage** | **int32** | Codepage | 
**Lang** | **string** | Language | 
**MagicType** | **string** | Resource type | 
**Offset** | **int32** | Resource offset | 
**Size** | **int32** | Resource size | 
**Sublang** | **string** | Sub-Language | 

## Methods

### NewResource

`func NewResource(codepage int32, lang string, magicType string, offset int32, size int32, sublang string, ) *Resource`

NewResource instantiates a new Resource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceWithDefaults

`func NewResourceWithDefaults() *Resource`

NewResourceWithDefaults instantiates a new Resource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodepage

`func (o *Resource) GetCodepage() int32`

GetCodepage returns the Codepage field if non-nil, zero value otherwise.

### GetCodepageOk

`func (o *Resource) GetCodepageOk() (*int32, bool)`

GetCodepageOk returns a tuple with the Codepage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodepage

`func (o *Resource) SetCodepage(v int32)`

SetCodepage sets Codepage field to given value.


### GetLang

`func (o *Resource) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *Resource) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *Resource) SetLang(v string)`

SetLang sets Lang field to given value.


### GetMagicType

`func (o *Resource) GetMagicType() string`

GetMagicType returns the MagicType field if non-nil, zero value otherwise.

### GetMagicTypeOk

`func (o *Resource) GetMagicTypeOk() (*string, bool)`

GetMagicTypeOk returns a tuple with the MagicType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagicType

`func (o *Resource) SetMagicType(v string)`

SetMagicType sets MagicType field to given value.


### GetOffset

`func (o *Resource) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Resource) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *Resource) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetSize

`func (o *Resource) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Resource) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Resource) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSublang

`func (o *Resource) GetSublang() string`

GetSublang returns the Sublang field if non-nil, zero value otherwise.

### GetSublangOk

`func (o *Resource) GetSublangOk() (*string, bool)`

GetSublangOk returns a tuple with the Sublang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSublang

`func (o *Resource) SetSublang(v string)`

SetSublang sets Sublang field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


