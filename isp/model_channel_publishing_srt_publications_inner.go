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

// checks if the ChannelPublishingSrtPublicationsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelPublishingSrtPublicationsInner{}

// ChannelPublishingSrtPublicationsInner struct for ChannelPublishingSrtPublicationsInner
type ChannelPublishingSrtPublicationsInner struct {
	AudioEncoders []ChannelPublishingSrtPublicationsInnerAudioEncodersInner `json:"audio_encoders,omitempty" minItems:"1"`
	// SRT publication ID. Must be unique.
	Id *string `json:"id,omitempty" validate:"regexp=^([a-z0-9]+(-*[a-z0-9]+)*)$" minLength:"1" pattern:"/^([a-z0-9]+(-*[a-z0-9]+)*)$/" doc:"SRT publication ID. Must be unique."`
	// MPEG-TS PMT PID. PIDs should be set on the PMT, SCTE-35 and all encoders or none. Valid PIDs must 13-bit values greater than 31. If no PIDs are provided (pid == 0) then they will be generated automatically.
	PmtPid *int32 `json:"pmt_pid,omitempty" format:"int32" exclusiveMaximum:"8191" doc:"MPEG-TS PMT PID. PIDs should be set on the PMT, SCTE-35 and all encoders or none. Valid PIDs must 13-bit values greater than 31. If no PIDs are provided (pid == 0) then they will be generated automatically."`
	// MPEG-TS SCTE-35 PID. PIDs should be set on the PMT, SCTE-35, and all encoders or none. Valid PIDs must 13-bit values greater than 31. If no PIDs are provided (pid == 0) then they will be generated automatically.
	Scte35Pid     *int32                                                    `json:"scte35_pid,omitempty" format:"int32" exclusiveMaximum:"8191" doc:"MPEG-TS SCTE-35 PID. PIDs should be set on the PMT, SCTE-35, and all encoders or none. Valid PIDs must 13-bit values greater than 31. If no PIDs are provided (pid == 0) then they will be generated automatically."`
	Url           *string                                                   `json:"url,omitempty" validate:"regexp=^srt:\\/\\/" format:"uri" minLength:"1" pattern:"/^srt:\/\//"`
	VideoEncoders []ChannelPublishingSrtPublicationsInnerAudioEncodersInner `json:"video_encoders,omitempty" minItems:"1"`
}

// NewChannelPublishingSrtPublicationsInner instantiates a new ChannelPublishingSrtPublicationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelPublishingSrtPublicationsInner() *ChannelPublishingSrtPublicationsInner {
	this := ChannelPublishingSrtPublicationsInner{}
	return &this
}

// NewChannelPublishingSrtPublicationsInnerWithDefaults instantiates a new ChannelPublishingSrtPublicationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelPublishingSrtPublicationsInnerWithDefaults() *ChannelPublishingSrtPublicationsInner {
	this := ChannelPublishingSrtPublicationsInner{}
	return &this
}

// GetAudioEncoders returns the AudioEncoders field value if set, zero value otherwise.
func (o *ChannelPublishingSrtPublicationsInner) GetAudioEncoders() []ChannelPublishingSrtPublicationsInnerAudioEncodersInner {
	if o == nil || IsNil(o.AudioEncoders) {
		var ret []ChannelPublishingSrtPublicationsInnerAudioEncodersInner
		return ret
	}
	return o.AudioEncoders
}

// GetAudioEncodersOk returns a tuple with the AudioEncoders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingSrtPublicationsInner) GetAudioEncodersOk() ([]ChannelPublishingSrtPublicationsInnerAudioEncodersInner, bool) {
	if o == nil || IsNil(o.AudioEncoders) {
		return nil, false
	}
	return o.AudioEncoders, true
}

// HasAudioEncoders returns a boolean if a field has been set.
func (o *ChannelPublishingSrtPublicationsInner) HasAudioEncoders() bool {
	if o != nil && !IsNil(o.AudioEncoders) {
		return true
	}

	return false
}

// SetAudioEncoders gets a reference to the given []ChannelPublishingSrtPublicationsInnerAudioEncodersInner and assigns it to the AudioEncoders field.
func (o *ChannelPublishingSrtPublicationsInner) SetAudioEncoders(v []ChannelPublishingSrtPublicationsInnerAudioEncodersInner) {
	o.AudioEncoders = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ChannelPublishingSrtPublicationsInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingSrtPublicationsInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ChannelPublishingSrtPublicationsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ChannelPublishingSrtPublicationsInner) SetId(v string) {
	o.Id = &v
}

// GetPmtPid returns the PmtPid field value if set, zero value otherwise.
func (o *ChannelPublishingSrtPublicationsInner) GetPmtPid() int32 {
	if o == nil || IsNil(o.PmtPid) {
		var ret int32
		return ret
	}
	return *o.PmtPid
}

