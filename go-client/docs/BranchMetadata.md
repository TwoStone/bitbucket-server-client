# BranchMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComAtlassianBitbucketServerBitbucketBranchaheadBehindMetadataProvider** | Pointer to [**BranchMetadataAheadBehind**](branchMetadataAheadBehind.md) |  | [optional] 
**ComAtlassianBitbucketServerBitbucketBranchlatestCommitMetadata** | Pointer to [**Commit**](commit.md) |  | [optional] 
**ComAtlassianBitbucketServerBitbucketBuildbuildStatusMetadata** | Pointer to [**BranchMetadataBuildStatus**](branchMetadataBuildStatus.md) |  | [optional] 
**ComAtlassianBitbucketServerBitbucketRefMetadataoutgoingPullRequestMetadata** | Pointer to [**BranchMetadataOutgoingPullRequest**](branchMetadataOutgoingPullRequest.md) |  | [optional] 
**ComAtlassianBitbucketServerBitbucketJirabranchListJiraIssues** | Pointer to [**[]BranchMetadataJiraIssue**](branchMetadataJiraIssue.md) |  | [optional] 

## Methods

### NewBranchMetadata

`func NewBranchMetadata() *BranchMetadata`

NewBranchMetadata instantiates a new BranchMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchMetadataWithDefaults

`func NewBranchMetadataWithDefaults() *BranchMetadata`

NewBranchMetadataWithDefaults instantiates a new BranchMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComAtlassianBitbucketServerBitbucketBranchaheadBehindMetadataProvider

`func (o *BranchMetadata) GetComAtlassianBitbucketServerBitbucketBranchaheadBehindMetadataProvider() BranchMetadataAheadBehind`

GetComAtlassianBitbucketServerBitbucketBranchaheadBehindMetadataProvider returns the ComAtlassianBitbucketServerBitbucketBranchaheadBehindMetadataProvider field if non-nil, zero value otherwise.

### GetComAtlassianBitbucketServerBitbucketBranchaheadBehindMetadataProviderOk

`func (o *BranchMetadata) GetComAtlassianBitbucketServerBitbucketBranchaheadBehindMetadataProviderOk() (*BranchMetadataAheadBehind, bool)`

GetComAtlassianBitbucketServerBitbucketBranchaheadBehindMetadataProviderOk returns a tuple with the ComAtlassianBitbucketServerBitbucketBranchaheadBehindMetadataProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComAtlassianBitbucketServerBitbucketBranchaheadBehindMetadataProvider

`func (o *BranchMetadata) SetComAtlassianBitbucketServerBitbucketBranchaheadBehindMetadataProvider(v BranchMetadataAheadBehind)`

SetComAtlassianBitbucketServerBitbucketBranchaheadBehindMetadataProvider sets ComAtlassianBitbucketServerBitbucketBranchaheadBehindMetadataProvider field to given value.

### HasComAtlassianBitbucketServerBitbucketBranchaheadBehindMetadataProvider

`func (o *BranchMetadata) HasComAtlassianBitbucketServerBitbucketBranchaheadBehindMetadataProvider() bool`

HasComAtlassianBitbucketServerBitbucketBranchaheadBehindMetadataProvider returns a boolean if a field has been set.

### GetComAtlassianBitbucketServerBitbucketBranchlatestCommitMetadata

`func (o *BranchMetadata) GetComAtlassianBitbucketServerBitbucketBranchlatestCommitMetadata() Commit`

GetComAtlassianBitbucketServerBitbucketBranchlatestCommitMetadata returns the ComAtlassianBitbucketServerBitbucketBranchlatestCommitMetadata field if non-nil, zero value otherwise.

### GetComAtlassianBitbucketServerBitbucketBranchlatestCommitMetadataOk

`func (o *BranchMetadata) GetComAtlassianBitbucketServerBitbucketBranchlatestCommitMetadataOk() (*Commit, bool)`

GetComAtlassianBitbucketServerBitbucketBranchlatestCommitMetadataOk returns a tuple with the ComAtlassianBitbucketServerBitbucketBranchlatestCommitMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComAtlassianBitbucketServerBitbucketBranchlatestCommitMetadata

`func (o *BranchMetadata) SetComAtlassianBitbucketServerBitbucketBranchlatestCommitMetadata(v Commit)`

SetComAtlassianBitbucketServerBitbucketBranchlatestCommitMetadata sets ComAtlassianBitbucketServerBitbucketBranchlatestCommitMetadata field to given value.

### HasComAtlassianBitbucketServerBitbucketBranchlatestCommitMetadata

`func (o *BranchMetadata) HasComAtlassianBitbucketServerBitbucketBranchlatestCommitMetadata() bool`

HasComAtlassianBitbucketServerBitbucketBranchlatestCommitMetadata returns a boolean if a field has been set.

### GetComAtlassianBitbucketServerBitbucketBuildbuildStatusMetadata

`func (o *BranchMetadata) GetComAtlassianBitbucketServerBitbucketBuildbuildStatusMetadata() BranchMetadataBuildStatus`

