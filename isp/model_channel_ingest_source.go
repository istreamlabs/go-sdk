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

// ChannelIngestSource Source provides a reference to the input media stream for this channel.
type ChannelIngestSource struct {
	AudioSources []ChannelIngestSourceAudioSourcesInner `json:"audio_sources,omitempty"`
	// Closed captions source embedding. If unspecified, defaults to ATSC_A53.
	CaptionsSource *string `json:"captions_source,omitempty"`
	// Unique identifier for this source.
	Id *string `json:"id,omitempty"`
	// Self link for the source.
	Self *string `json:"self,omitempty"`
}

// NewChannelIngestSource instantiates a new ChannelIngestSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelIngestSource() *ChannelIngestSource {
	this := ChannelIngestSource{}
	return &this
}

// NewChannelIngestSourceWithDefaults instantiates a new ChannelIngestSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelIngestSourceWithDefaults() *ChannelIngestSource {
	this := ChannelIngestSource{}
	return &this
}

// GetAudioSources returns the AudioSources field value if set, zero value otherwise.
func (o *ChannelIngestSource) GetAudioSources() []ChannelIngestSourceAudioSourcesInner {
	if o == nil || o.AudioSources == nil {
		var ret []ChannelIngestSourceAudioSourcesInner
		return ret
	}
	return o.AudioSources
}

// GetAudioSourcesOk returns a tuple with the AudioSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelIngestSource) GetAudioSourcesOk() ([]ChannelIngestSourceAudioSourcesInner, bool) {
	if o == nil || o.AudioSources == nil {
		return nil, false
	}
	return o.AudioSources, true
}

// HasAudioSources returns a boolean if a field has been set.
func (o *ChannelIngestSource) HasAudioSources() bool {
	if o != nil && o.AudioSources != nil {
		return true
	}

	return false
}

// SetAudioSources gets a reference to the given []ChannelIngestSourceAudioSourcesInner and assigns it to the AudioSources field.
func (o *ChannelIngestSource) SetAudioSources(v []ChannelIngestSourceAudioSourcesInner) {
	o.AudioSources = v
}

// GetCaptionsSource returns the CaptionsSource field value if set, zero value otherwise.
func (o *ChannelIngestSource) GetCaptionsSource() string {
	if o == nil || o.CaptionsSource == nil {
		var ret string
		return ret
	}
	return *o.CaptionsSource
}

// GetCaptionsSourceOk returns a tuple with the CaptionsSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelIngestSource) GetCaptionsSourceOk() (*string, bool) {
	if o == nil || o.CaptionsSource == nil {
		return nil, false
	}
	return o.CaptionsSource, true
}

// HasCaptionsSource returns a boolean if a field has been set.
func (o *ChannelIngestSource) HasCaptionsSource() bool {
	if o != nil && o.CaptionsSource != nil {
		return true
	}

	return false
}

// SetCaptionsSource gets a reference to the given string and assigns it to the CaptionsSource field.
func (o *ChannelIngestSource) SetCaptionsSource(v string) {
	o.CaptionsSource = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ChannelIngestSource) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelIngestSource) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ChannelIngestSource) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ChannelIngestSource) SetId(v string) {
	o.Id = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ChannelIngestSource) GetSelf() string {
	if o == nil || o.Self == nil {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelIngestSource) GetSelfOk() (*string, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ChannelIngestSource) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *ChannelIngestSource) SetSelf(v string) {
	o.Self = &v
}

func (o ChannelIngestSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AudioSources != nil {
		toSerialize["audio_sources"] = o.AudioSources
	}
	if o.CaptionsSource != nil {
		toSerialize["captions_source"] = o.CaptionsSource
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	return json.Marshal(toSerialize)
}

type NullableChannelIngestSource struct {
	value *ChannelIngestSource
	isSet bool
}

func (v NullableChannelIngestSource) Get() *ChannelIngestSource {
	return v.value
}

func (v *NullableChannelIngestSource) Set(val *ChannelIngestSource) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelIngestSource) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelIngestSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelIngestSource(val *ChannelIngestSource) *NullableChannelIngestSource {
	return &NullableChannelIngestSource{value: val, isSet: true}
}

func (v NullableChannelIngestSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelIngestSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


