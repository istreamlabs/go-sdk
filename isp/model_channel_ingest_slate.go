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

// checks if the ChannelIngestSlate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelIngestSlate{}

// ChannelIngestSlate Slate configures default slating behavior.
type ChannelIngestSlate struct {
	// Source loss URL defines the location of the TS file to play when no source is available. It must have one audio and one video stream. Either MPEG2 or H.264 can be used.
	SourceLossUrl *string `json:"source_loss_url,omitempty" format:"uri-reference" doc:"Source loss URL defines the location of the TS file to play when no source is available. It must have one audio and one video stream. Either MPEG2 or H.264 can be used."`
}

// NewChannelIngestSlate instantiates a new ChannelIngestSlate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelIngestSlate() *ChannelIngestSlate {
	this := ChannelIngestSlate{}
	return &this
}

// NewChannelIngestSlateWithDefaults instantiates a new ChannelIngestSlate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelIngestSlateWithDefaults() *ChannelIngestSlate {
	this := ChannelIngestSlate{}
	return &this
}

// GetSourceLossUrl returns the SourceLossUrl field value if set, zero value otherwise.
func (o *ChannelIngestSlate) GetSourceLossUrl() string {
	if o == nil || IsNil(o.SourceLossUrl) {
		var ret string
		return ret
	}
	return *o.SourceLossUrl
}

// GetSourceLossUrlOk returns a tuple with the SourceLossUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelIngestSlate) GetSourceLossUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SourceLossUrl) {
		return nil, false
	}
	return o.SourceLossUrl, true
}

// HasSourceLossUrl returns a boolean if a field has been set.
func (o *ChannelIngestSlate) HasSourceLossUrl() bool {
	if o != nil && !IsNil(o.SourceLossUrl) {
		return true
	}

	return false
}

// SetSourceLossUrl gets a reference to the given string and assigns it to the SourceLossUrl field.
func (o *ChannelIngestSlate) SetSourceLossUrl(v string) {
	o.SourceLossUrl = &v
}

func (o ChannelIngestSlate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelIngestSlate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SourceLossUrl) {
		toSerialize["source_loss_url"] = o.SourceLossUrl
	}
	return toSerialize, nil
}

type NullableChannelIngestSlate struct {
	value *ChannelIngestSlate
	isSet bool
}

func (v NullableChannelIngestSlate) Get() *ChannelIngestSlate {
	return v.value
}

func (v *NullableChannelIngestSlate) Set(val *ChannelIngestSlate) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelIngestSlate) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelIngestSlate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelIngestSlate(val *ChannelIngestSlate) *NullableChannelIngestSlate {
	return &NullableChannelIngestSlate{value: val, isSet: true}
}

func (v NullableChannelIngestSlate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelIngestSlate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


