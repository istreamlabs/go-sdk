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

// checks if the ChannelTranscode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelTranscode{}

// ChannelTranscode Transcode configures audio/video conversion settings.
type ChannelTranscode struct {
	// Audio encoders specify audio conversion settings, e.g. channels, samples, codec, bitrate, etc.
	AudioEncoders []ChannelTranscodeAudioEncodersInner `json:"audio_encoders,omitempty" minItems:"1" doc:"Audio encoders specify audio conversion settings, e.g. channels, samples, codec, bitrate, etc."`
	// Deprecated: Do not use. Debug_overlay overlays debugging information from the transcoder into the top right of the video output. The overlay is burned into the video and will be visible to end-users if enabled. Do not enable on customer facing channels. Requires a transcoder restart if the state is changed. The default value is false, which disables the overlay. This setting is deprecated in favour of debug_overlay2.
	// Deprecated
	DebugOverlay *bool `json:"debug_overlay,omitempty" deprecated:"true" doc:"Deprecated: Do not use. Debug_overlay overlays debugging information from the transcoder into the top right of the video output. The overlay is burned into the video and will be visible to end-users if enabled. Do not enable on customer facing channels. Requires a transcoder restart if the state is changed. The default value is false, which disables the overlay. This setting is deprecated in favour of debug_overlay2."`
	DebugOverlays *ChannelTranscodeDebugOverlays `json:"debug_overlays,omitempty"`
	// Feature flag strings enable experimental transcode features or functionality that are not yet or never will be promoted to the channeldoc model proper.
	FeatureFlags []string `json:"feature_flags,omitempty" doc:"Feature flag strings enable experimental transcode features or functionality that are not yet or never will be promoted to the channeldoc model proper."`
	// Specify how to process ID3 tags from the input source. If not specified, ID3 tags in the source will be ignored.
	Id3Mode *string `json:"id3_mode,omitempty" enum:"PASSTHROUGH" doc:"Specify how to process ID3 tags from the input source. If not specified, ID3 tags in the source will be ignored."`
	// List of overlays. An overlay is an image that will be rendered on top of the source video. Only one overlay is supported at the moment. If specified, the overlay will be always rendered unless a video slate is on.
	Overlays []ChannelTranscodeOverlaysInner `json:"overlays,omitempty" maxItems:"1" doc:"List of overlays. An overlay is an image that will be rendered on top of the source video. Only one overlay is supported at the moment. If specified, the overlay will be always rendered unless a video slate is on."`
	// Resize mode specifies how to scale a video up or down to match the output dimensions.
	ResizeMode *string `json:"resize_mode,omitempty" enum:"STRETCH,LETTERBOX" doc:"Resize mode specifies how to scale a video up or down to match the output dimensions."`
	Segmenter *ChannelTranscodeSegmenter `json:"segmenter,omitempty"`
	// Subtitle encoders specify how text-based subtitles are extracted into separate segments. They are not used to describe CEA 608 captions, which remain part of the video codec.
	SubtitleEncoders []ChannelTranscodeSubtitleEncodersInner `json:"subtitle_encoders,omitempty" doc:"Subtitle encoders specify how text-based subtitles are extracted into separate segments. They are not used to describe CEA 608 captions, which remain part of the video codec."`
	// Thumbnail encoders specify how to create image snapshots of the video stream.
	ThumbnailEncoders []ChannelTranscodeThumbnailEncodersInner `json:"thumbnail_encoders,omitempty" doc:"Thumbnail encoders specify how to create image snapshots of the video stream."`
	// Video encoders specify video conversion settings, e.g. dimensions, codec, bitrate, etc.
	VideoEncoders []ChannelTranscodeVideoEncodersInner `json:"video_encoders,omitempty" minItems:"1" doc:"Video encoders specify video conversion settings, e.g. dimensions, codec, bitrate, etc."`
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
	if o == nil || IsNil(o.AudioEncoders) {
		var ret []ChannelTranscodeAudioEncodersInner
		return ret
	}
	return o.AudioEncoders
}

