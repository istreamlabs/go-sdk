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

// ChannelPackagingPackagers struct for ChannelPackagingPackagers
type ChannelPackagingPackagers struct {
	ContentProtection *ChannelPackagingContentProtection `json:"content_protection,omitempty"`
	Mp2t *ChannelPackagingMp2t `json:"mp2t,omitempty"`
	Mp4 *ChannelPackagingMp4 `json:"mp4,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChannelPackagingPackagers ChannelPackagingPackagers

// NewChannelPackagingPackagers instantiates a new ChannelPackagingPackagers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelPackagingPackagers() *ChannelPackagingPackagers {
	this := ChannelPackagingPackagers{}
	return &this
}

// NewChannelPackagingPackagersWithDefaults instantiates a new ChannelPackagingPackagers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelPackagingPackagersWithDefaults() *ChannelPackagingPackagers {
	this := ChannelPackagingPackagers{}
	return &this
}

// GetContentProtection returns the ContentProtection field value if set, zero value otherwise.
func (o *ChannelPackagingPackagers) GetContentProtection() ChannelPackagingContentProtection {
	if o == nil || o.ContentProtection == nil {
		var ret ChannelPackagingContentProtection
		return ret
	}
	return *o.ContentProtection
}

// GetContentProtectionOk returns a tuple with the ContentProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPackagingPackagers) GetContentProtectionOk() (*ChannelPackagingContentProtection, bool) {
	if o == nil || o.ContentProtection == nil {
		return nil, false
	}
	return o.ContentProtection, true
}

// HasContentProtection returns a boolean if a field has been set.
func (o *ChannelPackagingPackagers) HasContentProtection() bool {
	if o != nil && o.ContentProtection != nil {
		return true
	}

	return false
}

// SetContentProtection gets a reference to the given ChannelPackagingContentProtection and assigns it to the ContentProtection field.
func (o *ChannelPackagingPackagers) SetContentProtection(v ChannelPackagingContentProtection) {
	o.ContentProtection = &v
}

// GetMp2t returns the Mp2t field value if set, zero value otherwise.
func (o *ChannelPackagingPackagers) GetMp2t() ChannelPackagingMp2t {
	if o == nil || o.Mp2t == nil {
		var ret ChannelPackagingMp2t
		return ret
	}
	return *o.Mp2t
}

// GetMp2tOk returns a tuple with the Mp2t field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPackagingPackagers) GetMp2tOk() (*ChannelPackagingMp2t, bool) {
	if o == nil || o.Mp2t == nil {
		return nil, false
	}
	return o.Mp2t, true
}

// HasMp2t returns a boolean if a field has been set.
func (o *ChannelPackagingPackagers) HasMp2t() bool {
	if o != nil && o.Mp2t != nil {
		return true
	}

	return false
}

// SetMp2t gets a reference to the given ChannelPackagingMp2t and assigns it to the Mp2t field.
func (o *ChannelPackagingPackagers) SetMp2t(v ChannelPackagingMp2t) {
	o.Mp2t = &v
}

// GetMp4 returns the Mp4 field value if set, zero value otherwise.
func (o *ChannelPackagingPackagers) GetMp4() ChannelPackagingMp4 {
	if o == nil || o.Mp4 == nil {
		var ret ChannelPackagingMp4
		return ret
	}
	return *o.Mp4
}

// GetMp4Ok returns a tuple with the Mp4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPackagingPackagers) GetMp4Ok() (*ChannelPackagingMp4, bool) {
	if o == nil || o.Mp4 == nil {
		return nil, false
	}
	return o.Mp4, true
}

// HasMp4 returns a boolean if a field has been set.
func (o *ChannelPackagingPackagers) HasMp4() bool {
	if o != nil && o.Mp4 != nil {
		return true
	}

	return false
}

// SetMp4 gets a reference to the given ChannelPackagingMp4 and assigns it to the Mp4 field.
func (o *ChannelPackagingPackagers) SetMp4(v ChannelPackagingMp4) {
	o.Mp4 = &v
}

func (o ChannelPackagingPackagers) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ChannelPackagingPackagers) UnmarshalJSON(bytes []byte) (err error) {
	varChannelPackagingPackagers := _ChannelPackagingPackagers{}

	if err = json.Unmarshal(bytes, &varChannelPackagingPackagers); err == nil {
		*o = ChannelPackagingPackagers(varChannelPackagingPackagers)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "content_protection")
		delete(additionalProperties, "mp2t")
		delete(additionalProperties, "mp4")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChannelPackagingPackagers struct {
	value *ChannelPackagingPackagers
	isSet bool
}

func (v NullableChannelPackagingPackagers) Get() *ChannelPackagingPackagers {
	return v.value
}

func (v *NullableChannelPackagingPackagers) Set(val *ChannelPackagingPackagers) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelPackagingPackagers) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelPackagingPackagers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelPackagingPackagers(val *ChannelPackagingPackagers) *NullableChannelPackagingPackagers {
	return &NullableChannelPackagingPackagers{value: val, isSet: true}
}

func (v NullableChannelPackagingPackagers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelPackagingPackagers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


