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

// checks if the ChannelPublishingSrtPublicationsInnerAudioEncodersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelPublishingSrtPublicationsInnerAudioEncodersInner{}

// ChannelPublishingSrtPublicationsInnerAudioEncodersInner struct for ChannelPublishingSrtPublicationsInnerAudioEncodersInner
type ChannelPublishingSrtPublicationsInnerAudioEncodersInner struct {
	Id *string `json:"id,omitempty" minLength:"1"`
	// Output MPEG-TS PID. PIDs should be set on the PMT, SCTE-35 and all encoders or none. Valid PIDs must 13-bit values greater than 31. If no PIDs are provided (pid == 0) then they will be generated automatically.
	Pid *int32 `json:"pid,omitempty" format:"int32" exclusiveMaximum:"8191" doc:"Output MPEG-TS PID. PIDs should be set on the PMT, SCTE-35 and all encoders or none. Valid PIDs must 13-bit values greater than 31. If no PIDs are provided (pid == 0) then they will be generated automatically."`
}

// NewChannelPublishingSrtPublicationsInnerAudioEncodersInner instantiates a new ChannelPublishingSrtPublicationsInnerAudioEncodersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelPublishingSrtPublicationsInnerAudioEncodersInner() *ChannelPublishingSrtPublicationsInnerAudioEncodersInner {
	this := ChannelPublishingSrtPublicationsInnerAudioEncodersInner{}
	return &this
}

// NewChannelPublishingSrtPublicationsInnerAudioEncodersInnerWithDefaults instantiates a new ChannelPublishingSrtPublicationsInnerAudioEncodersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelPublishingSrtPublicationsInnerAudioEncodersInnerWithDefaults() *ChannelPublishingSrtPublicationsInnerAudioEncodersInner {
	this := ChannelPublishingSrtPublicationsInnerAudioEncodersInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ChannelPublishingSrtPublicationsInnerAudioEncodersInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingSrtPublicationsInnerAudioEncodersInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ChannelPublishingSrtPublicationsInnerAudioEncodersInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ChannelPublishingSrtPublicationsInnerAudioEncodersInner) SetId(v string) {
	o.Id = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *ChannelPublishingSrtPublicationsInnerAudioEncodersInner) GetPid() int32 {
	if o == nil || IsNil(o.Pid) {
		var ret int32
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingSrtPublicationsInnerAudioEncodersInner) GetPidOk() (*int32, bool) {
	if o == nil || IsNil(o.Pid) {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *ChannelPublishingSrtPublicationsInnerAudioEncodersInner) HasPid() bool {
	if o != nil && !IsNil(o.Pid) {
		return true
	}

	return false
}

// SetPid gets a reference to the given int32 and assigns it to the Pid field.
func (o *ChannelPublishingSrtPublicationsInnerAudioEncodersInner) SetPid(v int32) {
	o.Pid = &v
}

func (o ChannelPublishingSrtPublicationsInnerAudioEncodersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelPublishingSrtPublicationsInnerAudioEncodersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Pid) {
		toSerialize["pid"] = o.Pid
	}
	return toSerialize, nil
}

type NullableChannelPublishingSrtPublicationsInnerAudioEncodersInner struct {
	value *ChannelPublishingSrtPublicationsInnerAudioEncodersInner
	isSet bool
}

func (v NullableChannelPublishingSrtPublicationsInnerAudioEncodersInner) Get() *ChannelPublishingSrtPublicationsInnerAudioEncodersInner {
	return v.value
}

func (v *NullableChannelPublishingSrtPublicationsInnerAudioEncodersInner) Set(val *ChannelPublishingSrtPublicationsInnerAudioEncodersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelPublishingSrtPublicationsInnerAudioEncodersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelPublishingSrtPublicationsInnerAudioEncodersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelPublishingSrtPublicationsInnerAudioEncodersInner(val *ChannelPublishingSrtPublicationsInnerAudioEncodersInner) *NullableChannelPublishingSrtPublicationsInnerAudioEncodersInner {
	return &NullableChannelPublishingSrtPublicationsInnerAudioEncodersInner{value: val, isSet: true}
}

func (v NullableChannelPublishingSrtPublicationsInnerAudioEncodersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelPublishingSrtPublicationsInnerAudioEncodersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


