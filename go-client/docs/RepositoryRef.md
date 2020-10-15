# RepositoryRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | 
**Repository** | Pointer to [**RepositoryRefRepository**](repositoryRef_repository.md) |  | 

## Methods

### NewRepositoryRef

`func NewRepositoryRef(id string, repository RepositoryRefRepository, ) *RepositoryRef`

NewRepositoryRef instantiates a new RepositoryRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryRefWithDefaults

`func NewRepositoryRefWithDefaults() *RepositoryRef`

NewRepositoryRefWithDefaults instantiates a new RepositoryRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RepositoryRef) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RepositoryRef) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RepositoryRef) SetId(v string)`

SetId sets Id field to given value.


### GetRepository

`func (o *RepositoryRef) GetRepository() RepositoryRefRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *RepositoryRef) GetRepositoryOk() (*RepositoryRefRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *RepositoryRef) SetRepository(v RepositoryRefRepository)`

SetRepository sets Repository field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