// GetPmtPidOk returns a tuple with the PmtPid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingSrtPublicationsInner) GetPmtPidOk() (*int32, bool) {
	if o == nil || IsNil(o.PmtPid) {
		return nil, false
	}
	return o.PmtPid, true
}

// HasPmtPid returns a boolean if a field has been set.
func (o *ChannelPublishingSrtPublicationsInner) HasPmtPid() bool {
	if o != nil && !IsNil(o.PmtPid) {
		return true
	}

	return false
}

// SetPmtPid gets a reference to the given int32 and assigns it to the PmtPid field.
func (o *ChannelPublishingSrtPublicationsInner) SetPmtPid(v int32) {
	o.PmtPid = &v
}

// GetScte35Pid returns the Scte35Pid field value if set, zero value otherwise.
func (o *ChannelPublishingSrtPublicationsInner) GetScte35Pid() int32 {
	if o == nil || IsNil(o.Scte35Pid) {
		var ret int32
		return ret
	}
	return *o.Scte35Pid
}

// GetScte35PidOk returns a tuple with the Scte35Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingSrtPublicationsInner) GetScte35PidOk() (*int32, bool) {
	if o == nil || IsNil(o.Scte35Pid) {
		return nil, false
	}
	return o.Scte35Pid, true
}

// HasScte35Pid returns a boolean if a field has been set.
func (o *ChannelPublishingSrtPublicationsInner) HasScte35Pid() bool {
	if o != nil && !IsNil(o.Scte35Pid) {
		return true
	}

	return false
}

// SetScte35Pid gets a reference to the given int32 and assigns it to the Scte35Pid field.
func (o *ChannelPublishingSrtPublicationsInner) SetScte35Pid(v int32) {
	o.Scte35Pid = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ChannelPublishingSrtPublicationsInner) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingSrtPublicationsInner) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ChannelPublishingSrtPublicationsInner) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ChannelPublishingSrtPublicationsInner) SetUrl(v string) {
	o.Url = &v
}

// GetVideoEncoders returns the VideoEncoders field value if set, zero value otherwise.
func (o *ChannelPublishingSrtPublicationsInner) GetVideoEncoders() []ChannelPublishingSrtPublicationsInnerAudioEncodersInner {
	if o == nil || IsNil(o.VideoEncoders) {
		var ret []ChannelPublishingSrtPublicationsInnerAudioEncodersInner
		return ret
	}
	return o.VideoEncoders
}

// GetVideoEncodersOk returns a tuple with the VideoEncoders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingSrtPublicationsInner) GetVideoEncodersOk() ([]ChannelPublishingSrtPublicationsInnerAudioEncodersInner, bool) {
	if o == nil || IsNil(o.VideoEncoders) {
		return nil, false
	}
	return o.VideoEncoders, true
}

// HasVideoEncoders returns a boolean if a field has been set.
func (o *ChannelPublishingSrtPublicationsInner) HasVideoEncoders() bool {
	if o != nil && !IsNil(o.VideoEncoders) {
		return true
	}

	return false
}

// SetVideoEncoders gets a reference to the given []ChannelPublishingSrtPublicationsInnerAudioEncodersInner and assigns it to the VideoEncoders field.
func (o *ChannelPublishingSrtPublicationsInner) SetVideoEncoders(v []ChannelPublishingSrtPublicationsInnerAudioEncodersInner) {
	o.VideoEncoders = v
}

func (o ChannelPublishingSrtPublicationsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelPublishingSrtPublicationsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AudioEncoders) {
		toSerialize["audio_encoders"] = o.AudioEncoders
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.PmtPid) {
		toSerialize["pmt_pid"] = o.PmtPid
	}
	if !IsNil(o.Scte35Pid) {
		toSerialize["scte35_pid"] = o.Scte35Pid
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.VideoEncoders) {
		toSerialize["video_encoders"] = o.VideoEncoders
	}
	return toSerialize, nil
}

type NullableChannelPublishingSrtPublicationsInner struct {
	value *ChannelPublishingSrtPublicationsInner
	isSet bool
}

func (v NullableChannelPublishingSrtPublicationsInner) Get() *ChannelPublishingSrtPublicationsInner {
	return v.value
}

func (v *NullableChannelPublishingSrtPublicationsInner) Set(val *ChannelPublishingSrtPublicationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelPublishingSrtPublicationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelPublishingSrtPublicationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelPublishingSrtPublicationsInner(val *ChannelPublishingSrtPublicationsInner) *NullableChannelPublishingSrtPublicationsInner {
	return &NullableChannelPublishingSrtPublicationsInner{value: val, isSet: true}
}

func (v NullableChannelPublishingSrtPublicationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelPublishingSrtPublicationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
