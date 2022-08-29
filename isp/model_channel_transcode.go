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

// ChannelTranscode Transcode configures audio/video conversion settings.
type ChannelTranscode struct {
	// Audio encoders specify audio conversion settings, e.g. channels, samples, codec, bitrate, etc.
	AudioEncoders []ChannelTranscodeAudioEncodersInner `json:"audio_encoders,omitempty"`
	// Feature flag strings enable experimental transcode features or functionality that are not yet or never will be promoted to the channeldoc model proper.
	FeatureFlags []string `json:"feature_flags,omitempty"`
	// Specify how to process ID3 tags from the input source. If not specified, ID3 tags in the source will be ignored.
	Id3Mode *string `json:"id3_mode,omitempty"`
	// Resize mode specifies how to scale a video up or down to match the output dimensions.
	ResizeMode *string `json:"resize_mode,omitempty"`
	Segmenter *ChannelTranscodeSegmenter `json:"segmenter,omitempty"`
	// Thumbnail encoders specify how to create image snapshots of the video stream.
	ThumbnailEncoders []ChannelTranscodeThumbnailEncodersInner `json:"thumbnail_encoders,omitempty"`
	// Video encoders specify video conversion settings, e.g. dimensions, codec, bitrate, etc.
	VideoEncoders []ChannelTranscodeVideoEncodersInner `json:"video_encoders,omitempty"`
}

// NewChannelTranscode instantiates a new ChannelTranscode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelTranscode() *ChannelTranscode {
	this := ChannelTranscode{}
	return &this
}

// NewChannelTranscodeWithDefaults instantiates a new ChannelTranscode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelTranscodeWithDefaults() *ChannelTranscode {
	this := ChannelTranscode{}
	return &this
}

// GetAudioEncoders returns the AudioEncoders field value if set, zero value otherwise.
func (o *ChannelTranscode) GetAudioEncoders() []ChannelTranscodeAudioEncodersInner {
	if o == nil || o.AudioEncoders == nil {
		var ret []ChannelTranscodeAudioEncodersInner
		return ret
	}
	return o.AudioEncoders
}

// GetAudioEncodersOk returns a tuple with the AudioEncoders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscode) GetAudioEncodersOk() ([]ChannelTranscodeAudioEncodersInner, bool) {
	if o == nil || o.AudioEncoders == nil {
		return nil, false
	}
	return o.AudioEncoders, true
}

// HasAudioEncoders returns a boolean if a field has been set.
func (o *ChannelTranscode) HasAudioEncoders() bool {
	if o != nil && o.AudioEncoders != nil {
		return true
	}

	return false
}

// SetAudioEncoders gets a reference to the given []ChannelTranscodeAudioEncodersInner and assigns it to the AudioEncoders field.
func (o *ChannelTranscode) SetAudioEncoders(v []ChannelTranscodeAudioEncodersInner) {
	o.AudioEncoders = v
}

// GetFeatureFlags returns the FeatureFlags field value if set, zero value otherwise.
func (o *ChannelTranscode) GetFeatureFlags() []string {
	if o == nil || o.FeatureFlags == nil {
		var ret []string
		return ret
	}
	return o.FeatureFlags
}

// GetFeatureFlagsOk returns a tuple with the FeatureFlags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscode) GetFeatureFlagsOk() ([]string, bool) {
	if o == nil || o.FeatureFlags == nil {
		return nil, false
	}
	return o.FeatureFlags, true
}

// HasFeatureFlags returns a boolean if a field has been set.
func (o *ChannelTranscode) HasFeatureFlags() bool {
	if o != nil && o.FeatureFlags != nil {
		return true
	}

	return false
}

// SetFeatureFlags gets a reference to the given []string and assigns it to the FeatureFlags field.
func (o *ChannelTranscode) SetFeatureFlags(v []string) {
	o.FeatureFlags = v
}

// GetId3Mode returns the Id3Mode field value if set, zero value otherwise.
func (o *ChannelTranscode) GetId3Mode() string {
	if o == nil || o.Id3Mode == nil {
		var ret string
		return ret
	}
	return *o.Id3Mode
}

// GetId3ModeOk returns a tuple with the Id3Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscode) GetId3ModeOk() (*string, bool) {
	if o == nil || o.Id3Mode == nil {
		return nil, false
	}
	return o.Id3Mode, true
}

// HasId3Mode returns a boolean if a field has been set.
func (o *ChannelTranscode) HasId3Mode() bool {
	if o != nil && o.Id3Mode != nil {
		return true
	}

	return false
}

// SetId3Mode gets a reference to the given string and assigns it to the Id3Mode field.
func (o *ChannelTranscode) SetId3Mode(v string) {
	o.Id3Mode = &v
}

// GetResizeMode returns the ResizeMode field value if set, zero value otherwise.
func (o *ChannelTranscode) GetResizeMode() string {
	if o == nil || o.ResizeMode == nil {
		var ret string
		return ret
	}
	return *o.ResizeMode
}

// GetResizeModeOk returns a tuple with the ResizeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscode) GetResizeModeOk() (*string, bool) {
	if o == nil || o.ResizeMode == nil {
		return nil, false
	}
	return o.ResizeMode, true
}

