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

// checks if the ChannelPublishingPublicationsInnerHls type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelPublishingPublicationsInnerHls{}

// ChannelPublishingPublicationsInnerHls HLS configures publication settings. Only one of HLS or DASH can be set.
type ChannelPublishingPublicationsInnerHls struct {
	// Defines how audio only variant streams are included in the master playlist, where the variant streams are defined by #EXT-X-STREAM-INF tag, the tag attributes provide information about the Stream. The INCLUDE_DEFAULT option - only the default 'audio only variant stream' is included in master playlist. This is the most common use case. INCLUDE_NONE - no audio only variant streams are included in the master playlist. INCLUDE_ALL - include all audio only variant streams in the master playlist.
	AudioOnlyVariants *string `json:"audio_only_variants,omitempty" enum:"INCLUDE_DEFAULT,INCLUDE_NONE,INCLUDE_ALL" doc:"Defines how audio only variant streams are included in the master playlist, where the variant streams are defined by #EXT-X-STREAM-INF tag, the tag attributes provide information about the Stream. The INCLUDE_DEFAULT option - only the default 'audio only variant stream' is included in master playlist. This is the most common use case. INCLUDE_NONE - no audio only variant streams are included in the master playlist. INCLUDE_ALL - include all audio only variant streams in the master playlist."`
	// Allows turning gap tags ON/OFF. When turned ON - the tag '#EXT-X-GAP' is inserted into media playlist for a missing segment. When turned OFF - Discontinuity is inserted into the playlist for missing segment(s). The default option UNDEFINED is mapped to OFF. Note: Gap tags are always inserted for the missing thumbnail segments independently of this setting
	GapTags *string `json:"gap_tags,omitempty" enum:"ON,OFF" doc:"Allows turning gap tags ON/OFF. When turned ON - the tag '#EXT-X-GAP' is inserted into media playlist for a missing segment. When turned OFF - Discontinuity is inserted into the playlist for missing segment(s). The default option UNDEFINED is mapped to OFF. Note: Gap tags are always inserted for the missing thumbnail segments independently of this setting"`
	// How often the multi-variant playlist(s) should be published (in seconds) once a channel is started. If 0 (not provided), defaults to 1 hour. The following can also result in re-publishing of multi-variant playlists: - Configuration changes related to transcoding configuration (i.e. adding, removing, changing encoders). - Service code changes related to manifest generation.
	MasterPublishFrequencySecs *int32 `json:"master_publish_frequency_secs,omitempty" format:"int32" doc:"How often the multi-variant playlist(s) should be published (in seconds) once a channel is started. If 0 (not provided), defaults to 1 hour. The following can also result in re-publishing of multi-variant playlists: - Configuration changes related to transcoding configuration (i.e. adding, removing, changing encoders). - Service code changes related to manifest generation."`
	// Allows specifying url type for HLS master playlists. If not provided, playlist generation will use 'RELATIVE'.
	MasterUrlType *string `json:"master_url_type,omitempty" enum:"RELATIVE,ABSOLUTE,HOST_RELATIVE" doc:"Allows specifying url type for HLS master playlists. If not provided, playlist generation will use 'RELATIVE'."`
	// Allows specifying url type for HLS media playlists. If not provided, playlist generation will use 'RELATIVE'.
	MediaUrlType *string `json:"media_url_type,omitempty" enum:"RELATIVE,ABSOLUTE,HOST_RELATIVE" doc:"Allows specifying url type for HLS media playlists. If not provided, playlist generation will use 'RELATIVE'."`
	// Specify which partial presentations should be used for this presentation. Partial presentations are additional master playlists that point to a subset of the parent presentation's media streams/variant playlists.
	PartialPresentations []ChannelPublishingPublicationsInnerHlsPartialPresentationsInner `json:"partial_presentations,omitempty" doc:"Specify which partial presentations should be used for this presentation. Partial presentations are additional master playlists that point to a subset of the parent presentation's media streams/variant playlists."`
	// When true a #EXT-X-PROGRAM-DATE-TIME tag will be placed on every media segment in media playlists. When false, the default behavior, the PDT tag is set according to the HLS specification.
	PdtOnEverySegment *bool `json:"pdt_on_every_segment,omitempty" doc:"When true a #EXT-X-PROGRAM-DATE-TIME tag will be placed on every media segment in media playlists. When false, the default behavior, the PDT tag is set according to the HLS specification."`
	// Signaling formats specifies which SCTE-35 timeline marker formatting to use when rendering playlists.
	SignalingFormats []string `json:"signaling_formats,omitempty" uniqueItems:"true" enum:"APPLE_SCTE35,AD_SIMPLE,SCTE35" doc:"Signaling formats specifies which SCTE-35 timeline marker formatting to use when rendering playlists."`
	// Include a UTC timestamp (that is equivalent in value to #EXT-X-PROGRAM-DATE-TIME) in the title of each media segment in media playlists. Ex. #EXTINF:6.006,LTC=2020-01-01T12:00:00.000Z
	UtcInSegmentTitle *bool `json:"utc_in_segment_title,omitempty" doc:"Include a UTC timestamp (that is equivalent in value to #EXT-X-PROGRAM-DATE-TIME) in the title of each media segment in media playlists. Ex. #EXTINF:6.006,LTC=2020-01-01T12:00:00.000Z"`
}

