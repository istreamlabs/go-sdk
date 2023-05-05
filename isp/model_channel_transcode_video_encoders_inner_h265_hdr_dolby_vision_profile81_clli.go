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

// checks if the ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli{}

// ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli struct for ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli
type ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli struct {
	// Maximum Content Light Level (nits).
	MaxCll *float64 `json:"max_cll,omitempty"`
	// Maximum Frame-Average Light Level (nits).
	MaxFall *float64 `json:"max_fall,omitempty"`
}

// NewChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli instantiates a new ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli() *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli {
	this := ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli{}
	return &this
}

// NewChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81ClliWithDefaults instantiates a new ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81ClliWithDefaults() *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli {
	this := ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli{}
	return &this
}

// GetMaxCll returns the MaxCll field value if set, zero value otherwise.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli) GetMaxCll() float64 {
	if o == nil || IsNil(o.MaxCll) {
		var ret float64
		return ret
	}
	return *o.MaxCll
}

// GetMaxCllOk returns a tuple with the MaxCll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli) GetMaxCllOk() (*float64, bool) {
	if o == nil || IsNil(o.MaxCll) {
		return nil, false
	}
	return o.MaxCll, true
}

// HasMaxCll returns a boolean if a field has been set.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli) HasMaxCll() bool {
	if o != nil && !IsNil(o.MaxCll) {
		return true
	}

	return false
}

// SetMaxCll gets a reference to the given float64 and assigns it to the MaxCll field.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli) SetMaxCll(v float64) {
	o.MaxCll = &v
}

// GetMaxFall returns the MaxFall field value if set, zero value otherwise.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli) GetMaxFall() float64 {
	if o == nil || IsNil(o.MaxFall) {
		var ret float64
		return ret
	}
	return *o.MaxFall
}

// GetMaxFallOk returns a tuple with the MaxFall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli) GetMaxFallOk() (*float64, bool) {
	if o == nil || IsNil(o.MaxFall) {
		return nil, false
	}
	return o.MaxFall, true
}

// HasMaxFall returns a boolean if a field has been set.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli) HasMaxFall() bool {
	if o != nil && !IsNil(o.MaxFall) {
		return true
	}

	return false
}

// SetMaxFall gets a reference to the given float64 and assigns it to the MaxFall field.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli) SetMaxFall(v float64) {
	o.MaxFall = &v
}

func (o ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxCll) {
		toSerialize["max_cll"] = o.MaxCll
	}
	if !IsNil(o.MaxFall) {
		toSerialize["max_fall"] = o.MaxFall
	}
	return toSerialize, nil
}

type NullableChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli struct {
	value *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli
	isSet bool
}

func (v NullableChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli) Get() *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli {
	return v.value
}

func (v *NullableChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli) Set(val *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli(val *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli) *NullableChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli {
	return &NullableChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli{value: val, isSet: true}
}

func (v NullableChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81Clli) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


