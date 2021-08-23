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

// ChannelPackagingMp2T Uses MP2T format for each segments. Only one of ['mp2_t', 'mp4'] may be set.
type ChannelPackagingMp2T struct {
	// Forces the Video and Audio Encodings to be unmuxed when there is one audio encodings. This setting will have to be uniformed across MP2T packagers within a config. When there are two or more audio encodings, unmuxed will be used automatically.
	ForceUnmuxedAudio *bool `json:"force_unmuxed_audio,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChannelPackagingMp2T ChannelPackagingMp2T

// NewChannelPackagingMp2T instantiates a new ChannelPackagingMp2T object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelPackagingMp2T() *ChannelPackagingMp2T {
	this := ChannelPackagingMp2T{}
	return &this
}

// NewChannelPackagingMp2TWithDefaults instantiates a new ChannelPackagingMp2T object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelPackagingMp2TWithDefaults() *ChannelPackagingMp2T {
	this := ChannelPackagingMp2T{}
	return &this
}

// GetForceUnmuxedAudio returns the ForceUnmuxedAudio field value if set, zero value otherwise.
func (o *ChannelPackagingMp2T) GetForceUnmuxedAudio() bool {
	if o == nil || o.ForceUnmuxedAudio == nil {
		var ret bool
		return ret
	}
	return *o.ForceUnmuxedAudio
}

// GetForceUnmuxedAudioOk returns a tuple with the ForceUnmuxedAudio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPackagingMp2T) GetForceUnmuxedAudioOk() (*bool, bool) {
	if o == nil || o.ForceUnmuxedAudio == nil {
		return nil, false
	}
	return o.ForceUnmuxedAudio, true
}

// HasForceUnmuxedAudio returns a boolean if a field has been set.
func (o *ChannelPackagingMp2T) HasForceUnmuxedAudio() bool {
	if o != nil && o.ForceUnmuxedAudio != nil {
		return true
	}

	return false
}

// SetForceUnmuxedAudio gets a reference to the given bool and assigns it to the ForceUnmuxedAudio field.
func (o *ChannelPackagingMp2T) SetForceUnmuxedAudio(v bool) {
	o.ForceUnmuxedAudio = &v
}

func (o ChannelPackagingMp2T) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ForceUnmuxedAudio != nil {
		toSerialize["force_unmuxed_audio"] = o.ForceUnmuxedAudio
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ChannelPackagingMp2T) UnmarshalJSON(bytes []byte) (err error) {
	varChannelPackagingMp2T := _ChannelPackagingMp2T{}

	if err = json.Unmarshal(bytes, &varChannelPackagingMp2T); err == nil {
		*o = ChannelPackagingMp2T(varChannelPackagingMp2T)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "force_unmuxed_audio")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChannelPackagingMp2T struct {
	value *ChannelPackagingMp2T
	isSet bool
}

func (v NullableChannelPackagingMp2T) Get() *ChannelPackagingMp2T {
	return v.value
}

func (v *NullableChannelPackagingMp2T) Set(val *ChannelPackagingMp2T) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelPackagingMp2T) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelPackagingMp2T) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelPackagingMp2T(val *ChannelPackagingMp2T) *NullableChannelPackagingMp2T {
	return &NullableChannelPackagingMp2T{value: val, isSet: true}
}

func (v NullableChannelPackagingMp2T) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelPackagingMp2T) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


