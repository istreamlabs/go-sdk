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

// ChannelTranscodeVideoEncodersInnerH265 Configure the encoder to use the H.265 codec. Only one of ['H264Settings', 'H265Settings'] may be set.
type ChannelTranscodeVideoEncodersInnerH265 struct {
	Hdr *ChannelTranscodeVideoEncodersInnerH265Hdr `json:"hdr,omitempty"`
	// H.265 video profile, which defines various encoder features and settings. See https://en.wikipedia.org/wiki/High_Efficiency_Video_Coding#Profiles for details.
	Profile *string `json:"profile,omitempty"`
}

// NewChannelTranscodeVideoEncodersInnerH265 instantiates a new ChannelTranscodeVideoEncodersInnerH265 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelTranscodeVideoEncodersInnerH265() *ChannelTranscodeVideoEncodersInnerH265 {
	this := ChannelTranscodeVideoEncodersInnerH265{}
	return &this
}

// NewChannelTranscodeVideoEncodersInnerH265WithDefaults instantiates a new ChannelTranscodeVideoEncodersInnerH265 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelTranscodeVideoEncodersInnerH265WithDefaults() *ChannelTranscodeVideoEncodersInnerH265 {
	this := ChannelTranscodeVideoEncodersInnerH265{}
	return &this
}

// GetHdr returns the Hdr field value if set, zero value otherwise.
func (o *ChannelTranscodeVideoEncodersInnerH265) GetHdr() ChannelTranscodeVideoEncodersInnerH265Hdr {
	if o == nil || o.Hdr == nil {
		var ret ChannelTranscodeVideoEncodersInnerH265Hdr
		return ret
	}
	return *o.Hdr
}

// GetHdrOk returns a tuple with the Hdr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscodeVideoEncodersInnerH265) GetHdrOk() (*ChannelTranscodeVideoEncodersInnerH265Hdr, bool) {
	if o == nil || o.Hdr == nil {
		return nil, false
	}
	return o.Hdr, true
}

// HasHdr returns a boolean if a field has been set.
func (o *ChannelTranscodeVideoEncodersInnerH265) HasHdr() bool {
	if o != nil && o.Hdr != nil {
		return true
	}

	return false
}

// SetHdr gets a reference to the given ChannelTranscodeVideoEncodersInnerH265Hdr and assigns it to the Hdr field.
func (o *ChannelTranscodeVideoEncodersInnerH265) SetHdr(v ChannelTranscodeVideoEncodersInnerH265Hdr) {
	o.Hdr = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *ChannelTranscodeVideoEncodersInnerH265) GetProfile() string {
	if o == nil || o.Profile == nil {
		var ret string
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscodeVideoEncodersInnerH265) GetProfileOk() (*string, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *ChannelTranscodeVideoEncodersInnerH265) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given string and assigns it to the Profile field.
func (o *ChannelTranscodeVideoEncodersInnerH265) SetProfile(v string) {
	o.Profile = &v
}

func (o ChannelTranscodeVideoEncodersInnerH265) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hdr != nil {
		toSerialize["hdr"] = o.Hdr
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}
	return json.Marshal(toSerialize)
}

type NullableChannelTranscodeVideoEncodersInnerH265 struct {
	value *ChannelTranscodeVideoEncodersInnerH265
	isSet bool
}

func (v NullableChannelTranscodeVideoEncodersInnerH265) Get() *ChannelTranscodeVideoEncodersInnerH265 {
	return v.value
}

func (v *NullableChannelTranscodeVideoEncodersInnerH265) Set(val *ChannelTranscodeVideoEncodersInnerH265) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelTranscodeVideoEncodersInnerH265) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelTranscodeVideoEncodersInnerH265) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelTranscodeVideoEncodersInnerH265(val *ChannelTranscodeVideoEncodersInnerH265) *NullableChannelTranscodeVideoEncodersInnerH265 {
	return &NullableChannelTranscodeVideoEncodersInnerH265{value: val, isSet: true}
}

func (v NullableChannelTranscodeVideoEncodersInnerH265) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelTranscodeVideoEncodersInnerH265) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

