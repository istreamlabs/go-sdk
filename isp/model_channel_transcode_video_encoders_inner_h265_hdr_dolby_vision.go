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

// checks if the ChannelTranscodeVideoEncodersInnerH265HdrDolbyVision type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelTranscodeVideoEncodersInnerH265HdrDolbyVision{}

// ChannelTranscodeVideoEncodersInnerH265HdrDolbyVision Only one of ['hlg', 'hdr10', 'dolby_vision'] may be set.
type ChannelTranscodeVideoEncodersInnerH265HdrDolbyVision struct {
	// Only one of ['profile5', 'profile81', 'profile84'] may be set.
	Profile5 *map[string]interface{} `json:"profile5,omitempty"`
	Profile81 *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81 `json:"profile81,omitempty"`
	// Only one of ['profile5', 'profile81', 'profile84'] may be set.
	Profile84 *map[string]interface{} `json:"profile84,omitempty"`
}

// NewChannelTranscodeVideoEncodersInnerH265HdrDolbyVision instantiates a new ChannelTranscodeVideoEncodersInnerH265HdrDolbyVision object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelTranscodeVideoEncodersInnerH265HdrDolbyVision() *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVision {
	this := ChannelTranscodeVideoEncodersInnerH265HdrDolbyVision{}
	return &this
}

// NewChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionWithDefaults instantiates a new ChannelTranscodeVideoEncodersInnerH265HdrDolbyVision object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionWithDefaults() *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVision {
	this := ChannelTranscodeVideoEncodersInnerH265HdrDolbyVision{}
	return &this
}

// GetProfile5 returns the Profile5 field value if set, zero value otherwise.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVision) GetProfile5() map[string]interface{} {
	if o == nil || IsNil(o.Profile5) {
		var ret map[string]interface{}
		return ret
	}
	return o.Profile5
}

// GetProfile5Ok returns a tuple with the Profile5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVision) GetProfile5Ok() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Profile5) {
		return map[string]interface{}{}, false
	}
	return o.Profile5, true
}

// HasProfile5 returns a boolean if a field has been set.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVision) HasProfile5() bool {
	if o != nil && !IsNil(o.Profile5) {
		return true
	}

	return false
}

// SetProfile5 gets a reference to the given map[string]interface{} and assigns it to the Profile5 field.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVision) SetProfile5(v map[string]interface{}) {
	o.Profile5 = v
}

// GetProfile81 returns the Profile81 field value if set, zero value otherwise.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVision) GetProfile81() ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81 {
	if o == nil || IsNil(o.Profile81) {
		var ret ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81
		return ret
	}
	return *o.Profile81
}

// GetProfile81Ok returns a tuple with the Profile81 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVision) GetProfile81Ok() (*ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81, bool) {
	if o == nil || IsNil(o.Profile81) {
		return nil, false
	}
	return o.Profile81, true
}

// HasProfile81 returns a boolean if a field has been set.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVision) HasProfile81() bool {
	if o != nil && !IsNil(o.Profile81) {
		return true
	}

	return false
}

// SetProfile81 gets a reference to the given ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81 and assigns it to the Profile81 field.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVision) SetProfile81(v ChannelTranscodeVideoEncodersInnerH265HdrDolbyVisionProfile81) {
	o.Profile81 = &v
}

// GetProfile84 returns the Profile84 field value if set, zero value otherwise.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVision) GetProfile84() map[string]interface{} {
	if o == nil || IsNil(o.Profile84) {
		var ret map[string]interface{}
		return ret
	}
	return o.Profile84
}

// GetProfile84Ok returns a tuple with the Profile84 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVision) GetProfile84Ok() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Profile84) {
		return map[string]interface{}{}, false
	}
	return o.Profile84, true
}

// HasProfile84 returns a boolean if a field has been set.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVision) HasProfile84() bool {
	if o != nil && !IsNil(o.Profile84) {
		return true
	}

	return false
}

// SetProfile84 gets a reference to the given map[string]interface{} and assigns it to the Profile84 field.
func (o *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVision) SetProfile84(v map[string]interface{}) {
	o.Profile84 = v
}

func (o ChannelTranscodeVideoEncodersInnerH265HdrDolbyVision) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelTranscodeVideoEncodersInnerH265HdrDolbyVision) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Profile5) {
		toSerialize["profile5"] = o.Profile5
	}
	if !IsNil(o.Profile81) {
		toSerialize["profile81"] = o.Profile81
	}
	if !IsNil(o.Profile84) {
		toSerialize["profile84"] = o.Profile84
	}
	return toSerialize, nil
}

type NullableChannelTranscodeVideoEncodersInnerH265HdrDolbyVision struct {
	value *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVision
	isSet bool
}

func (v NullableChannelTranscodeVideoEncodersInnerH265HdrDolbyVision) Get() *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVision {
	return v.value
}

func (v *NullableChannelTranscodeVideoEncodersInnerH265HdrDolbyVision) Set(val *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVision) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelTranscodeVideoEncodersInnerH265HdrDolbyVision) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelTranscodeVideoEncodersInnerH265HdrDolbyVision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelTranscodeVideoEncodersInnerH265HdrDolbyVision(val *ChannelTranscodeVideoEncodersInnerH265HdrDolbyVision) *NullableChannelTranscodeVideoEncodersInnerH265HdrDolbyVision {
	return &NullableChannelTranscodeVideoEncodersInnerH265HdrDolbyVision{value: val, isSet: true}
}

func (v NullableChannelTranscodeVideoEncodersInnerH265HdrDolbyVision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelTranscodeVideoEncodersInnerH265HdrDolbyVision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


