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

// ChannelPublishing Publishing configures playlist formats and where to send video and playlist data.
type ChannelPublishing struct {
	// Configures how captioning information is published.
	ClosedCaptionStreams *[]ChannelPublishingClosedCaptionStreams `json:"closed_caption_streams,omitempty"`
	// Set of string identifiers corresponding to features that this Channel is opting in.
	FeatureFlags *[]string `json:"feature_flags,omitempty"`
	Live2vod *ChannelPublishingLive2vod `json:"live2vod,omitempty"`
	// A set of individual configurations that each can configure a specific destination and mechanism of delivery for segments and/or playlists.
	Publications *[]ChannelPublishingPublications `json:"publications,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChannelPublishing ChannelPublishing

// NewChannelPublishing instantiates a new ChannelPublishing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelPublishing() *ChannelPublishing {
	this := ChannelPublishing{}
	return &this
}

// NewChannelPublishingWithDefaults instantiates a new ChannelPublishing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelPublishingWithDefaults() *ChannelPublishing {
	this := ChannelPublishing{}
	return &this
}

// GetClosedCaptionStreams returns the ClosedCaptionStreams field value if set, zero value otherwise.
func (o *ChannelPublishing) GetClosedCaptionStreams() []ChannelPublishingClosedCaptionStreams {
	if o == nil || o.ClosedCaptionStreams == nil {
		var ret []ChannelPublishingClosedCaptionStreams
		return ret
	}
	return *o.ClosedCaptionStreams
}

// GetClosedCaptionStreamsOk returns a tuple with the ClosedCaptionStreams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishing) GetClosedCaptionStreamsOk() (*[]ChannelPublishingClosedCaptionStreams, bool) {
	if o == nil || o.ClosedCaptionStreams == nil {
		return nil, false
	}
	return o.ClosedCaptionStreams, true
}

// HasClosedCaptionStreams returns a boolean if a field has been set.
func (o *ChannelPublishing) HasClosedCaptionStreams() bool {
	if o != nil && o.ClosedCaptionStreams != nil {
		return true
	}

	return false
}

// SetClosedCaptionStreams gets a reference to the given []ChannelPublishingClosedCaptionStreams and assigns it to the ClosedCaptionStreams field.
func (o *ChannelPublishing) SetClosedCaptionStreams(v []ChannelPublishingClosedCaptionStreams) {
	o.ClosedCaptionStreams = &v
}

// GetFeatureFlags returns the FeatureFlags field value if set, zero value otherwise.
func (o *ChannelPublishing) GetFeatureFlags() []string {
	if o == nil || o.FeatureFlags == nil {
		var ret []string
		return ret
	}
	return *o.FeatureFlags
}

// GetFeatureFlagsOk returns a tuple with the FeatureFlags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishing) GetFeatureFlagsOk() (*[]string, bool) {
	if o == nil || o.FeatureFlags == nil {
		return nil, false
	}
	return o.FeatureFlags, true
}

// HasFeatureFlags returns a boolean if a field has been set.
func (o *ChannelPublishing) HasFeatureFlags() bool {
	if o != nil && o.FeatureFlags != nil {
		return true
	}

	return false
}

// SetFeatureFlags gets a reference to the given []string and assigns it to the FeatureFlags field.
func (o *ChannelPublishing) SetFeatureFlags(v []string) {
	o.FeatureFlags = &v
}

// GetLive2vod returns the Live2vod field value if set, zero value otherwise.
func (o *ChannelPublishing) GetLive2vod() ChannelPublishingLive2vod {
	if o == nil || o.Live2vod == nil {
		var ret ChannelPublishingLive2vod
		return ret
	}
	return *o.Live2vod
}

// GetLive2vodOk returns a tuple with the Live2vod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishing) GetLive2vodOk() (*ChannelPublishingLive2vod, bool) {
	if o == nil || o.Live2vod == nil {
		return nil, false
	}
	return o.Live2vod, true
}

// HasLive2vod returns a boolean if a field has been set.
func (o *ChannelPublishing) HasLive2vod() bool {
	if o != nil && o.Live2vod != nil {
		return true
	}

	return false
}

// SetLive2vod gets a reference to the given ChannelPublishingLive2vod and assigns it to the Live2vod field.
func (o *ChannelPublishing) SetLive2vod(v ChannelPublishingLive2vod) {
	o.Live2vod = &v
}

// GetPublications returns the Publications field value if set, zero value otherwise.
func (o *ChannelPublishing) GetPublications() []ChannelPublishingPublications {
	if o == nil || o.Publications == nil {
		var ret []ChannelPublishingPublications
		return ret
	}
	return *o.Publications
}

// GetPublicationsOk returns a tuple with the Publications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishing) GetPublicationsOk() (*[]ChannelPublishingPublications, bool) {
	if o == nil || o.Publications == nil {
		return nil, false
	}
	return o.Publications, true
}

// HasPublications returns a boolean if a field has been set.
func (o *ChannelPublishing) HasPublications() bool {
	if o != nil && o.Publications != nil {
		return true
	}

	return false
}

// SetPublications gets a reference to the given []ChannelPublishingPublications and assigns it to the Publications field.
func (o *ChannelPublishing) SetPublications(v []ChannelPublishingPublications) {
	o.Publications = &v
}

func (o ChannelPublishing) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClosedCaptionStreams != nil {
		toSerialize["closed_caption_streams"] = o.ClosedCaptionStreams
	}
	if o.FeatureFlags != nil {
		toSerialize["feature_flags"] = o.FeatureFlags
	}
	if o.Live2vod != nil {
		toSerialize["live2vod"] = o.Live2vod
	}
	if o.Publications != nil {
		toSerialize["publications"] = o.Publications
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ChannelPublishing) UnmarshalJSON(bytes []byte) (err error) {
	varChannelPublishing := _ChannelPublishing{}

	if err = json.Unmarshal(bytes, &varChannelPublishing); err == nil {
		*o = ChannelPublishing(varChannelPublishing)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "closed_caption_streams")
		delete(additionalProperties, "feature_flags")
		delete(additionalProperties, "live2vod")
		delete(additionalProperties, "publications")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChannelPublishing struct {
	value *ChannelPublishing
	isSet bool
}

func (v NullableChannelPublishing) Get() *ChannelPublishing {
	return v.value
}

func (v *NullableChannelPublishing) Set(val *ChannelPublishing) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelPublishing) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelPublishing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelPublishing(val *ChannelPublishing) *NullableChannelPublishing {
	return &NullableChannelPublishing{value: val, isSet: true}
}

func (v NullableChannelPublishing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelPublishing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


