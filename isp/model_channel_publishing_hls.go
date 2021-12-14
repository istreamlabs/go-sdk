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

// ChannelPublishingHls HLS configures publication settings. Only one of HLS or DASH can be set.
type ChannelPublishingHls struct {
	// Defines how audio only variant streams are included in the master playlist, where the variant streams are defined by #EXT-X-STREAM-INF tag, the tag attributes provide information about the Stream. If NOT_SET - honor the deprecated 'exclude_audio_only' flag. Later when the deprecated flag is removed, the NOT_SET would mean INCLUDE_DEFAULT The INCLUDE_DEFAULT option - only the default 'audio only variant stream' is included in master playlist. This is the most common use case. INCLUDE_NONE - no audio only variant streams are included in the master playlist, it replaces 'exclude_audio_only' setting. INCLUDE_ALL - include all audio only variant streams in the master playlist.
	AudioOnlyVariants *string `json:"audio_only_variants,omitempty"`
	// Allows turning gap tags ON/OFF. When turned ON - the tag '#EXT-X-GAP' is inserted into media playlist for a missing segment. When turned OFF - Discontinuity is inserted into the playlist for missing segment(s). The default option UNDEFINED is mapped to OFF. Note: Gap tags are always inserted for the missing thumbnail segments independently of this setting
	GapTags *string `json:"gap_tags,omitempty"`
	// How often the master playlist(s) should be published in seconds. A value of 0 means the master playlist will only be published once at channel start.
	MasterPublishFrequencySecs *int32 `json:"master_publish_frequency_secs,omitempty"`
	// Allows specifying url type for HLS master playlists. If not provided, playlist generation will use 'RELATIVE'.
	MasterUrlType *string `json:"master_url_type,omitempty"`
	// Allows specifying url type for HLS media playlists. If not provided, playlist generation will use 'RELATIVE'.
	MediaUrlType *string `json:"media_url_type,omitempty"`
	// Specify which partial presentations should be used for this presentation. Partial presentations are additional master playlists that point to a subset of the parent presentation's media streams/variant playlists.
	PartialPresentations *[]ChannelPublishingHlsPartialPresentations `json:"partial_presentations,omitempty"`
	// When true a #EXT-X-PROGRAM-DATE-TIME tag will be placed on every media segment in media playlists. When false, the default behavior, the PDT tag is set according to the HLS specification.
	PdtOnEverySegment *bool `json:"pdt_on_every_segment,omitempty"`
	// Signaling formats specifies which SCTE-35 timeline marker formatting to use when rendering playlists.
	SignalingFormats *[]string `json:"signaling_formats,omitempty"`
	// Include a UTC timestamp (that is equivalent in value to #EXT-X-PROGRAM-DATE-TIME) in the title of each media segment in media playlists. Ex. #EXTINF:6.006,LTC=2020-01-01T12:00:00.000Z
	UtcInSegmentTitle *bool `json:"utc_in_segment_title,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChannelPublishingHls ChannelPublishingHls

// NewChannelPublishingHls instantiates a new ChannelPublishingHls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelPublishingHls() *ChannelPublishingHls {
	this := ChannelPublishingHls{}
	return &this
}

// NewChannelPublishingHlsWithDefaults instantiates a new ChannelPublishingHls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelPublishingHlsWithDefaults() *ChannelPublishingHls {
	this := ChannelPublishingHls{}
	return &this
}

// GetAudioOnlyVariants returns the AudioOnlyVariants field value if set, zero value otherwise.
func (o *ChannelPublishingHls) GetAudioOnlyVariants() string {
	if o == nil || o.AudioOnlyVariants == nil {
		var ret string
		return ret
	}
	return *o.AudioOnlyVariants
}

// GetAudioOnlyVariantsOk returns a tuple with the AudioOnlyVariants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingHls) GetAudioOnlyVariantsOk() (*string, bool) {
	if o == nil || o.AudioOnlyVariants == nil {
		return nil, false
	}
	return o.AudioOnlyVariants, true
}

// HasAudioOnlyVariants returns a boolean if a field has been set.
func (o *ChannelPublishingHls) HasAudioOnlyVariants() bool {
	if o != nil && o.AudioOnlyVariants != nil {
		return true
	}

	return false
}

// SetAudioOnlyVariants gets a reference to the given string and assigns it to the AudioOnlyVariants field.
func (o *ChannelPublishingHls) SetAudioOnlyVariants(v string) {
	o.AudioOnlyVariants = &v
}

// GetGapTags returns the GapTags field value if set, zero value otherwise.
func (o *ChannelPublishingHls) GetGapTags() string {
	if o == nil || o.GapTags == nil {
		var ret string
		return ret
	}
	return *o.GapTags
}

// GetGapTagsOk returns a tuple with the GapTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingHls) GetGapTagsOk() (*string, bool) {
	if o == nil || o.GapTags == nil {
		return nil, false
	}
	return o.GapTags, true
}

// HasGapTags returns a boolean if a field has been set.
func (o *ChannelPublishingHls) HasGapTags() bool {
	if o != nil && o.GapTags != nil {
		return true
	}

	return false
}

// SetGapTags gets a reference to the given string and assigns it to the GapTags field.
func (o *ChannelPublishingHls) SetGapTags(v string) {
	o.GapTags = &v
}

// GetMasterPublishFrequencySecs returns the MasterPublishFrequencySecs field value if set, zero value otherwise.
func (o *ChannelPublishingHls) GetMasterPublishFrequencySecs() int32 {
	if o == nil || o.MasterPublishFrequencySecs == nil {
		var ret int32
		return ret
	}
	return *o.MasterPublishFrequencySecs
}

// GetMasterPublishFrequencySecsOk returns a tuple with the MasterPublishFrequencySecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingHls) GetMasterPublishFrequencySecsOk() (*int32, bool) {
	if o == nil || o.MasterPublishFrequencySecs == nil {
		return nil, false
	}
	return o.MasterPublishFrequencySecs, true
}

// HasMasterPublishFrequencySecs returns a boolean if a field has been set.
func (o *ChannelPublishingHls) HasMasterPublishFrequencySecs() bool {
	if o != nil && o.MasterPublishFrequencySecs != nil {
		return true
	}

	return false
}

// SetMasterPublishFrequencySecs gets a reference to the given int32 and assigns it to the MasterPublishFrequencySecs field.
func (o *ChannelPublishingHls) SetMasterPublishFrequencySecs(v int32) {
	o.MasterPublishFrequencySecs = &v
}

// GetMasterUrlType returns the MasterUrlType field value if set, zero value otherwise.
func (o *ChannelPublishingHls) GetMasterUrlType() string {
	if o == nil || o.MasterUrlType == nil {
		var ret string
		return ret
	}
	return *o.MasterUrlType
}

// GetMasterUrlTypeOk returns a tuple with the MasterUrlType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingHls) GetMasterUrlTypeOk() (*string, bool) {
	if o == nil || o.MasterUrlType == nil {
		return nil, false
	}
	return o.MasterUrlType, true
}

// HasMasterUrlType returns a boolean if a field has been set.
func (o *ChannelPublishingHls) HasMasterUrlType() bool {
	if o != nil && o.MasterUrlType != nil {
		return true
	}

	return false
}

// SetMasterUrlType gets a reference to the given string and assigns it to the MasterUrlType field.
func (o *ChannelPublishingHls) SetMasterUrlType(v string) {
	o.MasterUrlType = &v
}

// GetMediaUrlType returns the MediaUrlType field value if set, zero value otherwise.
func (o *ChannelPublishingHls) GetMediaUrlType() string {
	if o == nil || o.MediaUrlType == nil {
		var ret string
		return ret
	}
	return *o.MediaUrlType
}

// GetMediaUrlTypeOk returns a tuple with the MediaUrlType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingHls) GetMediaUrlTypeOk() (*string, bool) {
	if o == nil || o.MediaUrlType == nil {
		return nil, false
	}
	return o.MediaUrlType, true
}

// HasMediaUrlType returns a boolean if a field has been set.
func (o *ChannelPublishingHls) HasMediaUrlType() bool {
	if o != nil && o.MediaUrlType != nil {
		return true
	}

	return false
}

// SetMediaUrlType gets a reference to the given string and assigns it to the MediaUrlType field.
func (o *ChannelPublishingHls) SetMediaUrlType(v string) {
	o.MediaUrlType = &v
}

// GetPartialPresentations returns the PartialPresentations field value if set, zero value otherwise.
func (o *ChannelPublishingHls) GetPartialPresentations() []ChannelPublishingHlsPartialPresentations {
	if o == nil || o.PartialPresentations == nil {
		var ret []ChannelPublishingHlsPartialPresentations
		return ret
	}
	return *o.PartialPresentations
}

// GetPartialPresentationsOk returns a tuple with the PartialPresentations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingHls) GetPartialPresentationsOk() (*[]ChannelPublishingHlsPartialPresentations, bool) {
	if o == nil || o.PartialPresentations == nil {
		return nil, false
	}
	return o.PartialPresentations, true
}

// HasPartialPresentations returns a boolean if a field has been set.
func (o *ChannelPublishingHls) HasPartialPresentations() bool {
	if o != nil && o.PartialPresentations != nil {
		return true
	}

	return false
}

// SetPartialPresentations gets a reference to the given []ChannelPublishingHlsPartialPresentations and assigns it to the PartialPresentations field.
func (o *ChannelPublishingHls) SetPartialPresentations(v []ChannelPublishingHlsPartialPresentations) {
	o.PartialPresentations = &v
}

// GetPdtOnEverySegment returns the PdtOnEverySegment field value if set, zero value otherwise.
func (o *ChannelPublishingHls) GetPdtOnEverySegment() bool {
	if o == nil || o.PdtOnEverySegment == nil {
		var ret bool
		return ret
	}
	return *o.PdtOnEverySegment
}

// GetPdtOnEverySegmentOk returns a tuple with the PdtOnEverySegment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingHls) GetPdtOnEverySegmentOk() (*bool, bool) {
	if o == nil || o.PdtOnEverySegment == nil {
		return nil, false
	}
	return o.PdtOnEverySegment, true
}

// HasPdtOnEverySegment returns a boolean if a field has been set.
func (o *ChannelPublishingHls) HasPdtOnEverySegment() bool {
	if o != nil && o.PdtOnEverySegment != nil {
		return true
	}

	return false
}

// SetPdtOnEverySegment gets a reference to the given bool and assigns it to the PdtOnEverySegment field.
func (o *ChannelPublishingHls) SetPdtOnEverySegment(v bool) {
	o.PdtOnEverySegment = &v
}

// GetSignalingFormats returns the SignalingFormats field value if set, zero value otherwise.
func (o *ChannelPublishingHls) GetSignalingFormats() []string {
	if o == nil || o.SignalingFormats == nil {
		var ret []string
		return ret
	}
	return *o.SignalingFormats
}

// GetSignalingFormatsOk returns a tuple with the SignalingFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingHls) GetSignalingFormatsOk() (*[]string, bool) {
	if o == nil || o.SignalingFormats == nil {
		return nil, false
	}
	return o.SignalingFormats, true
}

// HasSignalingFormats returns a boolean if a field has been set.
func (o *ChannelPublishingHls) HasSignalingFormats() bool {
	if o != nil && o.SignalingFormats != nil {
		return true
	}

	return false
}

// SetSignalingFormats gets a reference to the given []string and assigns it to the SignalingFormats field.
func (o *ChannelPublishingHls) SetSignalingFormats(v []string) {
	o.SignalingFormats = &v
}

// GetUtcInSegmentTitle returns the UtcInSegmentTitle field value if set, zero value otherwise.
func (o *ChannelPublishingHls) GetUtcInSegmentTitle() bool {
	if o == nil || o.UtcInSegmentTitle == nil {
		var ret bool
		return ret
	}
	return *o.UtcInSegmentTitle
}

// GetUtcInSegmentTitleOk returns a tuple with the UtcInSegmentTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingHls) GetUtcInSegmentTitleOk() (*bool, bool) {
	if o == nil || o.UtcInSegmentTitle == nil {
		return nil, false
	}
	return o.UtcInSegmentTitle, true
}

// HasUtcInSegmentTitle returns a boolean if a field has been set.
func (o *ChannelPublishingHls) HasUtcInSegmentTitle() bool {
	if o != nil && o.UtcInSegmentTitle != nil {
		return true
	}

	return false
}

// SetUtcInSegmentTitle gets a reference to the given bool and assigns it to the UtcInSegmentTitle field.
func (o *ChannelPublishingHls) SetUtcInSegmentTitle(v bool) {
	o.UtcInSegmentTitle = &v
}

func (o ChannelPublishingHls) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AudioOnlyVariants != nil {
		toSerialize["audio_only_variants"] = o.AudioOnlyVariants
	}
	if o.GapTags != nil {
		toSerialize["gap_tags"] = o.GapTags
	}
	if o.MasterPublishFrequencySecs != nil {
		toSerialize["master_publish_frequency_secs"] = o.MasterPublishFrequencySecs
	}
	if o.MasterUrlType != nil {
		toSerialize["master_url_type"] = o.MasterUrlType
	}
	if o.MediaUrlType != nil {
		toSerialize["media_url_type"] = o.MediaUrlType
	}
	if o.PartialPresentations != nil {
		toSerialize["partial_presentations"] = o.PartialPresentations
	}
	if o.PdtOnEverySegment != nil {
		toSerialize["pdt_on_every_segment"] = o.PdtOnEverySegment
	}
	if o.SignalingFormats != nil {
		toSerialize["signaling_formats"] = o.SignalingFormats
	}
	if o.UtcInSegmentTitle != nil {
		toSerialize["utc_in_segment_title"] = o.UtcInSegmentTitle
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ChannelPublishingHls) UnmarshalJSON(bytes []byte) (err error) {
	varChannelPublishingHls := _ChannelPublishingHls{}

	if err = json.Unmarshal(bytes, &varChannelPublishingHls); err == nil {
		*o = ChannelPublishingHls(varChannelPublishingHls)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "audio_only_variants")
		delete(additionalProperties, "gap_tags")
		delete(additionalProperties, "master_publish_frequency_secs")
		delete(additionalProperties, "master_url_type")
		delete(additionalProperties, "media_url_type")
		delete(additionalProperties, "partial_presentations")
		delete(additionalProperties, "pdt_on_every_segment")
		delete(additionalProperties, "signaling_formats")
		delete(additionalProperties, "utc_in_segment_title")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChannelPublishingHls struct {
	value *ChannelPublishingHls
	isSet bool
}

func (v NullableChannelPublishingHls) Get() *ChannelPublishingHls {
	return v.value
}

func (v *NullableChannelPublishingHls) Set(val *ChannelPublishingHls) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelPublishingHls) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelPublishingHls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelPublishingHls(val *ChannelPublishingHls) *NullableChannelPublishingHls {
	return &NullableChannelPublishingHls{value: val, isSet: true}
}

func (v NullableChannelPublishingHls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelPublishingHls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


