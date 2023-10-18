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

// checks if the ChannelSignalingSegmentSettingsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelSignalingSegmentSettingsInner{}

// ChannelSignalingSegmentSettingsInner struct for ChannelSignalingSegmentSettingsInner
type ChannelSignalingSegmentSettingsInner struct {
	// Specifies the duration of a segment when the in-band SCTE-35 that initiates it (e.g. Distributor Placement Opportunity Start) is missing an explicit duration. N.B. for program and ad types, this also affects 'Simple Program' and 'Simple Ad' markers, respectively.
	DefaultDurationSecs *int32 `json:"default_duration_secs,omitempty" format:"int32" minimum:"0" maximum:"172800" doc:"Specifies the duration of a segment when the in-band SCTE-35 that initiates it (e.g. Distributor Placement Opportunity Start) is missing an explicit duration. N.B. for program and ad types, this also affects 'Simple Program' and 'Simple Ad' markers, respectively."`
	// Determines whether to include the default duration in the output SCTE-35 messages when the input SCTE-35 message did not specify a duration.
	EmitDefaultDuration *bool `json:"emit_default_duration,omitempty" doc:"Determines whether to include the default duration in the output SCTE-35 messages when the input SCTE-35 message did not specify a duration."`
	// Specifies a 'correction' to the splice_time of in-band SCTE-35 in milliseconds.
	OffsetMillis *int32 `json:"offset_millis,omitempty" format:"int32" minimum:"-4000" maximum:"4000" doc:"Specifies a 'correction' to the splice_time of in-band SCTE-35 in milliseconds."`
	// Determines which Segment End signaling mode to use for the provided segments. If unspecified, defaults to MATCH_END_EVENT_ID.
	SegmentEndMode *string `json:"segment_end_mode,omitempty" enum:"MATCH_END_EVENT_ID,IGNORE_END_EVENT_ID,IGNORE_END_SIGNAL" doc:"Determines which Segment End signaling mode to use for the provided segments. If unspecified, defaults to MATCH_END_EVENT_ID."`
	// Specifies the list of which segment types this setting applies to. Any segment type defined here _must_ also be present in the parent signaling configuration.
	Segments []string `json:"segments,omitempty" uniqueItems:"true" enum:"SPLICE_INSERT,CONTENT_ID,PROGRAM,PROGRAM_BLACKOUT_OVERRIDE,PROGRAM_BREAKAWAY,CHAPTER,BREAK,OPENING_CREDIT,CLOSING_CREDIT,PROVIDER_PLACEMENT_OP,DISTRIBUTOR_PLACEMENT_OP,PROVIDER_OVERLAY_OP,DISTRIBUTOR_OVERLAY_OP,PROVIDER_AD,DISTRIBUTOR_AD,UNSCHEDULED_EVENT,NETWORK,SLATE" doc:"Specifies the list of which segment types this setting applies to. Any segment type defined here _must_ also be present in the parent signaling configuration."`
	TierFilter *ChannelSignalingSegmentSettingsInnerTierFilter `json:"tier_filter,omitempty"`
}

// NewChannelSignalingSegmentSettingsInner instantiates a new ChannelSignalingSegmentSettingsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelSignalingSegmentSettingsInner() *ChannelSignalingSegmentSettingsInner {
	this := ChannelSignalingSegmentSettingsInner{}
	return &this
}

// NewChannelSignalingSegmentSettingsInnerWithDefaults instantiates a new ChannelSignalingSegmentSettingsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelSignalingSegmentSettingsInnerWithDefaults() *ChannelSignalingSegmentSettingsInner {
	this := ChannelSignalingSegmentSettingsInner{}
	return &this
}

// GetDefaultDurationSecs returns the DefaultDurationSecs field value if set, zero value otherwise.
func (o *ChannelSignalingSegmentSettingsInner) GetDefaultDurationSecs() int32 {
	if o == nil || IsNil(o.DefaultDurationSecs) {
		var ret int32
		return ret
	}
	return *o.DefaultDurationSecs
}

