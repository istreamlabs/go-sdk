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

// checks if the GetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig{}

// GetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig struct for GetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig
type GetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig struct {
	Scte35 *CollapseVODRequestFilterconfigScte35 `json:"scte_35,omitempty"`
	// Time-based filtering
	Time []CollapseVODRequestFilterconfigTimeInner `json:"time,omitempty" doc:"Time-based filtering"`
}

// NewGetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig instantiates a new GetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig() *GetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig {
	this := GetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig{}
	return &this
}

// NewGetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfigWithDefaults instantiates a new GetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfigWithDefaults() *GetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig {
	this := GetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig{}
	return &this
}

// GetScte35 returns the Scte35 field value if set, zero value otherwise.
func (o *GetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig) GetScte35() CollapseVODRequestFilterconfigScte35 {
	if o == nil || IsNil(o.Scte35) {
		var ret CollapseVODRequestFilterconfigScte35
		return ret
	}
	return *o.Scte35
}

// GetScte35Ok returns a tuple with the Scte35 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig) GetScte35Ok() (*CollapseVODRequestFilterconfigScte35, bool) {
	if o == nil || IsNil(o.Scte35) {
		return nil, false
	}
	return o.Scte35, true
}

// HasScte35 returns a boolean if a field has been set.
func (o *GetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig) HasScte35() bool {
	if o != nil && !IsNil(o.Scte35) {
		return true
	}

	return false
}

// SetScte35 gets a reference to the given CollapseVODRequestFilterconfigScte35 and assigns it to the Scte35 field.
func (o *GetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig) SetScte35(v CollapseVODRequestFilterconfigScte35) {
	o.Scte35 = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig) GetTime() []CollapseVODRequestFilterconfigTimeInner {
	if o == nil || IsNil(o.Time) {
		var ret []CollapseVODRequestFilterconfigTimeInner
		return ret
	}
	return o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig) GetTimeOk() ([]CollapseVODRequestFilterconfigTimeInner, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given []CollapseVODRequestFilterconfigTimeInner and assigns it to the Time field.
func (o *GetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig) SetTime(v []CollapseVODRequestFilterconfigTimeInner) {
	o.Time = v
}

func (o GetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Scte35) {
		toSerialize["scte_35"] = o.Scte35
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	return toSerialize, nil
}

type NullableGetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig struct {
	value *GetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig
	isSet bool
}

func (v NullableGetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig) Get() *GetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig {
	return v.value
}

func (v *NullableGetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig) Set(val *GetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig(val *GetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig) *NullableGetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig {
	return &NullableGetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig{value: val, isSet: true}
}

func (v NullableGetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


