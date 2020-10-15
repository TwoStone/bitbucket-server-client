# Repository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Slug** | Pointer to **string** |  | 
**Name** | Pointer to **string** |  | 
**Description** | Pointer to **string** |  | 
**HierarchyId** | Pointer to **string** |  | [optional] 
**ScmId** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**Forkable** | Pointer to **bool** |  | 
**Public** | Pointer to **bool** |  | 
**Project** | Pointer to [**Project**](project.md) |  | 
**Links** | Pointer to [**RepositoryLinks**](repository_links.md) |  | [optional] 

## Methods

### NewRepository

`func NewRepository(slug string, name string, description string, forkable bool, public bool, project Project, ) *Repository`

NewRepository instantiates a new Repository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryWithDefaults

`func NewRepositoryWithDefaults() *Repository`

NewRepositoryWithDefaults instantiates a new Repository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Repository) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Repository) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Repository) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Repository) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSlug

`func (o *Repository) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Repository) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Repository) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetName

`func (o *Repository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Repository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Repository) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Repository) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Repository) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Repository) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHierarchyId

`func (o *Repository) GetHierarchyId() string`

GetHierarchyId returns the HierarchyId field if non-nil, zero value otherwise.

### GetHierarchyIdOk

`func (o *Repository) GetHierarchyIdOk() (*string, bool)`

GetHierarchyIdOk returns a tuple with the HierarchyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchyId

`func (o *Repository) SetHierarchyId(v string)`

SetHierarchyId sets HierarchyId field to given value.

### HasHierarchyId

`func (o *Repository) HasHierarchyId() bool`

HasHierarchyId returns a boolean if a field has been set.

### GetScmId

`func (o *Repository) GetScmId() string`

GetScmId returns the ScmId field if non-nil, zero value otherwise.

### GetScmIdOk

`func (o *Repository) GetScmIdOk() (*string, bool)`

GetScmIdOk returns a tuple with the ScmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmId

`func (o *Repository) SetScmId(v string)`

SetScmId sets ScmId field to given value.

### HasScmId

`func (o *Repository) HasScmId() bool`

HasScmId returns a boolean if a field has been set.

### GetState

`func (o *Repository) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Repository) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Repository) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Repository) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatusMessage

`func (o *Repository) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *Repository) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *Repository) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *Repository) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetForkable

`func (o *Repository) GetForkable() bool`

GetForkable returns the Forkable field if non-nil, zero value otherwise.

### GetForkableOk

`func (o *Repository) GetForkableOk() (*bool, bool)`

GetForkableOk returns a tuple with the Forkable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForkable

`func (o *Repository) SetForkable(v bool)`

SetForkable sets Forkable field to given value.


### GetPublic

`func (o *Repository) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *Repository) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *Repository) SetPublic(v bool)`

SetPublic sets Public field to given value.


### GetProject

`func (o *Repository) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Repository) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Repository) SetProject(v Project)`

SetProject sets Project field to given value.


### GetLinks

`func (o *Repository) GetLinks() RepositoryLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Repository) GetLinksOk() (*RepositoryLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Repository) SetLinks(v RepositoryLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Repository) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


