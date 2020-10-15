# Branch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | 
**DisplayId** | Pointer to **string** |  | 
**Type** | Pointer to **string** |  | 
**LatestCommit** | Pointer to **string** |  | 
**LatestChangeset** | Pointer to **string** |  | 
**IsDefault** | Pointer to **bool** |  | 
**Metadata** | Pointer to [**BranchMetadata**](branchMetadata.md) |  | [optional] 

## Methods

### NewBranch

`func NewBranch(id string, displayId string, type_ string, latestCommit string, latestChangeset string, isDefault bool, ) *Branch`

NewBranch instantiates a new Branch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchWithDefaults

`func NewBranchWithDefaults() *Branch`

NewBranchWithDefaults instantiates a new Branch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Branch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Branch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Branch) SetId(v string)`

SetId sets Id field to given value.


### GetDisplayId

`func (o *Branch) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *Branch) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *Branch) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.


### GetType

`func (o *Branch) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Branch) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Branch) SetType(v string)`

SetType sets Type field to given value.


### GetLatestCommit

`func (o *Branch) GetLatestCommit() string`

GetLatestCommit returns the LatestCommit field if non-nil, zero value otherwise.

### GetLatestCommitOk

`func (o *Branch) GetLatestCommitOk() (*string, bool)`

GetLatestCommitOk returns a tuple with the LatestCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestCommit

`func (o *Branch) SetLatestCommit(v string)`

SetLatestCommit sets LatestCommit field to given value.


### GetLatestChangeset

`func (o *Branch) GetLatestChangeset() string`

GetLatestChangeset returns the LatestChangeset field if non-nil, zero value otherwise.

### GetLatestChangesetOk

`func (o *Branch) GetLatestChangesetOk() (*string, bool)`

GetLatestChangesetOk returns a tuple with the LatestChangeset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestChangeset

`func (o *Branch) SetLatestChangeset(v string)`

SetLatestChangeset sets LatestChangeset field to given value.


### GetIsDefault

`func (o *Branch) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *Branch) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *Branch) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetMetadata

`func (o *Branch) GetMetadata() BranchMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Branch) GetMetadataOk() (*BranchMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Branch) SetMetadata(v BranchMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Branch) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


