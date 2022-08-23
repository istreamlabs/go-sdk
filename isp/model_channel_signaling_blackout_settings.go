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

// ChannelSignalingBlackoutSettings Configure blackout: replacing content with custom slates based on program signaling.
type ChannelSignalingBlackoutSettings struct {
	// Default slate URL to use for blackouts. Can be overridden by the 'slates' field.
	DefaultBlackoutSlateUrl *string `json:"default_blackout_slate_url,omitempty"`
	// List of signaling segment types to force blackout, e.g. add 'SPLICE_INSERT' to blackout all ads signaled via SCTE-35 splice_insert.
	ForceBlackoutSegments []string `json:"force_blackout_segments,omitempty"`
	// Determines whether to honor the web_delivery_allowed attribute in SCTE-35 segmentation descriptors. When this is enabled, a segmentation descriptor with web_delivery_allowed=false will trigger a blackout.
	HonorWebDeliveryRestriction *bool `json:"honor_web_delivery_restriction,omitempty"`
	// Per-segment type slate overrides.
	Slates []ChannelSignalingBlackoutSettingsSlatesInner `json:"slates,omitempty"`
}

// NewChannelSignalingBlackoutSettings instantiates a new ChannelSignalingBlackoutSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelSignalingBlackoutSettings() *ChannelSignalingBlackoutSettings {
	this := ChannelSignalingBlackoutSettings{}
	return &this
}

// NewChannelSignalingBlackoutSettingsWithDefaults instantiates a new ChannelSignalingBlackoutSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelSignalingBlackoutSettingsWithDefaults() *ChannelSignalingBlackoutSettings {
	this := ChannelSignalingBlackoutSettings{}
	return &this
}

// GetDefaultBlackoutSlateUrl returns the DefaultBlackoutSlateUrl field value if set, zero value otherwise.
func (o *ChannelSignalingBlackoutSettings) GetDefaultBlackoutSlateUrl() string {
	if o == nil || o.DefaultBlackoutSlateUrl == nil {
		var ret string
		return ret
	}
	return *o.DefaultBlackoutSlateUrl
}

// GetDefaultBlackoutSlateUrlOk returns a tuple with the DefaultBlackoutSlateUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelSignalingBlackoutSettings) GetDefaultBlackoutSlateUrlOk() (*string, bool) {
	if o == nil || o.DefaultBlackoutSlateUrl == nil {
		return nil, false
	}
	return o.DefaultBlackoutSlateUrl, true
}

// HasDefaultBlackoutSlateUrl returns a boolean if a field has been set.
func (o *ChannelSignalingBlackoutSettings) HasDefaultBlackoutSlateUrl() bool {
	if o != nil && o.DefaultBlackoutSlateUrl != nil {
		return true
	}

	return false
}

// SetDefaultBlackoutSlateUrl gets a reference to the given string and assigns it to the DefaultBlackoutSlateUrl field.
func (o *ChannelSignalingBlackoutSettings) SetDefaultBlackoutSlateUrl(v string) {
	o.DefaultBlackoutSlateUrl = &v
}

// GetForceBlackoutSegments returns the ForceBlackoutSegments field value if set, zero value otherwise.
func (o *ChannelSignalingBlackoutSettings) GetForceBlackoutSegments() []string {
	if o == nil || o.ForceBlackoutSegments == nil {
		var ret []string
		return ret
	}
	return o.ForceBlackoutSegments
}

// GetForceBlackoutSegmentsOk returns a tuple with the ForceBlackoutSegments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelSignalingBlackoutSettings) GetForceBlackoutSegmentsOk() ([]string, bool) {
	if o == nil || o.ForceBlackoutSegments == nil {
		return nil, false
	}
	return o.ForceBlackoutSegments, true
}

// HasForceBlackoutSegments returns a boolean if a field has been set.
func (o *ChannelSignalingBlackoutSettings) HasForceBlackoutSegments() bool {
	if o != nil && o.ForceBlackoutSegments != nil {
		return true
	}

	return false
}

