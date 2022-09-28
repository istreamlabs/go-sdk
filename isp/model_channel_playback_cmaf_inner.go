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

// ChannelPlaybackCmafInner struct for ChannelPlaybackCmafInner
type ChannelPlaybackCmafInner struct {
	ContentProtection *ChannelPlaybackCmafInnerContentProtection `json:"content_protection,omitempty"`
	// Partial Presentation Playback URLs
	PartialPresentations []string `json:"partial_presentations,omitempty"`
	// Playback URL
	Url string `json:"url"`
}

// NewChannelPlaybackCmafInner instantiates a new ChannelPlaybackCmafInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelPlaybackCmafInner(url string) *ChannelPlaybackCmafInner {
	this := ChannelPlaybackCmafInner{}
	this.Url = url
	return &this
}

// NewChannelPlaybackCmafInnerWithDefaults instantiates a new ChannelPlaybackCmafInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelPlaybackCmafInnerWithDefaults() *ChannelPlaybackCmafInner {
	this := ChannelPlaybackCmafInner{}
	return &this
}

// GetContentProtection returns the ContentProtection field value if set, zero value otherwise.
func (o *ChannelPlaybackCmafInner) GetContentProtection() ChannelPlaybackCmafInnerContentProtection {
	if o == nil || o.ContentProtection == nil {
		var ret ChannelPlaybackCmafInnerContentProtection
		return ret
	}
	return *o.ContentProtection
}

// GetContentProtectionOk returns a tuple with the ContentProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPlaybackCmafInner) GetContentProtectionOk() (*ChannelPlaybackCmafInnerContentProtection, bool) {
	if o == nil || o.ContentProtection == nil {
		return nil, false
	}
	return o.ContentProtection, true
}

// HasContentProtection returns a boolean if a field has been set.
func (o *ChannelPlaybackCmafInner) HasContentProtection() bool {
	if o != nil && o.ContentProtection != nil {
		return true
	}

	return false
}

// SetContentProtection gets a reference to the given ChannelPlaybackCmafInnerContentProtection and assigns it to the ContentProtection field.
func (o *ChannelPlaybackCmafInner) SetContentProtection(v ChannelPlaybackCmafInnerContentProtection) {
	o.ContentProtection = &v
}

// GetPartialPresentations returns the PartialPresentations field value if set, zero value otherwise.
func (o *ChannelPlaybackCmafInner) GetPartialPresentations() []string {
	if o == nil || o.PartialPresentations == nil {
		var ret []string
		return ret
	}
	return o.PartialPresentations
}

// GetPartialPresentationsOk returns a tuple with the PartialPresentations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPlaybackCmafInner) GetPartialPresentationsOk() ([]string, bool) {
	if o == nil || o.PartialPresentations == nil {
		return nil, false
	}
	return o.PartialPresentations, true
}

// HasPartialPresentations returns a boolean if a field has been set.
func (o *ChannelPlaybackCmafInner) HasPartialPresentations() bool {
	if o != nil && o.PartialPresentations != nil {
		return true
	}

	return false
}

// SetPartialPresentations gets a reference to the given []string and assigns it to the PartialPresentations field.
func (o *ChannelPlaybackCmafInner) SetPartialPresentations(v []string) {
	o.PartialPresentations = v
}

// GetUrl returns the Url field value
func (o *ChannelPlaybackCmafInner) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ChannelPlaybackCmafInner) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ChannelPlaybackCmafInner) SetUrl(v string) {
	o.Url = v
}

func (o ChannelPlaybackCmafInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContentProtection != nil {
		toSerialize["content_protection"] = o.ContentProtection
	}
	if o.PartialPresentations != nil {
		toSerialize["partial_presentations"] = o.PartialPresentations
	}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableChannelPlaybackCmafInner struct {
	value *ChannelPlaybackCmafInner
	isSet bool
}

func (v NullableChannelPlaybackCmafInner) Get() *ChannelPlaybackCmafInner {
	return v.value
}

func (v *NullableChannelPlaybackCmafInner) Set(val *ChannelPlaybackCmafInner) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelPlaybackCmafInner) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelPlaybackCmafInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelPlaybackCmafInner(val *ChannelPlaybackCmafInner) *NullableChannelPlaybackCmafInner {
	return &NullableChannelPlaybackCmafInner{value: val, isSet: true}
}

func (v NullableChannelPlaybackCmafInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelPlaybackCmafInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

