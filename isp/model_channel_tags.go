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

// ChannelTags Use ChannelMetadata when possible instead of tags.
type ChannelTags struct {
	// Indicates whether this channel is monitored by automation.
	Monitored *bool `json:"monitored,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChannelTags ChannelTags

// NewChannelTags instantiates a new ChannelTags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelTags() *ChannelTags {
	this := ChannelTags{}
	return &this
}

// NewChannelTagsWithDefaults instantiates a new ChannelTags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelTagsWithDefaults() *ChannelTags {
	this := ChannelTags{}
	return &this
}

// GetMonitored returns the Monitored field value if set, zero value otherwise.
func (o *ChannelTags) GetMonitored() bool {
	if o == nil || o.Monitored == nil {
		var ret bool
		return ret
	}
	return *o.Monitored
}

// GetMonitoredOk returns a tuple with the Monitored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTags) GetMonitoredOk() (*bool, bool) {
	if o == nil || o.Monitored == nil {
		return nil, false
	}
	return o.Monitored, true
}

// HasMonitored returns a boolean if a field has been set.
func (o *ChannelTags) HasMonitored() bool {
	if o != nil && o.Monitored != nil {
		return true
	}

	return false
}

// SetMonitored gets a reference to the given bool and assigns it to the Monitored field.
func (o *ChannelTags) SetMonitored(v bool) {
	o.Monitored = &v
}

func (o ChannelTags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Monitored != nil {
		toSerialize["monitored"] = o.Monitored
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ChannelTags) UnmarshalJSON(bytes []byte) (err error) {
	varChannelTags := _ChannelTags{}

	if err = json.Unmarshal(bytes, &varChannelTags); err == nil {
		*o = ChannelTags(varChannelTags)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "monitored")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChannelTags struct {
	value *ChannelTags
	isSet bool
}

func (v NullableChannelTags) Get() *ChannelTags {
	return v.value
}

func (v *NullableChannelTags) Set(val *ChannelTags) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelTags) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelTags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelTags(val *ChannelTags) *NullableChannelTags {
	return &NullableChannelTags{value: val, isSet: true}
}

func (v NullableChannelTags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelTags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


