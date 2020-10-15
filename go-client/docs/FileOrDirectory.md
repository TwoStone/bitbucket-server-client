# FileOrDirectory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **int32** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**IsLastPage** | Pointer to **bool** |  | [optional] 
**Lines** | Pointer to [**[]FileLines**](file_lines.md) |  | [optional] 
**Path** | Pointer to [**Path**](path.md) |  | [optional] 
**Revision** | Pointer to **string** |  | [optional] 
**Children** | Pointer to [**Children**](children.md) |  | [optional] 

## Methods

### NewFileOrDirectory

`func NewFileOrDirectory() *FileOrDirectory`

NewFileOrDirectory instantiates a new FileOrDirectory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileOrDirectoryWithDefaults

`func NewFileOrDirectoryWithDefaults() *FileOrDirectory`

NewFileOrDirectoryWithDefaults instantiates a new FileOrDirectory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *FileOrDirectory) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *FileOrDirectory) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *FileOrDirectory) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *FileOrDirectory) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetSize

`func (o *FileOrDirectory) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FileOrDirectory) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FileOrDirectory) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *FileOrDirectory) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetIsLastPage

`func (o *FileOrDirectory) GetIsLastPage() bool`

GetIsLastPage returns the IsLastPage field if non-nil, zero value otherwise.

### GetIsLastPageOk

`func (o *FileOrDirectory) GetIsLastPageOk() (*bool, bool)`

GetIsLastPageOk returns a tuple with the IsLastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLastPage

`func (o *FileOrDirectory) SetIsLastPage(v bool)`

SetIsLastPage sets IsLastPage field to given value.

### HasIsLastPage

`func (o *FileOrDirectory) HasIsLastPage() bool`

HasIsLastPage returns a boolean if a field has been set.

### GetLines

`func (o *FileOrDirectory) GetLines() []FileLines`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *FileOrDirectory) GetLinesOk() (*[]FileLines, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *FileOrDirectory) SetLines(v []FileLines)`

SetLines sets Lines field to given value.

### HasLines

`func (o *FileOrDirectory) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetPath

`func (o *FileOrDirectory) GetPath() Path`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FileOrDirectory) GetPathOk() (*Path, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FileOrDirectory) SetPath(v Path)`

SetPath sets Path field to given value.

### HasPath

`func (o *FileOrDirectory) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRevision

`func (o *FileOrDirectory) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *FileOrDirectory) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *FileOrDirectory) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *FileOrDirectory) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetChildren

`func (o *FileOrDirectory) GetChildren() Children`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *FileOrDirectory) GetChildrenOk() (*Children, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *FileOrDirectory) SetChildren(v Children)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *FileOrDirectory) HasChildren() bool`

HasChildren returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


