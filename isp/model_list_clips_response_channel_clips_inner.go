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

// ListClipsResponseChannelClipsInner struct for ListClipsResponseChannelClipsInner
type ListClipsResponseChannelClipsInner struct {
	// List of clips associated with the VOD
	Clips []ListClipsResponseChannelClipsInnerClipsInner `json:"clips"`
	// ID of the VOD from which the clips were created
	VodId string `json:"vod_id"`
}

// NewListClipsResponseChannelClipsInner instantiates a new ListClipsResponseChannelClipsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListClipsResponseChannelClipsInner(clips []ListClipsResponseChannelClipsInnerClipsInner, vodId string) *ListClipsResponseChannelClipsInner {
	this := ListClipsResponseChannelClipsInner{}
	this.Clips = clips
	this.VodId = vodId
	return &this
}

// NewListClipsResponseChannelClipsInnerWithDefaults instantiates a new ListClipsResponseChannelClipsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListClipsResponseChannelClipsInnerWithDefaults() *ListClipsResponseChannelClipsInner {
	this := ListClipsResponseChannelClipsInner{}
	return &this
}

// GetClips returns the Clips field value
func (o *ListClipsResponseChannelClipsInner) GetClips() []ListClipsResponseChannelClipsInnerClipsInner {
	if o == nil {
		var ret []ListClipsResponseChannelClipsInnerClipsInner
		return ret
	}

	return o.Clips
}

// GetClipsOk returns a tuple with the Clips field value
// and a boolean to check if the value has been set.
func (o *ListClipsResponseChannelClipsInner) GetClipsOk() ([]ListClipsResponseChannelClipsInnerClipsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Clips, true
}

// SetClips sets field value
func (o *ListClipsResponseChannelClipsInner) SetClips(v []ListClipsResponseChannelClipsInnerClipsInner) {
	o.Clips = v
}

// GetVodId returns the VodId field value
func (o *ListClipsResponseChannelClipsInner) GetVodId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VodId
}

// GetVodIdOk returns a tuple with the VodId field value
// and a boolean to check if the value has been set.
func (o *ListClipsResponseChannelClipsInner) GetVodIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VodId, true
}

// SetVodId sets field value
func (o *ListClipsResponseChannelClipsInner) SetVodId(v string) {
	o.VodId = v
}

func (o ListClipsResponseChannelClipsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["clips"] = o.Clips
	}
	if true {
		toSerialize["vod_id"] = o.VodId
	}
	return json.Marshal(toSerialize)
}

type NullableListClipsResponseChannelClipsInner struct {
	value *ListClipsResponseChannelClipsInner
	isSet bool
}

func (v NullableListClipsResponseChannelClipsInner) Get() *ListClipsResponseChannelClipsInner {
	return v.value
}

func (v *NullableListClipsResponseChannelClipsInner) Set(val *ListClipsResponseChannelClipsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListClipsResponseChannelClipsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListClipsResponseChannelClipsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListClipsResponseChannelClipsInner(val *ListClipsResponseChannelClipsInner) *NullableListClipsResponseChannelClipsInner {
	return &NullableListClipsResponseChannelClipsInner{value: val, isSet: true}
}

func (v NullableListClipsResponseChannelClipsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListClipsResponseChannelClipsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


