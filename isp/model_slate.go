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

// Slate struct for Slate
type Slate struct {
	// Slate duration (ms)
	Duration *int32 `json:"duration,omitempty"`
	// Slate url
	Uri string `json:"uri"`
}

// NewSlate instantiates a new Slate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlate(uri string) *Slate {
	this := Slate{}
	var duration int32 = 0
	this.Duration = &duration
	this.Uri = uri
	return &this
}

// NewSlateWithDefaults instantiates a new Slate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlateWithDefaults() *Slate {
	this := Slate{}
	var duration int32 = 0
	this.Duration = &duration
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *Slate) GetDuration() int32 {
	if o == nil || o.Duration == nil {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Slate) GetDurationOk() (*int32, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *Slate) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *Slate) SetDuration(v int32) {
	o.Duration = &v
}

// GetUri returns the Uri field value
func (o *Slate) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *Slate) GetUriOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *Slate) SetUri(v string) {
	o.Uri = v
}

func (o Slate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if true {
		toSerialize["uri"] = o.Uri
	}
	return json.Marshal(toSerialize)
}

type NullableSlate struct {
	value *Slate
	isSet bool
}

func (v NullableSlate) Get() *Slate {
	return v.value
}

func (v *NullableSlate) Set(val *Slate) {
	v.value = val
	v.isSet = true
}

func (v NullableSlate) IsSet() bool {
	return v.isSet
}

func (v *NullableSlate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlate(val *Slate) *NullableSlate {
	return &NullableSlate{value: val, isSet: true}
}

func (v NullableSlate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


