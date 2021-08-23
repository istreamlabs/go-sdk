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

// ChannelPublishingHlsPartialPresentations struct for ChannelPublishingHlsPartialPresentations
type ChannelPublishingHlsPartialPresentations struct {
	// Specify which audio encoders should be used for this presentation. If none are specified, all audio encoders configured for the parent Publication will be used.
	AudioEncoderIds *[]string `json:"audio_encoder_ids,omitempty"`
	// List of video encoder IDs that should have I-Frame only playlists generated for them.
	IframeOnlyEncoderIds *[]string `json:"iframe_only_encoder_ids,omitempty"`
	// Sub-path that will be appended onto the publish and playback base URLs of HTTP PublishPoints for published playlist files.
	PlaylistPath *string `json:"playlist_path,omitempty"`
	// Specify which thumbnail encoders should be used for this presentation. If none are specified, all thumbnail encoders configured for the parent Publication will be used.
	ThumbnailEncoderIds *[]string `json:"thumbnail_encoder_ids,omitempty"`
	// Specify which video encoders should be used for this presentation. If none are specified, all video encoders configured for the parent Publication will be used.
	VideoEncoderIds *[]string `json:"video_encoder_ids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChannelPublishingHlsPartialPresentations ChannelPublishingHlsPartialPresentations

// NewChannelPublishingHlsPartialPresentations instantiates a new ChannelPublishingHlsPartialPresentations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelPublishingHlsPartialPresentations() *ChannelPublishingHlsPartialPresentations {
	this := ChannelPublishingHlsPartialPresentations{}
	return &this
}

// NewChannelPublishingHlsPartialPresentationsWithDefaults instantiates a new ChannelPublishingHlsPartialPresentations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelPublishingHlsPartialPresentationsWithDefaults() *ChannelPublishingHlsPartialPresentations {
	this := ChannelPublishingHlsPartialPresentations{}
	return &this
}

// GetAudioEncoderIds returns the AudioEncoderIds field value if set, zero value otherwise.
func (o *ChannelPublishingHlsPartialPresentations) GetAudioEncoderIds() []string {
	if o == nil || o.AudioEncoderIds == nil {
		var ret []string
		return ret
	}
	return *o.AudioEncoderIds
}

// GetAudioEncoderIdsOk returns a tuple with the AudioEncoderIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingHlsPartialPresentations) GetAudioEncoderIdsOk() (*[]string, bool) {
	if o == nil || o.AudioEncoderIds == nil {
		return nil, false
	}
	return o.AudioEncoderIds, true
}

// HasAudioEncoderIds returns a boolean if a field has been set.
func (o *ChannelPublishingHlsPartialPresentations) HasAudioEncoderIds() bool {
	if o != nil && o.AudioEncoderIds != nil {
		return true
	}

	return false
}

// SetAudioEncoderIds gets a reference to the given []string and assigns it to the AudioEncoderIds field.
func (o *ChannelPublishingHlsPartialPresentations) SetAudioEncoderIds(v []string) {
	o.AudioEncoderIds = &v
}

// GetIframeOnlyEncoderIds returns the IframeOnlyEncoderIds field value if set, zero value otherwise.
func (o *ChannelPublishingHlsPartialPresentations) GetIframeOnlyEncoderIds() []string {
	if o == nil || o.IframeOnlyEncoderIds == nil {
		var ret []string
		return ret
	}
	return *o.IframeOnlyEncoderIds
}

// GetIframeOnlyEncoderIdsOk returns a tuple with the IframeOnlyEncoderIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingHlsPartialPresentations) GetIframeOnlyEncoderIdsOk() (*[]string, bool) {
	if o == nil || o.IframeOnlyEncoderIds == nil {
		return nil, false
	}
	return o.IframeOnlyEncoderIds, true
}

// HasIframeOnlyEncoderIds returns a boolean if a field has been set.
func (o *ChannelPublishingHlsPartialPresentations) HasIframeOnlyEncoderIds() bool {
	if o != nil && o.IframeOnlyEncoderIds != nil {
		return true
	}

	return false
}

// SetIframeOnlyEncoderIds gets a reference to the given []string and assigns it to the IframeOnlyEncoderIds field.
func (o *ChannelPublishingHlsPartialPresentations) SetIframeOnlyEncoderIds(v []string) {
	o.IframeOnlyEncoderIds = &v
}

// GetPlaylistPath returns the PlaylistPath field value if set, zero value otherwise.
func (o *ChannelPublishingHlsPartialPresentations) GetPlaylistPath() string {
	if o == nil || o.PlaylistPath == nil {
		var ret string
		return ret
	}
	return *o.PlaylistPath
}

// GetPlaylistPathOk returns a tuple with the PlaylistPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingHlsPartialPresentations) GetPlaylistPathOk() (*string, bool) {
	if o == nil || o.PlaylistPath == nil {
		return nil, false
	}
	return o.PlaylistPath, true
}

// HasPlaylistPath returns a boolean if a field has been set.
func (o *ChannelPublishingHlsPartialPresentations) HasPlaylistPath() bool {
	if o != nil && o.PlaylistPath != nil {
		return true
	}

	return false
}

// SetPlaylistPath gets a reference to the given string and assigns it to the PlaylistPath field.
func (o *ChannelPublishingHlsPartialPresentations) SetPlaylistPath(v string) {
	o.PlaylistPath = &v
}

// GetThumbnailEncoderIds returns the ThumbnailEncoderIds field value if set, zero value otherwise.
func (o *ChannelPublishingHlsPartialPresentations) GetThumbnailEncoderIds() []string {
	if o == nil || o.ThumbnailEncoderIds == nil {
		var ret []string
		return ret
	}
	return *o.ThumbnailEncoderIds
}

// GetThumbnailEncoderIdsOk returns a tuple with the ThumbnailEncoderIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingHlsPartialPresentations) GetThumbnailEncoderIdsOk() (*[]string, bool) {
	if o == nil || o.ThumbnailEncoderIds == nil {
		return nil, false
	}
	return o.ThumbnailEncoderIds, true
}

// HasThumbnailEncoderIds returns a boolean if a field has been set.
func (o *ChannelPublishingHlsPartialPresentations) HasThumbnailEncoderIds() bool {
	if o != nil && o.ThumbnailEncoderIds != nil {
		return true
	}

	return false
}

// SetThumbnailEncoderIds gets a reference to the given []string and assigns it to the ThumbnailEncoderIds field.
func (o *ChannelPublishingHlsPartialPresentations) SetThumbnailEncoderIds(v []string) {
	o.ThumbnailEncoderIds = &v
}

// GetVideoEncoderIds returns the VideoEncoderIds field value if set, zero value otherwise.
func (o *ChannelPublishingHlsPartialPresentations) GetVideoEncoderIds() []string {
	if o == nil || o.VideoEncoderIds == nil {
		var ret []string
		return ret
	}
	return *o.VideoEncoderIds
}

// GetVideoEncoderIdsOk returns a tuple with the VideoEncoderIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingHlsPartialPresentations) GetVideoEncoderIdsOk() (*[]string, bool) {
	if o == nil || o.VideoEncoderIds == nil {
		return nil, false
	}
	return o.VideoEncoderIds, true
}

// HasVideoEncoderIds returns a boolean if a field has been set.
func (o *ChannelPublishingHlsPartialPresentations) HasVideoEncoderIds() bool {
	if o != nil && o.VideoEncoderIds != nil {
		return true
	}

	return false
}

// SetVideoEncoderIds gets a reference to the given []string and assigns it to the VideoEncoderIds field.
func (o *ChannelPublishingHlsPartialPresentations) SetVideoEncoderIds(v []string) {
	o.VideoEncoderIds = &v
}

func (o ChannelPublishingHlsPartialPresentations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AudioEncoderIds != nil {
		toSerialize["audio_encoder_ids"] = o.AudioEncoderIds
	}
	if o.IframeOnlyEncoderIds != nil {
		toSerialize["iframe_only_encoder_ids"] = o.IframeOnlyEncoderIds
	}
	if o.PlaylistPath != nil {
		toSerialize["playlist_path"] = o.PlaylistPath
	}
	if o.ThumbnailEncoderIds != nil {
		toSerialize["thumbnail_encoder_ids"] = o.ThumbnailEncoderIds
	}
	if o.VideoEncoderIds != nil {
		toSerialize["video_encoder_ids"] = o.VideoEncoderIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ChannelPublishingHlsPartialPresentations) UnmarshalJSON(bytes []byte) (err error) {
	varChannelPublishingHlsPartialPresentations := _ChannelPublishingHlsPartialPresentations{}

	if err = json.Unmarshal(bytes, &varChannelPublishingHlsPartialPresentations); err == nil {
		*o = ChannelPublishingHlsPartialPresentations(varChannelPublishingHlsPartialPresentations)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "audio_encoder_ids")
		delete(additionalProperties, "iframe_only_encoder_ids")
		delete(additionalProperties, "playlist_path")
		delete(additionalProperties, "thumbnail_encoder_ids")
		delete(additionalProperties, "video_encoder_ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChannelPublishingHlsPartialPresentations struct {
	value *ChannelPublishingHlsPartialPresentations
	isSet bool
}

func (v NullableChannelPublishingHlsPartialPresentations) Get() *ChannelPublishingHlsPartialPresentations {
	return v.value
}

func (v *NullableChannelPublishingHlsPartialPresentations) Set(val *ChannelPublishingHlsPartialPresentations) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelPublishingHlsPartialPresentations) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelPublishingHlsPartialPresentations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelPublishingHlsPartialPresentations(val *ChannelPublishingHlsPartialPresentations) *NullableChannelPublishingHlsPartialPresentations {
	return &NullableChannelPublishingHlsPartialPresentations{value: val, isSet: true}
}

func (v NullableChannelPublishingHlsPartialPresentations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelPublishingHlsPartialPresentations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


