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

// checks if the ChannelSignalingSegmentSettingsInnerTierFilterExplicitTier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelSignalingSegmentSettingsInnerTierFilterExplicitTier{}

// ChannelSignalingSegmentSettingsInnerTierFilterExplicitTier Only one of ['explicit_tier'] may be set.
type ChannelSignalingSegmentSettingsInnerTierFilterExplicitTier struct {
	Values []int32 `json:"values,omitempty" format:"int32"`
}

// NewChannelSignalingSegmentSettingsInnerTierFilterExplicitTier instantiates a new ChannelSignalingSegmentSettingsInnerTierFilterExplicitTier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelSignalingSegmentSettingsInnerTierFilterExplicitTier() *ChannelSignalingSegmentSettingsInnerTierFilterExplicitTier {
	this := ChannelSignalingSegmentSettingsInnerTierFilterExplicitTier{}
	return &this
}

// NewChannelSignalingSegmentSettingsInnerTierFilterExplicitTierWithDefaults instantiates a new ChannelSignalingSegmentSettingsInnerTierFilterExplicitTier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelSignalingSegmentSettingsInnerTierFilterExplicitTierWithDefaults() *ChannelSignalingSegmentSettingsInnerTierFilterExplicitTier {
	this := ChannelSignalingSegmentSettingsInnerTierFilterExplicitTier{}
	return &this
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *ChannelSignalingSegmentSettingsInnerTierFilterExplicitTier) GetValues() []int32 {
	if o == nil || IsNil(o.Values) {
		var ret []int32
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelSignalingSegmentSettingsInnerTierFilterExplicitTier) GetValuesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *ChannelSignalingSegmentSettingsInnerTierFilterExplicitTier) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []int32 and assigns it to the Values field.
func (o *ChannelSignalingSegmentSettingsInnerTierFilterExplicitTier) SetValues(v []int32) {
	o.Values = v
}

func (o ChannelSignalingSegmentSettingsInnerTierFilterExplicitTier) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelSignalingSegmentSettingsInnerTierFilterExplicitTier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableChannelSignalingSegmentSettingsInnerTierFilterExplicitTier struct {
	value *ChannelSignalingSegmentSettingsInnerTierFilterExplicitTier
	isSet bool
}

func (v NullableChannelSignalingSegmentSettingsInnerTierFilterExplicitTier) Get() *ChannelSignalingSegmentSettingsInnerTierFilterExplicitTier {
	return v.value
}

func (v *NullableChannelSignalingSegmentSettingsInnerTierFilterExplicitTier) Set(val *ChannelSignalingSegmentSettingsInnerTierFilterExplicitTier) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelSignalingSegmentSettingsInnerTierFilterExplicitTier) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelSignalingSegmentSettingsInnerTierFilterExplicitTier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelSignalingSegmentSettingsInnerTierFilterExplicitTier(val *ChannelSignalingSegmentSettingsInnerTierFilterExplicitTier) *NullableChannelSignalingSegmentSettingsInnerTierFilterExplicitTier {
	return &NullableChannelSignalingSegmentSettingsInnerTierFilterExplicitTier{value: val, isSet: true}
}

func (v NullableChannelSignalingSegmentSettingsInnerTierFilterExplicitTier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelSignalingSegmentSettingsInnerTierFilterExplicitTier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


