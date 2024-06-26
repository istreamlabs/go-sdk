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

// checks if the ChannelTranscodeVideoEncodersInnerH265HdrHdr10 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelTranscodeVideoEncodersInnerH265HdrHdr10{}

// ChannelTranscodeVideoEncodersInnerH265HdrHdr10 Only one of ['hdr10', 'dolby_vision'] may be set.
type ChannelTranscodeVideoEncodersInnerH265HdrHdr10 struct {
	Clli *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli `json:"clli,omitempty"`
	Mdcv *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv `json:"mdcv,omitempty"`
}

// NewChannelTranscodeVideoEncodersInnerH265HdrHdr10 instantiates a new ChannelTranscodeVideoEncodersInnerH265HdrHdr10 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelTranscodeVideoEncodersInnerH265HdrHdr10() *ChannelTranscodeVideoEncodersInnerH265HdrHdr10 {
	this := ChannelTranscodeVideoEncodersInnerH265HdrHdr10{}
	return &this
}

// NewChannelTranscodeVideoEncodersInnerH265HdrHdr10WithDefaults instantiates a new ChannelTranscodeVideoEncodersInnerH265HdrHdr10 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelTranscodeVideoEncodersInnerH265HdrHdr10WithDefaults() *ChannelTranscodeVideoEncodersInnerH265HdrHdr10 {
	this := ChannelTranscodeVideoEncodersInnerH265HdrHdr10{}
	return &this
}

// GetClli returns the Clli field value if set, zero value otherwise.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrHdr10) GetClli() ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli {
	if o == nil || IsNil(o.Clli) {
		var ret ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli
		return ret
	}
	return *o.Clli
}

// GetClliOk returns a tuple with the Clli field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrHdr10) GetClliOk() (*ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli, bool) {
	if o == nil || IsNil(o.Clli) {
		return nil, false
	}
	return o.Clli, true
}

// HasClli returns a boolean if a field has been set.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrHdr10) HasClli() bool {
	if o != nil && !IsNil(o.Clli) {
		return true
	}

	return false
}

// SetClli gets a reference to the given ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli and assigns it to the Clli field.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrHdr10) SetClli(v ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli) {
	o.Clli = &v
}

// GetMdcv returns the Mdcv field value if set, zero value otherwise.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrHdr10) GetMdcv() ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv {
	if o == nil || IsNil(o.Mdcv) {
		var ret ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv
		return ret
	}
	return *o.Mdcv
}

// GetMdcvOk returns a tuple with the Mdcv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrHdr10) GetMdcvOk() (*ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv, bool) {
	if o == nil || IsNil(o.Mdcv) {
		return nil, false
	}
	return o.Mdcv, true
}

// HasMdcv returns a boolean if a field has been set.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrHdr10) HasMdcv() bool {
	if o != nil && !IsNil(o.Mdcv) {
		return true
	}

	return false
}

// SetMdcv gets a reference to the given ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv and assigns it to the Mdcv field.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrHdr10) SetMdcv(v ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv) {
	o.Mdcv = &v
}

func (o ChannelTranscodeVideoEncodersInnerH265HdrHdr10) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelTranscodeVideoEncodersInnerH265HdrHdr10) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Clli) {
		toSerialize["clli"] = o.Clli
	}
	if !IsNil(o.Mdcv) {
		toSerialize["mdcv"] = o.Mdcv
	}
	return toSerialize, nil
}

type NullableChannelTranscodeVideoEncodersInnerH265HdrHdr10 struct {
	value *ChannelTranscodeVideoEncodersInnerH265HdrHdr10
	isSet bool
}

func (v NullableChannelTranscodeVideoEncodersInnerH265HdrHdr10) Get() *ChannelTranscodeVideoEncodersInnerH265HdrHdr10 {
	return v.value
}

func (v *NullableChannelTranscodeVideoEncodersInnerH265HdrHdr10) Set(val *ChannelTranscodeVideoEncodersInnerH265HdrHdr10) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelTranscodeVideoEncodersInnerH265HdrHdr10) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelTranscodeVideoEncodersInnerH265HdrHdr10) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelTranscodeVideoEncodersInnerH265HdrHdr10(val *ChannelTranscodeVideoEncodersInnerH265HdrHdr10) *NullableChannelTranscodeVideoEncodersInnerH265HdrHdr10 {
	return &NullableChannelTranscodeVideoEncodersInnerH265HdrHdr10{value: val, isSet: true}
}

func (v NullableChannelTranscodeVideoEncodersInnerH265HdrHdr10) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelTranscodeVideoEncodersInnerH265HdrHdr10) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


