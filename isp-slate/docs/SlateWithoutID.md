# SlateWithoutID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**Description** | **string** | A friendly slate description. | 
**Url** | **string** | The url where the slate can be accessed. Must be a valid URL | 

## Methods

### NewSlateWithoutID

`func NewSlateWithoutID(description string, url string, ) *SlateWithoutID`

NewSlateWithoutID instantiates a new SlateWithoutID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlateWithoutIDWithDefaults

`func NewSlateWithoutIDWithDefaults() *SlateWithoutID`

NewSlateWithoutIDWithDefaults instantiates a new SlateWithoutID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *SlateWithoutID) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *SlateWithoutID) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *SlateWithoutID) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *SlateWithoutID) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetDescription

`func (o *SlateWithoutID) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SlateWithoutID) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SlateWithoutID) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetUrl

`func (o *SlateWithoutID) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SlateWithoutID) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SlateWithoutID) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


