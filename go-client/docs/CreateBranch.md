# CreateBranch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | 
**StartPoint** | Pointer to **string** |  | 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateBranch

`func NewCreateBranch(name string, startPoint string, ) *CreateBranch`

NewCreateBranch instantiates a new CreateBranch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBranchWithDefaults

`func NewCreateBranchWithDefaults() *CreateBranch`

NewCreateBranchWithDefaults instantiates a new CreateBranch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateBranch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateBranch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateBranch) SetName(v string)`

SetName sets Name field to given value.


### GetStartPoint

`func (o *CreateBranch) GetStartPoint() string`

GetStartPoint returns the StartPoint field if non-nil, zero value otherwise.

### GetStartPointOk

`func (o *CreateBranch) GetStartPointOk() (*string, bool)`

GetStartPointOk returns a tuple with the StartPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPoint

`func (o *CreateBranch) SetStartPoint(v string)`

SetStartPoint sets StartPoint field to given value.


### GetMessage

`func (o *CreateBranch) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateBranch) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateBranch) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateBranch) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


