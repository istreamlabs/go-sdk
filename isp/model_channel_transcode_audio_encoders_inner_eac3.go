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

// checks if the ChannelTranscodeAudioEncodersInnerEac3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelTranscodeAudioEncodersInnerEac3{}

// ChannelTranscodeAudioEncodersInnerEac3 Only one of ['eac3'] may be set.
type ChannelTranscodeAudioEncodersInnerEac3 struct {
	// Indicates that the source will contain Joint Object Coding metadata (Dolby Atmos) and that the encoder should operate in passthrough mode. https://learning.dolby.com/hc/en-us/articles/4406039180564-Appendix-C-Dolby-Atmos-Delivery-Codecs- When the audio encoder is configured with this setting, the source audio will be repackaged without being decoded and re-encoded. If the source is not DD+JOC, it will be replaced with silence.
	DdpJocPassthrough *bool `json:"ddp_joc_passthrough,omitempty" doc:"Indicates that the source will contain Joint Object Coding metadata (Dolby Atmos) and that the encoder should operate in passthrough mode. https://learning.dolby.com/hc/en-us/articles/4406039180564-Appendix-C-Dolby-Atmos-Delivery-Codecs- When the audio encoder is configured with this setting, the source audio will be repackaged without being decoded and re-encoded. If the source is not DD+JOC, it will be replaced with silence."`
}

// NewChannelTranscodeAudioEncodersInnerEac3 instantiates a new ChannelTranscodeAudioEncodersInnerEac3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelTranscodeAudioEncodersInnerEac3() *ChannelTranscodeAudioEncodersInnerEac3 {
	this := ChannelTranscodeAudioEncodersInnerEac3{}
	return &this
}

// NewChannelTranscodeAudioEncodersInnerEac3WithDefaults instantiates a new ChannelTranscodeAudioEncodersInnerEac3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelTranscodeAudioEncodersInnerEac3WithDefaults() *ChannelTranscodeAudioEncodersInnerEac3 {
	this := ChannelTranscodeAudioEncodersInnerEac3{}
	return &this
}

// GetDdpJocPassthrough returns the DdpJocPassthrough field value if set, zero value otherwise.
func (o *ChannelTranscodeAudioEncodersInnerEac3) GetDdpJocPassthrough() bool {
	if o == nil || IsNil(o.DdpJocPassthrough) {
		var ret bool
		return ret
	}
	return *o.DdpJocPassthrough
}

// GetDdpJocPassthroughOk returns a tuple with the DdpJocPassthrough field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscodeAudioEncodersInnerEac3) GetDdpJocPassthroughOk() (*bool, bool) {
	if o == nil || IsNil(o.DdpJocPassthrough) {
		return nil, false
	}
	return o.DdpJocPassthrough, true
}

// HasDdpJocPassthrough returns a boolean if a field has been set.
func (o *ChannelTranscodeAudioEncodersInnerEac3) HasDdpJocPassthrough() bool {
	if o != nil && !IsNil(o.DdpJocPassthrough) {
		return true
	}

	return false
}

// SetDdpJocPassthrough gets a reference to the given bool and assigns it to the DdpJocPassthrough field.
func (o *ChannelTranscodeAudioEncodersInnerEac3) SetDdpJocPassthrough(v bool) {
	o.DdpJocPassthrough = &v
}

func (o ChannelTranscodeAudioEncodersInnerEac3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelTranscodeAudioEncodersInnerEac3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DdpJocPassthrough) {
		toSerialize["ddp_joc_passthrough"] = o.DdpJocPassthrough
	}
	return toSerialize, nil
}

type NullableChannelTranscodeAudioEncodersInnerEac3 struct {
	value *ChannelTranscodeAudioEncodersInnerEac3
	isSet bool
}

func (v NullableChannelTranscodeAudioEncodersInnerEac3) Get() *ChannelTranscodeAudioEncodersInnerEac3 {
	return v.value
}

func (v *NullableChannelTranscodeAudioEncodersInnerEac3) Set(val *ChannelTranscodeAudioEncodersInnerEac3) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelTranscodeAudioEncodersInnerEac3) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelTranscodeAudioEncodersInnerEac3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelTranscodeAudioEncodersInnerEac3(val *ChannelTranscodeAudioEncodersInnerEac3) *NullableChannelTranscodeAudioEncodersInnerEac3 {
	return &NullableChannelTranscodeAudioEncodersInnerEac3{value: val, isSet: true}
}

func (v NullableChannelTranscodeAudioEncodersInnerEac3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelTranscodeAudioEncodersInnerEac3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


