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

// checks if the ChannelSignalingBlackoutSettingsSlatesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelSignalingBlackoutSettingsSlatesInner{}

// ChannelSignalingBlackoutSettingsSlatesInner struct for ChannelSignalingBlackoutSettingsSlatesInner
type ChannelSignalingBlackoutSettingsSlatesInner struct {
	// Blackout slate URL to use for the specified segments. It must have one audio and one video stream. Either MPEG2 or H.264 can be used.
	BlackoutSlateUrl *string `json:"blackout_slate_url,omitempty" format:"uri-reference" doc:"Blackout slate URL to use for the specified segments. It must have one audio and one video stream. Either MPEG2 or H.264 can be used."`
	// Segment types that shall utilize the blackout slate URL. Any segment type defined here _must_ also be present in the parent signaling configuration.
	Segments []string `json:"segments,omitempty" enum:"SPLICE_INSERT,CONTENT_ID,PROGRAM,PROGRAM_BLACKOUT_OVERRIDE,PROGRAM_BREAKAWAY,CHAPTER,BREAK,OPENING_CREDIT,CLOSING_CREDIT,PROVIDER_PLACEMENT_OP,DISTRIBUTOR_PLACEMENT_OP,PROVIDER_OVERLAY_OP,DISTRIBUTOR_OVERLAY_OP,PROVIDER_AD,DISTRIBUTOR_AD,UNSCHEDULED_EVENT,NETWORK,SLATE" doc:"Segment types that shall utilize the blackout slate URL. Any segment type defined here _must_ also be present in the parent signaling configuration."`
	// Exclusive list of hex string encoded colon separated UPID Type:ID pairs (e.g. '0A:1A2B3C4D') to enable this blackout slate on.
	Upids []string `json:"upids,omitempty" doc:"Exclusive list of hex string encoded colon separated UPID Type:ID pairs (e.g. '0A:1A2B3C4D') to enable this blackout slate on."`
}

// NewChannelSignalingBlackoutSettingsSlatesInner instantiates a new ChannelSignalingBlackoutSettingsSlatesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelSignalingBlackoutSettingsSlatesInner() *ChannelSignalingBlackoutSettingsSlatesInner {
	this := ChannelSignalingBlackoutSettingsSlatesInner{}
	return &this
}

// NewChannelSignalingBlackoutSettingsSlatesInnerWithDefaults instantiates a new ChannelSignalingBlackoutSettingsSlatesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelSignalingBlackoutSettingsSlatesInnerWithDefaults() *ChannelSignalingBlackoutSettingsSlatesInner {
	this := ChannelSignalingBlackoutSettingsSlatesInner{}
	return &this
}

// GetBlackoutSlateUrl returns the BlackoutSlateUrl field value if set, zero value otherwise.
func (o *ChannelSignalingBlackoutSettingsSlatesInner) GetBlackoutSlateUrl() string {
	if o == nil || IsNil(o.BlackoutSlateUrl) {
		var ret string
		return ret
	}
	return *o.BlackoutSlateUrl
}

// GetBlackoutSlateUrlOk returns a tuple with the BlackoutSlateUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelSignalingBlackoutSettingsSlatesInner) GetBlackoutSlateUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BlackoutSlateUrl) {
		return nil, false
	}
	return o.BlackoutSlateUrl, true
}

// HasBlackoutSlateUrl returns a boolean if a field has been set.
func (o *ChannelSignalingBlackoutSettingsSlatesInner) HasBlackoutSlateUrl() bool {
	if o != nil && !IsNil(o.BlackoutSlateUrl) {
		return true
	}

	return false
}

// SetBlackoutSlateUrl gets a reference to the given string and assigns it to the BlackoutSlateUrl field.
func (o *ChannelSignalingBlackoutSettingsSlatesInner) SetBlackoutSlateUrl(v string) {
	o.BlackoutSlateUrl = &v
}

// GetSegments returns the Segments field value if set, zero value otherwise.
func (o *ChannelSignalingBlackoutSettingsSlatesInner) GetSegments() []string {
	if o == nil || IsNil(o.Segments) {
		var ret []string
		return ret
	}
	return o.Segments
}

// GetSegmentsOk returns a tuple with the Segments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelSignalingBlackoutSettingsSlatesInner) GetSegmentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Segments) {
		return nil, false
	}
	return o.Segments, true
}

// HasSegments returns a boolean if a field has been set.
func (o *ChannelSignalingBlackoutSettingsSlatesInner) HasSegments() bool {
	if o != nil && !IsNil(o.Segments) {
		return true
	}

	return false
}

// SetSegments gets a reference to the given []string and assigns it to the Segments field.
func (o *ChannelSignalingBlackoutSettingsSlatesInner) SetSegments(v []string) {
	o.Segments = v
}

// GetUpids returns the Upids field value if set, zero value otherwise.
func (o *ChannelSignalingBlackoutSettingsSlatesInner) GetUpids() []string {
	if o == nil || IsNil(o.Upids) {
		var ret []string
		return ret
	}
	return o.Upids
}

// GetUpidsOk returns a tuple with the Upids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelSignalingBlackoutSettingsSlatesInner) GetUpidsOk() ([]string, bool) {
	if o == nil || IsNil(o.Upids) {
		return nil, false
	}
	return o.Upids, true
}

// HasUpids returns a boolean if a field has been set.
func (o *ChannelSignalingBlackoutSettingsSlatesInner) HasUpids() bool {
	if o != nil && !IsNil(o.Upids) {
		return true
	}

	return false
}

// SetUpids gets a reference to the given []string and assigns it to the Upids field.
func (o *ChannelSignalingBlackoutSettingsSlatesInner) SetUpids(v []string) {
	o.Upids = v
}

func (o ChannelSignalingBlackoutSettingsSlatesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelSignalingBlackoutSettingsSlatesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlackoutSlateUrl) {
		toSerialize["blackout_slate_url"] = o.BlackoutSlateUrl
	}
	if !IsNil(o.Segments) {
		toSerialize["segments"] = o.Segments
	}
	if !IsNil(o.Upids) {
		toSerialize["upids"] = o.Upids
	}
	return toSerialize, nil
}

type NullableChannelSignalingBlackoutSettingsSlatesInner struct {
	value *ChannelSignalingBlackoutSettingsSlatesInner
	isSet bool
}

func (v NullableChannelSignalingBlackoutSettingsSlatesInner) Get() *ChannelSignalingBlackoutSettingsSlatesInner {
	return v.value
}

func (v *NullableChannelSignalingBlackoutSettingsSlatesInner) Set(val *ChannelSignalingBlackoutSettingsSlatesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelSignalingBlackoutSettingsSlatesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelSignalingBlackoutSettingsSlatesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelSignalingBlackoutSettingsSlatesInner(val *ChannelSignalingBlackoutSettingsSlatesInner) *NullableChannelSignalingBlackoutSettingsSlatesInner {
	return &NullableChannelSignalingBlackoutSettingsSlatesInner{value: val, isSet: true}
}

func (v NullableChannelSignalingBlackoutSettingsSlatesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelSignalingBlackoutSettingsSlatesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


