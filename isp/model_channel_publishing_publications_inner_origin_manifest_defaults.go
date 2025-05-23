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

// checks if the ChannelPublishingPublicationsInnerOriginManifestDefaults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelPublishingPublicationsInnerOriginManifestDefaults{}

// ChannelPublishingPublicationsInnerOriginManifestDefaults ManifestDefaults specifies the default behavior of the dynamic manifest generator. Specific behaviors can be overwritten using appropriate query string parameters when making the request for the manifest.
type ChannelPublishingPublicationsInnerOriginManifestDefaults struct {
	// DASH Signaling formats specifies which SCTE-35 timeline marker formatting to use when rendering DASH manifests.
	DashSignalingFormats []string `json:"dash_signaling_formats,omitempty" uniqueItems:"true" enum:"SCTE35_SPLICE_INFO_SECTION,SCTE35_BIN" doc:"DASH Signaling formats specifies which SCTE-35 timeline marker formatting to use when rendering DASH manifests."`
	// Duration is the length of content that will be included in the manifest, in seconds. The max supported DVR window is 12 hours. If not specified, the default duration will be 30 seconds.
	DurationSeconds *int32 `json:"duration_seconds,omitempty" format:"int32" minimum:"0" maximum:"43200" doc:"Duration is the length of content that will be included in the manifest, in seconds. The max supported DVR window is 12 hours. If not specified, the default duration will be 30 seconds."`
	// When true enables HLS Delta updates for this particular origin.
	EnableHlsDeltaUpdates *bool `json:"enable_hls_delta_updates,omitempty" doc:"When true enables HLS Delta updates for this particular origin."`
	// HLS signaling formats specifies which SCTE-35 timeline marker formatting to use when rendering playlists.
	HlsSignalingFormats []string `json:"hls_signaling_formats,omitempty" uniqueItems:"true" enum:"APPLE_SCTE35,AD_SIMPLE,SCTE35" doc:"HLS signaling formats specifies which SCTE-35 timeline marker formatting to use when rendering playlists."`
	// Allows specifying url type for HLS media playlists and DASH manifests. If not provided, playlist generation will use 'RELATIVE'.
	MediaUrlType *string `json:"media_url_type,omitempty" enum:"RELATIVE,ABSOLUTE,HOST_RELATIVE" doc:"Allows specifying url type for HLS media playlists and DASH manifests. If not provided, playlist generation will use 'RELATIVE'."`
	// Sets the minimumUpdatePeriod field in MPD to be this value. If set to 0 (default), segment duration is used. The value shall not exceed the 'suggested_presentation_delay_secs'.
	MinimumUpdatePeriodSeconds *int32 `json:"minimum_update_period_seconds,omitempty" format:"int32" minimum:"0" doc:"Sets the minimumUpdatePeriod field in MPD to be this value. If set to 0 (default), segment duration is used. The value shall not exceed the 'suggested_presentation_delay_secs'."`
	// Allows specifying url type for HLS multi-variant playlists. If not provided, playlist generation will use 'RELATIVE'.
	MultiVariantUrlType *string `json:"multi_variant_url_type,omitempty" enum:"RELATIVE,ABSOLUTE,HOST_RELATIVE" doc:"Allows specifying url type for HLS multi-variant playlists. If not provided, playlist generation will use 'RELATIVE'."`
	// When true a #EXT-X-PROGRAM-DATE-TIME tag will be placed on every media segment in HLS media playlists. When false, the default behavior, the PDT tag is set according to the HLS specification.
	PdtOnEverySegment *bool `json:"pdt_on_every_segment,omitempty" doc:"When true a #EXT-X-PROGRAM-DATE-TIME tag will be placed on every media segment in HLS media playlists. When false, the default behavior, the PDT tag is set according to the HLS specification."`
	// Pre Start Duration is the length of content that will be included in the manifest when the start time of a manifest is in the future. When a manifest is near the start time but does not contain 'pre_start_duration_seconds' content Manifest Origin will keep content before the start time to make the manifest length at least 'pre_start_duration_seconds'. If not specified requesting a manifest in the future will result in a 404.
	PreStartDurationSeconds *int32 `json:"pre_start_duration_seconds,omitempty" format:"int32" minimum:"0" maximum:"43200" doc:"Pre Start Duration is the length of content that will be included in the manifest when the start time of a manifest is in the future. When a manifest is near the start time but does not contain 'pre_start_duration_seconds' content Manifest Origin will keep content before the start time to make the manifest length at least 'pre_start_duration_seconds'. If not specified requesting a manifest in the future will result in a 404."`
	// Sets the suggestedPresentationDelay field in MPD to be this value. This value must be greater or equal to 'minimum_update_period_secs'. If unset, the default value will be calculated as 3 * segment target duration.
	SuggestedPresentationDelaySeconds *int32 `json:"suggested_presentation_delay_seconds,omitempty" format:"int32" minimum:"0" doc:"Sets the suggestedPresentationDelay field in MPD to be this value. This value must be greater or equal to 'minimum_update_period_secs'. If unset, the default value will be calculated as 3 * segment target duration."`
}

