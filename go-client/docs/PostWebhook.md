# PostWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchCreated** | Pointer to **bool** |  | [optional] 
**BranchDeleted** | Pointer to **bool** |  | [optional] 
**BranchesToIgnore** | Pointer to **string** |  | 
**BuildStatus** | Pointer to **bool** |  | [optional] 
**CommittersToIgnore** | Pointer to **string** |  | 
**Enabled** | Pointer to **bool** |  | 
**Id** | Pointer to **int32** |  | 
**PrCommented** | Pointer to **bool** |  | [optional] 
**PrCreated** | Pointer to **bool** |  | [optional] 
**PrDeclined** | Pointer to **bool** |  | [optional] 
**PrMerged** | Pointer to **bool** |  | [optional] 
**PrReopened** | Pointer to **bool** |  | [optional] 
**PrRescoped** | Pointer to **bool** |  | [optional] 
**PrUpdated** | Pointer to **bool** |  | [optional] 
**RepoPush** | Pointer to **bool** |  | [optional] 
**TagCreated** | Pointer to **bool** |  | [optional] 
**Title** | Pointer to **string** |  | 
**Url** | Pointer to **string** |  | 

## Methods

### NewPostWebhook

`func NewPostWebhook(branchesToIgnore string, committersToIgnore string, enabled bool, id int32, title string, url string, ) *PostWebhook`

NewPostWebhook instantiates a new PostWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostWebhookWithDefaults

`func NewPostWebhookWithDefaults() *PostWebhook`

NewPostWebhookWithDefaults instantiates a new PostWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranchCreated

`func (o *PostWebhook) GetBranchCreated() bool`

GetBranchCreated returns the BranchCreated field if non-nil, zero value otherwise.

### GetBranchCreatedOk

`func (o *PostWebhook) GetBranchCreatedOk() (*bool, bool)`

GetBranchCreatedOk returns a tuple with the BranchCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchCreated

`func (o *PostWebhook) SetBranchCreated(v bool)`

SetBranchCreated sets BranchCreated field to given value.

### HasBranchCreated

`func (o *PostWebhook) HasBranchCreated() bool`

HasBranchCreated returns a boolean if a field has been set.

### GetBranchDeleted

`func (o *PostWebhook) GetBranchDeleted() bool`

GetBranchDeleted returns the BranchDeleted field if non-nil, zero value otherwise.

### GetBranchDeletedOk

`func (o *PostWebhook) GetBranchDeletedOk() (*bool, bool)`

GetBranchDeletedOk returns a tuple with the BranchDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchDeleted

`func (o *PostWebhook) SetBranchDeleted(v bool)`

SetBranchDeleted sets BranchDeleted field to given value.

### HasBranchDeleted

`func (o *PostWebhook) HasBranchDeleted() bool`

HasBranchDeleted returns a boolean if a field has been set.

### GetBranchesToIgnore

`func (o *PostWebhook) GetBranchesToIgnore() string`

GetBranchesToIgnore returns the BranchesToIgnore field if non-nil, zero value otherwise.

### GetBranchesToIgnoreOk

`func (o *PostWebhook) GetBranchesToIgnoreOk() (*string, bool)`

GetBranchesToIgnoreOk returns a tuple with the BranchesToIgnore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchesToIgnore

`func (o *PostWebhook) SetBranchesToIgnore(v string)`

SetBranchesToIgnore sets BranchesToIgnore field to given value.


### GetBuildStatus

`func (o *PostWebhook) GetBuildStatus() bool`

GetBuildStatus returns the BuildStatus field if non-nil, zero value otherwise.

### GetBuildStatusOk

`func (o *PostWebhook) GetBuildStatusOk() (*bool, bool)`

GetBuildStatusOk returns a tuple with the BuildStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildStatus

`func (o *PostWebhook) SetBuildStatus(v bool)`

SetBuildStatus sets BuildStatus field to given value.

### HasBuildStatus

`func (o *PostWebhook) HasBuildStatus() bool`

HasBuildStatus returns a boolean if a field has been set.

### GetCommittersToIgnore

`func (o *PostWebhook) GetCommittersToIgnore() string`

GetCommittersToIgnore returns the CommittersToIgnore field if non-nil, zero value otherwise.

### GetCommittersToIgnoreOk

`func (o *PostWebhook) GetCommittersToIgnoreOk() (*string, bool)`

GetCommittersToIgnoreOk returns a tuple with the CommittersToIgnore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommittersToIgnore

`func (o *PostWebhook) SetCommittersToIgnore(v string)`

SetCommittersToIgnore sets CommittersToIgnore field to given value.


### GetEnabled

