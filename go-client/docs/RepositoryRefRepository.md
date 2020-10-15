# RepositoryRefRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slug** | Pointer to **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Project** | Pointer to [**RepositoryRefRepositoryProject**](repositoryRef_repository_project.md) |  | 

## Methods

### NewRepositoryRefRepository

`func NewRepositoryRefRepository(slug string, project RepositoryRefRepositoryProject, ) *RepositoryRefRepository`

NewRepositoryRefRepository instantiates a new RepositoryRefRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryRefRepositoryWithDefaults

`func NewRepositoryRefRepositoryWithDefaults() *RepositoryRefRepository`

NewRepositoryRefRepositoryWithDefaults instantiates a new RepositoryRefRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlug

`func (o *RepositoryRefRepository) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *RepositoryRefRepository) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *RepositoryRefRepository) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetName

`func (o *RepositoryRefRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RepositoryRefRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RepositoryRefRepository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RepositoryRefRepository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProject

`func (o *RepositoryRefRepository) GetProject() RepositoryRefRepositoryProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *RepositoryRefRepository) GetProjectOk() (*RepositoryRefRepositoryProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *RepositoryRefRepository) SetProject(v RepositoryRefRepositoryProject)`

SetProject sets Project field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


