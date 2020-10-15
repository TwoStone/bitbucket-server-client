# Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | 
**CreateDate** | Pointer to **int64** |  | [optional] 
**UpdatedDate** | Pointer to **int64** |  | [optional] 
**Events** | Pointer to [**[]WebhookEvent**](webhookEvent.md) |  | 
**Configuration** | Pointer to [**WebhookConfiguration**](webhook_configuration.md) |  | [optional] 
**Url** | Pointer to **string** |  | 

## Methods

### NewWebhook

`func NewWebhook(name string, events []WebhookEvent, url string, ) *Webhook`

NewWebhook instantiates a new Webhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookWithDefaults

`func NewWebhookWithDefaults() *Webhook`

NewWebhookWithDefaults instantiates a new Webhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Webhook) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Webhook) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Webhook) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Webhook) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Webhook) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Webhook) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Webhook) SetName(v string)`

SetName sets Name field to given value.


### GetCreateDate

`func (o *Webhook) GetCreateDate() int64`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *Webhook) GetCreateDateOk() (*int64, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *Webhook) SetCreateDate(v int64)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *Webhook) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetUpdatedDate

`func (o *Webhook) GetUpdatedDate() int64`

GetUpdatedDate returns the UpdatedDate field if non-nil, zero value otherwise.

### GetUpdatedDateOk

`func (o *Webhook) GetUpdatedDateOk() (*int64, bool)`

GetUpdatedDateOk returns a tuple with the UpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDate

`func (o *Webhook) SetUpdatedDate(v int64)`

SetUpdatedDate sets UpdatedDate field to given value.

### HasUpdatedDate

`func (o *Webhook) HasUpdatedDate() bool`

HasUpdatedDate returns a boolean if a field has been set.

### GetEvents

`func (o *Webhook) GetEvents() []WebhookEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Webhook) GetEventsOk() (*[]WebhookEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Webhook) SetEvents(v []WebhookEvent)`

SetEvents sets Events field to given value.


### GetConfiguration

`func (o *Webhook) GetConfiguration() WebhookConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *Webhook) GetConfigurationOk() (*WebhookConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *Webhook) SetConfiguration(v WebhookConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *Webhook) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetUrl

`func (o *Webhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Webhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Webhook) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


