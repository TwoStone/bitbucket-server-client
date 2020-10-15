# ProjectsPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | Pointer to **int32** |  | 
**Limit** | Pointer to **int32** |  | 
**Start** | Pointer to **int32** |  | 
**IsLastPage** | Pointer to **bool** |  | 
**NextPageStart** | Pointer to **int32** |  | [optional] 
**Values** | Pointer to [**[]Project**](project.md) |  | 

## Methods

### NewProjectsPage

`func NewProjectsPage(size int32, limit int32, start int32, isLastPage bool, values []Project, ) *ProjectsPage`

NewProjectsPage instantiates a new ProjectsPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsPageWithDefaults

`func NewProjectsPageWithDefaults() *ProjectsPage`

NewProjectsPageWithDefaults instantiates a new ProjectsPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *ProjectsPage) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ProjectsPage) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ProjectsPage) SetSize(v int32)`

SetSize sets Size field to given value.


### GetLimit

`func (o *ProjectsPage) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ProjectsPage) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ProjectsPage) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetStart

`func (o *ProjectsPage) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ProjectsPage) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ProjectsPage) SetStart(v int32)`

SetStart sets Start field to given value.


### GetIsLastPage

`func (o *ProjectsPage) GetIsLastPage() bool`

GetIsLastPage returns the IsLastPage field if non-nil, zero value otherwise.

### GetIsLastPageOk

`func (o *ProjectsPage) GetIsLastPageOk() (*bool, bool)`

GetIsLastPageOk returns a tuple with the IsLastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLastPage

`func (o *ProjectsPage) SetIsLastPage(v bool)`

SetIsLastPage sets IsLastPage field to given value.


### GetNextPageStart

`func (o *ProjectsPage) GetNextPageStart() int32`

GetNextPageStart returns the NextPageStart field if non-nil, zero value otherwise.

### GetNextPageStartOk

`func (o *ProjectsPage) GetNextPageStartOk() (*int32, bool)`

GetNextPageStartOk returns a tuple with the NextPageStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageStart

`func (o *ProjectsPage) SetNextPageStart(v int32)`

SetNextPageStart sets NextPageStart field to given value.

### HasNextPageStart

`func (o *ProjectsPage) HasNextPageStart() bool`

HasNextPageStart returns a boolean if a field has been set.

### GetValues

`func (o *ProjectsPage) GetValues() []Project`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ProjectsPage) GetValuesOk() (*[]Project, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ProjectsPage) SetValues(v []Project)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


