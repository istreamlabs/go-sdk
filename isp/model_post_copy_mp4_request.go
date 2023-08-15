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

// checks if the PostCopyMP4Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCopyMP4Request{}

// PostCopyMP4Request struct for PostCopyMP4Request
type PostCopyMP4Request struct {
	// An optional URL to a JSON Schema document describing this resource
	Schema *string `json:"$schema,omitempty"`
	// override path for writing an mp4 to destination.
	OverrideMp4Path *string `json:"override_mp4_path,omitempty"`
}

// NewPostCopyMP4Request instantiates a new PostCopyMP4Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCopyMP4Request() *PostCopyMP4Request {
	this := PostCopyMP4Request{}
	return &this
}

// NewPostCopyMP4RequestWithDefaults instantiates a new PostCopyMP4Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCopyMP4RequestWithDefaults() *PostCopyMP4Request {
	this := PostCopyMP4Request{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *PostCopyMP4Request) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCopyMP4Request) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *PostCopyMP4Request) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *PostCopyMP4Request) SetSchema(v string) {
	o.Schema = &v
}

// GetOverrideMp4Path returns the OverrideMp4Path field value if set, zero value otherwise.
func (o *PostCopyMP4Request) GetOverrideMp4Path() string {
	if o == nil || IsNil(o.OverrideMp4Path) {
		var ret string
		return ret
	}
	return *o.OverrideMp4Path
}

// GetOverrideMp4PathOk returns a tuple with the OverrideMp4Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCopyMP4Request) GetOverrideMp4PathOk() (*string, bool) {
	if o == nil || IsNil(o.OverrideMp4Path) {
		return nil, false
	}
	return o.OverrideMp4Path, true
}

// HasOverrideMp4Path returns a boolean if a field has been set.
func (o *PostCopyMP4Request) HasOverrideMp4Path() bool {
	if o != nil && !IsNil(o.OverrideMp4Path) {
		return true
	}

	return false
}

// SetOverrideMp4Path gets a reference to the given string and assigns it to the OverrideMp4Path field.
func (o *PostCopyMP4Request) SetOverrideMp4Path(v string) {
	o.OverrideMp4Path = &v
}

func (o PostCopyMP4Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCopyMP4Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	if !IsNil(o.OverrideMp4Path) {
		toSerialize["override_mp4_path"] = o.OverrideMp4Path
	}
	return toSerialize, nil
}

type NullablePostCopyMP4Request struct {
	value *PostCopyMP4Request
	isSet bool
}

func (v NullablePostCopyMP4Request) Get() *PostCopyMP4Request {
	return v.value
}

func (v *NullablePostCopyMP4Request) Set(val *PostCopyMP4Request) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCopyMP4Request) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCopyMP4Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCopyMP4Request(val *PostCopyMP4Request) *NullablePostCopyMP4Request {
	return &NullablePostCopyMP4Request{value: val, isSet: true}
}

func (v NullablePostCopyMP4Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCopyMP4Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


