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

// ChannelPackagingPackagersValue struct for ChannelPackagingPackagersValue
type ChannelPackagingPackagersValue struct {
	ContentProtection *ChannelPackagingPackagersValueContentProtection `json:"content_protection,omitempty"`
	Mp2t *ChannelPackagingPackagersValueMp2t `json:"mp2t,omitempty"`
	Mp4 *ChannelPackagingPackagersValueMp4 `json:"mp4,omitempty"`
}

// NewChannelPackagingPackagersValue instantiates a new ChannelPackagingPackagersValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelPackagingPackagersValue() *ChannelPackagingPackagersValue {
	this := ChannelPackagingPackagersValue{}
	return &this
}

// NewChannelPackagingPackagersValueWithDefaults instantiates a new ChannelPackagingPackagersValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelPackagingPackagersValueWithDefaults() *ChannelPackagingPackagersValue {
	this := ChannelPackagingPackagersValue{}
	return &this
}

// GetContentProtection returns the ContentProtection field value if set, zero value otherwise.
func (o *ChannelPackagingPackagersValue) GetContentProtection() ChannelPackagingPackagersValueContentProtection {
	if o == nil || o.ContentProtection == nil {
		var ret ChannelPackagingPackagersValueContentProtection
		return ret
	}
	return *o.ContentProtection
}

// GetContentProtectionOk returns a tuple with the ContentProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPackagingPackagersValue) GetContentProtectionOk() (*ChannelPackagingPackagersValueContentProtection, bool) {
	if o == nil || o.ContentProtection == nil {
		return nil, false
	}
	return o.ContentProtection, true
}

// HasContentProtection returns a boolean if a field has been set.
func (o *ChannelPackagingPackagersValue) HasContentProtection() bool {
	if o != nil && o.ContentProtection != nil {
		return true
	}

	return false
}

// SetContentProtection gets a reference to the given ChannelPackagingPackagersValueContentProtection and assigns it to the ContentProtection field.
func (o *ChannelPackagingPackagersValue) SetContentProtection(v ChannelPackagingPackagersValueContentProtection) {
	o.ContentProtection = &v
}

// GetMp2t returns the Mp2t field value if set, zero value otherwise.
func (o *ChannelPackagingPackagersValue) GetMp2t() ChannelPackagingPackagersValueMp2t {
	if o == nil || o.Mp2t == nil {
		var ret ChannelPackagingPackagersValueMp2t
		return ret
	}
	return *o.Mp2t
}

// GetMp2tOk returns a tuple with the Mp2t field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPackagingPackagersValue) GetMp2tOk() (*ChannelPackagingPackagersValueMp2t, bool) {
	if o == nil || o.Mp2t == nil {
		return nil, false
	}
	return o.Mp2t, true
}

// HasMp2t returns a boolean if a field has been set.
func (o *ChannelPackagingPackagersValue) HasMp2t() bool {
	if o != nil && o.Mp2t != nil {
		return true
	}

	return false
}

// SetMp2t gets a reference to the given ChannelPackagingPackagersValueMp2t and assigns it to the Mp2t field.
func (o *ChannelPackagingPackagersValue) SetMp2t(v ChannelPackagingPackagersValueMp2t) {
	o.Mp2t = &v
}

// GetMp4 returns the Mp4 field value if set, zero value otherwise.
func (o *ChannelPackagingPackagersValue) GetMp4() ChannelPackagingPackagersValueMp4 {
	if o == nil || o.Mp4 == nil {
		var ret ChannelPackagingPackagersValueMp4
		return ret
	}
	return *o.Mp4
}

// GetMp4Ok returns a tuple with the Mp4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPackagingPackagersValue) GetMp4Ok() (*ChannelPackagingPackagersValueMp4, bool) {
	if o == nil || o.Mp4 == nil {
		return nil, false
	}
	return o.Mp4, true
}

// HasMp4 returns a boolean if a field has been set.
func (o *ChannelPackagingPackagersValue) HasMp4() bool {
	if o != nil && o.Mp4 != nil {
		return true
	}

	return false
}

// SetMp4 gets a reference to the given ChannelPackagingPackagersValueMp4 and assigns it to the Mp4 field.
func (o *ChannelPackagingPackagersValue) SetMp4(v ChannelPackagingPackagersValueMp4) {
	o.Mp4 = &v
}

func (o ChannelPackagingPackagersValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContentProtection != nil {
		toSerialize["content_protection"] = o.ContentProtection
	}
	if o.Mp2t != nil {
		toSerialize["mp2t"] = o.Mp2t
	}
	if o.Mp4 != nil {
		toSerialize["mp4"] = o.Mp4
	}
	return json.Marshal(toSerialize)
}

type NullableChannelPackagingPackagersValue struct {
	value *ChannelPackagingPackagersValue
	isSet bool
}

func (v NullableChannelPackagingPackagersValue) Get() *ChannelPackagingPackagersValue {
	return v.value
}

func (v *NullableChannelPackagingPackagersValue) Set(val *ChannelPackagingPackagersValue) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelPackagingPackagersValue) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelPackagingPackagersValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelPackagingPackagersValue(val *ChannelPackagingPackagersValue) *NullableChannelPackagingPackagersValue {
	return &NullableChannelPackagingPackagersValue{value: val, isSet: true}
}

func (v NullableChannelPackagingPackagersValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelPackagingPackagersValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


