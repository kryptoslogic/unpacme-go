# ImportEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Library** | **string** | Library name | 
**Functions** | [**[]Function**](Function.md) |  | 

## Methods

### NewImportEntity

`func NewImportEntity(library string, functions []Function, ) *ImportEntity`

NewImportEntity instantiates a new ImportEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportEntityWithDefaults

`func NewImportEntityWithDefaults() *ImportEntity`

NewImportEntityWithDefaults instantiates a new ImportEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLibrary

`func (o *ImportEntity) GetLibrary() string`

GetLibrary returns the Library field if non-nil, zero value otherwise.

### GetLibraryOk

`func (o *ImportEntity) GetLibraryOk() (*string, bool)`

GetLibraryOk returns a tuple with the Library field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibrary

`func (o *ImportEntity) SetLibrary(v string)`

SetLibrary sets Library field to given value.


### GetFunctions

`func (o *ImportEntity) GetFunctions() []Function`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *ImportEntity) GetFunctionsOk() (*[]Function, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *ImportEntity) SetFunctions(v []Function)`

SetFunctions sets Functions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