// HasResizeMode returns a boolean if a field has been set.
func (o *ChannelTranscode) HasResizeMode() bool {
	if o != nil && o.ResizeMode != nil {
		return true
	}

	return false
}

// SetResizeMode gets a reference to the given string and assigns it to the ResizeMode field.
func (o *ChannelTranscode) SetResizeMode(v string) {
	o.ResizeMode = &v
}

// GetSegmenter returns the Segmenter field value if set, zero value otherwise.
func (o *ChannelTranscode) GetSegmenter() ChannelTranscodeSegmenter {
	if o == nil || o.Segmenter == nil {
		var ret ChannelTranscodeSegmenter
		return ret
	}
	return *o.Segmenter
}

// GetSegmenterOk returns a tuple with the Segmenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscode) GetSegmenterOk() (*ChannelTranscodeSegmenter, bool) {
	if o == nil || o.Segmenter == nil {
		return nil, false
	}
	return o.Segmenter, true
}

// HasSegmenter returns a boolean if a field has been set.
func (o *ChannelTranscode) HasSegmenter() bool {
	if o != nil && o.Segmenter != nil {
		return true
	}

	return false
}

// SetSegmenter gets a reference to the given ChannelTranscodeSegmenter and assigns it to the Segmenter field.
func (o *ChannelTranscode) SetSegmenter(v ChannelTranscodeSegmenter) {
	o.Segmenter = &v
}

// GetThumbnailEncoders returns the ThumbnailEncoders field value if set, zero value otherwise.
func (o *ChannelTranscode) GetThumbnailEncoders() []ChannelTranscodeThumbnailEncodersInner {
	if o == nil || o.ThumbnailEncoders == nil {
		var ret []ChannelTranscodeThumbnailEncodersInner
		return ret
	}
	return o.ThumbnailEncoders
}

// GetThumbnailEncodersOk returns a tuple with the ThumbnailEncoders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscode) GetThumbnailEncodersOk() ([]ChannelTranscodeThumbnailEncodersInner, bool) {
	if o == nil || o.ThumbnailEncoders == nil {
		return nil, false
	}
	return o.ThumbnailEncoders, true
}

// HasThumbnailEncoders returns a boolean if a field has been set.
func (o *ChannelTranscode) HasThumbnailEncoders() bool {
	if o != nil && o.ThumbnailEncoders != nil {
		return true
	}

	return false
}

// SetThumbnailEncoders gets a reference to the given []ChannelTranscodeThumbnailEncodersInner and assigns it to the ThumbnailEncoders field.
func (o *ChannelTranscode) SetThumbnailEncoders(v []ChannelTranscodeThumbnailEncodersInner) {
	o.ThumbnailEncoders = v
}

// GetVideoEncoders returns the VideoEncoders field value if set, zero value otherwise.
func (o *ChannelTranscode) GetVideoEncoders() []ChannelTranscodeVideoEncodersInner {
	if o == nil || o.VideoEncoders == nil {
		var ret []ChannelTranscodeVideoEncodersInner
		return ret
	}
	return o.VideoEncoders
}

// GetVideoEncodersOk returns a tuple with the VideoEncoders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscode) GetVideoEncodersOk() ([]ChannelTranscodeVideoEncodersInner, bool) {
	if o == nil || o.VideoEncoders == nil {
		return nil, false
	}
	return o.VideoEncoders, true
}

// HasVideoEncoders returns a boolean if a field has been set.
func (o *ChannelTranscode) HasVideoEncoders() bool {
	if o != nil && o.VideoEncoders != nil {
		return true
	}

	return false
}

// SetVideoEncoders gets a reference to the given []ChannelTranscodeVideoEncodersInner and assigns it to the VideoEncoders field.
func (o *ChannelTranscode) SetVideoEncoders(v []ChannelTranscodeVideoEncodersInner) {
	o.VideoEncoders = v
}

func (o ChannelTranscode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AudioEncoders != nil {
		toSerialize["audio_encoders"] = o.AudioEncoders
	}
	if o.FeatureFlags != nil {
		toSerialize["feature_flags"] = o.FeatureFlags
	}
	if o.Id3Mode != nil {
		toSerialize["id3_mode"] = o.Id3Mode
	}
	if o.ResizeMode != nil {
		toSerialize["resize_mode"] = o.ResizeMode
	}
	if o.Segmenter != nil {
		toSerialize["segmenter"] = o.Segmenter
	}
	if o.ThumbnailEncoders != nil {
		toSerialize["thumbnail_encoders"] = o.ThumbnailEncoders
	}
	if o.VideoEncoders != nil {
		toSerialize["video_encoders"] = o.VideoEncoders
	}
	return json.Marshal(toSerialize)
}

type NullableChannelTranscode struct {
	value *ChannelTranscode
	isSet bool
}

func (v NullableChannelTranscode) Get() *ChannelTranscode {
	return v.value
}

func (v *NullableChannelTranscode) Set(val *ChannelTranscode) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelTranscode) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelTranscode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelTranscode(val *ChannelTranscode) *NullableChannelTranscode {
	return &NullableChannelTranscode{value: val, isSet: true}
}

func (v NullableChannelTranscode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelTranscode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