// GetAudioEncodersOk returns a tuple with the AudioEncoders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscode) GetAudioEncodersOk() ([]ChannelTranscodeAudioEncodersInner, bool) {
	if o == nil || IsNil(o.AudioEncoders) {
		return nil, false
	}
	return o.AudioEncoders, true
}

// HasAudioEncoders returns a boolean if a field has been set.
func (o *ChannelTranscode) HasAudioEncoders() bool {
	if o != nil && !IsNil(o.AudioEncoders) {
		return true
	}

	return false
}

// SetAudioEncoders gets a reference to the given []ChannelTranscodeAudioEncodersInner and assigns it to the AudioEncoders field.
func (o *ChannelTranscode) SetAudioEncoders(v []ChannelTranscodeAudioEncodersInner) {
	o.AudioEncoders = v
}

// GetDebugOverlay returns the DebugOverlay field value if set, zero value otherwise.
// Deprecated
func (o *ChannelTranscode) GetDebugOverlay() bool {
	if o == nil || IsNil(o.DebugOverlay) {
		var ret bool
		return ret
	}
	return *o.DebugOverlay
}

// GetDebugOverlayOk returns a tuple with the DebugOverlay field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ChannelTranscode) GetDebugOverlayOk() (*bool, bool) {
	if o == nil || IsNil(o.DebugOverlay) {
		return nil, false
	}
	return o.DebugOverlay, true
}

// HasDebugOverlay returns a boolean if a field has been set.
func (o *ChannelTranscode) HasDebugOverlay() bool {
	if o != nil && !IsNil(o.DebugOverlay) {
		return true
	}

	return false
}

// SetDebugOverlay gets a reference to the given bool and assigns it to the DebugOverlay field.
// Deprecated
func (o *ChannelTranscode) SetDebugOverlay(v bool) {
	o.DebugOverlay = &v
}

// GetDebugOverlays returns the DebugOverlays field value if set, zero value otherwise.
func (o *ChannelTranscode) GetDebugOverlays() ChannelTranscodeDebugOverlays {
	if o == nil || IsNil(o.DebugOverlays) {
		var ret ChannelTranscodeDebugOverlays
		return ret
	}
	return *o.DebugOverlays
}

// GetDebugOverlaysOk returns a tuple with the DebugOverlays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscode) GetDebugOverlaysOk() (*ChannelTranscodeDebugOverlays, bool) {
	if o == nil || IsNil(o.DebugOverlays) {
		return nil, false
	}
	return o.DebugOverlays, true
}

// HasDebugOverlays returns a boolean if a field has been set.
func (o *ChannelTranscode) HasDebugOverlays() bool {
	if o != nil && !IsNil(o.DebugOverlays) {
		return true
	}

	return false
}

// SetDebugOverlays gets a reference to the given ChannelTranscodeDebugOverlays and assigns it to the DebugOverlays field.
func (o *ChannelTranscode) SetDebugOverlays(v ChannelTranscodeDebugOverlays) {
	o.DebugOverlays = &v
}

// GetFeatureFlags returns the FeatureFlags field value if set, zero value otherwise.
func (o *ChannelTranscode) GetFeatureFlags() []string {
	if o == nil || IsNil(o.FeatureFlags) {
		var ret []string
		return ret
	}
	return o.FeatureFlags
}

// GetFeatureFlagsOk returns a tuple with the FeatureFlags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscode) GetFeatureFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.FeatureFlags) {
		return nil, false
	}
	return o.FeatureFlags, true
}

// HasFeatureFlags returns a boolean if a field has been set.
func (o *ChannelTranscode) HasFeatureFlags() bool {
	if o != nil && !IsNil(o.FeatureFlags) {
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
	if o == nil || IsNil(o.Id3Mode) {
		var ret string
		return ret
	}
	return *o.Id3Mode
}

// GetId3ModeOk returns a tuple with the Id3Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscode) GetId3ModeOk() (*string, bool) {
	if o == nil || IsNil(o.Id3Mode) {
		return nil, false
	}
	return o.Id3Mode, true
}