// NewChannelPublishingPublicationsInnerHls instantiates a new ChannelPublishingPublicationsInnerHls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelPublishingPublicationsInnerHls() *ChannelPublishingPublicationsInnerHls {
	this := ChannelPublishingPublicationsInnerHls{}
	return &this
}

// NewChannelPublishingPublicationsInnerHlsWithDefaults instantiates a new ChannelPublishingPublicationsInnerHls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelPublishingPublicationsInnerHlsWithDefaults() *ChannelPublishingPublicationsInnerHls {
	this := ChannelPublishingPublicationsInnerHls{}
	return &this
}

// GetAudioOnlyVariants returns the AudioOnlyVariants field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInnerHls) GetAudioOnlyVariants() string {
	if o == nil || IsNil(o.AudioOnlyVariants) {
		var ret string
		return ret
	}
	return *o.AudioOnlyVariants
}

// GetAudioOnlyVariantsOk returns a tuple with the AudioOnlyVariants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInnerHls) GetAudioOnlyVariantsOk() (*string, bool) {
	if o == nil || IsNil(o.AudioOnlyVariants) {
		return nil, false
	}
	return o.AudioOnlyVariants, true
}

// HasAudioOnlyVariants returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInnerHls) HasAudioOnlyVariants() bool {
	if o != nil && !IsNil(o.AudioOnlyVariants) {
		return true
	}

	return false
}

// SetAudioOnlyVariants gets a reference to the given string and assigns it to the AudioOnlyVariants field.
func (o *ChannelPublishingPublicationsInnerHls) SetAudioOnlyVariants(v string) {
	o.AudioOnlyVariants = &v
}

// GetGapTags returns the GapTags field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInnerHls) GetGapTags() string {
	if o == nil || IsNil(o.GapTags) {
		var ret string
		return ret
	}
	return *o.GapTags
}

// GetGapTagsOk returns a tuple with the GapTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInnerHls) GetGapTagsOk() (*string, bool) {
	if o == nil || IsNil(o.GapTags) {
		return nil, false
	}
	return o.GapTags, true
}

// HasGapTags returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInnerHls) HasGapTags() bool {
	if o != nil && !IsNil(o.GapTags) {
		return true
	}

	return false
}

// SetGapTags gets a reference to the given string and assigns it to the GapTags field.
func (o *ChannelPublishingPublicationsInnerHls) SetGapTags(v string) {
	o.GapTags = &v
}

// GetMasterPublishFrequencySecs returns the MasterPublishFrequencySecs field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInnerHls) GetMasterPublishFrequencySecs() int32 {
	if o == nil || IsNil(o.MasterPublishFrequencySecs) {
		var ret int32
		return ret
	}
	return *o.MasterPublishFrequencySecs
}

// GetMasterPublishFrequencySecsOk returns a tuple with the MasterPublishFrequencySecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInnerHls) GetMasterPublishFrequencySecsOk() (*int32, bool) {
	if o == nil || IsNil(o.MasterPublishFrequencySecs) {
		return nil, false
	}
	return o.MasterPublishFrequencySecs, true
}

// HasMasterPublishFrequencySecs returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInnerHls) HasMasterPublishFrequencySecs() bool {
	if o != nil && !IsNil(o.MasterPublishFrequencySecs) {
		return true
	}

	return false
}

