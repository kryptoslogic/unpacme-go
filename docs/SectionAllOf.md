# SectionAllOf

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

### NewSectionAllOf

`func NewSectionAllOf(characteristics int32, entropy float32, mD5 string, name string, nameHex string, numberOfLinenumbers int32, numberOfRelocations int32, pointerToLinenumbers int32, pointerToRawData int32, pointerToRelocations int32, sHA1 string, sHA256 string, sizeOfRawData int32, virtualAddress int32, virtualSize int32, ) *SectionAllOf`

NewSectionAllOf instantiates a new SectionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectionAllOfWithDefaults

`func NewSectionAllOfWithDefaults() *SectionAllOf`

NewSectionAllOfWithDefaults instantiates a new SectionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCharacteristics

`func (o *SectionAllOf) GetCharacteristics() int32`

GetCharacteristics returns the Characteristics field if non-nil, zero value otherwise.

### GetCharacteristicsOk

`func (o *SectionAllOf) GetCharacteristicsOk() (*int32, bool)`

GetCharacteristicsOk returns a tuple with the Characteristics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacteristics

`func (o *SectionAllOf) SetCharacteristics(v int32)`

SetCharacteristics sets Characteristics field to given value.


### GetEntropy

`func (o *SectionAllOf) GetEntropy() float32`

GetEntropy returns the Entropy field if non-nil, zero value otherwise.

### GetEntropyOk

`func (o *SectionAllOf) GetEntropyOk() (*float32, bool)`

GetEntropyOk returns a tuple with the Entropy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntropy

`func (o *SectionAllOf) SetEntropy(v float32)`

SetEntropy sets Entropy field to given value.


### GetMD5

`func (o *SectionAllOf) GetMD5() string`

GetMD5 returns the MD5 field if non-nil, zero value otherwise.

### GetMD5Ok

`func (o *SectionAllOf) GetMD5Ok() (*string, bool)`

GetMD5Ok returns a tuple with the MD5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMD5

`func (o *SectionAllOf) SetMD5(v string)`

SetMD5 sets MD5 field to given value.


### GetName

`func (o *SectionAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SectionAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SectionAllOf) SetName(v string)`

SetName sets Name field to given value.


### GetNameHex

`func (o *SectionAllOf) GetNameHex() string`

GetNameHex returns the NameHex field if non-nil, zero value otherwise.

### GetNameHexOk

`func (o *SectionAllOf) GetNameHexOk() (*string, bool)`

GetNameHexOk returns a tuple with the NameHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameHex

`func (o *SectionAllOf) SetNameHex(v string)`

SetNameHex sets NameHex field to given value.


### GetNumberOfLinenumbers

`func (o *SectionAllOf) GetNumberOfLinenumbers() int32`

GetNumberOfLinenumbers returns the NumberOfLinenumbers field if non-nil, zero value otherwise.

### GetNumberOfLinenumbersOk

`func (o *SectionAllOf) GetNumberOfLinenumbersOk() (*int32, bool)`

GetNumberOfLinenumbersOk returns a tuple with the NumberOfLinenumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfLinenumbers

`func (o *SectionAllOf) SetNumberOfLinenumbers(v int32)`

SetNumberOfLinenumbers sets NumberOfLinenumbers field to given value.


### GetNumberOfRelocations

`func (o *SectionAllOf) GetNumberOfRelocations() int32`

GetNumberOfRelocations returns the NumberOfRelocations field if non-nil, zero value otherwise.

### GetNumberOfRelocationsOk

`func (o *SectionAllOf) GetNumberOfRelocationsOk() (*int32, bool)`

GetNumberOfRelocationsOk returns a tuple with the NumberOfRelocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfRelocations

`func (o *SectionAllOf) SetNumberOfRelocations(v int32)`

SetNumberOfRelocations sets NumberOfRelocations field to given value.


### GetPointerToLinenumbers

`func (o *SectionAllOf) GetPointerToLinenumbers() int32`

GetPointerToLinenumbers returns the PointerToLinenumbers field if non-nil, zero value otherwise.

