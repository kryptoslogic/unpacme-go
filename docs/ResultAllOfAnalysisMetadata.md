# ResultAllOfAnalysisMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Characteristics** | Pointer to **[]string** |  | [optional] 
**Checksum** | Pointer to **int32** | PE file checksum | [optional] 
**CompileTime** | Pointer to **string** | PE file compile time | [optional] 
**ContainsCompressedData** | Pointer to **bool** | PE file contains compressed data | [optional] 
**EPBytes** | Pointer to **string** | Entry point first 16 bytes | [optional] 
**EntryPoint** | Pointer to **int32** | PE file entry point | [optional] 
**ImageBase** | Pointer to **int32** | PE file image base | [optional] 
**LinkerVersion** | Pointer to **string** | PE file linker version | [optional] 
**PDBPath** | Pointer to **string** | PE file program database file path | [optional] 
**Sections** | Pointer to **int32** | Number of sections | [optional] 
**Signature** | Pointer to **int32** | PE file signature | [optional] 
**Size** | Pointer to **int32** | PE file size | [optional] 
**Subsystem** | Pointer to **string** | PE file subsystem | [optional] 
**Type** | Pointer to **string** | PE file type | [optional] 
**VersionInfo** | Pointer to [**ResultAllOfAnalysisMetadataVersionInfo**](result_allOf_analysis_metadata_VersionInfo.md) |  | [optional] 

## Methods

### NewResultAllOfAnalysisMetadata

`func NewResultAllOfAnalysisMetadata() *ResultAllOfAnalysisMetadata`

NewResultAllOfAnalysisMetadata instantiates a new ResultAllOfAnalysisMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultAllOfAnalysisMetadataWithDefaults

`func NewResultAllOfAnalysisMetadataWithDefaults() *ResultAllOfAnalysisMetadata`

NewResultAllOfAnalysisMetadataWithDefaults instantiates a new ResultAllOfAnalysisMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCharacteristics

`func (o *ResultAllOfAnalysisMetadata) GetCharacteristics() []string`

GetCharacteristics returns the Characteristics field if non-nil, zero value otherwise.

### GetCharacteristicsOk

`func (o *ResultAllOfAnalysisMetadata) GetCharacteristicsOk() (*[]string, bool)`

GetCharacteristicsOk returns a tuple with the Characteristics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacteristics

`func (o *ResultAllOfAnalysisMetadata) SetCharacteristics(v []string)`

SetCharacteristics sets Characteristics field to given value.

### HasCharacteristics

`func (o *ResultAllOfAnalysisMetadata) HasCharacteristics() bool`

HasCharacteristics returns a boolean if a field has been set.

### GetChecksum

`func (o *ResultAllOfAnalysisMetadata) GetChecksum() int32`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *ResultAllOfAnalysisMetadata) GetChecksumOk() (*int32, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *ResultAllOfAnalysisMetadata) SetChecksum(v int32)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *ResultAllOfAnalysisMetadata) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetCompileTime

`func (o *ResultAllOfAnalysisMetadata) GetCompileTime() string`

GetCompileTime returns the CompileTime field if non-nil, zero value otherwise.

### GetCompileTimeOk

`func (o *ResultAllOfAnalysisMetadata) GetCompileTimeOk() (*string, bool)`

GetCompileTimeOk returns a tuple with the CompileTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompileTime

`func (o *ResultAllOfAnalysisMetadata) SetCompileTime(v string)`

SetCompileTime sets CompileTime field to given value.

### HasCompileTime

`func (o *ResultAllOfAnalysisMetadata) HasCompileTime() bool`

HasCompileTime returns a boolean if a field has been set.

### GetContainsCompressedData

`func (o *ResultAllOfAnalysisMetadata) GetContainsCompressedData() bool`

GetContainsCompressedData returns the ContainsCompressedData field if non-nil, zero value otherwise.

### GetContainsCompressedDataOk

`func (o *ResultAllOfAnalysisMetadata) GetContainsCompressedDataOk() (*bool, bool)`

GetContainsCompressedDataOk returns a tuple with the ContainsCompressedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsCompressedData

`func (o *ResultAllOfAnalysisMetadata) SetContainsCompressedData(v bool)`

SetContainsCompressedData sets ContainsCompressedData field to given value.

### HasContainsCompressedData

`func (o *ResultAllOfAnalysisMetadata) HasContainsCompressedData() bool`

HasContainsCompressedData returns a boolean if a field has been set.

### GetEPBytes

`func (o *ResultAllOfAnalysisMetadata) GetEPBytes() string`

GetEPBytes returns the EPBytes field if non-nil, zero value otherwise.

### GetEPBytesOk

`func (o *ResultAllOfAnalysisMetadata) GetEPBytesOk() (*string, bool)`

GetEPBytesOk returns a tuple with the EPBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPBytes

`func (o *ResultAllOfAnalysisMetadata) SetEPBytes(v string)`

SetEPBytes sets EPBytes field to given value.

### HasEPBytes

`func (o *ResultAllOfAnalysisMetadata) HasEPBytes() bool`

HasEPBytes returns a boolean if a field has been set.

### GetEntryPoint

`func (o *ResultAllOfAnalysisMetadata) GetEntryPoint() int32`

GetEntryPoint returns the EntryPoint field if non-nil, zero value otherwise.

