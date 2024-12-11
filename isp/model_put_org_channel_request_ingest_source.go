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

// checks if the PutOrgChannelRequestIngestSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutOrgChannelRequestIngestSource{}

// PutOrgChannelRequestIngestSource Source provides a reference to the input media stream for this channel.
type PutOrgChannelRequestIngestSource struct {
	AudioSources []ChannelIngestSourceAudioSourcesInner `json:"audio_sources,omitempty" minItems:"1"`
	// Closed captions source embedding. If unspecified, defaults to ATSC_A53.
	CaptionsSource *string `json:"captions_source,omitempty" enum:"ATSC_A53,SMPTE_2038" doc:"Closed captions source embedding. If unspecified, defaults to ATSC_A53."`
	// Unique identifier for this source.
	Id string `json:"id" doc:"Unique identifier for this source."`
	// Self link for the source.
	Self        *string                         `json:"self,omitempty" format:"uri-reference" doc:"Self link for the source."`
	VideoSource *ChannelIngestSourceVideoSource `json:"video_source,omitempty"`
}

// NewPutOrgChannelRequestIngestSource instantiates a new PutOrgChannelRequestIngestSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutOrgChannelRequestIngestSource(id string) *PutOrgChannelRequestIngestSource {
	this := PutOrgChannelRequestIngestSource{}
	this.Id = id
	return &this
}

// NewPutOrgChannelRequestIngestSourceWithDefaults instantiates a new PutOrgChannelRequestIngestSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutOrgChannelRequestIngestSourceWithDefaults() *PutOrgChannelRequestIngestSource {
	this := PutOrgChannelRequestIngestSource{}
	return &this
}

// GetAudioSources returns the AudioSources field value if set, zero value otherwise.
func (o *PutOrgChannelRequestIngestSource) GetAudioSources() []ChannelIngestSourceAudioSourcesInner {
	if o == nil || IsNil(o.AudioSources) {
		var ret []ChannelIngestSourceAudioSourcesInner
		return ret
	}
	return o.AudioSources
}

// GetAudioSourcesOk returns a tuple with the AudioSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutOrgChannelRequestIngestSource) GetAudioSourcesOk() ([]ChannelIngestSourceAudioSourcesInner, bool) {
	if o == nil || IsNil(o.AudioSources) {
		return nil, false
	}
	return o.AudioSources, true
}

// HasAudioSources returns a boolean if a field has been set.
func (o *PutOrgChannelRequestIngestSource) HasAudioSources() bool {
	if o != nil && !IsNil(o.AudioSources) {
		return true
	}

	return false
}

// SetAudioSources gets a reference to the given []ChannelIngestSourceAudioSourcesInner and assigns it to the AudioSources field.
func (o *PutOrgChannelRequestIngestSource) SetAudioSources(v []ChannelIngestSourceAudioSourcesInner) {
	o.AudioSources = v
}

// GetCaptionsSource returns the CaptionsSource field value if set, zero value otherwise.
func (o *PutOrgChannelRequestIngestSource) GetCaptionsSource() string {
	if o == nil || IsNil(o.CaptionsSource) {
		var ret string
		return ret
	}
	return *o.CaptionsSource
}

// GetCaptionsSourceOk returns a tuple with the CaptionsSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutOrgChannelRequestIngestSource) GetCaptionsSourceOk() (*string, bool) {
	if o == nil || IsNil(o.CaptionsSource) {
		return nil, false
	}
	return o.CaptionsSource, true
}

// HasCaptionsSource returns a boolean if a field has been set.
func (o *PutOrgChannelRequestIngestSource) HasCaptionsSource() bool {
	if o != nil && !IsNil(o.CaptionsSource) {
		return true
	}

	return false
}

// SetCaptionsSource gets a reference to the given string and assigns it to the CaptionsSource field.
func (o *PutOrgChannelRequestIngestSource) SetCaptionsSource(v string) {
	o.CaptionsSource = &v
}

// GetId returns the Id field value
func (o *PutOrgChannelRequestIngestSource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PutOrgChannelRequestIngestSource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PutOrgChannelRequestIngestSource) SetId(v string) {
	o.Id = v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *PutOrgChannelRequestIngestSource) GetSelf() string {
	if o == nil || IsNil(o.Self) {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutOrgChannelRequestIngestSource) GetSelfOk() (*string, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *PutOrgChannelRequestIngestSource) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *PutOrgChannelRequestIngestSource) SetSelf(v string) {
	o.Self = &v
}

// GetVideoSource returns the VideoSource field value if set, zero value otherwise.
func (o *PutOrgChannelRequestIngestSource) GetVideoSource() ChannelIngestSourceVideoSource {
	if o == nil || IsNil(o.VideoSource) {
		var ret ChannelIngestSourceVideoSource
		return ret
	}
	return *o.VideoSource
}

// GetVideoSourceOk returns a tuple with the VideoSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutOrgChannelRequestIngestSource) GetVideoSourceOk() (*ChannelIngestSourceVideoSource, bool) {
	if o == nil || IsNil(o.VideoSource) {
		return nil, false
	}
	return o.VideoSource, true
}

// HasVideoSource returns a boolean if a field has been set.
func (o *PutOrgChannelRequestIngestSource) HasVideoSource() bool {
	if o != nil && !IsNil(o.VideoSource) {
		return true
	}

	return false
}

// SetVideoSource gets a reference to the given ChannelIngestSourceVideoSource and assigns it to the VideoSource field.
func (o *PutOrgChannelRequestIngestSource) SetVideoSource(v ChannelIngestSourceVideoSource) {
	o.VideoSource = &v
}

func (o PutOrgChannelRequestIngestSource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutOrgChannelRequestIngestSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AudioSources) {
		toSerialize["audio_sources"] = o.AudioSources
	}
	if !IsNil(o.CaptionsSource) {
		toSerialize["captions_source"] = o.CaptionsSource
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.VideoSource) {
		toSerialize["video_source"] = o.VideoSource
	}
	return toSerialize, nil
}

type NullablePutOrgChannelRequestIngestSource struct {
	value *PutOrgChannelRequestIngestSource
	isSet bool
}

func (v NullablePutOrgChannelRequestIngestSource) Get() *PutOrgChannelRequestIngestSource {
	return v.value
}

func (v *NullablePutOrgChannelRequestIngestSource) Set(val *PutOrgChannelRequestIngestSource) {
	v.value = val
	v.isSet = true
}

func (v NullablePutOrgChannelRequestIngestSource) IsSet() bool {
	return v.isSet
}

func (v *NullablePutOrgChannelRequestIngestSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutOrgChannelRequestIngestSource(val *PutOrgChannelRequestIngestSource) *NullablePutOrgChannelRequestIngestSource {
	return &NullablePutOrgChannelRequestIngestSource{value: val, isSet: true}
}

func (v NullablePutOrgChannelRequestIngestSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutOrgChannelRequestIngestSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
