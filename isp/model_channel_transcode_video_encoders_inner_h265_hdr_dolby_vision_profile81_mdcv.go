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

// ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv struct for ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv
type ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv struct {
	// Identifies color primaries and white point.
	ColorProfile *string `json:"color_profile,omitempty"`
	// Maximum display mastering luminance (nits).
	MaxDml *float64 `json:"max_dml,omitempty"`
	// Minimum display mastering luminance (nits).
	MinDml *float64 `json:"min_dml,omitempty"`
}

// NewChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv instantiates a new ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv() *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv {
	this := ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv{}
	return &this
}

// NewChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81MdcvWithDefaults instantiates a new ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81MdcvWithDefaults() *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv {
	this := ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv{}
	return &this
}

// GetColorProfile returns the ColorProfile field value if set, zero value otherwise.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv) GetColorProfile() string {
	if o == nil || o.ColorProfile == nil {
		var ret string
		return ret
	}
	return *o.ColorProfile
}

// GetColorProfileOk returns a tuple with the ColorProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv) GetColorProfileOk() (*string, bool) {
	if o == nil || o.ColorProfile == nil {
		return nil, false
	}
	return o.ColorProfile, true
}

// HasColorProfile returns a boolean if a field has been set.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv) HasColorProfile() bool {
	if o != nil && o.ColorProfile != nil {
		return true
	}

	return false
}

// SetColorProfile gets a reference to the given string and assigns it to the ColorProfile field.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv) SetColorProfile(v string) {
	o.ColorProfile = &v
}

// GetMaxDml returns the MaxDml field value if set, zero value otherwise.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv) GetMaxDml() float64 {
	if o == nil || o.MaxDml == nil {
		var ret float64
		return ret
	}
	return *o.MaxDml
}

// GetMaxDmlOk returns a tuple with the MaxDml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv) GetMaxDmlOk() (*float64, bool) {
	if o == nil || o.MaxDml == nil {
		return nil, false
	}
	return o.MaxDml, true
}

// HasMaxDml returns a boolean if a field has been set.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv) HasMaxDml() bool {
	if o != nil && o.MaxDml != nil {
		return true
	}

	return false
}

// SetMaxDml gets a reference to the given float64 and assigns it to the MaxDml field.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv) SetMaxDml(v float64) {
	o.MaxDml = &v
}

// GetMinDml returns the MinDml field value if set, zero value otherwise.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv) GetMinDml() float64 {
	if o == nil || o.MinDml == nil {
		var ret float64
		return ret
	}
	return *o.MinDml
}

// GetMinDmlOk returns a tuple with the MinDml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv) GetMinDmlOk() (*float64, bool) {
	if o == nil || o.MinDml == nil {
		return nil, false
	}
	return o.MinDml, true
}

// HasMinDml returns a boolean if a field has been set.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv) HasMinDml() bool {
	if o != nil && o.MinDml != nil {
		return true
	}

	return false
}

// SetMinDml gets a reference to the given float64 and assigns it to the MinDml field.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv) SetMinDml(v float64) {
	o.MinDml = &v
}

func (o ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ColorProfile != nil {
		toSerialize["color_profile"] = o.ColorProfile
	}
	if o.MaxDml != nil {
		toSerialize["max_dml"] = o.MaxDml
	}
	if o.MinDml != nil {
		toSerialize["min_dml"] = o.MinDml
	}
	return json.Marshal(toSerialize)
}

type NullableChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv struct {
	value *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv
	isSet bool
}

func (v NullableChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv) Get() *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv {
	return v.value
}

func (v *NullableChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv) Set(val *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv(val *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv) *NullableChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv {
	return &NullableChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv{value: val, isSet: true}
}

func (v NullableChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Mdcv) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

