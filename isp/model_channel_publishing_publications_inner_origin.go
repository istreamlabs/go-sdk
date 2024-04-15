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

// checks if the ChannelPublishingPublicationsInnerOrigin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelPublishingPublicationsInnerOrigin{}

// ChannelPublishingPublicationsInnerOrigin Configures defaults used when generating manifests or playlist using the dynamic origin. Cannot be set if this is a playlist-only publication (i.e. contains publish points that specify 'playlist_only_for').
type ChannelPublishingPublicationsInnerOrigin struct {
	FallbackDefaults *ChannelPublishingPublicationsInnerOriginFallbackDefaults `json:"fallback_defaults,omitempty"`
	ManifestDefaults *ChannelPublishingPublicationsInnerOriginManifestDefaults `json:"manifest_defaults,omitempty"`
	// RetentionMinutes specifies how long data is retained, in minutes. Live linear (24/7) channels should set this to the longest expected DVR window (a few hours). Live event channels should set this to how Live2VOD playlists are expected to be available. If unspecified, the default will be 60 minutes. The maximum value is 15 days (21600 minutes).
	RetentionMinutes *int32 `json:"retention_minutes,omitempty" format:"int32" minimum:"0" maximum:"21600" doc:"RetentionMinutes specifies how long data is retained, in minutes. Live linear (24/7) channels should set this to the longest expected DVR window (a few hours). Live event channels should set this to how Live2VOD playlists are expected to be available. If unspecified, the default will be 60 minutes. The maximum value is 15 days (21600 minutes)."`
}

// NewChannelPublishingPublicationsInnerOrigin instantiates a new ChannelPublishingPublicationsInnerOrigin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelPublishingPublicationsInnerOrigin() *ChannelPublishingPublicationsInnerOrigin {
	this := ChannelPublishingPublicationsInnerOrigin{}
	return &this
}

// NewChannelPublishingPublicationsInnerOriginWithDefaults instantiates a new ChannelPublishingPublicationsInnerOrigin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelPublishingPublicationsInnerOriginWithDefaults() *ChannelPublishingPublicationsInnerOrigin {
	this := ChannelPublishingPublicationsInnerOrigin{}
	return &this
}

// GetFallbackDefaults returns the FallbackDefaults field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInnerOrigin) GetFallbackDefaults() ChannelPublishingPublicationsInnerOriginFallbackDefaults {
	if o == nil || IsNil(o.FallbackDefaults) {
		var ret ChannelPublishingPublicationsInnerOriginFallbackDefaults
		return ret
	}
	return *o.FallbackDefaults
}

// GetFallbackDefaultsOk returns a tuple with the FallbackDefaults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInnerOrigin) GetFallbackDefaultsOk() (*ChannelPublishingPublicationsInnerOriginFallbackDefaults, bool) {
	if o == nil || IsNil(o.FallbackDefaults) {
		return nil, false
	}
	return o.FallbackDefaults, true
}

// HasFallbackDefaults returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInnerOrigin) HasFallbackDefaults() bool {
	if o != nil && !IsNil(o.FallbackDefaults) {
		return true
	}

	return false
}

// SetFallbackDefaults gets a reference to the given ChannelPublishingPublicationsInnerOriginFallbackDefaults and assigns it to the FallbackDefaults field.
func (o *ChannelPublishingPublicationsInnerOrigin) SetFallbackDefaults(v ChannelPublishingPublicationsInnerOriginFallbackDefaults) {
	o.FallbackDefaults = &v
}

// GetManifestDefaults returns the ManifestDefaults field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInnerOrigin) GetManifestDefaults() ChannelPublishingPublicationsInnerOriginManifestDefaults {
	if o == nil || IsNil(o.ManifestDefaults) {
		var ret ChannelPublishingPublicationsInnerOriginManifestDefaults
		return ret
	}
	return *o.ManifestDefaults
}

// GetManifestDefaultsOk returns a tuple with the ManifestDefaults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInnerOrigin) GetManifestDefaultsOk() (*ChannelPublishingPublicationsInnerOriginManifestDefaults, bool) {
	if o == nil || IsNil(o.ManifestDefaults) {
		return nil, false
	}
	return o.ManifestDefaults, true
}

// HasManifestDefaults returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInnerOrigin) HasManifestDefaults() bool {
	if o != nil && !IsNil(o.ManifestDefaults) {
		return true
	}

	return false
}

// SetManifestDefaults gets a reference to the given ChannelPublishingPublicationsInnerOriginManifestDefaults and assigns it to the ManifestDefaults field.
func (o *ChannelPublishingPublicationsInnerOrigin) SetManifestDefaults(v ChannelPublishingPublicationsInnerOriginManifestDefaults) {
	o.ManifestDefaults = &v
}

// GetRetentionMinutes returns the RetentionMinutes field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInnerOrigin) GetRetentionMinutes() int32 {
	if o == nil || IsNil(o.RetentionMinutes) {
		var ret int32
		return ret
	}
	return *o.RetentionMinutes
}

// GetRetentionMinutesOk returns a tuple with the RetentionMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInnerOrigin) GetRetentionMinutesOk() (*int32, bool) {
	if o == nil || IsNil(o.RetentionMinutes) {
		return nil, false
	}
	return o.RetentionMinutes, true
}

// HasRetentionMinutes returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInnerOrigin) HasRetentionMinutes() bool {
	if o != nil && !IsNil(o.RetentionMinutes) {
		return true
	}

	return false
}

// SetRetentionMinutes gets a reference to the given int32 and assigns it to the RetentionMinutes field.
func (o *ChannelPublishingPublicationsInnerOrigin) SetRetentionMinutes(v int32) {
	o.RetentionMinutes = &v
}

func (o ChannelPublishingPublicationsInnerOrigin) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelPublishingPublicationsInnerOrigin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FallbackDefaults) {
		toSerialize["fallback_defaults"] = o.FallbackDefaults
	}
	if !IsNil(o.ManifestDefaults) {
		toSerialize["manifest_defaults"] = o.ManifestDefaults
	}
	if !IsNil(o.RetentionMinutes) {
		toSerialize["retention_minutes"] = o.RetentionMinutes
	}
	return toSerialize, nil
}

type NullableChannelPublishingPublicationsInnerOrigin struct {
	value *ChannelPublishingPublicationsInnerOrigin
	isSet bool
}

func (v NullableChannelPublishingPublicationsInnerOrigin) Get() *ChannelPublishingPublicationsInnerOrigin {
	return v.value
}

func (v *NullableChannelPublishingPublicationsInnerOrigin) Set(val *ChannelPublishingPublicationsInnerOrigin) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelPublishingPublicationsInnerOrigin) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelPublishingPublicationsInnerOrigin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelPublishingPublicationsInnerOrigin(val *ChannelPublishingPublicationsInnerOrigin) *NullableChannelPublishingPublicationsInnerOrigin {
	return &NullableChannelPublishingPublicationsInnerOrigin{value: val, isSet: true}
}

func (v NullableChannelPublishingPublicationsInnerOrigin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelPublishingPublicationsInnerOrigin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


