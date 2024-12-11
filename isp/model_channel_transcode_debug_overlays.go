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

// checks if the ChannelTranscodeDebugOverlays type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelTranscodeDebugOverlays{}

// ChannelTranscodeDebugOverlays The debug_overlays specifies the position of the debugging overlay information from the transcoder into each video output. The overlay is burned into the video and will be visible to end-users if enabled. The default position value is OFF, which results in no debug overlay. Do not enable on customer facing channels. Requires a transcoder restart if the state is changed.
type ChannelTranscodeDebugOverlays struct {
	// Position of the debug overlay within the output frame.
	Position *string `json:"position,omitempty" enum:"OFF,TOP_LEFT,TOP_RIGHT,BOTTOM_LEFT,BOTTOM_RIGHT" doc:"Position of the debug overlay within the output frame."`
}

// NewChannelTranscodeDebugOverlays instantiates a new ChannelTranscodeDebugOverlays object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelTranscodeDebugOverlays() *ChannelTranscodeDebugOverlays {
	this := ChannelTranscodeDebugOverlays{}
	return &this
}

// NewChannelTranscodeDebugOverlaysWithDefaults instantiates a new ChannelTranscodeDebugOverlays object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelTranscodeDebugOverlaysWithDefaults() *ChannelTranscodeDebugOverlays {
	this := ChannelTranscodeDebugOverlays{}
	return &this
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ChannelTranscodeDebugOverlays) GetPosition() string {
	if o == nil || IsNil(o.Position) {
		var ret string
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscodeDebugOverlays) GetPositionOk() (*string, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ChannelTranscodeDebugOverlays) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given string and assigns it to the Position field.
func (o *ChannelTranscodeDebugOverlays) SetPosition(v string) {
	o.Position = &v
}

func (o ChannelTranscodeDebugOverlays) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelTranscodeDebugOverlays) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	return toSerialize, nil
}

type NullableChannelTranscodeDebugOverlays struct {
	value *ChannelTranscodeDebugOverlays
	isSet bool
}

func (v NullableChannelTranscodeDebugOverlays) Get() *ChannelTranscodeDebugOverlays {
	return v.value
}

func (v *NullableChannelTranscodeDebugOverlays) Set(val *ChannelTranscodeDebugOverlays) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelTranscodeDebugOverlays) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelTranscodeDebugOverlays) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelTranscodeDebugOverlays(val *ChannelTranscodeDebugOverlays) *NullableChannelTranscodeDebugOverlays {
	return &NullableChannelTranscodeDebugOverlays{value: val, isSet: true}
}

func (v NullableChannelTranscodeDebugOverlays) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelTranscodeDebugOverlays) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
