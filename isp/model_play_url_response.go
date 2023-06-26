/*
 * iStreamPlanet Channels API
 *
 * API version: 0.0.0
 * Contact: support@istreamplanet.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package isp

import (
	"encoding/json"
)

// checks if the PlayURLResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlayURLResponse{}

// PlayURLResponse struct for PlayURLResponse
type PlayURLResponse struct {
	// An optional URL to a JSON Schema document describing this resource
	Schema *string `json:"$schema,omitempty"`
	// URL for playback of a clip
	Url string `json:"url"`
}

// NewPlayURLResponse instantiates a new PlayURLResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlayURLResponse(url string) *PlayURLResponse {
	this := PlayURLResponse{}
	this.Url = url
	return &this
}

// NewPlayURLResponseWithDefaults instantiates a new PlayURLResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlayURLResponseWithDefaults() *PlayURLResponse {
	this := PlayURLResponse{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *PlayURLResponse) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayURLResponse) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *PlayURLResponse) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *PlayURLResponse) SetSchema(v string) {
	o.Schema = &v
}

// GetUrl returns the Url field value
func (o *PlayURLResponse) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *PlayURLResponse) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *PlayURLResponse) SetUrl(v string) {
	o.Url = v
}

func (o PlayURLResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlayURLResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

type NullablePlayURLResponse struct {
	value *PlayURLResponse
	isSet bool
}

func (v NullablePlayURLResponse) Get() *PlayURLResponse {
	return v.value
}

func (v *NullablePlayURLResponse) Set(val *PlayURLResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePlayURLResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePlayURLResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlayURLResponse(val *PlayURLResponse) *NullablePlayURLResponse {
	return &NullablePlayURLResponse{value: val, isSet: true}
}

func (v NullablePlayURLResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlayURLResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