// HasId3Mode returns a boolean if a field has been set.
func (o *ChannelTranscode) HasId3Mode() bool {
	if o != nil && !IsNil(o.Id3Mode) {
		return true
	}

	return false
}

// SetId3Mode gets a reference to the given string and assigns it to the Id3Mode field.
func (o *ChannelTranscode) SetId3Mode(v string) {
	o.Id3Mode = &v
}

// GetOverlays returns the Overlays field value if set, zero value otherwise.
func (o *ChannelTranscode) GetOverlays() []ChannelTranscodeOverlaysInner {
	if o == nil || IsNil(o.Overlays) {
		var ret []ChannelTranscodeOverlaysInner
		return ret
	}
	return o.Overlays
}

// GetOverlaysOk returns a tuple with the Overlays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscode) GetOverlaysOk() ([]ChannelTranscodeOverlaysInner, bool) {
	if o == nil || IsNil(o.Overlays) {
		return nil, false
	}
	return o.Overlays, true
}

// HasOverlays returns a boolean if a field has been set.
func (o *ChannelTranscode) HasOverlays() bool {
	if o != nil && !IsNil(o.Overlays) {
		return true
	}

	return false
}

// SetOverlays gets a reference to the given []ChannelTranscodeOverlaysInner and assigns it to the Overlays field.
func (o *ChannelTranscode) SetOverlays(v []ChannelTranscodeOverlaysInner) {
	o.Overlays = v
}

// GetResizeMode returns the ResizeMode field value if set, zero value otherwise.
func (o *ChannelTranscode) GetResizeMode() string {
	if o == nil || IsNil(o.ResizeMode) {
		var ret string
		return ret
	}
	return *o.ResizeMode
}

// GetResizeModeOk returns a tuple with the ResizeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscode) GetResizeModeOk() (*string, bool) {
	if o == nil || IsNil(o.ResizeMode) {
		return nil, false
	}
	return o.ResizeMode, true
}

// HasResizeMode returns a boolean if a field has been set.
func (o *ChannelTranscode) HasResizeMode() bool {
	if o != nil && !IsNil(o.ResizeMode) {
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
	if o == nil || IsNil(o.Segmenter) {
		var ret ChannelTranscodeSegmenter
		return ret
	}
	return *o.Segmenter
}

// GetSegmenterOk returns a tuple with the Segmenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscode) GetSegmenterOk() (*ChannelTranscodeSegmenter, bool) {
	if o == nil || IsNil(o.Segmenter) {
		return nil, false
	}
	return o.Segmenter, true
}

// HasSegmenter returns a boolean if a field has been set.
func (o *ChannelTranscode) HasSegmenter() bool {
	if o != nil && !IsNil(o.Segmenter) {
		return true
	}

	return false
}

// SetSegmenter gets a reference to the given ChannelTranscodeSegmenter and assigns it to the Segmenter field.
func (o *ChannelTranscode) SetSegmenter(v ChannelTranscodeSegmenter) {
	o.Segmenter = &v
}

// GetSubtitleEncoders returns the SubtitleEncoders field value if set, zero value otherwise.
func (o *ChannelTranscode) GetSubtitleEncoders() []ChannelTranscodeSubtitleEncodersInner {
	if o == nil || IsNil(o.SubtitleEncoders) {
		var ret []ChannelTranscodeSubtitleEncodersInner
		return ret
	}
	return o.SubtitleEncoders
}

// GetSubtitleEncodersOk returns a tuple with the SubtitleEncoders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscode) GetSubtitleEncodersOk() ([]ChannelTranscodeSubtitleEncodersInner, bool) {
	if o == nil || IsNil(o.SubtitleEncoders) {
		return nil, false
	}
	return o.SubtitleEncoders, true
}

