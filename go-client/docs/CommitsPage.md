# CommitsPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | Pointer to **int32** |  | 
**Limit** | Pointer to **int32** |  | 
**Start** | Pointer to **int32** |  | 
**IsLastPage** | Pointer to **bool** |  | 
**NextPageStart** | Pointer to **int32** |  | [optional] 
**Values** | Pointer to [**[]Commit**](commit.md) |  | 
**AuthorCount** | Pointer to **int32** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewCommitsPage

`func NewCommitsPage(size int32, limit int32, start int32, isLastPage bool, values []Commit, ) *CommitsPage`

NewCommitsPage instantiates a new CommitsPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitsPageWithDefaults

`func NewCommitsPageWithDefaults() *CommitsPage`

NewCommitsPageWithDefaults instantiates a new CommitsPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *CommitsPage) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CommitsPage) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CommitsPage) SetSize(v int32)`

SetSize sets Size field to given value.


### GetLimit

`func (o *CommitsPage) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CommitsPage) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CommitsPage) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetStart

`func (o *CommitsPage) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *CommitsPage) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *CommitsPage) SetStart(v int32)`

SetStart sets Start field to given value.


### GetIsLastPage

`func (o *CommitsPage) GetIsLastPage() bool`

GetIsLastPage returns the IsLastPage field if non-nil, zero value otherwise.

### GetIsLastPageOk

`func (o *CommitsPage) GetIsLastPageOk() (*bool, bool)`

GetIsLastPageOk returns a tuple with the IsLastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLastPage

`func (o *CommitsPage) SetIsLastPage(v bool)`

SetIsLastPage sets IsLastPage field to given value.


### GetNextPageStart

`func (o *CommitsPage) GetNextPageStart() int32`

GetNextPageStart returns the NextPageStart field if non-nil, zero value otherwise.

### GetNextPageStartOk

`func (o *CommitsPage) GetNextPageStartOk() (*int32, bool)`

GetNextPageStartOk returns a tuple with the NextPageStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageStart

`func (o *CommitsPage) SetNextPageStart(v int32)`

SetNextPageStart sets NextPageStart field to given value.

### HasNextPageStart

`func (o *CommitsPage) HasNextPageStart() bool`

HasNextPageStart returns a boolean if a field has been set.

### GetValues

`func (o *CommitsPage) GetValues() []Commit`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CommitsPage) GetValuesOk() (*[]Commit, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CommitsPage) SetValues(v []Commit)`

SetValues sets Values field to given value.


### GetAuthorCount

`func (o *CommitsPage) GetAuthorCount() int32`

GetAuthorCount returns the AuthorCount field if non-nil, zero value otherwise.

### GetAuthorCountOk

`func (o *CommitsPage) GetAuthorCountOk() (*int32, bool)`

GetAuthorCountOk returns a tuple with the AuthorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorCount

`func (o *CommitsPage) SetAuthorCount(v int32)`

SetAuthorCount sets AuthorCount field to given value.

### HasAuthorCount

`func (o *CommitsPage) HasAuthorCount() bool`

HasAuthorCount returns a boolean if a field has been set.

### GetTotalCount

`func (o *CommitsPage) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CommitsPage) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CommitsPage) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *CommitsPage) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


