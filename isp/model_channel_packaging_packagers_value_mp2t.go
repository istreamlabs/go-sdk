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

// ChannelPackagingPackagersValueMp2t Uses MP2T format for each segments. Only one of ['mp2t', 'mp4'] may be set.
type ChannelPackagingPackagersValueMp2t struct {
	// Forces the Video and Audio Encodings to be unmuxed when there is one audio encodings. This setting will have to be uniformed across MP2T packagers within a config. When there are two or more audio encodings, unmuxed will be used automatically.
	ForceUnmuxedAudio *bool `json:"force_unmuxed_audio,omitempty"`
	// If true, insert ID3 tags that include a UTC timestamp. This is a Turner/WM-specific extension.
	InsertId3UtcTime *bool `json:"insert_id3_utc_time,omitempty"`
}

// NewChannelPackagingPackagersValueMp2t instantiates a new ChannelPackagingPackagersValueMp2t object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelPackagingPackagersValueMp2t() *ChannelPackagingPackagersValueMp2t {
	this := ChannelPackagingPackagersValueMp2t{}
	return &this
}

// NewChannelPackagingPackagersValueMp2tWithDefaults instantiates a new ChannelPackagingPackagersValueMp2t object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelPackagingPackagersValueMp2tWithDefaults() *ChannelPackagingPackagersValueMp2t {
	this := ChannelPackagingPackagersValueMp2t{}
	return &this
}

// GetForceUnmuxedAudio returns the ForceUnmuxedAudio field value if set, zero value otherwise.
func (o *ChannelPackagingPackagersValueMp2t) GetForceUnmuxedAudio() bool {
	if o == nil || o.ForceUnmuxedAudio == nil {
		var ret bool
		return ret
	}
	return *o.ForceUnmuxedAudio
}

// GetForceUnmuxedAudioOk returns a tuple with the ForceUnmuxedAudio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPackagingPackagersValueMp2t) GetForceUnmuxedAudioOk() (*bool, bool) {
	if o == nil || o.ForceUnmuxedAudio == nil {
		return nil, false
	}
	return o.ForceUnmuxedAudio, true
}

// HasForceUnmuxedAudio returns a boolean if a field has been set.
func (o *ChannelPackagingPackagersValueMp2t) HasForceUnmuxedAudio() bool {
	if o != nil && o.ForceUnmuxedAudio != nil {
		return true
	}

	return false
}

// SetForceUnmuxedAudio gets a reference to the given bool and assigns it to the ForceUnmuxedAudio field.
func (o *ChannelPackagingPackagersValueMp2t) SetForceUnmuxedAudio(v bool) {
	o.ForceUnmuxedAudio = &v
}

// GetInsertId3UtcTime returns the InsertId3UtcTime field value if set, zero value otherwise.
func (o *ChannelPackagingPackagersValueMp2t) GetInsertId3UtcTime() bool {
	if o == nil || o.InsertId3UtcTime == nil {
		var ret bool
		return ret
	}
	return *o.InsertId3UtcTime
}

// GetInsertId3UtcTimeOk returns a tuple with the InsertId3UtcTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPackagingPackagersValueMp2t) GetInsertId3UtcTimeOk() (*bool, bool) {
	if o == nil || o.InsertId3UtcTime == nil {
		return nil, false
	}
	return o.InsertId3UtcTime, true
}

// HasInsertId3UtcTime returns a boolean if a field has been set.
func (o *ChannelPackagingPackagersValueMp2t) HasInsertId3UtcTime() bool {
	if o != nil && o.InsertId3UtcTime != nil {
		return true
	}

	return false
}

// SetInsertId3UtcTime gets a reference to the given bool and assigns it to the InsertId3UtcTime field.
func (o *ChannelPackagingPackagersValueMp2t) SetInsertId3UtcTime(v bool) {
	o.InsertId3UtcTime = &v
}

func (o ChannelPackagingPackagersValueMp2t) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ForceUnmuxedAudio != nil {
		toSerialize["force_unmuxed_audio"] = o.ForceUnmuxedAudio
	}
	if o.InsertId3UtcTime != nil {
		toSerialize["insert_id3_utc_time"] = o.InsertId3UtcTime
	}
	return json.Marshal(toSerialize)
}

type NullableChannelPackagingPackagersValueMp2t struct {
	value *ChannelPackagingPackagersValueMp2t
	isSet bool
}

func (v NullableChannelPackagingPackagersValueMp2t) Get() *ChannelPackagingPackagersValueMp2t {
	return v.value
}

func (v *NullableChannelPackagingPackagersValueMp2t) Set(val *ChannelPackagingPackagersValueMp2t) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelPackagingPackagersValueMp2t) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelPackagingPackagersValueMp2t) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelPackagingPackagersValueMp2t(val *ChannelPackagingPackagersValueMp2t) *NullableChannelPackagingPackagersValueMp2t {
	return &NullableChannelPackagingPackagersValueMp2t{value: val, isSet: true}
}

func (v NullableChannelPackagingPackagersValueMp2t) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelPackagingPackagersValueMp2t) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

