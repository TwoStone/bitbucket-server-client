# DeleteBranch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | 
**DryRun** | Pointer to **bool** |  | 
**EndPoint** | Pointer to **string** |  | [optional] 

## Methods

### NewDeleteBranch

`func NewDeleteBranch(name string, dryRun bool, ) *DeleteBranch`

NewDeleteBranch instantiates a new DeleteBranch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteBranchWithDefaults

`func NewDeleteBranchWithDefaults() *DeleteBranch`

NewDeleteBranchWithDefaults instantiates a new DeleteBranch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeleteBranch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeleteBranch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeleteBranch) SetName(v string)`

SetName sets Name field to given value.


### GetDryRun

`func (o *DeleteBranch) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteBranch) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteBranch) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.


### GetEndPoint

`func (o *DeleteBranch) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *DeleteBranch) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *DeleteBranch) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.

### HasEndPoint

`func (o *DeleteBranch) HasEndPoint() bool`

HasEndPoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


