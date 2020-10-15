# Directory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to [**Path**](path.md) |  | [optional] 
**Revision** | Pointer to **string** |  | [optional] 
**Children** | Pointer to [**Children**](children.md) |  | [optional] 

## Methods

### NewDirectory

`func NewDirectory() *Directory`

NewDirectory instantiates a new Directory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectoryWithDefaults

`func NewDirectoryWithDefaults() *Directory`

NewDirectoryWithDefaults instantiates a new Directory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *Directory) GetPath() Path`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Directory) GetPathOk() (*Path, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Directory) SetPath(v Path)`

SetPath sets Path field to given value.

### HasPath

`func (o *Directory) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRevision

`func (o *Directory) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *Directory) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *Directory) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *Directory) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetChildren

`func (o *Directory) GetChildren() Children`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *Directory) GetChildrenOk() (*Children, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *Directory) SetChildren(v Children)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *Directory) HasChildren() bool`

HasChildren returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