### GetPointerToLinenumbersOk

`func (o *SectionAllOf) GetPointerToLinenumbersOk() (*int32, bool)`

GetPointerToLinenumbersOk returns a tuple with the PointerToLinenumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointerToLinenumbers

`func (o *SectionAllOf) SetPointerToLinenumbers(v int32)`

SetPointerToLinenumbers sets PointerToLinenumbers field to given value.


### GetPointerToRawData

`func (o *SectionAllOf) GetPointerToRawData() int32`

GetPointerToRawData returns the PointerToRawData field if non-nil, zero value otherwise.

### GetPointerToRawDataOk

`func (o *SectionAllOf) GetPointerToRawDataOk() (*int32, bool)`

GetPointerToRawDataOk returns a tuple with the PointerToRawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointerToRawData

`func (o *SectionAllOf) SetPointerToRawData(v int32)`

SetPointerToRawData sets PointerToRawData field to given value.


### GetPointerToRelocations

`func (o *SectionAllOf) GetPointerToRelocations() int32`

GetPointerToRelocations returns the PointerToRelocations field if non-nil, zero value otherwise.

### GetPointerToRelocationsOk

`func (o *SectionAllOf) GetPointerToRelocationsOk() (*int32, bool)`

GetPointerToRelocationsOk returns a tuple with the PointerToRelocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointerToRelocations

`func (o *SectionAllOf) SetPointerToRelocations(v int32)`

SetPointerToRelocations sets PointerToRelocations field to given value.


### GetSHA1

`func (o *SectionAllOf) GetSHA1() string`

GetSHA1 returns the SHA1 field if non-nil, zero value otherwise.

### GetSHA1Ok

`func (o *SectionAllOf) GetSHA1Ok() (*string, bool)`

GetSHA1Ok returns a tuple with the SHA1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSHA1

`func (o *SectionAllOf) SetSHA1(v string)`

SetSHA1 sets SHA1 field to given value.


### GetSHA256

`func (o *SectionAllOf) GetSHA256() string`

GetSHA256 returns the SHA256 field if non-nil, zero value otherwise.

### GetSHA256Ok

`func (o *SectionAllOf) GetSHA256Ok() (*string, bool)`

GetSHA256Ok returns a tuple with the SHA256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSHA256

`func (o *SectionAllOf) SetSHA256(v string)`

SetSHA256 sets SHA256 field to given value.


### GetSizeOfRawData

`func (o *SectionAllOf) GetSizeOfRawData() int32`

GetSizeOfRawData returns the SizeOfRawData field if non-nil, zero value otherwise.

### GetSizeOfRawDataOk

`func (o *SectionAllOf) GetSizeOfRawDataOk() (*int32, bool)`

GetSizeOfRawDataOk returns a tuple with the SizeOfRawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeOfRawData

`func (o *SectionAllOf) SetSizeOfRawData(v int32)`

SetSizeOfRawData sets SizeOfRawData field to given value.


### GetVirtualAddress

`func (o *SectionAllOf) GetVirtualAddress() int32`

GetVirtualAddress returns the VirtualAddress field if non-nil, zero value otherwise.

### GetVirtualAddressOk

`func (o *SectionAllOf) GetVirtualAddressOk() (*int32, bool)`

GetVirtualAddressOk returns a tuple with the VirtualAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualAddress

`func (o *SectionAllOf) SetVirtualAddress(v int32)`

SetVirtualAddress sets VirtualAddress field to given value.


### GetVirtualSize

`func (o *SectionAllOf) GetVirtualSize() int32`

GetVirtualSize returns the VirtualSize field if non-nil, zero value otherwise.

### GetVirtualSizeOk

`func (o *SectionAllOf) GetVirtualSizeOk() (*int32, bool)`

GetVirtualSizeOk returns a tuple with the VirtualSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualSize

`func (o *SectionAllOf) SetVirtualSize(v int32)`

SetVirtualSize sets VirtualSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