// NewChannelPublishingPublicationsInnerOriginManifestDefaults instantiates a new ChannelPublishingPublicationsInnerOriginManifestDefaults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelPublishingPublicationsInnerOriginManifestDefaults() *ChannelPublishingPublicationsInnerOriginManifestDefaults {
	this := ChannelPublishingPublicationsInnerOriginManifestDefaults{}
	return &this
}

// NewChannelPublishingPublicationsInnerOriginManifestDefaultsWithDefaults instantiates a new ChannelPublishingPublicationsInnerOriginManifestDefaults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelPublishingPublicationsInnerOriginManifestDefaultsWithDefaults() *ChannelPublishingPublicationsInnerOriginManifestDefaults {
	this := ChannelPublishingPublicationsInnerOriginManifestDefaults{}
	return &this
}

// GetDashSignalingFormats returns the DashSignalingFormats field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) GetDashSignalingFormats() []string {
	if o == nil || IsNil(o.DashSignalingFormats) {
		var ret []string
		return ret
	}
	return o.DashSignalingFormats
}

// GetDashSignalingFormatsOk returns a tuple with the DashSignalingFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) GetDashSignalingFormatsOk() ([]string, bool) {
	if o == nil || IsNil(o.DashSignalingFormats) {
		return nil, false
	}
	return o.DashSignalingFormats, true
}

// HasDashSignalingFormats returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) HasDashSignalingFormats() bool {
	if o != nil && !IsNil(o.DashSignalingFormats) {
		return true
	}

	return false
}

// SetDashSignalingFormats gets a reference to the given []string and assigns it to the DashSignalingFormats field.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) SetDashSignalingFormats(v []string) {
	o.DashSignalingFormats = v
}

// GetDurationSeconds returns the DurationSeconds field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) GetDurationSeconds() int32 {
	if o == nil || IsNil(o.DurationSeconds) {
		var ret int32
		return ret
	}
	return *o.DurationSeconds
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) GetDurationSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.DurationSeconds) {
		return nil, false
	}
	return o.DurationSeconds, true
}

// HasDurationSeconds returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) HasDurationSeconds() bool {
	if o != nil && !IsNil(o.DurationSeconds) {
		return true
	}

	return false
}

// SetDurationSeconds gets a reference to the given int32 and assigns it to the DurationSeconds field.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) SetDurationSeconds(v int32) {
	o.DurationSeconds = &v
}

// GetEnableHlsDeltaUpdates returns the EnableHlsDeltaUpdates field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) GetEnableHlsDeltaUpdates() bool {
	if o == nil || IsNil(o.EnableHlsDeltaUpdates) {
		var ret bool
		return ret
	}
	return *o.EnableHlsDeltaUpdates
}

// GetEnableHlsDeltaUpdatesOk returns a tuple with the EnableHlsDeltaUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) GetEnableHlsDeltaUpdatesOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableHlsDeltaUpdates) {
		return nil, false
	}
	return o.EnableHlsDeltaUpdates, true
}

// HasEnableHlsDeltaUpdates returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) HasEnableHlsDeltaUpdates() bool {
	if o != nil && !IsNil(o.EnableHlsDeltaUpdates) {
		return true
	}

	return false
}

// SetEnableHlsDeltaUpdates gets a reference to the given bool and assigns it to the EnableHlsDeltaUpdates field.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) SetEnableHlsDeltaUpdates(v bool) {
	o.EnableHlsDeltaUpdates = &v
}

// GetHlsSignalingFormats returns the HlsSignalingFormats field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) GetHlsSignalingFormats() []string {
	if o == nil || IsNil(o.HlsSignalingFormats) {
		var ret []string
		return ret
	}
	return o.HlsSignalingFormats
}

// GetHlsSignalingFormatsOk returns a tuple with the HlsSignalingFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) GetHlsSignalingFormatsOk() ([]string, bool) {
	if o == nil || IsNil(o.HlsSignalingFormats) {
		return nil, false
	}
	return o.HlsSignalingFormats, true
}