// SetForceBlackoutSegments gets a reference to the given []string and assigns it to the ForceBlackoutSegments field.
func (o *ChannelSignalingBlackoutSettings) SetForceBlackoutSegments(v []string) {
	o.ForceBlackoutSegments = v
}

// GetHonorWebDeliveryRestriction returns the HonorWebDeliveryRestriction field value if set, zero value otherwise.
func (o *ChannelSignalingBlackoutSettings) GetHonorWebDeliveryRestriction() bool {
	if o == nil || o.HonorWebDeliveryRestriction == nil {
		var ret bool
		return ret
	}
	return *o.HonorWebDeliveryRestriction
}

// GetHonorWebDeliveryRestrictionOk returns a tuple with the HonorWebDeliveryRestriction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelSignalingBlackoutSettings) GetHonorWebDeliveryRestrictionOk() (*bool, bool) {
	if o == nil || o.HonorWebDeliveryRestriction == nil {
		return nil, false
	}
	return o.HonorWebDeliveryRestriction, true
}

// HasHonorWebDeliveryRestriction returns a boolean if a field has been set.
func (o *ChannelSignalingBlackoutSettings) HasHonorWebDeliveryRestriction() bool {
	if o != nil && o.HonorWebDeliveryRestriction != nil {
		return true
	}

	return false
}

// SetHonorWebDeliveryRestriction gets a reference to the given bool and assigns it to the HonorWebDeliveryRestriction field.
func (o *ChannelSignalingBlackoutSettings) SetHonorWebDeliveryRestriction(v bool) {
	o.HonorWebDeliveryRestriction = &v
}

// GetSlates returns the Slates field value if set, zero value otherwise.
func (o *ChannelSignalingBlackoutSettings) GetSlates() []ChannelSignalingBlackoutSettingsSlatesInner {
	if o == nil || o.Slates == nil {
		var ret []ChannelSignalingBlackoutSettingsSlatesInner
		return ret
	}
	return o.Slates
}

// GetSlatesOk returns a tuple with the Slates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelSignalingBlackoutSettings) GetSlatesOk() ([]ChannelSignalingBlackoutSettingsSlatesInner, bool) {
	if o == nil || o.Slates == nil {
		return nil, false
	}
	return o.Slates, true
}

// HasSlates returns a boolean if a field has been set.
func (o *ChannelSignalingBlackoutSettings) HasSlates() bool {
	if o != nil && o.Slates != nil {
		return true
	}

	return false
}

// SetSlates gets a reference to the given []ChannelSignalingBlackoutSettingsSlatesInner and assigns it to the Slates field.
func (o *ChannelSignalingBlackoutSettings) SetSlates(v []ChannelSignalingBlackoutSettingsSlatesInner) {
	o.Slates = v
}

func (o ChannelSignalingBlackoutSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultBlackoutSlateUrl != nil {
		toSerialize["default_blackout_slate_url"] = o.DefaultBlackoutSlateUrl
	}
	if o.ForceBlackoutSegments != nil {
		toSerialize["force_blackout_segments"] = o.ForceBlackoutSegments
	}
	if o.HonorWebDeliveryRestriction != nil {
		toSerialize["honor_web_delivery_restriction"] = o.HonorWebDeliveryRestriction
	}
	if o.Slates != nil {
		toSerialize["slates"] = o.Slates
	}
	return json.Marshal(toSerialize)
}

type NullableChannelSignalingBlackoutSettings struct {
	value *ChannelSignalingBlackoutSettings
	isSet bool
}

func (v NullableChannelSignalingBlackoutSettings) Get() *ChannelSignalingBlackoutSettings {
	return v.value
}

func (v *NullableChannelSignalingBlackoutSettings) Set(val *ChannelSignalingBlackoutSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelSignalingBlackoutSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelSignalingBlackoutSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelSignalingBlackoutSettings(val *ChannelSignalingBlackoutSettings) *NullableChannelSignalingBlackoutSettings {
	return &NullableChannelSignalingBlackoutSettings{value: val, isSet: true}
}

func (v NullableChannelSignalingBlackoutSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelSignalingBlackoutSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


