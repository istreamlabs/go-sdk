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

// checks if the ChannelTranscodeSegmenter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelTranscodeSegmenter{}

// ChannelTranscodeSegmenter Segmenter configures how video GOPs and segments get generated.
type ChannelTranscodeSegmenter struct {
	// GOP (group of pictures) duration specifies the amount of time between I-frames. Shorter durations can lower quality slightly as each I-frame uses more bits than P- & B-frames but can provide a better seeking experience when enabling thumbnail encoders and/or I-Frame Only playlists.
	GopDurationSecs *float64 `json:"gop_duration_secs,omitempty"`
	// Not public because we haven't shipped low latency HLS yet and we probably need to update the naming.
	PartialsMode *string `json:"partials_mode,omitempty"`
	// Segment duration specifies the target duration of a single segment. Segments shorter than this duration can occur at signaling boundaries. This value _must_ be a multiple of the GOP duration value.
	SegmentDurationSecs *float64 `json:"segment_duration_secs,omitempty"`
	// Include TEMI (Timeline and External Media Information ISO/IEC 13818-1:2019 Annex U) to mpeg-ts segments.
	Temi *bool `json:"temi,omitempty"`
}

// NewChannelTranscodeSegmenter instantiates a new ChannelTranscodeSegmenter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelTranscodeSegmenter() *ChannelTranscodeSegmenter {
	this := ChannelTranscodeSegmenter{}
	return &this
}

// NewChannelTranscodeSegmenterWithDefaults instantiates a new ChannelTranscodeSegmenter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelTranscodeSegmenterWithDefaults() *ChannelTranscodeSegmenter {
	this := ChannelTranscodeSegmenter{}
	return &this
}

// GetGopDurationSecs returns the GopDurationSecs field value if set, zero value otherwise.
func (o *ChannelTranscodeSegmenter) GetGopDurationSecs() float64 {
	if o == nil || IsNil(o.GopDurationSecs) {
		var ret float64
		return ret
	}
	return *o.GopDurationSecs
}

// GetGopDurationSecsOk returns a tuple with the GopDurationSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscodeSegmenter) GetGopDurationSecsOk() (*float64, bool) {
	if o == nil || IsNil(o.GopDurationSecs) {
		return nil, false
	}
	return o.GopDurationSecs, true
}

// HasGopDurationSecs returns a boolean if a field has been set.
func (o *ChannelTranscodeSegmenter) HasGopDurationSecs() bool {
	if o != nil && !IsNil(o.GopDurationSecs) {
		return true
	}

	return false
}

// SetGopDurationSecs gets a reference to the given float64 and assigns it to the GopDurationSecs field.
func (o *ChannelTranscodeSegmenter) SetGopDurationSecs(v float64) {
	o.GopDurationSecs = &v
}

// GetPartialsMode returns the PartialsMode field value if set, zero value otherwise.
func (o *ChannelTranscodeSegmenter) GetPartialsMode() string {
	if o == nil || IsNil(o.PartialsMode) {
		var ret string
		return ret
	}
	return *o.PartialsMode
}

// GetPartialsModeOk returns a tuple with the PartialsMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscodeSegmenter) GetPartialsModeOk() (*string, bool) {
	if o == nil || IsNil(o.PartialsMode) {
		return nil, false
	}
	return o.PartialsMode, true
}

// HasPartialsMode returns a boolean if a field has been set.
func (o *ChannelTranscodeSegmenter) HasPartialsMode() bool {
	if o != nil && !IsNil(o.PartialsMode) {
		return true
	}

	return false
}

// SetPartialsMode gets a reference to the given string and assigns it to the PartialsMode field.
func (o *ChannelTranscodeSegmenter) SetPartialsMode(v string) {
	o.PartialsMode = &v
}

// GetSegmentDurationSecs returns the SegmentDurationSecs field value if set, zero value otherwise.
func (o *ChannelTranscodeSegmenter) GetSegmentDurationSecs() float64 {
	if o == nil || IsNil(o.SegmentDurationSecs) {
		var ret float64
		return ret
	}
	return *o.SegmentDurationSecs
}

// GetSegmentDurationSecsOk returns a tuple with the SegmentDurationSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscodeSegmenter) GetSegmentDurationSecsOk() (*float64, bool) {
	if o == nil || IsNil(o.SegmentDurationSecs) {
		return nil, false
	}
	return o.SegmentDurationSecs, true
}

// HasSegmentDurationSecs returns a boolean if a field has been set.
func (o *ChannelTranscodeSegmenter) HasSegmentDurationSecs() bool {
	if o != nil && !IsNil(o.SegmentDurationSecs) {
		return true
	}

	return false
}

// SetSegmentDurationSecs gets a reference to the given float64 and assigns it to the SegmentDurationSecs field.
func (o *ChannelTranscodeSegmenter) SetSegmentDurationSecs(v float64) {
	o.SegmentDurationSecs = &v
}

// GetTemi returns the Temi field value if set, zero value otherwise.
func (o *ChannelTranscodeSegmenter) GetTemi() bool {
	if o == nil || IsNil(o.Temi) {
		var ret bool
		return ret
	}
	return *o.Temi
}

// GetTemiOk returns a tuple with the Temi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscodeSegmenter) GetTemiOk() (*bool, bool) {
	if o == nil || IsNil(o.Temi) {
		return nil, false
	}
	return o.Temi, true
}

// HasTemi returns a boolean if a field has been set.
func (o *ChannelTranscodeSegmenter) HasTemi() bool {
	if o != nil && !IsNil(o.Temi) {
		return true
	}

	return false
}

// SetTemi gets a reference to the given bool and assigns it to the Temi field.
func (o *ChannelTranscodeSegmenter) SetTemi(v bool) {
	o.Temi = &v
}

func (o ChannelTranscodeSegmenter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelTranscodeSegmenter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GopDurationSecs) {
		toSerialize["gop_duration_secs"] = o.GopDurationSecs
	}
	if !IsNil(o.PartialsMode) {
		toSerialize["partials_mode"] = o.PartialsMode
	}
	if !IsNil(o.SegmentDurationSecs) {
		toSerialize["segment_duration_secs"] = o.SegmentDurationSecs
	}
	if !IsNil(o.Temi) {
		toSerialize["temi"] = o.Temi
	}
	return toSerialize, nil
}

type NullableChannelTranscodeSegmenter struct {
	value *ChannelTranscodeSegmenter
	isSet bool
}

func (v NullableChannelTranscodeSegmenter) Get() *ChannelTranscodeSegmenter {
	return v.value
}

func (v *NullableChannelTranscodeSegmenter) Set(val *ChannelTranscodeSegmenter) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelTranscodeSegmenter) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelTranscodeSegmenter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelTranscodeSegmenter(val *ChannelTranscodeSegmenter) *NullableChannelTranscodeSegmenter {
	return &NullableChannelTranscodeSegmenter{value: val, isSet: true}
}

func (v NullableChannelTranscodeSegmenter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelTranscodeSegmenter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


