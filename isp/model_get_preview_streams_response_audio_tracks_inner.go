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

// GetPreviewStreamsResponseAudioTracksInner struct for GetPreviewStreamsResponseAudioTracksInner
type GetPreviewStreamsResponseAudioTracksInner struct {
	Bitrate *int32 `json:"bitrate,omitempty"`
	// codec_string will be set to the RFC 6381 compliant string that represents the specific codec in this AudioTrackInfo. ex. 'mp4a.40.2' for AAC Low-Complexity.
	CodecString *string `json:"codec_string,omitempty"`
	Id *int32 `json:"id,omitempty"`
	NumChannels *int32 `json:"num_channels,omitempty"`
	SampleRate *int32 `json:"sample_rate,omitempty"`
}

// NewGetPreviewStreamsResponseAudioTracksInner instantiates a new GetPreviewStreamsResponseAudioTracksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPreviewStreamsResponseAudioTracksInner() *GetPreviewStreamsResponseAudioTracksInner {
	this := GetPreviewStreamsResponseAudioTracksInner{}
	return &this
}

// NewGetPreviewStreamsResponseAudioTracksInnerWithDefaults instantiates a new GetPreviewStreamsResponseAudioTracksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPreviewStreamsResponseAudioTracksInnerWithDefaults() *GetPreviewStreamsResponseAudioTracksInner {
	this := GetPreviewStreamsResponseAudioTracksInner{}
	return &this
}

// GetBitrate returns the Bitrate field value if set, zero value otherwise.
func (o *GetPreviewStreamsResponseAudioTracksInner) GetBitrate() int32 {
	if o == nil || o.Bitrate == nil {
		var ret int32
		return ret
	}
	return *o.Bitrate
}

// GetBitrateOk returns a tuple with the Bitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPreviewStreamsResponseAudioTracksInner) GetBitrateOk() (*int32, bool) {
	if o == nil || o.Bitrate == nil {
		return nil, false
	}
	return o.Bitrate, true
}

// HasBitrate returns a boolean if a field has been set.
func (o *GetPreviewStreamsResponseAudioTracksInner) HasBitrate() bool {
	if o != nil && o.Bitrate != nil {
		return true
	}

	return false
}

// SetBitrate gets a reference to the given int32 and assigns it to the Bitrate field.
func (o *GetPreviewStreamsResponseAudioTracksInner) SetBitrate(v int32) {
	o.Bitrate = &v
}

// GetCodecString returns the CodecString field value if set, zero value otherwise.
func (o *GetPreviewStreamsResponseAudioTracksInner) GetCodecString() string {
	if o == nil || o.CodecString == nil {
		var ret string
		return ret
	}
	return *o.CodecString
}

// GetCodecStringOk returns a tuple with the CodecString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPreviewStreamsResponseAudioTracksInner) GetCodecStringOk() (*string, bool) {
	if o == nil || o.CodecString == nil {
		return nil, false
	}
	return o.CodecString, true
}

// HasCodecString returns a boolean if a field has been set.
func (o *GetPreviewStreamsResponseAudioTracksInner) HasCodecString() bool {
	if o != nil && o.CodecString != nil {
		return true
	}

	return false
}

// SetCodecString gets a reference to the given string and assigns it to the CodecString field.
func (o *GetPreviewStreamsResponseAudioTracksInner) SetCodecString(v string) {
	o.CodecString = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetPreviewStreamsResponseAudioTracksInner) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPreviewStreamsResponseAudioTracksInner) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetPreviewStreamsResponseAudioTracksInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetPreviewStreamsResponseAudioTracksInner) SetId(v int32) {
	o.Id = &v
}

// GetNumChannels returns the NumChannels field value if set, zero value otherwise.
func (o *GetPreviewStreamsResponseAudioTracksInner) GetNumChannels() int32 {
	if o == nil || o.NumChannels == nil {
		var ret int32
		return ret
	}
	return *o.NumChannels
}

// GetNumChannelsOk returns a tuple with the NumChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPreviewStreamsResponseAudioTracksInner) GetNumChannelsOk() (*int32, bool) {
	if o == nil || o.NumChannels == nil {
		return nil, false
	}
	return o.NumChannels, true
}

// HasNumChannels returns a boolean if a field has been set.
func (o *GetPreviewStreamsResponseAudioTracksInner) HasNumChannels() bool {
	if o != nil && o.NumChannels != nil {
		return true
	}

	return false
}

// SetNumChannels gets a reference to the given int32 and assigns it to the NumChannels field.
func (o *GetPreviewStreamsResponseAudioTracksInner) SetNumChannels(v int32) {
	o.NumChannels = &v
}

// GetSampleRate returns the SampleRate field value if set, zero value otherwise.
func (o *GetPreviewStreamsResponseAudioTracksInner) GetSampleRate() int32 {
	if o == nil || o.SampleRate == nil {
		var ret int32
		return ret
	}
	return *o.SampleRate
}

// GetSampleRateOk returns a tuple with the SampleRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPreviewStreamsResponseAudioTracksInner) GetSampleRateOk() (*int32, bool) {
	if o == nil || o.SampleRate == nil {
		return nil, false
	}
	return o.SampleRate, true
}

// HasSampleRate returns a boolean if a field has been set.
func (o *GetPreviewStreamsResponseAudioTracksInner) HasSampleRate() bool {
	if o != nil && o.SampleRate != nil {
		return true
	}

	return false
}

// SetSampleRate gets a reference to the given int32 and assigns it to the SampleRate field.
func (o *GetPreviewStreamsResponseAudioTracksInner) SetSampleRate(v int32) {
	o.SampleRate = &v
}

func (o GetPreviewStreamsResponseAudioTracksInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bitrate != nil {
		toSerialize["bitrate"] = o.Bitrate
	}
	if o.CodecString != nil {
		toSerialize["codec_string"] = o.CodecString
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.NumChannels != nil {
		toSerialize["num_channels"] = o.NumChannels
	}
	if o.SampleRate != nil {
		toSerialize["sample_rate"] = o.SampleRate
	}
	return json.Marshal(toSerialize)
}

type NullableGetPreviewStreamsResponseAudioTracksInner struct {
	value *GetPreviewStreamsResponseAudioTracksInner
	isSet bool
}

func (v NullableGetPreviewStreamsResponseAudioTracksInner) Get() *GetPreviewStreamsResponseAudioTracksInner {
	return v.value
}

func (v *NullableGetPreviewStreamsResponseAudioTracksInner) Set(val *GetPreviewStreamsResponseAudioTracksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPreviewStreamsResponseAudioTracksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPreviewStreamsResponseAudioTracksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPreviewStreamsResponseAudioTracksInner(val *GetPreviewStreamsResponseAudioTracksInner) *NullableGetPreviewStreamsResponseAudioTracksInner {
	return &NullableGetPreviewStreamsResponseAudioTracksInner{value: val, isSet: true}
}

func (v NullableGetPreviewStreamsResponseAudioTracksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPreviewStreamsResponseAudioTracksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