// HasHlsSignalingFormats returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) HasHlsSignalingFormats() bool {
	if o != nil && !IsNil(o.HlsSignalingFormats) {
		return true
	}

	return false
}

// SetHlsSignalingFormats gets a reference to the given []string and assigns it to the HlsSignalingFormats field.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) SetHlsSignalingFormats(v []string) {
	o.HlsSignalingFormats = v
}

// GetMediaUrlType returns the MediaUrlType field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) GetMediaUrlType() string {
	if o == nil || IsNil(o.MediaUrlType) {
		var ret string
		return ret
	}
	return *o.MediaUrlType
}

// GetMediaUrlTypeOk returns a tuple with the MediaUrlType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) GetMediaUrlTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MediaUrlType) {
		return nil, false
	}
	return o.MediaUrlType, true
}

// HasMediaUrlType returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) HasMediaUrlType() bool {
	if o != nil && !IsNil(o.MediaUrlType) {
		return true
	}

	return false
}

// SetMediaUrlType gets a reference to the given string and assigns it to the MediaUrlType field.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) SetMediaUrlType(v string) {
	o.MediaUrlType = &v
}

// GetMinimumUpdatePeriodSeconds returns the MinimumUpdatePeriodSeconds field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) GetMinimumUpdatePeriodSeconds() int32 {
	if o == nil || IsNil(o.MinimumUpdatePeriodSeconds) {
		var ret int32
		return ret
	}
	return *o.MinimumUpdatePeriodSeconds
}

// GetMinimumUpdatePeriodSecondsOk returns a tuple with the MinimumUpdatePeriodSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) GetMinimumUpdatePeriodSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.MinimumUpdatePeriodSeconds) {
		return nil, false
	}
	return o.MinimumUpdatePeriodSeconds, true
}

// HasMinimumUpdatePeriodSeconds returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) HasMinimumUpdatePeriodSeconds() bool {
	if o != nil && !IsNil(o.MinimumUpdatePeriodSeconds) {
		return true
	}

	return false
}

// SetMinimumUpdatePeriodSeconds gets a reference to the given int32 and assigns it to the MinimumUpdatePeriodSeconds field.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) SetMinimumUpdatePeriodSeconds(v int32) {
	o.MinimumUpdatePeriodSeconds = &v
}

// GetMultiVariantUrlType returns the MultiVariantUrlType field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) GetMultiVariantUrlType() string {
	if o == nil || IsNil(o.MultiVariantUrlType) {
		var ret string
		return ret
	}
	return *o.MultiVariantUrlType
}

// GetMultiVariantUrlTypeOk returns a tuple with the MultiVariantUrlType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) GetMultiVariantUrlTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MultiVariantUrlType) {
		return nil, false
	}
	return o.MultiVariantUrlType, true
}

// HasMultiVariantUrlType returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) HasMultiVariantUrlType() bool {
	if o != nil && !IsNil(o.MultiVariantUrlType) {
		return true
	}

	return false
}

// SetMultiVariantUrlType gets a reference to the given string and assigns it to the MultiVariantUrlType field.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) SetMultiVariantUrlType(v string) {
	o.MultiVariantUrlType = &v
}

// GetPdtOnEverySegment returns the PdtOnEverySegment field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) GetPdtOnEverySegment() bool {
	if o == nil || IsNil(o.PdtOnEverySegment) {
		var ret bool
		return ret
	}
	return *o.PdtOnEverySegment
}

// GetPdtOnEverySegmentOk returns a tuple with the PdtOnEverySegment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) GetPdtOnEverySegmentOk() (*bool, bool) {
	if o == nil || IsNil(o.PdtOnEverySegment) {
		return nil, false
	}
	return o.PdtOnEverySegment, true
}

// HasPdtOnEverySegment returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) HasPdtOnEverySegment() bool {
	if o != nil && !IsNil(o.PdtOnEverySegment) {
		return true
	}

	return false
}

// SetPdtOnEverySegment gets a reference to the given bool and assigns it to the PdtOnEverySegment field.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) SetPdtOnEverySegment(v bool) {
	o.PdtOnEverySegment = &v
}

// GetPreStartDurationSeconds returns the PreStartDurationSeconds field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) GetPreStartDurationSeconds() int32 {
	if o == nil || IsNil(o.PreStartDurationSeconds) {
		var ret int32
		return ret
	}
	return *o.PreStartDurationSeconds
}

