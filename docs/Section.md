# Section

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Characteristics** | **int32** | Section characteristics | 
**Entropy** | **float32** | Section entropy | 
**MD5** | **string** | MD5 of section data | 
**Name** | **string** | Section name | 
**NameHex** | **string** | Hex encoded section name | 
**NumberOfLinenumbers** | **int32** | Number of line numbers | 
**NumberOfRelocations** | **int32** | Number of relocations | 
**PointerToLinenumbers** | **int32** | Pointer to line numbers | 
**PointerToRawData** | **int32** | Pointer to raw data | 
**PointerToRelocations** | **int32** | Pointer to relocations | 
**SHA1** | **string** | SHA1 of section data | 
**SHA256** | **string** | SHA256 of section data | 
**SizeOfRawData** | **int32** | Size of raw data | 
**VirtualAddress** | **int32** | Virtual address of section | 
**VirtualSize** | **int32** | Virtual size of section | 

## Methods

### NewSection

`func NewSection(characteristics int32, entropy float32, mD5 string, name string, nameHex string, numberOfLinenumbers int32, numberOfRelocations int32, pointerToLinenumbers int32, pointerToRawData int32, pointerToRelocations int32, sHA1 string, sHA256 string, sizeOfRawData int32, virtualAddress int32, virtualSize int32, ) *Section`

NewSection instantiates a new Section object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectionWithDefaults

`func NewSectionWithDefaults() *Section`

NewSectionWithDefaults instantiates a new Section object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCharacteristics

`func (o *Section) GetCharacteristics() int32`

GetCharacteristics returns the Characteristics field if non-nil, zero value otherwise.

### GetCharacteristicsOk

`func (o *Section) GetCharacteristicsOk() (*int32, bool)`

GetCharacteristicsOk returns a tuple with the Characteristics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacteristics

`func (o *Section) SetCharacteristics(v int32)`

SetCharacteristics sets Characteristics field to given value.


### GetEntropy

`func (o *Section) GetEntropy() float32`

GetEntropy returns the Entropy field if non-nil, zero value otherwise.

### GetEntropyOk

`func (o *Section) GetEntropyOk() (*float32, bool)`

GetEntropyOk returns a tuple with the Entropy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntropy

`func (o *Section) SetEntropy(v float32)`

SetEntropy sets Entropy field to given value.


### GetMD5

`func (o *Section) GetMD5() string`

GetMD5 returns the MD5 field if non-nil, zero value otherwise.

### GetMD5Ok

`func (o *Section) GetMD5Ok() (*string, bool)`

GetMD5Ok returns a tuple with the MD5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMD5

`func (o *Section) SetMD5(v string)`

SetMD5 sets MD5 field to given value.


### GetName

`func (o *Section) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Section) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Section) SetName(v string)`

SetName sets Name field to given value.


### GetNameHex

`func (o *Section) GetNameHex() string`

GetNameHex returns the NameHex field if non-nil, zero value otherwise.

### GetNameHexOk

`func (o *Section) GetNameHexOk() (*string, bool)`

GetNameHexOk returns a tuple with the NameHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameHex

`func (o *Section) SetNameHex(v string)`

SetNameHex sets NameHex field to given value.


### GetNumberOfLinenumbers

`func (o *Section) GetNumberOfLinenumbers() int32`

GetNumberOfLinenumbers returns the NumberOfLinenumbers field if non-nil, zero value otherwise.

### GetNumberOfLinenumbersOk

`func (o *Section) GetNumberOfLinenumbersOk() (*int32, bool)`

GetNumberOfLinenumbersOk returns a tuple with the NumberOfLinenumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfLinenumbers

`func (o *Section) SetNumberOfLinenumbers(v int32)`

SetNumberOfLinenumbers sets NumberOfLinenumbers field to given value.


### GetNumberOfRelocations

`func (o *Section) GetNumberOfRelocations() int32`

GetNumberOfRelocations returns the NumberOfRelocations field if non-nil, zero value otherwise.

### GetNumberOfRelocationsOk

`func (o *Section) GetNumberOfRelocationsOk() (*int32, bool)`