// HasSubtitleEncoders returns a boolean if a field has been set.
func (o *ChannelTranscode) HasSubtitleEncoders() bool {
	if o != nil && !IsNil(o.SubtitleEncoders) {
		return true
	}

	return false
}

// SetSubtitleEncoders gets a reference to the given []ChannelTranscodeSubtitleEncodersInner and assigns it to the SubtitleEncoders field.
func (o *ChannelTranscode) SetSubtitleEncoders(v []ChannelTranscodeSubtitleEncodersInner) {
	o.SubtitleEncoders = v
}

// GetThumbnailEncoders returns the ThumbnailEncoders field value if set, zero value otherwise.
func (o *ChannelTranscode) GetThumbnailEncoders() []ChannelTranscodeThumbnailEncodersInner {
	if o == nil || IsNil(o.ThumbnailEncoders) {
		var ret []ChannelTranscodeThumbnailEncodersInner
		return ret
	}
	return o.ThumbnailEncoders
}

// GetThumbnailEncodersOk returns a tuple with the ThumbnailEncoders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscode) GetThumbnailEncodersOk() ([]ChannelTranscodeThumbnailEncodersInner, bool) {
	if o == nil || IsNil(o.ThumbnailEncoders) {
		return nil, false
	}
	return o.ThumbnailEncoders, true
}

// HasThumbnailEncoders returns a boolean if a field has been set.
func (o *ChannelTranscode) HasThumbnailEncoders() bool {
	if o != nil && !IsNil(o.ThumbnailEncoders) {
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
	if o == nil || IsNil(o.VideoEncoders) {
		var ret []ChannelTranscodeVideoEncodersInner
		return ret
	}
	return o.VideoEncoders
}

// GetVideoEncodersOk returns a tuple with the VideoEncoders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscode) GetVideoEncodersOk() ([]ChannelTranscodeVideoEncodersInner, bool) {
	if o == nil || IsNil(o.VideoEncoders) {
		return nil, false
	}
	return o.VideoEncoders, true
}

// HasVideoEncoders returns a boolean if a field has been set.
func (o *ChannelTranscode) HasVideoEncoders() bool {
	if o != nil && !IsNil(o.VideoEncoders) {
		return true
	}

	return false
}

// SetVideoEncoders gets a reference to the given []ChannelTranscodeVideoEncodersInner and assigns it to the VideoEncoders field.
func (o *ChannelTranscode) SetVideoEncoders(v []ChannelTranscodeVideoEncodersInner) {
	o.VideoEncoders = v
}

func (o ChannelTranscode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelTranscode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AudioEncoders) {
		toSerialize["audio_encoders"] = o.AudioEncoders
	}
	if !IsNil(o.DebugOverlay) {
		toSerialize["debug_overlay"] = o.DebugOverlay
	}
	if !IsNil(o.DebugOverlays) {
		toSerialize["debug_overlays"] = o.DebugOverlays
	}
	if !IsNil(o.FeatureFlags) {
		toSerialize["feature_flags"] = o.FeatureFlags
	}
	if !IsNil(o.Id3Mode) {
		toSerialize["id3_mode"] = o.Id3Mode
	}
	if !IsNil(o.Overlays) {
		toSerialize["overlays"] = o.Overlays
	}
	if !IsNil(o.ResizeMode) {
		toSerialize["resize_mode"] = o.ResizeMode
	}
	if !IsNil(o.Segmenter) {
		toSerialize["segmenter"] = o.Segmenter
	}
	if !IsNil(o.SubtitleEncoders) {
		toSerialize["subtitle_encoders"] = o.SubtitleEncoders
	}
	if !IsNil(o.ThumbnailEncoders) {
		toSerialize["thumbnail_encoders"] = o.ThumbnailEncoders
	}
	if !IsNil(o.VideoEncoders) {
		toSerialize["video_encoders"] = o.VideoEncoders
	}
	return toSerialize, nil
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