// GetPreStartDurationSecondsOk returns a tuple with the PreStartDurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) GetPreStartDurationSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.PreStartDurationSeconds) {
		return nil, false
	}
	return o.PreStartDurationSeconds, true
}

// HasPreStartDurationSeconds returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) HasPreStartDurationSeconds() bool {
	if o != nil && !IsNil(o.PreStartDurationSeconds) {
		return true
	}

	return false
}

// SetPreStartDurationSeconds gets a reference to the given int32 and assigns it to the PreStartDurationSeconds field.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) SetPreStartDurationSeconds(v int32) {
	o.PreStartDurationSeconds = &v
}

// GetSuggestedPresentationDelaySeconds returns the SuggestedPresentationDelaySeconds field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) GetSuggestedPresentationDelaySeconds() int32 {
	if o == nil || IsNil(o.SuggestedPresentationDelaySeconds) {
		var ret int32
		return ret
	}
	return *o.SuggestedPresentationDelaySeconds
}

// GetSuggestedPresentationDelaySecondsOk returns a tuple with the SuggestedPresentationDelaySeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) GetSuggestedPresentationDelaySecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.SuggestedPresentationDelaySeconds) {
		return nil, false
	}
	return o.SuggestedPresentationDelaySeconds, true
}

// HasSuggestedPresentationDelaySeconds returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) HasSuggestedPresentationDelaySeconds() bool {
	if o != nil && !IsNil(o.SuggestedPresentationDelaySeconds) {
		return true
	}

	return false
}

// SetSuggestedPresentationDelaySeconds gets a reference to the given int32 and assigns it to the SuggestedPresentationDelaySeconds field.
func (o *ChannelPublishingPublicationsInnerOriginManifestDefaults) SetSuggestedPresentationDelaySeconds(v int32) {
	o.SuggestedPresentationDelaySeconds = &v
}

func (o ChannelPublishingPublicationsInnerOriginManifestDefaults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelPublishingPublicationsInnerOriginManifestDefaults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DashSignalingFormats) {
		toSerialize["dash_signaling_formats"] = o.DashSignalingFormats
	}
	if !IsNil(o.DurationSeconds) {
		toSerialize["duration_seconds"] = o.DurationSeconds
	}
	if !IsNil(o.EnableHlsDeltaUpdates) {
		toSerialize["enable_hls_delta_updates"] = o.EnableHlsDeltaUpdates
	}
	if !IsNil(o.HlsSignalingFormats) {
		toSerialize["hls_signaling_formats"] = o.HlsSignalingFormats
	}
	if !IsNil(o.MediaUrlType) {
		toSerialize["media_url_type"] = o.MediaUrlType
	}
	if !IsNil(o.MinimumUpdatePeriodSeconds) {
		toSerialize["minimum_update_period_seconds"] = o.MinimumUpdatePeriodSeconds
	}
	if !IsNil(o.MultiVariantUrlType) {
		toSerialize["multi_variant_url_type"] = o.MultiVariantUrlType
	}
	if !IsNil(o.PdtOnEverySegment) {
		toSerialize["pdt_on_every_segment"] = o.PdtOnEverySegment
	}
	if !IsNil(o.PreStartDurationSeconds) {
		toSerialize["pre_start_duration_seconds"] = o.PreStartDurationSeconds
	}
	if !IsNil(o.SuggestedPresentationDelaySeconds) {
		toSerialize["suggested_presentation_delay_seconds"] = o.SuggestedPresentationDelaySeconds
	}
	return toSerialize, nil
}

type NullableChannelPublishingPublicationsInnerOriginManifestDefaults struct {
	value *ChannelPublishingPublicationsInnerOriginManifestDefaults
	isSet bool
}

func (v NullableChannelPublishingPublicationsInnerOriginManifestDefaults) Get() *ChannelPublishingPublicationsInnerOriginManifestDefaults {
	return v.value
}

func (v *NullableChannelPublishingPublicationsInnerOriginManifestDefaults) Set(val *ChannelPublishingPublicationsInnerOriginManifestDefaults) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelPublishingPublicationsInnerOriginManifestDefaults) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelPublishingPublicationsInnerOriginManifestDefaults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelPublishingPublicationsInnerOriginManifestDefaults(val *ChannelPublishingPublicationsInnerOriginManifestDefaults) *NullableChannelPublishingPublicationsInnerOriginManifestDefaults {
	return &NullableChannelPublishingPublicationsInnerOriginManifestDefaults{value: val, isSet: true}
}

func (v NullableChannelPublishingPublicationsInnerOriginManifestDefaults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelPublishingPublicationsInnerOriginManifestDefaults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


