/*
 * iStreamPlanet Slate Management API
 *
 * API version: 1.0.0
 * Contact: support@istreamplanet.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package isp

import (
	"encoding/json"
)

// checks if the Slate2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Slate2{}

// Slate2 struct for Slate2
type Slate2 struct {
	// A friendly slate description.
	Description string `json:"description" doc:"A friendly slate description."`
	// Unique identifier for this slate
	Id string `json:"id" doc:"Unique identifier for this slate"`
	// The url where the slate can be accessed.
	Url string `json:"url" doc:"The url where the slate can be accessed."`
}

// NewSlate2 instantiates a new Slate2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlate2(description string, id string, url string) *Slate2 {
	this := Slate2{}
	this.Description = description
	this.Id = id
	this.Url = url
	return &this
}

// NewSlate2WithDefaults instantiates a new Slate2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlate2WithDefaults() *Slate2 {
	this := Slate2{}
	return &this
}

// GetDescription returns the Description field value
func (o *Slate2) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Slate2) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Slate2) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value
func (o *Slate2) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Slate2) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Slate2) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *Slate2) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Slate2) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Slate2) SetUrl(v string) {
	o.Url = v
}

func (o Slate2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Slate2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

type NullableSlate2 struct {
	value *Slate2
	isSet bool
}

func (v NullableSlate2) Get() *Slate2 {
	return v.value
}

func (v *NullableSlate2) Set(val *Slate2) {
	v.value = val
	v.isSet = true
}

func (v NullableSlate2) IsSet() bool {
	return v.isSet
}

func (v *NullableSlate2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlate2(val *Slate2) *NullableSlate2 {
	return &NullableSlate2{value: val, isSet: true}
}

func (v NullableSlate2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlate2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


