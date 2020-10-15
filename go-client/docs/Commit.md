# Commit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | 
**DisplayId** | Pointer to **string** |  | 
**Author** | Pointer to [**User**](user.md) |  | 
**AuthorTimestamp** | Pointer to **int64** |  | 
**Commiter** | Pointer to [**User**](user.md) |  | 
**CommiterTimestamp** | Pointer to **int64** |  | 
**Message** | Pointer to **string** |  | 
**Parents** | Pointer to [**[]CommitParents**](commit_parents.md) |  | [optional] 

## Methods

### NewCommit

`func NewCommit(id string, displayId string, author User, authorTimestamp int64, commiter User, commiterTimestamp int64, message string, ) *Commit`

NewCommit instantiates a new Commit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitWithDefaults

`func NewCommitWithDefaults() *Commit`

NewCommitWithDefaults instantiates a new Commit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Commit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Commit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Commit) SetId(v string)`

SetId sets Id field to given value.


### GetDisplayId

`func (o *Commit) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *Commit) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *Commit) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.


### GetAuthor

`func (o *Commit) GetAuthor() User`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *Commit) GetAuthorOk() (*User, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *Commit) SetAuthor(v User)`

SetAuthor sets Author field to given value.


### GetAuthorTimestamp

`func (o *Commit) GetAuthorTimestamp() int64`

GetAuthorTimestamp returns the AuthorTimestamp field if non-nil, zero value otherwise.

### GetAuthorTimestampOk

`func (o *Commit) GetAuthorTimestampOk() (*int64, bool)`

GetAuthorTimestampOk returns a tuple with the AuthorTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorTimestamp

`func (o *Commit) SetAuthorTimestamp(v int64)`

SetAuthorTimestamp sets AuthorTimestamp field to given value.


### GetCommiter

`func (o *Commit) GetCommiter() User`

GetCommiter returns the Commiter field if non-nil, zero value otherwise.

### GetCommiterOk

`func (o *Commit) GetCommiterOk() (*User, bool)`

GetCommiterOk returns a tuple with the Commiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommiter

`func (o *Commit) SetCommiter(v User)`

SetCommiter sets Commiter field to given value.


### GetCommiterTimestamp

`func (o *Commit) GetCommiterTimestamp() int64`

GetCommiterTimestamp returns the CommiterTimestamp field if non-nil, zero value otherwise.

### GetCommiterTimestampOk

`func (o *Commit) GetCommiterTimestampOk() (*int64, bool)`

GetCommiterTimestampOk returns a tuple with the CommiterTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommiterTimestamp

`func (o *Commit) SetCommiterTimestamp(v int64)`

SetCommiterTimestamp sets CommiterTimestamp field to given value.


### GetMessage

`func (o *Commit) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Commit) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Commit) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetParents

`func (o *Commit) GetParents() []CommitParents`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *Commit) GetParentsOk() (*[]CommitParents, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *Commit) SetParents(v []CommitParents)`

SetParents sets Parents field to given value.

### HasParents

`func (o *Commit) HasParents() bool`

HasParents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