GetComAtlassianBitbucketServerBitbucketBuildbuildStatusMetadata returns the ComAtlassianBitbucketServerBitbucketBuildbuildStatusMetadata field if non-nil, zero value otherwise.

### GetComAtlassianBitbucketServerBitbucketBuildbuildStatusMetadataOk

`func (o *BranchMetadata) GetComAtlassianBitbucketServerBitbucketBuildbuildStatusMetadataOk() (*BranchMetadataBuildStatus, bool)`

GetComAtlassianBitbucketServerBitbucketBuildbuildStatusMetadataOk returns a tuple with the ComAtlassianBitbucketServerBitbucketBuildbuildStatusMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComAtlassianBitbucketServerBitbucketBuildbuildStatusMetadata

`func (o *BranchMetadata) SetComAtlassianBitbucketServerBitbucketBuildbuildStatusMetadata(v BranchMetadataBuildStatus)`

SetComAtlassianBitbucketServerBitbucketBuildbuildStatusMetadata sets ComAtlassianBitbucketServerBitbucketBuildbuildStatusMetadata field to given value.

### HasComAtlassianBitbucketServerBitbucketBuildbuildStatusMetadata

`func (o *BranchMetadata) HasComAtlassianBitbucketServerBitbucketBuildbuildStatusMetadata() bool`

HasComAtlassianBitbucketServerBitbucketBuildbuildStatusMetadata returns a boolean if a field has been set.

### GetComAtlassianBitbucketServerBitbucketRefMetadataoutgoingPullRequestMetadata

`func (o *BranchMetadata) GetComAtlassianBitbucketServerBitbucketRefMetadataoutgoingPullRequestMetadata() BranchMetadataOutgoingPullRequest`

GetComAtlassianBitbucketServerBitbucketRefMetadataoutgoingPullRequestMetadata returns the ComAtlassianBitbucketServerBitbucketRefMetadataoutgoingPullRequestMetadata field if non-nil, zero value otherwise.

### GetComAtlassianBitbucketServerBitbucketRefMetadataoutgoingPullRequestMetadataOk

`func (o *BranchMetadata) GetComAtlassianBitbucketServerBitbucketRefMetadataoutgoingPullRequestMetadataOk() (*BranchMetadataOutgoingPullRequest, bool)`

GetComAtlassianBitbucketServerBitbucketRefMetadataoutgoingPullRequestMetadataOk returns a tuple with the ComAtlassianBitbucketServerBitbucketRefMetadataoutgoingPullRequestMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComAtlassianBitbucketServerBitbucketRefMetadataoutgoingPullRequestMetadata

`func (o *BranchMetadata) SetComAtlassianBitbucketServerBitbucketRefMetadataoutgoingPullRequestMetadata(v BranchMetadataOutgoingPullRequest)`

SetComAtlassianBitbucketServerBitbucketRefMetadataoutgoingPullRequestMetadata sets ComAtlassianBitbucketServerBitbucketRefMetadataoutgoingPullRequestMetadata field to given value.

### HasComAtlassianBitbucketServerBitbucketRefMetadataoutgoingPullRequestMetadata

`func (o *BranchMetadata) HasComAtlassianBitbucketServerBitbucketRefMetadataoutgoingPullRequestMetadata() bool`

HasComAtlassianBitbucketServerBitbucketRefMetadataoutgoingPullRequestMetadata returns a boolean if a field has been set.

### GetComAtlassianBitbucketServerBitbucketJirabranchListJiraIssues

`func (o *BranchMetadata) GetComAtlassianBitbucketServerBitbucketJirabranchListJiraIssues() []BranchMetadataJiraIssue`

GetComAtlassianBitbucketServerBitbucketJirabranchListJiraIssues returns the ComAtlassianBitbucketServerBitbucketJirabranchListJiraIssues field if non-nil, zero value otherwise.

### GetComAtlassianBitbucketServerBitbucketJirabranchListJiraIssuesOk

`func (o *BranchMetadata) GetComAtlassianBitbucketServerBitbucketJirabranchListJiraIssuesOk() (*[]BranchMetadataJiraIssue, bool)`

GetComAtlassianBitbucketServerBitbucketJirabranchListJiraIssuesOk returns a tuple with the ComAtlassianBitbucketServerBitbucketJirabranchListJiraIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComAtlassianBitbucketServerBitbucketJirabranchListJiraIssues

`func (o *BranchMetadata) SetComAtlassianBitbucketServerBitbucketJirabranchListJiraIssues(v []BranchMetadataJiraIssue)`

SetComAtlassianBitbucketServerBitbucketJirabranchListJiraIssues sets ComAtlassianBitbucketServerBitbucketJirabranchListJiraIssues field to given value.

### HasComAtlassianBitbucketServerBitbucketJirabranchListJiraIssues

`func (o *BranchMetadata) HasComAtlassianBitbucketServerBitbucketJirabranchListJiraIssues() bool`

HasComAtlassianBitbucketServerBitbucketJirabranchListJiraIssues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