// SetMasterPublishFrequencySecs gets a reference to the given int32 and assigns it to the MasterPublishFrequencySecs field.
func (o *ChannelPublishingPublicationsInnerHls) SetMasterPublishFrequencySecs(v int32) {
	o.MasterPublishFrequencySecs = &v
}

// GetMasterUrlType returns the MasterUrlType field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInnerHls) GetMasterUrlType() string {
	if o == nil || IsNil(o.MasterUrlType) {
		var ret string
		return ret
	}
	return *o.MasterUrlType
}

// GetMasterUrlTypeOk returns a tuple with the MasterUrlType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInnerHls) GetMasterUrlTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MasterUrlType) {
		return nil, false
	}
	return o.MasterUrlType, true
}

// HasMasterUrlType returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInnerHls) HasMasterUrlType() bool {
	if o != nil && !IsNil(o.MasterUrlType) {
		return true
	}

	return false
}

// SetMasterUrlType gets a reference to the given string and assigns it to the MasterUrlType field.
func (o *ChannelPublishingPublicationsInnerHls) SetMasterUrlType(v string) {
	o.MasterUrlType = &v
}

// GetMediaUrlType returns the MediaUrlType field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInnerHls) GetMediaUrlType() string {
	if o == nil || IsNil(o.MediaUrlType) {
		var ret string
		return ret
	}
	return *o.MediaUrlType
}

// GetMediaUrlTypeOk returns a tuple with the MediaUrlType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInnerHls) GetMediaUrlTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MediaUrlType) {
		return nil, false
	}
	return o.MediaUrlType, true
}

// HasMediaUrlType returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInnerHls) HasMediaUrlType() bool {
	if o != nil && !IsNil(o.MediaUrlType) {
		return true
	}

	return false
}

// SetMediaUrlType gets a reference to the given string and assigns it to the MediaUrlType field.
func (o *ChannelPublishingPublicationsInnerHls) SetMediaUrlType(v string) {
	o.MediaUrlType = &v
}

// GetPartialPresentations returns the PartialPresentations field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInnerHls) GetPartialPresentations() []ChannelPublishingPublicationsInnerHlsPartialPresentationsInner {
	if o == nil || IsNil(o.PartialPresentations) {
		var ret []ChannelPublishingPublicationsInnerHlsPartialPresentationsInner
		return ret
	}
	return o.PartialPresentations
}

// GetPartialPresentationsOk returns a tuple with the PartialPresentations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInnerHls) GetPartialPresentationsOk() ([]ChannelPublishingPublicationsInnerHlsPartialPresentationsInner, bool) {
	if o == nil || IsNil(o.PartialPresentations) {
		return nil, false
	}
	return o.PartialPresentations, true
}

// HasPartialPresentations returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInnerHls) HasPartialPresentations() bool {
	if o != nil && !IsNil(o.PartialPresentations) {
		return true
	}

	return false
}

// SetPartialPresentations gets a reference to the given []ChannelPublishingPublicationsInnerHlsPartialPresentationsInner and assigns it to the PartialPresentations field.
func (o *ChannelPublishingPublicationsInnerHls) SetPartialPresentations(v []ChannelPublishingPublicationsInnerHlsPartialPresentationsInner) {
	o.PartialPresentations = v
}

// GetPdtOnEverySegment returns the PdtOnEverySegment field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInnerHls) GetPdtOnEverySegment() bool {
	if o == nil || IsNil(o.PdtOnEverySegment) {
		var ret bool
		return ret
	}
	return *o.PdtOnEverySegment
}

// GetPdtOnEverySegmentOk returns a tuple with the PdtOnEverySegment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInnerHls) GetPdtOnEverySegmentOk() (*bool, bool) {
	if o == nil || IsNil(o.PdtOnEverySegment) {
		return nil, false
	}
	return o.PdtOnEverySegment, true
}

// HasPdtOnEverySegment returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInnerHls) HasPdtOnEverySegment() bool {
	if o != nil && !IsNil(o.PdtOnEverySegment) {
		return true
	}

	return false
}