// GetDefaultDurationSecsOk returns a tuple with the DefaultDurationSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelSignalingSegmentSettingsInner) GetDefaultDurationSecsOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultDurationSecs) {
		return nil, false
	}
	return o.DefaultDurationSecs, true
}

// HasDefaultDurationSecs returns a boolean if a field has been set.
func (o *ChannelSignalingSegmentSettingsInner) HasDefaultDurationSecs() bool {
	if o != nil && !IsNil(o.DefaultDurationSecs) {
		return true
	}

	return false
}

// SetDefaultDurationSecs gets a reference to the given int32 and assigns it to the DefaultDurationSecs field.
func (o *ChannelSignalingSegmentSettingsInner) SetDefaultDurationSecs(v int32) {
	o.DefaultDurationSecs = &v
}

// GetEmitDefaultDuration returns the EmitDefaultDuration field value if set, zero value otherwise.
func (o *ChannelSignalingSegmentSettingsInner) GetEmitDefaultDuration() bool {
	if o == nil || IsNil(o.EmitDefaultDuration) {
		var ret bool
		return ret
	}
	return *o.EmitDefaultDuration
}

// GetEmitDefaultDurationOk returns a tuple with the EmitDefaultDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelSignalingSegmentSettingsInner) GetEmitDefaultDurationOk() (*bool, bool) {
	if o == nil || IsNil(o.EmitDefaultDuration) {
		return nil, false
	}
	return o.EmitDefaultDuration, true
}

// HasEmitDefaultDuration returns a boolean if a field has been set.
func (o *ChannelSignalingSegmentSettingsInner) HasEmitDefaultDuration() bool {
	if o != nil && !IsNil(o.EmitDefaultDuration) {
		return true
	}

	return false
}

// SetEmitDefaultDuration gets a reference to the given bool and assigns it to the EmitDefaultDuration field.
func (o *ChannelSignalingSegmentSettingsInner) SetEmitDefaultDuration(v bool) {
	o.EmitDefaultDuration = &v
}

// GetOffsetMillis returns the OffsetMillis field value if set, zero value otherwise.
func (o *ChannelSignalingSegmentSettingsInner) GetOffsetMillis() int32 {
	if o == nil || IsNil(o.OffsetMillis) {
		var ret int32
		return ret
	}
	return *o.OffsetMillis
}

// GetOffsetMillisOk returns a tuple with the OffsetMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelSignalingSegmentSettingsInner) GetOffsetMillisOk() (*int32, bool) {
	if o == nil || IsNil(o.OffsetMillis) {
		return nil, false
	}
	return o.OffsetMillis, true
}

// HasOffsetMillis returns a boolean if a field has been set.
func (o *ChannelSignalingSegmentSettingsInner) HasOffsetMillis() bool {
	if o != nil && !IsNil(o.OffsetMillis) {
		return true
	}

	return false
}

// SetOffsetMillis gets a reference to the given int32 and assigns it to the OffsetMillis field.
func (o *ChannelSignalingSegmentSettingsInner) SetOffsetMillis(v int32) {
	o.OffsetMillis = &v
}

// GetSegmentEndMode returns the SegmentEndMode field value if set, zero value otherwise.
func (o *ChannelSignalingSegmentSettingsInner) GetSegmentEndMode() string {
	if o == nil || IsNil(o.SegmentEndMode) {
		var ret string
		return ret
	}
	return *o.SegmentEndMode
}

// GetSegmentEndModeOk returns a tuple with the SegmentEndMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelSignalingSegmentSettingsInner) GetSegmentEndModeOk() (*string, bool) {
	if o == nil || IsNil(o.SegmentEndMode) {
		return nil, false
	}
	return o.SegmentEndMode, true
}

// HasSegmentEndMode returns a boolean if a field has been set.
func (o *ChannelSignalingSegmentSettingsInner) HasSegmentEndMode() bool {
	if o != nil && !IsNil(o.SegmentEndMode) {
		return true
	}

	return false
}

