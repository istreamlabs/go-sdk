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

// checks if the ChannelPublishingRtmpPublicationsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelPublishingRtmpPublicationsInner{}

// ChannelPublishingRtmpPublicationsInner struct for ChannelPublishingRtmpPublicationsInner
type ChannelPublishingRtmpPublicationsInner struct {
	// Only AAC encoders are supported
	AudioEncoderId *string `json:"audio_encoder_id,omitempty" minLength:"1" doc:"Only AAC encoders are supported"`
	// RTMP publication ID. Must be unique.
	Id *string `json:"id,omitempty" minLength:"1" pattern:"/^([a-z0-9]+(-*[a-z0-9]+)*)$/" doc:"RTMP publication ID. Must be unique."`
	Url *string `json:"url,omitempty" format:"uri" minLength:"1" pattern:"/^rtmps?:\/\//"`
	// Only h264 encoders are supported
	VideoEncoderId *string `json:"video_encoder_id,omitempty" minLength:"1" doc:"Only h264 encoders are supported"`
}

// NewChannelPublishingRtmpPublicationsInner instantiates a new ChannelPublishingRtmpPublicationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelPublishingRtmpPublicationsInner() *ChannelPublishingRtmpPublicationsInner {
	this := ChannelPublishingRtmpPublicationsInner{}
	return &this
}

// NewChannelPublishingRtmpPublicationsInnerWithDefaults instantiates a new ChannelPublishingRtmpPublicationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelPublishingRtmpPublicationsInnerWithDefaults() *ChannelPublishingRtmpPublicationsInner {
	this := ChannelPublishingRtmpPublicationsInner{}
	return &this
}

// GetAudioEncoderId returns the AudioEncoderId field value if set, zero value otherwise.
func (o *ChannelPublishingRtmpPublicationsInner) GetAudioEncoderId() string {
	if o == nil || IsNil(o.AudioEncoderId) {
		var ret string
		return ret
	}
	return *o.AudioEncoderId
}

// GetAudioEncoderIdOk returns a tuple with the AudioEncoderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingRtmpPublicationsInner) GetAudioEncoderIdOk() (*string, bool) {
	if o == nil || IsNil(o.AudioEncoderId) {
		return nil, false
	}
	return o.AudioEncoderId, true
}

// HasAudioEncoderId returns a boolean if a field has been set.
func (o *ChannelPublishingRtmpPublicationsInner) HasAudioEncoderId() bool {
	if o != nil && !IsNil(o.AudioEncoderId) {
		return true
	}

	return false
}

// SetAudioEncoderId gets a reference to the given string and assigns it to the AudioEncoderId field.
func (o *ChannelPublishingRtmpPublicationsInner) SetAudioEncoderId(v string) {
	o.AudioEncoderId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ChannelPublishingRtmpPublicationsInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingRtmpPublicationsInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ChannelPublishingRtmpPublicationsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ChannelPublishingRtmpPublicationsInner) SetId(v string) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ChannelPublishingRtmpPublicationsInner) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingRtmpPublicationsInner) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ChannelPublishingRtmpPublicationsInner) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ChannelPublishingRtmpPublicationsInner) SetUrl(v string) {
	o.Url = &v
}

// GetVideoEncoderId returns the VideoEncoderId field value if set, zero value otherwise.
func (o *ChannelPublishingRtmpPublicationsInner) GetVideoEncoderId() string {
	if o == nil || IsNil(o.VideoEncoderId) {
		var ret string
		return ret
	}
	return *o.VideoEncoderId
}

// GetVideoEncoderIdOk returns a tuple with the VideoEncoderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingRtmpPublicationsInner) GetVideoEncoderIdOk() (*string, bool) {
	if o == nil || IsNil(o.VideoEncoderId) {
		return nil, false
	}
	return o.VideoEncoderId, true
}

// HasVideoEncoderId returns a boolean if a field has been set.
func (o *ChannelPublishingRtmpPublicationsInner) HasVideoEncoderId() bool {
	if o != nil && !IsNil(o.VideoEncoderId) {
		return true
	}

	return false
}

// SetVideoEncoderId gets a reference to the given string and assigns it to the VideoEncoderId field.
func (o *ChannelPublishingRtmpPublicationsInner) SetVideoEncoderId(v string) {
	o.VideoEncoderId = &v
}

func (o ChannelPublishingRtmpPublicationsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelPublishingRtmpPublicationsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AudioEncoderId) {
		toSerialize["audio_encoder_id"] = o.AudioEncoderId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.VideoEncoderId) {
		toSerialize["video_encoder_id"] = o.VideoEncoderId
	}
	return toSerialize, nil
}

type NullableChannelPublishingRtmpPublicationsInner struct {
	value *ChannelPublishingRtmpPublicationsInner
	isSet bool
}

func (v NullableChannelPublishingRtmpPublicationsInner) Get() *ChannelPublishingRtmpPublicationsInner {
	return v.value
}

func (v *NullableChannelPublishingRtmpPublicationsInner) Set(val *ChannelPublishingRtmpPublicationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelPublishingRtmpPublicationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelPublishingRtmpPublicationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelPublishingRtmpPublicationsInner(val *ChannelPublishingRtmpPublicationsInner) *NullableChannelPublishingRtmpPublicationsInner {
	return &NullableChannelPublishingRtmpPublicationsInner{value: val, isSet: true}
}

func (v NullableChannelPublishingRtmpPublicationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelPublishingRtmpPublicationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