// SetPdtOnEverySegment gets a reference to the given bool and assigns it to the PdtOnEverySegment field.
func (o *ChannelPublishingPublicationsInnerHls) SetPdtOnEverySegment(v bool) {
	o.PdtOnEverySegment = &v
}

// GetSignalingFormats returns the SignalingFormats field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInnerHls) GetSignalingFormats() []string {
	if o == nil || IsNil(o.SignalingFormats) {
		var ret []string
		return ret
	}
	return o.SignalingFormats
}

// GetSignalingFormatsOk returns a tuple with the SignalingFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInnerHls) GetSignalingFormatsOk() ([]string, bool) {
	if o == nil || IsNil(o.SignalingFormats) {
		return nil, false
	}
	return o.SignalingFormats, true
}

// HasSignalingFormats returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInnerHls) HasSignalingFormats() bool {
	if o != nil && !IsNil(o.SignalingFormats) {
		return true
	}

	return false
}

// SetSignalingFormats gets a reference to the given []string and assigns it to the SignalingFormats field.
func (o *ChannelPublishingPublicationsInnerHls) SetSignalingFormats(v []string) {
	o.SignalingFormats = v
}

// GetUtcInSegmentTitle returns the UtcInSegmentTitle field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInnerHls) GetUtcInSegmentTitle() bool {
	if o == nil || IsNil(o.UtcInSegmentTitle) {
		var ret bool
		return ret
	}
	return *o.UtcInSegmentTitle
}

// GetUtcInSegmentTitleOk returns a tuple with the UtcInSegmentTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInnerHls) GetUtcInSegmentTitleOk() (*bool, bool) {
	if o == nil || IsNil(o.UtcInSegmentTitle) {
		return nil, false
	}
	return o.UtcInSegmentTitle, true
}

// HasUtcInSegmentTitle returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInnerHls) HasUtcInSegmentTitle() bool {
	if o != nil && !IsNil(o.UtcInSegmentTitle) {
		return true
	}

	return false
}

// SetUtcInSegmentTitle gets a reference to the given bool and assigns it to the UtcInSegmentTitle field.
func (o *ChannelPublishingPublicationsInnerHls) SetUtcInSegmentTitle(v bool) {
	o.UtcInSegmentTitle = &v
}

func (o ChannelPublishingPublicationsInnerHls) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelPublishingPublicationsInnerHls) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AudioOnlyVariants) {
		toSerialize["audio_only_variants"] = o.AudioOnlyVariants
	}
	if !IsNil(o.GapTags) {
		toSerialize["gap_tags"] = o.GapTags
	}
	if !IsNil(o.MasterPublishFrequencySecs) {
		toSerialize["master_publish_frequency_secs"] = o.MasterPublishFrequencySecs
	}
	if !IsNil(o.MasterUrlType) {
		toSerialize["master_url_type"] = o.MasterUrlType
	}
	if !IsNil(o.MediaUrlType) {
		toSerialize["media_url_type"] = o.MediaUrlType
	}
	if !IsNil(o.PartialPresentations) {
		toSerialize["partial_presentations"] = o.PartialPresentations
	}
	if !IsNil(o.PdtOnEverySegment) {
		toSerialize["pdt_on_every_segment"] = o.PdtOnEverySegment
	}
	if !IsNil(o.SignalingFormats) {
		toSerialize["signaling_formats"] = o.SignalingFormats
	}
	if !IsNil(o.UtcInSegmentTitle) {
		toSerialize["utc_in_segment_title"] = o.UtcInSegmentTitle
	}
	return toSerialize, nil
}

type NullableChannelPublishingPublicationsInnerHls struct {
	value *ChannelPublishingPublicationsInnerHls
	isSet bool
}

func (v NullableChannelPublishingPublicationsInnerHls) Get() *ChannelPublishingPublicationsInnerHls {
	return v.value
}

func (v *NullableChannelPublishingPublicationsInnerHls) Set(val *ChannelPublishingPublicationsInnerHls) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelPublishingPublicationsInnerHls) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelPublishingPublicationsInnerHls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelPublishingPublicationsInnerHls(val *ChannelPublishingPublicationsInnerHls) *NullableChannelPublishingPublicationsInnerHls {
	return &NullableChannelPublishingPublicationsInnerHls{value: val, isSet: true}
}

func (v NullableChannelPublishingPublicationsInnerHls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelPublishingPublicationsInnerHls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