GetNumberOfRelocationsOk returns a tuple with the NumberOfRelocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfRelocations

`func (o *Section) SetNumberOfRelocations(v int32)`

SetNumberOfRelocations sets NumberOfRelocations field to given value.


### GetPointerToLinenumbers

`func (o *Section) GetPointerToLinenumbers() int32`

GetPointerToLinenumbers returns the PointerToLinenumbers field if non-nil, zero value otherwise.

### GetPointerToLinenumbersOk

`func (o *Section) GetPointerToLinenumbersOk() (*int32, bool)`

GetPointerToLinenumbersOk returns a tuple with the PointerToLinenumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointerToLinenumbers

`func (o *Section) SetPointerToLinenumbers(v int32)`

SetPointerToLinenumbers sets PointerToLinenumbers field to given value.


### GetPointerToRawData

`func (o *Section) GetPointerToRawData() int32`

GetPointerToRawData returns the PointerToRawData field if non-nil, zero value otherwise.

### GetPointerToRawDataOk

`func (o *Section) GetPointerToRawDataOk() (*int32, bool)`

GetPointerToRawDataOk returns a tuple with the PointerToRawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointerToRawData

`func (o *Section) SetPointerToRawData(v int32)`

SetPointerToRawData sets PointerToRawData field to given value.


### GetPointerToRelocations

`func (o *Section) GetPointerToRelocations() int32`

GetPointerToRelocations returns the PointerToRelocations field if non-nil, zero value otherwise.

### GetPointerToRelocationsOk

`func (o *Section) GetPointerToRelocationsOk() (*int32, bool)`

GetPointerToRelocationsOk returns a tuple with the PointerToRelocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointerToRelocations

`func (o *Section) SetPointerToRelocations(v int32)`

SetPointerToRelocations sets PointerToRelocations field to given value.


### GetSHA1

`func (o *Section) GetSHA1() string`

GetSHA1 returns the SHA1 field if non-nil, zero value otherwise.

### GetSHA1Ok

`func (o *Section) GetSHA1Ok() (*string, bool)`

GetSHA1Ok returns a tuple with the SHA1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSHA1

`func (o *Section) SetSHA1(v string)`

SetSHA1 sets SHA1 field to given value.


### GetSHA256

`func (o *Section) GetSHA256() string`

GetSHA256 returns the SHA256 field if non-nil, zero value otherwise.

### GetSHA256Ok

`func (o *Section) GetSHA256Ok() (*string, bool)`

GetSHA256Ok returns a tuple with the SHA256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSHA256

`func (o *Section) SetSHA256(v string)`

SetSHA256 sets SHA256 field to given value.


### GetSizeOfRawData

`func (o *Section) GetSizeOfRawData() int32`

GetSizeOfRawData returns the SizeOfRawData field if non-nil, zero value otherwise.

### GetSizeOfRawDataOk

`func (o *Section) GetSizeOfRawDataOk() (*int32, bool)`

GetSizeOfRawDataOk returns a tuple with the SizeOfRawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeOfRawData

`func (o *Section) SetSizeOfRawData(v int32)`

SetSizeOfRawData sets SizeOfRawData field to given value.


### GetVirtualAddress

`func (o *Section) GetVirtualAddress() int32`

GetVirtualAddress returns the VirtualAddress field if non-nil, zero value otherwise.

### GetVirtualAddressOk

`func (o *Section) GetVirtualAddressOk() (*int32, bool)`

GetVirtualAddressOk returns a tuple with the VirtualAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualAddress

`func (o *Section) SetVirtualAddress(v int32)`

SetVirtualAddress sets VirtualAddress field to given value.


### GetVirtualSize

`func (o *Section) GetVirtualSize() int32`

GetVirtualSize returns the VirtualSize field if non-nil, zero value otherwise.

### GetVirtualSizeOk

`func (o *Section) GetVirtualSizeOk() (*int32, bool)`

GetVirtualSizeOk returns a tuple with the VirtualSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualSize

`func (o *Section) SetVirtualSize(v int32)`

SetVirtualSize sets VirtualSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


