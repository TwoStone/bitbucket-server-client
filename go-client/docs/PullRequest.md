# PullRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | 
**Description** | Pointer to **string** |  | 
**State** | Pointer to [**PullRequestState**](pullRequestState.md) |  | [optional] 
**Open** | Pointer to **bool** |  | [optional] 
**Closed** | Pointer to **bool** |  | [optional] 
**CreatedDate** | Pointer to **int64** |  | [optional] 
**UpdatedDate** | Pointer to **int64** |  | [optional] 
**FromRef** | Pointer to [**RepositoryRef**](repositoryRef.md) |  | 
**ToRef** | Pointer to [**RepositoryRef**](repositoryRef.md) |  | 
**Locked** | Pointer to **bool** |  | [optional] 
**Author** | Pointer to [**UserRole**](userRole.md) |  | [optional] 
**Reviewers** | Pointer to [**[]UserRole**](userRole.md) |  | [optional] 
**Participants** | Pointer to [**[]UserRole**](userRole.md) |  | [optional] 
**Links** | Pointer to [**ProjectLinks**](project_links.md) |  | [optional] 

## Methods

### NewPullRequest

`func NewPullRequest(title string, description string, fromRef RepositoryRef, toRef RepositoryRef, ) *PullRequest`

NewPullRequest instantiates a new PullRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullRequestWithDefaults

`func NewPullRequestWithDefaults() *PullRequest`

NewPullRequestWithDefaults instantiates a new PullRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PullRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PullRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PullRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PullRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersion

`func (o *PullRequest) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PullRequest) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PullRequest) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PullRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetTitle

`func (o *PullRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PullRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PullRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *PullRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PullRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PullRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetState

`func (o *PullRequest) GetState() PullRequestState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PullRequest) GetStateOk() (*PullRequestState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PullRequest) SetState(v PullRequestState)`

SetState sets State field to given value.

### HasState

`func (o *PullRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetOpen

`func (o *PullRequest) GetOpen() bool`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *PullRequest) GetOpenOk() (*bool, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *PullRequest) SetOpen(v bool)`

SetOpen sets Open field to given value.

### HasOpen

`func (o *PullRequest) HasOpen() bool`

HasOpen returns a boolean if a field has been set.

### GetClosed

`func (o *PullRequest) GetClosed() bool`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *PullRequest) GetClosedOk() (*bool, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *PullRequest) SetClosed(v bool)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *PullRequest) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetCreatedDate

`func (o *PullRequest) GetCreatedDate() int64`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *PullRequest) GetCreatedDateOk() (*int64, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *PullRequest) SetCreatedDate(v int64)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *PullRequest) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetUpdatedDate

`func (o *PullRequest) GetUpdatedDate() int64`

GetUpdatedDate returns the UpdatedDate field if non-nil, zero value otherwise.

### GetUpdatedDateOk

`func (o *PullRequest) GetUpdatedDateOk() (*int64, bool)`

GetUpdatedDateOk returns a tuple with the UpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDate

`func (o *PullRequest) SetUpdatedDate(v int64)`

SetUpdatedDate sets UpdatedDate field to given value.

### HasUpdatedDate

`func (o *PullRequest) HasUpdatedDate() bool`

HasUpdatedDate returns a boolean if a field has been set.

### GetFromRef

`func (o *PullRequest) GetFromRef() RepositoryRef`

GetFromRef returns the FromRef field if non-nil, zero value otherwise.

### GetFromRefOk

`func (o *PullRequest) GetFromRefOk() (*RepositoryRef, bool)`

GetFromRefOk returns a tuple with the FromRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromRef

`func (o *PullRequest) SetFromRef(v RepositoryRef)`

SetFromRef sets FromRef field to given value.


### GetToRef

`func (o *PullRequest) GetToRef() RepositoryRef`

GetToRef returns the ToRef field if non-nil, zero value otherwise.

### GetToRefOk

`func (o *PullRequest) GetToRefOk() (*RepositoryRef, bool)`

GetToRefOk returns a tuple with the ToRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToRef

`func (o *PullRequest) SetToRef(v RepositoryRef)`

SetToRef sets ToRef field to given value.


### GetLocked

`func (o *PullRequest) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *PullRequest) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *PullRequest) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *PullRequest) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetAuthor

`func (o *PullRequest) GetAuthor() UserRole`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *PullRequest) GetAuthorOk() (*UserRole, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *PullRequest) SetAuthor(v UserRole)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *PullRequest) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetReviewers

`func (o *PullRequest) GetReviewers() []UserRole`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *PullRequest) GetReviewersOk() (*[]UserRole, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *PullRequest) SetReviewers(v []UserRole)`

SetReviewers sets Reviewers field to given value.

### HasReviewers

`func (o *PullRequest) HasReviewers() bool`

HasReviewers returns a boolean if a field has been set.

### GetParticipants

`func (o *PullRequest) GetParticipants() []UserRole`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *PullRequest) GetParticipantsOk() (*[]UserRole, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *PullRequest) SetParticipants(v []UserRole)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *PullRequest) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.

### GetLinks

`func (o *PullRequest) GetLinks() ProjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PullRequest) GetLinksOk() (*ProjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PullRequest) SetLinks(v ProjectLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PullRequest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


