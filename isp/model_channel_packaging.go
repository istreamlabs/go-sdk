/*
 * WBD Aventus Channels API
 *
 * API version: 0.0.0
 * Contact: live-control-plane-devs@wbd.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package isp

import (
	"encoding/json"
)

// checks if the ChannelPackaging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelPackaging{}

// ChannelPackaging Packaging configures media format and content protection settings.
type ChannelPackaging struct {
	// Defines the list of packagers available for Publications.
	Packagers *map[string]ChannelPackagingPackagersValue `json:"packagers,omitempty" doc:"Defines the list of packagers available for Publications."`
}

// NewChannelPackaging instantiates a new ChannelPackaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelPackaging() *ChannelPackaging {
	this := ChannelPackaging{}
	return &this
}

// NewChannelPackagingWithDefaults instantiates a new ChannelPackaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelPackagingWithDefaults() *ChannelPackaging {
	this := ChannelPackaging{}
	return &this
}

// GetPackagers returns the Packagers field value if set, zero value otherwise.
func (o *ChannelPackaging) GetPackagers() map[string]ChannelPackagingPackagersValue {
	if o == nil || IsNil(o.Packagers) {
		var ret map[string]ChannelPackagingPackagersValue
		return ret
	}
	return *o.Packagers
}

// GetPackagersOk returns a tuple with the Packagers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPackaging) GetPackagersOk() (*map[string]ChannelPackagingPackagersValue, bool) {
	if o == nil || IsNil(o.Packagers) {
		return nil, false
	}
	return o.Packagers, true
}

// HasPackagers returns a boolean if a field has been set.
func (o *ChannelPackaging) HasPackagers() bool {
	if o != nil && !IsNil(o.Packagers) {
		return true
	}

	return false
}

// SetPackagers gets a reference to the given map[string]ChannelPackagingPackagersValue and assigns it to the Packagers field.
func (o *ChannelPackaging) SetPackagers(v map[string]ChannelPackagingPackagersValue) {
	o.Packagers = &v
}

func (o ChannelPackaging) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelPackaging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Packagers) {
		toSerialize["packagers"] = o.Packagers
	}
	return toSerialize, nil
}

type NullableChannelPackaging struct {
	value *ChannelPackaging
	isSet bool
}

func (v NullableChannelPackaging) Get() *ChannelPackaging {
	return v.value
}

func (v *NullableChannelPackaging) Set(val *ChannelPackaging) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelPackaging) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelPackaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelPackaging(val *ChannelPackaging) *NullableChannelPackaging {
	return &NullableChannelPackaging{value: val, isSet: true}
}

func (v NullableChannelPackaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelPackaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


