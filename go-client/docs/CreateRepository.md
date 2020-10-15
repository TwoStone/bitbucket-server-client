# CreateRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | 
**ScmId** | Pointer to **string** |  | [optional] [default to "git"]
**Forkable** | Pointer to **bool** |  | [optional] 
**DefaultBranch** | Pointer to **string** |  | [optional] 
**Public** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateRepository

`func NewCreateRepository(name string, ) *CreateRepository`

NewCreateRepository instantiates a new CreateRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRepositoryWithDefaults

`func NewCreateRepositoryWithDefaults() *CreateRepository`

NewCreateRepositoryWithDefaults instantiates a new CreateRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRepository) SetName(v string)`

SetName sets Name field to given value.


### GetScmId

`func (o *CreateRepository) GetScmId() string`

GetScmId returns the ScmId field if non-nil, zero value otherwise.

### GetScmIdOk

`func (o *CreateRepository) GetScmIdOk() (*string, bool)`

GetScmIdOk returns a tuple with the ScmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmId

`func (o *CreateRepository) SetScmId(v string)`

SetScmId sets ScmId field to given value.

### HasScmId

`func (o *CreateRepository) HasScmId() bool`

HasScmId returns a boolean if a field has been set.

### GetForkable

`func (o *CreateRepository) GetForkable() bool`

GetForkable returns the Forkable field if non-nil, zero value otherwise.

### GetForkableOk

`func (o *CreateRepository) GetForkableOk() (*bool, bool)`

GetForkableOk returns a tuple with the Forkable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForkable

`func (o *CreateRepository) SetForkable(v bool)`

SetForkable sets Forkable field to given value.

### HasForkable

`func (o *CreateRepository) HasForkable() bool`

HasForkable returns a boolean if a field has been set.

### GetDefaultBranch

`func (o *CreateRepository) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *CreateRepository) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *CreateRepository) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.

### HasDefaultBranch

`func (o *CreateRepository) HasDefaultBranch() bool`

HasDefaultBranch returns a boolean if a field has been set.

### GetPublic

`func (o *CreateRepository) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *CreateRepository) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *CreateRepository) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *CreateRepository) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetDescription

`func (o *CreateRepository) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateRepository) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateRepository) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateRepository) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