// SetSegmentEndMode gets a reference to the given string and assigns it to the SegmentEndMode field.
func (o *ChannelSignalingSegmentSettingsInner) SetSegmentEndMode(v string) {
	o.SegmentEndMode = &v
}

// GetSegments returns the Segments field value if set, zero value otherwise.
func (o *ChannelSignalingSegmentSettingsInner) GetSegments() []string {
	if o == nil || IsNil(o.Segments) {
		var ret []string
		return ret
	}
	return o.Segments
}

// GetSegmentsOk returns a tuple with the Segments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelSignalingSegmentSettingsInner) GetSegmentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Segments) {
		return nil, false
	}
	return o.Segments, true
}

// HasSegments returns a boolean if a field has been set.
func (o *ChannelSignalingSegmentSettingsInner) HasSegments() bool {
	if o != nil && !IsNil(o.Segments) {
		return true
	}

	return false
}

// SetSegments gets a reference to the given []string and assigns it to the Segments field.
func (o *ChannelSignalingSegmentSettingsInner) SetSegments(v []string) {
	o.Segments = v
}

// GetTierFilter returns the TierFilter field value if set, zero value otherwise.
func (o *ChannelSignalingSegmentSettingsInner) GetTierFilter() ChannelSignalingSegmentSettingsInnerTierFilter {
	if o == nil || IsNil(o.TierFilter) {
		var ret ChannelSignalingSegmentSettingsInnerTierFilter
		return ret
	}
	return *o.TierFilter
}

// GetTierFilterOk returns a tuple with the TierFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelSignalingSegmentSettingsInner) GetTierFilterOk() (*ChannelSignalingSegmentSettingsInnerTierFilter, bool) {
	if o == nil || IsNil(o.TierFilter) {
		return nil, false
	}
	return o.TierFilter, true
}

// HasTierFilter returns a boolean if a field has been set.
func (o *ChannelSignalingSegmentSettingsInner) HasTierFilter() bool {
	if o != nil && !IsNil(o.TierFilter) {
		return true
	}

	return false
}

// SetTierFilter gets a reference to the given ChannelSignalingSegmentSettingsInnerTierFilter and assigns it to the TierFilter field.
func (o *ChannelSignalingSegmentSettingsInner) SetTierFilter(v ChannelSignalingSegmentSettingsInnerTierFilter) {
	o.TierFilter = &v
}

func (o ChannelSignalingSegmentSettingsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelSignalingSegmentSettingsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultDurationSecs) {
		toSerialize["default_duration_secs"] = o.DefaultDurationSecs
	}
	if !IsNil(o.EmitDefaultDuration) {
		toSerialize["emit_default_duration"] = o.EmitDefaultDuration
	}
	if !IsNil(o.OffsetMillis) {
		toSerialize["offset_millis"] = o.OffsetMillis
	}
	if !IsNil(o.SegmentEndMode) {
		toSerialize["segment_end_mode"] = o.SegmentEndMode
	}
	if !IsNil(o.Segments) {
		toSerialize["segments"] = o.Segments
	}
	if !IsNil(o.TierFilter) {
		toSerialize["tier_filter"] = o.TierFilter
	}
	return toSerialize, nil
}

type NullableChannelSignalingSegmentSettingsInner struct {
	value *ChannelSignalingSegmentSettingsInner
	isSet bool
}

func (v NullableChannelSignalingSegmentSettingsInner) Get() *ChannelSignalingSegmentSettingsInner {
	return v.value
}

func (v *NullableChannelSignalingSegmentSettingsInner) Set(val *ChannelSignalingSegmentSettingsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelSignalingSegmentSettingsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelSignalingSegmentSettingsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelSignalingSegmentSettingsInner(val *ChannelSignalingSegmentSettingsInner) *NullableChannelSignalingSegmentSettingsInner {
	return &NullableChannelSignalingSegmentSettingsInner{value: val, isSet: true}
}

func (v NullableChannelSignalingSegmentSettingsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelSignalingSegmentSettingsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