`func (o *PostWebhook) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PostWebhook) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PostWebhook) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetId

`func (o *PostWebhook) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostWebhook) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostWebhook) SetId(v int32)`

SetId sets Id field to given value.


### GetPrCommented

`func (o *PostWebhook) GetPrCommented() bool`

GetPrCommented returns the PrCommented field if non-nil, zero value otherwise.

### GetPrCommentedOk

`func (o *PostWebhook) GetPrCommentedOk() (*bool, bool)`

GetPrCommentedOk returns a tuple with the PrCommented field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrCommented

`func (o *PostWebhook) SetPrCommented(v bool)`

SetPrCommented sets PrCommented field to given value.

### HasPrCommented

`func (o *PostWebhook) HasPrCommented() bool`

HasPrCommented returns a boolean if a field has been set.

### GetPrCreated

`func (o *PostWebhook) GetPrCreated() bool`

GetPrCreated returns the PrCreated field if non-nil, zero value otherwise.

### GetPrCreatedOk

`func (o *PostWebhook) GetPrCreatedOk() (*bool, bool)`

GetPrCreatedOk returns a tuple with the PrCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrCreated

`func (o *PostWebhook) SetPrCreated(v bool)`

SetPrCreated sets PrCreated field to given value.

### HasPrCreated

`func (o *PostWebhook) HasPrCreated() bool`

HasPrCreated returns a boolean if a field has been set.

### GetPrDeclined

`func (o *PostWebhook) GetPrDeclined() bool`

GetPrDeclined returns the PrDeclined field if non-nil, zero value otherwise.

### GetPrDeclinedOk

`func (o *PostWebhook) GetPrDeclinedOk() (*bool, bool)`

GetPrDeclinedOk returns a tuple with the PrDeclined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrDeclined

`func (o *PostWebhook) SetPrDeclined(v bool)`

SetPrDeclined sets PrDeclined field to given value.

### HasPrDeclined

`func (o *PostWebhook) HasPrDeclined() bool`

HasPrDeclined returns a boolean if a field has been set.

### GetPrMerged

`func (o *PostWebhook) GetPrMerged() bool`

GetPrMerged returns the PrMerged field if non-nil, zero value otherwise.

### GetPrMergedOk

`func (o *PostWebhook) GetPrMergedOk() (*bool, bool)`

GetPrMergedOk returns a tuple with the PrMerged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrMerged

`func (o *PostWebhook) SetPrMerged(v bool)`

SetPrMerged sets PrMerged field to given value.

### HasPrMerged

`func (o *PostWebhook) HasPrMerged() bool`

HasPrMerged returns a boolean if a field has been set.

### GetPrReopened

`func (o *PostWebhook) GetPrReopened() bool`

GetPrReopened returns the PrReopened field if non-nil, zero value otherwise.

### GetPrReopenedOk

`func (o *PostWebhook) GetPrReopenedOk() (*bool, bool)`

GetPrReopenedOk returns a tuple with the PrReopened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrReopened

`func (o *PostWebhook) SetPrReopened(v bool)`

SetPrReopened sets PrReopened field to given value.

### HasPrReopened

`func (o *PostWebhook) HasPrReopened() bool`

HasPrReopened returns a boolean if a field has been set.

### GetPrRescoped

`func (o *PostWebhook) GetPrRescoped() bool`

GetPrRescoped returns the PrRescoped field if non-nil, zero value otherwise.

### GetPrRescopedOk

`func (o *PostWebhook) GetPrRescopedOk() (*bool, bool)`

GetPrRescopedOk returns a tuple with the PrRescoped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrRescoped

`func (o *PostWebhook) SetPrRescoped(v bool)`

SetPrRescoped sets PrRescoped field to given value.

### HasPrRescoped

`func (o *PostWebhook) HasPrRescoped() bool`

HasPrRescoped returns a boolean if a field has been set.

### GetPrUpdated

`func (o *PostWebhook) GetPrUpdated() bool`

GetPrUpdated returns the PrUpdated field if non-nil, zero value otherwise.

### GetPrUpdatedOk

`func (o *PostWebhook) GetPrUpdatedOk() (*bool, bool)`

GetPrUpdatedOk returns a tuple with the PrUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrUpdated

`func (o *PostWebhook) SetPrUpdated(v bool)`

SetPrUpdated sets PrUpdated field to given value.

### HasPrUpdated

`func (o *PostWebhook) HasPrUpdated() bool`

HasPrUpdated returns a boolean if a field has been set.

### GetRepoPush

`func (o *PostWebhook) GetRepoPush() bool`

GetRepoPush returns the RepoPush field if non-nil, zero value otherwise.

### GetRepoPushOk

`func (o *PostWebhook) GetRepoPushOk() (*bool, bool)`

GetRepoPushOk returns a tuple with the RepoPush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoPush

`func (o *PostWebhook) SetRepoPush(v bool)`

SetRepoPush sets RepoPush field to given value.

### HasRepoPush

`func (o *PostWebhook) HasRepoPush() bool`

HasRepoPush returns a boolean if a field has been set.

### GetTagCreated

`func (o *PostWebhook) GetTagCreated() bool`

GetTagCreated returns the TagCreated field if non-nil, zero value otherwise.

### GetTagCreatedOk

`func (o *PostWebhook) GetTagCreatedOk() (*bool, bool)`

GetTagCreatedOk returns a tuple with the TagCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagCreated

`func (o *PostWebhook) SetTagCreated(v bool)`

SetTagCreated sets TagCreated field to given value.

### HasTagCreated

`func (o *PostWebhook) HasTagCreated() bool`

HasTagCreated returns a boolean if a field has been set.

### GetTitle

`func (o *PostWebhook) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PostWebhook) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PostWebhook) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUrl

`func (o *PostWebhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PostWebhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PostWebhook) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


