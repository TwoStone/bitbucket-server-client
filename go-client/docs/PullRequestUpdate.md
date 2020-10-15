# PullRequestUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Reviewers** | Pointer to [**[]UserRole**](userRole.md) |  | [optional] 
**ToRef** | Pointer to [**RepositoryRef**](repositoryRef.md) |  | [optional] 

## Methods

### NewPullRequestUpdate

`func NewPullRequestUpdate() *PullRequestUpdate`

NewPullRequestUpdate instantiates a new PullRequestUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullRequestUpdateWithDefaults

`func NewPullRequestUpdateWithDefaults() *PullRequestUpdate`

NewPullRequestUpdateWithDefaults instantiates a new PullRequestUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PullRequestUpdate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PullRequestUpdate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PullRequestUpdate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PullRequestUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersion

`func (o *PullRequestUpdate) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PullRequestUpdate) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PullRequestUpdate) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PullRequestUpdate) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetTitle

`func (o *PullRequestUpdate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PullRequestUpdate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PullRequestUpdate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PullRequestUpdate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *PullRequestUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PullRequestUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PullRequestUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PullRequestUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetReviewers

`func (o *PullRequestUpdate) GetReviewers() []UserRole`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *PullRequestUpdate) GetReviewersOk() (*[]UserRole, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *PullRequestUpdate) SetReviewers(v []UserRole)`

SetReviewers sets Reviewers field to given value.

### HasReviewers

`func (o *PullRequestUpdate) HasReviewers() bool`

HasReviewers returns a boolean if a field has been set.

### GetToRef

`func (o *PullRequestUpdate) GetToRef() RepositoryRef`

GetToRef returns the ToRef field if non-nil, zero value otherwise.

### GetToRefOk

`func (o *PullRequestUpdate) GetToRefOk() (*RepositoryRef, bool)`

GetToRefOk returns a tuple with the ToRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToRef

`func (o *PullRequestUpdate) SetToRef(v RepositoryRef)`

SetToRef sets ToRef field to given value.

### HasToRef

`func (o *PullRequestUpdate) HasToRef() bool`

HasToRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