### GetEntryPointOk

`func (o *ResultAllOfAnalysisMetadata) GetEntryPointOk() (*int32, bool)`

GetEntryPointOk returns a tuple with the EntryPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryPoint

`func (o *ResultAllOfAnalysisMetadata) SetEntryPoint(v int32)`

SetEntryPoint sets EntryPoint field to given value.

### HasEntryPoint

`func (o *ResultAllOfAnalysisMetadata) HasEntryPoint() bool`

HasEntryPoint returns a boolean if a field has been set.

### GetImageBase

`func (o *ResultAllOfAnalysisMetadata) GetImageBase() int32`

GetImageBase returns the ImageBase field if non-nil, zero value otherwise.

### GetImageBaseOk

`func (o *ResultAllOfAnalysisMetadata) GetImageBaseOk() (*int32, bool)`

GetImageBaseOk returns a tuple with the ImageBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBase

`func (o *ResultAllOfAnalysisMetadata) SetImageBase(v int32)`

SetImageBase sets ImageBase field to given value.

### HasImageBase

`func (o *ResultAllOfAnalysisMetadata) HasImageBase() bool`

HasImageBase returns a boolean if a field has been set.

### GetLinkerVersion

`func (o *ResultAllOfAnalysisMetadata) GetLinkerVersion() string`

GetLinkerVersion returns the LinkerVersion field if non-nil, zero value otherwise.

### GetLinkerVersionOk

`func (o *ResultAllOfAnalysisMetadata) GetLinkerVersionOk() (*string, bool)`

GetLinkerVersionOk returns a tuple with the LinkerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkerVersion

`func (o *ResultAllOfAnalysisMetadata) SetLinkerVersion(v string)`

SetLinkerVersion sets LinkerVersion field to given value.

### HasLinkerVersion

`func (o *ResultAllOfAnalysisMetadata) HasLinkerVersion() bool`

HasLinkerVersion returns a boolean if a field has been set.

### GetPDBPath

`func (o *ResultAllOfAnalysisMetadata) GetPDBPath() string`

GetPDBPath returns the PDBPath field if non-nil, zero value otherwise.

### GetPDBPathOk

`func (o *ResultAllOfAnalysisMetadata) GetPDBPathOk() (*string, bool)`

GetPDBPathOk returns a tuple with the PDBPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPDBPath

`func (o *ResultAllOfAnalysisMetadata) SetPDBPath(v string)`

SetPDBPath sets PDBPath field to given value.

### HasPDBPath

`func (o *ResultAllOfAnalysisMetadata) HasPDBPath() bool`

HasPDBPath returns a boolean if a field has been set.

### GetSections

`func (o *ResultAllOfAnalysisMetadata) GetSections() int32`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *ResultAllOfAnalysisMetadata) GetSectionsOk() (*int32, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *ResultAllOfAnalysisMetadata) SetSections(v int32)`

SetSections sets Sections field to given value.

### HasSections

`func (o *ResultAllOfAnalysisMetadata) HasSections() bool`

HasSections returns a boolean if a field has been set.

### GetSignature

`func (o *ResultAllOfAnalysisMetadata) GetSignature() int32`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *ResultAllOfAnalysisMetadata) GetSignatureOk() (*int32, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *ResultAllOfAnalysisMetadata) SetSignature(v int32)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *ResultAllOfAnalysisMetadata) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetSize

`func (o *ResultAllOfAnalysisMetadata) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ResultAllOfAnalysisMetadata) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ResultAllOfAnalysisMetadata) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ResultAllOfAnalysisMetadata) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSubsystem

`func (o *ResultAllOfAnalysisMetadata) GetSubsystem() string`

GetSubsystem returns the Subsystem field if non-nil, zero value otherwise.

### GetSubsystemOk

`func (o *ResultAllOfAnalysisMetadata) GetSubsystemOk() (*string, bool)`

GetSubsystemOk returns a tuple with the Subsystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsystem

`func (o *ResultAllOfAnalysisMetadata) SetSubsystem(v string)`

SetSubsystem sets Subsystem field to given value.

### HasSubsystem

`func (o *ResultAllOfAnalysisMetadata) HasSubsystem() bool`

HasSubsystem returns a boolean if a field has been set.

### GetType

`func (o *ResultAllOfAnalysisMetadata) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResultAllOfAnalysisMetadata) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResultAllOfAnalysisMetadata) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResultAllOfAnalysisMetadata) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersionInfo

`func (o *ResultAllOfAnalysisMetadata) GetVersionInfo() ResultAllOfAnalysisMetadataVersionInfo`

GetVersionInfo returns the VersionInfo field if non-nil, zero value otherwise.

### GetVersionInfoOk

`func (o *ResultAllOfAnalysisMetadata) GetVersionInfoOk() (*ResultAllOfAnalysisMetadataVersionInfo, bool)`

GetVersionInfoOk returns a tuple with the VersionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionInfo

`func (o *ResultAllOfAnalysisMetadata) SetVersionInfo(v ResultAllOfAnalysisMetadataVersionInfo)`

SetVersionInfo sets VersionInfo field to given value.

### HasVersionInfo

`func (o *ResultAllOfAnalysisMetadata) HasVersionInfo() bool`

HasVersionInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


